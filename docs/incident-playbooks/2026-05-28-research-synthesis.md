# Research Synthesis: Incident Playbooks for Non-Security Staff

**Date:** 2026-05-28  
**Status:** Final — 3 Rounds Completed  
**Method:** Iterative debate — Product Owner × Technical Advisor × Project Manager  
**Bias constraint:** Zero reference to Track 1 / Track 2 implementation plans during research

---

## Executive Summary

The market offers zero products that let non-security staff (IT admins, HR managers, company owners) actually **execute** incident response. Every competitor either:
- Assumes security expertise (PagerDuty, SOAR platforms)
- Provides compliance documentation but not execution (Vanta, Drata, Secureframe)
- Outsources the problem to an external team (Huntress MDR, Arctic Wolf)

This is the gap SMESec fills. The feature is defensible, differentiated, and — if scoped correctly — buildable in v1.

The 3-round debate **significantly revised** the original plan. Key changes are documented in Section 7.

---

## 1. Market Landscape

### 1.1 Incident Management (DevOps-Oriented)

| Product | Playbook Approach | Target User | SME Non-Security Gap |
|---------|------------------|-------------|----------------------|
| **PagerDuty** | Runbooks = bash scripts + API calls | DevOps/SRE engineers | ❌ Unusable by HR/CEO — requires command-line knowledge |
| **Opsgenie** | Alert-linked runbooks | DevOps/SRE engineers | ❌ Same — technical dashboards, zero plain-language guidance |
| **FireHydrant** | Runbooks + retrospectives (Slack-first) | Engineering teams | ❌ Slack-native, technical team assumed |
| **Rootly** | Incident workflows via Slack commands | Engineering/DevOps | ❌ `/incident declare` makes no sense to an HR manager |

**Category verdict:** Every product assumes a technical on-call person. Zero UX consideration for HR managers, company owners, or non-technical IT admins.

### 1.2 Compliance + Security Platforms

| Product | Incident Response | What it Actually Is |
|---------|------------------|---------------------|
| **Vanta** | Incident tracker: create/close tickets | Compliance documentation — proves you have a plan, not a tool to execute it |
| **Drata** | Risk register + incident log | PDF policy templates — read them, still don't know what to do |
| **Secureframe** | Incident policy templates (PDF) | PDF library — the non-expert reads it, remains lost |
| **Sprinto** | Basic incident module (log + status) | Cheapest option still treats incident response as a compliance checkbox |

**Category verdict:** These solve "prove you have an IR plan." They do not solve "actually respond to an incident." A non-security IT admin using Vanta during a breach finds a PDF to read, not a wizard to follow.

### 1.3 SOAR Platforms

| Product | Approach | Price | SME Non-Security Gap |
|---------|----------|-------|----------------------|
| **CrowdStrike Falcon** | Full SOAR + MDR, 500+ integrations | $15-30/endpoint/mo | ❌ Requires Tier-2 security analysts |
| **Rapid7 InsightConnect** | Visual workflow builder, 300+ plugins | $5K-$50K+/year | ❌ Configuration requires security knowledge |
| **Swimlane** | No-code SOAR | $50K-$200K/year | ❌ "No-code" = no-code for security engineers — still needs security knowledge |
| **Tines** | Hyperautomation, API-first | $25K-$100K/year | ❌ Extremely powerful, completely wrong for SMEs |
| **Blink Ops** | AI-generated playbooks from natural language | $20K+/year | ⚠️ Closest to right direction — still targets security teams |

**Category verdict:** SOAR tools universally assume a security expert operates them. Even "no-code" SOAR requires security knowledge to configure. No SOAR product has solved for a non-security person **running** a pre-built playbook.

### 1.4 Managed Security for SMEs

| Product | Model | Self-Service Gap |
|---------|-------|------------------|
| **Huntress** | MDR — Huntress team responds on your behalf | ❌ Creates dependency: non-security staff still doesn't know what to do independently |
| **Todyl** | SASE + basic incident alerts | ⚠️ Better than most; some guided actions; still requires IT knowledge |
| **Arctic Wolf** | Managed SOC (external team) | ❌ $50K-$200K/year; outsources judgment, doesn't build capability |

