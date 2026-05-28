# Asset Inventory — Feature Specification

**Version:** 1.0  
**Date:** 2026-05-28  
**Status:** Draft — Pending Approval  
**Author:** SMESec Product Team  
**Sources:** [01-research-synthesis.md](../../access-inventory/01-research-synthesis.md) · [02-decision-record.md](../../access-inventory/02-decision-record.md)

---

## Problem Statement

SMBs với 50–500 nhân viên đang vận hành trong bóng tối kỹ thuật số: công ty 100 người thực sự dùng 150-300 ứng dụng SaaS nhưng IT chỉ biết tới 40-50. Delta này là nguồn gốc của:

- **69%** sự cố bảo mật liên quan đến ex-employee còn access sau khi nghỉ (Ponemon 2023)
- **68%** breach liên quan đến third-party app chưa được phát hiện (Verizon DBIR 2025)
- **$127K/năm** lãng phí license trung bình tại công ty 200 người (Zluri 2025)
- **52%** nhân viên dùng AI tools không được IT cấp phép; 11% paste dữ liệu nhạy cảm vào ChatGPT (Cyberhaven 2025)

IT admin không có công cụ để trả lời câu hỏi cơ bản nhất: *"Ai đang có access vào cái gì trong công ty chúng ta?"* — không phải vì dữ liệu không tồn tại, mà vì dữ liệu đó nằm rải rác ở Google Admin Console, Azure AD, AWS IAM, và hàng chục vendor portals khác.

**Target user**: IT admin hoặc IT manager tại công ty 50-500 nhân viên. Không phải CISO. Không phải security engineer. Người phụ trách mọi thứ và không có 20 giờ/tuần để dành cho security.

---

## Solution Overview

Asset Inventory là foundation layer của SMESec — một real-time, unified catalog tự động discover, classify, và duy trì tất cả digital assets của một công ty từ một nơi duy nhất.

**Product story**: Khi IT admin lần đầu connect Google Workspace, trong vòng 7 phút họ thấy: "147 users · 23 connected apps · 4 apps need review · 2 AI tools with Drive access." Đây là aha moment. Từ đây trở đi, mọi action — offboarding, shadow IT blocking, compliance export — đều được thực hiện trong context của catalog này.

**3 luồng chính:**

```
1. CONTINUOUS DISCOVERY:
   OAuth consent → Agentless API sync (15 min) + event streams (<5 min critical)
   → Classify via rule-based engine (500+ vendor catalog)
   → Surface in unified dashboard with risk scores

2. SHADOW IT MANAGEMENT:
   New OAuth app detected → Alert (Slack/email <15 min)
   → IT admin: Approve / Block / Pending
   → Blocked apps: revocation via Google Admin SDK / MS Graph
   → Audit trail: decision + timestamp + reviewer

3. COMPLIANCE EVIDENCE:
   6 deterministic findings run continuously
   → Dashboard + "Fix it" buttons
   → One-click export: ISO 27001 A.8/A.9 + SOC 2 CC6.1
```

**Scope v1:**
- Identity assets: Google Workspace + M365 (users, groups, service accounts, admin status)
- SaaS assets: OAuth-connected apps + shadow IT detection
- Shadow AI: AI tools với OAuth access (ChatGPT, Claude, Copilot, Gemini, etc.)
- Cloud security posture: 5 AWS checks (IAM, S3, root MFA, key rotation, security groups)
- Automated offboarding: Google + M365 + flagged OAuth apps (<5 min)
- Compliance evidence: exportable audit trail + PDF reports

---

## Architecture / Data Flow

### Discovery Architecture: Agentless, Dual-Speed Pipeline

