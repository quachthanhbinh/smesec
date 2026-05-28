# Access Governance — Feature Research Synthesis
**Date:** 2026-05-28  
**Method:** 3 subagents × 2 rounds (Round 1: parallel research | Round 2: cross-iteration)  
**Agents:** Product Owner · Technical Advisor · Project Manager  
**Constraint:** Fresh research — NOT biased by prior Track 1/2 plan

---

## Executive Summary

After 2 rounds of multi-agent iteration, the 3 agents reached consensus on the following:

1. **Lumos is the real competitor** (not Vanta/Drata). Lumos = access governance SaaS, $5–10/user/mo. Vanta = compliance evidence tool. These are different markets.
2. **The product anchor is "nothing falls through the cracks" offboarding** — but scoped to Google + M365 automated + checklist for all others. AWS and Slack deprovisioning automation explicitly deferred.
3. **Shadow IT allow-list (discover → human review → block) closes a gap no competitor covers** at SMB pricing. Auto-remediation deferred.
4. **Risk scoring replaced by compliance findings** — deterministic, binary, zero false positive risk. ML-based risk scores are a v3 feature after 6+ months of customer data.
5. **SOC 2 Type 1 is on the critical path** — must start Week 1, not Month 4. It gates the first paying customer.
6. **Mobile app is a genuine differentiator** — zero competitors offer it; already funded within sprint capacity.

---

## Part 1 — Market & Competitive Intelligence

### 1.1 Competitor Feature Matrix (Top 10 Products)

| Feature | Vanta | Drata | Nudge Security | Okta Lifecycle | BetterCloud | Zluri | **Lumos** | Entra ID Gov | SailPoint Biz+ | Cerby |
|---|---|---|---|---|---|---|---|---|---|---|
| SaaS asset inventory | ⚠️ | ⚠️ | ✅ | ⚠️ | ✅ | ✅ | ✅ | ⚠️ M365 only | ⚠️ | ❌ |
| Shadow IT discovery | ⚠️ | ❌ | ✅ Core | ❌ | ✅ | ✅ | ⚠️ | ⚠️ | ❌ | ❌ |
| Shadow IT remediation | ❌ | ❌ | ❌ nudge only | ⚠️ | ✅ | ✅ | ⚠️ | ⚠️ Complex | ❌ | ❌ |
| Automated offboarding | ❌ | ❌ | ❌ | ✅ SCIM | ✅ | ✅ | ✅ | ✅ Complex | ✅ | ❌ |
| Self-service access requests | ❌ | ❌ | ❌ | ⚠️ | ⚠️ | ✅ | ✅ **Core** | ✅ | ✅ | ❌ |
| Access reviews | ✅ | ✅ | ❌ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ❌ |
| JIT / privileged access | ❌ | ❌ | ❌ | ✅ | ❌ | ⚠️ | ⚠️ | ✅ PIM | ✅ | ❌ |
| RBAC enforcement | ❌ | ❌ | ❌ | ✅ | ⚠️ | ⚠️ | ⚠️ | ✅ | ✅ | ❌ |
| MFA enforcement check | ✅ | ✅ | ⚠️ | ✅ | ⚠️ | ❌ | ❌ | ✅ | ✅ | ❌ |
| Compliance evidence (SOC 2) | ✅ Core | ✅ Core | ❌ | ⚠️ | ❌ | ❌ | ❌ | ⚠️ | ⚠️ | ❌ |
| Audit trail | ✅ | ✅ | ⚠️ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| **Mobile app** | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |
| Non-expert UX (SMB-native) | ✅ | ✅ | ✅ | ❌ Expert | ⚠️ | ✅ | ✅ **Best** | ❌ Complex | ❌ Complex | ⚠️ |
| SMB pricing (50 users) | ~$8–15K/yr | ~$8–15K/yr | ~$2.4K/yr | ~$3–6K/yr + Okta | ~$6–12/usr/mo | ~$3–5K/yr | ~$5–10/usr/mo | ~$3.6K/yr | ~$8–15K/yr | ~$4–8/usr/mo |

**Legend:** ✅ Strong | ⚠️ Partial | ❌ Not available

**Key Takeaway:** Lumos has the best UX and access automation for SMBs, but lacks compliance evidence (SOC 2/ISO 27001) and shadow IT remediation. No competitor has a mobile app.

---

