# Feature Requirements Specification: Asset Inventory & Classification

**Document version:** 1.0  
**Status:** APPROVED — Synthesis of 3-Agent Debate (Rounds 1 & 2)  
**Date:** 2026-05-28  
**Authors:** Product Owner · Project Manager · Technical Advisor (3-agent iterative debate)  
**Applies to:** SMESec — Key Requirement: Asset Inventory & Classification  
**Method:** Round 1 (independent research) → Round 2 (cross-review & challenge) → Round 3 (synthesis)

---

## 1. Executive Summary

The Asset Inventory & Classification module is the foundational visibility layer of SMESec — it discovers every user account, OAuth-authorized application, device, and cloud resource connected to an SME's identity providers, then classifies each asset by criticality and data sensitivity. Without this foundation, access governance, automated offboarding, JIT access, and compliance reporting are impossible. The primary buyer is an IT admin or founder at a 10–500 person company who currently has no systematic answer to "what do we have, who has access to it, and which of it involves AI tools?" — and needs that answer in minutes, not months, without a security team.

---

## 2. Market Research Findings

### 2.1 Competitor Feature Matrix

| Feature | Vanta | Drata | Secureframe | Nudge Security | Axonius | JupiterOne |
|---|---|---|---|---|---|---|
| SaaS/OAuth app discovery | ⚠️ Compliance-scoped | ⚠️ Compliance-scoped | ⚠️ Compliance-scoped | ✅ Best-in-class (440+ apps avg) | ✅ via connectors | ✅ Graph-based |
| Device inventory | ❌ Agent required | ❌ Agent required | ❌ | ❌ | ✅ 800+ connectors | ⚠️ Cloud only |
| Shadow IT detection | ❌ | ❌ | ❌ | ✅ Email nudges | ⚠️ Detects, no workflow | ⚠️ No workflow |
| Shadow AI detection | ❌ | ❌ | ❌ | ⚠️ OAuth only | ❌ | ❌ |
| Asset criticality classification | ❌ Manual only | ❌ Manual only | ❌ Manual only | ❌ | ⚠️ Rule-based | ✅ Policy-as-code |
| Data sensitivity classification | ❌ | ❌ | ❌ | ❌ | ❌ | ⚠️ Custom queries |
| User-to-asset relationship map | ⚠️ Basic | ⚠️ Basic | ⚠️ Basic | ⚠️ OAuth grants only | ✅ | ✅ Graph |
| Compliance evidence auto-mapping | ✅ Primary feature | ✅ Primary feature | ✅ Primary feature | ❌ | ❌ | ⚠️ Queries required |
| Shadow IT remediation (in-platform) | ❌ Links out | ❌ Links out | ❌ Links out | ⚠️ Email nudge only | ❌ | ❌ |
| SME-friendly pricing | ⚠️ ~$800/mo | ⚠️ ~$700/mo | ⚠️ ~$600/mo | ⚠️ ~$4/user/mo | ❌ $50K+/yr | ❌ $2K-8K/mo |

### 2.2 Market Gaps

**Gap 1 — Shadow AI Tool Detection (Critical, Unserved)**  
No product comprehensively discovers employee AI tool usage at SME price points. Nudge Security detects OAuth-connected AI apps but misses browser-based tools. Gartner estimates 78% of SME employees use unsanctioned AI tools weekly.

**Gap 2 — Unified Asset Intelligence**  
No SME product unifies devices + SaaS apps + cloud resources + data into a single risk-ranked inventory. Customers need Lansweeper + Nudge + Vanta — three bills, no correlation.

**Gap 3 — Automated Data Sensitivity Classification**  
Every competitor requires manual classification. No product auto-scans Google Drive, Confluence, or Notion for PII / financial / IP sensitivity at SME pricing.

**Gap 4 — In-Platform Shadow IT Remediation**  
All competitors link out to the provider console to revoke OAuth grants. No competitor revokes directly via API with confirmation, audit log, and user notification.

**Gap 5 — Affordable Entry (<$200/month)**  
No product serves 10–50 employees with basic asset visibility, shadow IT detection, and compliance readiness at under $200/month.

### 2.3 Customer Jobs-to-be-Done

| Job | Frequency | Current Solution | WTP |
|---|---|---|---|
| Know what touches business data | Triggered by near-miss or new IT hire | Manual spreadsheet, 2-3 weeks, always stale | $300-800/mo |
| Pass SOC 2 / ISO 27001 without consultant | Annual | $15-30K consultant engagement | Very high ($1,500/mo saves $15K/yr) |
| Stop ex-employees & zombie apps from accessing systems | Every offboarding (2-4/mo) | Manual checklist, misses shadow IT | $200-400/mo |
| Stop employees leaking data to AI tools | Continuous background concern | IT policy nobody reads. Zero enforcement | Medium-high |
| Answer "what's the blast radius if this account is compromised?" | Rare deliberate exercise | Not answerable today | Medium — decisive differentiator |

---

## 3. Scope Definition

### 3.1 In Scope for v1

| Feature | Early Access (W8) | Commercial Launch (W13–14) | Full v1 (W26) |
|---|:---:|:---:|:---:|
| User Account Discovery & Inventory (Google + M365) | ✅ | | |
| OAuth App Discovery & Shadow IT Detection | ✅ | | |
| Basic Device Inventory (agentless, admin-reported) | ✅ | | |
| Asset Classification Engine (rule-based) | ✅ | | |
| OAuth Scope Risk Scoring | ✅ | | |
| Shadow AI Tool Authorization Detection | ✅ | | |
| New OAuth App Alert Pipeline (<15 min) | ✅ | | |
| Shadow IT Remediation Workflow (approve/block/notify) | ✅ | | |
| Asset Inventory Dashboard — basic (Google + M365) | ✅ | | |
| Slack + AWS Asset Discovery | | ✅ | |
| Asset Inventory Dashboard — full (all 4 providers) | | ✅ | |
| Data Sensitivity Labeling (metadata-based, no content scan) | | ✅ | |
| Plain-English Remediation Guidance | | ✅ | |
| User-to-Asset Dependency Map (PostgreSQL graph) | | | ✅ |
| Compliance Evidence Export (ISO 27001, GDPR, SOC 2) | | | ✅ |

