# SMESec Platform — System Design Document (Structured)

**Date:** 2026-05-29 | **Version:** 3.0 | **Methodology:** Problem → Gap → Solution → Plan → Measure → Optimize

---

## 0. Scope Declaration

This document deliberately scopes **Phase 1 to 4 of 7 key requirements**. The remaining 3 are deferred to Phase 2 with explicit rationale. Tackling all 7 at once in a 6-month window would require a full team from Day 1, carry unvalidated ML accuracy risk, and produce a half-built product across all areas rather than a production-ready product in the highest-priority ones.

| # | Requirement | Phase 1 | Phase 2 | Rationale |
|---|---|---|---|---|
| R1 | Asset inventory & classification | ✅ | — | Foundation for everything; fastest to deliver standalone value |
| R2 | Access governance + offboarding | ✅ | — | Highest pain point; single most-asked question by SME IT admins |
| R3 | Incident playbooks (non-security staff) | ✅ | — | Differentiates from pure-monitoring tools; no ML dependency |
| R4 | Continuous compliance posture | ✅ | — | Unlocks paying customers; SOC 2 evidence is a sales gate |
| R5 | AI-specific threat surface (LLM DLP, deepfake, prompt injection) | ❌ | ✅ | ML accuracy unvalidated; must not delay Phase 1 |
| R6 | Cost model / tiered pricing | ❌ | ✅ | Pricing strategy, not a feature; monetization follows product-market fit |
| R7 | Full integration suite (Slack, IAM, QuickBooks) | ❌ | ✅ | Core integrations (Google, M365) already cover 80% of SME workforce |

---

## 1. Current State

### 1.1 Who We Are Serving

Target: SMEs with **10–500 employees**, using cloud-first tooling (Google Workspace or Microsoft 365), with **no dedicated security team**. Security is handled by the IT admin (often a generalist), a senior developer, or the founder directly.

### 1.2 Observed Security Posture (Baseline Assumptions)

These assumptions will be validated with 5 pilot interviews before Sprint 3. They represent the baseline from which we measure improvement.

| Assumption | Source / Basis |
|---|---|
| **A1:** The average SME has 20+ unapproved third-party apps connected to company Google/M365 accounts | Nudge Security — 2024 State of SaaS Security Report |
| **A2:** When an employee leaves, full access revocation takes 1–3 business days on average | Assumption — to be validated in pilot interviews (Week 3–4) |
| **A3:** 70% of SMEs have never audited which OAuth apps have read/write access to employee email and files | Assumption — to be validated via onboarding cohort data (Month 2) |
| **A4:** No SME in our target segment has a documented, tested incident response playbook | Assumption — to be validated in pilot interviews |
| **A5:** 11% of content pasted into public AI tools (ChatGPT, Gemini) contains confidential company data | Cyberhaven — 2024 Data Exposure Report (1.4M workers) |

### 1.3 Current Pain — What Happens Without SMESec

1. **Orphaned access:** Ex-employees retain Google Drive access, Slack channels, and SaaS app logins for days or weeks after departure.
2. **OAuth sprawl:** No IT admin has a live view of which third-party apps can read employee email or access company files.
3. **No playbook:** When a phishing attack hits, the company improvises. No documented steps, no assigned roles, no audit trail.
4. **Compliance gap:** SOC 2 / ISO 27001 evidence is collected manually (spreadsheets), if at all. The audit preparation process alone costs 3+ months of engineering time.

---

## 2. Goals & Objectives

### 2.1 North Star Goal

> Enable a non-technical IT admin at a 50-employee company to achieve enterprise-grade security governance — asset visibility, access control, and compliance readiness — without hiring a dedicated security professional.

### 2.2 Phase 1 OKRs (6 months, measurable)

