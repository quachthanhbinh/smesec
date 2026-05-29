# SMESec — Feasibility Assessment & Remediation Plan

**Ngày:** 2026-05-29  
**Phiên bản:** 1.1  
**Trạng thái:** Working Document  
**Nguồn:** Tổng hợp đánh giá 3-agent (Product Owner · Technical Advisor · Project Manager)  
**Liên quan:** [01-system-architecture.md](01-system-architecture.md) · [02-design-document.md](02-design-document.md) · [04-delivery-plan-original.md](04-delivery-plan-original.md) · [06-delivery-plan-adjusted-2x.md](06-delivery-plan-adjusted-2x.md) · [07-delivery-plan-realistic-hiring.md](07-delivery-plan-realistic-hiring.md)

---

## ⚠️ Timeline Context

Risk assessments reference the **original 12-month plan**. For adjusted timelines:
- **[2x Adjusted Plan](06-delivery-plan-adjusted-2x.md)** — Multiply week/month references by ~2x
- **[Realistic Hiring Plan](07-delivery-plan-realistic-hiring.md)** — See specific hiring timeline in that document

**Key difference:** Realistic Hiring Plan has ML Engineer #1 joining Month 8 (not Day 1), significantly impacting Track 2 risks.

---

## Mục Lục

1. [Verdict Tổng Quan](#1-verdict-tổng-quan)
2. [Rủi Ro CRITICAL — Có thể kill sản phẩm](#2-rủi-ro-critical--có-thể-kill-sản-phẩm)
3. [Rủi Ro HIGH — Có thể derail kế hoạch](#3-rủi-ro-high--có-thể-derail-kế-hoạch)
4. [Vùng Sương Mù — Blind Spots Chưa Được Tính Tới](#4-vùng-sương-mù--blind-spots-chưa-được-tính-tới)
5. [Giải Pháp Theo Từng Rủi Ro](#5-giải-pháp-theo-từng-rủi-ro)
6. [Revised Sprint Plan](#6-revised-sprint-plan)
7. [Missing Components Checklist Trước Launch](#7-missing-components-checklist-trước-launch)
8. [Top 5 Quyết Định Tuần 1](#8-top-5-quyết-định-tuần-1)

---

## 1. Verdict Tổng Quan

**Confidence tổng hợp:** PO: 6/10 · TA: 8/10 · PM: 5.5/10

| Track | Khả thi? | Điều kiện |
|---|---|---|
| **Track 1 MVP (T3)** | ✅ Có, với thu hẹp scope | Chỉ Google WS + M365; defer Slack/AWS IAM sang T4 |
| **Track 1 v1 (T6)** | ⚠️ Có nhưng thực tế W30–32 | Phải giải quyết 5 blockers kỹ thuật + legal trước |
| **Track 2 AI (T9+)** | ⚠️ Có điều kiện | Phụ thuộc accuracy gate; không được overmarket trước khi qua gate |
| **SOC 2 Type 2 + ISO 27001 (T12)** | ❌ Rủi ro cao | Nếu Type 1 audit trễ quá W30, observation window 6 tháng không đủ |

**Xác suất thực tế (PM estimate):**
- MVP tại W12: **35–45%** (W14–16 là kết quả khả dĩ hơn)
- v1 tại W26 với 5+ customers: **25–35%** (W30–32 là kết quả khả dĩ hơn)
- v1.5 tại W38: **40–50%**
- v2 (SOC 2 Type 2 + ISO 27001) tại W52: **30–40%**

---

## 2. Rủi Ro CRITICAL — Có thể kill sản phẩm

### R-C1: False Positive Automated Offboarding → Legal Liability

**Nguồn:** Technical Advisor + Product Owner  
**Mô tả:** Hệ thống tự động revoke access của nhân viên trong <5 phút khi HR trigger. Không có manual approval gate, không có rollback workflow, không có grace period.

**Kịch bản thực tế:**
- HR sync error → CEO mất access trong board meeting → kiện
- Nhân viên bị offboard nhầm → access bị revoke trước khi check → compliance incident

**Thiếu hoàn toàn trong design hiện tại:**
- Rollback/re-enable workflow
- Grace period (e.g., 30 phút) với override capability
- Manual approval gate (optional per policy)
- ToS liability cap cho automated actions
- Audit trail liên kết revocation với HR event cụ thể

---

### R-C2: Google Workspace API Rate Limit — Sẽ breach ở scale

**Nguồn:** Technical Advisor  
**Mô tả:** Google Admin SDK quota: **1,500 requests/100s per project** (không phải per tenant — toàn bộ SMESec GCP project share quota này).

```
50 tenants × 15-min sync × ~30 API calls/sync = 1,500 calls/15 phút
```

Đúng bằng ceiling, zero margin. Ở 100 tenants: **vượt liên tục.**

**Hậu quả:** Sync fails → stale inventory → offboarding trigger không fire → ex-employee vẫn có access. Đây là lý do cốt lõi khách hàng mua sản phẩm.

---

### R-C3: M365 Webhook Subscription Expiry 3 Ngày — Silent Failure

**Nguồn:** Technical Advisor + Project Manager  
**Mô tả:** Microsoft Graph webhook subscriptions expire mỗi **3 ngày** và phải renew. Nếu renewal job fail → toàn bộ M365 change events ngừng mà **không có error, không có alert, chỉ im lặng.**

**Hậu quả:** Inventory đóng băng trong khi IT admin nghĩ data đang fresh. Offboarding triggers miss M365 deprovisioning → SOC 2 evidence gaps.

**Đây là known Microsoft Graph gotcha** mà mọi sản phẩm dùng delta sync đều gặp. Phải thiết kế ở S1 (infra), không phải phát hiện ở S3.

---

### R-C4: GDPR Article 17 vs. S3 Object Lock — Mâu Thuẫn Kiến Trúc

**Nguồn:** Technical Advisor + Product Owner  
**Mô tả:** S3 Object Lock WORM = không ai xóa được trong 7 năm. GDPR Article 17 = quyền erasure. Hai yêu cầu này trực tiếp mâu thuẫn.

Legal workaround (encrypt data + destroy key → data becomes unreadable = effective erasure) phải được thiết kế vào kiến trúc **trước khi onboard customer EU đầu tiên.** Hiện tại không có mention nào trong design.

---

### R-C5: Không có Customer Acquisition Plan — CAC Model Rỗng

**Nguồn:** Product Owner  
**Mô tả:** Ở $399/mo (LTV ~$7,500 tại 24 tháng): CAC tối đa $2,500 để viable.

| Channel | Estimated CAC | Viable? |
|---|---|---|
| Outbound SDR | $3,000–8,000 | ❌ Không |
| Google/LinkedIn ads | $800–2,000 | ⚠️ Barely, zero margin |
| MSP channel | $500–1,500 | ✅ Viable |
| Freemium/PLG | $50–300 | ✅ Viable |

**Hiện trạng:** Team không có sales/BD resource. PM là 0.5 FTE. Không có pipeline, không có owner cho 5 pilot customers. Margin 91% là vô nghĩa nếu không thể acquire customers profitably.

---

### R-C6: Keycloak Self-Hosted — Single Point of Failure Authentication

**Nguồn:** Technical Advisor  
**Mô tả:** Toàn bộ auth đi qua Keycloak ECS. Nếu Keycloak crash → không ai login được, kể cả IT admin đang cần revoke access khẩn cấp.

**Thiếu:**
- JWKS caching cho JWT validation độc lập với Keycloak uptime
- HA config (multiple task instances, không chỉ multi-AZ)
- RTO/RPO không được định nghĩa
- DevSecOps contract = không ai patch quarterly CVEs

---

## 3. Rủi Ro HIGH — Có thể derail kế hoạch

### R-H1: S1 Sprint Over-Scoped 230%

**Nguồn:** Project Manager  
**Mô tả:** AWS infra + Keycloak OIDC + multi-tenant schema + CI/CD trong 2 tuần với 4 người.

```
Required: 50–65 dev-days
Available: 4 FTE × 7 effective days = 28 dev-days
Utilization: 179–232%
```

**Kết quả thực tế:** MVP trễ W14–W16 thay vì W12. Toàn bộ Phase 2 cascade theo.

---

### R-H2: Chrome MV3 Service Worker Kill mid-WASM-scan

**Nguồn:** Technical Advisor  
**Mô tả:** Chrome giết service worker sau ~5 phút idle. Nếu WASM BERT-tiny inference đang chạy bị kill → scan interrupted. Fail-closed behavior khi interrupted không được spec rõ — nếu default là "allow" thì đây là bypass vector hoàn chỉnh.

Thêm: First-scan latency sau idle = **2–4 giây** (service worker cold start + WASM compilation), không phải 50–80ms như spec — chỉ đúng sau warm-up.

---

### R-H3: PgBouncer Không Được Nhắc Đến

**Nguồn:** Technical Advisor  
**Math:**
```
100 tenants × 10 concurrent req × 10 ECS tasks × 4 pg connections = 40,000 connections
RDS PostgreSQL db.r6g.2xlarge max connections: ~3,200
Oversubscription: 12.5×
```

PgBouncer (transaction-pooling mode) hoặc RDS Proxy là **bắt buộc** để đạt 50+ tenants.

---

### R-H4: EventBridge At-Least-Once → Duplicate Offboarding

**Nguồn:** Technical Advisor  
**Mô tả:** EventBridge guarantee là at-least-once. Duplicate `OffboardingTriggered` → 2 offboarding workflows chạy concurrent cho cùng user → duplicate access revocation entries trong audit log, potential customer alert noise.

**Fix đơn giản:** `StartExecution` với deterministic `name` parameter (derived từ event correlation ID). Không được mention.

---

### R-H5: WASM Model Update Mechanism Không Tồn Tại

**Nguồn:** Technical Advisor  
**Mô tả:** ONNX model BERT-tiny được nhúng vào extension. Khi có model mới:
- Chrome Web Store review: 2–7 ngày để ship update
- Model trong production lag hàng tuần so với server
- Không có hot-swap mechanism, không có model versioning
- Không có signature verification khi download model từ S3 → **supply chain attack vector**

---

### R-H6: Shadow AI Risk Score Formula — Gameable + Wrong Direction

**Nguồn:** Technical Advisor  
**Vấn đề:**
1. Weights (30%/20%/15%...) chọn theo intuition, chưa validated empirically
2. `vendor_DPA_available` (20%) + `security_certifications` (15%) = self-reported bởi vendors → predatory vendor có thể publish DPA + mua ISO cert để score tốt
3. `app_age_days` direction **sai**: app mới = rủi ro cao hơn, phải tăng score, không giảm

---

### R-H7: S1 Không Bao Gồm Test Tenant Provisioning

**Nguồn:** Project Manager  
**Mô tả:** CI gate tại S6 yêu cầu timed offboarding test (<5 phút) end-to-end. Test này cần mock hoặc dedicated test tenants cho cả 4 providers. Nếu dùng real test tenants (preferred), phải provision từ S1 (không có trong scope S1).

---

### R-H8: M4 Onboarding Shock — 3 New Hires Đồng Thời

**Nguồn:** Project Manager  
**Mô tả:** ML Eng + BE#3 + FE#2 đều join M4 trong cùng sprint S7. Tech Lead mất 2–3 ngày per person cho onboarding = **25–40% capacity loss** trong S7, trong khi S7 đã scoped JIT access (feature phức tạp).

---

### R-H9: Pentest Timeline Không Đủ

**Nguồn:** Project Manager  
**Math:**
```
Scoping + kickoff: 1 tuần
Active testing (multi-tenant SaaS): 2+ tuần
Report delivery: 1 tuần
Remediation: 1–3 tuần
Retest: 1 tuần
Minimum: 6–8 tuần
```

Hiện tại plan: W21 → W26 = **5 tuần**. Không đủ nếu có critical finding (và gần như chắc chắn sẽ có với multi-tenant SaaS).

---

### R-H10: Chrome Web Store Review — Ngoài tầm kiểm soát

**Nguồn:** Project Manager  
**Mô tả:** Security extension với `tabs`, `webRequest`, `scripting`, `host permissions` và LLM DLP content scanning sẽ kích hoạt manual security review của Google. Timeline: **2–6 tuần**. Submit S12 (W24), cần approval W26 = **window 2 tuần** = xác suất miss ~50%.

---

## 4. Vùng Sương Mù — Blind Spots Chưa Được Tính Tới

### BS-1: Alert Fatigue (Critical UX blind spot)

Ngày 2, IT admin mở dashboard thấy 847 shadow IT flags. Reaction: panic → gọi support, hoặc dismiss all → không mở lại. Không có trong design:
- Calibration period (14 ngày đầu = learning mode, không alert)
- Smart default suppression rules
- Phân biệt "act today" vs "review monthly"
- Noise baseline per tenant

Đây là **#1 lý do security tools bị ignore** trong 30 ngày đầu. Module C (Shadow AI) fires từ Day 1 — cần thiết kế alert strategy trước khi MVP.

---

### BS-2: Module F (Employee Privacy Dashboard) — GDPR + Labor Law Trap

Module F transparent dashboard cho employees về monitoring data tạo:
1. **GDPR Article 13/14:** Explicitly acknowledges SMESec là data processor. Customer phải có documented lawful basis. Hầu hết SME không có.
2. **Employment law variance:** France (rất restrictive), Germany (works council approval required), UK, US (state by state) — không thể có one-size-fits-all design.
3. **HR liability trigger:** Employee thấy DLP log → biết bị monitor → không có trong employment contract → SMESec tạo legal exposure cho customer.

**Đề xuất:** Remove Module F khỏi v1 hoàn toàn. Admin-only transparency report thay thế. Reintroduce ở Enterprise tier sau legal review per geography.

---

### BS-3: Switching Cost Gần Bằng 0

Customer có thể connect Google Workspace OAuth trong 30 phút với competitor và rời đi. Không có gì giữ lại ngoài custom playbooks đã tuned (switching cost duy nhất đáng kể trong design hiện tại).

Playbook builder là feature tạo switching cost cao nhất → phải vào MVP scope và được marketing là "your company's security procedures, automated," không phải "access management tool."

---

### BS-4: Prompt Injection BERT Target (TPR >85%, FPR <2%) — Không Thực Tế Trong 6 Tháng

Cần: ~50,000+ labeled examples đa dạng (GPT-4/Claude/Gemini/open-source), adversarial test set, fine-tuning infrastructure, production feedback loop. Dataset này không tồn tại publicly ở đủ quality.

**Practical alternative:** Dùng Lakera Guard (đã cited ở Tier 3) cho prompt injection, bỏ internal BERT model target ở v1.

---

### BS-5: Deepfake Detection Liability — Không Có Insurance Coverage

Không có cyber insurance nào cover AI-assisted false positive liability trong security tooling (market gap 2026).

Kịch bản: Module D flag legitimate CEO video call là deepfake → IT admin không proceed → deal falls through → CEO sues → demand letter > annual revenue.

**Required trước EU/US launch:**
- E&O insurance covering AI decision outputs
- UI language: "advisory only — human verification required" (không phải "deepfake detected")
- ToS liability cap tại 1 month subscription
- Customer acknowledgment: all AI alerts require human verification before action

---

### BS-6: New AI Tool Recognition Lag

Khi GPT-5, Claude-next, hoặc enterprise AI tool mới ra mắt, extension không recognize form fields / URL patterns. Server-push rule update cần bypass Chrome Web Store review (declarativeNetRequest dynamic rules trong MV3 có giới hạn). Trong window đó DLP không hoạt động với tool mới nhất.

Extension là **Chrome-only DLP cho AI tools đã biết** — không phải universal DLP layer. Cần honest marketing và expectation setting.

---

### BS-7: Compliance Report Quality ≠ Audit-Ready Evidence

Growth tier customer ($799/mo) mua SMESec để pass ISO 27001. Auditor reject evidence format (SMESec chưa có auditor relationships như Vanta). Customer churn với chargeback demand.

**Disclaimer bắt buộc trong sales materials:** "compliance preparation assistance — not audit-ready evidence without auditor validation."

---

### BS-8: EU Legal Opinion → Track 2 Architecture Risk

EU legal opinion về voice deepfake = biometric (Article 9) commissioned Day 1. Nếu opinion đến W6–W8 với kết quả "yes = biometric" → Track 2 deepfake detection phải redesign. Không có sprint capacity reserve cho kịch bản này trong plan hiện tại.

---

### BS-9: Không Có QA Role Trong Phase 1

Phase 1 không có dedicated QA engineer. Không rõ:
- Ai viết integration tests cho Google/M365/Slack API mocks?
- Ai owns test coverage standards?
- Ai validates 2 playbooks end-to-end trước S6?

Thực tế: engineers test code của chính mình = acceptable cho unit tests, không đủ cho multi-tenant offboarding system integration test.

---

### BS-10: Slack Tier Constraint Không Được Acknowledge

Slack Admin API (user management + offboarding) yêu cầu **Business+ tier** ($12.50/user/month). ~80% SME Slack users dùng Free/Pro tier → Slack offboarding silently fails cho 80% target market mà không có error hay notification. UI phải detect tier và communicate limitation rõ ràng.

---

## 5. Giải Pháp Theo Từng Rủi Ro

### Sol-C1: Automated Offboarding với Safety Net (R-C1)

**Thiết kế lại offboarding flow:**

```
Offboarding Request (HR trigger / manual)
    │
    ▼
[GRACE PERIOD: 30 min — configurable per org, 0 for emergency]
    │  Admin nhận alert: "Access revocation scheduled for [user] at [time]"
    │  One-click CANCEL available via Slack/email
    │
    ├── No cancel → Proceed to execution
    │
    ▼
[EXECUTION: Step Functions workflow]
    │  Revoke Google WS → M365 → Slack → AWS IAM (parallel, retry x3)
    │  Log each step với correlation_id + hr_event_id
    │
    ▼
[ROLLBACK CAPABILITY: 24h window]
    Admin có thể trigger "Re-enable access" → reverse workflow
    Re-enable logged separately cho audit trail

ToS update: "Automated actions are advisory executions based on HR system signals.
SMESec is not liable for access changes resulting from incorrect HR data."
```

**Sprint:** Add offboarding rollback workflow ở S6 (không defer). Add grace period config ở S7.

---

### Sol-C2: Google API Rate Limit Management (R-C2)

**Per-tenant quota distribution + adaptive sync:**

```go
// Sync scheduler: distribute tenants across 15-min window
type SyncScheduler struct {
    tenants       []TenantID
    windowSeconds int  // 900 seconds = 15 min
    apiCallBudget int  // 1400 calls/15min (safety margin below 1500)
}

// Spread tenants: tenant[0] syncs at t=0, tenant[1] at t=18s, etc.
// Each tenant gets apiCallBudget/len(tenants) calls per window

// Retry policy: exponential backoff on 429 with jitter
// Quota monitoring: alert when aggregate usage > 80% of quota
// Fallback: degrade to 30-min sync when >70 tenants active
```

**Config:** Separate GCP service account per tenant cluster (20 tenants/project) → multiply quota by N projects. Cost: $0 (free GCP service accounts).

**Sprint:** Design ở S1, implement ở S2 (Google sync). CI test: simulate 100 tenant quota scenario.

---

### Sol-C3: M365 Webhook Renewal Service (R-C3)

**Renewal service phải là S1 infra deliverable, không phải S3 feature:**

```
Architecture:
  SubscriptionRegistry (RDS):
    - subscription_id, tenant_id, resource_type, expiry_timestamp, status

  RenewalJob (EventBridge Scheduler, runs every 12 hours):
    - Query subscriptions expiring in next 24 hours
    - Renew via Graph API: PATCH /subscriptions/{id}
    - On 410 Gone → trigger full delta sync for that tenant
    - On renewal failure → DLQ → alert + fallback to polling mode
    - Update expiry_timestamp on success

  UI indicator:
    - "Last synced: X minutes ago" per tenant
    - Amber warning if last sync > 20 minutes
    - Red alert if last sync > 60 minutes
```

**Sprint:** Add to S1 infra scope (schema + scheduler setup). Implement renewal logic ở S3.

---

### Sol-C4: GDPR Erasure + S3 Object Lock (R-C4)

**Envelope encryption với key destruction:**

```
Architecture:
  Per-tenant audit log encryption:
    - Each tenant has a dedicated KMS key (data key, wrapped by CMK)
    - Audit log entries encrypted with tenant data key before writing to S3
    - S3 Object Lock prevents deletion of ciphertext (compliance safe)

  Erasure workflow (GDPR Art. 17):
    - Customer submits erasure request → /api/v1/gdpr/erasure
    - SMESec schedules KMS key deletion (7-day pending window, KMS minimum)
    - After key deletion: ciphertext exists but is permanently unreadable
    - This satisfies GDPR "effective erasure" standard (ICO guidance, EDPB guidance)
    - Erasure certificate issued with key deletion timestamp

  Legal documentation:
    - DPA template explicitly states: audit logs stored in encrypted WORM storage;
      erasure performed via key destruction per EDPB Recommendation 01/2020
```

**Sprint:** Design ở S1 (key management architecture). Implement erasure endpoint ở S11 (GDPR automation sprint).

---

### Sol-C5: Customer Acquisition Plan (R-C5)

**Primary motion: MSP Channel + Freemium health check**

```
Channel 1 — MSP Partner Program:
  Target: 10 MSPs serving SMEs in US/EU
  Product requirement: MSP portal (multi-tenant dashboard, white-label option)
  Economics: MSP marks up 20-30%; SMESec gets $280-320/mo per customer
  CAC via MSP: ~$500-800 (MSP handles customer relationship)
  Requirement: MSP portal as S14-S15 deliverable (Phase 3)

Channel 2 — Freemium "Security Health Check":
  Free tier: Asset inventory + shadow IT scan for ≤25 users, 14 days
  Gate to paid: offboarding automation + compliance reporting
  Conversion: user sees risk report → needs remediation → upgrade
  CAC target: <$300 via self-serve
  Requirement: Free tier billing config ở S13

Pilot Customer Pipeline (immediate):
  Owner: Assign full-time BD consultant from Week 1 (3 days/week, $60-80/hr contract)
  Funnel: 100 outreach → 30 qualified → 15 demo → 5 pilot agreement
  Timeline: Pipeline starts W1, close first 2 at W16, 5 at W26
  Qualification criteria: Google WS or M365 admin, 25-200 employees, active compliance pressure
```

---

### Sol-C6: Keycloak Resilience (R-C6)

**JWT validation independence:**

```go
// Each service caches JWKS locally with 6-hour TTL
// JWT validation does NOT call Keycloak at runtime
type JWTValidator struct {
    jwksCache    *jwk.Cache  // auto-refresh every 6h, serve stale on failure
    keycloakURL  string
}

// If Keycloak is down:
// - New logins fail (expected)
// - Existing valid JWTs (15-min TTL) continue to work for their remaining lifetime
// - Services remain fully functional for up to 15 minutes without Keycloak

// Keycloak HA config:
// - Minimum 2 ECS tasks (not just multi-AZ placement)
// - Separate RDS instance (not shared with app DB)
// - Health check: /health/ready endpoint, 30s interval, 3 failure threshold
// - RTO target: <2 min (ECS task replacement)
```

**Alternative for v1:** Evaluate WorkOS or Auth0 as managed alternative. Cost: ~$500-1,000/mo at 50 customers (vs Keycloak $50/mo but requires dedicated ops). Trade-off: operational simplicity vs cost. Revisit at v2.

---

### Sol-H1: Split S1 (R-H1)

**Revised sprint structure:**

```
S1a (W1–W2): Infrastructure Foundation
  ✅ AWS VPC + ECS Fargate skeleton + RDS multi-tenant schema
  ✅ CI/CD pipeline (GitHub Actions + Terraform)
  ✅ M365 webhook renewal service schema + EventBridge scheduler setup
  ✅ Test tenant provisioning (Google WS + M365 dev accounts)
  ✅ PgBouncer / RDS Proxy configuration
  ❌ Defer: Keycloak full config (JWT custom claims, SAML)

S1b (W3–W4): Auth + Security Layer
  ✅ Keycloak OIDC setup + MFA TOTP + tenant provisioning flow
  ✅ JWT RS256 middleware + JWKS caching
  ✅ RLS policies + tenant isolation CI test (MUST PASS to merge)
  ✅ WAF + GuardDuty + Security Hub baseline
  ✅ W4 usability test: Google Workspace OAuth wizard (non-technical user)

Impact: MVP moves to W14. All downstream milestones shift +2 weeks.
Benefit: Foundation is stable. No architectural surprises in S2+.
```

---

### Sol-H3: PgBouncer (R-H3)

Add PgBouncer sidecar to all ECS service task definitions hoặc dùng RDS Proxy (managed, simpler):

```hcl
# RDS Proxy (recommended for ECS Fargate)
resource "aws_db_proxy" "smesec_proxy" {
  name                   = "smesec-${var.env}"
  engine_family          = "POSTGRESQL"
  role_arn               = aws_iam_role.rds_proxy.arn
  vpc_subnet_ids         = var.private_subnet_ids
  require_tls            = true

  auth {
    auth_scheme = "SECRETS"
    secret_arn  = aws_secretsmanager_secret.db_credentials.arn
    iam_auth    = "REQUIRED"
  }
  
  # Connection pooling handles the 40,000→3,200 oversubscription
}
```

**Sprint:** S1a deliverable.

---

### Sol-H4: EventBridge Idempotency (R-H4)

```go
// Step Functions StartExecution with deterministic name
func triggerOffboarding(event ThreatDetectionEvent) error {
    executionName := fmt.Sprintf("offboard-%s-%s",
        event.TenantID,
        event.CorrelationID,  // UUID from originating HR event
    )
    _, err := sfnClient.StartExecution(&sfn.StartExecutionInput{
        StateMachineArn: aws.String(offboardingStateMachineARN),
        Name:            aws.String(executionName),  // Duplicate = ExecutionAlreadyExists error, ignored
        Input:           aws.String(mustMarshal(event)),
    })
    if isAlreadyExists(err) {
        return nil  // Idempotent: already running, skip
    }
    return err
}
```

**Sprint:** S6 (offboarding workflow). Required for MVP.

---

### Sol-BS1: Alert Fatigue — Tiered Alert Strategy (BS-1)

```
Alert Strategy Design:

Phase 1 — Calibration (Days 1–14):
  - All detections logged internally, NO alerts shown
  - System learns baseline: which OAuth apps are "normal" for this tenant
  - Dashboard shows: "Calibrating... your security baseline is being established"

Phase 2 — Curated Alerts (Day 15+):
  Priority 1 (act today): New HIGH/CRITICAL OAuth apps added in last 24h
  Priority 2 (this week): Policy violations + over-provisioned access
  Priority 3 (monthly review): Low-risk shadow IT inventory

Default suppression rules:
  - Auto-whitelist all apps installed >90 days ago with >5 users
  - Auto-whitelist Microsoft-native + Google-native apps
  - Suppress duplicate alerts for same app/user within 7 days

Dashboard: "You have 3 things to review today" (not "847 alerts")
```

---

### Sol-BS4: Prompt Injection — Use Lakera Guard (BS-4)

Replace internal BERT model target với Lakera Guard API (already cited in Tier 3 DLP):

```
v1: Lakera Guard API (server-side, async) + OWASP regex ruleset (Tier 1, sync)
  - No internal ML model needed
  - Lakera maintains model updates
  - Cost: ~$0.001/request
  - Accuracy: production-validated

v2 (Enterprise, Sprint 23+): Fine-tuned BERT only if Lakera Guard
  - Cannot meet FPR <2% requirement (measure in production first)
  - Internal model only for tenants with enough volume for fine-tuning
  
Remove from plan: "BERT prompt injection at TPR >85%, FPR <2% in Sprint 23-24"
Replace with: "Lakera Guard integration at Sprint 8, internal model evaluation at Sprint 23"
```

---

## 6. Revised Sprint Plan

### Phase 1 Revised: Foundation → MVP (S1a–S6, Month 1–3.5)

| Sprint | Tuần | Deliverable | Gate |
|---|---|---|---|
| **S1a** | W1–2 | AWS infra (VPC/ECS/RDS), PgBouncer/RDS Proxy, CI/CD, M365 webhook renewal schema, test tenant provisioning | Tenant isolation CI test green; test tenants provisioned |
| **S1b** | W3–4 | Keycloak OIDC + JWT JWKS caching, RLS policies + CI test, WAF baseline. **W4 usability test: Google WS OAuth wizard** | Keycloak HA live; JWT independent of Keycloak; W4 test: <45 min setup |
| **S2** | W5–6 | Google Workspace sync (users, OAuth apps, shadow IT). Dashboard skeleton. **Setup guide written.** | First-value demo <30 min from OAuth grant |
| **S3** | W7–8 | M365 delta link + webhook renewal implementation + 410 Gone fallback. Unified dashboard. Staleness indicators. | M365 sync stable; webhook auto-renews without intervention |
| **S4** | W9–10 | Asset classification engine, OAuth scope risk scoring (formula corrected + calibration period), shadow IT alerts. Flutter mobile scaffold. | Calibration period fires; no alert before Day 15 |
| **S5** | W11–12 | Slack tier detection + graceful degradation. AWS IAM discovery. RBAC model + least-privilege. Identity graph. | Slack tier shown in UI; 4 providers unified |
| **S6** | W13–14 | **🏁 MVP**: Offboarding với grace period + rollback workflow + idempotency key. 2 playbooks. Immutable audit log (envelope encryption). Mobile beta → TestFlight W11. | Offboarding timed test <5 min in CI (với real test tenants); rollback verified |

**MVP = W14 (revised từ W12)**

---

### Phase 2 Revised: MVP → v1 (S7–S14, Month 4–7)

| Sprint | Tuần | Track 1 | Track 2 | Gate |
|---|---|---|---|---|
| **S7** | W15–16 | JIT access + access reviews. **Onboarding buffer: 3 new hires ramp.** | ML Eng onboards; Lakera Guard integration POC | Onboarding docs complete; Lakera Guard API connected |
| **S8** | W17–18 | Playbook engine (Step Functions), 3 playbooks. Alert fatigue strategy implemented. | LLM DLP browser ext v0.1 (Presidio local, Tier 1+2 only). **Submit stripped-down extension v0 to Chrome Web Store W18** | Calibration → curated alerts live; extension Tier 1 detects PII |
| **S9** | W19–20 | 5 playbooks complete, mobile push notifications, MSP portal foundation. | Shadow AI governance v1: risk score (formula v2, corrected), attestation workflow | Risk score recalc on incident event; Slack tier detection live |
| **S10** | W21–22 | ISO 27001 + SOC 2 compliance dashboard, Vanta integration. **Pentest begins W21.** | Deepfake defense POC (Hive API). `ThreatDetectionEvent` schema v1 **FROZEN** | Schema locked; pentest scope agreed |
| **S11** | W23–24 | Compliance reports (PDF export), GDPR erasure endpoint (envelope key deletion), GDPR Art. 17 flow. | T1-T2 integration: AI threat events → EventBridge → Step Functions. **Tech Lead full-time.** Fallback: manual trigger preserved | **Pentest remediation: Critical/High issues fixed** |
| **S12** | W25–26 | SaaS dependency map. Pentest retest. **Remove Module F from all non-Enterprise tiers.** | Full T1-T2 end-to-end test (automated). Extension full version submitted to Chrome Web Store. | 0 Critical/High pentest findings open |
| **S13** | W27–28 | Pilot feedback sprint (**reserved as buffer/feedback only, no new features**). Chrome Web Store approval tracking. | Chrome extension approval or contingency (web-based DLP fallback for blocked extension) | Extension approved OR fallback deployed |
| **S14** | W29–30 | **🏁 v1**: Production launch, 5+ pilot customers, SOC 2 Type 1 audit engagement signed | Track 2: Shadow AI + LLM DLP in v1 (Chrome) | No new features; hardening only |

**v1 = W30 (revised từ W26)**

---

### Phase 3 & 4 Adjusted Timeline

| Milestone | Tuần (revised) | Key Additions |
|---|---|---|
| **v1.5** | W42 (Month 10) | MSP portal GA, deepfake v2, AI phishing (M365 Defender), billing tiers enforced, 10+ paying customers |
| **v2** | W56 (Month 14) | SOC 2 Type 2 ✅ · ISO 27001 ✅ (nếu Type 1 signed W30, 6-month observation = W56) · BERT prompt injection (via Lakera baseline) · Enterprise tier |

> **Note:** v2 SOC 2 Type 2 completion phụ thuộc vào observation window bắt đầu từ v1 launch. Nếu v1 = W30, earliest completion = W56. Mọi slip thêm đẩy tiếp.

---

## 7. Missing Components Checklist Trước Launch

### Sprint 0 / Pre-Sprint (Trước khi viết dòng code đầu tiên)

- [ ] **Legal:** Engage EU/US employment + data protection counsel → scope Module B/D/F liability
- [ ] **Legal:** EU legal opinion: voice deepfake = GDPR Article 9 biometric? (Day 1)
- [ ] **Legal:** GDPR Article 17 vs S3 Object Lock resolution → envelope encryption design approved
- [ ] **Insurance:** E&O insurance covering AI decision outputs (Module D deepfake, Module B prompt injection)
- [ ] **Hiring:** ML Engineer JD posted (Day 1); go/no-go W6
- [ ] **BD:** BD consultant engaged (Week 1); pilot customer funnel started
- [ ] **Architecture:** PgBouncer / RDS Proxy decision documented in ADR
- [ ] **Architecture:** M365 webhook renewal service design reviewed
- [ ] **Architecture:** Offboarding grace period + rollback flow design reviewed
- [ ] **Architecture:** Envelope encryption per-tenant KMS key strategy reviewed

### Sprint 1 (S1a) — Infrastructure Must-Haves

- [ ] PgBouncer / RDS Proxy deployed (not deferred)
- [ ] M365 webhook renewal schema + EventBridge scheduler configured
- [ ] Test tenants provisioned: Google WS dev domain + M365 dev tenant + Slack dev workspace + AWS sandbox
- [ ] `tenant_id + data_residency` columns on ALL tables from day 1 (no exceptions)
- [ ] EU region (`eu-west-1`) infra configured in Terraform from S1 (irreversible decision)
- [ ] RLS CI test: creates 2 tenants, inserts for A, queries as B → must return 0 rows → blocks merge

### Sprint 1b — Auth Must-Haves

- [ ] Keycloak JWKS caching in all API services (JWT validation independent of Keycloak uptime)
- [ ] Keycloak: minimum 2 ECS task instances (HA, not just multi-AZ placement)
- [ ] RLS policy CI enforcement: new table without RLS policy = build failure
- [ ] SCA pipeline: `govulncheck` (Go) + `pip-audit` (Python) in CI, block on HIGH CVE

### Before v1 Launch (W30)

- [ ] Google rate limit: per-tenant quota distribution implemented + aggregate monitoring
- [ ] EventBridge idempotency keys on all Step Functions StartExecution calls
- [ ] Offboarding rollback workflow tested end-to-end
- [ ] GDPR erasure endpoint live + key deletion tested
- [ ] Slack tier detection in UI (clear message when Business+ not available)
- [ ] Sync staleness indicators in UI (amber >20 min, red >60 min)
- [ ] Module F removed from all non-Enterprise tiers
- [ ] Compliance report disclaimer: "preparation assistance, not audit-ready evidence"
- [ ] ToS: liability cap for automated actions + AI decision advisory language
- [ ] Deepfake UI: "advisory only — human verification required before acting"
- [ ] Alert calibration period (14-day learning mode) live
- [ ] Chrome extension: stripped-down v0 submitted W18 for early review buffer

---

## 8. Top 5 Quyết Định Tuần 1

| # | Quyết định | Hạn chót | Lý do bắt buộc ngay |
|---|---|---|---|
| 🔴 **1** | **Submit Google Workspace + Microsoft 365 verification** (Week -3, before project start) → OAuth consent screen + publisher verification | **Week -3** | Google: 2-4 weeks lead time. Microsoft: 3-6 weeks lead time. Delay → S2/S3 blocked → use unverified (limited users) → production delayed to W16-18. See [11-third-party-integration-principles.md](11-third-party-integration-principles.md) Gates 1 & 2. |
| 🔴 **2** | **Post ML Engineer JD ngay hôm nay** + đặt W6 go/no-go checkpoint (nếu chưa có candidate → engage ML contractor) | **Day 1** | Lead time hire ML Eng = 8–15 tuần; post W4 = arrive W13–18 = Track 2 delayed 4–9 tuần |
| 🔴 **3** | **Submit all Category B 3rd-party API access requests** (Slack, Hive, Lakera, Apple, Google Play) Day 1-2 of Sprint 1 | **Week 1 Day 1-2** | 1-2 week lead time each. Delay → S5/S8/S10 features blocked or cut. See [11-third-party-integration-principles.md](11-third-party-integration-principles.md) Category B. |
| 🔴 **4** | **Engage BD consultant** (3 days/week từ W1) để own pilot customer pipeline. Define funnel: 100 leads → 30 qualified → 15 demo → 5 pilot | **Day 3** | Pipeline cần bắt đầu W1 để close 5 customers ở W30. 0.5 FTE PM không thể vừa run sprints vừa do sales |
| 🟠 **5** | **Add M365 webhook renewal service vào S1a scope** + resolve GDPR Article 17 vs S3 Object Lock bằng envelope encryption architecture | **Day 1** | Cả hai là irreversible architecture decisions. Phát hiện ở S3 = costly rework. Phát hiện sau EU launch = GDPR penalty |
| 🟠 **6** | **Engage legal counsel** cho Module B/D/F liability review và EU voice biometric opinion. Target: opinion delivered trước S8 (W17) để không block Track 2 design | **Week 1** | Nếu legal opinion W8 = "yes biometric" và không có sprint buffer → Track 2 stalls với không có plan |
| 🟠 **7** | **Vanta account setup** (Week 8) + pentest vendor RFP (Week 8) + pentest LOI signed (Week 14 hard deadline) | **Week 8 start, Week 14 LOI** | Vanta: 2-3 weeks to active. Pentest: 6-8 weeks from RFP to kickoff. Delay → SOC 2 Type 1 insufficient evidence OR pentest delayed → v1 delayed. See [11-third-party-integration-principles.md](11-third-party-integration-principles.md) Gates 4 & 5. |

---

## Appendix: Risk Register Tổng Hợp

| ID | Risk | Severity | Source | Mitigation | Sprint |
|---|---|---|---|---|---|
| R-C1 | False positive automated offboarding → legal liability | CRITICAL | TA + PO | Grace period + rollback + ToS cap | S6 |
| R-C2 | Google API rate limit breach at scale | CRITICAL | TA | Per-tenant quota distribution | S1a → S2 |
| R-C3 | M365 webhook expiry silent failure | CRITICAL | TA + PM | Renewal service + DLQ + fallback | S1a → S3 |
| R-C4 | GDPR erasure vs S3 Object Lock | CRITICAL | TA + PO | Envelope encryption + key deletion | S1a design → S11 impl |
| R-C5 | No CAC model / pilot pipeline | CRITICAL | PO | MSP channel + freemium + BD hire | Week 1 |
| R-C6 | Keycloak SPOF | CRITICAL | TA | JWKS caching + HA + Keycloak upgrade | S1b |
| R-H1 | S1 over-scoped 230% | HIGH | PM | Split S1a/S1b | Day 1 |
| R-H2 | Chrome MV3 service worker kill mid-scan | HIGH | TA | Explicit fail-closed on interrupt; service worker keepalive | S8 |
| R-H3 | PgBouncer missing | HIGH | TA | RDS Proxy in S1a | S1a |
| R-H4 | EventBridge duplicate events | HIGH | TA | Idempotency keys on Step Functions | S6 |
| R-H5 | WASM model no signature verification | HIGH | TA | Cosign model artifacts before S3 load | S8 |
| R-H6 | Shadow AI score formula gameable | HIGH | TA | Fix `app_age_days` direction; annual weight review | S4 |
| R-H7 | Test tenants not in S1 | HIGH | PM | Add to S1a scope | S1a |
| R-H8 | M4 onboarding shock (3 new hires) | HIGH | PM | Onboarding docs ready W12; S7 scoped conservatively | Pre-S7 |
| R-H9 | Pentest timeline insufficient | HIGH | PM | Start pentest W21 (not W21 scoping + W21 start). **LOI must be signed W14 (hard gate).** See [11-third-party-integration-principles.md](11-third-party-integration-principles.md) Gate 4. | W21 |
| R-H10 | Chrome Web Store review risk | HIGH | PM | Submit stripped v0 extension at W18. **Register developer account W10.** See [11-third-party-integration-principles.md](11-third-party-integration-principles.md) Category B. | S8 |
| R-H11 | Google Workspace verification delayed >6 weeks | HIGH | PM | Submitted Week -3. **Fallback:** Unverified OAuth (100 user limit) for pilot, defer production to W16. See [11-third-party-integration-principles.md](11-third-party-integration-principles.md) Gate 1. | Week -3 → Week 2-4 |
| R-H12 | Microsoft 365 publisher verification delayed >8 weeks | HIGH | PM | Submitted Week -3. **Fallback:** Unverified app (10 user limit) for pilot, defer production to W18. See [11-third-party-integration-principles.md](11-third-party-integration-principles.md) Gate 2. | Week -3 → Week 3-6 |
| R-H13 | Lakera Guard pricing not viable (>$0.05/request) | MEDIUM | PM + ML Eng | **Go/No-go decision Week 2 (S1 end).** Fallback: WASM-only BERT (TPR ~75%, FPR ~10%). See [11-third-party-integration-principles.md](11-third-party-integration-principles.md) Gate 3. | Week 2 |
| R-H14 | Vanta setup delayed >3 weeks | MEDIUM | PM | Start Week 8. **Fallback:** Manual evidence collection (higher PM workload ~20h/week). See [11-third-party-integration-principles.md](11-third-party-integration-principles.md) Gate 5. | Week 8 → Week 11 |
| BS-1 | Alert fatigue Day 2 | HIGH | PO | 14-day calibration + curated alert strategy | S4 |
| BS-2 | Module F legal trap | HIGH | PO | Remove from v1; admin-only transparency | Pre-launch |
| BS-3 | Zero switching cost | MEDIUM | PO | Playbook builder in MVP as primary retention feature | S6 |
| BS-4 | BERT prompt injection not feasible | MEDIUM | TA | Use Lakera Guard; defer internal BERT | S8 |
| BS-5 | Deepfake liability no insurance | HIGH | PO | E&O insurance + advisory UI language | Pre-launch |
| BS-6 | New AI tool recognition lag | MEDIUM | TA | Honest marketing; server-push rules | S8 |
| BS-7 | Compliance report not audit-ready | MEDIUM | PO | Sales disclaimer; auditor relationship build | v1 comms |
| BS-8 | EU biometric opinion triggers redesign | HIGH | PM | Legal opinion delivered pre-S8 | Week 1 |
| BS-9 | No QA role in Phase 1 | HIGH | PM | Pair programming + rotating review; E2E test owner assigned | S1a |
| BS-10 | Slack tier not acknowledged | MEDIUM | TA | UI tier detection + graceful degradation | S5 |
