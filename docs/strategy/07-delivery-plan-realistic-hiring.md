# SMESec Platform — Realistic Hiring Delivery Plan

**Date:** 2026-05-29  
**Status:** Realistic Scenario — Rolling recruitment from zero  
**Version:** 1.0  
**Scope:** Full roadmap with realistic hiring constraints (30+ months)  
**Context:** Tech Lead starts alone, recruits team progressively

---

## Executive Summary


This is the most realistic plan for the scenario: **Tech Lead starts with zero team, recruiting while building**.

**Realistic hiring assumptions:**
- **Sourcing + Interview:** 2-4 weeks per position
- **Notice period:** 30 days (2-4 weeks) after offer accepted
- **Onboarding:** 1-2 weeks per person
- **Total time to productivity:** 6-10 weeks from start of recruiting to productive contributor

**Overall timeline:**
- **Month 1-3:** TL solo, infrastructure + first hire
- **Month 4-6:** 2-3 people, foundation work
- **Month 7-12:** 4-5 people, MVP development
- **Month 13-18:** 6-7 people, v1 development
- **Month 19-24:** 8-9 people, v1.5 development
- **Month 25-30+:** 10-11 people, v2 development

**Milestones:**
- **MVP**: Month 12 (vs Month 6 in 2x plan)
- **v1**: Month 20 (vs Month 13 in 2x plan)
- **v1.5**: Month 28 (vs Month 19 in 2x plan)
- **v2**: Month 36+ (vs Month 26 in 2x plan)

---

## 1. Hiring Timeline & Team Ramp

### 1.1 Recruitment Assumptions

```
REALISTIC HIRING TIMELINE PER ROLE:

Week 1-2:   Job posting, sourcing, initial screening
Week 3-4:   Technical interviews (2-3 rounds)
Week 5-6:   Offer negotiation, background check
Week 7-10:  Notice period at current company (30 days standard)
Week 11-12: Onboarding, environment setup, codebase walkthrough

Total: 10-12 weeks from job posting to productive contributor
```

**Parallel recruitment constraint:** TL có thể tuyển 2-3 vị trí song song, nhưng không thể interview 5+ người cùng lúc (time constraint).

### 1.2 Progressive Team Build-up

```
MONTH 1-3: TECH LEAD SOLO (1 FTE)
├── Week 1-4:   Infrastructure foundation (AWS, Keycloak, CI/CD)
├── Week 5-8:   Start recruiting: BE1, BE2, FE1 (parallel)
└── Week 9-12:  Continue infrastructure + interview candidates

MONTH 4-6: FIRST HIRES JOIN (2-3 FTE)
├── Month 4:    BE1 joins (Week 16) — onboarding 2 weeks
├── Month 5:    BE2 joins (Week 20) — onboarding 2 weeks
├── Month 6:    FE1 joins (Week 24) — onboarding 2 weeks
└── Parallel:   Start recruiting Flutter, ML Eng #1

MONTH 7-9: CORE TEAM FORMS (4-5 FTE)
├── Month 7:    Flutter joins (Week 28)
├── Month 8:    ML Eng #1 joins (Week 32) — CRITICAL HIRE
├── Month 9:    Team stabilizes, MVP development begins
└── Parallel:   Start recruiting DevSecOps (contract)

MONTH 10-12: MVP PUSH (5-6 FTE)
├── Month 10:   DevSecOps contract starts (0.5 FTE)
├── Month 11:   PM contract starts (0.5 FTE)
├── Month 12:   MVP LAUNCH (limited scope)
└── Parallel:   Start recruiting BE3, FE2 for Track 2

MONTH 13-18: v1 DEVELOPMENT (7-8 FTE)
├── Month 13:   BE3 joins (Track 2)
├── Month 15:   FE2 joins (Browser Extension)
├── Month 18:   Start recruiting Customer Success
└── Parallel:   v1 feature development

MONTH 19-24: v1 LAUNCH & v1.5 (8-9 FTE)
├── Month 20:   v1 LAUNCH
├── Month 21:   Customer Success Eng joins
├── Month 22:   DevSecOps → FTE (from contract)
└── Month 24:   Start recruiting ML Eng #2

MONTH 25-30+: v1.5 → v2 (10-11 FTE)
├── Month 26:   ML Eng #2 joins (optional)
├── Month 28:   v1.5 LAUNCH
├── Month 30:   Compliance Consultant contract
└── Month 36+:  v2 LAUNCH
```