**Category verdict:** Managed security outsources the problem. When the managed service alerts at 3am, the company owner still doesn't know what to do. The self-service gap for non-security staff remains unaddressed.

---

## 2. Critical Market Gaps — What No Product Has Built

These gaps exist across ALL competitor categories:

| Gap | Description |
|-----|-------------|
| **Execution, not documentation** | No product guides non-security staff to actually execute containment — they all document that you have a plan or outsource the work |
| **"What does this mean?" plain-language steps** | No product explains *why* a step matters in accessible language — all assume security vocabulary |
| **Real-time verification that human steps worked** | No product checks that a manual step was completed correctly before proceeding |
| **GDPR notification triage** | Art. 33's 72-hour clock. No product connects incident detection → regulatory notification workflow |
| **"Stop the bleeding" minimal mode** | No product offers a 3-step safe containment for uncertain non-experts while waiting for an expert |
| **BEC (Business Email Compromise) playbook** | #1 financial loss threat for SMEs ($2.9B in 2023 per FBI IC3). Zero self-service products cover it |
| **Escalation that goes somewhere specific** | No product provides concrete escalation targets (who do you call?) for non-security staff — just a generic "escalate" button going nowhere |
| **Insurance-ready evidence packaging** | No product auto-packages incident evidence into a format usable for cyber insurance claims |

---

## 3. What "Incident Playbooks for Non-Security Staff" Actually Requires

The core promise, precisely stated:

> **Within 10 minutes, an IT admin with no security background can: identify the correct incident type, launch the relevant playbook, execute all automated remediation actions across connected providers, and have an immutable audit log recording every action taken — for an Account Compromise or Phishing Response incident.**

This requires:

### 3.1 The 6 Playbooks for v1

After 3 rounds of debate, the final playbook set is:

| # | Playbook | Incident Type | API Automation Level | Priority |
|---|----------|---------------|----------------------|----------|
| 1 | **Account Compromise** | Suspicious login, credential stuffing | Full — revoke sessions, suspend account, notify | P0 |
| 2 | **Offboarding Emergency** | Termination, insider threat | Full — revoke all access across providers | P0 |
| 3 | **Unauthorized Access** | Privilege escalation, lateral movement | Partial — suspend resource access, notify manager | P1 |
| 4 | **Shadow IT Remediation** | Unapproved OAuth app, AI tool with data access | Full — revoke OAuth token, notify user | P1 |
| 5 | **Phishing Response** | Email phishing, BEC detected | Partial — quarantine, reset credentials, alert recipients | P0 (**new — missing from original plan**) |
| 6 | **GDPR Data Breach Guidance** | Personal data breach requiring Art. 33 consideration | Fact-gathering only — no automated legal determination | P1 (**new — replaces Inactive Account**) |

**Explicitly NOT in v1:**
- Ransomware Containment as an action playbook — requires MDM/EDR (not in stack); network isolation is not automatable without network admin access. Ships as "Credential Containment Guidance" page with escalation to IT/MSP.
- Inactive Account — moved to scheduled governance automation (not an incident response)
- BEC wire fraud verification — requires banking integration; v2

### 3.2 Why This Differs from Original 5 Playbooks

The original 5 were primarily **access governance** playbooks — they reflect what identity control can already automate. After market research:

| Original | Status | Reason |
|----------|--------|--------|
| Account Compromise | ✅ Keep | Core playbook |
| Offboarding Emergency | ✅ Keep | Core playbook |
| Shadow IT Detected | ✅ Keep (renamed) | Current #1 SME threat with AI tool explosion |
| Unauthorized Access | ✅ Keep | Core playbook |
| Inactive Account | ❌ Replace | This is maintenance, not incident response — move to scheduled automation |
| **Phishing Response** | ✅ Add | #1-2 SME incident type; missing from original plan |
| **GDPR Data Breach Guidance** | ✅ Add | Required to sell to EU customers; legal obligation playbook |

