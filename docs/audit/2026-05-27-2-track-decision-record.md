# SMESec Platform - 2-Track Approach Decision Record

**Date:** 2026-05-27  
**Status:** Approved  
**Context:** Phân tích lại thiết kế ban đầu và quyết định chia thành 2 tracks song song

---

## Executive Summary

Session này review lại thiết kế ban đầu (2026-05-26) và phát hiện **critical risk**: AI threat detection chỉ đạt 85% accuracy, trong khi 15% còn lại có thể chứa các threats nghiêm trọng hoặc false positives gây frustration.

**Quyết định:** Chia development thành 2 tracks song song:
- **Track 1 (Foundation):** Asset + Access + Playbooks - accuracy ~100%, launch sau 6 tháng
- **Track 2 (AI Detection):** R&D để đạt >95% accuracy, pilot sau 6 tháng, chỉ launch khi đạt quality standards

---

## 1. Vấn Đề Với Thiết Kế Ban Đầu

### 1.1 Phân Tích Thiết Kế 2026-05-26

**Documents reviewed:**
- [2026-05-26-sme-ai-security-platform-design.md](../superpowers/specs/2026-05-26-sme-ai-security-platform-design.md)
- [2026-05-26-sme-platform-decision-record.md](../superpowers/specs/2026-05-26-sme-platform-decision-record.md)
- [2026-05-26-sme-platform-synthesis-en.md](../others/2026-05-26-sme-platform-synthesis-en.md)

**Phát hiện:**
1. Thiết kế tập trung quá nhiều vào **Requirement 2 (AI Threat Surface)**
2. AI detection target: 85% precision, 15% false positive
3. Các requirements khác (Asset, Access, Compliance) chưa được detail đầy đủ

### 1.2 Critical Risk Analysis

**15% gap trong AI detection có thể chứa:**

**False Negatives (bỏ sót threats thật):**
- Deepfake fraud bypass detection → CEO voice clone thành công → wire transfer $500K
- Prompt injection không phát hiện → attacker extract sensitive data từ LLM
- Data leakage không block → customer PII leaked to ChatGPT → GDPR violation

**False Positives (block nhầm công việc hợp lệ):**
- Block nhầm legitimate AI usage → employees frustrated → productivity giảm
- Alert fatigue → IT manager ignore alerts → miss real threats
- Over-blocking → employees tìm cách bypass system → shadow IT tăng

**Hậu quả:**
- **Mất lòng tin khách hàng:** "Hệ thống không bắt được deepfake" hoặc "Block nhầm công việc quan trọng"
- **Không ai dám dùng:** SMEs sẽ không adopt nếu không tin tưởng 100%
- **Reputation damage:** Một incident nghiêm trọng có thể kill product

### 1.3 Root Cause

**Tại sao thiết kế ban đầu có risk cao:**
1. **AI detection là hardest problem:** ML models inherently probabilistic, không thể đạt 100%
2. **Thiếu validation plan:** Không có clear plan để validate accuracy trước khi launch
3. **All-or-nothing approach:** Nếu AI detection fail, toàn bộ product fail
4. **Scope quá rộng:** Cố gắng solve tất cả problems cùng lúc trong 6 tháng

---

## 2. Brainstorming Session: Tìm Giải Pháp

### 2.1 User Input

**Câu hỏi từ user:**
> "Tôi thấy nội dung đó chỉ tập trung giải quyết cho mục 2 ai-threat-surface, nhưng giải quyết ở mức 85%, nhưng nếu các vấn đề critical lại nằm ở 15% không giải quyết được thì sẽ không ai có thể tin tưởng sử dụng công cụ này. Tôi nghĩ nên chia cách tiếp cận thành 2 hướng là: hướng 1 bao gồm mục 1 và 3, hướng 2 là mục 2, sau đó triển khai tiếp các yêu cầu khác."

**Key insights:**
- User nhận ra AI detection là highest risk
- Đề xuất tách riêng AI detection để R&D kỹ hơn
- Ưu tiên Foundation (Asset + Access) vì có high confidence

### 2.2 Clarifying Questions & Answers

**Q1: Về 15% critical cases - lo ngại cụ thể về false negatives hay false positives?**
- **A:** Cả hai

**Q2: Timeline mới - launch Hướng 1 trước hay develop song song?**
- **A:** Chia thành 2 team nhỏ để develop song song, nhưng sẽ cùng trong 1 sản phẩm chính

**Q3: Scope Hướng 1 - có cần bao gồm Incident Playbooks, Cost Model, Integrations?**
- **A:** Nếu cần để thực hiện hướng 1 thì cần bao gồm vào. Về cost model sẽ bàn riêng sau khi hoàn thành xong các tính năng

**Q4: Về compliance - Hướng 1 có đủ để đạt ISO 27001/GDPR/SOC2 không?**
- **A:** Cần thiết phần nào thì phải bổ sung phần đó vào

### 2.3 Options Considered

**Option A: Continue với thiết kế ban đầu**
- ❌ Rejected: Risk quá cao, 15% gap không acceptable

**Option B: Defer AI detection hoàn toàn, focus vào Foundation**
- ❌ Rejected: AI detection là core value proposition, không thể bỏ

**Option C: 2-track approach (SELECTED)**
- ✅ Track 1: Foundation (high confidence, launch sau 6 tháng)
- ✅ Track 2: AI Detection (R&D riêng, chỉ launch khi đạt >95% accuracy)
- ✅ Develop song song, integrate sau

---

## 3. Quyết Định: 2-Track Approach

### 3.1 Track 1: Foundation & Governance

**Scope:**
1. **Asset Inventory & Classification** (Requirement 1)
2. **Access Governance** (Requirement 3)
3. **Incident Playbooks** (Requirement 4 - subset cho access incidents)
4. **Compliance Foundation** (ISO 27001, GDPR, SOC 2 - Asset + Access controls)
5. **Core Integrations** (Requirement 6 - Google, M365, Slack, AWS)

**Why high confidence:**
- Proven technology: OAuth 2.0, RBAC, API integrations
- Deterministic logic (không phụ thuộc ML)
- Có thể test 100% với automated tests
- Immediate value: visibility + control

