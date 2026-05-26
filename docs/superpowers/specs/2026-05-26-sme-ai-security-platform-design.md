# SME AI Security Platform Design (Hybrid Model)

Date: 2026-05-26  
Scope: Strategic proposal for a unified SME protection platform (10–500 employees)

## 1) System Architecture Diagram

### 1.1 Logical View

```mermaid
flowchart LR
    U[Admin / Ops / Non-security Staff] --> UI[Web Dashboard\nReact/Next.js]
    U --> FD[Flutter Mobile/Desktop App]

    UI --> APIGW[API Gateway]
    FD --> APIGW

    APIGW --> INV[Asset Inventory Service\nGo]
    APIGW --> POL[Policy Orchestration Service\nGo + OPA]
    APIGW --> AIT[AI Threat Detection Service\nPython]
    APIGW --> CMP[Compliance Engine\nPython]
    APIGW --> IR[Incident Playbook Service\nGo]
    APIGW --> INT[Integration Hub\nPlugin-based]

    INT --> GW[Google Workspace]
    INT --> M365[Microsoft 365]
    INT --> SLK[Slack]
    INT --> QB[QuickBooks]
    INT --> AWSI[AWS Accounts]

    AIT --> DFK[Deepfake Detection Vendor API]

    INV --> DB[(Aurora PostgreSQL)]
    POL --> DB
    CMP --> DB
    IR --> DB

    AIT --> SQS[SQS/EventBridge]
    POL --> SQS
    IR --> SQS

    CMP --> S3[(S3 Evidence Store)]
    AIT --> CW[CloudWatch]
    IR --> SNS[SNS/Email/Slack Alerts]
```

### 1.2 Deployment View (AWS-first)

```mermaid
flowchart TB
    subgraph AWS_VPC[VPC]
      subgraph Edge[Edge]
        CF[CloudFront] --> WAF[AWS WAF]
        WAF --> APIGW2[API Gateway / ALB]
      end

      subgraph App[ECS Fargate Cluster]
        S1[Asset Inventory]
        S2[Policy + OPA]
        S3A[AI Threat Detection]
        S4[Compliance Engine]
        S5[Incident Playbook]
        S6[Integration Hub]
      end

      subgraph Data[Data Services]
        RDS[(Aurora PostgreSQL)]
        RED[ElastiCache Redis]
        S3B[(S3)]
      end

      subgraph Async[Eventing]
        EB[EventBridge]
        Q[SQS]
        SF[Step Functions]
        L[Lambda Workers]
      end

      subgraph ML[ML Services]
        SM[SageMaker Endpoint/Registry]
      end

      APIGW2 --> S1
      APIGW2 --> S2
      APIGW2 --> S3A
      APIGW2 --> S4
      APIGW2 --> S5
      APIGW2 --> S6

      S1 --> RDS
      S2 --> RDS
      S3A --> RDS
      S4 --> RDS
      S5 --> RDS
      S3A --> SM

      S2 --> EB
      S3A --> Q
      S5 --> SF
      SF --> L

      S4 --> S3B
      L --> S3B
    end

    IdP[Keycloak / SSO] --> APIGW2
```

## 2) Design Document (<=600 words)

This platform uses a hybrid architecture: build core differentiators and integrate commodity capabilities. The goal is to protect SME assets against AI-era threats without requiring a full-time security team.

Build vs buy is the primary decision. The system builds five core modules in-house: AI Threat Detection, Asset Inventory, Policy Orchestration, Compliance Engine, and Incident Playbook Service. Commodity functions are integrated instead of rebuilt: identity and SSO (Keycloak/IdP), SaaS connectors (Google Workspace, Microsoft 365, Slack, QuickBooks), and managed deepfake APIs. This reduces delivery risk in 6 months while preserving unique IP where it matters.

The multi-tenancy model is tenant-aware by default, with logical isolation in shared infrastructure using tenant-scoped access controls and row-level data isolation in Aurora PostgreSQL. For customers requiring stronger separation, the same services can be deployed in dedicated environments (private cloud/on-prem pattern). This supports a tiered commercial model without maintaining separate codebases.

