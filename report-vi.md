# Nền tảng SMESec — Tài liệu Thiết kế Hệ thống

**Ngày:** 28-05-2026 | **Phiên bản:** 2.0 | **Trạng thái:** Hoàn tất  
**Nguồn:** Tổng hợp từ nghiên cứu đa đại lý (Chủ sở hữu Sản phẩm · Quản lý Dự án · Cố vấn Kỹ thuật)

---

## Tóm tắt dự án (Executive Summary)

Các doanh nghiệp vừa và nhỏ (10–500 nhân viên) đang phải đối mặt với các rủi ro bảo mật leo thang do AI thúc đẩy — tấn công giả mạo có chủ đích tự động (automated spear-phishing), rò rỉ dữ liệu của nhân viên sang các LLM công cộng, sự bùng nổ các công cụ AI bóng tối (shadow AI), gian lận deepfake và xâm nhập chuỗi cung ứng — nhưng lại thiếu các đội ngũ bảo mật chuyên trách và ngân sách dành cho doanh nghiệp lớn. **SMESec** là một nền tảng bảo vệ SaaS thống nhất bao phủ toàn bộ bề mặt tài sản của doanh nghiệp vừa và nhỏ (SME): dữ liệu, con người, sở hữu trí tuệ, tài khoản tài chính và tính liên tục trong vận hành.

**Chiến lược hai luồng (Two-Track Strategy):** Tất cả hoạt động phát triển được chia thành các luồng song song để loại bỏ rủi ro về độ chính xác của việc phát hiện bằng AI.

- **Luồng 1 — Nền tảng & Quản trị (mô hình tất định/deterministic, độ chính xác ~100%):** Quản lý tài sản (Asset inventory), quản trị truy cập (access governance), tự động hóa quy trình cho thôi việc (automated offboarding), kịch bản ứng phó sự cố (incident playbooks), báo cáo tuân thủ. Phát hành độc lập tại phiên bản MVP (Tháng 3) và v1 (Tháng 6).
- **Luồng 2 — Phát hiện Mối đe dọa bằng AI (kiểm soát bởi ML):** Phòng chống mất mát dữ liệu (DLP) trên trình duyệt, quản trị shadow AI, phòng chống deepfake, phát hiện tấn công chèn lệnh (prompt injection). **Bắt đầu từ Sprint 1, song song với Luồng 1.** Kỹ sư ML #1 gia nhập ngay Ngày 1 để bắt đầu R&D (nghiên cứu, thu thập bộ dữ liệu, xây dựng mô hình mẫu). Chỉ tích hợp vào sản phẩm sau khi vượt qua bốn cổng kiểm định độ chính xác. Nếu không đạt yêu cầu, Luồng 1 sẽ được phát hành độc lập.

**Các tùy chọn Timeline:** Có bốn kế hoạch bàn giao. **Kế hoạch điều chỉnh 1.5x là kế hoạch được khuyến nghị:**
- **✅ Điều chỉnh 1.5x (19.5 tháng) — KHUYẾN NGHỊ:** Khởi đầu gọn nhẹ 5 nhân sự → phiên bản thương mại v1 (chỉ Luồng 1 + tính năng thanh toán Stripe) tại Tháng 4.5 → v1 đầy đủ với các tính năng AI tại Tháng 9.75. Mức sử dụng sprint 60–75%. Cân bằng tốt nhất giữa tốc độ ra thị trường và kiểm soát rủi ro.
- **Gốc (12 tháng):** Kế hoạch gốc KHÔNG KHUYẾN NGHỊ: yêu cầu 7 FTE từ Ngày 1 (rất khó tuyển dụng), sử dụng sprint 75–90% (có sprint lên tới 88%), Kỹ sư ML #1 phải có ngay Ngày 1, không có buffer cho SOC 2 Type 2 (chỉ vừa đủ 24 tuần). Rủi ro kiệt sức, burnout và cascade-failure rất cao.
- **Điều chỉnh 2x (26 tháng):** Timeline bền vững, 50–60% sử dụng sprint, phù hợp với đội ngũ cần thêm thời gian tuyển dụng.
- **Tuyển dụng thực tế (36+ tháng):** Trưởng nhóm kỹ thuật bắt đầu một mình, xây dựng đội ngũ dần dần, Kỹ sư ML #1 gia nhập Tháng 8.

---

## 1. Sơ đồ Kiến trúc Hệ thống

### 1.1 Kiến trúc Logic — Các lớp Kiến trúc Sạch (Clean Architecture Layers)

SMESec áp dụng **Kiến trúc Sạch (Clean Architecture)** (Robert C. Martin) + **Kiến trúc Lục giác (Hexagonal Architecture / Ports & Adapters)**. Quy tắc Phụ thuộc (Dependency Rule) bắt buộc: `Interface → Application → Domain ← Infrastructure`. Lớp Domain không có bất kỳ phụ thuộc bên ngoài nào.

```
┌──────────────────────────────────────────────────────────────────────┐
│  LỚP GIAO DIỆN (INTERFACE LAYER)                                      │
│  Web App (React/Next.js) · Mobile App (Flutter) · Browser Ext (MV3)  │
│  REST/gRPC/WebSocket ← API Gateway (AWS) + Keycloak JWT auth         │
├──────────────────────────────────────────────────────────────────────┤
│  LỚP ỨNG DỤNG (APPLICATION LAYER - Use Cases)                        │
│  AssetInventorySvc · AccessGovernanceSvc · IncidentPlaybookSvc        │
│  ComplianceSvc · IntegrationSyncSvc · ThreatDetectionSvc (Luồng 2)   │
├──────────────────────────────────────────────────────────────────────┤
│  LỚP NGHIỆP VỤ (DOMAIN LAYER - Không phụ thuộc bên ngoài)             │
│  Entities: Asset · TenantUser · ThreatEvent · Playbook · AccessPolicy│
│  Domain Services: RiskScorer · AccessGovernor · ComplianceAuditor     │
│  Domain Events: AssetDiscovered · ThreatDetected · AccessRevoked      │
├──────────────────────────────────────────────────────────────────────┤
│  LỚP HẠ TẦNG (INFRASTRUCTURE LAYER - Adapters triển khai Domain ports) │
│  PostgreSQL Repos (RLS) · GoogleWorkspaceAdapter · M365Adapter        │
│  SlackAdapter · AWSIAMAdapter · EventBridgePublisher · HiveClient     │
│  VantaClient · SageMakerClient · SecretsManagerClient                 │
└──────────────────────────────────────────────────────────────────────┘

                 Luồng 1 và Luồng 2 chia sẻ:
           ThreatDetectionEvent schema + EventBridge event bus
           Sự kiện Luồng 2 có thể kích hoạt playbook Luồng 1.
           Luồng 1 không bao giờ phụ thuộc vào tính sẵn sàng của Luồng 2.
```

### 1.2 Kiến trúc Triển khai — AWS Multi-Region

```
INTERNET
  │ HTTPS (Chỉ TLS 1.3)
  ▼
VÙNG BIÊN (EDGE ZONE)
  Route 53 (GeoDNS: US → us-east-1, EU → eu-west-1)
  → CloudFront CDN → WAF (OWASP rules) → ALB

AWS VPC (Chỉ các mạng con riêng tư - không sử dụng IP công cộng cho các máy chủ tính toán)
  ├── XÁC THỰC: Keycloak ECS Fargate (Tối thiểu 2 tác vụ active-active, JWKS được lưu cache 6 giờ — xác thực JWT độc lập với thời gian hoạt động của Keycloak) [R-C6]
  │
  ├── ỨNG DỤNG — Dịch vụ ECS Fargate (Go):
  │     Luồng 1: AssetSvc · AccessSvc · PlaybookSvc · ComplianceSvc · SyncSvc
  │     Luồng 2: ThreatDetectionSvc · DLPSvc · DeepfakeSvc (Python/FastAPI)
  │
  ├── DỮ LIỆU:
  │     RDS PostgreSQL Multi-AZ (Row-Level Security, tenant_id trên mọi bảng)
  │     ElastiCache Redis (session tokens, TTL 15 phút)
  │     S3 Object Lock (WORM, lưu trữ nhật ký kiểm toán trong 7 năm)
  │
  └── CÁC DỊCH VỤ AWS ĐƯỢC QUẢN LÝ (ngoài VPC):
        EventBridge · Step Functions · SNS/SQS
        SageMaker (ML training + inference, Luồng 2)
        Secrets Manager · KMS (CMK trên mỗi vùng) · GuardDuty · Security Hub
        CloudWatch · CloudTrail · IAM

ỨNG DỤNG PHÍA KHÁCH (CLIENTS):
  Web Dashboard (Next.js) · Mobile App (Flutter iOS+Android) · Browser Extension (Chrome MV3 + Edge)
```

**Công nghệ sử dụng (Technology Stack):**
- **Backend:** Go (các dịch vụ API chính, đồng bộ hóa tích hợp) · Python/FastAPI (các dịch vụ ML/AI)
- **Frontend:** React/Next.js (web) · Flutter (iOS, Android) · Chrome MV3 (browser extension)
- **Auth (Xác thực):** Keycloak (tự lưu trữ trên ECS, OIDC/SAML 2.0, bắt buộc sử dụng MFA TOTP, JWT RS256)
- **ML (Học máy):** AWS SageMaker (mô hình rủi ro shadow AI, bộ phân loại prompt injection dựa trên BERT-tiny)
- **Tự động hóa Tuân thủ (Compliance Automation):** Vanta (kết nối AWS + GitHub, thu thập minh chứng SOC 2 + ISO 27001)

### 1.3 Các Điểm chạm Tích hợp (Integration Touchpoints)

| Dịch vụ | Phương thức | OAuth Scopes (tối thiểu) | Tần suất | Tính năng được kích hoạt |
|---|---|---|---|---|
| **Google Workspace** | OAuth 2.0 + Admin SDK | `admin.directory.user.readonly` `admin.directory.userschema.readonly` `admin.reports.audit.readonly` | Đồng bộ hóa chênh lệch mọi 15 phút. **⚠️ R-C2 (Bắt buộc từ Sprint 1):** Quota = 1,500 yêu cầu/100 giây cho mỗi dự án GCP. v1 nhắm mục tiêu 1K tenant — cần 50 dự án GCP (1K / 20 tenant mỗi dự án). Triển khai GCP project pool từ Sprint 1; cột `gcp_project_id` trong schema `tenant_config` từ ngày đầu. SyncScheduler gán project cho tenant khi onboarding. | Kho lưu trữ người dùng, phát hiện ứng dụng OAuth, phát hiện shadow IT, quy trình thôi việc (offboarding) |
| **Microsoft 365** | OAuth 2.0 + Graph API + webhook | `User.Read.All` `Application.Read.All` `AuditLog.Read.All` `SecurityEvents.Read.All` | Đồng bộ chênh lệch mỗi 15 phút + webhook. **⚠️ R-C3:** Đăng ký Webhook hết hạn sau mỗi **3 ngày** — cần công việc gia hạn định kỳ (EventBridge Scheduler, mỗi 12 giờ) + cơ chế dự phòng đồng bộ toàn phần khi gặp lỗi 410 Gone + giao diện cảnh báo dữ liệu cũ. Schema được thiết kế trong S1. | Kho lưu trữ người dùng, ứng dụng OAuth, cảnh báo lừa đảo M365 Defender, quy trình thôi việc (offboarding) |
| **Slack** | OAuth 2.0 + Admin API | `admin.users:read` `admin.apps:read` `channels:read` | Đồng bộ chênh lệch mỗi 30 phút | Kho lưu trữ ứng dụng, vô hiệu hóa người dùng (gói Business+ trở lên), kiểm toán kênh |
| **AWS IAM** | Giả lập vai trò IAM (liên tài khoản) | `iam:ListUsers` `iam:ListRoles` `cloudtrail:LookupEvents` `config:ListDiscoveredResources` | Đồng bộ chênh lệch mỗi 30 phút | Kho lưu trữ tài nguyên đám mây, so sánh chính sách IAM, sự kiện CloudTrail |
| **Hive Moderation API** | REST (trả phí theo lượt sử dụng) | Khóa API (Secrets Manager) | Theo yêu cầu | Phát hiện giọng nói/video deepfake (<$0.01/lượt kiểm tra) |
| **Vanta** | Trình kết nối AWS + GitHub gốc | Chỉ đọc cho mục đích SOC 2 | Liên tục | Thu thập bằng chứng tuân thủ, cổng thông tin cho kiểm toán viên |

