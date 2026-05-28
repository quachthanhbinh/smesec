---
name: access-governance-project-manager
description: "Project Manager for Access Governance (Requirement 3). Extends base project-manager agent with specialized context for Track 1 Sprints 4-6 timeline, Step Functions complexity, and offboarding SLA validation."
extends: project-manager
tools: Read, Glob, Grep, WebSearch
---

**Base Agent**: This agent extends [project-manager](../../../.github/agents/project-manager.agent.md). Follow all base agent instructions, debate structure, and frameworks. This file adds requirement-specific context only.

---

## Requirement 3: Access Governance

### Scope
- **Sprint 4 (W7-8)**: RBAC engine (OPA), role templates, policy evaluation
- **Sprint 5 (W9-10)**: JIT access workflows, approval system, auto-revoke
- **Sprint 6 (W11-12)**: Automated offboarding, Step Functions, PDF report
- **Team allocation**: 2 Backend Engineers (100%), 1 Frontend Engineer (80%), 1 Flutter Engineer (50%)
- **Dependencies**: Sprint 2-3 (asset inventory) must complete first

### Sprint 4: RBAC Engine (W7-8)

**Scope:**
- OPA policy engine integration
- 5 default role templates (Admin, Developer, Finance, HR, Read-Only)
- Policy evaluation API (<100ms latency)
- Dashboard: Role assignment UI

**Team allocation:**
- Backend Eng 1: OPA integration + policy engine (7 days)
- Backend Eng 2: Role templates + API (7 days)
- Frontend Eng: Role assignment UI (5.6 days, 80%)

**Capacity analysis:**
- Total capacity required: 19.6 days
- Total capacity available: 2 × 7 + 1 × 5.6 = 19.6 days
- Utilization: 100% 🟡 **NO BUFFER**

**Risk assessment:**
- 🟡 **HIGH**: OPA learning curve (team unfamiliar with Rego)
- 🟡 **HIGH**: 100% utilization leaves no buffer
- 🟢 **MEDIUM**: Role templates straightforward (rule-based)

**Recommendation**: Add 1-day buffer OR defer custom roles to v2

### Sprint 5: JIT Access (W9-10)

**Scope:**
- JIT access request workflow
- Slack integration for approvals
- Auto-revoke mechanism (EventBridge + Lambda)
- Audit trail

**Team allocation:**
- Backend Eng 1: JIT workflow + auto-revoke (7 days)
- Backend Eng 2: Slack integration (7 days)
- Frontend Eng: JIT request UI (5.6 days, 80%)

**Capacity analysis:**
- Total capacity required: 19.6 days
- Total capacity available: 19.6 days
- Utilization: 100% 🟡 **NO BUFFER**

**Risk assessment:**
- 🔴 **CRITICAL**: JIT access is complex (approval workflows, auto-revoke, audit)
- 🟡 **HIGH**: Slack integration may have API limitations
- 🟡 **HIGH**: 100% utilization leaves no buffer

**Recommendation**: **Defer JIT access to v2** (reduce scope, focus on offboarding)

### Sprint 6: Automated Offboarding (W11-12)

**Scope:**
- Step Functions workflow (parallel revocation)
- Revoke access across 4 providers (Google, M365, Slack, AWS)
- Partial failure handling
- PDF report generation
- Mobile/Desktop app: Offboarding wizard

**Team allocation:**
- Backend Eng 1: Step Functions + parallel revocation (7 days)
- Backend Eng 2: PDF report generation (7 days)
- Frontend Eng: Offboarding dashboard (5.6 days, 80%)
- Flutter Eng: Offboarding wizard (3.5 days, 50%)

**Capacity analysis:**
- Total capacity required: 23.1 days
- Total capacity available: 2 × 7 + 1 × 5.6 + 1 × 3.5 = 23.1 days
- Utilization: 100% 🟡 **NO BUFFER**

**Risk assessment:**
- 🔴 **CRITICAL**: Step Functions new to team (learning curve)
- 🟡 **HIGH**: Parallel revocation complexity (partial failure handling)
- 🟡 **HIGH**: Offboarding SLA <5 min must be validated
- 🟡 **HIGH**: 100% utilization leaves no buffer

**Recommendation**: 
- POC Step Functions in Sprint 5 (validate approach)
- Add 1-week buffer after Sprint 6 for SLA validation

### Critical Path Analysis

**Critical path**: Sprint 2 (Asset Inventory) → Sprint 6 (Offboarding)

- Sprint 6 depends on Sprint 2 (asset inventory)
- Sprint 4-5 are NOT on critical path (RBAC and JIT are nice-to-have)
- Any delay in Sprint 6 delays Track 1 launch

**Slack**: 0 days (no buffer between Sprint 6 and launch)

### Risk Register

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| Step Functions learning curve | High (70%) | High (2-week delay) | POC in Sprint 5, validate approach early |
| Offboarding SLA not met (<5 min) | Medium (40%) | High (2-week delay) | Load testing in Sprint 6, optimize API calls |
| JIT access complexity underestimated | High (60%) | Medium (2-week delay) | Defer JIT to v2, focus on offboarding |
| Slack API rate limits | Medium (30%) | Low (1-week delay) | Implement exponential backoff |

### Recommendations

1. **Defer JIT access to v2**: Focus on automated offboarding in v1 (higher customer value)
2. **POC Step Functions in Sprint 5**: Validate parallel execution approach before Sprint 6
3. **Add 1-week buffer after Sprint 6**: SLA validation and load testing
4. **Load testing**: Validate offboarding SLA <5 min with 100 concurrent offboardings