**Team:** 5 FTE
- 1 Tech Lead / Architect
- 2 Backend Engineers (Go + Python)
- 1 Frontend Engineer (React/Next.js)
- 1 Flutter Engineer (Mobile/Desktop)

**Timeline:** 6 tháng → Launch-ready

### 3.2 Track 2: AI Threat Detection

**Scope:**
1. **Prompt Injection Detection** (target >95% precision)
2. **LLM Data Leakage Prevention** (target 0 false negatives on critical data)
3. **Deepfake Detection** (target >90% detection rate)
4. **Shadow AI Discovery** (target >95% discovery rate)

**Why needs R&D:**
- ML models non-deterministic
- Cần pilot data để tune thresholds
- False positive/negative trade-offs phức tạp

**Team:** 3 FTE
- 1 ML Engineer / Security Researcher
- 1 Backend Engineer (Python/FastAPI)
- 1 Frontend Engineer (Browser Extension + Desktop Agent)

**Timeline:** 6 tháng → Pilot-ready
- **Decision gate:** Nếu đạt >95% precision → merge vào product
- **If not:** Continue iterate, không launch sản phẩm chưa hoàn thiện

### 3.3 Shared Resources

**2 FTE:**
- 1 Product Manager / Security Analyst (coordinate both tracks)
- 1 DevSecOps / QA (CI/CD, testing, infrastructure)

**Total team:** 10 FTE

---

## 4. Key Decisions Made

### 4.1 Team Size
**Decision:** TBD - sẽ thảo luận riêng sau

**Rationale:** Cần assess available resources và budget trước khi finalize

### 4.2 Pilot Customers

**Decision:** ⚠️ **CRITICAL** - Chưa có sẵn SMEs để pilot

**Action required:**
- Cần tìm 2-3 SMEs sẵn sàng pilot Track 2 trong Month 4-6
- **Criteria:**
  - 50-200 employees (sweet spot cho pilot)
  - Đang dùng AI tools (ChatGPT, Copilot, etc.)
  - Có IT manager/admin có thể collaborate
  - Sẵn sàng share feedback và telemetry data
- **Timeline:** Cần identify pilot customers trước Month 4

**Rationale:** Không có pilot customers = không thể validate AI accuracy trong real-world conditions

### 4.3 Budget Track 2

**Decision:** ✅ Approved

**Budget items:**
- Deepfake detection API ($3K-5K/year)
- ML training infrastructure (SageMaker)
- Labeled datasets (nếu cần mua)

**Rationale:** AI detection cần specialized tools và infrastructure để đạt accuracy cao

### 4.4 Quality Standards

**Decision:** ⚠️ **QUALITY FIRST**

**Policy:**
- Nếu Track 2 chỉ đạt 90% accuracy sau 6 tháng → **KHÔNG launch**
- Tiếp tục iterate cho đến khi đạt >95% precision, <5% false positive
- **Nguyên tắc:** Không đưa sản phẩm chưa hoàn thiện ra thị trường
- Track 1 có thể launch độc lập nếu Track 2 chưa sẵn sàng

**Rationale:**
- Reputation risk quá cao nếu launch với accuracy thấp
- Better to delay than to damage brand trust
- Track 1 vẫn có value proposition mạnh (Asset + Access governance)

---

## 5. Dependencies Giữa 2 Tracks

### 5.1 Track 1 → Track 2 (Foundation provides to AI)

| Foundation Component | AI Detection Usage |
|---------------------|-------------------|
| Asset Inventory | AI detection cần biết user roles, device context để risk scoring |
| Access Governance | AI alerts cần trigger offboarding workflows |
| Incident Playbooks | AI threats cần automated response actions |
| Integrations | AI detection cần Google/M365 APIs để revoke access |

### 5.2 Track 2 → Track 1 (AI provides to Foundation)

| AI Detection Output | Foundation Usage |
|--------------------|------------------|
| AI Threat Events | Feed vào incident playbook engine |
| Risk Scores | Enrich asset classification (high-risk users/devices) |
| Shadow AI Discovery | Feed vào access governance (unapproved apps) |

### 5.3 Integration Architecture

```
[Track 2: AI Detection Service]
         │ Publishes events
         ▼
[EventBridge / SQS]
         │ Routes events by severity
         ▼
[Track 1: Playbook Engine]
         │ Executes response workflow
         ▼
[Track 1: Access Governance]
         │ Revokes access, locks accounts
         ▼
[Track 1: Integration Hub]
         │ Calls Google/M365/Slack APIs
```

**Key design principle:** Loose coupling via event-driven architecture
- Track 2 có thể develop và test độc lập
- Track 1 có thể launch mà không cần Track 2
- Integration point rõ ràng (EventBridge)

---

## 6. Compliance Requirements cho Track 1

### 6.1 ISO 27001 Controls

Track 1 đủ để đạt các controls sau:

| Control | Requirement | Track 1 Implementation |
|---------|-------------|----------------------|
| **A.8.1** Asset Management | Inventory of assets | Asset discovery + classification engine |
| **A.8.2** Information Classification | Classify by sensitivity | Criticality levels (Critical/High/Medium/Low) |
| **A.9.1** Access Control Policy | Least privilege | RBAC + JIT access |
| **A.9.2** User Access Management | Provisioning/deprovisioning | Automated offboarding workflows |
| **A.9.4** Access Review | Periodic review | Quarterly access review reports |
| **A.12.4** Logging & Monitoring | Audit trails | All access events logged to S3 |

### 6.2 GDPR Requirements

| Article | Requirement | Track 1 Implementation |
|---------|-------------|----------------------|
| **Art. 30** Records of Processing | Data inventory | Asset inventory includes data assets |
| **Art. 32** Security Measures | Access controls | RBAC + MFA + encryption |
| **Art. 17** Right to Erasure | Delete personal data | Offboarding workflow includes data deletion |
| **Art. 33** Breach Notification | 72-hour reporting | Incident playbooks include notification workflows |

### 6.3 SOC 2 Trust Services Criteria

