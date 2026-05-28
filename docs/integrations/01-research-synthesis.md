# Research Synthesis: Integrations Key Requirement
## Google Workspace + Microsoft 365 + Slack

**Date:** 2026-05-28  
**Process:** 5-round multi-agent debate (Product Owner ↔ Technical Advisor ↔ Project Manager)  
**Status:** Final — Approved with conditions

---

## 1. Competitor Landscape

| Capability | SMESec v1 | Vanta | Drata | Secureframe | Nudge Security | JupiterOne |
|---|---|---|---|---|---|---|
| Google Workspace sync | ✅ Full | ✅ Users/groups | ✅ Users/groups | ✅ Users/groups | ✅ OAuth apps + users | ✅ Deep graph |
| M365 sync | ✅ Full | ✅ Users/groups | ✅ Users/groups | ✅ Users/groups | ✅ OAuth apps + users | ✅ Deep graph |
| Slack sync | ✅ Tiered | ⚠️ Members only | ⚠️ Members only | ⚠️ Members only | ⚠️ App inventory | ✅ Full |
| Automated offboarding | ✅ Google+M365 full; Slack tiered | ❌ Manual checklist | ❌ Manual checklist | ❌ Manual checklist | ❌ | ❌ |
| Shadow IT detection | ✅ OAuth apps all 3 | ⚠️ Manual review | ⚠️ Manual review | ❌ | ✅ Best-in-class | ✅ |
| Guided onboarding wizard | ✅ <30 min | ❌ Manual docs | ❌ Manual docs | ❌ Manual docs | ✅ Good UX | ❌ |
| Compliance evidence | ✅ | ✅ Primary focus | ✅ Primary focus | ✅ Primary focus | ❌ | ⚠️ |
| AI threat detection | ✅ (Track 2) | ❌ | ❌ | ❌ | ❌ | ❌ |
| SME pricing | $200-5K/mo | $3K-15K/yr | $3K-15K/yr | $3K-15K/yr | $5K-20K/yr | $5K-25K/yr |

**Market gap SMESec owns:** Automated offboarding + shadow IT detection + compliance evidence + AI detection — single unified platform at SME pricing. No competitor combines all four.

---

## 2. Customer Pain Points (Priority Order)

| # | Pain Point | Severity |
|---|---|---|
| 1 | Invisible OAuth shadow apps (140+ apps connected, no visibility) | Critical |
| 2 | Offboarding takes hours/days — ex-employees retain access | Critical |
| 3 | Setup requires developer / CTO — competitors take 2-4 hours minimum | High |
| 4 | Permission consent screens trigger legal objection ("read all email") | High |
| 5 | Daily sync → alerts arrive 18 hours after incident | High |
| 6 | No cross-provider view — 3 separate admin consoles | High |
| 7 | Silent sync failures — integration stops, nobody notices for weeks | Medium |
| 8 | Manual compliance evidence export before every audit | Medium |

---

## 3. Final Feature Requirements — v1

### 3.1 Must-Have (Commit)

