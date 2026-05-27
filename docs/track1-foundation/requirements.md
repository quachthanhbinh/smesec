# Track 1: Foundation & Governance -- Sprint Plan

**Date:** 2026-05-27
**Status:** Revised
**Timeline:** 6 thang -- 13 sprints x 2 tuan
**Team:** 5 FTE (1 Tech Lead/Architect, 2 Backend, 1 Frontend, 1 Flutter/Mobile)

---

## Summary

Track 1 xay dung **nen tang bao mat deterministic** cho SMESec -- khong phu thuoc ML/AI, accuracy gan 100%. Moi sprint ket thuc bang mot **deliverable co the test noi bo** truoc khi tiep tuc.

**Value proposition:** SMEs (10-500 nhan vien) co the quan ly tai san, kiem soat truy cap, va chung minh compliance (ISO 27001, GDPR, SOC 2) ma khong can dedicated security team.

---

## Scope

| Requirement (tu topic.md) | Track 1 |
|---------------------------|---------|
| Asset inventory and classification | Yes - Full |
| Access governance (least-privilege, offboarding, shadow IT) | Yes - Full |
| Compliance posture (ISO 27001, GDPR, SOC 2) | Yes - Foundation |
| Incident playbooks cho non-security staff | Yes - 5 core playbooks |
| Integration (Google Workspace, M365, Slack, AWS) | Yes - Full |
| AI-specific threat surface | No - Track 2 |
| Cost model / pricing | No - Workstream rieng |

---

## Sprint Overview

| Sprint | Tuan | Focus | End-of-Sprint Output |
|--------|------|-------|----------------------|
| S1 | W1-2 | Infrastructure and Auth | Dev env chay duoc, dang nhap SSO thanh cong |
| S2 | W3-4 | Google Workspace sync | Thay users + OAuth apps tu Google tenant |
| S3 | W5-6 | M365 sync + Dashboard v1 | Dashboard hien thi assets tu ca Google + M365 |
| S4 | W7-8 | Classification + Shadow IT alerts | IT admin phan loai duoc asset, nhan alert OAuth moi |
| S5 | W9-10 | Slack + AWS discovery + RBAC | Inventory bao gom Slack + AWS; RBAC ap dung duoc |
| S6 | W11-12 | Automated Offboarding | Offboard test user <5 phut, co report PDF |
| S7 | W13-14 | JIT Access | JIT request -> approve -> auto-revoke hoat dong end-to-end |
| S8 | W15-16 | Playbook engine + 3 playbooks | 3 playbooks chay duoc qua wizard UI |
| S9 | W17-18 | 2 playbooks con lai + Mobile app | 5 playbooks + mobile app hoan chinh |
| S10 | W19-20 | Compliance mapping + Evidence | Evidence tu dong thu thap, link duoc toi control |
| S11 | W21-22 | Compliance reports + Dashboard | Xuat bao cao ISO 27001 / GDPR / SOC 2 mot click |
| S12 | W23-24 | Dependency map + Lifecycle | Graph hien thi; zombie asset alert hoat dong |
| S13 | W25-26 | Hardening + Pen-test + Launch | Pen-test pass, 5-10 pilot customers onboard |

---

## Dieu chinh tu Debate (Solution Architect vs PM/Risk Manager)

> 2 vong debate (Round 1: doc lap; Round 2: phan hoi cheo) giua **Solution Architect 30 nam** (chuyen gia cybersecurity) va **PM/Risk Manager 30 nam**.

