# Decision Record: SMESec Pricing Model

**Date:** 2026-05-28  
**Status:** Proposed  
**Deciders:** Product Team  
**Context:** Key Requirement #5 from topic.md - Cost model design for SME market

---

## Context and Problem Statement

SMESec platform cần một pricing model vừa đảm bảo tính kinh tế cho SMEs (10-500 nhân viên), vừa cover được chi phí infrastructure và vận hành, đồng thời duy trì gross margin khỏe mạnh (>70%) để sustainable.

**Challenges:**
- SMEs có budget hạn chế, không thể chi $50K+/năm như enterprise
- Platform có 2 tracks: Track 1 (deterministic) + Track 2 (AI/ML) với chi phí khác nhau
- Competitors (Drata, Vanta, Secureframe) chỉ cover compliance, không có AI threat detection
- Infrastructure costs cao ban đầu (SageMaker, deepfake APIs, multi-tenant architecture)
- Cần balance giữa predictability (cho customers) và flexibility (cho growth)

---

## Decision Drivers

1. **Market competitiveness:** Pricing phải nằm trong range $60-150/employee/year (competitor benchmark)
2. **Gross margin target:** >70% để sustainable cho SaaS business
3. **Customer predictability:** SMEs cần biết trước chi phí hàng tháng
4. **Scalability:** Model phải work từ 10 employees đến 500 employees
5. **Value differentiation:** Track 2 (AI detection) là unique value prop so với competitors

---

## Options Considered

### Option A: Per-Employee Tiered Pricing

**Structure:**
- Starter (10-50 emp): $499/month
- Growth (51-150 emp): $999/month  
- Business (151-300 emp): $1,799/month
- Enterprise (301-500 emp): $2,999/month

**Pros:**
- ✅ Simple, predictable cho customers
- ✅ Aligns với competitor models
- ✅ Gross margin 63-81%

**Cons:**
- ❌ Không flexible cho customers có usage thấp
- ❌ Hard ceiling tại mỗi tier → friction khi scale
- ❌ Không incentivize adoption của Track 2 features

**Gross Margin:** 63-81% depending on tier

---

### Option B: Usage-Based Pricing

**Structure:**
- Base platform: $299/month
- Per user: $8/month
- Per 100 assets: $50/month
- Per 1K AI prompts: $10/month
- Per deepfake check: $0.01

**Pros:**
- ✅ True pay-as-you-grow
- ✅ High gross margin (~75%)
- ✅ Aligns costs với value delivered

**Cons:**
- ❌ Unpredictable bills → SME CFOs hate this
- ❌ Complex to explain và forecast
- ❌ Risk of bill shock → churn

**Gross Margin:** ~75% but high variance

---

### Option C: Hybrid Model ✅ **RECOMMENDED**

**Structure:**
- Base price với included usage
- Overage charges cho usage vượt quota
- Annual discount 15%

| Tier | Monthly | Included | Overage |
|------|---------|----------|---------|
| **Starter** | $399 | 50 users, 5K prompts, 100 deepfake checks | $5/user, $5/1K prompts, $0.01/check |
| **Growth** | $799 | 150 users, 20K prompts, 500 checks | $4/user, $4/1K prompts, $0.008/check |
| **Business** | $1,499 | 300 users, 50K prompts, 2K checks | $3/user, $3/1K prompts, $0.005/check |
| **Enterprise** | Custom | Unlimited | N/A |

**Pros:**
- ✅ Predictable base price (SME-friendly)
- ✅ Flexible overage (covers edge cases)
- ✅ Gross margin 70-75%
- ✅ Competitive với Drata/Vanta
- ✅ Incentivizes annual commits (15% discount)

**Cons:**
- ⚠️ Slightly more complex than pure tiered
- ⚠️ Need good usage monitoring/alerting

**Gross Margin:** 70-75% sustainable

---

## Decision Outcome

**Chosen option:** **Option C - Hybrid Model**

### Rationale

