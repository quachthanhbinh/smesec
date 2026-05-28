# Access Governance — Decision Record

**Date:** 2026-05-28  
**Status:** Approved  
**Session:** Multi-agent research (3 agents × 2 rounds — 6 invocations total)  
**Stakeholders:** Product Owner · Technical Advisor · Project Manager  
**Related:** [02-feature-research-synthesis.md](02-feature-research-synthesis.md)

---

## Executive Summary

Session này thực hiện fresh research về Access Governance key requirement thông qua 3 agents song song với 2 vòng iterate (Round 1: independent research; Round 2: cross-challenge). Kết quả: consensus về product approach, feature set, technical architecture, và các quyết định scope bị đảo ngược so với plan cũ. **Đối thủ thực sự là Lumos (không phải Vanta)**. Product anchor là "nothing falls through the cracks" offboarding cho Google + M365 — không phải full automation.

---

## 1. Bối Cảnh & Vấn Đề

### 1.1 Tại Sao Cần Research Lại

- Plan cũ (Track 1, Sprints 4-6) thiết kế dựa trên assumptions chưa được validate bởi market research
- Scope quá rộng: 4 providers automated offboarding + AWS IAM trong 2 sprints → không khả thi
- Competitor benchmark sai: so sánh với Vanta/Drata (compliance tools) thay vì Lumos (access governance)
- Risk scoring được đề xuất trong v1 mà không xét đến false positive problem

### 1.2 Pain Points Được Validate (SMB 50–500 nhân viên)

| Rank | Pain Point | Evidence |
|---|---|---|
| #1 | **Orphaned access sau khi nhân viên nghỉ** | 69% tổ chức có incident từ ex-employee (Ponemon 2023); avg 3–5 ngày để revoke đầy đủ |
| #2 | **Shadow IT / unauthorized SaaS** | Avg 4.5 unauthorized apps/employee; 63% dữ liệu nhạy cảm ở unsanctioned apps (Netskope 2024) |
| #3 | **Access sprawl / quá nhiều quyền** | Avg 17+ SaaS apps/employee; chỉ 36% license được dùng thực sự |
| #4 | **Không có audit trail cho compliance** | 73% SMB fail SOC 2 lần đầu vì thiếu access control evidence (Vanta 2023) |
| #5 | **Over-provisioning khi onboarding** | "Copy access từ người trước" = kế thừa 2 năm access creep |

---

## 2. Phân Tích Đối Thủ

### 2.1 Ma Trận Tính Năng

| Feature | Vanta | Drata | Nudge Security | BetterCloud | Zluri | **Lumos** | Entra ID Gov |
|---|---|---|---|---|---|---|---|
| SaaS asset inventory | ⚠️ | ⚠️ | ✅ | ✅ | ✅ | ✅ | ⚠️ |
| Shadow IT discovery | ⚠️ | ❌ | ✅ Core | ✅ | ✅ | ⚠️ | ⚠️ |
| Shadow IT remediation | ❌ | ❌ | ❌ nudge only | ✅ | ✅ | ⚠️ | ⚠️ |
| Automated offboarding | ❌ | ❌ | ❌ | ✅ | ✅ | ✅ | ✅ Complex |
| Access reviews | ✅ | ✅ | ❌ | ✅ | ✅ | ✅ | ✅ |
| JIT / privileged access | ❌ | ❌ | ❌ | ❌ | ⚠️ | ⚠️ | ✅ PIM |
| Compliance evidence (SOC 2) | ✅ Core | ✅ Core | ❌ | ❌ | ❌ | ❌ | ⚠️ |
| **Mobile app** | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |
| Non-expert UX (SMB-native) | ✅ | ✅ | ✅ | ⚠️ | ✅ | ✅ **Best** | ❌ |
| SMB pricing (50 users) | ~$8–15K/yr | ~$8–15K/yr | ~$2.4K/yr | ~$6–12/usr/mo | ~$3–5K/yr | ~$5–10/usr/mo | ~$3.6K/yr |

### 2.2 Competitor Thực Sự: Lumos (Không Phải Vanta)

**Quyết định:** Benchmark primary competitor là **Lumos**, không phải Vanta/Drata.