---

## 4. Essential Feature Set

Ordered by: **if missing, product fails its promise**.

### Tier 1 — Launch Blockers (missing = product fails)

**1. Step-by-step Wizard UI with Zero Security Jargon**
- Every decision gate answerable by a 45-year-old office manager without Googling
- Each step: WHY we're doing this + WHAT will happen + HOW to confirm it worked
- Maximum 3 human decision points per playbook
- Progress indicator: "Step 3 of 6"
- Mobile-responsive: incidents happen outside office hours
- **Validation requirement:** Every decision gate label must pass "no jargon" review before Sprint 8 ships

**2. Pre-Flight Confirmation UX (replaces undo)**
- Before any irreversible action: show name, email, photo, role, last login of affected entity
- Plain-language consequence: "Suspending this account will remove Sarah's access to email, Google Drive, Slack, and all OAuth apps immediately."
- Confirm + proceed OR cancel
- This is NOT a generic "are you sure?" — it is entity-aware and consequence-specific
- Undo removed entirely — it creates false confidence for security-critical actions

**3. Concrete Escalation with Pre-Configured Contacts**
- Onboarding captures: Security Lead, Legal Contact, CEO/CTO, IT Admin backup, HR Contact
- Every decision gate that requires expertise beyond IT admin authority surfaces the SPECIFIC person + pre-filled message
- Escalation type is declared per-playbook (legal, technical, business, compliance)
- No generic "escalate" button — specific contact for specific decision type
- Hard dependency: playbook blocks at first escalation gate if onboarding contacts are not configured

**4. Automated Actions That Actually Work**
- Account suspension, session revocation, OAuth app removal execute in <2 minutes
- Parallel execution across all connected providers (Google Workspace, M365, Slack, AWS)
- Step-level status: per-provider, not per-playbook
- Partial failure handling: show which providers succeeded and which failed, with "Retry [provider]" button for failed steps only (not full playbook re-run)
- M365 caveat displayed explicitly: modern auth clients revoked in <5 min; legacy auth clients up to 60 min

**5. Immutable Audit Log**
- Every action: what, who, when, result, API call made, provider response
- S3 Object Lock COMPLIANCE mode, minimum 3 years (7-year option for Enterprise)
- SHA-256 hash of evidence JSON stored alongside S3 key in PostgreSQL — tamper-evident, cryptographically verifiable
- Required for GDPR, ISO 27001, SOC 2 compliance claims
- Generated automatically — zero manual effort from IT admin

**6. Basic Symptom Navigation**
- 6 symptom categories: "Suspicious account activity / Employee leaving / Strange emails / Unauthorized app found / Data may have been accessed / Access control issue"
- Maps to correct playbook (simple lookup, not ML routing)
- Disambiguation screen when symptom matches >1 playbook — users cannot be silently mis-routed to a wrong playbook
- Every routing decision logged with the rule that fired it

### Tier 2 — High Value, Sprint 9-11

**7. Inline Action Log (evidence timeline in UI)**
- Persistent read-only feed in the playbook UI: every action taken with timestamp, result status, provider
- Visible during AND after incident — not just in post-incident report
- Replaces the "PDF fallback" concept — more useful, simpler to build
- Doubles as compliance evidence display

**8. CEO/Executive Summary Export (Sprint 9)**
- One-page plain-language summary: what happened, what we did, risk to the business, prevention recommendations
- Generated from frozen evidence snapshot (not live data) — anchored to SHA-256 hash of source evidence
- Export: PDF only, timestamped
- Separate from compliance PDF which remains the legal record
- **Condition:** Requires Scenario A data model — each playbook step tagged with semantic outcome (SUCCESS / SKIPPED / ESCALATED / FAILED / PARTIAL)

