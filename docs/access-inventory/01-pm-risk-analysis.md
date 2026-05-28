# PM / Risk Manager Analysis: Asset Inventory & Classification

**Date:** 2026-05-28  
**Analyst:** PM / Risk Manager (30yr experience)  
**Scope:** Sprints 1-5, 12 (Asset Inventory & Classification)

---

## Executive Risk Summary

🔴 **CRITICAL RISKS:**

1. **Sprint 5 Overload (Probability: 85% | Impact: HIGH)** — Slack + AWS discovery + OPA RBAC engine in one 2-week sprint with 2 backend engineers is a 3-sprint scope compressed into 1. AWS Config alone is 1 week of integration work.

2. **External API Dependency Chain (Probability: 70% | Impact: MEDIUM-HIGH)** — Success depends on 4 vendor APIs (Google, Microsoft, Slack, AWS) with different approval processes, rate limits, and breaking change policies. One API change mid-sprint = cascade delay.

3. **Single Point of Failure: Tech Lead (Probability: 60% | Impact: HIGH)** — Tech Lead is critical path for S1 (infrastructure), S2 (Google OAuth design), S3 (multi-provider abstraction), S5 (OPA architecture). No backup. Sick leave = sprint slip.

4. **Pilot Customer Availability Gap (Probability: 50% | Impact: MEDIUM)** — Sprint plan assumes pilot customers have all 4 providers (Google Workspace, M365, Slack Enterprise Grid, AWS). Reality: SMEs rarely have all 4. Testing will be partial.

5. **Slack Enterprise Grid Requirement (Probability: 40% | Impact: MEDIUM)** — Full Slack discovery requires Enterprise Grid ($$$). Most SME pilot customers will have Standard/Plus plans. Fallback to OAuth app discovery only = reduced coverage.

---

## Sprint-by-Sprint Assessment

### Sprint 1: Infrastructure and Auth (W1-2)
**Status:** 🟢 **GREEN**

**Scope:**
- AWS infra (VPC, ECS, RDS, S3, EventBridge, Secrets Manager)
- Multi-tenant DB schema + tenant isolation
- Keycloak SSO (Google + M365)
- MFA (TOTP)
- CI/CD pipeline

**Capacity:**
- Tech Lead: 90% (9d)
- Backend Eng: 90% (9d)
- Frontend Eng: 25% (2.5d)
- Flutter Eng: 15% (1.5d)
- DevSecOps (5d): 90% (4.5d)
- PM (5d): 80% (4d)
- **Total: 30.5 / 50 person-days (61%)**

**Assessment:**
- ✅ Realistic scope for infrastructure sprint
- ✅ DevSecOps contractor brings AWS expertise
- ✅ Keycloak SSO is well-documented
- ⚠️ **Risk:** RDS Multi-AZ + tenant isolation CI test added post-debate — adds 1-2 days. Still achievable but buffer is thin.

**Recommendations:**
1. Pre-provision AWS accounts in Week 0 (lead time: 2-3 days for Organizations setup)
2. Tech Lead should scaffold Terraform modules before sprint starts
3. Have Keycloak Docker Compose config ready for local dev

---

### Sprint 2: Google Workspace Asset Sync (W3-4)
**Status:** 🟡 **AMBER**

**Scope:**
- Google Admin SDK integration (OAuth 2.0 service account)
- Sync: users, groups, OAuth apps
- Background job (15-min incremental sync)
- API endpoint for verification

**Capacity:**
- Tech Lead: 60% (6d)
- Backend Eng: 100% (10d)
- Frontend Eng: 70% (7d)
- Flutter Eng: 20% (2d)
- DevSecOps (5d): 30% (1.5d)
- PM (5d): 70% (3.5d)
- **Total: 30.0 / 50 person-days (60%)**