**Mô hình bảo mật tích hợp:** Tất cả các mã thông báo OAuth (OAuth tokens) được lưu trữ trong AWS Secrets Manager (mã hóa AES-256, tự động xoay vòng). Mặc định là chỉ đọc; quyền ghi (thu hồi truy cập) được yêu cầu riêng biệt và phải có sự đồng ý rõ ràng của quản trị viên CNTT. Mỗi cuộc gọi API đều được ghi nhật ký kèm theo thông tin `tenant_id + user_id + hành động + mốc thời gian`.

---

## 2. Tài liệu Thiết kế — Các Quyết định Kiến trúc Cốt lõi

### 2.1 Tự xây dựng (Build) so với Mua ngoài (Buy) (Mô hình Lai)

**Quyết định:** Tự xây dựng các yếu tố khác biệt cốt lõi (bất kỳ tính năng nào khách hàng sẵn sàng trả tiền); mua các dịch vụ phổ thông (bất kỳ thứ gì mất >3 tháng để xây dựng nhưng chi phí thuê ngoài chỉ <$5K/năm).

| Thành phần | Quyết định | Đối tác / Công nghệ | Lý do lựa chọn |
|---|---|---|---|
| **Kho lưu trữ tài sản & Công cụ đồng bộ** | **Tự xây dựng** (Go) | Google Admin SDK, Graph API | Logic phát hiện Shadow IT là lợi thế cạnh tranh cốt lõi; không có đối thủ nào cung cấp tính năng này ở mức giá cho SME. |
| **Quản trị truy cập (RBAC + JIT)** | **Tự xây dựng** (Go) | Chính sách OPA/Rego | Tự động hóa quy trình thôi việc tối ưu cho SME là điểm khác biệt chính so với Vanta/Drata. |
| **Công cụ kịch bản ứng phó sự cố** | **Xây dựng trên Step Functions** | AWS Step Functions | Step Functions là công cụ điều phối đã được kiểm chứng; giao diện wizard trực quan cho nhân sự không thuộc mảng bảo mật là điểm khác biệt. |
| **Phòng chống mất mát dữ liệu (DLP) trên trình duyệt** | **Tự xây dựng** (Chrome MV3) | Microsoft Presidio WASM | Suy luận PII cục bộ — nội dung không bao giờ rời khỏi trình duyệt. Rào cản quyền riêng tư mà không đối thủ nào sánh được ở mức giá SME. |
| **Phân loại rủi ro công cụ AI** | **Xây dựng & Tuyển chọn** | SageMaker + cơ sở dữ liệu nội bộ | Không có giải pháp tính điểm rủi ro chuyên biệt cho AI sẵn có nào phù hợp với ngữ cảnh và giá cả cho SME. |
| **SSO / MFA** | **Mua ngoài: Keycloak** (tự lưu trữ trên ECS) | Keycloak | Chi phí bằng 0 trên mỗi người dùng (~$150/tháng cho 4 tác vụ ECS) so với Auth0 (~$115,000+/tháng với 1K tenant × 500 người dùng = 500K MAU × $0,23/MAU). Keycloak tiết kiệm hơn ~$500K+/năm ở quy mô mục tiêu v1. **⚠️ Yêu cầu R-C6:** Tối thiểu 2 tác vụ ECS chạy song song (active-active); bắt buộc lưu cache JWKS; cơ sở dữ liệu Keycloak phải tách biệt với cơ sở dữ liệu ứng dụng. **Đánh giá WorkOS/Auth0 trước phiên bản v1.5** nếu năng lực DevSecOps không đủ đáp ứng. |
| **Phát hiện tấn công chèn lệnh (Prompt injection)** | **Mua ngoài: API Lakera Guard (v1)** | Lakera Guard | Đã được kiểm chứng thực tế (~$0.001/yêu cầu). Không cần dữ liệu huấn luyện. Mục tiêu tự xây dựng mô hình BERT nội bộ được chuyển sang kiểm tra dành riêng cho gói Enterprise tại Sprint 23–24. Điều kiện (Gate): Tỷ lệ dương tính giả (FPR) <2% + Tỷ lệ dương tính thật (TPR) >85% trên dữ liệu thử nghiệm độc lập 30 ngày trước khi chuyển từ beta lên chính thức. [Sửa đổi BS-4] |
| **Tự động hóa tuân thủ** | **Mua ngoài: Vanta** | Gói Vanta Startup | Chi phí $4–6K/năm so với 3 tháng kỹ thuật (~$60K+). Đã có sẵn uy tín với các kiểm toán viên. Đạt SOC 2 Type 1 trong 60 ngày. |
| **Phát hiện Deepfake** | **Mua ngoài: API Hive Moderation** | Hive Moderation | Trả phí theo lượt sử dụng (<$0.01/lượt kiểm tra). Không yêu cầu dữ liệu huấn luyện. Đối tác chịu trách nhiệm cập nhật mô hình. |
| **Nền tảng ML** | **Mua ngoài: AWS SageMaker** | SageMaker | Quản lý huấn luyện, tự động mở rộng endpoint, giám sát độ lệch mô hình (drift monitoring). So với 6 tháng tự xây dựng MLOps tùy chỉnh. |
| **Cảnh báo lừa đảo bằng AI** | **Đối tác: M365 Defender** | Microsoft Security Graph API | Tính năng phát hiện cấp doanh nghiệp đã có sẵn trong M365. SMESec bổ sung thêm bối cảnh phong phú + kích hoạt kịch bản ứng phó sự cố (playbook). |

**Tổng chi phí sở hữu (TCO) năm thứ 1 (năng lực 1K tenant):** Hạ tầng được thiết kế cho 1K tenant từ Sprint 1. Chi phí hạ tầng + mua ngoài ~$7,500/tháng ở đủ 1K tenant; biên lợi nhuận gộp ~99% ở 1K tenant ($800K/tháng MRR). Ở giai đoạn tăng trưởng ban đầu (50 tenant, cùng hạ tầng): $40K MRR so với ~$5,500/tháng hạ tầng (chi phí biến đổi thấp hơn ở quy mô nhỏ) = ~86% biên lợi nhuận gộp. Biên lợi nhuận tăng dần khi số tenant tăng về mức 1K.

### 2.2 Mô hình Đa khách thuê (Multi-Tenancy Model)

**Quyết định:** Sử dụng cụm PostgreSQL dùng chung với cơ chế Bảo mật cấp dòng (Row-Level Security - RLS) được thực thi ở lớp cơ sở dữ liệu.

**Các giải pháp thay thế đã bị loại bỏ:**
- *Silo (Mỗi khách thuê một DB riêng biệt):* Chi phí hạ tầng khoảng ~$100–200/tháng cho mỗi khách thuê — không khả thi với mức giá dành cho SME.
- *Lược đồ dùng chung, cô lập ở cấp ứng dụng:* Lỗi ứng dụng có thể dẫn đến rò rỉ dữ liệu giữa các khách thuê. Không đảm bảo độ tin cậy.

**Triển khai:**

Mỗi bảng dữ liệu nghiệp vụ (domain table) bắt buộc phải có hai cột sau đây, không có ngoại lệ:

```sql
tenant_id      UUID        NOT NULL  -- bắt buộc thực thi RLS
data_residency VARCHAR(10) NOT NULL  -- 'US' | 'EU' | 'APAC'

-- PostgreSQL RLS policy (áp dụng ngay cả với chủ sở hữu bảng):
CREATE POLICY tenant_isolation ON assets
  FOR ALL TO app_role
  USING (tenant_id = current_setting('app.tenant_id', TRUE)::UUID);

ALTER TABLE assets ENABLE ROW LEVEL SECURITY;
ALTER TABLE assets FORCE ROW LEVEL SECURITY;  -- chặn cả quyền siêu người dùng (superuser)
```

Lớp trung gian (middleware) của Go API sẽ chèn `tenant_id` vào mỗi phiên PostgreSQL thông qua lệnh `SET LOCAL app.tenant_id` trước khi bất kỳ truy vấn nào được thực thi. Các thông tin khai báo của JWT (JWT claims) được xác thực, định dạng UUID được kiểm tra (ngăn chặn tấn công injection), sau đó biến phiên sẽ được thiết lập. Một bài kiểm tra CI bắt buộc sẽ tạo ra hai tenant, chèn dữ liệu cho Tenant A, truy vấn với tư cách là Tenant B — kết quả phải trả về 0 dòng. Quy trình hợp nhất mã nguồn (merge) sẽ bị chặn nếu bài kiểm tra này thất bại.

**Định tuyến lưu trữ dữ liệu (Data residency routing):** Khách thuê khu vực EU sẽ được định tuyến đến cụm ECS + RDS ở vùng `eu-west-1`. Dữ liệu của EU không bao giờ được ghi vào vùng `us-east-1`. Điều này được thực thi tại: DB schema, chính sách S3 bucket, vùng khóa KMS và vùng Secrets Manager. Đây là một bất biến bắt buộc ngay từ Sprint 1 — việc bổ sung sau này sẽ yêu cầu chuyển đổi toàn bộ lược đồ dữ liệu (schema migration).

### 2.3 Chiến lược Phát hiện Mối đe dọa từ AI

**Kiến trúc:** Phân tách làm 2 luồng độc lập — luồng tất định (Luồng 1) và luồng ML/AI (Luồng 2) — chỉ chia sẻ một giao ước cấu trúc dữ liệu sự kiện `ThreatDetectionEvent` (schema contract) và EventBridge event bus.

**Tại sao lại chia làm 2 luồng (mà không gộp chung thành một dịch vụ):**
- Luồng 1 có cam kết dịch vụ (SLA) mang tính tất định (hoàn thành thôi việc <5 phút). Hoạt động suy luận ML của Luồng 2 có độ trễ không tất định dao động từ 100 mili giây đến 2 giây.
- Các lỗi phát sinh ở Luồng 2 (lệch mô hình - model drift, khởi động nguội SageMaker) tuyệt đối không được làm ảnh hưởng đến tính sẵn sàng của Luồng 1.
- Các sự kiện từ Luồng 2 có thể kích hoạt các kịch bản (playbook) của Luồng 1 — nhưng Luồng 1 không bao giờ phải chờ đợi Luồng 2.