| # | Dieu chinh | Sprint | Ly do dong thuan |
|---|-----------|--------|-----------------|
| 1 | Them 3 acceptance criteria vao S1: tenant isolation CI test + secrets rotation policy + RDS Multi-AZ BCP stub | S1 | SA: cross-tenant leak = existential risk. PM: compliance blocker ISO 27001. |
| 2 | Flutter Eng bat dau mobile scaffold tu S1 (khong doi den S9) | S1 -> S8 | PM: 8 sprints idle la lang phi. SA: scaffold sach thi S9 scope giam kha thi hon. |
| 3 | S9 mobile scope giam: Android+iOS only, JIT + read-only (khong Desktop, khong incident wizard -> S10, FCM only -> APNs sang S11) | S9, S10 | SA: 1 Flutter Eng khong build full mobile trong 1 sprint. PM: demo milestone can sach. |
| 4 | Pen-test STARTS tai S12 (khong phai S13). S13 = remediation + launch only | S12, S13 | SA+PM dong y: neu Critical finding, can sprint de fix. Launch voi unfixed Critical = ISO 27001 violation. |
| 5 | Joint T1-T2 schema session Tuan 1: define ThreatDetectionEvent interface, freeze by end of S2 | S1 | SA: incompatible schemas khi S10 integrate = refactor ca 2 phia. PM: khong co ai flag dieu nay. |
| 6 | PM phai lua chon va ky LOI voi pen-test vendor truoc cuoi S8 (lead time 2-3 tuan) | S8, S11 | PM: external vendor khong the fast-track. SA: pen-test phai chay khi infra stable (prerequisites S11). |


## Sprint Details

---

### Sprint 1 -- Infrastructure and Auth (W1-2)

**Goal:** Dung nen tang ky thuat, team co the deploy va dang nhap duoc.

**Deliverables:**
- AWS infrastructure: VPC, ECS Fargate, RDS PostgreSQL, S3, EventBridge, Secrets Manager
- Multi-tenant data model (tenant_id tren moi bang, enforced tai API middleware)
- Keycloak SSO: dang nhap bang Google account va Microsoft account
- MFA bat buoc (TOTP)
- CI/CD pipeline: GitHub Actions -> staging auto-deploy; production manual gate

**End-of-Sprint Output:**
> Team engineer dang nhap vao web app bang tai khoan Google/M365 that, thay dashboard rong nhung hoat dong. Staging environment deploy thanh cong tu CI.


**Phan bo nhan luc:**

| Thanh vien | Vai tro | Nhiem vu chinh trong sprint | Capacity |
|-----------|---------|----------------------------|----------|
| Tech Lead | Go/AWS | VPC+ECS+RDS+S3 + Keycloak SSO config + code review | **90%** |
| Backend Eng | Go/Python | Multi-tenant DB schema + tenant isolation CI test + API scaffold | **90%** |
| Frontend Eng | React/Next.js | Design system + component library (early start) | **25%** |
| Flutter Eng | Flutter/Dart | Flutter project init + mobile scaffold architecture | **15%** |
| DevSecOps (5d) | Infra/SecOps | CI/CD (GitHub Actions->ECR->ECS) + Terraform + RDS Multi-AZ | **90%** |
| PM (5d) | PMO | Sprint kickoff + pilot outreach W1 + vendor API requests | **80%** |

> *Sprint workload: **30.5 / 50 person-days** (61% utilization)*
>
> **[DEBATE]** Them vao acceptance criteria: tenant isolation CI test + secrets rotation policy doc + RDS Multi-AZ BCP stub. Joint T1-T2 schema session trong Tuan 1.

---

### Sprint 2 -- Google Workspace Asset Sync (W3-4)

**Goal:** Ket noi Google Workspace tenant that, pull ve danh sach users va OAuth apps.

**Deliverables:**
- Google Admin SDK integration (OAuth 2.0 service account)
- Sync: users, groups, OAuth apps duoc authorize boi tung user
- Background job chay moi 15 phut (incremental sync)
- API endpoint `GET /api/assets?type=user&provider=google` de verify du lieu

**Acceptance Criteria:**
- Discover >90% users trong 1 gio dau
- Discover >90% OAuth apps trong 1 gio sau khi authorize
- Partial failure (rate limit, token loi) khong crash toan bo sync

**End-of-Sprint Output:**
> Chay `GET /api/assets?type=user&provider=google` tra ve list users that tu test tenant Google Workspace.


**Phan bo nhan luc:**

