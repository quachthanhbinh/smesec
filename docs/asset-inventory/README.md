# Tài Liệu Chiến Lược: Asset Inventory & Classification

## Tổng Quan

Bộ tài liệu này mô tả chiến lược quản lý và phân loại tài sản (assets) cho hệ thống SMESec, bao gồm data, devices, accounts, và third-party integrations.

**Thông tin dự án:**
- **Người phụ trách:** Quách Thanh Bình - Senior Backend Engineer
- **Giai đoạn hệ thống:** Build from scratch (Dự kiến v1 sau 6 tháng)
- **Ngày cập nhật:** Tháng 5/2026
- **Thời gian triển khai:** 6 tháng (song song với v1)

## Mục Tiêu

Xây dựng hệ thống quản lý tài sản toàn diện để:
- Tự động phát hiện và phân loại tất cả tài sản trong tổ chức SME
- Theo dõi vòng đời tài sản từ khởi tạo đến hủy bỏ
- Đánh giá rủi ro và độ quan trọng của từng tài sản
- Hỗ trợ compliance và audit requirements
- Phát hiện shadow IT và tài sản không được quản lý

## Phạm Vi Tài Sản

### 1. Data Assets
- **Structured Data**: Databases, data warehouses
- **Unstructured Data**: Files, documents, emails
- **Sensitive Data**: PII, financial records, IP
- **Data Flows**: Data movement between systems

### 2. Device Assets
- **Endpoints**: Laptops, desktops, mobile devices
- **Servers**: Physical and virtual servers
- **Network Devices**: Routers, switches, firewalls
- **IoT Devices**: Smart devices, sensors

### 3. Account Assets
- **User Accounts**: Employee accounts across systems
- **Service Accounts**: API keys, service principals
- **Admin Accounts**: Privileged access accounts
- **External Accounts**: Contractor, vendor access

### 4. Third-Party Integrations
- **SaaS Applications**: Google Workspace, Microsoft 365, Slack
- **APIs**: External API connections
- **Cloud Services**: AWS, Azure, GCP resources
- **Vendor Systems**: QuickBooks, CRM, etc.

## Cấu Trúc Tài Liệu

### [01. Architecture Decision Record (ADR)](01-adr.md)
Ghi nhận các quyết định kiến trúc về:
- Discovery mechanisms (agent-based vs agentless)
- Classification taxonomy và metadata schema
- Storage và indexing strategy
- Integration patterns với existing systems

### [02. Asset Classification Framework](02-classification-framework.md)
Chi tiết về:
- Asset types và categories
- Criticality levels (Critical, High, Medium, Low)
- Sensitivity classifications (Public, Internal, Confidential, Restricted)
- Risk scoring methodology

### [03. Discovery & Collection Strategy](03-discovery-strategy.md)
Phương pháp thu thập thông tin:
- Automated discovery mechanisms
- Integration với cloud providers (AWS, Azure, GCP)
- SaaS application discovery
- Network scanning và endpoint detection

### [04. Lộ Trình Triển Khai](04-roadmap.md)
Timeline chi tiết 6 tháng:
- **Tháng 1-2**: Core inventory system + basic discovery
- **Tháng 3-4**: Classification engine + risk scoring
- **Tháng 5-6**: Advanced integrations + reporting

### [05. Technical Implementation Guide](05-technical-guide.md)
Hướng dẫn kỹ thuật:
- Database schema cho asset inventory
- API design cho asset management
- Discovery agents implementation
- Integration với compliance module

### [06. Phân Bổ Nguồn Lực](06-resources.md)
Kế hoạch nguồn lực:
- Team roles và responsibilities
- Tools và services cần thiết
- Ngân sách ước tính

## Công Nghệ & Công Cụ

### Discovery Tools
- **Cloud Asset Discovery**: AWS Config, Azure Resource Graph, GCP Asset Inventory
- **SaaS Discovery**: BetterCloud, Torii, Productiv
- **Network Discovery**: Nmap, Lansweeper
- **Endpoint Discovery**: Osquery, Fleet

### Asset Management Platform
- **Core Database**: PostgreSQL với JSONB cho flexible metadata
- **Search & Indexing**: Elasticsearch cho fast querying
- **API Layer**: GraphQL cho flexible data access
- **UI Dashboard**: React-based inventory dashboard

### Integration Framework
- **Cloud Providers**: AWS SDK, Azure SDK, GCP SDK
- **SaaS Apps**: OAuth 2.0 integrations
- **Identity Providers**: SCIM, LDAP, Azure AD
- **CMDB Integration**: ServiceNow, Jira Service Management

## Nguyên Tắc Chính

### 1. Continuous Discovery
Tự động phát hiện tài sản mới 24/7, không phụ thuộc vào manual input.

### 2. Automated Classification
Sử dụng ML và rule-based systems để tự động phân loại tài sản.

### 3. Real-Time Visibility
Dashboard cập nhật real-time về trạng thái tài sản và rủi ro.

### 4. Integration-First
Tích hợp sâu với existing tools thay vì yêu cầu thay đổi workflow.

## Metrics & KPIs

| Metric | Target | Measurement |
|--------|--------|-------------|
| Asset Discovery Coverage | >95% | % of actual assets discovered |
| Classification Accuracy | >90% | % correctly classified assets |
| Time to Discovery | <24 hours | Time from asset creation to discovery |
| Shadow IT Detection | >80% | % of unauthorized apps detected |
| Data Freshness | <1 hour | Age of asset information |

## Ngân Sách Ước Tính

| Hạng mục | Chi phí/năm |
|----------|-------------|
| Discovery Tools (Osquery, Fleet) | $2,000 - $3,000 |
| SaaS Discovery Platform | $3,000 - $5,000 |
| Cloud Asset Management APIs | $1,000 - $2,000 |
| Storage & Infrastructure | $2,000 - $3,000 |

**Tổng ước tính năm đầu:** ~$8,000 - $13,000

## Milestone Chính

- **Milestone 1 (Tháng 2)**: Core inventory system operational
- **Milestone 2 (Tháng 4)**: Automated classification working
- **Milestone 3 (Tháng 6)**: Full integration với compliance module

## Liên Hệ & Hỗ Trợ

**Người phụ trách:** Quách Thanh Bình  
**Email:** [Thêm email]  
**Slack:** [Thêm channel]

## Tài Liệu Tham Khảo

- [NIST Cybersecurity Framework - Asset Management](https://www.nist.gov/cyberframework)
- [ISO 27001 Asset Management Controls](https://www.iso.org/standard/27001)
- [CIS Controls - Inventory and Control of Assets](https://www.cisecurity.org/controls)
- [SANS Asset Management Best Practices](https://www.sans.org/)

---

**Lưu ý:** Tài liệu này là living document và sẽ được cập nhật thường xuyên theo tiến độ dự án.
