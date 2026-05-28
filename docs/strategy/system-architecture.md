# SMESec Platform — System Architecture

**Ngày tạo:** 2026-05-28  
**Trạng thái:** Approved  
**Phiên bản:** 1.0  
**Tác giả:** Technical Advisor (30 năm cybersecurity + cloud architecture)

---

## Mục Lục

1. [Tổng Quan Kiến Trúc](#1-tổng-quan-kiến-trúc)
2. [Clean Architecture Principles](#2-clean-architecture-principles)
3. [Logical Architecture — Các Tầng Hệ Thống](#3-logical-architecture--các-tầng-hệ-thống)
4. [AWS Deployment Architecture](#4-aws-deployment-architecture)
5. [Integration Touchpoints](#5-integration-touchpoints)
6. [Data Architecture & Multi-Tenancy](#6-data-architecture--multi-tenancy)
7. [Security Architecture](#7-security-architecture)
8. [Track 1 vs Track 2 — Separation of Concerns](#8-track-1-vs-track-2--separation-of-concerns)
9. [Non-Functional Requirements](#9-non-functional-requirements)
10. [Build vs Buy Decisions](#10-build-vs-buy-decisions)
11. [Sơ Đồ Kiến Trúc (go-diagrams)](#11-sơ-đồ-kiến-trúc-go-diagrams)

---

## 1. Tổng Quan Kiến Trúc

SMESec là một **unified security platform cho SMEs (10–500 nhân viên)** được xây dựng theo mô hình SaaS multi-tenant. Platform bảo vệ các tài sản quan trọng nhất của doanh nghiệp vừa và nhỏ: data, accounts, intellectual property, và operational continuity — trong bối cảnh AI-driven threats ngày càng gia tăng.

### Quyết Định Kiến Trúc Cốt Lõi

| Quyết định | Lựa chọn | Lý do |
|---|---|---|
| **Build vs Buy** | Hybrid | Build core domain logic (differentiator); Buy commodity services (Keycloak, Vanta, Hive) |
| **Multi-tenancy** | Shared infrastructure, isolated data (Row-Level Security) | Cost-efficient; PostgreSQL RLS đủ mạnh cho SME scale; physical isolation quá tốn kém |
| **AI-threat detection** | 2-track: deterministic (Track 1) + ML/AI (Track 2) | Track 1 = 100% accuracy, high trust; Track 2 = R&D-gated, không ảnh hưởng core reliability |
| **Data privacy** | `data_residency` trên mọi bảng; EU data ở `eu-west-1`; no training on customer data | GDPR compliance; customer trust; không dùng customer data để train ML models |
| **Architecture pattern** | Clean Architecture + Event-Driven | Domain logic độc lập; adapter pattern cho integration; event sourcing cho audit trail |
| **Infrastructure** | AWS (primary) + Cloudflare R2 (CDN/storage) | AWS SLA + managed services; Cloudflare giảm egress cost |
| **Runtime** | Go (backend) + Python (ML/scripts) + React (web) + Flutter (mobile) | Go cho concurrency + type safety; Python cho ML ecosystem |

---

## 2. Clean Architecture Principles

SMESec áp dụng **Clean Architecture** (Robert C. Martin) kết hợp với **Hexagonal Architecture** (Ports & Adapters):

```
┌──────────────────────────────────────────────────────────┐
│  INTERFACE LAYER (Adapters, Controllers, Presenters)      │
│  ┌────────────────────────────────────────────────────┐  │
│  │  APPLICATION LAYER (Use Cases, Orchestrators)       │  │
│  │  ┌──────────────────────────────────────────────┐  │  │
│  │  │  DOMAIN LAYER (Entities, Domain Services)    │  │  │
│  │  │  ● Asset    ● TenantUser   ● ThreatEvent    │  │  │
│  │  │  ● Playbook ● ComplianceCtrl ● AccessPolicy │  │  │
│  │  └──────────────────────────────────────────────┘  │  │
│  └────────────────────────────────────────────────────┘  │
│                                                            │
│  INFRASTRUCTURE LAYER (Repositories, External Adapters)   │
└──────────────────────────────────────────────────────────┘
```

### Dependency Rule

> **Mọi source code dependency chỉ được phép trỏ VÀO TRONG** — về phía Domain Layer. Domain không biết gì về Infrastructure, Application không biết gì về Interface.

```
Interface → Application → Domain ← Infrastructure
```

### Ports & Adapters

**Ports (interfaces định nghĩa trong Domain):**

```go
// Domain/Ports — Primary Ports (driven by Interface Layer)
type AssetInventoryUseCase interface {
    DiscoverAssets(ctx context.Context, tenantID string) ([]Asset, error)
    ClassifyAsset(ctx context.Context, assetID string, classification Classification) error
    GetInventorySnapshot(ctx context.Context, tenantID string) (*InventorySnapshot, error)
}

// Domain/Ports — Secondary Ports (driven by Application Layer, implemented by Infrastructure)
type AssetRepository interface {
    Save(ctx context.Context, asset Asset) error
    FindByTenant(ctx context.Context, tenantID string) ([]Asset, error)
    FindByID(ctx context.Context, id string) (*Asset, error)
}

type GoogleWorkspacePort interface {
    ListUsers(ctx context.Context, tenantID string) ([]ExternalUser, error)
    ListOAuthApps(ctx context.Context, tenantID string) ([]OAuthApp, error)
    RevokeUserAccess(ctx context.Context, userID string) error
}

type ThreatEventPublisher interface {
    Publish(ctx context.Context, event ThreatDetectionEvent) error
}
```

### Layer Responsibilities

| Layer | Trách nhiệm | Dependencies |
|---|---|---|
| **Domain** | Entities, Value Objects, Domain Services, Aggregates, Domain Events, Repository Interfaces | Không có — zero external dependencies |
| **Application** | Use Cases, Orchestration, Transaction boundaries, DTO mapping | Domain only |
| **Infrastructure** | Repository implementations (PostgreSQL), Integration adapters (Google, M365, Slack), Event publishers (EventBridge), External clients (Vanta, Hive) | Application + Domain interfaces |
| **Interface** | REST handlers (Echo), gRPC handlers, WebSocket server, React SSR, Flutter bridge, Browser Extension service worker | Application Use Cases only |

---

## 3. Logical Architecture — Các Tầng Hệ Thống

### 3.1 Interface Layer

```
┌─────────────────────────────────────────────────────────────────┐
│                        INTERFACE LAYER                          │
│                                                                 │
│  ┌─────────────┐  ┌──────────────┐  ┌────────────────────────┐ │
│  │ Web App     │  │ Mobile App   │  │ Browser Extension       │ │
│  │ React/Next.js│ │ Flutter      │  │ Chrome MV3 + Edge       │ │
│  │ Port: 443   │  │ iOS + Android│  │ Content Script + SW     │ │
│  └──────┬──────┘  └──────┬───────┘  └──────────┬─────────────┘ │
│         │                │                       │               │
│         └────────────────┴───────────────────────┘               │
│                          │ HTTPS / WSS                           │
│  ┌───────────────────────▼──────────────────────────────────┐   │
│  │              API Gateway (AWS API Gateway + Kong)         │   │
│  │  REST (v1) · gRPC (internal) · WebSocket (real-time)     │   │
│  │  Auth: JWT (Keycloak) · Rate Limiting · WAF               │   │
│  └──────────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────────┘
```

### 3.2 Application Layer — Use Cases

```
┌─────────────────────────────────────────────────────────────────┐
│                       APPLICATION LAYER                         │
│                                                                 │
│  ┌──────────────────┐  ┌───────────────────┐                   │
│  │ AssetInventory   │  │ AccessGovernance   │                   │
│  │ Service          │  │ Service            │                   │
│  │ · DiscoverAssets │  │ · OffboardUser     │                   │
│  │ · ClassifyAssets │  │ · EnforceJIT       │                   │
│  │ · DetectShadowIT │  │ · ReviewAccess     │                   │
│  │ · DetectShadowAI │  │ · RemediatePolicy  │                   │
│  └──────────────────┘  └───────────────────┘                   │
│  ┌──────────────────┐  ┌───────────────────┐                   │
│  │ ThreatDetection  │  │ IncidentPlaybook   │                   │
│  │ Service (T2)     │  │ Service            │                   │
│  │ · DetectLLMLeak  │  │ · ExecutePlaybook  │                   │
│  │ · ScoreShadowAI  │  │ · TriggerResponse  │                   │
│  │ · VerifyDeepfake │  │ · CollectEvidence  │                   │
│  │ · FlagInjection  │  │ · NotifyStakeholder│                   │
│  └──────────────────┘  └───────────────────┘                   │
│  ┌──────────────────┐  ┌───────────────────┐                   │
│  │ Compliance       │  │ Integration        │                   │
│  │ Service          │  │ Sync Service       │                   │
│  │ · MapControls    │  │ · SyncGoogle       │                   │
│  │ · CollectEvidence│  │ · SyncM365         │                   │
│  │ · GenerateReport │  │ · SyncSlack        │                   │
│  │ · TrackPosture   │  │ · SyncAWSIAM       │                   │
│  └──────────────────┘  └───────────────────┘                   │
└─────────────────────────────────────────────────────────────────┘
```

### 3.3 Domain Layer — Core Business Logic

```
┌─────────────────────────────────────────────────────────────────┐
│                         DOMAIN LAYER                            │
│                                                                 │
│  AGGREGATES & ENTITIES                                          │
│  ┌──────────┐ ┌──────────┐ ┌───────────┐ ┌──────────────────┐  │
│  │  Asset   │ │TenantUser│ │ThreatEvent│ │ ComplianceControl│  │
│  │ id       │ │ id       │ │ id        │ │ id               │  │
│  │ tenantID │ │ tenantID │ │ tenantID  │ │ framework        │  │
│  │ type     │ │ email    │ │ type      │ │ controlID        │  │
│  │ critical.│ │ roles[]  │ │ severity  │ │ status           │  │
│  │ owner    │ │ providers│ │ source    │ │ evidenceRefs[]   │  │
│  └──────────┘ └──────────┘ └───────────┘ └──────────────────┘  │
│  ┌──────────┐ ┌──────────┐ ┌───────────┐                        │
│  │ Playbook │ │AccessPolicy│ │TenantConfig│                     │
│  │ id       │ │ id       │ │ id        │                        │
│  │ tenantID │ │ tenantID │ │ plan      │                        │
│  │ steps[]  │ │ subject  │ │ dataResid.│                        │
│  │ triggers │ │ resource │ │ providers │                        │
│  └──────────┘ └──────────┘ └───────────┘                        │
│                                                                 │
│  DOMAIN SERVICES                                                │
│  ┌──────────────┐ ┌───────────────┐ ┌─────────────────────┐    │
│  │ RiskScorer   │ │ AccessGovernor│ │ ComplianceAuditor   │    │
│  │ (pure logic) │ │ (pure logic)  │ │ (pure logic)        │    │
│  └──────────────┘ └───────────────┘ └─────────────────────┘    │
│                                                                 │
│  DOMAIN EVENTS (published via ThreatEventPublisher port)       │
│  AssetDiscovered · ThreatDetected · AccessRevoked              │
│  PlaybookTriggered · ComplianceViolated · OffboardingCompleted  │
└─────────────────────────────────────────────────────────────────┘
```

### 3.4 Infrastructure Layer — Adapters

```
┌─────────────────────────────────────────────────────────────────┐
│                     INFRASTRUCTURE LAYER                        │
│                                                                 │
│  REPOSITORIES (implements Domain ports)                         │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │ PostgresAssetRepo · PostgresUserRepo · PostgresThreatRepo│   │
│  │ PostgresPlaybookRepo · PostgresComplianceRepo            │   │
│  │ [All enforce tenant_id + data_residency RLS policies]    │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
│  INTEGRATION ADAPTERS (implements Domain ports)                 │
│  ┌──────────────────────────┐  ┌───────────────────────────┐   │
│  │ GoogleWorkspaceAdapter   │  │ M365Adapter               │   │
│  │ - Admin SDK              │  │ - Graph API               │   │
│  │ - Workspace Events API   │  │ - Azure AD Webhooks       │   │
│  │ - 15-min delta sync      │  │ - Delta Link sync         │   │
│  └──────────────────────────┘  └───────────────────────────┘   │
│  ┌──────────────────────────┐  ┌───────────────────────────┐   │
│  │ SlackAdapter             │  │ AWSIAMAdapter             │   │
│  │ - Admin API (tier-gated) │  │ - IAM API                 │   │
│  │ - Events API             │  │ - CloudTrail events       │   │
│  │ - Webhook receiver       │  │ - STS assume role         │   │
│  └──────────────────────────┘  └───────────────────────────┘   │
│                                                                 │
│  EVENT PUBLISHERS                                               │
│  ┌──────────────────────────┐  ┌───────────────────────────┐   │
│  │ EventBridgePublisher     │  │ SNSNotificationPublisher  │   │
│  │ - ThreatDetectionEvent   │  │ - Email alerts            │   │
│  │ - AccessEvent            │  │ - Slack webhook alerts    │   │
│  │ - PlaybookTrigger        │  │ - PagerDuty P1            │   │
│  └──────────────────────────┘  └───────────────────────────┘   │
│                                                                 │
│  EXTERNAL CLIENTS                                               │
│  VantaClient · HiveModerationClient · SageMakerClient          │
│  KeycloakClient · CloudflareR2Client                           │
└─────────────────────────────────────────────────────────────────┘
```

---

## 4. AWS Deployment Architecture

### 4.1 Logical Deployment Zones

```
┌─────────────────────────────────────────────────────────────────┐
│  INTERNET ZONE                                                  │
│  Clients: Web App · Mobile App · Browser Extension             │
│                          │ HTTPS                                │
│  ┌────────────────────────▼────────────────────────────────┐   │
│  │  EDGE ZONE (AWS Global)                                  │   │
│  │  Route 53 → CloudFront → WAF → ALB (us-east-1)          │   │
│  └────────────────────────┬────────────────────────────────┘   │
│                            │                                    │
│  ┌─────────────────────────▼───────────────────────────────┐   │
│  │  AWS VPC (us-east-1 / eu-west-1)                         │   │
│  │                                                           │   │
│  │  ┌──────────────────────────────────────────────────┐   │   │
│  │  │  PRIVATE SUBNET — AUTH & API                      │   │   │
│  │  │  Keycloak (ECS Fargate, Multi-AZ)                 │   │   │
│  │  │  API Gateway (ECS Fargate, Multi-AZ)              │   │   │
│  │  └──────────────────────────────────────────────────┘   │   │
│  │  ┌──────────────────────────────────────────────────┐   │   │
│  │  │  PRIVATE SUBNET — APPLICATION SERVICES (ECS)     │   │   │
│  │  │  Track 1: AssetSvc · AccessSvc · PlaybookSvc      │   │   │
│  │  │           ComplianceSvc · SyncSvc                 │   │   │
│  │  │  Track 2: ThreatDetectionSvc · DLPSvc             │   │   │
│  │  │           DeepfakeSvc · PhishingSvc               │   │   │
│  │  └──────────────────────────────────────────────────┘   │   │
│  │  ┌──────────────────────────────────────────────────┐   │   │
│  │  │  PRIVATE SUBNET — DATA LAYER                      │   │   │
│  │  │  RDS PostgreSQL Multi-AZ · ElastiCache Redis      │   │   │
│  │  │  S3 (Object Lock, WORM, 7yr) · ECR               │   │   │
│  │  └──────────────────────────────────────────────────┘   │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                  │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │  AWS MANAGED SERVICES (outside VPC)                       │   │
│  │  EventBridge · Step Functions · SNS/SQS · SageMaker       │   │
│  │  Secrets Manager · KMS · GuardDuty · Security Hub         │   │
│  │  CloudWatch · CloudTrail · IAM                             │   │
│  └──────────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────────┘
```

### 4.2 Service Communication Pattern

```
Client Request Flow:
  Browser/Mobile
      ↓ HTTPS
  CloudFront (edge cache, SSL termination)
      ↓ HTTPS
  WAF (OWASP Top 10 rules, rate limiting)
      ↓ HTTPS
  ALB (Layer 7, health checks, SSL offload)
      ↓ HTTP
  API Gateway Service (ECS Fargate)
      → JWT validation (Keycloak introspect)
      → Route to downstream service
      ↓ HTTP (internal VPC)
  Application Service (ECS Fargate)
      → PostgreSQL (sync read/write, RLS enforced)
      → ElastiCache Redis (session, cache, rate limit)
      → EventBridge (async event publish)

Async Flow:
  Application Service → EventBridge → Step Functions (playbook)
                                    → Lambda (scheduled jobs)
                                    → SNS → Email/Slack/PagerDuty
```

### 4.3 ECS Fargate Service Configuration

| Service | CPU | Memory | Min Tasks | Max Tasks | Track |
|---|---|---|---|---|---|
| API Gateway | 512 | 1024 MB | 2 | 10 | Shared |
| Asset Inventory | 256 | 512 MB | 2 | 8 | 1 |
| Access Governance | 256 | 512 MB | 2 | 8 | 1 |
| Playbook Engine | 512 | 1024 MB | 2 | 6 | 1 |
| Compliance | 256 | 512 MB | 1 | 4 | 1 |
| Integration Sync | 512 | 1024 MB | 2 | 8 | 1 |
| Threat Detection | 1024 | 2048 MB | 1 | 6 | 2 |
| LLM DLP | 512 | 1024 MB | 1 | 4 | 2 |
| Deepfake Defense | 256 | 512 MB | 1 | 4 | 2 |
| Keycloak | 512 | 1024 MB | 2 | 4 | Auth |

---

## 5. Integration Touchpoints

### 5.1 Identity Provider Integrations

| Provider | Protocol | Data Collected | Write Operations | Sync Frequency |
|---|---|---|---|---|
| **Google Workspace** | OAuth 2.0 (Admin SDK) | Users, Groups, OAuth Apps, Devices, Audit Logs | Disable user, Revoke OAuth, Suspend account | 15-min delta via Workspace Events API |
| **Microsoft 365** | OAuth 2.0 (Graph API) | Users, Groups, OAuth Apps, Devices, SignIn Logs | Disable user, Revoke sessions, Block signin | 15-min delta via Delta Link + Webhooks |
| **Slack** | OAuth 2.0 (Admin API) | Users, Channels, OAuth Apps, Audit Logs | Deactivate user (Business+ only), Remove app | 15-min poll + Events API webhooks |
| **AWS IAM** | AWS SDK (assumed role) | Users, Roles, Policies, Access Keys, CloudTrail | Disable access key, Remove policy (dry-run first) | 30-min full pull |

### 5.2 OAuth Scope Policy (Minimum Permissions)

```
Google Workspace (Service Account):
  admin.directory.user.readonly
  admin.directory.group.readonly
  admin.reports.audit.readonly
  admin.directory.device.chromeos.readonly
  → Write (offboarding): admin.directory.user (Suspend)
  → Write (oauth revoke): admin.directory.userschema

Microsoft 365 (App Registration):
  User.Read.All
  Application.Read.All
  AuditLog.Read.All
  DeviceManagementManagedDevices.Read.All
  → Write (offboarding): User.ReadWrite.All (scoped)

Slack (App):
  admin.users:read
  admin.apps:read
  auditlogs:read
  → Write (offboarding): admin.users:write (Business+)
```

### 5.3 Third-Party Security Services

| Service | Purpose | Integration Type | Data Shared |
|---|---|---|---|
| **Vanta** | Compliance automation (SOC 2, ISO 27001) | API + OAuth (AWS, GitHub) | Infrastructure metadata, not customer data |
| **Hive Moderation** | Deepfake detection (voice/video) | REST API (pay-per-use) | Audio/video hash, not raw content |
| **Keycloak** | SSO + MFA (self-hosted on ECS) | Internal service | Session tokens, user credentials |
| **SageMaker** | ML model inference (shadow AI, risk scoring) | AWS SDK | Anonymized feature vectors, not PII |
| **Cloudflare R2** | Audit log archive, asset storage | S3-compatible API | Encrypted audit logs |

---

## 6. Data Architecture & Multi-Tenancy

### 6.1 Multi-Tenancy Model: Shared Database, Isolated Data

**Chosen approach:** Single PostgreSQL cluster, Row-Level Security (RLS) enforced at database level.

```sql
-- Mọi bảng domain đều có hai cột bắt buộc
CREATE TABLE assets (
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id     UUID NOT NULL,          -- FK to tenants table
    data_residency VARCHAR(10) NOT NULL   -- 'US' | 'EU' | 'APAC'
        CHECK (data_residency IN ('US', 'EU', 'APAC')),
    -- ... domain columns ...
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- RLS Policy — mọi query tự động scoped theo tenant_id
ALTER TABLE assets ENABLE ROW LEVEL SECURITY;
CREATE POLICY tenant_isolation ON assets
    USING (tenant_id = current_setting('app.tenant_id')::UUID);

-- Index bắt buộc trên tenant_id để performance
CREATE INDEX idx_assets_tenant ON assets(tenant_id);
```

**Enforcement tại API Middleware (Go):**

```go
// Middleware inject tenant_id vào mọi PostgreSQL connection
func TenantMiddleware(db *sql.DB) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            tenantID := c.Get("tenant_id").(string)
            // Set PostgreSQL session variable — RLS reads this
            _, err := db.ExecContext(c.Request().Context(),
                "SET LOCAL app.tenant_id = $1", tenantID)
            if err != nil {
                return echo.ErrInternalServerError
            }
            return next(c)
        }
    }
}
```

### 6.2 ThreatDetectionEvent — Shared Schema Contract (T1 ↔ T2)

```go
// Định nghĩa trong domain/events package — cả Track 1 và Track 2 đều dùng
type ThreatDetectionEvent struct {
    EventID        string          `json:"event_id"`    // UUID v4
    TenantID       string          `json:"tenant_id"`   // Non-nullable
    DataResidency  string          `json:"data_residency"` // 'US'|'EU'|'APAC'
    Source         string          `json:"source"`      // 'track1'|'track2'
    EventType      string          `json:"event_type"`  // shadow_it|llm_dlp|deepfake|phishing
    Severity       string          `json:"severity"`    // LOW|MEDIUM|HIGH|CRITICAL
    ActorUserID    string          `json:"actor_user_id"`
    AssetRefs      []string        `json:"asset_refs"`  // FKs into asset_inventory
    RawPayload     json.RawMessage `json:"raw_payload"`
    MLMetadata     *MLEventMeta    `json:"ml_metadata,omitempty"` // nil for Track 1
    OccurredAt     time.Time       `json:"occurred_at"`
    DetectedAt     time.Time       `json:"detected_at"`
    SchemaVersion  string          `json:"schema_version"` // 'v1' — semver
}

// Không được phép thay đổi schema sau Sprint 6 mà không có RFC + migration plan
```

### 6.3 Audit Log Architecture (Immutable)

```
Write Path:
  Application Event → PostgreSQL (append-only table, no UPDATE/DELETE)
                    → S3 Object Lock (WORM, 7-year retention, AES-256)
                    → CloudWatch Logs (operational, 90-day retention)

Read Path:
  Compliance Dashboard → PostgreSQL (indexed queries)
  Auditor Export → S3 pre-signed URL (time-limited, tenant-scoped)
  SIEM Integration → CloudWatch Logs Insights
```

### 6.4 Data Residency Routing

```
EU Tenants (data_residency = 'EU'):
  → ECS Tasks: eu-west-1
  → RDS: eu-west-1 (primary) + eu-central-1 (read replica)
  → S3: eu-west-1 bucket (separate from US)
  → Vanta: EU data never leaves EU

US/APAC Tenants (data_residency = 'US'|'APAC'):
  → ECS Tasks: us-east-1
  → RDS: us-east-1 (primary) + us-west-2 (read replica)
  → S3: us-east-1 bucket
```

---

## 7. Security Architecture

### 7.1 Defense in Depth

```
Layer 1 — Network:
  CloudFront (DDoS protection, edge caching)
  WAF (OWASP Top 10, custom rules, geo-blocking)
  Security Groups (deny-by-default, least-privilege)
  VPC (private subnets, no public IPs on app servers)
  NAT Gateway (outbound-only internet access for app tier)

Layer 2 — Authentication & Authorization:
  Keycloak (OIDC/OAuth 2.0 + SAML 2.0)
  MFA: TOTP bắt buộc cho tất cả users
  JWT (RS256, 15-min access token, 7-day refresh token)
  RBAC: Tenant Admin > Security Admin > IT Staff > Read-Only
  API Gateway: JWT validation on every request

Layer 3 — Application:
  RLS: tenant_id enforced at PostgreSQL level
  Input validation: struct tags + custom validators (no raw SQL)
  CSRF protection: SameSite=Strict cookies
  CSP headers: strict-dynamic + nonce
  Rate limiting: ElastiCache Redis token bucket (per tenant + per endpoint)

Layer 4 — Data:
  Encryption at rest: RDS AES-256, S3 AES-256 (KMS CMK)
  Encryption in transit: TLS 1.3 everywhere (no TLS 1.0/1.1)
  Secrets: AWS Secrets Manager (auto-rotation enabled)
  PII anonymization: SHA-256 hash for user identifiers in ML feature vectors

Layer 5 — Operations:
  GuardDuty: anomaly detection on API calls + data access
  Security Hub: centralized security findings
  CloudTrail: immutable API audit trail
  Dependabot + CodeQL: automated dependency + SAST scanning
  Pen-test: external vendor, bi-annual
```

### 7.2 OAuth Token Security

```
Google/M365/Slack OAuth tokens:
  - Stored: AWS Secrets Manager (not DB) — encrypted + auto-rotated
  - Access: Only by Integration Sync Service (IAM role scoped)
  - Rotation: Refresh tokens rotated on every use
  - Revocation: Immediate on tenant deactivation

Keycloak Sessions:
  - Access token: 15 minutes (short-lived)
  - Refresh token: 7 days (sliding, revocable)
  - Stored: Redis (not PostgreSQL) — eviction on logout
  - MFA: Required on every new device/IP
```

---

## 8. Track 1 vs Track 2 — Separation of Concerns

### Architecture Boundary

```
TRACK 1 (Deterministic, High Confidence)          TRACK 2 (ML/AI, R&D-Gated)
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━            ━━━━━━━━━━━━━━━━━━━━━━━━━━━━
 Rule-based logic (100% accuracy)                  ML models (target >90%)
 PostgreSQL-centric                                SageMaker + Python/FastAPI
 Sync + Async processing                           Async-only (inference queue)
 Available from Sprint 1                           Available from Sprint 7
 Powers: Asset Inventory, Access                   Powers: Shadow AI scoring,
         Governance, Playbooks,                            LLM DLP, Deepfake,
         Compliance, Integrations                          Prompt injection

                          SHARED CONTRACT
                  ThreatDetectionEvent (schema v1)
                  EventBridge (event bus)
                  PostgreSQL (shared asset_inventory)
                  Tenant isolation (RLS)
```

### Track 2 Data Pipeline

```
Runtime Flow:
  Track 1 Event (e.g., new OAuth app discovered)
      → EventBridge bus: "track1.asset.discovered"
      → Track 2 ThreatDetection Service (ECS Fargate)
          → Enrich with asset context (query Track 1 PostgreSQL read replica)
          → Run SageMaker inference endpoint (shadow AI risk model)
          → Publish ThreatDetectionEvent to EventBridge
      → EventBridge: "track2.threat.detected"
          → Track 1 Playbook Service (Step Functions trigger)
          → SNS notification (email + Slack)
          → Compliance Service (evidence collection)

Browser Extension Flow:
  User types in ChatGPT textarea
      → Content Script (Chrome MV3)
      → Service Worker: PII detection (local Presidio WASM)
      → If HIGH risk: block submit + show warning
      → Async: POST /api/v1/dlp-events (pseudonymized)
      → Track 2 LLM DLP Service: log + alert IT admin
```

---

## 9. Non-Functional Requirements

### Performance Targets

| Metric | Target | Measured From |
|---|---|---|
| API P99 latency | <200ms | v1 launch |
| Shadow IT alert delay | <15 min from OAuth grant | v1 launch |
| Offboarding completion | <5 min end-to-end | MVP |
| JIT access revocation | <60 sec | v1 |
| Dashboard load time | <2 sec (P95) | v1 launch |
| ML inference latency | <500ms (shadow AI scoring) | v1 Track 2 |
| Compliance report generation | <10 sec | v1 |

### Availability & Reliability

| Metric | Target | Phase |
|---|---|---|
| Uptime SLA | Best effort (pilot) | MVP |
| Uptime SLA | 99.9% (43 min/mo) | v1 launch |
| Uptime SLA | 99.95% | v2 |
| RTO (Recovery Time Objective) | <4 hours | v1 |
| RPO (Recovery Point Objective) | <1 hour | v1 |
| Multi-AZ | Active-active (ECS + RDS) | Sprint 1 |

### Scalability

```
Current target (v1): 100 tenants × 500 users = 50,000 users
Scale assumption: 15-min sync × 50,000 users = ~56 events/sec peak

Horizontal scaling:
  ECS Fargate: Auto-scaling (CloudWatch → target 60% CPU)
  RDS: Read replicas (add for analytics queries)
  Redis: ElastiCache cluster mode (shard if >100K keys)
  EventBridge: Serverless (scales automatically)
  SageMaker: Endpoint auto-scaling (provisioned concurrency)
```

---

## 10. Build vs Buy Decisions

| Component | Decision | Rationale |
|---|---|---|
| **Authentication (SSO + MFA)** | Buy: Keycloak (self-hosted) | Open-source, OIDC/SAML, no per-user cost, full control |
| **Compliance automation** | Buy: Vanta | $4-6K/yr vs 3 months engineer time. Evidence collection + auditor portal. |
| **Deepfake detection** | Buy: Hive Moderation API | Pay-per-use (<$0.01/check), no model training cost, maintained by vendor |
| **ML platform** | Buy: SageMaker | Managed training, endpoints, model registry — vs 6 months infra build |
| **Browser extension DLP** | Build: Chrome MV3 | No SME-priced competitor. Local inference (privacy). Custom allow-list per tenant. |
| **Integration sync engine** | Build: Custom (Go) | Provider-specific quirks (Google rate limits, M365 delta links, Slack tier detection) require custom handling |
| **Asset inventory** | Build: Custom (Go) | Core differentiator. No competitor at SME price with Shadow AI detection. |
| **Incident playbooks** | Build on AWS Step Functions | Step Functions = proven orchestration. Build the playbook logic + wizard UI. |
| **Audit logging** | Build on S3 Object Lock | S3 Object Lock = WORM compliance-ready at near-zero cost. |
| **Observability** | Buy: CloudWatch + Datadog | CloudWatch = native AWS. Datadog for APM if budget allows. |

---

## 11. Sơ Đồ Kiến Trúc (go-diagrams)

Các sơ đồ được sinh tự động bằng Go programs sử dụng thư viện [blushft/go-diagrams](https://github.com/blushft/go-diagrams). Source code trong thư mục `diagrams/`.

### Cách Chạy

```bash
# Yêu cầu: Go 1.21+, Graphviz (dot)
# Cài Graphviz: https://graphviz.org/download/

cd diagrams

# Diagram 1: Logical Architecture (Clean Architecture layers)
go run cmd/logical/main.go
dot -Tpng go-diagrams/out/logical-architecture.dot -o go-diagrams/out/logical-architecture.png

# Diagram 2: AWS Deployment Architecture
go run cmd/deployment/main.go
dot -Tpng go-diagrams/out/deployment-architecture.dot -o go-diagrams/out/deployment-architecture.png

# Diagram 3: Integration Touchpoints
go run cmd/integrations/main.go
dot -Tpng go-diagrams/out/integration-touchpoints.dot -o go-diagrams/out/integration-touchpoints.png

# Hoặc chạy tất cả 1 lần:
make diagrams   # xem diagrams/Makefile
```

### Sơ Đồ 1: Logical Architecture

> **File:** `diagrams/cmd/logical/main.go`  
> **Mô tả:** Clean Architecture layers — dependency flow từ Interface → Application → Domain ← Infrastructure

### Sơ Đồ 2: AWS Deployment Architecture

> **File:** `diagrams/cmd/deployment/main.go`  
> **Mô tả:** Physical deployment trên AWS — VPC, ECS Fargate, RDS, EventBridge, SageMaker, Security services

### Sơ Đồ 3: Integration Touchpoints

> **File:** `diagrams/cmd/integrations/main.go`  
> **Mô tả:** Integration với third-party SaaS (Google Workspace, M365, Slack, AWS IAM) và security services (Vanta, Hive, Keycloak)
