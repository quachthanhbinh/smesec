---
name: incident-playbooks-project-manager
description: "Project Manager for Incident Playbooks (Requirement 5). Extends base project-manager agent with specialized context for Track 1 Sprint 8 timeline, Step Functions learning curve, and Flutter app development."
extends: project-manager
tools: Read, Glob, Grep, WebSearch
---

**Base Agent**: This agent extends [project-manager](../../../.github/agents/project-manager.agent.md). Follow all base agent instructions, debate structure, and frameworks. This file adds requirement-specific context only.

---

## Requirement 5: Incident Playbooks

### Scope
- **Sprint 8 (W15-16)**: Playbook engine, 4 pre-built playbooks, mobile app
- **Team allocation**: 2 Backend Engineers (100%), 1 Frontend Engineer (80%), 1 Flutter Engineer (100%)
- **Dependencies**: Sprint 6 (offboarding) must complete first (reuse Step Functions patterns)

### Sprint 8: Incident Playbooks (W15-16)

**Scope:**
- Playbook engine (Step Functions state machines)
- 4 pre-built playbooks (account compromise, unauthorized access, shadow IT, offboarding emergency)
- Web UI: Playbook wizard
- Mobile app: Playbook execution on mobile (Flutter)
- Notification system (Slack, email, push)

**Team allocation:**
- Backend Eng 1: Playbook engine + Step Functions (7 days)
- Backend Eng 2: Notification system + audit trail (7 days)
- Frontend Eng: Playbook wizard UI (5.6 days, 80%)
- Flutter Eng: Mobile app (7 days, 100%)

**Capacity analysis:**
- Total capacity required: 26.6 days
- Total capacity available: 2 × 7 + 1 × 5.6 + 1 × 7 = 26.6 days
- Utilization: 100% 🟡 **NO BUFFER**

**Risk assessment:**
- 🟡 **HIGH**: Step Functions already used in Sprint 6 (offboarding), but playbooks are more complex
- 🟡 **HIGH**: Flutter app is new (team unfamiliar with Flutter)
- 🟡 **HIGH**: 100% utilization leaves no buffer
- 🟢 **MEDIUM**: Notification system is straightforward (SNS, SES)

**Recommendation**: 
- Reuse Step Functions patterns from Sprint 6 (reduce learning curve)
- POC Flutter app in Sprint 7 (validate approach)
- Add 1-week buffer after Sprint 8 for mobile app testing

### Risk Register

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| Flutter learning curve | High (60%) | Medium (2-week delay) | POC in Sprint 7, validate approach early |
| Step Functions complexity underestimated | Medium (40%) | Medium (1-week delay) | Reuse patterns from Sprint 6 (offboarding) |
| Mobile app testing insufficient | Medium (40%) | Low (1-week delay) | Add 1-week buffer for mobile testing |

### Recommendations

1. **POC Flutter app in Sprint 7**: Validate Flutter approach before Sprint 8
2. **Reuse Step Functions patterns**: Leverage Sprint 6 offboarding workflows
3. **Add 1-week buffer**: Mobile app testing and bug fixes