| Thanh vien | Vai tro | Nhiem vu chinh trong sprint | Capacity |
|-----------|---------|----------------------------|----------|
| Tech Lead | Go/AWS | Google OAuth design review + API architecture | **60%** |
| Backend Eng | Go/Python | Google Admin SDK + background sync + retry logic | **100%** |
| Frontend Eng | React/Next.js | Google sync status UI + asset list skeleton | **70%** |
| Flutter Eng | Flutter/Dart | Mobile navigation shell + Flutter architecture patterns | **20%** |
| DevSecOps (5d) | Infra/SecOps | CloudWatch logging + alarm setup + secrets baseline | **30%** |
| PM (5d) | PMO | Pilot customer outreach (active) + sprint ceremonies | **70%** |

> *Sprint workload: **30.0 / 50 person-days** (60% utilization)*

---

### Sprint 3 -- M365 Sync + Dashboard v1 (W5-6)

**Goal:** Them M365, build dashboard dau tien cho IT admin nhin thay toan bo asset.

**Deliverables:**
- Microsoft Graph API integration (OAuth 2.0 app registration)
- Sync: users, groups, OAuth apps, M365 licensed apps
- Web dashboard: bang asset inventory (filter, search, sort theo type/provider)
- CSV export
- User list voi trang thai account va provider

**End-of-Sprint Output:**
> IT admin dang nhap, thay bang hien thi tat ca users + OAuth apps tu ca Google va M365. Bam "Export CSV" download duoc file.
>
> **Milestone: Visibility checkpoint** -- team co the dung noi bo de kiem tra coverage.


**Phan bo nhan luc:**

| Thanh vien | Vai tro | Nhiem vu chinh trong sprint | Capacity |
|-----------|---------|----------------------------|----------|
| Tech Lead | Go/AWS | Microsoft Graph design + multi-provider abstraction | **50%** |
| Backend Eng | Go/Python | Microsoft Graph API + sync service + dashboard API | **100%** |
| Frontend Eng | React/Next.js | Dashboard v1: asset table, filter, sort, CSV export | **100%** |
| Flutter Eng | Flutter/Dart | Flutter basic nav + asset list + data model | **25%** |
| DevSecOps (5d) | Infra/SecOps | Monitoring setup + structured logging | **30%** |
| PM (5d) | PMO | Visibility checkpoint prep + pilot outreach | **70%** |

> *Sprint workload: **32.5 / 50 person-days** (65% utilization)*

---

### Sprint 4 -- Classification + Shadow IT Alerts (W7-8)

**Goal:** IT admin co the phan loai asset va nhan alert khi co OAuth app moi chua duoc approve.

**Deliverables:**
- Auto-classification theo rule: account type (admin / standard / service / contractor)
- Sensitivity levels: Restricted / Confidential / Internal / Public (default: Internal)
- Manual override per asset; bulk update qua CSV import
- Allow-list quan ly boi IT admin (approved / pending-review / blocked)
- Alert email + Slack khi OAuth app moi duoc detect va chua co trong allow-list
- Classification history logged (ai thay doi, khi nao, tu gi sang gi)

**End-of-Sprint Output:**
> IT admin authorize mot OAuth app moi tren Google -> trong 15 phut nhan Slack alert "New OAuth app detected: Notion -- Pending Review". Vao dashboard approve/block duoc.


**Phan bo nhan luc:**

| Thanh vien | Vai tro | Nhiem vu chinh trong sprint | Capacity |
|-----------|---------|----------------------------|----------|
| Tech Lead | Go/AWS | Classification logic design + allow-list architecture | **50%** |
| Backend Eng | Go/Python | Auto-classification engine + email/Slack alert service | **90%** |
| Frontend Eng | React/Next.js | Classification UI + allow-list management + bulk CSV import | **90%** |
| Flutter Eng | Flutter/Dart | Flutter asset list with classification badges | **25%** |
| DevSecOps (5d) | Infra/SecOps | Alert infra (SES/SNS) + notification routing | **20%** |
| PM (5d) | PMO | Legal review kickoff (T2 browser ext) + ceremonies | **60%** |

