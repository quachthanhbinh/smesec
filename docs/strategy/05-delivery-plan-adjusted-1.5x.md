# SMESec Platform — Delivery Plan (Adjusted 1.5x Timeline)

**Date:** 2026-05-29  
**Status:** Revised — Timeline extended 1.5x for realistic execution  
**Version:** 2.0  
**Scope:** Full roadmap from Sprint 1 to v2 (19.5 months)  
**Previous Version:** [04-delivery-plan-original.md](04-delivery-plan-original.md) (12 months, aggressive timeline)

---

## Executive Summary

Bản kế hoạch này điều chỉnh timeline gốc (12 tháng) ra **1.5x = 19.5 tháng** để giảm rủi ro và tăng tính khả thi. Các thay đổi chính:

- **MVP**: W12 → **W18** (Month 3 → Month 4.5)
- **v1**: W26 → **W39** (Month 6 → Month 9.75)
- **v1.5**: W38 → **W57** (Month 9 → Month 14.25)
- **v2**: W52 → **W78** (Month 12 → Month 19.5)

**Lý do điều chỉnh:**
- Giảm sprint overload (nhiều sprint trong plan gốc có utilization >85%)
- Thêm buffer cho tuyển dụng và onboarding
- Thêm thời gian cho pentest remediation và compliance audit
- Realistic về Google/M365 API rate limit handling
- Buffer cho Chrome Web Store review delays

---

## Table of Contents