**Assessment:**
- ✅ Google Admin SDK is mature and well-documented
- ✅ Backend Eng at 100% capacity is appropriate
- ⚠️ **Risk:** OAuth 2.0 service account setup requires domain-wide delegation approval from Google Workspace admin. If pilot customer's IT admin is slow to approve = sprint blocker.
- ⚠️ **Risk:** Incremental sync state tracking (what changed since last run?) is non-trivial. Google API supports `pageToken` but not delta queries for all resources.

**Recommendations:**
1. **CRITICAL:** PM must secure Google Workspace admin access from pilot customer by end of Sprint 1. Approval lead time: 3-5 days.
2. Backend Eng should use Google's `AdminReportsService` for audit logs to detect changes (fallback if delta queries unavailable)
3. Start with full sync every 15 min (simpler), optimize to incremental in Sprint 3 if performance issues arise

---

### Sprint 3: M365 Sync + Dashboard v1 (W5-6)
**Status:** 🔴 **RED** → **AMBER** (with scope reduction)

**Scope:**
- Microsoft Graph API integration
- Sync: users, groups, OAuth apps, M365 licensed apps
- Web dashboard: asset inventory table (filter, search, sort)
- CSV export
- User list with account status

**Capacity:**
- Tech Lead: 50% (5d)
- Backend Eng: 100% (10d)
- Frontend Eng: 100% (10d)
- Flutter Eng: 25% (2.5d)
- DevSecOps (5d): 30% (1.5d)
- PM (5d): 70% (3.5d)
- **Total: 32.5 / 50 person-days (65%)**

**Assessment:**
- ❌ **Original scope too large:** M365 Graph API + full-featured dashboard in 2 weeks is a stretch
- ⚠️ **Risk:** Microsoft Graph API has 10+ different auth scopes. Choosing wrong scopes = re-approval cycle (3-5 days)
- ⚠️ **Risk:** Dashboard UI (filter, search, sort, CSV export) is 5-7 days of frontend work. Frontend Eng at 100% but also building design system components in parallel.
- ✅ Multi-provider abstraction (Tech Lead 50%) is appropriate architectural investment

**Recommendations:**
1. **SCOPE REDUCTION:** Defer CSV export to Sprint 4. Focus on read-only table with basic filter/search.
2. **SCOPE REDUCTION:** Defer "User list with account status" to Sprint 4. Sprint 3 = asset table only.
3. Frontend Eng should reuse existing table component library (shadcn/ui or similar) — don't build from scratch
4. PM must secure M365 admin consent by end of Sprint 2 (lead time: 3-5 days)

**Revised Status:** 🟡 **AMBER** (with scope cuts)

---

### Sprint 4: Classification + Shadow IT Alerts (W7-8)
**Status:** 🟢 **GREEN**

**Scope:**
- Auto-classification (account type, sensitivity levels)
- Manual override + bulk CSV import
- Allow-list management (approved/pending/blocked)
- Alert email + Slack for new OAuth apps
- Classification history audit log

**Capacity:**
- Tech Lead: 50% (5d)
- Backend Eng: 90% (9d)
- Frontend Eng: 90% (9d)
- Flutter Eng: 25% (2.5d)
- DevSecOps (5d): 20% (1d)
- PM (5d): 60% (3d)
- **Total: 29.5 / 50 person-days (59%)**

**Assessment:**
- ✅ Realistic scope — classification rules are deterministic (no ML)
- ✅ Shadow IT detection via OAuth app monitoring is well-scoped
- ✅ Email + Slack alerts are straightforward (AWS SES + Slack webhook)
- ⚠️ **Risk:** Bulk CSV import for classification override = edge case handling (malformed CSV, invalid values). Add 1-2 days buffer.

**Recommendations:**
1. Use existing CSV parsing library (Papa Parse for frontend, Python csv module for backend)
2. Classification rules should be configurable via JSON file (not hardcoded) for future customization
3. PM should draft alert message templates in Sprint 3 to avoid last-minute copywriting

---

### Sprint 5: Slack + AWS Discovery + RBAC Engine (W9-10)
**Status:** 🔴 **RED** → **AMBER** (with scope reduction or timeline extension)