| # | Feature | Customer Value | Sprint |
|---|---|---|---|
| 1 | **Guided OAuth wizard** (<30 min, 3 providers, no developer) | "Under 30 minutes, no consultants required" vs competitors' 2-4 hours | S2-S5 |
| 2 | **15-min incremental sync** (Google: full pull+diff+Workspace Events API; M365: delta link+webhook; Slack: Events API+poll) | Detect shadow app in 15 min, not 18 hours | S2-S5 |
| 3 | **Cross-provider identity matching** (email-based canonical; "possible duplicate" flag for mismatches) | Unified risk view per person across all 3 providers | S3 |
| 4 | **Minimum-permission OAuth scopes** with plain-English explainers (never read email/file content) | Overcomes legal consent objection | S2-S5 |
| 5 | **Sync health monitoring** (auth failure, scope degradation, webhook expiry alerts) | Zero silent failures — always know integration is healthy | S4 |
| 6 | **Google + M365 automated offboarding** (parallel, <5 min account disable; M365 sessions drain 60 min — disclosed) | "Disable access across Google and M365 in under 5 minutes" | S6 |
| 7 | **Dry-run + 2-step confirmation** for ALL write operations (hard gate — no bypass) | "Preview what will happen before any access is revoked" | S6 |
| 8 | **Slack tier detection + feature gating** | Transparent capability boundary, upgrade path shown | S5 |
| 9 | **Slack Business+ automated offboarding** | Full automation when plan supports it | S6 (conditional on API review) |
| 10 | **Slack Free/Pro guided offboarding** (deep-link + step-by-step + audit log) | Better than all competitors (which offer nothing) | S6 |
| 11 | **Shadow IT detection** (OAuth app inventory, new app alert <15 min) | "Alerted within 15 minutes when employee authorises unapproved app" | S4 |
| 12 | **Immutable audit log** (PostgreSQL append-only + S3 archival) | ISO 27001 A.12.4 evidence, tamper-proof | S5 (hard gate before offboarding) |
| 13 | **data_residency_region schema column** (Sprint 1 — 0.5 days) | EU GDPR compliance, blocks EU launch if missing | S1 |
| 14 | **Per-person risk indicators** per provider (not composite score yet) | Simple status: active/suspended/no-MFA/excessive-permissions | S4 |
| 15 | **Composite risk score per person** (weighted sum across all providers) | Enterprise tier differentiator | S6 (conditional on data quality) |

### 3.2 Conditional (Require Pilot Data)

| Feature | Condition | Fallback |
|---|---|---|
| Slack Business+ automated offboarding | Slack API admin review approved by Sprint 5 + pilot customers have Business+ plan | Free/Pro guided path (always ships) |
| AWS offboarding | >50% pilot customers have AWS + >20 IAM users per tenant (Sprint 5 go/no-go) | Deferred to v1.1 |
| Composite risk score | Signal coverage ≥80% of assets (Sprint 6 evaluation) | Per-provider status indicators only |

### 3.3 Deferred to v1.1 (Month 7-8)

- M365 Continuous Access Evaluation (CAE) for near-real-time session revocation
- Google Workspace Events API full integration
- AWS offboarding (IAM disable + access key revocation)
- Multi-region EU infrastructure deployment (schema is ready from Sprint 1)
- M365 webhook renewal hardening with full CloudWatch coverage

### 3.4 Deferred to v2

- SCIM provisioning (SMESec as source of truth)
- File content scanning (Drive, SharePoint — consent risk + scope creep)
- Slack DLP (requires Enterprise Grid)
- Webhook real-time push for all providers (15-min polling delivers 90% value at 10% complexity)

### 3.5 Cut List (PO + TA + PM unanimous)

| Feature | Reason |
|---|---|
| `<15 min onboarding` as a contractual SLA | Not achievable — replace with "guided wizard, typical 20-30 min" |
| `<5 min offboarding` for M365 as universal SLA | False — M365 sessions take up to 60 min to drain. Disclose transparently. |
| Step Functions orchestration for v1 | Over-engineered — Go errgroup + DB state table achieves same result |
| Full Object Lock WORM in Sprint 5 | PostgreSQL append-only + S3 archival is sufficient for v1 audit compliance |
| Broad OAuth scopes ("read all email") | Anti-feature that kills consent and triggers GDPR objections |
| Silent "offboarding complete" when partial failure | Trust destroyer — never show green if any provider failed |

---

## 4. Technical Architecture Decisions (Finalized)

### 4.1 Authentication per Provider

| Provider | Auth Method | Rationale |
|---|---|---|
| Google Workspace | 3-legged OAuth 2.0 (admin consent) | Service account + DWD requires manual Google Admin Console step that cannot be automated. 3-legged OAuth is fully wizard-able. |
| Microsoft 365 | App registration, client credentials (application permissions) | Background service, no user login required, admin consent once for tenant |
| Slack | OAuth 2.0, bot token (all tiers) + admin token (Business+ only) | Two separate apps registered: bot (standard) + admin (requires Slack API review) |

### 4.2 Sync Mechanism per Provider

