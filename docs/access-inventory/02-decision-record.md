# Asset Inventory — Decision Record

**Date:** 2026-05-28  
**Status:** Approved  
**Session:** Multi-agent research (3 agents × 2 rounds — 6 invocations total)  
**Stakeholders:** Product Owner · Technical Advisor · Project Manager  
**Related:** [01-research-synthesis.md](01-research-synthesis.md) | [01-pm-risk-analysis.md](01-pm-risk-analysis.md)

---

## Executive Summary

Session này thực hiện fresh research về Asset Inventory key requirement thông qua 3 agents song song với 2 vòng cross-iterate. **Kết quả**: consensus về product approach, feature set, và technical architecture — với 5 quyết định lớn khác so với PM risk analysis cũ.

**Product anchor**: "Your team used 47 apps. 3 AI tools had access to your customer data folder. You didn't know. Now you do — and you can fix it in 5 minutes."

---

## 1. Bối Cảnh & Vấn Đề

### Tại Sao Cần Research Lại

- PM risk analysis cũ (`01-pm-risk-analysis.md`) tập trung vào sprint feasibility, không phải market validation
- Old plan thiếu: competitor analysis, customer pain point ranking, feature value vs effort
- 3 sprint plan cũ có Sprint 5 overload (Slack + AWS + OPA RBAC trong 1 sprint) — needs re-scoping
- Asset Inventory là *foundation layer* — tất cả Access Governance, compliance, offboarding đều phụ thuộc vào độ chính xác của layer này

### Scope của Key Requirement

Asset Inventory = continuous, automatic discovery và classification của:
- **Identity assets**: users, groups, service accounts (Google Workspace, M365)
- **SaaS assets**: mọi app employees dùng — OAuth-connected + shadow IT
- **Cloud assets**: AWS resources (v1 scoped: 5 security posture checks)
- **Shadow AI assets**: AI tools employees dùng với company data

---

## 2. Phân Tích Đối Thủ

### Primary Competitor: JupiterOne (Không Phải Vanta/Drata)

**Quyết định**: Benchmark primary competitor là **JupiterOne** ($500-3K/month SMB tier).

**Lý do**:
- Vanta/Drata = compliance evidence tools; ITAM tools (Lansweeper) = device inventory only
- JupiterOne = full-stack asset inventory với identity + SaaS + cloud; closest to SMESec vision
- JupiterOne's weakness: JQL graph query language requires security literacy — IT admin at 100-person SMB không dùng được

**SMESec phải beat JupiterOne trên:**
1. Time to first value: JupiterOne 2-4 weeks; SMESec <30 minutes
2. UX: Pre-answered findings; không cần viết JQL
3. Shadow AI discovery: JupiterOne partial; SMESec first SMB-priced tool
4. Mobile app: Zero competitors; SMESec có
5. Giá: $3-8/user/mo vs JupiterOne $5-30/user/mo

### White Space SMESec Khai Thác

| Gap | Severity | Đây là cơ hội của SMESec |
|---|---|---|
| Shadow AI discovery ở SMB price | **Critical 2026** | Không ai làm; first-mover advantage |
| Non-expert UX cho full-stack inventory | High | JupiterOne cần security literacy; SMESec không |
| Mobile app incident response | Medium | Zero competitors |
| Asset inventory + offboarding + compliance bundle | High | Requires 3+ separate tools hiện tại |

---

## 3. Quyết Định: Product Approach

### Core Approach

**"Nothing in your digital environment is invisible to IT — and everything actionable has a one-click fix."**

Không phải "comprehensive asset management platform." Không phải "security dashboard." Là cái cảm giác: *IT admin luôn có câu trả lời chính xác cho câu hỏi "ai có access vào cái gì?"*

**The product story (v1):**
> "Reza just resigned. One click. Here are all 47 apps he had access to, including 3 AI tools that were connected to your customer data folder. Offboarding complete in 4 minutes 23 seconds. PDF audit trail attached."

### 5 Design Principles (Converged by All 3 Agents)

