# SMESec Platform — Project Health Metrics Scorecard

**Date:** 2026-05-28  
**Status:** Active — Review weekly (PM) · Review monthly (CTO/CPO)  
**Purpose:** Single source of truth for whether this project is on track. Every number here is either a hard gate, an early warning signal, or a survival threshold.

---

## Table of Contents

1. [How to Use This Document](#1-how-to-use-this-document)
2. [Tier 1 — Survival Metrics (Kill-Switch Indicators)](#2-tier-1--survival-metrics-kill-switch-indicators)
3. [Tier 2 — Delivery Gates (Pass/Fail at Each Milestone)](#3-tier-2--delivery-gates-passfail-at-each-milestone)
4. [Tier 3 — Track 2 ML Accuracy Gates](#4-tier-3--track-2-ml-accuracy-gates)
5. [Tier 4 — Weekly Operational Health](#5-tier-4--weekly-operational-health)
6. [Tier 5 — Business & Revenue Health](#6-tier-5--business--revenue-health)
7. [Tier 6 — Team & Execution Velocity](#7-tier-6--team--execution-velocity)
8. [Early Warning System — Red Flags](#8-early-warning-system--red-flags)
9. [Monthly Review Checklist](#9-monthly-review-checklist)
10. [How These Numbers Connect](#10-how-these-numbers-connect)

---

## 1. How to Use This Document

This document defines three response levels:

| Signal | Threshold | Response |
|---|---|---|
| 🟢 **Green** | Within target | Continue. Note in weekly standup. |
| 🟡 **Yellow** | Within 10–15% of threshold | PM escalates. Root cause analysis within 48h. |
| 🔴 **Red** | Threshold breached | Executive decision required within 24h. Stop/pivot/escalate. |

**Reading order if you have 5 minutes:** Read Tier 1 → Tier 2 gate for the current milestone → that's all you need.

**Reading order if you have 30 minutes:** Read all tiers sequentially. The dependencies between tiers are explained in Section 10.

---

## 2. Tier 1 — Survival Metrics (Kill-Switch Indicators)

These are the five numbers that determine if the project lives or dies. If any one of these turns red and stays red for 2+ consecutive weeks, the project needs a formal pivot discussion.

### S1. Offboarding Time — Core Product Promise

| Metric | Target | Yellow | Red |
|---|---|---|---|
| Automated offboarding end-to-end time | **<5 minutes** | 5–8 min | >8 min |
| Emergency offboarding (override mode) | **<2 minutes** | 2–4 min | >4 min |
| Offboarding success rate (no manual fallback) | **>98%** | 95–98% | <95% |

**Why it kills you:** This is the product's core promise to every customer. It is the demo you show every prospect. It is the reason Track 1 exists. A single failure in a customer demo costs the deal. A regression from CI-green to prod-yellow means something is broken in Step Functions or the downstream API integration — fix immediately.

**How to measure:** Automated E2E test in CI (every PR). Manual timed test on staging before every customer demo. Monitor `OffboardingWorkflow` Step Functions execution history in AWS Console — p95 duration.

---

### S2. Gross Margin — Business Model Validity

| Metric | Target | Yellow | Red |
|---|---|---|---|
| Gross margin (revenue − COGS) / revenue | **>65%** at launch → **>85%** at 100 customers → **>99%** at 1K tenants | 60–65% | <60% |
| COGS per tenant/month (Track 1 + Track 2) | **<$354** (base) → **<$291** (6-month optimized) | $354–$420 | >$420 |
| SageMaker inference cost per tenant/month | **<$50** | $50–$75 | >$75 |

**Why it kills you:** Below 60% gross margin, the unit economics break before you reach 100 customers. You cannot fund growth, hire support, or pay for the SOC 2 audit. The cost model assumes ~$70K/yr infra at 1K tenant capacity — infrastructure is pre-provisioned for scale from Sprint 1. At early growth (50 tenants), gross margin is ~85%; at 1K tenants capacity, gross margin reaches ~99%. If AWS costs run 30% over, gross margin at early stage drops to ~80% (still viable). See cost-analysis.md Section 1.2 Risk Scenarios.

**How to measure:** Monthly P&L review. Tag all AWS resources with `tenant_id` and use Cost Explorer to get per-tenant actual costs. Compare against the cost-analysis.md baseline monthly.

---

### S3. Monthly Churn Rate — Customer Retention

| Metric | Target | Yellow | Red |
|---|---|---|---|
| Monthly logo churn (customers lost) | **<3%** | 3–5% | >5% |
| Monthly revenue churn (MRR lost) | **<2%** | 2–4% | >4% |
| Net Revenue Retention (NRR) | **>110%** | 100–110% | <100% |

**Why it kills you:** At 5% monthly churn, you lose 46% of customers every year. CAC of $800–$5,000 requires at minimum 12–18 months to recover. A churn rate above 5% means the product does not solve the problem well enough to retain customers — no amount of sales velocity fixes this.

**How to measure:** Track in CRM from Month 3 onward. Weekly: count active vs churned tenants. Monthly: calculate MRR churn. NRR = (Starting MRR + expansion − contraction − churn) / Starting MRR × 100.

---

### S4. Pilot-to-Paid Conversion — Sales Funnel Validity

| Metric | Target | Yellow | Red |
|---|---|---|---|
| Pilot funnel: outreach → qualified | **30%** (30/100) | 20–30% | <20% |
| Qualified → demo | **50%** (15/30) | 35–50% | <35% |
| Demo → pilot | **33%** (5/15) | 20–33% | <20% |
| **Pilot → paying customer** | **>60%** | 40–60% | <40% |
| Time to "first insight" (post-OAuth grant) | **<30 minutes** | 30–60 min | >60 min |

**Why it kills you:** The entire revenue model depends on converting pilots to paying customers. If pilot → paid conversion is below 40%, the product either (a) does not demonstrate enough value quickly enough, or (b) is priced wrong. "Time to first insight <30 min" is the activation metric — it predicts conversion rate. If this degrades, conversion follows within 30 days.

**How to measure:** BD Consultant tracks funnel in CRM. "First insight" = first time customer sees a real shadow IT detection or risk score on their own data. Automated event logged in backend (`user.first_value_event`).

---

### S5. Runway vs. Burn — Financial Survival

| Metric | Target | Yellow | Red |
|---|---|---|---|
| Runway (months of cash remaining) | **>12 months** | 6–12 months | <6 months |
| Monthly burn rate vs plan | **Within ±15%** | 15–30% over | >30% over |
| Headcount cost as % of total burn | **<70%** | 70–80% | >80% |

**Why it kills you:** You need to reach v1 (Month 6) with runway intact to close Series A or bridge round. If burn runs 30% over plan due to unplanned hires or AWS cost overruns, you may not survive to v1.

**How to measure:** Monthly finance review. Compare actual headcount costs against the delivery-plan.md headcount ramp table. AWS Cost Explorer weekly.

---

## 3. Tier 2 — Delivery Gates (Pass/Fail at Each Milestone)

These are binary gates. The project does not advance to the next phase without passing them. Partial credit is not accepted.

### MVP Gate — Week 12

> The question: "Can you revoke all access for a departing employee in 5 minutes?"

| Gate Item | Pass Condition | Notes |
|---|---|---|
| Offboarding <5 min | ✅ Automated timed test <5 min in CI | Includes Google WS + M365 revocation |
| Grace period & rollback | ✅ 30-min grace period configurable · rollback 24h working | R-C1 |
| Idempotency | ✅ Duplicate offboarding requests safe | R-C1 |
| Tenant isolation CI | ✅ No cross-tenant data leakage | CI test green on every PR |
| M365 webhook renewal | ✅ `subscription_registry` + 12h renewal job live | R-C3 |
| Google rate limit strategy | ✅ Per-cluster GCP service account quota distribution | R-C2 |
| Keycloak HA | ✅ 2 ECS tasks min · JWKS cache active | R-C6 |
| 3+ pilot customers staged | ✅ Seen "first insight" on staging in <30 min | MVP is not MVP without at least 3 real customers |
| Audit log active | ✅ S3 Object Lock + per-tenant KMS from Day 1 | Non-negotiable compliance evidence |
| 2 incident playbooks | ✅ Offboarding + Credential Compromise working | Wizard UI |

**Consequence of failure:** MVP launch delayed. Each week of delay at this stage = 1 week less SOC 2 evidence collection. SOC 2 Type 2 requires 6 months of evidence from a stable production system.

---

### v1 Gate — Week 26

| Gate Item | Pass Condition | Notes |
|---|---|---|
| All MVP gates still passing | ✅ No regression | |
| JIT access + access reviews | ✅ Request → approval → auto-revoke workflow | |
| 5 incident playbooks | ✅ Including wizard UI | |
| Slack full + AWS IAM basic | ✅ 4 providers fully integrated | |
| Vanta evidence collection | ✅ Live from W13 (audit collection starts here) | |
| SOC 2 Type 1 audit scheduled | ✅ Auditor engaged | |
| ISO 27001 report-ready | ✅ Control mapping complete | |
| LLM DLP browser ext (beta) | ✅ Chrome extension sideloaded, PII detection working | |
| Shadow AI governance v1 (beta) | ✅ Top-100 AI tool classification | |
| Track 2 Accuracy Gates 1–3 passed | ✅ See Tier 3 | |
| GDPR erasure endpoint | ✅ `/api/v1/gdpr/erasure` live · KMS key destruction workflow | |

---

### v1.5 Gate — Week 38

| Gate Item | Pass Condition | Notes |
|---|---|---|
| Pricing tiers enforced (Starter / Growth / Business) | ✅ Backend quota enforcement + Stripe billing | |
| 10+ paying customers on production | ✅ Revenue > $0 | |
| Billing integration live (Stripe) | ✅ | |
| Chrome Web Store published | ✅ Not sideloaded | |
| AI detection accuracy | ✅ Deepfake + LLM DLP >90% combined | |
| Prompt injection (Lakera Guard) production | ✅ TPR >85%, FPR <2% on 30-day holdout | Independently evaluated |
| SOC 2 Type 2 evidence collection running | ✅ Running continuously since W26 (>12 weeks) | |
| SageMaker model monitoring (drift detection) | ✅ Active | |
| AWS v1.1 (CloudTrail, IAM deep) | ✅ | |

---

### v2 Gate — Week 52

| Gate Item | Pass Condition | Notes |
|---|---|---|
| SOC 2 Type 2 report received | ✅ From auditor | |
| ISO 27001 certificate received | ✅ | |
| BERT prompt injection | ✅ TPR >85%, FPR <2% on holdout · or opt-in preview with Lakera GA | |
| Enterprise tier live | ✅ Custom pricing · dedicated CSM · SIEM integration | |
| All Track 2 features graduated from beta | ✅ SLA guarantees applied | |
| Multi-region DR drill completed | ✅ RTO/RPO documented and tested | |
| 99.95% uptime target achievable | ✅ Verified from 12 months of monitoring data | |

---

## 4. Tier 3 — Track 2 ML Accuracy Gates

**Critical clarification:** All accuracy gates are evaluated independently by the SMESec ML team on SMESec-specific production holdout data. Vendor API uptime SLAs (Lakera Guard, Hive Moderation) do not satisfy accuracy gates. Failed gates → feature stays `beta`, opt-in only. Track 1 is never blocked by Track 2.

### Accuracy Thresholds

| Feature | Metric | Production Gate | Amber Warning | Sprint |
|---|---|---|---|---|
| **Prompt injection** | TPR (True Positive Rate) | **>85%** | <87% (approaching floor) | Gate: S5 (W10), BERT upgrade: S24 |
| **Prompt injection** | FPR (False Positive Rate) | **<2%** | >1.5% (approaching ceiling) | Same |
| **LLM DLP (Critical PII)** | Detection rate | **>99%** | <99.3% | S5 (W10) |
| **LLM DLP (Critical PII)** | False Positive rate | **<5%** | >4% | S5 (W10) |
| **Shadow AI classification** | Top-200 AI tool accuracy | **>95%** | <96% | S9 (W18) |
| **Deepfake detection** | Voice deepfake TPR | **>80%** (Hive API) | <82% | S10 (W20) |
| **Deepfake + OOV combined** | Fraud prevention rate | **~99%** | <97% | S10 (W20) |

### Why 80% Deepfake Gate Is Correct

The 80% gate is not a low bar — it is the right bar because:

1. Hive Moderation alone: ~80% detection of synthetic audio in controlled test conditions
2. Out-of-band verification (OOV) workflow: employees must confirm high-value decisions (finance transfers, credential changes) through a second independent channel (SMS/phone)
3. Combined effectiveness: an attacker must beat both Hive detection AND defeat the OOV workflow simultaneously → ~99% combined fraud prevention rate
4. Raising the gate to 95% for Hive alone is not realistic (adversarial arms race; vendors cannot guarantee it)

**Do not change the 80% Hive gate.** The defense-in-depth layered approach is architecturally correct.

### Accuracy Gate Roadmap — Prompt Injection

| Version | Approach | Expected TPR | Sprint |
|---|---|---|---|
| **v1** | Lakera Guard API alone | ~85–90% | S8–S11 |
| **v1.5** | Lakera Guard + regex pre-filter | ~92–93% | S14–S19 (no new model) |
| **v2 / Enterprise** | Lakera Guard + fine-tuned BERT ensemble | **>95%** TPR, FPR <2% | S23–S24 (only if: cost prohibitive + ≥50K labeled samples available) |

### Monthly Re-Evaluation Protocol

Every 30 days after production launch:
- Collect 500 labeled production samples (mix of real positives, benign, edge cases)
- Run each feature model against the labeled set
- Alert if:
  - Prompt injection TPR drops below **82%** (3pp below gate — early warning)
  - Deepfake TPR drops below **75%** (5pp below gate — early warning)
  - Shadow AI accuracy drops below **93%**
  - LLM DLP critical PII detection drops below **98%**

If alert fires → ML Engineer runs root cause analysis within 48h. Likely causes: data drift (new injection patterns), vendor model change (Lakera API update), or distribution shift in customer prompts.

---

## 5. Tier 4 — Weekly Operational Health

These metrics are checked every week by the PM and Tech Lead. They are leading indicators of problems that will hit delivery gates in 2–4 weeks if not addressed.

### System Reliability

| Metric | Target | Yellow | Red |
|---|---|---|---|
| API uptime (all endpoints) | **>99.9%** | 99.5–99.9% | <99.5% |
| Offboarding Step Functions p95 latency | **<3 min** | 3–5 min | >5 min |
| Google/M365 sync lag (delta sync) | **<15 min** | 15–30 min | >30 min |
| Tenant isolation CI test | **100% green on every PR** | — | Any red = merge blocked |
| Keycloak JWKS cache hit rate | **>95%** | 90–95% | <90% |

### Security

| Metric | Target | Yellow | Red |
|---|---|---|---|
| OWASP-detected vulnerabilities (SAST/DAST) | **0 Critical, 0 High** | — | Any Critical/High = release blocked |
| Secrets Manager rotation (90-day) | **100% compliant** | — | Any expired secret = P1 |
| RLS test coverage (cross-tenant queries) | **100%** | — | <100% = release blocked |
| WAF block rate (anomalous traffic) | Baseline ±20% | ±20–50% spike | >50% spike (investigate) |

### Track 2 ML Operational

| Metric | Target | Yellow | Red |
|---|---|---|---|
| Lakera Guard API response time (p95) | **<300ms** | 300–500ms | >500ms (switch to WASM fallback) |
| WASM BERT-tiny fallback activation rate | **<1% of sessions** | 1–5% | >5% (Lakera API issue) |
| SageMaker inference endpoint uptime | **>99.5%** | 99–99.5% | <99% |
| Prompt injection false positive incidents (customer complaints) | **0 per week** | 1–2 | >2 (FPR spiking, investigate immediately) |

---

## 6. Tier 5 — Business & Revenue Health

### Revenue Trajectory

| Metric | Target | Yellow | Red |
|---|---|---|---|
| MRR growth month-over-month | **>15%** M4–M9 · **>10%** M10–M12 | 8–15% | <8% |
| ARR at v1 (W26) | **>$50K** (approx 5 paying customers) | $25–50K | <$25K |
| ARR at v1.5 (W38) | **>$120K** (10+ paying customers) | $60–120K | <$60K |
| ARR at v2 (W52) | **>$480K** (50+ customers model) | $240–480K | <$240K |
| Average Contract Value (ACV) | **>$8K/year** (Growth tier baseline) | $5–8K | <$5K |

**ARR breakeven reference:** $480K ARR at 50+ customers (Growth tier average) covers ~$70K/yr infra cost with >85% gross margin on infrastructure. Architecture is designed for 1K tenants — growing to 1K customers generates $9.6M ARR at ~99% gross margin. Per-tenant infra cost falls as customer count grows toward 1K capacity.

### Customer Acquisition

| Metric | Target | Yellow | Red |
|---|---|---|---|
| CAC (direct sales) | **<$3,000** | $3,000–5,000 | >$5,000 |
| CAC (via MSP partner) | **<$800** | $800–1,200 | >$1,200 |
| Time to first value (post-onboarding) | **<30 min** | 30–60 min | >60 min |
| Payback period (months to recover CAC) | **<12 months** | 12–18 months | >18 months |
| MSP partner agreements signed (by M6) | **≥3 partners** | 1–2 partners | 0 partners |

### Customer Health (Post-MVP)

| Metric | Target | Yellow | Red |
|---|---|---|---|
| NPS (Net Promoter Score) | **>40** | 20–40 | <20 |
| Customer health score (CHS) | **>70/100** for 80% of customers | 60–70% at threshold | >20% below 60 |
| Feature adoption — offboarding automation | **>80% of customers use it** | 60–80% | <60% |
| Feature adoption — shadow IT alerts | **>60% of customers** | 40–60% | <40% |
| Support ticket volume per customer/month | **<2 tickets** | 2–5 tickets | >5 tickets |

---

## 7. Tier 6 — Team & Execution Velocity

### Sprint Velocity

| Metric | Target | Yellow | Red |
|---|---|---|---|
| Sprint goal completion rate | **>80%** of committed stories | 65–80% | <65% |
| Sprint carryover rate | **<20%** of points carry to next sprint | 20–35% | >35% |
| Unplanned work (interrupt-driven) | **<15%** of sprint capacity | 15–25% | >25% |
| Tech debt % of sprint capacity | **<20%** allocated deliberately | >30% unplanned | Debt blocking feature work |

### Critical Hiring Gates

| Role | Must Be Onboarded By | If Late |
|---|---|---|
| **ML Engineer #1** | **Week 1, Day 1** | Track 2 R&D falls behind — Gate 1 (W10) becomes at risk |
| **BD Consultant (contract)** | **Week 1** | Pilot pipeline is empty at MVP — no customers to onboard |
| **Backend Eng #3 (Track 2)** | **Month 4 (T4)** | Track 2 feature delivery S8–S11 becomes single-threaded |
| **Frontend Eng #2 (Browser Ext)** | **Month 4.5** | Browser extension sprint at S8 has no FE owner |
| **DevSecOps → FTE** | **Month 7** | Security CI coverage degrades in Phase 3 |
| **Customer Success Engineer** | **Month 7** | Churn increases — no one owns retention at v1 scale |

### Track 2 R&D Milestones (leading indicators for accuracy gates)

| Milestone | Week | If Missed |
|---|---|---|
| PromptBench dataset evaluation complete | W4 | Gate 1 (W10) at risk |
| Lakera Guard API: cost baseline + first 1,000 test calls | W6 | S8 integration blocked |
| Presidio WASM compile pipeline confirmed in browser | W6 | LLM DLP Gate 2 (W10) at risk |
| BERT-tiny ONNX first accuracy measurement vs production holdout | W8 | S5 Gate 1 report will show unknown risk |
| Shadow AI tool registry v0.1 (100+ tools) | W2 | Gate 3 (W18) timeline compressed |
| Hive Moderation API first test audio analysis | W10 | Gate 4 (W20) timeline compressed |

---

## 8. Early Warning System — Red Flags

These are specific combinations of metrics that — when appearing together — signal a systemic problem, not a one-off issue.

### Red Flag 1: Product-Market Fit Problem

**Pattern:** Pilot → paid conversion <40% **AND** NPS <30 **AND** Feature adoption (offboarding) <60%

**Diagnosis:** The product is not solving the pain point compellingly enough. Customers see it, try it, and don't find enough value to pay.

**Response:** Stop new feature development. Run 20 customer interviews in 2 weeks. Identify the top 3 underdelivered promises. Treat as P0.

---

### Red Flag 2: Unit Economics Breaking

**Pattern:** COGS/tenant >$420 **AND** Churn >5% **AND** NRR <100%

**Diagnosis:** You are spending more to serve customers than you earn, and customers are leaving. Both levers working against you simultaneously.

**Response:** Freeze Track 2 spending (SageMaker, Hive API). Cut to Track 1 only. Renegotiate AWS contracts. Do not scale sales until unit economics are fixed.

---

### Red Flag 3: Delivery Derail

**Pattern:** Sprint completion <65% for 3 consecutive sprints **AND** Unplanned work >25% **AND** Any milestone slipping by >2 weeks

**Diagnosis:** Scope creep, technical debt, or team capacity is underwater. The roadmap will miss v1 (W26) which breaks the SOC 2 Type 2 evidence collection window.

**Response:** Emergency scope triage. Cut everything that is not on the milestone gate checklist. SOC 2 Type 2 depends on evidence collection starting by W26 — a 4-week delay here causes a 4-week delay in the audit, potentially missing the v2 window entirely.

---

### Red Flag 4: ML Accuracy Plateau

**Pattern:** Prompt injection TPR stuck at 83–85% after 3 evaluation cycles **AND** FPR creeping up toward 2%

**Diagnosis:** Lakera Guard is not improving, and you don't have enough labeled data for BERT yet.

**Response:** Activate the v1.5 playbook: add regex pre-filter layer (OWASP + custom rules, no new model required). This lifts combined TPR to ~92–93% without S23–S24 BERT work. Do not attempt BERT fine-tuning prematurely — you need ≥50K labeled samples for it to be meaningful.

---

### Red Flag 5: Compliance Evidence Gap

**Pattern:** Vanta evidence coverage <80% at W34 (8 weeks before v1.5)

**Diagnosis:** SOC 2 Type 2 audit evidence window is filling with gaps. Auditors will challenge control effectiveness if evidence is sparse.

**Response:** DevSecOps engineer focuses 100% on Vanta evidence hooks for 2 sprints. Defer all feature work for that engineer. Evidence gaps close faster than they open — 2 focused sprints can cover 6 weeks of gaps.

---

## 9. Monthly Review Checklist

Use this checklist at the first Monday of every month.

### Business Health (PM + CPO, 30 min)
- [ ] MRR vs plan — on track / behind / ahead?
- [ ] Active customers count vs target — W12: 3+ pilots · W26: 5+ paying · W38: 10+ paying · W52: 50
- [ ] Churn: any customers left this month? Root cause documented?
- [ ] Funnel: outreach → qualified → demo → pilot → paid — which stage is the bottleneck?
- [ ] NPS collected from all customers who have been live >30 days

### Technical Health (Tech Lead, 30 min)
- [ ] All CI tests green? Tenant isolation test specifically?
- [ ] Offboarding p95 latency within 5-min target?
- [ ] AWS cost vs budget (tag-based, Cost Explorer)? Any surprise line items?
- [ ] SAST/DAST scan results — any Critical/High open?
- [ ] Lakera Guard and Hive API invoices — usage within model? Any overage?

### ML Health (ML Engineer #1, 30 min — starts Month 4)
- [ ] 500-sample accuracy re-evaluation completed?
- [ ] Prompt injection TPR / FPR vs gates — any degradation?
- [ ] Shadow AI classification — new AI tools to add to registry this month?
- [ ] Model drift alerts from SageMaker — any firing?
- [ ] Deepfake test: new adversarial sample types detected by Hive?

### Delivery Health (PM, 15 min)
- [ ] Current sprint completion rate — on track to hit gate?
- [ ] Next milestone gate checklist — what items are at risk?
- [ ] Hiring plan — any critical roles behind schedule?
- [ ] External dependencies (auditor engagement, Chrome Web Store review, Vanta onboarding) — all on track?

---

## 10. How These Numbers Connect

The metrics in this document form a dependency chain. Understanding the chain lets you predict cascading failures before they happen.

```
Offboarding <5 min (S1)
    ↓ enables
Pilot → Paid Conversion >60% (S4)
    ↓ drives
MRR growth + ARR targets (Tier 5)
    ↓ funds
Team headcount ramp (Tier 6)
    ↓ enables
Track 2 ML R&D (Tier 3)
    ↓ which must pass
Accuracy Gates 1–4 (Tier 3)
    ↓ to enable
Track 2 features graduate from beta → NRR >110% (S3)
    ↓ which funds
SOC 2 Type 2 audit engagement (Tier 2, v2 Gate)
    ↓ which unlocks
Enterprise tier sales → ARR target $480K (Tier 5)
```

**The single most fragile link:** Pilot → Paid Conversion. Everything above it is engineering (controllable). Everything below it is compounded revenue. If conversion is weak, the chain breaks before it starts.

**The single most time-sensitive constraint:** SOC 2 Type 2 evidence collection must start by **W26**. It requires 6 continuous months of evidence. A 4-week slip in v1 = a 4-week slip in the v2 certification date = enterprise deals that were counting on your SOC 2 certification in Month 12 close 4 weeks late or not at all.

---

*This document is a living scorecard. Update thresholds if the business model or architecture changes materially. Last updated: 2026-05-28.*
