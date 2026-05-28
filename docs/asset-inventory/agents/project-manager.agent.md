---
name: asset-inventory-project-manager
description: "Project Manager for Asset Inventory & Classification (Requirement 1). Extends base project-manager agent with specialized context for Track 1 Sprints 2-3 timeline, capacity allocation, and integration complexity."
extends: project-manager
tools: Read, Glob, Grep, WebSearch
---

**Base Agent**: This agent extends [project-manager](../../../.github/agents/project-manager.agent.md). Follow all base agent instructions, debate structure, and frameworks. This file adds requirement-specific context only.

---

## Requirement 1: Asset Inventory & Classification

### Scope
- **Sprint 2 (W3-4)**: Asset discovery engine for Google Workspace, M365, Slack, AWS
- **Sprint 3 (W5-6)**: Classification framework (criticality, sensitivity), dependency mapping
- **Team allocation**: 2 Backend Engineers (100%), 1 Frontend Engineer (80%)
- **Dependencies**: Sprint 1 (infrastructure, auth, integrations) must complete first

### Sprint 2: Asset Discovery Engine (W3-4)

**Scope:**
- Asset discovery for 4 providers (Google, M365, Slack, AWS)
- Incremental sync every 15 minutes
- Database schema for asset metadata
- API integration with rate limit handling
- Web Dashboard: Asset inventory table view

**Team allocation:**
- Backend Eng 1: Google + M365 integrations (10 days)
- Backend Eng 2: Slack + AWS integrations (10 days)
- Frontend Eng: Asset inventory dashboard (8 days, 80%)

**Capacity analysis:**
- Total capacity required: 28 days
- Total capacity available: 2 × 7 + 1 × 5.6 = 19.6 days
- Utilization: 28 / 19.6 = **143%** 🔴 **OVER-ALLOCATED**

**Risk assessment:**
- 🔴 **CRITICAL**: Over-allocated by 43% — unrealistic timeline
- 🟡 **HIGH**: 4 provider integrations in 2 weeks is aggressive
- 🟡 **HIGH**: No buffer for API rate limit issues or vendor API changes
- 🟢 **MEDIUM**: Dependency on Sprint 1 (infrastructure) completing on time

**Recommendation**: 
- **Option 1**: Reduce scope to 2 providers (Google + M365) in Sprint 2, defer Slack + AWS to Sprint 3
- **Option 2**: Extend Sprint 2 to 3 weeks (W3-5) to accommodate 4 providers
- **Option 3**: Add 1 Backend Engineer to team (increase from 2 to 3)
- **PREFERRED**: Option 1 (reduce scope)

### Sprint 3: Classification & Dependency Mapping (W5-6)

**Scope:**
- Classification framework (criticality levels)
- Dependency mapping (User→App→Resource)
- Dashboard: Classification UI, dependency list view
- Compliance reports: Asset inventory CSV export

**Team allocation:**
- Backend Eng 1: Classification engine (7 days)
- Backend Eng 2: Dependency mapping (7 days)
- Frontend Eng: Classification UI + dependency view (7 days, 100%)

**Capacity analysis:**
- Total capacity required: 21 days
- Total capacity available: 2 × 7 + 1 × 7 = 21 days
- Utilization: 21 / 21 = **100%** 🟡 **NO BUFFER**

**Risk assessment:**
- 🟡 **HIGH**: 100% utilization leaves no buffer for bugs or unknowns
- 🟡 **HIGH**: Dependency mapping complexity may be underestimated
- 🟢 **MEDIUM**: Classification framework is straightforward (rule-based)

**Recommendation**:
- Add 1-2 days buffer by deferring dependency graph visualization to v2 (provide list view only in v1)

### Critical Path Analysis

**Critical path**: Sprint 1 → Sprint 2 → Sprint 3 → Sprint 6 (Access Governance)

- Sprint 2 is on critical path: Access Governance (Sprint 6) depends on asset inventory
- Any delay in Sprint 2 delays Sprint 6 (automated offboarding)
- Sprint 3 is NOT on critical path: Dependency mapping is nice-to-have, not blocking

**Slack**: 0 days (no buffer between Sprint 2 and Sprint 6)

### Resource Bottlenecks

**Backend Engineers (2 FTE):**
- Sprint 2: 143% utilization 🔴 **OVER-ALLOCATED**
- Sprint 3: 100% utilization 🟡 **NO BUFFER**
- Sprint 4-5: 80% utilization ✅ **COMFORTABLE**

**Frontend Engineer (1 FTE):**
- Sprint 2: 80% utilization ✅ **COMFORTABLE**
- Sprint 3: 100% utilization 🟡 **NO BUFFER**
- Sprint 4-5: 60% utilization ✅ **COMFORTABLE**

**Bottleneck**: Backend Engineers in Sprint 2-3 are over-allocated.

### External Dependencies

**Vendor APIs:**
- Google Admin SDK: Stable, well-documented, 99.9% uptime
- Microsoft Graph API: Stable, but throttling issues on large tenants
- Slack Admin API: Requires Enterprise Grid (customer must have this)
- AWS Config: Customer must enable Config (additional cost ~$2/region/month)

**Lead times:**
- OAuth app registration: 1-2 days (Google, M365, Slack)
- AWS cross-account IAM role setup: 1-2 days (customer must configure)

**Risk**: If customer doesn't have Slack Enterprise Grid or AWS Config enabled, asset discovery will be incomplete.

### Risk Register

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| Sprint 2 over-allocated (143%) | High (90%) | High (2-week delay) | Reduce scope to 2 providers or extend sprint to 3 weeks |
| API rate limits hit during sync | Medium (50%) | Medium (1-week delay) | Implement exponential backoff + caching in Sprint 2 |
| Dependency mapping complexity underestimated | Medium (40%) | Medium (1-week delay) | Defer graph visualization to v2, provide list view in v1 |
| Customer doesn't have Slack Enterprise Grid | Medium (30%) | Low (feature incomplete) | Document requirement in onboarding, offer workaround |
| AWS Config not enabled | Medium (30%) | Low (feature incomplete) | Document requirement in onboarding, offer setup guide |

### Recommendations

1. **Reduce Sprint 2 scope**: Focus on Google + M365 in Sprint 2, defer Slack + AWS to Sprint 3
2. **Add buffer to Sprint 3**: Defer dependency graph visualization to v2, provide list view only in v1
3. **Document external dependencies**: Create onboarding checklist for Slack Enterprise Grid and AWS Config requirements
4. **Monitor API rate limits**: Implement rate limit monitoring in Sprint 2 to catch issues early