| Provider | Primary Trigger | Delta Mechanism | Fallback |
|---|---|---|---|
| Google | 15-min cron poll | `syncToken` from Directory API (users, groups). Reports API: `startTime` filter for audit events | Full resync on 410 Gone |
| M365 | Graph webhook change notification | `$deltaLink` for `/users/delta`, `/groups/delta`, `/applications/delta` | 15-min poll when webhook inactive. Webhook subscription auto-renewed every 2 days (max 3-day expiry). |
| Slack | Events API webhook | No native delta — hash-based manifest diff. Webhook events trigger targeted refresh | 15-min poll with cursor pagination |

**Key correction (TA Round 2):** Google Directory API has **no delta query support** — `syncToken` is available only for `users.list` and returns changed records, but `tokens.list` (OAuth app grants) requires per-user enumeration (1 API call per user per sync cycle). For large tenants (500+ users), this may exceed 15-min window — monitor in pilot.

### 4.3 Token Storage

```
AWS Secrets Manager path format:
  /smesec/{workspace_id}/google/oauth_credentials
  /smesec/{workspace_id}/m365/app_credentials
  /smesec/{workspace_id}/slack/bot_token
  /smesec/{workspace_id}/slack/admin_token  (Business+ only)
  /smesec/{workspace_id}/slack/signing_secret

Security model:
  - Per-tenant CMK (AWS KMS Customer Managed Key)
  - Token decrypted in memory only at API call time — never logged, never in DB
  - RDS stores only the Secrets Manager ARN per tenant-provider pair
  - ECS task IAM role scoped to own tenant's secrets only
  - 90-day rotation for long-lived credentials
```

### 4.4 Write Operation Protection

```
Two scoped clients per provider:
  ReadClient:  read-only scopes — used by sync engine always
  WriteClient: write scopes — constructed ONLY during offboarding, gated by:
    (1) offboarding_jobs.status = 'awaiting_confirm'
    (2) Initiating user has RBAC role >= Manager
    (3) Dry-run result validated against current state

Hard-coded allow-list of permitted write operations (no other writes possible):
  Google: suspend account, revoke OAuth token, unsuspend (rollback only)
  M365:   disable account, revoke sessions, invalidate refresh tokens, re-enable (rollback only)
  Slack:  setInactive (Business+ only)
```

### 4.5 Offboarding Engine

```
Pattern: Go errgroup (parallel, partial failure tolerant)
  - All providers run concurrently
  - One provider failure does NOT cancel others
  - Per-provider status tracked in offboarding_tasks table
  - No "complete" status unless all providers confirm

Dry-run → 2-step confirm → Execute:
  Step 1: POST /offboarding/{id}/dry-run → returns predicted actions + confirm_token (30 min TTL)
  Step 2: POST /offboarding/{id}/confirm → validates token → executes
  
Partial failure handling:
  - status = 'partial_failure'
  - Immediate alert to IT admin with deep-link to manual fix
  - Manual "Resume" button retries failed providers only
  - Completed providers are idempotent (skipped on retry)
  - PDF report generated regardless, marking failures in RED

Slack tier branching:
  Business+: admin.users.setInactive via API
  Free/Pro:  deep-link to admin console + guided email + audit log entry
```

### 4.6 Key Schemas

**sync_state table:**
```sql
(workspace_id, provider) PRIMARY KEY
status: idle | running | completed | failed | auth_failed | rate_limited
provider_cursor: syncToken (Google) | deltaLink (M365) | hash manifest (Slack)
consecutive_failures: INT
last_error, last_error_at
```

**offboarding_audit_events table (append-only):**
```sql
workspace_id, offboarding_id, target_user_id, provider, operation, 
status: dry_run | executed | failed | skipped
request_payload (JSONB, sanitized), response_code, executed_by, executed_at
PostgreSQL rules: NO UPDATE, NO DELETE
Nightly S3 export with Object Lock (WORM) for 7-year retention
```

