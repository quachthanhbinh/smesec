# Access Governance Research Synthesis
**Date:** 2026-05-28  
**Participants:** Product Owner, Project Manager, Technical Advisor

## Executive Summary

3 agents completed parallel research on competitive products, implementation feasibility, and technical architecture for Access Governance (Requirement 3). Key finding: **Current 6-month timeline with proposed scope is NOT FEASIBLE** - requires 680-930 person-days vs 396 available.

## Key Conflicts Requiring Resolution

### 1. RBAC Engine Selection: OPA vs Casbin

**PM Recommendation:** Casbin
- Simpler, faster implementation (1 week learning curve vs 2-3 weeks for OPA)
- Native Node.js/TypeScript support
- Lower risk for timeline-constrained project

**TA Recommendation:** OPA
- Performance exceeds requirements by 750x (40μs vs 100ms target)
- Better for multi-service architecture
- Industry standard (CNCF graduated)
- Policy-as-code enables centralized governance

**Conflict:** PM prioritizes speed-to-delivery, TA prioritizes long-term architecture quality.

### 2. Timeline Feasibility

**PM Assessment:** NOT FEASIBLE
- Full scope requires 680-930 person-days
- Available capacity: 396 person-days (6 months)
- **Gap: 284-534 person-days SHORT**

**TA Assessment:** FEASIBLE
- Estimates 16 weeks (4 months) for 4 phases
- Leaves 2 months buffer in 6-month window

**Conflict:** Different scope assumptions - PM analyzed all features, TA focused on core features only.

### 3. MVP Scope Definition

**PO Must-Have (6 features):**
1. Automated access reviews
2. Automated offboarding
3. Shadow IT discovery (OAuth + API keys)
4. Basic RBAC
5. Real-time compliance dashboard
6. Quick setup (<1 week)

**PM Recommended Scope (3 features):**
1. IAM foundation
2. Basic RBAC
3. Automated offboarding

**Defer to Phase 2:**
- JIT Access
- Shadow IT Detection
- Access Reviews

**Conflict:** PO wants competitive feature parity, PM wants achievable timeline.

## Competitive Intelligence Summary

### Market Positioning

| Competitor | Strengths | Weaknesses | Price Point |
|------------|-----------|------------|-------------|
| **Vanta/Drata/Secureframe** | Strong compliance automation, 300+ integrations | Manual offboarding, basic RBAC, complex deployment | Mid ($$$) |
| **Nudge Security** | Advanced shadow IT/AI detection, behavioral nudging | Discovery-only (no enforcement), no PAM | Mid ($$$) |
| **Okta/Auth0** | Enterprise IAM, strong RBAC | Complex for SMEs, expensive, no shadow IT | High ($$$$) |
| **BeyondTrust/CyberArk** | Advanced PAM, session monitoring | Very expensive, complex deployment, enterprise-only | Very High ($$$$$) |

### SMESec Differentiation Opportunities

