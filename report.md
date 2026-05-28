# SMESec Platform - System Design Document

**Date:** 2026-05-28  
**Version:** 1.0  
**Status:** Draft

---

## Executive Summary

Small and medium enterprises (10-500 employees) face escalating AI-driven security risks—automated spear-phishing, AI-generated disinformation, data leakage to public LLMs, shadow AI adoption, and supply-chain compromise—yet lack dedicated security teams and enterprise budgets. SMESec addresses this gap with a unified protection platform covering critical assets: data, people, intellectual property, financial accounts, and operational continuity.

**Two-Track Development Strategy:** To mitigate the inherent risk of AI detection accuracy (false positives frustrate users; false negatives cause breaches), we split development into parallel tracks. **Track 1 (Foundation & Governance)** delivers deterministic, high-confidence capabilities—asset inventory, access governance, automated offboarding, incident playbooks, and compliance reporting—achieving near-100% accuracy with proven technologies. **Track 2 (AI Threat Detection)** pursues high-value, high-risk ML-based detection for prompt injection, data leakage, and deepfakes, with strict validation gates (>95% precision, <5% false positive) before launch.

This approach ensures Track 1 can launch independently after 6 months (13 sprints) with 5-10 pilot customers, providing immediate value while Track 2 continues R&D. Track 2 merges into the product only after passing four validation gates confirming production-ready accuracy. If Track 2 accuracy targets are not met, it remains in beta while Track 1 delivers compliance-ready asset and access management to the SME market.

---

## 1. System Architecture Diagram

### 1.1 Logical Architecture View

```
┌─────────────────────────────────────────────────────────────────────┐
│                         SMESec Platform                              │
├─────────────────────────────────────────────────────────────────────┤
│                                                                       │
│  ┌──────────────────────┐         ┌──────────────────────┐         │
│  │   Track 1: Foundation │         │  Track 2: AI Detection│         │
│  │   & Governance        │         │                       │         │
│  ├──────────────────────┤         ├──────────────────────┤         │
│  │                       │         │                       │         │
│  │ • Asset Inventory     │         │ • Prompt Injection    │         │
│  │ • Classification      │         │   Detection (3-layer) │         │
│  │ • Access Governance   │         │ • DLP Engine          │         │
│  │   - RBAC (OPA/Rego)   │         │ • Dynamic Redaction   │         │
│  │   - JIT Access        │         │ • Deepfake Detection  │         │
│  │ • Offboarding Engine  │         │   - Voice (vendor API)│         │
│  │ • Playbook Engine     │         │   - Video (vendor API)│         │
│  │   (Step Functions)    │         │ • Shadow AI Discovery │         │
│  │ • Compliance Mapping  │         │ • Risk Scoring        │         │
│  │ • Evidence Collection │         │                       │         │
│  └──────────┬────────────┘         └──────────┬───────────┘         │
│             │                                  │                      │
│             └──────────────┬───────────────────┘                      │
│                            │                                          │
│                   ┌────────▼────────┐                                │
│                   │  EventBridge    │                                │
│                   │  Integration    │                                │
│                   └────────┬────────┘                                │
│                            │                                          │
└────────────────────────────┼──────────────────────────────────────────┘
                             │
        ┌────────────────────┼────────────────────┐
        │                    │                    │
   ┌────▼─────┐      ┌──────▼──────┐     ┌──────▼──────┐
   │ Google   │      │ Microsoft   │     │   Slack     │
   │Workspace │      │    365      │     │             │
   └──────────┘      └─────────────┘     └─────────────┘
        │                    │                    │
   ┌────▼─────┐      ┌──────▼──────┐     ┌──────▼──────┐
   │   AWS    │      │   Azure     │     │    GCP      │
   └──────────┘      └─────────────┘     └─────────────┘
```

**Key Components:**
- **Asset Inventory Engine**: Automated discovery of devices, accounts, SaaS apps, cloud resources
- **Access Governance Layer**: RBAC policy engine (OPA), JIT access workflows, automated offboarding
- **Playbook Engine**: AWS Step Functions-based incident response automation
- **AI Detection Service**: 3-layer ML pipeline (regex → BERT classifier → context analysis)
- **DLP & Redaction Engine**: PII/IP detection with dynamic masking and de-redaction
- **Compliance Module**: ISO 27001, GDPR, SOC 2 control mapping and evidence automation
- **Integration Hub**: OAuth 2.0-based connectors for Google, M365, Slack, AWS, Azure, GCP

### 1.2 Deployment Architecture View

```
┌─────────────────────────────────────────────────────────────────────┐
│                         AWS Cloud (Multi-Region)                     │
├─────────────────────────────────────────────────────────────────────┤
│                                                                       │
│  ┌──────────────────────────────────────────────────────────────┐  │
│  │                    VPC (Tenant Isolated)                      │  │
│  │                                                                │  │
│  │  ┌─────────────────┐         ┌─────────────────┐            │  │
│  │  │  ECS Fargate    │         │  ECS Fargate    │            │  │
│  │  │  (Track 1 API)  │◄────────┤  (Track 2 API)  │            │  │
│  │  │  Go + Python    │         │  Python/FastAPI │            │  │
│  │  └────────┬────────┘         └────────┬────────┘            │  │
│  │           │                            │                      │  │
│  │           └──────────┬─────────────────┘                      │  │
│  │                      │                                        │  │
│  │           ┌──────────▼──────────┐                            │  │
│  │           │   RDS PostgreSQL    │                            │  │
│  │           │   (Multi-AZ)        │                            │  │
│  │           │   Row-Level Security│                            │  │
│  │           └─────────────────────┘                            │  │
│  │                                                                │  │
│  └────────────────────────────────────────────────────────────────┘
│                                                                       │
│  ┌──────────────────────────────────────────────────────────────┐  │
│  │                    Shared Services                            │  │
│  │                                                                │  │
│  │  ┌─────────────┐  ┌──────────────┐  ┌──────────────┐        │  │
│  │  │ EventBridge │  │ Step Functions│  │  SageMaker   │        │  │
│  │  │  (Events)   │  │  (Playbooks)  │  │  (ML Models) │        │  │
│  │  └─────────────┘  └──────────────┘  └──────────────┘        │  │
│  │                                                                │  │
│  │  ┌─────────────┐  ┌──────────────┐  ┌──────────────┐        │  │
│  │  │     S3      │  │    KMS       │  │   Secrets    │        │  │
│  │  │  (Evidence) │  │ (Encryption) │  │   Manager    │        │  │
│  │  └─────────────┘  └──────────────┘  └──────────────┘        │  │
│  │                                                                │  │
│  └────────────────────────────────────────────────────────────────┘
│                                                                       │
│  ┌──────────────────────────────────────────────────────────────┐  │
│  │                    Auth & Gateway                             │  │
│  │                                                                │  │
│  │  ┌─────────────┐  ┌──────────────┐  ┌──────────────┐        │  │
│  │  │  Keycloak   │  │  API Gateway │  │  CloudFront  │        │  │
│  │  │  (SSO/MFA)  │  │  (Rate Limit)│  │     (CDN)    │        │  │
│  │  └─────────────┘  └──────────────┘  └──────────────┘        │  │
│  │                                                                │  │
│  └────────────────────────────────────────────────────────────────┘
│                                                                       │
└─────────────────────────────────────────────────────────────────────┘
                                │
                                │ HTTPS/TLS 1.3
                                │
        ┌───────────────────────┼───────────────────────┐
        │                       │                       │
   ┌────▼─────┐         ┌──────▼──────┐        ┌──────▼──────┐
   │   Web    │         │   Mobile    │        │  Browser    │
   │Dashboard │         │    App      │        │ Extension   │
   │(Next.js) │         │  (Flutter)  │        │(Chrome/Edge)│
   └──────────┘         └─────────────┘        └─────────────┘
```

**Infrastructure:**
- **Compute**: ECS Fargate (serverless containers), auto-scaling based on CPU/memory
- **Database**: RDS PostgreSQL Multi-AZ with row-level security for tenant isolation
- **Storage**: S3 with Object Lock (7-year retention for compliance evidence)
- **Security**: AWS KMS for encryption at rest, Secrets Manager for credentials
- **Orchestration**: Step Functions for playbook workflows, EventBridge for event routing
- **ML Infrastructure**: SageMaker for model training and inference (Track 2)
- **Auth**: Keycloak for SSO/MFA, supporting Google and Microsoft identity providers
- **CDN**: CloudFront for global content delivery and DDoS protection
- **Monitoring**: CloudWatch for logs/metrics, X-Ray for distributed tracing

### 1.3 Integration Touchpoints

#### Third-Party SaaS Integrations
| Service | Integration Method | Purpose | APIs Used |
|---------|-------------------|---------|-----------|
| Google Workspace | OAuth 2.0 + Admin SDK | Asset discovery, access control | Admin API, Audit API |
| Microsoft 365 | OAuth 2.0 + Graph API | Asset discovery, access control | Graph API, Azure AD |
| Slack | OAuth 2.0 + Admin API | Asset discovery, notifications | Admin API, Audit Logs |
| AWS | IAM + Config API | Cloud asset discovery | Config, IAM, CloudTrail |

#### On-Premise Asset Integration
- **Network Discovery**: Agent-based discovery for on-premise devices (Windows, Linux, macOS)
- **Active Directory**: LDAP integration for user/group synchronization
- **File Servers**: SMB/NFS scanning for sensitive data classification
- **Databases**: JDBC/ODBC connectors for PostgreSQL, MySQL, SQL Server, Oracle
- **VPN Gateway**: Integration for remote worker device inventory

#### Integration Security Model
All integrations follow OAuth 2.0 with least-privilege scopes:
- **Read-only by default**: Asset discovery requires only read permissions
- **Write permissions**: Explicitly requested for access revocation (offboarding)
- **Credential storage**: All tokens encrypted in AWS Secrets Manager with KMS
- **Token rotation**: Automatic refresh before expiration, manual rotation every 90 days
- **Audit trail**: All API calls logged to CloudWatch with tenant_id, user, action, timestamp

---

## 2. Design Document - Core Architectural Decisions

### 2.1 Build vs Buy (Hybrid Approach)

**Decision:** SMESec adopts a **hybrid strategy** that builds core differentiation while integrating proven third-party services for commodity capabilities. This balances time-to-market, cost efficiency, and competitive moat.

