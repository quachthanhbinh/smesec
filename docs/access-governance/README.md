# Tài Liệu Chiến Lược: Access Governance

## Tổng Quan

Bộ tài liệu này mô tả chiến lược quản trị truy cập (access governance) cho hệ thống SMESec, bao gồm least-privilege enforcement, offboarding automation, và shadow IT detection.

**Thông tin dự án:**
- **Người phụ trách:** Quách Thanh Bình - Senior Backend Engineer
- **Giai đoạn hệ thống:** Build from scratch (Dự kiến v1 sau 6 tháng)
- **Ngày cập nhật:** Tháng 5/2026
- **Thời gian triển khai:** 6 tháng (song song với v1)

## Mục Tiêu

Xây dựng hệ thống quản trị truy cập toàn diện để:
- Thực thi nguyên tắc least privilege tự động
- Tự động hóa quy trình offboarding nhân viên
- Phát hiện và quản lý shadow IT
- Đảm bảo compliance với access control policies
- Giảm thiểu rủi ro từ excessive permissions

## Phạm Vi Quản Trị

### 1. Least-Privilege Enforcement
- **Role-Based Access Control (RBAC)**: Phân quyền theo vai trò
- **Just-In-Time (JIT) Access**: Cấp quyền tạm thời khi cần
- **Privilege Escalation Control**: Kiểm soát nâng quyền
- **Access Reviews**: Định kỳ review và revoke permissions

### 2. Offboarding Automation
- **Account Deactivation**: Tự động vô hiệu hóa accounts
- **Access Revocation**: Thu hồi tất cả quyền truy cập
- **Data Transfer**: Chuyển giao dữ liệu cho người kế nhiệm
- **Device Wipe**: Xóa dữ liệu từ thiết bị công ty

### 3. Shadow IT Detection
- **Unauthorized SaaS**: Phát hiện ứng dụng không được phê duyệt
- **Personal Accounts**: Phát hiện tài khoản cá nhân được dùng cho công việc
- **Data Exfiltration**: Giám sát dữ liệu được upload lên dịch vụ ngoài
- **Compliance Violations**: Phát hiện vi phạm chính sách IT

### 4. Identity & Access Management (IAM)
- **Single Sign-On (SSO)**: Tích hợp SSO cho tất cả ứng dụng
- **Multi-Factor Authentication (MFA)**: Bắt buộc MFA cho tất cả users
- **Password Policies**: Chính sách mật khẩu mạnh
- **Session Management**: Quản lý phiên đăng nhập

## Cấu Trúc Tài Liệu

### [01. Architecture Decision Record (ADR)](01-adr.md)
Ghi nhận các quyết định kiến trúc về:
- IAM platform selection (Okta, Auth0, Azure AD)
- RBAC vs ABAC (Attribute-Based Access Control)
- JIT access implementation approach
- Offboarding workflow automation

### [02. Least-Privilege Framework](02-least-privilege.md)
Chi tiết về:
- Role definitions và permission mappings
- JIT access request/approval workflows
- Privilege escalation policies
- Access review schedules và procedures

### [03. Offboarding Automation Strategy](03-offboarding.md)
Phương pháp tự động hóa:
- Trigger mechanisms (HR system integration)
- Account deactivation workflows
- Access revocation checklists
- Data transfer procedures

### [04. Shadow IT Detection & Management](04-shadow-it.md)
Chiến lược phát hiện và quản lý:
- Discovery methods (network monitoring, browser extensions)
- Risk assessment framework
- Remediation workflows
- User education programs

### [05. Lộ Trình Triển Khai](05-roadmap.md)
Timeline chi tiết 6 tháng:
- **Tháng 1-2**: IAM foundation + SSO/MFA
- **Tháng 3-4**: RBAC + JIT access + offboarding automation
- **Tháng 5-6**: Shadow IT detection + access reviews

### [06. Technical Implementation Guide](06-technical-guide.md)
Hướng dẫn kỹ thuật:
- IAM platform integration
- RBAC policy engine implementation
- Offboarding workflow automation
- Shadow IT detection mechanisms

### [07. Phân Bổ Nguồn Lực](07-resources.md)
Kế hoạch nguồn lực:
- Team roles (IAM admin, security analyst)
- Tools và services cần thiết
- Training programs
- Ngân sách ước tính

## Công Nghệ & Công Cụ

### Identity & Access Management
- **IAM Platform**: Okta, Auth0, Azure AD, Google Workspace
- **SSO Protocol**: SAML 2.0, OAuth 2.0, OpenID Connect
- **MFA Solutions**: Duo Security, Okta Verify, Google Authenticator
- **Privileged Access Management (PAM)**: CyberArk, BeyondTrust (for larger SMEs)