**Luồng 1 — Tất định (Bàn giao vào Tháng 3, độ chính xác 100%):**

| Mối đe dọa | Hướng phát hiện | Phương thức xử lý |
|---|---|---|
| Phát hiện Shadow IT | Rà soát danh mục ứng dụng OAuth — chấm điểm rủi ro theo phạm vi (dựa trên ma trận quy tắc) | Cảnh báo + áp dụng danh sách cho phép (allow-list) |
| Quyền truy cập mồ côi (tài khoản cũ không dùng) | Mô hình máy trạng thái tất định: nhân viên bị vô hiệu hóa trong HR ≠ vẫn hoạt động trong SaaS | Quy trình tự động hóa thôi việc (offboarding) qua Step Functions |
| Cấp thừa quyền | Công cụ so sánh RBAC: quyền thực tế so với chính sách vai trò được định nghĩa | Đưa ra khuyến nghị đặc quyền tối thiểu |
| Vi phạm tuân thủ | Danh sách kiểm tra đối chiếu kiểm soát theo ISO 27001 / SOC 2 / GDPR | Phát hiện các lỗ hổng tuân thủ |

**Luồng 2 — Phát hiện bằng ML/AI (Bàn giao vào Tháng 9.75/v1 đầy đủ, được kiểm soát bởi các cổng kiểm định độ chính xác):**

| Tính năng | Công nghệ | Cổng kiểm định độ chính xác (Accuracy Gate) | Điều kiện phát hành |
|---|---|---|---|
| **Chấm điểm rủi ro Shadow AI** | SageMaker endpoint (vectơ đặc trưng: các phạm vi OAuth, văn bản DPA của nhà cung cấp, tuổi đời ứng dụng) | Phân loại công cụ AI đạt độ chính xác >95% | Đánh giá tại Sprint 9 |
| **LLM DLP (browser ext)** | Presidio WASM (Lớp 1: regex) + BERT-tiny ONNX (Lớp 2: ngữ nghĩa) | Phát hiện thông tin cá nhân (PII) quan trọng >99%, tỷ lệ dương tính giả (FP) <5% | Triển khai môi trường staging tại Sprint 8 |
| **Phòng chống Deepfake** | API Hive Moderation + xác thực kênh phụ ngoài luồng (Step Functions) | Phát hiện deepfake giọng nói >80% (do đội ML của SMESec tự đánh giá độc lập trên bộ dữ liệu kiểm thử có gán nhãn — không phải do nhà cung cấp cam kết); kết hợp quy trình xác minh kênh phụ (OOV) ≈ 99% tỷ lệ ngăn chặn gian lận | Đánh giá tại Sprint 10 |
| **Tấn công chèn lệnh (Prompt injection)** | **API Lakera Guard (v1, Sprint 8)** → Tinh chỉnh BERT (v2, gói Enterprise, Sprint 23–24, chỉ áp dụng nếu chi phí Lakera quá đắt đỏ + có đủ dữ liệu gán nhãn) | TPR >85%, FPR <2% — do đội ML của SMESec tự đánh giá độc lập trên bộ dữ liệu holdout 30 ngày đặc thù của SMESec (cho cả Lakera v1 lẫn BERT v2; SLA của Lakera chỉ cam kết uptime API, không phải độ chính xác phát hiện) | Sprint 8 / Sprint 24 |

**Chính sách cổng kiểm định độ chính xác:** Không có tính năng nào thuộc Luồng 2 được phát hành dưới dạng GA (bản chính thức) cho đến khi đáp ứng cổng kiểm định độ chính xác. **Tất cả cổng kiểm định đều được đội ML của SMESec tự đánh giá độc lập trên dữ liệu holdout đặc thù của SMESec — SLA uptime API của nhà cung cấp không được chấp nhận thay thế cổng kiểm định độ chính xác.** Nếu không đạt cổng kiểm định → tính năng sẽ giữ nguyên ở trạng thái `beta` (chỉ cho phép đăng ký dùng thử, không áp dụng SLA). Hoạt động phát hành Luồng 1 không bao giờ bị trì hoãn bởi Luồng 2.

### 2.4 Cam kết Quyền riêng tư Dữ liệu

Bốn cam kết theo hợp đồng và được thực thi chặt chẽ bằng kiến trúc hệ thống:

| Cam kết | Triển khai thực tế | Xác minh |
|---|---|---|
| **Không huấn luyện mô hình trên dữ liệu của khách hàng** | SageMaker chỉ huấn luyện trên các tập dữ liệu công khai + dữ liệu tổng hợp. Dữ liệu của khách hàng không bao giờ được dùng để huấn luyện mô hình. | Công bố tài liệu đặc tính mô hình (model card); đánh giá kiến trúc định kỳ. |
| **Suy luận cục bộ trên tiện ích mở rộng trình duyệt** | Presidio WASM chạy hoàn toàn trên trình duyệt của người dùng. Nội dung nhập vào các công cụ AI không bao giờ rời khỏi thiết bị của người dùng. Chỉ có các siêu dữ liệu ẩn danh hóa (loại, mức độ nghiêm trọng, mốc thời gian) được gửi về máy chủ. | Mã nguồn tiện ích mở rộng dạng mở; kiểm toán lưu lượng mạng. |
| **Nhật ký kiểm toán bất biến (Có thể xóa theo chuẩn GDPR)** | Bản mã khóa đối tượng S3 Object Lock WORM (lưu giữ 7 năm). **⚠️ R-C4:** Mã hóa phong bì (envelope encryption) bằng khóa KMS riêng biệt cho mỗi khách thuê — khi hủy khóa = không thể truy cập vĩnh viễn = tương đương với "xóa bỏ thực tế" theo chuẩn GDPR (Khuyến nghị EDPB 01/2020). Bản mã vẫn nằm trên kho lưu trữ nhưng vĩnh viễn không thể đọc được nếu không có khóa giải mã. | Cấu hình AWS Object Lock; nhật ký xóa khóa KMS; chứng nhận xóa dữ liệu; minh chứng phục vụ kiểm toán SOC 2. |
| **Cách ly khu vực lưu trữ dữ liệu** | Cột `data_residency` là bắt buộc từ Sprint 1. Dữ liệu của các khách thuê thuộc EU chỉ nằm ở vùng `eu-west-1` — được thực thi tại các lớp cơ sở dữ liệu, S3, KMS và Secrets Manager. | Bài kiểm tra CI cô lập khách thuê; kiểm tra xâm nhập (penetration test). |

**Mã hóa:** RDS AES-256 (KMS CMK), S3 SSE-KMS, TLS 1.3 (bên ngoài), tất cả các bí mật được lưu trong Secrets Manager (tự động xoay vòng, không lưu văn bản thuần túy trong biến môi trường). Quyền truy cập các bí mật tuân thủ đặc quyền tối thiểu IAM: mỗi dịch vụ chỉ có thể truy cập không gian tên bí mật của riêng mình.

**Sự phù hợp với GDPR:** Điều 17 (Xóa dữ liệu) thông qua điểm cuối `/api/v1/gdpr/erasure` — PII được ẩn danh hóa trong vòng 30 ngày + khóa KMS CMK được lên lịch xóa (cửa sổ chờ của AWS là 7 ngày); bản mã vĩnh viễn không thể truy cập sau khi xóa khóa; cấp chứng nhận xóa dữ liệu (Khuyến nghị EDPB 01/2020). Điều 20 (Quyền chuyển đổi dữ liệu) thông qua điểm cuối xuất dữ liệu dạng JSON. Điều 25 (Bảo mật ngay từ khâu thiết kế - Privacy by design) thông qua cấu trúc dữ liệu `data_residency` ngay từ ngày đầu tiên và kiến trúc suy luận cục bộ trên máy khách. [R-C4]

**Lộ trình tuân thủ:** Vanta được thiết lập tại Tháng 4.5 (Tuần 18), thu thập minh chứng chính thức từ Tuần 20. SOC 2 Type 1 tại v1 đầy đủ (Tháng 9.75/Tuần 39, cửa sổ minh chứng Tuần 20→39 = 19 tuần). SOC 2 Type 2 + ISO 27001 tại v2 (Tháng 19.5/Tuần 78, cửa sổ minh chứng Tuần 39→68 = 29 tuần — vượt yêu cầu tối thiểu 24 tuần 5 tuần).

---

## 3. Đội ngũ & Kế hoạch Bàn giao

### 3.1 Nhân sự — Tăng trưởng theo từng Cột mốc

| Giai đoạn | Tháng | Nhân sự tương đương toàn thời gian (FTE) | Cơ cấu đội ngũ | Cột mốc bàn giao |
|---|---|---|---|---|
| **Giai đoạn 1** | 1–4.5 | **5** | Trưởng nhóm kỹ thuật (Tech Lead) · BE#1 · BE#2 · FE#1 · DevSecOps (Hợp đồng) + PM (0.5) + **Cố vấn BD (Hợp đồng, Tuần 1, 3 ngày/tuần) [R-C5]** | **v1 Thương mại** (Tuần 18/Tháng 4.5) — Chỉ Luồng 1 · Gói Starter Stripe live |
| **Giai đoạn 2** | 5–9.75 | **7 → 9** | +Kỹ sư Flutter (Tháng 5) · **+Kỹ sư ML #1 (Tháng 5 — R&D Luồng 2 bắt đầu)** · +BE#3 Python (Tháng 8) · +FE#2 Browser Extension (Tháng 9) | **v1 đầy đủ** (Tuần 39/Tháng 9.75) — Thêm các tính năng AI |
| **Giai đoạn 3** | 10–14.25 | **11** | +Kỹ sư hỗ trợ khách hàng (Tháng 10.5) · +Kỹ sư ML #2 (Tháng 12, tùy chọn) · Chuyển DevSecOps sang FTE | **v1.5** (Tuần 57/Tháng 14.25) |
| **Giai đoạn 4** | 15–19.5 | **11.5** | +Cố vấn tuân thủ (Hợp đồng Tháng 15–19.5) | **v2** (Tuần 78/Tháng 19.5) |

**Phân chia công việc đội ngũ từ Giai đoạn 3+ (2 luồng):** Luồng A (65%): phát triển tính năng mới + chuẩn bị cho SOC 2 Type 2 + cải tiến độ chính xác của AI. Luồng B (35%): tiếp nhận phản hồi từ khách hàng thử nghiệm, sửa lỗi, tinh chỉnh trải nghiệm người dùng (UX). Cả hai luồng hội tụ tại mỗi cột mốc quan trọng.

### 3.2 Kế hoạch bàn giao 19.5 tháng — Điều chỉnh 1.5x (39 Sprint, mỗi Sprint 2 tuần)

#### Giai đoạn 1: Xây dựng Nền tảng → v1 Thương mại (Sprint 1–9, Tháng 1–4.5) — 5 FTE, Chỉ Luồng 1

> **Luồng 2 (AI/ML) được hoãn sang Giai đoạn 2.** Kỹ sư ML #1 gia nhập Tháng 5. Điều này giúp Giai đoạn 1 chỉ cần 5 FTE, sử dụng sprint 60%, và cho ra sản phẩm thương mại có thanh toán Stripe tại Tuần 18 — không có rủi ro sprint 88% hay yêu cầu tuyển ML trước Ngày 1.

