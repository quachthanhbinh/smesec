# Asset Inventory — Research Synthesis

**Date:** 2026-05-28  
**Method:** 3 agents × 2 rounds (6 agent invocations total)  
**Agents:** Product Owner · Technical Advisor · Project Manager  
**Scope:** IT Asset Inventory key requirement — fresh research, no prior plan assumed  
**Next:** [02-decision-record.md](02-decision-record.md)

---

## Executive Summary (Converged Across All 3 Agents)

1. **The "aha moment" is universal**: Every SMB that runs its first discovery says the same thing — *"We had no idea we were using this many applications."* 100-person SMB runs 150–300 SaaS apps; IT knows about fewer than 50. This 3x–6x gap is the activation hook.
2. **Depth beats breadth every time**: Teams that launched with 8 integrations at 60% coverage consistently shipped products admins couldn't trust. Winners (Zluri 12 months, BetterCloud 18 months) started with 1 IdP at >95% coverage and expanded methodically.
3. **The product is the onboarding experience, not the feature list**: 80% 90-day retention if first asset seen in <30 minutes; 30% if >2 hours. This is not a UX metric — it is the activation mechanism.
4. **Shadow AI is the unoccupied 2026 market gap**: Zero SMB-priced tools inventory AI tool usage. OAuth-based discovery (no browser extension needed for v1) delivers immediate value at the highest-growth risk category.
5. **The primary competitor is JupiterOne** at SMB tier ($500-3K/month) — but its JQL query language creates a UX barrier SMBs can't cross. SMESec wins by being the pre-answered version, not the open-ended query version.
6. **Timeline reality**: 14-16 weeks to pilot, 5 months to first paying customer. The old 10-week estimate under-counts M365 complexity (2 sprints minimum, confirmed by all 3 agents + market pattern data).

---

## Part 1: Competitive Landscape

### Competitor Feature Matrix

| Product | Category | SaaS Discovery | Identity Inventory | Shadow IT | Shadow AI | Cloud Inventory | Offboarding | Compliance | Mobile App | SMB Price |
|---|---|---|---|---|---|---|---|---|---|---|
| **Torii** | SaaS Mgmt | ✅ 300+ | ⚠️ User-app | ✅ Strong | ⚠️ Partial | ❌ | ❌ | ⚠️ Basic | ❌ | $5-20K/yr |
| **Zluri** | SaaS Mgmt | ✅ 800+ | ⚠️ User-app | ✅ Strong | ⚠️ Partial | ❌ | ❌ | ⚠️ Basic | ❌ | $5-15K/yr |
| **BetterCloud** | SaaSOps | ✅ G/M365 | ⚠️ Limited | ✅ | ❌ | ❌ | ✅ G/M365 | ❌ | ❌ | $10K+/yr |
| **JupiterOne** | CAASM | ✅ | ✅ Graph | ✅ | ⚠️ | ✅ Multi-cloud | ⚠️ | ✅ | ❌ | $6-36K/yr |
| **Axonius** | CAASM | ✅ | ✅ | ✅ | ❌ | ✅ | ⚠️ | ✅ | ❌ | $50K+/yr |
| **AppOmni** | SSPM | ✅ Config | ⚠️ | ⚠️ | ❌ | ❌ | ❌ | ✅ | ❌ | $30K+/yr |
| **DoControl** | SSPM | ✅ Data | ⚠️ | ⚠️ | ❌ | ❌ | ✅ | ⚠️ | ❌ | $40K+/yr |
| **Adaptive Shield** | SSPM | ✅ 150+ | ✅ | ✅ | ❌ | ❌ | ⚠️ | ✅ | ❌ | $30K+/yr |
| **Lansweeper** | ITAM | ❌ | ⚠️ AD only | ❌ | ❌ | ⚠️ Basic | ❌ | ⚠️ | ❌ | $2-10K/yr |
| **SailPoint** | IGA | ❌ | ✅ Deep | ❌ | ❌ | ❌ | ✅ Complex | ✅ | ❌ | $100K+/yr |
| **SMESec Target** | Unified | ✅ | ✅ | ✅ + enforce | **✅ First SMB** | ✅ Posture | ✅ <5min | ✅ | **✅ First** | $3-8/usr/mo |

