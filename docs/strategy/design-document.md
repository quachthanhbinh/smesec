# SMESec Platform — Design Document

**Date:** 2026-05-28  
**Status:** Approved  
**Version:** 1.0  
**Scope:** Core architectural decisions — Build vs Buy · Multi-tenancy · AI-threat detection · Data privacy  
**Related:** [system-architecture.md](system-architecture.md) · [delivery-plan.md](delivery-plan.md)

---

## Table of Contents

1. [Executive Summary (600 words — Deliverable)](#1-executive-summary-600-words--deliverable)
2. [Build vs Buy — Component Details](#2-build-vs-buy--component-details)
3. [Multi-Tenancy Model](#3-multi-tenancy-model)
4. [AI-Threat Detection Strategy](#4-ai-threat-detection-strategy)
5. [Data Privacy Guarantees](#5-data-privacy-guarantees)
6. [AI Governance Module](#6-ai-governance-module)
7. [Architectural Trade-offs & Rejected Alternatives](#7-architectural-trade-offs--rejected-alternatives)

---

## 1. Executive Summary (600 words — Deliverable)

> *This section is the 600-word summary for Deliverable #2. Subsequent sections are detailed internal documentation.*

---

### Build vs Buy: Hybrid Approach

SMESec applies a deliberate **Hybrid Build/Buy** strategy: buy proven commodity services, build what creates real differentiation.

**Buy:** Authentication (Keycloak self-hosted — OIDC/SAML, zero per-user cost), compliance automation (Vanta — $4–6K/yr vs 3 months of engineering), deepfake detection API (Hive Moderation — pay-per-use $0.01/check), ML platform (AWS SageMaker — managed training and inference), and cloud infrastructure (AWS ECS Fargate, RDS, EventBridge, Step Functions).

**Build:** Integration sync engine (Google Workspace + M365 + Slack + AWS IAM — rate limit handling, delta sync, shadow IT detection are the core differentiators), asset inventory and classification engine (no competitor detects shadow AI tools at SME pricing), browser extension DLP (local PII inference, tenant-scoped allow-list, privacy-preserving), incident playbook wizard (domain-specific UX for non-security staff), and all domain logic (Clean Architecture — domain is vendor-independent).

**Why Hybrid wins:** Pure Buy = vendor lock-in to tools ill-suited for SMEs. Pure Build = 18+ months before product exists. Hybrid = v1 in 6 months with gross margin ~91% (infra costs $44K/yr vs $480K ARR at 50 customers).

---

### Multi-Tenancy Model: Shared Infrastructure, Isolated Data

SMESec uses a **Shared PostgreSQL with Row-Level Security (RLS)** model — a single database cluster serves all tenants, but every query is automatically scoped by `tenant_id` at the database level (not application level).

Every domain table has two mandatory columns: `tenant_id UUID NOT NULL` and `data_residency VARCHAR(10) NOT NULL` (values: `'US'`, `'EU'`, `'APAC'`). The PostgreSQL RLS policy blocks all cross-tenant access even when there are application code bugs. API middleware injects `tenant_id` into every database session via `SET LOCAL app.tenant_id`. Tenant isolation CI tests run on every PR — failures block the merge.

EU tenants are routed to `eu-west-1` ECS tasks and a separate RDS cluster. US/APAC tenants go to `us-east-1`. Data never leaves the committed region — this is a hard requirement of GDPR Article 46.

---

### AI-Threat Detection Strategy: 2-Track Architecture

SMESec solves the AI detection reliability problem by completely separating **Track 1** (deterministic, 100% accuracy) and **Track 2** (ML/AI, target >90% accuracy).

Track 1 is the backbone: asset inventory, access governance, offboarding automation, incident playbooks, compliance reporting — all rule-based logic, with no ML dependency. Track 2 is the AI detection layer: shadow AI risk scoring, LLM data leakage prevention, deepfake fraud defense, prompt injection detection — R&D-gated, ships only when accuracy thresholds are met.

Both tracks share the `ThreatDetectionEvent` schema contract and EventBridge event bus. Track 2 events can automatically trigger Track 1 playbooks — but Track 1 does not depend on Track 2. If Track 2 produces a false positive, Track 1 continues operating normally.

---

### Data Privacy Guarantees

Four non-negotiable commitments:

1. **No training on customer data:** ML models are trained on public datasets and synthetic data — customer data is never used for training.
2. **Local inference for browser extension:** PII detection runs locally in the browser (Presidio WASM) — content the user types never leaves their device.
3. **Immutable audit logs:** S3 Object Lock (WORM, 7-year retention) — nobody, including SMESec engineers, can delete audit evidence.
4. **Data residency isolation:** `data_residency` column on every table from Sprint 1 — EU data stays in `eu-west-1`, never replicated to US.

---

## 2. Build vs Buy — Component Details

### 2.1 Decision Matrix

| Component | Decision | Detailed Rationale | Cost |
|---|---|---|---|
| **Authentication (SSO + MFA)** | **Buy: Keycloak (self-hosted ECS)** | OIDC/SAML 2.0, Google + M365 federation built-in. Zero per-user cost (vs Auth0 $0.23/MAU = $1,380/mo at 500 users/tenant × 10 tenants). Full control: custom MFA flows, branding, GDPR DPA. **⚠️ R-C6 requirements:** (1) Min 2 ECS tasks (active-active, not just multi-AZ placement); (2) JWKS caching mandatory — JWT validation must not depend on Keycloak uptime; (3) Keycloak DB separate from app DB. **Alternative (revisit at v1 launch):** If DevSecOps ops capacity is insufficient, evaluate WorkOS/Auth0 (~$500–1,000/mo) before v1.5. | ~$50/mo (ECS only); WorkOS/Auth0 alt: ~$500–1,000/mo |
| **Compliance automation** | **Buy: Vanta (Startup plan)** | $4–6K/yr vs 3 months engineer time ($60K+ cost). AWS + GitHub + Cloudflare connectors native. Evidence collection 24/7. Auditor portal. SOC 2 Type 1 in 60 days. | $4–6K/yr |
| **Deepfake detection** | **Buy: Hive Moderation API** | Pay-per-use (<$0.01/check). No training data required. Vendor maintains model updates. Voice + Video. Only SME-accessible tool with a real-time API. | ~$0.01/check |
| **ML platform** | **Buy: AWS SageMaker** | Managed training jobs, endpoint auto-scaling, model registry, A/B testing. vs 6 months building custom MLOps infra. Drift monitoring built-in. | ~$200–500/mo |
| **Cloud infrastructure** | **Buy: AWS (ECS, RDS, EventBridge, Step Functions, S3)** | 99.9%+ SLA. Managed scaling. Compliance certifications (ISO 27001, SOC 2) inherited. Single vendor = simplified compliance scope. | ~$2–4K/mo (v1 launch) |
| **Integration sync engine** | **Build (Go)** | Google Admin SDK rate limits, M365 delta link quirks, Slack tier detection, AWS IAM pagination — all require custom handling. This is the core differentiator (shadow IT detection). | 2 engineers × 3 sprints |
| **Asset inventory + classification** | **Build (Go)** | No competitor has Shadow AI detection at SME pricing. Rule-based classification engine is the moat. | 1 engineer × 4 sprints |
| **Browser extension DLP** | **Build (Chrome MV3 + Edge)** | Local inference (privacy-preserving). Tenant-scoped allow-list. Prompt Security (closest competitor) costs $15–30K/yr. | 1 FE engineer × 3 sprints |
| **Incident playbook engine** | **Build on Step Functions** | Step Functions = proven orchestration (retry, state, parallel). Build playbook logic + wizard UI. Domain-specific UX for non-security staff is the differentiator. | 2 engineers × 2 sprints |
| **Audit logging** | **Build on S3 Object Lock** | S3 Object Lock = WORM compliance-ready at near-zero cost. No vendor to depend on for immutability. | ~$10–50/mo (storage) |
| **AI phishing detection** | **Buy + Thin wrapper: M365 Defender / Google Workspace security** | Enterprise-grade phishing detection. SMESec adds: alert routing + playbook trigger + compliance evidence. No need to build ML classifier for known phishing. | Included in M365/Google subscription |
| **Prompt injection detection** | **Buy: Lakera Guard API (v1)** | Production-validated (~$0.001/request). No training data required. Covers known injection patterns + novel variants. **Internal BERT target moved to Sprint 23–24 Enterprise-only evaluation** — not a v1 target. Gate: FPR <2% AND TPR >85% on 30-day production holdout before graduating from beta. | ~$0.001/request; ~$50/mo at 50K daily checks |
| **Observability** | **Buy: CloudWatch (primary) + Datadog (optional v1.5)** | CloudWatch = zero additional cost (AWS native). Datadog APM if budget allows post-v1. | CloudWatch: included; Datadog: ~$200/mo |

### 2.2 Build Decision Criteria

Build only when ≥2 of 4 criteria are met:

```
✅ BUILD CRITERIA:
  1. Is a core differentiator vs competitors (customers pay for this)
  2. No affordable alternative (<$5K/yr) meets requirements
  3. Requires deep customization that SaaS tools don't support
  4. High vendor risk (lock-in, price increase, sunset)

✅ BUY CRITERIA:
  1. Commodity service (authentication, monitoring, storage)
  2. Compliance/security domain already certified by vendor
  3. Time-to-market: usable immediately in Sprint 1
  4. Cost < 3 months engineering time
```

### 2.3 Total Cost of Ownership (Year 1, v1 Launch)

```
BUY costs (monthly, ~50 customers):
  Vanta:               $400/mo  ($4,800/yr)
  AWS infrastructure:  $3,000/mo
  Keycloak ECS:        $50/mo (compute only)
  Hive Moderation:     $200/mo  (usage-based est.)
  Cloudflare R2:       $50/mo
  ─────────────────────────────────────────────
  Total Buy:          ~$3,700/mo = $44,400/yr

BUILD cost (amortized, 9 FTE × 6 months):
  Engineer cost:      ~$540K (year 1 build cost)
  Amortized over 3yr: $180K/yr

Revenue at 50 customers (avg $800/mo):
  MRR:                $40,000/mo = $480,000/yr
  Gross Margin:       ($480K - $44K infra) / $480K = ~91%
  (after subtracting engineer salaries: ~35–45% net margin)
```

---

## 3. Multi-Tenancy Model

### 3.1 Approach: Shared Database, Isolated Data

**Decision:** PostgreSQL Row-Level Security (RLS) — Shared cluster, data isolation at the database level.

**Three approaches considered:**

| Approach | Description | SMESec Decision |
|---|---|---|
| **Silo (Separate DB per tenant)** | Each tenant has its own DB | ❌ Too costly (~$100/mo/tenant), not viable at SME pricing |
| **Shared Schema (App-level isolation)** | Shared DB, application code filters by tenant | ❌ Application bug → cross-tenant data leak. Not sufficiently trustworthy. |
| **Shared Schema (DB-level RLS)** | Shared DB, PostgreSQL RLS enforces isolation | ✅ **Chosen** — Defense in depth: both DB and app enforce isolation |

### 3.2 Mandatory Schema Design

```sql
-- ═══════════════════════════════════════════════════════════════
-- RULE: Every domain table MUST have these 2 columns. No exceptions.
-- Enforced by: migration validator script + code review checklist
-- ═══════════════════════════════════════════════════════════════

CREATE TABLE assets (
    id             UUID         PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id      UUID         NOT NULL REFERENCES tenants(id),
    data_residency VARCHAR(10)  NOT NULL
        CHECK (data_residency IN ('US', 'EU', 'APAC')),

    -- Domain columns
    asset_type     TEXT         NOT NULL,  -- 'user_account'|'oauth_app'|'device'|'cloud_resource'
    name           TEXT         NOT NULL,
    criticality    TEXT         NOT NULL DEFAULT 'MEDIUM'
        CHECK (criticality IN ('LOW', 'MEDIUM', 'HIGH', 'CRITICAL')),
    data_sensitivity TEXT       NOT NULL DEFAULT 'INTERNAL'
        CHECK (data_sensitivity IN ('PUBLIC', 'INTERNAL', 'CONFIDENTIAL', 'SECRET')),
    owner_user_id  UUID,
    provider       TEXT         NOT NULL,  -- 'google'|'m365'|'slack'|'aws'|'manual'
    provider_id    TEXT         NOT NULL,  -- External ID (e.g. Google user ID)
    metadata       JSONB        NOT NULL DEFAULT '{}',
    is_shadow_it   BOOLEAN      NOT NULL DEFAULT FALSE,
    is_shadow_ai   BOOLEAN      NOT NULL DEFAULT FALSE,
    last_seen_at   TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    created_at     TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

-- ─── RLS Policy ──────────────────────────────────────────────────
ALTER TABLE assets ENABLE ROW LEVEL SECURITY;
ALTER TABLE assets FORCE ROW LEVEL SECURITY;  -- Applies even to table owner

CREATE POLICY tenant_isolation_assets ON assets
    AS PERMISSIVE
    FOR ALL
    TO app_role  -- Application DB user
    USING (tenant_id = current_setting('app.tenant_id', TRUE)::UUID);

-- ─── Performance ─────────────────────────────────────────────────
CREATE INDEX idx_assets_tenant_type    ON assets(tenant_id, asset_type);
CREATE INDEX idx_assets_tenant_created ON assets(tenant_id, created_at DESC);
CREATE INDEX idx_assets_shadow         ON assets(tenant_id) WHERE is_shadow_it = TRUE OR is_shadow_ai = TRUE;
```

### 3.3 Application-Level Enforcement

```go
// ═══════════════════════════════════════════════════════════════
// infrastructure/middleware/tenant.go
// ═══════════════════════════════════════════════════════════════

// TenantMiddleware extracts tenant_id from JWT claims and sets
// PostgreSQL session variable so RLS policy activates.
// Every handler runs inside this middleware — no bypass possible.
func TenantMiddleware(db *pgxpool.Pool) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            claims, ok := c.Get("jwt_claims").(JWTClaims)
            if !ok || claims.TenantID == "" {
                return echo.NewHTTPError(http.StatusUnauthorized, "missing tenant context")
            }

            // Validate UUID format — prevent injection
            if _, err := uuid.Parse(claims.TenantID); err != nil {
                return echo.NewHTTPError(http.StatusUnauthorized, "invalid tenant_id format")
            }

            // Acquire connection from pool and set session variable
            conn, err := db.Acquire(c.Request().Context())
            if err != nil {
                return echo.ErrInternalServerError
            }
            defer conn.Release()

            _, err = conn.Exec(c.Request().Context(),
                "SELECT set_config('app.tenant_id', $1, TRUE)", // LOCAL = transaction-scoped
                claims.TenantID,
            )
            if err != nil {
                return echo.ErrInternalServerError
            }

            c.Set("db_conn", conn)
            c.Set("tenant_id", claims.TenantID)
            c.Set("data_residency", claims.DataResidency)
            return next(c)
        }
    }
}
```

### 3.4 Tenant Isolation CI Test (Must Be Green Before Every Merge)

```go
// tests/integration/tenant_isolation_test.go
// Runs in CI/CD pipeline. Blocks merge on failure.

func TestCrossTenantIsolation(t *testing.T) {
    tenantA := createTestTenant(t)
    tenantB := createTestTenant(t)

    // Tenant A creates an asset
    assetA := createAsset(t, tenantA, "test-oauth-app")

    // Tenant B tries to read Tenant A's asset — must return empty/error
    t.Run("tenant_B_cannot_read_tenant_A_assets", func(t *testing.T) {
        assets := queryAssetsAsTeant(t, tenantB)
        for _, a := range assets {
            assert.NotEqual(t, assetA.ID, a.ID,
                "CRITICAL: Cross-tenant data leak detected!")
        }
    })

    // Direct UUID query with wrong tenant must return 0 rows
    t.Run("direct_id_query_scoped_by_tenant", func(t *testing.T) {
        result := queryAssetByIDAsTeant(t, tenantB, assetA.ID)
        assert.Nil(t, result, "RLS must block cross-tenant direct ID access")
    })

    // Admin bypass test — even using superuser connection
    t.Run("app_role_cannot_bypass_rls", func(t *testing.T) {
        // app_role does not have BYPASSRLS privilege
        count := countAllAssetsWithAppRole(t)
        expectedCount := countAssetsForTenant(t, tenantB)
        assert.Equal(t, expectedCount, count, "app_role must not see other tenants' data")
    })
}
```

### 3.5 Data Residency Routing Architecture

```
                    ┌─────────────────────────────────┐
                    │      Tenant Onboarding Wizard    │
                    │  "Where is your team located?"  │
                    │  ○ United States                 │
                    │  ○ European Union (GDPR)         │
                    │  ○ Asia Pacific                  │
                    └──────────────┬──────────────────┘
                                   │ data_residency set at tenant creation
                                   │ IMMUTABLE after creation
                    ┌──────────────▼──────────────────┐
                    │         Route 53 (GeoDNS)        │
                    └──────┬──────────────┬────────────┘
                           │              │
              ┌────────────▼───┐    ┌─────▼──────────────┐
              │  us-east-1     │    │  eu-west-1          │
              │  ECS Cluster   │    │  ECS Cluster        │
              │  RDS Primary   │    │  RDS Primary        │
              │  S3 (us)       │    │  S3 (eu)            │
              │  KMS (us)      │    │  KMS (eu)           │
              └────────────────┘    └─────────────────────┘

Invariant:
  ● EU tenant data NEVER written to us-east-1 RDS
  ● EU tenant OAuth tokens stored in eu-west-1 Secrets Manager
  ● EU tenant audit logs ONLY in eu-west-1 S3 bucket
  ● Vanta evidence collection: EU data never leaves EU
```

---

## 4. AI-Threat Detection Strategy

### 4.1 The Core Problem: Accuracy vs Trust

> Incorrect AI detection has two equally bad consequences:
> - **False Negative:** Miss a real deepfake fraud → money lost, data lost
> - **False Positive:** Block legitimate work → frustrate employees → product becomes ineffective

**Solution:** 2-track architecture with separate trust gateways for each track.

### 4.2 Track 1 — Deterministic Security Controls

```
TRACK 1: DETERMINISTIC (No ML/AI)
══════════════════════════════════════════════════════════

Threat Surface Covered:
  ✅ Shadow IT discovery       — OAuth app inventory (rule-based scope analysis)
  ✅ Orphaned access           — Offboarding automation (deterministic state machine)
  ✅ Over-provisioning         — RBAC diff engine (compare actual vs policy)
  ✅ Compliance violations     — Control mapping (ISO 27001 / SOC 2 / GDPR checklist)
  ✅ Access anomalies          — Deterministic rules (e.g., ex-employee with active session)

Detection Approach:
  Input:  Provider API responses (Google, M365, Slack, AWS)
  Logic:  Rule engine in Go (no ML)
           ● OAuth scope risk scoring: scopes × sensitivity matrix
           ● Shadow IT: app NOT in approved-list → alert
           ● Stale access: last login >90 days + critical permission → alert
           ● Offboarding gap: employee deactivated in HR ≠ deactivated in SaaS → CRITICAL
  Output: ThreatDetectionEvent{Source:"track1", MLMetadata:nil}

Accuracy: ~100% (rules are deterministic)
Latency:  <5 sec from sync to alert
Trust:    Full automation allowed (automated offboarding, playbook auto-trigger)
```

### 4.3 Track 2 — AI/ML Threat Detection

#### 4.3.1 Shadow AI Governance

```
FEATURE: Shadow AI Risk Scoring
Goal: Detect and score AI tools authorized by employees

Pipeline:
  1. Track 1 discovers OAuth apps (every 15 min)
  2. Classification step: is this an AI tool?
     ● Lookup against AI tool registry (ChatGPT, Copilot, Gemini, Claude,
       Mistral, Perplexity, etc.) — maintained dataset, versioned
     ● OAuth scope analysis: does app request file/email access?
  3. Risk scoring (SageMaker endpoint: shadow-ai-scorer-v1):

     Feature vector:
       - app_category: 'ai_assistant'|'ai_coding'|'ai_image'|...
       - oauth_scopes_requested: [list of scopes]
       - scope_sensitivity_score: float (0.0–1.0)
       - num_users_authorized: int
       - app_age_days: int (new app = higher risk)
       - has_known_data_exfil_cve: bool
       - tenant_approved: bool

     Output: risk_score (0.0–1.0), risk_tier ('LOW'|'MEDIUM'|'HIGH'|'CRITICAL')

  4. Policy enforcement:
     LOW/MEDIUM:  Alert IT admin, no block
     HIGH:        Alert + require employee attestation (confirm/deny usage)
     CRITICAL:    Alert + auto-block OAuth grant (dry-run first, 2-step confirm)

Accuracy target: >95% AI tool classification
False positive policy: Conservative threshold — unknown apps default to MEDIUM, not CRITICAL
```

#### 4.3.2 LLM Data Leakage Prevention (Browser Extension)

```
FEATURE: Real-time DLP before submission to ChatGPT / Copilot / Gemini
Architecture: Client-side (local) + Server-side (async logging)

Client-side (Chrome Extension Content Script):
  1. Monitor: input/textarea in AI tool domains (chatgpt.com, copilot.microsoft.com, etc.)
  2. On submit (form submit / Enter keypress):
     → Run local PII detector (Microsoft Presidio compiled to WASM)
         Detects: Email, Phone, Credit Card, SSN, IBAN, IP Address,
                  API Keys (regex patterns), Company-specific keywords
     → Risk assessment:
         SAFE:     Submit normally (no block)
         MEDIUM:   Show warning + require click-through confirmation
         HIGH:     Block submit + show "Sensitive data detected" modal
                   + offer to redact automatically
         CRITICAL: Block submit + notify IT admin (async, pseudonymized)

Server-side (async, privacy-preserving):
  → POST /api/v1/dlp-events (only sends: category, severity, timestamp, tenant_id)
  → NEVER sends: actual content, user's text, the prompt itself
  → IT admin dashboard: "3 HIGH-risk AI submissions blocked today"

Privacy guarantee: Browser extension NEVER sends content to SMESec servers.
                   Only sends pseudonymized metadata (type of PII detected, not the PII).
```

#### 4.3.3 Deepfake Fraud Defense

```
FEATURE: Voice/video verification before executing sensitive transactions
Use case: "Is this my CEO actually asking me to wire $50K?"

Architecture: Out-of-Band Verification Workflow (AWS Step Functions)

Flow:
  1. Employee receives suspicious voice/video call claiming to be exec/partner
  2. Employee triggers "Verify this person" from SMESec mobile app
  3. SMESec initiates OOB verification:
     a. Send pre-agreed verification code via separate channel (SMS to employee's registered phone)
     b. Audio/video analysis via Hive Moderation API:
        - POST audio/video hash to Hive (NOT raw content — privacy)
        - Hive returns: deepfake_score (0.0–1.0), detection_confidence
     c. Combined decision:
        - OOB code NOT received + deepfake_score > 0.7 → LIKELY DEEPFAKE
        - OOB code received + deepfake_score < 0.3   → LIKELY LEGITIMATE
        - Ambiguous → escalate to IT admin + do not proceed

  4. Outcome logged as ThreatDetectionEvent → compliance evidence
  5. If confirmed deepfake → auto-trigger Incident Playbook #6 (Deepfake Fraud Response)

Cost: ~$0.01/check via Hive API
Spend cap: $50/mo per tenant (configurable, with CloudWatch alarm)

Limitation (transparent to customer): Real-time call interception is NOT supported.
The verification workflow requires manual trigger by the employee.
This is by design — automated interception of calls raises legal concerns in most jurisdictions.
```

#### 4.3.4 Prompt Injection Detection

```
FEATURE: Detect attempts to manipulate AI tools via injected instructions
Applicable to: SMEs deploying internal AI assistants (Enterprise tier)

v1 — Lakera Guard API (Sprint 8, replacing internal rule-based target):
  Service: Lakera Guard REST API
  Integration: API call per user-submitted prompt before forwarding to AI assistant
  Cost: ~$0.001/request
  Latency: <50ms (p99)
  Accuracy: Production-validated by Lakera (covers known + novel injection patterns)
  ⚠️ BS-4 fix: Internal BERT model target was unrealistic for v1 timeline.
     Lakera Guard is already cited in competitive analysis — use it.

v2 — Internal BERT Classifier (Sprint 23–24, Enterprise tier only):
  Trigger: Only if Lakera Guard pricing becomes prohibitive at Enterprise volume
           AND sufficient labeled production data available (>50K samples)
  Model: Fine-tuned BERT (bert-base-uncased) on production data (opt-in tenants only)
  Accuracy target: TPR >85%, FPR <2%
  Infrastructure: SageMaker endpoint (async queue for non-real-time use cases)
  Gate: Must achieve FPR <2% AND TPR >85% on 30-day production holdout set
        before graduating from beta. If not achieved → Lakera Guard remains GA.
```

### 4.4 Track 2 Accuracy Gates (Ship Criteria)

| Feature | Minimum Accuracy | False Positive Limit | Gate Evaluation |
|---|---|---|---|
| Shadow AI classification | >95% AI tool identification | <10% (unknown apps miscategorized as AI) | Sprint 9 production evaluation |
| LLM DLP (PII detection) | >99% for CRITICAL data (credit card, SSN) | <5% for INTERNAL data | Sprint 8 staging test |
| Deepfake detection | >80% voice deepfake detection (independently evaluated by SMESec ML team on labeled test dataset — not vendor-asserted); combined with OOV workflow ≈ 99% fraud prevention rate | <15% (ambiguous → escalate, not block) | Sprint 10 evaluation |
| Prompt injection (Lakera Guard v1) | TPR >85%, FPR <2% — independently evaluated by SMESec ML team on 30-day SMESec-specific production holdout | Not vendor-managed — SMESec evaluates independently | Sprint 8 integration + 30-day holdout evaluation |
| Prompt injection (BERT internal) | TPR >85% | FPR <2% | Sprint 24 production holdout (Enterprise only) |

> **Policy:** If accuracy gate is not met → feature remains `beta`, opt-in only. Track 2 features are never shipped as GA without passing the accuracy gate. **All accuracy gates are evaluated independently by the SMESec ML team on SMESec-specific production holdout data — vendor API uptime SLAs (e.g., Lakera Guard) do not satisfy accuracy gates.** Monthly re-evaluation: 500 labeled samples; alert if TPR drops below 82% (prompt injection) or 75% (deepfake).

---

## 5. Data Privacy Guarantees

### 5.1 Four Core Commitments

```
╔══════════════════════════════════════════════════════════════════╗
║  SMESEC DATA PRIVACY COMMITMENTS (contractual, verifiable)      ║
╠══════════════════════════════════════════════════════════════════╣
║  1. NO TRAINING ON CUSTOMER DATA                                 ║
║     ML models trained on public datasets + synthetic data only. ║
║     Customer data is NEVER used for model training or fine-tuning║
║     without explicit opt-in consent (Enterprise tier, future).  ║
║                                                                  ║
║  2. LOCAL INFERENCE FOR BROWSER EXTENSION                        ║
║     Content typed in AI tools NEVER leaves the user's browser.  ║
║     PII detection runs via Presidio WASM (fully local).         ║
║     Only pseudonymized event metadata sent to SMESec servers.   ║
║                                                                  ║
║  3. IMMUTABLE AUDIT LOGS (tamper-proof + GDPR-erasable)          ║
║     S3 Object Lock (WORM) — ciphertext cannot be deleted by anyone.       ║
║     7-year retention for compliance evidence.                    ║
║     **GDPR Art. 17 resolution:** Per-tenant KMS envelope         ║
║     encryption — key destruction = permanent inaccessibility    ║
║     = "effective erasure" per EDPB Recommendation 01/2020.      ║
║     Audit ciphertext persists but is permanently unreadable.   ║
║                                                                  ║
║  4. DATA RESIDENCY ISOLATION                                     ║
║     EU tenant data STAYS in eu-west-1. No exceptions.           ║
║     Enforced at: DB schema, S3 bucket policy, ECS task routing, ║
║     KMS key region, and Secrets Manager region.                 ║
╚══════════════════════════════════════════════════════════════════╝
```

### 5.2 Data Classification & Handling

| Data Class | Examples | Storage | Encryption | Retention | Access |
|---|---|---|---|---|---|
| **Customer PII** | Employee names, emails, job titles | RDS (tenant-scoped) | AES-256 (KMS CMK) | Tenant lifetime + 30 days | App role (RLS) only |
| **OAuth Tokens** | Google refresh tokens, M365 app secrets | AWS Secrets Manager | AES-256 (Secrets Manager CMK) | Auto-rotated | Integration Sync Service only (IAM role) |
| **Audit Logs** | Access revocations, playbook executions | S3 Object Lock + PostgreSQL | AES-256 | 7 years (WORM) | Read-only via pre-signed URL (tenant-scoped) |
| **ML Feature Vectors** | OAuth scope scores, app risk signals | SageMaker (ephemeral training) | AES-256 | 30 days training job retention | ML Engineer IAM role |
| **DLP Events (browser ext)** | "HIGH risk PII type detected" | PostgreSQL (metadata only) | AES-256 | 12 months | Security Admin within tenant |
| **Deepfake Verification** | Audio/video HASH only (not content) | Event log only | AES-256 | 12 months | Security Admin within tenant |
| **Compliance Evidence** | SOC 2 control screenshots, access logs | S3 + Vanta | AES-256 | 7 years | Vanta auditor portal (tenant-authorized) |
| **Telemetry / Metrics** | API latency, error rates, feature usage | CloudWatch (anonymized) | In-transit TLS | 90 days | SMESec operations team |

### 5.3 GDPR Compliance Architecture

```
GDPR Article Mapping → SMESec Controls:

Art. 5 (Principles):
  ● Lawfulness: Legitimate interest + contract basis for employee monitoring
  ● Purpose limitation: Each data class has defined purpose (see table 5.2)
  ● Data minimisation: OAuth scopes minimum-required; DLP sends metadata only
  ● Accuracy: 15-min sync ensures current state
  ● Storage limitation: Defined retention per class (see table 5.2)
  ● Integrity/confidentiality: AES-256 + TLS 1.3 + RLS

Art. 13/14 (Transparency):
  ● Privacy notice template provided to SMESec customers for their employees
  ● Browser extension shows consent dialog on first install

Art. 17 (Right to erasure):
  ● /api/v1/gdpr/erasure endpoint: anonymizes PII within 30 days
  ● Audit logs: WORM (S3 Object Lock, 7-year) — cannot be deleted through normal means
  ● **Erasure mechanism for WORM audit logs:** Per-tenant KMS CMK
    → Customer submits DELETE → Schedule KMS key deletion (7-day pending window)
    → After key deletion: DEK cannot be decrypted → ciphertext unreadable forever
    → Issue erasure certificate with key_deletion_timestamp
    → Satisfies GDPR "effective erasure" standard (EDPB Recommendation 01/2020)
    → DPA contract must explicitly reference this cryptographic erasure approach
  ● Automated erasure pipeline (Sprint 11): Google → M365 cascade

Art. 20 (Portability):
  ● /api/v1/gdpr/export: JSON export of all tenant data (Sprint 11)

Art. 25 (Privacy by design):
  ● data_residency column mandatory from Sprint 1
  ● Local inference for browser extension (no data leaves device)
  ● Pseudonymization in ML feature vectors (SHA-256 hash of user IDs)

Art. 32 (Security):
  ● See section 5.4 (Encryption architecture)
  ● Pen-test bi-annual
  ● Security Hub + GuardDuty monitoring

Art. 33/34 (Breach notification):
  ● Incident Playbook #7: Data Breach Response
  ● 72h GDPR notification SLA automated via Step Functions
  ● DPA notification template pre-filled from audit evidence
```

### 5.4 Encryption Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                    ENCRYPTION AT REST                           │
│                                                                 │
│  RDS PostgreSQL:                                                │
│    ● AES-256 via AWS KMS CMK (per-region, per-environment)     │
│    ● CMK rotation: automatic, annual                            │
│    ● Read replicas: encrypted with same CMK                     │
│                                                                 │
│  S3 Buckets (audit logs, evidence):                             │
│    ● SSE-KMS (AES-256, KMS CMK)                                 │
│    ● Object Lock: COMPLIANCE mode, 7-year retention             │
│    ● Bucket policy: deny unencrypted PutObject                  │
│                                                                 │
│  AWS Secrets Manager:                                           │
│    ● KMS CMK (separate from data CMK)                           │
│    ● Auto-rotation enabled (OAuth tokens: on-use rotation)      │
│                                                                 │
│  ElastiCache Redis:                                             │
│    ● Encryption at rest (AES-256)                               │
│    ● In-transit encryption (TLS 1.3)                            │
│    ● No sensitive data stored — session tokens only (15-min TTL)│
└─────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────┐
│                    ENCRYPTION IN TRANSIT                        │
│                                                                 │
│  External:   TLS 1.3 (minimum), TLS 1.0/1.1 disabled           │
│              HSTS: max-age=31536000; includeSubDomains          │
│              Certificate: ACM (auto-renewal), OCSP stapling     │
│                                                                 │
│  Internal:   VPC internal traffic: TLS (mTLS for service-to-   │
│              service in v1.5), private subnets (no public IPs)  │
│                                                                 │
│  Browser Extension:  Local Presidio WASM (no external calls     │
│              for PII detection). Metadata POST: TLS 1.3 only.  │
└─────────────────────────────────────────────────────────────────┘
```

### 5.5 Secrets Management

```go
// infrastructure/secrets/manager.go
// RULE: Zero plaintext secrets in environment variables, config files, or code.
// ALL secrets accessed via AWS Secrets Manager.

type SecretManager struct {
    client *secretsmanager.Client
    cache  *sync.Map // Short-lived local cache (30 sec TTL)
}

// GetOAuthToken retrieves a tenant's provider OAuth token.
// Accessed ONLY by Integration Sync Service — IAM policy enforces this.
func (sm *SecretManager) GetOAuthToken(ctx context.Context, tenantID, provider string) (*OAuthTokens, error) {
    secretID := fmt.Sprintf("smesec/%s/oauth/%s", tenantID, provider)
    // IAM role for IntegrationSyncService has GetSecretValue permission ONLY for smesec/*/oauth/*
    // All other services have ZERO access to OAuth tokens — principle of least privilege
    result, err := sm.client.GetSecretValue(ctx, &secretsmanager.GetSecretValueInput{
        SecretId: &secretID,
    })
    if err != nil {
        return nil, fmt.Errorf("failed to retrieve oauth token for tenant %s: %w", tenantID, err)
    }
    // ... unmarshal and return
}
```

**IAM Least Privilege — Secret Access Matrix:**

| Service | Secret Type | Permission |
|---|---|---|
| Integration Sync Service | `smesec/*/oauth/*` (provider tokens) | `secretsmanager:GetSecretValue` |
| API Gateway Service | `smesec/*/jwt-signing-key` | `secretsmanager:GetSecretValue` |
| All other services | — | ❌ No access |
| ML Services | `smesec/sagemaker/*` | `secretsmanager:GetSecretValue` (ML keys only) |
| Compliance Service | `smesec/*/vanta-api-key` | `secretsmanager:GetSecretValue` |

---

## 6. AI Governance Module

### 6.1 Scope & Objectives

> **The problem:** 75% of knowledge workers use AI tools at work. 78% bring their own tools without employer approval. 11% of data pasted into ChatGPT contains confidential company information. (Cyberhaven, 2024)

The SMESec AI Governance Module is not designed to **ban** AI — but to **see, understand, and intentionally control** AI usage in the organization.

**Three governance levels:**

```
LEVEL 1 — DISCOVER (Passive)
  Know which AI tools employees are using
  → No intervention, inventory only

LEVEL 2 — GOVERN (Active)
  Policy: approved list, required attestation
  → Block unauthorized CRITICAL risk apps
  → Require justification for HIGH risk apps
  → Allow LOW/MEDIUM with logging

LEVEL 3 — PROTECT (Real-time)
  Browser extension: LLM DLP
  → Block submission of sensitive data to AI tools
  → Alert + educate, not just block
```

### 6.2 Discovery Layer — AI Tool Usage Detection

```
SIGNAL 1: OAuth App Inventory (Track 1)
  ● Source: Google Workspace + M365 + Slack Admin API
  ● What it finds: AI apps authorized via OAuth (users clicked "Connect with Google")
  ● Examples: ChatGPT Team, Microsoft Copilot, Jasper, Notion AI, Grammarly
  ● Coverage: Any AI tool using OAuth
  ● Blind spot: Browser-based tools WITHOUT OAuth (users go directly to chatgpt.com)

SIGNAL 2: Browser Extension Telemetry (Track 2)
  ● Source: Chrome Extension installed on employee devices (opt-in or MDM push)
  ● What it finds: Direct usage of AI websites (chatgpt.com, claude.ai, etc.)
  ● Privacy: Only domain + timestamp logged — NOT the content
  ● Coverage: Closes the OAuth blind spot
  ● Blind spot: Non-Chrome browsers, mobile

SIGNAL 3: DNS / Network Telemetry (Future — v2)
  ● Source: DNS query logs via Pi-hole / CIRA (if SME controls DNS)
  ● What it finds: AI tool domain resolution (indirect signal)
  ● Privacy: DNS-only, not content
  ● Coverage: All devices on corporate network

Coverage at v1: Signals 1 + 2 = >80% of AI tool usage detected
```

### 6.3 Risk Classification Framework

```
AI Tool Risk Tiers:

CRITICAL — Immediate action required:
  ● AI tools with unclear data export / training opt-out policies
  ● Apps requesting Google Drive / M365 email READ access
  ● Tools with known CVEs or data breach history
  → Response: Auto-revoke OAuth (dry-run → 2-step confirm), alert IT admin

HIGH — Attention needed:
  ● AI tools requesting broad API scopes (files, calendar, contacts)
  ● New AI tools (<6 months old, unverified privacy policy)
  ● AI coding assistants with repository access (GitHub Copilot alternative = OK; unknown tool = HIGH)
  → Response: Require employee attestation ("I understand and accept responsibility")
               IT admin can approve or revoke

MEDIUM — Visible and managed:
  ● AI writing assistants (text only, no file access)
  ● AI image generators (Midjourney, DALL-E — no corporate data access)
  ● AI search (Perplexity — no corporate data integration)
  → Response: Log usage, include in monthly AI usage report to IT admin

LOW — Approved or negligible risk:
  ● Microsoft Copilot (M365 tenant = IT approved, data stays in tenant)
  ● GitHub Copilot (code only, no corporate data in prompts unless dev does it)
  ● Google Duet AI (Google Workspace tenant = IT approved)
  → Response: Inventory only, no alert
```

### 6.4 Policy Enforcement Workflow

```
                 New OAuth App Detected
                          │
                          ▼
              ┌───────────────────────┐
              │  AI Tool Classifier   │
              │  (SageMaker + registry)│
              └───────────┬───────────┘
                          │
              ┌───────────▼────────────────────────────────┐
              │         Risk Tier Decision                  │
              ├─────────────┬──────────────┬───────────────┤
              ▼             ▼              ▼               ▼
            LOW           MEDIUM          HIGH          CRITICAL
              │             │              │               │
           Add to        Log +          Send             Auto-revoke
           approved      monthly        attestation      OAuth (dry-run)
           inventory     report         request to         + alert IT
                                        employee           admin
                                           │
                                    ┌──────▼──────┐
                                    │  Employee   │
                                    │  Response   │
                                    └──────┬──────┘
                                    ┌──────▼──────────────────┐
                                    │ Confirm: I accept risk  │  → Log + allow (90 days)
                                    │ Deny: Remove this access│  → IT revokes OAuth
                                    │ Escalate to IT          │  → IT admin review queue
                                    └─────────────────────────┘
```

### 6.5 AI Governance Dashboard (IT Admin View)

```
AI Usage Summary — Last 30 Days
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  Total AI tools discovered:    47 OAuth apps
  Approved (IT):                12
  Pending attestation:           8
  Auto-revoked (CRITICAL):       3
  Blocked by DLP extension:     126 submissions

Top AI Tools by User Count:
  1. Microsoft Copilot          (38 users) — APPROVED
  2. Grammarly                  (22 users) — LOW (approved)
  3. ChatGPT Team               (15 users) — MEDIUM (logged)
  4. Jasper (AI Writing)         (8 users) — HIGH (awaiting attestation)
  5. [Unknown AI App]            (3 users) — CRITICAL (revoked)

DLP Events This Month:
  HIGH risk blocked:             12 (prevented potential IP leakage)
  MEDIUM warnings shown:         89 (user awareness)
  Content categories blocked:
    ● Source code / credentials:  7
    ● Customer PII:               3
    ● Financial data:             2
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
```

### 6.6 Risks Introduced by AI Tool Usage (and SMESec Mitigation)

| Risk | Description | SMESec Mitigation |
|---|---|---|
| **Inadvertent IP disclosure** | Developers paste proprietary code into ChatGPT | LLM DLP extension blocks/warns. AI governance policy requires attestation. |
| **Training data opt-out gap** | Some AI tools train on user input by default | Risk tier system flags tools with unclear training policies as HIGH/CRITICAL. |
| **Credential leakage** | API keys, passwords pasted into AI prompts | LLM DLP detects credential patterns (API key regex, password patterns) — blocks submission. |
| **PII/GDPR violation** | Employee data, customer data paste into AI | LLM DLP blocks PII (email, phone, SSN, IBAN). GDPR evidence logged. |
| **Shadow AI creating audit gaps** | Decisions made by AI without documentation | AI governance attestation creates audit trail: who used what AI tool, when. |
| **AI-generated disinformation** | Competitor uses AI to create fake news about company | Out of scope for AI governance module. Covered by brand monitoring (v2 roadmap). |
| **Supply chain AI risk** | Third-party vendor using AI in ways that affect SME data | Integration risk scoring in asset inventory flags high-risk vendors. |

---

## 7. Architectural Trade-offs & Rejected Alternatives

### 7.1 Multi-Tenancy: Rejected "Silo" Approach

**Rejected:** Separate database per tenant (Silo model)

**Reasons for rejection:**
- Cost: $100–200/mo/tenant infrastructure × 100 tenants = $10–20K/mo infrastructure for databases alone — not viable at SME pricing ($399/mo Starter tier)
- Operational complexity: 100 databases = 100 patch cycles, 100 backup policies, 100 connection pools
- Scaling: Adding a new tenant requires provisioning, not just a database insert

**Why RLS is sufficient:** PostgreSQL RLS with `FORCE ROW LEVEL SECURITY` applies even to table owners. A bypass is only possible if: (a) an attacker has direct database access (mitigated by VPC private subnet + no public endpoint), or (b) the application explicitly runs as superuser (mitigated by `app_role` without BYPASSRLS privilege).

### 7.2 Track 2 Integration: Rejected "Unified Service" Approach

**Rejected:** Merging Track 1 and Track 2 into a single service

**Reasons for rejection:**
- Track 2 ML models have non-deterministic latency (SageMaker inference: 100ms–2s)
- Merging would mean Track 1 deterministic operations (offboarding <5 min SLA) could be impacted by Track 2 ML latency
- Track 2 failure modes (model drift, SageMaker endpoint cold start) would affect Track 1 availability

**Trade-off accepted:** More complex event-driven integration (EventBridge contract between tracks), but Track 1 SLA is fully independent of Track 2.

### 7.3 Auth: Rejected Auth0/Cognito

**Rejected:** Auth0 or AWS Cognito (managed auth)

**Auth0 rejection reasons:**
- Cost: $0.23/MAU × 500 users/tenant × 50 tenants = $5,750/mo at v1 launch — not sustainable given gross margin target
- SAML 2.0 enterprise feature: Auth0 B2C does not support it. Auth0 B2B (Enterprise): additional cost.
- Data residency: Auth0 does not guarantee EU data stays in EU at Startup tier

**Cognito rejection reasons:**
- Limited SAML customization (SME IT admins need custom SAML attributes for Google federation)
- No built-in TOTP MFA enforcement policy per-tenant
- Cognito User Pools scale poorly for multi-tenant B2B (separate pool per tenant = complexity)

**Keycloak decision rationale:** Self-hosted ECS ($50/mo compute), full OIDC/SAML control, Google + M365 federation native, TOTP enforcement per-realm, no per-user pricing.

### 7.4 Compliance Automation: Rejected "Build Custom"

**Rejected:** Build custom compliance evidence collection system

**Reasons for rejection:**
- Vanta $4–6K/yr vs 3 months engineering ($60K+ cost + ongoing maintenance)
- Vanta has pre-built auditor relationships — auditor trusts Vanta evidence. Custom-built system requires auditor to validate the tool itself.
- Vanta AWS + GitHub connectors collect evidence automatically — replaces hundreds of manual screenshots

**Trade-off accepted:** Vanta vendor dependency. Mitigation: all raw evidence also in S3 Object Lock. If Vanta sunsets, evidence is preserved independently.