### 1.2 Top SMB Pain Points (Validated, Ranked)

| Rank | Pain | Evidence | ROI |
|---|---|---|---|
| **#1** | **Orphaned access after employee departure** | 69% of orgs have breach/incident from current/former employee (Ponemon 2023); avg revocation time 3–5 business days | $4.99M avg incident cost (IBM 2024) |
| **#2** | **Shadow IT / unauthorized SaaS proliferation** | Avg SMB employee uses 4.5 unauthorized apps (Nudge Security 2024); 63% of sensitive data in unsanctioned apps (Netskope 2024); AI tools (ChatGPT, Claude) uploading customer data | Direct GDPR / ISO 27001 liability |
| **#3** | **Access sprawl / excessive permissions** | Avg employee has 17+ SaaS apps; only 36% of licenses actively used (BetterCloud 2024); "just give admin" is SMB default | SOC 2 CC6.3 / ISO 27001 A.9 violation |
| **#4** | **No compliance-ready audit trail** | Manual evidence compilation takes 40–80 hrs/audit; 73% of SMBs fail first SOC 2 attempt (Vanta 2023) | Blocks enterprise sales contracts |
| **#5** | **Provisioning delays + over-provisioning on onboarding** | Avg new employee waits 1.4 days for correct access; "copy previous person's access" = inheriting 2 years of access creep | Security debt compound over time |

---

### 1.3 Market Gaps (White Space Analysis)

| Gap | Size | Description |
|---|---|---|
| **Offboarding automation at SMB pricing** | Critical | No compliance tool does multi-provider revocation <5 min. BetterCloud/Zluri do it but cost $6–12/user/mo. Okta requires Okta as IdP. |
| **Shadow IT discover → remediate loop** | High | Nudge Security discovers but cannot remediate. BetterCloud remediates but is IT ops, not security-focused. No affordable tool closes the loop. |
| **Unified platform (compliance + access + shadow IT)** | High | SMBs need Vanta + Nudge Security + BetterCloud = $23–32K/year. No single tool at <$6/user/month. |
| **Non-expert UX** | High | All tools except Lumos require dedicated IT admin. No tool for "dev who spends 20% time on IT". |
| **Mobile app for incident response** | Medium | Zero competitors. IT admin cannot respond to breach at 10pm from phone. |
| **Compliance evidence + access automation bundled** | High | Lumos has access automation but no SOC 2 evidence. Vanta has SOC 2 but no access automation. |

---

### 1.4 Competitive Positioning

```
High Compliance ↑
                │
   Vanta  Drata │
                │         [SMESec Target Zone]
                │         Compliance + Access +
   ─────────────┼──────────Shadow IT───────────→ High Access Automation
                │                      Lumos
                │
  Nudge Security│
(Shadow IT only)│
                │
Low Compliance  ↓
```