#### Components to Build
| Component | Rationale |
|-----------|-----------|
| **Asset Inventory Engine** | Core IP: multi-source aggregation logic, dependency mapping, classification rules tailored to SME context |
| **Access Governance (RBAC + JIT)** | Differentiation: SME-optimized policies, automated offboarding workflows, zero-touch provisioning |
| **Playbook Engine** | Unique value: non-security staff can execute incident response via wizard UI, pre-built SME playbooks |
| **Compliance Mapping** | Competitive advantage: automated evidence collection, one-click audit reports for ISO 27001/GDPR/SOC 2 |
| **AI Detection Pipeline (3-layer)** | Core R&D: proprietary ML models fine-tuned on SME threat patterns, context-aware risk scoring |
| **DLP + Dynamic Redaction** | Differentiation: redact-send-de-redact workflow preserves LLM utility while protecting data |
| **Browser Extension** | Privacy-first monitoring: AI tool domains only, user opt-out, transparent data handling |

#### Components to Buy/Integrate
| Component | Vendor/Service | Rationale |
|-----------|---------------|-----------|
| **SSO/MFA** | Keycloak (open-source) | Proven, standards-compliant (SAML, OAuth 2.0), self-hosted for data sovereignty |
| **Deepfake Detection** | Sensity AI / Reality Defender | Specialized vendor APIs achieve >90% accuracy; building in-house would require 12+ months R&D |
| **ML Infrastructure** | AWS SageMaker | Managed training/inference, auto-scaling, experiment tracking — faster than building MLOps from scratch |
| **Workflow Orchestration** | AWS Step Functions | Serverless, fault-tolerant, visual workflow editor — proven for incident playbooks |
| **Identity Providers** | Google Workspace, Microsoft 365 | Leverage existing SME identity infrastructure via OAuth 2.0; no need to replicate |
| **Notification Services** | AWS SES (email), Slack API, FCM/APNs (push) | Commodity services with high deliverability and reliability |

### 2.2 Multi-Tenancy Model

**Architecture Pattern:** Row-Level Security (RLS) with tenant_id enforcement at database and application layers

**Rationale:** RLS provides strong isolation guarantees while maintaining operational simplicity (single database, single schema) and cost efficiency for SME scale (10-500 employees per tenant). Schema-per-tenant or database-per-tenant would introduce operational overhead (migrations, backups, monitoring) that exceeds SME requirements.

#### Tenant Isolation Strategy

**Data Layer (PostgreSQL RLS):**
- Every table includes `tenant_id UUID NOT NULL` column with index
- Row-Level Security policies enforce `tenant_id = current_setting('app.current_tenant')::uuid`
- Database connection sets `app.current_tenant` immediately after authentication
- Foreign key constraints include `tenant_id` to prevent cross-tenant references
- Automated CI tests verify zero cross-tenant data leakage (critical acceptance criterion Sprint 1)

**Application Layer (API Middleware):**
- JWT token contains `tenant_id` claim, validated on every request
- API middleware extracts `tenant_id` from JWT and sets PostgreSQL session variable
- All queries automatically scoped to tenant via RLS policies
- Audit log records `tenant_id` + `user_id` + `action` + `resource` for every API call
- Rate limiting applied per-tenant (prevents noisy neighbor problem)

**Infrastructure Layer (AWS):**
- Shared ECS Fargate cluster with tenant-scoped resource tagging
- S3 evidence storage: `s3://smesec-evidence/{tenant_id}/{year}/{month}/{evidence_id}.json`
- S3 bucket policies enforce tenant_id prefix matching
- KMS encryption keys: one master key per region, tenant_id in encryption context
- CloudWatch logs tagged with `tenant_id` for per-tenant observability

#### Security Controls

**Tenant Isolation Verification:**
- **Automated CI test (Sprint 1):** Create 2 test tenants, insert data for Tenant A, query as Tenant B → must return 0 rows
- **Penetration test (Sprint 13):** External auditor attempts cross-tenant access via API manipulation, JWT tampering, SQL injection
- **Runtime monitoring:** CloudWatch alarm triggers if any query returns rows with mismatched `tenant_id`

**Data Encryption:**
- **At rest:** AWS KMS encryption for RDS (AES-256), S3 (SSE-KMS), EBS volumes
- **In transit:** TLS 1.3 for all API traffic, certificate pinning for mobile app
- **Tenant-scoped encryption context:** KMS encrypt/decrypt operations include `tenant_id` in context, preventing key reuse across tenants

**Access Control:**
- **Database credentials:** Stored in AWS Secrets Manager, rotated every 90 days, never logged
- **API authentication:** Keycloak SSO with MFA enforced, JWT tokens expire after 1 hour
- **Service-to-service:** IAM roles with least-privilege policies, no long-lived credentials

### 2.3 AI Threat Detection Strategy

**Approach:** 3-layer hybrid detection combining deterministic rules, ML classification, and contextual risk scoring to achieve >95% precision with <5% false positive rate.

**Rationale:** Single-layer detection fails for AI threats. Rule-based alone misses novel attacks; ML alone produces too many false positives for SME tolerance; context-free scoring blocks legitimate workflows. The 3-layer cascade provides defense-in-depth while maintaining usability.

#### Detection Layers

1. **Layer 1: Rule-Based Detection (Deterministic)**
   - **Technology:** Regex patterns + Luhn algorithm validation
   - **Coverage:** 50+ OWASP LLM Top 10 patterns (direct injection, role manipulation, jailbreaks)
   - **PII Detection:** Credit cards, SSN, email, phone, API keys, credentials
   - **Performance:** <10ms latency, can run client-side (WASM)
   - **Accuracy target:** >90% on known patterns, 0% false negatives on critical data (credit cards)
   - **Rationale:** Catches obvious attacks instantly with zero false positives on well-defined patterns

2. **Layer 2: ML Classifier (BERT Fine-Tuned)**
   - **Model:** BERT-base-uncased fine-tuned on 5K+ labeled prompt injection examples
   - **Input:** Prompt text (max 512 tokens)
   - **Output:** Risk score 0-100 + confidence
   - **NER Model:** spaCy + custom entities for PII/IP detection
   - **Performance:** <500ms inference latency (p95) via SageMaker
   - **Accuracy target:** >95% precision on severe class, >95% DLP precision on critical data
   - **Rationale:** Detects novel injection attempts and semantic data leakage that regex misses

3. **Layer 3: Context Analysis (Risk Multiplier)**
   - **Context Sources:** User role (from Track 1 RBAC), data sensitivity, historical patterns, application context
   - **Risk Multipliers:**
     - Admin user: 0.5x (less suspicious)
     - Employee with PII access: 2.0x (critical)
     - First-time AI tool user: 1.5x
     - Repeated similar prompts: 0.7x (likely legitimate workflow)
   - **Performance:** <100ms context enrichment
   - **Accuracy target:** Reduce false positives by >30% vs Layer 2 alone
   - **Rationale:** Same prompt has different risk depending on who, what, where — context prevents blocking legitimate work

#### Response Actions by Risk Score
| Score Range | Level | Action | Playbook |
|-------------|-------|--------|----------|
| 0-30 | Low | Log only | None |
| 31-60 | Medium | Advisory alert + request justification | None |
| 61-85 | High | Block + require manager approval | Shadow IT Detected |
| 86-100 | Critical | Block + immediate IT admin alert + revoke access | Account Compromise |

#### Validation Gates

| Gate | Sprint | Metrics | Pass Criteria |
|------|--------|---------|---------------|
| **Gate 1** | S4 (W8) | Prompt injection precision, DLP precision | Injection >90% severe class; DLP >95% critical data |
| **Gate 2** | S6 (W12) | False positive rates | Injection FP <10%; DLP FP <5% |
| **Gate 3** | S9 (W18) | All targets met | Injection >95%; DLP >99%; Deepfake voice >90%, video >85% |
| **Gate 4** | S12 (W24) | Real-world pilot validation | Customer satisfaction >4.0/5.0; 0 critical incidents; FP complaints <10% |

**Quality-First Policy:** Track 2 launches ONLY if all Gate 4 criteria met. If accuracy insufficient after 6 months, Track 2 continues as beta while Track 1 launches independently. No compromises on accuracy — false negatives cause breaches, false positives destroy trust.

### 2.4 Data Privacy Guarantees

**Privacy-by-Design Principles:**

1. **Data Minimization**: Collect only what's necessary for security functions. Asset discovery reads metadata (usernames, device IDs, app names) but not file contents or email bodies unless explicitly required for DLP scanning.

2. **Purpose Limitation**: Data used solely for stated security purposes. AI prompt monitoring limited to AI tool domains (chatgpt.com, copilot.microsoft.com, etc.) — never social media, banking, or personal browsing.

3. **User Control & Transparency**: 
   - Browser extension displays active monitoring status
   - Users can disable extension (IT admin receives notification)
   - Dashboard shows what data was collected and why
   - Data deletion requests honored within 30 days (GDPR Article 17)

4. **Encryption Everywhere**: All customer data encrypted at rest (AWS KMS AES-256) and in transit (TLS 1.3). Redaction mappings encrypted with tenant-scoped KMS context, expire after 24 hours.

5. **Audit Trail Immutability**: All access events logged to S3 with Object Lock (append-only, 7-year retention). Logs cannot be deleted or modified, even by SMESec administrators.

#### Data Handling

| Data Type | Storage | Encryption | Retention | Access Control |
|-----------|---------|------------|-----------|----------------|
| **Customer credentials** | AWS Secrets Manager | KMS (AES-256) | Rotated every 90 days | IAM roles only, no human access |
| **Audit logs** | S3 (Object Lock) | SSE-KMS | 7 years (compliance) | Read-only via API, tenant-scoped |
| **AI prompts (monitored)** | S3 (tenant-scoped) | SSE-KMS | 90 days (configurable) | Tenant admins only, encrypted in transit |
| **Asset metadata** | RDS PostgreSQL | KMS (at-rest) | Active + 1 year post-deletion | Row-level security (tenant_id) |
| **Evidence artifacts** | S3 (Object Lock) | SSE-KMS | 7 years (compliance) | Immutable, tenant-scoped |
| **Redaction mappings** | Redis (encrypted) | TLS + KMS | 24 hours (auto-expire) | Ephemeral, tenant-scoped |
| **ML training data** | S3 (anonymized) | SSE-KMS | Indefinite (research) | Anonymized (no PII), aggregated across tenants |

#### Compliance Alignment

**GDPR (General Data Protection Regulation):**
- **Article 25 (Privacy by Design)**: Encryption, data minimization, purpose limitation built into architecture
- **Article 30 (Records of Processing)**: Asset inventory includes data processing activities
- **Article 32 (Security Measures)**: Encryption, access controls, audit logs, MFA enforcement
- **Article 17 (Right to Erasure)**: Offboarding workflow includes data deletion; manual deletion requests honored within 30 days
- **Article 33 (Breach Notification)**: Incident playbooks include 72-hour breach notification workflow
- **Article 13/88 (Browser Extension Disclosure)**: Extension privacy policy displayed on install, explains monitoring scope and user rights