```
┌─────────────────────────────────────────────────────────────────┐
│                    FAST PATH (<5 minutes)                       │
│  AWS EventBridge (CloudTrail) ──────────────────────────┐      │
│  M365 Graph Change Notifications ───────────────────────┤      │
│  Slack Events API ──────────────────────────────────────┼──>  │
│                                                          SQS   │
│                                                         FIFO   │
│                    STANDARD PATH (15–60 min)              │    │
│  Google Workspace Admin SDK (15-min poll) ───────────────┤    │
│  M365 Graph delta queries (15-min poll) ─────────────────┤    │
│  GitHub API (30-min poll) ───────────────────────────────┤    │
│  AWS IAM/S3 reconciliation (60-min) ─────────────────────┤    │
└─────────────────────────────────────────────────────────────────┘
                              │
                    ECS Worker Pool
                              │
                    ┌─────────┴──────────┐
                    │  Asset Processor   │
                    │  - Classify        │
                    │  - Upsert assets   │
                    │  - Detect delta    │
                    │  - Trigger alerts  │
                    └─────────┬──────────┘
                              │
                    PostgreSQL RDS (assets table, RLS)
                              │
                    ┌─────────┴──────────┐
                    │   Alert Engine     │
                    │   SES + Slack      │
                    └────────────────────┘
```

### First-Sync Fast Path (<30 Minutes to First Value)

```
T+0:00  IT admin clicks "Connect Google Workspace" → OAuth consent screen
T+0:30  OAuth callback processed; ECS sync task spawned
        Dashboard shows: "Syncing your workspace..." (loading skeleton)

T+3:00  PHASE 1 — Users visible
        Admin SDK: admin.users.list (1-4 API calls for ≤500 users)
        PostgreSQL COPY bulk insert (~10K rows/sec)
        SSE push: {event: "users_ready", count: 147}
        Dashboard: user table renders immediately

T+7:00  PHASE 2 — OAuth apps visible (AHA MOMENT)
        Reports API: activities.list (applicationName=token, last 30 days)
        Vendor catalog match (trigram index, batch lookup)
        SSE push: {event: "apps_ready", count: 23, shadow_candidates: 4}
        Dashboard: "23 apps connected · 4 need review · 2 AI tools"

T+25:00 PHASE 3 (background) — Full history + findings
        tokens.list per user (full historical — 500 calls for 500 users)
        6 deterministic findings calculated
        Risk scores derived and stored
        SSE push: {event: "sync_complete", total_apps: 31}
```

**Why this architecture enables <30 min**: Reports API (1-3 calls for 30-day grant history) runs before `tokens.list` (500 individual calls). 80-90% of apps visible at T+7. Risk scores never block first render.

### Classification Pipeline (6-Stage Rule Engine)

```
Stage 1: Asset type detection (source → type, deterministic)
Stage 2: Vendor catalog match (domain/clientId → app name + category + sensitivity)
Stage 3: OAuth scope classification (drive/mail = CONFIDENTIAL; profile only = INTERNAL)
Stage 4: AWS resource rules (S3 public = RESTRICTED; RDS instance = CONFIDENTIAL)
Stage 5: Shadow detection (not in catalog → SHADOW; untagged AWS → SHADOW)
Stage 6: Orphan detection (user departed AND active OAuth grants remain)

Confidence: Stages 1-2 → 1.0 | Stages 3-4 → 0.85 | Stages 5-6 → 0.70
Below 0.60 → status = UNCLASSIFIED → surface to IT admin for manual review
```

**Acceptance criterion**: <10% Unknown for top-200 common SMB tools (hard gate before any customer goes live).

---

## Database Changes

### Core Schema