### 3.2 Explicitly Out of Scope for v1

| Item | Reason | Target |
|---|---|---|
| MDM API device inventory (Jamf, Intune, Kandji) | Each MDM = 1 full sprint. Admin-reported device data from Google/M365 provides minimum viable device visibility. | v2 Month 7+ |
| Content-level PII scanning (file content) | DLP-grade scanning is a product category (18+ months). Creates GDPR exposure if content accessed without consent. | v2 |
| Slack full app inventory (Admin API) | Requires Slack Business+ ($15/user/mo). Majority of SME target customers are on Pro plan. OAuth-centric detection covers same use case. | Full in v2 with tier detection |
| Shadow AI prompt monitoring (browser-level) | Track 2 browser extension, Gate 3 dependent. Marketing constraint: v1 detects AI tools via OAuth only (~40% coverage). Must not claim full monitoring. | v2/Track 2 conditional |
| Network-level discovery (NMAP, passive DNS) | Targets SaaS-first SMEs. On-prem network scanning is a different product category. | Not planned |
| Asset Inbox request workflow | No sprint budget in v1. | v2 backlog |
| Integration health score | No customer validation yet. | Post-v1 roadmap |
| Deepfake detection | Track 2, S7-S8. | v2/Track 2 |

---

## 4. Feature Catalog

---

### Feature 1 — User Account Discovery & Inventory (Google + M365)

**Description:** Automatically discovers all user accounts, service accounts, and groups from Google Workspace and Microsoft 365 tenants via admin APIs. 15-minute polling with incremental sync.

**User Story:** As an IT admin, I want to see every user account and group across our Google and M365 tenants in a single view so that I can identify orphaned accounts, ex-employee access, and unauthorized users without manually exporting from each provider.

**Acceptance Criteria:**
1. Within 60 minutes of initial OAuth connection, ≥90% of active user accounts are discovered and visible.
2. Incremental sync detects new accounts within 15 minutes of creation in source directory.
3. Each user record shows: email, display name, account status (active/suspended/deleted), account type (admin/standard/service/contractor), provider, groups, last sign-in date (where available), MFA enrollment status.
4. Partial sync failure (rate limit, token expiry for one provider) does not prevent the other provider's data from displaying. Per-provider error state shown in UI.
5. Account deletion/suspension in source directory reflected in SMESec within 15 minutes.
6. `GET /api/assets?type=user&provider=google` returns verified data from test tenant with correct pagination and filtering.

**Technical Notes:**
- Google: Admin SDK Directory API, domain-wide delegation service account. Scopes: `admin.directory.user.readonly`, `admin.directory.group.readonly`.
- M365: Microsoft Graph application permissions. Scopes: `User.Read.All`, `Group.Read.All`. Delta sync via `$deltaLink` token.
- Multi-tenancy: `workspace_id` on every row, RLS enforced at DB level. Tenant isolation CI test required on every PR.
- Rate limits: Google 1,500 req/min, Graph 10,000 req/10 min. Exponential backoff + delta sync mandatory — full re-sync on every poll is a critical trap.

**Dependencies:** S1 infrastructure (RDS, ECS, Keycloak, Secrets Manager).  
**Milestone:** Early Access (W8 / S2-S3). Google S2, M365 S3.  
**Confidence:** 88%

---

### Feature 2 — OAuth App Discovery & Shadow IT Detection

**Description:** Discovers all third-party applications authorized via OAuth grants on Google Workspace and M365, creating a complete shadow SaaS map across the tenant.

**User Story:** As an IT admin, I want to see every third-party app our employees have authorized with company credentials so that I can identify unauthorized SaaS, data-exfiltration risks, and OAuth apps I never approved.

**Acceptance Criteria:**
1. Within 60 minutes of initial connection, ≥90% of OAuth-authorized apps across all users are discovered.
2. Each OAuth app record shows: app name, publisher, OAuth scopes granted, number of users who authorized it, first/last authorized date, and risk score (Feature 5).
3. New OAuth app authorizations detected within 15 minutes.
4. Apps not on allow-list automatically flagged as "Pending Review."
5. Classification is deterministic: identical scopes + app ID = identical risk score across all tenants.
6. Apps authorized by multiple users deduplicated at app-ID level, per-user grant detail preserved.

**Technical Notes:**
- Google: `admin.directory.tokens` API (requires super admin consent at onboarding). Bulk listing — do not enumerate per-user in a loop.
- M365: `GET /users/{id}/oauth2PermissionGrants` + `GET /applications`. Requires `Directory.Read.All`.
- App identity normalized by `clientId` not display name.
- Webhooks (Graph Change Notifications + Google Push) added in S4 for sub-5-min latency.

**Dependencies:** Feature 1 (user inventory).  
**Milestone:** Early Access (W8 / S2-S3).  
**Confidence:** 88%

---

### Feature 3 — Basic Device Inventory (Agentless, Admin-Reported)

**Description:** Collects device enrollment data from Google Workspace (Chrome/Mobile Device Management) and M365 (Entra/Intune-enrolled devices) via admin APIs. No endpoint agent required. Provides minimum viable device visibility to satisfy ISO 27001 A.8.1.

**User Story:** As an IT admin, I want to see which devices are registered against our identity providers so that I can identify unmanaged or non-compliant devices without deploying a full MDM solution.

**Acceptance Criteria:**
1. Enrolled devices discovered and displayed within 60 minutes of initial connection.
2. Each device record shows: device name, OS type/version, last sync date, compliance status (per IDP), associated user, enrollment date.
3. Devices labeled "Admin-Reported (No Agent)" in UI to set correct expectations.
4. "Device Coverage" metric = (admin-reported devices / total active users) × 100%. Alert if coverage drops below 70%.
5. Read-only — no device management actions in v1.

**Technical Notes:**
- Google: `admin.directory.chromeosdevices`, `admin.directory.mobiledevices`. Scopes: `.device.chromeos.readonly`, `.device.mobile.readonly`.
- M365: `GET /devices`, `GET /deviceManagement/managedDevices`. Requires `DeviceManagementManagedDevices.Read.All`.
- **No new sprints required** — bundled into S2/S3 Google/M365 sync work using already-authorized API surface.
- Known limitation: Non-enrolled BYOD devices will not appear. Communicated clearly in UI and onboarding.