**ISO 27001 (Information Security Management):**
- **A.8.1 (Asset Management)**: Automated asset inventory with classification
- **A.8.2 (Information Classification)**: Sensitivity levels (Restricted/Confidential/Internal/Public)
- **A.9.1 (Access Control Policy)**: RBAC with least-privilege enforcement
- **A.9.2 (User Access Management)**: Automated provisioning/deprovisioning
- **A.9.4 (Access Review)**: Quarterly access review reports
- **A.12.4 (Logging & Monitoring)**: All access events logged, immutable audit trail

**SOC 2 (Trust Services Criteria):**
- **CC6.1 (Logical Access)**: RBAC restricts access based on role
- **CC6.2 (Access Provisioning)**: JIT access with auto-expiration
- **CC6.3 (Access Removal)**: Automated offboarding <5 minutes
- **CC7.2 (System Monitoring)**: CloudWatch monitoring, real-time alerts for anomalies
- **CC7.3 (Audit Logs)**: Immutable logs with 7-year retention

**Data Sovereignty**: All customer data stored in AWS region selected by customer (default: us-east-1). Multi-region support planned for v1.2.

---

## 3. Team & Delivery Plan

### 3.1 Team Structure

#### Track 1: Foundation & Governance (5 FTE)
| Role | Count | Responsibilities |
|------|-------|------------------|
| **Tech Lead / Architect** | 1 | AWS infrastructure design, VPC/ECS/RDS architecture, Keycloak SSO configuration, multi-tenant data model, API architecture review, security threat modeling, code review for both backend engineers |
| **Backend Engineer** | 2 | Go/Python services, Google Workspace & M365 integration, RBAC engine (OPA/Rego), JIT access workflows, automated offboarding (Step Functions), playbook engine, compliance mapping, evidence collection automation |
| **Frontend Engineer** | 1 | React/Next.js web dashboard, asset inventory UI, access management UI, compliance reports UI, incident wizard UI, design system & component library |
| **Flutter Engineer** | 1 | Mobile/desktop app (iOS, Android, Desktop), asset inventory mobile view, JIT approval from mobile, incident wizard mobile, push notifications (FCM/APNs) |

#### Track 2: AI Threat Detection (3 FTE)
| Role | Count | Responsibilities |
|------|-------|------------------|
| **ML Engineer / Security Researcher** | 1 | Dataset collection & labeling, BERT fine-tuning for prompt injection, NER model for PII detection, deepfake detection evaluation, accuracy benchmarking, threshold tuning, SageMaker model deployment |
| **Backend Engineer (Python/FastAPI)** | 1 | Detection API (3-layer pipeline), DLP engine, dynamic redaction/de-redaction, risk scoring algorithm, deepfake API integration (Sensity/Reality Defender), EventBridge event publishing |
| **Frontend Engineer (Browser Extension)** | 1 | Chrome/Edge extension (MV3), prompt interception, domain monitoring, risk assessment UI, warning banners, usage analytics, privacy controls |

#### Shared Resources (2 FTE)
| Role | Count | Responsibilities |
|------|-------|------------------|
| **Product Manager / Security Analyst** | 1 | Sprint planning & ceremonies, pilot customer outreach & onboarding, vendor API procurement, legal/compliance review coordination, validation gate checkpoints, stakeholder communication |
| **DevSecOps / QA** | 1 | CI/CD pipeline (GitHub Actions), Terraform infrastructure-as-code, monitoring (CloudWatch/X-Ray), security hardening, penetration test coordination, load testing, deployment automation |

**Total Team Size:** 10 FTE

### 3.2 6-Month Delivery Sequence

#### Track 1: Foundation & Governance (Launch-Ready)

| Month | Focus | Key Deliverables | Milestone |
|-------|-------|------------------|-----------|
| **Month 1** (S1-S2) | **Infrastructure + Google Workspace** | AWS VPC/ECS/RDS, Keycloak SSO, multi-tenant DB schema, Google Workspace sync (users + OAuth apps), CI/CD pipeline | Dev environment operational, Google asset discovery >90% |
| **Month 2** (S3-S4) | **M365 Integration + Classification** | Microsoft 365 sync, Dashboard v1 (asset inventory table), auto-classification engine, shadow IT alerts, allow-list management | **Visibility Checkpoint**: Dashboard shows all assets from Google + M365 |
| **Month 3** (S5-S6) | **Access Governance + Offboarding** | Slack + AWS discovery, RBAC engine (OPA/Rego), JIT access workflows, automated offboarding (<5 min), offboarding PDF reports | **Offboarding Checkpoint**: Demo-ready for external stakeholders |
| **Month 4** (S7-S8) | **Playbooks + Mobile** | JIT access with auto-revoke, Playbook engine (Step Functions), 3 core playbooks (Account Compromise, Offboarding Emergency, Shadow IT), wizard UI, mobile app scaffold | Playbooks executable by non-security staff |
| **Month 5** (S9-S11) | **Compliance + Reports** | 2 additional playbooks, mobile app (iOS/Android), compliance mapping (ISO 27001/GDPR/SOC 2), evidence auto-collection, compliance dashboard, audit reports (PDF+JSON) | **Access Control Checkpoint**: Full governance + playbooks demo |
| **Month 6** (S12-S13) | **Hardening + Launch** | Dependency graph, lifecycle tracking, penetration test (external), load testing (500 assets, 50 users), security hardening, 5-10 pilot customers onboarded | **Beta Launch**: Production-ready, pen-test passed (0 Critical/High findings) |

#### Track 2: AI Threat Detection (Pilot-Ready)

| Month | Focus | Key Deliverables | Validation Gate |
|-------|-------|------------------|-----------------|
| **Month 1-2** (S1-S4) | **Research + Layer 1-2 Detection** | Dataset collection (5K+ labeled examples), ML infra (SageMaker), regex patterns (50+ OWASP LLM), BERT fine-tuning, NER model, DLP engine with dynamic redaction, browser extension scaffold | **Gate 1 (W8)**: Injection >90%, DLP >95% critical data |
| **Month 3-4** (S5-S6) | **Browser Extension + Context** | Extension v1 (Chrome/Edge, prompt interception, domain monitoring), Layer 3 context analysis, risk multipliers, response actions (0-100 scale), policy engine (OPA/Rego) | **Gate 2 (W12)**: Injection FP <10%, DLP FP <5% |
| **Month 5** (S7-S9) | **Deepfake + Shadow AI** | Voice deepfake detection (vendor API), video deepfake detection, out-of-band verification, incident response integration (EventBridge), shadow AI discovery, risk scoring framework | **Gate 3 (W18)**: All targets met (Injection >95%, DLP >99%, Deepfake voice >90%, video >85%) |
| **Month 6** (S10-S13) | **Integration + Pilot + Launch Decision** | Track 1-Track 2 integration (event-driven), 2-3 pilot customers onboarded, threshold tuning, real-world accuracy validation, pilot report | **Gate 4 (W24)**: Customer satisfaction >4.0/5.0, 0 critical incidents, FP <10% → Launch decision |

### 3.3 Riskiest Assumptions to Validate First

#### Critical Assumption #1: AI Detection Accuracy Achievable in 6 Months
- **Risk Level:** **CRITICAL** (Probability: Medium, Impact: Existential)
- **Impact if Wrong:** Track 2 fails all validation gates → 6 months of ML Engineer + Backend Eng + Extension Eng wasted → Track 1 launches alone (acceptable) but competitive differentiation lost → market positioning as "AI security platform" untenable
- **Validation Approach:** 
  - **Week 1 (Sprint 1):** Chrome MV3 service worker persistence prototype (HARD GATE) — if workers terminate after 30s, entire extension monitoring architecture breaks
  - **Week 6 (Gate 1):** Prompt injection >90% on test dataset — if fails, indicates dataset quality issue or model capacity insufficient
  - **Week 12 (Gate 2):** False positive rate <10% — if fails, indicates context layer insufficient or thresholds need major tuning
  - **Week 18 (Gate 3):** All accuracy targets met simultaneously — if fails, indicates fundamental approach flaw
- **Timeline:** Validate incrementally at each gate; final go/no-go decision at Week 24 (Gate 4)
- **Mitigation:** 
  - If Gate 1 fails: Add 2 weeks for dataset quality improvement + model re-training (delay Gate 2)
  - If Gate 2 fails: Add 4 weeks for threshold tuning + context enrichment (delay Gate 3)
  - If Gate 3 fails: Track 2 continues as beta, Track 1 launches independently
  - If Gate 4 fails: Track 2 remains beta indefinitely until accuracy proven in production

#### Critical Assumption #2: Pilot Customers Available for Track 2 Validation
- **Risk Level:** **HIGH** (Probability: Medium, Impact: High)
- **Impact if Wrong:** Gate 4 cannot execute → no real-world accuracy validation → Track 2 cannot launch even if benchmark accuracy met → 6-month investment yields no production-ready AI detection
- **Validation Approach:**
  - **Week 1 (Sprint 1):** PM begins pilot customer outreach (2-3 SMEs, 50-200 employees, using ChatGPT/Copilot, willing to share telemetry)
  - **Week 13 (Sprint 7):** 2+ signed NDAs/DPAs (HARD DEADLINE) — legal procurement takes 4-8 weeks, cannot be fast-tracked
  - **Week 18 (Sprint 9):** Pilot customers CONFIRMED and ready to onboard
  - **Week 21 (Sprint 11):** Pilot execution begins (4-week pilot: 2 weeks in S11, 2 weeks in S12)
- **Timeline:** Outreach starts Week 1; contracts signed by Week 13; pilot runs Week 21-24
- **Mitigation:**
  - If no customers by Week 13: Extend Track 2 timeline by 8 weeks to allow procurement + pilot
  - If customers drop out mid-pilot: Have 3 customers (not 2) to absorb 1 dropout
  - If zero customers available: Track 2 pivots to internal dogfooding only (lower confidence, longer beta period)

#### Critical Assumption #3: Vendor APIs (Deepfake Detection) Available & Affordable
- **Risk Level:** **HIGH** (Probability: Low, Impact: High)
- **Impact if Wrong:** Deepfake detection accuracy <70% (open-source fallback) → Gate 3 fails on deepfake criterion → Track 2 scope reduced (no deepfake in v1) → value proposition weakened
- **Validation Approach:**
  - **Week 1 (Sprint 1):** PM requests trial accounts from Sensity AI + Reality Defender
  - **Week 1 (Sprint 1):** Extension Eng implements DeepfakeDetector abstraction interface + Resemblyzer open-source fallback (~78-82% accuracy)
  - **Week 2-4:** Evaluate vendor APIs on FaceForensics++ benchmark (voice + video)
  - **Week 8 (Sprint 4-5):** Vendor contract signed (12-week lead time from Week 1 request)
  - **Week 13 (Sprint 7):** Vendor API integrated and tested