### Primary Competitor: JupiterOne (Not Vanta/Drata)

JupiterOne ($500-3K/mo SMB tier) is the closest full-stack competitor at accessible price. **Why SMESec wins:**

| Dimension | JupiterOne | SMESec |
|---|---|---|
| Query model | JQL (graph query language — requires security literacy) | Pre-answered findings + playbooks |
| Time to first value | 2–4 weeks onboarding | <30 minutes |
| Mobile app | ❌ | ✅ |
| Shadow AI discovery | ⚠️ Partial | ✅ (OAuth) + browser ext (Track 2) |
| Offboarding automation | ⚠️ Query-driven, not workflow | ✅ <5 min, human-initiated |
| Compliance evidence | ✅ Deep | ✅ |

**The SMESec positioning weapon**: JupiterOne shows everything. SMESec tells you what to do about it.

### White Space (Market Gaps)

| Gap | Severity | Who Has It | Price |
|---|---|---|---|
| Shadow AI tool discovery at SMB price | **Critical 2026** | No one | N/A |
| Unified cross-class inventory + offboarding | High | JupiterOne (but complex UX) | $6K+ complex |
| Non-expert UX for full-stack inventory | High | None | N/A |
| Mobile app for incident response | Medium | Zero competitors | N/A |
| Asset inventory + compliance evidence bundle | High | None at <$10K | N/A |

---

## Part 2: SMB Pain Points (Ranked)

| Rank | Pain | Evidence |
|---|---|---|
| #1 | **Shadow IT / unknown SaaS** | 254 apps avg 200-person SMB; IT approved 41. 68% breaches involved undetected third-party app (Verizon DBIR 2025) |
| #2 | **Orphaned access after offboarding** | 74% breaches involve privileged credential misuse; ex-employee accounts are #1 source. Avg offboarding takes 3–5 days, 1-in-5 SMBs had incident from ex-employee (BetterCloud) |
| #3 | **Shadow AI (2026 emerging critical)** | 78% knowledge workers use AI tools at work; 52% use tools employer didn't provide. 11% of data pasted to ChatGPT is confidential company data (Cyberhaven 2025) |
| #4 | **No data classification / "where is our sensitive data?"** | 34% GDPR fines involve failure to maintain processing records. Zero SMB-priced data classification tools exist below $15K/yr |
| #5 | **License waste / SaaS spend blindness** | Avg SMB wastes $127K/yr (200-person). 31% of SaaS licenses are unused >90 days (Torii 2025) |
| #6 | **Cloud asset blindness (tech companies)** | 58% cloud incidents traced to misconfigured IAM; 70% SMBs have at least one public cloud storage bucket (Wiz 2025) |

### The "Unlock Event" (PM Research)
SMBs switch from DIY to tooling after ONE of:
- Former employee found with active app access 3 months after leaving
- SOC 2/ISO 27001 audit reveals no asset inventory
- Incident: file-sharing to unknown OAuth app

**Product positioning must target this emotional context, not "comprehensive asset management."**

---

## Part 3: Tension Resolutions (Round 2 Cross-Iteration)

### Tension A: Shadow AI Discovery — Browser Extension vs. OAuth

**Resolution (Consensus across all 3 agents):**

Shadow AI discovery is **two separate features** that were conflated:

| Feature | Mechanism | Timeline | Value |
|---|---|---|---|
| **Shadow AI tool discovery** | OAuth scope analysis — which AI apps have granted access + what scopes | v1 Track 1 (Sprint 4, data already collected) | "ChatGPT has read access to your Google Drive — it's not on your approved list" |
| **Shadow AI content monitoring** | Browser extension (Manifest V3) — domain visits + form POST detection | Track 2 v1 conditional (MV3 hard gate Week 1) | "3 employees pasted data into Claude.ai last week" |