**Dependencies:** Feature 1 (same Google/M365 connection).  
**Milestone:** Early Access (W8 / bundled into S3).  
**Confidence:** 85%

---

### Feature 4 — Asset Classification Engine (Rule-Based Criticality + Sensitivity)

**Description:** Automatically assigns a criticality tier (Critical/High/Medium/Low) and sensitivity level (Restricted/Confidential/Internal/Public) to every discovered asset using a deterministic, auditable rule set. Manual override with immutable audit log.

**User Story:** As an IT admin, I want every asset to have an automatic classification I can review and correct so that I can prioritize which assets to protect first and demonstrate to auditors that we have a documented asset register.

**Acceptance Criteria:**
1. 100% of newly discovered assets receive automatic classification within 5 minutes of discovery.
2. Criticality tiers assigned by deterministic rules (no ML). Rule set version-controlled.
3. Default sensitivity = "Internal" for unmatched assets.
4. Manual override recorded with: who changed, when, previous value, new value, optional justification. Audit log is append-only (immutable).
5. Bulk update via CSV (`asset_id, criticality, sensitivity, justification`). Applied within 2 minutes; per-row error reporting.
6. Classification rules configurable by IT admin at tenant level. Changes take effect within 1 minute and are logged.

**Criticality Rule Set (v1 defaults):**

| Asset Type | Rule | Criticality |
|---|---|---|
| Admin accounts (Google/M365 global admin) | account_type = admin | Critical |
| Service accounts with OAuth grants | has active tokens + service type | High |
| Active users, MFA disabled | mfa_enrolled = false AND status = active | High |
| Standard users, MFA enabled | Default | Medium |
| Suspended / inactive accounts (>90 days) | last_sign_in > 90 days OR status = suspended | Low |
| AWS IAM root account | iam_type = root | Critical |
| AWS S3 bucket with public read | public_access = true | Critical |

**Sensitivity Rule Set (v1 defaults):**

| Signal | Sensitivity |
|---|---|
| OAuth app with financial scopes (QuickBooks, Xero) | Confidential |
| OAuth app with mail read/write scopes | Confidential |
| OAuth app with admin API scopes | Restricted |
| AWS S3 bucket tagged sensitivity=pii | per tag |
| Google Drive app with drive scope | Confidential |
| Unmatched | Internal |

**Technical Notes:**
- Go struct-based rule evaluator. Rules stored in `classification_rules` JSONB table. Priority-ordered evaluation, highest priority wins.
- Classification runs event-driven: new asset written to DB → EventBridge → classification job.
- Override model: explicit overrides take precedence over rule engine re-runs (`manual_override = true` flag).
- Audit log: append-only `classification_history` table. Backed to S3 with Object Lock (7-year retention).

**Dependencies:** Feature 1, 2 (assets must exist).  
**Milestone:** Early Access (W8 / S4).  
**Confidence:** 88%

---

### Feature 5 — OAuth Scope Risk Scoring

**Description:** Assigns a deterministic risk score (0–100) to each OAuth-authorized application based on OAuth permissions held, using a curated scope risk registry. Same app + same scopes = same score, always.

**User Story:** As an IT admin, I want each OAuth app to show a risk score based on its permissions so that I can immediately identify high-risk apps without reading OAuth scope documentation.

**Acceptance Criteria:**
1. Every OAuth app has a risk score 0–100.
2. Score computed exclusively from declared OAuth scopes — no ML. Identical inputs = identical output.
3. Scope risk registry covers ≥50 OAuth scopes at launch. Version-controlled, updatable without code deployment.
4. Score changes when app's scope list changes. Change triggers re-classification and is logged.
5. UI badge: Low (0-30 green), Medium (31-60 amber), High (61-85 orange), Critical (86-100 red). Tooltip lists highest-risk scopes.
6. Dashboard filterable and sortable by risk score.

**Scope Risk Registry (representative examples):**

| Scope | Risk Points | Rationale |
|---|---|---|
| `admin` / `*.admin` (Google) | +40 | Full tenant admin |
| `mail.readwrite` (M365) | +30 | Read/write all email |
| `https://mail.google.com/` | +30 | Read/write all Gmail |
| `drive` (Google Drive full) | +25 | Access all Drive files |
| `Directory.ReadWrite.All` (Graph) | +35 | Read/write entire directory |
| `User.ReadWrite.All` (Graph) | +25 | Modify all user accounts |
| `offline_access` + high-risk scope | +10 | Persistent background access |
| Financial API scopes (Xero, QuickBooks) | +35 | Access financial records |
| `openid`, `profile`, `email` only | +5 | Login-only, minimal risk |

**Technical Notes:**
- Registry in `oauth_scope_risk_registry` table. Admin-updatable via API with audit log.
- Additive scoring model (sum all matching points, cap at 100) — intentionally simple for auditability.
- Unknown scopes default to `+5` and flagged for registry review. Do not generate/guess scores.

**Dependencies:** Feature 2 (OAuth apps with scope lists).  
**Milestone:** Early Access (W8 / S4).  
**Confidence:** 88%

---

### Feature 6 — Shadow AI Tool Authorization Detection (OAuth-Based)

**Description:** Identifies AI tools (ChatGPT Enterprise, GitHub Copilot, Notion AI, Grammarly, Gemini apps, etc.) authorized via OAuth grants. Surfaces unsanctioned AI tool usage without browser-level monitoring. Coverage: ~40% of AI tools (OAuth-authenticated only).

**User Story:** As an IT admin, I want to see which AI tools our employees have authorized with company accounts so that I can assess data leakage risk from AI tool usage before browser-level monitoring is deployed.