| Criteria | Requirement | Track 1 Implementation |
|----------|-------------|----------------------|
| **CC6.1** Logical Access | Restrict access | RBAC + least privilege |
| **CC6.2** Access Provisioning | Timely provisioning/deprovisioning | Automated offboarding <5 minutes |
| **CC6.3** Access Removal | Remove access when no longer needed | JIT access auto-expires |
| **CC7.2** System Monitoring | Monitor security events | All access events logged + alerted |

**Conclusion:** Track 1 alone đủ để đạt ISO 27001, GDPR, SOC 2 compliance cho Asset + Access Management domain. Track 2 (AI detection) là value-add, không bắt buộc cho compliance.

---

## 7. Success Criteria

### 7.1 Track 1 Launch Criteria (Month 6)

**Technical:**
- ✅ Asset discovery coverage >95% (all devices, accounts, SaaS apps)
- ✅ Offboarding completion time <5 minutes (all platforms)
- ✅ Shadow IT detection rate >90% (OAuth apps)
- ✅ Zero cross-tenant data leakage (security audit passed)
- ✅ All services deploy to AWS production
- ✅ CI/CD pipeline green

**Compliance:**
- ✅ Compliance reports generated (ISO 27001, GDPR, SOC 2)
- ✅ Evidence collection automated
- ✅ Audit trails complete

**Business:**
- ✅ 5-10 pilot customers onboarded successfully
- ✅ Customer feedback validates value proposition
- ✅ Pricing model accepted

### 7.2 Track 2 Pilot Criteria (Month 6)

**Accuracy Metrics:**
- ✅ Prompt injection precision >95% (on test dataset)
- ✅ DLP false negative rate <1% (on critical data: credit cards, passwords)
- ✅ Deepfake detection >90% (on benchmark dataset)
- ✅ False positive rate <5% (on pilot customer data)

**Validation:**
- ✅ 2-3 pilot customers validate accuracy in production
- ✅ Real-world telemetry collected
- ✅ Threshold tuning completed

**Decision Gate:**
- **If metrics met:** Merge into main product, full launch
- **If metrics not met:** Continue iterate, Track 1 launches independently

---

## 8. Risk Analysis

### 8.1 Track 1 Risks

| Risk | Probability | Impact | Mitigation |
|------|------------|--------|------------|
| Integration API changes | Medium | Medium | Daily integration tests, version pinning |
| Performance issues at scale | Medium | Medium | Load testing, caching, rate limiting |
| Compliance audit failure | Low | High | External audit review at Month 5 |
| Team skill gaps | Low | Medium | Hire experienced architect early |

### 8.2 Track 2 Risks

| Risk | Probability | Impact | Mitigation |
|------|------------|--------|------------|
| AI accuracy insufficient | **High** | **CRITICAL** | Early validation gates (Week 6, 12, 18) |
| False positives frustrate users | Medium | High | Pilot with friendly customers, tune thresholds |
| Deepfake detection too slow | Medium | Medium | Use vendor APIs, optimize inference |
| Cannot find pilot customers | Medium | High | Start outreach in Month 1, not Month 4 |
| Labeled datasets insufficient | Medium | Medium | Budget for dataset purchase, synthetic data generation |

### 8.3 Cross-Track Risks

| Risk | Probability | Impact | Mitigation |
|------|------------|--------|------------|
| Integration complexity | Low | Medium | Clear event schema, integration tests |
| Resource contention | Low | Low | Separate teams, shared DevOps only |
| Timeline misalignment | Low | Medium | Weekly sync meetings, shared roadmap |

---

## 9. Timeline: 6 Months Parallel Development

### Track 1: Foundation (Launch-Ready)

```
Month 1: FOUNDATION
├── AWS infrastructure + VPC + security baseline
├── Tenant model + auth (Keycloak/SSO)
├── CI/CD pipeline
└── Integration skeletons (Google, M365, Slack)

Month 2: ASSET INVENTORY
├── Asset discovery engine (cloud, SaaS, devices)
├── Classification framework (criticality, sensitivity)
├── Dependency mapping
└── Web Dashboard: Asset inventory view

Month 3: ACCESS GOVERNANCE
├── RBAC engine (policy evaluation with OPA)
├── JIT access workflows (request/approve/revoke)
├── Automated offboarding (parallel API calls)
└── Shadow IT detection (OAuth app inventory)

Month 4: INCIDENT PLAYBOOKS
├── Playbook engine (Step Functions workflows)
├── Pre-built playbooks (account compromise, shadow IT)
├── Mobile/Desktop app: Incident wizards
└── Notification system (Slack, email, push)

Month 5: COMPLIANCE & INTEGRATIONS
├── Compliance control mappings (ISO 27001, GDPR, SOC 2)
├── Evidence collection automation
├── Audit report generation
└── Full integration testing (Google, M365, Slack)

Month 6: HARDENING & LAUNCH
├── Security hardening + penetration testing
├── Performance optimization
├── Documentation + training materials
└── Beta launch with 5-10 pilot customers
```

### Track 2: AI Detection (Pilot-Ready)

```
Month 1-2: RESEARCH & PROTOTYPING
├── Literature review (OWASP LLM Top 10, research papers)
├── Dataset collection (prompt injection, DLP test cases)
├── Baseline models (BERT, GPT-based classifiers)
└── Accuracy benchmarking (precision, recall, F1)

Month 3-4: CORE DETECTION ENGINE
├── Prompt injection detection (rules + ML)
├── DLP engine (PII, credentials, IP patterns)
├── Risk scoring algorithm (0-100 scale)
└── Browser extension v1 (prompt interceptor)

Month 5: DEEPFAKE & ADVANCED FEATURES
├── Deepfake detection API integration (Sensity, Reality Defender)
├── Dynamic redaction (mask sensitive data)
├── Desktop agent v1 (clipboard monitoring)
└── Behavioral analysis (user baseline)

Month 6: VALIDATION & TUNING
├── Pilot with 2-3 customers (collect real-world data)
├── False positive/negative analysis
├── Threshold tuning (optimize for SME use cases)
└── Decision: Launch or iterate based on metrics
```

---

## 10. Next Steps

### Immediate Actions (Week 1)

1. ✅ **Approve 2-track strategy** (this document)
2. ⏳ **Create detailed requirements for Track 1:**
   - Asset Inventory & Classification spec
   - Access Governance spec
   - Incident Playbooks spec
   - Compliance mapping spec
