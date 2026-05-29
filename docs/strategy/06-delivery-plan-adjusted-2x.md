# SMESec Platform — Delivery Plan (Adjusted 2x Timeline)

**Date:** 2026-05-29  
**Status:** Revised — Timeline extended 2x for sustainable execution  
**Version:** 3.0  
**Scope:** Full roadmap from Sprint 1 to v2 (24 months)  
**Previous Version:** [05-delivery-plan-adjusted-1.5x.md](05-delivery-plan-adjusted-1.5x.md) (19.5 months)

---

## Executive Summary

Bản kế hoạch này điều chỉnh timeline gốc (12 tháng) ra **2x = 24 tháng** để đảm bảo tính bền vững và giảm thiểu rủi ro burnout. Đây là timeline thực tế cho team nhỏ với workload hợp lý.

**Thay đổi chính so với plan 1.5x:**
- **MVP**: W18 → **W24** (Month 4.5 → Month 6)
- **v1**: W39 → **W52** (Month 9.75 → Month 13)
- **v1.5**: W57 → **W76** (Month 14.25 → Month 19)
- **v2**: W78 → **W104** (Month 19.5 → Month 26)

**Lý do điều chỉnh thêm:**
- Sprint utilization trong plan 1.5x vẫn còn 70-80% (vẫn cao)
- Cần thêm buffer cho integration testing và bug fixing
- Thêm thời gian cho team onboarding và knowledge transfer
- Realistic về compliance audit timeline (SOC 2 Type 2 cần 6+ tháng observation)
- Buffer cho external dependencies (Google/M365 API verification, pentest scheduling)

---

## 1. Roadmap Overview

```
Month:  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26
Sprint: S1 S2 S3 S4 S5 S6 S7 S8 S9 S10S11S12S13S14S15S16S17S18S19S20S21S22S23S24S25S26S27...S52
        |------------Phase 1: Foundation------------|----------Phase 2: MVP→v1----------|----------Phase 3: v1→v1.5----------|----------Phase 4: v1.5→v2----------|
                                        ↑                                               ↑                                   ↑                                       ↑
                                       MVP                                              v1                                 v1.5                                    v2
                                    (W24/M6)                                        (W52/M13)                           (W76/M19)                              (W104/M26)
```

| Milestone | Week | Month | Description | Change from 1.5x |
|-----------|------|-------|-------------|------------------|
| **MVP** | W24 | M6 | Asset inventory + automated offboarding + shadow IT | +6 weeks (was W18) |
| **v1** | W52 | M13 | All key requirements delivered. SOC 2 Type 1 audit scheduled. | +13 weeks (was W39) |
| **v1.5** | W76 | M19 | Advanced AI detection + AWS v1.1 + pilot feedback. Billing live. | +19 weeks (was W57) |
| **v2** | W104 | M26 | SOC 2 Type 2 + ISO 27001 certified. Enterprise tier. | +26 weeks (was W78) |

---

## 2. Team & Headcount Ramp

### Headcount Timeline (Adjusted 2x)

```
Month 1 (Sprint 1):                ML Engineer #1 onboards Day 1 — Track 2 R&D starts in parallel
Month 1–6 (Phase 1 / MVP):         7 FTE core + DevSecOps contract
Month 8 (Sprint 16):               + Backend Engineer #3 (Track 2)
Month 9 (Sprint 18):               + Frontend Engineer #2 (Browser Extension)
Month 14 (start of Phase 3):       DevSecOps → FTE (no longer contract)
Month 14 (start of Phase 3):       + Customer Success Engineer
Month 16 (mid Phase 3):            + ML Engineer #2 (optional, depending on v1 velocity)
Month 20–26 (Phase 4):             + Compliance Consultant (contract)
```

### Detailed Headcount by Phase