**Acceptance Criteria:**
1. AI tool registry pre-loaded with ≥30 known OAuth-based AI tools at launch (name, OAuth client ID, vendor, risk category).
2. Any OAuth app matching the registry is automatically labeled "AI Tool — OAuth Authorized."
3. Dashboard has dedicated "AI Tools" filter showing only AI-tool-matching OAuth apps.
4. Each AI tool entry shows: tool name, users who authorized it, risk score, authorization date range, sanctioned status (IT admin-controlled).
5. Unsanctioned AI tool authorizations trigger the same alert pipeline as other Shadow IT (Feature 7) with additional `ai_tool: true` flag.
6. **Hard product requirement**: UI copy explicitly states "AI tools detected via OAuth authorization" — NOT "all AI tool usage monitored." This constraint applies to in-product copy, API responses, email alerts, and all marketing materials.

**AI Tool Registry (seed list):**

| Tool | Detection Signal | Risk |
|---|---|---|
| ChatGPT Enterprise | chat.openai.com OAuth app | High |
| GitHub Copilot | github.com copilot scopes | High |
| Google Gemini (Workspace Add-on) | workspace.googleapis.com/gemini | High |
| Notion AI | api.notion.so OAuth | Medium |
| Grammarly | grammarly.com OAuth | Medium |
| Otter.ai | otter.ai OAuth | Medium |
| Zapier AI | zapier.com OAuth | High |
| Slack AI | Slack app with AI capabilities | Medium |

**Coverage limitation (documented in product):** Direct web access to ChatGPT/Claude without OAuth — typing in browser — is NOT detectable via this method. Browser extension (Track 2, S5) required for full coverage. This distinction must be maintained in all external communications.

**Dependencies:** Feature 2, Feature 5.  
**Milestone:** Early Access (W8 / S4).  
**Confidence:** 85%

---

### Feature 7 — New OAuth App Alert Pipeline (<15-Minute Detection)

**Description:** Detects newly authorized OAuth apps within 15 minutes (polling) or sub-5 minutes (webhooks, S4) and delivers alerts via email and Slack. Polling-primary architecture; webhooks are a latency enhancement, not a reliability dependency.

**User Story:** As an IT admin, I want to receive an alert within 15 minutes whenever any employee authorizes a new OAuth app so that I can review and block data-exfiltration risks before significant exposure occurs.

**Acceptance Criteria:**
1. New OAuth app authorization → alert delivered within 15 minutes via email (SES) and Slack.
2. Alert contains: app name, authorizing user email, OAuth scopes, risk score, authorization timestamp, direct link to app review page.
3. Alert sent only if app is NOT on tenant allow-list. Approved apps do not re-alert.
4. Webhook-based sub-5-min detection (Google Push Notifications / MS Graph Change Notifications) delivered in S4. If webhook fails, polling fallback ensures ≤15 min delivery. Webhook failures monitored separately via CloudWatch alarm.
5. IT admin can configure alert threshold: all new apps / Medium+ risk / High+ risk only.
6. Alert delivery failure retried 3× with exponential backoff. All-retry failure → CloudWatch alarm.

**Technical Notes:**
- Polling: background sync (15-min cadence) diffs newly synced apps against `oauth_apps` table.
- Webhooks (S4): Google Push expires every 7 days. M365 Graph subscriptions expire every ~3 days. **Renewal is non-optional** — silent expiry = silent monitoring gap = trust breach. Automated renewal via Step Functions scheduled task with CloudWatch alarm on failure.
- Alert delivery: SES (email), SNS-to-Slack (Slack). Templates in DB, not hardcoded.
- Idempotency: duplicate webhook delivery of same event must not produce duplicate alerts. Use `event_id` deduplication.

**Dependencies:** Feature 2, SES/Slack integration (S1).  
**Milestone:** Polling alerts live in S2/S3. Webhook enhancement in S4 (Early Access W8).  
**Confidence:** 85%

---

### Feature 8 — Shadow IT Remediation Workflow (Approve / Block / Notify)

**Description:** Structured workflow for IT admin to review each newly detected OAuth app and take action: approve (add to allow-list), block (revoke OAuth grant at provider level), or notify user. All decisions are logged immutably.

**User Story:** As an IT admin, I want to take action on a newly detected OAuth app directly from the alert so that I can remediate shadow IT risk in under 5 minutes without manually navigating to each identity provider's admin console.

**Acceptance Criteria:**
1. From alert email or Slack message, IT admin reaches app review page in ≤2 clicks.
2. Three actions: **Approve** (adds to allow-list), **Block** (revokes all OAuth grants via provider API, with confirmation dialog listing affected users), **Notify User** (sends pre-templated email to authorizing user).
3. Block action: confirmation dialog shown before execution. Revocation performed via Google Admin SDK `tokens.delete` or MS Graph `DELETE /oauth2PermissionGrants/{id}`. Per-user outcome reported within 60 seconds. **M365 caveat displayed**: revocation revokes the grant but existing access tokens remain valid for 60-90 minutes.
4. Approve action: app added to allow-list with approver identity, timestamp, justification. Retroactively clears pending-review for all existing grants.
5. All actions logged: who acted, which app, action, timestamp, outcome. Immutable (S3 Object Lock).
6. Bulk remediation: select multiple pending apps, apply same action. Parallel processing; per-app outcome reported.

**Technical Notes:**
- State machine: `pending_review → approved | blocked | notified`. Stored in `oauth_app_decisions` table.
- Notify template: pre-written templates in DB; admin can customize. No LLM at runtime in v1.
- RBAC guard: only `shadow_it_reviewer` or `admin` role can execute Block or Approve.

**Dependencies:** Feature 2, Feature 7.  
**Milestone:** Early Access (W8 / S4).  
**Confidence:** 82%

---

### Feature 9 — Asset Inventory Dashboard (Search, Filter, Sort, Export)

**Description:** Unified web dashboard presenting all discovered assets in a searchable, filterable, sortable table with CSV export. Single pane of glass for IT admin visibility.

**User Story:** As an IT admin, I want a single dashboard where I can search and filter all our assets by type, provider, risk level, and classification so that I can answer any "what do we have?" question in under 30 seconds.