**9. Phishing Response Playbook (Sprint 11, conditional)**
- Shares ~80% code with Account Compromise
- Steps: quarantine suspicious email, reset credentials, check for forwarding rules, alert other employees
- Condition for v1 inclusion: Backend capacity confirmed ≤75% effective in Sprint 11

**10. Static Playbook Walkthrough (onboarding)**
- Read-only preview of each playbook before first real incident
- Annotated screenshots showing each step and what to expect
- Adds to onboarding flow — not a separate feature
- Build cost: ~1 person-day; replaces the unfeasible "practice mode" for v1

### Tier 3 — v2 Commitments (not optional, just deferred)

**11. DRY_RUN Practice Mode (v2)**
- `mode: "DRY_RUN"` flag on playbook execution API
- Executes full SFN state machine without real API calls — logs what would fire
- Returns `DRY_RUN_SIMULATED` status on all steps
- No sandbox accounts required
- Requires: each provider Lambda to check execution context mode before API call
- Build cost: ~5-6 Backend days + 1 Frontend day — not feasible in v1 sprint capacity

**12. Ransomware Credential Containment Guidance (v2)**
- Steps that ARE automatable: suspend accounts, revoke sessions across providers, snapshot evidence
- Steps that are NOT automatable: network segmentation, device isolation (requires MDM/EDR)
- Named correctly: "Credential Containment (Ransomware Response)" — not "Ransomware Containment"
- v1 delivers: escalation guidance page with "call your IT provider NOW" and cyber insurance hotline

**13. BEC Wire Fraud Verification (v2)**
- "We received a suspicious wire transfer request" workflow
- Requires: out-of-band verification flow, banking integration or manual steps
- High SME financial value ($2.9B annual loss sector)

**14. Symptom Disambiguation (Sprint 10)**
- Smart disambiguation when symptom maps to multiple playbooks
- Rule-based JSON config (not ML) — versioned in git, human-readable
- Deferred from launch to gather real usage data first

---

## 5. Architecture Decisions (Finalized)

### 5.1 Execution Engine

**AWS Step Functions Standard Workflows — fixed constraint, not a decision**

| Factor | Decision |
|--------|----------|
| Workflow type | Standard (not Express) — mandatory for pause/resume up to 24h |
| Cost | $0.025/state transition — negligible at SME scale (~$2/month for 100 tenants) |
| Pause/resume | `.waitForTaskToken` pattern with 1-hour heartbeat |
| Fault tolerance | Built-in retry with exponential backoff per step |
| Audit trail | All state transitions logged via EventBridge → PostgreSQL |

SFN is used for: external API orchestration only. It does NOT model wizard navigation state.

### 5.2 Wizard State Layer

**PostgreSQL FSM (via execution log table) — not a custom state machine**

```
playbook_step_executions table:
  id, execution_id (FK), step_id, provider_id,
  status (PENDING | IN_PROGRESS | COMPLETED | FAILED | DRY_RUN_SIMULATED),
  started_at, completed_at, error_code, api_call_summary

UNIQUE INDEX on (execution_id, step_id, provider_id)  -- idempotency key
```

EventBridge routes SFN completion events → Lambda → PostgreSQL. No custom FSM code. No Redis.

### 5.3 Idempotency

PostgreSQL unique constraint on `(execution_id, tenant_id)`. Duplicate execution request → `23505` constraint violation → HTTP 409. Zero additional infrastructure.

### 5.4 Evidence Storage

- S3 Object Lock COMPLIANCE mode, 3-year default (7-year Enterprise)
- SHA-256 computed before S3 write, stored in `evidence_records.sha256_hash`
- CloudTrail enabled on evidence bucket
- Verification endpoint: `GET /api/evidence/{id}/verify`

### 5.5 What Was Rejected (and Why)

