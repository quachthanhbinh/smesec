# SMESec Platform - 2-Track Development Strategy

**Date:** 2026-05-27  
**Status:** Proposed  
**Context:** Phân chia development thành 2 tracks song song để giảm risk và tăng confidence

---

## Vấn Đề Với Thiết Kế Trước

**Critical Risk:** AI threat detection chỉ đạt 85% accuracy → 15% còn lại có thể chứa:
- **False Negatives**: Bỏ sót deepfake fraud, prompt injection thực sự → mất tiền, mất data
- **False Positives**: Block nhầm công việc hợp lệ → frustrate employees, giảm productivity

→ **Mất lòng tin khách hàng**, không ai dám dùng nếu không tin tưởng 100%

---

## Chiến Lược Mới: 2 Tracks Song Song

### Track 1: Foundation & Governance (High Confidence)
**Mục tiêu:** Xây nền tảng vững chắc với accuracy gần 100%

**Scope:**
1. **Asset Inventory & Classification** (Requirement 1)
   - Tự động discover: devices, accounts, SaaS apps, cloud resources
   - Classification: criticality, sensitivity, owner
   - Dependency mapping
   
2. **Access Governance** (Requirement 3)
   - Least-privilege enforcement (RBAC + JIT access)
   - Automated offboarding (revoke all access trong <5 phút)
   - Shadow IT detection (OAuth apps, unapproved SaaS)
   - SSO + MFA enforcement

3. **Incident Playbooks** (Requirement 4 - subset)
   - Account compromise response
   - Unauthorized access response
   - Shadow IT remediation
   - Offboarding emergency procedures

4. **Compliance Foundation** (Requirement 4 - subset)
   - ISO 27001 controls: Asset Management (A.8), Access Control (A.9)
   - GDPR: Data inventory, access logs, right to erasure
   - SOC 2: Access provisioning/deprovisioning, audit trails

5. **Core Integrations** (Requirement 6)
   - Google Workspace (Admin API, Audit API)
   - Microsoft 365 (Graph API, Azure AD)
   - Slack (Admin API, Audit Logs)
   - AWS/Azure/GCP (asset discovery APIs)

**Tại sao Track 1 có High Confidence:**
- Công nghệ proven: OAuth 2.0, RBAC, API integrations
- Không phụ thuộc vào ML/AI (deterministic logic)
- Có thể test 100% với automated tests
- Immediate value: visibility + control ngay lập tức

**Deliverables:**
- Web Dashboard: Asset inventory, access management, policy config
- Mobile/Desktop App: Alerts, quick actions, incident wizards
- API Gateway: Integration hub cho Google/M365/Slack
- Compliance Reports: ISO 27001, GDPR, SOC 2 evidence

---

### Track 2: AI Threat Detection (High Risk, High Value)
**Mục tiêu:** R&D để đạt accuracy >95% trước khi launch

**Scope:**
1. **Prompt Injection Detection**
   - Rule-based patterns (regex)
   - ML classifier (fine-tuned BERT/GPT)
   - Contextual analysis
   - **Target:** >95% precision, <5% false positive

2. **LLM Data Leakage Prevention**
   - DLP patterns (PII, credentials, IP)
   - Dynamic redaction (mask sensitive data)
   - Semantic analysis (detect trade secrets in prose)
   - **Target:** 0 false negatives on critical data (credit cards, passwords)

3. **Deepfake Detection**
   - Voice cloning detection (vendor API + custom model)
   - Video deepfake detection (FaceForensics++ models)
   - Liveness detection (challenge-response)
   - **Target:** >90% detection rate, <10% false positive

4. **Shadow AI Discovery**
   - Browser telemetry (ChatGPT, Copilot usage)
   - OAuth app inventory (AI tools authorized)
   - Network traffic analysis (DNS, API calls)
   - **Target:** >95% discovery rate

**Tại sao Track 2 cần R&D riêng:**
- Phụ thuộc vào ML models (non-deterministic)
- Cần pilot data để tune thresholds
- False positive/negative trade-offs phức tạp
- Cần thời gian để validate accuracy