| Role | M1–M6 | M7–M13 | M14–M19 | M20–M26 | Track |
|---------|-------|-------|-------|---------|-------|
| Tech Lead / Architect | ✅ 1.0 FTE | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | 1 |
| Backend Eng #1 (Go) | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | 1 |
| Backend Eng #2 (Go/Python) | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | 1 |
| Frontend Eng #1 (React) | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | 1 |
| Flutter / Mobile Eng | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | 1 |
| DevSecOps | Contract (0.5) | Contract (0.5) | **FTE (1.0)** | **FTE (1.0)** | Shared |
| PM | 0.5 | 0.5 | 0.5 | 0.5 | Shared |
| **ML Engineer #1** | **✅ 1.0 (M1, Day 1)** | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | 2 |
| **Backend Eng #3 (Python/FastAPI)** | — | **✅ 1.0 (M8)** | ✅ 1.0 | ✅ 1.0 | 2 |
| **Frontend Eng #2 (Browser Ext)** | — | **✅ 1.0 (M9)** | ✅ 1.0 | ✅ 1.0 | 2 |
| **Customer Success Engineer** | — | — | **✅ 1.0 (M14)** | ✅ 1.0 | Customer |
| **ML Engineer #2** | — | — | **✅ 1.0 (M16, opt.)** | ✅ 1.0 | 2 |
| **Compliance Consultant** | — | — | — | **Contract (M20–M26)** | Compliance |
| **Total FTE** | **7** | **9 → 9.5** | **10 → 11** | **11.5** | |

---

## 3. Sprint Breakdown

### Phase 1: Foundation → MVP (Month 1–6, S1–S12)

**Duration:** 24 weeks (vs 18 weeks in 1.5x plan) — **+33% time**

**Team Phase 1:** Tech Lead · BE1 · BE2 · FE1 · Flutter · ML Eng #1 · DevSecOps(contract) · PM = **7 FTE**

#### Key Changes from 1.5x Plan:
- Extended from 9 sprints to 12 sprints
- Added 2 buffer sprints for integration testing
- Added 1 sprint for pilot onboarding and feedback
- MVP moved from S9 (W18) to S12 (W24)
- Target sprint utilization: 50-60% (sustainable pace)

#### Sprint Summary:

| Sprint | Week | Focus | Key Deliverables | Utilization |
|--------|------|-------|------------------|-------------|
| **S1** | W1–2 | Infrastructure Foundation | AWS VPC, ECS, RDS + **RDS Proxy**, Redis **r6g.large**, CI/CD, multi-tenant schema (`tenant_id`, `data_residency`, **`gcp_project_id`**, **`shard_id`** on `tenant_config`), M365 webhook schema, **50 GCP projects provisioned**, **bounded sync worker pool (200 workers)**, **batch secrets schema** | 55% |
| **S2** | W3–4 | Auth + Security | Keycloak HA (4 ECS tasks), JWT JWKS caching, RLS policies, WAF, GCP project assignment logic in SyncScheduler | 55% |
| **S3** | W5–6 | Google Workspace Sync | User/OAuth sync, dashboard skeleton | 60% |
| **S4** | W7–8 | M365 Sync | M365 delta link, webhook renewal | 60% |
| **S5** | W9–10 | Dashboard + Classification | Unified dashboard, asset classification | 55% |
| **S6** | W11–12 | Slack + AWS + RBAC | 4 providers unified, RBAC model | 60% |
| **S7** | W13–14 | Offboarding Workflow | Automated offboarding <5 min, grace period | 65% |
| **S8** | W15–16 | Playbooks + Audit Log | 2 playbooks, immutable audit log | 60% |
| **S9** | W17–18 | **Integration Testing** | Cross-provider integration tests, rate limit validation | 50% |
| **S10** | W19–20 | **Mobile App Beta** | Mobile app TestFlight/Play Console, push notifications | 55% |
| **S11** | W21–22 | **Pilot Onboarding** | 3+ pilot customers onboarded, feedback collection | 40% |
| **S12** | W23–24 | **🏁 MVP Launch** | MVP production launch, offboarding verified <5 min | 50% |

> **MVP = W24** (vs W18 in 1.5x plan) — **+6 weeks buffer**

---

### Phase 2: MVP → v1 (Month 7–13, S13–S26)

**Duration:** 28 weeks (vs 22 weeks in 1.5x plan) — **+27% time**

**Team Phase 2:** 7 FTE (Phase 1) + BE3 (M8) + FE2 (M9) = **9 FTE**

#### Key Changes from 1.5x Plan:
- Extended from 11 sprints to 14 sprints
- Added 2 sprints for Track 2 accuracy validation
- Added 1 sprint for compliance audit preparation
- v1 moved from S20 (W39) to S26 (W52)