**Acceptance Criteria:**
1. Dashboard loads (LCP) in ≤3 seconds for tenants with up to 10,000 assets on 50 Mbps. API pages return in ≤500ms (p95).
2. Filter panel: asset type, identity provider, risk score range (slider), criticality tier, sensitivity level, account status, date ranges.
3. Full-text search on `name`, `email`, `app_name`. Results within 500ms. Minimum 2 characters.
4. Table sortable by all displayed columns. Default sort: criticality descending, then discovery date descending.
5. CSV export of current filtered view. Generated within 10 seconds for ≤10,000 rows. Includes `asset_id`.
6. Each asset row links to detail view: full metadata, classification history, associated OAuth grants, remediation guidance.
7. Early Access: Google + M365 assets. Commercial Launch: adds Slack + AWS. Unavailable providers show "Not connected" state.

**Technical Notes:**
- PostgreSQL full-text search (`tsvector`/`tsquery`). Composite index on `(workspace_id, asset_type, criticality, last_seen)`.
- Cursor-based pagination (not offset). Cursor = `(criticality, discovery_date, id)` tuple. Offset pagination is a trap at 10K+ rows.
- React/Next.js frontend with table virtualization (`react-virtual`).
- Server-side streaming CSV — do not buffer full dataset in memory.

**Dependencies:** Features 1–6.  
**Milestone:** Basic (Google + M365) at Early Access (W8 / S3). Full at Commercial Launch (W14 / S5).  
**Confidence:** 90%

---

### Feature 10 — Data Sensitivity Labeling (Metadata-Based, No Content Scan)

**Description:** Applies data sensitivity labels to assets using observable metadata signals — OAuth scopes, resource tags, account type, provider-reported classification — without reading any file or message content. Privacy-safe by design.

**User Story:** As a compliance officer or IT admin, I want each asset to carry a data sensitivity label derived from metadata so that I can produce a GDPR Article 30 record of processing and demonstrate which assets handle sensitive data — without accessing employee files.

**Acceptance Criteria:**
1. Sensitivity label assigned to 100% of assets using metadata rules only. Zero content read or transmitted to SMESec servers. This is a privacy guarantee.
2. 4-tier scale: Restricted (PII, financial, admin credentials), Confidential (internal communications, contracts), Internal (standard business data), Public (marketing, public-facing).
3. Metadata signals: OAuth scopes, AWS resource tags, M365 Purview sensitivity labels (read-only, E3+ tenants), Google Workspace shared drive classification (if admin-published).
4. Manual override by IT admin with justification. Override logged.
5. "Data Sensitivity Summary" widget shows asset counts per tier. Highlights Restricted assets with no access review in past 90 days.
6. Labels exported in compliance evidence (Feature 12) and mapped to ISO 27001 A.8.2.

**Technical Notes:**
- Same classification rule engine as Feature 4 (shared code path; sensitivity rules are a separate rule set).
- M365 Purview: `InformationProtectionPolicy.Read.All`. Requires E3/E5 or equivalent. Graceful degradation if not available.
- AWS: read S3 bucket tags and EC2 instance tags. No S3 object content reads.
- Privacy policy and onboarding flow must explicitly state: SMESec never reads email bodies, document contents, or chat messages in v1.

**Dependencies:** Features 1–3, Feature 4 (shared rule engine).  
**Milestone:** Commercial Launch (W14 / S4-S5 with Slack/AWS signals).  
**Confidence:** 85%

---

### Feature 11 — User-to-Asset Dependency Map (PostgreSQL Graph, Basic)

**Description:** Queryable dependency graph showing which users have access to which assets and the transitive access chain. Implemented in PostgreSQL with recursive CTEs — no graph database needed at SME scale (≤500 users, ≤50K edges per tenant).

**User Story:** As an IT admin or security officer, I want to visualize the access dependency chain for any user or asset so that I can perform blast radius analysis before offboarding and answer "what data could this user access?" in an audit.

**Acceptance Criteria:**
1. "Show Dependency Graph" from any user's detail page loads within ≤5 seconds for users with ≤500 access edges.
2. Graph shows: user → OAuth app → data scope, user → group → resource, user → IAM role → AWS resource. Maximum depth: 5 hops.
3. Blast radius view for offboarding: lists all assets that will lose access. Count displayed prominently.
4. API: `GET /api/graph/blast-radius?user_id={id}` returns JSON adjacency list in ≤2 seconds (≤10,000 nodes).
5. Zombie asset detection: assets with no active user connections flagged in graph view and "Zombie Assets" dashboard section.
6. Graph refreshed on each sync cycle (15-min cadence). Edge creation/deletion events logged.

**Technical Notes:**
- Graph storage: `asset_edges(workspace_id, source_id, target_id, edge_type, metadata JSONB)`. Indexed on `(workspace_id, source_id)` and `(workspace_id, target_id)`.
- PostgreSQL recursive CTE traversal. At 500 users × 50 OAuth apps = 26K edges — handles trivially with proper indexing. **No graph database required for SME scale.**
- **Repository abstraction layer required** (Sprint S12 acceptance criterion): `AssetGraphRepository` interface so Year 2 migration to graph DB (if needed for 2K+ users or enterprise tier) swaps implementation without touching business logic.
- Frontend: D3.js force-directed graph. Render max 200 nodes at once; collapse groups for larger graphs.

**Dependencies:** Features 1–3, 10; RBAC engine (S5) for role edges.  
**Milestone:** Full v1 (W26 / S12).  
**Confidence:** 65%

---

### Feature 12 — Compliance Evidence Export (ISO 27001, GDPR, SOC 2)

**Description:** One-click generation of machine-readable (JSON) and human-readable (PDF) compliance evidence packages mapping SMESec asset data to ISO 27001 controls A.8/A.9, GDPR Articles 30/32, and SOC 2 CC6.x controls.

**User Story:** As an IT admin preparing for ISO 27001 or GDPR review, I want to export a compliance evidence package with one click so that I can provide auditors with complete asset register evidence without spending days manually compiling spreadsheets.