1. **<30 minutes to first value** — Nếu customer không thấy first asset trong 30 phút, họ churn trước khi thấy deep value. Onboarding wizard là product, không phải feature.
2. **Depth before breadth** — Google Workspace at >95% coverage. Không thêm M365 cho đến khi Google có zero bugs 30 ngày. Không thêm provider tiếp theo cho đến khi provider trước stable.
3. **Pre-answered questions, not open-ended queries** — Mọi data point phải trả lời "so what?" với một action button. Không build query interface (JupiterOne trap).
4. **Agentless first** — Zero endpoint agents trong v1. 90%+ SMB asset surface discoverable via OAuth + API grants.
5. **Human-confirmed destructive actions** — Không auto-revoke bất kỳ thứ gì. Mọi bulk action phải show blast radius trước khi confirm.

---

## 4. Quyết Định: Feature Scope

### 4.1 Must-Have v1

| Feature | Customer Pain Solved | Sprint |
|---|---|---|
| **Onboarding wizard** (<30 min first asset) | 80% vs 30% retention difference | 7 |
| **SaaS discovery — Google Workspace** (users + OAuth apps + last login) | Foundation; aha moment | 2 |
| **SaaS discovery — M365** (users + app grants + activity) | 50% of SMB market | 3-4 (2 sprints) |
| **Shadow IT allow-list + alerts** (new app → Slack/email <15 min) | Closes discover→action loop | 4 |
| **Shadow AI discovery via OAuth** (which AI tools have what scopes) | 2026 critical gap; data already collected | 4 (free) |
| **Pre-seeded vendor catalog 500+ apps** (<10% Unknown acceptance criterion) | Prevents alert fatigue on day 1 | 4-5 |
| **Identity inventory** (users, groups, admin status, service accounts) | Foundation for offboarding + compliance | 2-3 |
| **Stale account detection** (>90 days inactive) | Ex-employee access — unlock event #1 | 5 |
| **"Forgotten access" dashboard** (ex-employees + active accounts) | Buyer emotional resonance; zero extra data | 5 |
| **6 deterministic findings** + **risk score rollup** (0-100) | SOC 2 evidence; actionable for non-security staff | 5 |
| **Zombie Account Cost Recovery** (inactive × license price) | Immediate ROI; no billing API needed | 5 |
| **Automated offboarding** (Google + M365 + flagged OAuth, <5 min) | Anchor feature; closest competitor: BetterCloud (Google/M365 only) | 6 |
| **Compliance evidence export** (ISO 27001 A.8/A.9 + SOC 2 CC6.1) | Saves $5-20K audit prep per customer | 6 |
| **AWS Security Posture** (5 checks: IAM orphaned users, public S3, root MFA, key rotation, SG 0.0.0.0/0) | Critical cloud findings; complete coverage within bounded scope | 7 |

### 4.2 Should-Have v1 (Include nếu capacity)

| Feature | Notes |
|---|---|
| Ramp + Brex expense integration | Discovers paid-but-no-OAuth apps; license waste signal. 1-2 days each. |
| License waste — M365/Google seat calculation | Requires Usage Reports APIs; "Zombie Account" expansion |
| Dependency graph (basic) | User → app → cloud resource (3 hops); already Sprint 12 in old plan |
| Mobile app (Flutter) | Zero competitor; incident response + offboarding approval |

### 4.3 Deferred to v2

| Feature | Lý Do | Effort |
|---|---|---|
| Full cloud inventory (AWS Config + Azure + GCP) | Multi-cloud coverage gap = trust erosion in v1 | 6-8 weeks |
| Data classification — SaaS PII scanning | 83hr/tenant initial scan; $150/tenant DLP cost; GDPR legal review; 5-15% false positive rate | 4-6 months |
| Full license waste (billing APIs: Zoom, Salesforce, etc.) | Per-vendor API; 2-3 sprints each | 4-6 sprints |
| Browser extension content monitoring | Conditional on MV3 gate + GDPR legal review (Track 2) | Track 2 |
| Expensify/Concur integration | Older APIs; v2 after Ramp/Brex validated | 3-4 days |