| Objective | Key Result | Gate |
|---|---|---|
| **O1: Reduce orphaned access exposure** | Mean time to fully revoke all access for a departing employee ≤ 5 minutes | MVP (Month 3) |
| **O2: Surface hidden OAuth risk** | ≥ 80% of connected third-party apps discovered and risk-classified within 15 minutes of granting OAuth consent | MVP (Month 3) |
| **O3: Enable non-expert incident response** | IT admin can initiate a full phishing response playbook (5 steps) with 0 prior security training, measured by usability test completion rate ≥ 80% | v1 (Month 6) |
| **O4: Achieve compliance readiness** | SOC 2 Type 1 audit engagement signed with 0 manual evidence collection required | v1 (Month 6) |
| **O5: Validate product-market fit** | 5 paying pilot customers on production before Month 6 ends | v1 (Month 6) |

### 2.3 What Success Looks Like at v1

A 50-employee company with one IT admin who has no security background can:
- See every OAuth app connected to their Google and M365 tenants within 15 minutes of signup
- Revoke all access for a departing employee in under 5 minutes
- Run a phishing response playbook without calling an external consultant
- Generate a SOC 2-ready compliance report with one click

---

## 3. GAP Analysis

```
CURRENT STATE                          TARGET STATE (v1, Month 6)
─────────────────────────────────      ─────────────────────────────────────
No asset inventory                  →  Live inventory: all users, OAuth apps,
                                        devices, third-party integrations
                                        (Google + M365, updated every 15 min)

Offboarding = manual Jira ticket    →  Automated offboarding workflow:
(1–3 business days)                     all access revoked in < 5 minutes
                                        via single trigger

No incident playbooks               →  5 guided playbooks (wizard UI),
                                        executable by non-security staff

Compliance = spreadsheets           →  SOC 2 Type 1 evidence collected
(3+ months manual prep)                 automatically via Vanta integration
                                        (continuous, zero manual effort)
```

**Key Gaps to Close:**

| Gap | Root Cause | Our Solution |
|---|---|---|
| **G1: No visibility** | No tool aggregates OAuth grants across providers | Integration sync engine (Google Admin SDK + M365 Graph API) |
| **G2: Slow offboarding** | Manual, multi-tool process with no single trigger | Automated Step Functions workflow with idempotency + rollback |
| **G3: No playbooks** | Creating them requires security expertise most SMEs don't have | Pre-built playbook library with a wizard UI — zero expertise needed |
| **G4: Compliance overhead** | Evidence collection is manual, fragmented | Vanta connector (automated SOC 2 evidence from AWS + GitHub) |
| **G5: No audit trail** | Actions are undocumented; GDPR erasure is unsupported | Immutable S3 WORM log + KMS envelope encryption (GDPR-compliant) |

**What Phase 1 does NOT close** (intentional deferral):
- AI-generated phishing / deepfake fraud detection (Phase 2 — ML accuracy gate required)
- LLM data leakage via browser extension (Phase 2)
- Slack and AWS IAM integration (Phase 2)

---

## 4. Solution

### 4.1 Architecture Overview

SMESec is a multi-tenant SaaS platform built on **Clean Architecture** (Dependency Rule: Interface → Application → Domain ← Infrastructure) and deployed on AWS.

```
┌──────────────────────────────────────────────────────────────────────┐
│  INTERFACE LAYER                                                      │
│  Web Dashboard (React/Next.js)  ·  Mobile App (Flutter)              │
│  REST/WebSocket ← AWS API Gateway + Keycloak JWT (MFA mandatory)     │
├──────────────────────────────────────────────────────────────────────┤
│  APPLICATION LAYER (Phase 1 Use Cases)                                │
│  AssetInventorySvc  ·  AccessGovernanceSvc  ·  PlaybookSvc           │
│  ComplianceSvc  ·  IntegrationSyncSvc                                 │
├──────────────────────────────────────────────────────────────────────┤
│  DOMAIN LAYER  (zero external dependencies)                           │
│  Entities: Asset · TenantUser · Playbook · AccessPolicy              │
│  Services:  AccessGovernor · ComplianceAuditor · RiskScorer           │
│  Events:    AssetDiscovered · AccessRevoked · PlaybookTriggered       │
├──────────────────────────────────────────────────────────────────────┤
│  INFRASTRUCTURE LAYER                                                 │
│  PostgreSQL (RLS, tenant_id on every table)  ·  S3 Object Lock WORM  │
│  GoogleWorkspaceAdapter  ·  M365Adapter                               │
│  VantaClient  ·  EventBridgePublisher  ·  SecretsManagerClient        │
└──────────────────────────────────────────────────────────────────────┘
```