#### Sprint Summary:

| Sprint | Week | Track 1 | Track 2 | Gate |
|--------|------|---------|---------|------|
| **S13** | W25–26 | JIT Access + Access Reviews | BE3 onboard, Shadow AI v0.2 | Vanta setup W26 |
| **S14** | W27–28 | Playbook Engine | FE2 onboard, LLM DLP v0.1 | 3 playbooks operational |
| **S15** | W29–30 | 5 Playbooks Complete | Shadow AI v1, DLP extension | Mobile push notifications |
| **S16** | W31–32 | Compliance Mapping | Deepfake POC, schema v1 locked | Gate 3: Shadow AI >95% |
| **S17** | W33–34 | Compliance Reports | T1-T2 integration, Lakera Guard | Gate 4: Deepfake >80% |
| **S18** | W35–36 | **Track 2 Accuracy Validation** | Full accuracy gate evaluation, 30-day holdout testing | All gates documented |
| **S19** | W37–38 | GDPR Automation | T1-T2 integration testing | Pentest vendor LOI signed |
| **S20** | W39–40 | **Pentest Preparation** | Security hardening, SAST/DAST clean | Pentest starts W40 |
| **S21** | W41–42 | **Pentest Remediation #1** | Critical/High findings fixed | 0 Critical open |
| **S22** | W43–44 | **Pentest Remediation #2** | Medium findings, retest | 0 High open |
| **S23** | W45–46 | Dependency Map + Vanta Dry Run | Shadow AI policy enforcement | Vanta >90% pass |
| **S24** | W47–48 | **Chrome Extension Submit** | Extension to Web Store | 2-week review buffer |
| **S25** | W49–50 | **Store Review Buffer** | Track 2 hardening, mobile app submit | Extension approved |
| **S26** | W51–52 | **🏁 v1 Launch** | 5+ pilot customers, SOC 2 Type 1 scheduled | Production cutover |

> **v1 = W52** (vs W39 in 1.5x plan) — **+13 weeks buffer**

---

### Phase 3: v1 → v1.5 (Month 14–19, S27–S38)

**Duration:** 24 weeks (vs 18 weeks in 1.5x plan) — **+33% time**

**Team Phase 3:** 9 FTE + Customer Success Eng (M14) + ML Eng #2 (M16) + DevSecOps → FTE = **11 FTE**

#### Sprint Summary:

| Sprint | Week | Stream A (65%) | Stream B (35%) |
|--------|------|----------------|----------------|
| **S27–S28** | W53–56 | Post-launch Stabilization + AWS v1.1 | Top 20 bugs, M365 wizard UX |
| **S29–S30** | W57–60 | Advanced AI Detection v2 | Dashboard UX redesign |
| **S31** | W61–62 | **AWS Deep Integration** | API documentation |
| **S32–S33** | W63–66 | Business Tier + SOC 2 Type 2 Prep | Billing integration (Stripe) |
| **S34** | W67–68 | **Billing Integration Testing** | Customer portal |
| **S35–S36** | W69–72 | **SOC 2 Type 2 Evidence Review** | Compliance gap remediation |
| **S37** | W73–74 | **Pre-launch Hardening** | Final bug fixes |
| **S38** | W75–76 | **🏁 v1.5 Launch** | Pricing tiers enforced, 10+ paying customers |

> **v1.5 = W76** (vs W57 in 1.5x plan) — **+19 weeks buffer**

---

### Phase 4: v1.5 → v2 (Month 20–26, S39–S52)

**Duration:** 28 weeks (vs 20 weeks in 1.5x plan) — **+40% time**

**Team Phase 4:** 11 FTE + Compliance Consultant (contract M20–M26) = **11.5 FTE (peak)**

#### Sprint Summary:

