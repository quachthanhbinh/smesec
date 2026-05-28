# SMESec Platform — System Design Document

**Date:** 2026-05-28 | **Version:** 2.0 | **Status:** Final  
**Source:** Synthesized from multi-agent research (Product Owner · Project Manager · Technical Advisor)

---

## Executive Summary

Small and medium enterprises (10–500 employees) face escalating AI-driven security risks — automated spear-phishing, employee data leakage to public LLMs, shadow AI tool sprawl, deepfake fraud, and supply-chain compromise — yet lack dedicated security teams and enterprise budgets. **SMESec** is a unified SaaS protection platform covering the full SME asset surface: data, people, intellectual property, financial accounts, and operational continuity.

**Two-Track Strategy:** All development splits into parallel tracks to eliminate the accuracy risk of AI detection.

- **Track 1 — Foundation & Governance (deterministic, ~100% accuracy):** Asset inventory, access governance, automated offboarding, incident playbooks, compliance reporting. Ships at MVP (Month 3) and v1 (Month 6) independently.
- **Track 2 — AI Threat Detection (ML-gated):** Browser DLP, shadow AI governance, deepfake defense, prompt injection detection. Merges into product only after four accuracy validation gates. If gates not met, Track 1 ships alone.

This document covers all four deliverables: System Architecture, Design Document, Team & Delivery Plan, and AI Governance Module.

---

## 1. System Architecture Diagram

### 1.1 Logical Architecture — Clean Architecture Layers

SMESec applies **Clean Architecture** (Robert C. Martin) + **Hexagonal Architecture** (Ports & Adapters). The Dependency Rule enforces: `Interface → Application → Domain ← Infrastructure`. Domain has zero external dependencies.

```
┌──────────────────────────────────────────────────────────────────────┐
│  INTERFACE LAYER                                                      │
│  Web App (React/Next.js) · Mobile App (Flutter) · Browser Ext (MV3)  │
│  REST/gRPC/WebSocket ← API Gateway (AWS) + Keycloak JWT auth         │
├──────────────────────────────────────────────────────────────────────┤
│  APPLICATION LAYER (Use Cases)                                        │
│  AssetInventorySvc · AccessGovernanceSvc · IncidentPlaybookSvc        │
│  ComplianceSvc · IntegrationSyncSvc · ThreatDetectionSvc (Track 2)   │
├──────────────────────────────────────────────────────────────────────┤
│  DOMAIN LAYER  (zero external dependencies)                           │
│  Entities: Asset · TenantUser · ThreatEvent · Playbook · AccessPolicy│
│  Domain Services: RiskScorer · AccessGovernor · ComplianceAuditor     │
│  Domain Events: AssetDiscovered · ThreatDetected · AccessRevoked      │
├──────────────────────────────────────────────────────────────────────┤
│  INFRASTRUCTURE LAYER (Adapters, implements Domain ports)             │
│  PostgreSQL Repos (RLS) · GoogleWorkspaceAdapter · M365Adapter        │
│  SlackAdapter · AWSIAMAdapter · EventBridgePublisher · HiveClient     │
│  VantaClient · SageMakerClient · SecretsManagerClient                 │
└──────────────────────────────────────────────────────────────────────┘

                 Track 1 and Track 2 share:
          ThreatDetectionEvent schema + EventBridge event bus
          Track 2 events can trigger Track 1 playbooks.
          Track 1 never depends on Track 2 availability.
```

### 1.2 Deployment Architecture — AWS Multi-Region

```
INTERNET
  │ HTTPS (TLS 1.3 only)
  ▼
EDGE ZONE
  Route 53 (GeoDNS: US → us-east-1, EU → eu-west-1)
  → CloudFront CDN → WAF (OWASP rules) → ALB

AWS VPC (private subnets only — no public IPs on compute)
  ├── AUTH: Keycloak ECS Fargate (Multi-AZ, self-hosted SSO/MFA)
  │
  ├── APPLICATION — ECS Fargate services (Go):
  │     Track 1: AssetSvc · AccessSvc · PlaybookSvc · ComplianceSvc · SyncSvc
  │     Track 2: ThreatDetectionSvc · DLPSvc · DeepfakeSvc (Python/FastAPI)
  │
  ├── DATA:
  │     RDS PostgreSQL Multi-AZ (Row-Level Security, tenant_id on every table)
  │     ElastiCache Redis (session tokens, 15-min TTL)
  │     S3 Object Lock (WORM, 7-year audit log retention)
  │
  └── AWS MANAGED SERVICES (outside VPC):
        EventBridge · Step Functions · SNS/SQS
        SageMaker (ML training + inference, Track 2)
        Secrets Manager · KMS (CMK per region) · GuardDuty · Security Hub
        CloudWatch · CloudTrail · IAM

CLIENTS:
  Web Dashboard (Next.js) · Mobile App (Flutter iOS+Android) · Browser Extension (Chrome MV3 + Edge)
```