| Rejected | Alternative | Reason |
|----------|-------------|--------|
| Step Functions Express Workflows | Standard Workflows | Express max 5 min — incompatible with 24h human pause |
| Custom PostgreSQL FSM | EventBridge → PostgreSQL log | Custom FSM = 4.5-6.5 unplanned Backend days for no additional value |
| Redis for idempotency | PostgreSQL unique constraint | Adds Redis cluster ops burden; PG constraint is sufficient at SME scale |
| OPA/Rego for playbook authorization | Simple `allowed_roles JSONB` column | OPA adds 100ms+ latency per step for a permission model that's a 3-row lookup table |
| WebSocket for step status | HTTP polling every 3s | SSE/WebSocket complexity unnecessary when steps take 5-30s; polling is identical UX |
| SHA-256 hash chains (TA v1 proposal) | SHA-256 + S3 Object Lock | Hash chain per evidence object is sufficient; full hash chain infrastructure is over-engineering |
| Undo support | Pre-flight confirmation | Distributed rollback across 4 providers = 2+ sprint complexity; creates false security confidence |
| Inactive Account as incident playbook | Scheduled governance automation | This is maintenance, not incident response — wrong product category |

---

## 6. Security Requirements (Non-Negotiable)

| # | Requirement | Specification |
|---|-------------|---------------|
| SR-1 | Tenant isolation | RLS on all `playbook_*` tables; `workspace_id` injected from Keycloak JWT only; CI test verifies cross-tenant query returns zero rows |
| SR-2 | Audit log immutability | S3 Object Lock COMPLIANCE mode; write-via-Lambda only (no direct put from app); CloudTrail on bucket |
| SR-3 | OAuth tokens encrypted at rest | AWS Secrets Manager + KMS CMK; zero plaintext in DB; 90-day auto-rotation |
| SR-4 | Playbook execution authorization | Keycloak JWT with `playbook:execute` scope; checked at API gateway AND SFN execution start |
| SR-5 | Privileged action rate limiting | `MaxConcurrency` on parallel SFN states; max 50 provider API calls per minute per execution |
| SR-6 | Pre-flight entity display | Name, email, photo, role, last login shown before any destructive action — no generic "are you sure?" |
| SR-7 | workspace_id enforcement | Injected from JWT claim only; never from URL parameter or request body |
| SR-8 | Symptom routing audit | Every routing decision logged with the rule/branch that triggered it |
| SR-9 | MFA gate | Tenant-configurable, **default OFF** — hard gate would block response during crisis for 40-60% of SMEs without full MFA rollout |
| SR-10 | M365 CAE caveat | UI must display after every M365 revocation: modern auth <5 min; legacy auth up to 60 min — this is NOT optional |

---

## 7. Scope Changes vs. Original Plan

### Added (missing from original)

| Feature | Source | Justification |
|---------|--------|---------------|
| Phishing Response playbook | Round 1 PO + TA research | #1-2 SME incident type; table stakes against all competitors |
| GDPR Data Breach Guidance playbook | Round 1 TA research | Required for EU compliance claims; ISO 27001 A.16 compliance blocker |
| Symptom-based entry (basic) | Round 1 PO gap analysis | Users don't know playbook names during active incident |
| Concrete escalation UX with onboarding contacts | Round 1 TA + Round 2 PO | "Escalate" button with no contact is a dead end for 80% of target users |
| CEO/Executive Summary report | Round 1 PO gap analysis | Compliance PDF is for auditors; owners need plain language |
| Per-provider step status + partial failure retry | Round 1 TA threat model | "Partially completed" state with no recovery path = product failure |
| Pre-flight confirmation UX | Round 2 consensus | Replaces undo with safer, simpler pattern |
| SHA-256 hash before S3 write | Round 1 TA, accepted Round 2 | Cryptographic tamper-evidence for litigation/SOC 2 scenarios |

### Removed (over-engineered or wrong scope)

