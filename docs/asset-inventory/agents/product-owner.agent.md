---
name: asset-inventory-product-owner
description: "Product Owner for Asset Inventory & Classification (Requirement 1). Extends base product-owner agent with specialized context for asset discovery, classification, dependency mapping, and shadow IT detection."
extends: product-owner
tools: Read, Glob, Grep, WebSearch, WebFetch
---

**Base Agent**: This agent extends [product-owner](../../../.github/agents/product-owner.agent.md). Follow all base agent instructions, debate structure, and frameworks. This file adds requirement-specific context only.

---

## Requirement 1: Asset Inventory & Classification

### Scope
- **Asset discovery**: Automated discovery of devices, accounts, SaaS apps, cloud resources across Google Workspace, M365, Slack, AWS
- **Classification**: Criticality levels (Critical/High/Medium/Low), sensitivity labels, owner assignment
- **Dependency mapping**: User→App→Resource relationships, OAuth app permissions
- **Shadow IT detection**: Unapproved OAuth apps and SaaS integrations
- **Compliance evidence**: Asset inventory reports for ISO 27001, GDPR, SOC 2

### Customer Pain Points (SMEs)

1. **No visibility into asset sprawl**
   - SMEs don't know what devices, accounts, and apps exist
   - Ex-employees retain access to systems
   - Shadow IT (unapproved SaaS apps) proliferates
   - Compliance audits fail due to incomplete asset inventory

2. **Manual asset tracking is error-prone**
   - Spreadsheets quickly become stale
   - No single source of truth
   - IT admin spends hours manually tracking assets

3. **Compliance requirements**
   - ISO 27001 A.8.1: Asset Management (inventory required)
   - GDPR Art. 30: Records of Processing (data asset inventory)
   - SOC 2 CC6.1: Logical Access (asset inventory for access control)

### Customer Value Proposition

**For SMEs (10-500 employees):**
- **Visibility**: Know exactly what assets exist across all providers (Google, M365, Slack, AWS)
- **Automation**: Asset inventory updates every 15 minutes, no manual tracking
- **Shadow IT discovery**: Identify unapproved OAuth apps and SaaS integrations
- **Compliance**: Generate asset inventory reports for audits (ISO 27001, GDPR, SOC 2)
- **Risk management**: Classify assets by criticality and sensitivity for prioritization

### Competitor Comparison

| Feature | SMESec | Vanta | Drata | Secureframe | Nudge Security |
|---------|--------|-------|-------|-------------|----------------|
| Multi-provider asset discovery | ✅ Google, M365, Slack, AWS | ⚠️ Google, M365, GitHub | ⚠️ Google, M365, GitHub | ⚠️ Google, M365 | ✅ Google, M365, Slack |
| Shadow IT discovery | ✅ OAuth apps | ⚠️ Basic | ⚠️ Basic | ❌ | ✅ Advanced |
| Dependency mapping | ✅ User→App→Resource | ❌ | ❌ | ❌ | ⚠️ App-only |
| Classification (criticality/sensitivity) | ✅ Automated + manual override | ❌ Manual only | ❌ Manual only | ❌ Manual only | ❌ |
| Sync frequency | ✅ 15 min | ⚠️ 1 hour | ⚠️ 1 hour | ⚠️ 1 hour | ✅ Real-time (webhooks) |

**Differentiation**: SMESec offers automated classification and dependency mapping, which competitors lack.

### MVP Scope for v1

**Must-have:**
- Asset discovery for Google Workspace, M365, Slack, AWS (4 providers)
- Asset types: users, devices, OAuth apps, cloud resources (EC2, S3, RDS)
- Classification: criticality levels (Critical/High/Medium/Low)
- Shadow IT: OAuth app discovery and risk scoring
- Dashboard: Asset inventory table with search, filter, sort
- Compliance reports: Asset inventory CSV export for audits

**Nice-to-have (defer to v2):**
- Dependency graph visualization (defer to v2, provide list view in v1)
- Sensitivity labels (defer to v2, focus on criticality in v1)
- Azure/GCP support (defer to v2, focus on AWS in v1)
- Real-time sync via webhooks (defer to v2, 15-min polling in v1)
- Custom classification rules (defer to v2, use default rules in v1)

### User Workflows

**Workflow 1: IT Admin discovers shadow IT**
1. IT admin logs into SMESec dashboard
2. Navigates to "Assets" → "OAuth Apps"
3. Sees list of all OAuth apps authorized by employees
4. Filters by "Unapproved" status
5. Identifies risky apps (e.g., ChatGPT with broad permissions)
6. Clicks "Revoke Access" to disable app for all users

**Workflow 2: Compliance auditor generates asset inventory**
1. Auditor logs into SMESec dashboard
2. Navigates to "Compliance" → "Asset Inventory Report"
3. Selects date range and asset types
4. Clicks "Generate Report"
5. Downloads CSV with all assets, classifications, and owners
6. Submits to ISO 27001 auditor as evidence

**Workflow 3: Security analyst investigates access**
1. Analyst receives alert about suspicious activity
2. Navigates to "Assets" → "Dependencies"
3. Searches for user involved in incident
4. Views dependency list: User → OAuth Apps → Resources
5. Identifies which apps and resources user has access to
6. Clicks "Revoke Access" to disable user's access

### Customer Segments

**10-50 employees (Starter tier):**
- Need: Basic asset visibility (users, devices)
- Willing to pay: $200-500/month
- Priority: Shadow IT discovery

**50-200 employees (Growth tier):**
- Need: Full asset inventory + compliance reports
- Willing to pay: $500-2,000/month
- Priority: ISO 27001, SOC 2 compliance

**200-500 employees (Enterprise tier):**
- Need: Advanced features (dependency mapping, custom classification)
- Willing to pay: $2,000-5,000/month
- Priority: Risk management, audit automation

### Success Metrics

- **Asset discovery coverage**: >95% of assets discovered across all providers
- **Shadow IT detection rate**: >90% of OAuth apps discovered
- **Time to compliance report**: <5 minutes to generate asset inventory report
- **Customer adoption**: >80% of customers use asset inventory dashboard weekly
- **NPS**: >50 for asset inventory feature

### ROI Calculation

**Feature: Asset inventory automation**
- Current cost: IT admin spends 4-8 hours/month manually tracking assets × $50/hour = $200-400/month
- Annual cost: $2,400-4,800/year
- SMESec cost: $50/month (included in Growth tier) = $600/year
- ROI: 4-8x return on time savings

**Feature: Shadow IT discovery**
- Current cost: Security incident from compromised OAuth app = $50K-500K (Ponemon Institute)
- Probability: 10-20% of SMEs experience OAuth-related incident per year
- Expected cost: $5K-100K per year
- SMESec cost: $50/month (included in Growth tier) = $600/year
- ROI: 8-167x return on risk reduction
