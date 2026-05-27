# Tổng Quan Các Tiêu Chuẩn Tuân Thủ

## Giới Thiệu

Tài liệu này so sánh chi tiết ba tiêu chuẩn/quy định chính mà SMESec cần đạt được:
- **ISO 27001:** Tiêu chuẩn quốc tế về Hệ thống Quản lý An ninh Thông tin (ISMS)
- **SOC 2:** Báo cáo kiểm toán về kiểm soát bảo mật (Hoa Kỳ)
- **GDPR:** Quy định bảo vệ dữ liệu cá nhân (EU)

## So Sánh Tổng Quan

| Tiêu chuẩn | Bản chất | Phạm vi | Kết quả | Thời gian | Chi phí |
|------------|----------|---------|---------|-----------|---------|
| **ISO 27001** | Tiêu chuẩn quốc tế | Toàn tổ chức | Chứng chỉ (Certificate) | 3-6 tháng | $3K-5K |
| **SOC 2 Type 1** | Báo cáo kiểm toán | Hệ thống IT | Báo cáo (Report) | 2-3 tháng | $5K-8K |
| **SOC 2 Type 2** | Báo cáo kiểm toán | Hệ thống IT | Báo cáo (Report) | 6-12 tháng | $8K-15K |
| **GDPR** | Quy định pháp lý | Dữ liệu cá nhân | Tuân thủ (Compliance) | Ongoing | Varies |

## ISO 27001: Information Security Management System

### Tổng Quan
ISO 27001 là tiêu chuẩn quốc tế về Hệ thống Quản lý An ninh Thông tin (ISMS). Đây là framework tổng thể giúp tổ chức quản lý và bảo vệ thông tin một cách có hệ thống.

### Trọng Tâm Kiểm Tra

#### 1. Khung Quản Lý (Management Framework)
- **Chính sách bảo mật:** Tài liệu chính sách tổng thể
- **Vai trò & trách nhiệm:** Phân công rõ ràng
- **Đánh giá rủi ro:** Risk assessment methodology
- **Mục tiêu bảo mật:** Security objectives và KPIs

#### 2. Kiểm Soát Kỹ Thuật (Technical Controls)
- **Kiểm soát truy cập:** Access control policies
- **Mã hóa:** Encryption at rest và in transit
- **Quản lý lỗ hổng:** Vulnerability management
- **Backup & recovery:** Disaster recovery plans

#### 3. Kiểm Soát Tổ Chức (Organizational Controls)
- **Nhân sự:** Background checks, training
- **Nhà cung cấp:** Vendor management
- **Sự cố:** Incident response procedures
- **Tuân thủ:** Legal và regulatory compliance

#### 4. Kiểm Soát Vật Lý (Physical Controls)
- **Bảo mật văn phòng:** Physical security
- **Thiết bị:** Device management
- **Môi trường:** Environmental controls

### Quy Trình Chứng Nhận

```
Stage 1: Document Review (1-2 tuần)
├─ Kiểm tra tài liệu chính sách
├─ Review risk assessment
└─ Đánh giá readiness

Stage 2: On-site Audit (1-2 tuần)
├─ Phỏng vấn nhân viên
├─ Kiểm tra hệ thống
├─ Xác minh kiểm soát
└─ Tạo báo cáo phát hiện

Corrective Actions (2-4 tuần)
├─ Fix non-conformities
└─ Submit evidence

Certification Decision (1-2 tuần)
└─ Issue certificate (valid 3 years)

Annual Surveillance (ongoing)
└─ Yearly audits to maintain
```

### Yêu Cầu Đối Với SMESec

**Tài liệu cần chuẩn bị:**
- Information Security Policy
- Risk Assessment & Treatment Plan
- Statement of Applicability (SoA)
- Asset Inventory
- Access Control Policy
- Incident Response Plan
- Business Continuity Plan
- Supplier Management Policy

**Bằng chứng vận hành:**
- Access logs (AWS CloudTrail)
- Security scan results (Dependabot, CodeQL)
- Training records
- Incident logs
- Change management records

### Lợi Ích
- ✅ Công nhận quốc tế, đặc biệt ở EU và châu Á
- ✅ Framework tổng thể cho toàn tổ chức
- ✅ Tạo nền tảng cho các chứng nhận khác
- ✅ Chứng minh commitment với bảo mật

### Thách Thức
- ⚠️ Yêu cầu nhiều tài liệu và chính sách
- ⚠️ Cần đánh giá rủi ro toàn diện
- ⚠️ Audit hàng năm để duy trì

## SOC 2: Service Organization Control

### Tổng Quan
SOC 2 là báo cáo kiểm toán do AICPA (American Institute of CPAs) phát triển, tập trung vào kiểm soát bảo mật của service providers.

### Trust Services Criteria (TSC)

SOC 2 có 5 TSC, nhưng SMESec tập trung vào **Security** (bắt buộc):

