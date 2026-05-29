# SMESec Strategy Documents

**Date:** 2026-05-29  
**Status:** Organized with numbered sequence

---

## 📚 Document Index

### Core Architecture & Design (01-03)

#### [01-system-architecture.md](01-system-architecture.md)
**System Architecture Diagram — Logical and Deployment View**
- Clean Architecture layers (Interface → Application → Domain ← Infrastructure)
- AWS Multi-Region deployment architecture
- Integration touchpoints with 3rd-party SaaS tools
- Technology stack and infrastructure components
- **Use when:** Understanding overall system structure, deployment topology, or integration points

#### [02-design-document.md](02-design-document.md)
**Core Architectural Decisions**
- Build vs Buy decisions (Hybrid approach)
- Multi-tenancy model (Shared PostgreSQL with RLS)
- AI-threat detection strategy (2-Track architecture)
- Data privacy guarantees (4 core commitments)
- **Use when:** Making architectural decisions, understanding design rationale, or evaluating trade-offs

#### [03-two-track-approach.md](03-two-track-approach.md)
**2-Track Development Strategy**
- Track 1: Foundation & Governance (deterministic, ~100% accuracy)
- Track 2: AI Threat Detection (ML-gated, >90% accuracy target)
- Team structure and parallel development approach
- Dependencies and integration points between tracks
- **Use when:** Understanding why Track 1 and Track 2 are separated, or planning feature development

---

### Delivery Plans (04-07)

**Choose the plan that matches your constraints:**

#### [04-delivery-plan-original.md](04-delivery-plan-original.md)
**Original Plan — 12 months (Aggressive)**
- **MVP:** Month 3 | **v1:** Month 6 | **v1.5:** Month 9 | **v2:** Month 12
- **Assumptions:** Full team (7 FTE) available Day 1, ML Engineer #1 joins Day 1
- **Sprint utilization:** 75-90% (high intensity)
- **Best for:** Well-funded startup with experienced team ready to start
- **Risk:** Burnout, external dependencies may delay

#### [05-delivery-plan-adjusted-1.5x.md](05-delivery-plan-adjusted-1.5x.md)
**1.5x Adjusted Plan — 19.5 months (Moderate)**
- **MVP:** Month 4.5 | **v1:** Month 9.75 | **v1.5:** Month 14.25 | **v2:** Month 19.5
- **Assumptions:** Full team available Day 1, ML Engineer #1 joins Day 1
- **Sprint utilization:** 60-75% (moderate pace)
- **Best for:** Team wanting balance between speed and sustainability
- **Risk:** Medium risk, some buffer for external dependencies

#### [06-delivery-plan-adjusted-2x.md](06-delivery-plan-adjusted-2x.md)
**2x Adjusted Plan — 26 months (Sustainable)**
- **MVP:** Month 6 | **v1:** Month 13 | **v1.5:** Month 19 | **v2:** Month 26
- **Assumptions:** Full team available Day 1, ML Engineer #1 joins Day 1
- **Sprint utilization:** 50-60% (sustainable pace)
- **Best for:** Team prioritizing quality and work-life balance
- **Risk:** Lower risk, realistic external dependencies

#### [07-delivery-plan-realistic-hiring.md](07-delivery-plan-realistic-hiring.md)
**Realistic Hiring Plan — 36+ months (Progressive Build-up)**
- **MVP:** Month 12 | **v1:** Month 20 | **v1.5:** Month 28 | **v2:** Month 36+
- **Assumptions:** Tech Lead starts solo, progressive team recruitment
- **ML Engineer #1:** Joins Month 8 (not Day 1) → Track 2 delayed 8 months
- **Sprint utilization:** 40-60% (sustainable with onboarding overhead)
- **Best for:** Solo founder or small team building from scratch
- **Risk:** Slowest time-to-market, but most realistic for bootstrapped startup

**Comparison Table:**

| Metric | Original | 1.5x | 2x | Realistic Hiring |
|--------|----------|------|----|--------------------|
| **MVP** | M3 | M4.5 | M6 | M12 |
| **v1** | M6 | M9.75 | M13 | M20 |
| **v1.5** | M9 | M14.25 | M19 | M28 |
| **v2** | M12 | M19.5 | M26 | M36+ |
| **Team at start** | 7 FTE | 7 FTE | 7 FTE | 1 FTE (TL solo) |
| **ML Eng #1 join** | Day 1 | Day 1 | Day 1 | Month 8 |
| **Sprint util** | 75-90% | 60-75% | 50-60% | 40-60% |

---

### Team & Execution (08-10)

#### [08-team-scope-of-work.md](08-team-scope-of-work.md)
**Detailed Role Responsibilities and Phase-by-Phase Ownership**
- 15 role definitions with hiring profiles
- Primary ownership matrix by domain
- Phase-by-phase focus areas for each role
- Cross-role responsibilities and coordination
- **Use when:** Hiring, defining role boundaries, or understanding who owns what

