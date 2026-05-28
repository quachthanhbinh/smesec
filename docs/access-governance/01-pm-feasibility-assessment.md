# Access Governance Implementation Feasibility Assessment

**Project Manager:** Quách Thanh Bình  
**Assessment Date:** 2026-05-28  
**Timeline:** 6 months (Sprints 4-6, weeks 7-12)  
**Team Capacity:** 2 Backend Engineers (100%), 1 Frontend Engineer (80%), 1 Flutter Engineer (50%)

---

## Executive Summary

**CRITICAL FINDING:** The proposed 6-month timeline for comprehensive access governance is **NOT FEASIBLE** with current team capacity and scope. The team is already at 100% utilization with NO BUFFER, and the access governance scope requires an estimated 520-680 person-days of effort, while only 360 person-days are available.

**Recommendation:** Reduce scope by 40-50% or extend timeline to 9-10 months.

---

## 1. Feature Complexity Matrix

| Feature | Complexity | Effort (Person-Days) | Dependencies | Risk Level |
|---------|-----------|---------------------|--------------|------------|
| **IAM Foundation (SSO/MFA)** | Medium | 40-60 | None | Medium |
| **RBAC Engine (OPA/Casbin)** | High | 80-120 | IAM Foundation | High |
| **JIT Access System** | Very High | 100-140 | RBAC, IAM | Very High |
| **Offboarding Automation** | High | 60-80 | IAM, HR Integration | High |
| **Shadow IT Detection (CASB)** | Very High | 120-160 | Network/Browser Integration | Very High |
| **Access Review Automation** | High | 80-100 | RBAC, Workflow Engine | High |
| **HR System Integration** | Medium | 40-60 | Offboarding | Medium |
| **Browser Extension (Shadow IT)** | High | 60-80 | Backend API | High |
| **Testing & QA** | - | 80-100 | All features | High |
| **Documentation & Training** | - | 20-30 | All features | Low |
| **TOTAL** | - | **680-930** | - | - |

### Complexity Breakdown

#### IAM Foundation (40-60 days)
- **Components:** SSO integration (SAML/OAuth), MFA setup, user provisioning
- **Rationale:** Standard integration with proven platforms (Okta/Auth0), but requires careful security configuration
- **Learning curve:** 1-2 weeks for team unfamiliar with enterprise IAM

#### RBAC Engine (80-120 days)
- **Components:** Policy engine selection, role definitions, permission mappings, enforcement points
- **Rationale:** 
  - OPA: More flexible but steeper learning curve (Rego language)
  - Casbin: Simpler for pure RBAC but less extensible
  - Requires integration across all services
- **Learning curve:** 2-3 weeks for OPA/Rego, 1 week for Casbin
- **Research finding:** OPA is widely adopted but requires dedicated policy engineering expertise

#### JIT Access System (100-140 days)
- **Components:** Request workflow, approval engine, time-bounded grants, auto-revocation, audit logging
- **Rationale:** Most complex feature due to:
  - Multi-step approval workflows with escalation
  - Real-time revocation mechanisms (critical for security)
  - Integration with RBAC for temporary privilege elevation
  - Risk-adaptive approval logic (can reduce review volume by 50-70%)
- **Research finding:** Revocation speed is more critical than approval speed; if access outlasts the task, you create standing privilege
- **Common pitfall:** Manual revocation processes lead to privilege accumulation

#### Offboarding Automation (60-80 days)
- **Components:** HR system integration, AWS Step Functions workflow, account deactivation, access revocation, device wipe, data transfer
- **Rationale:**
  - 2026 best practice: 15-minute countdown for critical systems
  - Requires orchestration across multiple systems (IAM, SaaS apps, device management)
  - AWS Step Functions suitable for workflow orchestration
- **Research finding:** Industry gap - 1/3 of organizations still take >24 hours to fully offboard
- **Dependencies:** Requires HR API integration (BambooHR, Workday, Gusto)