- **Timeline:** Vendor evaluation Week 1-4; contract signed by Week 8; integration complete Week 13
- **Mitigation:**
  - If vendor pricing too high (>$5K/year): Use Resemblyzer fallback, accept 78-82% accuracy (still useful as tripwire)
  - If vendor contract delayed: Continue with Resemblyzer, upgrade to vendor API post-launch
  - If vendor API unreliable: Implement hybrid (vendor primary, Resemblyzer fallback on vendor failure)

#### Critical Assumption #4: Track 1-Track 2 Schema Compatibility
- **Risk Level:** **MEDIUM** (Probability: Low, Impact: Medium)
- **Impact if Wrong:** Sprint 10 (Track 1-Track 2 integration) discovers incompatible EventBridge schemas → 2-4 weeks of refactoring both tracks → delays Track 1 launch OR Track 2 ships without integration (reduced value)
- **Validation Approach:**
  - **Week 1 (Sprint 1):** Joint T1-T2 schema session defines ThreatDetectionEvent interface
  - **Week 4 (End Sprint 2):** Schema frozen (no breaking changes allowed after this point)
  - **Week 15-16 (Sprint 8):** T1-T2 schema validation meeting confirms compatibility before S10 integration sprint
- **Timeline:** Schema defined Week 1, frozen Week 4, validated Week 15-16
- **Mitigation:**
  - If schema incompatibility discovered Week 15-16: Emergency 1-week sprint to align schemas before S10
  - If discovered during S10: Extend S10 by 1 sprint (delay Track 1 launch by 2 weeks)

#### Critical Assumption #5: Team Capacity Realistic (No Attrition, Sick Days, Context-Switching)
- **Risk Level:** **MEDIUM** (Probability: High, Impact: Medium)
- **Impact if Wrong:** Sprint velocity drops 20-30% → critical path extends → Track 1 launch delayed 2-4 weeks OR features cut to meet deadline
- **Validation Approach:**
  - **Sprint 1-2:** Measure actual velocity vs planned (person-days delivered vs allocated)
  - **Sprint 3:** Adjust sprint scope if velocity <80% of plan
  - **Monthly:** Review team utilization, identify bottlenecks (e.g., Tech Lead over-allocated at 90-100% across multiple sprints)
- **Timeline:** Continuous monitoring; adjust after Sprint 2 if velocity gap detected
- **Mitigation:**
  - If velocity <80%: Reduce sprint scope by 20%, extend timeline by 1 sprint (2 weeks)
  - If key person unavailable (sick, attrition): Flutter Eng (Dart/TS background) can assist Extension Eng; Backend Eng 2 can assist Backend Eng 1
  - If Tech Lead bottleneck: Delegate code reviews to senior Backend Eng, focus Tech Lead on architecture decisions only

### 3.4 Key Requirements Coverage

#### Requirement 1: Asset Inventory and Classification
- **Track:** Track 1 (Foundation & Governance)
- **Sprints:** S1-S4 (W1-W8)
- **Approach:** 
  - **Discovery:** Automated multi-source aggregation (Google Workspace, M365, Slack, AWS, Azure, GCP) via OAuth 2.0 APIs
  - **Classification:** Auto-classification by account type (admin/standard/service/contractor) + sensitivity levels (Restricted/Confidential/Internal/Public)
  - **Manual Override:** IT admin can override classification per asset or bulk via CSV import
  - **Dependency Mapping (S12):** Graph visualization showing user → OAuth app → cloud resource relationships
- **Success Criteria:** 
  - Asset discovery coverage >95% within 1 hour for 500-asset org
  - Classification accuracy >90% (auto-classification vs manual review)
  - Dashboard load time <2s for 10,000 assets

#### Requirement 2: AI-Specific Threat Surface
- **Track:** Track 2 (AI Threat Detection)
- **Sprints:** S1-S9 (W1-W18)
- **Approach:**
  - **Prompt Injection:** 3-layer detection (regex → BERT → context analysis)
  - **Data Leakage:** DLP engine with dynamic redaction (PII, credentials, IP)
  - **Deepfake Detection:** Vendor API integration (Sensity AI / Reality Defender) for voice + video
  - **Shadow AI Discovery:** OAuth app classification + browser extension monitoring
- **Success Criteria:**
  - Prompt injection precision >95%, false positive <5%
  - DLP precision >99% on critical data (credit cards, SSN), false negative <1%
  - Deepfake detection: voice >90%, video >85%
  - Shadow AI discovery >95% within 1 hour

#### Requirement 3: Access Governance
- **Track:** Track 1 (Foundation & Governance)
- **Sprints:** S5-S7 (W9-W14)
- **Approach:**
  - **RBAC Engine:** OPA/Rego policy evaluation (<100ms latency)
  - **Built-in Roles:** Admin, Manager, Employee, Contractor, Service Account
  - **JIT Access:** Request → Approve → Auto-revoke workflow with Slack/email notifications
  - **Automated Offboarding:** Parallel revocation across Google/M365/Slack/AWS via Step Functions (<5 min)
  - **Shadow IT Detection:** OAuth app allow-list with alerts for unapproved apps
- **Success Criteria:**
  - Offboarding completion <5 minutes (all platforms)
  - JIT access auto-revoked within <1 minute after expiration
  - Shadow IT detection rate >90%
  - Zero cross-tenant data leakage (verified by automated CI tests + penetration test)

#### Requirement 4: Continuous Compliance Posture
- **Track:** Track 1 (Foundation & Governance)
- **Sprints:** S10-S11 (W19-W22)
- **Approach:**
  - **Control Mapping:** ISO 27001 (A.8.1, A.8.2, A.9.1, A.9.2, A.9.4, A.12.4), GDPR (Art. 30, 32, 17, 33), SOC 2 (CC6.1, CC6.2, CC6.3, CC7.2)
  - **Evidence Auto-Collection:** Daily asset snapshots, access event logs, offboarding reports, incident reports → S3 Object Lock (7-year retention)
  - **Compliance Dashboard:** Control status (Implemented/Partial/Not Implemented), Audit Readiness Score (0-100), evidence links
  - **On-Demand Reports:** ISO 27001, GDPR, SOC 2 reports (PDF + JSON) generated in <5 minutes
- **Success Criteria:**
  - All Track 1 features mapped to relevant controls
  - Evidence auto-collected for 100% of mapped controls
  - Report generation <5 minutes
  - Audit Readiness Score >80/100 at launch

**Note:** V1 launch is **compliance-ready** (technical controls + evidence automation in place). **Certification audits** (SOC 2 Type II, ISO 27001) occur **post-launch** (6-12 months) as they require operational history.

#### Requirement 5: Incident Playbooks
- **Track:** Track 1 (Foundation & Governance)
- **Sprints:** S8-S9 (W15-W18)
- **Approach:**
  - **Playbook Engine:** AWS Step Functions (stateful, fault-tolerant, resume after restart)
  - **Wizard UI:** Step-by-step guidance with decision gates (Yes/No), progress indicator, undo support
  - **5 Pre-Built Playbooks:**
    1. Account Compromise (suspicious login detected)
    2. Offboarding Emergency (immediate termination)
    3. Shadow IT Detected (unapproved OAuth app)
    4. Unauthorized Access (user accessed resource without permission)
    5. Inactive Account (account unused >90 days)
  - **Notifications:** P0/P1 incidents → Email + Slack + Push; P2/P3 → Email only
  - **Track 2 Integration (S10):** AI threats trigger Track 1 playbooks via EventBridge
- **Success Criteria:**
  - Non-security staff can execute playbooks in <10 minutes without additional guidance
  - All playbook steps logged for audit
  - Mobile app supports playbook execution (iOS/Android)

#### Requirement 6: Cost Model
- **Track:** Separate workstream (parallel to Track 1/Track 2)
- **Approach:** Tiered, pay-as-you-grow pricing aligned with SME constraints
- **Pricing Strategy:** 
  - **Starter:** 10-50 employees, $X/user/month (Track 1 only)
  - **Professional:** 51-200 employees, $Y/user/month (Track 1 + Track 2 AI detection)
  - **Enterprise:** 201-500 employees, $Z/user/month (Track 1 + Track 2 + priority support)
  - **Add-ons:** Additional integrations, extended evidence retention, custom playbooks
- **Cost Drivers:** AWS infrastructure (ECS, RDS, S3, SageMaker), vendor APIs (deepfake detection), support overhead
- **Reference:** See [docs/pricing/cost-analysis.md](docs/pricing/cost-analysis.md) for detailed cost breakdown

#### Requirement 7: Integration with SME Tools
- **Track:** Track 1 (Foundation & Governance)
- **Sprints:** S1-S5 (W1-W10)
- **Approach:**
  - **Google Workspace (S2):** Admin SDK + Audit API for users, groups, OAuth apps
  - **Microsoft 365 (S3):** Graph API + Azure AD for users, groups, licensed apps
  - **Slack (S5):** Admin API + Audit Logs for users, channels, installed apps
  - **AWS (S5):** Config API + IAM API for EC2, S3, RDS, Lambda, IAM users
  - **Azure/GCP:** Planned for v1.1
  - **QuickBooks:** Planned for v1.1 (financial account monitoring)
- **Success Criteria:**
  - All integrations use OAuth 2.0 with least-privilege scopes
  - Incremental sync every 15 minutes (background job)
  - Partial failure (rate limit, token error) does not crash entire sync
  - >90% asset discovery within 1 hour for each integration

---

## 4. AI Governance Module

### 4.1 Shadow AI Discovery

**Objective:** Identify all AI tools (approved and unapproved) used within the organization to provide visibility into AI adoption patterns, data exposure risks, and compliance gaps.

#### Discovery Methods

1. **OAuth App Inventory (Track 1, Sprint 4)**
   - **Approach:** Scan Google Workspace and M365 OAuth app authorizations for AI tool patterns
   - **Classification Logic:** App name matching (ChatGPT, Copilot, Gemini, Claude, etc.) + domain analysis (openai.com, anthropic.com, cohere.ai, mistral.ai)
   - **Coverage:** >95% of OAuth-based AI tools within 1 hour
   - **Alert:** IT admin receives Slack + email notification when new AI tool detected and not in allow-list

2. **Browser Extension Monitoring (Track 2, Sprint 5)**
   - **Approach:** Chrome/Edge extension (MV3) intercepts prompts on AI tool domains before submission
   - **Supported Platforms:** ChatGPT (chatgpt.com), Microsoft Copilot (copilot.microsoft.com), Google Gemini (gemini.google.com), Kiro (kiro.ai), Anthropic Claude (claude.ai)
   - **Privacy Controls:** 
     - Monitoring limited to AI tool domains only (whitelist-based)
     - Does NOT monitor social media, banking, news, or personal browsing
     - User can disable extension (IT admin receives notification)
   - **Performance:** <50ms overhead per prompt, <5% CPU usage
   - **Detection Rate:** >95% of browser-based AI tool usage

