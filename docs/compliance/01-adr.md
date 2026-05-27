# Architecture Decision Record (ADR)

## Metadata
- **Ngày tạo:** 2026-05-27
- **Người phụ trách:** Quách Thanh Bình - Senior Backend Engineer
- **Trạng thái:** Approved
- **Phiên bản:** 1.0

## Context (Bối cảnh)

### Tình Hình Hiện Tại
Hệ thống SMESec đang được xây dựng từ đầu với các đặc điểm:
- **Giai đoạn:** Build from scratch
- **Timeline:** Dự kiến v1 sau 6 tháng
- **Quy mô:** Phục vụ lượng lớn dữ liệu tải trọng cao
- **Hạ tầng:** AWS (compute & services) + Cloudflare R2 (storage)
- **Source Control:** GitHub

### Thách Thức
1. **Yêu cầu tuân thủ:** Cần đạt các chứng nhận bảo mật quốc tế (ISO 27001, SOC 2, GDPR)
2. **Tốc độ phát triển:** Không được làm chậm SDLC (Software Development Life Cycle)
3. **Độ tin cậy:** Tăng độ tin cậy với khách hàng doanh nghiệp
4. **Nguồn lực hạn chế:** Team nhỏ, cần tối ưu hóa effort

### Ràng Buộc
- Ngân sách startup (~$12K-19K/năm cho compliance)
- Team kỹ thuật nhỏ (Senior Backend Engineer lead)
- Cần release v1 trong 6 tháng
- Không có dedicated security team

## Decisions (Quyết định)

### 1. Chiến Lược Chứng Nhận

**Quyết định:** Tiếp cận theo lộ trình "ISO 27001 & SOC 2 Type 1 (Security)" trước, sau đó duy trì để lấy SOC 2 Type 2.

**Lý do:**
- **SOC 2 Type 1:** Kiểm tra tại một thời điểm (point-in-time), nhanh hơn, phù hợp với timeline v1
- **ISO 27001:** Framework tổng thể, tạo nền tảng cho các chứng nhận khác
- **SOC 2 Type 2:** Yêu cầu 6 tháng vận hành, phù hợp sau khi v1 stable
- **GDPR:** 70% yêu cầu được đáp ứng tự động khi có ISO 27001 + SOC 2

**Lộ trình:**
```
Tháng 6: SOC 2 Type 1 (cùng v1 release)
    ↓
Tháng 8: ISO 27001 Certificate
    ↓
Tháng 12: SOC 2 Type 2 (sau 6 tháng vận hành)
    ↓
Ongoing: GDPR compliance maintenance
```

**Alternatives Considered:**
- ❌ **Chỉ ISO 27001:** Không đủ cho khách hàng Hoa Kỳ (cần SOC 2)
- ❌ **Chỉ SOC 2:** Không đủ cho khách hàng quốc tế (cần ISO 27001)
- ❌ **GDPR trước:** Quá phức tạp, không có framework tổng thể

### 2. Công Cụ Automation

**Quyết định:** Sử dụng Vanta (gói Startups) làm nền tảng quản lý tuân thủ trung tâm.

**Lý do:**
- **Tích hợp API:** Kết nối trực tiếp AWS, GitHub, Cloudflare
- **Tự động thu thập bằng chứng:** 24/7 monitoring, giảm 80% công việc thủ công
- **Audit-ready:** Dashboard sẵn sàng cho kiểm toán viên
- **Cost-effective:** $4K-6K/năm, phù hợp ngân sách startup
- **Proven:** Được nhiều startup sử dụng thành công

**Alternatives Considered:**
- ❌ **Drata:** Tương tự Vanta nhưng đắt hơn (~$8K-10K/năm)
- ❌ **Manual compliance:** Quá tốn thời gian, dễ sai sót
- ❌ **Secureframe:** Ít tích hợp hơn với AWS/GitHub

**Integration Points:**
```
Vanta Hub
    ├─ AWS (CloudTrail, IAM, S3, RDS)
    ├─ GitHub (repos, security, access)
    ├─ Cloudflare (R2 buckets, access logs)
    └─ HR Systems (employee onboarding/offboarding)
```

### 3. Bảo Mật Mã Nguồn

**Quyết định:** Tích hợp GitHub Dependabot + GitHub Advanced Security (CodeQL) vào GitHub Actions.

**Lý do:**
- **Native integration:** Tích hợp sẵn với GitHub, không cần setup phức tạp
- **Miễn phí:** Cho public repos và private repos (với GitHub Team/Enterprise)
- **Automated:** Chạy tự động trên mỗi PR
- **Blocking:** Block PR nếu phát hiện lỗi High/Critical
- **Comprehensive:**
  - **Dependabot:** Quét thư viện bên thứ 3 (SCA - Software Composition Analysis)
  - **CodeQL:** Quét mã nguồn tự viết (SAST - Static Application Security Testing)

**CI/CD Pipeline:**
```yaml
# .github/workflows/security.yml
on: [pull_request]
jobs:
  security:
    - Dependabot scan (dependencies)
    - CodeQL analysis (source code)
    - Secret scanning (credentials)
    → Block merge if High/Critical found
```

