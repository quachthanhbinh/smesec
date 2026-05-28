# Access Governance — Feature Specification

**Version:** 1.0  
**Date:** 2026-05-28  
**Status:** Draft — Pending Approval  
**Author:** SMESec Product Team  
**Sources:** [02-feature-research-synthesis.md](../../access-governance/02-feature-research-synthesis.md) · [03-decision-record.md](../../access-governance/03-decision-record.md)

---

## Problem Statement

SMBs với 50–500 nhân viên không có công cụ đảm bảo rằng **không có access nào bị bỏ sót** — khi nhân viên nghỉ, khi app không được phép, hay khi compliance audit yêu cầu bằng chứng. Cụ thể:

- **69%** tổ chức ghi nhận sự cố liên quan đến ex-employee còn active access sau khi nghỉ (Ponemon 2023); thời gian revoke đầy đủ trung bình là **3–5 ngày làm việc**
- **4.5 unauthorized apps/nhân viên** trung bình tại SMB; **63%** dữ liệu nhạy cảm nằm ở unsanctioned apps (Netskope 2024)
- **73%** SMB fail SOC 2 lần đầu vì thiếu access control evidence; 40–80 giờ/audit để compile bằng chứng thủ công (Vanta 2023)
- **17+ SaaS apps/nhân viên** trung bình; chỉ **36%** license được dùng thực sự; "copy access từ người trước" khi onboard = kế thừa 2 năm access creep

IT admin tại SMB không thể trả lời 3 câu hỏi cơ bản: *(1) Ai có access vào đâu? (2) App nào đang được dùng mà IT không biết? (3) Tôi có thể chứng minh điều này cho auditor không?*

**Target user**: IT admin hoặc IT manager tại công ty 50–500 nhân viên. Developer dành 20% thời gian cho IT. Người không có background security, không có thời gian configure webhook, không muốn học policy language.

---

## Solution Overview

Access Governance là control layer của SMESec — đảm bảo mọi access được kiểm soát, mọi hành động được ghi lại, và không có gì bị bỏ sót.

**Product anchor**: "Nothing falls through the cracks" — không phải "full automation". Automated khi có thể (Google + M365); human-confirmed với deep-links khi automation không khả thi (Slack, AWS, GitHub).

**Product story**: Một nhân viên vừa nộp đơn nghỉ. IT admin click "Mark as leaver." Trong 4 phút 30 giây: Google Workspace suspended, M365 sessions revoked, tất cả OAuth grants deleted. Dashboard hiển thị: Slack (link mở thẳng Slack admin — 1 click), AWS (checklist với direct console link), GitHub (1 click remove org member). PDF report tự động tạo ra: *"Offboarding complete — 2 automated, 2 manual (confirmed). Audit-ready."*

**3 luồng chính:**

```
1. OFFBOARDING:
   IT admin mark as leaver
   → Parallel: Google suspend + M365 disable + OAuth revoke (automated, <5 min)
   → Checklist: Slack + AWS + GitHub (manual, deep-links, tracked)
   → PDF report: ✅ Automated | ⚠️ Manual (confirmed) | ❌ Failed (retry)

2. SHADOW IT MANAGEMENT:
   OAuth app detected via polling/webhook
   → IT admin alerted (Slack + email <15 min)
   → Review: Approved / Blocked / Pending
   → Blocked: API revocation via Google Admin SDK / MS Graph
   → Audit trail: decision + timestamp + reviewer

3. COMPLIANCE EVIDENCE:
   6 deterministic findings run continuously
   → Dashboard: finding + severity + "Fix it" button + SOC 2 control mapping
   → One-click export: SOC 2 CC6.1–CC6.8 + ISO 27001 A.7/A.9 evidence package
```

**Scope v1:**
- Asset inventory: Google Workspace + M365 (users, groups, OAuth apps, license status)
- Asset inventory: Slack + AWS (discovery + read-only display — không phải automation)
- Shadow IT: OAuth discovery + allow-list management (approve / block / pending)
- Automated offboarding: Google + M365 (<5 min, human-initiated, saga pattern)
- Offboarding checklist: Slack + AWS + GitHub (manual, deep-links, audit trail)
- Compliance findings: 6 deterministic rules, zero ML, zero false positives
- Compliance evidence export: PDF + CSV, SOC 2 + ISO 27001

---

## Architecture / Data Flow

### Discovery Architecture: Pull-Based Polling + Hybrid Webhooks

```
┌─────────────────────────────────────────────────────────────────────┐
│                  REAL-TIME LAYER (<1 minute)                        │
│  Google Pub/Sub Push Subscription ──────────────────────────┐      │
│  (registered automatically during OAuth consent)            │      │
│  Microsoft Graph Change Notifications ──────────────────────┤      │
│  (registered automatically during OAuth consent)            ├──>  │
│                                                             SQS    │
│                  FALLBACK POLLING (15 min)                   │    │
│  Google Admin SDK (users, tokens, audit logs) ───────────────┤    │
│  Microsoft Graph (users, oauth2PermissionGrants, delta) ─────┤    │
│  GitHub org API (members, oauth apps) ───────────────────────┤    │
│  Slack (Business+/Grid: admin.users.list) ───────────────────┤    │
│  AWS CloudTrail (IAM events, read-only) ─────────────────────┤    │
└─────────────────────────────────────────────────────────────────────┘
                               │
                     ECS Worker Pool
                               │
                     ┌─────────┴──────────┐
                     │  Asset Processor   │
                     │  - Upsert assets   │
                     │  - Diff detection  │
                     │  - Run 6 findings  │
                     │  - Trigger alerts  │
                     └─────────┬──────────┘
                               │
                     PostgreSQL RDS (RLS per workspace_id)
                               │
                     ┌─────────┴──────────┐
                     │   Alert Engine     │
                     │   SES + Slack      │
                     └────────────────────┘
```