#### Shadow IT Detection (120-160 days)
- **Components:** CASB integration OR browser extension, network monitoring, SaaS discovery, risk assessment, remediation workflows
- **Rationale:** HIGHEST COMPLEXITY feature due to:
  - Traditional CASBs have significant gaps (catalog only ~30k apps vs 200M active web apps)
  - Browser-centric approaches more effective in 2026
  - Shadow AI detection particularly challenging (AI tools exist within approved apps)
  - Mid-market companies use 200-800 SaaS apps but IT only knows about 40-80
- **Research finding:** Browser extensions + network monitoring hybrid approach recommended
- **Common pitfall:** CASB alone insufficient; requires multiple detection methods

#### Access Review Automation (80-100 days)
- **Components:** Review workflow engine, attestation system, reporting, remediation tracking
- **Rationale:**
  - 75% of access requests can be automated with proper workflow design
  - Requires continuous certification rather than one-time reviews
  - Integration with RBAC for permission analysis
- **Research finding:** Manual reviews cannot meet modern compliance requirements (SOC 2, ISO 27001)

---

## 2. Risk Register

| Risk ID | Risk Description | Probability | Impact | Severity | Mitigation Strategy |
|---------|-----------------|-------------|--------|----------|---------------------|
| **R1** | **Scope Creep** | High | Critical | **CRITICAL** | Strict scope boundary; defer shadow IT and access reviews to Phase 2 |
| **R2** | **OPA/Rego Learning Curve** | High | High | **HIGH** | 2-week dedicated training; consider Casbin for simpler RBAC needs |
| **R3** | **JIT Auto-Revocation Failures** | Medium | Critical | **HIGH** | Implement fail-safe revocation with monitoring; test extensively |
| **R4** | **HR System Integration Delays** | High | High | **HIGH** | Start integration early (Sprint 4); have manual fallback |
| **R5** | **Shadow IT Detection False Positives** | High | Medium | **MEDIUM** | Phased rollout; tune detection rules; user feedback loop |
| **R6** | **Team Capacity Overload** | Very High | Critical | **CRITICAL** | Reduce scope by 40-50% or extend timeline |
| **R7** | **AWS Step Functions Complexity** | Medium | Medium | **MEDIUM** | Use managed services; start with simple workflows |
| **R8** | **CASB Vendor Lock-in** | Medium | Medium | **MEDIUM** | Evaluate open-source alternatives; design abstraction layer |
| **R9** | **Testing Insufficient** | High | High | **HIGH** | Allocate 15-20% of timeline to testing; automate where possible |
| **R10** | **Offboarding Race Conditions** | Medium | Critical | **HIGH** | Implement idempotent workflows; comprehensive error handling |
| **R11** | **Browser Extension Compatibility** | Medium | Medium | **MEDIUM** | Test across Chrome, Edge, Firefox; graceful degradation |
| **R12** | **No Buffer for Unknowns** | Very High | Critical | **CRITICAL** | Add 20% buffer or reduce scope |

### Risk Details

#### R1: Scope Creep (CRITICAL)
- **Scenario:** Feature requests expand beyond core RBAC/JIT/Offboarding
- **Indicators:** Shadow IT and access reviews are "nice-to-have" not "must-have" for v1
- **Mitigation:** 
  - Defer shadow IT detection to Phase 2 (saves 120-160 days)
  - Defer access review automation to Phase 2 (saves 80-100 days)
  - Focus on core: IAM + RBAC + JIT + Offboarding

#### R2: OPA/Rego Learning Curve (HIGH)
- **Scenario:** Team unfamiliar with policy-as-code paradigm
- **Indicators:** Rego is declarative language requiring mindset shift
- **Mitigation:**
  - 2-week dedicated training before Sprint 4
  - Consider Casbin if team prefers imperative approach
  - Hire consultant for initial policy design

#### R3: JIT Auto-Revocation Failures (HIGH)
- **Scenario:** Time-bounded access not revoked, creating standing privileges
- **Indicators:** Research shows revocation failures are common attack vector
- **Mitigation:**
  - Implement fail-safe revocation with monitoring
  - Alert on revocation failures
  - Manual revocation fallback
  - Extensive testing of edge cases

#### R6: Team Capacity Overload (CRITICAL)
- **Scenario:** 680-930 person-days required vs 360 available
- **Calculation:**
  - 2 Backend Engineers × 100% × 120 days = 240 days
  - 1 Frontend Engineer × 80% × 120 days = 96 days
  - 1 Flutter Engineer × 50% × 120 days = 60 days
  - **Total: 396 days** (assuming 20 working days/month × 6 months)