AI threat detection strategy is phased. V1 prioritizes practical controls: prompt injection detection (rule + lightweight classifier), LLM data leakage controls (DLP patterns at browser/desktop endpoint), shadow AI discovery (domain/API/OAuth app detection), and deepfake fraud mitigation via vendor APIs plus out-of-band callback verification workflow. This approach targets high-risk scenarios while controlling false positives and implementation complexity. Model-heavy custom detection is deferred to later versions after telemetry and pilot feedback.

Data privacy guarantees are designed for SME trust and compliance: metadata-first collection, minimum necessary content retention, encryption in transit (TLS) and at rest (KMS-managed keys), tenant-scoped access boundaries, configurable retention, and audit trails for all high-risk actions. GDPR-aligned controls are included for export/delete workflows and evidence traceability.

Operationally, the platform is AWS-first and event-driven. ECS Fargate runs core services; EventBridge/SQS/Step Functions orchestrate asynchronous actions; S3 stores compliance evidence; CloudWatch/SNS provide monitoring and alerting. This keeps ops overhead manageable for a lean team while allowing future scale-out.

The architecture is explicitly designed for non-security operators. Incident response is playbook-driven with guided wizard steps, decision gates, and prebuilt communication templates. Automated offboarding and least-privilege enforcement reduce manual errors. Compliance posture is continuously measured against ISO 27001, GDPR, and SOC2-lite mappings with clear gap dashboards.

Economically, the platform fits SME constraints with tiered pricing and pay-as-you-grow adoption. Shared multi-tenant delivery keeps baseline costs low; premium tiers unlock stronger isolation, retention, and advanced integrations. This aligns product capability, security maturity, and budget progression over time.

## 3) Team & Delivery Plan (6 months)

Team (8 FTE): Product/Security Analyst (1), Solution Architect/Tech Lead (1), Backend (2), Frontend Web (1), Flutter (2), DevSecOps/QA (1).

- Month 1: AWS foundation, tenant model, auth baseline, integration skeletons.
- Month 2: Asset inventory + classification, policy engine v1, offboarding automation.
- Month 3: AI governance v1 (shadow AI, prompt guard, DLP pattern controls).
- Month 4: Incident playbooks + compliance control mappings.
- Month 5: Unified dashboard + Flutter mobile/desktop workflows.
- Month 6: Hardening, pilot with 2–3 SMEs, false-positive tuning, launch readiness.

Riskiest assumption to validate first: AI detection quality is actionable for SMEs without overwhelming false positives. Validate in first 6 weeks with pilot telemetry and acceptance thresholds.

## 4) AI Governance Module

Detection channels:
- Endpoint/browser telemetry for AI tool usage.
- Domain/API and OAuth app discovery for shadow AI.
- Prompt/data inspection via DLP patterns.

Governance controls:
- Approved AI tool catalog by department.
- Policy levels: advisory, justification-required, block.
- Sensitive data guardrails for PII/financial/IP/secrets.
- Automated offboarding: revoke tokens/sessions and third-party app access.

Operator UX:
- Risk scoring dashboard.
- Guided incident wizards for non-security staff.
- Full audit trail for compliance evidence.

## 5) Next Version (Large Scale) Tech Stack Notes

- Stream processing: migrate EventBridge/SQS-heavy paths to Amazon MSK + Managed Flink.
- Compute: expand from ECS/Lambda to EKS for high-density services.
- Data: add OpenSearch (hunt/search), DynamoDB (high-throughput state), S3 data lake + Glue/Athena, Redshift for advanced BI.
- AI/ML: SageMaker Pipelines + Registry, optional KServe on EKS for custom serving.
- Tenancy: shard-by-segment, key-per-tenant for high isolation tiers, dedicated tenant deployment options.
- Reliability: OpenTelemetry, SLO/SLI, canary deployments, multi-region DR.