3. ⏳ **Create research plan for Track 2:**
   - AI detection accuracy targets
   - Dataset requirements
   - Validation methodology
   - Benchmark selection
4. ⏳ **Start pilot customer outreach:**
   - Define ideal customer profile
   - Create outreach materials
   - Identify 5-10 prospects
   - Target: 2-3 confirmed by Month 4

### Month 1 Actions

5. ⏳ **Finalize team structure:**
   - Assess available resources
   - Identify skill gaps
   - Hire if needed (especially ML Engineer for Track 2)
6. ⏳ **Set up infrastructure:**
   - AWS accounts + VPC
   - CI/CD pipeline
   - Monitoring + alerting
7. ⏳ **Kick off both tracks:**
   - Track 1: Start AWS foundation
   - Track 2: Start literature review + dataset collection

---

## 11. Open Questions

### 11.1 Team & Resources

1. **Team size:** 10 FTE đủ không? Cần hire thêm không?
2. **ML expertise:** Có sẵn ML Engineer với security background không?
3. **Budget:** Total budget cho 6 tháng là bao nhiêu?

### 11.2 Pilot Customers

4. **Outreach strategy:** Làm sao tìm và qualify pilot customers?
5. **Incentives:** Có offer gì để attract pilot customers? (free tier, early access, etc.)
6. **Legal:** Cần NDA hoặc pilot agreement không?

### 11.3 Technical

7. **Deepfake API vendor:** Chọn Sensity hay Reality Defender? Criteria?
8. **ML infrastructure:** SageMaker hay self-hosted? Cost trade-offs?
9. **Dataset licensing:** Có thể dùng public datasets hay cần mua commercial?

---

## 12. Lessons Learned

### 12.1 From Previous Design (2026-05-26)

**What went wrong:**
- Tried to solve too many problems at once
- Underestimated AI detection complexity
- No clear validation plan for accuracy
- All-or-nothing approach (high risk)

**What to keep:**
- Hybrid architecture (build core, integrate commodity)
- Event-driven design
- Multi-tenancy model
- AWS-native infrastructure

### 12.2 Principles for This Design

**Risk management:**
- Separate high-risk (AI) from low-risk (Foundation) components
- Validate early and often (Week 6, 12, 18 gates)
- Quality first - no launch until metrics met

**Incremental value:**
- Track 1 delivers value independently
- Track 2 is additive, not blocking
- Can launch Track 1 while iterating Track 2

**Customer-centric:**
- Pilot customers involved early (Month 4)
- Real-world validation before launch
- Feedback loop built into timeline

---

## Document History

- **2026-05-27:** Initial version - decision record for 2-track approach
- **Session context:** Review of 2026-05-26 design, identified critical risk in AI detection accuracy, decided to split into 2 parallel tracks
- **2026-05-27 (update):** Added Section 13 -- Risk Assessment Review of sprint plans for both tracks. Identified 6 CRITICAL risks, 5 HIGH risks, 8 MEDIUM risks. Recommended revised timeline: Track 1 = 15 sprints (7.5 months), Track 2 = 14 sprints (7 months).

---

## Related Documents

- [2-track-approach.md](2-track-approach.md) - Strategic overview of 2-track approach
- [2026-05-26-sme-ai-security-platform-design.md](../superpowers/specs/2026-05-26-sme-ai-security-platform-design.md) - Original design (superseded)
- [2026-05-26-sme-platform-decision-record.md](../superpowers/specs/2026-05-26-sme-platform-decision-record.md) - Original decision record (superseded)


---

## 13. Risk Assessment Review -- Sprint Plan (2026-05-27)

**Reviewer:** Risk Management & Project Management (30 years experience)
**Review scope:** Track 1 Sprint Plan (13 sprints), Track 2 Sprint Plan (13 sprints)
**Date:** 2026-05-27

---

### 13.1 Phuong phap danh gia

Review tung sprint theo 3 chieu: **(1) Sprint scope co kha thi trong 2 tuan?** **(2) Dependencies co duoc giai quyet dung thu tu?** **(3) Risk cua tung sprint neu bi delay?**

---

### 13.2 Rui ro Cap Do CRITICAL (Phai xu ly truoc khi kick off)

---

#### CRIT-01: Track 1 Sprint 1 -- Scope qua lon cho 2 tuan

**Van de:** 1 sprint yeu cau: VPC + ECS + RDS + S3 + EventBridge + Secrets Manager + Multi-tenant data model + Keycloak SSO (Google + M365) + MFA + CI/CD pipeline.

**Kinh nghiem thuc te:** Keycloak setup dung cach (HA mode, realm config, client scopes, token policies) mat toi thieu 3-5 ngay cua 1 senior engineer. Multi-tenant data model neu sai se gay cross-tenant data leakage -- la loi nghiem trong nhat co the xay ra, khong the voi vang. AWS infrastructure co the setup nhanh hon nhung van can review boi TL truoc khi merge.

**Hau qua neu bi tre:** Day la foundation cua ca he thong. Moi sprint sau deu phu thuoc vao Sprint 1. Delay 1 tuan o day = delay 1 tuan cho ca 12 sprints con lai.

**Khuynh nghi:** Tach Sprint 1 thanh 2 sprints:
- **Sprint 1a (W1-2):** AWS infrastructure + CI/CD + multi-tenant data model (core DB schema)
- **Sprint 1b (W3-4):** Keycloak SSO + MFA + auth middleware

**Tac dong:** Tang tong timeline len 14 sprints.

---

#### CRIT-02: Track 1 Sprint 5 -- 3 integrations lon + RBAC trong 1 sprint

**Van de:** Sprint 5 yeu cau: Slack Admin API + AWS Config/IAM + OPA RBAC engine + audit logging -- tat ca trong 2 tuan.

**Kinh nghiem thuc te:** OPA/Rego policy engine de tich hop dung la 1 sprint rieng (phai hieu Rego language, viet policies, test edge cases, setup cache invalidation). Slack Admin API yeu cau Enterprise Grid subscription -- neu khach hang chua co, phai wait. AWS IAM cross-account access co trust policy phuc tap, thuong mat 1-2 ngay debug permissions.