| Sprint | Luồng 1 (Track 1) | Luồng 2 (Track 2) | Cổng kiểm tra (Gate) |
|---|---|---|---|
| **S1** (Tuần 1–2) | Hạ tầng AWS: VPC + ECS + RDS + **RDS Proxy** (bắt buộc — 1K×10×4=40K kết nối >> 3.200 giới hạn) + Redis **cache.r6g.large** + S3 Object Lock · Keycloak SSO (4 tác vụ ECS, JWKS cache) · Schema đa thuê: `tenant_id` + `data_residency` + **`gcp_project_id`** + **`shard_id`** + **`renewal_bucket`** trên `tenant_config` từ ngày đầu · CI/CD · **50 dự án GCP được provisioned** (R-C2: quota vượt tại ~70 tenant) · **Bounded sync worker pool** (200 worker + job queue) · **Schema Secrets Manager theo lô** (1 JSON/tenant) · **`subscription_registry`** + EventBridge Scheduler gia hạn webhook M365 [R-C3] | Thiết kế schema `ThreatDetectionEvent` v0.1 (phối hợp cả hai luồng) · Đánh giá tài liệu nghiên cứu (OWASP LLM Top 10, PromptBench) · Lập kế hoạch thu thập bộ dữ liệu · Cài đặt môi trường SageMaker · Danh mục công cụ AI v0.1 (100+ công cụ) | CI cô lập khách thuê đạt xanh · Luồng 2: schema v0.1 được hai luồng review |
| **S2** (Tuần 3–4) | Đồng bộ Google Workspace — người dùng, ứng dụng OAuth, phát hiện shadow IT. Khung dashboard. | Đánh giá mô hình nền tảng (baseline): BERT-tiny + regex so với bộ dữ liệu gán nhãn (PromptBench, Presidio test suite) · Thiết kế tiêu chí chấm điểm rủi ro Shadow AI | Demo giá trị đầu tiên <30 phút sau OAuth · Luồng 2: benchmark độ chính xác nền tảng được ghi lại |
| **S3** (Tuần 5–6) | Đồng bộ M365 + delta link, dashboard hợp nhất (Google + M365), chỉ số rủi ro trên mỗi người dùng/ứng dụng | Nguyên mẫu phát hiện chèn lệnh v0.1 (fine-tune BERT-tiny — đo TPR/FPR baseline; khoảng cách so với cổng sản xuất TPR >85%/FPR <2% được xác định rõ) · Thiết lập pipeline biên dịch Presidio WASM · **Tài khoản API Lakera Guard + đo lường chi phí/request — được chỉ định là giải pháp triển khai chính cho v1** | Toàn bộ tài sản từ cả hai nhà cung cấp hiển thị · Luồng 2: TPR/FPR baseline được ghi lại |
| **S4** (Tuần 7–8) | Công cụ phân loại tài sản, chấm điểm rủi ro OAuth, cảnh báo shadow IT (<15 phút), khung Flutter mobile | Scaffold browser extension (Chrome MV3): Tier 1 regex DLP hoạt động trên Chrome dev · Mô hình chấm điểm rủi ro Shadow AI v0.1 (SageMaker training job) | Pipeline cảnh báo shadow IT hoạt động · Luồng 2: DLP chặn email/thẻ tín dụng trong dev Chrome |
| **S5** (Tuần 9–10) | Phát hiện Slack + AWS IAM, RBAC + khuynế nghị đặc quyền tối thiểu, biểu đồ định danh tổng hợp | **Cổng kiểm định độ chính xác 1 & 2 (Tuần 10):** Cổng 1 — Chèn lệnh: Lakera Guard TPR >85%, FPR <2% trên holdout 30 ngày (do đội ML SMESec tự đánh giá độc lập) · Cổng 2 — LLM DLP: PII quan trọng >99%, FP <5% · Shadow AI classification >95% trên top-100 công cụ · Hive API live | 4 nhà cung cấp hợp nhất · Luồng 2: Báo cáo cổng kiểm định 1 & 2 |
| **S6** (Tuần 11–12) | **🏁 MVP**: Tự động hóa thôi việc <5 phút (Step Functions) + **grace period 30 phút cấu hình được (khẩn cấp=0) + rollback 24h + idempotency key [R-C1]**, 2 kịch bản ứng phó sự cố (wizard UI), nhật ký kiểm toán bất biến, mobile app beta | DLP extension v0.3 kiểm thử trên ChatGPT/Gemini thực (môi trường staging) · Schema `ThreatDetectionEvent` v1 draft · Tổng kết R&D Giai đoạn 1 của Luồng 2 | Thôi việc <5 phút trên CI đạt · grace period/rollback tests đạt · Luồng 2: DLP xác nhận end-to-end trong staging |

**v1 Thương mại (Tuần 18/Tháng 4.5) = Chỉ Luồng 1. Thanh toán Starter ($399/tháng) qua Stripe đã hoạt động. "Bạn có thể thu hồi toàn bộ quyền truy cập của nhân viên nghỉ việc trong 5 phút không?" — và tính tiền dịch vụ.**

> **Bổ sung Sprint S7–S9:** S7 = Slack cơ bản + onboard 3+ khách hàng thử nghiệm. S8 = Tích hợp Stripe + hoàn thiện nhật ký kiểm toán. S9 = Tăng cường bảo mật + cổng kiểm tra phát hành thương mại.**

#### Giai đoạn 2: v1 Thương mại → v1 Đầy đủ (Sprint 10–20, Tháng 5–9.75) — 7→9 FTE, R&D Luồng 2 Bắt đầu

| Sprint | Luồng 1 (Track 1) | Luồng 2 (Track 2) | Cổng kiểm tra (Gate) |
|---|---|---|---|
| **S10** | Truy cập JIT + tự động thu hồi, đánh giá định kỳ · Vanta thiết lập Tuần 20 | Kỹ sư ML #1 + Flutter gia nhập Tháng 5 — R&D Shadow AI trên dữ liệu OAuth thực | Vanta thu thập minh chứng từ Tuần 20 |
| **S11** | Công cụ kịch bản (Step Functions), phát triển xong 3 kịch bản | LLM DLP browser extension v0.1 (Presidio + Tier 2 BERT-tiny inference cục bộ) | Tiện ích mở rộng phát hiện PII trong trường nhập liệu |
| **S12** | Hoàn thành 5 kịch bản, thông báo đẩy trên di động | Quản trị Shadow AI v1: phân loại công cụ AI + chấm điểm rủi ro + cam kết | Shadow AI >95% độ chính xác |
| **S13** | Dashboard tuân thủ ISO 27001 + SOC 2, tích hợp Vanta | POC phòng chống deepfake (Hive API), schema `ThreatDetectionEvent` v1 **đóng băng** | Schema khóa — không thay đổi tương thích |
| **S14** | Xuất báo cáo tuân thủ PDF, tự động hóa GDPR · **Pentest bắt đầu Tuần 27** | Tích hợp Luồng 1–Luồng 2: Sự kiện AI → EventBridge → Step Functions | LOI pentest ký Tuần 21 (hạn chót bắt buộc) |
| **S15** | **Khắc phục Pentest #1** — Lỗi Nghiêm trọng/Cao | Kiểm thử tích hợp T1–T2 | Tất cả lỗi Nghiêm trọng/Cao đã xử lý |
| **S16** | **Khắc phục Pentest #2** — kiểm tra lại + Vanta dry run | Kiểm thử tích hợp đầu–cuối tự động | 0 lỗi Nghiêm trọng/Cao còn tồn đọng |
| **S17** | Bản đồ phụ thuộc SaaS · Vanta >90% tỷ lệ qua | Chế độ thực thi chính sách Shadow AI | Vanta dry run ✅ |
| **S18** | **Nộp Chrome Extension lên Web Store** | Phiên bản đầy đủ extension nộp lên | Bộ đệm review 2 tuần |
| **S19** | Bộ đệm review Store · tăng cường bảo mật | Hoàn thiện Luồng 2 | Extension được duyệt hoặc kế hoạch dự phòng |
| **S20** | **🏁 v1 ĐẦY ĐỦ (Tuần 39)**: Ra mắt production, 5+ khách hàng trả phí, ký SOC 2 Type 1 | Shadow AI + LLM DLP trong v1 | Không phát triển tính năng mới — chỉ tăng cường |

**Cổng bàn giao v1 (Tuần 39/Tháng 9.75):** Hoàn thành 7 yêu cầu chính. 5+ khách hàng trả phí trên production. Lên lịch kiểm toán SOC 2 Type 1.

#### Giai đoạn 3 & 4: v1 → v1.5 → v2 (Sprint 21–39, Tháng 10–19.5)

| Cột mốc | Tháng | Các hạng mục bổ sung chính |
|---|---|---|
| **v1.5** (Tuần 57) | 14.25 | Tích hợp sâu AWS (CloudTrail), phòng chống deepfake v2 + lọc lừa đảo AI (M365 Defender), tiện ích mở rộng trên Chrome Web Store, gói Growth/Business áp dụng, Stripe live, 10+ khách hàng trả phí |
| **v2** (Tuần 78) | 19.5 | Chứng nhận SOC 2 Type 2 ✅ · ISO 27001 ✅ · BERT prompt injection (TPR >85%, FPR <2%) · Gói Enterprise (SIEM, RBAC tùy chỉnh, CSM riêng) · Tất cả tính năng Luồng 2 chính thức hoàn tất thử nghiệm |

### 3.3 Bản đồ Đáp ứng Yêu cầu Cốt lõi

| Yêu cầu | Cột mốc | Sprint | Ghi chú |
|---|---|---|---|
| **Định danh & phân loại tài sản** | v1 đầy đủ (Tháng 9.75) | Tính năng cốt lõi tại S2–S4 | Google+M365 hoàn thành ở v1 thương mại (Tháng 4.5). Slack+AWS tại S5. Phát hiện Shadow AI (Luồng 2) bổ sung trong Giai đoạn 2. |
| **Bề mặt đe dọa chuyên biệt từ AI** | v1 đầy đủ (Tháng 9.75) | Giai đoạn 2 S10–S19 | Quản trị Shadow AI, LLM DLP, deepfake + chèn lệnh đều trong Giai đoạn 2. |
| **Quản trị truy cập** | v1 thương mại (Tháng 4.5) → đầy đủ (Tháng 9.75) | S5–S10 | RBAC + thôi việc tại v1 thương mại. JIT + đánh giá định kỳ trong Giai đoạn 2 (S10). |
| **Đánh giá hiện trạng tuân thủ liên tục** | Sẵn sàng xuất báo cáo tại v1 (Tháng 9.75) | Giai đoạn 2 S13–S20 | SOC 2 Type 1 + ISO 27001 sẵn sàng xuất báo cáo tại v1. Chứng nhận chính thức tại v2 (Tháng 19.5). |
| **Kịch bản ứng phó sự cố** | v1 thương mại (Tháng 4.5) → đầy đủ (Tháng 9.75) | S6 (3 kịch bản); Giai đoạn 2 (5 kịch bản) | 3 kịch bản tại v1 thương mại. 5 kịch bản trong Giai đoạn 2 (S11–S12). |
| **Mô hình giá (phân cấp gói dịch vụ)** | Starter live tại v1 thương mại (Tháng 4.5); Growth/Business tại v1.5 (Tháng 14.25) | S8–S9 Stripe Starter; Giai đoạn 3 các gói còn lại | Starter ($399/tháng) live Tháng 4.5 · Growth ($799/tháng) + Business ($1,499/tháng) tại v1.5 (Tháng 14.25) · Enterprise (tùy chỉnh) tại v2. |
| **Tích hợp các công cụ SME phổ biến** | v1 đầy đủ (Tháng 9.75) | S2–S5 + Giai đoạn 2 | Google Workspace + M365 tại v1 thương mại. Slack tại S5. AWS IAM trong Giai đoạn 2. Hoãn QuickBooks sang v2. |