**Key webhook considerations:**
- Google Push Notifications expire every **7 days** — auto-renewal CloudWatch alarm required (<24h warning)
- Microsoft Graph subscriptions expire every **3 days to 3 years** depending on resource type — auto-renewal mandatory
- Silent expiry = alerting down without any warning; treated as production outage

### Offboarding Architecture: Saga Pattern (Non-Atomic, Per-Provider Isolation)

```
IT Admin → POST /offboarding/:user_id/execute (confirm: true)
                │
     AWS Step Functions — Offboarding State Machine
                │
        ┌───────┼──────────────┐
        │       │              │
   Google    M365         GitHub (if connected)
   Suspend   Disable      DELETE /orgs/{org}/members
   + revoke  + revoke     (1 API call, free tier)
   tokens    sessions
        │       │              │
        └───────┼──────────────┘
                │
        Checklist Items Created:
        - Slack: deep-link to admin console
        - AWS IAM: deep-link to IAM console
        (IT admin manually confirms each)
                │
        Per-step status recorded in offboarding_records
                │
        PDF report generated (Puppeteer):
        ✅ Automated | ⚠️ Manual (link) | ❌ Failed (retry)
```

**Saga semantics:**
- Each provider step commits independently — failure in M365 does NOT undo Google suspension
- Per-step failure: alert admin, mark step as `failed`, single-click retry available
- Never roll back successful suspensions
- Blast radius preview shown **before** confirmation: "This will revoke access for Sarah Chen across 3 automated providers and create 2 manual checklist items."

### Sync State Machine

```
INITIAL → SYNCING_FAST_PATH → FAST_PATH_COMPLETE → SYNCING_FULL → ACTIVE
                                                                      ↑
ACTIVE → (15-min polling schedule) ────────────────────────────────> │
                                                                      │
ACTIVE → (webhook event received) → PROCESSING_EVENT → ────────────> │
                                                                      │
ACTIVE → (delta token expired / 410 Gone) → FULL_RESYNC → ─────────> │
```

---

## Database Changes

### Core Schema