**Hau qua neu bi tre:** Sprint 6 (Offboarding) phu thuoc truc tiep vao Slack + AWS integration. Neu S5 bi delay, S6 bi block.

**Khuyen nghi:** Tach thanh 2 sprints:
- **Sprint 5a:** Slack integration + AWS discovery
- **Sprint 5b:** OPA RBAC engine + audit logging

---

#### CRIT-03: Track 1 Sprint 9 -- Flutter mobile app trong 1 sprint

**Van de:** 1 Flutter engineer phai build trong 2 tuan: asset inventory view + incident wizard (tat ca 5 playbooks) + JIT approval + push notifications (iOS + Android + Desktop).

**Kinh nghiem thuc te:** Mot mobile app co push notifications va offline support tren iOS/Android can it nhat 6-8 tuan de build, test, va submit. Push notification setup tren iOS (APNs certificates, provisioning profiles) va Android (FCM) mat 2-3 ngay rieng. App Store / Play Store review mat 1-3 ngay neu khong co van de, co the lon hon.

**Hau qua:** Flutter engineer se deliver mot app thieu tinh nang, chua duoc test ky, co the anh huong den demo voi pilot customers.

**Khuyen nghi:** Tach mobile app thanh 3 sprints (S9, S10, S11). Day la thay doi scope lon nhat can xu ly.

---

#### CRIT-04: Track 1 Sprint 13 -- Pen-test va Launch trong cung 1 sprint

**Van de:** Penetration test + fix findings + load test + documentation + onboard 5-10 pilot customers trong 2 tuan.

**Kinh nghiem thuc te:** Penetration test phai schedule truoc voi external vendor 2-3 tuan. Ban than viec test mat 5-7 ngay. Neu phat hien Critical/High finding (rat co the xay ra lan dau), team can 3-7 ngay de fix va re-test. Neu pen-test va fix deu nam trong Sprint 13, khong co thoi gian nen neu gap van de.

**Hau qua:** Launch voi unresolved Critical findings (security risk), hoac delay launch (no plan B).

**Khuyen nghi:** Schedule pen-test vao Sprint 12 (W23-24). Sprint 13 chi la remediation + launch. Bo Dependency Mapping ra khoi S12, dua xuong post-launch (v1.1) -- day la feature "nice to have", khong phai blocker cho launch.

---

#### CRIT-05: Track 2 -- 1 ML Engineer la Single Point of Failure

**Van de:** Toan bo Track 2 phu thuoc vao 1 ML Engineer: data collection, BERT fine-tuning, NER model, deepfake integration, accuracy tuning, pilot support. Neu nguoi nay nghi om, nghi phep, hoac roi cong ty, Track 2 dung lai.

**Kinh nghiem thuc te:** ML work khong the transfer nhanh cho nguoi khac. Model training, hyperparameter tuning, va debugging la kien thuc "in-head" rat kho document. Trong 30 nam lam viec, toi da thay nhieu du an that bai vi single point of failure nhu the nay.

**Hau qua:** Neu mat ML Engineer sau Gate 2 (tuan 12), Track 2 bi delay it nhat 4-6 tuan de onboard nguoi moi.

**Khuyen nghi:** Thue them 1 ML Engineer hoac train Backend Engineer (Python) de co backup capacity. It nhat, Backend Engineer phai co kha nang chay lai training jobs va interpret metrics.

---

#### CRIT-06: Track 2 Gate 1 tai W6 -- Qua som voi qua nhieu yeu cau

**Van de:** Gate 1 (W6) yeu cau: prompt injection precision >90% va DLP >95% -- nhung ML Engineer chi co 4 tuan thuc su (S2 + S3) sau khi hoan thanh research va infra setup (S1).

**Kinh nghiem thuc te:** Fine-tuning BERT tren custom dataset, evaluate, iterate, va reach >90% precision tren novel injection attacks thuong mat 3-4 iterations. 4 tuan la qua ngan de co cai gi to ra "production-credible". Neu dat duoc 90% chi tren OWASP known patterns, dieu nay khong co nghia la 90% tren real-world novel attacks.

**Hau qua:** Gate 1 "pass" tren test data co the tao false confidence. Real-world performance se thap hon dang ke (benchmark-to-production gap), dan den that bai o Gate 3 hoac Gate 4.

**Khuyen nghi:** Doi Gate 1 sang W8 (sau Sprint 4), thu them 2 sprints cho ML iteration. Dong thoi, them yeu cau adversarial test (test tren patterns KHONG co trong training data) vao Gate 1.

---

### 13.3 Rui ro Cap Do HIGH (Phai co plan xu ly truoc khi bat dau sprint lien quan)

---

#### HIGH-01: Khong co buffer sprints -- 0% contingency

**Van de:** Ca 2 tracks deu co 13 sprints lien tuc, khong co sprint nao lam buffer. Trong project management, khong co du an phan mem nao chay dung 100% ke hoach trong 6 thang.

**Lich su thuc te:** Theo Standish Group Chaos Report, chi 31% IT projects deliver dung han. Voi 10 FTE, 6 thang, tich hop voi 4+ external APIs, xac suat delay it nhat 1 sprint la >80%.

**Hau qua:** Khong co buffer = moi delay nho tich luy thanh delay lon o cuoi. Sprint 13 (launch) se bi cap du lai.

**Khuyen nghi:** Them 1 buffer sprint sau S6 (T1) va 1 buffer sprint sau S9 (T2). Tong: 15 sprints (~7.5 thang). Dung buffer sprint de: fix bugs tu sprint truoc, re-test integration, prepare documentation.

---

#### HIGH-02: Google Admin SDK -- Customer Consent Mat Thoi Gian

**Van de:** De tich hop Google Admin SDK, moi customer SME phai cap quyen super-admin cho SMESec service account. Quy trinh nay thuong mat 1-2 tuan vi:
- Customer IT team phai review permissions (security review noi bo)
- Phai co GCP project setup
- OAuth consent screen phai duoc Google review (neu publish, mat 4-6 tuan)

**Hau qua:** Sprint 2 (T1) yeu cau co real Google Workspace tenant de test. Neu khong chuan bi tu W1, S2 se bi block.