```sql
-- ── Tenant isolation (applies to ALL tables) ──────────────────────────────
-- Every table has workspace_id.
-- PostgreSQL RLS enforced at DB level (not application layer).
-- CI test: two tenants, cross-tenant query must return 0 rows. Required on every PR.

-- ── Universal asset entity ────────────────────────────────────────────────
CREATE TABLE assets (
  id               UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  workspace_id     UUID NOT NULL REFERENCES workspaces(id),
  asset_type       TEXT NOT NULL CHECK (asset_type IN (
                     'human_user', 'service_account', 'group', 'role',
                     'saas_app', 'ai_tool', 'cloud_resource', 'data_store', 'device'
                   )),
  source           TEXT NOT NULL CHECK (source IN (
                     'google_workspace', 'm365', 'aws', 'slack', 'github',
                     'browser_extension', 'expense_report', 'manual'
                   )),
  external_id      TEXT NOT NULL,           -- provider's native ID
  display_name     TEXT NOT NULL,
  classification   TEXT NOT NULL DEFAULT 'internal'
                   CHECK (classification IN ('public','internal','confidential','restricted')),
  status           TEXT NOT NULL DEFAULT 'active'
                   CHECK (status IN ('active','inactive','orphaned','shadow','decommissioned')),
  lifecycle_state  TEXT NOT NULL DEFAULT 'discovered'
                   CHECK (lifecycle_state IN ('discovered','active','inactive','orphaned','decommissioned')),
  is_admin         BOOLEAN DEFAULT FALSE,
  mfa_enabled      BOOLEAN,
  last_seen_at     TIMESTAMPTZ,
  last_login_at    TIMESTAMPTZ,             -- for users: last IdP login
  owner_id         UUID REFERENCES assets(id),
  metadata         JSONB NOT NULL DEFAULT '{}',
  first_seen_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
  created_at       TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at       TIMESTAMPTZ NOT NULL DEFAULT now(),
  UNIQUE (workspace_id, source, external_id)
);

CREATE INDEX idx_assets_workspace_type    ON assets(workspace_id, asset_type);
CREATE INDEX idx_assets_workspace_status  ON assets(workspace_id, status);
CREATE INDEX idx_assets_workspace_shadow  ON assets(workspace_id) WHERE status = 'shadow';
CREATE INDEX idx_assets_workspace_orphan  ON assets(workspace_id) WHERE status = 'orphaned';
CREATE INDEX idx_assets_display_name_gin  ON assets USING GIN (to_tsvector('english', display_name));
CREATE INDEX idx_assets_last_login        ON assets(workspace_id, last_login_at);

ALTER TABLE assets ENABLE ROW LEVEL SECURITY;
CREATE POLICY tenant_isolation ON assets
  USING (workspace_id = current_setting('app.workspace_id')::UUID);

-- ── Asset relationships (edges) ───────────────────────────────────────────
CREATE TABLE asset_relationships (
  id                 UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  workspace_id       UUID NOT NULL REFERENCES workspaces(id),
  source_asset_id    UUID NOT NULL REFERENCES assets(id),
  target_asset_id    UUID NOT NULL REFERENCES assets(id),
  relationship_type  TEXT NOT NULL CHECK (relationship_type IN (
                       'member_of', 'has_access_to', 'owns', 'granted_by', 'depends_on'
                     )),
  granted_scopes     TEXT[],               -- OAuth scopes when has_access_to
  granted_at         TIMESTAMPTZ,
  revoked_at         TIMESTAMPTZ,          -- NULL = still active
  metadata           JSONB NOT NULL DEFAULT '{}',
  created_at         TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_rel_source    ON asset_relationships(workspace_id, source_asset_id) WHERE revoked_at IS NULL;
CREATE INDEX idx_rel_target    ON asset_relationships(workspace_id, target_asset_id) WHERE revoked_at IS NULL;

ALTER TABLE asset_relationships ENABLE ROW LEVEL SECURITY;
CREATE POLICY tenant_isolation ON asset_relationships
  USING (workspace_id = current_setting('app.workspace_id')::UUID);

-- ── Append-only state changes (compliance audit trail) ───────────────────
CREATE TABLE asset_snapshots (
  id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  workspace_id UUID NOT NULL,
  asset_id     UUID NOT NULL REFERENCES assets(id),
  snapshot_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
  change_type  TEXT NOT NULL CHECK (change_type IN ('created','updated','deleted','reclassified','status_changed')),
  delta        JSONB NOT NULL,             -- {field: {old: x, new: y}}
  source_event TEXT                        -- e.g. 'google_push_notification'
);

CREATE INDEX idx_snapshots_asset ON asset_snapshots(workspace_id, asset_id, snapshot_at DESC);

ALTER TABLE asset_snapshots ENABLE ROW LEVEL SECURITY;
CREATE POLICY tenant_isolation ON asset_snapshots
  USING (workspace_id = current_setting('app.workspace_id')::UUID);

-- ── Sync job state tracking ───────────────────────────────────────────────
CREATE TABLE sync_jobs (
  id               UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  workspace_id     UUID NOT NULL,
  source           TEXT NOT NULL,
  phase            TEXT NOT NULL CHECK (phase IN ('fast_path','standard')),
  status           TEXT NOT NULL CHECK (status IN ('running','success','partial_failure','failed')),
  started_at       TIMESTAMPTZ NOT NULL DEFAULT now(),
  completed_at     TIMESTAMPTZ,
  assets_added     INTEGER DEFAULT 0,
  assets_updated   INTEGER DEFAULT 0,
  assets_removed   INTEGER DEFAULT 0,
  next_sync_token  TEXT,                   -- delta token for incremental sync
  error_details    JSONB
);

CREATE INDEX idx_sync_jobs_workspace ON sync_jobs(workspace_id, source, started_at DESC);

-- ── Vendor catalog (platform-wide, not tenant-specific) ──────────────────
CREATE TABLE vendor_catalog (
  id                UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  client_id         TEXT UNIQUE,           -- Google OAuth clientId
  domain            TEXT,                  -- e.g. 'notion.so'
  display_name      TEXT NOT NULL,
  category          TEXT NOT NULL CHECK (category IN (
                      'productivity','communication','dev_tool','ai_tool',
                      'finance','hr','security','analytics','storage','other'
                    )),
  default_sensitivity TEXT NOT NULL DEFAULT 'internal'
                    CHECK (default_sensitivity IN ('public','internal','confidential','restricted')),
  soc2_certified    BOOLEAN DEFAULT FALSE,
  iso27001_certified BOOLEAN DEFAULT FALSE,
  aliases           TEXT[],               -- normalized merchant name variants
  updated_at        TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_vendor_domain      ON vendor_catalog(domain);
CREATE INDEX idx_vendor_client_id   ON vendor_catalog(client_id);
CREATE INDEX idx_vendor_name_trgm   ON vendor_catalog USING GIN (display_name gin_trgm_ops);

-- ── Deterministic findings ────────────────────────────────────────────────
CREATE TABLE asset_findings (
  id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  workspace_id    UUID NOT NULL,
  asset_id        UUID NOT NULL REFERENCES assets(id),
  finding_type    TEXT NOT NULL,           -- 'admin_without_mfa', 'orphaned_account', etc.
  severity        TEXT NOT NULL CHECK (severity IN ('critical','high','medium','low')),
  title           TEXT NOT NULL,
  description     TEXT,
  remediation_url TEXT,                    -- deep link to fix location
  compliance_ref  TEXT[],                  -- ['SOC2_CC6.1', 'ISO27001_A.9.2']
  detected_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
  resolved_at     TIMESTAMPTZ,
  suppressed_by   UUID REFERENCES assets(id),
  suppressed_until TIMESTAMPTZ,
  suppression_reason TEXT
);

CREATE INDEX idx_findings_workspace ON asset_findings(workspace_id, resolved_at) WHERE resolved_at IS NULL;
CREATE INDEX idx_findings_asset     ON asset_findings(workspace_id, asset_id);

ALTER TABLE asset_findings ENABLE ROW LEVEL SECURITY;
CREATE POLICY tenant_isolation ON asset_findings
  USING (workspace_id = current_setting('app.workspace_id')::UUID);

-- ── Risk scores (derived, deterministic rollup) ───────────────────────────
CREATE TABLE asset_risk_scores (
  id                  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  workspace_id        UUID NOT NULL,
  asset_id            UUID NOT NULL REFERENCES assets(id),
  score               INTEGER NOT NULL CHECK (score BETWEEN 0 AND 100),
  severity            TEXT NOT NULL CHECK (severity IN ('critical','high','medium','low')),
  factors             JSONB NOT NULL,      -- {"admin_without_mfa": 25, "orphaned": 30}
  calculated_at       TIMESTAMPTZ NOT NULL DEFAULT now(),
  overridden          BOOLEAN DEFAULT FALSE,
  override_reason     TEXT,
  override_by         UUID REFERENCES assets(id),
  override_expires_at TIMESTAMPTZ,
  UNIQUE (workspace_id, asset_id)
);

ALTER TABLE asset_risk_scores ENABLE ROW LEVEL SECURITY;
CREATE POLICY tenant_isolation ON asset_risk_scores
  USING (workspace_id = current_setting('app.workspace_id')::UUID);

-- ── License / zombie account tracking ─────────────────────────────────────
CREATE TABLE app_license_records (
  id                   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  workspace_id         UUID NOT NULL,
  asset_id             UUID NOT NULL REFERENCES assets(id),
  plan_name            TEXT,
  license_count        INTEGER,
  cost_monthly_cents   INTEGER,            -- set manually by admin in v1
  billing_period       TEXT CHECK (billing_period IN ('monthly','annual')),
  renewal_date         DATE,
  source               TEXT NOT NULL CHECK (source IN ('manual','expense_report','vendor_api')),
  created_at           TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at           TIMESTAMPTZ NOT NULL DEFAULT now()
);

ALTER TABLE app_license_records ENABLE ROW LEVEL SECURITY;
CREATE POLICY tenant_isolation ON app_license_records
  USING (workspace_id = current_setting('app.workspace_id')::UUID);

-- ── Usage events (daily granularity, stale account detection) ─────────────
CREATE TABLE app_usage_events (
  id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  workspace_id    UUID NOT NULL,
  app_asset_id    UUID NOT NULL REFERENCES assets(id),
  user_asset_id   UUID NOT NULL REFERENCES assets(id),
  event_date      DATE NOT NULL,
  event_count     INTEGER DEFAULT 1,
  source          TEXT NOT NULL CHECK (source IN (
                    'google_audit','m365_audit','slack_audit',
                    'browser_extension','expense_report'
                  )),
  UNIQUE (workspace_id, app_asset_id, user_asset_id, event_date, source)
);

CREATE INDEX idx_usage_user   ON app_usage_events(workspace_id, user_asset_id, event_date DESC);
CREATE INDEX idx_usage_app    ON app_usage_events(workspace_id, app_asset_id, event_date DESC);

ALTER TABLE app_usage_events ENABLE ROW LEVEL SECURITY;
CREATE POLICY tenant_isolation ON app_usage_events
  USING (workspace_id = current_setting('app.workspace_id')::UUID);

-- ── Allow-list management (shadow IT) ────────────────────────────────────
CREATE TABLE app_allow_list (
  id             UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  workspace_id   UUID NOT NULL,
  asset_id       UUID NOT NULL REFERENCES assets(id),
  status         TEXT NOT NULL CHECK (status IN ('approved','blocked','pending')),
  reviewed_by    UUID REFERENCES assets(id),
  reviewed_at    TIMESTAMPTZ,
  review_note    TEXT,
  UNIQUE (workspace_id, asset_id)
);

ALTER TABLE app_allow_list ENABLE ROW LEVEL SECURITY;
CREATE POLICY tenant_isolation ON app_allow_list
  USING (workspace_id = current_setting('app.workspace_id')::UUID);
```