**Acceptance Criteria:**
1. Export available for: ISO 27001 (A.8.1, A.8.2, A.9.2, A.9.4, A.12.4), GDPR (Art. 30, 32), SOC 2 (CC6.1, CC6.2, CC6.3, CC7.2).
2. Package contains: asset inventory snapshot (timestamped), classification summary, access review log (last 90 days), OAuth app allow-list decisions, offboarding records, JIT access audit trail.
3. JSON format follows versioned SMESec API schema. PDF is human-readable with tenant name.
4. Export generated within 30 seconds for ≤10,000 assets.
5. Each evidence item carries: timestamp, source system, responsible user, SHA-256 hash of evidence payload.
6. Exports stored in tenant-scoped S3 with 7-year retention (Object Lock, COMPLIANCE mode). Download link expires after 48 hours.
7. **UI and documentation must state**: SMESec evidence supports audit preparation — it does not constitute ISO/SOC certification. Certification requires an accredited third-party auditor.

**Control Mapping (partial):**

| Control | Evidence Provided |
|---|---|
| ISO 27001 A.8.1 | Full asset register with discovery date, owner, type |
| ISO 27001 A.8.2 | Classification per asset, rule applied, override log |
| ISO 27001 A.9.2 | Provisioned access per user, group memberships |
| ISO 27001 A.12.4 | Audit log completeness declaration, S3 retention proof |
| GDPR Art.30 | Data sensitivity labels, processing basis per asset type |
| SOC 2 CC6.1 | RBAC policy snapshot, access denied log sample |
| SOC 2 CC6.3 | Offboarding records with revocation timestamps |

**Technical Notes:**
- Evidence collection from audit log + asset snapshot tables. PostgreSQL `REPEATABLE READ` isolation for point-in-time consistency.
- PDF: HTML template → PDF via headless Chrome (chromedp) for auditable formatting.
- S3 path: `s3://smesec-evidence-{workspace_id}/compliance/{framework}/{year}/{month}/{export_id}.json`.
- SHA-256 computed on canonical JSON (sorted keys, no whitespace). Hash stored in DB and in export itself.

**Dependencies:** All other features. Compliance mapping (S10), report (S11).  
**Milestone:** Full v1 (W26 / S11).  
**Confidence:** 85%

---

### Feature 13 — Plain-English Remediation Guidance (Template-Based)

**Description:** For every detected risk — high-risk OAuth app, MFA-disabled admin, zombie account, shadow AI tool — surfaces a plain-English explanation and specific, actionable remediation step with action buttons. Tailored for non-security IT admin users.

**User Story:** As an IT admin without a security background, I want the system to tell me not just what the risk is but exactly what to do about it in plain language so that I can act immediately without researching security practices.

**Acceptance Criteria:**
1. Every alert, risk flag, and remediation recommendation includes a plain-English explanation (≤3 sentences) of why it's a risk and what to do.
2. Guidance templates maintained in DB table (not hardcoded). Updatable without code deployment.
3. Each template has an action button where applicable: "Revoke App", "Disable Account", "Enable MFA", "Start Offboarding", "Mark as Reviewed". Buttons wire directly to action APIs.
4. No jargon, no CVE numbers, no RFC references in primary guidance. Technical details in collapsible "Technical Details" section.
5. All templates reviewed by a non-technical tester before v1 launch. Acceptance test: non-security IT admin at 50-person company completes remediation without asking for help.

**Representative Templates:**

| Risk | Guidance |
|---|---|
| New OAuth app — High risk scopes | "An app called {app_name} was just authorized by {user_email} and can read all email in your company. Review it now: approve if it's legitimate, or block it to revoke access immediately." |
| MFA disabled on admin account | "{user_email} is an admin without two-factor authentication. If this account is compromised, an attacker has full control of your directory. Enable MFA now — it takes 2 minutes." |
| Zombie account (90+ days inactive) | "{user_email} hasn't signed in for {days_inactive} days. If this person left your company, their account should be disabled to prevent unauthorized access." |
| Shadow AI tool detected | "{app_name} is an AI tool that {user_email} has connected to your company account. AI tools can receive company data in their prompts. Review whether this is authorized." |
| AWS S3 bucket public read | "The S3 bucket '{bucket_name}' is publicly readable by anyone on the internet. If it contains company data, disable public access immediately." |

**Technical Notes:**
- `remediation_templates` table with fields: `risk_type`, `severity`, `title`, `guidance_text`, `technical_details`, `action_type`, `action_endpoint`.
- Template variables: Mustache-style. Server-side rendering — no client-side templating (XSS risk).
- No LLM-generated content at runtime in v1. LLM-enhanced guidance is v2 enhancement.

**Dependencies:** Features 4–8.  
**Milestone:** Basic templates at Early Access (W8 / S4). Full library at Commercial Launch (W14).  
**Confidence:** 90%

---

## 5. Integration Requirements

### 5.1 Google Workspace

| Attribute | Detail |
|---|---|
| APIs | Google Admin SDK Directory API (`admin.googleapis.com`) |
| Auth | OAuth 2.0 service account with domain-wide delegation |
| Required scopes | `admin.directory.user.readonly`, `admin.directory.group.readonly`, `admin.directory.group.member.readonly`, `admin.directory.tokens`, `admin.directory.device.chromeos.readonly`, `admin.directory.device.mobile.readonly` |
| Who must consent | Google Workspace super admin — prerequisite for onboarding. Onboarding wizard must guide customer through this step. |
| Rate limits | Directory API: 1,500 queries/min. Tokens API: empirically test. Mitigation: exponential backoff, incremental delta sync. |
| Sync frequency | 15-min polling (incremental). Push notifications from S4. Subscriptions expire every 7 days — **automated renewal mandatory**. |
| Key limitations | (1) Per-user token listing requires domain-wide delegation. (2) No last-OAuth-use timestamps. (3) No native delta sync for token grants — must diff against previous snapshot. |

### 5.2 Microsoft 365

| Attribute | Detail |
|---|---|
| APIs | Microsoft Graph API v1.0 (`graph.microsoft.com`) |
| Auth | Azure AD app registration, application permissions (no user sign-in required for sync) |
| Required permissions | `User.Read.All`, `Group.Read.All`, `GroupMember.Read.All`, `AuditLog.Read.All` (P1 required), `Directory.Read.All`, `Application.Read.All`, `Device.Read.All`, `DeviceManagementManagedDevices.Read.All`, `InformationProtectionPolicy.Read.All` (optional, E3+) |
| License trap | `AuditLog.Read.All` (last sign-in date) requires **Azure AD P1/P2** (M365 Business Premium, E3, E5). Detect at onboarding; gracefully hide last-sign-in column if P1 not present. |
| Rate limits | 10,000 requests/10 minutes. Use `$batch` (20 requests/batch) for bulk operations. Delta sync via `$deltaLink`. |
| Sync frequency | 15-min polling (delta sync). Graph Change Notifications from S4. Subscriptions expire every ~3 days — **automated renewal via Step Functions**. |
| Key limitations | (1) OAuth grant revocation revokes the grant but existing tokens valid for 60-90 more minutes — display this caveat to admin. (2) M365 guests must be labeled distinctly (type: `guest`). (3) Eventually consistent — new users may not appear for up to 15 min after creation. |