### 1.3 Detailed Hiring Schedule

| Role | Start Recruiting | Offer Accepted | Join Date | Onboard Complete | Notes |
|------|------------------|----------------|-----------|------------------|-------|
| **Tech Lead** | Pre-project | — | **Month 1, Week 1** | Immediate | Already hired |
| **Backend Eng #1** | Month 2, Week 5 | Month 3, Week 12 | **Month 4, Week 16** | Month 4, Week 18 | First hire, critical |
| **Backend Eng #2** | Month 2, Week 5 | Month 4, Week 14 | **Month 5, Week 20** | Month 5, Week 22 | Parallel with BE1 |
| **Frontend Eng #1** | Month 3, Week 9 | Month 4, Week 16 | **Month 6, Week 24** | Month 6, Week 26 | Dashboard critical |
| **Flutter / Mobile** | Month 4, Week 13 | Month 5, Week 20 | **Month 7, Week 28** | Month 7, Week 30 | Mobile app |
| **ML Engineer #1** | Month 5, Week 17 | Month 6, Week 24 | **Month 8, Week 32** | Month 8, Week 34 | Track 2 lead, hard to find |
| **DevSecOps** | Month 7, Week 25 | Month 8, Week 30 | **Month 10, Week 40** | Month 10, Week 41 | Contract 0.5 FTE |
| **PM** | Month 8, Week 29 | Month 9, Week 34 | **Month 11, Week 44** | Month 11, Week 45 | Contract 0.5 FTE |
| **Backend Eng #3** | Month 10, Week 37 | Month 11, Week 42 | **Month 13, Week 52** | Month 13, Week 54 | Track 2 API |
| **Frontend Eng #2** | Month 11, Week 41 | Month 12, Week 46 | **Month 15, Week 60** | Month 15, Week 62 | Browser extension |
| **Customer Success** | Month 16, Week 61 | Month 17, Week 66 | **Month 21, Week 84** | Month 21, Week 86 | Post-v1 launch |
| **ML Engineer #2** | Month 20, Week 77 | Month 21, Week 82 | **Month 26, Week 104** | Month 26, Week 106 | Optional, v1.5+ |
| **Compliance Consultant** | Month 27, Week 105 | Month 28, Week 108 | **Month 30, Week 120** | Month 30, Week 121 | Contract, v2 audit |

**Critical observation:** ML Engineer #1 join Month 8 (vs Day 1 trong plan gốc) = **Track 2 delayed 8 months**. Đây là risk lớn nhất.

---

## 2. Adjusted Milestones

### 2.1 Milestone Timeline

```
Month:  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36
Team:   1  1  1  2  3  3  4  5  5  5  6  6  7  7  8  8  8  8  8  8  9  9  9  9  9 10 10 10 10 10 11 11 11 11 11 11
        |--------Solo--------|-----Ramp Up-----|--------MVP Dev--------|----------v1 Dev----------|-----v1.5 Dev-----|-----v2 Dev-----|
                                                            ↑                                       ↑                  ↑                ↑
                                                           MVP                                      v1                v1.5             v2
                                                        (M12)                                    (M20)              (M28)           (M36+)
```