### 4.2 Deployment (AWS, Phase 1)

```
INTERNET → Route 53 → CloudFront → WAF (OWASP rules) → ALB
  │
  └── AWS VPC (private subnets only)
        ├── Keycloak ECS Fargate (min 2 tasks, JWKS cache)
        ├── App Services: ECS Fargate (Go) — AssetSvc · AccessSvc · PlaybookSvc
        ├── RDS PostgreSQL Multi-AZ (Row-Level Security)
        ├── ElastiCache Redis (session tokens, 15-min TTL)
        ├── S3 Object Lock (WORM, 7-year audit log)
        └── AWS Managed: EventBridge · Step Functions · Secrets Manager · KMS · GuardDuty
```

**Technology choices:**
- **Backend:** Go (API, sync engine) — chosen for concurrency model matching integration sync workloads
- **Frontend:** React/Next.js (web), Flutter (mobile)
- **Auth:** Keycloak self-hosted on ECS — $0/user vs Auth0 at ~$115,000+/mo for 1K tenants (500K MAU × $0.23/MAU). Keycloak saves ~$500K/yr at v1 target scale.
- **Compliance:** Vanta — $4–6K/yr vs 3 months engineering ($60K+)

### 4.3 Core Architectural Decisions

#### Build vs Buy

**Rule:** Build what customers pay for. Buy what takes >3 months to build for <$5K/yr in vendor cost.

| Component | Decision | Why |
|---|---|---|
| Asset Inventory + sync engine | **Build** (Go) | Shadow IT detection is our core moat — no vendor covers it at SME pricing |
| Access Governance (RBAC + offboarding) | **Build** (Go + OPA/Rego) | SME-optimized automation is the primary differentiator |
| Incident Playbook Engine | **Build** on AWS Step Functions | Wizard UI for non-security staff is our differentiator; orchestration itself is commodity |
| SSO / MFA | **Buy:** Keycloak (self-hosted ECS) | Zero per-user cost; enterprise features out of the box |
| Compliance automation | **Buy:** Vanta | $4–6K/yr vs $60K+ to build; auditor trust is built in |

#### Multi-Tenancy

Shared PostgreSQL cluster with **Row-Level Security (RLS)** enforced at the database layer — not the application layer.

```sql
-- Every table has these two mandatory columns:
tenant_id      UUID        NOT NULL
data_residency VARCHAR(10) NOT NULL  -- 'US' | 'EU' | 'APAC'

-- RLS policy (applies to table owner too):
CREATE POLICY tenant_isolation ON assets
  FOR ALL TO app_role
  USING (tenant_id = current_setting('app.tenant_id', TRUE)::UUID);
ALTER TABLE assets FORCE ROW LEVEL SECURITY;
```

**Why not DB-per-tenant:** $100–200/mo/tenant infrastructure cost at SME pricing = unviable unit economics.

**Mandatory CI test:** Creates two tenants, inserts data for Tenant A, queries as Tenant B → must return 0 rows. Merges blocked if this test fails.

#### Data Privacy (4 commitments)

| Commitment | Implementation |
|---|---|
| No training on customer data | SageMaker trains on public datasets + synthetic data only |
| Immutable audit logs (GDPR-erasable) | S3 Object Lock WORM + per-tenant KMS key → key destruction = GDPR "effective erasure" (EDPB Rec 01/2020) |
| Data residency isolation | `data_residency` mandatory from Sprint 1; EU data stays in `eu-west-1` — enforced at DB, S3, KMS layers |
| Secrets management | All OAuth tokens in AWS Secrets Manager (AES-256, auto-rotation). Zero plaintext in env vars. |