> *Sprint workload: **29.5 / 50 person-days** (59% utilization)*
>
> **[DEBATE]** PM bat dau legal review kickoff cho T2 browser extension (GDPR Article 13/88) trong sprint nay.

---

### Sprint 5 -- Slack + AWS Discovery + RBAC Engine (W9-10)

**Goal:** Mo rong inventory sang Slack va AWS; ap dung RBAC cho chinh SMESec platform.

**Deliverables:**
- Slack Admin API: sync users, channels, installed apps
- AWS integration: EC2, S3, RDS, Lambda, IAM users (via AWS Config + IAM API)
- RBAC engine (OPA/Rego): built-in roles Admin / Manager / Employee / Contractor / Service Account
- Policy evaluation <100ms; changes co hieu luc trong <1 phut
- Audit log moi access decision (who, resource, result, timestamp)

**End-of-Sprint Output:**
> Dashboard inventory hien thi ca Slack users + AWS resources (EC2, S3, IAM). User co role Employee khong truy cap duoc trang Admin Settings -- bi chan va ghi log.


**Phan bo nhan luc:**

| Thanh vien | Vai tro | Nhiem vu chinh trong sprint | Capacity |
|-----------|---------|----------------------------|----------|
| Tech Lead | Go/AWS | OPA RBAC policy design + AWS integration architecture | **80%** |
| Backend Eng | Go/Python | Slack Admin API + AWS Config/IAM + OPA/Rego engine | **100%** |
| Frontend Eng | React/Next.js | Slack/AWS assets dashboard + RBAC role management UI | **70%** |
| Flutter Eng | Flutter/Dart | Flutter provider filter views + permission-aware UI | **30%** |
| DevSecOps (5d) | Infra/SecOps | AWS Config + Terraform IAM + OPA deployment | **50%** |
| PM (5d) | PMO | Slack Enterprise Grid validation + sprint ceremonies | **50%** |

> *Sprint workload: **33.0 / 50 person-days** (66% utilization)*
>
> **[DEBATE]** Kiem tra Slack Enterprise Grid availability voi pilot customers truoc khi commit. Neu khong co, fallback sang Slack OAuth app discovery.

---

### Sprint 6 -- Automated Offboarding (W11-12)

**Goal:** Trigger offboarding -> toan bo access bi revoke trong <5 phut, co report de audit.

**Deliverables:**
- Webhook endpoint nhan trigger tu HR system (BambooHR, Workday) hoac manual trigger
- Parallel revocation via AWS Step Functions:
  - Google Workspace: suspend account, revoke OAuth tokens, transfer Drive ownership
  - Microsoft 365: disable account, revoke sessions, convert mailbox to shared
  - Slack: deactivate account, remove from channels
  - AWS: disable IAM user, revoke access keys
- Offboarding report tu dong: PDF + JSON (proof of revocation cho compliance)
- Failed revocation -> alert ngay lap tuc cho IT admin de xu ly manual

**Acceptance Criteria:**
- Tat ca access revoked trong <5 phut tu luc trigger
- Report PDF generated trong <1 phut sau khi hoan thanh

**End-of-Sprint Output:**
> Trigger offboard cho test user `bob@company.com` -> sau 5 phut: account bi suspend tren Google, M365, Slack; IAM user disabled tren AWS. Download PDF report co timestamp cua tung buoc.
>
> **Milestone: Offboarding checkpoint** -- demo duoc cho external stakeholder.


**Phan bo nhan luc:**

| Thanh vien | Vai tro | Nhiem vu chinh trong sprint | Capacity |
|-----------|---------|----------------------------|----------|
| Tech Lead | Go/AWS | Step Functions workflow design + parallel revocation arch | **70%** |
| Backend Eng | Go/Python | Webhook + Step Functions orchestration + 4-provider revocation | **100%** |
| Frontend Eng | React/Next.js | Offboarding trigger UI + report download + status tracking | **80%** |
| Flutter Eng | Flutter/Dart | Flutter offboarding status view + P0 push notifications | **35%** |
| DevSecOps (5d) | Infra/SecOps | Step Functions deploy + alert routing + DR test | **40%** |
| PM (5d) | PMO | External stakeholder demo prep + checkpoint coordination | **70%** |