**Technology Stack:**
- **Backend:** Go (primary API services, integration sync) · Python/FastAPI (ML/AI services)
- **Frontend:** React/Next.js (web) · Flutter (iOS, Android) · Chrome MV3 (browser extension)
- **Auth:** Keycloak (self-hosted ECS, OIDC/SAML 2.0, MFA TOTP mandatory, JWT RS256)
- **ML:** AWS SageMaker (shadow AI risk model, BERT prompt injection classifier)
- **Compliance Automation:** Vanta (connects AWS + GitHub, SOC 2 + ISO 27001 evidence)

### 1.3 Integration Touchpoints

| Service | Method | OAuth Scopes (minimum) | Cadence | Features Enabled |
|---|---|---|---|---|
| **Google Workspace** | OAuth 2.0 + Admin SDK | `admin.directory.user.readonly` `admin.directory.userschema.readonly` `admin.reports.audit.readonly` | 15-min delta sync | User inventory, OAuth app discovery, shadow IT detection, offboarding |
| **Microsoft 365** | OAuth 2.0 + Graph API + webhook | `User.Read.All` `Application.Read.All` `AuditLog.Read.All` `SecurityEvents.Read.All` | 15-min delta + webhook | User inventory, OAuth apps, M365 Defender phishing alerts, offboarding |
| **Slack** | OAuth 2.0 + Admin API | `admin.users:read` `admin.apps:read` `channels:read` | 30-min delta | App inventory, user deactivation (Business+ tier), channel audit |
| **AWS IAM** | IAM assumed role (cross-account) | `iam:ListUsers` `iam:ListRoles` `cloudtrail:LookupEvents` `config:ListDiscoveredResources` | 30-min delta | Cloud resource inventory, IAM policy diff, CloudTrail events |
| **Hive Moderation API** | REST (pay-per-use) | API key (Secrets Manager) | On-demand | Deepfake voice/video detection (<$0.01/check) |
| **Vanta** | Native AWS + GitHub connector | SOC 2 read-only | Continuous | Compliance evidence collection, auditor portal |

**Integration Security Model:** All OAuth tokens stored in AWS Secrets Manager (AES-256, auto-rotation). Read-only by default; write permissions (revocation) requested separately with explicit IT admin consent. Every API call logged with `tenant_id + user_id + action + timestamp`.

---

## 2. Design Document — Core Architectural Decisions

### 2.1 Build vs Buy (Hybrid)

**Decision:** Build the core differentiators (anything customers pay for); buy commodity services (anything that takes >3 months to build for <$5K/yr in vendor cost).

| Component | Decision | Vendor / Technology | Rationale |
|---|---|---|---|
| **Asset Inventory + sync engine** | **Build** (Go) | Google Admin SDK, Graph API | Shadow IT detection logic is core moat; no competitor covers it at SME pricing |
| **Access Governance (RBAC + JIT)** | **Build** (Go) | OPA/Rego policies | SME-optimized offboarding automation is the primary differentiator vs Vanta/Drata |
| **Incident Playbook Engine** | **Build on Step Functions** | AWS Step Functions | Step Functions = proven orchestration; wizard UI for non-security staff is differentiator |
| **Browser Extension DLP** | **Build** (Chrome MV3) | Microsoft Presidio WASM | Local PII inference — content never leaves browser. Privacy moat no competitor matches at SME price |
| **AI tool risk classification** | **Build + curate** | SageMaker + internal registry | No off-the-shelf AI-specific risk scoring at SME context and pricing |
| **SSO / MFA** | **Buy: Keycloak** (self-hosted ECS) | Keycloak | Zero per-user cost ($50/mo compute only) vs Auth0 $5,750/mo at 50 tenants × 500 users |
| **Compliance automation** | **Buy: Vanta** | Vanta Startup plan | $4–6K/yr vs 3 months engineering ($60K+). Auditor trust built in. SOC 2 Type 1 in 60 days. |
| **Deepfake detection** | **Buy: Hive Moderation API** | Hive Moderation | Pay-per-use (<$0.01/check). No training data required. Vendor maintains model updates. |
| **ML platform** | **Buy: AWS SageMaker** | SageMaker | Managed training, endpoint auto-scaling, drift monitoring. vs 6 months custom MLOps. |
| **AI phishing alerts** | **Partner: M365 Defender** | Microsoft Security Graph API | Enterprise-grade detection already in M365. SMESec adds context enrichment + playbook trigger. |

