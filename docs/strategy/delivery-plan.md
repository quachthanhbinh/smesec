# SMESec Platform — Delivery Plan

**Date:** 2026-05-28  
**Status:** Approved — Synthesized from 3 agents (Product Owner · Project Manager · Technical Advisor)  
**Version:** 1.0  
**Scope:** Full roadmap from Sprint 1 to v2 (12 months)

---

## Table of Contents

1. [Roadmap Overview](#1-roadmap-overview)
2. [Scope by Milestone](#2-scope-by-milestone)
3. [Team & Headcount Ramp](#3-team--headcount-ramp)
4. [Sprint Breakdown](#4-sprint-breakdown)
   - [Phase 1: Foundation → MVP (Month 1–3, S1–S6)](#phase-1-foundation--mvp-month-13-s1s6)
   - [Phase 2: MVP → v1 (Month 4–6, S7–S13)](#phase-2-mvp--v1-month-46-s7s13)
   - [Phase 3: v1 → v1.5 (Month 7–9, S14–S20)](#phase-3-v1--v15-month-79-s14s20)
   - [Phase 4: v1.5 → v2 (Month 10–12, S21–S26)](#phase-4-v15--v2-month-1012-s21s26)
5. [Two-Stream Team Split (Post-v1)](#5-two-stream-team-split-post-v1)
6. [Key Requirements Coverage](#6-key-requirements-coverage)
7. [Riskiest Assumption to Validate First](#7-riskiest-assumption-to-validate-first)
8. [Compliance Certification Timeline](#8-compliance-certification-timeline)
9. [External Dependencies & Hard Deadlines](#9-external-dependencies--hard-deadlines)

---

## 1. Roadmap Overview

```
Month:  1    2    3    4    5    6    7    8    9   10   11   12
Sprint: S1  S2  S3  S4  S5  S6  S7  S8  S9  S10 S11 S12 S13 S14 S15 S16 S17 S18 S19 S20 S21 S22 S23 S24 S25 S26
        |------Phase 1: Foundation-------|----Phase 2: MVP→v1----|---Phase 3: v1→v1.5-----|----Phase 4: v1.5→v2----|
                            ↑                                   ↑                   ↑                              ↑
                           MVP                                  v1                 v1.5                            v2
                        (W12/M3)                           (W26/M6)           (W38/M9)                       (W52/M12)
```

| Milestone | Week | Month | Description |
|-----------|------|-------|-------------|
| **MVP** | W12 | M3 | Asset inventory + automated offboarding + shadow IT — pilot customers onboard |
| **v1** | W26 | M6 | All key requirements delivered. SOC 2 Type 1 audit scheduled. |
| **v1.5** | W38 | M9 | Advanced AI detection + AWS v1.1 + pilot feedback integrated. Billing tiers live. |
| **v2** | W52 | M12 | SOC 2 Type 2 + ISO 27001 certified. Enterprise tier. ML features production-ready. |

---

## 2. Scope by Milestone

### Feature Map by Phase

| Feature Domain | MVP (T3) | v1 (T6) | v1.5 (T9) | v2 (T12) |
|---|---|---|---|---|
| **Asset Inventory** | Google WS + M365: users, OAuth apps, basic devices | + Slack, AWS, Shadow AI detection | + Custom asset types, dependency map | + Full cloud posture, peer anomaly |
| **Access Governance** | Automated offboarding <5 min, RBAC dashboard | + JIT access, access reviews, shadow IT remediation | + Risk scoring, access policy templates | + Peer group anomaly, insider threat signal |
| **AI Threat Surface** | ❌ (Track 2 in R&D) | Shadow AI governance + LLM DLP browser ext (beta) | + Deepfake defense, AI phishing, prompt injection v1 | + Prompt injection ML (BERT), advanced analytics |
| **Compliance Posture** | Evidence collection begins (silent, no UI) | Dashboard compliance, SOC 2 Type 1 + ISO 27001 report-ready | SOC 2 Type 2 evidence running (90 days) | SOC 2 Type 2 certified + ISO 27001 certified |
| **Incident Playbooks** | 2 playbooks (Offboarding, Cred Compromise) | 5 playbooks, wizard UI, AWS Step Functions | + Custom playbook builder, mobile triggers | + Playbook analytics, ML suggestions |
| **Integrations** | Google WS + M365 (OAuth wizard <30 min) | + Slack full + AWS IAM basic | + AWS CloudTrail, S3 audit, IAM deep | + SIEM (Splunk/QRadar), custom webhooks |
| **Mobile App** | ❌ TestFlight/Beta | Alerts + playbook trigger (iOS + Android) | Full incident response mobile | Full feature parity |
| **Billing / Pricing** | Manual invoicing (pilot free) | Starter + Growth tiers code-ready | Pricing tiers enforced, billing live | Enterprise custom + usage-based |

### MVP — Minimum Viable Value Definition

> **Question MVP must answer:** *"Do you know how many applications are connected to your Google Workspace / M365, and can you revoke all access for a departing employee in 5 minutes?"*

```
MVP = Sprint 6 complete (end of Week 12)

✅ OAuth wizard: Google Workspace + M365, setup <30 min
✅ Asset inventory dashboard: users, OAuth apps, basic devices
✅ Shadow IT detection: alert on new OAuth app in <15 min
✅ Automated offboarding: revoke all access <5 min
✅ 2 incident playbooks: Offboarding Emergency + Credential Compromise
✅ RBAC dashboard: view permissions, least-privilege recommendations
✅ Keycloak SSO + MFA mandatory
✅ Compliance evidence collection: begins running silently from day 1
✅ Mobile app beta: TestFlight + Play Console

❌ NOT in MVP:
  - AI/ML detection (Track 2 in R&D)
  - Full compliance reports
  - JIT access
  - Slack integration
  - Billing system
  - Deepfake / prompt injection
```

---

## 3. Team & Headcount Ramp

### Headcount Timeline

```
Month 1–3 (Phase 1 / MVP):         6 FTE core + DevSecOps contract
Month 4 (start of Phase 2):        + ML Engineer #1 (Track 2)
Month 4 (Sprint 7):                + Backend Engineer #3 (Track 2)
Month 4–5 (Sprint 8):              + Frontend Engineer #2 (Browser Extension)
Month 6 (Sprint 13 / v1 launch):   DevSecOps → FTE (no longer contract)
Month 7 (start of Phase 3):        + Customer Success Engineer
Month 8 (mid Phase 3):             + ML Engineer #2 (optional, depending on v1 velocity)
Month 10–12 (Phase 4):             + Compliance Consultant (contract)
```

### Detailed Headcount by Phase

| Role | T1–T3 | T4–T6 | T7–T9 | T10–T12 | Track |
|---------|-------|-------|-------|---------|-------|
| Tech Lead / Architect | ✅ 1.0 FTE | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | 1 |
| Backend Eng #1 (Go) | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | 1 |
| Backend Eng #2 (Go/Python) | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | 1 |
| Frontend Eng #1 (React) | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | 1 |
| Flutter / Mobile Eng | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | 1 |
| DevSecOps | Contract (0.5) | Contract (0.5) | **FTE (1.0)** | **FTE (1.0)** | Shared |
| PM | 0.5 | 0.5 | 0.5 | 0.5 | Shared |
| **ML Engineer #1** | — | **✅ 1.0 (T4)** | ✅ 1.0 | ✅ 1.0 | 2 |
| **Backend Eng #3 (Python/FastAPI)** | — | **✅ 1.0 (T4)** | ✅ 1.0 | ✅ 1.0 | 2 |
| **Frontend Eng #2 (Browser Ext)** | — | **✅ 1.0 (T4.5)** | ✅ 1.0 | ✅ 1.0 | 2 |
| **Customer Success Engineer** | — | — | **✅ 1.0 (T7)** | ✅ 1.0 | Customer |
| **ML Engineer #2** | — | — | **✅ 1.0 (T8, opt.)** | ✅ 1.0 | 2 |
| **Compliance Consultant** | — | — | — | **Contract (T10–T12)** | Compliance |
| **Total FTE** | **6** | **8.5 → 9** | **10 → 11** | **11.5** | |

> **Hiring principle:** ML Engineer must have SageMaker or managed ML platform experience — not an academic researcher. Begin recruiting from Week 5 (Sprint 3) to onboard in time for Sprint 6.

> **⚠️ R-C5 (No customer acquisition plan):** BD Consultant (contract 3 days/week, $60–80/hr) onboard **Week 1** — not after product is ready. Year 1 goals:
> - **MSP Partner Program:** Partner with 3 MSP/IT consultant firms in first 6 months (CAC via MSP: $500–$800 vs $3,000–$5,000 direct)
> - **Freemium "Security Health Check":** Free tier (5 users, read-only, 14 days) → demo-to-paid conversion funnel (CAC <$300)
> - **Pilot funnel target:** 100 outreach → 30 qualified → 15 demo → 5 pilot (from Month 1)

---

## 4. Sprint Breakdown

### Phase 1: Foundation → MVP (Month 1–3, S1–S6)

**Team Phase 1:** Tech Lead · BE1 · BE2 · FE1 · Flutter · DevSecOps(contract) · PM = **6 FTE**

---

#### S1 — W1–2: Infrastructure & Auth

| | |
|---|---|
| **Goal** | Technical foundation: deployable, users can log in |
| **Sprint deliverable** | Engineer logs into web app with real Google/M365. Staging deployed from CI automatically. |
| **Scope** | AWS VPC + ECS Fargate + RDS PostgreSQL Multi-AZ · S3 Object Lock (audit log, envelope encryption per-tenant KMS key) · Keycloak SSO (Google + M365) · JWKS cache (6-hour TTL, serve-stale-on-failure) · MFA TOTP mandatory · CI/CD GitHub Actions · Multi-tenant schema (`tenant_id` + `data_residency` on all tables, RLS enforced) · `ThreatDetectionEvent` interface draft (T1-T2 schema contract) · **⚠️ R-C3 (Mandatory):** `subscription_registry` table schema for M365 webhook renewal service · EventBridge Scheduler skeleton for 12-hour renewal job · **R-C2 (Design):** Google rate limit strategy — per-cluster GCP service account layout, quota distribution algorithm |
| **Key risks** | Auth provider decision (Auth0 vs Cognito vs Keycloak self-host) must be finalized Day 1 |
| **PM action** | Begin recruiting ML Engineer #1. Prepare list of potential pilot customers. **BD Consultant (contract 3 days/week) must onboard Week 1 (R-C5).** |

> **Mandatory gate:** `data_residency` column must be present from S1 — if skipped, full schema refactor required near MVP. Tenant isolation CI test must be green before merging any code.

---

#### S2 — W3–4: Google Workspace Sync

| | |
|---|---|
| **Goal** | View users + OAuth apps from Google Workspace |
| **Sprint deliverable** | Dashboard displays user list + OAuth apps from real Google tenant. First-value demo <30 min from OAuth grant. |
| **Scope** | Google Admin SDK: user/group/device sync · OAuth app discovery (scope risk analysis) · 15-min incremental sync (delta pull) · Asset inventory DB schema v1 · Shadow IT detection rules v1 (high-risk OAuth scopes) · Dashboard skeleton (data visible, no styling required) |
| **Key risks** | Google Admin SDK pagination + rate limits — validate on real tenant >100 users in S1 skeleton |
| **PM action** | Pilot outreach begins. Target 3–5 SMEs (50–200 employees) for Month 3 onboard. |

---

#### S3 — W5–6: M365 Sync + Dashboard v1

| | |
|---|---|
| **Goal** | Unified dashboard: Google + M365 on one screen |
| **Sprint deliverable** | Dashboard displays assets from both Google + M365. Risk indicators per user/app. Export CSV. |
| **Scope** | Microsoft Graph API + Azure AD: user/app/device sync · M365 delta link + webhook · **⚠️ R-C3: Deploy webhook renewal job** (already designed from S1 schema) — 410 Gone handler + DLQ + polling fallback + staleness UI indicator · Cross-provider identity matching (email canonical) · Unified risk indicators (per-provider, not composite) · Dashboard polish: filter, search, sort |
| **Key risks** | M365 OAuth permission consent — need detailed IT Admin guide. Prepare "minimum-permission scope explainer" for customers. **Webhook renewal CANNOT be skipped — silent failure here cascades permanently.** |
| **PM action** | Sign ML Engineer #1 LOI this week. Pilot customer list must have at least 5 leads. |

---

#### S4 — W7–8: Classification + Shadow IT Alerts

| | |
|---|---|
| **Goal** | IT admin can classify assets, receive alerts on new OAuth apps |
| **Sprint deliverable** | Shadow IT alerts firing correctly. Asset classifications visible. Flutter mobile scaffold running on iOS + Android. |
| **Scope** | Asset classification engine (criticality + data sensitivity, rule-based) · OAuth scope risk scoring (high/medium/low) · New OAuth app alert pipeline (<15 min) · Email + Slack notification system · Mobile scaffold (Flutter): auth flow Keycloak PKCE, navigation shell, push notification skeleton |
| **Key risks** | Alert noise too high → pilot users overwhelmed. Start with conservative threshold (only HIGH risk alerts). |

---

#### S5 — W9–10: Slack + AWS Discovery + RBAC

| | |
|---|---|
| **Goal** | 4 integrations (Google, M365, Slack, AWS). RBAC dashboard live. |
| **Sprint deliverable** | Unified inventory 4 providers. Least-privilege recommendations displayed. Slack deactivation tested. |
| **Scope** | Slack Admin API: users, apps, channels · Slack tier detection (Free/Pro/Business+ gating) · AWS IAM inventory: users, roles, policies · RBAC model: role assignment, permission diff engine · Least-privilege recommendations (rule-based) · Composite identity graph (cross-provider) |
| **Key risks** | Slack API tier limitation — Business+ required for automated offboarding. Detect tier early and set expectations with pilot customers. |
| **PM action** | ⚠️ **ML Engineer #1 must onboard this week (W9)**. Begin shadow AI governance R&D on synthetic data. |

---

#### S6 — W11–12: Automated Offboarding + 2 Playbooks — **🏁 MVP**

| | |
|---|---|
| **Goal** | Offboard employee in <5 min. 2 playbooks. Mobile app beta. |
| **Sprint deliverable** | **MVP: Offboarding test user <5 min via Google+M365+Slack. Mobile app on TestFlight/Play Console. PDF offboarding report.** |
| **Scope** | Automated offboarding workflow (AWS Step Functions): disable + revoke + notify · **⚠️ R-C1 (mandatory):** 30-min grace period with cancellation (configurable 0–60 min, emergency = 0) before execution — one-click cancel via Slack/email · **Rollback workflow:** Reactivate account within 24h post-offboard (manual Admin action, audit-logged) · **Idempotency key** on all offboarding requests (prevents double-execution) · Dry-run + 2-step confirmation (hard gate, no bypass) · Offboarding report PDF · 2 incident playbooks: (1) Offboarding Emergency (2) Credential Compromise · Playbook wizard UI (web) · Immutable audit log: PostgreSQL append-only + S3 (envelope-encrypted) · Mobile app v1: alerts, offboarding trigger, read-only inventory |
| **Key risks** | Sprint with highest utilization in Phase 1 (~89%). Mobile scope must be cut if needed — offboarding is absolute priority. **Grace period CANNOT be cut — this is a legal requirement.** |
| **PM action** | ✅ 3+ pilot customers must onboard on staging environment before end of W12. |

> **MVP Gate Checklist:**
> - [ ] Offboarding <5 min (timed automated test pass in CI)
> - [ ] **Grace period + Rollback workflow operational (cancel-within-30min test + rollback-within-24h test pass) (R-C1)**
> - [ ] **Idempotency key enforced — double-execution test pass (R-C1)**
> - [ ] Shadow IT alert <15 min from OAuth grant to notification
> - [ ] Tenant isolation CI test continuously green
> - [ ] 3+ pilot customers have seen "first insight" in <30 minutes setup
> - [ ] Zero plaintext secrets in environment variables
> - [ ] RDS Multi-AZ + S3 Object Lock active
> - [ ] **M365 webhook renewal job active + staleness UI indicator visible (R-C3)**
> - [ ] **Keycloak: 2 ECS tasks running (min HA), JWKS caching active (R-C6)**

---

### Phase 2: MVP → v1 (Month 4–6, S7–S13)

**Team Phase 2:** 6 FTE (Phase 1) + ML Eng #1 + BE3 + FE2 = **9 FTE** (ramping up from S7→S8)

**Phase 2 characteristics:** Track 2 (AI/ML) begins integrating with Track 1 from S7. Both tracks run in parallel, converging at S11.

---

#### S7 — W13–14: JIT Access + Track 2 Integration Begins + Vanta Setup

| | |
|---|---|
| **Goal** | JIT access end-to-end. Track 2 onboard live data. Vanta begins collecting evidence. |
| **Sprint deliverable** | JIT request → approve → auto-revoke operational. Track 2 shadow AI model receives live data from Track 1. Vanta dashboard green for configured controls. |
| **Scope — Track 1** | JIT access: request form + approval workflow + time-boxed grant + auto-revoke · Access review scheduling (periodic reminder) · Pilot feedback triage from MVP (top 10 bugs) |
| **Scope — Track 2** | BE3 onboard: environment setup, codebase walkthrough · ML Eng: shadow AI governance v1 connected to live `oauth_application` table from Track 1 · OAuth risk score model v0.2 training on live data |
| **Key risks** | JIT access approval workflow more complex than expected. Simplify: 1 approver, email-based, no self-service portal needed in v1. |
| **PM action** | ⚠️ **Sign pentest vendor LOI before end of W14 — hard deadline, no exceptions.** · ⚠️ **Vanta account provisioned W13 — evidence collection must begin immediately.** |

---

#### S8 — W15–16: Playbook Engine + 3 Playbooks + LLM DLP Prototype

| | |
|---|---|
| **Goal** | AWS Step Functions playbook engine. First 3 playbooks. Browser extension prototype. |
| **Sprint deliverable** | 3 playbooks running end-to-end on staging. LLM DLP browser extension can detect PII in text fields. |
| **Scope — Track 1** | AWS Step Functions playbook engine · Playbook wizard UI (web) · 3 playbooks: (1) Account Compromise (2) Phishing Response (3) Data Exfiltration · Playbook audit log (each step logged) |
| **Scope — Track 2** | LLM DLP browser extension v0.1 (Chrome Manifest V3): PII detection with Microsoft Presidio (local inference, no API calls) · FE2 onboard: environment setup, Chrome Extension CI/CD setup |
| **Key risks** | ⚠️ MOST LOADED sprint in entire plan (~88% utilization across both tracks). PM needs daily standups in S8. |
| **PM action** | FE2 must onboard at start of W15. If delayed → LLM DLP shifts to S9. |

---

#### S9 — W17–18: 5 Playbooks + Mobile + Shadow AI v1

| | |
|---|---|
| **Goal** | Full 5 playbooks. Mobile incident alerts. Shadow AI governance v1 in production. |
| **Sprint deliverable** | 5 playbooks complete. Push notifications from mobile for security alerts. Shadow AI risk scores live (OAuth apps classified as AI/non-AI). |
| **Scope — Track 1** | Remaining 2 playbooks: (4) Ransomware Response (5) Insider Threat Response · Mobile push notifications (FCM + APNs) · Incident alert from playbook → mobile |
| **Scope — Track 2** | Shadow AI governance v1: AI tool classification (ChatGPT, Copilot, Gemini, Claude, etc.) + risk score per OAuth app · Shadow AI attestation workflow: employee confirm/deny usage · LLM DLP extension: tenant-scoped allow-list, PII blocking before submit |
| **Key risks** | Shadow AI classification accuracy low → many false positives → pilot users complain. Start with conservative threshold (only block HIGH risk apps already confirmed as AI). |

---

#### S10 — W19–20: Compliance Mapping + T1-T2 Integration Contract

| | |
|---|---|
| **Goal** | Compliance dashboard with Vanta. T1-T2 API contract finalized. |
| **Sprint deliverable** | Compliance dashboard: coverage % ISO 27001 + SOC 2. Deepfake defense prototype. `ThreatDetectionEvent` schema v1 locked. |
| **Scope — Track 1** | ISO 27001 + SOC 2 control mapping in Vanta · Automated evidence collection hooks · Compliance dashboard (control status, evidence links) · Cross-provider composite risk score (per user, weighted) |
| **Scope — Track 2** | Deepfake defense: Hive Moderation API POC (pay-per-use, rate limit test) · Out-of-band verification workflow design · **`ThreatDetectionEvent` schema v1 finalized and locked** |
| **Tech action** | ⚠️ **`ThreatDetectionEvent` schema must be approved by end of S10.** Delay here cascades directly into S11 integration sprint. |

---

#### S11 — W21–22: Compliance Reports + T1-T2 Integration Live

| | |
|---|---|
| **Goal** | Compliance reports exportable. AI threats auto-trigger Track 1 playbooks. |
| **Sprint deliverable** | PDF compliance report (ISO 27001 + SOC 2 Type 1 evidence). Track 2 AI threat event → auto-trigger Step Functions playbook in staging. |
| **Scope — Track 1** | ISO 27001 + SOC 2 compliance reports (PDF export) · Audit trail UI · GDPR data subject request automation (export + delete) |
| **Scope — Track 2** | T1-T2 integration: `ThreatDetectionEvent` → EventBridge → Step Functions trigger · Prompt injection detection v1 (rule-based regex) · AI phishing: M365 Defender + Google Workspace threat feed connected |
| **Key risks** | ⚠️ **This is the HIGHEST technical risk sprint** — integration always takes 3x longer than estimated. Tech Lead must be full-time on this integration. Fallback: if auto-trigger is unstable → manual trigger (button) for v1, auto-trigger in v1.5. |
| **PM action** | ⚠️ **Pentest MUST START week W21** (per LOI signed in S7). Coordinate with vendor. |

---

#### S12 — W23–24: Dependency Map + Pentest Remediation + Vanta Dry Run

| | |
|---|---|
| **Goal** | Full T1-T2 integration validated. Pentest findings remediated. Vanta dry run pass >90%. |
| **Sprint deliverable** | App dependency map live (SaaS lifecycle management). Pentest findings: all Critical + High resolved. Vanta evidence dry run pass rate >90%. |
| **Scope — Track 1** | SaaS dependency mapping + lifecycle management (zombie app detection) · Pentest: remediate all Critical and High findings (Pentest runs W21–W23) · Vanta compliance evidence dry run |
| **Scope — Track 2** | Full T1-T2 end-to-end integration test (automated) · Shadow AI governance: policy enforcement mode (block vs alert) |
| **Key risks** | Pentest Critical finding in infrastructure (not code) → DevSecOps needs extra time. Buffer: S12 has 20% slack to handle. |

---

#### S13 — W25–26: Hardening + v1 Launch — **🏁 v1**

| | |
|---|---|
| **Goal** | Launch v1 to production. SOC 2 Type 1 audit scheduled. |
| **Sprint deliverable** | **v1 LIVE on production. 5+ pilot customers migrated to production. SOC 2 Type 1 audit engagement signed.** |
| **Scope** | NO new features · Performance hardening · Pentest Medium findings remediation · Launch runbook · Production cutover · SOC 2 Type 1 readiness review with Vanta auditor · Marketing launch brief |
| **Target utilization** | 60% — deliberate buffer sprint |

> **v1 Gate Checklist:**
> - [ ] All 5 incident playbooks in production
> - [ ] JIT access + offboarding <5 min (CI test)
> - [ ] 4 integrations (Google, M365, Slack, AWS)
> - [ ] Compliance dashboard: ISO 27001 + SOC 2 Type 1 report-ready
> - [ ] AI threat module: shadow AI governance + LLM DLP extension
> - [ ] Mobile app: iOS App Store + Google Play (submit S12, ~1 week review)
> - [ ] Pentest: zero Critical/High findings open
> - [ ] 5+ pilot customers on production
> - [ ] SOC 2 Type 1 audit scheduled with auditor
> - [ ] CloudWatch monitoring + PagerDuty alerting live
> - [ ] Disaster recovery runbook tested (RTO <4h)

---

### Phase 3: v1 → v1.5 (Month 7–9, S14–S20)

**Team Phase 3:** 9 FTE (Phase 2) + Customer Success Eng (M7) + ML Eng #2 (M8) + DevSecOps → FTE = **11 FTE**

**Phase 3 characteristics:** Team splits into 2 parallel streams. Stream A (planned roadmap development). Stream B (pilot feedback updates).

#### Phase 3 Team Split — Two Parallel Streams

| Stream | Split | Members | Focus |
|-------|-------|------------|-------|
| **Stream A — New Features** | **65%** | Tech Lead · BE1 · BE2 · FE1 · ML Eng #1 · ML Eng #2 | AWS v1.1, advanced AI detection, SOC 2 Type 2 prep, Business tier |
| **Stream B — Pilot Feedback** | **35%** | BE3 · FE2 · Customer Success Eng · Flutter (40%) | Bug fixes, UX polish, onboarding friction, customer requests |

> **Stream split principles:**
> - Each week: PM triages feedback queue on Monday. Issues >1 sprint → v1.5 backlog. Issues <0.5 sprint → Stream B handles immediately.
> - Stream B does NOT accept new features — fix and polish only.
> - Both streams converge at v1.5 release (W38).

---

#### S14–S15 — W27–30: Post-launch Stabilization + AWS v1.1

| | |
|---|---|
| **Stream A** | AWS IAM deep integration: CloudTrail events, S3 access auditing, IAM role recommendations · LLM DLP browser extension v1 (Chrome Web Store submit) |
| **Stream B** | Top 10 customer-reported bugs · M365 OAuth wizard UX improvements · Mobile crash fixes · Alert threshold tuning (reduce noise) |
| **Deliverable** | AWS v1.1 production. Browser extension submitted to Chrome Web Store. Sprint 14 utilization 65% (recovery sprint after 6 months of peak intensity). |
| **PM action** | Customer Success Engineer onboard M7W1. Sprint 14 kickoff includes v1 retrospective. |

---

#### S16–S17 — W31–34: Advanced AI Detection v2

| | |
|---|---|
| **Stream A** | LLM data leakage detection v2: real-time DLP (semantic analysis, not just regex) · Deepfake fraud defense v2: Hive API live + out-of-band verification workflow (SMS + Slack) · Prompt injection hardening (expanded ruleset) · ML Eng #2 onboard W32 |
| **Stream B** | Dashboard UX redesign (based on pilot feedback) · Custom alert rules UI · API documentation · Auditor-specific compliance export templates |
| **Deliverable** | AI detection accuracy >90% on test set. Customer-configurable alert rules. v2 UX pilot tested. |
| **Risk** | Browser extension rejected by Chrome Web Store (2–4 week review) → submit W29 (before S15), has buffer. |

---

#### S18–S19 — W35–38: Business Tier + SOC 2 Type 2 Prep — **🏁 v1.5**

| | |
|---|---|
| **Stream A** | Pricing tier enforcement (Starter/Growth/Business gates) · Vanta SOC 2 Type 2 evidence framework setup · Advanced compliance reporting · Custom playbook builder (Stream A) · ISO 27001 evidence continuation |
| **Stream B** | Pilot → paid customer conversion flow · Billing integration (Stripe) · Customer portal · Custom playbook builder (UX, coordinated with Stream A) |
| **Deliverable** | **v1.5 LAUNCH (W38).** Pricing tiers enforced. Billing live. 10+ paying customers. SOC 2 Type 2 evidence collection running continuously since W26. |

> **v1.5 Gate Checklist:**
> - [ ] AWS v1.1 production (CloudTrail, IAM deep)
> - [ ] Browser extension: Chrome Web Store published (not sideloaded)
> - [ ] AI detection accuracy >90% (deepfake + LLM DLP)
> - [ ] Prompt injection detection v1 (rule-based) production
> - [ ] Pricing tiers enforced (Starter / Growth / Business)
> - [ ] Billing integration live (Stripe)
> - [ ] 10+ paying customers on production
> - [ ] SOC 2 Type 2 evidence collection running since W26 (>12 weeks of evidence)
> - [ ] Custom playbook builder beta
> - [ ] SageMaker model monitoring (drift detection) active

---

### Phase 4: v1.5 → v2 (Month 10–12, S21–S26)

**Team Phase 4:** 11 FTE + Compliance Consultant (contract T10–T12) = **11.5 FTE (peak)**

**Phase 4 characteristics:** Feature freeze Month 10. Focus: SOC 2 Type 2 audit, ISO 27001 certification, Enterprise tier, BERT ML production.

---

#### S21–S22 — W39–44: Enterprise Features + SOC 2 Type 2 Audit Prep

| | |
|---|---|
| **Scope** | Enterprise tier features: multi-tenant enterprise, custom RBAC policies, SIEM integration (Splunk/QRadar webhooks) · Vanta SOC 2 Type 2 evidence final packaging · SOC 2 Type 2 audit engagement signed |
| **Deliverable** | Enterprise tier code-complete. SOC 2 Type 2 audit engagement signed with auditor. Evidence coverage >95% in Vanta. |
| **Timeline SOC 2 Type 2** | Evidence collection started W26 → audit window W26–W52 (26 weeks = 6 months ✅) · Audit fieldwork: W46–W48 · Report issued: W50–W52 |

---

#### S23–S24 — W45–48: ISO 27001 Certification + BERT Prompt Injection

| | |
|---|---|
| **Scope** | ISO 27001 Stage 2 audit prep + Statement of Applicability finalized · BERT prompt injection classifier: fine-tuned on 6 months of production data (Enterprise tier only) · Advanced analytics dashboard (SOC-level insights) · Peer group anomaly detection v1 (insider threat signal) |
| **Deliverable** | ISO 27001 Stage 2 audit complete. BERT model: FPR <2%, TPR >85% on 30-day holdout set. |
| **Risk** | BERT FPR too high → ship rule-based prompt injection (already in place) + BERT as opt-in preview, not GA. |

---

#### S25–S26 — W49–52: v2 Launch + Compliance Certified — **🏁 v2**

| | |
|---|---|
| **Scope** | Compliance certification received · Enterprise tier GA · White-label / MSSP foundation · Usage-based billing option · Multi-region DR test (not just documented) · All Track 2 features graduate from beta (SLA applies) |
| **Deliverable** | **v2 LAUNCH (W52).** SOC 2 Type 2 certified. ISO 27001 certified. Enterprise tier live. |

> **v2 Gate Checklist:**
> - [ ] SOC 2 Type 2 report received from auditor
> - [ ] ISO 27001 certificate received
> - [ ] BERT prompt injection: FPR <2%, TPR >85% (or opt-in preview if not yet achieved)
> - [ ] Enterprise tier: custom pricing, dedicated CSM, SIEM integration
> - [ ] Advanced analytics dashboard production
> - [ ] Peer group anomaly detection production
> - [ ] All Track 2 features: beta flag removed, SLA guarantees
> - [ ] Multi-region DR failover drill tested (RTO/RPO documented)
> - [ ] 99.95% uptime SLA target achievable (verified from monitoring data)

---

## 5. Two-Stream Team Split (Post-v1)

### Detailed Two-Stream Structure (Phase 3 & 4)

```
STREAM A — New Features (65%)          STREAM B — Pilot Feedback (35%)
────────────────────────────           ────────────────────────────────
Tech Lead                              Customer Success Engineer
Backend Eng #1                         Backend Eng #3
Backend Eng #2                         Frontend Eng #2
Frontend Eng #1                        Flutter Eng (40% time)
ML Eng #1
ML Eng #2

Responsibilities:                      Responsibilities:
  - Pre-planned roadmap features         - Bug queue from pilot/customers
  - v1.5 capabilities                    - UX friction from usage analytics
  - SOC 2 Type 2 evidence prep           - Feature requests <2 days work
  - Enterprise tier                      - Onboarding wizard improvements
  - AI accuracy improvements             - Alert noise tuning
  - New integrations                     - Performance issues

Cadence:                               Cadence:
  Standard 2-week sprints                  Weekly triage (Monday)
  Sprint planning Monday start of sprint   Continuous deployment
  Sprint demo                              SLA: P1 fix <24h, P2 <5 days
```

### Two-Stream Coordination Rules

| Rule | Description |
|---------|-------|
| **Weekly triage** | PM triages feedback queue every Monday. Assigns to Stream B or v1.5 backlog |
| **Escalation gate** | Issue >3 days estimate → move to backlog, do not disrupt Stream B sprint |
| **Feature creep guard** | Stream B does NOT accept new feature requests from customers. Fix and polish only. |
| **Convergence** | Both streams merge code to main daily (feature flags for non-GA features) |
| **Demo chung** | Sprint demo includes both streams. Customers are invited to see Stream B fixes. |

---

## 6. Key Requirements Coverage

| Key Requirement | Milestone | Sprint | Notes |
|---|---|---|---|
| **Asset inventory & classification** | v1 (T6) | S2–S4 core, S12 full | Google+M365 from MVP. Slack+AWS at S5. Shadow AI at S9 (Track 2). |
| **AI-specific threat surface** | v1 (T6) | S7–S11 (Track 2) | Shadow AI governance S9. LLM DLP extension S8–S9. Deepfake + prompt injection S11. Full AI detection package present in v1. |
| **Access governance** | v1 (T6) | S5–S7 core | RBAC S5. Offboarding S6 (MVP). JIT S7. Access reviews S7. Shadow IT remediation S9. |
| **Compliance posture** | v1 (T6) — report-ready | S10–S11 | SOC 2 Type 1 + ISO 27001 report exportable from v1. Certification (audit verified) at v2. |
| **Incident playbooks** | v1 (T6) | S6 (2 playbooks), S8–S9 (5 playbooks) | 5 playbooks, wizard UI, AWS Step Functions, non-security staff operable. |
| **Cost model** | v1.5 (T9) billing live | S13 code-ready, S18–S19 billing | Pricing tiers code-complete at v1. Billing Stripe integration at v1.5. Manual invoicing for pilot months 1–6. |
| **Integrations** (Google, M365, Slack, QuickBooks...) | v1 (T6) | S2–S5 | Google+M365 from MVP. Slack S5. AWS S5. QuickBooks → v2 backlog (insufficient AI security value for v1). |

> **Conclusion:** All 7 key requirements from the brief **are present in v1 (month 6)**, as required by "v1 after 5-6 months".

---

## 7. Riskiest Assumption to Validate First

### Risk #1 (Critical): Pilot Customers Cannot Onboard in <30 Minutes

> **Assumption:** SME IT admin (not a developer) can set up Google Workspace + M365 OAuth in 30 minutes using a guided wizard.

**Why this is the highest-risk assumption:**
- The entire MVP value prop depends on "first-value <30 min"
- If onboarding actually takes 3 hours (due to M365 permission complexity), the pilot program collapses
- Competitors take 2–4 hours — if SMESec is the same, there is no differentiation

**When to validate:** Sprint 2 end (W4) — test with 1–2 non-technical users on a real Google Workspace tenant
**How to validate:** Time-boxed usability test, no assistance from engineer
**Go/No-go:** If >45 min → redesign wizard before continuing S3

### Top 5 Risks by Phase

| # | Risk | Phase | Probability | Impact | Mitigation |
|---|--------|-----|----------|----------|------------|
| 1 | OAuth wizard >30 min for non-technical IT admin (M365) | MVP | High | Critical | Usability test W4. Prepared IT admin guide. Minimum-permission scopes. |
| 2 | ML Engineer not hired before W9 | Phase 2 | Medium | High | Begin recruiting W5. Contractor ML fallback if hire is delayed. Tech Lead builds SageMaker scaffold S5. |
| 3 | Pentest vendor LOI not signed before W14 | Phase 2 | Low | High | PM locks calendar from W8. Backup vendor list. |
| 4 | T1–T2 integration at S11 delayed >1 sprint | Phase 2 | High | High | Tech Lead full-time S11. API contract frozen S10. Fallback: manual trigger for v1. |
| 5 | SOC 2 Type 2 evidence gap at Month 9 review | Phase 3 | Low | High | Vanta weekly review from W13. PM owns Vanta. Zero gap policy from W22. |

---

## 8. Compliance Certification Timeline

```
Month 3 (W12):  Vanta account setup, evidence collection begins (silent)
Month 4 (W13):  Vanta OFFICIALLY active — SOC 2 control mapping begins
Month 5 (W21):  Pentest begins (6-month lead time from LOI signing W14)
Month 6 (W26):  v1 LAUNCH
                  → SOC 2 Type 1 audit: scheduled with auditor
                  → Evidence collection W13→W26 = ~13 weeks (sufficient for Type 1)
Month 7 (W27):  ISO 27001 gap analysis begins
Month 8 (W33):  ISO 27001 Stage 1 audit (documentation review)
Month 9 (W38):  v1.5 LAUNCH
                  → SOC 2 Type 2 evidence W26→W38 = 12 weeks (need 24 weeks total)
Month 10 (W41): ISO 27001 Stage 2 audit (implementation review)
Month 11 (W46): SOC 2 Type 2 audit fieldwork begins
                  → Evidence W26→W46 = 20 weeks (need 24 weeks — ⚠️ tight)
                  → Safer milestone: start audit W48
Month 12 (W52): v2 LAUNCH
                  → SOC 2 Type 2 report issued ✅
                  → ISO 27001 certificate issued ✅
```

> ⚠️ **SOC 2 Type 2 timing note:** To have a full 6-month (24-week) observation window before W52, evidence collection MUST begin no later than W26. Starting from W13 as planned gives a 10-week buffer, but the official SOC 2 Type 2 window counts from W26 (v1 launch date — production environment).

---

## 9. External Dependencies & Hard Deadlines

| Deadline | Week | Description | Consequence if delayed |
|----------|------|--------|----------------|
| Auth provider decision | W1D1 | Choose Keycloak self-host vs Auth0 vs Cognito | Delay S1 → cascade all sprints |
| Google test tenant available | W3 | Internal Google Workspace tenant for S2 development | S2 cannot demo |
| Pilot customer #1 onboard | W8 | At least 1 real customer using staging | MVP has no real validation |
| **ML Engineer #1 onboard** | **W9** | Begin recruiting W5 | Track 2 delay → AI features miss v1 |
| **Pentest vendor LOI signed** | **W14** | Hard deadline — 7-week lead time | Pentest does not start W21 → v1 delay |
| **Vanta setup active** | **W13** | Need 60+ days evidence for SOC 2 Type 1 | SOC 2 Type 1 insufficient evidence at v1 |
| Chrome Web Store submission | W29 | Browser extension needs 1–2 weeks review | Extension misses v1.5 launch |
| iOS App Store submission | W50 | App Store review 1–2 weeks | Mobile feature misses v2 window |
| SOC 2 Type 2 audit sign | W42 | Engage auditor firm | Audit does not complete before W52 |
| ISO 27001 Stage 2 audit | W45 | Certification 6–8 weeks after audit | Certificate not available at W52 |

---

## Summary Dashboard

```
MILESTONE OVERVIEW
══════════════════════════════════════════════════════════════════════

MVP    │ W12  │ T3  │ Asset inventory + offboarding + shadow IT
       │      │     │ Team: 6 FTE + contract
       │      │     │ Pilot: 3+ customers on staging (free)
───────┼──────┼─────┼──────────────────────────────────────────────
v1     │ W26  │ T6  │ All 7 key requirements delivered
       │      │     │ Team: 9 FTE
       │      │     │ Compliance: SOC 2 Type 1 audit scheduled
       │      │     │ Revenue: Starter ($399/mo) + Growth ($799/mo) tiers
───────┼──────┼─────┼──────────────────────────────────────────────
v1.5   │ W38  │ T9  │ AI detection v2 + AWS v1.1 + pilot feedback
       │      │     │ Team: 11 FTE (2-stream split: 65/35)
       │      │     │ Compliance: SOC 2 Type 2 evidence running
       │      │     │ Revenue: Business tier ($1,499/mo) live. Billing Stripe.
───────┼──────┼─────┼──────────────────────────────────────────────
v2     │ W52  │ T12 │ Compliance verified + Enterprise + BERT ML
       │      │     │ Team: 11.5 FTE (peak)
       │      │     │ Compliance: SOC 2 Type 2 ✅ + ISO 27001 ✅
       │      │     │ Revenue: Enterprise (custom) + usage-based option

══════════════════════════════════════════════════════════════════════
TEAM HEADCOUNT PROGRESSION
  T1–T3:  ██████░░░░░░  6 FTE
  T4–T6:  █████████░░░  9 FTE
  T7–T9:  ███████████░  11 FTE
  T10–T12:████████████  11.5 FTE (peak)
══════════════════════════════════════════════════════════════════════
```