### Migrations
All schema changes deployed via numbered migration files (`migrations/NNNN_description.sql`). No manual DDL in production. Every migration reviewed by Tech Lead for RLS completeness.

---

## API Contract

### Base URL
```
/api/v1/
```
All endpoints require: `Authorization: Bearer <JWT>` with `workspace_id` claim. JWT verified by Keycloak.

### Asset Endpoints

```
GET /assets
  Query params:
    asset_type: human_user | saas_app | ai_tool | cloud_resource | ...
    status: active | inactive | orphaned | shadow
    source: google_workspace | m365 | aws | ...
    classification: public | internal | confidential | restricted
    search: string (full-text, display_name)
    risk_min: integer 0-100
    page: integer (default 1)
    per_page: integer (default 50, max 200)
  Response 200:
    {
      assets: Asset[],
      total: integer,
      page: integer,
      per_page: integer
    }

GET /assets/:id
  Response 200: Asset (with relationships[] and findings[])
  Response 404: { error: "asset_not_found" }

GET /assets/:id/relationships
  Query params: relationship_type, active_only (default true)
  Response 200: { relationships: Relationship[] }

GET /assets/:id/history
  Response 200: { snapshots: Snapshot[] }
```

### Discovery / Sync Endpoints

```
POST /integrations/google/connect
  Body: { oauth_code: string, redirect_uri: string }
  Response 201:
    {
      integration_id: UUID,
      sync_job_id: UUID,
      status: "syncing",
      sse_channel: "/sse/sync/:sync_job_id"
    }

GET /sse/sync/:sync_job_id   (Server-Sent Events)
  Events:
    { event: "users_ready",    data: { count: 147 } }
    { event: "apps_ready",     data: { count: 23, shadow_candidates: 4 } }
    { event: "sync_progress",  data: { pct: 65 } }
    { event: "sync_complete",  data: { total_assets: 178, findings: 3 } }
    { event: "sync_error",     data: { source: "google_workspace", message: "..." } }

GET /integrations
  Response 200: { integrations: Integration[] }  — connection status per source

POST /integrations/:id/sync
  Triggers manual re-sync (queues sync job)
  Response 202: { sync_job_id: UUID }

DELETE /integrations/:id
  Disconnects source; marks assets as UNLINKED (not deleted)
  Response 204
```