**TCO Year 1 (~50 customers):** Buy costs ~$3,700/mo; gross margin ~91% at $800/mo avg ARR/customer.

### 2.2 Multi-Tenancy Model

**Decision:** Shared PostgreSQL cluster with Row-Level Security (RLS) enforced at the database layer.

**Rejected alternatives:**
- *Silo (DB per tenant):* ~$100–200/mo/tenant infrastructure cost — unviable at SME pricing
- *Shared schema, app-level isolation:* Application bug → cross-tenant data leak. Not trustworthy.

**Implementation:**

Every domain table has two mandatory columns with no exceptions:

```sql
tenant_id      UUID        NOT NULL  -- enforces RLS
data_residency VARCHAR(10) NOT NULL  -- 'US' | 'EU' | 'APAC'

-- PostgreSQL RLS policy (applies even to table owner):
CREATE POLICY tenant_isolation ON assets
  FOR ALL TO app_role
  USING (tenant_id = current_setting('app.tenant_id', TRUE)::UUID);

ALTER TABLE assets ENABLE ROW LEVEL SECURITY;
ALTER TABLE assets FORCE ROW LEVEL SECURITY;  -- blocks superuser too
```

Go API middleware injects `tenant_id` into every PostgreSQL session via `SET LOCAL app.tenant_id` before any query executes. JWT claims are validated, UUID format checked (injection prevention), then session variable is set. A mandatory CI test creates two tenants, inserts data for Tenant A, queries as Tenant B — must return 0 rows. Merges are blocked if this test fails.

**Data residency routing:** EU tenants route to `eu-west-1` ECS + RDS cluster. EU data never written to `us-east-1`. Enforced at: DB schema, S3 bucket policy, KMS key region, and Secrets Manager region. This is a hard invariant from Sprint 1 — retrofitting it later requires full schema migration.

### 2.3 AI-Threat Detection Strategy

**Architecture:** 2-Track separation — deterministic (Track 1) and ML/AI (Track 2) — sharing only a `ThreatDetectionEvent` schema contract and EventBridge event bus.

**Why 2 tracks (not a unified service):**
- Track 1 has deterministic SLA (offboarding <5 min). Track 2 ML inference has 100ms–2s latency non-determinism.
- Track 2 failure modes (model drift, SageMaker cold start) must never degrade Track 1 availability.
- Track 2 events can trigger Track 1 playbooks — but Track 1 never waits on Track 2.

**Track 1 — Deterministic (ships Month 3, 100% accuracy):**

| Threat | Detection Approach | Response |
|---|---|---|
| Shadow IT discovery | OAuth app inventory — scope risk scoring (rule-based matrix) | Alert + allow-list enforcement |
| Orphaned access | Deterministic state machine: employee deactivated in HR ≠ active in SaaS | Auto-offboarding Step Functions workflow |
| Over-provisioning | RBAC diff engine: actual permissions vs defined role policy | Least-privilege recommendation |
| Compliance violations | ISO 27001 / SOC 2 / GDPR control mapping checklist | Compliance gap finding |

**Track 2 — ML/AI Detection (ships Month 6, gated by accuracy):**

| Feature | Technology | Accuracy Gate | Ship Condition |
|---|---|---|---|
| **Shadow AI risk scoring** | SageMaker endpoint (feature vector: OAuth scopes, vendor DPA, app age) | >95% AI tool classification | Sprint 9 eval |
| **LLM DLP (browser ext)** | Presidio WASM (Tier 1 regex) + BERT-tiny ONNX (Tier 2 semantic) | >99% CRITICAL PII, <5% FP | Sprint 8 staging |
| **Deepfake defense** | Hive Moderation API + out-of-band verification (Step Functions) | >80% voice deepfake detection | Sprint 10 eval |
| **Prompt injection** | OWASP regex library (v1) → fine-tuned BERT (v2, Enterprise tier) | TPR >85%, FPR <2% (BERT) | Sprint 11 / Sprint 24 |