### 4.4 Removed / Never

| Feature | Lý Do |
|---|---|
| Agent-based endpoint discovery | MDM required; SMB deployment friction quá cao |
| Auto-revocation (shadow IT, inactive users) | Blast radius risk; trust killer |
| JQL / graph query interface | JupiterOne trap — builds for security team, not IT admin |
| ML risk scoring | No training data; false positive rate kills trust |
| Inline CASB / network traffic analysis | Proxy infrastructure; enterprise complexity and cost |
| Raw credit card number ingestion | PCI DSS risk |

### 4.5 Shadow AI: Hai Features Bị Nhầm Lẫn Thành Một

| Feature | Mechanism | Timeline | Verdict |
|---|---|---|---|
| **Shadow AI discovery** (what tools have OAuth access to what data) | OAuth scope analysis — data collected in Sprint 4 | v1 Must-Have | **Ship — zero extra work** |
| **Shadow AI content monitoring** (what data employees type/paste) | Browser extension MV3 — domain + form POST | Track 2 conditional | **Gate decision Sprint 1 Week 1** |

**Critical product disclosure**: Khi extension không được install, dashboard phải hiển thị rõ: *"AI tool discovery: partial coverage — browser activity monitoring not active."*

---

## 5. Quyết Định: Technical Architecture

### 5.1 Discovery Architecture

**Agentless-first, API polling + selective event streams.** Không cần endpoint agents cho v1.

| Source | Scope v1 | Auth Method | Key Trap |
|---|---|---|---|
| Google Workspace Admin SDK | Full identity + SaaS | Service account + DWD | Reports API 90-min delay; NO real-time events |
| Microsoft Graph | Full identity + SaaS | App registration + admin consent | Delta token expires 7 days if not consumed |
| AWS IAM/S3/EC2 | 5 security posture checks only | IAM role + CloudFormation one-click | Never ask SMBs to configure IAM manually |
| GitHub | Org members + repos | GitHub App (not OAuth App) | Fine-grained permissions; 2-min install |
| Slack | Users + apps (Business+ only) | OAuth + admin scopes | NO admin API on free tier — disclose clearly |

### 5.2 Sync Architecture: Dual-Speed Pipeline

```
FAST PATH (<5 min): AWS EventBridge, M365 Graph Change Notifications, Slack Events API
STANDARD PATH (15-60 min): Google 15-min poll, M365 delta queries, AWS reconciliation
Both → SQS → ECS workers → PostgreSQL assets table
```

### 5.3 Build vs. Buy

| Component | Quyết Định | Lý Do |
|---|---|---|
| Auth (platform SSO) | **Buy: Keycloak** (existing) | Không bao giờ build OAuth/OIDC |
| Multi-tenant isolation | **Build: PostgreSQL RLS** | `workspace_id` + CI cross-tenant test |
| Asset graph | **Build: PostgreSQL CTEs** | Đủ cho 100K assets, 1M relationships; graph DB premature |
| Background sync | **Buy: SQS + ECS workers** | Dead-letter queues; at-least-once delivery |
| Classification | **Build: rule-based 6-stage** | Không có training data cho ML; 80-85% accuracy với rules |
| Risk scoring | **Build: deterministic rollup** | Findings → weighted sum → score; auditable, no false positives |
| PDF reports | **Build: Puppeteer** (existing) | Full template control |
| Secrets | **Buy: AWS Secrets Manager** (existing) | KMS rotation, audit |

### 5.4 Technical Non-Negotiables (Phải Đúng Từ Day 1)