**Lý do:**
- Vanta/Drata = compliance evidence tools. SMESec Access Governance = access automation platform. Hai thị trường khác nhau.
- Lumos làm: self-service access requests, automated deprovisioning, app catalog, access reviews, modern UX ($5–10/user/mo)
- Lumos KHÔNG có: SOC 2/ISO 27001 compliance evidence, shadow IT remediation, mobile app

**SMESec phải beat Lumos trên:**
1. Compliance evidence (SOC 2 + ISO 27001 tích hợp) — Lumos không có
2. Shadow IT enforce (allow-list management) — Lumos chỉ discover
3. Mobile app cho incident response — Zero competitor có
4. Giá thấp hơn: target $3–5/user/mo vs Lumos $5–10/user/mo

### 2.3 Market Gaps (White Space)

| Gap | Mức Độ | Mô Tả |
|---|---|---|
| Offboarding automation ở SMB pricing | Critical | BetterCloud/Zluri làm được nhưng $6–12/user/mo. Okta cần Okta làm IdP. Không có ai <$5/user/mo. |
| Shadow IT discover → remediate loop | High | Nudge Security chỉ discover. BetterCloud remediate nhưng là IT ops tool. Không có affordable tool. |
| Unified platform | High | SMBs cần Vanta + Nudge + BetterCloud = $23–32K/yr. Không có 1 tool ở <$6/user/mo. |
| Mobile app incident response | Medium | Zero competitors. IT admin không thể respond breach lúc 10 giờ tối. |

---

## 3. Quyết Định: Product Approach

### 3.1 Core Approach

**"Nothing falls through the cracks"** — không phải "full automation"

Mọi access đều được kiểm soát và không có gì bị bỏ sót. Automated khi có thể; human-confirmed khi cần; checklist + deep-links khi automation không khả thi.

**3 luồng chính:**

```
1. OFFBOARDING:
   IT admin mark as leaver
        ↓
   Parallel: Google Workspace revoke + M365 revoke (automated, <5 min)
        ↓
   Checklist: Slack + AWS + GitHub (manual, deep-links, checked off)
        ↓
   PDF report: ✅ Automated | ⚠️ Manual (done) | ❌ Failed (retry)

2. SHADOW IT:
   System discovers OAuth apps (polling 15 min + webhooks)
        ↓
   IT admin receives alert: "New app: Notion - 3 users"
        ↓
   IT admin reviews → Approved / Blocked / Pending
        ↓
   Blocked apps: API revocation via Google Admin SDK / MS Graph
        ↓
   Audit trail: decision + timestamp + reviewer

3. COMPLIANCE FINDINGS:
   6 deterministic rules run continuously
        ↓
   Findings surface in dashboard (no ML, no false positives)
        ↓
   Each finding: direct "Fix it" button + SOC 2 control mapping
```

### 3.2 Design Principles (Converged by All 3 Agents)

1. **Human-confirmed destructive actions** — Mọi revocation đều cần explicit confirmation. Zero auto-revocations trong v1. Blast radius hiển thị trước khi confirm.
2. **Pull-based polling + hybrid webhooks** — IT admin không cần configure webhook. System tự đăng ký trong OAuth consent flow.
3. **Graceful degradation per provider** — Product hoạt động bình thường khi một provider down. Status hiển thị per-provider.
4. **Saga pattern cho offboarding** — Non-atomic: mỗi provider step commit độc lập. Failure không rollback step đã succeed.
5. **Zero configuration** — Connect qua 2 clicks OAuth consent → system tự discover mọi thứ.

---

## 4. Quyết Định: Feature Scope

### 4.1 Must-Have v1 (Launch Gate — Không Thể Thiếu)