```sql
-- ── Tenant isolation (applies to ALL tables) ──────────────────────────────
-- Every table has workspace_id.
-- PostgreSQL RLS enforced at DB level (not application layer).
-- CI test: two tenants, cross-tenant query must return 0 rows. Required on every PR.

-- ── Core identity entity ──────────────────────────────────────────────────
CREATE TABLE identities (
  id               UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  workspace_id     UUID NOT NULL REFERENCES workspaces(id),
  source           TEXT NOT NULL CHECK (source IN (
                     'google_workspace', 'm365', 'slack', 'github', 'aws', 'manual'
                   )),
  external_id      TEXT NOT NULL,           -- provider's native user ID
  email            TEXT NOT NULL,
  display_name     TEXT NOT NULL,
  status           TEXT NOT NULL DEFAULT 'active'
                   CHECK (status IN ('active','suspended','deprovisioned','leaver_pending','service_account')),
  is_admin         BOOLEAN NOT NULL DEFAULT FALSE,
  mfa_enabled      BOOLEAN,
  last_login_at    TIMESTAMPTZ,
  suspended_at     TIMESTAMPTZ,
  metadata         JSONB NOT NULL DEFAULT '{}',  -- provider-specific fields
  created_at       TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at       TIMESTAMPTZ NOT NULL DEFAULT now(),
  UNIQUE (workspace_id, source, external_id)
);

CREATE INDEX idx_identities_workspace       ON identities(workspace_id);
CREATE INDEX idx_identities_status          ON identities(workspace_id, status);
CREATE INDEX idx_identities_admin           ON identities(workspace_id) WHERE is_admin = TRUE;
CREATE INDEX idx_identities_mfa_off         ON identities(workspace_id) WHERE mfa_enabled = FALSE;
CREATE INDEX idx_identities_last_login      ON identities(workspace_id, last_login_at);
CREATE INDEX idx_identities_email_trgm      ON identities USING GIN (email gin_trgm_ops);

ALTER TABLE identities ENABLE ROW LEVEL SECURITY;
CREATE POLICY tenant_isolation ON identities
  USING (workspace_id = current_setting('app.workspace_id')::UUID);

-- ── OAuth / SaaS app catalog (per workspace) ──────────────────────────────
CREATE TABLE oauth_apps (
  id                 UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  workspace_id       UUID NOT NULL REFERENCES workspaces(id),
  source             TEXT NOT NULL CHECK (source IN ('google_workspace', 'm365', 'slack', 'github')),
  client_id          TEXT NOT NULL,           -- OAuth clientId from provider
  display_name       TEXT NOT NULL,
  vendor_category    TEXT NOT NULL DEFAULT 'uncategorized'
                     CHECK (vendor_category IN (
                       'productivity','communication','dev_tool','ai_tool',
                       'finance','hr','security','analytics','storage','other','uncategorized'
                     )),
  allow_list_status  TEXT NOT NULL DEFAULT 'pending'
                     CHECK (allow_list_status IN ('approved','blocked','pending')),
  reviewed_by        UUID REFERENCES identities(id),
  reviewed_at        TIMESTAMPTZ,
  review_note        TEXT,
  first_seen_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
  user_count         INTEGER NOT NULL DEFAULT 0,  -- denormalized, updated on sync
  is_ai_tool         BOOLEAN NOT NULL DEFAULT FALSE,
  metadata           JSONB NOT NULL DEFAULT '{}',
  created_at         TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at         TIMESTAMPTZ NOT NULL DEFAULT now(),
  UNIQUE (workspace_id, source, client_id)
);

CREATE INDEX idx_oauth_apps_workspace        ON oauth_apps(workspace_id);
CREATE INDEX idx_oauth_apps_status           ON oauth_apps(workspace_id, allow_list_status);
CREATE INDEX idx_oauth_apps_pending          ON oauth_apps(workspace_id) WHERE allow_list_status = 'pending';
CREATE INDEX idx_oauth_apps_ai               ON oauth_apps(workspace_id) WHERE is_ai_tool = TRUE;

ALTER TABLE oauth_apps ENABLE ROW LEVEL SECURITY;
CREATE POLICY tenant_isolation ON oauth_apps
  USING (workspace_id = current_setting('app.workspace_id')::UUID);

-- ── OAuth grants (user ↔ app relationship) ────────────────────────────────
CREATE TABLE oauth_grants (
  id               UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  workspace_id     UUID NOT NULL REFERENCES workspaces(id),
  identity_id      UUID NOT NULL REFERENCES identities(id),
  oauth_app_id     UUID NOT NULL REFERENCES oauth_apps(id),
  scopes           TEXT[] NOT NULL DEFAULT '{}',
  granted_at       TIMESTAMPTZ,
  revoked_at       TIMESTAMPTZ,              -- NULL = still active
  revoked_by       TEXT CHECK (revoked_by IN ('admin','system','user','offboarding')),
  last_used_at     TIMESTAMPTZ,
  metadata         JSONB NOT NULL DEFAULT '{}',
  UNIQUE (workspace_id, identity_id, oauth_app_id)
);

CREATE INDEX idx_grants_identity   ON oauth_grants(workspace_id, identity_id) WHERE revoked_at IS NULL;
CREATE INDEX idx_grants_app        ON oauth_grants(workspace_id, oauth_app_id) WHERE revoked_at IS NULL;
CREATE INDEX idx_grants_active     ON oauth_grants(workspace_id) WHERE revoked_at IS NULL;

ALTER TABLE oauth_grants ENABLE ROW LEVEL SECURITY;
CREATE POLICY tenant_isolation ON oauth_grants
  USING (workspace_id = current_setting('app.workspace_id')::UUID);

-- ── Offboarding records (saga orchestration state) ────────────────────────
CREATE TABLE offboarding_records (
  id               UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  workspace_id     UUID NOT NULL REFERENCES workspaces(id),
  identity_id      UUID NOT NULL REFERENCES identities(id),
  initiated_by     UUID NOT NULL REFERENCES identities(id),
  status           TEXT NOT NULL DEFAULT 'pending_confirmation'
                   CHECK (status IN ('pending_confirmation','in_progress','complete','partial','failed')),
  step_functions_execution_arn TEXT,
  blast_radius     JSONB NOT NULL DEFAULT '{}',  -- snapshot shown before confirmation
  started_at       TIMESTAMPTZ,
  completed_at     TIMESTAMPTZ,
  notes            TEXT,
  created_at       TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_offboarding_workspace  ON offboarding_records(workspace_id, status);
CREATE INDEX idx_offboarding_identity   ON offboarding_records(workspace_id, identity_id);

ALTER TABLE offboarding_records ENABLE ROW LEVEL SECURITY;
CREATE POLICY tenant_isolation ON offboarding_records
  USING (workspace_id = current_setting('app.workspace_id')::UUID);

-- ── Per-provider offboarding step results ────────────────────────────────
CREATE TABLE offboarding_steps (
  id                  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  workspace_id        UUID NOT NULL,
  offboarding_id      UUID NOT NULL REFERENCES offboarding_records(id),
  provider            TEXT NOT NULL CHECK (provider IN (
                        'google_workspace','m365','github','slack','aws','oauth_grants','manual'
                      )),
  step_type           TEXT NOT NULL CHECK (step_type IN ('automated','checklist')),
  status              TEXT NOT NULL DEFAULT 'pending'
                      CHECK (status IN ('pending','in_progress','success','failed','skipped','manual_confirmed')),
  action_taken        TEXT,                  -- e.g. 'suspend_user', 'revoke_sessions'
  error_message       TEXT,
  deep_link_url       TEXT,                  -- for checklist items
  manual_confirmed_by UUID REFERENCES identities(id),
  manual_confirmed_at TIMESTAMPTZ,
  retry_count         INTEGER NOT NULL DEFAULT 0,
  executed_at         TIMESTAMPTZ,
  created_at          TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_steps_offboarding  ON offboarding_steps(workspace_id, offboarding_id);

ALTER TABLE offboarding_steps ENABLE ROW LEVEL SECURITY;
CREATE POLICY tenant_isolation ON offboarding_steps
  USING (workspace_id = current_setting('app.workspace_id')::UUID);

-- ── Deterministic compliance findings ────────────────────────────────────
CREATE TABLE access_findings (
  id                 UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  workspace_id       UUID NOT NULL,
  identity_id        UUID REFERENCES identities(id),
  oauth_app_id       UUID REFERENCES oauth_apps(id),
  finding_type       TEXT NOT NULL CHECK (finding_type IN (
                       'user_without_mfa',
                       'inactive_admin',
                       'blocked_app_active_grant',
                       'suspended_user_active_grants',
                       'admin_without_onboarding_record',
                       'offboarding_incomplete'
                     )),
  severity           TEXT NOT NULL CHECK (severity IN ('critical','high','medium','low')),
  title              TEXT NOT NULL,
  description        TEXT,
  remediation_url    TEXT,                    -- deep link to fix location
  compliance_refs    TEXT[] NOT NULL DEFAULT '{}',  -- ['SOC2_CC6.1', 'ISO27001_A.9.4']
  detected_at        TIMESTAMPTZ NOT NULL DEFAULT now(),
  resolved_at        TIMESTAMPTZ,
  suppressed_by      UUID REFERENCES identities(id),
  suppressed_until   TIMESTAMPTZ,
  suppression_reason TEXT
);

CREATE INDEX idx_findings_workspace  ON access_findings(workspace_id, resolved_at) WHERE resolved_at IS NULL;
CREATE INDEX idx_findings_identity   ON access_findings(workspace_id, identity_id);
CREATE INDEX idx_findings_type       ON access_findings(workspace_id, finding_type) WHERE resolved_at IS NULL;

ALTER TABLE access_findings ENABLE ROW LEVEL SECURITY;
CREATE POLICY tenant_isolation ON access_findings
  USING (workspace_id = current_setting('app.workspace_id')::UUID);

-- ── Append-only audit trail (SOC 2 evidence) ─────────────────────────────
CREATE TABLE audit_events (
  id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  workspace_id  UUID NOT NULL,
  actor_id      UUID REFERENCES identities(id),   -- NULL if system action
  actor_type    TEXT NOT NULL CHECK (actor_type IN ('admin_user','system','api_key')),
  event_type    TEXT NOT NULL,                     -- 'offboarding.executed', 'app.blocked', etc.
  resource_type TEXT,                              -- 'identity', 'oauth_app', 'oauth_grant', etc.
  resource_id   UUID,
  payload       JSONB NOT NULL DEFAULT '{}',       -- event-specific details
  source_ip     INET,
  user_agent    TEXT,
  occurred_at   TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_audit_workspace     ON audit_events(workspace_id, occurred_at DESC);
CREATE INDEX idx_audit_actor         ON audit_events(workspace_id, actor_id, occurred_at DESC);
CREATE INDEX idx_audit_resource      ON audit_events(workspace_id, resource_type, resource_id);
CREATE INDEX idx_audit_event_type    ON audit_events(workspace_id, event_type, occurred_at DESC);

-- Append-only: no UPDATE, no DELETE via application code
-- Enforced via PostgreSQL policy:
ALTER TABLE audit_events ENABLE ROW LEVEL SECURITY;
CREATE POLICY tenant_isolation ON audit_events
  USING (workspace_id = current_setting('app.workspace_id')::UUID);
CREATE POLICY append_only ON audit_events
  AS RESTRICTIVE FOR UPDATE USING (FALSE);
CREATE POLICY no_delete ON audit_events
  AS RESTRICTIVE FOR DELETE USING (FALSE);

-- ── Webhook subscription tracking ────────────────────────────────────────
CREATE TABLE webhook_subscriptions (
  id               UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  workspace_id     UUID NOT NULL,
  source           TEXT NOT NULL CHECK (source IN ('google_workspace', 'm365')),
  resource_type    TEXT NOT NULL,              -- 'users', 'tokens', 'groups', etc.
  subscription_id  TEXT NOT NULL,              -- provider's subscription ID
  expires_at       TIMESTAMPTZ NOT NULL,
  last_renewed_at  TIMESTAMPTZ,
  status           TEXT NOT NULL DEFAULT 'active'
                   CHECK (status IN ('active','expired','renewal_failed')),
  UNIQUE (workspace_id, source, resource_type)
);

CREATE INDEX idx_webhooks_expiry ON webhook_subscriptions(expires_at) WHERE status = 'active';

-- ── Sync job state tracking ───────────────────────────────────────────────
CREATE TABLE sync_jobs (
  id               UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  workspace_id     UUID NOT NULL,
  source           TEXT NOT NULL,
  trigger          TEXT NOT NULL CHECK (trigger IN ('scheduled','webhook','manual','initial')),
  status           TEXT NOT NULL CHECK (status IN ('running','success','partial_failure','failed')),
  started_at       TIMESTAMPTZ NOT NULL DEFAULT now(),
  completed_at     TIMESTAMPTZ,
  identities_added    INTEGER DEFAULT 0,
  identities_updated  INTEGER DEFAULT 0,
  apps_added          INTEGER DEFAULT 0,
  apps_updated        INTEGER DEFAULT 0,
  grants_added        INTEGER DEFAULT 0,
  grants_revoked      INTEGER DEFAULT 0,
  next_delta_token TEXT,                       -- M365 $deltaLink or Google nextPageToken
  error_details    JSONB
);

CREATE INDEX idx_sync_jobs_workspace ON sync_jobs(workspace_id, source, started_at DESC);
```