**Accuracy gate policy:** No Track 2 feature ships as GA until its accuracy gate is met. Failed gates → feature stays `beta` (opt-in only, no SLA). Track 1 is never held back by Track 2.

### 2.4 Data Privacy Guarantees

Four contractual, architecturally-enforced commitments:

| Commitment | Implementation | Verification |
|---|---|---|
| **No training on customer data** | SageMaker trains on public datasets + synthetic data only. Customer data is never used for model training. | Model card published; architecture review. |
| **Local inference for browser extension** | Presidio WASM runs entirely in-browser. Content typed into AI tools never leaves the user's device. Only pseudonymized metadata (category, severity, timestamp) is sent to servers. | Open-source extension code; network traffic audit. |
| **Immutable audit logs** | S3 Object Lock, COMPLIANCE mode, 7-year retention. No deletion — not by customers, not by SMESec engineers. | AWS Object Lock settings; SOC 2 evidence. |
| **Data residency isolation** | `data_residency` column mandatory from Sprint 1. EU tenant data stays in `eu-west-1` exclusively — enforced at DB, S3, KMS, and Secrets Manager layers. | Tenant isolation CI test; penetration test. |

**Encryption:** RDS AES-256 (KMS CMK), S3 SSE-KMS, TLS 1.3 (external), all secrets in Secrets Manager (auto-rotation, zero plaintext in env vars). Secrets access follows IAM least-privilege: each service can only access its own secret namespace.

**GDPR alignment:** Art. 17 (erasure) via `/api/v1/gdpr/erasure` endpoint (30-day SLA, Sprint 11). Art. 20 (portability) via JSON export endpoint. Art. 25 (privacy by design) via `data_residency` from day 1 and local inference architecture.

**Compliance roadmap:** SOC 2 Type 1 at v1 (Month 6, Vanta evidence from Week 13). SOC 2 Type 2 + ISO 27001 at v2 (Month 12, 6-month observation window from Week 26).

---

## 3. Team & Delivery Plan

### 3.1 Staffing — Grow with Milestones

| Phase | Months | FTE | Team Composition | Milestone |
|---|---|---|---|---|
| **Phase 1** | 1–3 | **6** | Tech Lead · BE#1 · BE#2 · FE#1 · Flutter · DevSecOps(contract) + PM(0.5) | **MVP** (W12) |
| **Phase 2** | 4–6 | **9** | +ML Eng #1 (M4) · +BE#3 Python (M4) · +FE#2 Browser Ext (M4.5) | **v1** (W26) |
| **Phase 3** | 7–9 | **11** | +Customer Success Eng (M7) · +ML Eng #2 (M8, opt.) · DevSecOps → FTE | **v1.5** (W38) |
| **Phase 4** | 10–12 | **11.5** | +Compliance Consultant (contract M10–M12) | **v2** (W52) |

**Phase 3+ team split (2-stream):** Luồng A (65%): new features + SOC 2 Type 2 prep + AI accuracy improvements. Luồng B (35%): pilot feedback, bug fixes, UX polish. Converge at each milestone.

### 3.2 6-Month v1 Delivery Sequence (26 Sprints, 2-week each)

#### Phase 1: Foundation → MVP (S1–S6, Month 1–3)

| Sprint | Deliverable | Gate |
|---|---|---|
| **S1** (W1–2) | AWS infra (VPC/ECS/RDS), Keycloak SSO, multi-tenant schema (`tenant_id + data_residency` from day 1), CI/CD | Tenant isolation CI test green |
| **S2** (W3–4) | Google Workspace sync — users, OAuth apps, shadow IT detection. Dashboard skeleton. | First-value demo <30 min from OAuth grant |
| **S3** (W5–6) | M365 sync + delta link, unified dashboard (Google + M365), risk indicators per user/app | Visibility: all assets from both providers |
| **S4** (W7–8) | Asset classification engine, OAuth scope risk scoring, shadow IT alerts (<15 min), Flutter mobile scaffold | Shadow IT alert pipeline live |
| **S5** (W9–10) | Slack + AWS IAM discovery, RBAC model + least-privilege recommendations, composite identity graph | 4 providers unified in one view |
| **S6** (W11–12) | **🏁 MVP**: Automated offboarding <5 min (Step Functions), 2 incident playbooks (wizard UI), immutable audit log, mobile app beta | Offboarding timed test <5 min in CI |