| Feature | Status | Alternative |
|---------|--------|-------------|
| Undo support | **Removed** | Pre-flight confirmation UX |
| Inactive Account playbook | **Removed from IR** | Moved to scheduled governance automation |
| Ransomware Containment (full) | **Descoped** | Credential Containment guidance page + escalation (v2 for full) |
| 85% coverage percentage claim | **Removed** | Replaced with named incident list — specificity > statistics |
| Custom PostgreSQL FSM | **Not built** | EventBridge → PostgreSQL execution log is sufficient |
| Redis for idempotency | **Not built** | PostgreSQL unique constraint is sufficient |
| OPA for playbook RBAC | **Removed from playbook path** | OPA stays for asset/DLP decisions; playbooks use simple role column |
| "GDPR notification auto-determination" | **Descoped** | Fact-gathering + DPO escalation only — legal determination is not automatable |

---

## 8. Delivery Plan (Validated by PM Analysis)

### Pre-Sprint 8 Hard Gates (all must pass before Sprint 8 starts)

| Gate | Owner | Evidence Required |
|------|-------|-------------------|
| SFN Account Compromise state machine prototype built and executed in staging (normal + pause/resume + failure paths) | Tech Lead | Screenshot in SFN console + state machine definition in repo |
| Test tenant access: real API calls to Google Admin SDK or Microsoft Graph from staging | PM | API call log showing real user data from test tenant |
| Pre-flight confirmation UX design reviewed and signed off by PO + Frontend | PO + Frontend | Figma artifact with PO sign-off |
| Sprint 7 JIT: zero Sev1 open defects | DevSecOps/QA | GitHub Issues: no open `Sev1` + `JIT-access` tickets |
| PostgreSQL schema peer-reviewed and locked: `execution_records`, `execution_steps`, `evidence_records` | Backend + Tech Lead | Merged PR with migration files |

### Sprint 8 — Scope (2 playbooks, engine, wizard)

| Deliverable | Notes |
|-------------|-------|
| SFN Standard Workflows engine | Pause/resume, fault-tolerant, EventBridge → PostgreSQL audit |
| Wizard UI (steps, decision gates, progress, pre-flight confirmation) | Zero jargon validation required as AC |
| Account Compromise playbook | Full automation: revoke + suspend + notify + evidence |
| Offboarding Emergency playbook | Wraps Sprint 6 SFN workflow as playbook-triggered variant |
| Email + Slack notifications (P0/P1) | SES + Slack webhook |
| SHA-256 + S3 Object Lock evidence storage | 0.5d Backend |
| FCM Android push setup | Flutter infra only, no wizard |
| Basic symptom navigation (6-category dropdown) | Simple lookup, 2d total |

**Effective utilization: ~94% (down from 113%)**

### Sprint 9 — 3 Playbooks + Mobile

| Deliverable | Notes |
|-------------|-------|
| Shadow IT Remediation playbook | Reuses OAuth revoke engine from Sprint 4 |
| Unauthorized Access playbook | Suspend resource access + notify manager |
| Inactive Account playbook | Warn + grace period + suspend (governance automation variant) |
| Mobile REST APIs for all 5 playbooks | Flutter client endpoints |
| Flutter iOS + Android: asset inventory + JIT approve/deny + FCM push | Wizard deferred to Sprint 10 |
| CEO Summary report (if Scenario A data model confirmed in S8 design) | Template variant of compliance PDF |

**Effective utilization: ~91%**

### Sprint 11 — Phishing Playbook (conditional)

Condition: Backend capacity ≤75% effective + TA code diff confirms ≥80% reuse with Account Compromise. If either condition fails: Phishing is v1.1.

---

## 9. Success Metrics at v1 Launch

Measured within 30 days of pilot deployment. Source: system telemetry only (no self-reported data).

| Metric | Target | Source |
|--------|--------|--------|
| Time-to-contain for Account Compromise | ≤15 min for ≥80% of executions | PostgreSQL: `first_step.started_at` → `final_revocation.completed_at` |
| Completion rate (non-security user, no escalation) | ≥85% of runs complete without IT admin override | Execution log: `escalation_required = false` |
| Production safety incidents caused by playbook | 0 | Support tickets tagged `playbook-caused` |
| Evidence audit completeness | 100% of completed executions have S3 + SHA-256 + PDF | Nightly automated verification job |
| Pilot NPS for playbook feature | ≥40 | In-app survey after first 3 completed runs |