**Khuyen nghi:** Dung test tenant cua chinh team (mua Google Workspace trial), khong phu thuoc vao pilot customer. Setup process nay song song ngay tu S1.

---

#### HIGH-03: Pilot Customer Outreach -- Khong Co Timeline Cu The

**Van de:** Plan hien tai chi noi "can tim 2-3 customers" nhung khong co action plan cu the: ai phu trach, outreach khi nao, qualify khi nao, sign NDA khi nao. Track 2 S11 (W21-22) yeu cau customers da duoc onboard.

**Kinh nghiem thuc te:** Tim va qualify mot pilot SME (co IT admin, dang dung AI tools, san sang chia se data) mat 4-8 tuan. Ky NDA + Data Processing Agreement (GDPR) mat them 2-4 tuan. Neu bat dau outreach o W13-14, se khong kip cho S11 (W21-22).

**Hau qua:** Khong co pilot customers = Gate 4 bi cancel = Track 2 khong the validate = launch decision bi delay vo han.

**Khuyen nghi:** Bat dau pilot customer outreach ngay tu TUAN 1, song song voi development. Assign Product Manager phu trach full-time cho task nay trong Month 1-2.

---

#### HIGH-04: Browser Extension -- Chua Co Legal/GDPR Review

**Van de:** Browser extension (T2 S5, W9-10) intercept tat ca text user go tren AI tool domains. Day la data processing co y nghia phap ly nghiem trong:
- GDPR Article 6: Can co legal basis cho viec xu ly du lieu
- GDPR Article 13/14: Phai thong bao cho nguoi dung
- Mot so quoc gia co luat rieng ve workplace monitoring (e.g., DSVGO Germany)

**Hau qua:** Neu deploy browser extension ma khong co proper consent mechanism, SMESec co the bi kien vi vi pham GDPR. Pilot customers o EU se khong the dung extension.

**Khuyen nghi:** Add legal review task vao Sprint 4 (W7-8), truoc khi bat dau build extension. Can: (1) GDPR DPA template cho customers, (2) Employee consent notice, (3) Privacy policy update cho extension.

---

#### HIGH-05: Track 1 -- Track 2 Event Schema Chua Duoc Agreed

**Van de:** T2 S10 (W19-20) build EventBridge integration, nhung neu T1 va T2 chua agreed tren event schema thi 2 teams co the build incompatible systems trong 18 tuan song song.

**Reu qua:**
- T1 S8 (W15-16) build playbook engine -- se build triggers dua tren "expected event format"
- T2 S8 (W15-16) build deepfake detection -- se publish events theo "own format"
- T2 S10 (W19-20) khi integrate, phat hien schema khac nhau -> phai refactor ca 2 sides

**Khuyen nghi:** To chuc 1 buoi joint design session trong TUAN 1 giua Tech Lead T1 va ML Engineer/Backend T2 de agree tren EventBridge event schema va API contracts. Document thanh ADR (Architecture Decision Record).

---

### 13.4 Rui ro Cap Do MEDIUM (Can co mitigation plan, khong block ngay)

| ID | Mo ta | Track | Sprint | Mitigation |
|----|-------|-------|--------|------------|
| MED-01 | Vendor API access (Sensity, Reality Defender) mat 2-4 tuan onboarding | T2 | S7 | Request trial access ngay tu W1, truoc khi can (S7 = W13) |
| MED-02 | AWS Step Functions cho playbook engine -- team co the chua co experience | T1 | S8 | Tech Lead can validate experience truoc; neu chua co, spike 3 ngay o S7 |
| MED-03 | OPA/Rego language la learning curve cho backend engineers | T1 | S5 | 2 ngay training/spike session o S4 truoc khi viet policies chinh |
| MED-04 | S12 T2 ghi "pilot 4 tuan" nhung sprint chi 2 tuan -- mau thuan | T2 | S12 | Pilot phai bat dau tu S11 (W21), ket thuc o cuoi S12 (W24) -- 4 tuan tong |
| MED-05 | DevSecOps overload o S13 (pen-test T1 + security review T2 cung luc) | Both | S13 | Pen-test T1 doi len S12; T2 security review la internal, khong can DevSecOps full-time |
| MED-06 | FaceForensics++ dataset la 1.5TB -- download mat nhieu gio, can license | T2 | S1 | Dang ky download ngay W1; dung subset (100GB) cho prototype |
| MED-07 | Benchmark vs production accuracy gap -- Gate 1-3 tren test data khong dam bao real-world | T2 | S3/S6/S9 | Add adversarial examples vao test set; monitor production accuracy tu S12 |
| MED-08 | Keycloak version compatibility voi AWS ECS + RDS -- hay gap loi khi setup cluster | T1 | S1 | Dung Keycloak operator tren ECS; test voi Docker Compose truoc |

---

### 13.5 Revised Timeline Recommendation

**Van de goc:** 13 sprints x 2 tuan = 26 tuan = 6 thang, 0% buffer.

**De xuat sua:**

```
Track 1: 15 sprints x 2 tuan = 30 tuan = 7.5 thang
- S1a: AWS infra + DB (tach tu S1)
- S1b: Auth/SSO/MFA (tach tu S1)
- S5a: Slack + AWS discovery (tach tu S5)
- S5b: RBAC + audit (tach tu S5)
- S9a: 2 playbooks (giu nguyen)
- S9b/S10/S11: Mobile app (tach ra 3 sprints, hien tai la 1 sprint)
  -> Mobile sprint 1: Core screens (inventory, JIT approval)
  -> Mobile sprint 2: Incident wizard + push notifications
  -> Pen-test doi len tu S13 sang S12 (moi)
- Buffer sprint sau S6 (sau Offboarding checkpoint)
- S13 (moi): Launch + pilot onboard chi

Track 2: 14 sprints x 2 tuan = 28 tuan = 7 thang
- Gate 1 doi sang cuoi S4 (W8) thay vi W6
- Buffer sprint sau Gate 3 (S10 moi)
- Pilot thuc su chay du 4 tuan (S12-S13)
- S14: Launch decision + polish
```

**Tong ket impact:**