**MVP = "Can you revoke all access for a departing employee in 5 minutes?"**

#### Phase 2: MVP → v1 (S7–S13, Month 4–6)

| Sprint | Track 1 | Track 2 | Gate |
|---|---|---|---|
| **S7** | JIT access + auto-revoke, access reviews | ML Eng onboards; shadow AI model v0.2 on live OAuth data | Vanta evidence collection active |
| **S8** | Playbook engine (Step Functions), 3 playbooks | LLM DLP browser extension v0.1 (Presidio local inference) | Extension detects PII in text field |
| **S9** | 5 playbooks complete, mobile push notifications | Shadow AI governance v1: AI tool classification + risk scores + attestation workflow | Shadow AI risk scores live |
| **S10** | ISO 27001 + SOC 2 compliance dashboard, Vanta integration | Deepfake defense POC (Hive API), `ThreatDetectionEvent` schema v1 **frozen** | Schema locked — no breaking changes |
| **S11** | Compliance reports (PDF export), GDPR automation | T1-T2 integration: AI threat events → EventBridge → Step Functions playbook auto-trigger | **Highest-risk sprint** — Tech Lead full-time |
| **S12** | SaaS dependency map, penetration test remediation (all Critical/High) | Full T1-T2 end-to-end integration test (automated), Chrome Web Store submission | Pentest: 0 Critical/High open |
| **S13** | **🏁 v1**: Production launch, 5+ pilot customers, SOC 2 Type 1 audit engagement signed | Track 2 features: Shadow AI + LLM DLP extension in v1 | No new features — hardening only |

**v1 gate:** All 7 key requirements delivered. 5+ customers on production. SOC 2 Type 1 audit scheduled.

#### Phase 3 & 4: v1 → v1.5 → v2 (S14–S26, Month 7–12)

| Milestone | Month | Key Additions |
|---|---|---|
| **v1.5** (W38) | 9 | AWS deep integration (CloudTrail), deepfake v2 + AI phishing (M365 Defender), browser extension on Chrome Web Store, pricing tiers enforced, Stripe billing live, 10+ paying customers |
| **v2** (W52) | 12 | SOC 2 Type 2 ✅ · ISO 27001 ✅ · BERT prompt injection (TPR >85%, FPR <2%) · Enterprise tier (SIEM, custom RBAC, dedicated CSM) · All Track 2 features graduate from beta |

### 3.3 Key Requirements Coverage

| Requirement | Milestone | Sprint | Notes |
|---|---|---|---|
| **Asset inventory & classification** | v1 (T6) | S2–S4 core | Google+M365 at MVP. Slack+AWS at S5. Shadow AI detection (Track 2) at S9. |
| **AI-specific threat surface** | v1 (T6) | S7–S11 (Track 2) | Shadow AI governance S9. LLM DLP S8–S9. Deepfake defense + prompt injection S11. |
| **Access governance** | v1 (T6) | S5–S7 | RBAC S5. Offboarding S6 (MVP). JIT access S7. Shadow IT remediation S9. |
| **Continuous compliance posture** | v1 report-ready (T6) | S10–S11 | SOC 2 Type 1 + ISO 27001 reportable at v1. Certification audit at v2 (Month 12). |
| **Incident playbooks** | v1 (T6) | S6 (2), S8–S9 (5 total) | 5 playbooks, AWS Step Functions, wizard UI for non-security staff. |
| **Cost model (tiered pricing)** | v1.5 billing live (T9) | S13 code-ready; S18 Stripe | Starter ($399/mo) · Growth ($799/mo) · Business ($1,499/mo) · Enterprise (custom). |
| **SME tool integrations** | v1 (T6) | S2–S5 | Google Workspace + M365 at MVP. Slack + AWS at S5. QuickBooks deferred to v2. |

### 3.4 Riskiest Assumption to Validate First

> **#1 Risk: SME IT admin (non-technical) can complete Google Workspace + M365 OAuth setup in under 30 minutes using the wizard.**