### Migrations
All schema changes deployed via numbered migration files (`migrations/NNNN_description.sql`). No manual DDL in production. Every migration reviewed by Tech Lead for RLS completeness and audit trail coverage.

---

## API Contract

### Base URL
```
/api/v1/
```
All endpoints require: `Authorization: Bearer <JWT>` with `workspace_id` claim. JWT verified by Keycloak. `app.workspace_id` set in PostgreSQL session on each request.

### Identity / Asset Endpoints

```
GET /identities
  Query params:
    status: active | suspended | deprovisioned | leaver_pending | service_account
    source: google_workspace | m365 | slack | github | aws
    is_admin: boolean
    mfa_enabled: boolean
    search: string (email + display_name full-text)
    inactive_since: ISO8601 (filter last_login_at before date)
    page: integer (default 1)
    per_page: integer (default 50, max 200)
  Response 200:
    {
      identities: Identity[],
      total: integer,
      page: integer,
      per_page: integer
    }

GET /identities/:id
  Response 200: Identity (with oauth_grants[], active_findings[])
  Response 404: { error: "identity_not_found" }

GET /identities/:id/grants
  Query params: active_only (default true)
  Response 200: { grants: OAuthGrant[] }
```

### Integration / Sync Endpoints

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

POST /integrations/m365/connect
  Body: { oauth_code: string, redirect_uri: string, tenant_id: string }
  Response 201: (same as above)