| | Plan Hien Tai | Sau Sua |
|--|---------------|---------|
| Track 1 timeline | 6 thang | 7.5 thang |
| Track 2 timeline | 6 thang | 7 thang |
| Buffer | 0% | ~15% |
| Mobile app quality | Rui ro cao | Kha thi |
| Pen-test co thoi gian fix | Khong | Co |
| Pilot customer outreach | W13+ | W1 |

---

### 13.6 Top 5 Actions Uu Tien Tuan 1

1. **Tach Sprint 1 (T1)** -- Review lai sprint plan T1, tach S1 thanh S1a va S1b
2. **Joint schema session** -- T1 Tech Lead + T2 ML/Backend agree EventBridge event schema trong 2 gio
3. **Bat dau pilot customer outreach** -- PM/Product phu trach, target co 3 leads bi qualified truoc W12
4. **Request vendor API trial accounts** -- Sensity AI + Reality Defender ngay tuan nay (lead time 2-4 tuan)
5. **Legal review kickoff** -- Thue/consult legal counsel de review browser extension data processing va GDPR obligations

---

### 13.7 Ket Luan

Tong the, 2-track approach la quyet dinh dung. Viec tach Foundation (high confidence) khoi AI Detection (high risk) la bai hoc kinh nghiem dung dan.

Tuy nhien, **sprint plan hien tai overestimate capacity va underestimate integration complexity**. Dieu nay rat pho bien trong giai doan planning -- team thuong tinh optimistic khi moi va chua gap van de thuc te.

**Xac suat deliver dung han voi plan hien tai: ~25%**
**Xac suat deliver voi revised plan (+7 weeks): ~70%**

Nen ap dung revised timeline. 7.5 thang de deliver Track 1 production-ready tot hon nhieu so voi 6 thang deliver mot san pham co loi ky thuat tich luy lon.

> "Plans are useless, but planning is indispensable." -- Dwight D. Eisenhower
>
> Plan khong phai de follow mot cach cung nhac -- no la cong cu de hieu ro nhung gi co the sai va chuan bi truoc.

---

## 14. Ket Qua Debate: Solution Architect vs PM/Risk Manager (2026-05-27)

> **Phuong phap:** 2 vong debate doc lap giua **Solution Architect 30 nam** (chuyen gia cybersecurity) va **PM/Risk Manager 30 nam**.
> Round 1: moi ben review plan doc lap. Round 2: phan hoi cheo cac diem bat dong. Ket qua: **toan bo dieu chinh duoc ca 2 ben dong thuan**.

### 14.1 Tong Quan Ket Qua

| Hang muc | Truoc Debate | Sau Debate |
|----------|-------------|------------|
| T1 sprint plan (Acceptance Criteria S1) | Thieu tenant isolation + secrets rotation + BCP stub | 3 criteria bo sung, CI test bat buoc |
| T1 Flutter Eng allocation | Idle 8 sprints (S1-S8) | Bat dau mobile scaffold tu S1 |
| T1 S9 mobile scope | Full mobile (Android+iOS+Desktop+incident wizard) | Thu hep: Android+iOS only, JIT + read-only |
| T1 Pen-test timing | S13 (qua tre, khong co thoi gian fix) | Bat dau S12, S13 = remediation + launch |
| T1-T2 schema alignment | Khong co sprint nao define shared schema | Joint session Tuan 1, freeze by end S2 |
| T1 Pen-test vendor procurement | Chua plan | LOI ky truoc cuoi S8 (lead time 2-3 tuan) |
| T2 Chrome MV3 blocker | Khong phat hien (PM bo sot) | HARD GATE W1 Week 1 -- SA phat hien |
| T2 contracts-first | Extension Eng idle, spec den sau | Extension Eng own shared-types + OpenAPI tu S1 |
| T2 Gate 1 timing | W6 (cuoi S3) -- qua nhanh cho ML iteration | Chuyen sang W8 (cuoi S4) |
| T2 BERT load test | Khong co trong plan | NER chuyen sang S4, S3 = BERT training + load test |
| T2 Dataset quality | Khong co exit criteria | Inter-rater label agreement >85% tren 200-sample |
| T2 Pilot customer procurement | W13+ | NDA+DPA ky truoc S7 (W13), outreach bat dau W1 |
| T2 Vendor contract risk | Gia su vendor san sang | DeepfakeDetector abstraction + Resemblyzer fallback tu W1 |

### 14.2 Dieu Chinh Track 1 (6 items)

| # | Dieu chinh | Sprint bi anh huong | Nguon |
|---|-----------|---------------------|-------|
| T1-1 | Them 3 acceptance criteria vao S1: (a) tenant isolation CI test "Tenant A token cannot retrieve Tenant B's assets"; (b) secrets rotation policy doc + manual runbook + CloudWatch age alerting; (c) RDS Multi-AZ BCP stub | S1 | SA: cross-tenant leak = existential risk; PM: ISO 27001 compliance blocker |
| T1-2 | Flutter Eng bat dau mobile scaffold architecture tu S1 (khong doi den S9) | S1 -> S8 | PM: 8 sprints idle la lang phi. SA: scaffold sach thi S9 scope giam kha thi hon |
| T1-3 | S9 mobile scope giam: Android+iOS only, JIT + read-only. Desktop flutter va incident wizard -> S10. APNs -> S11 | S9, S10, S11 | SA: 1 Flutter Eng khong build full mobile trong 1 sprint. PM: demo milestone can scope sach |
| T1-4 | Pen-test STARTS tai S12 (khong phai S13). S13 = remediation + launch ONLY | S12, S13 | SA+PM dong y: neu co Critical finding, can it nhat 1 sprint de fix. Launch voi unfixed Critical = ISO 27001 violation |
| T1-5 | Joint T1-T2 schema session trong Tuan 1: define ThreatDetectionEvent interface, freeze by end of S2. T1 EventBridge bus provisioned trong S1 | S1, S2 | SA: incompatible schemas khi S10 integrate = refactor ca 2 phia. PM: khong co ai flag dieu nay trong plan |
| T1-6 | PM phai lua chon va ky LOI voi pen-test vendor truoc cuoi S8 (lead time 2-3 tuan). Vendor shortlist truoc cuoi S6 | S6, S8, S11 | PM: external vendor khong the fast-track. SA: pen-test prerequisites = stable infra (sau S11) |

