# SMESec Platform — Delivery Plan

**Ngày tạo:** 2026-05-28  
**Trạng thái:** Approved — Tổng hợp từ 3 agent (Product Owner · Project Manager · Technical Advisor)  
**Phiên bản:** 1.0  
**Phạm vi:** Toàn bộ lộ trình từ Sprint 1 đến v2 (12 tháng)

---

## Mục Lục

1. [Tổng Quan Lộ Trình](#1-tổng-quan-lộ-trình)
2. [Scope Theo Từng Milestone](#2-scope-theo-từng-milestone)
3. [Team & Nhân Sự Tăng Dần](#3-team--nhân-sự-tăng-dần)
4. [Chia Nhỏ Theo Sprint](#4-chia-nhỏ-theo-sprint)
   - [Phase 1: Foundation → MVP (Tháng 1–3, S1–S6)](#phase-1-foundation--mvp-tháng-13-s1s6)
   - [Phase 2: MVP → v1 (Tháng 4–6, S7–S13)](#phase-2-mvp--v1-tháng-46-s7s13)
   - [Phase 3: v1 → v1.5 (Tháng 7–9, S14–S20)](#phase-3-v1--v15-tháng-79-s14s20)
   - [Phase 4: v1.5 → v2 (Tháng 10–12, S21–S26)](#phase-4-v15--v2-tháng-1012-s21s26)
5. [Phân Bổ Team Sau v1 (Hai Luồng Song Song)](#5-phân-bổ-team-sau-v1-hai-luồng-song-song)
6. [Đáp Ứng Key Requirements Đề Bài](#6-đáp-ứng-key-requirements-đề-bài)
7. [Giả Định Rủi Ro Nhất Cần Validate Sớm Nhất](#7-giả-định-rủi-ro-nhất-cần-validate-sớm-nhất)
8. [Compliance Certification Timeline](#8-compliance-certification-timeline)
9. [External Dependencies & Hard Deadlines](#9-external-dependencies--hard-deadlines)

---

## 1. Tổng Quan Lộ Trình

```
Tháng:  1    2    3    4    5    6    7    8    9   10   11   12
Sprint: S1  S2  S3  S4  S5  S6  S7  S8  S9  S10 S11 S12 S13 S14 S15 S16 S17 S18 S19 S20 S21 S22 S23 S24 S25 S26
        |------Phase 1: Foundation-------|----Phase 2: MVP→v1----|---Phase 3: v1→v1.5-----|----Phase 4: v1.5→v2----|
                            ↑                                   ↑                   ↑                              ↑
                           MVP                                  v1                 v1.5                            v2
                        (W12/T3)                           (W26/T6)           (W38/T9)                       (W52/T12)
```

| Milestone | Tuần | Tháng | Mô Tả |
|-----------|------|-------|-------|
| **MVP** | W12 | T3 | Asset inventory + offboarding tự động + shadow IT — pilot customers onboard |
| **v1** | W26 | T6 | Tất cả key requirements từ đề bài đều có mặt. SOC 2 Type 1 audit lên lịch. |
| **v1.5** | W38 | T9 | AI detection nâng cao + AWS v1.1 + feedback pilot tích hợp. Billing tiers live. |
| **v2** | W52 | T12 | SOC 2 Type 2 + ISO 27001 certified. Enterprise tier. ML features production-ready. |

---

## 2. Scope Theo Từng Milestone

### Bản Đồ Feature Theo Giai Đoạn

| Feature Domain | MVP (T3) | v1 (T6) | v1.5 (T9) | v2 (T12) |
|---|---|---|---|---|
| **Asset Inventory** | Google WS + M365: users, OAuth apps, devices cơ bản | + Slack, AWS, Shadow AI detection | + Custom asset types, dependency map | + Full cloud posture, peer anomaly |
| **Access Governance** | Automated offboarding <5 min, RBAC dashboard | + JIT access, access reviews, shadow IT remediation | + Risk scoring, access policy templates | + Peer group anomaly, insider threat signal |
| **AI Threat Surface** | ❌ (Track 2 đang R&D) | Shadow AI governance + LLM DLP browser ext (beta) | + Deepfake defense, AI phishing, prompt injection v1 | + Prompt injection ML (BERT), advanced analytics |
| **Compliance Posture** | Evidence collection bắt đầu (silent, no UI) | Dashboard compliance, SOC 2 Type 1 + ISO 27001 report-ready | SOC 2 Type 2 evidence running (90 ngày) | SOC 2 Type 2 certified + ISO 27001 certified |
| **Incident Playbooks** | 2 playbooks (Offboarding, Cred Compromise) | 5 playbooks, wizard UI, AWS Step Functions | + Custom playbook builder, mobile triggers | + Playbook analytics, ML suggestions |
| **Integrations** | Google WS + M365 (OAuth wizard <30 min) | + Slack full + AWS IAM cơ bản | + AWS CloudTrail, S3 audit, IAM deep | + SIEM (Splunk/QRadar), custom webhooks |
| **Mobile App** | ❌ TestFlight/Beta | Alerts + playbook trigger (iOS + Android) | Full incident response mobile | Full feature parity |
| **Billing / Pricing** | Manual invoicing (pilot free) | Starter + Growth tiers code-ready | Pricing tiers enforced, billing live | Enterprise custom + usage-based |

### MVP — Định Nghĩa Giá Trị Tối Thiểu

> **Câu hỏi MVP cần trả lời được:** *"Bạn có biết bao nhiêu ứng dụng đang kết nối vào Google Workspace / M365 của bạn, và bạn có thể thu hồi toàn bộ quyền truy cập của một nhân viên nghỉ việc trong 5 phút không?"*

```
MVP = Sprint 6 hoàn thành (cuối Tuần 12)

✅ OAuth wizard: Google Workspace + M365, setup <30 phút
✅ Asset inventory dashboard: users, OAuth apps, device cơ bản
✅ Shadow IT detection: alert khi có OAuth app mới trong <15 phút
✅ Automated offboarding: revoke toàn bộ quyền truy cập <5 phút
✅ 2 incident playbooks: Offboarding Emergency + Credential Compromise
✅ RBAC dashboard: xem phân quyền, recommendation least-privilege
✅ Keycloak SSO + MFA bắt buộc
✅ Compliance evidence collection: bắt đầu chạy ngầm từ ngày 1
✅ Mobile app beta: TestFlight + Play Console

❌ KHÔNG có trong MVP:
  - AI/ML detection (Track 2 đang R&D)
  - Full compliance reports
  - JIT access
  - Slack integration
  - Billing system
  - Deepfake / prompt injection
```

---

## 3. Team & Nhân Sự Tăng Dần

### Timeline Nhân Sự

```
Tháng 1–3 (Phase 1 / MVP):         6 FTE core + DevSecOps contract
Tháng 4 (đầu Phase 2):             + ML Engineer #1 (Track 2)
Tháng 4 (Sprint 7):                + Backend Engineer #3 (Track 2)
Tháng 4–5 (Sprint 8):              + Frontend Engineer #2 (Browser Extension)
Tháng 6 (Sprint 13 / v1 launch):   DevSecOps → FTE (không còn contract)
Tháng 7 (đầu Phase 3):             + Customer Success Engineer
Tháng 8 (giữa Phase 3):            + ML Engineer #2 (optional, tùy v1 velocity)
Tháng 10–12 (Phase 4):             + Compliance Consultant (contract)
```

### Bảng Nhân Sự Chi Tiết Theo Giai Đoạn

| Vai Trò | T1–T3 | T4–T6 | T7–T9 | T10–T12 | Track |
|---------|-------|-------|-------|---------|-------|
| Tech Lead / Architect | ✅ 1.0 FTE | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | 1 |
| Backend Eng #1 (Go) | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | 1 |
| Backend Eng #2 (Go/Python) | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | 1 |
| Frontend Eng #1 (React) | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | 1 |
| Flutter / Mobile Eng | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | ✅ 1.0 | 1 |
| DevSecOps | Contract (0.5) | Contract (0.5) | **FTE (1.0)** | **FTE (1.0)** | Shared |
| PM | 0.5 | 0.5 | 0.5 | 0.5 | Shared |
| **ML Engineer #1** | — | **✅ 1.0 (T4)** | ✅ 1.0 | ✅ 1.0 | 2 |
| **Backend Eng #3 (Python/FastAPI)** | — | **✅ 1.0 (T4)** | ✅ 1.0 | ✅ 1.0 | 2 |
| **Frontend Eng #2 (Browser Ext)** | — | **✅ 1.0 (T4.5)** | ✅ 1.0 | ✅ 1.0 | 2 |
| **Customer Success Engineer** | — | — | **✅ 1.0 (T7)** | ✅ 1.0 | Customer |
| **ML Engineer #2** | — | — | **✅ 1.0 (T8, opt.)** | ✅ 1.0 | 2 |
| **Compliance Consultant** | — | — | — | **Contract (T10–T12)** | Compliance |
| **Tổng FTE** | **6** | **8.5 → 9** | **10 → 11** | **11.5** | |

> **Nguyên tắc hiring:** ML Engineer phải có kinh nghiệm SageMaker hoặc managed ML platform — không phải researcher hàn lâm. Bắt đầu tuyển từ Tuần 5 (Sprint 3) để onboard kịp Sprint 6.

---

## 4. Chia Nhỏ Theo Sprint

### Phase 1: Foundation → MVP (Tháng 1–3, S1–S6)

**Team Phase 1:** Tech Lead · BE1 · BE2 · FE1 · Flutter · DevSecOps(contract) · PM = **6 FTE**

---

#### S1 — W1–2: Infrastructure & Auth

| | |
|---|---|
| **Mục tiêu** | Nền tảng kỹ thuật: deploy được, đăng nhập được |
| **Deliverable cuối sprint** | Engineer đăng nhập vào web app bằng Google/M365 thật. Staging deploy từ CI tự động. |
| **Scope** | AWS VPC + ECS Fargate + RDS PostgreSQL Multi-AZ · S3 Object Lock (audit log) · Keycloak SSO (Google + M365) · MFA TOTP bắt buộc · CI/CD GitHub Actions · Multi-tenant schema (`tenant_id` + `data_residency` trên mọi bảng, RLS enforced) · `ThreatDetectionEvent` interface draft (T1-T2 schema contract) |
| **Rủi ro chính** | Quyết định Auth provider (Auth0 vs Cognito vs Keycloak self-host) phải xong ngày 1 |
| **Action PM** | Bắt đầu tuyển ML Engineer #1. Chuẩn bị danh sách pilot customers tiềm năng. |

> **Gate bắt buộc:** `data_residency` column phải có ngay từ S1 — nếu bỏ qua, phải refactor toàn bộ schema khi gần MVP. Tenant isolation CI test phải xanh trước khi merge bất kỳ code nào.

---

#### S2 — W3–4: Google Workspace Sync

| | |
|---|---|
| **Mục tiêu** | Nhìn thấy users + OAuth apps từ Google Workspace |
| **Deliverable cuối sprint** | Dashboard hiển thị user list + OAuth apps từ Google tenant thật. First-value demo <30 phút từ khi OAuth grant. |
| **Scope** | Google Admin SDK: user/group/device sync · OAuth app discovery (scope risk analysis) · 15-min incremental sync (delta pull) · Asset inventory DB schema v1 · Shadow IT detection rules v1 (high-risk OAuth scopes) · Dashboard skeleton (data visible, no styling required) |
| **Rủi ro chính** | Google Admin SDK pagination + rate limits — validate trên tenant thật >100 users ngay S1 skeleton |
| **Action PM** | Pilot outreach bắt đầu. Target 3–5 SME (50–200 nhân viên) cho onboard tháng 3. |

---

#### S3 — W5–6: M365 Sync + Dashboard v1

| | |
|---|---|
| **Mục tiêu** | Unified dashboard: Google + M365 cùng một màn hình |
| **Deliverable cuối sprint** | Dashboard hiển thị assets từ cả Google + M365. Risk indicators per user/app. Export CSV. |
| **Scope** | Microsoft Graph API + Azure AD: user/app/device sync · M365 delta link + webhook · Cross-provider identity matching (email canonical) · Unified risk indicators (per-provider, không phải composite) · Dashboard polish: filter, search, sort |
| **Rủi ro chính** | M365 OAuth permission consent — cần IT Admin guide chi tiết. Prepare "minimum-permission scope explainer" cho khách hàng. |
| **Action PM** | Ký LOI tuyển ML Engineer #1 trong tuần này. Pilot customer list phải có ít nhất 5 leads. |

---

#### S4 — W7–8: Classification + Shadow IT Alerts

| | |
|---|---|
| **Mục tiêu** | IT admin phân loại được asset, nhận alert khi có OAuth app mới |
| **Deliverable cuối sprint** | Shadow IT alerts firing đúng. Asset classifications visible. Flutter mobile scaffold chạy được trên iOS + Android. |
| **Scope** | Asset classification engine (criticality + data sensitivity, rule-based) · OAuth scope risk scoring (high/medium/low) · New OAuth app alert pipeline (<15 min) · Email + Slack notification system · Mobile scaffold (Flutter): auth flow Keycloak PKCE, navigation shell, push notification skeleton |
| **Rủi ro chính** | Alert noise quá cao → pilot user bị overwhelm. Bắt đầu với threshold conservative (chỉ HIGH risk alerts). |

---

#### S5 — W9–10: Slack + AWS Discovery + RBAC

| | |
|---|---|
| **Mục tiêu** | 4 nguồn tích hợp (Google, M365, Slack, AWS). RBAC dashboard live. |
| **Deliverable cuối sprint** | Unified inventory 4 providers. Least-privilege recommendations hiển thị. Slack deactivation tested. |
| **Scope** | Slack Admin API: users, apps, channels · Slack tier detection (Free/Pro/Business+ gating) · AWS IAM inventory: users, roles, policies · RBAC model: role assignment, permission diff engine · Least-privilege recommendations (rule-based) · Composite identity graph (cross-provider) |
| **Rủi ro chính** | Slack API tier limitation — Business+ required cho automated offboarding. Detect tier sớm và set expectation với pilot customers. |
| **Action PM** | ⚠️ **ML Engineer #1 phải onboard tuần này (W9)**. Bắt đầu R&D shadow AI governance trên synthetic data. |

---

#### S6 — W11–12: Automated Offboarding + 2 Playbooks — **🏁 MVP**

| | |
|---|---|
| **Mục tiêu** | Offboard nhân viên trong <5 phút. 2 playbooks. Mobile app beta. |
| **Deliverable cuối sprint** | **MVP: Offboarding test user <5 min qua Google+M365+Slack. Mobile app trên TestFlight/Play Console. PDF offboarding report.** |
| **Scope** | Automated offboarding workflow (AWS Step Functions): disable + revoke + notify · Dry-run + 2-step confirmation (hard gate, không bypass) · Offboarding report PDF · 2 incident playbooks: (1) Offboarding Emergency (2) Credential Compromise · Playbook wizard UI (web) · Immutable audit log: PostgreSQL append-only + S3 · Mobile app v1: alerts, offboarding trigger, read-only inventory |
| **Rủi ro chính** | Sprint có utilization cao nhất Phase 1 (~89%). Mobile scope phải cut nếu cần — offboarding là priority tuyệt đối. |
| **Action PM** | ✅ 3+ pilot customers phải onboard trên staging environment trước cuối W12. |

> **MVP Gate Checklist:**
> - [ ] Offboarding <5 min (timed automated test pass trong CI)
> - [ ] Shadow IT alert <15 min từ OAuth grant đến notification
> - [ ] Tenant isolation CI test xanh liên tục
> - [ ] 3+ pilot customers đã thấy "first insight" trong <30 phút setup
> - [ ] Zero plaintext secrets trong environment variables
> - [ ] RDS Multi-AZ + S3 Object Lock active

---

### Phase 2: MVP → v1 (Tháng 4–6, S7–S13)

**Team Phase 2:** 6 FTE (Phase 1) + ML Eng #1 + BE3 + FE2 = **9 FTE** (tăng dần từ S7→S8)

**Đặc điểm Phase 2:** Track 2 (AI/ML) bắt đầu tích hợp với Track 1 từ S7. Hai track song song, converge tại S11.

---

#### S7 — W13–14: JIT Access + Track 2 Integration Begins + Vanta Setup

| | |
|---|---|
| **Mục tiêu** | JIT access end-to-end. Track 2 onboard live data. Vanta bắt đầu thu evidence. |
| **Deliverable cuối sprint** | JIT request → approve → auto-revoke hoạt động. Track 2 shadow AI model nhận live data từ Track 1. Vanta dashboard green cho các controls đã setup. |
| **Scope — Track 1** | JIT access: request form + approval workflow + time-boxed grant + auto-revoke · Access review scheduling (periodic reminder) · Pilot feedback triage từ MVP (top 10 bugs) |
| **Scope — Track 2** | BE3 onboard: môi trường, codebase walkthrough · ML Eng: shadow AI governance v1 kết nối live `oauth_application` table từ Track 1 · OAuth risk score model v0.2 training trên live data |
| **Rủi ro chính** | JIT access approval workflow phức tạp hơn dự kiến. Simplify: 1 approver, email-based, không cần self-service portal trong v1. |
| **Action PM** | ⚠️ **Ký LOI pentest vendor trước cuối W14 — hard deadline không nhượng bộ.** · ⚠️ **Vanta account provisioned W13 — evidence collection phải bắt đầu ngay.** |

---

#### S8 — W15–16: Playbook Engine + 3 Playbooks + LLM DLP Prototype

| | |
|---|---|
| **Mục tiêu** | AWS Step Functions playbook engine. 3 playbooks đầu. Browser extension prototype. |
| **Deliverable cuối sprint** | 3 playbooks chạy end-to-end trên staging. LLM DLP browser extension có thể detect PII trong text field. |
| **Scope — Track 1** | AWS Step Functions playbook engine · Playbook wizard UI (web) · 3 playbooks: (1) Account Compromise (2) Phishing Response (3) Data Exfiltration · Playbook audit log (mỗi step được log) |
| **Scope — Track 2** | LLM DLP browser extension v0.1 (Chrome Manifest V3): PII detection với Microsoft Presidio (local inference, không call API) · FE2 onboard: môi trường, Chrome Extension CI/CD setup |
| **Rủi ro chính** | ⚠️ Sprint LOADED NHẤT trong toàn bộ plan (~88% utilization cả 2 tracks). PM cần daily standup S8. |
| **Action PM** | FE2 phải onboard đầu W15. Nếu chậm → LLM DLP shift sang S9. |

---

#### S9 — W17–18: 5 Playbooks + Mobile + Shadow AI v1

| | |
|---|---|
| **Mục tiêu** | Đủ 5 playbooks. Mobile incident alerts. Shadow AI governance v1 production. |
| **Deliverable cuối sprint** | 5 playbooks hoàn chỉnh. Push notification từ mobile cho security alerts. Shadow AI risk scores live (OAuth apps được classify AI/non-AI). |
| **Scope — Track 1** | 2 playbooks còn lại: (4) Ransomware Response (5) Insider Threat Response · Mobile push notifications (FCM + APNs) · Incident alert từ playbook → mobile |
| **Scope — Track 2** | Shadow AI governance v1: AI tool classification (ChatGPT, Copilot, Gemini, Claude, etc.) + risk score per OAuth app · Shadow AI attestation workflow: employee confirm/deny usage · LLM DLP extension: tenant-scoped allow-list, PII blocking trước khi submit |
| **Rủi ro chính** | Shadow AI classification accuracy thấp → nhiều false positive → pilot users phàn nàn. Bắt đầu với conservative threshold (only block HIGH risk apps đã confirmed AI). |

---

#### S10 — W19–20: Compliance Mapping + T1-T2 Integration Contract

| | |
|---|---|
| **Mục tiêu** | Compliance dashboard với Vanta. API contract T1-T2 được ký. |
| **Deliverable cuối sprint** | Compliance dashboard: coverage % ISO 27001 + SOC 2. Deepfake defense prototype. `ThreatDetectionEvent` schema v1 locked. |
| **Scope — Track 1** | ISO 27001 + SOC 2 control mapping trong Vanta · Automated evidence collection hooks · Compliance dashboard (control status, evidence links) · Cross-provider composite risk score (per user, weighted) |
| **Scope — Track 2** | Deepfake defense: Hive Moderation API POC (pay-per-use, rate limit test) · Out-of-band verification workflow design · **`ThreatDetectionEvent` schema v1 finalized và locked** |
| **Action kỹ thuật** | ⚠️ **`ThreatDetectionEvent` schema phải được approve cuối S10.** Delay ở đây cascade thẳng vào S11 integration sprint. |

---

#### S11 — W21–22: Compliance Reports + T1-T2 Integration Live

| | |
|---|---|
| **Mục tiêu** | Compliance reports có thể export. AI threats trigger Track 1 playbooks tự động. |
| **Deliverable cuối sprint** | PDF compliance report (ISO 27001 + SOC 2 Type 1 evidence). Track 2 AI threat event → auto-trigger Step Functions playbook trong staging. |
| **Scope — Track 1** | ISO 27001 + SOC 2 compliance reports (PDF export) · Audit trail UI · GDPR data subject request automation (export + delete) |
| **Scope — Track 2** | T1-T2 integration: `ThreatDetectionEvent` → EventBridge → Step Functions trigger · Prompt injection detection v1 (rule-based regex) · AI phishing: M365 Defender + Google Workspace threat feed connected |
| **Rủi ro chính** | ⚠️ **Đây là sprint rủi ro kỹ thuật CAO NHẤT** — integration luôn mất gấp 3x thời gian dự kiến. Tech Lead phải full-time trên integration này. Fallback: nếu auto-trigger không ổn định → manual trigger (button) cho v1, auto-trigger sang v1.5. |
| **Action PM** | ⚠️ **Pentest phải BẮT ĐẦU tuần W21** (theo LOI đã ký S7). Coordinate với vendor. |

---

#### S12 — W23–24: Dependency Map + Pentest Remediation + Vanta Dry Run

| | |
|---|---|
| **Mục tiêu** | Full T1-T2 integration validated. Pentest findings remediated. Vanta dry run pass >90%. |
| **Deliverable cuối sprint** | App dependency map live (SaaS lifecycle management). Pentest findings: tất cả Critical + High resolved. Vanta evidence dry run pass rate >90%. |
| **Scope — Track 1** | SaaS dependency mapping + lifecycle management (zombie app detection) · Pentest: remediate tất cả Critical và High findings (Pentest chạy W21–W23) · Vanta compliance evidence dry run |
| **Scope — Track 2** | Full T1-T2 end-to-end integration test (automated) · Shadow AI governance: policy enforcement mode (block vs alert) |
| **Rủi ro chính** | Pentest Critical finding trong infrastructure (không phải code) → cần DevSecOps thêm thời gian. Buffer: S12 có 20% slack để handle. |

---

#### S13 — W25–26: Hardening + v1 Launch — **🏁 v1**

| | |
|---|---|
| **Mục tiêu** | Launch v1 production. SOC 2 Type 1 audit lên lịch. |
| **Deliverable cuối sprint** | **v1 LIVE trên production. 5+ pilot customers chuyển sang production. SOC 2 Type 1 audit engagement signed.** |
| **Scope** | KHÔNG có feature mới · Performance hardening · Pentest Medium findings remediation · Launch runbook · Production cutover · SOC 2 Type 1 readiness review với Vanta auditor · Marketing launch brief |
| **Utilization mục tiêu** | 60% — deliberate sprint đệm |

> **v1 Gate Checklist:**
> - [ ] Tất cả 5 incident playbooks production
> - [ ] JIT access + offboarding <5 min (CI test)
> - [ ] 4 integrations (Google, M365, Slack, AWS)
> - [ ] Compliance dashboard: ISO 27001 + SOC 2 Type 1 report-ready
> - [ ] AI threat module: shadow AI governance + LLM DLP extension
> - [ ] Mobile app: iOS App Store + Google Play (submit S12, review ~1 tuần)
> - [ ] Pentest: zero Critical/High findings open
> - [ ] 5+ pilot customers trên production
> - [ ] SOC 2 Type 1 audit đã lên lịch với auditor
> - [ ] CloudWatch monitoring + PagerDuty alerting live
> - [ ] Disaster recovery runbook tested (RTO <4h)

---

### Phase 3: v1 → v1.5 (Tháng 7–9, S14–S20)

**Team Phase 3:** 9 FTE (Phase 2) + Customer Success Eng (T7) + ML Eng #2 (T8) + DevSecOps → FTE = **11 FTE**

**Đặc điểm Phase 3:** Team chia 2 luồng song song. Luồng A (phát triển theo plan). Luồng B (cập nhật theo feedback pilot).

#### Phân Bổ Team Phase 3 — Hai Luồng Song Song

| Luồng | Tỷ lệ | Thành viên | Focus |
|-------|-------|------------|-------|
| **Luồng A — New Features** | **65%** | Tech Lead · BE1 · BE2 · FE1 · ML Eng #1 · ML Eng #2 | AWS v1.1, AI detection nâng cao, SOC 2 Type 2 prep, Business tier |
| **Luồng B — Pilot Feedback** | **35%** | BE3 · FE2 · Customer Success Eng · Flutter (40%) | Bug fixes, UX polish, onboarding friction, customer requests |

> **Nguyên tắc chia luồng:**
> - Mỗi tuần: PM triage feedback queue vào thứ 2. Issues >1 sprint → backlog v1.5. Issues <0.5 sprint → Luồng B xử lý ngay.
> - Luồng B KHÔNG nhận feature mới — chỉ fix và polish.
> - Hai luồng converge tại v1.5 release (W38).

---

#### S14–S15 — W27–30: Post-launch Stabilization + AWS v1.1

| | |
|---|---|
| **Luồng A** | AWS IAM deep integration: CloudTrail events, S3 access auditing, IAM role recommendations · LLM DLP browser extension v1 (Chrome Web Store submit) |
| **Luồng B** | Top 10 customer-reported bugs · M365 OAuth wizard UX improvements · Mobile crash fixes · Alert threshold tuning (reduce noise) |
| **Deliverable** | AWS v1.1 production. Browser extension submitted to Chrome Web Store. Sprint 14 utilization 65% (recovery sprint sau 6 tháng cao điểm). |
| **Action PM** | Customer Success Engineer onboard T7W1. Sprint 14 kickoff bao gồm retrospective v1. |

---

#### S16–S17 — W31–34: Advanced AI Detection v2

| | |
|---|---|
| **Luồng A** | LLM data leakage detection v2: real-time DLP (semantic analysis, không chỉ regex) · Deepfake fraud defense v2: Hive API live + out-of-band verification workflow (SMS + Slack) · Prompt injection hardening (expanded ruleset) · ML Eng #2 onboard W32 |
| **Luồng B** | Dashboard UX redesign (based on pilot feedback) · Custom alert rules UI · API documentation · Auditor-specific compliance export templates |
| **Deliverable** | AI detection accuracy >90% trên test set. Customer-configurable alert rules. v2 UX pilot tested. |
| **Rủi ro** | Browser extension bị Chrome Web Store reject (review 2–4 tuần) → submit W29 (trước S15), có buffer. |

---

#### S18–S19 — W35–38: Business Tier + SOC 2 Type 2 Prep — **🏁 v1.5**

| | |
|---|---|
| **Luồng A** | Pricing tier enforcement (Starter/Growth/Business gates) · Vanta SOC 2 Type 2 evidence framework setup · Advanced compliance reporting · Custom playbook builder (Luồng A) · ISO 27001 evidence continuation |
| **Luồng B** | Pilot → paid customer conversion flow · Billing integration (Stripe) · Customer portal · Custom playbook builder (UX, phối hợp cùng Luồng A) |
| **Deliverable** | **v1.5 LAUNCH (W38).** Pricing tiers enforced. Billing live. 10+ paying customers. SOC 2 Type 2 evidence collection đang chạy liên tục kể từ W26. |

> **v1.5 Gate Checklist:**
> - [ ] AWS v1.1 production (CloudTrail, IAM deep)
> - [ ] Browser extension: Chrome Web Store published (không phải sideload)
> - [ ] AI detection accuracy >90% (deepfake + LLM DLP)
> - [ ] Prompt injection detection v1 (rule-based) production
> - [ ] Pricing tiers enforced (Starter / Growth / Business)
> - [ ] Billing integration live (Stripe)
> - [ ] 10+ paying customers trên production
> - [ ] SOC 2 Type 2 evidence collection chạy từ W26 (>12 tuần evidence rồi)
> - [ ] Custom playbook builder beta
> - [ ] SageMaker model monitoring (drift detection) active

---

### Phase 4: v1.5 → v2 (Tháng 10–12, S21–S26)

**Team Phase 4:** 11 FTE + Compliance Consultant (contract T10–T12) = **11.5 FTE (peak)**

**Đặc điểm Phase 4:** Feature freeze tháng 10. Tập trung SOC 2 Type 2 audit, ISO 27001 certification, Enterprise tier, BERT ML production.

---

#### S21–S22 — W39–44: Enterprise Features + SOC 2 Type 2 Audit Prep

| | |
|---|---|
| **Scope** | Enterprise tier features: multi-tenant enterprise, custom RBAC policies, SIEM integration (Splunk/QRadar webhooks) · Vanta SOC 2 Type 2 evidence final packaging · SOC 2 Type 2 audit engagement signed |
| **Deliverable** | Enterprise tier code-complete. SOC 2 Type 2 audit engagement signed với auditor. Evidence coverage >95% trong Vanta. |
| **Timeline SOC 2 Type 2** | Evidence collection started W26 → audit window W26–W52 (26 tuần = 6 tháng ✅) · Audit fieldwork: W46–W48 · Report issued: W50–W52 |

---

#### S23–S24 — W45–48: ISO 27001 Certification + BERT Prompt Injection

| | |
|---|---|
| **Scope** | ISO 27001 Stage 2 audit prep + Statement of Applicability finalized · BERT prompt injection classifier: fine-tuned trên 6 tháng production data (Enterprise tier only) · Advanced analytics dashboard (SOC-level insights) · Peer group anomaly detection v1 (insider threat signal) |
| **Deliverable** | ISO 27001 Stage 2 audit complete. BERT model: FPR <2%, TPR >85% trên 30-ngày holdout set. |
| **Rủi ro** | BERT FPR quá cao → ship rule-based prompt injection (đã có) + BERT as opt-in preview, không GA. |

---

#### S25–S26 — W49–52: v2 Launch + Compliance Certified — **🏁 v2**

| | |
|---|---|
| **Scope** | Compliance certification received · Enterprise tier GA · White-label / MSSP foundation · Usage-based billing option · Multi-region DR test (không chỉ documented) · All Track 2 features graduate từ beta (SLA applies) |
| **Deliverable** | **v2 LAUNCH (W52).** SOC 2 Type 2 certified. ISO 27001 certified. Enterprise tier live. |

> **v2 Gate Checklist:**
> - [ ] SOC 2 Type 2 report received từ auditor
> - [ ] ISO 27001 certificate received
> - [ ] BERT prompt injection: FPR <2%, TPR >85% (hoặc opt-in preview nếu chưa đạt)
> - [ ] Enterprise tier: custom pricing, dedicated CSM, SIEM integration
> - [ ] Advanced analytics dashboard production
> - [ ] Peer group anomaly detection production
> - [ ] All Track 2 features: beta flag removed, SLA guarantees
> - [ ] Multi-region DR failover drill tested (RTO/RPO documented)
> - [ ] 99.95% uptime SLA target achievable (verified from monitoring data)

---

## 5. Phân Bổ Team Sau v1 (Hai Luồng Song Song)

### Cấu Trúc Chi Tiết Hai Luồng (Phase 3 & 4)

```
LUỒNG A — New Features (65%)          LUỒNG B — Pilot Feedback (35%)
────────────────────────────           ────────────────────────────────
Tech Lead                              Customer Success Engineer
Backend Eng #1                         Backend Eng #3
Backend Eng #2                         Frontend Eng #2
Frontend Eng #1                        Flutter Eng (40% time)
ML Eng #1
ML Eng #2

Làm gì:                                Làm gì:
  - Pre-planned roadmap features         - Bug queue từ pilot/customers
  - v1.5 capabilities                    - UX friction từ usage analytics
  - SOC 2 Type 2 evidence prep           - Feature requests <2 ngày work
  - Enterprise tier                      - Onboarding wizard improvements
  - AI accuracy improvements             - Alert noise tuning
  - New integrations                     - Performance issues

Cadence:                               Cadence:
  2-tuần sprint bình thường               Weekly triage (thứ 2)
  Sprint planning thứ 2 đầu sprint        Continuous deployment
  Demo cuối sprint                        SLA: P1 fix <24h, P2 <5 ngày
```

### Quy Tắc Phối Hợp Hai Luồng

| Quy tắc | Mô tả |
|---------|-------|
| **Weekly triage** | PM triage feedback queue mỗi thứ 2. Assign vào Luồng B hoặc backlog v1.5 |
| **Escalation gate** | Issue >3 ngày estimate → đưa vào backlog, không phá sprint Luồng B |
| **Feature creep guard** | Luồng B KHÔNG nhận feature mới từ customers. Chỉ fix và polish hiện có. |
| **Convergence** | Hai luồng merge code vào main hằng ngày (feature flags cho features chưa GA) |
| **Demo chung** | Demo cuối sprint bao gồm cả hai luồng. Customers được invite demo Luồng B fixes. |

---

## 6. Đáp Ứng Key Requirements Đề Bài

| Key Requirement (topic.md) | Milestone | Sprint | Ghi Chú |
|---|---|---|---|
| **Asset inventory & classification** | v1 (T6) | S2–S4 core, S12 full | Google+M365 từ MVP. Slack+AWS tại S5. Shadow AI tại S9 (Track 2). |
| **AI-specific threat surface** | v1 (T6) | S7–S11 (Track 2) | Shadow AI governance S9. LLM DLP extension S8–S9. Deepfake + prompt injection S11. Full AI detection package có mặt trong v1. |
| **Access governance** | v1 (T6) | S5–S7 core | RBAC S5. Offboarding S6 (MVP). JIT S7. Access reviews S7. Shadow IT remediation S9. |
| **Compliance posture** | v1 (T6) — report-ready | S10–S11 | SOC 2 Type 1 + ISO 27001 report có thể xuất từ v1. Certification (audit verified) tại v2. |
| **Incident playbooks** | v1 (T6) | S6 (2 playbooks), S8–S9 (5 playbooks) | 5 playbooks, wizard UI, AWS Step Functions, non-security staff operable. |
| **Cost model** | v1.5 (T9) billing live | S13 code-ready, S18–S19 billing | Pricing tiers code-complete tại v1. Billing Stripe integration tại v1.5. Manual invoicing cho pilot tháng 1–6. |
| **Integrations** (Google, M365, Slack, QuickBooks...) | v1 (T6) | S2–S5 | Google+M365 từ MVP. Slack S5. AWS S5. QuickBooks → v2 backlog (không có AI security value đủ cao cho v1). |

> **Kết luận:** Tất cả 7 key requirements từ đề bài **đều có mặt trong v1 (tháng 6)**, đúng theo yêu cầu "v1 sau 5-6 tháng".

---

## 7. Giả Định Rủi Ro Nhất Cần Validate Sớm Nhất

### Rủi Ro #1 (Critical): Pilot Customers Không Thể Onboard <30 Phút

> **Assumption:** SME IT admin (không phải developer) có thể setup Google Workspace + M365 OAuth trong 30 phút bằng guided wizard.

**Tại sao đây là rủi ro nhất:**
- Toàn bộ MVP value prop phụ thuộc vào "first-value <30 min"
- Nếu onboarding thực tế mất 3 giờ (do M365 permission complexity), pilot program sụp đổ
- Competitors mất 2–4 giờ — nếu SMESec cũng vậy, không có differentiation

**Validate khi nào:** Sprint 2 end (W4) — test với 1–2 non-technical users thật trên Google Workspace tenant thật  
**Validate như thế nào:** Time-boxed usability test, no assistance from engineer  
**Go/No-go:** Nếu >45 phút → redesign wizard trước khi tiếp tục S3

### Top 5 Rủi Ro Theo Giai Đoạn

| # | Rủi Ro | Pha | Xác Suất | Tác Động | Mitigation |
|---|--------|-----|----------|----------|------------|
| 1 | OAuth wizard >30 min cho non-technical IT admin (M365) | MVP | Cao | Critical | Usability test W4. Prepared IT admin guide. Minimum-permission scopes. |
| 2 | ML Engineer không hire được trước W9 | Phase 2 | Trung bình | Cao | Bắt đầu tuyển W5. Contractor ML fallback nếu hire chậm. Tech Lead build SageMaker scaffold S5. |
| 3 | Pentest vendor LOI không ký trước W14 | Phase 2 | Thấp | Cao | PM lock calendar từ W8. Backup vendor list. |
| 4 | T1–T2 integration tại S11 bị delay >1 sprint | Phase 2 | Cao | Cao | Tech Lead full-time S11. API contract frozen S10. Fallback: manual trigger cho v1. |
| 5 | SOC 2 Type 2 evidence gap tại Month 9 review | Phase 3 | Thấp | Cao | Vanta weekly review từ W13. PM chịu trách nhiệm Vanta. Zero gap policy từ W22. |

---

## 8. Compliance Certification Timeline

```
Tháng 3 (W12):  Vanta account setup, evidence collection begins (silent)
Tháng 4 (W13):  Vanta CHÍNH THỨC active — SOC 2 control mapping bắt đầu
Tháng 5 (W21):  Pentest bắt đầu (6 tháng lead time từ LOI signing W14)
Tháng 6 (W26):  v1 LAUNCH
                  → SOC 2 Type 1 audit: scheduled với auditor
                  → Evidence collection W13→W26 = ~13 tuần (đủ cho Type 1)
Tháng 7 (W27):  ISO 27001 gap analysis bắt đầu
Tháng 8 (W33):  ISO 27001 Stage 1 audit (documentation review)
Tháng 9 (W38):  v1.5 LAUNCH
                  → SOC 2 Type 2 evidence W26→W38 = 12 tuần (cần 24 tuần total)
Tháng 10 (W41): ISO 27001 Stage 2 audit (implementation review)
Tháng 11 (W46): SOC 2 Type 2 audit fieldwork bắt đầu
                  → Evidence W26→W46 = 20 tuần (cần 24 tuần — ⚠️ tight)
                  → Mốc an toàn hơn: bắt đầu audit W48
Tháng 12 (W52): v2 LAUNCH
                  → SOC 2 Type 2 report issued ✅
                  → ISO 27001 certificate issued ✅
```

> ⚠️ **SOC 2 Type 2 timing note:** Để có đủ 6 tháng (24 tuần) observation window trước W52, evidence collection PHẢI bắt đầu không muộn hơn W26. Bắt đầu từ W13 như plan cho 10 tuần buffer nhưng chính thức SOC 2 Type 2 window tính từ W26 (ngày v1 launch — production environment).

---

## 9. External Dependencies & Hard Deadlines

| Deadline | Tuần | Mô Tả | Hậu Quả nếu trễ |
|----------|------|--------|----------------|
| Auth provider decision | W1D1 | Chọn Keycloak self-host vs Auth0 vs Cognito | Delay S1 → cascade tất cả sprints |
| Google test tenant có sẵn | W3 | Internal Google Workspace tenant cho S2 development | S2 không thể demo |
| Pilot customer #1 onboard | W8 | Ít nhất 1 real customer dùng staging | MVP không có validation thực tế |
| **ML Engineer #1 onboard** | **W9** | Bắt đầu tuyển W5 | Track 2 delay → AI features miss v1 |
| **Pentest vendor LOI signed** | **W14** | Hard deadline — lead time 7 tuần | Pentest không start W21 → v1 delay |
| **Vanta setup active** | **W13** | Cần 60+ ngày evidence cho SOC 2 Type 1 | SOC 2 Type 1 không đủ evidence tại v1 |
| Chrome Web Store submission | W29 | Browser extension cần 1–2 tuần review | Extension miss v1.5 launch |
| iOS App Store submission | W50 | App Store review 1–2 tuần | Mobile feature miss v2 window |
| SOC 2 Type 2 audit sign | W42 | Engage auditor firm | Audit không hoàn thành trước W52 |
| ISO 27001 Stage 2 audit | W45 | Certification 6–8 tuần sau audit | Certificate không có tại W52 |

---

## Summary Dashboard

```
MILESTONE OVERVIEW
══════════════════════════════════════════════════════════════════════

MVP    │ W12  │ T3  │ Asset inventory + offboarding + shadow IT
       │      │     │ Team: 6 FTE + contract
       │      │     │ Pilot: 3+ customers on staging (free)
───────┼──────┼─────┼──────────────────────────────────────────────
v1     │ W26  │ T6  │ All 7 key requirements delivered
       │      │     │ Team: 9 FTE
       │      │     │ Compliance: SOC 2 Type 1 audit scheduled
       │      │     │ Revenue: Starter ($399/mo) + Growth ($799/mo) tiers
───────┼──────┼─────┼──────────────────────────────────────────────
v1.5   │ W38  │ T9  │ AI detection v2 + AWS v1.1 + pilot feedback
       │      │     │ Team: 11 FTE (2-stream split: 65/35)
       │      │     │ Compliance: SOC 2 Type 2 evidence running
       │      │     │ Revenue: Business tier ($1,499/mo) live. Billing Stripe.
───────┼──────┼─────┼──────────────────────────────────────────────
v2     │ W52  │ T12 │ Compliance verified + Enterprise + BERT ML
       │      │     │ Team: 11.5 FTE (peak)
       │      │     │ Compliance: SOC 2 Type 2 ✅ + ISO 27001 ✅
       │      │     │ Revenue: Enterprise (custom) + usage-based option

══════════════════════════════════════════════════════════════════════
TEAM HEADCOUNT PROGRESSION
  T1–T3:  ██████░░░░░░  6 FTE
  T4–T6:  █████████░░░  9 FTE
  T7–T9:  ███████████░  11 FTE
  T10–T12:████████████  11.5 FTE (peak)
══════════════════════════════════════════════════════════════════════
```