1. **Multi-tenant RLS với CI enforcement** — Mọi table mới phải có RLS policy + cross-tenant isolation test trước khi PR merge. Không exceptions.
2. **Delta/incremental sync từ Sprint 2** — Không phải optimization. Full sync at 100 tenants × 500 users = rate limit guaranteed in production. `sync_jobs.next_sync_token` là first-class resource.
3. **Vendor catalog ≥500 apps trước Sprint 5** — Acceptance criterion: <10% Unknown cho top-200 SMB tools. Không có catalog = alert fatigue ngay day 1.
4. **Google Reports API 90-min delay** — Không thể real-time detect new OAuth grants. Hybrid: polling mỗi 15 min + manual "check now" button. Marketing không được claim "real-time."
5. **M365 delta token 7-day expiry recovery** — `410 Gone` response = fall back to full sync + log. Không để silent failure.
6. **AWS CloudFormation one-click** — Không bao giờ yêu cầu SMB admin configure IAM cross-account role manually. Một-click CFN template.
7. **Onboarding wizard acceptance criterion: timed walkthrough** — Sprint 7 acceptance test: ba new pilot customers kết nối Google Workspace và thấy first assets trong <30 phút, được ghi lại.

### 5.5 First-Sync Fast Path Architecture

```
T+0:00  Google OAuth consent completed
T+0:30  Phase 1: Users visible (bulk insert via PostgreSQL COPY)
T+3:00  Phase 2: Recent OAuth apps visible (30-day Reports API fast path)
T+7:00  "AHA MOMENT" — "23 apps discovered, 4 need review"
T+25:00 Phase 3 (background): Full history + risk scores + findings
```

**Key decisions:**
- Reports API (1-3 calls) trước `tokens.list` (500 calls) — fast path không block aha moment
- SSE progressive rendering — customer thấy users trước, apps sau, scores cuối
- Risk scores decouple từ first render — calculated T+20, không block T+7

---

## 6. Quyết Định: Delivery Approach

### 6.1 Revised Timeline (4-Person Team)

| Sprint | Tuần | Mục Tiêu | Key Deliverable |
|---|---|---|---|
| **Sprint 1** | W1-2 | Infrastructure + Auth + MV3 Gate | Keycloak SSO + multi-tenant DB + CloudFormation template |
| **Sprint 2** | W3-4 | Google Workspace Sync | Users + OAuth apps visible; >95% coverage target |
| **Sprint 3** | W5-6 | M365 Sync Phase 1 | Users + groups + basic app grants; delta token handling |
| **Sprint 4** | W7-8 | M365 Phase 2 + Shadow IT + Shadow AI tagging | Alerts live; AI tools tagged; allow-list management |
| **Sprint 5** | W9-10 | Findings + Risk Score + Zombie Accounts | 6 deterministic rules + score + cost recovery card |
| **Sprint 6** | W11-12 | Offboarding + Compliance Export | Automated offboard <5 min + ISO/SOC 2 evidence PDF |
| **Sprint 7** | W13-14 | Onboarding Wizard + AWS Posture | <30 min first value + 5 security checks |
| **Sprint 8** | W15-16 | Pilot Support + Security Hardening | Pen test remediation + pilot NPS >40 |

**Week 14: PILOT READY (3-5 design partners)**  
**Week 20: FIRST PAYING CUSTOMER**

### 6.2 "Ready for First Paying Customer" Definition

| Criterion | Threshold |
|---|---|
| Google Workspace coverage | >90% users + OAuth apps trong 1 giờ |
| M365 coverage | >90% users + licensed apps trong 1 giờ |
| Classification accuracy | <10% Unknown cho top-200 SaaS apps |
| Time to first asset | <30 phút từ OAuth consent |
| Automated offboarding | Full revocation <5 phút |
| Security posture | Zero Critical findings từ pen test |
| Uptime (30-day rolling) | >99.5% |
| Pilot NPS | >40 (3 pilot customer surveys) |

### 6.3 3 Highest-Risk Items

| Risk | Probability | Impact | Mitigation |
|---|---|---|---|
| M365 delta token expiry handling (Sprint 3) | 60% | High | `410 Gone` = full sync fallback + integration test với 8-day gap simulation |
| Pilot customer sourcing (external) | 40% | Critical | Outreach bắt đầu Week 1; 2 LOI trước Sprint 3; Google Workspace customers first |
| Onboarding wizard scope creep (Sprint 7) | 70% | Medium | Sprint 4 backlog grooming defines acceptance criteria; FE mockups in Sprint 5; no new scenarios post-Sprint 6 |