GET /sse/sync/:sync_job_id   (Server-Sent Events)
  Events:
    { event: "identities_ready",  data: { count: 147 } }
    { event: "apps_ready",        data: { count: 23, pending_review: 4 } }
    { event: "findings_ready",    data: { critical: 1, high: 3 } }
    { event: "sync_complete",     data: { total_identities: 147, total_apps: 23 } }
    { event: "sync_error",        data: { source: "m365", message: "..." } }

GET /integrations
  Response 200:
    {
      integrations: [
        {
          id: UUID,
          source: "google_workspace",
          status: "active" | "error" | "syncing" | "disconnected",
          last_synced_at: ISO8601,
          identity_count: integer,
          app_count: integer,
          error_message: string | null
        }
      ]
    }

POST /integrations/:id/sync
  Triggers manual full re-sync
  Response 202: { sync_job_id: UUID }
  Error 409: { error: "sync_in_progress" }

DELETE /integrations/:id
  Disconnects source; marks identities + apps as UNLINKED (not deleted)
  Requires explicit confirmation: Body: { confirm: true }
  Response 204
```

### Offboarding Endpoints

```
POST /offboarding/preview
  Body: { identity_id: UUID }
  Response 200:
    {
      identity: Identity,
      blast_radius: {
        automated: [
          { provider: "google_workspace", actions: ["suspend_user", "revoke_oauth_tokens"] },
          { provider: "m365",             actions: ["disable_account", "revoke_sessions"] }
        ],
        checklist: [
          { provider: "slack",  label: "Deactivate in Slack", deep_link: "https://..." },
          { provider: "aws",    label: "Remove IAM access",   deep_link: "https://..." },
          { provider: "github", label: "Remove org member",   deep_link: "https://..." }
        ],
        oauth_grants_to_revoke: integer,
        service_account_warnings: string[]  -- empty = safe to proceed
      }
    }

POST /offboarding/execute
  Body: { identity_id: UUID, confirm: true, notes: string? }
  Response 202:
    {
      offboarding_id: UUID,
      sse_channel: "/sse/offboarding/:offboarding_id"
    }
  Error 422: { error: "service_account_warning", warnings: string[] }
    (returned if unconfirmed service account detected; must pass confirmed_service_accounts: true)

GET /sse/offboarding/:offboarding_id   (Server-Sent Events)
  Events:
    { event: "step_started",    data: { provider: "google_workspace", action: "suspend_user" } }
    { event: "step_success",    data: { provider: "google_workspace", action: "suspend_user" } }
    { event: "step_failed",     data: { provider: "m365", action: "revoke_sessions", error: "..." } }
    { event: "checklist_ready", data: { items: ChecklistItem[] } }
    { event: "complete",        data: { status: "complete"|"partial", pdf_url: string } }

GET /offboarding
  Query params: status, page, per_page
  Response 200: { offboardings: OffboardingRecord[] }

GET /offboarding/:id
  Response 200: OffboardingRecord (with steps[])

POST /offboarding/:id/steps/:step_id/retry
  Retries a failed automated step
  Response 202: { step_id: UUID }

POST /offboarding/:id/steps/:step_id/confirm-manual
  Marks a checklist item as manually completed by IT admin
  Body: { confirmation_note: string? }
  Response 200: { step: OffboardingStep }

GET /offboarding/:id/report
  Returns PDF report of offboarding record
  Response 200: application/pdf