> *Sprint workload: **34.0 / 50 person-days** (68% utilization)*

---

### Sprint 7 -- JIT Access (W13-14)

**Goal:** Privileged access chi cap khi can, tu dong expire -- khong can IT admin can thiep thu cong.

**Deliverables:**
- JIT request flow: user dien form (resource, duration, justification)
- Approver nhan notification Slack + email voi link approve/deny
- Neu approved: access cap ngay, auto-revoke sau het thoi gian
- Warning notification 10 phut truoc khi expire
- Toan bo workflow logged cho audit
- Emergency bypass kha dung (can post-incident review)

**Acceptance Criteria:**
- Access auto-revoked trong <1 phut sau khi expire

**End-of-Sprint Output:**
> User request JIT admin access vao AWS S3 trong 2 gio -> manager approve qua Slack -> access active -> sau 2 gio tu bi revoke -> log ghi day du.


**Phan bo nhan luc:**

| Thanh vien | Vai tro | Nhiem vu chinh trong sprint | Capacity |
|-----------|---------|----------------------------|----------|
| Tech Lead | Go/AWS | JIT workflow design + approval flow + threat model review | **60%** |
| Backend Eng | Go/Python | JIT request API + auto-revoke scheduler + Slack/email approvals | **100%** |
| Frontend Eng | React/Next.js | JIT form + approver dashboard + expiry timer UI | **90%** |
| Flutter Eng | Flutter/Dart | Flutter JIT: approve/deny from mobile | **45%** |
| DevSecOps (5d) | Infra/SecOps | Notification infra + emergency bypass audit trail | **20%** |
| PM (5d) | PMO | Pilot qualification + pen-test vendor selection starts | **50%** |

> *Sprint workload: **33.0 / 50 person-days** (66% utilization)*
>
> **[DEBATE]** PM target: 3 signed LOIs tu pilot customers truoc cuoi S8 (W16). Pen-test vendor phai duoc chon truoc cuoi sprint nay.

---

### Sprint 8 -- Playbook Engine + 3 Core Playbooks (W15-16)

**Goal:** IT admin va non-security staff co the chay incident playbook qua wizard UI.

**Deliverables:**
- Playbook engine: AWS Step Functions (stateful, fault-tolerant, resume after restart)
- Wizard UI (web): step-by-step, decision gates Yes/No, progress indicator, undo support
- Notification system: Email + Slack (P0: ca hai; P1: ca hai; P2/P3: Email only)
- 3 playbooks:

| Playbook | Trigger |
|---------|---------|
| Account Compromise | Suspicious login / impossible travel duoc report |
| Offboarding Emergency | Employee terminated ngay lap tuc |
| Shadow IT Detected | Unapproved OAuth app bi phat hien |

**End-of-Sprint Output:**
> Chay playbook "Account Compromise" cho test user qua wizard -- hoan thanh trong <10 phut, khong can giai thich bo sung. Moi buoc logged.


**Phan bo nhan luc:**

| Thanh vien | Vai tro | Nhiem vu chinh trong sprint | Capacity |
|-----------|---------|----------------------------|----------|
| Tech Lead | Go/AWS | Playbook engine architecture + Step Functions design | **80%** |
| Backend Eng | Go/Python | Playbook engine + wizard API + notifications + 3 playbooks | **100%** |
| Frontend Eng | React/Next.js | Wizard UI: steps, decision gates, progress indicator, undo | **100%** |
| Flutter Eng | Flutter/Dart | Flutter incident wizard (basic) + FCM Android push setup | **65%** |
| DevSecOps (5d) | Infra/SecOps | Playbook infra + P0/P1 notification routing | **40%** |
| PM (5d) | PMO | Playbook content review + user testing + pen-test LOI ky | **60%** |

> *Sprint workload: **39.5 / 50 person-days** (79% utilization)*
>
> **[DEBATE]** 2 signed pilot LOIs phai co truoc cuoi sprint nay. Pen-test vendor phai duoc ky LOI truoc cuoi S8 de dam bao booking cho S12.