#### Security (Bắt buộc)
- **CC1:** Control Environment
- **CC2:** Communication and Information
- **CC3:** Risk Assessment
- **CC4:** Monitoring Activities
- **CC5:** Control Activities
- **CC6:** Logical and Physical Access Controls
- **CC7:** System Operations
- **CC8:** Change Management
- **CC9:** Risk Mitigation

#### Optional Criteria (Có thể thêm sau)
- **Availability:** System uptime và performance
- **Processing Integrity:** Data processing accuracy
- **Confidentiality:** Sensitive data protection
- **Privacy:** Personal information handling

### Type 1 vs Type 2

| Aspect | Type 1 | Type 2 |
|--------|--------|--------|
| **Thời điểm** | Point-in-time | Period of time (6-12 months) |
| **Kiểm tra** | Controls exist | Controls operate effectively |
| **Timeline** | 2-3 tháng | 6-12 tháng |
| **Chi phí** | $5K-8K | $8K-15K |
| **Phù hợp** | V1 launch | Production stable |

### Quy Trình Kiểm Toán

```
Scoping (1-2 tuần)
├─ Define system boundaries
├─ Select TSC criteria
└─ Agree on timeline

Readiness Assessment (2-4 tuần)
├─ Connect Vanta to systems
├─ Review control evidence
└─ Identify gaps

Audit Fieldwork (2-4 tuần)
├─ Auditor reviews Vanta dashboard
├─ Sample testing
├─ Interviews
└─ Evidence validation

Report Issuance (1-2 tuần)
└─ SOC 2 Report delivered
```

### Yêu Cầu Đối Với SMESec

**Kiểm soát kỹ thuật:**
- Multi-factor authentication (MFA)
- Encryption at rest và in transit
- Access logging (CloudTrail)
- Vulnerability scanning (Dependabot, CodeQL)
- Backup và disaster recovery
- Change management process

**Kiểm soát tổ chức:**
- Background checks
- Security awareness training
- Vendor risk management
- Incident response procedures
- Policy documentation

**Bằng chứng tự động (qua Vanta):**
- AWS security configurations
- GitHub access controls
- Employee training completion
- Security scan results
- Access logs

### Lợi Ích
- ✅ Yêu cầu bởi khách hàng Hoa Kỳ
- ✅ Tập trung vào technical controls
- ✅ Automation-friendly (Vanta)
- ✅ Type 1 nhanh, phù hợp v1 launch

### Thách Thức
- ⚠️ Type 2 yêu cầu 6-12 tháng vận hành
- ⚠️ Chi phí audit cao hơn ISO 27001
- ⚠️ Cần renew hàng năm

## GDPR: General Data Protection Regulation

### Tổng Quan
GDPR là quy định pháp lý của EU về bảo vệ dữ liệu cá nhân. Áp dụng cho mọi tổ chức xử lý dữ liệu của công dân EU, bất kể tổ chức ở đâu.

### Nguyên Tắc Chính

#### 1. Lawfulness, Fairness, Transparency
- Xử lý dữ liệu hợp pháp và minh bạch
- Thông báo rõ ràng về việc thu thập dữ liệu

#### 2. Purpose Limitation
- Chỉ thu thập dữ liệu cho mục đích cụ thể
- Không sử dụng cho mục đích khác

#### 3. Data Minimization
- Chỉ thu thập dữ liệu cần thiết
- Không thu thập dư thừa

#### 4. Accuracy
- Dữ liệu phải chính xác và cập nhật
- Cho phép người dùng sửa đổi

#### 5. Storage Limitation
- Chỉ lưu trữ trong thời gian cần thiết
- Xóa khi không còn cần

#### 6. Integrity and Confidentiality
- Bảo mật dữ liệu
- Bảo vệ khỏi truy cập trái phép

#### 7. Accountability
- Chứng minh tuân thủ
- Maintain records of processing

### Quyền Của Data Subject

| Quyền | Mô tả | Yêu cầu kỹ thuật |
|-------|-------|------------------|
| **Right to Access** | Xem dữ liệu cá nhân | Export API |
| **Right to Rectification** | Sửa dữ liệu sai | Update API |
| **Right to Erasure** | Xóa dữ liệu | Hard delete function |
| **Right to Portability** | Chuyển dữ liệu | Export in machine-readable format |
| **Right to Object** | Từ chối xử lý | Opt-out mechanism |
| **Right to Restriction** | Hạn chế xử lý | Pause processing flag |

### Yêu Cầu Đối Với SMESec

#### Technical Measures
```
Data Protection by Design:
├─ Encryption at rest (AES-256)
├─ Encryption in transit (TLS 1.3)
├─ Pseudonymization where possible
├─ Access controls (least privilege)
└─ Audit logging (CloudTrail)

Data Subject Rights Implementation:
├─ Export API (JSON format)
├─ Update API (self-service)
├─ Hard delete function (cascade)
└─ Consent management system
```