**Scope:**
- Slack Admin API: users, channels, installed apps
- AWS integration: EC2, S3, RDS, Lambda, IAM (via AWS Config + IAM API)
- OPA/Rego RBAC engine
- Policy evaluation <100ms
- Audit log for access decisions

**Capacity:**
- Tech Lead: 80% (8d)
- Backend Eng: 100% (10d)
- Frontend Eng: 70% (7d)
- Flutter Eng: 30% (3d)
- DevSecOps (5d): 50% (2.5d)
- PM (5d): 50% (2.5d)
- **Total: 33.0 / 50 person-days (66%)**

**Assessment:**
- ❌ **CRITICAL OVERLOAD:** This is 3 sprints of work compressed into 1:
  - **Slack integration:** 3-4 days (API auth, sync logic, Enterprise Grid vs Standard handling)
  - **AWS integration:** 5-7 days (AWS Config setup, IAM API, multi-account support, 5 resource types)
  - **OPA RBAC engine:** 4-5 days (policy design, OPA deployment, API middleware integration, <100ms requirement, audit log)
- ❌ **Single Backend Eng cannot deliver all 3 in 10 days**
- ⚠️ **Risk:** Slack Enterprise Grid requirement — most SME pilot customers won't have it. Fallback to OAuth app discovery only = reduced asset coverage.
- ⚠️ **Risk:** AWS multi-account setup requires AWS Organizations + cross-account IAM roles. If pilot customer doesn't have this = manual per-account setup (adds 2-3 days per account).

**Recommendations (CHOOSE ONE):**

**Option A: Split into 2 sprints (RECOMMENDED)**
- **Sprint 5A (W9-10):** Slack + AWS discovery only
- **Sprint 5B (W11-12):** OPA RBAC engine only
- Move original Sprint 6 (Offboarding) to Sprint 7

**Option B: Reduce scope within Sprint 5**
- **Keep:** Slack users + channels (defer installed apps to Sprint 6)
- **Keep:** AWS IAM users only (defer EC2, S3, RDS, Lambda to Sprint 6)
- **Keep:** OPA RBAC engine (core requirement for Sprint 6 Offboarding)
- **Defer:** Full AWS resource discovery to Sprint 6

**Option C: Add 3rd Backend Eng contractor for Sprint 5 only**
- Cost: ~$10k-15k for 2-week contract
- Assign: AWS integration (isolated work, clear API boundaries)
- Risk mitigation: Reduces single point of failure

**PM Decision Required:** Choose option by end of Sprint 3.

**Revised Status:** 🟡 **AMBER** (with Option B scope reduction)

---

### Sprint 12: Dependency Mapping + Lifecycle Tracking (W23-24)
**Status:** 🟡 **AMBER**

**Scope:**
- Dependency graph: user → OAuth app → cloud resource
- Blast radius view
- Asset lifecycle states (Discovered/Active/Inactive/Decommissioned)
- Auto-flag inactive >90 days
- Decommissioned asset retention in audit log

**Capacity:**
- Tech Lead: 70% (7d) — includes pen-test kickoff + daily findings review
- Backend Eng: 90% (9d)
- Frontend Eng: 90% (9d)
- Flutter Eng: 40% (4d)
- DevSecOps (5d): 80% (4d) — pen-test coordination
- PM (5d): 60% (3d)
- **Total: 36.0 / 50 person-days (72%)**

**Assessment:**
- ✅ Dependency graph is achievable with PostgreSQL JSONB (no need for dedicated graph DB at this scale)
- ⚠️ **Risk:** Pen-test starts in Sprint 12 (per debate outcome). If Critical/High findings in asset inventory = rework during Sprint 13.
- ⚠️ **Risk:** Dependency graph performance at 10k assets with complex relationships. Need indexed queries + caching.
- ⚠️ **Risk:** Blast radius calculation = recursive graph traversal. Could be slow without optimization.

**Recommendations:**
1. Use PostgreSQL recursive CTEs for graph traversal (built-in, no external dependencies)
2. Cache dependency graph in Redis with 15-min TTL (refresh on asset sync)
3. Limit blast radius depth to 3 levels (user → app → resource → sub-resource) to prevent infinite loops
4. Add database indexes on foreign key columns used in graph queries