| Milestone | Month | Team Size | Description | Delay vs 2x Plan |
|-----------|-------|-----------|-------------|------------------|
| **Infrastructure** | M3 | 1 FTE | AWS (VPC, ECS, RDS + RDS Proxy, Redis r6g.large), 50 GCP projects, Keycloak (4 ECS tasks), CI/CD, multi-tenant schema (`gcp_project_id`, `shard_id`, `renewal_bucket` from day 1) | +2 months |
| **Foundation** | M6 | 3 FTE | Google/M365 sync, basic dashboard | +3 months |
| **MVP** | M12 | 6 FTE | Asset inventory + offboarding + shadow IT | +6 months |
| **v1** | M20 | 8 FTE | All key requirements, SOC 2 Type 1 | +7 months |
| **v1.5** | M28 | 9 FTE | AI detection + billing live | +9 months |
| **v2** | M36+ | 11 FTE | SOC 2 Type 2 + ISO 27001 certified | +10 months |

### 2.2 Scope Adjustments by Milestone

**MVP (Month 12) — Reduced Scope:**
- ✅ Google Workspace sync (users, OAuth apps)
- ✅ M365 sync (basic, no webhook renewal yet)
- ✅ Automated offboarding <10 min (not <5 min)
- ✅ 1 incident playbook (Offboarding only)
- ✅ Basic dashboard (read-only)
- ❌ Slack integration (deferred to v1)
- ❌ AWS IAM integration (deferred to v1)
- ❌ Mobile app (deferred to v1)
- ❌ Track 2 features (ML Eng #1 only joined Month 8)

**v1 (Month 20) — Core Features:**
- ✅ All MVP features + Slack + AWS IAM
- ✅ 3 incident playbooks
- ✅ Mobile app (iOS + Android beta)
- ✅ RBAC dashboard
- ✅ JIT access
- ✅ Compliance dashboard (ISO 27001 + SOC 2 Type 1)
- ⚠️ Track 2: Shadow AI governance only (beta)
- ❌ LLM DLP extension (deferred to v1.5)
- ❌ Deepfake detection (deferred to v1.5)

**v1.5 (Month 28) — AI Features:**
- ✅ All v1 features
- ✅ LLM DLP browser extension
- ✅ Deepfake detection (Hive API)
- ✅ Prompt injection (Lakera Guard)
- ✅ Billing live (Stripe)
- ✅ 10+ paying customers

**v2 (Month 36+) — Enterprise:**
- ✅ SOC 2 Type 2 certified
- ✅ ISO 27001 certified
- ✅ Enterprise tier
- ✅ BERT prompt injection (if needed)

---

## 3. Phase-by-Phase Breakdown

### Phase 1: Solo Foundation (Month 1-3, TL only)

**Team:** Tech Lead (1 FTE)

**TL Activities:**
- **Week 1-4:** AWS infrastructure (VPC, ECS, RDS + **RDS Proxy mandatory**, Redis **r6g.large**, S3 Object Lock) + **50 GCP projects provisioned** (Terraform)
- **Week 5-8:** Keycloak setup (4 ECS tasks), multi-tenant schema (`tenant_id`, `data_residency`, **`gcp_project_id`**, **`shard_id`** on `tenant_config`), RLS policies, **batch secrets schema** (1 JSON/tenant), M365 webhook schema
- **Week 9-12:** CI/CD pipeline, **bounded sync worker pool scaffold** (200-worker pool, job queue), basic integration skeleton

**Parallel:** Recruiting BE1, BE2, FE1 (3 roles in parallel)

**Deliverables:**
- Deployable infrastructure (RDS Proxy + Redis r6g.large provisioned for 1K tenant capacity)
- Multi-tenant database schema (`gcp_project_id`, `shard_id`, `renewal_bucket` columns from day 1)
- 50 GCP projects provisioned (Terraform), each with dedicated service account
- Bounded sync worker pool (200-worker) scaffold ready for Google/M365 sync
- Keycloak SSO working (4 ECS tasks, HA)
- CI/CD green

**Risks:**
- TL burnout (working alone 3 months)
- Recruitment delays (if no candidates found)
- Infrastructure decisions made without team input

**Mitigation:**
- TL works 40-50 hours/week max (sustainable pace)
- Use contractors for non-critical work (DevOps setup)
- Document all decisions for future team review

---

### Phase 2: First Hires Onboard (Month 4-6, 2-3 FTE)

**Team:** TL + BE1 (M4) + BE2 (M5) + FE1 (M6)

**Month 4 (TL + BE1):**
- BE1 onboarding (2 weeks)
- Google Workspace sync (BE1 lead, TL review)
- TL: Continue recruiting + architecture decisions

**Month 5 (TL + BE1 + BE2):**
- BE2 onboarding (2 weeks)
- M365 sync (BE2 lead)
- Keycloak production hardening (BE1)
- TL: Start recruiting Flutter + ML Eng #1

**Month 6 (TL + BE1 + BE2 + FE1):**
- FE1 onboarding (2 weeks)
- Dashboard v1 (FE1 lead)
- Asset classification engine (BE1)
- OAuth scope risk scoring (BE2)

**Deliverables:**
- Google Workspace + M365 sync working
- Basic dashboard (asset inventory view)
- Shadow IT detection (rule-based)

**Risks:**
- New hires need significant TL time (onboarding overhead)
- Productivity drops during onboarding weeks
- Team dynamics forming (communication overhead)

---

### Phase 3: Core Team Forms (Month 7-9, 4-5 FTE)

**Team:** TL + BE1 + BE2 + FE1 + Flutter (M7) + ML Eng #1 (M8)

**Month 7:**
- Flutter onboarding
- Mobile app scaffold (Flutter lead)
- RBAC model (BE1)
- Dashboard polish (FE1)

**Month 8:**
- **ML Eng #1 onboarding (critical)**
- Track 2 kickoff: literature review, dataset collection
- Offboarding workflow design (BE1 + TL)
- M365 webhook renewal (BE2)

**Month 9:**
- Track 2: Baseline models, SageMaker setup
- Offboarding workflow implementation (BE1)
- 1 incident playbook (Offboarding)
- Mobile app beta (Flutter)

**Deliverables:**
- Automated offboarding workflow
- 1 incident playbook
- Mobile app beta (TestFlight/Play Console)
- Track 2 R&D started (8 months late)

**Risks:**
- ML Eng #1 joining 8 months late = Track 2 significantly delayed
- Team size still small (5 FTE) for ambitious scope
- MVP scope must be reduced

---

### Phase 4: MVP Development (Month 10-12, 5-6 FTE)

**Team:** TL + BE1 + BE2 + FE1 + Flutter + ML Eng #1 + DevSecOps (0.5) + PM (0.5)

**Month 10:**
- DevSecOps contract starts (CI/CD hardening, SAST)
- Compliance foundation (audit log, S3 Object Lock)
- Track 2: Prompt injection prototype

**Month 11:**
- PM contract starts (sprint planning, pilot outreach)
- Dashboard UX improvements
- Track 2: Shadow AI risk scoring model v0.1

**Month 12:**
- **MVP LAUNCH (reduced scope)**
- 2-3 pilot customers onboarded
- Offboarding verified <10 min (not <5 min yet)
- Track 2: Still in R&D phase (no production features)

**MVP Scope (Reduced):**
- ✅ Google Workspace + M365 sync
- ✅ Automated offboarding <10 min
- ✅ 1 incident playbook
- ✅ Basic dashboard
- ✅ Mobile app beta
- ❌ Slack (deferred)
- ❌ AWS IAM (deferred)
- ❌ Track 2 features (not ready)

**Risks:**
- MVP scope significantly reduced vs original plan
- Track 2 features not in MVP (competitive disadvantage)
- Only 2-3 pilot customers (limited feedback)

---

### Phase 5: v1 Development (Month 13-20, 7-8 FTE)

**Team:** TL + BE1 + BE2 + FE1 + Flutter + ML Eng #1 + DevSecOps + PM + BE3 (M13) + FE2 (M15)

**Month 13-14:**
- BE3 onboarding (Track 2 API)
- Slack integration (BE2)
- AWS IAM integration (BE1)
- Track 2: Shadow AI governance v1

**Month 15-16:**
- FE2 onboarding (Browser Extension)
- JIT access (BE1)
- 3 incident playbooks total (BE2)
- Track 2: LLM DLP extension v0.1

**Month 17-18:**
- Compliance dashboard (FE1)
- RBAC dashboard (FE1)
- Track 2: Deepfake POC (ML Eng #1)
- Vanta setup (DevSecOps)

**Month 19-20:**
- **v1 LAUNCH**
- 5+ pilot customers
- SOC 2 Type 1 audit scheduled
- Track 2: Shadow AI governance (beta)

**v1 Scope:**
- ✅ All MVP features + Slack + AWS IAM
- ✅ 3 incident playbooks
- ✅ Mobile app production
- ✅ RBAC + JIT access
- ✅ Compliance dashboard
- ⚠️ Track 2: Shadow AI only (beta)
- ❌ LLM DLP (not ready)
- ❌ Deepfake (not ready)

---

### Phase 6: v1.5 Development (Month 21-28, 8-9 FTE)

**Team:** v1 team + Customer Success (M21) + DevSecOps → FTE (M22)

**Month 21-22:**
- Customer Success onboarding
- Post-v1 stabilization
- Track 2: LLM DLP extension v1

**Month 23-24:**
- Track 2: Deepfake detection v1
- Track 2: Prompt injection (Lakera Guard)
- Billing integration (Stripe)

**Month 25-26:**
- Track 2: Full T1-T2 integration
- Chrome Web Store submission
- Pricing tiers enforced

**Month 27-28:**
- **v1.5 LAUNCH**
- 10+ paying customers
- All Track 2 features (beta → GA)
- Billing live

---

### Phase 7: v2 Development (Month 29-36+, 10-11 FTE)

**Team:** v1.5 team + ML Eng #2 (M26) + Compliance Consultant (M30)

**Month 29-32:**
- Enterprise features
- SIEM integration
- SOC 2 Type 2 prep

**Month 33-36:**
- SOC 2 Type 2 audit
- ISO 27001 certification
- **v2 LAUNCH**

---

## 4. Risk Analysis

### 4.1 Critical Risks

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| **ML Eng #1 not found by Month 6** | High | Critical | Start recruiting Month 1 (not Month 5). Use ML contractor if needed. Accept Track 2 delay. |
| **TL burnout (Month 1-3 solo)** | Medium | Critical | 40-50 hour weeks max. Use contractors for non-critical work. Take breaks. |
| **First hires quit during onboarding** | Low | High | Strong onboarding process. Clear expectations. Competitive compensation. |
| **Recruitment takes longer than 10 weeks** | Medium | High | Start recruiting earlier. Have backup candidates. Use recruiters. |
| **Team productivity lower than expected** | High | Medium | Reduce scope. Extend timeline. Focus on MVP first. |
| **MVP delayed beyond Month 12** | Medium | High | Cut scope aggressively. Focus on core value prop only. |

### 4.2 Hiring Risks

**Hardest roles to fill:**
1. **ML Engineer #1** — Specialized skill set, high demand
2. **Tech Lead** — Must be senior, hands-on, and willing to work solo initially
3. **DevSecOps** — Security + compliance expertise rare at SME pricing

**Mitigation strategies:**
- **ML Eng #1:** Start recruiting Month 1 (not Month 5). Offer equity. Remote OK.
- **Tech Lead:** Already hired (pre-project requirement)
- **DevSecOps:** Start with contract (0.5 FTE), convert to FTE later

---

## 5. Comparison with Other Plans

| Metric | Original Plan | 2x Plan | Realistic Hiring |
|--------|---------------|---------|------------------|
| **MVP** | Month 3 | Month 6 | **Month 12** |
| **v1** | Month 6 | Month 13 | **Month 20** |
| **v1.5** | Month 9 | Month 19 | **Month 28** |
| **v2** | Month 12 | Month 26 | **Month 36+** |
| **Team at MVP** | 7 FTE | 7 FTE | **6 FTE** |
| **Team at v1** | 9 FTE | 9 FTE | **8 FTE** |
| **ML Eng #1 join** | Day 1 | Day 1 | **Month 8** |
| **Track 2 at MVP** | R&D complete | R&D complete | **Just started** |
| **Track 2 at v1** | Beta features | Beta features | **1 feature only** |

**Key insight:** Realistic hiring adds **+6-10 months** to every milestone due to:
- Recruitment time (2-4 weeks)
- Notice period (4 weeks)
- Onboarding (2 weeks)
- Ramp-up to full productivity (4-8 weeks)

---

## 6. Recommendations

### 6.1 For Tech Lead Starting Solo

**Month 1-3 Priorities:**
1. **Infrastructure first** — AWS, Keycloak, CI/CD (foundation for team)
2. **Recruit aggressively** — Start recruiting BE1, BE2, FE1 immediately
3. **Document everything** — Future team needs context
4. **Sustainable pace** — 40-50 hours/week max (marathon, not sprint)
5. **Use contractors** — DevOps, design, non-critical work

**What NOT to do:**
- ❌ Try to build features alone (wait for team)
- ❌ Work 60-80 hour weeks (burnout risk)
- ❌ Skip documentation (team will be lost)
- ❌ Delay recruiting (every week matters)

### 6.2 Hiring Strategy

**Parallel recruitment:**
- Month 1-2: BE1, BE2, FE1 (3 roles)
- Month 4-5: Flutter, ML Eng #1 (2 roles)
- Month 7-8: DevSecOps, PM (2 roles)
- Month 10-11: BE3, FE2 (2 roles)

**Prioritization:**
1. **BE1** — Most critical (integration sync engine)
2. **BE2** — Second critical (auth, security)
3. **FE1** — Dashboard (customer-facing)
4. **Flutter** — Mobile app
5. **ML Eng #1** — Track 2 (can't start without this)

### 6.3 Scope Management

**MVP must be ruthlessly scoped:**
- Focus on 1 core value prop: "Automated offboarding <10 min"
- Google + M365 only (no Slack, no AWS)
- 1 playbook only (Offboarding)
- Basic dashboard (read-only)
- No Track 2 features (not ready)

**v1 adds breadth:**
- Slack + AWS IAM
- 3 playbooks
- RBAC + JIT access
- Mobile app
- 1 Track 2 feature (Shadow AI governance)

**v1.5 adds AI:**
- All Track 2 features
- Billing live
- 10+ paying customers

---

## 7. Success Criteria

### MVP (Month 12)
- ✅ 2-3 pilot customers using product
- ✅ Automated offboarding <10 min
- ✅ Google + M365 sync working
- ✅ 1 incident playbook operational
- ✅ Team size: 6 FTE

### v1 (Month 20)
- ✅ 5+ pilot customers
- ✅ All 7 key requirements (reduced scope)
- ✅ SOC 2 Type 1 audit scheduled
- ✅ 1 Track 2 feature (Shadow AI)
- ✅ Team size: 8 FTE

### v1.5 (Month 28)
- ✅ 10+ paying customers
- ✅ All Track 2 features GA
- ✅ Billing live (Stripe)
- ✅ Team size: 9 FTE

### v2 (Month 36+)
- ✅ SOC 2 Type 2 certified
- ✅ ISO 27001 certified
- ✅ Enterprise tier
- ✅ Team size: 11 FTE

---

## Conclusion

**Realistic hiring timeline = 3 years (36 months) to v2**, not 2 years.

**Key factors:**
- Recruitment takes time (10-12 weeks per role)
- Team ramps slowly (1-2 new hires per quarter sustainable)
- Onboarding overhead reduces productivity
- ML Eng #1 joining Month 8 (not Day 1) delays Track 2 significantly

**This is the most realistic plan** for a Tech Lead starting with zero team, building a complex security platform with compliance requirements.

**Trade-off:** Slower time-to-market, but sustainable team growth and lower burnout risk.