---

### Sprint 9 -- 2 Playbooks con lai + Mobile App (W17-18)

**Goal:** Hoan thien playbook library va dua toan bo chuc nang len mobile.

**Deliverables:**
- 2 playbooks bo sung:

| Playbook | Trigger |
|---------|---------|
| Unauthorized Access | User truy cap resource khong co quyen |
| Inactive Account | Account khong dung >90 ngay |

- Flutter mobile app (iOS + Android + Desktop):
  - Asset inventory view
  - Incident wizard (chay playbook tren mobile)
  - JIT approval (approve/deny tu mobile)
  - Push notifications cho P0/P1 incidents

**End-of-Sprint Output:**
> Cai app tren iPhone -> nhan push notification ve incident -> approve JIT request -> chay playbook "Inactive Account" end-to-end tren mobile.
>
> **Milestone: Access Control checkpoint** -- demo day du Access Governance + Playbooks.


**Phan bo nhan luc:**

| Thanh vien | Vai tro | Nhiem vu chinh trong sprint | Capacity |
|-----------|---------|----------------------------|----------|
| Tech Lead | Go/AWS | Code review + mobile arch review + scope sign-off | **50%** |
| Backend Eng | Go/Python | 2 additional playbooks + mobile REST APIs | **70%** |
| Frontend Eng | React/Next.js | UI polish + dashboard bug fixes from testing | **30%** |
| Flutter Eng | Flutter/Dart | iOS + Android: asset inventory + JIT approval + FCM push | **100%** |
| DevSecOps (5d) | Infra/SecOps | FCM push infra + APNs certificate preparation | **30%** |
| PM (5d) | PMO | Access Control checkpoint demo + pilot customer status update | **70%** |

> *Sprint workload: **30.0 / 50 person-days** (60% utilization)*
>
> **[DEBATE]** DEBATE OUTCOME: Mobile scope giam. Sprint 9 = Android+iOS only (Desktop -> post-launch). Khong co incident wizard (-> S10). FCM Android only (APNs -> S11). Kha thi vi Flutter Eng da co scaffold tu S1.

---

### Sprint 10 -- Compliance Mapping + Evidence Collection (W19-20)

**Goal:** Moi feature trong Track 1 duoc map toi control cu the; evidence tu dong thu thap.

**Deliverables:**
- Mapping ISO 27001 / GDPR / SOC 2 controls -> features da build:

| Standard | Controls |
|---------|---------|
| ISO 27001 | A.8.1 (Asset Mgmt), A.8.2 (Classification), A.9.1 (Access Policy), A.9.2 (User Mgmt), A.9.4 (Access Review), A.12.4 (Logging) |
| GDPR | Art. 30 (Records), Art. 32 (Security), Art. 17 (Erasure), Art. 33 (Breach Notification) |
| SOC 2 | CC6.1 (Logical Access), CC6.2 (Provisioning), CC6.3 (Access Removal), CC7.2 (Monitoring) |

- Evidence auto-collection vao S3 (append-only, 7-year retention, encrypted KMS):
  - Daily asset inventory snapshot (CSV)
  - Access event logs (immutable)
  - Offboarding completion reports
  - Incident reports voi timeline

**End-of-Sprint Output:**
> Tro vao control ISO 27001 A.9.2 -> thay link toi 3 offboarding reports gan nhat nhu evidence. Moi evidence stored tren S3, khong the delete.


**Phan bo nhan luc:**

| Thanh vien | Vai tro | Nhiem vu chinh trong sprint | Capacity |
|-----------|---------|----------------------------|----------|
| Tech Lead | Go/AWS | Evidence schema + S3 Object Lock architecture | **60%** |
| Backend Eng | Go/Python | Compliance mapping + evidence collection + S3 append-only | **90%** |
| Frontend Eng | React/Next.js | Compliance mapping UI + evidence browser (ISO/GDPR/SOC2) | **80%** |
| Flutter Eng | Flutter/Dart | Flutter compliance view + APNs push + mobile incident wizard | **60%** |
| DevSecOps (5d) | Infra/SecOps | S3 Object Lock + KMS policies + 7yr retention config | **50%** |
| PM (5d) | PMO | Compliance validation + auditor framework review | **60%** |