#### Organizational Measures
- **Privacy Policy:** Rõ ràng, dễ hiểu
- **Data Processing Agreement (DPA):** Với sub-processors (AWS, Cloudflare)
- **Data Breach Notification:** Quy trình thông báo trong 72h
- **Data Protection Impact Assessment (DPIA):** Cho high-risk processing
- **Records of Processing Activities (RoPA):** Maintain inventory

#### Documentation
- Privacy Policy
- Cookie Policy (nếu có)
- Data Processing Agreements
- Data Retention Policy
- Data Breach Response Plan
- GDPR Deletion Runbook

### Quy Trình Tuân Thủ

```
Assessment (2-4 tuần)
├─ Identify personal data
├─ Map data flows
├─ Assess legal basis
└─ Identify gaps

Implementation (4-8 tuần)
├─ Implement technical measures
├─ Create documentation
├─ Train staff
└─ Update privacy policy

Validation (2-4 tuần)
├─ Test data subject rights
├─ Review with legal
└─ Document compliance

Ongoing Maintenance
├─ Annual DPIA review
├─ Breach notification drills
└─ Policy updates
```

### Lợi Ích
- ✅ Bắt buộc nếu có khách hàng EU
- ✅ Tăng trust với users
- ✅ 70% overlap với ISO 27001 + SOC 2

### Thách Thức
- ⚠️ Phạt nặng (€20M hoặc 4% revenue)
- ⚠️ Yêu cầu legal expertise
- ⚠️ Cần implement data subject rights

## Mối Quan Hệ Giữa Các Tiêu Chuẩn

### Overlap Matrix

|  | ISO 27001 | SOC 2 | GDPR |
|--|-----------|-------|------|
| **Access Control** | ✅ | ✅ | ✅ |
| **Encryption** | ✅ | ✅ | ✅ |
| **Logging** | ✅ | ✅ | ✅ |
| **Incident Response** | ✅ | ✅ | ✅ |
| **Risk Assessment** | ✅ | ✅ | ⚠️ (DPIA) |
| **Data Subject Rights** | ❌ | ❌ | ✅ |
| **Vendor Management** | ✅ | ✅ | ✅ (DPA) |
| **Training** | ✅ | ✅ | ✅ |

### Synergy Strategy

**Phase 1: ISO 27001 Foundation**
- Tạo framework tổng thể
- Establish policies và procedures
- Risk assessment methodology

**Phase 2: SOC 2 Technical Implementation**
- Implement technical controls
- Automation với Vanta
- Evidence collection

**Phase 3: GDPR Specific Requirements**
- Data subject rights APIs
- Privacy-specific documentation
- Legal review

**Kết quả:** 70-80% công việc overlap, chỉ cần 20-30% effort bổ sung cho mỗi tiêu chuẩn.

## Checklist Tổng Hợp

### ISO 27001
- [ ] Information Security Policy
- [ ] Risk Assessment & Treatment Plan
- [ ] Statement of Applicability
- [ ] Asset Inventory
- [ ] Access Control Policy
- [ ] Incident Response Plan
- [ ] Business Continuity Plan
- [ ] Supplier Management Policy
- [ ] Training records
- [ ] Audit logs

### SOC 2
- [ ] Vanta setup và integration
- [ ] MFA enabled for all accounts
- [ ] Encryption at rest và in transit
- [ ] CloudTrail logging enabled
- [ ] Dependabot + CodeQL enabled
- [ ] Backup và DR procedures
- [ ] Change management process
- [ ] Vendor risk assessments
- [ ] Security awareness training
- [ ] Incident response runbook

### GDPR
- [ ] Privacy Policy published
- [ ] Cookie Policy (if applicable)
- [ ] Data Processing Agreements signed
- [ ] Data Retention Policy
- [ ] Export API implemented
- [ ] Hard delete function implemented
- [ ] Consent management system
- [ ] Data breach notification procedure
- [ ] DPIA for high-risk processing
- [ ] Records of Processing Activities

## Tài Liệu Tham Khảo

### ISO 27001
- [ISO 27001:2022 Official](https://www.iso.org/standard/27001)
- [ISO 27001 Controls](https://www.isms.online/iso-27001/)

### SOC 2
- [AICPA Trust Services Criteria](https://www.aicpa.org/interestareas/frc/assuranceadvisoryservices/aicpasoc2report)
- [SOC 2 Guide](https://www.vanta.com/resources/soc-2-compliance-guide)

### GDPR
- [GDPR Official Text](https://gdpr-info.eu/)
- [GDPR Checklist](https://gdpr.eu/checklist/)
- [ICO GDPR Guide](https://ico.org.uk/for-organisations/guide-to-data-protection/guide-to-the-general-data-protection-regulation-gdpr/)

---

**Cập nhật:** 2026-05-27  
**Người phụ trách:** Quách Thanh Bình