1. **Market fit:** Growth tier ($799/mo) competitive với Drata ($750-2,500/mo) nhưng includes AI detection
2. **Customer experience:** Base price predictable, overage charges rare (quotas generous)
3. **Financial health:** 70-75% gross margin sustainable, path to 20%+ operating margin tại 50+ customers
4. **Flexibility:** Can adjust base/overage ratio based on customer feedback post-launch
5. **Differentiation:** Track 2 (AI detection) included in base price → clear value prop vs competitors

### Implementation Details

**Pricing tiers (final):**
- **Starter:** $399/month ($339/month annual) - Target: 10-50 employees
- **Growth:** $799/month ($680/month annual) - Target: 51-150 employees  
- **Business:** $1,499/month ($1,276/month annual) - Target: 151-300 employees
- **Enterprise:** Custom pricing - Target: 301-500 employees

**Free trial:** 14 days, no credit card, up to 10 users

**Add-ons:**
- Premium support (1-hour SLA): +$500/month
- Custom playbooks: $2,500 one-time
- Dedicated CSM: +$1,000/month (Enterprise only)

---

## Cost Structure (Supporting Data)

### Infrastructure COGS Per Tenant

| Component | Monthly Cost | Optimization Potential |
|-----------|-------------|----------------------|
| Track 1 (Foundation) | $212 | → $160 (Reserved instances, Savings Plans) |
| Track 2 (AI Detection) | $53 | → $40 (Serverless inference, volume discounts) |
| Shared (allocated) | $19 | → $15 (Scale efficiency) |
| **Total COGS** | **$284** | **→ $215 (24% reduction)** |

**Target COGS at scale (100+ tenants):** $120/tenant/month

### Break-Even Analysis

**Growth tier ($799/month):**
- Revenue: $799
- COGS (optimized): $245
- Gross profit: $554 (69%)
- Operating expenses (S&M 40%, R&D 20%, G&A 10%): $560
- **Break-even:** ~18 customers = $14K MRR

**Healthy SaaS target:** 50+ customers = $40K MRR = 25%+ operating margin

---

## Consequences

### Positive

- **Competitive positioning:** Undercuts Drata/Vanta while offering more (AI detection)
- **Customer acquisition:** Predictable pricing removes friction in sales cycle
- **Gross margin:** 70-75% sustainable, industry-standard for SaaS
- **Scalability:** Model works từ 10 employees đến 500 employees
- **Flexibility:** Can adjust quotas/overage based on real usage patterns

### Negative

- **Complexity:** Need robust usage tracking và billing system
- **Support burden:** Customers may need help understanding overage charges
- **Competitive response:** Drata/Vanta có thể lower prices hoặc add AI features
- **Early-stage margin:** Low gross margin (17%) trong Year 1 due to small scale

### Risks and Mitigations

| Risk | Mitigation |
|------|-----------|
| Customers exceed quotas frequently → bill shock | Set quotas at 80th percentile usage; alert at 70% quota |
| Competitors lower prices | Emphasize Track 2 differentiation; bundle value |
| COGS higher than projected | Aggressive optimization roadmap (see cost-analysis.md) |
| Low conversion from free trial | Extend trial to 30 days; offer onboarding support |

---

## Validation Plan

### Pre-Launch (Sprint 11-12)
1. **Pilot pricing test:** Offer 2-3 pilot customers Growth tier at $799/month
2. **Usage monitoring:** Track actual usage vs. quotas during pilot
3. **Customer feedback:** Survey pilot customers on pricing clarity và value perception

### Post-Launch (Month 1-3)
1. **Conversion tracking:** Free trial → paid conversion rate (target: >20%)
2. **Overage analysis:** % of customers hitting overage (target: <15%)
3. **Churn analysis:** Price-related churn (target: <5%)
4. **Competitive win/loss:** Track deals lost to Drata/Vanta on price

### Adjustment Triggers

- **If overage rate >25%:** Increase base quotas by 20%
- **If conversion rate <15%:** Consider Starter tier at $299/month
- **If competitors undercut by >30%:** Re-evaluate Track 2 bundling (offer Track 1-only tier)

---

## Related Documents

- [Cost Analysis](../pricing/cost-analysis.md) - Detailed COGS breakdown và financial projections
- [Track 1 Requirements](../track1-foundation/requirements.md) - Foundation features
- [Track 2 Requirements](../track2-ai-detection/requirements.md) - AI detection features
- [2-Track Approach](../strategy/2-track-approach.md) - Strategic overview