> *Sprint workload: **34.5 / 50 person-days** (69% utilization)*
>
> **[DEBATE]** Flutter Eng them: APNs push (iOS) + incident wizard tren mobile (chuyen tu S9).

---

### Sprint 11 -- Compliance Reports + Dashboard (W21-22)

**Goal:** Mot click xuat bao cao compliance; IT admin va CEO thay duoc security posture realtime.

**Deliverables:**
- On-demand compliance reports: ISO 27001, GDPR, SOC 2 (PDF + JSON)
- Report bao gom: control status, evidence links, gaps (chua implement)
- Report generation <5 phut
- Compliance dashboard:
  - Control status: Implemented / Partial / Not Implemented
  - Audit Readiness Score (0-100)
  - Evidence collection progress
  - Recent compliance events (timeline)
- Quarterly access review workflow: tu dong tao review task, manager approve/revoke qua dashboard

**End-of-Sprint Output:**
> Bam "Generate ISO 27001 Report" -> 3 phut sau co file PDF voi Audit Readiness Score 82/100, danh sach gaps, va link evidence cho tung control.


**Phan bo nhan luc:**

| Thanh vien | Vai tro | Nhiem vu chinh trong sprint | Capacity |
|-----------|---------|----------------------------|----------|
| Tech Lead | Go/AWS | Report architecture + PDF generation design | **40%** |
| Backend Eng | Go/Python | Report engine (PDF+JSON) + compliance dashboard API + review WF | **90%** |
| Frontend Eng | React/Next.js | Compliance dashboard + Audit Readiness Score + report UI | **100%** |
| Flutter Eng | Flutter/Dart | Flutter compliance dashboard + mobile feature polish | **70%** |
| DevSecOps (5d) | Infra/SecOps | Monitoring dashboard + SLA alerting setup | **20%** |
| PM (5d) | PMO | UAT + compliance sign-off + pen-test vendor SELECTION DONE | **80%** |

> *Sprint workload: **35.0 / 50 person-days** (70% utilization)*
>
> **[DEBATE]** PM PHAI hoan tat lua chon pen-test vendor truoc cuoi sprint nay de dam bao lich booking cho S12.

---

### Sprint 12 -- Dependency Mapping + Lifecycle Tracking (W23-24)

**Goal:** IT admin hieu duoc "neu mat user X thi anh huong nhung gi" va phat hien zombie assets.

**Deliverables:**
- Dependency graph: user -> OAuth app -> cloud resource (dua tren OAuth scopes + IAM policies)
- Blast radius view: "Disabling user X affects Y resources"
- Asset lifecycle states: Discovered / Active / Inactive / Decommissioned
- Auto-flag asset inactive >90 ngay + alert
- Decommissioned assets retained trong audit log (khong xoa)

**End-of-Sprint Output:**
> Chon user `alice@company.com` -> thay graph: Alice -> 8 OAuth apps -> 3 S3 buckets, 1 RDS. Co 2 accounts khong dung 95 ngay duoc flag "Inactive -- review required".


**Phan bo nhan luc:**

| Thanh vien | Vai tro | Nhiem vu chinh trong sprint | Capacity |
|-----------|---------|----------------------------|----------|
| Tech Lead | Go/AWS | Dependency graph arch + pen-test kickoff + daily findings review | **70%** |
| Backend Eng | Go/Python | Dependency graph engine + lifecycle tracking + blast radius API | **90%** |
| Frontend Eng | React/Next.js | Dependency graph viz + lifecycle UI + blast radius view | **90%** |
| Flutter Eng | Flutter/Dart | Flutter dependency/lifecycle view | **40%** |
| DevSecOps (5d) | Infra/SecOps | External pen-test coordination + internal hardening + load test | **80%** |
| PM (5d) | PMO | Pen-test vendor management + docs + pilot final confirmation | **60%** |