| Sprint | Week | Focus | Deliverable |
|--------|------|-------|-------------|
| **S39–S40** | W77–80 | Enterprise Features | Multi-tenant enterprise, custom RBAC |
| **S41–S42** | W81–84 | SIEM Integration | Splunk/QRadar webhooks |
| **S43–S44** | W85–88 | **SOC 2 Type 2 Audit Prep** | Evidence packaging, auditor engagement |
| **S45–S46** | W89–92 | ISO 27001 Certification | Stage 2 audit prep, BERT prompt injection |
| **S47–S48** | W93–96 | **Compliance Audit Fieldwork** | Auditor fieldwork, findings remediation |
| **S49** | W97–98 | **Audit Findings Remediation** | All audit findings closed |
| **S50** | W99–100 | **Certification Finalization** | SOC 2 Type 2 report, ISO 27001 certificate |
| **S51** | W101–102 | **Pre-launch Validation** | Multi-region DR drill, final testing |
| **S52** | W103–104 | **🏁 v2 Launch** | Compliance certified, Enterprise tier GA |

> **v2 = W104** (vs W78 in 1.5x plan) — **+26 weeks buffer**

---

## 4. Key Requirements Coverage

| Key Requirement | Milestone | Sprint | Timeline Change |
|---|---|---|---|
| **Asset inventory & classification** | v1 (M13) | S3–S6 core, S23 full | +13 weeks |
| **AI-specific threat surface** | v1 (M13) | S13–S18 (Track 2) | +13 weeks |
| **Access governance** | v1 (M13) | S6–S13 core | +13 weeks |
| **Compliance posture** | v1 (M13) — report-ready | S16–S17 | +13 weeks |
| **Incident playbooks** | v1 (M13) | S7–S8 (2), S14–S15 (5) | +13 weeks |
| **Cost model** | v1.5 (M19) billing live | S26 code-ready, S32–S34 billing | +19 weeks |
| **Integrations** | v1 (M13) | S3–S6 | +13 weeks |

> **Conclusion:** All 7 key requirements delivered at v1 (Month 13), với timeline bền vững 2x.

---

## 5. Compliance Certification Timeline

```
Month 6 (W24):  MVP LAUNCH — Vanta account setup, evidence collection begins
Month 7 (W26):  Vanta OFFICIALLY active — SOC 2 control mapping begins
Month 10 (W40): Pentest begins (LOI signed W38)
Month 13 (W52): v1 LAUNCH
                → SOC 2 Type 1 audit: scheduled with auditor
                → Evidence collection W26→W52 = ~26 weeks (exceeds Type 1 minimum)
Month 14 (W56): ISO 27001 gap analysis begins
Month 16 (W64): ISO 27001 Stage 1 audit (documentation review)
Month 19 (W76): v1.5 LAUNCH
                → SOC 2 Type 2 evidence W52→W76 = 24 weeks (meets minimum)
Month 20 (W80): ISO 27001 Stage 2 audit (implementation review)
Month 22 (W88): SOC 2 Type 2 audit fieldwork begins
                → Evidence W52→W88 = 36 weeks (exceeds 24-week minimum ✅)
Month 26 (W104): v2 LAUNCH
                → SOC 2 Type 2 report issued ✅
                → ISO 27001 certificate issued ✅
```

> ✅ **SOC 2 Type 2 timing:** Evidence window W52→W88 = **36 weeks** (exceeds 24-week minimum with 12-week buffer)

---

## 6. Comparison with Previous Plans

### Timeline Comparison

| Milestone | Original | 1.5x Adjusted | 2x Adjusted | Delta from Original |
|-----------|----------|---------------|-------------|---------------------|
| **MVP** | W12 (M3) | W18 (M4.5) | **W24 (M6)** | **+12 weeks (+100%)** |
| **v1** | W26 (M6) | W39 (M9.75) | **W52 (M13)** | **+26 weeks (+100%)** |
| **v1.5** | W38 (M9) | W57 (M14.25) | **W76 (M19)** | **+38 weeks (+100%)** |
| **v2** | W52 (M12) | W78 (M19.5) | **W104 (M26)** | **+52 weeks (+100%)** |

### Sprint Count Comparison

| Phase | Original | 1.5x Adjusted | 2x Adjusted | Delta |
|-------|----------|---------------|-------------|-------|
| **Phase 1 (MVP)** | 6 sprints | 9 sprints | **12 sprints** | **+6 sprints** |
| **Phase 2 (v1)** | 7 sprints | 11 sprints | **14 sprints** | **+7 sprints** |
| **Phase 3 (v1.5)** | 7 sprints | 9 sprints | **12 sprints** | **+5 sprints** |
| **Phase 4 (v2)** | 6 sprints | 10 sprints | **14 sprints** | **+8 sprints** |
| **Total** | **26 sprints** | **39 sprints** | **52 sprints** | **+26 sprints** |