```

### Shadow IT / Allow-List Endpoints

```
GET /oauth-apps
  Query params:
    allow_list_status: approved | blocked | pending
    is_ai_tool: boolean
    source: google_workspace | m365 | slack
    search: string
    page, per_page
  Response 200: { apps: OAuthApp[], total: integer }

GET /oauth-apps/:id
  Response 200: OAuthApp (with grants[] and user details)

PUT /oauth-apps/:id/allow-list
  Body: { status: "approved"|"blocked"|"pending", review_note: string }
  Response 200: { app: OAuthApp }
  Side effect if blocked: queues OAuth revocation job for all active grants

POST /oauth-apps/:id/revoke-all-grants
  Revokes all active OAuth grants for this app across all users
  Body: { confirm: true }
  Response 202: { revocation_job_id: UUID }
  Error 422: { error: "confirmation_required" } (if confirm !== true)
```

### Findings Endpoints

```
GET /findings
  Query params: severity, finding_type, identity_id, resolved
  Response 200: { findings: AccessFinding[] }

GET /findings/summary
  Response 200:
    {
      total_open: integer,
      by_severity: { critical: 2, high: 5, medium: 3, low: 1 },
      by_type: { user_without_mfa: 3, inactive_admin: 2, ... }
    }

POST /findings/:id/resolve
  Marks finding as resolved (system will re-detect on next sync if condition still present)
  Body: { resolution_note: string }
  Response 200: { finding: AccessFinding }

POST /findings/:id/suppress
  Body: { reason: string, until: ISO8601 | null }
  Response 200: { finding: AccessFinding }
```

### Audit Trail Endpoints

```
GET /audit
  Query params:
    event_type: string
    actor_id: UUID
    resource_type: string
    resource_id: UUID
    from: ISO8601
    to: ISO8601
    page, per_page
  Response 200: { events: AuditEvent[], total: integer }

GET /audit/export
  Query params: from, to, format (csv|json)
  Response 200: Binary (CSV) or application/json
```

### Compliance Export Endpoints

```
GET /compliance/export
  Query params: framework (soc2|iso27001|gdpr|all), format (pdf|csv|json)
  Response 200: Binary (PDF) or application/json
  Report includes:
    - Asset inventory snapshot (users + apps) at time of export
    - Open findings with SOC 2 / ISO 27001 control mapping
    - Offboarding records (last 90 days)
    - Allow-list decision log
    - Access grant history
  Generation time: <2 minutes (async; 202 + poll or SSE)
```

### Error Response Format

```json
{
  "error": "error_code_snake_case",
  "message": "Human-readable description",
  "details": {}
}
```

**Error codes relevant to access governance:**
- `identity_not_found` — 404
- `integration_not_found` — 404
- `offboarding_already_in_progress` — 409
- `sync_in_progress` — 409
- `confirmation_required` — 422 (destructive action without `confirm: true`)
- `service_account_warning` — 422 (service account detected; pass `confirmed_service_accounts: true`)
- `provider_unavailable` — 503 (graceful degradation; other providers still work)
- `rate_limit_upstream` — 503 (upstream provider rate limited; retry-after header included)

---

## Platform-Specific Considerations

### Web

**Dashboard — Identity Inventory**
- Columns: Name · Email · Source · Admin · MFA · Status · Last Login · Open Findings · Actions
- Filter panel (sticky): Source · Admin status · MFA status · Status · Inactive since
- Search: full-text across email + display_name
- Actions per row: View grants · Mark as leaver → offboarding flow · Suppress findings

**Offboarding Flow (Web)**
1. Click "Mark as leaver" on identity
2. Modal: Blast radius preview — animated list of what will happen per provider
3. Service account warning (if applicable): "This user owns 1 service account. Confirm to proceed."
4. Confirmation button (active only after scrolling full blast radius)
5. Real-time progress via SSE: per-provider status chips update live
6. Completion screen: PDF download + checklist items remaining

**Shadow IT Alert Flow**
- Pending apps → persistent banner: "4 apps need review" (dismissible only by reviewing them)
- Email + Slack notification within 15 minutes of detection
- Notification: app name · user count · top scopes · one-click "Review" deep link
- Review page: app name · vendor category · users who authorized · OAuth scopes · approve/block

**Compliance Findings Dashboard**
- 6 finding cards with severity badges + open counts
- Each card: finding description · affected users/apps · "Fix it" button → deep link to admin console
- SOC 2 / ISO 27001 control mapped visibly per finding
- "Learning mode" banner: "Findings calculated — alerts suppressed for first 30 days until you review your inventory"
- One-click "Export evidence for SOC 2 Type 1" button

**Integration Health Bar** (persistent in sidebar)
```
Google Workspace  ✅ synced 8 min ago
Microsoft 365    ⚠️ throttled, retry in 3 min
GitHub           ✅ synced 12 min ago
Slack            ❌ token expired — reconnect
```

**Browser compatibility:** Chrome 120+, Edge 120+, Firefox 121+, Safari 17+

### Mobile (Flutter)

**Scope**: Incident response + JIT approval + passive monitoring

**Features:**
- **Push notification**: "New app detected: Notion — 3 users authorized it" → one-tap "Review"
- **Offboarding quick-action**: "Off-board [name]" — shows blast radius, requires biometric confirmation
- **JIT access approve/deny**: notification + one-tap decision (Sprint 7 if capacity)
- **Finding alerts**: Critical/High findings → push notification → deep link to identity detail
- **Read-only dashboard**: summary counts (active users · pending apps · open findings)
- **Identity search**: typeahead search across all providers (read-only)

**Offline behavior**: Last-synced snapshot cached in SQLite. Read-only offline. Sync on reconnect. Push notifications require connectivity.

**Security**: Biometric unlock (FaceID/TouchID) gates the app. Keycloak token stored in secure enclave. Session expires after 30 min background inactivity.

**Platforms**: iOS 16+ · Android 12+

### Desktop

Not in scope for Access Governance v1. Desktop app handles other features (see platform roadmap).

---

## Testing Strategy

### Unit Tests

**Offboarding state machine** (highest priority):
- Saga step execution: each provider success path
- Partial failure: Google succeeds, M365 fails → Google NOT rolled back; M365 flagged as failed
- Retry logic: failed step → exponential backoff → retry → success
- Service account detection: user owns service account → `service_account_warning` error returned
- Blast radius calculation: correct enumeration of automated vs checklist steps

**Compliance findings engine** (all 6 rules):
- `user_without_mfa`: `mfa_enabled = false` → Critical finding created
- `inactive_admin`: admin with `last_login_at > 90 days ago` → High finding
- `blocked_app_active_grant`: app blocked AND grant.revoked_at IS NULL → Critical finding
- `suspended_user_active_grants`: identity suspended AND grants active → High finding
- `admin_without_onboarding_record`: admin AND no onboarding record → Medium finding
- `offboarding_incomplete`: offboarding started >24h ago AND status ≠ complete → High finding
- Negative cases: all conditions not met → no finding created
- Resolution: condition corrected → finding marked resolved on next sync

**Sync state machine:**
- Delta token persistence: `next_delta_token` saved on success
- M365 delta token expiry (`410 Gone`) → fallback to full sync + log
- Webhook renewal: subscription expiry within 24h → renewal triggered
- Partial failure: 200 of 500 users fail → job status = `partial_failure`; successful rows committed

### Integration Tests

**Multi-tenant isolation** (CI gate — must pass before every PR merges):
```python
def test_cross_tenant_isolation():
    tenant_a = create_test_workspace()
    tenant_b = create_test_workspace()

    alice = create_identity(tenant_a, email="alice@acme.com")

    # Query from tenant_b context — must return 0 results
    set_db_context(workspace_id=tenant_b.id)
    results = db.query("SELECT * FROM identities WHERE email = 'alice@acme.com'")
    assert len(results) == 0

