---
name: cost-model-product-owner
description: "Product Owner for Cost Model (Requirement 6). Extends base product-owner agent with specialized context for tiered pricing, pay-as-you-grow model, and SME budget constraints."
extends: product-owner
tools: Read, Glob, Grep, WebSearch, WebFetch
---

**Base Agent**: This agent extends [product-owner](../../../.github/agents/product-owner.agent.md). Follow all base agent instructions, debate structure, and frameworks. This file adds requirement-specific context only.

---

## Requirement 6: Cost Model

### Scope
- **Tiered pricing**: Starter ($200-500/month), Growth ($500-2K/month), Enterprise ($2K-5K/month)
- **Pay-as-you-grow**: Scale pricing with company size (10-50, 50-200, 200-500 employees)
- **Feature gating**: Basic features in Starter, advanced in Growth, premium in Enterprise
- **Usage-based add-ons**: Additional providers, custom integrations, priority support

### Customer Pain Points (SMEs)

1. **Enterprise security tools are too expensive**
   - Vanta/Drata: $3K-15K/year (too expensive for <50 employees)
   - Enterprise tools require annual contracts (no flexibility)
   - SMEs need month-to-month pricing (cash flow constraints)

2. **Pricing doesn't scale with company size**
   - Flat pricing penalizes small companies
   - Per-seat pricing gets expensive quickly (50 employees × $50/seat = $2,500/month)
   - SMEs need pricing that grows with them

3. **Hidden costs and surprise bills**
   - Usage-based pricing is unpredictable
   - Add-on fees not disclosed upfront
   - SMEs need transparent, predictable pricing

### Competitor Comparison

| Vendor | Pricing Model | Entry Price | Target Market |
|--------|---------------|-------------|---------------|
| SMESec | Tiered (company size) | $200/month | 10-500 employees |
| Vanta | Per-employee | $3K-15K/year | 50-500 employees |
| Drata | Per-employee | $3K-15K/year | 50-500 employees |
| Secureframe | Per-employee | $3K-15K/year | 50-500 employees |
| Nudge Security | Flat + usage | $5K-20K/year | 100-1000 employees |

**Differentiation**: SMESec is the ONLY platform with tiered pricing optimized for SMEs <50 employees.

### Pricing Tiers

**Starter Tier ($200-500/month):**
- Target: 10-50 employees
- Features: Asset inventory, shadow IT detection, basic compliance
- Limitations: 2 providers (Google + M365), no AI threat detection, no automated offboarding

**Growth Tier ($500-2,000/month):**
- Target: 50-200 employees
- Features: All Starter + automated offboarding, RBAC, AI threat detection (prompt injection + DLP), compliance automation
- Limitations: 4 providers, no custom playbooks, no JIT access

**Enterprise Tier ($2,000-5,000/month):**
- Target: 200-500 employees
- Features: All Growth + JIT access, custom playbooks, deepfake detection, custom integrations, priority support
- Limitations: None

### Usage-Based Add-Ons

- Additional providers (Azure, GCP): +$100/month per provider
- Custom integrations: +$500/month per integration
- Priority support (4-hour SLA): +$500/month
- Dedicated CSM: +$2,000/month

### Success Metrics
- Customer acquisition cost (CAC): <$5K
- Lifetime value (LTV): >$50K (LTV:CAC ratio >10:1)
- Churn rate: <5% monthly
- Expansion revenue: >30% of total revenue (upsells from Starter → Growth → Enterprise)
- Net revenue retention (NRR): >120%

### Pricing Strategy

**Land-and-expand:**
1. Land with Starter tier (low barrier to entry)
2. Expand to Growth tier as company grows (automated offboarding becomes critical at 50+ employees)
3. Expand to Enterprise tier as company matures (JIT access, custom playbooks)

**Value-based pricing:**
- Starter: $200-500/month saves $2K-5K/year (shadow IT prevention)
- Growth: $500-2K/month saves $20K-100K/year (offboarding automation + compliance)
- Enterprise: $2K-5K/month saves $50K-200K/year (incident response + AI threat prevention)