3. **Network Traffic Analysis (Future: v1.1)**
   - **Approach:** DNS query analysis + API call pattern detection for AI service endpoints
   - **Detection Rate:** >90% of network-accessible AI tools
   - **Note:** Deferred to v1.1 due to complexity and infrastructure requirements

### 4.2 Risk Scoring Framework

**Risk Calculation:** Multi-factor weighted scoring (0-100 scale) combining data sensitivity, tool reputation, usage patterns, and user context.

**Formula:**
```
Base Risk Score = (Layer 1 Score × 0.2) + (Layer 2 Score × 0.5) + (Layer 3 Context × 0.3)

Final Risk Score = Base Risk Score × Context Multiplier

Context Multiplier = User Role Factor × Data Sensitivity Factor × Usage Pattern Factor
```

| Risk Factor | Weight | Measurement |
|-------------|--------|-------------|
| **Data Sensitivity** | 30% | Classification level of data in prompt (Restricted=2.0x, Confidential=1.5x, Internal=1.0x, Public=0.5x) |
| **Tool Reputation** | 25% | Known vendor (OpenAI, Anthropic, Google=0.8x) vs unknown (1.5x) vs blocked (3.0x) |
| **Usage Frequency** | 20% | First-time user=1.5x, regular user (>10 prompts/week)=0.7x, power user (>50/week)=0.9x |
| **Data Volume** | 15% | Prompt length: <500 chars=1.0x, 500-2000=1.2x, >2000=1.5x |
| **User Role** | 10% | Admin=0.5x, Manager=0.8x, Employee=1.0x, Contractor=1.5x, Service Account=2.0x |

#### Risk Levels & Actions
| Risk Level | Score Range | Automated Action | Human Review | Playbook Triggered |
|------------|-------------|------------------|--------------|-------------------|
| **Low** | 0-30 | Log only, no user interruption | None | None |
| **Medium** | 31-60 | Advisory alert (non-blocking), request justification | Optional (user can self-justify) | None |
| **High** | 61-85 | Block prompt, require manager approval via Slack/email | Manager approval required | Shadow IT Detected (if unapproved tool) |
| **Critical** | 86-100 | Block prompt, immediate IT admin alert, revoke access if repeated | IT admin review required | Account Compromise |

**Example Calculation:**
```
User: Employee (1.0x) with PII access (2.0x), regular AI user (0.7x)
Prompt: "Summarize this customer list: [500 names + emails]" to unknown AI tool
Layer 1: Detects PII (email regex) → 70
Layer 2: BERT detects data leakage intent → 85
Layer 3: Context enrichment → 80
Base Score: (70 × 0.2) + (85 × 0.5) + (80 × 0.3) = 80.5
Context Multiplier: 1.0 (role) × 2.0 (PII data) × 0.7 (regular user) = 1.4
Final Score: 80.5 × 1.4 = 112.7 → capped at 100 → CRITICAL
Action: Block + IT admin alert
```

### 4.3 Threat Detection Capabilities

#### 4.3.1 Prompt Injection Detection
- **Approach:** 3-layer cascade (regex → BERT → context analysis)
- **Accuracy Target:** >95% precision, <5% false positive
- **Response Time:** <500ms end-to-end (p95)

**Detection Patterns:**

**Layer 1 - Regex (50+ patterns, <10ms):**
- Direct injection: "Ignore previous instructions...", "Disregard all prior commands..."
- Role manipulation: "You are now DAN...", "Pretend you are..."
- System prompt extraction: "Repeat your instructions verbatim", "What are your rules?"
- Jailbreak patterns: "Developer mode enabled", "Sudo mode", "Unrestricted mode"
- Delimiter attacks: "---END SYSTEM---", "```system", "<!--OVERRIDE-->"

**Layer 2 - BERT Classifier (fine-tuned on 5K+ examples, <500ms):**
- Novel injection attempts not matching regex patterns
- Semantic manipulation (e.g., "As my grandmother used to say before bedtime...")
- Obfuscated attacks (e.g., base64-encoded instructions, Unicode tricks)
- Context-dependent attacks that require understanding intent

**Layer 3 - Context Analysis (<100ms):**
- User role: Admin users less likely to be attacking (0.5x multiplier)
- Historical patterns: User with 100+ legitimate prompts gets lower suspicion (0.7x)
- Data sensitivity: Prompt containing PII + injection pattern = higher risk (2.0x)
- Application context: ChatGPT vs unknown AI tool affects risk score

**Response Actions:**
- Score 0-30: Log only
- Score 31-60: Advisory warning (non-blocking)
- Score 61-85: Block + manager approval required
- Score 86-100: Block + IT admin alert + Account Compromise playbook triggered

#### 4.3.2 LLM Data Leakage Prevention (DLP)
- **Approach:** DLP engine with dynamic redaction/de-redaction
- **Accuracy Target:** >99% precision on critical data, <1% false negative
- **Response Time:** <100ms for redaction

**Protected Data Types:**
- **PII:** Credit cards (Luhn validation), SSN (XXX-XX-XXXX), email, phone, passport numbers
- **Credentials:** API keys (regex patterns for AWS, OpenAI, GitHub, etc.), passwords, tokens
- **Intellectual Property:** Source code snippets (detected via syntax patterns), confidential document fingerprints
- **Financial Data:** Bank account numbers, routing numbers, credit card CVV

**Dynamic Redaction Workflow:**
1. **Scan:** Detect PII/credentials in prompt before submission
2. **Redact:** Replace with tokens: `John Smith, card 4111-1111-1111-1111` → `[PERSON_1], card [CARD_1]`
3. **Send:** Redacted prompt sent to LLM
4. **Receive:** LLM response may contain `[PERSON_1]` or `[CARD_1]`
5. **De-redact:** Replace tokens back: `[PERSON_1]` → `John Smith`
6. **Deliver:** User sees response with original data restored

**Redaction Mapping Security:**
- Encrypted with AWS KMS, tenant-scoped encryption context
- Stored in Redis (in-memory, encrypted in transit)
- Auto-expire after 24 hours
- Never logged or persisted to disk

**Policy Engine Integration:**
- IT admin defines policies: "Block all credit cards", "Redact PII but allow", "Role-based: admins bypass"
- Policy evaluation <50ms via OPA/Rego
- Policies versioned and rollback-able

#### 4.3.3 Deepfake Detection

**Voice Detection (Vendor API + Fallback):**
- **Primary:** Sensity AI or Reality Defender API
  - Accuracy: >90% on FaceForensics++ benchmark
  - Response Time: <5 seconds per audio sample
  - Features: Liveness detection, voice cloning detection, synthetic speech detection
- **Fallback:** Resemblyzer (open-source)
  - Accuracy: ~78-82% (acceptable as tripwire)
  - Response Time: <3 seconds
  - Use case: Vendor API unavailable or budget constraints

**Video Detection (Vendor API):**
- **Primary:** Sensity AI or Reality Defender API
  - Accuracy: >85% on FaceForensics++ benchmark
  - Response Time: <30 seconds per video (up to 2 minutes)
  - Features: Face boundary artifacts, lighting inconsistency, blinking patterns, lip-sync analysis

**Out-of-Band Verification Workflow:**
1. Employee receives suspicious voice/video request (e.g., "CEO" requesting wire transfer)
2. Employee flags "verify" in SMESec app
3. System sends verification SMS + email to claimed sender
4. Sender clicks one-click link to confirm or deny
5. Employee receives result within <5 minutes
6. If denied: Deepfake incident report auto-generated, IT admin alerted

**Liveness Detection Challenge:**
- System requests user to speak random phrase (e.g., "The quick brown fox jumps over the lazy dog at 3:47 PM")
- Prevents replay attacks and pre-recorded deepfakes

### 4.4 Policy Engine

**Policy Framework:** OPA (Open Policy Agent) with Rego policy language for declarative, auditable policy management

**Architecture:**
- Policies stored as code (version-controlled, rollback-able)
- Policy evaluation <50ms (in-memory, compiled Rego)
- Changes take effect within <1 minute (policy cache refresh)
- IT admin defines policies via UI (generates Rego behind the scenes)

#### Policy Types

1. **Data Protection Policies**
   - **Block PII:** Never allow credit cards, SSN, or passwords to any LLM (hard block, no override)
   - **Redact PII:** Allow names, emails, phone numbers but mask them before sending to LLM
   - **IP Protection:** Block source code snippets, confidential document fingerprints
   - **Conditional Blocking:** Block PII to unknown AI tools, allow to approved tools (e.g., ChatGPT Enterprise with BAA)

2. **Access Control Policies**
   - **Role-Based:** Admins can bypass DLP warnings, employees cannot
   - **Data Sensitivity:** Users with "Restricted" data access face stricter DLP rules
   - **Tool Allow-List:** Only approved AI tools allowed (e.g., ChatGPT, Copilot); block all others
   - **Justification Required:** Medium-risk prompts require user to provide business justification

3. **Usage Policies**
   - **Rate Limiting:** Max 100 prompts/day per user to prevent abuse
   - **Time-Based:** Block AI tool usage outside business hours (9am-6pm) for contractors
   - **Audit Trail:** All blocked prompts logged with user, prompt hash, reason, timestamp
   - **Escalation:** 3+ blocks in 1 hour triggers IT admin alert

**Policy Examples (Rego):**
```rego
# Block credit cards to all LLMs
deny[msg] {
  input.dlp_findings[_].type == "credit_card"
  msg := "Credit card detected - blocked by policy"
}

# Allow admins to bypass DLP warnings
allow {
  input.user.role == "admin"
  input.risk_score < 85
}

# Require justification for medium-risk prompts
require_justification {
  input.risk_score >= 31
  input.risk_score <= 60
  not input.justification
}
```

### 4.5 Incident Response Integration

**Event-Driven Architecture:**

```
┌─────────────────────────────────────────────────────────────┐
│              Track 2: AI Threat Detection                    │
│                                                              │
│  [Prompt Injection] [Data Leakage] [Deepfake] [Shadow AI]  │
│         │                 │              │           │       │
│         └─────────────────┴──────────────┴───────────┘       │
│                           │                                  │
│                  ┌────────▼────────┐                        │
│                  │ ThreatDetection │                        │
│                  │ Event Publisher │                        │
│                  └────────┬────────┘                        │
└───────────────────────────┼──────────────────────────────────┘
                            │
                   ┌────────▼────────┐
                   │  EventBridge    │
                   │  (Event Router) │
                   └────────┬────────┘
                            │
        ┌───────────────────┼───────────────────┐
        │                   │                   │
   ┌────▼─────┐      ┌─────▼──────┐     ┌─────▼──────┐
   │ Playbook │      │   Slack    │     │  Evidence  │
   │  Engine  │      │   Alert    │     │ Collection │
   └────┬─────┘      └────────────┘     └────────────┘
        │
   ┌────▼─────┐
   │ Access   │
   │Revocation│
   └──────────┘
```