**v1 without browser extension coverage:**
- OAuth-discoverable AI tools: ~65-75% of shadow AI inventory
- Browser extension adds: ~90-95% total coverage
- Critical insight: AI tools that *caused data leakage* (OAuth-integrated) ARE detectable without extension

**Browser extension gates (must pass before shipping):**
1. MV3 service worker persistence test: prototype passes Week 1 Sprint 1
2. GDPR Article 13/88 legal opinion received
3. Privacy dashboard implemented (what data collected, retention period)

**Product disclosure requirement**: When extension not installed, dashboard MUST show "AI tool discovery: partial coverage — browser activity monitoring not active."

---

### Tension B: Cloud Asset Inventory — v1 Scope

**Resolution: "AWS Security Posture (5 checks)" — not "Cloud Asset Inventory"**

PM's coverage-gap concern is real. TA's API quality argument is also real. The resolution: **rename and scope the feature to eliminate the coverage gap problem.**

| Scope | Problem | Solution |
|---|---|---|
| "Cloud Asset Inventory" (full) | SMBs expect Azure + GCP too; 60% coverage → trust erosion | ❌ Not v1 |
| Nothing | Misses IAM orphaned users + public S3 — both critical security findings | ❌ Too little |
| **"AWS Security Posture (5 checks)"** | Bounded, actionable, complete coverage within its scope | ✅ v1 |

**5 Checks (v1):**

| Check | API | Why Critical |
|---|---|---|
| IAM users inactive >90 days | `iam:ListUsers` + `GetLoginProfile` | Identity governance — same category as Google users |
| S3 buckets with public read/write | `s3:GetBucketAcl` + `GetBucketPolicyStatus` | Nightmare scenario every SMB fears |
| Root account MFA not enabled | `iam:GetAccountSummary` | Critical finding, zero ambiguity |
| Access keys not rotated >90 days | `iam:ListAccessKeys` | Common CISO finding, easy win |
| Security groups with 0.0.0.0/0 on 22/3389 | `ec2:DescribeSecurityGroups` | SSH/RDP exposed to internet — exploitable today |

**v2:** Full AWS Config inventory + Azure + GCP + EC2/RDS/Lambda resource graph  
**Never:** Use the phrase "Cloud Asset Inventory" in v1 customer-facing materials (implies multi-cloud completeness)

---

### Tension C: Asset Risk Score vs. Deterministic Findings

**Resolution: Deterministic findings FIRST, risk score as derived display. No false positive problem.**

Context difference from Access Governance: Access Governance scoring requires *judgment* on permission graph complexity → false positives. Asset Inventory scoring = weighted sum of rule-based findings → deterministic, auditable, transparent.

**Architecture (TA-designed, all agents converged):**

```
User Risk Score (0–100):
  admin_without_mfa          → +25
  orphaned_account           → +30
  privileged_no_login_60d    → +20
  inactive_account_90d       → +10
  pending_offboarding        → +25
  unreviewed_shadow_apps     → +5 each (capped +20)

App Risk Score (0–100):
  approval_status_blocked    → +35
  approval_status_unknown    → +25
  high_risk_oauth_scopes     → +20 (mail.read, admin, files.readwrite)
  no_vendor_security_cert    → +10
  ai_tool_unreviewed         → +15
  data_access_restricted     → +20
  review_overdue_180d        → +10
```

**False positive mitigations:**
- Service accounts: `account_type = service_account` suppresses MFA/login rules
- Leave of absence: `leave_status` column auto-suppresses inactive rules
- Day 1 "learning mode": scores calculated but not surfaced as alerts until >50% catalog reviewed
- Score is informational ONLY in v1 — NEVER triggers automated enforcement

---

## Part 4: Feature Set (Final)

### Must-Have v1 (Launch Gate)