### 4.4 Phase 1 Integration Touchpoints

| Integration | Method | Sync Cadence | Value Delivered |
|---|---|---|---|
| **Google Workspace** | OAuth 2.0 + Admin SDK | 15-min delta | User inventory, OAuth app discovery, shadow IT detection, offboarding trigger |
| **Microsoft 365** | OAuth 2.0 + Graph API + webhook | 15-min delta + webhook | User inventory, OAuth apps, M365 Defender alerts, offboarding |
| **Vanta** | Native AWS + GitHub connector | Continuous | Automated SOC 2 evidence collection; auditor portal |

*Phase 2 additions: Slack, AWS IAM, Hive Moderation API (deepfake)*

---

## 5. Delivery Plan

### 5.1 Riskiest Assumption — Validate First

> **#1 Risk: A non-technical IT admin can complete Google Workspace + M365 OAuth setup in under 30 minutes using our onboarding wizard.**

This is the gating assumption for the entire MVP value proposition. If onboarding takes 3 hours, the pilot program fails before the product is evaluated.

**Validation plan:** Week 4 — timed usability test with 2 non-technical users on real Google Workspace tenants. No engineer assistance.

**Go/No-go:** > 45 minutes → redesign wizard before Sprint 3. No feature work proceeds until this is resolved.

### 5.2 Team — Phase 1 (Month 1–3, MVP)

| Role | Type | When |
|---|---|---|
| Tech Lead | FTE | Day 1 |
| Backend Engineer #1 (Go) | FTE | Day 1 |
| Backend Engineer #2 (Go) | FTE | Day 1 |
| Frontend Engineer #1 (React) | FTE | Day 1 |
| Flutter Engineer | FTE | Day 1 |
| DevSecOps | Contract | Day 1 |
| PM | 0.5 FTE | Day 1 |
| BD Consultant | Contract (3 days/wk) | Week 1 — recruits pilot customers |

*Note: ML Engineer #1 intentionally deferred to Phase 2. Phase 1 is 100% deterministic — no ML dependency.*

### 5.3 Sprint Plan — Phase 1: Foundation → MVP (Sprints 1–6, Month 1–3)

| Sprint | Deliverable | Gate |
|---|---|---|
| **S1** (W1–2) | AWS infra (VPC/ECS/RDS), Keycloak SSO (2 ECS tasks, JWKS cache), multi-tenant schema (`tenant_id + data_residency`), CI/CD pipeline, S3 Object Lock with per-tenant KMS key | Tenant isolation CI test green |
| **S2** (W3–4) | Google Workspace sync — users, OAuth apps, shadow IT detection. Dashboard skeleton. | **Usability test: onboarding ≤ 30 min** with real non-technical user |
| **S3** (W5–6) | M365 sync + delta link, unified dashboard (Google + M365), risk indicators per user/app | All assets visible from both providers |
| **S4** (W7–8) | Asset classification engine, OAuth scope risk scoring, shadow IT alerts (< 15 min), Flutter mobile scaffold | Shadow IT alert pipeline live |
| **S5** (W9–10) | RBAC model, least-privilege recommendations, composite identity graph | Recommendations generated for ≥ 3 pilot accounts |
| **S6** (W11–12) | **🏁 MVP:** Automated offboarding < 5 min (Step Functions, grace period configurable, rollback 24h, idempotency key), 2 incident playbooks (wizard UI), immutable audit log | Offboarding timed test < 5 min in CI · Rollback test passes |

**MVP definition:** *"Can you revoke all access for a departing employee in 5 minutes?"* — yes, with audit trail.

### 5.4 Sprint Plan — Phase 2: MVP → v1 (Sprints 7–13, Month 4–6)