**data_residency_region column (Sprint 1):**
```sql
ALTER TABLE workspaces ADD COLUMN data_residency_region TEXT DEFAULT 'us';
-- Values: 'us' | 'eu' | 'ap'
-- All secrets, sync jobs, and data routed based on this value
```

---

## 5. OAuth Scopes Summary

### Google Workspace (3-legged OAuth, admin consent)
```
admin.directory.user.readonly          — User inventory
admin.directory.group.readonly         — Group inventory
admin.directory.orgunit.readonly       — OU classification
admin.directory.user.security          — OAuth app discovery + token revocation
admin.reports.audit.readonly           — Audit events (token grants, admin changes)
admin.reports.usage.readonly           — Usage analytics

Write scope (requested separately on offboarding feature activation):
admin.directory.user                   — Suspend/unsuspend accounts
```

### Microsoft 365 (App registration, application permissions)
```
User.Read.All                          — Read all users
Group.Read.All                         — Read all groups
Directory.Read.All                     — Directory objects
AuditLog.Read.All                      — Sign-in + audit logs
Application.Read.All                   — Service principals (OAuth apps)
DelegatedPermissionGrant.ReadWrite.All — Per-user delegated permissions
User.EnableDisableAccount.All          — Disable accounts (offboarding)
UserAuthenticationMethod.ReadWrite.All — Revoke MFA sessions
Reports.Read.All                       — Usage reports
```

### Slack (OAuth 2.0)
```
Free/Pro tier (bot app):
  users:read, users:read.email, team:read, channels:read, groups:read

Business+ only (admin app — requires Slack API review):
  admin.users:read, admin.users:write, admin.apps:read, admin.conversations:read
```

---

## 6. Critical Constraint: Slack Plan Tier

**~80% of SME customers are on Slack Free or Pro plans.**  
`admin.users:write` (required for automated deactivation) requires **Slack Business+ minimum**.

| Slack Plan | SMESec Capability |
|---|---|
| Free / Pro | User discovery, app inventory (limited), guided offboarding (deep-link + audit log) |
| Business+ | Full automated offboarding, org-wide app inventory, channel admin |
| Enterprise Grid | Full audit logs, org-wide policies, DLP (v2) |

**PM action required (this week):** Survey all pilot customers: "What is your Slack plan?"  
If <2 customers have Business+, do NOT demo automated Slack offboarding at v1 launch. Demo Google + M365 offboarding only.

**UI requirement:** During Slack OAuth wizard, display plan tier and capability gate clearly. Never show "offboarding complete" if Slack task was actually a guided manual prompt.

---

## 7. M365 Session Revocation Constraint

`POST /users/{id}/revokeSignInSessions` invalidates refresh tokens immediately.  
**Active access tokens remain valid for up to 60 minutes** (Microsoft architectural constraint, not SMESec limitation).

**Customer-facing language:**
> "Account disabled immediately — no new logins possible. Active sessions clear within 60 minutes, consistent with Microsoft's architecture. Enable Conditional Access with CAE to reduce this to near-real-time."

CAE integration is a v1.1 feature. Document as an upgrade path, not a gap.

---

## 8. Onboarding Time Claim

| Old claim | Revised claim | Rationale |
|---|---|---|
| `<15 min` | `<30 min guided, no developers` | Google wizard requires domain-wide delegation confirmation step in Google Admin Console (5-10 min). M365 requires admin consent grant. Realistic for IT-literate admin: 20-30 min total. |

Competitive context: Vanta/Drata average 3-7 **days** for initial setup. "<30 min" is still 10-15x better than competition. Do not over-promise.

---

## 9. Open Questions (Require Pilot Customer Data)

| Question | Why It Matters | Action | Owner |
|---|---|---|---|
| What Slack plan are pilot customers on? | Determines automated vs guided offboarding story | Survey this week | PM |
| Is M365 CAE enabled on pilot tenants? | If yes, 60-min session window becomes near-real-time | Add to onboarding questionnaire | PM |
| Do pilot customers have AWS in their stack? | Triggers/cancels AWS offboarding in v1 | Qualify during LOI conversations | PM |
| Are any pilot customers in EU? | Cross-region infra (Sprint 4-5) cannot slip if EU customer exists | Confirm during LOI signing | PM |
| What is offboarding frequency? (events/month) | Validates ROI of offboarding feature | Discovery call question | PO |
| Google token API performance at scale? | 500 users × 20 OAuth apps = 10K API calls per sync | Measure on first 3 pilot tenants | TA |