#### [09-ai-governance-module.md](09-ai-governance-module.md)
**AI Tool Detection and Governance Approach**
- 3-tier governance framework (Discover → Govern → Protect)
- Shadow AI discovery mechanisms
- Risk classification framework (CRITICAL → HIGH → MEDIUM → LOW)
- Browser DLP architecture (local inference, privacy-preserving)
- Deepfake fraud defense workflow
- **Use when:** Implementing AI governance features, understanding Track 2 scope

#### [10-feasibility-assessment-and-remediation-plan.md](10-feasibility-assessment-and-remediation-plan.md)
**Risk Identification and Mitigation Strategies**
- CRITICAL risks that could kill the product
- HIGH risks that could derail the plan
- Blind spots not yet accounted for
- Remediation solutions for each risk
- Top 5 Week 1 decisions
- **Use when:** Risk assessment, sprint planning, or decision-making under uncertainty

---

### Implementation Details (11-13)

#### [11-third-party-integration-principles.md](11-third-party-integration-principles.md)
**Integration Strategy and Vendor Selection**
- Integration architecture patterns
- Vendor evaluation criteria
- API rate limit handling strategies
- OAuth scope management
- **Use when:** Evaluating new integrations or troubleshooting existing ones

#### [12-third-party-preparation-plan.md](12-third-party-preparation-plan.md)
**Long-Lead-Time Vendor Registrations and API Access**
- Priority 1 (CRITICAL): Google Workspace, M365, Vanta (start before Sprint 1)
- Priority 2 (HIGH): Slack, AWS IAM, Hive, Lakera (start Sprint 1-2)
- Priority 3 (MEDIUM): Chrome Web Store, App Store, Pentest vendor (start Sprint 3-5)
- Timeline and hard deadlines for each vendor
- **Use when:** Project kickoff, sprint planning, or unblocking external dependencies

#### [13-metrics-scorecard.md](13-metrics-scorecard.md)
**Success Metrics and KPIs**
- 5-tier metrics framework (Product → Engineering → Business → Compliance → AI/ML)
- Success criteria by milestone (MVP, v1, v1.5, v2)
- Monthly tracking cadence
- **Use when:** Measuring progress, reporting to stakeholders, or setting goals

---

## 🗂️ Archive

The `archive/` folder contains older checklist documents that have been superseded by the numbered documents above:
- `INTEGRATION_SYNC_STATUS.md`
- `PRIORITIZED_3RD_PARTY_PREP_PLAN.md`
- `THIRD_PARTY_READINESS_CHECKLIST.md`

---

## 📖 Reading Order Recommendations

### For New Team Members
1. Start with [01-system-architecture.md](01-system-architecture.md) — understand the big picture
2. Read [02-design-document.md](02-design-document.md) — understand key decisions
3. Read [03-two-track-approach.md](03-two-track-approach.md) — understand development strategy
4. Choose your delivery plan (04-07) based on your context
5. Read [08-team-scope-of-work.md](08-team-scope-of-work.md) — understand your role

### For Project Managers
1. Read all delivery plans (04-07) — understand timeline options
2. Read [10-feasibility-assessment-and-remediation-plan.md](10-feasibility-assessment-and-remediation-plan.md) — understand risks
3. Read [12-third-party-preparation-plan.md](12-third-party-preparation-plan.md) — understand external dependencies
4. Read [13-metrics-scorecard.md](13-metrics-scorecard.md) — understand success metrics

### For Engineers
1. Read [01-system-architecture.md](01-system-architecture.md) — understand architecture
2. Read [03-two-track-approach.md](03-two-track-approach.md) — understand Track 1 vs Track 2
3. Read [08-team-scope-of-work.md](08-team-scope-of-work.md) — understand your responsibilities
4. Read [11-third-party-integration-principles.md](11-third-party-integration-principles.md) — understand integration patterns

### For ML Engineers
1. Read [03-two-track-approach.md](03-two-track-approach.md) — understand Track 2 scope
2. Read [09-ai-governance-module.md](09-ai-governance-module.md) — understand AI features
3. Read [08-team-scope-of-work.md](08-team-scope-of-work.md) — understand ML Engineer roles
4. Read [02-design-document.md](02-design-document.md) — understand AI-threat detection strategy

---

## 🔄 Document Maintenance

- **Last updated:** 2026-05-29
- **Maintained by:** Tech Lead + PM
- **Update frequency:** As needed when major decisions change
- **Version control:** All documents are in git — see commit history for changes

---

*All documents use kebab-case lowercase naming with numbered prefixes for easy navigation.*