**ThreatDetectionEvent Schema (defined Sprint 1, frozen Sprint 2):**
```json
{
  "version": "1.0",
  "event_id": "uuid",
  "tenant_id": "uuid",
  "timestamp": "ISO8601",
  "threat_type": "prompt_injection | data_leakage | deepfake | shadow_ai",
  "severity": "low | medium | high | critical",
  "risk_score": 0-100,
  "user": {
    "user_id": "uuid",
    "email": "string",
    "role": "admin | manager | employee | contractor"
  },
  "context": {
    "ai_tool": "chatgpt | copilot | gemini | unknown",
    "prompt_hash": "sha256",
    "detection_layers": ["regex", "ml", "context"],
    "evidence_s3_url": "s3://..."
  },
  "recommended_action": "log | alert | block | revoke_access"
}
```

#### Response Workflows
| Threat Type | Severity | Automated Action | Playbook Triggered | Response Time |
|-------------|----------|------------------|-------------------|---------------|
| **Prompt Injection** | Low (0-30) | Log only | None | N/A |
| **Prompt Injection** | Medium (31-60) | Advisory alert + justification request | None | <1 min |
| **Prompt Injection** | High (61-85) | Block + manager approval required | Shadow IT Detected (if unapproved tool) | <2 min |
| **Prompt Injection** | Critical (86-100) | Block + IT admin alert + access review | Account Compromise | <1 min |
| **Data Leakage** | Medium (31-60) | Redact + advisory alert | None | <1 min |
| **Data Leakage** | High (61-85) | Block + manager approval | None | <2 min |
| **Data Leakage** | Critical (86-100) | Block + IT admin alert + incident report | Account Compromise | <1 min |
| **Deepfake** | High (>70% confidence) | Alert employee + manager + IT admin | Unauthorized Access | <1 min |
| **Deepfake** | Critical (>90% confidence) | Block transaction + out-of-band verification | Account Compromise | <1 min |
| **Shadow AI** | Medium (unapproved tool) | Alert IT admin + add to review queue | Shadow IT Detected | <15 min |
| **Shadow AI** | High (unapproved + PII) | Block + IT admin alert + access review | Shadow IT Detected | <1 min |

**Integration Points (Sprint 10):**
- Track 2 publishes events to EventBridge topic `smesec-threat-detection`
- Track 1 Playbook Engine subscribes to topic, filters by `severity` and `threat_type`
- Evidence artifacts (prompt hash, detection metadata) stored in S3, linked in event
- All events logged to CloudWatch for audit trail
- Slack/email notifications sent in parallel with playbook execution

### 4.6 Privacy & Compliance

**Privacy-by-Design:**

1. **Minimal Monitoring Scope**: Browser extension ONLY monitors AI tool domains (whitelist: chatgpt.com, copilot.microsoft.com, gemini.google.com, kiro.ai). Does NOT monitor social media, banking, news, personal browsing, or any non-AI websites.

2. **Transparent Data Collection**: Extension displays active monitoring status in browser toolbar. Privacy policy shown on install explains exactly what data is collected (prompt text, AI tool domain, timestamp, user ID) and why (security threat detection).

3. **User Control & Consent**: 
   - User can disable extension at any time (IT admin receives notification but cannot force re-enable)
   - User can view all collected data via dashboard
   - User can request data deletion (honored within 30 days per GDPR Article 17)

4. **Data Minimization**: Prompts analyzed in real-time, only metadata stored long-term (prompt hash, risk score, timestamp). Full prompt text stored only for High/Critical risk events (evidence for incident investigation).

5. **Purpose Limitation**: Data used solely for security threat detection. Never used for employee surveillance, performance monitoring, or productivity tracking.

#### Browser Extension Privacy Controls

- **Monitoring Scope:** AI tool domains only (whitelist-based). Explicitly excludes: social media, banking, healthcare, news, personal email, shopping, entertainment sites.

- **Data Collection:** 
  - **Collected:** Prompt text (analyzed in real-time), AI tool domain, timestamp, user ID, risk score
  - **NOT Collected:** Browsing history, non-AI website content, keystrokes outside AI tools, screenshots, clipboard data (except when explicitly pasted into AI tool)

- **Data Retention:** 
  - Low/Medium risk: Metadata only (prompt hash, risk score) for 90 days
  - High/Critical risk: Full prompt text for 7 years (compliance evidence)
  - User can configure retention period (30-365 days) for their tenant

- **User Control:** 
  - Disable extension: User can turn off monitoring (IT admin notified but cannot override)
  - Opt-out: User can request permanent opt-out (requires manager approval)
  - Data access: User can view all their collected data via dashboard
  - Data deletion: User can request deletion of non-compliance data (processed within 30 days)

#### Compliance Alignment

**GDPR Article 13/88 (Browser Extension Disclosure):**
- Privacy policy displayed on extension install (before any data collection)
- Explains: what data collected, why, how long retained, user rights (access, deletion, opt-out)
- Legal basis: Legitimate interest (employer security) + user consent (can disable)
- Data controller: Employer (SME customer), not SMESec
- DPO contact information provided in policy

**GDPR Article 32 (Security Measures):**
- Encryption: All prompt data encrypted in transit (TLS 1.3) and at rest (AWS KMS)
- Access control: Only tenant admins can view collected data, row-level security enforced
- Pseudonymization: Prompt hashes used for low-risk events (not full text)
- Audit trail: All data access logged with user ID, timestamp, purpose

**ISO 27001 A.12.4 (Logging & Monitoring):**
- All AI threat events logged to immutable audit trail (S3 Object Lock)
- Logs include: user, timestamp, threat type, risk score, action taken
- Log retention: 7 years (compliance requirement)
- Log integrity: SHA-256 hashes prevent tampering

**Legal Review (Sprint 4-5):**
- PM submits browser extension privacy policy to legal team for GDPR compliance review
- Legal review covers: data collection scope, user consent mechanism, retention periods, cross-border data transfer (if applicable)
- Review must complete before Sprint 5 (browser extension deployment)

---

## 5. Risk Assessment & Mitigation

### 5.1 Technical Risks

| Risk | Probability | Impact | Mitigation | Owner |
|------|-------------|--------|------------|-------|
| **AI detection accuracy insufficient (<95%)** | Medium | **CRITICAL** | Validation gates at W8, W12, W18, W24; early dataset quality checks; fallback to Track 1-only launch | ML Engineer + PM |
| **Cross-tenant data leakage** | Low | **CRITICAL** | Automated CI tests (Sprint 1); RLS policies enforced at DB layer; penetration test (Sprint 13); runtime monitoring with CloudWatch alarms | Tech Lead + DevSecOps |
| **Chrome MV3 service worker persistence fails** | Low | High | HARD GATE Week 1 Sprint 1; if fails, cut browser extension from v1 (ship API-only detection) | Extension Eng |
| **Integration API changes break sync** | Medium | Medium | Daily integration tests; version pinning; fallback to manual sync if API down | Backend Eng |
| **Performance degradation at scale** | Medium | Medium | Load testing Sprint 13 (500 assets, 50 users); caching layer; rate limiting; auto-scaling | Tech Lead + DevSecOps |
| **Vendor API (deepfake) unavailable or too expensive** | Low | High | DeepfakeDetector abstraction interface; Resemblyzer open-source fallback (~78-82% accuracy); negotiate pricing early | PM + ML Engineer |
| **Track 1-Track 2 schema incompatibility** | Low | Medium | Joint schema session Week 1; schema frozen Week 4; validation meeting Week 15-16 | Tech Lead (both tracks) |
| **SageMaker cold-start latency >10s** | Medium | Medium | Provisioned concurrency for ML endpoints; load test Sprint 3; optimize model size | ML Engineer + DevSecOps |

### 5.2 Project Risks

| Risk | Probability | Impact | Mitigation | Owner |
|------|-------------|--------|------------|-------|
| **Pilot customers unavailable for Track 2** | Medium | High | Outreach starts Week 1; 2+ signed NDAs/DPAs by Week 13 (HARD DEADLINE); have 3 customers (not 2) to absorb dropout | PM |
| **Team velocity <80% of plan** | High | Medium | Measure velocity Sprint 1-2; adjust scope Sprint 3 if gap detected; reduce sprint scope by 20% if needed | PM + Tech Lead |
| **Key person unavailable (sick, attrition)** | Medium | Medium | Cross-training: Flutter Eng can assist Extension Eng; Backend Eng 2 can assist Backend Eng 1; document critical knowledge | PM + Tech Lead |
| **Tech Lead over-allocated (90-100% utilization)** | High | Medium | Delegate code reviews to senior Backend Eng; focus Tech Lead on architecture decisions only; reduce Sprint 1 scope | PM + Tech Lead |
| **Pen-test vendor booking delayed** | Medium | High | PM selects vendor by Sprint 7 (W14); signs LOI by Sprint 8 (W16); 12-week lead time from request to execution | PM |
| **Legal review delays browser extension** | Low | Medium | Submit privacy policy for legal review Sprint 4; 4-week buffer before Sprint 5 deployment | PM |
| **Compliance audit failure (ISO 27001, SOC 2)** | Low | High | External audit review at Month 5 (Sprint 11); address findings before launch; v1 is compliance-ready, certification post-launch | PM + Tech Lead |

### 5.3 Market Risks

| Risk | Probability | Impact | Mitigation | Owner |
|------|-------------|--------|------------|-------|
| **Competitor launches similar AI security platform** | Medium | Medium | Track 1 provides immediate value (asset + access governance); Track 2 differentiation (3-layer detection, context-aware); focus on SME market (10-500 employees) underserved by enterprise solutions | PM |
| **SMEs unwilling to pay for security platform** | Low | High | Tiered pricing aligned with SME budgets; free tier for <10 employees; ROI calculator (cost of breach vs platform cost); compliance-ready messaging (ISO 27001, GDPR, SOC 2) | PM |
| **AI threat landscape evolves faster than detection models** | High | Medium | Continuous model retraining with new attack patterns; community threat intelligence feeds; quarterly model updates post-launch | ML Engineer |
| **Privacy concerns around browser extension monitoring** | Medium | Medium | Privacy-by-design: AI tool domains only, user opt-out, transparent data handling; legal review (GDPR compliance); clear privacy policy on install | PM + Extension Eng |
| **Vendor API pricing increases post-launch** | Low | Medium | Negotiate multi-year contract with price lock; maintain open-source fallback (Resemblyzer); evaluate alternative vendors annually | PM |