### Shadow IT / Allow-List Endpoints

```
GET /allow-list
  Query params: status (approved|blocked|pending), page, per_page
  Response 200: { apps: AllowListEntry[] }

PUT /allow-list/:asset_id
  Body: { status: "approved"|"blocked"|"pending", review_note: string }
  Response 200: { app: AllowListEntry }
  Side effect: If blocked → revocation job queued (Google Admin SDK / MS Graph)

POST /allow-list/:asset_id/revoke-all-grants
  Confirms bulk OAuth revocation for all users (requires explicit action)
  Body: { confirm: true }
  Response 202: { revocation_job_id: UUID }
```

### Findings & Risk Score Endpoints

```
GET /findings
  Query params: severity, asset_id, finding_type, resolved
  Response 200: { findings: Finding[] }

PUT /findings/:id/suppress
  Body: { reason: string, until: ISO8601 | null }
  Response 200: { finding: Finding }

GET /risk-scores
  Query params: asset_type, min_score, sort_by (score|updated)
  Response 200: { scores: RiskScore[] }
```

### Compliance Export Endpoints

```
GET /compliance/export
  Query params: framework (iso27001|soc2|all), format (pdf|csv|json)
  Response 200: Binary (PDF) or application/json
  Includes: asset inventory snapshot + findings summary + access review evidence
```