**Why this is the highest-risk assumption:**
- The entire MVP value prop depends on "first-value in <30 min". If onboarding takes 3 hours (due to M365 admin consent complexity), the pilot program fails before it starts.
- Competitors take 2–4 hours for equivalent setup. If SMESec also takes that long, there is no differentiation.
- This assumption is untestable in a controlled environment — must be tested with real non-technical users on real Google Workspace tenants.

**Validation plan:** End of Sprint 2 (Week 4) — time-boxed usability test with 1–2 non-technical users, no engineer assistance.  
**Go/No-go:** If >45 minutes → redesign wizard before Sprint 3. Block all feature work until this is resolved.

**Top 5 risks (all phases):**

| # | Risk | Phase | Probability | Impact | Mitigation |
|---|---|---|---|---|---|
| 1 | OAuth wizard >30 min for non-technical IT admin | MVP | High | Critical | Usability test W4. IT admin setup guide. Minimum-permission scope explainer. |
| 2 | ML Engineer not hired before W9 | Phase 2 | Medium | High | Begin recruiting W5. Contractor ML fallback. Tech Lead builds SageMaker scaffold at S5. |
| 3 | Track 1–Track 2 integration at S11 delayed >1 sprint | Phase 2 | High | High | Tech Lead full-time S11. API contract frozen S10. Fallback: manual playbook trigger for v1. |
| 4 | Pentest vendor LOI not signed before W14 | Phase 2 | Low | High | PM locks calendar W8. Backup vendor list ready. Hard deadline: no extensions. |
| 5 | SOC 2 Type 2 evidence gap at Month 9 review | Phase 3 | Low | High | Vanta weekly review from W13. PM owns Vanta. Zero-gap policy from W22 onward. |

---

## 4. AI Governance Module

### 4.1 The Problem

78% of knowledge workers use AI tools at work. 52% use tools their employer didn't provide. 11% of content pasted into ChatGPT contains confidential company data (Cyberhaven 2025). The average SME now has 20+ unapproved AI tools connected to company accounts (Nudge Security 2024). BEC losses from AI-powered CEO voice impersonation: $2.9B in 2023 (FBI IC3), avg SME loss $140K/incident.

**No vendor has an affordable, unified solution for the "SME as AI consumer" threat model.** Every serious AI security vendor (HiddenLayer, Wiz, Prompt Security) targets companies deploying LLMs — not companies using them. Nudge Security discovers shadow AI but cannot block it. Prompt Security has browser DLP but costs $15–30K/yr and requires IT admin/developer setup.

### 4.2 Governance Framework: 3 Tiers

```
TIER 3 — PROTECT (Real-time prevention)
  Browser Extension: intercepts before submission, blocks sensitive data
  Deepfake detection: out-of-band verification before acting on suspicious requests
  Prompt injection detection: rule-based (v1) + BERT classifier (v2, Enterprise)

TIER 2 — GOVERN (Policy enforcement)
  AI tool risk scoring + policy engine: block/allow/attest based on OAuth scopes + vendor posture
  Employee attestation workflow: self-reported AI tool usage fills the OAuth blind spot
  Manager approval workflow for risk score 61–85

TIER 1 — DISCOVER (Passive inventory)
  OAuth app inventory (Google + M365 + Slack, every 15 min)
  AI tool classification: 100+ known tools in maintained registry
  Usage telemetry: domain + timestamp only (zero content)

Tier 1 feeds context into Tier 2. Tier 2 policy feeds risk thresholds into Tier 3.
Tier 3 (browser extension) works independently — fails closed if backend is unavailable.
```

### 4.3 Module A — AI Submission Gate (Browser DLP)

The core privacy-preserving architectural decision: **prompt content never leaves the user's browser**.

**3-tier scanning pipeline (all runs in-browser):**

| Tier | Technology | Latency | What It Detects | Accuracy |
|---|---|---|---|---|
| **Tier 1 (Regex)** | OWASP patterns + custom rules, server-push updatable | <1ms | Credit cards (Luhn), SSN, email, phone, API keys (AWS/GitHub/Stripe regex), JWT tokens, IBAN | >99% CRITICAL PII, <1% FP |
| **Tier 2 (WASM BERT-tiny)** | Microsoft Presidio compiled to ONNX WASM (17MB, lazy-loaded) | 50–80ms | Semantic confidential data: "Q3 revenue forecast", M&A discussions, client-specific IP | >85% semantic, <10% FP |
| **Tier 3 (Context, async)** | FastAPI → Lakera Guard API (server-side, non-blocking) | Async | Novel injection patterns, context-aware risk scoring (user role + asset sensitivity multiplier) | >90% precision |