**Revised Status:** 🟡 **AMBER** (achievable with caching + query optimization)

---

## Resource Bottlenecks

### Tech Lead (Critical Path)
**Sprints at risk:** S1 (90%), S2 (60%), S3 (50%), S5 (80%), S12 (70%)

**Bottleneck analysis:**
- S1: Infrastructure design + Keycloak SSO = foundational, cannot parallelize
- S2: Google OAuth architecture = must be right first time (refactoring OAuth is painful)
- S3: Multi-provider abstraction = architectural decision with long-term impact
- S5: OPA RBAC design = security-critical, cannot delegate
- S12: Pen-test findings review = cannot delegate

**Mitigation:**
1. **Pre-sprint prep:** Tech Lead should draft architecture docs 1 week before each sprint starts
2. **Backup plan:** Identify external architect consultant who can step in if Tech Lead is unavailable (sick leave, emergency)
3. **Knowledge transfer:** Tech Lead must document architectural decisions in ADRs (Architecture Decision Records) so team can continue if blocked

### Backend Eng (High Utilization)
**Sprints at risk:** S2 (100%), S3 (100%), S4 (90%), S5 (100%)

**Bottleneck analysis:**
- 4 consecutive sprints at 90-100% utilization = burnout risk
- No buffer for unexpected bugs or scope creep
- Single Backend Eng in S5 cannot deliver 3 integrations