### Error Response Format

```json
{
  "error": "error_code_snake_case",
  "message": "Human-readable description",
  "details": {}  // optional field-level errors
}
```

**Error codes relevant to asset inventory:**
- `integration_not_found` — 404
- `sync_in_progress` — 409 (cannot trigger manual sync while one is running)
- `revocation_requires_confirm` — 422 (bulk revocation without `confirm: true`)
- `rate_limit_upstream` — 503 (upstream provider rate limited; retry via SSE)
- `provider_unavailable` — 503 (graceful degradation: other providers still work)

---

## Platform-Specific Considerations

### Web

**Dashboard — Asset Inventory Table**
- Columns: Name · Type · Source · Classification · Status · Risk Score · Last Seen · Actions
- Filter panel (sticky): Asset Type · Source · Status · Classification · Risk Score range
- Search: full-text across display_name, materialized to PostgreSQL tsvector
- Sort: any column, server-side
- Pagination: 50/page default; infinite scroll option
- Actions per row: View details · Approve/Block (shadow apps) · Revoke access

**Shadow IT Alert Flow**
- New unreviewed OAuth app → banner alert on dashboard ("3 apps need review")
- Email + Slack notification within 15 minutes of detection
- Notification includes: app name · user count · top OAuth scopes · one-click "Review" deep link

**Onboarding Wizard** (Sprint 7 deliverable — highest retention impact)
- Step 1: "Connect your identity provider" → Google / M365 OAuth button
- Step 2: "Grant admin access" → inline explanation of why each permission is needed
- Step 3: Loading screen with progressive counter: "Discovering users... Discovering apps..."
- Step 4: First inventory view with explainer overlays ("You have 3 apps pending review")
- Acceptance criterion: 3 new pilot customers complete Step 1→4 in <30 minutes (timed recording)

**Risk Score Display**
- Per user: colored badge (Critical 🔴 / High 🟠 / Medium 🟡 / Low 🟢) + factor breakdown tooltip
- Per app: same badge + "why" expansion panel
- "Learning mode" first 30 days: scores computed but not surfaced as alerts until >50% of apps reviewed