---

## 6. Success Criteria

### 6.1 Track 1 Launch Criteria (Beta Launch - Month 6)
- ✅ **Asset discovery coverage >95%**: All devices, accounts, SaaS apps, cloud resources discovered within 1 hour for 500-asset organization
- ✅ **Offboarding completion <5 minutes**: All access revoked across Google/M365/Slack/AWS in parallel via Step Functions
- ✅ **Shadow IT detection rate >90%**: OAuth apps and unapproved SaaS tools detected and alerted
- ✅ **Zero cross-tenant data leakage**: Verified by automated CI tests (Sprint 1) and external penetration test (Sprint 13)
- ✅ **Compliance reports generated**: ISO 27001, GDPR, SOC 2 reports (PDF + JSON) in <5 minutes
- ✅ **5-10 pilot customers onboarded**: Successfully using Track 1 features end-to-end
- ✅ **Penetration test passed**: 0 Critical findings, 0 High findings (all remediated before launch)
- ✅ **Performance targets met**: Dashboard load <2s, search <1s, offboarding <5 min, asset discovery <5 min
- ✅ **5 playbooks operational**: Account Compromise, Offboarding Emergency, Shadow IT, Unauthorized Access, Inactive Account — executable by non-security staff in <10 minutes

### 6.2 Track 2 Pilot Criteria (Pilot-Ready - Month 6)
- ✅ **Prompt injection precision >95%**: False positive rate <5% on test dataset
- ✅ **DLP precision >99%**: On critical data (credit cards, SSN, passwords), false negative rate <1%
- ✅ **Deepfake detection accuracy**: Voice >90%, video >85% on FaceForensics++ benchmark
- ✅ **False positive rate <5%**: On pilot customer production data (real-world validation)
- ✅ **2-3 pilot customers validated**: 4-week pilot completed with telemetry collection
- ✅ **Customer satisfaction >4.0/5.0**: NPS or CSAT score from pilot customers
- ✅ **0 critical incidents**: No false negatives causing actual data breaches or financial loss
- ✅ **User frustration <10%**: False positive complaints from pilot users below threshold
- ✅ **Performance targets met**: End-to-end detection <1s (p95), ML inference <500ms, browser extension overhead <50ms

**Launch Decision (Gate 4 - Week 24):**
- If ALL Track 2 criteria met → Merge into main product, full launch with Track 1
- If criteria NOT met → Track 2 continues as beta, Track 1 launches independently

### 6.3 Business Metrics (Post-Launch - 6-12 Months)
- **Customer Acquisition**: 50+ paying customers within 6 months of launch
- **Revenue**: $500K ARR within 12 months (assuming $10K average contract value × 50 customers)
- **Churn Rate**: <10% annual churn (SME market benchmark: 15-20%)
- **NPS Score**: >40 (promoters - detractors, industry benchmark for B2B SaaS: 30-40)
- **Time to Value**: <7 days from signup to first asset discovered and first policy enforced
- **Compliance Certification**: SOC 2 Type II audit completed within 12 months of launch (requires 6+ months operational history)
- **ISO 27001 Certification**: Audit completed within 12 months of launch
- **Support Ticket Volume**: <5 tickets per customer per month (indicates product usability)
- **Feature Adoption**: >80% of customers using offboarding automation, >60% using compliance reports, >40% using AI detection (if Track 2 launched)

---

## 7. Next Steps

### Immediate Actions (Week 1 - Sprint 1)

1. ✅ **Approve 2-track strategy** (this document) — Stakeholder sign-off on parallel Track 1 + Track 2 development
2. ⏳ **Team hiring/allocation** — Confirm 10 FTE availability (5 Track 1, 3 Track 2, 2 shared); backfill if needed
3. ⏳ **Chrome MV3 persistence prototype** (HARD GATE) — Extension Eng validates service worker persistence within Week 1; if fails, cut browser extension from v1
4. ⏳ **Joint T1-T2 schema session** — Define ThreatDetectionEvent interface, freeze by end of Sprint 2 (Week 4)
5. ⏳ **Pilot customer outreach** — PM begins identifying 2-3 SMEs (50-200 employees, using ChatGPT/Copilot) for Track 2 pilot
6. ⏳ **Vendor API trial requests** — PM requests trial accounts from Sensity AI + Reality Defender for deepfake detection evaluation
7. ⏳ **AWS infrastructure setup** — DevSecOps provisions VPC, ECS, RDS, S3, Secrets Manager, CloudWatch for Sprint 1 development

### Sprint 1-2 Deliverables (Week 1-4)

8. ⏳ **Track 1 Sprint 1** — AWS infra, Keycloak SSO, multi-tenant DB schema, CI/CD pipeline, tenant isolation CI test
9. ⏳ **Track 1 Sprint 2** — Google Workspace sync (users + OAuth apps), background job (15-min incremental sync)
10. ⏳ **Track 2 Sprint 1** — Dataset collection (5K+ labeled examples), ML infra (SageMaker), DeepfakeDetector abstraction interface
11. ⏳ **Track 2 Sprint 2** — Regex patterns (50+ OWASP LLM), PII rules, detection API scaffold

### Key Milestones & Gates

12. ⏳ **Week 4 (End Sprint 2)** — ThreatDetectionEvent schema frozen (no breaking changes after this point)
13. ⏳ **Week 8 (Gate 1)** — Track 2 validation: Prompt injection >90%, DLP >95% critical data
14. ⏳ **Week 12 (Gate 2)** — Track 2 validation: Injection FP <10%, DLP FP <5%
15. ⏳ **Week 13 (Sprint 7)** — 2+ pilot customer NDAs/DPAs signed (HARD DEADLINE for Track 2)
16. ⏳ **Week 16 (Sprint 8)** — Pen-test vendor LOI signed (12-week lead time to Sprint 12 execution)
17. ⏳ **Week 18 (Gate 3)** — Track 2 validation: All targets met (Injection >95%, DLP >99%, Deepfake >90%/85%)
18. ⏳ **Week 24 (Gate 4)** — Track 2 pilot complete, launch decision (merge to product OR continue beta)
19. ⏳ **Week 26 (Sprint 13)** — Track 1 beta launch with 5-10 pilot customers

### Post-Launch (Month 7-12)

20. ⏳ **Customer onboarding** — Scale from 5-10 pilot customers to 50+ paying customers
21. ⏳ **SOC 2 Type II audit** — Requires 6+ months operational history, target completion within 12 months
22. ⏳ **ISO 27001 certification** — External audit, target completion within 12 months
23. ⏳ **Track 2 iteration** (if beta) — Continue accuracy tuning based on production telemetry until >95% precision achieved
24. ⏳ **v1.1 features** — QuickBooks integration, MDM integration (Intune, Workspace ONE), custom playbook builder UI, multi-language support

---

## Appendices

### Appendix A: Technology Stack

#### Backend Services
| Component | Technology | Rationale |
|-----------|-----------|-----------|
| **API Gateway** | AWS API Gateway | Managed service, built-in rate limiting, request/response transformation |
| **Application Runtime** | ECS Fargate | Serverless containers, auto-scaling, no EC2 management overhead |
| **Track 1 API** | Go 1.21+ | Performance, concurrency (goroutines), strong typing, excellent AWS SDK |
| **Track 2 API** | Python 3.11+ FastAPI | ML ecosystem compatibility, async support, automatic OpenAPI docs |
| **Database** | PostgreSQL 15 (RDS Multi-AZ) | ACID compliance, row-level security, JSON support, proven at scale |
| **Cache** | Redis 7 (ElastiCache) | In-memory performance for redaction mappings, session data |
| **Object Storage** | AWS S3 | Durable evidence storage, Object Lock for compliance, lifecycle policies |
| **Secrets Management** | AWS Secrets Manager | Automatic rotation, KMS integration, audit trail |
| **Encryption** | AWS KMS | Tenant-scoped encryption context, FIPS 140-2 Level 3 validated |

#### ML & AI Infrastructure
| Component | Technology | Rationale |
|-----------|-----------|-----------|
| **ML Training** | AWS SageMaker | Managed Jupyter notebooks, distributed training, experiment tracking |
| **ML Inference** | SageMaker Endpoints | Auto-scaling, A/B testing, model monitoring, provisioned concurrency |
| **ML Framework** | PyTorch 2.0 + Transformers | BERT fine-tuning, active community, production-ready |
| **NER Model** | spaCy 3.7 | Fast inference, custom entity training, production-proven |
| **Experiment Tracking** | MLflow | Model registry, versioning, reproducibility |
| **Deepfake Detection** | Sensity AI / Reality Defender API | Specialized vendor, >90% accuracy, faster than building in-house |
| **Fallback Deepfake** | Resemblyzer (open-source) | Voice embedding similarity, ~78-82% accuracy, no vendor dependency |

#### Frontend & Client
| Component | Technology | Rationale |
|-----------|-----------|-----------|
| **Web Framework** | Next.js 14 (React 18) | SSR/SSG, API routes, file-based routing, excellent DX |
| **UI Library** | Tailwind CSS + shadcn/ui | Utility-first CSS, accessible components, rapid prototyping |
| **State Management** | React Query + Zustand | Server state (React Query), client state (Zustand), minimal boilerplate |
| **Charts & Viz** | Recharts + D3.js | Dependency graphs, risk scoring visualization, compliance dashboards |
| **Mobile/Desktop** | Flutter 3.16+ | Single codebase for iOS/Android/Desktop, native performance |
| **Browser Extension** | Chrome MV3 (TypeScript) | Latest manifest version, service worker architecture, cross-browser (Chrome/Edge) |

#### Orchestration & Events
| Component | Technology | Rationale |
|-----------|-----------|-----------|
| **Workflow Engine** | AWS Step Functions | Visual workflow editor, fault-tolerant, retry logic, state persistence |
| **Event Bus** | AWS EventBridge | Event-driven architecture, schema registry, cross-account routing |
| **Message Queue** | AWS SQS | Decoupling, at-least-once delivery, dead-letter queues |
| **Scheduler** | EventBridge Scheduler | Cron-based scheduling, one-time events, timezone support |

#### Auth & Security
| Component | Technology | Rationale |
|-----------|-----------|-----------|
| **SSO/MFA** | Keycloak 23+ | Open-source, SAML/OAuth 2.0/OIDC, self-hosted, Google/M365 federation |
| **Policy Engine** | Open Policy Agent (OPA) | Declarative policies (Rego), <50ms evaluation, version-controlled |
| **WAF** | AWS WAF | DDoS protection, SQL injection prevention, rate limiting |
| **Certificate Management** | AWS Certificate Manager | Free SSL/TLS certificates, auto-renewal |