**Supported AI tools (v1, expandable via server-push config):** chatgpt.com · copilot.microsoft.com · gemini.google.com · claude.ai · perplexity.ai · github.com/copilot · notion.so

**Fail-closed guarantee:** If extension cannot complete Tier 1 scan → submission is **blocked** with explicit notice. Never silent pass-through.

**Pre-send Redaction Review UI:**
When sensitive data is detected, the extension shows a blocking modal (not dismissable by Esc):
- Highlights flagged tokens: `[API_KEY_1]` `[EMAIL_1]` `[PHONE_1]` with category labels
- Default action: **"Send with redactions applied"** (placeholders preserve prompt grammar)
- Override: Requires explicit justification text (logged to IT admin dashboard within 60 seconds)
- IT admin sees: type of PII detected, risk severity, action taken — never the actual content

**What is sent to SMESec servers (zero-knowledge architecture):**

```
✅ Sent:   risk_tier, pattern_category, target_domain, timestamp, tenant_id (hashed)
❌ Never:  actual prompt content, flagged tokens, user's text
```

### 4.4 Module C — Shadow AI Governance

**OAuth AI Tool Inventory (C1):** Every 15 minutes, SMESec pulls OAuth app grants from Google Admin SDK + M365 Graph API + Slack Admin API. Each app is classified against a curated registry of 100+ AI tools and risk-scored on a weighted formula:

```
risk_score = (oauth_scopes_sensitivity × 30%) +
             (vendor_DPA_available × 20%) +
             (data_residency_compliance × 15%) +
             (security_certifications × 15%) +
             (known_incidents × 10%) +
             (app_age_days × 5%) +
             (user_count_in_tenant × 5%)
```

| Risk Tier | Example | Response |
|---|---|---|
| **CRITICAL** | Unknown app requesting `gmail.modify` + `drive.readwrite`, no DPA | Alert + auto-revoke (dry-run → 2-step confirm) |
| **HIGH** | Jasper AI with Gmail read access, <6 months old | Alert + require employee attestation ("I understand and accept responsibility") |
| **MEDIUM** | ChatGPT text-only, no file access | Log + monthly AI usage report to IT admin |
| **LOW/PRE-APPROVED** | Microsoft Copilot (M365-native), GitHub Copilot | Inventory only, no alert |

**Attestation Workflow (C2):** Quarterly self-survey cross-references self-reported AI tool usage against the OAuth inventory. Closes the "OAuth blind spot" — employees using ChatGPT directly via browser (no OAuth grant to company account). Non-response after 5 business days = compliance gap finding.

### 4.5 Module D — Deepfake Fraud Defense

**Use case:** "Is this my CEO actually asking me to wire $50K?"

**D1 — Voice Deepfake Detection (non-EU first, legal gate for EU):**
Employee uploads ≤60 second audio clip → Hive Moderation API analyzes (audio NOT stored raw — deleted within 60 seconds). Results shown as probability bands, never binary: *"Likely authentic"* / *"Inconclusive"* / *"Likely synthetic — exercise caution"*. EU deployment requires GDPR Article 9 legal opinion (voice = biometric data) — commissioned Day 1, ships US/UK/AU first.

**D2 — Out-of-Band Verification Workflow (independent of D1, always available):**

```
1. Employee triggers "Verify this person" from mobile app (3 taps)
2. SMESec sends via TWO independent channels to alleged sender:
   - Email: "Did you contact [employee] at [time]?" → [YES / NO] link (no SMESec account needed)
   - SMS: One-time verification code to registered phone → employee asks caller to read it back
3. Combined result within 5 minutes:
   Email "NOT ME" + code not provided → "⚠️ LIKELY FRAUD — Do NOT proceed"
   Email "YES" + code provided → "✅ VERIFIED — Identity confirmed"
   Ambiguous → "⚠️ INCONCLUSIVE — Escalate to IT admin"
4. If fraud confirmed → one-tap trigger Incident Playbook #6 (Deepfake Fraud Response)
5. Full verification timeline stored as compliance evidence (audit log)
```