---

## 7. Quyết Định Đảo Ngược So Với Plan Cũ

| PM Risk Analysis Cũ | Quyết Định Mới | Lý Do |
|---|---|---|
| Sprint 5: Slack + AWS + OPA RBAC cùng 1 sprint | Tách riêng: Slack/AWS (Sprint 4-5) + RBAC (separate concern) | 3-sprint scope compressed into 1; confirmed infeasible by TA |
| AWS = full inventory (EC2, S3, RDS, Lambda) | AWS = 5 security posture checks only (IAM, S3, root MFA, key rotation, SG) | Coverage gap = trust erosion; PMO rule: don't ship 60% products |
| Sprint 12: Dependency mapping | Giữ nguyên Sprint 12 | No change; already correct |
| M365 trong 1 sprint (Sprint 3) | M365 trong 2 sprints (Sprint 3-4) | All market evidence confirms this; APIs fundamentally different |
| Risk scoring không được đề cập trong PM risk analysis | 6 deterministic findings + risk score rollup | Deterministic beats ML; no false positive problem if findings-based |
| Pilot estimate: 10 weeks | Revised: 14-16 weeks | M365 complexity + onboarding wizard scope |
| Shadow IT = feature trong Sprint 4 | Shadow IT + Shadow AI = same sprint (Sprint 4, free) | Shadow AI via OAuth requires zero extra API calls |

---

## 8. Open Questions

| # | Question | Owner | Deadline | Blocking? |
|---|---|---|---|---|
| 1 | **MV3 browser extension gate**: Prototype pass Week 1? | TA + Track 2 | Sprint 1 W1 | Determines browser ext v1 vs v2 |
| 2 | **GDPR legal review** for browser extension (EU markets = 40-60% SMB target) | PM + Legal | Sprint 1 W1 | If fails → extension never in EU |
| 3 | **v1 ACV pricing**: $3-8K (realistic v1 features) vs $8-25K (full vision)? Pilot LOI framing? | PO + PM | Sprint 3 | Pilot LOI |
| 4 | **Pilot customer sourcing**: 2 LOI trước Sprint 3. Google Workspace customers first. | PM | Sprint 3 | Sprint 7 milestone |
| 5 | **Pen test vendor**: Contract bởi Sprint 4 (6-tuần scheduling lead time) | PM | Sprint 4 | Sprint 8 hardening |
| 6 | **Ramp/Brex integration**: Confirm Should-Have v1 trong Sprint 3 backlog grooming | All | Sprint 3 | No |
| 7 | **Slack Business+ qualification**: % of pilot customers qualify? Determines v1 vs v2 automated offboarding | PM | Sprint 4 | No |

---

## 9. Success Metrics

### Activation Metrics (First 30 Days)
| Metric | Target |
|---|---|
| Time to first inventory completed | <30 minutes from OAuth consent |
| Apps discovered vs customer estimate | >3x (the aha moment) |
| % of customers connecting 2+ providers within 7 days | >70% |
| Shadow apps flagged per customer | >15 on first scan |

### Engagement Metrics (90-Day Health)
| Metric | Target |
|---|---|
| 90-day retention (cohort) | >80% |
| Stale accounts actioned per month | >5 per customer |
| Shadow IT allow-list actions (approve/block) | >10 per customer per month |
| Pilot NPS | >40 |

### Security Outcome Metrics
| Metric | Target |
|---|---|
| Time to full offboarding | <5 minutes (vs 3-5 day industry baseline) |
| Shadow apps remediated within 60 days | >50% of flagged apps |
| MFA coverage post-product adoption | >95% (from baseline 40-60%) |
| Orphaned accounts post-90 days | <5% (from baseline ~20%) |

---

*Session: 2026-05-28 | Method: 3 agents × 2 rounds | Files: [01-research-synthesis.md](01-research-synthesis.md) | Next step: SPEC*