> *Sprint workload: **36.0 / 50 person-days** (72% utilization)*
>
> **[DEBATE]** DEBATE OUTCOME: Pen-test STARTS tai sprint nay (chuyen tu S13). Prerequisites truoc khi bat dau pen-test: (1) API auth complete + deployed, (2) RDS chi accessible tu VPC, (3) S3 BlockPublicAccess = ON, (4) tat ca secrets trong Secrets Manager (khong co .env trong ECS tasks).

---

### Sprint 13 -- Hardening + Pen-test + Launch (W25-26)

**Goal:** Production-ready. Pass security audit. Onboard pilot customers.

**Deliverables:**
- External penetration test (fix all Critical/High findings)
- Load test: 500 assets, 50 concurrent users -- dashboard <2s, offboarding <5 min
- Zero cross-tenant data leakage (verified bang automated tests)
- Onboarding guide + documentation cho IT admin
- 5-10 pilot customers onboarded va chay duoc Track 1 end-to-end

**End-of-Sprint Output:**
> Pen-test report: 0 Critical, 0 High findings. 5 pilot customers da onboard, moi customer co it nhat 1 offboarding + 1 compliance report duoc tao.
>
> **Milestone: Beta Launch** -- Track 1 production-ready.


**Phan bo nhan luc:**

| Thanh vien | Vai tro | Nhiem vu chinh trong sprint | Capacity |
|-----------|---------|----------------------------|----------|
| Tech Lead | Go/AWS | Critical/High pen-test fixes + zero-trust verification | **100%** |
| Backend Eng | Go/Python | Security fixes + load test + multi-tenant isolation final check | **100%** |
| Frontend Eng | React/Next.js | UI security fixes + onboarding guide + documentation | **80%** |
| Flutter Eng | Flutter/Dart | Mobile security patches + iOS App Store submission | **65%** |
| DevSecOps (5d) | Infra/SecOps | Deployment hardening + pilot customer infra + load test exec | **100%** |
| PM (5d) | PMO | Pilot onboarding (5-10 customers) + beta launch coordination | **90%** |

> *Sprint workload: **44.0 / 50 person-days** (88% utilization)*
>
> **[DEBATE]** DEBATE OUTCOME: Sprint nay = pen-test remediation + launch. Pen-test da chay tu S12. Neu co Critical/High finding: fix va re-test TRONG sprint nay truoc khi launch. Khong co finding nao duoc defer sang post-launch.

---

## Non-Functional Requirements

| Category | Requirement |
|---------|------------|
| Performance | Asset discovery <5 phut cho 500-asset org; dashboard load <2s; search <1s |
| Scale | Toi da 10,000 assets va 500 users moi tenant |
| Reliability | Uptime >99.5%; failed jobs auto-retry toi da 5 lan voi exponential backoff |
| Security | Encryption at rest (AWS KMS) va in transit (TLS 1.3); audit logs immutable (S3 Object Lock) |
| Multi-tenancy | Zero cross-tenant data leakage -- enforced tai DB row level + API middleware |
| Data retention | Logs va evidence giu 7 nam |

---

## Out of Scope (Track 1)

| Feature | Deferred |
|---------|---------|
| AI threat detection (prompt injection, deepfake, DLP) | Track 2 |
| Network-level scanning (Nmap) | v1.1 |
| MDM integration (Intune, Workspace ONE) | v1.1 |
| Custom playbook builder (UI) | v1.1 |
| QuickBooks integration | v1.1 |
| Advanced compliance auto-remediation | v1.2 |
| Multi-region deployment | v1.2 |
| Voice call escalation (Twilio) | v1.1 |

---

## Related Documents

- [2-track-approach.md](../strategy/2-track-approach.md) -- Strategic overview
- [2026-05-27-2-track-decision-record.md](../strategy/2026-05-27-2-track-decision-record.md) -- Decision record
- [Track 2 Requirements](../track2-ai-detection/requirements.md) -- AI detection R&D plan