| Sprint | Deliverable | Gate |
|---|---|---|
| **S7** | JIT access + auto-revoke, access review workflow | Vanta evidence collection starts |
| **S8** | Playbook engine (Step Functions), 3 playbooks total | Extension detects PII in text field (Track 2 parallel start) |
| **S9** | 5 playbooks complete, mobile push notifications | Playbook usability test: 80% completion rate without guidance |
| **S10** | ISO 27001 + SOC 2 compliance dashboard, Vanta integration | Vanta dashboard shows 0 manual gaps |
| **S11** | Compliance PDF report export, GDPR erasure endpoint | PDF report generated in < 30 seconds |
| **S12** | SaaS dependency map, penetration test remediation | Pentest: 0 Critical/High open |
| **S13** | **🏁 v1:** Production launch, 5+ pilot customers, SOC 2 Type 1 audit engagement signed | 5 customers live · SOC 2 audit booked |

### 5.5 Top Risks

| # | Risk | Mitigation |
|---|---|---|
| 1 | OAuth wizard > 30 min for non-technical IT admin | Usability test W4. Redesign before S3 if needed. |
| 2 | M365 webhook subscriptions expire (3-day TTL) | EventBridge Scheduler renewal job + 410 Gone → full sync fallback, built in S1 |
| 3 | Track 1–Track 2 integration sprint (S11) delayed | Tech Lead full-time S11. API contract frozen S10. Fallback: manual playbook trigger for v1. |
| 4 | Pentest vendor not engaged before W14 | PM locks calendar W8. Hard deadline — no extensions. |
| 5 | 5 pilot customers not recruited by Month 5 | BD Consultant starts Week 1 with explicit target: 3 LOIs by Month 3. |

---

## 6. Measurement & Evaluation

### 6.1 KPIs per Milestone

| KPI | MVP Target (Month 3) | v1 Target (Month 6) |
|---|---|---|
| Onboarding time (OAuth wizard, non-technical user) | ≤ 30 min | ≤ 20 min |
| Mean time to offboard departing employee | ≤ 5 min | ≤ 3 min |
| OAuth app discovery latency | ≤ 15 min from grant | ≤ 10 min |
| Playbook completion rate (untrained user, no help) | — | ≥ 80% |
| Compliance evidence gaps (Vanta) | — | 0 critical gaps |
| Pilot customers on production | 0 (internal staging only) | ≥ 5 |
| Pentest: open Critical/High findings | 0 | 0 |

### 6.2 Instrumentation Plan

Every KPI above has a defined measurement source:

| KPI | Measurement Source | Owner |
|---|---|---|
| Onboarding time | Usability test recording (W4) + production funnel analytics | PM |
| Offboarding time | Automated CI test (Step Functions execution log); production p95 latency | Tech Lead |
| OAuth discovery latency | Sync job timestamp delta (logged per tenant) | BE#1 |
| Playbook completion rate | Wizard step-completion event in product analytics | FE#1 |
| Compliance gaps | Vanta weekly report export | PM |
| Pilot customers | CRM (BD Consultant) | BD |

### 6.3 Go/No-Go Gates

These are hard gates — no bypass, no extension:

| Gate | Condition | If Failed |
|---|---|---|
| **G1 (W4):** Onboarding | Non-technical user completes OAuth setup ≤ 30 min | Redesign wizard; delay S3 by 1 sprint |
| **G2 (W12):** MVP | Offboarding CI test passes ≤ 5 min; rollback test passes | MVP delayed; no Phase 2 start |
| **G3 (W22):** Pentest | 0 Critical/High open findings | No v1 launch; fix-then-retest |
| **G4 (W26):** v1 | 5 customers live; SOC 2 audit engagement signed | v1 delayed by 2 weeks max, then go with what's live |

---

## 7. Optimization & Iteration

### 7.1 Post-MVP Feedback Loop (Month 3–6)

After MVP, two parallel streams run simultaneously:

| Stream | Allocation | Focus |
|---|---|---|
| **Stream A** | 65% | New Phase 2 features (playbooks, compliance, access reviews) |
| **Stream B** | 35% | Pilot feedback, UX polish, bug fixes, KPI gap closure |