| Feature | Lý Do Must-Have |
|---|---|
| Asset inventory: Google + M365 users, groups, OAuth apps | Foundation cho mọi thứ; "wow moment" ngày đầu |
| Asset inventory: Slack + AWS (discovery/read-only) | Hiển thị trong inventory; KHÔNG phải deprovisioning |
| Shadow IT discovery + allow-list management | Closes discover→act loop mà không competitor nào làm được ở SMB pricing |
| Automated offboarding: Google + M365 (<5 min, human-initiated) | #1 ROI demonstrable trong market; anchor feature |
| Offboarding checklist: Slack + AWS + GitHub (manual, deep-links) | Audit trail completeness; không bỏ sót bất kỳ hệ thống nào |
| Offboarding PDF report (per-provider status) | SOC 2 evidence artifact |
| MFA coverage check (per user) | Table stakes; drives immediate action |
| Compliance findings dashboard (6 deterministic rules) | Thay risk scoring; zero false positives; map thẳng SOC 2 controls |
| Real-time alerts (new OAuth app, incomplete offboarding, admin MFA off) | Không có alerts = customers không bao giờ đăng nhập lại |
| Exportable audit trail (PDF + CSV) | Blocks enterprise sales nếu không có |
| Multi-tenant data model (PostgreSQL RLS, workspace_id) | Must build from day 1; retrofitting = 6–8 tuần |

### 4.2 Should-Have v1 (Include Nếu Capacity)

| Feature | Notes |
|---|---|
| RBAC engine (DB-native PostgreSQL RLS + OPA cho policy eval) | Platform integrity |
| JIT access (Google Workspace groups, admin-initiated, 1h/4h/8h/24h) | Sprint 7 if capacity; v2 first item nếu không |
| Access review workflows (manager attestation, evidence export) | SOC 2 Type II compliance |
| Compliance reports (SOC 2, ISO 27001, GDPR one-click export) | Revenue driver Growth/Enterprise tier |
| Mobile app (JIT approve + push notifications + read-only dashboard) | Zero competitor; đã funded trong sprint plan |
| Onboarding templates by role | Ngăn access creep từ ngày đầu |
| Dependency mapping (blast radius per user) | Safety check trước automated actions |

### 4.3 Deferred to v2 (Explicitly Out of v1)

| Feature | Lý Do | Effort |
|---|---|---|
| AWS IAM automated deprovisioning | TA: 4–6 tuần một mình; multi-resource; non-atomic | 4–6 tuần |
| Slack automated deprovisioning | Cần Enterprise Grid ($12.50+/user/mo); ~60% SMB không có | 1–2 tuần sau qualification |
| GitHub SCIM deprovisioning | Cần GitHub Enterprise Cloud ($21/user/mo) | 1 tuần |
| HRIS integration | 60% SMB không có HRIS chuẩn; dùng manual trigger trước | 6–10 tuần |
| Periodic access reviews (campaign engine) | SOC 2 Type II, not Type 1; sau khi offboarding stable | 2–3 tuần |
| Custom workflow builder | 80% customers dùng defaults; platform play | 12–16 tuần |

### 4.4 Removed Entirely (Không Làm)

| Feature | Lý Do |
|---|---|
| Auto provisioning (SCIM inbound) | SMBs disable ngay ngày đầu; over-provisioning risk |
| RBAC role suggestions / role mining | PM: never acted on; tạo noise; damages trust |
| Automated shadow IT revocation (không confirm) | Blast radius risk = existential churn |
| Risk score composite (0–100) | Replaced by 6 compliance findings (deterministic, 0% false positive) |
| ML-based behavioral analytics | 60+ ngày cold start; no day-1 value; v3 |

### 4.5 Compliance Findings (Thay Thế Risk Scoring)

| Finding | Rule | SOC 2 | False Positive |
|---|---|---|---|
| User không có MFA | `mfa_enabled = false` | CC6.1 | 0% |
| Admin inactive >90 ngày | `last_login < now()-90d AND role = admin` | CC6.2 | <1% |
| OAuth grant của blocked app vẫn active | `app.status = BLOCKED AND grant.active = true` | CC6.3 | 0% |
| Suspended user còn OAuth grants | `user.suspended = true AND grant.count > 0` | CC6.8 | 0% |
| Admin không có onboarding record | `role = admin AND onboarding_record = null` | CC6.1 | ~2% |
| Offboarding incomplete >24h | `status = IN_PROGRESS AND started > 24h ago` | ISO A.7.3 | 0% |

---

## 5. Quyết Định: Technical Architecture

### 5.1 Integration Priority