### 3.4 Giả định Rủi ro Nhất cần Kiểm chứng Trước tiên

> **Rủi ro số #1:** Quản trị viên CNTT của SME (không chuyên về kỹ thuật sâu) có thể hoàn thành cấu hình OAuth cho Google Workspace + M365 trong vòng chưa đầy 30 phút bằng giao diện wizard.

**Tại sao đây lại là giả định có rủi ro cao nhất:**
- Toàn bộ giá trị cốt lõi của MVP phụ thuộc vào việc "mang lại giá trị đầu tiên trong dưới 30 phút". Nếu quá trình tích hợp ban đầu mất tới 3 giờ (do sự phức tạp trong việc xin phê duyệt từ quản trị viên M365), chương trình chạy thử nghiệm sẽ thất bại ngay trước khi bắt đầu.
- Các đối thủ cạnh tranh mất từ 2–4 giờ cho các cấu hình tương đương. Nếu SMESec cũng mất nhiều thời gian như vậy, sản phẩm sẽ mất đi tính độc đáo cạnh tranh.
- Giả định này không thể kiểm chứng trong môi trường phòng thí nghiệm khép kín — bắt buộc phải kiểm tra thực tế với những người dùng không chuyên kỹ thuật trên các tenant Google Workspace thực tế.

**Kế hoạch kiểm chứng:** Cuối Sprint 2 (Tuần 4) — thực hiện bài kiểm tra khả năng sử dụng giới hạn thời gian với 1-2 người dùng không chuyên về kỹ thuật, không có sự hỗ trợ từ kỹ sư phần mềm.  
**Quyết định Đi tiếp/Dừng lại (Go/No-go):** Nếu thời gian hoàn thành >45 phút → bắt buộc phải thiết kế lại giao diện thuật sĩ (wizard) trước Sprint 3. Đình chỉ tất cả các công việc phát triển tính năng khác cho đến khi vấn đề này được giải quyết triệt để.

**Top 5 rủi ro hàng đầu (qua tất cả các giai đoạn):**

| # | Rủi ro | Giai đoạn | Xác suất | Tác động | Giải pháp giảm thiểu |
|---|---|---|---|---|---|
| 1 | Giao diện OAuth wizard mất >30 phút với quản trị viên CNTT không chuyên kỹ thuật | MVP | Cao | Nghiêm trọng | Thực hiện bài kiểm tra khả năng sử dụng ở Tuần 4. Viết tài liệu hướng dẫn cụ thể. Giải thích lý do yêu cầu quyền ở mức tối thiểu. |
| 2 | Kỹ sư ML #1 chưa được tuyển dụng trước Tháng 5 (bắt đầu Giai đoạn 2) | Giai đoạn 2 | Trung bình | Cao | Bắt đầu tuyển dụng từ Tháng 3. Phải hoạt động hiệu quả trước Tháng 5 khi R&D Luồng 2 bắt đầu. Giảm áp lực so với yêu cầu tuyển trước Ngày 1. |
| 3 | Việc tích hợp Luồng 1 - Luồng 2 tại S11 bị chậm trễ >1 sprint | Giai đoạn 2 | Cao | Cao | Trưởng nhóm kỹ thuật tập trung 100% tại S11. Đóng băng API contract tại S10. Phương án dự phòng: kích hoạt kịch bản thủ công cho v1. |
| 4 | LOI với nhà cung cấp Pentest chưa được ký trước Tuần 14 | Giai đoạn 2 | Thấp | Cao | PM chốt lịch trước Tuần 8. Chuẩn bị danh sách nhà cung cấp dự phòng. Thời hạn nghiêm ngặt: tuyệt đối không gia hạn. |
| 5 | Phát hiện lỗ hổng minh chứng SOC 2 Type 2 khi đánh giá Tháng 9 | Giai đoạn 3 | Thấp | Cao | Rà soát Vanta hàng tuần từ Tuần 20. PM chịu trách nhiệm. Chính sách không chấp nhận lỗ hổng từ Tuần 39 trở đi. |

---

## 4. Mô-đun Quản trị AI

### 4.1 Vấn đề Thực trạng