**Alternatives Considered:**
- ❌ **Snyk:** Tốt nhưng có chi phí (~$500-1K/năm)
- ❌ **SonarQube:** Quá phức tạp cho team nhỏ
- ❌ **Manual code review only:** Không đủ coverage, dễ miss

### 4. Hạ Tầng Tuân Thủ (Infrastructure as Code)

**Quyết định:** Áp dụng nguyên tắc "Compliance-by-Design" - mọi tài nguyên phải tuân thủ ngay từ khi khởi tạo.

**Lý do:**
- **Shift-left security:** Bảo mật từ đầu, không phải fix sau
- **Consistency:** Đảm bảo mọi resource đều tuân thủ
- **Audit trail:** IaC code là documentation và audit trail
- **Automation:** Vanta có thể quét IaC code tự động

**Nguyên Tắc Bắt Buộc:**

1. **Encryption at Rest:**
   ```terraform
   # ✅ Đúng
   resource "aws_s3_bucket" "data" {
     server_side_encryption_configuration {
       rule {
         apply_server_side_encryption_by_default {
           sse_algorithm = "AES256"
         }
       }
     }
   }
   
   # ❌ Sai - không có encryption
   resource "aws_s3_bucket" "data" {
     # missing encryption
   }
   ```

2. **Logging Enabled:**
   ```terraform
   # ✅ Đúng
   resource "aws_cloudtrail" "main" {
     enable_logging = true
     s3_bucket_name = aws_s3_bucket.logs.id
   }
   ```

3. **Private by Default:**
   ```terraform
   # ✅ Đúng
   resource "aws_db_instance" "main" {
     publicly_accessible = false
     db_subnet_group_name = aws_db_subnet_group.private.name
   }
   ```

4. **Least Privilege IAM:**
   ```terraform
   # ✅ Đúng - specific permissions
   resource "aws_iam_policy" "app" {
     policy = jsonencode({
       Statement = [{
         Effect = "Allow"
         Action = ["s3:GetObject", "s3:PutObject"]
         Resource = "${aws_s3_bucket.data.arn}/*"
       }]
     })
   }
   
   # ❌ Sai - quá rộng
   Action = ["s3:*"]
   ```

**Validation:**
- Pre-commit hooks: Validate IaC trước khi commit
- CI/CD checks: Scan IaC trong pipeline
- Vanta integration: Continuous compliance monitoring

**Alternatives Considered:**
- ❌ **Manual configuration:** Không consistent, dễ sai sót
- ❌ **Fix sau khi deploy:** Tốn thời gian, có thể miss resources
- ❌ **Không có IaC:** Không có audit trail, khó maintain

## Consequences (Hệ quả)

### Positive
1. **Tốc độ:** Automation giảm 80% effort cho compliance
2. **Chất lượng:** Compliance-by-design đảm bảo consistency
3. **Chi phí:** $12K-19K/năm, phù hợp ngân sách startup
4. **Credibility:** Có chứng nhận quốc tế tăng độ tin cậy
5. **Scalability:** Framework sẵn sàng cho growth

### Negative
1. **Learning curve:** Team cần học Vanta, IaC best practices
2. **Initial setup:** 1-2 tháng setup ban đầu
3. **Maintenance:** Cần maintain policies, runbooks
4. **Audit cost:** $8K-13K cho audits (SOC 2 + ISO 27001)

### Risks & Mitigation

| Risk | Impact | Probability | Mitigation |
|------|--------|-------------|------------|
| Vanta không đủ coverage | High | Low | Bổ sung manual checks, có fallback plan |
| Audit fail lần đầu | Medium | Medium | Mock audit trước, fix issues sớm |
| Team overload | Medium | High | Phân bổ rõ ràng, không rush |
| Cost overrun | Low | Low | Budget buffer 20%, track monthly |

## Implementation Notes

### Phase 1: Foundation (Tháng 1-4)
- Setup Vanta account
- Connect AWS, GitHub, Cloudflare
- Enable Dependabot + CodeQL
- Write IaC templates với compliance rules
- Create initial policies & runbooks

### Phase 2: Certification (Tháng 5-8)
- Complete Vanta dashboard
- Mock audit với Vanta
- Invite SOC 2 auditor (Tháng 6)
- Invite ISO 27001 auditor (Tháng 7-8)

### Phase 3: Maintenance (Tháng 9-12)
- Continuous monitoring
- Quarterly policy review
- Prepare for SOC 2 Type 2 (Tháng 12)

## References
- [Vanta Documentation](https://www.vanta.com/resources)
- [AWS Security Best Practices](https://aws.amazon.com/security/best-practices/)
- [GitHub Security Features](https://docs.github.com/en/code-security)
- [ISO 27001 Standard](https://www.iso.org/isoiec-27001-information-security.html)
- [SOC 2 Trust Services Criteria](https://www.aicpa.org/interestareas/frc/assuranceadvisoryservices/aicpasoc2report)

## Approval
- **Approved by:** [Founder/CTO Name]
- **Date:** 2026-05-27
- **Next review:** 2026-11-27 (6 months)