def test_audit_trail_append_only():
    # Attempt UPDATE on audit_events must be rejected by PostgreSQL policy
    with pytest.raises(psycopg2.errors.InsufficientPrivilege):
        db.execute("UPDATE audit_events SET event_type = 'tampered' WHERE TRUE")
```

**Offboarding integration** (mock provider SDKs):
- Google Admin SDK mock: verify `users.update(suspended:true)` called + `tokens.delete` per grant
- Microsoft Graph mock: verify `PATCH accountEnabled:false` + `POST revokeSignInSessions`
- GitHub mock: verify `DELETE /orgs/{org}/members/{username}` called exactly once
- Partial failure: M365 mock returns 503 → step marked failed, Google step unaffected
- PDF report generated: all ✅/⚠️/❌ statuses rendered correctly

**Shadow IT detection and revocation** (mock Google Admin SDK):
- New OAuth app in polling response → `oauth_apps` row created with status = `pending`
- Alert: mock SES + Slack webhook called within 15 minutes (simulated time)
- Block action → revocation job queued → mock SDK receives `tokens.delete` per affected user
- Block without `confirm: true` → 422 error

**Webhook renewal** (simulated expiry):
- Subscription `expires_at = now() + 20h` → renewal CloudWatch alarm triggers
- Auto-renewal API call made; `expires_at` updated
- Renewal failure → status = `renewal_failed`; admin alerted

### Security Tests

- **OWASP Top 10 scan** via OWASP ZAP on every staging deploy
- **Pen test** engaged Sprint 4 (before pilot customers), results Sprint 8 — Critical/High findings block release
- **Dependency audit**: `npm audit` + `pip-audit` in CI; no known Critical CVEs in production
- **Audit trail tamper test**: PostgreSQL policy blocks UPDATE + DELETE on `audit_events` (automated test)
- **SOC 2 Type 1**: observation period starts Week 8; auditor engaged by Week 6

### Manual / Acceptance Tests

**Offboarding acceptance test** (Sprint 6):
- Real Google Workspace test tenant; test user has 3 OAuth grants
- Trigger offboarding from web UI
- Criteria: Google suspension confirmed + all OAuth grants revoked + PDF generated within 5 minutes
- Criteria: Checklist items visible with working deep links

**Shadow IT detection test** (Sprint 4):
- Test user authorizes new OAuth app in test tenant
- Criteria: Slack alert received within 15 minutes of authorization

**Onboarding acceptance test** (Sprint 8, pilot):
- 3 pilot customers connect Google Workspace via OAuth
- Criteria: First inventory (users + apps) visible within 30 minutes
- Recorded (screen share)

---

## Acceptance Criteria

### Asset Inventory & Discovery
- [ ] IT admin can connect Google Workspace via OAuth in <5 minutes (no documentation needed)
- [ ] IT admin can connect M365 via OAuth in <5 minutes (no documentation needed)
- [ ] Initial sync discovers >90% of users and active OAuth apps (verified against known tenant data)
- [ ] First inventory visible within 30 minutes of OAuth consent (recorded in 3 pilot sessions)
- [ ] Slack + AWS assets visible in inventory (read-only, no automation) when those integrations connected

### Shadow IT Management
- [ ] New OAuth app detected → Slack + email notification delivered within 15 minutes
- [ ] IT admin can approve / block / mark pending any app from dashboard in 1 click
- [ ] Blocked app → OAuth grants revoked for all users (with explicit confirmation required)
- [ ] All allow-list decisions recorded in audit trail with reviewer identity, timestamp, and note
- [ ] AI tools tagged as category `ai_tool` automatically based on vendor catalog

### Offboarding
- [ ] Full Google Workspace offboarding (suspend + revoke all OAuth tokens) completes in <5 minutes from confirmation
- [ ] Full M365 offboarding (disable account + revoke sign-in sessions) completes in <5 minutes from confirmation
- [ ] GitHub org member removal automated (no Enterprise subscription required)
- [ ] Checklist items generated for Slack + AWS with working deep links to admin consoles
- [ ] Blast radius preview shown before every offboarding confirmation (no destructive action without preview)
- [ ] Service account detection: warning shown if offboarded user owns service accounts
- [ ] Per-provider PDF report generated: ✅ Automated / ⚠️ Manual confirmed / ❌ Failed status per provider
- [ ] Failed steps can be retried individually without re-running entire offboarding

### Compliance Findings
- [ ] All 6 deterministic findings run on every sync cycle (no ML, no false positives >2%)
- [ ] Each finding displays: severity · SOC 2 control mapping · "Fix it" deep link
- [ ] Findings dashboard shows open count per finding type + per severity
- [ ] Suppression: IT admin can suppress a finding with a reason and optional expiry date
- [ ] Compliance export: SOC 2 CC6.1–CC6.8 evidence package generated in <2 minutes as PDF
- [ ] ISO 27001 A.7/A.9 evidence export available in same export flow

### Security & Data Integrity
- [ ] Multi-tenant isolation: cross-tenant CI test passes (0 rows returned across tenant boundaries)
- [ ] Audit trail append-only: PostgreSQL policy blocks UPDATE + DELETE (automated test)
- [ ] Data encrypted at rest (AES-256, AWS KMS) and in transit (TLS 1.3)
- [ ] Zero Critical/High findings from pen test before first customer go-live
- [ ] Webhook renewal automated: <24h expiry triggers renewal; failure alerts admin

### Performance
- [ ] Identity search query response <200ms (p95) for workspace with 500 users
- [ ] Dashboard initial load <2s (p95) on standard broadband
- [ ] Sync does not hit Google Admin SDK rate limit (1,500 req/100s) for workspace ≤500 users
- [ ] M365 delta token sync handles `410 Gone` gracefully (fallback to full sync, no data loss)
- [ ] Uptime >99.5% (30-day rolling average)

### Mobile
- [ ] Push notification delivered within 2 minutes of new shadow app detected
- [ ] Offboarding initiated from mobile app completes successfully on server
- [ ] App locks after 30 minutes of background inactivity (biometric re-auth required)

---

## Out of Scope

### Explicitly Deferred to v2
- **AWS IAM automated deprovisioning** — Non-atomic; multiple resource types; 4–6 weeks alone; v2 first item after offboarding stable
- **Slack deprovisioning automation** — Requires Enterprise Grid ($12.50+/user/mo); ~60% SMBs don't qualify; add after pilot qualification check
- **GitHub SCIM deprovisioning** — Requires GitHub Enterprise Cloud ($21/user/mo)
- **HRIS integration** (BambooHR, Rippling, Gusto) — 60% of SMBs don't have HRIS; manual trigger in v1; 5+ connectors = 4–6 month detour
- **Periodic access review campaigns** (SOC 2 Type II) — After offboarding stable and Type 1 complete
- **Custom workflow builder** — 80% of customers use defaults; platform play
- **Salesforce / Jira deprovisioning** — v2 after Google + M365 depth-first

### Never In This Feature
- **Auto-provisioning (SCIM inbound)** — SMBs disable it; creates over-provisioning risk
- **RBAC role suggestions / role mining** — Never acted on; creates noise; damages trust
- **Automated shadow IT revocation without human confirmation** — Blast radius risk = instant churn
- **ML-based risk scoring** — No training data; false positive rate destroys trust in first 2 weeks
- **Behavioral analytics / UEBA** — 60+ day cold start; separate ML pipeline; v3 minimum
- **PAM (Privileged Access Management)** — Separate product category; enterprise-only

### Out of Mobile Scope (v1)
- Full admin UI on mobile (web is the primary admin surface)
- Creating or editing allow-list policies on mobile
- Bulk operations on mobile (only single-identity actions)
- Offboarding checklist item confirmation from mobile (web only)

---

*Specification version 1.0 — 2026-05-28*  
*Requires approval before implementation per AGENTS.md hard gate*  
*Next step: PLAN — sprint-by-sprint task breakdown for implementation*