**Mitigation:**
1. **Sprint 5:** Add contractor Backend Eng (see Option C above) OR split into 2 sprints
2. **Sprint 6-7:** Reduce Backend Eng to 70-80% to allow recovery time
3. **Code review:** Tech Lead must review Backend Eng PRs daily to catch issues early (don't batch at end of sprint)

### Frontend Eng (Moderate Utilization)
**Sprints at risk:** S3 (100%), S4 (90%), S11 (100%)

**Bottleneck analysis:**
- S3: Dashboard v1 is large scope (filter, search, sort, CSV export)
- S11: Compliance dashboard + reports = complex UI

**Mitigation:**
1. **Component reuse:** Use shadcn/ui or similar library (don't build from scratch)
2. **Design handoff:** PM should provide Figma mockups 1 week before sprint starts
3. **Sprint 3 scope reduction:** Defer CSV export to Sprint 4 (see above)

---

## Critical Path Analysis

### Dependency Chain

```
S1 (Infra) → S2 (Google) → S3 (M365 + Dashboard) → S4 (Classification) → S5 (Slack + AWS + RBAC) → S6 (Offboarding)
                                                                                                    ↓
                                                                                                  S12 (Dependency Graph)
```

**Critical path sprints (delay = launch delay):**
1. **Sprint 1:** Infrastructure — blocks everything
2. **Sprint 5:** RBAC engine — blocks Sprint 6 (Offboarding requires RBAC for access revocation)
3. **Sprint 6:** Offboarding — blocks Sprint 10 (Compliance evidence collection requires offboarding reports)

**Non-critical path sprints (can slip without launch delay):**
- Sprint 12: Dependency graph (nice-to-have for v1, can defer to v1.1)
- Sprint 4: Classification (can use manual classification as workaround)

### Minimum Viable Asset Inventory for Sprint 6 (Offboarding)

**Required by Sprint 6:**
- ✅ Google Workspace sync (users, OAuth apps)
- ✅ M365 sync (users, OAuth apps)
- ✅ Slack sync (users)
- ✅ AWS sync (IAM users)
- ✅ RBAC engine (to enforce who can trigger offboarding)

**Not required by Sprint 6:**
- ❌ Classification (offboarding works without it)
- ❌ Shadow IT alerts (offboarding works without it)
- ❌ Dependency graph (offboarding works without it)

**Implication:** If Sprint 5 slips, Sprint 6 is blocked. If Sprint 4 slips, Sprint 6 can proceed.

---

## External Dependencies & Lead Times

### Vendor API Approvals

| Provider | Approval Type | Lead Time | Owner | Deadline |
|----------|--------------|-----------|-------|----------|
| Google Workspace | Domain-wide delegation | 3-5 days | Pilot customer IT admin | End of S1 |
| Microsoft 365 | Admin consent (Graph API) | 3-5 days | Pilot customer IT admin | End of S2 |
| Slack | OAuth app approval | 1-2 days | Pilot customer Workspace admin | End of S4 |
| AWS | Cross-account IAM role | 2-3 days per account | Pilot customer AWS admin | End of S4 |

**Risk:** If pilot customer IT admin is slow to respond = sprint blocker.

**Mitigation:**
1. PM must identify pilot customers with responsive IT admins (SLA: <24hr response time)
2. Pre-draft approval request emails with step-by-step instructions (reduce back-and-forth)
3. Have backup pilot customer ready if primary is unresponsive

### Pen-Test Vendor

| Milestone | Lead Time | Owner | Deadline |
|-----------|-----------|-------|----------|
| Vendor selection | 2-3 weeks | PM | End of S8 |
| LOI signing | 1 week | PM | End of S8 |
| Pen-test scheduling | 2-3 weeks | Vendor | S12 start |

**Risk:** If vendor not booked by end of S8 = pen-test delayed to post-launch (ISO 27001 violation).

**Mitigation:**
1. PM must start vendor outreach in Sprint 7 (not Sprint 8)
2. Have 3 vendor quotes ready by end of Sprint 7
3. Budget approval secured before Sprint 8 starts

---

## Pilot Customer Readiness

### Provider Coverage Gap

**Assumption in sprint plan:** Pilot customers have all 4 providers (Google Workspace, M365, Slack Enterprise Grid, AWS)

**Reality check:**
- **Google Workspace OR M365:** 95% of SMEs have one (not both)
- **Slack Enterprise Grid:** <10% of SMEs (most have Standard/Plus)
- **AWS:** 60% of SMEs (others use Azure, GCP, or no cloud)

**Implication:** Testing will be partial. No single pilot customer can test all integrations.

**Mitigation:**
1. **Recruit 3 pilot customers with different provider mixes:**
   - Pilot A: Google Workspace + Slack Standard + AWS
   - Pilot B: M365 + Slack Plus + Azure (use Azure as "AWS equivalent" for testing)
   - Pilot C: Google Workspace + M365 + no Slack + AWS
2. **Slack fallback:** If no Enterprise Grid, test OAuth app discovery only (reduced coverage but still valuable)
3. **AWS fallback:** If no AWS, test with GCP or Azure (similar IAM concepts, different APIs)

### Onboarding Timeline

| Milestone | Lead Time | Owner | Deadline |
|-----------|-----------|-------|----------|
| Pilot customer recruitment | 4-6 weeks | PM | End of S3 |
| NDA + pilot agreement signing | 1-2 weeks | PM + Legal | End of S4 |
| IT admin access provisioning | 1 week | Pilot customer | End of S5 |
| First asset sync test | N/A | Team | Sprint 6 |

**Risk:** If pilot customers not onboarded by Sprint 6 = no real-world testing until Sprint 10+.

**Mitigation:**
1. PM must have 2 signed pilot LOIs by end of Sprint 8 (per debate outcome)
2. Start pilot recruitment in Sprint 2 (not Sprint 7)
3. Offer pilot customers free 6-month subscription as incentive

---

## Risk Register

| # | Risk | Probability | Impact | Mitigation | Owner | Status |
|---|------|-------------|--------|------------|-------|--------|
| 1 | Sprint 5 overload (3 integrations in 1 sprint) | 85% | HIGH | Split into 2 sprints OR add contractor | PM | Open |
| 2 | Tech Lead single point of failure | 60% | HIGH | Backup architect consultant + ADRs | PM | Open |
| 3 | Google Workspace admin approval delay | 50% | MEDIUM | Pre-draft approval email, backup pilot | PM | Open |
| 4 | M365 admin consent delay | 50% | MEDIUM | Pre-draft approval email, backup pilot | PM | Open |
| 5 | Slack Enterprise Grid unavailable | 40% | MEDIUM | Fallback to OAuth app discovery | Tech Lead | Open |
| 6 | AWS multi-account setup complexity | 40% | MEDIUM | Test with single-account first | Backend Eng | Open |
| 7 | Pilot customer provider coverage gap | 70% | MEDIUM | Recruit 3 pilots with different mixes | PM | Open |
| 8 | Pen-test vendor booking delay | 30% | HIGH | Start outreach in S7 (not S8) | PM | Open |
| 9 | Backend Eng burnout (4 sprints at 90-100%) | 50% | MEDIUM | Reduce to 70-80% in S6-7 | PM | Open |
| 10 | Sprint 3 dashboard scope too large | 60% | MEDIUM | Defer CSV export to S4 | PM | Open |

---

## Recommendations

### Immediate Actions (Before Sprint 1)

1. **[PM]** Recruit 3 pilot customers with different provider mixes (Google/M365/Slack/AWS). Target: 2 signed LOIs by end of Sprint 8.
2. **[PM]** Pre-provision AWS accounts and Organizations setup (lead time: 2-3 days).
3. **[Tech Lead]** Draft Terraform modules for VPC, ECS, RDS, S3 before Sprint 1 starts.
4. **[PM]** Draft Google Workspace domain-wide delegation approval email (ready to send in Sprint 1).

### Sprint-Specific Actions

**Sprint 1:**
- **[PM]** Send Google Workspace approval request to pilot customer IT admin (target: approval by end of S1).

**Sprint 2:**
- **[PM]** Send M365 admin consent request to pilot customer IT admin (target: approval by end of S2).
- **[Backend Eng]** Use Google AdminReportsService for change detection (fallback if delta queries unavailable).

**Sprint 3:**
- **[PM]** DECISION REQUIRED: Approve scope reduction (defer CSV export to S4).
- **[Frontend Eng]** Reuse shadcn/ui table component (don't build from scratch).

**Sprint 4:**
- **[PM]** Draft alert message templates for Shadow IT detection.
- **[Backend Eng]** Make classification rules configurable via JSON file.

**Sprint 5:**
- **[PM]** CRITICAL DECISION REQUIRED: Choose Option A (split into 2 sprints), Option B (reduce scope), or Option C (add contractor). Deadline: end of Sprint 3.

**Sprint 7:**
- **[PM]** Start pen-test vendor outreach (target: 3 quotes by end of S7).

**Sprint 8:**
- **[PM]** Sign pen-test vendor LOI (target: booking confirmed for S12).

**Sprint 12:**
- **[Backend Eng]** Use PostgreSQL recursive CTEs + Redis caching for dependency graph.
- **[DevSecOps]** Coordinate pen-test start (prerequisites: API auth complete, RDS in VPC, S3 BlockPublicAccess ON).

---

## Conclusion

**Overall Assessment:** 🟡 **AMBER** (achievable with scope adjustments and risk mitigation)

**Key Takeaways:**
1. **Sprint 5 is the highest-risk sprint** — must be split or scoped down.
2. **Tech Lead is critical path** — backup plan required.
3. **Pilot customer readiness is underestimated** — start recruitment in Sprint 2, not Sprint 7.
4. **External API approvals have 3-5 day lead times** — must be secured 1 sprint in advance.
5. **Pen-test vendor booking must start in Sprint 7** — not Sprint 8.

**Go/No-Go Recommendation:** **GO** (with mandatory Sprint 5 scope reduction and pilot recruitment acceleration)

---

**Next Steps:**
1. PM to review and approve Sprint 5 scope reduction (Option A or B) by end of Sprint 3
2. PM to draft pilot customer recruitment email by end of Sprint 1
3. Tech Lead to document architectural decisions in ADRs starting Sprint 1
4. PM to identify backup architect consultant by end of Sprint 2