**Deliverables:**
- Browser Extension: Prompt interceptor, DLP scanner
- Desktop Agent: Clipboard monitoring, app traffic inspection
- AI Threat Detection Service: ML inference, risk scoring
- Deepfake Detection API: Voice/video analysis

**Validation Gates:**
- **Week 6:** Prompt injection precision >90% on test dataset
- **Week 12:** DLP false negative rate <1% on critical data
- **Week 18:** Deepfake detection >85% on benchmark dataset
- **Week 24:** Pilot with 2-3 customers, collect real-world metrics

---

## Phân Chia Team

### Team 1: Foundation & Governance (5 FTE)
- 1 Tech Lead / Architect
- 2 Backend Engineers (Go + Python)
- 1 Frontend Engineer (React/Next.js)
- 1 Flutter Engineer (Mobile/Desktop)

**Focus:**
- Asset discovery & classification
- Access governance (RBAC, offboarding)
- Integrations (Google, M365, Slack)
- Compliance reporting

### Team 2: AI Threat Detection (3 FTE)
- 1 ML Engineer / Security Researcher
- 1 Backend Engineer (Python/FastAPI)
- 1 Frontend Engineer (Browser Extension + Desktop Agent)

**Focus:**
- Prompt injection detection
- LLM DLP
- Deepfake detection
- Browser extension + desktop agent

### Shared Resources (2 FTE)
- 1 Product Manager / Security Analyst (coordinate both tracks)
- 1 DevSecOps / QA (CI/CD, testing, infrastructure)

**Total:** 10 FTE

---

## Timeline: 6 Tháng Song Song

### Track 1: Foundation (Launch-Ready sau 6 tháng)

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

### Track 2: AI Detection (Pilot-Ready sau 6 tháng)

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

## Dependencies Giữa 2 Tracks

### Track 1 → Track 2 (Foundation cung cấp cho AI)
- **Asset Inventory:** AI detection cần biết user roles, device context
- **Access Governance:** AI alerts cần trigger offboarding workflows
- **Incident Playbooks:** AI threats cần automated response
- **Integrations:** AI detection cần Google/M365 APIs để revoke access

### Track 2 → Track 1 (AI cung cấp cho Foundation)
- **AI Threat Events:** Feed vào incident playbook engine
- **Risk Scores:** Enrich asset classification (high-risk users/devices)
- **Shadow AI Discovery:** Feed vào access governance (unapproved apps)

### Integration Point: Event-Driven Architecture
```
[Track 2: AI Detection] 
         │ Publishes events
         ▼
[EventBridge / SQS]
         │ Routes events
         ▼
[Track 1: Playbook Engine]
         │ Executes response
         ▼
[Track 1: Access Governance]
         │ Revokes access
         ▼
[Track 1: Integrations]
```

---

## Compliance Requirements cho Track 1

### ISO 27001 Controls (Cần thiết cho Track 1)

| Control | Requirement | Implementation |
|---------|-------------|----------------|
| **A.8.1** Asset Management | Inventory of assets | Asset discovery + classification engine |
| **A.8.2** Information Classification | Classify by sensitivity | Criticality levels (Critical/High/Medium/Low) |
| **A.9.1** Access Control Policy | Least privilege | RBAC + JIT access |
| **A.9.2** User Access Management | Provisioning/deprovisioning | Automated offboarding workflows |
| **A.9.4** Access Review | Periodic review | Quarterly access review reports |
| **A.12.4** Logging & Monitoring | Audit trails | All access events logged to S3 |

### GDPR Requirements (Cần thiết cho Track 1)

| Article | Requirement | Implementation |
|---------|-------------|----------------|
| **Art. 30** Records of Processing | Data inventory | Asset inventory includes data assets |
| **Art. 32** Security Measures | Access controls | RBAC + MFA + encryption |
| **Art. 17** Right to Erasure | Delete personal data | Offboarding workflow includes data deletion |
| **Art. 33** Breach Notification | 72-hour reporting | Incident playbooks include notification workflows |

### SOC 2 Trust Services Criteria (Cần thiết cho Track 1)