1. **Speed-to-Value:** <1 week deployment vs 3+ months for competitors
2. **Unified Platform:** Shadow IT discovery + enforcement + governance (vs Nudge's discovery-only)
3. **SME-Native:** Built for 50-200 employees, no security team required
4. **Automated Offboarding:** <5 min across all providers (competitors are manual)
5. **Cost-Effective:** $15K-24K/year vs $50K+ for enterprise solutions

## Customer Pain Points (Validated)

### Top 5 Pain Points for SMEs (50-200 employees)

1. **Visibility Gaps Despite "Good" Metrics** (State of Identity Governance 2026)
   - Organizations report strong metrics while real risks remain unaddressed
   - ROI: Early breach detection saves $4.45M average (IBM 2025)

2. **Delayed De-provisioning & Orphan Accounts**
   - Ex-employees retain access, become attack vectors
   - ROI: 60-70% reduction in security incidents with automation

3. **Shadow IT & Unmanaged SaaS Sprawl**
   - OAuth grants, API keys, GenAI tools untracked
   - ROI: 30% reduction in SaaS spend through duplicate elimination

4. **Complex Deployment & High TCO**
   - Weeks-to-months deployment, incomplete coverage
   - ROI: Fast deployment = faster time-to-value

5. **Weak Governance Frameworks**
   - Preventable breaches due to lagging controls
   - ROI: 75% reduction in audit prep time

## Technical Architecture Recommendations

### Automated Offboarding (Highest Priority)

**Architecture:** AWS Step Functions with parallel execution
- Google Workspace, M365, Slack, AWS IAM revocation in parallel
- Expected time: 30-45 seconds (well under <5 min SLA)
- Partial failure handling with audit trail

**API Rate Limits Analysis:**
- Google: 2,400 req/min → <5 sec
- M365: 130,000 req/10s → <5 sec
- Slack: 20+ req/min → <10 sec
- AWS IAM: Variable → <30 sec

**Verdict:** ✅ <5 min SLA is easily achievable

### RBAC Policy Engine

**OPA Performance:**
- Mean latency: 40 microseconds
- 99th percentile: 134 microseconds
- **750x faster than 100ms requirement**

**Implementation Complexity:**
- OPA: 2-3 week learning curve (Rego DSL)
- Casbin: 1 week learning curve (code-based)

### Shadow IT Detection

**Recommended:** Hybrid approach
- OAuth token scanning (80% coverage, real-time)
- Expense report correlation (20-30% coverage, monthly)
- Defer: Network traffic analysis (requires CASB infrastructure)

**Risk Scoring Algorithm:**
```
Risk Score = (OAuth Scope × 40%) + (Vendor Reputation × 30%) 
           + (User Count × 20%) + (Data Sensitivity × 10%)
```

## Capacity Analysis

### Current Team Allocation (Sprints 4-6)
- 2 Backend Engineers: 100% (14 days/sprint × 3 sprints = 42 days each = 84 days total)
- 1 Frontend Engineer: 80% (11.2 days/sprint × 3 sprints = 33.6 days)
- 1 Flutter Engineer: 50% (7 days/sprint × 3 sprints = 21 days)
- **Total Available:** 138.6 person-days for 6 weeks (Sprints 4-6)

### Full 6-Month Capacity
- 2 Backend Engineers: 100% × 120 days = 240 days
- 1 Frontend Engineer: 80% × 120 days = 96 days
- 1 Flutter Engineer: 50% × 120 days = 60 days
- **Total Available:** 396 person-days

### Feature Effort Estimates (from PM)

| Feature | Effort (person-days) | Priority |
|---------|---------------------|----------|
| IAM Foundation (SSO/MFA) | 60-80 | Must-have |
| RBAC Engine | 80-120 | Must-have |
| Automated Offboarding | 60-80 | Must-have |
| JIT Access | 100-140 | Defer |
| Shadow IT Detection | 120-160 | Defer |
| Access Reviews | 80-120 | Defer |
| Compliance Dashboard | 40-60 | Nice-to-have |
| Mobile/Desktop UI | 60-80 | Must-have |
| Integration Testing | 80-100 | Must-have |

**Core MVP (Must-have only):** 380-520 person-days  
**Available capacity:** 396 person-days  
**Verdict:** ⚠️ TIGHT - requires perfect execution, no buffer

## Recommendations for Debate

### Questions for Product Owner
1. Can we defer Shadow IT Detection to v2 given 120-160 day effort?
2. Is automated offboarding alone sufficient differentiation vs competitors?
3. What's minimum viable compliance dashboard (reduce from 40-60 days)?

### Questions for Project Manager
1. Can we reduce RBAC scope to 5 default roles only (no custom roles in v1)?
2. What's the risk of 0-day buffer? Should we extend to 8 months?
3. Can we phase rollout: Core features in Month 6, polish in Month 7-8?

### Questions for Technical Advisor
1. Does OPA's 2-3 week learning curve justify the long-term benefits?
2. Can we prototype Step Functions in Sprint 5 to validate approach?
3. What's the minimum viable shadow IT detection (OAuth only, no network analysis)?

## Next Steps

1. **Facilitate 3-way debate** to resolve OPA vs Casbin, scope, and timeline
2. **Define final MVP scope** with clear must-have/defer boundaries
3. **Create revised sprint plan** with realistic effort estimates
4. **Document architecture decisions** in ADR format
5. **Get stakeholder sign-off** on reduced scope or extended timeline