**Browser compatibility**: Chrome 120+, Edge 120+, Firefox 121+, Safari 17+

### Mobile (Flutter)

**Scope**: Incident response + passive monitoring (not full admin UI)

**Features:**
- Push notification: "New app detected: Notion — 3 users have authorized it" → one-tap "Review"
- Asset search: typeahead search across identity + SaaS assets (read-only)
- Finding alerts: push notification for Critical/High findings → deep link to asset detail
- Offboarding quick-action: "Off-board [name]" initiated from mobile (confirmation required, full execution on server)
- Dashboard widget: summary counts (users · shadow apps · open findings)

**Offline behavior**: Last-synced snapshot cached locally (SQLite). Read-only offline. Sync on reconnect.

**Platforms**: iOS 16+ · Android 12+

**Biometric auth**: FaceID/TouchID for app unlock (Keycloak mobile token, biometric-gated)

### Desktop

Not in scope for Asset Inventory v1. Desktop app handles other features (see platform roadmap).

---

## Testing Strategy

### Unit Tests

**Classification engine** (highest coverage priority):
- All 6 classification stages with known inputs/expected outputs
- Vendor catalog lookup: exact match, trigram fuzzy match, no-match
- OAuth scope sensitivity mapping: all scope patterns documented
- Risk score calculation: all factor combinations; edge cases (all zeros, max score)
- Finding detection: each of the 6 deterministic rules, positive and negative cases

**Sync job state machine**:
- `sync_jobs.next_sync_token` persistence and recovery
- Delta token expiry (M365 `410 Gone` → full sync fallback)
- Partial failure handling: some assets fail to insert, job status = `partial_failure`

### Integration Tests

**Multi-tenant isolation** (CI gate — runs on every PR):
```python
# Must pass before any PR merges
def test_cross_tenant_isolation():
    tenant_a = create_test_workspace()
    tenant_b = create_test_workspace()
    
    asset_a = create_asset(tenant_a, display_name="Secret App A")
    
    # Query from tenant_b context — must return 0 results
    set_db_context(workspace_id=tenant_b.id)
    results = db.query("SELECT * FROM assets WHERE display_name = 'Secret App A'")
    assert len(results) == 0
```

**Sync pipeline**:
- Full sync with synthetic 500-user Google Workspace tenant (mock SDK)
- Delta sync recovery after simulated 8-day token expiry gap
- Rate limit handling: mock 429 responses → exponential backoff → resume
- Fast path timing: Phase 1 (users) visible within 3 minutes of OAuth consent
- Phase 2 (apps) visible within 7 minutes of OAuth consent

**Allow-list + revocation**:
- Block app → OAuth revocation job queued → mock Google Admin SDK receives revocation call
- Revocation with `confirm: false` → 422 error (blast radius protection)
- Bulk revocation → each user's grant revoked independently (not all-or-nothing)

**Classification accuracy gate** (runs before Sprint 5):
- Benchmark: 200 common SMB apps (known display names + client IDs)
- Acceptance: <10% classified as "unknown" or "shadow" incorrectly
- Fail = block release until catalog updated

### Security Tests

- **OWASP Top 10 scan** via OWASP ZAP on every staging deploy
- **Pen test** engaged Sprint 4, results Sprint 8 — Critical/High findings block release
- **Dependency audit**: `npm audit` + `pip-audit` in CI; no known Critical CVEs in production

### Manual / Acceptance Tests

**Onboarding acceptance test** (Sprint 7):
- 3 pilot customers perform fresh onboarding
- Criteria: Time from OAuth button click → first assets visible ≤ 30 minutes
- Recorded (screen share) for documentation

**Shadow IT detection test** (Sprint 4):
- Test user authorizes new OAuth app in test tenant
- Criteria: Slack alert received within 15 minutes

**Offboarding test** (Sprint 6):
- Test user revoked across Google + M365 + all flagged OAuth apps
- Criteria: Full revocation confirmed within 5 minutes; PDF report generated

---

## Acceptance Criteria