**SMESec differentiation against Lumos (the real benchmark):**
1. SOC 2/ISO 27001 compliance evidence (Lumos has none)
2. Shadow IT remediation — allow-list enforcement (Lumos discovers but doesn't enforce)
3. Mobile app for incident response (no competitor has this)
4. Lower price target: $3–5/user/mo vs Lumos $5–10/user/mo

---

## Part 2 — Feature Requirements (Converged)

### 2.1 Non-Negotiable Features (Will Not Pay Without These)

These 5 features are what SMB customers evaluate in their first 30 days. Without them, no contract:

| # | Feature | Minimum Viable Form | Why Non-Negotiable |
|---|---|---|---|
| **1** | **Complete "who has access to what" inventory** | Live data across Google Workspace + M365; users, groups, OAuth app grants | Every post-incident, every audit, every onboarding starts here. Without this, everything else is guesswork. |
| **2** | **Automated offboarding: Google + M365 in <5 min** | Human-initiated ("mark as leaver") → automated parallel revocation → PDF report with results | #1 demonstrable ROI in the market. Must be automated, not just a checklist. Must work for the 2 most common IdPs. |
| **3** | **Shadow IT: OAuth discovery + allow-list management** | IT admin sees all OAuth apps → reviews → marks Approved / Blocked / Pending → blocked apps trigger revocation via API | Closes the discover → act loop no competitor does affordably. Discovery-only (Nudge Security model) is not sufficient. |
| **4** | **Exportable audit trail (PDF + CSV)** | Timestamped access events, offboarding records, allow-list decisions — exportable on demand | Required for SOC 2 Type I/II access control evidence. Blocks enterprise sales without it. |
| **5** | **Real-time alerts** | New OAuth app authorized; offboarding incomplete after 30 min; MFA disabled for admin user | These 2 alerts alone justify the subscription for SMBs with no security team. Without alerts, customers never log in. |

---

### 2.2 Revised Feature Set (Post-Iteration)

#### 🔴 Must-Have v1 (Launch Gate)

| Feature | Agent Consensus | Notes |
|---|---|---|
| Asset inventory: Google Workspace users, groups, OAuth apps | Unanimous ✅ | Foundation for everything; "wow" moment on Day 1 |
| Asset inventory: M365 users, licenses, OAuth grants | Unanimous ✅ | Paired with Google; covers 85–90% of SMB IdP market |
| Asset inventory: Slack + AWS (discovery/read-only) | Unanimous ✅ | Display in inventory; NOT deprovisioning in v1 |
| OAuth / Shadow IT discovery (70% coverage via APIs) | Unanimous ✅ | TA: Google Admin SDK + MS Graph covers ~70%; API keys are permanent blind spot |
| Shadow IT allow-list management (human review, approve/block) | Unanimous ✅ | No auto-scoring; no auto-blocking. IT admin reviews and decides. |
| Automated offboarding: Google + M365 parallel revocation | PO+TA ✅ / PM conceded | Non-atomic saga pattern; per-provider failure isolation; human-confirmed before execution |
| Offboarding checklist: Slack + AWS + GitHub (manual, deep-links) | Unanimous ✅ | Audit trail completeness without over-promising; direct console deep-links per provider |
| Offboarding PDF report (what was revoked, by whom, when, what failed) | Unanimous ✅ | SOC 2 evidence artifact; closes audit loop |
| MFA coverage check (per user: MFA on/off, SSO enforced) | Unanimous ✅ | Table stakes; drives immediate customer action |
| Real-time alerts: new OAuth app, incomplete offboarding, admin MFA off | Unanimous ✅ | Low engineering cost, high daily value |
| Compliance findings dashboard (deterministic, binary findings) | TA + PM replaced risk scoring | 6 rule-based findings mapped to SOC 2 / ISO 27001 controls; see §3.3 |
| Audit trail: timestamped events, exportable PDF/CSV | Unanimous ✅ | Designed with SOC 2 auditor input from Sprint 1 |
| Multi-tenant data model with row-level isolation | PM critical path | Must be built from day 1; retrofitting = 6–8 weeks |
| SOC 2 Type 1 process: start Week 1 | PM critical path | Prerequisites: auditor selected by W6; controls operational from W8 |

#### 🟡 Should-Have v1 (Strong Differentiators — Include If Capacity)

| Feature | Notes |
|---|---|
| RBAC engine (Admin/Manager/Employee/Contractor/Service Account) | DB-native PostgreSQL RLS for platform RBAC; OPA for customer policy evaluation in JIT + allow-list |
| JIT access (Google Workspace group-based, admin-initiated, fixed durations: 1h/4h/8h/24h) | Sprint 7 if capacity; v2 first item if not. NOT self-service portal. NOT AWS IAM JIT. |
| Access review workflows (periodic, manager attestation, evidence export) | Compliance evidence requirement; 2–3 weeks effort |
| Onboarding provisioning templates by role | Paired with offboarding for full lifecycle; prevents access creep from day 1 |
| Compliance report: ISO 27001, GDPR, SOC 2 lite (one-click export) | Revenue driver for Growth/Enterprise tier |
| Mobile app: JIT approve/deny + push notifications + read-only inventory (Android + iOS) | Already funded in sprint plan; zero competitors; incident response differentiator |
| Dependency mapping: who uses what, blast radius per user | Needed before automated actions; prevents business disruption from offboarding |
| Self-service access request portal | Lumos's core differentiator; competitive gap; consider after JIT is stable |

#### 🟢 Defer to v2 (Explicitly Out of v1 Scope)

| Feature | Reason | Estimated Effort |
|---|---|---|
| AWS IAM automated deprovisioning | TA: 4–6 weeks alone; multiple resource types; non-atomic | 4–6 weeks |
| Slack automated deprovisioning | Requires Enterprise Grid ($12.50+/user/mo); ~60% SMBs don't qualify | 1–2 weeks after qualification check |
| GitHub SCIM deprovisioning | Requires GitHub Enterprise Cloud ($21/user/mo) | 1 week |
| ML-based risk scoring | False positive hell in first 2 weeks without behavioral baseline (60+ days data) | 16–20 weeks (v3) |
| HRIS integration (BambooHR, Rippling, Gusto) | 60% of SMBs don't have HRIS; 5+ connectors = 4–6 month detour; use manual trigger | 6–10 weeks |
| Periodic access reviews (formal campaign engine) | SOC 2 Type 2 requirement, not Type 1; 2–3 weeks; after core offboarding is stable | 2–3 weeks |
| Custom workflow builder | 80% of customers use defaults; custom builder is platform play | 12–16 weeks |
| Behavioral analytics / UEBA | 60+ day cold start; no day-1 value; separate ML pipeline required | 16–20 weeks |
| PAM (Privileged Access Management) | Separate product category; enterprise-only | — |
| White-label / multi-tenant MSP portal | Channel strategy, not v1 | — |

#### ❌ Removed Entirely

| Feature | Reason |
|---|---|
| Auto provisioning (SCIM inbound) | SMBs disable it; creates more problems than it solves; over-provisioning risk |
| RBAC role suggestions / role mining | PM: never acted on; creates noise; damages trust |
| Automated shadow IT revocation (no human confirmation) | Blast radius risk is existential; one mis-revocation = instant churn |
| Risk score (0–100 composite) | Replaced by compliance findings (deterministic, zero false positives) |

---

### 2.3 Compliance Findings Dashboard (Replaces Risk Scoring)

Instead of a composite risk score (which generates false positives before baseline data exists), ship **deterministic compliance findings** mapped directly to SOC 2 / ISO 27001 controls:

| Finding | Detection Rule | False Positive Rate | Control |
|---|---|---|---|
| User without MFA | `mfa_enabled = false` | 0% — factual | SOC 2 CC6.1, ISO A.9.4 |
| Inactive admin account (>90 days no login) | `last_login < now()-90d AND role = admin` | <1% | SOC 2 CC6.2 |
| OAuth grant for blocked app still active | `app.status = BLOCKED AND grant.active = true` | 0% | SOC 2 CC6.3 |
| Suspended user with active OAuth grants | `user.suspended = true AND grant.count > 0` | 0% | SOC 2 CC6.8 |
| Admin account without onboarding record | `role = admin AND onboarding_record = null` | ~2% | SOC 2 CC6.1 |
| Offboarding incomplete after 24 hours | `offboarding.status = IN_PROGRESS AND started > 24h ago` | 0% — factual | ISO A.7.3 |

Each finding has: direct remediation button, SOC 2 control mapping, fix-it guidance. This is what auditors actually review — and it is achievable on day 1.

---

## Part 3 — Technical Architecture (Converged)

### 3.1 Integration Priority Map

| Provider | Market Share (SMB) | API Type | Offboarding Capability | Shadow IT Discoverability | Priority | v1 Scope |
|---|---|---|---|---|---|---|
| **Google Workspace** | ~40% | REST (Admin SDK) | ✅ Suspend + revoke tokens | ✅ Reports API (OAuth grants) | P0 | Full automation |
| **Microsoft 365** | ~50% | Microsoft Graph | ✅ Disable + revoke sessions | ✅ `oauth2PermissionGrants` | P0 | Full automation |
| **GitHub** (free org) | ~60% SMB devs | REST | ✅ Remove org member (1 API call) | ✅ OAuth apps via API | P1 | Auto (1 API call) + checklist |
| **Slack** (free/Pro) | ~75% | REST | ❌ No deactivation API | ✅ (paid tier only) | P1 | Discovery + manual checklist |
| **Slack** (Business+/Grid) | ~40% of Slack users | REST + SCIM | ✅ SCIM or admin API | ✅ | P1.5 | Conditional automation |
| **AWS IAM** | ~40% SMB devs | AWS SDK | ⚠️ Complex (multiple types) | ⚠️ Partial (CloudTrail) | P2 | Discovery only; checklist |
| **Salesforce** | ~20% | REST + SCIM | ✅ SCIM deactivate | ✅ Connected Apps | P2 | v2 |
| **Jira/Confluence** | ~50% | REST | ✅ Deactivate user | ✅ Atlassian OAuth | P2 | v2 |

**Critical Slack caveat:** Free and Pro Slack plans have NO programmatic user deactivation API. `admin.users.*` methods require Business+ or Enterprise Grid. Product must disclose this clearly — show "Slack (manual — upgrade to automate)" in offboarding report.

**Critical GitHub caveat:** SCIM requires GitHub Enterprise Cloud ($21/user/mo). Free org member removal via `DELETE /orgs/{org}/members/{username}` works fine and should be automated.

---

### 3.2 Architecture Patterns for SMB (Converged)

**Pattern 1: Pull-based polling + hybrid webhooks**
- Baseline: scheduled polling every 15 min (self-healing, handles API downtime)
- Real-time layer: Google Pub/Sub push subscriptions + Microsoft Graph change notifications registered programmatically during OAuth consent (no customer webhook configuration required)
- Alert latency: <1 min for new OAuth app or user events (via webhooks); 15 min max fallback

**Pattern 2: Sensible defaults, zero configuration**
- IT admin connects via OAuth2 consent (2 clicks per provider)
- System immediately discovers all users + OAuth grants
- Pre-built policies operate out of the box (no policy builder required)
- Adjustable thresholds (90-day inactive admin, not a policy editor)

**Pattern 3: Human-confirmed destructive actions**
- Every revocation (offboarding or allow-list block) requires explicit confirmation
- Bulk revocation shows blast radius: affected users by name before confirmation
- Stripe-style confirmation dialog for high-impact actions: "Type app name to confirm"
- No auto-revocations in v1. Zero exceptions.

**Pattern 4: Saga pattern for offboarding (non-atomic)**
- Each provider step commits independently
- Per-step failure recorded; alert admin; single-click retry for failed steps
- Never roll back successful suspensions on partial failure
- PDF report shows per-provider status: ✅ Automated | ⚠️ Manual (link) | ❌ Failed (retry)

**Pattern 5: Graceful degradation per provider**
- Each integration has independent health status
- Product continues operating when one provider is unhealthy
- Dashboard shows: `Google Workspace ✅ synced 12 min ago | M365 ⚠️ throttled, retry in 3 min | GitHub ❌ token expired — reconnect`

---

### 3.3 Build vs. Buy (Converged)

| Component | Decision | Rationale |
|---|---|---|
| Auth (SSO login for SMESec platform) | **Buy: Keycloak** (self-hosted) | Never build OAuth/OIDC; Keycloak = zero per-MAU cost |
| RBAC (platform-level: who can use SMESec features) | **Build: PostgreSQL RLS** | Simple `workspace_id` + `role` columns; sufficient for 95% of SMB use cases; no runtime dependency |
| Policy evaluation (JIT rules, allow-list logic) | **Build: OPA on ECS** | Policy-as-code enables auditable, versioned access policies; OPA lightweight container |
| Offboarding workflow orchestration | **Buy: AWS Step Functions** | Durable execution, compensating transactions, parallel states; zero ops overhead vs Temporal for this team size |
| Background jobs / scheduled sync | **Buy: SQS + ECS workers** | Dead-letter queues, at-least-once delivery, visibility timeout built-in |
| PDF report generation | **Build: Puppeteer** | HTML→PDF, full control over template; no 3rd-party SaaS dependency |
| Access graph (user → app → resource) | **Build: PostgreSQL recursive CTEs** | Sufficient to 50K nodes; plan Neptune/AGE migration path at 100K+ nodes |
| Secrets (provider OAuth tokens, service account keys) | **Buy: AWS Secrets Manager** | KMS rotation, audit, never .env files |
| Notifications (email + Slack alerts) | **Build thin wrapper: SES + Slack Webhooks** | SES for email; Slack incoming webhooks for Slack alerts; full Slack bot in v2 |
| SCIM server (inbound: if needed for Okta customers) | **Build on library** | Use `node-scim` or `scim2` library; don't implement RFC 7643 from scratch |

---

### 3.4 Technical Non-Negotiables (Must Be Right from Day 1)

1. **Multi-tenant isolation from Sprint 1** — `workspace_id` on every table, PostgreSQL RLS policies, never skip. Retrofitting = 6–8 weeks of rework.
2. **Incremental sync with delta tokens** — Full sync at 100 tenants × 500 users × 4 providers = 200K API calls/sync cycle. `$deltaLink` (M365) and `nextPageToken` (Google) are architecturally required, not optional.
3. **SOC 2 audit log schema designed before Sprint 2** — Evidence must capture: who did what, to whom, when, why (approval record). Retrofitting evidence format for auditors is expensive.
4. **Webhook renewal automation from Sprint 4** — Google Push Notifications expire in 7 days; Microsoft Graph subscriptions expire in 3 days–3 years. Silent expiry = alerting down without warning. CloudWatch alarm required: "webhook expiry in <24h."
5. **M365 token propagation caveat disclosed** — `POST revokeSignInSessions` revokes refresh tokens; access tokens remain valid up to 60 min. SLA language: "access revoked within 5 min; existing sessions expire within [X] min." Customers with Azure AD P1 can set token lifetime to 15 min.
6. **Service account detection before offboarding** — Before executing offboarding, check: does this user own any service accounts / app credentials? Flag for manual review. One undetected service account dependency can cause production outage.
7. **Google OAuth verification started Week 1** — 4–6 week lead time for Google's OAuth app review. Without it, consent screen shows security warnings that kill SMB conversions. Block marketing demos until verified.

---

## Part 4 — Delivery Reality Check (Converged)

### 4.1 Market Timeline Benchmarks

| Company | Scope | Time to Payable MVP |
|---|---|---|
| Nudge Security | Observation-only (no automation) | ~15 months ← fastest |
| Zluri | SaaS discovery + basic governance | ~14 months |
| Torii | Browser extension discovery | ~12 months |
| Lumos | Full access governance (requests, deprovisioning, reviews) | ~18 months |
| BetterCloud | Full SaaS management | ~24 months |
| Veza | Enterprise access graph | ~24+ months |

**Key lesson:** Nudge Security reached payable MVP fastest by deliberate scope cuts (no write-back automation, no workflow engine, no HRIS). Their "observation-only" model created immediate value but limited their total addressable market.

**SMESec advantage:** By committing to Google + M365 automated offboarding in v1 (not observation-only), SMESec starts with a stronger value proposition than Nudge Security at launch — but must execute the integration correctly.

### 4.2 Feature Value vs Delivery Effort

```
                    HIGH VALUE
                         │
         ┌───────────────┼───────────────┐
    Q2   │  OAuth/Shadow │  Asset        │  Q1
  QUICK  │  IT Allow-List│  Inventory    │  STRATEGIC
   WINS  │               │  (Google+M365)│
         │  Compliance   │               │
         │  Findings     │  Offboarding  │
         │  Dashboard    │  (G+M365)     │
─────────┼───────────────┼───────────────┼─────────
 LOW     │               │               │  HIGH
 EFFORT  │  Real-time    │  JIT Access   │  EFFORT
         │  Alerts       │  AWS Offboard │
    Q3   │  MFA Check    │  HRIS Sync    │  Q4
  LOW    │               │  Workflow     │  LONG-TERM
  VALUE  │               │  Engine       │
         │               │               │
         └───────────────┴───────────────┘
                    LOW VALUE
```

**Build in this order:** Q1 (strategic) → Q2 (quick wins) → Q1 strategic continuation → Q4 only after revenue validates need.

### 4.3 Common Failure Modes to Avoid

| Failure | Frequency | Mitigation |
|---|---|---|
| Integration breadth over depth | Very High | Commit to Google + M365 depth-first; no 3rd provider until both are solid |
| Auto-revocation before customer trust established | Very High | Human confirmation required for ALL write operations in v1, no exceptions |
| AWS deprovisioning in v1 | High | Explicitly deferred; must not re-enter v1 scope |
| HRIS dependency trap | High | Manual "mark as leaver" trigger only; HRIS in v2 after validating which systems customers actually use |
| Multi-tenant underestimation | Medium-High | Row-level security from Sprint 1 schema design |
| Compliance theater (reports don't close the loop) | High | Audit trail must record: decision made → action taken → verified. Must close the loop. |
| SOC 2 process started too late | High | Start Week 1; auditor engaged by W6; observation period begins W8 |

### 4.4 Trust Infrastructure (Critical Path — Non-Feature)

These are prerequisites for the first paying customer. They are NOT features. They are product readiness requirements:

| Item | Timeline | Cost | Blocking? |
|---|---|---|---|
| Google OAuth app verification | 4–6 weeks lead; start Week 1 | Free | Yes — demo consent screen |
| Data Processing Agreement (DPA) | 2–4 weeks with legal | $5–10K legal | Yes — GDPR compliance |
| SOC 2 Type 1 audit kickoff | W1 (observation); W23 (audit) | $15–40K | Yes — first enterprise customer |
| Penetration test | W16 engage; W24 results | $10–20K | Yes — 70% of SMB customers require before admin access |
| Cyber liability insurance | 2–3 weeks | $5–10K/yr | No — but enterprise-adjacent customers require |

---

## Part 5 — Open Questions (Require Team Decision)

| # | Question | Owner | Blocking? | Deadline |
|---|---|---|---|---|
| 1 | **Google + M365 automated offboarding: is this Sprint 6 launch gate or v2?** PO says v1 non-negotiable. PM accepted this with human-confirmation requirement. Needs formal alignment sign-off from all stakeholders. | PO + PM | **Yes** | Sprint 1 end |
| 2 | **Marketing language for offboarding SLA.** "Automated offboarding across all providers" vs "Google + M365 automated in <5 min, full audit trail for others." Must be locked before external demos. | PM + PO | **Yes** | Sprint 4 end |
| 3 | **Pilot customer qualification.** Does the first paying customer need to waive SOC 2 Type 1 requirement, or do we find design partners who accept early access in exchange? 3 signed LOIs needed by Sprint 8. | PM | **Yes** | Sprint 8 |
| 4 | **Google OAuth app verification.** 4–6 week lead time. Does this block Sprint 2 production testing? Do we build against test tenants until verified? | TA + PM | **Yes — Sprint 2** | Week 1 |
| 5 | **Slack customer qualification.** What % of pilot customers have Slack Business+ or Enterprise Grid? Determines whether Slack automation is a v1.5 feature or a sales disqualifier. | PM | No | Sprint 4 |
| 6 | **M365 token propagation caveat.** SLA documentation: "sessions expire within [X] min" — what is the disclosed [X]? Requires Azure AD P1 ($6/user/mo) for 15-min tokens. Is this a customer requirement? | TA + PO | No | Sprint 6 |
| 7 | **JIT Access: Sprint 7 or v2?** PO: include if capacity allows. PM: first v2 item. Decision gate = Sprint 5 utilization review. | PM | No | Sprint 5 end |
| 8 | **SOC 2 budget.** $15–40K for Type 1 audit not currently in sprint plan budget. Needs confirmation. | PM | **Yes** | Week 1 |

---

## Summary: Feature Requirements for Access Governance

### The 3-Sentence Product Requirement

> SMESec Access Governance solves the #1 unmet SMB security problem: **orphaned access after employee departure** — by automating Google Workspace + M365 revocation in under 5 minutes with a human-initiated trigger, a per-provider audit trail, and a checklist for all other systems. It simultaneously closes the **shadow IT discovery-to-action loop** no competitor closes at SMB pricing: discover all OAuth apps, let IT admin approve/block, track the policy. The entire product is designed to be operated by a developer spending 20% of their time on IT — no security expertise required, no webhook configuration, no policy language to learn.

### Feature Tiers (Final)

| Tier | Features |
|---|---|
| **Must-Have** | Asset inventory (G+M365+Slack+AWS discovery), OAuth/shadow IT discovery + allow-list management, automated offboarding (G+M365, human-initiated), offboarding checklist (Slack+AWS+GitHub, manual with deep-links), PDF audit report, MFA coverage check, compliance findings dashboard (6 deterministic rules), real-time alerts (new OAuth app, incomplete offboarding, admin MFA off) |
| **Should-Have** | RBAC engine (DB-native + OPA for policy eval), JIT access (Google Workspace groups, admin-initiated), access review workflow (manager attestation, evidence export), compliance reports (SOC 2/ISO 27001 one-click), mobile app (JIT approve + push notifications + read-only), onboarding templates by role |
| **v2** | AWS IAM deprovisioning, Slack deprovisioning (Business+), access review campaign engine, HRIS integration, GitHub SCIM, self-service access request portal |
| **Removed** | Auto-provisioning, RBAC role suggestions, automated shadow IT revocation, ML-based risk scores |

---

*Generated by 3-agent research loop: Product Owner + Technical Advisor + Project Manager × 2 rounds (6 agent invocations total). Date: 2026-05-28.*