---

## 10. Sprint Allocation Summary

| Sprint | Weeks | Key Deliverables |
|---|---|---|
| S1 | W1-2 | Infrastructure foundation + `data_residency_region` schema (0.5 days BE — non-negotiable) |
| S2 | W3-4 | Google Workspace sync (full pull + client-side diff + Workspace Events API) + OAuth wizard Step 1 |
| S3 | W5-6 | M365 sync ($deltaLink + webhook + 3-day subscription auto-renewal) + OAuth wizard Step 2 + identity matching |
| S4 | W7-8 | Sync health monitoring + scope degradation detection + shadow IT alerts + per-provider status indicators |
| S5 | W9-10 | Slack integration (Events API + poll fallback + tier gating) + OAuth wizard Step 3 + Redis distributed locking + immutable audit log MVA (PostgreSQL) |
| S6 | W11-12 | Automated offboarding (Go errgroup + per-provider DB state + dry-run + 2-step confirm + partial failure handling) + composite risk score |

**Critical path:** Sprint 5 is at 79% BE utilization. AWS cross-account IAM discovery in Sprint 5 is the most likely candidate to slip — explicitly documented as v1.1 contingency if Sprint 5 bleeds.

**External deadline:** Slack admin API review submission by end of Sprint 4 (W8). Submit early — approval takes 2-6 weeks.

---

## 11. Anti-Features (Do NOT Copy from Competitors)

| Anti-Feature | Who Does It | SMESec Approach |
|---|---|---|
| Broad OAuth scopes ("read all email") | Vanta, Drata, Secureframe | Minimum necessary scopes only; never read email/file content |
| Daily full sync | Most competitors | 15-min incremental sync |
| Read-only — no automated action | Vanta, Drata, Secureframe, Scrut | Bi-directional: detect AND act |
| Silent sync failures | Vanta, Drata | Proactive health monitoring + re-auth alerts |
| No explanation of WHY permissions are needed | All competitors | Plain-English permission explainer before every consent |
| Silent "offboarding complete" on partial failure | — | Never green unless all providers confirmed |
| Slack deactivation attempt on Free/Pro (results in silent failure) | — | Detect plan tier, route to guided path, never attempt impossible API call |

---

## 12. Differentiation Summary (Post-Constraint Validation)

Despite all constraints identified in debate, SMESec differentiation remains intact:

| Differentiator | Status After Debate |
|---|---|
| Automated Google + M365 offboarding | ✅ Intact — no competitor offers this |
| Slack offboarding (tiered) | ✅ Intact — guided path still leads market for Free/Pro (competitors offer nothing) |
| <30 min guided onboarding | ✅ Intact — 10-15x better than competition |
| Shadow IT detection (OAuth apps) | ✅ Intact |
| AI threat detection (Track 2) | ✅ Completely uncontested |
| Unified platform (compliance + access + AI) | ✅ No competitor combines all three at SME pricing |

---

## 13. Confidence Scores

| Agent | Round | Score | Key Concern |
|---|---|---|---|
| Product Owner | R1 → R4 | 8/10 → 7.5/10 | Slack plan constraint; AWS offboarding deferred |
| Technical Advisor | R2 → R5 | 5/10 → 8.5/10 | Google token scale; Slack API review outcome |
| Project Manager | R3 | 6.5/10 | Sprint 5 capacity; Slack plan distribution unknown |

**Overall: APPROVED with conditions** — 3 hard gates before Sprint 6 offboarding ships:
1. Cross-tenant RLS CI test in Sprint 1
2. Slack offboarding does not block Sprint 6 milestone if admin review pending (Free/Pro path demo-ready independently)
3. Dry-run + 2-step confirm is a hard gate on ALL write operations — no bypass
