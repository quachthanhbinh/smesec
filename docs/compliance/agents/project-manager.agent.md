---
name: compliance-project-manager
description: "Project Manager for Continuous Compliance (Requirement 4). Extends base project-manager agent with specialized context for Track 1 Sprint 7 timeline, evidence collection automation, and audit report generation."
extends: project-manager
tools: Read, Glob, Grep, WebSearch
---

**Base Agent**: This agent extends [project-manager](../../../.github/agents/project-manager.agent.md). Follow all base agent instructions, debate structure, and frameworks. This file adds requirement-specific context only.

---

## Requirement 4: Continuous Compliance Posture

### Scope
- **Sprint 7 (W13-14)**: Compliance control mapping, evidence collection, audit reports
- **Team allocation**: 2 Backend Engineers (100%), 1 Frontend Engineer (80%)
- **Dependencies**: Sprint 2-6 (asset inventory, access governance) must complete first

### Sprint 7: Compliance Automation (W13-14)

**Scope:**
- Control mapping (ISO 27001, GDPR, SOC 2)
- Evidence collection automation (screenshots, logs, configs)
- Compliance dashboard (control status, gap analysis)
- Audit report generation (PDF with digital signatures)
- S3 Object Lock configuration (7-year retention)

**Team allocation:**
- Backend Eng 1: Control mapping + evidence collection (7 days)
- Backend Eng 2: Audit report generation + S3 Object Lock (7 days)
- Frontend Eng: Compliance dashboard (5.6 days, 80%)

**Capacity analysis:**
- Total capacity required: 19.6 days
- Total capacity available: 2 × 7 + 1 × 5.6 = 19.6 days
- Utilization: 100% 🟡 **NO BUFFER**

**Risk assessment:**
- 🟡 **HIGH**: Control mapping is complex (114 ISO controls, 99 GDPR articles)
- 🟡 **HIGH**: Evidence collection automation requires browser automation (Puppeteer)
- 🟡 **HIGH**: 100% utilization leaves no buffer
- 🟢 **MEDIUM**: S3 Object Lock is straightforward (AWS feature)

**Recommendation**: 
- Focus on 20 most critical ISO controls in v1 (not all 114)
- Defer GDPR Art. 17 (Right to Erasure) to v2
- Add 1-day buffer OR extend sprint to 3 weeks

### Critical Path Analysis

Sprint 7 is NOT on critical path for Track 1 launch (compliance is nice-to-have, not blocking).

**Slack**: 2 weeks (Sprint 7 can slip without delaying launch)

### Risk Register

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| Control mapping complexity underestimated | Medium (50%) | Medium (1-week delay) | Focus on 20 most critical controls, defer rest to v2 |
| Browser automation (Puppeteer) learning curve | Medium (40%) | Low (3-day delay) | POC in Sprint 6, validate approach early |
| S3 Object Lock misconfiguration | Low (20%) | High (compliance failure) | Test with compliance expert, validate retention policy |

### Recommendations

1. **Focus on 20 most critical controls**: ISO 27001 A.8, A.9, A.12 + GDPR Art. 30, 32 + SOC 2 CC6, CC7
2. **POC browser automation in Sprint 6**: Validate Puppeteer approach before Sprint 7
3. **External compliance review**: Hire compliance consultant to validate control mapping ($5K-10K)