| Criteria | Requirement | Implementation |
|----------|-------------|----------------|
| **CC6.1** Logical Access | Restrict access | RBAC + least privilege |
| **CC6.2** Access Provisioning | Timely provisioning/deprovisioning | Automated offboarding <5 minutes |
| **CC6.3** Access Removal | Remove access when no longer needed | JIT access auto-expires |
| **CC7.2** System Monitoring | Monitor security events | All access events logged + alerted |

**Kết luận:** Track 1 đủ để đạt ISO 27001, GDPR, SOC 2 compliance cho phần Asset + Access Management. Track 2 (AI detection) là bonus, không bắt buộc cho compliance.

---

## Success Criteria

### Track 1 (Launch Criteria)
- ✅ Asset discovery coverage >95% (all devices, accounts, SaaS apps)
- ✅ Offboarding completion time <5 minutes (all platforms)
- ✅ Shadow IT detection rate >90% (OAuth apps)
- ✅ Zero cross-tenant data leakage (security audit passed)
- ✅ Compliance reports generated (ISO 27001, GDPR, SOC 2)
- ✅ 5-10 pilot customers onboarded successfully

### Track 2 (Pilot Criteria)
- ✅ Prompt injection precision >95% (on test dataset)
- ✅ DLP false negative rate <1% (on critical data: credit cards, passwords)
- ✅ Deepfake detection >90% (on benchmark dataset)
- ✅ False positive rate <5% (on pilot customer data)
- ✅ 2-3 pilot customers validate accuracy in production

**Decision Gate (Month 6):**
- If Track 2 meets criteria → Merge into main product
- If Track 2 needs more work → Continue as beta feature, iterate based on pilot feedback

---

## Risk Mitigation

### Track 1 Risks
| Risk | Impact | Mitigation |
|------|--------|------------|
| Integration API changes | Medium | Daily integration tests, version pinning |
| Performance issues at scale | Medium | Load testing, caching, rate limiting |
| Compliance audit failure | High | External audit review at Month 5 |

### Track 2 Risks
| Risk | Impact | Mitigation |
|------|--------|------------|
| AI accuracy insufficient | **CRITICAL** | Early validation gates (Week 6, 12, 18) |
| False positives frustrate users | High | Pilot with friendly customers, tune thresholds |
| Deepfake detection too slow | Medium | Use vendor APIs, optimize inference |

---

## Next Steps

1. ✅ **Approve 2-track strategy** (this document)
2. ⏳ **Create detailed requirements** for Track 1 (Asset + Access + Playbooks)
3. ⏳ **Create research plan** for Track 2 (AI detection accuracy targets)
4. ⏳ **Set up 2 team structures** (hire if needed)
5. ⏳ **Kick off Month 1** for both tracks in parallel

---

## Decisions Made

1. **Team size:** TBD - sẽ thảo luận riêng

2. **Pilot customers:** ⚠️ **CRITICAL** - Chưa có sẵn SMEs để pilot
   - **Action required:** Cần tìm 2-3 SMEs sẵn sàng pilot Track 2 trong Month 4-6
   - **Criteria:** SMEs có:
     - 50-200 employees (sweet spot cho pilot)
     - Đang dùng AI tools (ChatGPT, Copilot, etc.)
     - Có IT manager/admin có thể collaborate
     - Sẵn sàng share feedback và telemetry data
   - **Timeline:** Cần identify pilot customers trước Month 4

3. **Budget Track 2:** ✅ Approved
   - Deepfake detection API ($3K-5K/year)
   - ML training infrastructure (SageMaker)
   - Labeled datasets (nếu cần mua)

4. **Decision criteria:** ⚠️ **QUALITY FIRST**
   - Nếu Track 2 chỉ đạt 90% accuracy sau 6 tháng → **KHÔNG launch**
   - Tiếp tục iterate cho đến khi đạt >95% precision, <5% false positive
   - **Nguyên tắc:** Không đưa sản phẩm chưa hoàn thiện ra thị trường
   - Track 1 có thể launch độc lập nếu Track 2 chưa sẵn sàng