| Feature | Why | Data Source | Sprint |
|---|---|---|---|
| **Onboarding wizard** (<30 min first asset) | 80% vs 30% 90-day retention | - | 7 |
| **SaaS discovery — Google Workspace** (users + OAuth apps + last login) | Foundation + aha moment | Admin SDK | 2 |
| **SaaS discovery — M365** (users + OAuth app grants + activity) | 50% SMB market | Microsoft Graph | 3-4 (2 sprints) |
| **Shadow IT allow-list + alerts** | Closes discover→action loop | OAuth tokens | 4 |
| **Shadow AI discovery via OAuth** | 2026 differentiator; data already collected in Sprint 4 | OAuth scopes + vendor catalog | 4 (zero extra work) |
| **Pre-seeded vendor catalog (500+ apps)** | Hard gate — <10% Unknown; without it every app scores high → alert fatigue | Platform | 4-5 |
| **Identity inventory** (users, groups, roles, admin status) | Foundation for all downstream features | IdP APIs | 2-3 |
| **Stale account detection** (>90 days inactive) | Ex-employee access — unlock event #1 | IdP last-login | 5 |
| **"Forgotten access" dashboard** (ex-employees still with active accounts) | Unlock event #1 UX; high buyer emotional resonance | Query on offboarding + asset data | 5 |
| **6 deterministic findings** (MFA missing, orphaned, public S3, etc.) | SOC 2 evidence; 0% false positives | Rule engine | 5 |
| **Asset risk score** (0-100 per user + per app, deterministic rollup) | Makes findings actionable for non-security staff | Derived from findings | 5 |
| **Zombie Account Cost Recovery** (inactive users × license price) | Immediate ROI visible; no billing API needed | Last-login + manual price entry | 5 |
| **Automated offboarding** (Google + M365 + flagged OAuth apps, <5 min) | Anchor feature; product's defining capability | IdP + OAuth APIs | 6 |
| **Compliance evidence export** (ISO 27001 A.8/A.9 + SOC 2 CC6.1) | Audit unlock event; $5-20K saved per SMB | Asset snapshot + PDF | 6 |
| **AWS Security Posture** (5 checks — IAM, S3, root MFA, key rotation, SG) | Compliance completeness; high-value findings | AWS SDK | 7 |

### Should-Have v1 (Include If Capacity)

| Feature | Notes | Sprint |
|---|---|---|
| Ramp + Brex expense integration | Discovers paid-but-no-OAuth apps; powers license waste signal. 1-2 days each. | 5 |
| Full license waste — M365/Google seat calculation | Requires M365 Usage Reports + Google Reports APIs; possible in same sprint as stale account | 11-12 |
| Dependency graph / blast radius (basic) | Already in S12 current plan; user → app → cloud resource (3 hops) | 12 |
| Mobile app (Flutter — incident response + offboarding approval) | Zero competitor; already funded in sprint plan | 4-7 |

### v2 (Explicitly Out of v1)

| Feature | Reason | Effort Estimate |
|---|---|---|
| Full cloud inventory (AWS Config + Azure + GCP) | Multi-cloud coverage gap; trust erosion in v1 | 6-8 weeks |
| Data classification — SaaS PII scanning (Google Drive / SharePoint) | 83hr/tenant initial scan; $150/tenant Google DLP cost; GDPR legal review; 5-15% false positive rate | 4-6 months |
| Full license waste dashboard (billing APIs — Zoom, Salesforce, HubSpot) | Per-vendor API builds; 2-3 sprints each | 4-6 sprints |
| Expensify / Concur expense integration | Older APIs; lower SMB relevance | 3-4 days |
| Browser extension for content monitoring | Conditional on MV3 gate + GDPR review | Track 2 |

### Removed / Never

| Feature | Reason |
|---|---|
| Agent-based endpoint discovery | MDM required; deployment friction too high for SMBs |
| Auto-revocation of shadow IT | Blast radius risk = instant churn (same decision as Access Governance) |
| JQL / graph query interface | JupiterOne trap — builds for security team, not IT admin |
| ML-based risk scoring | No labeled training data; false positive rate kills trust |
| Network-layer CASB | Proxy infrastructure; $15-25/user/mo; enterprise only |
| Custom playbook builder (v1) | 2 pre-built playbooks sufficient; engine in v2 |
| Raw credit card number ingestion | PCI DSS risk; not worth it |

---

## Part 5: Technical Architecture Decisions

