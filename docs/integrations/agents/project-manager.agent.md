---
name: integrations-project-manager
description: "Project Manager for Integration Requirements (Requirement 7). Extends base project-manager agent with specialized context for Track 1 Sprint 1 timeline, OAuth setup complexity, and integration testing."
extends: project-manager
tools: Read, Glob, Grep, WebSearch
---

**Base Agent**: This agent extends [project-manager](../../../.github/agents/project-manager.agent.md). Follow all base agent instructions, debate structure, and frameworks. This file adds requirement-specific context only.

---

## Requirement 7: Integration Requirements

### Scope
- **Sprint 1 (W1-2)**: OAuth 2.0 setup, integration skeletons for 4 providers
- **Team allocation**: 2 Backend Engineers (100%), 1 Frontend Engineer (80%)
- **Dependencies**: None (Sprint 1 is foundation for all other sprints)

### Sprint 1: Integration Foundation (W1-2)

**Scope:**
- OAuth 2.0 setup wizard (Google, M365, Slack, AWS)
- Integration skeletons (API clients, token management)
- Token storage (AWS Secrets Manager)
- Integration health dashboard
- Rate limit handling framework

**Team allocation:**
- Backend Eng 1: Google + M365 OAuth setup (7 days)
- Backend Eng 2: Slack + AWS OAuth setup (7 days)
- Frontend Eng: OAuth wizard UI + health dashboard (5.6 days, 80%)

**Capacity analysis:**
- Total capacity required: 19.6 days
- Total capacity available: 2 × 7 + 1 × 5.6 = 19.6 days
- Utilization: 100% 🟡 **NO BUFFER**

**Risk assessment:**
- 🟡 **HIGH**: OAuth 2.0 complexity (different flows per provider)
- 🟡 **HIGH**: AWS cross-account IAM role setup (customer must configure)
- 🟡 **HIGH**: 100% utilization leaves no buffer
- 🟢 **MEDIUM**: Integration skeletons are straightforward (API clients)

**Recommendation**: 
- Focus on OAuth setup in Sprint 1 (integration skeletons can be minimal)
- Add 1-day buffer OR extend sprint to 3 weeks
- Document OAuth setup for customers (Google domain-wide delegation, AWS cross-account role)

### External Dependencies

**Customer setup required:**
- Google Workspace: Domain-wide delegation (IT admin must configure)
- Microsoft 365: Azure AD app registration (IT admin must configure)
- Slack: Enterprise Grid required (customer must have this plan)
- AWS: Cross-account IAM role (customer must create role)

**Lead times:**
- OAuth app registration: 1-2 days (Google, M365, Slack)
- AWS cross-account IAM role setup: 1-2 days (customer must configure)

**Risk**: If customer doesn't complete setup, integration will fail.

### Risk Register

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| OAuth setup complexity underestimated | High (60%) | Medium (1-week delay) | Focus on OAuth in Sprint 1, defer integration skeletons to Sprint 2 |
| Customer doesn't complete OAuth setup | Medium (40%) | Low (integration incomplete) | Document setup steps, offer setup assistance |
| Slack Enterprise Grid not available | Medium (30%) | Low (Slack integration unavailable) | Document requirement, offer workaround (manual Slack integration) |

### Recommendations

1. **Focus on OAuth setup in Sprint 1**: Integration skeletons can be minimal (just API clients)
2. **Document customer setup steps**: Google domain-wide delegation, AWS cross-account role
3. **Add 1-day buffer**: OAuth testing and troubleshooting