---

## Research Sources

- [Secureframe Pricing](https://sprinto.com/blog/secureframe-alternatives/)
- [Drata Pricing Analysis](https://www.complyjet.com/blog/drata-pricing-plans)
- [GRC Platforms Comparison 2026](https://guptadeepak.com/tools/top-5-grc-platforms-2026/)
- [AWS SageMaker Pricing](https://aws.amazon.com/sagemaker/ai/pricing/)
- [Reality Defender Deepfake API](https://www.realitydefender.com/insights/reality-defender-launches-free-access-to-deepfake-detection-api)
- [Sensity AI Pricing](https://indibloghub.com/ai-tools/compare/sensity-ai-vs-murf-ai)

---

## CTO-CPO Debate Results (2026-05-28)

A formal debate was conducted between CTO (technical/financial perspective) and CPO (product/market perspective) to validate the cost model and pricing strategy. Full debate transcript: [2026-05-28-cto-cpo-debate.md](2026-05-28-cto-cpo-debate.md)

### Key Findings

**COGS Revisions:**
- Base COGS (launch): $354/tenant/month (revised from $324)
- Optimized COGS (6 months): $291/tenant/month (revised from $245)
- At-scale COGS (100+ tenants): $255/tenant/month (revised from $180)

**Reasons for revision:**
- Support costs increased to 20% (industry standard vs 15% assumed)
- Added 10% usage buffer for AI usage spikes
- Added $10/tenant hidden costs (onboarding, failed tenants, compliance audits)

**Gross Margin Revisions:**
- Launch (10 tenants): 56% (revised from 64%)
- Scale (50 tenants): 64% (revised from 70%)
- Scale (100+ tenants): 68% (revised from 77%)

**Path to 70% margin requires:**
- 100+ tenants for shared infrastructure efficiency
- Self-hosted deepfake models (eliminate vendor API costs)
- Kubernetes instead of Fargate
- Serverless SageMaker for 50% of tenants

### Risk Scenarios Identified

**Worst-case scenario (all risks combined):**
- Small tenants (avg 30 employees)
- High AI usage (5x prompts, 4x deepfake checks)
- AWS price increase 20%
- **Worst-case COGS: $608/tenant/month**
- **Impact:** Starter tier ($399/month) becomes unprofitable

**Mitigation strategies:**
- Usage quotas with overage charges (already in Hybrid model ✓)
- Minimum pricing tier enforcement
- Reserved Instances lock AWS pricing for 1-3 years

### Consensus Reached

**CTO Final Confidence:** 7/10 (up from 6/10)  
**CPO Final Confidence:** 8/10 (unchanged)

**Decision:** **APPROVED WITH MODIFICATIONS**

**Conditions for approval:**
1. ✅ COGS estimates updated to $354 base, $291 optimized, $255 at scale
2. ✅ Gross margin targets revised to 55-60% at launch, 65-70% at scale
3. ⏳ Pilot customer validation completed before final launch (Sprint 11-12)
4. ⏳ Cost optimization roadmap executed to reach 70% margin at 100+ tenants

**Non-negotiables agreed:**
- Gross margin >50% at launch (minimum for SaaS viability)
- Gross margin >65% at 50+ tenants
- Break-even <25 customers at Growth tier
- Reserved Instances / Savings Plans commitment within 6 months of launch

---

## Approval

- [x] CTO (Technical/Financial) - Approved with modifications (2026-05-28)
- [x] CPO (Product/Market) - Approved with modifications (2026-05-28)
- [ ] Finance/CFO - Pending pilot validation
- [ ] Sales Lead - Pending pilot validation

**Next Steps:**
1. ✅ Update cost-analysis.md with revised COGS ($354/$291/$255)
2. ✅ Update decision record with debate findings
3. ⏳ Implement billing system với usage tracking (Sprint 10-11)
4. ⏳ Create pricing page và sales collateral (Sprint 11)
5. ⏳ Pilot pricing test với 5-10 customers (Sprint 11-12)
6. ⏳ Monitor actual COGS vs projections monthly
7. ⏳ Launch với validated pricing (Sprint 13)