- **Gap:** 284-534 person-days SHORT
- **Mitigation:** Reduce scope by 40-50% OR extend timeline to 9-10 months

---

## 3. Revised Sprint Allocation (6-Month Realistic Scope)

### Current Proposed Scope (NOT FEASIBLE)
- Sprint 4 (Weeks 7-8): RBAC
- Sprint 5 (Weeks 9-10): JIT Access
- Sprint 6 (Weeks 11-12): Offboarding

### Revised Scope Option A: Core Features Only (FEASIBLE)

**Total Effort: 360-440 person-days (fits in 396 available)**

#### Sprint 4 (Weeks 7-8): IAM Foundation
- **Effort:** 40-60 days
- **Deliverables:**
  - SSO integration (Okta/Auth0)
  - MFA enforcement for all users
  - User provisioning workflows
  - Basic audit logging
- **Success Criteria:** 100% MFA adoption, SSO for all apps

#### Sprint 5 (Weeks 9-10): RBAC Engine
- **Effort:** 80-120 days
- **Deliverables:**
  - Policy engine implementation (Casbin recommended for timeline)
  - Role definitions and permission mappings
  - Enforcement points in API gateway
  - Admin UI for role management
- **Success Criteria:** Least privilege enforced, <5% excessive permissions

#### Sprint 6 (Weeks 11-12): Offboarding Automation
- **Effort:** 60-80 days
- **Deliverables:**
  - HR system integration (BambooHR/Workday)
  - AWS Step Functions workflow
  - Account deactivation automation
  - Access revocation across all systems
  - Audit trail and reporting
- **Success Criteria:** <4 hours offboarding completion time

#### Sprint 7 (Weeks 13-14): Testing & Hardening
- **Effort:** 80-100 days
- **Deliverables:**
  - End-to-end testing
  - Security testing
  - Performance testing
  - Documentation
  - Training materials
- **Success Criteria:** All critical paths tested, zero P0 bugs

#### Sprint 8 (Weeks 15-16): Buffer & Polish
- **Effort:** 40-60 days
- **Deliverables:**
  - Bug fixes
  - Performance optimization
  - User feedback incorporation
  - Production readiness review
- **Success Criteria:** Production deployment approved

**DEFERRED TO PHASE 2 (Months 7-12):**
- JIT Access System (100-140 days)
- Shadow IT Detection (120-160 days)
- Access Review Automation (80-100 days)

### Revised Scope Option B: Include JIT Access (REQUIRES 8 MONTHS)

**Total Effort: 460-580 person-days (requires 8 months)**

#### Months 1-2: IAM Foundation (40-60 days)
#### Months 3-4: RBAC Engine (80-120 days)
#### Months 5-6: JIT Access System (100-140 days)
#### Month 7: Offboarding Automation (60-80 days)
#### Month 8: Testing & Buffer (80-100 days)

**DEFERRED TO PHASE 2:**
- Shadow IT Detection
- Access Review Automation

---

## 4. Recommendations

### Primary Recommendation: Reduce Scope (Option A)

**Implement Core Features in 6 Months:**
1. IAM Foundation (SSO/MFA)
2. RBAC Engine
3. Offboarding Automation
4. Testing & Buffer

**Defer to Phase 2 (Months 7-12):**
1. JIT Access System
2. Shadow IT Detection
3. Access Review Automation

**Rationale:**
- Fits within 396 available person-days
- Delivers immediate security value (SSO, RBAC, offboarding)
- Reduces risk of incomplete features
- Allows learning from Phase 1 before tackling complex JIT/Shadow IT

### Alternative Recommendation: Extend Timeline (Option B)

**Extend to 8 Months to Include JIT Access:**
- Critical for zero standing privileges principle
- High security value for privileged access
- Requires additional 2 months

**Defer to Phase 2:**
- Shadow IT Detection (most complex, 120-160 days)
- Access Review Automation (can be manual initially)

### Technology Recommendations