### Key Benefits of 2x Timeline

✅ **Sustainable Sprint Utilization:**
- Original plan: 75–90% utilization (burnout risk)
- 1.5x plan: 60–75% utilization (still high)
- 2x plan: 50–60% utilization (sustainable long-term)

✅ **Realistic External Dependencies:**
- Google/M365 API verification: 4–6 weeks buffer
- Pentest scheduling: 8-week buffer from LOI to start
- Chrome Web Store review: 4-week buffer
- SOC 2 Type 2 evidence: 12-week buffer above minimum

✅ **Team Health & Knowledge Transfer:**
- Time for proper onboarding (2 weeks per new hire)
- Time for knowledge sharing and documentation
- Time for code review and quality assurance
- Time for learning and skill development

✅ **Risk Mitigation:**
- Buffer for unexpected technical challenges
- Buffer for integration issues
- Buffer for compliance audit findings
- Buffer for customer feedback iterations

---

## 7. External Dependencies & Hard Deadlines

| Deadline | Week | Description | Buffer vs 1.5x |
|----------|------|-------------|----------------|
| Auth provider decision | W1D1 | Choose Keycloak self-host vs Auth0 vs Cognito | No change |
| **ML Engineer #1 onboard** | **W1D1** | Must be hired before project kick-off | No change |
| Google test tenant available | W5 | Internal Google Workspace tenant for S3 development | No change |
| Pilot customer #1 onboard | W20 | At least 1 real customer using staging | +8 weeks |
| **Pentest vendor LOI signed** | **W38** | Hard deadline — 8-week lead time | +24 weeks |
| **Vanta setup active** | **W26** | Need 60+ days evidence for SOC 2 Type 1 | +13 weeks |
| Chrome Web Store submission | W47 | Browser extension needs 2–4 weeks review | +12 weeks |
| iOS App Store submission | W99 | App Store review 1–2 weeks | +25 weeks |
| SOC 2 Type 2 audit sign | W84 | Engage auditor firm | +22 weeks |
| ISO 27001 Stage 2 audit | W80 | Certification 6–8 weeks after audit | +20 weeks |

---

## Summary Dashboard

```
MILESTONE OVERVIEW (ADJUSTED 2x)
══════════════════════════════════════════════════════════════════════

MVP    │ W24  │ M6   │ Asset inventory + offboarding + shadow IT
       │      │      │ Team: 7 FTE + contract (ML Eng #1 from Day 1)
       │      │      │ Buffer: +12 weeks vs original plan
───────┼──────┼──────┼──────────────────────────────────────────────
v1     │ W52  │ M13  │ All 7 key requirements delivered
       │      │      │ Team: 9 FTE
       │      │      │ Compliance: SOC 2 Type 1 audit scheduled
       │      │      │ Buffer: +26 weeks vs original plan
───────┼──────┼──────┼──────────────────────────────────────────────
v1.5   │ W76  │ M19  │ AI detection v2 + AWS v1.1 + pilot feedback
       │      │      │ Team: 11 FTE (2-stream split: 65/35)
       │      │      │ Billing: Stripe live, 10+ paying customers
       │      │      │ Buffer: +38 weeks vs original plan
───────┼──────┼──────┼──────────────────────────────────────────────
v2     │ W104 │ M26  │ Compliance verified + Enterprise + BERT ML
       │      │      │ Team: 11.5 FTE (peak)
       │      │      │ Compliance: SOC 2 Type 2 ✅ + ISO 27001 ✅
       │      │      │ Buffer: +52 weeks vs original plan

══════════════════════════════════════════════════════════════════════
TOTAL PROJECT DURATION
  Original:  52 weeks (12 months)
  1.5x:      78 weeks (19.5 months)
  2x:        104 weeks (26 months)
  Increase:  +100% timeline, sustainable execution
══════════════════════════════════════════════════════════════════════
```

---

*This plan prioritizes team health, quality, and sustainability over speed. A 2-year timeline is realistic for a complex security platform with compliance requirements.*