### Discovery Architecture

**Agentless-first + API polling + selective event streams.** No endpoint agents required for v1.

```
Discovery Sources (v1):
  Google Workspace Admin SDK    → users, groups, org units, OAuth tokens
  Microsoft Graph               → users, groups, service principals, app grants
  AWS IAM + S3 + EC2            → 5 security posture checks (not full Config)
  GitHub App                    → org members, repos, teams
  Slack Admin API               → users (Business+ only for full coverage)
  
Optional v1:
  Ramp / Brex APIs              → expense merchant data for paid-app discovery
  
Track 2 conditional:
  Chrome Extension (MV3)        → domain visits + form POST detection (shadow AI content)
```

### Sync Architecture: Dual-Speed Pipeline

```
FAST PATH (<5 minutes — critical events):
  AWS EventBridge → CloudTrail events (S3 public, IAM new admin)
  M365 Graph Change Notifications
  Slack Events API

STANDARD PATH (15–60 minutes — routine sync):
  Google Workspace 15-min poll (Reports API has 90-min indexing delay → polling only)
  M365 Graph delta queries (delta token must be consumed within 7 days)
  GitHub 30-min poll
  AWS reconciliation 60-min full sync

Both converge: SQS → ECS workers → PostgreSQL assets table
```

### First-Sync Fast Path (<30 Minutes to First Value)

```
T+0:00  Google OAuth consent completed
T+0:30  ECS task spawned — Phase 1 begins
T+3:00  Phase 1: Users visible in dashboard (bulk insert via PostgreSQL COPY)
T+7:00  Phase 2: OAuth apps visible — "23 apps, 4 need review" — AHA MOMENT
T+25:00 Phase 3 (background): Full token history + risk scores + findings calculated
```

**Key design decisions:**
- Reports API (fast path) shows 30-day OAuth grants in 1-3 API calls; full `tokens.list` takes 500 calls for 500 users
- Progressive SSE rendering: user sees partial data live, full data arrives while browsing
- PostgreSQL COPY protocol for bulk insert (20× faster than ORM row-by-row)
- Risk scores decouple from first render (calculated at T+20, not blocking T+7 moment)

### Data Model

```sql
-- Core universal entity
assets (
  id, workspace_id, asset_type, source, external_id,
  display_name, classification, status, lifecycle_state,
  owner_id, metadata JSONB,
  first_seen_at, last_seen_at
)

-- Edges: user → app, app → resource, group → user
asset_relationships (
  source_asset_id, target_asset_id, relationship_type,
  granted_scopes TEXT[], granted_at, revoked_at, metadata JSONB
)

-- Audit trail: append-only state changes
asset_snapshots (asset_id, snapshot_at, change_type, delta JSONB)

-- Lifecycle state machine
asset_lifecycle_history (asset_id, from_state, to_state, changed_by, reason)

-- Usage events (daily granularity)
app_usage_events (app_asset_id, user_asset_id, event_date, event_count, source)

-- License records
app_license_records (asset_id, plan_name, license_count, cost_monthly_cents, renewal_date)

-- Deterministic risk scores (derived from findings)
asset_risk_scores (asset_id, score 0-100, severity, factors JSONB, overridden, override_expires_at)

-- Expense-based discovery (optional)
expense_merchant_events (merchant_raw, merchant_normalized, matched_asset_id, spend_cents)

-- Browser extension events (domain only — GDPR minimization)
browser_domain_events (user_asset_id, domain, event_date, data_submitted)
```

**All tables: `workspace_id` + PostgreSQL RLS + CI cross-tenant isolation tests**

### Classification Engine (Rule-Based, 6 Stages)

```
Stage 1: Asset type rules (deterministic by source)
Stage 2: Vendor catalog match (500+ pre-classified vendors, trigram index)
Stage 3: OAuth scope classification (drive/mail = CONFIDENTIAL; profile only = INTERNAL)
Stage 4: AWS resource rules (S3 public = RESTRICTED; RDS = CONFIDENTIAL)
Stage 5: Shadow detection (not in catalog → SHADOW; untagged AWS = SHADOW)
Stage 6: Orphan detection (user departed AND still has active OAuth grants)

Confidence: Stage 1-2 = 1.0, Stage 3-4 = 0.85, Stage 5-6 = 0.70
Below 0.60 → surface to IT admin for manual review
Acceptance criterion: <10% Unknown for top-200 SMB apps (hard gate before launch)
```