### 5.3 Slack

| Attribute | Detail |
|---|---|
| APIs | Slack Web API (`api.slack.com`) |
| Auth | OAuth 2.0 app installation by Slack workspace admin |
| Required scopes | `users:read`, `users:read.email`, `channels:read`, `team:read`, `apps.discoveries:read` (Business+ / Enterprise Grid only) |
| License trap | `apps.discoveries:read` requires **Slack Business+** ($12.50/user/mo). Pro and Free plans cannot enumerate installed apps. Majority of SME target customers (10-200 employees) will be on Pro. |
| Degraded mode | For tenants without Business+: Slack app inventory disabled. OAuth-centric detection via Google/M365 covers same shadow IT use case. UI displays: "Slack app inventory requires Business+. AI tool coverage via identity provider OAuth is still active." |
| Sync frequency | 15-min polling. Slack Events API webhooks optional (S5). |

### 5.4 Amazon Web Services

| Attribute | Detail |
|---|---|
| APIs | AWS Config, IAM, EC2, S3, RDS, Lambda |
| Auth | Cross-account IAM role (`sts:AssumeRole`). No long-term credentials stored. Customer executes SMESec CloudFormation template to create role. |
| Required permissions | `config:Get*`, `config:List*`, `iam:ListUsers`, `iam:GetAccountAuthorizationDetails` (bulk IAM snapshot — single call), `iam:ListAccessKeys`, `ec2:Describe*`, `s3:ListAllMyBuckets`, `s3:GetBucketAcl`, `s3:GetBucketTagging`, `rds:Describe*`, `lambda:List*` |
| Rate limits | Use `GetAccountAuthorizationDetails` for bulk IAM — far more efficient than per-user enumeration. S3 bucket ACL/tag reads are per-bucket (N calls for N buckets). |
| Key limitations | (1) AWS Config must be enabled in customer's account. Graceful fallback to direct describe APIs if Config disabled. (2) v1 supports single-account only — multi-account Organizations support deferred to v2. |

---

## 6. Non-Functional Requirements

### 6.1 Performance

| Metric | Requirement |
|---|---|
| Dashboard initial load (LCP) | ≤3 seconds (10,000 assets, 50 Mbps) |
| API page response (p95) | ≤500ms |
| Full-text search response (p95) | ≤500ms |
| CSV export (10,000 rows) | ≤10 seconds |
| New OAuth app alert delivery | ≤15 min (polling); ≤5 min (webhook) |
| OAuth grant revocation confirmation | ≤60 seconds |
| Compliance evidence export (10,000 assets) | ≤30 seconds |
| Dependency graph query (500 edges) | ≤5 seconds |
| Asset classification (new asset) | ≤5 minutes from discovery |

### 6.2 Security

| Requirement | Control |
|---|---|
| Tenant isolation | RLS on every table (`workspace_id`). CI test: cross-tenant read returns 0 rows. Must pass every PR. |
| Data encryption at rest | RDS KMS-encrypted. S3 SSE-KMS. EBS encrypted. |
| Data encryption in transit | TLS 1.2 minimum (1.3 preferred). HSTS enforced. |
| OAuth token storage | Encrypted with AWS KMS per-tenant CMK. Never logged or returned in API responses. Rotated every 90 days. |
| API authentication | Keycloak SSO, MFA mandatory (TOTP), session tokens expire 8 hours. |
| Secrets management | All secrets in AWS Secrets Manager. No `.env` files or hardcoded secrets. `git-secrets` pre-commit hook. |
| Audit log immutability | Security-relevant events → S3 Object Lock (COMPLIANCE mode, 7-year retention). |
| Input validation | Parameterized queries only (no SQL concatenation). Request bodies validated against OpenAPI schema. SAST (Semgrep) in CI. |

### 6.3 Scalability

| Dimension | Requirement |
|---|---|
| Assets per tenant | 10,000 (users + OAuth apps + devices + cloud resources) |
| Concurrent tenants | 100 active simultaneously |
| Graph edges per tenant | 50,000 (500 users × 100 assets each) |
| DB connections | 1,000 peak via PgBouncer (transaction mode) |

### 6.4 Availability

| Component | SLA |
|---|---|
| Asset Inventory API | 99.9% (43 min/month downtime) |
| Discovery pipeline | 99.5% (3.6 hours/month) |
| RDS PostgreSQL | 99.95% with Multi-AZ |
| Alert notifications | 99% delivery within SLA |

---

## 7. Open Questions (Must Resolve Before Relevant Sprint)

**OQ-1: MDM Scope Freeze — Decision required W1**  
Device inventory v1 = admin-reported only (no Jamf/Intune API). PO must formally accept this in writing and update customer-facing docs. If pilot customer requires Jamf integration, treat as paid add-on engagement. Owner: PM. Impact: S2/S3 scope.

**OQ-2: Slack Degraded Mode UX — Decision required before S5 (W9)**  
When Slack tenant is on Free/Pro (no `apps.discoveries:read`), what exactly does the user see? Options: (a) "Upgrade to Business+ to enable app inventory," (b) "Slack app inventory unavailable on your plan — AI tool coverage via Google/M365 OAuth still active," (c) hide Slack app inventory entirely. Owner: PO + Frontend Eng.

**OQ-3: M365 P1/P2 Adoption Among Pilot Customers — Validate by S3**  
Last sign-in date requires Azure AD P1/P2. If majority of pilot customers are on M365 Business Basic (no P1), zombie account detection degrades. PM to survey pilot customers in W1-W2. Owner: PM.

