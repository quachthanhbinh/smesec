---
name: ai-threat-project-manager
description: "Project Manager for AI-Specific Threat Surface (Requirement 2). Extends base project-manager agent with specialized context for Track 2 Sprints 1-12, ML validation gates, and browser extension development timeline."
extends: project-manager
tools: Read, Glob, Grep, WebSearch
---

**Base Agent**: This agent extends [project-manager](../../../.github/agents/project-manager.agent.md). Follow all base agent instructions, debate structure, and frameworks. This file adds requirement-specific context only.

---

## Requirement 2: AI-Specific Threat Surface

### Scope
- **Track 2 (Sprints 1-12)**: AI threat detection R&D and validation
- **Team allocation**: 1 ML Engineer, 1 Backend Engineer, 1 Frontend Engineer (3 FTE)
- **Dependencies**: Track 1 (asset inventory, access governance) for incident response integration
- **Validation gates**: 4 gates at Week 6, 12, 18, 24 to validate ML accuracy

### Key Sprint Milestones

**Sprints 1-2 (W1-4): Research & Prototyping**
- Utilization: 100% ⚠️ NO BUFFER
- Risk: Dataset quality unknown
- Recommendation: Acceptable — research phase can slip without impacting Track 1

**Sprints 3-4 (W5-8): Core Detection Engine**
- Utilization: 100% ⚠️ NO BUFFER
- Validation Gate 1 (Week 6): Prompt injection precision >90%
- Risk: Browser extension complexity (Chrome MV3 restrictions)
- Recommendation: Add 1 week buffer after Sprint 4

**Sprint 5 (W9-10): Deepfake & Advanced Features**
- Utilization: 100% ⚠️ NO BUFFER
- Validation Gate 2 (Week 12): DLP false negative <1%
- Risk: 🔴 CRITICAL — Deepfake detection complex, vendor API may not meet requirements
- Recommendation: **Defer deepfake to v2** (reduce scope)

**Sprints 6-12 (W11-24): Validation & Tuning**
- Utilization: 100% ⚠️ NO BUFFER
- Validation Gate 3 (Week 18): Deepfake >85% (if not deferred)
- Validation Gate 4 (Week 24): Pilot validation — >95% precision, <5% false positive
- Risk: 🔴 CRITICAL — Pilot customers not identified yet, need 2-3 SMEs by Week 11
- Recommendation: Identify pilot customers by Week 8, add 2-week buffer after Sprint 12

### Critical Path Analysis

Track 2 is NOT on critical path for Track 1 launch (Track 1 can launch independently).

**Decision Gate (Week 24)**:
- If Track 2 meets criteria → Merge into main product
- If Track 2 needs more work → Continue as beta feature

### Risk Register

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| ML accuracy insufficient (<95%) | High (60%) | Critical (4-month delay) | Early validation gates, iterate based on feedback |
| Pilot customers not identified | Medium (40%) | High (2-month delay) | Identify by Week 8, offer incentives |
| Deepfake API doesn't meet requirements | Medium (30%) | High (3-month delay) | Defer deepfake to v2 |
| Browser extension complexity underestimated | Medium (30%) | Medium (2-week delay) | POC in Sprint 3, validate Chrome MV3 early |
| False positives frustrate pilots | High (50%) | Medium (1-month delay) | Start with alerting-only, tune thresholds |