75% lao động tri thức trên toàn cầu sử dụng AI trong công việc — và 78% trong số họ đang tự mang các công cụ AI cá nhân vào làm việc mà không có sự chấp thuận của người sử dụng lao động (BYOAI - Bring Your Own AI), tỷ lệ này tăng lên tới 80% tại các doanh nghiệp vừa và nhỏ.<sup>[[1]](#src-1)</sup> 52% số người dùng ngần ngại thừa nhận việc sử dụng AI cho các nhiệm vụ quan trọng nhất của họ.<sup>[[1]](#src-1)</sup> 11% nội dung được dán vào ChatGPT có chứa dữ liệu bảo mật của công ty.<sup>[[2]](#src-2)</sup> Trung bình mỗi doanh nghiệp SME hiện có hơn 20 công cụ AI chưa được phê duyệt đang kết nối với tài khoản của công ty.<sup>[[3]](#src-3)</sup> Thiệt hại do lừa đảo email doanh nghiệp (BEC) từ việc giả mạo giọng nói của CEO bằng AI đạt 2,9 tỷ USD vào năm 2023, với thiệt hại trung bình của mỗi doanh nghiệp SME là $140K cho mỗi vụ việc.<sup>[[4]](#src-4)</sup>

**Chưa có nhà cung cấp nào đưa ra được một giải pháp thống nhất với mức giá phải chăng cho mô hình đe dọa "doanh nghiệp SME với tư cách là người tiêu dùng AI".** Hầu hết các nhà cung cấp bảo mật AI lớn (HiddenLayer, Wiz, Prompt Security) đều nhắm tới các công ty tự triển khai phát triển LLM — chứ không phải các công ty sử dụng chúng. Nudge Security phát hiện được shadow AI nhưng không thể ngăn chặn. Prompt Security có tính năng DLP trên trình duyệt nhưng chi phí lên tới $15–30K/năm và yêu cầu cấu hình phức tạp từ quản trị viên CNTT hoặc lập trình viên.

<a name="src-1"></a>**[1]** Microsoft & LinkedIn — [Báo cáo Chỉ số Xu hướng Công việc Thường niên 2024: AI trong Công việc đã Hiện hữu. Giờ là Lúc cho Phần Khó khăn](https://www.microsoft.com/en-us/worklab/work-trend-index/ai-at-work-is-here-now-comes-the-hard-part) (Tháng 5/2024, khảo sát trên 31.000 lao động tri thức tại 31 quốc gia)

<a name="src-2"></a>**[2]** Cyberhaven — [Báo cáo Rò rỉ Dữ liệu Cyberhaven 2024](https://www.cyberhaven.com/resources/data-exposure-report-2024) — phân tích hành vi luân chuyển dữ liệu của 1,4 triệu nhân viên thông qua hệ thống Cyberhaven DLP

<a name="src-3"></a>**[3]** Nudge Security — [Báo cáo Hiện trạng Bảo mật SaaS 2024](https://www.nudgesecurity.com/post/state-of-saas-security-2024-report) — dữ liệu đo lường phát hiện shadow AI trên nhóm khách hàng SME

<a name="src-4"></a>**[4]** Trung tâm Khiếu nại Tội phạm Internet của FBI — [Báo cáo Thường niên IC3 năm 2023](https://www.ic3.gov/AnnualReport/Reports/2023_IC3Report.pdf) — Phần Xâm nhập Email Doanh nghiệp, trang 14–15

### 4.2 Khung Quản trị: 3 Lớp Bảo vệ

```
LỚP 3 — BẢO VỆ (Ngăn chặn thời gian thực / Real-time prevention)
  Tiện ích mở rộng trình duyệt (Browser Ext): chặn trước khi gửi, ngăn dữ liệu nhạy cảm rò rỉ
  Phát hiện Deepfake: xác thực kênh phụ ngoài luồng trước khi thực hiện các yêu cầu đáng ngờ
  Phát hiện chèn lệnh (Prompt injection): dựa trên quy tắc (v1) + bộ phân loại BERT (v2, Enterprise)

LỚP 2 — QUẢN TRỊ (Thực thi chính sách / Policy enforcement)
  Chấm điểm rủi ro công cụ AI + công cụ chính sách: chặn/cho phép/cam kết dựa trên phạm vi OAuth + vị thế nhà cung cấp
  Quy trình cam kết tuân thủ của nhân viên: tự báo cáo việc dùng công cụ AI để bù đắp điểm mù của OAuth
  Quy trình phê duyệt của người quản lý đối với điểm rủi ro từ 61–85

LỚP 1 — PHÁT HIỆN (Thống kê tài sản thụ động / Passive inventory)
  Danh mục ứng dụng OAuth (Google + M365 + Slack, quét mỗi 15 phút)
  Phân loại công cụ AI: hơn 100 công cụ đã biết trong danh mục được cập nhật liên tục
  Thông tin đo lường sử dụng: chỉ lưu tên miền + mốc thời gian (hoàn toàn không lưu nội dung)

Lớp 1 cung cấp bối cảnh cho Lớp 2. Chính sách Lớp 2 cung cấp các ngưỡng rủi ro cho Lớp 3.
Lớp 3 (tiện ích mở rộng trình duyệt) hoạt động độc lập — tự động chuyển sang chế độ chặn (fails closed) nếu máy chủ gặp sự cố.
```

### 4.3 Mô-đun A — Cổng kiểm soát gửi dữ liệu AI (DLP trên Trình duyệt)

Quyết định kiến trúc cốt lõi nhằm bảo vệ quyền riêng tư: **nội dung prompt không bao giờ rời khỏi trình duyệt của người dùng**.

**Quy trình quét 3 lớp (tất cả đều chạy trực tiếp trong trình duyệt):**

| Lớp quét | Công nghệ sử dụng | Độ trễ | Nội dung được phát hiện | Độ chính xác |
|---|---|---|---|---|
| **Lớp 1 (Regex)** | Các mẫu quy tắc OWASP + quy tắc tùy chỉnh, cập nhật tự động từ máy chủ | <1 mili giây | Thẻ tín dụng (thuật toán Luhn), số định danh cá nhân (SSN/CCCD), email, số điện thoại, khóa API (regex của AWS/GitHub/Stripe), mã JWT, mã IBAN | Phát hiện PII quan trọng >99%, dương tính giả (FP) <1% |
| **Lớp 2 (WASM BERT-tiny)** | Microsoft Presidio được biên dịch sang ONNX WASM (dung lượng 17MB, tải chậm khi cần) | 50–80 mili giây | Dữ liệu bảo mật dạng ngữ nghĩa: "Dự báo doanh thu Quý 3", thông tin thảo luận mua bán sáp nhập (M&A), sở hữu trí tuệ của khách hàng | Phát hiện ngữ nghĩa >85%, dương tính giả (FP) <10% |
| **Lớp 3 (Bối cảnh, bất đồng bộ)** | FastAPI → API Lakera Guard (chạy phía máy chủ, không chặn tiến trình chính) | Bất đồng bộ | Các dạng tấn công chèn lệnh mới xuất hiện, chấm điểm rủi ro nhận biết bối cảnh (kết hợp vai trò người dùng + hệ số nhạy cảm của tài sản) | Độ chuẩn xác (precision) >90% |

**Các công cụ AI được hỗ trợ (phiên bản v1, có thể mở rộng thông qua cấu hình cập nhật từ máy chủ):** chatgpt.com · copilot.microsoft.com · gemini.google.com · claude.ai · perplexity.ai · github.com/copilot · notion.so

**Cơ chế tự động chặn khi có sự cố (Fail-closed guarantee):** Nếu tiện ích mở rộng không thể hoàn thành lượt quét ở Lớp 1 → quá trình gửi dữ liệu sẽ bị **chặn** lập tức và hiển thị thông báo rõ ràng cho người dùng. Tuyệt đối không cho phép dữ liệu đi qua một cách âm thầm.

**Giao diện Kiểm tra và Che giấu thông tin trước khi gửi (Pre-send Redaction Review UI):**
Khi phát hiện dữ liệu nhạy cảm, tiện ích mở rộng sẽ hiển thị một cửa sổ bật lên chặn thao tác (không thể tắt bằng nút Esc):
- Đánh dấu các thông tin bị phát hiện: `[API_KEY_1]` `[EMAIL_1]` `[PHONE_1]` kèm theo nhãn phân loại.
- Hành động mặc định: **"Gửi kèm các phần đã che giấu"** (các từ khóa giả định được thay thế giúp giữ nguyên cấu trúc ngữ pháp của câu lệnh).
- Ghi đè (Override): Yêu cầu người dùng nhập lý do rõ ràng (được ghi nhận vào bảng điều khiển của quản trị viên CNTT trong vòng 60 giây).
- Quản trị viên CNTT chỉ nhìn thấy: Loại dữ liệu cá nhân (PII) bị phát hiện, mức độ rủi ro, hành động đã xử lý — tuyệt đối không xem được nội dung văn bản thực tế.

**Thông tin được gửi về máy chủ SMESec (kiến trúc không tri thức - zero-knowledge):**

```
✅ Được phép gửi:  mức rủi ro (risk_tier), danh mục mẫu (pattern_category), tên miền đích (target_domain), mốc thời gian (timestamp), id khách thuê đã băm bảo mật (tenant_id - hashed)
❌ Không bao giờ gửi: nội dung prompt thực tế, các đoạn text bị đánh dấu, văn bản người dùng nhập
```

### 4.4 Mô-đun C — Quản trị Shadow AI

**Kiểm kê ứng dụng OAuth AI (C1):** Mỗi 15 phút, SMESec sẽ lấy danh sách các quyền ứng dụng OAuth được cấp từ Google Admin SDK + M365 Graph API + Slack Admin API. Mỗi ứng dụng sẽ được phân loại đối chiếu với danh mục hơn 100 công cụ AI đã được chọn lọc và chấm điểm rủi ro theo công thức trọng số sau:

```
risk_score = (oauth_scopes_sensitivity × 30%) +
             (vendor_DPA_available × 20%) +
             (data_residency_compliance × 15%) +
             (security_certifications × 15%) +
             (known_incidents × 10%) +
             (app_age_days × 5%) +
             (user_count_in_tenant × 5%)
```

| Phân nhóm Rủi ro | Ví dụ thực tế | Phương án xử lý |
|---|---|---|
| **NGUY HIỂM (CRITICAL)** | Ứng dụng không xác định yêu cầu quyền `gmail.modify` + `drive.readwrite`, không có văn bản DPA | Cảnh báo + Tự động thu hồi quyền truy cập (chạy thử nghiệm đánh giá → xác nhận 2 bước để thực hiện thực tế) |
| **CAO (HIGH)** | Jasper AI yêu cầu quyền đọc Gmail, ứng dụng hoạt động <6 tháng | Cảnh báo + Yêu cầu nhân viên thực hiện cam kết tuân thủ ("Tôi hiểu và tự chịu trách nhiệm") |
| **TRUNG BÌNH (MEDIUM)** | Sử dụng ChatGPT chỉ để xử lý văn bản, không có quyền truy cập tệp tin | Ghi nhận nhật ký hệ thống + Gửi báo cáo sử dụng AI hàng tháng cho quản trị viên CNTT |
| **THẤP/ĐÃ PHÊ DUYỆT TRƯỚC (LOW)** | Microsoft Copilot (tích hợp gốc M365), GitHub Copilot | Chỉ lưu trữ danh mục kiểm kê tài sản, không phát cảnh báo |

**Quy trình cam kết tuân thủ của nhân viên (C2):** Khảo sát tự báo cáo hàng quý sẽ đối chiếu việc sử dụng công cụ AI của nhân viên với danh mục ứng dụng OAuth thực tế. Việc này giúp loại bỏ "điểm mù OAuth" — những trường hợp nhân viên sử dụng ChatGPT trực tiếp qua trình duyệt cá nhân (không yêu cầu cấp quyền OAuth vào tài khoản công ty). Nếu nhân viên không phản hồi sau 5 ngày làm việc, hệ thống sẽ tự động ghi nhận là một lỗ hổng trong tuân thủ bảo mật.

### 4.5 Mô-đun D — Phòng chống Gian lận Deepfake

**Tình huống sử dụng thực tế:** "Có phải CEO của tôi thực sự đang yêu cầu tôi chuyển khoản gấp $50.000 không?"

**D1 — Phát hiện Deepfake Giọng nói (Ưu tiên các quốc gia ngoài EU trước, cần ý kiến pháp lý cho khu vực EU):**
Nhân viên tải lên một đoạn âm thanh dài ≤60 giây → API Hive Moderation sẽ phân tích (tệp âm thanh gốc KHÔNG được lưu trữ trên hệ thống và sẽ bị xóa trong vòng 60 giây). Kết quả phân tích hiển thị theo các khoảng xác suất chứ không trả về kết quả nhị phân (đúng/sai): *"Nhiều khả năng là thật"* / *"Không thể kết luận"* / *"Nhiều khả năng là giả lập bằng AI — hãy cẩn trọng"*. Việc triển khai tại khu vực EU yêu cầu phải có ý kiến pháp lý theo Điều 9 của GDPR (giọng nói được tính là dữ liệu sinh trắc học) — nhiệm vụ này được bắt đầu thực hiện từ Ngày 1, bàn giao trước tại các thị trường Mỹ, Anh, Úc.

**D2 — Quy trình xác thực thông qua kênh phụ ngoài luồng (Độc lập với D1, luôn luôn khả dụng):**

```
1. Nhân viên nhấn kích hoạt "Xác thực người này" trên ứng dụng di động (chỉ mất 3 lần chạm)
2. SMESec sẽ gửi thông báo qua HAI kênh độc lập đến người được cho là người gửi yêu cầu:
   - Email: "Bạn có liên hệ với [nhân viên] vào lúc [thời gian] không?" → Liên kết phản hồi nhanh [ĐÚNG / KHÔNG] (không cần tài khoản SMESec)
   - Tin nhắn SMS: Gửi mã xác thực một lần (OTP) tới số điện thoại đăng ký → nhân viên yêu cầu người gọi đọc lại mã này
3. Kết quả tổng hợp trong vòng 5 phút:
   Email phản hồi "KHÔNG PHẢI TÔI" + không cung cấp được mã xác thực → "⚠️ NHIỀU KHẢ NĂNG GIAN LẬN — Tuyệt đối KHÔNG tiếp tục thực hiện"
   Email phản hồi "ĐÚNG" + cung cấp đúng mã xác thực → "✅ ĐÃ XÁC THỰC — Danh tính chính xác"
   Thông tin không đồng nhất hoặc mơ hồ → "⚠️ KHÔNG THỂ KẾT LUẬN — Hãy chuyển giao báo cáo cho quản trị viên CNTT"
4. Nếu xác nhận có gian lận → kích hoạt nhanh bằng một nút chạm Kịch bản Ứng phó sự cố số 6 (Deepfake Fraud Response)
5. Toàn bộ tiến trình xác thực được lưu lại làm minh chứng tuân thủ bảo mật (nhật ký kiểm toán)
```

### 4.6 Mô-đun B — Phát hiện Tấn công Chèn lệnh (Prompt Injection)

**v1 (Sprint 8, API Lakera Guard):** Gọi REST API cho mỗi prompt được gửi lên trước khi chuyển tiếp tới trợ lý AI nội bộ. Chi phí khoảng ~$0.001/yêu cầu. Đã được kiểm chứng thực tế bởi Lakera đối với cả các mẫu tấn công đã biết và mới xuất hiện. Độ trễ <50 mili giây (p99). **[Sửa đổi BS-4 — thay thế giải pháp ban đầu dựa trên quy tắc regex chỉ bao phủ được ~75% các mẫu tấn công đã biết.]**

**v2 (Sprint 23–24, BERT, chỉ dành riêng cho gói Enterprise):** Chỉ được kích hoạt nếu (a) chi phí sử dụng Lakera Guard trở nên quá cao khi quy mô Enterprise tăng lên VÀ (b) tích lũy được ≥50K mẫu dữ liệu thực tế được gán nhãn chính xác. Tiến hành tinh chỉnh mô hình BERT dựa trên dữ liệu của các khách thuê gói Enterprise đã đồng ý tham gia. Điều kiện kiểm định (Gate): Đạt TPR >85% VÀ FPR <2% trên tập dữ liệu thử nghiệm thực tế độc lập trong 30 ngày. Nếu không đạt yêu cầu → tiếp tục duy trì sử dụng Lakera Guard cho bản GA, mô hình BERT giữ nguyên ở trạng thái thử nghiệm giới hạn.

**4 cấp độ xử lý dựa trên điểm số rủi ro tổng hợp (0–100):**

| Khoảng Điểm rủi ro | Hành động xử lý | Chế độ thông báo |
|---|---|---|
| 0–30 | Chỉ ghi nhật ký hệ thống | Báo cáo tổng hợp hàng tuần |
| 31–60 | Hiển thị thông báo cảnh báo nhẹ + yêu cầu giải trình | Báo cáo tổng hợp hàng ngày |
| 61–85 | Chặn hoàn toàn + Yêu cầu người quản lý phê duyệt (qua Slack/email, duyệt nhanh bằng một cú nhấp chuột) | Cảnh báo thời gian thực |
| 86–100 | Chặn hoàn toàn, không có quyền ghi đè hay bỏ qua | Cảnh báo mức P1 khẩn cấp, thông báo tức thì cho quản trị viên CNTT |

### 4.7 Mô-đun E — Phòng chống Tấn công giả mạo (Phishing) dựa trên AI

**Tích hợp M365 Defender (E1):** Lấy dữ liệu cảnh báo lừa đảo/mã độc từ API Microsoft Security Graph (`/v1.0/security/alerts_v2`) sau mỗi 5 phút. Làm phong phú thêm thông tin dựa trên bối cảnh của Luồng 1: Vai trò của người dùng bị ảnh hưởng, cấp độ truy cập dữ liệu, danh sách báo cáo trực tiếp. Hỗ trợ kích hoạt kịch bản ứng phó sự cố số 3 (Phishing Response) chỉ bằng một cú nhấp chuột ngay trên thông tin cảnh báo phong phú đó. Chỉ áp dụng cho các khách thuê có bản quyền sử dụng M365.

**Kiểm tra trạng thái cấu hình xác thực email (E2):** Thực hiện kiểm toán cấu hình DMARC/DKIM/SPF hàng tuần của Google Workspace thông qua Admin SDK. Đưa ra các hướng dẫn khắc phục cụ thể cho các cấu hình sai lệch (ví dụ: "Chính sách DMARC hiện là 'none' — email của bạn có nguy cơ bị giả mạo. Hãy cập nhật bản ghi DNS: p=quarantine").

### 4.8 Mô-đun F — Đảm bảo Quyền riêng tư và Tính minh bạch của Nhân viên

**Trang quản lý tính minh bạch (F1, bắt buộc theo tiêu chuẩn của EU):** Nhân viên luôn có thể truy cập phần này từ cửa sổ tiện ích mở rộng và trên ứng dụng di động (quản trị viên CNTT không có quyền vô hiệu hóa tính năng này). Hiển thị rõ ràng các nội dung: Những gì đang được giám sát (tên miền công cụ AI + ngày sử dụng), những gì KHÔNG bị giám sát (lịch sử duyệt web cá nhân, nội dung cụ thể trong prompt, màn hình làm việc/các phím bấm). Nhân viên có thể xem lịch sử 10 sự kiện bị gắn cờ gần nhất của mình.

**Tính năng tạm dừng giám sát (F2, tuân thủ mô hình chấp thuận của EU):** Nhân viên có quyền tạm dừng toàn bộ hoạt động giám sát trong khoảng thời gian 15/30/60 phút. Khi ở trạng thái tạm dừng: hoàn toàn không thực hiện quét dữ liệu, không đo lường thông tin, không truyền dữ liệu về máy chủ. Quản trị viên CNTT chỉ nhận được thông báo về khoảng thời gian tạm dừng — lý do tạm dừng tuyệt đối không được ghi nhận. Thời gian tạm dừng tối đa có thể được thiết lập cấu hình trên mỗi vai trò (ví dụ: vai trò CFO có thời gian tạm dừng tối đa bằng 0 phút).

### 4.9 Vị thế Cạnh tranh trên Thị trường

| Tính năng năng lực | Nudge Security (đối thủ SME gần nhất) | Prompt Security | SMESec |
|---|---|---|---|
| Phát hiện Shadow AI | ✅ | ❌ | ✅ |
| Thực thi chính sách (chặn/cho phép) | ❌ chỉ hiển thị nhắc nhở | ✅ | ✅ |
| DLP trên trình duyệt (kiến trúc không tri thức - zero-knowledge) | ❌ | ✅ | ✅ |
| Phòng chống gian lận deepfake | ❌ | ❌ | ✅ |
| Trải nghiệm người dùng đơn giản (không yêu cầu thiết lập CNTT phức tạp) | ✅ | ❌ đòi hỏi lập trình viên thiết lập | ✅ |
| Cung cấp bằng chứng tuân thủ (SOC 2) | ❌ | ❌ | ✅ |
| **Giá cho doanh nghiệp SME (~50 người dùng)** | **~$2,400/năm** | **$15–30K/năm** | **~$4,800/năm (gói trọn gói)** |

**SMESec là nền tảng duy nhất kết hợp cả 5 năng lực bảo mật trên với mức giá phù hợp cho doanh nghiệp SME và hoàn toàn không yêu cầu kỹ năng quản trị CNTT chuyên sâu để cài đặt.**

---

## Phụ lục: Lộ trình Đạt Chứng nhận Tuân thủ

```
Tháng 4.5  (Tuần 18): RA MẮT v1 Thương mại — Thiết lập tài khoản Vanta, bắt đầu thu thập minh chứng
Tháng 5    (Tuần 20): Vanta CHÍNH THỨC hoạt động — bắt đầu thiết lập kiểm soát theo chuẩn SOC 2
Tháng 5.25 (Tuần 21): Ký hợp đồng LOI kiểm tra xâm nhập (HẠN CHÓT BẮT BUỘC)
Tháng 7    (Tuần 27): Bắt đầu kiểm tra xâm nhập
Tháng 9.75 (Tuần 39): RA MẮT v1 ĐẦY ĐỦ
                       → Ký kết thỏa thuận kiểm toán SOC 2 Type 1
                       → Cửa sổ minh chứng Tuần 20→39 = ~19 tuần (đủ cho Type 1)
Tháng 10.5 (Tuần 42): Bắt đầu phân tích sai lệch (gap analysis) ISO 27001
Tháng 12   (Tuần 48): Kiểm toán ISO 27001 Giai đoạn 1 (đánh giá tài liệu quy trình)
Tháng 14.25(Tuần 57): RA MẮT v1.5 → Minh chứng SOC 2 Type 2 tích lũy từ Tuần 39
Tháng 15   (Tuần 60): Kiểm toán ISO 27001 Giai đoạn 2 (đánh giá triển khai thực tế)
Tháng 17   (Tuần 68): Bắt đầu đánh giá thực địa SOC 2 Type 2
                       → Cửa sổ minh chứng Tuần 39→68 = 29 tuần ✅ (vượt yêu cầu tối thiểu 24 tuần)
Tháng 19.5 (Tuần 78): RA MẮT v2 → Đạt cả hai chứng nhận SOC 2 Type 2 ✅ + ISO 27001 ✅
```

**Lưu ý:** Chứng nhận SOC 2 Type 2 yêu cầu khoảng thời gian theo dõi liên tục tối thiểu là 6 tháng (24 tuần). Cửa sổ minh chứng Tuần 39→68 = **29 tuần** — cung cấp 5 tuần dự phòng, an toàn hơn đáng kể so với kế hoạch gốc (cửa sổ 26 tuần không có dự phòng).

---

## 5. Lộ trình Sau v2 & Các Nghĩa vụ Duy trì Thường xuyên (Tháng 13 trở đi)

> Phiên bản v2 (Tuần 52) là cột mốc khẳng định sản phẩm đã "thương mại hóa khả thi" — chứ không phải đã hoàn thành toàn bộ. Các hạng mục dưới đây là những nghĩa vụ bắt buộc, đi kèm mức độ rủi ro cao nếu không được lên kế hoạch chu đáo trước khi kết thúc Năm 1.

### 5.1 Duy trì Tuân thủ & Các Chứng nhận Định kỳ (Compliance & Certifications)

| Nghĩa vụ tuân thủ | Tần suất | Điều kiện kích hoạt / Thời hạn cuối |
|---|---|---|
| **Tái kiểm toán SOC 2 Type 2 (Năm thứ 2)** | Hàng năm | Tuần 104 (Tháng 24) — dữ liệu minh chứng trong khoảng Tuần 52→Tuần 104 phải đảm bảo hoàn toàn sạch lỗi |
| **Đánh giá giám sát ISO 27001 lần #1** | 12 tháng sau khi được cấp chứng nhận (Tuần 52 + 12 tháng) | Hạng mục bắt buộc để duy trì hiệu lực của chứng nhận — không phải tùy chọn |
| **Tái đánh giá chứng nhận ISO 27001** | Mỗi 3 năm kể từ ngày cấp chứng nhận đầu tiên | Lên kế hoạch chuẩn bị năng lực kỹ thuật và chi phí thuê chuyên gia tư vấn trước |
| **Đánh giá bảo mật xâm nhập (Pentest định kỳ)** | Mỗi 6 tháng | Lượt pentest tiếp theo diễn ra vào Tháng 18. Thay đổi đối tác pentest hàng quan năm. |
| **Đánh giá định kỳ hàng năm thỏa thuận DPA GDPR** | Hàng năm | Thỏa thuận xử lý dữ liệu (DPA) ký kết với khách hàng phải phản ánh chính xác các thay đổi trong kiến trúc hệ thống. Phương pháp mã hóa phong bì qua KMS phải được ghi nhận rõ trong DPA. |
| **Duy trì tính liên tục của minh chứng Vanta** | Liên tục | Đảm bảo hoàn toàn không có lỗ hổng minh chứng kể từ Tuần 26 trở đi. Bất kỳ lỗ hổng nào xuất hiện sẽ reset lại thời gian theo dõi của SOC 2 Type 2. Quản lý Dự án (PM) chịu trách nhiệm kiểm tra định kỳ Vanta hàng tuần. |

### 5.2 Hạ tầng Kỹ thuật

| Hạng mục hạ tầng | Độ ưu tiên | Ghi chú chi tiết |
|---|---|---|
| **Vá lỗi CVE của Keycloak hàng quý** | 🔴 Rất cấp bách (Duy trì liên tục) | Keycloak phát hành các bản vá lỗi bảo mật hàng tháng. Quy trình nâng cấp cuốn chiếu không gây gián đoạn hệ thống (zero-downtime rolling upgrade) phải được viết thành tài liệu và thực hiện diễn tập định kỳ trước khi ra mắt bản v1. Nếu chi phí vận hành quá lớn → xem xét chuyển đổi sang WorkOS/Auth0 (điểm ra quyết định: cuộc họp tổng kết bản v1.5). |
| **Xoay vòng khóa KMS & Quản lý chứng nhận xóa dữ liệu** | 🔴 Cấp bách | Tự động hóa việc xoay vòng khóa KMS CMK riêng của mỗi khách thuê hàng năm. Các chứng nhận xóa dữ liệu theo tiêu chuẩn GDPR phải được lưu trữ và có thể kết xuất nhanh khi có yêu cầu. Hoàn thiện quy trình lưu nhật ký xóa dữ liệu trước khi ra mắt bản v1. |
| **Tối ưu kiến trúc giới hạn tần suất gọi API Google** | 🔴 Rất cấp bách — Bắt buộc Sprint 1 | v1 nhắm mục tiêu 1K tenant. Quota Google API vượt tải tại ~70 tenant (1.500 req/100s mỗi GCP project). 1K tenant cần 50 GCP project (1K / 20 tenant/project). GCP project pool được triển khai trong hạ tầng Sprint 1; cột `gcp_project_id` trong schema `tenant_config` từ ngày đầu. SyncScheduler gán project cho tenant khi onboarding. Cảnh báo khi đạt 80% quota mỗi project. |
| **Cơ chế gom hồ kết nối RDS (connection pooling)** | 🔴 Rất cấp bách — Bắt buộc Sprint 1 | v1 nhắm mục tiêu 1K tenant. RDS Proxy vượt tải tại ~500 tenant (kết nối trực tiếp làm đầy max_connections). Ở 1K tenant: 1K × 10 tác vụ ECS × 4 kết nối postgres = 40K >> giới hạn tối đa 3.200 của RDS. RDS Proxy không phải tùy chọn — phải triển khai trong hạ tầng Sprint 1 trước khi có tenant sản xuất đầu tiên. |
| **Cơ chế dự phòng Multi-region active-active (khu vực EU)** | 🟡 Cao | Phiên bản v2 bao gồm tài liệu khôi phục thảm họa (DR runbook) + thực hành chuyển vùng dự phòng (failover drill). Sau v2: nếu doanh thu từ thị trường EU đạt trên 30% ARR, nâng cấp vùng `eu-west-1` từ chế độ DR dự phòng thụ động sang chế độ chạy song song active-active. Đòi hỏi thiết lập cụm ECS độc lập + kiểm toán quy trình sao chép dữ liệu RDS xuyên vùng. |
| **Giám sát độ lệch mô hình BERT (Model drift)** | 🟡 Trung bình | Nếu tính năng phát hiện prompt injection bằng mô hình BERT được phát hành ở bản v2: sử dụng công cụ SageMaker Model Monitor (để phát hiện độ lệch dữ liệu và độ lệch khái niệm). Điều kiện kích hoạt huấn luyện lại: Tỷ lệ dương tính giả (FPR) vượt quá mức 3% trong khoảng thời gian theo dõi liên tục 30 ngày trên môi trường thực tế. Yêu cầu xây dựng sẵn đường ống dữ liệu gán nhãn. |
| **Nâng cao chất lượng đường ống kiểm tra bảo mật SCA/SAST** | 🟡 Trung bình | Tích hợp công cụ `govulncheck` (cho Go) + `pip-audit` (cho Python) vào CI trước phiên bản v1. Sau v2: bổ sung công cụ kiểm tra bảo mật động DAST (OWASP ZAP) vào môi trường staging, tự động hóa quy trình merge các PR từ Dependabot đối với các bản vá lỗi bảo mật có rủi ro thấp. |

### 5.3 Mở rộng Sản phẩm

| Hạng mục phát triển | Cột mốc mục tiêu | Ghi chú chi tiết |
|---|---|---|
| **Phát hiện deepfake cho khu vực EU (Mô-đun D1)** | Tháng 15–18 | Đặc trưng sinh trắc học giọng nói = dữ liệu thuộc danh mục đặc biệt theo Điều 9 GDPR. Đòi hỏi phải có ý kiến pháp lý độc lập + xây dựng cơ chế chấp thuận rõ ràng của nhân viên trước khi triển khai thực tế tại khu vực EU. Bắt đầu đánh giá pháp lý từ Tháng 10 (trước phiên bản v1.5 chứ không phải Tháng 6). Bàn giao tính năng tại Anh, Úc ở bản v2; thị trường EU chỉ bàn giao sau khi nhận được sự chấp thuận pháp lý đầy đủ. |
| **Tuân thủ Đạo luật AI của EU (EU AI Act)** | Tháng 15–18 | Nếu các tính năng phát hiện đe dọa từ AI của SMESec bị phân loại là "AI có độ rủi ro cao" theo Đạo luật AI của EU (Phụ lục III) đối với các khách hàng Enterprise thuộc EU: đòi hỏi phải có đánh giá mức độ tuân thủ, hoàn thiện tài liệu kỹ thuật, bổ sung các quyền kiểm soát giám sát của con người và đăng ký vào cơ sở dữ liệu của EU. Thuê tư vấn chuyên gia pháp lý từ Tháng 10. |
| **Tích hợp phần mềm kế toán QuickBooks / Xero** | Tháng 14–18 | Trì hoãn từ bản v2. Dành riêng cho các khách hàng SME có đội ngũ kế toán/tài chính: Phát hiện gian lận hóa đơn + phát hiện điểm bất thường trong ủy quyền thanh toán (hỗ trợ thêm cho giải pháp phòng chống deepfake). Đòi hỏi thực hiện phân tích đánh giá phạm vi tuân thủ PCI DSS riêng biệt. |
| **Gói dịch vụ MSSP / Giao diện tùy chỉnh thương hiệu (White-label)** | Tháng 15 trở đi | Phiên bản v2 cung cấp nền tảng cơ sở (hỗ trợ Enterprise đa thuê, tích hợp SIEM). Sau v2: phát triển giao diện tùy chỉnh thương hiệu, trang quản trị dành riêng cho MSP, tích hợp API tính phí dựa trên lượng tiêu thụ thực tế, cổng thông tin đối tác. Chương trình đối tác MSP được khởi động từ Tháng 1 (thông qua BD Consultant) — sản phẩm phải sẵn sàng ngay khi hợp đồng kinh doanh MSSP đầu tiên được ký kết. |
| **Tích hợp QuickBooks / Xero bổ sung** | Tháng 16 trở đi | Phát hiện điểm bất thường trong tài chính + lớp phòng chống gian lận hóa đơn. Đã được hoãn lại. Yêu cầu đánh giá phạm vi xử lý dữ liệu tài chính riêng biệt. |

### 5.4 Bảo mật & Quyền riêng tư

| Hạng mục bảo mật | Mốc thời gian | Ghi chú chi tiết |
|---|---|---|
| **Quy trình Thực thi Quyền được xóa dữ liệu (GDPR) ở quy mô lớn** | Duy trì liên tục | Khi lượng khách hàng tăng lên: các yêu cầu xóa dữ liệu số lượng lớn (khi xóa một tenant hoàn toàn) phải hoàn tất trong vòng cam kết SLA 30 ngày. Thực hiện bài kiểm tra mô phỏng xóa 100+ tenant đồng thời trước Tháng 12. Việc xóa khóa KMS + ẩn danh hóa dữ liệu PII liên quan phải được tự động hóa hoàn toàn. |
| **Duy trì minh chứng kiểm soát liên tục theo ISO 27001** | Duy trì liên tục | Đánh giá giám sát định kỳ đòi hỏi các minh chứng về việc thực thi kiểm soát liên tục (chứ không chỉ mang tính thời điểm). PM bắt buộc phải duy trì việc kiểm tra Vanta + đối chiếu nhật ký kiểm toán nội bộ. Cần lấy ngẫu nhiên minh chứng cho 10% các mục kiểm soát hàng tháng. |
| **Cam kết SLA xử lý các lỗi bảo mật CVE của Keycloak** | Duy trì liên tục | Lỗi CVE mức Nghiêm trọng trong Keycloak: phải vá lỗi trong vòng 72 giờ (triển khai nâng cấp cuốn chiếu ECS). Lỗi CVE mức Cao: phải vá lỗi trong vòng 7 ngày. Quy trình xử lý phải được viết thành tài liệu và thực hiện diễn tập định kỳ. Nếu chuyển đổi sang WorkOS/Auth0 → phải tiến hành thu hồi hệ thống Keycloak cũ một cách bảo mật (chuyển đổi phiên làm việc của người dùng, thu hồi hiệu lực token cũ, xóa sạch dữ liệu liên quan). |
| **Đánh giá bảo mật các nhà cung cấp bên thứ ba hàng năm** | Hàng năm | Thực hiện thu thập báo cáo SOC 2 + đánh giá tài liệu DPA hàng năm của các bên: Hive Moderation, Lakera Guard, Vanta. Nếu bất kỳ đối tác nào bị tước chứng nhận SOC 2 → lập tức lên kế hoạch tìm kiếm đối tác thay thế. |

### 5.5 Chiến lược Tiếp thị và Phát triển Thị trường (Tăng trưởng Sau v2)

- **Phát triển chương trình đối tác MSP:** Cố vấn Phát triển Kinh doanh (BD Consultant) thiết lập 3 mối quan hệ hợp tác MSP trong Năm 1 (Tháng 1-12). Năm 2: chính thức hóa các cấp độ đối tác, xây dựng mô hình chia sẻ doanh thu và đồng hành quảng bá thương hiệu (co-marketing). Chi phí sở hữu khách hàng (CAC) qua kênh MSP ($500–800) so với bán hàng trực tiếp ($3,000–5,000) khiến đây trở thành đòn bẩy tăng trưởng cốt lõi.
- **Tối ưu hóa chuyển đổi từ miễn phí sang trả phí (Freemium → paid conversion):** Gói miễn phí "Kiểm tra Sức khỏe Bảo mật" (trải nghiệm trong 14 ngày, tối đa 5 người dùng) phải đạt tỷ lệ chuyển đổi sang gói trả phí >15%. Giám sát tỷ lệ chuyển đổi của các nhóm khách hàng theo từng kênh tiếp cận. Thực hiện thử nghiệm A/B kiểm tra thời gian dùng thử và phân cấp các tính năng bị khóa.
- **Chăm sóc khách hàng ở quy mô lớn (Customer Success at scale):** Kỹ sư hỗ trợ khách hàng bắt đầu làm việc từ Tháng 7 (v1.5). Đến Tháng 13: định nghĩa các dấu hiệu cảnh báo sớm nguy cơ mất khách hàng (tỷ lệ DAU/MAU, tốc độ xác nhận cảnh báo rủi ro, trạng thái hoạt động của các tích hợp). Xây dựng hệ thống tính điểm sức khỏe khách hàng tự động.
- **Tính phí dựa trên mức độ sử dụng thực tế (Usage-based billing):** Phiên bản v2 bao gồm cấu hình tùy chọn này. Sau v2: Triển khai mô hình tính phí phân cấp theo lượng tiêu thụ thực tế đối với các tính năng của Luồng 2 (ví dụ: số lượt kiểm tra deepfake, số sự kiện DLP) — giúp giảm bớt rào cản tài chính ban đầu cho doanh nghiệp SME đồng thời tối ưu doanh thu khi quy mô khách hàng mở rộng.

### 5.6 Top 5 Rủi ro Sau v2 cần Giám sát chặt chẽ

| # | Rủi ro | Xác suất | Tác động | Tín hiệu cảnh báo cần theo dõi |
|---|---|---|---|---|
| 1 | **Lỗ hổng minh chứng tuân thủ SOC 2 Type 2 trong Năm thứ 2** — đội ngũ lơ là chủ quan sau khi đạt chứng nhận v2 | Trung bình | Nghiêm trọng | Điểm số đánh giá hàng tuần trên Vanta giảm xuống dưới 90% |
| 2 | **Khai thác lỗi bảo mật CVE chưa được vá của Keycloak** — gánh nặng vận hành làm trì trệ việc nâng cấp | Trung bình | Nghiêm trọng | Xuất hiện lỗi CVE được công bố với điểm số CVSS >8.0 nhưng hệ thống chưa được áp dụng bản vá lỗi trong vòng 72 giờ |
| 3 | **Thất bại trong việc tuân thủ Đạo luật AI của EU** — các tính năng deepfake/DLP bị phân loại là rủi ro cao nhưng chưa qua đánh giá mức độ tuân thủ | Trung bình | Cao | Khách hàng gói Enterprise thuộc khu vực EU ký kết hợp đồng trước khi ý kiến đánh giá pháp lý hoàn tất |
| 4 | **Vượt quá giới hạn tần suất gọi API của Google khi đạt trên 70+ khách thuê** — dẫn đến lỗi đồng bộ diện rộng + tăng nguy cơ hủy dịch vụ | Cao (phụ thuộc vào tốc độ tăng quy mô) | Cao | Tổng lượng tài nguyên API sử dụng vượt quá ngưỡng 70% định mức hiển thị trong CloudWatch |
| 5 | **Lakera Guard tăng giá → dẫn đến lỗ hổng phòng chống chèn lệnh** — nhà cung cấp tăng giá hoặc dừng cung cấp gói giá rẻ cho SME | Thấp | Trung bình | Nhà cung cấp thay đổi giá bán hoặc chất lượng dịch vụ SLA giảm sút → đòi hỏi đẩy nhanh quy trình đánh giá mô hình BERT tự phát triển nội bộ |