**OQ-4: Repository Abstraction Layer Design — Required as S12 Acceptance Criterion**  
`AssetGraphRepository` interface must be designed in S12 (not retrofitted). At what edge count or scale milestone triggers graph DB migration? Must be documented before S12 begins. Owner: Tech Lead.

**OQ-5: Shadow AI Marketing Copy Governance — Required before Early Access (W8)**  
All product copy, sales materials, and in-product text reviewed against this constraint: "Does this text imply we detect AI tool usage that is not via OAuth grants?" Phrases requiring sign-off: "detect AI tool usage," "see which AI tools employees use," "AI governance." Must be qualified with "via OAuth authorization" in all primary contexts. Owner: PO + Legal.

---

## 8. Competitor Gap Analysis (Final)

| Feature | SMESec v1 | Vanta | Drata | Nudge Security | JupiterOne |
|---|---|---|---|---|---|
| User account discovery (Google + M365) | ✅ 15-min sync | ✅ ~1hr | ✅ ~1hr | ✅ Real-time | ✅ Full |
| OAuth app discovery (all Slack tiers) | ✅ All tiers | ⚠️ Compliance-only | ⚠️ Compliance-only | ✅ Primary feature | ⚠️ Enterprise |
| Shadow AI detection (OAuth-based) | ✅ AI registry, ~40% coverage | ❌ | ❌ | ✅ Advanced | ⚠️ Ad hoc query |
| Shadow AI detection (browser-level) | 🔜 Track 2 v2 | ❌ | ❌ | ✅ | ❌ |
| OAuth scope risk scoring | ✅ Deterministic 0-100 | ⚠️ Basic flag | ⚠️ Basic flag | ✅ | ✅ Policy-based |
| Shadow IT remediation in-platform | ✅ Revoke from UI ≤60s | ❌ Links to console | ❌ Links to console | ⚠️ Alert only | ❌ Alert only |
| New app alert latency | ✅ ≤15 min | ⚠️ ~1hr | ⚠️ ~1hr | ✅ Real-time | ⚠️ Near real-time |
| Basic device inventory (agentless) | ✅ Admin-reported | ✅ Agent required | ✅ Agent required | ❌ | ⚠️ Cloud only |
| Asset classification (auto + override) | ✅ Rule-based, bulk CSV | ⚠️ Manual only | ⚠️ Manual only | ❌ | ✅ Policy-based |
| Data sensitivity labeling (no content scan) | ✅ Metadata, privacy-safe | ⚠️ Agent content scan | ⚠️ Content scan | ❌ | ✅ Tag-based |
| Slack app inventory | ⚠️ Business+ only | ⚠️ Business+ only | ⚠️ Business+ only | ✅ All plans (DNS) | ⚠️ Business+ only |
| AWS cloud asset inventory | ✅ EC2, S3, RDS, IAM | ✅ Full | ✅ Full | ❌ | ✅ Deep (primary) |
| User-to-asset dependency graph | ✅ PostgreSQL, basic | ❌ | ❌ | ❌ | ✅ Advanced |
| Compliance evidence export (one-click) | ✅ ISO 27001, GDPR, SOC 2 | ✅ SOC 2 focused | ✅ SOC 2 + ISO | ❌ | ⚠️ Query-based |
| Plain-English remediation | ✅ Template + action buttons | ⚠️ Generic | ⚠️ Generic | ✅ Strong | ❌ Technical only |
| SME entry pricing | ✅ Target <$200/mo | ❌ ~$800+/mo | ❌ ~$700+/mo | ⚠️ ~$300-500/mo | ❌ $1K+/mo |
| Automated offboarding (<5 min) | ✅ All 4 providers | ❌ Manual | ❌ Manual | ❌ Manual | ❌ Manual |

**SMESec differentiators:**
1. **Shadow IT remediation in-platform** — only product that revokes OAuth grants directly via API with audit trail
2. **Agentless device inventory** — no endpoint agent install friction (Vanta/Drata require agent)
3. **Shadow AI tool registry** — dedicated AI tool detection overlaid on OAuth discovery; no SME competitor offers this
4. **Automated offboarding + JIT access** — no direct SME competitor at this price point
5. **Entry pricing** — only viable option for 10-100 employee segment at <$200/month

**Gaps to acknowledge honestly:**
- Nudge Security has broader Shadow AI coverage (DNS-based, catches direct web access). SMESec v1 covers OAuth only (~40%). Must not over-claim.
- JupiterOne has significantly more mature dependency graph for complex AWS environments. SMESec targets simple SME environments.
- Slack app inventory is Business+ only on ALL platforms — not a competitive disadvantage, but must be communicated at sales.

---

## 9. Pre-Sprint-1 Decision Checklist

These are non-optional decisions that must be completed by end of Week 1:

| Decision | Owner | Deadline | Consequence if Missed |
|---|---|---|---|
| MDM scope formally frozen as v2 (PO sign-off in writing) | PM | W1 Day 5 | S2 risks scope creep |
| Google Admin SDK service account created + APIs enabled | Tech Lead | W1 Day 3 | API approval takes 48-72h; blocks S2 start |
| Microsoft Graph app registration + admin consent flow designed | Tech Lead | W1 Day 3 | M365 auth UX underestimated; S3 delays |
| AWS `smesec-readonly` IAM role created in test accounts | Backend Eng | W1 Day 5 | Blocks S5 AWS discovery |
| Pilot customer outreach: 10 SME targets contacted | PM | W1 Day 5 | Qualification by W8 impossible otherwise |
| Slack degraded-mode product decision documented | PO + Frontend | W1 Day 5 | S5 frontend scope undetermined |
| Track 2 extension fallback confirmed (regex-only if Gate 1 fails?) | Tech Lead | W1 Day 5 | S5 contingency-dependent |
| Shadow AI marketing copy constraint acknowledged by all stakeholders | PO + Legal | W1 Day 5 | Early Access (W8) marketing materials must be reviewed |

---

*This document is the authoritative specification for Asset Inventory & Classification. All implementation in Sprints S2–S12 related to this module must be traceable to a feature, acceptance criterion, or NFR in this document. Deviations require a change request reviewed by PO + TA.*