#### RBAC Engine: Choose Casbin over OPA
- **Rationale:** Simpler learning curve (1 week vs 2-3 weeks)
- **Trade-off:** Less flexible than OPA, but sufficient for RBAC needs
- **Timeline impact:** Saves 2-3 weeks of learning curve

#### Offboarding: Use AWS Step Functions
- **Rationale:** Managed service, visual workflow designer, proven for orchestration
- **Research finding:** Used successfully by OneMain Financial for security workflows (97.5% speed improvement)

#### Shadow IT: Defer or Use Browser Extension Approach
- **Rationale:** CASB alone insufficient (only covers ~30k of 200M apps)
- **Alternative:** Browser extension for AI/SaaS monitoring (60-80 days vs 120-160 for full CASB)
- **Recommendation:** Defer to Phase 2 when team has more capacity

#### IAM Platform: Okta or Auth0
- **Rationale:** Proven enterprise platforms, good SME pricing, comprehensive features
- **Cost:** $5,000-8,000/year (within budget)

### Capacity Recommendations

**Option 1: Reduce Scope (Recommended)**
- Keep current team
- Implement Option A (Core Features)
- Deliver Phase 2 in months 7-12

**Option 2: Increase Capacity**
- Add 1 Backend Engineer (50% allocation) for 6 months
- Allows Option B (include JIT Access)
- Cost: ~$30,000-40,000 for 6 months

**Option 3: Extend Timeline**
- Extend to 8 months
- Implement Option B
- No additional cost

### Risk Mitigation Priorities

**Week 1-2 (Before Sprint 4):**
1. OPA/Casbin training for backend team
2. HR system API access and documentation
3. AWS Step Functions proof-of-concept
4. Finalize scope decision (Option A vs B)

**Sprint 4 (Weeks 7-8):**
1. Start HR integration early (parallel with IAM)
2. Set up monitoring and alerting infrastructure
3. Establish testing framework

**Sprint 5 (Weeks 9-10):**
1. Weekly scope reviews to prevent creep
2. Performance testing for RBAC enforcement
3. Security review of policy engine

**Sprint 6 (Weeks 11-12):**
1. Extensive offboarding testing (race conditions, failures)
2. Disaster recovery procedures
3. Runbook documentation

---

## 5. Dependencies and Critical Path

### Critical Path (Option A - 6 Months)

```
IAM Foundation (Weeks 7-8)
    ↓
RBAC Engine (Weeks 9-10)
    ↓
Offboarding Automation (Weeks 11-12)
    ↓
Testing & Hardening (Weeks 13-14)
    ↓
Buffer & Polish (Weeks 15-16)
```

### External Dependencies

1. **HR System API Access** (Week 7)
   - Required for offboarding automation
   - Risk: Delays if HR team unresponsive
   - Mitigation: Start request in Week 6

2. **IAM Platform Selection** (Week 6)
   - Required before Sprint 4
   - Risk: Procurement delays
   - Mitigation: Pre-approve budget and vendor

3. **AWS Account Setup** (Week 7)
   - Required for Step Functions
   - Risk: Corporate approval process
   - Mitigation: Start request in Week 6

4. **Device Management Platform** (Week 11)
   - Required for device wipe in offboarding
   - Risk: Integration complexity
   - Mitigation: Use existing MDM if available

### Parallel Work Streams

**Backend Team (2 engineers):**
- IAM integration
- RBAC policy engine
- Offboarding workflow
- API endpoints

**Frontend Team (1 engineer, 80%):**
- Admin UI for role management
- Offboarding dashboard
- Audit log viewer
- User self-service portal

**Flutter Team (1 engineer, 50%):**
- Mobile admin app
- Push notifications for approvals
- Offline capability for critical functions

---

## 6. Success Metrics and Gates

### Sprint 4 Gate (IAM Foundation)
- [ ] 100% MFA adoption across organization
- [ ] SSO integrated for all critical applications
- [ ] User provisioning automated
- [ ] Audit logging operational
- [ ] <2 second authentication latency

### Sprint 5 Gate (RBAC Engine)
- [ ] All users assigned to roles (no direct permissions)
- [ ] <5% users with excessive permissions
- [ ] Policy enforcement at API gateway
- [ ] Admin UI functional
- [ ] <100ms policy evaluation latency