---

## 10. The Product Promise — Final Statement

### What non-security staff can do in 10 minutes at launch:

> Within 10 minutes of opening SMESec, an IT admin with no security training can: identify the correct incident type, launch the relevant playbook, execute all automated remediation actions across connected providers, and have an immutable audit log recording every action taken — for an Account Compromise or Phishing Response incident.

### What requires more time (honestly documented):

- Offboarding Emergency: ~8-12 minutes (more providers, more steps)
- GDPR Data Breach Guidance: ~45-90 minutes across multiple sessions (regulatory fact-gathering cannot be rushed)
- Any playbook that requires escalation to legal/insurance: time depends on response from escalation contact

### What v1 does NOT do (stated prominently):

- Network isolation or device quarantine (requires MDM/EDR — not in scope)
- Automated determination of GDPR reportability (legal judgment — not automatable)
- Ransomware full containment (operational security scope beyond identity controls)
- BEC wire fraud verification (requires banking integration — v2)

### The 3 genuine differentiators vs. all competitors:

1. **Execution, not documentation** — Vanta/Drata give you a PDF. SMESec suspends the compromised account in 90 seconds.
2. **GDPR evidence generated automatically during response** — every time you respond to an incident, compliance evidence is created as a byproduct with zero extra effort.
3. **Non-security staff closes incidents without calling a security consultant** — the gap between "this product tells you about a problem" and "this product lets you solve the problem" is where SMESec sits alone in the SME market.

---

## Appendix A — Open Validation Questions (Customer Research Required)

| Question | Why It Matters | Target Sample |
|----------|---------------|---------------|
| Who does a non-security IT admin actually call when they hit an escalation gate? Is there a contracted MSSP, or are they the most technical person available? | Determines whether escalation feature is usable for Starter tier | 10+ pilot candidates, Starter profile |
| When an incident occurs, does the IT admin open the app and run a playbook, or do they call their manager first? | Changes entire UX flow if playbooks are post-incident documentation tools, not real-time response tools | 5+ user interviews before Sprint 8 |
| Do 50-200 employee companies require 7-year evidence retention, or is 3 years sufficient for their actual auditors? | Affects pricing tier design and S3 cost model | 3 conversations with ISO 27001 lead auditors |

---

## Appendix B — Competitor Feature Matrix (Final)

| Feature | SMESec v1 | Vanta | Drata | Secureframe | Huntress | PagerDuty |
|---------|-----------|-------|-------|-------------|---------|-----------|
| Pre-built playbooks | ✅ 6 playbooks | ❌ PDF only | ❌ PDF only | ❌ PDF only | ❌ Managed service | ❌ Technical runbooks |
| Executable by non-security staff | ✅ Core design | ❌ | ❌ | ❌ | ❌ | ❌ Engineer required |
| Automated remediation (suspend, revoke) | ✅ <2 min | ❌ | ❌ | ❌ | ⚠️ Huntress team does it | ✅ Engineer-executed |
| Compliance evidence auto-generated | ✅ SHA-256 + S3 Lock | ✅ Primary feature | ✅ Primary feature | ✅ | ❌ | ❌ |
| Mobile execution | ✅ (Sprint 9) | ❌ | ❌ | ❌ | ⚠️ Alerts only | ⚠️ Alerts only |
| GDPR notification guidance | ✅ Fact-gathering + DPO | ⚠️ Documentation only | ⚠️ Documentation only | ❌ | ❌ | ❌ |
| Concrete escalation path | ✅ Per-gate contacts | ❌ | ❌ | ❌ | ❌ | ❌ |
| Phishing response | ✅ (Sprint 11) | ❌ | ❌ | ❌ | ⚠️ Managed alert | ❌ |
| AI/Shadow IT remediation | ✅ OAuth revoke | ❌ | ❌ | ❌ | ❌ | ❌ |
| SME pricing (<$2K/month) | ✅ | ✅ | ✅ | ✅ | ✅ | ❌ |