Stream B is not optional. Without it, pilot customers churn before v1 launches. Both streams converge at the v1 gate (S13).

### 7.2 Phase 2 Decisions — What We Will Reassess After v1

These decisions will be made at the v1 retrospective (Month 6), informed by real data:

| Decision | Criteria for Phase 2 Inclusion |
|---|---|
| **AI threat surface (LLM DLP, deepfake)** | ML Engineer #1 hired; Track 2 accuracy gate results reviewed; pilot customer demand confirmed |
| **Slack + AWS IAM integrations** | ≥ 3 pilot customers request it before Month 5 |
| **Prompt injection detection (Lakera Guard)** | ≥ 2 customers report prompt injection incidents; Lakera API cost validated at pilot volume |
| **Tiered pricing / Stripe billing** | ≥ 5 paying customers confirm willingness to upgrade; pricing model A/B test complete |
| **BERT prompt injection (internal model)** | Only if Lakera cost > $2K/mo at Phase 2 volume AND ≥ 50K labeled samples available |

### 7.3 Architecture Evolution Path

Phase 1 makes deliberate, documented trade-offs that have known migration paths:

| Trade-off | Phase 1 Decision | Phase 2+ Migration |
|---|---|---|
| Shared DB (RLS) vs DB-per-tenant | Shared PostgreSQL + RLS — viable up to ~2,000 tenants (RDS Proxy mandatory from Sprint 1 for 1K target) | If > 2K tenants: shard clusters (see [16-large-scale-architecture.md](16-large-scale-architecture.md) §4) |
| Keycloak self-hosted | 2 ECS tasks, $50/mo — zero per-user cost | v1.5 retrospective: if ops load high → evaluate WorkOS/Auth0 |
| Vanta for compliance | $4–6K/yr buy — faster than build | If custom compliance rules needed at v2 → extend via Vanta API + custom evidence collectors |
| Track 2 deferred | No ML in Phase 1 — zero accuracy risk | Phase 2: ML Eng #1 joins M1; 3 months R&D before any GA feature |

---

## Appendix A: Requirement Coverage Traceability

| Requirement | Phase | Sprint | Delivered Feature |
|---|---|---|---|
| Asset inventory & classification | Phase 1 | S2–S4 | Google + M365 user/app inventory, OAuth risk scoring, shadow IT alerts |
| Access governance + offboarding | Phase 1 | S5–S7 | RBAC model, least-privilege recommendations, automated offboarding (Step Functions) |
| Incident playbooks | Phase 1 | S6–S9 | 5 playbooks, wizard UI, Step Functions orchestration |
| Compliance posture | Phase 1 | S10–S11 | SOC 2 + ISO 27001 dashboard, Vanta integration, PDF export, GDPR erasure endpoint |
| AI threat surface | Phase 2 | S14+ | Shadow AI governance, LLM DLP (browser extension), deepfake defense, prompt injection |
| Cost model / pricing | Phase 2 | S18 | Starter/Growth/Business tiers, Stripe billing live |
| Full integrations (Slack, IAM) | Phase 2 | S14–S16 | Slack Admin API, AWS IAM cross-account assumed role |

## Appendix B: Key Assumptions to Validate

| ID | Assumption | Validation Method | Owner | Deadline |
|---|---|---|---|---|
| A1 | Average SME has 20+ unapproved OAuth apps | Onboarding cohort data (Month 2 avg) | BE#1 | W8 |
| A2 | Manual offboarding takes 1–3 business days | 5 pilot interviews | BD | W4 |
| A3 | 70% of SMEs have never audited OAuth grants | Onboarding cohort: % with 0 prior revocations | PM | W8 |
| A4 | No SME has a documented incident playbook | Pilot interview question #3 | BD | W4 |
| A5 | Non-technical IT admin can complete OAuth setup ≤ 30 min | Timed usability test (W4) | Tech Lead | W4 |