### Sprint 6 Gate (Offboarding Automation)
- [ ] <4 hours full offboarding completion
- [ ] 100% account deactivation success rate
- [ ] Zero manual steps required
- [ ] Complete audit trail
- [ ] HR system integration tested

### Sprint 7 Gate (Testing & Hardening)
- [ ] 100% critical path test coverage
- [ ] Zero P0/P1 bugs
- [ ] Security review passed
- [ ] Performance benchmarks met
- [ ] Documentation complete

### Sprint 8 Gate (Production Readiness)
- [ ] Disaster recovery tested
- [ ] Runbooks complete
- [ ] Team trained
- [ ] Monitoring and alerting operational
- [ ] Stakeholder sign-off

---

## 7. Budget Impact

### Option A (Core Features - 6 Months)

| Item | Cost |
|------|------|
| IAM Platform (Okta/Auth0) | $2,500-4,000 (6 months) |
| MFA Solution | $1,000-1,500 (6 months) |
| Workflow Automation | $500-1,000 (6 months) |
| AWS Services (Step Functions, etc.) | $500-1,000 (6 months) |
| Training & Consulting | $5,000-10,000 |
| **Total** | **$9,500-17,500** |

### Option B (Include JIT - 8 Months)

| Item | Cost |
|------|------|
| IAM Platform | $3,300-5,300 (8 months) |
| MFA Solution | $1,300-2,000 (8 months) |
| JIT Access Tool | $2,000-3,300 (8 months) |
| Workflow Automation | $700-1,300 (8 months) |
| AWS Services | $700-1,300 (8 months) |
| Training & Consulting | $8,000-15,000 |
| **Total** | **$16,000-28,200** |

### Additional Capacity Option

| Item | Cost |
|------|------|
| 1 Backend Engineer (50%, 6 months) | $30,000-40,000 |

---

## 8. Conclusion

**The proposed 6-month timeline for comprehensive access governance is NOT FEASIBLE with current team capacity.**

**Recommended Path Forward:**

1. **Adopt Option A (Core Features)** - Implement IAM, RBAC, and Offboarding in 6 months
2. **Defer JIT Access, Shadow IT, and Access Reviews** to Phase 2 (months 7-12)
3. **Choose Casbin over OPA** for faster implementation
4. **Start HR integration early** (Week 7) to avoid delays
5. **Allocate 20% of timeline to testing** (Sprints 7-8)
6. **Add 20% buffer** for unknowns (Sprint 8)

**This approach:**
- Delivers immediate security value (SSO, RBAC, offboarding)
- Fits within available capacity (360-440 days vs 396 available)
- Reduces risk of incomplete features
- Allows learning from Phase 1 before tackling complex features
- Maintains team morale by avoiding burnout

**Alternative:** If JIT Access is critical for v1, extend timeline to 8 months (Option B).

---

## Sources

- [Automation of User Onboarding and Offboarding Workflows](https://aws.amazon.com/blogs/apn/automation-of-user-onboarding-and-offboarding-workflows/)
- [Risk-Adaptive Just-In-Time Access](https://multiplierhq.com/blog/risk-adaptive-just-in-time-access-policy-approval-playbook)
- [Understanding Approval Workflows in Access Management](https://multiplierhq.com/blog/understanding-approval-workflows-in-access-management)
- [Top Open-Source Authorization Tools for Enterprises in 2026](https://www.permit.io/blog/top-open-source-authorization-tools-for-enterprises-in-2026)
- [Think Your IdP or CASB Covers Shadow IT? These 5 Risks Prove Otherwise](https://thehackernews.com/2025/06/think-your-idp-or-casb-covers-shadow-it.html)
- [Beyond CASB: A Browser-Centric Approach to SaaS Security](https://layerxsecurity.com/blog/beyond-casb-a-browser-centric-approach-to-saas-security/)
- [How to Build a SaaS Access Review Workflow That Scales in 2026](https://www.toriihq.com/articles/saas-access-review-workflow/)
- [Centralized Workflows for Faster Access Management [2026]](https://multiplierhq.com/blog/centralized-workflows-for-faster-identity-provisioning)