| Provider | v1 Scope | API Method | Lý Do |
|---|---|---|---|
| **Google Workspace** (P0) | Full automation | Admin SDK: `users.update(suspended:true)` + `tokens.delete` | 40% SMB market; cleanest API; <3s latency |
| **Microsoft 365** (P0) | Full automation | Graph: `PATCH accountEnabled:false` + `revokeSignInSessions` | 50% SMB market; token propagation lag cần disclose |
| **GitHub** free org (P1) | Automated org removal | `DELETE /orgs/{org}/members/{username}` — 1 API call | Trivial; không cần SCIM |
| **Slack** free/Pro (P1) | Discovery + manual checklist | No deactivation API on free tier | Disclose clearly; upsell to Business+ |
| **Slack** Business+/Grid (P1.5) | Conditional automation | SCIM hoặc admin API | ~40% Slack users qualify |
| **AWS IAM** (P2) | Discovery only + checklist | CloudTrail read + manual link | 4–6 tuần complexity; explicit v2 |

### 5.2 Build vs. Buy

| Component | Quyết Định | Lý Do |
|---|---|---|
| Auth (platform SSO) | **Buy: Keycloak** | Never build OAuth/OIDC |
| RBAC platform-level | **Build: PostgreSQL RLS** | `workspace_id` + role; no runtime dependency |
| Policy evaluation (JIT + allow-list) | **Build: OPA on ECS** | Policy-as-code; auditable; versioned |
| Offboarding workflow | **Buy: AWS Step Functions** | Durable execution; compensating transactions; zero ops vs Temporal |
| Background jobs | **Buy: SQS + ECS workers** | Dead-letter queues; at-least-once delivery |
| PDF reports | **Build: Puppeteer** | Full template control |
| Access graph | **Build: PostgreSQL CTEs** | Sufficient đến 50K nodes |
| Secrets | **Buy: AWS Secrets Manager** | KMS; audit; không bao giờ .env |
| Notifications | **Build thin: SES + Slack Webhooks** | Đủ cho v1 |

### 5.3 Technical Non-Negotiables (Phải Đúng Từ Day 1)

1. **Multi-tenant isolation** — `workspace_id` mọi table + PostgreSQL RLS. Retrofitting = 6–8 tuần.
2. **Incremental sync với delta tokens** — `$deltaLink` (M365) + `nextPageToken` (Google). Full sync ở 100 tenants = 200K API calls. Mandatory.
3. **SOC 2 audit log schema trước Sprint 2** — Retrofitting evidence format cho auditors rất tốn kém.
4. **Webhook renewal automation** — Google Push expire 7 ngày; Graph expire 3 ngày. Silent expiry = alerting down.
5. **M365 token propagation caveat** — Access tokens valid đến 60 phút sau revoke. SLA language phải disclose rõ.
6. **Service account detection trước offboarding** — Undetected service account = production outage.
7. **Google OAuth verification Week 1** — 4–6 tuần lead time. Consent screen cảnh báo kill SMB conversions.

---

## 6. Quyết Định: Delivery Approach

### 6.1 Phased Timeline (3 Phases — 26 Weeks)

| Phase | Tuần | Mục Tiêu | Key Deliverable |
|---|---|---|---|
| **Phase 1: Visibility** | W1–8 | "Bạn có thể thấy mọi thứ" | Google + M365 users/apps visible; Shadow IT alerts live |
| **Phase 2: Control** | W9–18 | "Bạn có thể kiểm soát mọi thứ" | Automated offboarding G+M365; JIT access; Playbooks |
| **Phase 3: Trust** | W19–26 | "Bạn có thể chứng minh mọi thứ" | Compliance reports; Pen test pass; 5+ pilot customers |

**Revenue milestone: Sprint 6 (W11-12)** — Đây là lần đầu tiên có thể demo + charge.

### 6.2 Trust Infrastructure (Critical Path — Không Phải Feature)

Những thứ này là prerequisites cho first paying customer. Phải chạy song song từ Week 1:

| Item | Timeline | Chi Phí | Blocking? |
|---|---|---|---|
| Google OAuth app verification | Start W1; kết quả W5–7 | Free | **Yes** — demo |
| Data Processing Agreement (DPA) | W1–4 với legal | $5–10K | **Yes** — GDPR |
| SOC 2 Type 1: observation period | Bắt đầu W8 | $15–40K audit | **Yes** — enterprise customer |
| Penetration test | Engage W16; results W24 | $10–20K | **Yes** — 70% SMB yêu cầu |
| Cyber liability insurance | 2–3 tuần | $5–10K/yr | No |