#### Monitoring & Observability
| Component | Technology | Rationale |
|-----------|-----------|-----------|
| **Logs** | CloudWatch Logs | Centralized logging, log insights queries, retention policies |
| **Metrics** | CloudWatch Metrics | Custom metrics, dashboards, alarms |
| **Tracing** | AWS X-Ray | Distributed tracing, service map, latency analysis |
| **Alerting** | CloudWatch Alarms + SNS | Threshold-based alerts, Slack/email integration |
| **Uptime Monitoring** | AWS CloudWatch Synthetics | Canary tests, API endpoint monitoring |

#### CI/CD & Infrastructure
| Component | Technology | Rationale |
|-----------|-----------|-----------|
| **CI/CD** | GitHub Actions | Native GitHub integration, matrix builds, secrets management |
| **Infrastructure as Code** | Terraform 1.6+ | Declarative, state management, multi-cloud support (future) |
| **Container Registry** | AWS ECR | Private Docker registry, vulnerability scanning, lifecycle policies |
| **Deployment** | AWS CodeDeploy | Blue/green deployments, automatic rollback, traffic shifting |

### Appendix B: Compliance Mapping

#### ISO 27001:2022 Controls

| Control | Title | SMESec Implementation | Evidence |
|---------|-------|----------------------|----------|
| **A.8.1** | Inventory of assets | Automated asset discovery (Google/M365/Slack/AWS), classification engine, dependency mapping | Daily asset snapshots (S3), asset inventory reports |
| **A.8.2** | Information classification | Auto-classification (Restricted/Confidential/Internal/Public), manual override, bulk CSV import | Classification history logs, audit trail |
| **A.9.1** | Access control policy | RBAC engine (OPA/Rego), least-privilege enforcement, policy versioning | Policy documents, OPA Rego files |
| **A.9.2** | User access management | Automated provisioning/deprovisioning, JIT access with auto-expiration, offboarding <5 min | Offboarding reports (PDF), access event logs |
| **A.9.4** | Access review | Quarterly access review workflow, manager approval, access certification reports | Access review reports, approval audit trail |
| **A.12.4** | Logging and monitoring | All access events logged (immutable S3 Object Lock), 7-year retention, CloudWatch monitoring | Audit logs (S3), CloudWatch dashboards |
| **A.16.1** | Incident management | 5 pre-built playbooks, wizard UI for non-security staff, automated response workflows | Incident reports (PDF), playbook execution logs |
| **A.18.1** | Compliance with legal requirements | GDPR, SOC 2, ISO 27001 control mapping, evidence auto-collection, compliance dashboard | Compliance reports (PDF+JSON), evidence artifacts |

#### GDPR Articles

| Article | Requirement | SMESec Implementation | Evidence |
|---------|-------------|----------------------|----------|
| **Art. 30** | Records of processing activities | Asset inventory includes data processing activities, data flow mapping | Asset inventory reports, data flow diagrams |
| **Art. 32** | Security of processing | Encryption (KMS at-rest, TLS 1.3 in-transit), access controls (RBAC), MFA enforcement, audit logs | Encryption policies, access control policies, MFA enforcement logs |
| **Art. 17** | Right to erasure | Offboarding workflow includes data deletion, manual deletion requests honored within 30 days | Data deletion logs, offboarding reports |
| **Art. 33** | Breach notification | Incident playbooks include 72-hour breach notification workflow, automated alert escalation | Incident reports, notification logs |
| **Art. 25** | Data protection by design | Privacy-by-design principles, data minimization, purpose limitation, encryption everywhere | Architecture documentation, privacy policy |
| **Art. 13/88** | Browser extension disclosure | Privacy policy displayed on install, explains data collection scope, user rights (access/deletion/opt-out) | Extension privacy policy, consent logs |

#### SOC 2 Trust Services Criteria

| Criteria | Title | SMESec Implementation | Evidence |
|----------|-------|----------------------|----------|
| **CC6.1** | Logical and physical access controls | RBAC restricts access based on role, MFA enforced, least-privilege principle | Access control policies, MFA logs, role definitions |
| **CC6.2** | Prior to issuing system credentials | Automated provisioning via JIT access, manager approval required, access auto-expires | JIT access logs, approval workflows |
| **CC6.3** | Removes access when no longer required | Automated offboarding <5 min, JIT access auto-revocation, quarterly access reviews | Offboarding reports, access revocation logs |
| **CC7.2** | System monitoring | CloudWatch monitoring, real-time alerts for anomalies, security event correlation | CloudWatch dashboards, alert logs, incident reports |
| **CC7.3** | Evaluates security events | All access events logged (immutable), 7-year retention, audit trail for investigations | Audit logs (S3 Object Lock), investigation reports |
| **CC8.1** | Authorizes, designs, develops changes | CI/CD pipeline with automated tests, code review required, staging environment | GitHub Actions logs, PR approval history, deployment logs |

### Appendix C: Integration Specifications

#### ThreatDetectionEvent Schema (Track 1-Track 2 Integration)

**EventBridge Topic:** `smesec-threat-detection`

**Schema Version:** 1.0 (defined Sprint 1, frozen Sprint 2)

```json
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "type": "object",
  "required": ["version", "event_id", "tenant_id", "timestamp", "threat_type", "severity", "risk_score", "user"],
  "properties": {
    "version": {
      "type": "string",
      "const": "1.0"
    },
    "event_id": {
      "type": "string",
      "format": "uuid"
    },
    "tenant_id": {
      "type": "string",
      "format": "uuid"
    },
    "timestamp": {
      "type": "string",
      "format": "date-time"
    },
    "threat_type": {
      "type": "string",
      "enum": ["prompt_injection", "data_leakage", "deepfake", "shadow_ai"]
    },
    "severity": {
      "type": "string",
      "enum": ["low", "medium", "high", "critical"]
    },
    "risk_score": {
      "type": "integer",
      "minimum": 0,
      "maximum": 100
    },
    "user": {
      "type": "object",
      "required": ["user_id", "email", "role"],
      "properties": {
        "user_id": {"type": "string", "format": "uuid"},
        "email": {"type": "string", "format": "email"},
        "role": {"type": "string", "enum": ["admin", "manager", "employee", "contractor"]}
      }
    },
    "context": {
      "type": "object",
      "properties": {
        "ai_tool": {"type": "string"},
        "prompt_hash": {"type": "string"},
        "detection_layers": {"type": "array", "items": {"type": "string"}},
        "evidence_s3_url": {"type": "string", "format": "uri"}
      }
    },
    "recommended_action": {
      "type": "string",
      "enum": ["log", "alert", "block", "revoke_access"]
    }
  }
}
```

#### Google Workspace Integration

**API:** Google Admin SDK + Audit API  
**Auth:** OAuth 2.0 with service account  
**Scopes Required:**
- `https://www.googleapis.com/auth/admin.directory.user.readonly` (user discovery)
- `https://www.googleapis.com/auth/admin.directory.group.readonly` (group discovery)
- `https://www.googleapis.com/auth/admin.reports.audit.readonly` (audit logs)
- `https://www.googleapis.com/auth/admin.directory.user` (offboarding - write access)

**Sync Frequency:** Every 15 minutes (incremental)  
**Rate Limits:** 1,500 requests/minute per project

#### Microsoft 365 Integration

**API:** Microsoft Graph API + Azure AD  
**Auth:** OAuth 2.0 with app registration  
**Permissions Required:**
- `User.Read.All` (user discovery)
- `Group.Read.All` (group discovery)
- `AuditLog.Read.All` (audit logs)
- `User.ReadWrite.All` (offboarding - write access)

**Sync Frequency:** Every 15 minutes (incremental)  
**Rate Limits:** 10,000 requests/10 minutes per app

### Appendix D: Glossary

| Term | Definition |
|------|------------|
| **Asset** | Any resource owned or managed by the organization: devices, accounts, SaaS apps, cloud resources, data stores |
| **BERT** | Bidirectional Encoder Representations from Transformers - ML model architecture used for prompt injection detection |
| **Classification** | Assigning sensitivity levels (Restricted/Confidential/Internal/Public) to assets based on data criticality |
| **Deepfake** | AI-generated synthetic media (voice or video) that impersonates a real person |
| **DLP** | Data Loss Prevention - detecting and blocking sensitive data (PII, credentials, IP) from leaving the organization |
| **EventBridge** | AWS service for event-driven architecture, routing events between Track 1 and Track 2 |
| **JIT Access** | Just-In-Time access - temporary elevated permissions that auto-expire after a specified duration |
| **KMS** | AWS Key Management Service - encryption key management with tenant-scoped encryption context |
| **MFA** | Multi-Factor Authentication - requiring 2+ verification factors (password + TOTP) for login |
| **NER** | Named Entity Recognition - ML technique for identifying PII (names, emails, phone numbers) in text |
| **OAuth 2.0** | Industry-standard protocol for authorization, used for integrating with Google/M365/Slack |
| **OPA** | Open Policy Agent - policy engine using Rego language for declarative access control |
| **Playbook** | Automated incident response workflow with step-by-step guidance for non-security staff |
| **Prompt Injection** | Attack technique where malicious instructions are embedded in prompts to manipulate LLM behavior |
| **RBAC** | Role-Based Access Control - restricting access based on user roles (Admin/Manager/Employee/Contractor) |
| **Redaction** | Replacing sensitive data with tokens before sending to LLM, then restoring in response (de-redaction) |
| **RLS** | Row-Level Security - PostgreSQL feature enforcing tenant_id filtering at database layer |
| **SageMaker** | AWS managed ML service for training and deploying models at scale |
| **Shadow AI** | Unapproved AI tools (ChatGPT, Copilot, etc.) used by employees without IT approval |
| **Shadow IT** | Unapproved SaaS applications or cloud services used without IT knowledge or approval |
| **SME** | Small and Medium Enterprise - organizations with 10-500 employees (target market) |
| **SOC 2** | Service Organization Control 2 - audit framework for security, availability, confidentiality |
| **SSO** | Single Sign-On - centralized authentication using Google or Microsoft identity providers |
| **Step Functions** | AWS service for orchestrating workflows (playbooks) with visual editor and fault tolerance |
| **Tenant** | A single customer organization in the multi-tenant SMESec platform |
| **Track 1** | Foundation & Governance development track - deterministic, high-confidence features |
| **Track 2** | AI Threat Detection development track - ML-based, requires validation gates |
| **Validation Gate** | Quality checkpoint with specific accuracy metrics that must pass before proceeding |

---

## Document History

| Version | Date | Author | Changes |
|---------|------|--------|---------|
| 1.0 | 2026-05-28 | SMESec Platform Team | Initial comprehensive design document synthesizing Track 1 (Foundation & Governance) and Track 2 (AI Threat Detection) requirements, 2-track development strategy, team structure, 6-month delivery plan, risk assessment, and compliance mapping |