### Integration Complexity Matrix

| Source | What Discovered | Rate Limit | Auth Complexity | v1 Scope |
|---|---|---|---|---|
| **Google Workspace Admin SDK** | Users, groups, OAuth tokens | 1,500 req/100s (Reports: 250 req/100s) | Service account + DWD | Full identity + SaaS |
| **Microsoft Graph** | Users, groups, app grants, service principals | 10,000 req/10min; delta tokens expire 7 days | App registration + admin consent | Full identity + SaaS |
| **AWS IAM/S3/EC2** | IAM users, S3 ACLs, security groups | No meaningful limits | IAM role + cross-account trust (CloudFormation one-click) | 5 posture checks only |
| **GitHub** | Org members, repos, teams | 5,000 req/hr | GitHub App installation (2 min) | Member inventory |
| **Slack** | Users, installed apps, channels | 50 req/min; NO admin API on free tier | OAuth + admin scopes | Business+ only |

**Integration traps:**
- Google Reports API: 90-minute indexing delay — events appear up to 90 min after they occur
- M365 delta tokens: expire after 7 days if not consumed — silent failure if sync job interrupted
- Slack free/Pro: NO `admin.*` API access — flag clearly in onboarding
- AWS cross-account: provide one-click CloudFormation template; never ask SMBs to configure IAM manually
- GitHub: use GitHub App (fine-grained permissions), NOT OAuth App

---

## Part 6: Delivery Analysis

### Market Timeline Benchmarks

| Product | Approach | Time to First Revenue | Key Accelerator |
|---|---|---|---|
| Torii | Browser ext + OAuth | 14 months | Shadow IT "aha moment" |
| **Zluri** | OAuth + narrow (1 IdP deep) | **12 months (fastest)** | Narrow + deep |
| BetterCloud | Google-only API | 18 months | Depth before breadth |
| JupiterOne | Graph model | 14 months | Cloud-native focus |
| Axonius | Agentless connectors | 16 months | Read-only trust model |
| AppOmni | SSPM + event logs | 18 months | Compliance framing |

**Average: 13-16 months to first revenue. SMESec target: 5 months (faster due to 3 agents, mature 2026 APIs, lessons from above).**

### Feature Value vs Effort Quadrant

```
HIGH VALUE │ Q1: Do First             │ Q2: Do Later
           │─────────────────────────│──────────────────────────
           │ Google OAuth discovery   │ Shadow AI (browser ext)
           │ M365 sync               │ Data classification (PII)
           │ Shadow IT alerts        │ Full license waste (APIs)
           │ Offboarding automation  │ Full cloud inventory
           │ Shadow AI (OAuth)       │ Compliance reports deep
           │ Zombie account recovery │
           │ Risk score (findings)   │
           │ Onboarding wizard       │
           ├─────────────────────────┼──────────────────────────
           │ Q3: Nice to Have        │ Q4: Don't Build v1
           │                         │
LOW VALUE  │ Ramp/Brex integration   │ JQL/query interface
           │ AWS posture checks      │ Agent-based endpoint
           │ Dependency graph        │ ML risk scoring
           │ Mobile app (approvals)  │ Auto-revocation
           └─────────────────────────┴──────────────────────────
              LOW EFFORT                HIGH EFFORT
```

### Day 1 vs Day 90 Usage Patterns

| Feature | Day 1 Usage | Day 90 Usage | Notes |
|---|---|---|---|
| Shadow IT / new app alerts | 70% | **80%** | Stickiness driver — ongoing value |
| Offboarding automation | 15% | **90%** | Event-triggered; critical when needed |
| User inventory view | 95% | 40% | Novelty; becomes background |
| Risk score / findings | 60% | 35% | High initial, stabilizes |
| Compliance reports | 10% | 30% | Value grows near audit |
| Dependency graph | 20% | 5% | "Cool demo, never used again" |
| Cloud resource list | 30% | 10% | SMBs don't act on this data |