### Access Governance
- **RBAC Engine**: Open Policy Agent (OPA), Casbin
- **JIT Access**: Teleport, StrongDM, Okta Advanced Server Access
- **Access Reviews**: Okta Lifecycle Management, SailPoint (enterprise)
- **Audit Logging**: Splunk, ELK Stack, AWS CloudTrail

### Offboarding Automation
- **HR Integration**: BambooHR, Workday, Gusto APIs
- **Workflow Automation**: Zapier, n8n, custom scripts
- **Device Management**: Jamf (Mac), Intune (Windows), Google Workspace
- **Data Transfer**: Google Takeout, Microsoft 365 Admin Center

### Shadow IT Detection
- **Cloud Access Security Broker (CASB)**: Netskope, McAfee MVISION Cloud
- **Browser Extensions**: Custom monitoring extensions
- **Network Monitoring**: Cisco Umbrella, Zscaler
- **SaaS Discovery**: Torii, BetterCloud, Productiv

## Nguyên Tắc Chính

### 1. Least Privilege by Default
Mọi user chỉ có quyền tối thiểu cần thiết, không có quyền admin mặc định.

### 2. Zero Standing Privileges
Không có quyền vĩnh viễn cho privileged access, chỉ JIT khi cần.

### 3. Automated Enforcement
Policies được enforce tự động, không phụ thuộc vào manual processes.

### 4. Continuous Monitoring
Giám sát liên tục access patterns và phát hiện anomalies.

## Metrics & KPIs

| Metric | Target | Measurement |
|--------|--------|-------------|
| MFA Adoption Rate | 100% | % of users with MFA enabled |
| Offboarding Completion Time | <4 hours | Time from termination to full access revocation |
| Excessive Permissions | <5% | % of users with more permissions than needed |
| Shadow IT Apps Detected | >80% | % of unauthorized apps discovered |
| Access Review Completion | 100% | % of quarterly reviews completed on time |
| JIT Access Approval Time | <15 minutes | Average time to approve JIT requests |

## Ngân Sách Ước Tính

| Hạng mục | Chi phí/năm |
|----------|-------------|
| IAM Platform (Okta/Auth0) | $5,000 - $8,000 |
| MFA Solution | $2,000 - $3,000 |
| JIT Access Tool | $3,000 - $5,000 |
| Shadow IT Detection (CASB) | $4,000 - $6,000 |
| Workflow Automation | $1,000 - $2,000 |

**Tổng ước tính năm đầu:** ~$15,000 - $24,000

## Milestone Chính

- **Milestone 1 (Tháng 2)**: SSO + MFA deployed for all users
- **Milestone 2 (Tháng 4)**: RBAC + offboarding automation operational
- **Milestone 3 (Tháng 6)**: Shadow IT detection + JIT access live

## Use Cases Cụ Thể cho SMEs

### Scenario 1: Employee Termination
**Challenge**: Employee bị sa thải, cần revoke access ngay lập tức.
**Solution**: HR marks employee as terminated → Automated workflow deactivates all accounts, revokes access, wipes devices trong 4 giờ.

### Scenario 2: Contractor Temporary Access
**Challenge**: Contractor cần access vào production database để debug trong 2 giờ.
**Solution**: Contractor requests JIT access → Manager approves → Access granted for 2 hours → Auto-revoked sau đó.

### Scenario 3: Shadow IT Discovery
**Challenge**: Employees dùng personal Dropbox để share customer data.
**Solution**: CASB detects unauthorized Dropbox usage → Alert security team → Block access → Migrate to approved solution.

### Scenario 4: Quarterly Access Review
**Challenge**: Cần review và cleanup excessive permissions.
**Solution**: Automated report lists all users with admin rights → Managers review → Revoke unnecessary permissions → Audit trail maintained.

## Liên Hệ & Hỗ Trợ

**Người phụ trách:** Quách Thanh Bình  
**Email:** [Thêm email]  
**Slack:** [Thêm channel]

## Tài Liệu Tham Khảo

- [NIST SP 800-53 Access Control](https://csrc.nist.gov/publications/detail/sp/800-53/rev-5/final)
- [CIS Controls - Access Control Management](https://www.cisecurity.org/controls)
- [OWASP Access Control Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Access_Control_Cheat_Sheet.html)
- [Okta Identity Governance Best Practices](https://www.okta.com/identity-governance/)
- [Zero Standing Privileges (ZSP) Framework](https://www.cyberark.com/what-is/zero-standing-privileges/)

---

**Lưu ý:** Tài liệu này là living document và sẽ được cập nhật thường xuyên theo tiến độ dự án.