### 14.3 Dieu Chinh Track 2 (7 items)

| # | Dieu chinh | Sprint bi anh huong | Nguon |
|---|-----------|---------------------|-------|
| T2-1 | Chrome MV3 service worker persistence prototype la HARD GATE cuoi S1 Week 1. Neu khong co giai phap kha thi -> cut extension khoi V1, ship API-only detection | S1 | SA phat hien (PM da bo sot): MV3 workers terminate sau 30s inactivity = pha vo toan bo extension monitoring |
| T2-2 | Extension Eng owned `libs/shared-types/src/events/threat-event.ts` + OpenAPI spec tu S1 (contracts-first approach) | S1 -> S4 | SA: define contracts truoc khi 2 tracks build doc lap. PM: Extension Eng idle 6 sprints = lang phi capacity |
| T2-3 | Gate 1 chuyen tu W6 (cuoi S3) sang W8 (cuoi S4) | S3, S4 | PM: 4 tuan qua ngan cho ML training iteration cycle. SA: data quality risk cao hon model capability risk |
| T2-4 | NER model chuyen tu S3 sang S4; S3 chi = BERT training + BERT load test (SageMaker cold-start validation tai 100 concurrent requests) | S3, S4 | SA: cold-start 10-20s phai validate truoc khi extension ship (S5). PM: cat NER de add load test scope |
| T2-5 | Them dataset quality exit criteria vao S1: inter-rater label agreement >85% tren 200-sample blind check | S1, Gate 1 | SA: Gate 1 failure la data quality issue, khong phai model capability. PM: can exit criteria cu the de Go/No-go |
| T2-6 | Pilot customer NDA+DPA phai ky truoc S7 (W13). Outreach bat dau W1. Target: 3 qualified leads truoc W12 | S1, S7 | PM: procurement mat 4-8 tuan. SA: neu khong co customers, Gate 4 void |
| T2-7 | DeepfakeDetector abstraction interface tu Week 1. Open-source fallback (Resemblyzer, ~78-82% accuracy) cho den khi vendor ky hop dong chinh thuc | S1, S7 | SA: Resemblyzer la acceptable tripwire/fallback. PM: vendor procurement co the mat den S4-S5 |

### 14.4 Phat Hien Quan Trong Nhat (Top 3)

**1. Chrome MV3 Service Worker Termination (T2-1) -- CRITICAL, da bo sot**
> Chrome Manifest V3 service workers terminate sau 30 giay inactivity. Browser extension monitor AI prompts theo thoi gian thuc se bi terminate ngay truoc khi bam Submit. Day la architectural blocker chua ai trong team flag truoc khi SA review. Neu prototype W1 that bai -> toan bo extension approach phai bo, chi con API-only detection. Decision must be made by end Week 1.

**2. ThreatDetectionEvent Schema Contract (T1-5 / T2-2) -- HIGH**
> T1 va T2 dang build doc lap voi gia dinh Event schema khac nhau. Neu khong freeze schema truoc khi build, S10 integration se yeu cau refactor ca 2 phia (2-3 week delay). Schema phai bao gom: `eventId`, `sourceTrack`, `assetId`, `actorId`, `threatClass`, `severity`, `confidenceScore`, `evidenceRef`, `detectedAt`, `policyMatchId`. EventBridge bus provisioned by T1 trong S1, owned by Extension Eng trong T2.

**3. Pen-Test Timeline Compress (T1-4) -- HIGH**
> Plan cu: pen-test trong S13 (final sprint). Neu phat hien Critical finding -> khong co thoi gian fix -> launch voi known Critical = ISO 27001 violation va reputational risk. Plan moi: bat dau pen-test S12, danh toan bo S13 cho remediation + hardening + launch preparation.

### 14.5 Tac Dong Den Xac Suat Thanh Cong

> Phan tich tu Section 13.7: xac suat deliver voi revised plan = ~70%.
> Sau khi ap dung 13 dieu chinh tu debate:

| Rui ro | Truoc Debate | Sau Debate |
|--------|-------------|------------|
| Cross-tenant data leak | Khong co CI test | CI test bat buoc S1 |
| Extension broken by MV3 | Khong biet (undetected) | HARD GATE W1 -- detect som nhat |
| T1-T2 schema mismatch tao ra S10 refactor | Cao (khong co session) | Giam (freeze by S2) |
| Pen-test qua tre, khong co thoi gian fix | Cao | Giam (bat dau S12) |
| ML Gate 1 failure vi data quality | Cao (khong co exit criteria) | Giam (inter-rater >85%) |
| Flutter Eng idle waste | 8 sprints lang phi | S1 bat dau, dung capacity tot hon |
| Vendor API khong san sang | Khong co fallback | Resemblyzer abstraction layer |
| Pilot customer khong co | Outreach qua tre | Outreach W1, NDA ky truoc S7 |

**Xac suat deliver sau debate adjustments: ~75-80%** (tang tu ~70% o Section 13.7)

### 14.6 Actions Tuan 1 (Cap Nhat Tu Debate)

> Thay the Section 13.6 "Top 5 Actions Uu Tien Tuan 1":

1. **[HARD GATE]** Extension Eng: Chrome MV3 service worker persistence prototype -- binary Go/No-go decision cuoi Week 1
2. **[CRITICAL]** T1 Tech Lead + T2 ML Eng + T2 Extension Eng: 2-gio joint schema session -- define va commit `ThreatDetectionEvent` interface, T1 provision EventBridge bus
3. **[HIGH]** PM: Pilot customer outreach bat dau W1 -- target 3 qualified leads truoc W12
4. **[HIGH]** PM: Vendor API trial accounts -- Sensity AI + Reality Defender request ngay W1 (lead time 2-4 tuan)
5. **[HIGH]** T2 ML Eng: Dataset labeling guide + inter-rater agreement protocol -- 200-sample blind check truoc khi training bat dau
6. **[MEDIUM]** PM: Pen-test vendor shortlist (target: chon truoc cuoi S6, ky LOI truoc cuoi S8)
7. **[MEDIUM]** Legal review kickoff: browser extension data processing obligations + GDPR Art.25 privacy by design

---