### 6.3 Common Failure Modes (Phải Tránh)

| Failure | Frequency | Mitigation |
|---|---|---|
| Integration breadth over depth | Very High | Google + M365 depth-first; không sang provider 3 cho đến khi cả hai solid |
| Auto-revocation trước khi có trust | Very High | Human confirmation cho MỌI write operation trong v1 |
| AWS deprovisioning trong v1 | High | Không bao giờ. Explicitly v2. |
| HRIS dependency | High | Manual trigger only; HRIS sau khi biết customers dùng system gì |
| SOC 2 start quá muộn | High | Start W1; bị skip = no paying customer đến Month 8 |

---

## 7. Open Questions (Cần Quyết Định)

| # | Question | Owner | Blocking? | Deadline |
|---|---|---|---|---|
| 1 | **Offboarding marketing language:** "Automated across all providers" vs "Google + M365 automated <5 min, full audit trail for others"? | PO + PM | **Yes** — trước demo | Sprint 4 end |
| 2 | **Pilot customer qualification:** First customers waive SOC 2 requirement hay design partners? 3 LOIs cần bởi Sprint 8. | PM | **Yes** | Sprint 8 |
| 3 | **Google OAuth verification:** Block Sprint 2 production testing không? Build against test tenants cho đến khi verified? | TA + PM | **Yes — Sprint 2** | Week 1 |
| 4 | **SOC 2 budget:** $15–40K Type 1 audit chưa được confirm. | PM | **Yes** | Week 1 |
| 5 | **Slack customer base:** % pilot customers có Slack Business+/Grid? Determines v1.5 vs cut entirely. | PM | No | Sprint 4 |
| 6 | **M365 token propagation SLA:** Disclosed [X] min là bao nhiêu? Azure AD P1 required cho 15-min tokens? | TA + PO | No | Sprint 6 |
| 7 | **JIT Access timing:** Sprint 7 hay v2 first item? Decision gate = Sprint 5 utilization review. | PM | No | Sprint 5 end |

---

## 8. Kết Quả Session

### Quyết Định Đã Được Đảo Ngược So Với Plan Cũ

| Plan Cũ (Track 1) | Quyết Định Mới | Lý Do Thay Đổi |
|---|---|---|
| OPA cho RBAC engine | DB-native PostgreSQL RLS (platform) + OPA (policy eval only) | OPA overkill cho platform RBAC; chỉ cần cho JIT + allow-list logic |
| 4 providers automated offboarding trong Sprint 6 | Google + M365 only automated; checklist cho phần còn lại | AWS = 4–6 tuần một mình; Slack = Enterprise Grid required |
| Risk scoring (0–100) trong dashboard | 6 deterministic compliance findings | Risk scores → false positive hell trong 2 tuần đầu |
| Vanta/Drata là primary competitor | Lumos là primary competitor | Khác category hoàn toàn |
| JIT access Sprint 5 | JIT access Sprint 7 (nếu capacity) hoặc v2 | RBAC + offboarding phải stable trước |
| AWS trong v1 offboarding | AWS discovery only; deprovisioning v2 | Non-atomic; 4–6 tuần effort |
| Automated shadow IT revocation | Human-confirmed allow-list management | Blast radius risk = instant churn |

### Success Metrics (Customer-Facing)

| Metric | Target | Benchmark |
|---|---|---|
| Time to offboard | <5 min (Google + M365) | Industry avg: 3–5 days |
| Shadow IT discovery rate | >95% OAuth apps | Nudge Security benchmark |
| Offboarding coverage | >99% (automated + checklist) | Zero orphaned accounts |
| Access review completion | >90% (vs 40% manual) | SOC 2 compliance |
| MFA coverage post-onboarding | 100% | From baseline ~40–60% |
| Time-to-first-value | <1 hour from signup | Lumos benchmark ~4 hours |
| Audit evidence prep time | <2 hours (vs 40–80 hours) | SOC 2 Type I audit |

---

*Session: 2026-05-28 | Method: 3 agents × 2 rounds | Files: [02-feature-research-synthesis.md](02-feature-research-synthesis.md) | Next step: SPEC*