### Common Failure Patterns (All Products in Category)

| Pattern | Frequency | SMESec Mitigation |
|---|---|---|
| API rate limits in production | 80% | Delta sync from Sprint 1; exponential backoff; test with 500-user synthetic tenant |
| Data quality sacrificed for breadth | 70% | >95% coverage per provider before adding next one |
| Dashboard without action | 60% | Every finding has "Fix it" button; alerts before dashboards |
| Alert fatigue (30%+ Unknown apps) | 55% | Pre-seed 500-app catalog; <10% Unknown hard gate |
| Underestimating Google→M365 jump | 65% | 2 sprints allocated to M365 (not 1) |
| Onboarding too manual | 50% | Onboarding wizard is Sprint 7 named deliverable, not assumption |
| SMB admin revokes API access | 15% | Graceful degradation UX; "connection interrupted" not silent failure |

### Revised Timeline (4-Person Team)

```
Sprint 1 (W1-2):   Infrastructure + Auth + MV3 hard gate decision
Sprint 2 (W3-4):   Google Workspace sync (users + OAuth apps, >95% coverage)
Sprint 3 (W5-6):   M365 sync Phase 1 (users + groups + basic app grants)
Sprint 4 (W7-8):   M365 sync Phase 2 + classification + shadow IT alerts + Shadow AI tagging
Sprint 5 (W9-10):  6 deterministic findings + risk score + zombie account cost recovery
Sprint 6 (W11-12): Automated offboarding + compliance evidence export
Sprint 7 (W13-14): Onboarding wizard + AWS Security Posture (5 checks)
Sprint 8 (W15-16): Pilot support + security hardening + pen test remediation
──────────────────────────────────────────────────────────────────
Week 14: PILOT READY (3–5 design partners)
Week 20: FIRST PAYING CUSTOMER
```

---

## Part 7: Open Questions (Requiring Team Decisions)

| # | Question | Owner | Deadline | Blocking? |
|---|---|---|---|---|
| 1 | **MV3 browser extension gate**: Does MV3 service worker pass persistence test? (Week 1) | TA + Track 2 | Sprint 1 W1 | Determines v1 vs v2 |
| 2 | **GDPR legal review for browser extension**: Starts Week 1, not Week 7. Written legal memo on employee monitoring lawful basis for EU markets (40-60% of SMB target). | PM + Legal | Sprint 1 W1 | If fails, extension = never for EU |
| 3 | **v1 ACV pricing**: PO's $8-25K requires full license waste + Shadow AI via extension. v1 actually ships Google + M365 + offboarding + 6 findings. Realistic v1 ACV = $3-8K with upsell path. Confirm pilot pricing. | PO + PM | Sprint 3 | Pilot LOI framing |
| 4 | **Pilot customer sourcing**: 2 LOI from pilot customers before Sprint 3 ends. Google Workspace customers first (before M365 ready). 3 total by Sprint 7. | PM | Sprint 3 | Sprint 7 milestone gate |
| 5 | **Pen test vendor**: Must be contracted by Sprint 4 (6-week scheduling lead time). Without it, Sprint 8 security hardening cannot be validated. | PM | Sprint 4 | Sprint 8 milestone gate |
| 6 | **Ramp/Brex integration**: Is expense data (paid apps not visible via OAuth) worth 1-2 sprints? PO downgraded to Should-Have; PM/TA support it. Confirm in Sprint 3 backlog grooming. | All | Sprint 3 | No |
| 7 | **Slack Business+ qualification**: What % of target pilot customers have Slack Business+/Grid? Determines whether Slack automated offboarding ships in v1 or v2. | PM | Sprint 4 | No |

---

*Session: 2026-05-28 | 3 agents × 2 rounds | 6 agent invocations | Decision record: [02-decision-record.md](02-decision-record.md)*