### 4.6 Module B — Prompt Injection Detection

**v1 (Sprint 11, rule-based, all tiers):** 50+ OWASP LLM Top 10 patterns covering jailbreaks ("Ignore all previous instructions"), role overrides ("You are now DAN"), system prompt extraction ("Print your system prompt"), and data exfiltration instructions. Server-push pattern updates — no extension reinstall required. Accuracy ~75% of known patterns, <5% FP.

**v2 (Sprint 23–24, BERT, Enterprise tier only):** Fine-tuned BERT on 6 months of production data (opt-in Enterprise tenants). Gate: TPR >85% AND FPR <2% on 30-day production holdout set. If not achieved → rule-based remains GA, BERT stays opt-in preview.

**4-tier response by combined risk score (0–100):**

| Score | Action | Notification |
|---|---|---|
| 0–30 | Log only | Weekly digest |
| 31–60 | Warning toast + justification | Daily summary |
| 61–85 | Hard block + manager approval (Slack/email, one-click) | Real-time alert |
| 86–100 | Hard block, no override | P1 alert, IT admin immediate |

### 4.7 Module E — AI Phishing Defense

**M365 Defender Integration (E1):** Pulls phishing/malware alerts from Microsoft Security Graph API (`/v1.0/security/alerts_v2`) every 5 minutes. Enriches with Track 1 context: affected user's role, data access level, direct reports. One-click trigger to Incident Playbook #3 (Phishing Response) from enriched alert. M365-licensed tenants only.

**Email Authentication Posture (E2):** Weekly Google Workspace DMARC/DKIM/SPF audit via Admin SDK. Actionable remediation guidance for misconfigurations (e.g., "DMARC policy is 'none' — emails can be spoofed. Update DNS: p=quarantine").

### 4.8 Module F — Employee Privacy & Transparency

**Transparency Dashboard (F1, mandatory, EU-required):** Always accessible from extension popup and mobile app (IT admin cannot disable). Clearly states: what is monitored (AI tool domain + date), what is NOT (personal browsing, content of prompts, screen/keystrokes). Employee can view their last 10 flagged events.

**Pause Capability (F2, EU consent model):** Employee can pause all monitoring for 15/30/60 min. When paused: zero scanning, zero telemetry, zero transmission. IT admin notified of pause duration only — reason is never logged. Pause duration configurable per role (e.g., CFO role: 0 min max pause).

### 4.9 Competitive Position

| Capability | Nudge Security (closest SME competitor) | Prompt Security | SMESec |
|---|---|---|---|
| Shadow AI discovery | ✅ | ❌ | ✅ |
| Policy enforcement (block/allow) | ❌ nudge only | ✅ | ✅ |
| Browser DLP (zero-knowledge) | ❌ | ✅ | ✅ |
| Deepfake fraud defense | ❌ | ❌ | ✅ |
| Non-expert UX (no IT setup) | ✅ | ❌ requires dev setup | ✅ |
| Compliance evidence (SOC 2) | ❌ | ❌ | ✅ |
| **SME pricing (~50 users)** | **~$2,400/yr** | **$15–30K/yr** | **~$4,800/yr (bundled)** |

**SMESec is the only platform combining all 5 capabilities at SME pricing with zero IT expertise required for setup.**

---

## Appendix: Compliance Certification Timeline

```
Month 3  (W12): Vanta provisioned, compliance evidence collection begins (silent)
Month 4  (W13): Vanta active — SOC 2 control mapping starts
Month 5  (W21): Penetration test begins (vendor LOI signed W14)
Month 6  (W26): v1 LAUNCH → SOC 2 Type 1 audit engagement signed
Month 7  (W27): ISO 27001 gap analysis begins
Month 8  (W33): ISO 27001 Stage 1 audit (documentation review)
Month 9  (W38): v1.5 LAUNCH → SOC 2 Type 2 evidence running since W26
Month 10 (W41): ISO 27001 Stage 2 audit (implementation review)
Month 11 (W46): SOC 2 Type 2 audit fieldwork
Month 12 (W52): v2 LAUNCH → SOC 2 Type 2 ✅ + ISO 27001 ✅ both certified
```

**Note:** SOC 2 Type 2 requires a minimum 6-month observation window. Evidence collection **must start no later than W26** (v1 production launch date) to complete audit by W52. Starting W13 provides a 10-week buffer over the minimum.