### Onboarding & Discovery
- [ ] IT admin can connect Google Workspace via OAuth in <5 minutes (no documentation needed)
- [ ] IT admin can connect M365 via OAuth in <5 minutes (no documentation needed)
- [ ] First assets visible within 30 minutes of OAuth consent (recorded in 3 pilot sessions)
- [ ] Initial sync discovers >90% of users and OAuth apps (verified against known tenant data)
- [ ] Vendor catalog: <10% of top-200 common SMB apps classified as "Unknown" (hard gate)

### Shadow IT
- [ ] New OAuth app detected → Slack + email notification delivered within 15 minutes
- [ ] IT admin can approve / block / mark pending any app from dashboard in 1 click
- [ ] Blocked app → OAuth grants revoked for all users (with explicit confirmation required)
- [ ] All allow-list decisions recorded in audit trail (who, what, when, why)
- [ ] AI tools (ChatGPT, Claude, Copilot, Gemini, etc.) tagged as category "ai_tool" automatically

### Findings & Risk Scores
- [ ] 6 deterministic findings run on every sync cycle: admin without MFA, orphaned account, IAM inactive >90d, blocked app with active OAuth, public S3 bucket, root MFA disabled
- [ ] Risk score (0-100) displayed per user and per app with factor breakdown
- [ ] "Learning mode": scores computed but alerts suppressed until >50% of apps reviewed
- [ ] Each finding has a "Fix it" deep link to the appropriate admin console

### Security & Compliance
- [ ] Multi-tenant isolation: cross-tenant CI test passes (0 rows returned across tenant boundaries)
- [ ] Data encrypted at rest (AES-256, AWS KMS) and in transit (TLS 1.3)
- [ ] Zero Critical/High findings from pen test before first customer go-live
- [ ] Compliance export: ISO 27001 A.8/A.9 evidence generated in <2 minutes
- [ ] SOC 2 CC6.1 evidence (access inventory snapshot) exportable as PDF

### Performance
- [ ] Asset search query response <200ms (p95) for workspace with 1,000 assets
- [ ] Dashboard initial load <2s (p95) on standard broadband
- [ ] Sync does not hit Google Admin SDK rate limit (1,500 req/100s) for any workspace ≤500 users
- [ ] Uptime >99.5% (30-day rolling average)

### Mobile
- [ ] Push notification delivered within 2 minutes of new shadow app detected
- [ ] Asset search returns results within 1 second on mobile (offline: cached data)
- [ ] Offboarding action initiated from mobile completes successfully on server

---

## Out of Scope

### Explicitly Deferred to v2
- **Full cloud asset inventory** (AWS Config resource graph, Azure, GCP) — multi-cloud coverage gap creates trust erosion in v1; 5 security posture checks are the v1 scope
- **SaaS PII data classification** (Google Drive/SharePoint scanning) — 83hr/tenant initial scan, $150/tenant DLP cost, GDPR legal review; v2 minimum
- **Full license waste dashboard** (billing APIs for Zoom, Salesforce, HubSpot) — per-vendor API integrations; v1 has inactive user cost recovery only
- **Browser extension (content monitoring)** — conditional on MV3 persistence gate and GDPR legal review; Track 2 delivery
- **Expensify/Concur integration** — after Ramp/Brex validated
- **Periodic access review campaigns** — SOC 2 Type II; after asset inventory stable

### Never In This Feature
- **Agent-based endpoint discovery** — MDM required; SMB deployment friction
- **Auto-revocation without human confirmation** — blast radius risk = instant churn
- **JQL or graph query interface** — builds for security engineers, not IT admins
- **ML-based anomaly detection or risk scoring** — no training data; false positive rate kills trust
- **Inline network CASB** — requires proxy infrastructure; enterprise complexity
- **Raw credit card number ingestion** — PCI DSS risk

### Out of Mobile Scope (v1)
- Full admin UI on mobile (web is the primary admin surface)
- Creating/editing classification rules on mobile
- Bulk actions on mobile (only single-asset actions)

---

*Specification version 1.0 — 2026-05-28*  
*Requires approval before implementation per AGENTS.md hard gate*  
*Next step: PLAN — sprint-by-sprint task breakdown for implementation*