1. [Roadmap Overview](#1-roadmap-overview)
2. [Scope by Milestone](#2-scope-by-milestone)
3. [Team & Headcount Ramp](#3-team--headcount-ramp)
4. [Sprint Breakdown](#4-sprint-breakdown)
5. [Key Requirements Coverage](#5-key-requirements-coverage)
6. [Compliance Certification Timeline](#6-compliance-certification-timeline)
7. [External Dependencies & Hard Deadlines](#7-external-dependencies--hard-deadlines)
8. [Comparison with Original Plan](#8-comparison-with-original-plan)

---

## 1. Roadmap Overview

```
Month:  1    2    3    4    5    6    7    8    9   10   11   12   13   14   15   16   17   18   19   20
Sprint: S1  S2  S3  S4  S5  S6  S7  S8  S9  S10 S11 S12 S13 S14 S15 S16 S17 S18 S19 S20 S21 S22 S23 S24 S25 S26 S27 S28 S29 S30 S31 S32 S33 S34 S35 S36 S37 S38 S39
        |----------Phase 1: Foundation----------|--------Phase 2: MVP→v1---------|----------Phase 3: v1→v1.5----------|----------Phase 4: v1.5→v2----------|
                                    ↑                                           ↑                                   ↑                                       ↑
                                   MVP                                          v1                                 v1.5                                    v2
                                (W18/M4.5)                                  (W39/M9.75)                         (W57/M14.25)                           (W78/M19.5)
```

| Milestone | Week | Month | Description | Change from Original |
|-----------|------|-------|-------------|---------------------|
| **MVP** | W18 | M4.5 | Asset inventory + automated offboarding + shadow IT | +6 weeks (was W12) |
| **v1** | W39 | M9.75 | All key requirements delivered. SOC 2 Type 1 audit scheduled. | +13 weeks (was W26) |
| **v1.5** | W57 | M14.25 | Advanced AI detection + AWS v1.1 + pilot feedback. Billing live. | +19 weeks (was W38) |
| **v2** | W78 | M19.5 | SOC 2 Type 2 + ISO 27001 certified. Enterprise tier. | +26 weeks (was W52) |

---

## 2. Scope by Milestone

### Feature Map by Phase (Unchanged from original)

| Feature Domain | MVP (M4.5) | v1 (M9.75) | v1.5 (M14.25) | v2 (M19.5) |
|---|---|---|---|---|
| **Asset Inventory** | Google WS + M365: users, OAuth apps, basic devices | + Slack, AWS, Shadow AI detection | + Custom asset types, dependency map | + Full cloud posture, peer anomaly |
| **Access Governance** | Automated offboarding <5 min, RBAC dashboard | + JIT access, access reviews, shadow IT remediation | + Risk scoring, access policy templates | + Peer group anomaly, insider threat signal |
| **AI Threat Surface** | ❌ (Track 2 in R&D) | Shadow AI governance + LLM DLP browser ext (beta) | + Deepfake defense, AI phishing, prompt injection v1 | + Prompt injection ML (BERT), advanced analytics |
| **Compliance Posture** | Immutable audit log active (S3 Object Lock + per-tenant KMS) | Vanta evidence collection live · SOC 2 Type 1 + ISO 27001 report-ready | SOC 2 Type 2 evidence running (6 months) | SOC 2 Type 2 certified + ISO 27001 certified |
| **Incident Playbooks** | 2 playbooks (Offboarding, Cred Compromise) | 5 playbooks, wizard UI, AWS Step Functions | + Custom playbook builder, mobile triggers | + Playbook analytics, ML suggestions |
| **Integrations** | Google WS + M365 (OAuth wizard <30 min) | + Slack full + AWS IAM basic | + AWS CloudTrail, S3 audit, IAM deep | + SIEM (Splunk/QRadar), custom webhooks |
| **Mobile App** | ❌ TestFlight/Beta | Alerts + playbook trigger (iOS + Android) | Full incident response mobile | Full feature parity |
| **Billing / Pricing** | Manual invoicing (pilot free) | Starter + Growth tiers code-ready | Pricing tiers enforced, billing live | Enterprise custom + usage-based |

---

## 3. Team & Headcount Ramp

### Headcount Timeline (Adjusted)

```
Month 1 (Sprint 1):                ML Engineer #1 onboards Day 1 — Track 2 R&D starts in parallel
Month 1–4.5 (Phase 1 / MVP):       7 FTE core + DevSecOps contract
Month 6 (Sprint 12):               + Backend Engineer #3 (Track 2)
Month 7 (Sprint 14):               + Frontend Engineer #2 (Browser Extension)
Month 10.5 (start of Phase 3):     DevSecOps → FTE (no longer contract)
Month 10.5 (start of Phase 3):     + Customer Success Engineer
Month 12 (mid Phase 3):            + ML Engineer #2 (optional, depending on v1 velocity)
Month 15–19.5 (Phase 4):           + Compliance Consultant (contract)
```

### Detailed Headcount by Phase

| Role | M1–M4.5 | M5–M9.75 | M10–M14.25 | M15–M19.5 | Track |
|---------|-------|-------|-------|---------|-------|
| Tech Lead / Architect | ✅ 1.0 FTE | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | 1 |
| Backend Eng #1 (Go) | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | 1 |
| Backend Eng #2 (Go/Python) | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | 1 |
| Frontend Eng #1 (React) | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | 1 |
| Flutter / Mobile Eng | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | 1 |
| DevSecOps | Contract (0.5) | Contract (0.5) | **FTE (1.0)** | **FTE (1.0)** | Shared |
| PM | 0.5 | 0.5 | 0.5 | 0.5 | Shared |
| **ML Engineer #1** | **✅ 1.0 (M1, Day 1)** | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | 2 |
| **Backend Eng #3 (Python/FastAPI)** | — | **✅ 1.0 (M6)** | ✅ 1.0 | ✅ 1.0 | 2 |
| **Frontend Eng #2 (Browser Ext)** | — | **✅ 1.0 (M7)** | ✅ 1.0 | ✅ 1.0 | 2 |
| **Customer Success Engineer** | — | — | **✅ 1.0 (M10.5)** | ✅ 1.0 | Customer |
| **ML Engineer #2** | — | — | **✅ 1.0 (M12, opt.)** | ✅ 1.0 | 2 |
| **Compliance Consultant** | — | — | — | **Contract (M15–M19.5)** | Compliance |
| **Total FTE** | **7** | **9 → 9.5** | **10 → 11** | **11.5** | |

---

## 4. Sprint Breakdown

### Phase 1: Foundation → MVP (Month 1–4.5, S1–S9)

**Duration:** 18 weeks (vs 12 weeks original) — **+50% time**

**Team Phase 1:** Tech Lead · BE1 · BE2 · FE1 · Flutter · ML Eng #1 · DevSecOps(contract) · PM = **7 FTE**

#### Key Changes from Original:
- Split S1 into S1a + S1b (infrastructure vs auth) — **now 4 weeks instead of 2**
- **S1a expanded for 1K tenant target:** GCP project pool (50 projects), RDS Proxy, Redis r6g.large, bounded sync worker pool, and batch secrets schema are all Sprint 1 mandatory (not v2 concerns)
- Add buffer sprint S7 for integration stabilization
- Add buffer sprint S8 for pilot feedback
- MVP moved from S6 (W12) to S9 (W18)

#### Sprint Summary:

| Sprint | Week | Focus | Key Deliverables |
|--------|------|-------|------------------|
| **S1a** | W1–2 | Infrastructure Foundation | AWS VPC, ECS, RDS + **RDS Proxy**, Redis **r6g.large**, CI/CD, M365 webhook schema (`subscription_registry` + `renewal_bucket`), **50 GCP projects provisioned**, `gcp_project_id`+`shard_id` in `tenant_config` schema, **batch secrets schema** (1 JSON/tenant), **bounded sync worker pool scaffold** (200-worker pool) |
| **S1b** | W3–4 | Auth + Security | Keycloak HA (4 ECS tasks), JWT JWKS caching, RLS policies on all tables, WAF, GCP project assignment logic in SyncScheduler |
| **S2** | W5–6 | Google Workspace Sync | User/OAuth sync, dashboard skeleton, shadow IT detection v1 |
| **S3** | W7–8 | M365 Sync + Dashboard | M365 delta link, webhook renewal, unified dashboard |
| **S4** | W9–10 | Classification + Alerts | Asset classification, OAuth risk scoring, mobile scaffold |
| **S5** | W11–12 | Slack + AWS + RBAC | 4 providers unified, RBAC model, least-privilege |
| **S6** | W13–14 | Offboarding + Playbooks | Automated offboarding <5 min, 2 playbooks, audit log |
| **S7** | W15–16 | **Integration Stabilization** | Fix integration bugs, rate limit testing, mobile beta polish |
| **S8** | W17–18 | **Pilot Feedback Buffer** | Top 10 pilot issues, UX improvements, performance tuning |
| **S9** | W18 | **🏁 MVP Launch** | 3+ pilot customers onboarded, offboarding verified <5 min |

> **MVP = W18** (vs W12 original) — **+6 weeks buffer**

---

### Phase 2: MVP → v1 (Month 5–9.75, S10–S20)

**Duration:** 22 weeks (vs 14 weeks original) — **+57% time**

**Team Phase 2:** 7 FTE (Phase 1) + BE3 (M6) + FE2 (M7) = **9 FTE**

#### Key Changes from Original:
- Extended from 7 sprints (S7–S13) to 11 sprints (S10–S20)
- Added 2 buffer sprints for pentest remediation
- Added 1 buffer sprint for Chrome Web Store review delays
- v1 moved from S13 (W26) to S20 (W39)

#### Sprint Summary:

| Sprint | Week | Track 1 | Track 2 | Gate |
|--------|------|---------|---------|------|
| **S10** | W19–20 | JIT Access + Vanta Setup | BE3 onboard, Shadow AI v0.2 | Vanta active W20 |
| **S11** | W21–22 | Playbook Engine + 3 Playbooks | FE2 onboard, LLM DLP v0.1 | 3 playbooks end-to-end |
| **S12** | W23–24 | 5 Playbooks + Mobile | Shadow AI v1, LLM DLP extension | Gate 3: Shadow AI >95% |
| **S13** | W25–26 | Compliance Mapping | Deepfake POC, schema v1 locked | Gate 4: Deepfake >80% |
| **S14** | W27–28 | Compliance Reports + GDPR | T1-T2 integration, Lakera Guard | **Pentest starts W27** |
| **S15** | W29–30 | **Pentest Remediation #1** | T1-T2 integration testing | Critical/High findings fixed |
| **S16** | W31–32 | **Pentest Remediation #2** | Full integration validation | 0 Critical/High open |
| **S17** | W33–34 | Dependency Map + Vanta Dry Run | Shadow AI policy enforcement | Vanta >90% pass rate |
| **S18** | W35–36 | **Chrome Extension Submit** | Extension full version to Store | **2-week review buffer** |
| **S19** | W37–38 | **Store Review Buffer** | Track 2 hardening | Extension approved or fallback |
| **S20** | W39 | **🏁 v1 Launch** | 5+ pilot customers, SOC 2 Type 1 scheduled | Production cutover |

> **v1 = W39** (vs W26 original) — **+13 weeks buffer**

---

### Phase 3: v1 → v1.5 (Month 10–14.25, S21–S29)

**Duration:** 18 weeks (vs 12 weeks original) — **+50% time**

**Team Phase 3:** 9 FTE + Customer Success Eng (M10.5) + ML Eng #2 (M12) + DevSecOps → FTE = **11 FTE**

#### Key Changes from Original:
- Extended from 7 sprints (S14–S20) to 9 sprints (S21–S29)
- Added 1 sprint for AWS deep integration
- Added 1 sprint for billing integration testing
- v1.5 moved from S19 (W38) to S29 (W57)

#### Sprint Summary:

| Sprint | Week | Stream A (65%) | Stream B (35%) |
|--------|------|----------------|----------------|
| **S21–S22** | W40–43 | Post-launch Stabilization + AWS v1.1 | Top 10 bugs, M365 wizard UX |
| **S23–S24** | W44–47 | Advanced AI Detection v2 | Dashboard UX redesign, custom alerts |
| **S25** | W48–49 | **AWS Deep Integration** | API documentation, compliance templates |
| **S26–S27** | W50–53 | Business Tier + SOC 2 Type 2 Prep | Billing integration (Stripe) |
| **S28** | W54–55 | **Billing Integration Testing** | Customer portal, conversion flow |
| **S29** | W57 | **🏁 v1.5 Launch** | Pricing tiers enforced, 10+ paying customers |

> **v1.5 = W57** (vs W38 original) — **+19 weeks buffer**

---

### Phase 4: v1.5 → v2 (Month 15–19.5, S30–S39)

**Duration:** 20 weeks (vs 12 weeks original) — **+67% time**

**Team Phase 4:** 11 FTE + Compliance Consultant (contract M15–M19.5) = **11.5 FTE (peak)**

#### Key Changes from Original:
- Extended from 6 sprints (S21–S26) to 10 sprints (S30–S39)
- Added 2 sprints for SOC 2 Type 2 audit preparation
- Added 2 sprints for ISO 27001 certification buffer
- v2 moved from S26 (W52) to S39 (W78)

#### Sprint Summary:

| Sprint | Week | Focus | Deliverable |
|--------|------|-------|-------------|
| **S30–S31** | W58–61 | Enterprise Features | Multi-tenant enterprise, custom RBAC, SIEM integration |
| **S32–S33** | W62–65 | **SOC 2 Type 2 Audit Prep** | Evidence packaging, auditor engagement signed |
| **S34–S35** | W66–69 | ISO 27001 Certification + BERT | Stage 2 audit prep, BERT prompt injection (Enterprise) |
| **S36–S37** | W70–73 | **Compliance Audit Buffer** | Audit fieldwork, findings remediation |
| **S38** | W74–75 | **Certification Finalization** | SOC 2 Type 2 report, ISO 27001 certificate |
| **S39** | W78 | **🏁 v2 Launch** | Compliance certified, Enterprise tier GA |

> **v2 = W78** (vs W52 original) — **+26 weeks buffer**

---

## 5. Key Requirements Coverage

| Key Requirement | Milestone | Sprint | Timeline Change |
|---|---|---|---|
| **Asset inventory & classification** | v1 (M9.75) | S2–S4 core, S17 full | +13 weeks |
| **AI-specific threat surface** | v1 (M9.75) | S10–S14 (Track 2) | +13 weeks |
| **Access governance** | v1 (M9.75) | S5–S10 core | +13 weeks |
| **Compliance posture** | v1 (M9.75) — report-ready | S13–S14 | +13 weeks |
| **Incident playbooks** | v1 (M9.75) | S6 (2), S11–S12 (5) | +13 weeks |
| **Cost model** | v1.5 (M14.25) billing live | S20 code-ready, S26–S28 billing | +19 weeks |
| **Integrations** | v1 (M9.75) | S2–S5 | +13 weeks |

> **Conclusion:** All 7 key requirements delivered at v1 (Month 9.75), với timeline realistic hơn 50–67%.

---

## 6. Compliance Certification Timeline

```
Month 4.5 (W18):  MVP LAUNCH — Vanta account setup, evidence collection begins
Month 5 (W20):    Vanta OFFICIALLY active — SOC 2 control mapping begins
Month 7 (W27):    Pentest begins (LOI signed W21)
Month 9.75 (W39): v1 LAUNCH
                  → SOC 2 Type 1 audit: scheduled with auditor
                  → Evidence collection W20→W39 = ~19 weeks (sufficient for Type 1)
Month 10.5 (W42): ISO 27001 gap analysis begins
Month 12 (W48):   ISO 27001 Stage 1 audit (documentation review)
Month 14.25 (W57): v1.5 LAUNCH
                  → SOC 2 Type 2 evidence W39→W57 = 18 weeks (need 24 weeks total)
Month 15 (W60):   ISO 27001 Stage 2 audit (implementation review)
Month 17 (W68):   SOC 2 Type 2 audit fieldwork begins
                  → Evidence W39→W68 = 29 weeks (exceeds 24-week minimum ✅)
Month 19.5 (W78): v2 LAUNCH
                  → SOC 2 Type 2 report issued ✅
                  → ISO 27001 certificate issued ✅
```

> ✅ **SOC 2 Type 2 timing:** Evidence window W39→W68 = **29 weeks** (exceeds 24-week minimum requirement with 5-week buffer)

---

## 7. External Dependencies & Hard Deadlines

| Deadline | Week | Description | Buffer vs Original |
|----------|------|-------------|-------------------|
| Auth provider decision | W1D1 | Choose Keycloak self-host vs Auth0 vs Cognito | No change |
| **ML Engineer #1 onboard** | **W1D1** | Must be hired before project kick-off | No change |
| Google test tenant available | W5 | Internal Google Workspace tenant for S2 development | +2 weeks |
| Pilot customer #1 onboard | W12 | At least 1 real customer using staging | +4 weeks |
| **Pentest vendor LOI signed** | **W21** | Hard deadline — 7-week lead time | +7 weeks |
| **Vanta setup active** | **W20** | Need 60+ days evidence for SOC 2 Type 1 | +7 weeks |
| Chrome Web Store submission | W35 | Browser extension needs 1–2 weeks review | +6 weeks |
| iOS App Store submission | W74 | App Store review 1–2 weeks | +24 weeks |
| SOC 2 Type 2 audit sign | W62 | Engage auditor firm | +20 weeks |
| ISO 27001 Stage 2 audit | W60 | Certification 6–8 weeks after audit | +15 weeks |

---

## 8. Comparison with Original Plan

### Timeline Comparison

| Milestone | Original | Adjusted 1.5x | Delta | % Increase |
|-----------|----------|---------------|-------|------------|
| **MVP** | W12 (M3) | W18 (M4.5) | +6 weeks | +50% |
| **v1** | W26 (M6) | W39 (M9.75) | +13 weeks | +50% |
| **v1.5** | W38 (M9) | W57 (M14.25) | +19 weeks | +50% |
| **v2** | W52 (M12) | W78 (M19.5) | +26 weeks | +50% |

### Sprint Count Comparison

| Phase | Original Sprints | Adjusted Sprints | Delta |
|-------|------------------|------------------|-------|
| **Phase 1 (MVP)** | 6 sprints | 9 sprints | +3 sprints |
| **Phase 2 (v1)** | 7 sprints | 11 sprints | +4 sprints |
| **Phase 3 (v1.5)** | 7 sprints | 9 sprints | +2 sprints |
| **Phase 4 (v2)** | 6 sprints | 10 sprints | +4 sprints |
| **Total** | **26 sprints** | **39 sprints** | **+13 sprints** |

### Key Benefits of 1.5x Timeline

✅ **Reduced Sprint Overload:**
- Original plan had multiple sprints with >85% utilization
- Adjusted plan targets 60–75% utilization for sustainable pace

✅ **Realistic Hiring & Onboarding:**
- Added buffer for ML Engineer #1 recruitment (if delayed)
- Added onboarding time for BE3 and FE2 (M6–M7 vs M4–M4.5)

✅ **Pentest & Compliance Buffer:**
- Original: 5 weeks pentest → remediation → launch
- Adjusted: 13 weeks pentest → remediation → retest → launch

✅ **Chrome Web Store Review Buffer:**
- Original: Submit W29, need approval W26 (impossible)
- Adjusted: Submit W35, 2-week buffer before v1 (W39)

✅ **SOC 2 Type 2 Evidence Window:**
- Original: W26→W52 = 26 weeks (barely meets 24-week minimum)
- Adjusted: W39→W68 = 29 weeks (5-week buffer above minimum)

---

## Summary Dashboard

```
MILESTONE OVERVIEW (ADJUSTED 1.5x)
══════════════════════════════════════════════════════════════════════

MVP    │ W18  │ M4.5  │ Asset inventory + offboarding + shadow IT
       │      │       │ Team: 7 FTE + contract (ML Eng #1 from Day 1)
       │      │       │ Buffer: +6 weeks vs original plan
───────┼──────┼───────┼──────────────────────────────────────────────
v1     │ W39  │ M9.75 │ All 7 key requirements delivered
       │      │       │ Team: 9 FTE
       │      │       │ Compliance: SOC 2 Type 1 audit scheduled
       │      │       │ Buffer: +13 weeks vs original plan
───────┼──────┼───────┼──────────────────────────────────────────────
v1.5   │ W57  │ M14.25│ AI detection v2 + AWS v1.1 + pilot feedback
       │      │       │ Team: 11 FTE (2-stream split: 65/35)
       │      │       │ Billing: Stripe live, 10+ paying customers
       │      │       │ Buffer: +19 weeks vs original plan
───────┼──────┼───────┼──────────────────────────────────────────────
v2     │ W78  │ M19.5 │ Compliance verified + Enterprise + BERT ML
       │      │       │ Team: 11.5 FTE (peak)
       │      │       │ Compliance: SOC 2 Type 2 ✅ + ISO 27001 ✅
       │      │       │ Buffer: +26 weeks vs original plan

══════════════════════════════════════════════════════════════════════
TOTAL PROJECT DURATION
  Original:  52 weeks (12 months)
  Adjusted:  78 weeks (19.5 months)
  Increase:  +50% timeline, significantly reduced risk
══════════════════════════════════════════════════════════════════════
```
