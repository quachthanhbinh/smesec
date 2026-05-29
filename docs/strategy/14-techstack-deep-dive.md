# SMESec Platform — Tech Stack Deep Dive

**Date:** 2026-05-29
**Purpose:** Presentation reference for Solution Architect (SA) review. Covers every technology in the system: what it is, why it was chosen, pros/cons, trade-offs, and likely follow-up questions with answers.
**Audience:** Engineering leadership, Solution Architect
**Related:** [01-system-architecture.md](01-system-architecture.md) · [02-design-document.md](02-design-document.md) · [03-two-track-approach.md](03-two-track-approach.md)

---

## Table of Contents

1. [Architecture Pattern Choices](#1-architecture-pattern-choices)
2. [Backend — Go (Primary API Services)](#2-backend--go-primary-api-services)
3. [Backend — Python / FastAPI (ML/AI Services)](#3-backend--python--fastapi-mlai-services)
4. [Frontend — React / Next.js (Web Dashboard)](#4-frontend--react--nextjs-web-dashboard)
5. [Frontend — Flutter (Mobile iOS + Android)](#5-frontend--flutter-mobile-ios--android)
6. [Frontend — Chrome MV3 (Browser Extension)](#6-frontend--chrome-mv3-browser-extension)
7. [Auth — Keycloak (Self-Hosted SSO + MFA)](#7-auth--keycloak-self-hosted-sso--mfa)
8. [Database — PostgreSQL with Row-Level Security](#8-database--postgresql-with-row-level-security)
9. [Cache — ElastiCache Redis](#9-cache--elasticache-redis)
10. [Event Bus — AWS EventBridge](#10-event-bus--aws-eventbridge)
11. [Workflow Orchestration — AWS Step Functions](#11-workflow-orchestration--aws-step-functions)
12. [ML Platform — AWS SageMaker](#12-ml-platform--aws-sagemaker)
13. [Deepfake Detection — Hive Moderation API](#13-deepfake-detection--hive-moderation-api)
14. [Prompt Injection Detection — Lakera Guard API](#14-prompt-injection-detection--lakera-guard-api)
15. [Compliance Automation — Vanta](#15-compliance-automation--vanta)
16. [Audit Storage — S3 Object Lock (WORM)](#16-audit-storage--s3-object-lock-worm)
17. [Infrastructure — AWS ECS Fargate](#17-infrastructure--aws-ecs-fargate)
18. [CDN + Edge — CloudFront + WAF](#18-cdn--edge--cloudfront--waf)
19. [Backup Storage — Cloudflare R2](#19-backup-storage--cloudflare-r2)
20. [Observability — CloudWatch + AWS Security Services](#20-observability--cloudwatch--aws-security-services)
21. [Multi-Tenancy Strategy — Shared DB + RLS](#21-multi-tenancy-strategy--shared-db--rls)
22. [Local DLP — Microsoft Presidio (WASM)](#22-local-dlp--microsoft-presidio-wasm)
23. [Rejected Alternatives (Why We Didn't Choose X)](#23-rejected-alternatives-why-we-didnt-choose-x)
24. [SA Question Bank — Anticipated Questions & Answers](#24-sa-question-bank--anticipated-questions--answers)

---

## 1. Architecture Pattern Choices

### Clean Architecture + Hexagonal Architecture (Ports & Adapters)

**What it is:**
Clean Architecture (Robert C. Martin) organizes code into concentric layers: Domain (innermost) → Application → Infrastructure → Interface (outermost). The Dependency Rule: all source code dependencies point inward. The Domain layer has zero knowledge of frameworks, databases, or external services.

Hexagonal Architecture (Alistair Cockburn) expresses the same idea as "Ports & Adapters": the domain defines ports (interfaces), and all external systems are adapters implementing those ports.

**Why chosen:**
- Core business logic (asset classification, access governance, risk scoring) is completely independent of AWS, Keycloak, PostgreSQL, or any vendor
- Swapping infrastructure (e.g., PostgreSQL → DynamoDB) requires zero changes to domain or application layers
- All integration adapters (Google, M365, Slack) are isolated — vendor API changes don't ripple into domain logic
- Testability: domain logic is pure Go structs, 100% unit-testable without mocks of any infrastructure

**Pros:**
- High maintainability over time
- True separation of concerns
- Independent deployability of layers
- Vendor-neutral domain (no vendor lock-in in core logic)
- Full test coverage achievable without integration environment

**Cons:**
- More initial boilerplate (port interfaces, DTO mappers between layers)
- Steeper learning curve for engineers new to the pattern
- Risk of over-engineering if applied too rigidly to simple CRUD paths

**Trade-offs:**
- Accepting more upfront code structure for long-term maintainability. For SMESec, where integrations (Google, M365, Slack, AWS IAM) are the highest risk area for vendor change, this trade-off is correct.

**SA likely asks:** *"Why not a simpler layered architecture or MVC?"*
Answer: MVC collapses business logic into controllers or services that have direct DB/API dependencies. When any of the 4+ vendor integrations changes their API, you'd need to rewrite business logic. The port abstraction means the adapter changes, the domain stays intact.

---

### Event-Driven Architecture (EDA)

**What it is:**
Services communicate asynchronously via domain events published to AWS EventBridge. Track 1 playbooks are triggered by Track 2 threat detection events — but Track 1 never calls Track 2 directly. Loose coupling is enforced at the event schema level (`ThreatDetectionEvent`).

**Why chosen:**
- Track 1 and Track 2 must be independently deployable. Direct synchronous calls would create coupling — if Track 2 is down, Track 1 incidents would fail to trigger.
- Audit log: every domain event is immutable and replayable
- Scalability: events queue up during load spikes, services process at their own pace
- Enables future integrations (webhook delivery, SIEM export) without changing producers

**Pros:**
- Decoupled services — Track 1 continues operating even when Track 2 is down
- Natural audit trail (every event is a business fact)
- Easy to add new consumers without touching producers

**Cons:**
- Eventual consistency: event consumers process asynchronously, so there's a delay between event publish and effect
- Harder to debug (need distributed tracing)
- Schema evolution (event versioning) must be handled explicitly

**Trade-offs:**
- For real-time alerting, eventual consistency is acceptable (seconds, not hours). Incident playbooks tolerate a 1-5 second delay between threat detection and playbook trigger. The decoupling benefit outweighs the complexity cost.

---

## 2. Backend — Go (Primary API Services)

**What it is:**
Go (Golang) is a statically typed, compiled language by Google. Used for all Track 1 services: Asset Inventory, Access Governance, Integration Sync, Playbook Engine, Compliance, and API Gateway.

**Why chosen:**
- **Concurrency:** Go's goroutines and channels are designed for high-concurrency I/O work — exactly what integration sync (polling Google, M365, Slack, AWS simultaneously) requires. Goroutines are ~2KB stack each; Java threads are ~1MB.
- **Type safety:** Catches integration bugs at compile time, not runtime
- **Performance:** Compiled binary; typical API response <10ms at p99. Critical for real-time access governance checks.
- **Ecosystem:** First-class AWS SDK (`aws-sdk-go-v2`), excellent HTTP frameworks (`Echo`, `Gin`), strong OAuth library support
- **Operations:** Single binary deployment to ECS Fargate — no JVM warmup, no Python package hell

**Pros:**
- Fast compile times (seconds, not minutes)
- Excellent concurrency primitives for sync workers
- Low memory footprint (important for ECS Fargate cost)
- Strong standard library (HTTP, crypto, JSON)
- Easy to containerize (single static binary)

**Cons:**
- No built-in generics before Go 1.18 (we're on 1.22+ — resolved)
- Verbose error handling (`if err != nil` boilerplate)
- Smaller ML/data science ecosystem than Python
- Less talent pool than Java/JavaScript (hiring risk)

**Trade-offs:**
- Go vs Java: Java has a larger talent pool but higher memory footprint (important for Fargate cost) and slower startup (cold-start latency). Go wins for a cost-sensitive SME SaaS.
- Go vs Node.js: Node.js is single-threaded and needs worker_threads for CPU work; Go handles CPU and I/O concurrency natively. For integration sync (50+ concurrent API calls), Go is more appropriate.

**Specific use in SMESec:**
```
Integration Sync: 4 adapters (Google, M365, Slack, AWS) run in parallel goroutines
                  Each polling on 15-min intervals, delta-sync
Asset Discovery:  Concurrent fan-out across all connected providers
                  Results merged and classified in-process
Access Governance: Synchronous check path <100ms SLA (login gate check)
```

---

## 3. Backend — Python / FastAPI (ML/AI Services)

**What it is:**
Python with FastAPI framework for Track 2 services: ThreatDetectionSvc, DLPSvc, DeepfakeSvc. Python is the dominant language in ML/AI.

**Why chosen:**
- Python is the only practical choice for ML inference: PyTorch, transformers (HuggingFace), scikit-learn, Presidio — all Python-native
- FastAPI is ASGI-based (async), matching Go's performance profile for I/O-bound services
- FastAPI auto-generates OpenAPI docs — contracts shared with Go services

**Pros:**
- Unmatched ML library ecosystem (PyTorch, HuggingFace, SageMaker SDK)
- FastAPI is async-native, handles concurrent inference requests
- Auto-documentation (OpenAPI spec autogenerated)
- Type annotations (Pydantic) provide runtime validation

**Cons:**
- GIL (Global Interpreter Lock) limits true CPU parallelism — mitigated by running multiple Fargate tasks
- Slower cold-start than Go binaries
- Package management complexity (requirements.txt, virtual envs, Docker layer caching)
- Higher memory footprint than Go for the same task

**Trade-offs:**
- Using two languages (Go + Python) introduces operational complexity — two runtimes, two CI pipelines. The decision is justified because the ML ecosystem in Python has no viable alternative. The inter-service boundary is a clean REST/gRPC contract, so the two languages never mix.

---

## 4. Frontend — React / Next.js (Web Dashboard)

**What it is:**
React is the dominant UI component library. Next.js adds Server-Side Rendering (SSR), static generation, file-based routing, and API routes on top of React.

**Why chosen:**
- **SSR for security dashboards:** Server-rendered HTML means the security dashboard is readable without JavaScript — important for compliance audit reports
- **App Router (Next.js 13+):** React Server Components reduce client-side JavaScript payload; better performance on SME employee hardware (which may be slower)
- **Ecosystem:** Largest frontend talent pool; rich component library ecosystem (shadcn/ui, Radix UI, TailwindCSS)
- **API Routes:** Thin BFF (Backend-for-Frontend) layer in Next.js reduces direct exposure of internal Go services to browser

**Pros:**
- SSR improves SEO (relevant for marketing pages) and performance on slow devices
- Server Components shift heavy data-fetching server-side (no waterfall requests)
- Large talent pool and community
- Vercel deployment as alternative to ECS (escape hatch if needed)

**Cons:**
- Next.js App Router is still maturing (some APIs changed between v13-v15)
- Server Components add mental model complexity (server vs client boundary)
- Node.js server required for SSR (ECS Fargate service vs static hosting)
- Bundle size can grow if not carefully managed

**Trade-offs:**
- Next.js vs pure React SPA: An SPA loads all JS upfront and renders on the client. For a security dashboard shown to non-technical SME owners, fast initial paint (SSR) and ability to work in slow-network conditions outweigh the SPA simplicity trade-off.
- Next.js vs Vue/SvelteKit: React ecosystem depth (component libraries, chart libraries, talent pool) wins for a product that needs to ship quickly.

---

## 5. Frontend — Flutter (Mobile iOS + Android)

**What it is:**
Flutter is Google's cross-platform UI framework using the Dart language. A single codebase compiles to native iOS, Android, and can also target web and desktop.

**Why chosen:**
- **Single codebase:** One Flutter codebase produces iOS + Android apps. The alternative (React Native, or two native apps) costs 2x or has performance compromises.
- **Performance:** Flutter renders directly via Skia/Impeller GPU engine (bypasses native UI components) — consistent 60-120fps animations
- **Push notifications + biometrics:** Native plugins (`flutter_local_notifications`, `local_auth`) needed for security app functionality
- **Background processing:** Dart isolates for background sync jobs

**Pros:**
- True native performance (not a WebView)
- Single codebase for iOS + Android saves ~40% mobile development cost
- Hot reload for fast iteration
- Strong widget library (Material + Cupertino)

**Cons:**
- Dart language has smaller talent pool than JavaScript/Swift/Kotlin
- Flutter app binary sizes are larger than native (typically +10-15MB)
- Some platform-specific plugins have delayed support (e.g., new iOS features)
- Not ideal for desktop-heavy features (web/macOS Flutter is less mature)

**Trade-offs:**
- Flutter vs React Native: React Native bridges to native components which introduces bridge latency. Flutter's own rendering pipeline avoids this. For an app with security alerts and biometric auth, reliability > React Native's JS familiarity benefit.
- Flutter vs native (Swift + Kotlin): Native gives maximum platform capability but doubles the mobile team required. SMESec is resource-constrained — one Flutter engineer can maintain both platforms.

---

## 6. Frontend — Chrome MV3 (Browser Extension)

**What it is:**
Chrome Manifest V3 (MV3) is the current Chrome extension platform. The extension runs as a Service Worker (background), Content Script (in-page), and Popup (toolbar UI). Also compatible with Edge (Chromium-based).

**Why chosen:**
- The DLP requirement (intercept LLM prompts before they leave the browser) is **only achievable via a browser extension**. No server-side solution can inspect what a user types into chatgpt.com before submission.
- MV3 is the mandatory standard for new Chrome extensions (MV2 deprecated)
- Edge compatibility: Same codebase works on Edge (Chromium), covering both Chrome and Edge users

**Key architectural decision — Local inference:**
```
User types in ChatGPT → Content Script intercepts
→ Send text to Service Worker
→ Service Worker runs Presidio WASM (local PII detection)
→ If PII found → show warning UI + optional block
→ NO content sent to SMESec servers (privacy-preserving)
```

**Pros:**
- Only technology capable of local browser-side DLP
- Privacy-preserving: text never leaves the browser for DLP checks
- Works across all AI tools (ChatGPT, Copilot, Gemini, Claude) without vendor-specific integrations
- Tenant-scoped allow-list configurable per company

**Cons:**
- **MV3 Service Worker terminates after 30 seconds of inactivity** — hard constraint. Requires keep-alive pattern (alarm API). This is a critical technical risk (SA discovery from debate).
- Limited persistent storage in Service Workers
- User must install extension manually (distribution friction, especially on non-managed devices)
- Cannot intercept HTTPS traffic on WebSocket-only AI tools (workaround: DOM mutation observer)
- Chrome Web Store review for security extensions: 2-6 weeks (planning risk)

**Trade-offs:**
- The 30-second Service Worker termination is a genuine platform constraint. Mitigation: use `chrome.alarms` API to schedule keep-alive pings every 25 seconds. This must be prototyped and validated in Sprint 1 Week 1 (hard gate from SA).
- Local WASM inference (Presidio) vs server-round-trip: WASM is slower (~50-200ms) but preserves privacy. For a DLP product, privacy is non-negotiable. Accept the latency.

---

## 7. Auth — Keycloak (Self-Hosted SSO + MFA)

**What it is:**
Keycloak is an open-source Identity and Access Management (IAM) solution by Red Hat. It provides OIDC (OpenID Connect), SAML 2.0, social login (Google, Microsoft), and MFA (TOTP, WebAuthn). Deployed on ECS Fargate.

**Why chosen:**
- **Zero per-user cost:** Auth0 charges ~$0.23/MAU. At 10 tenants × 500 users = 5,000 MAU, Auth0 = $1,150/mo. Keycloak on ECS Fargate = ~$50/mo (compute only). Saves ~$13K/yr.
- **OIDC + SAML 2.0:** Google Workspace and M365 federation built-in — no custom code needed for SSO
- **Full control:** Custom MFA flows, branding per tenant, GDPR Data Processing Agreement (DPA) with no third-party data processor risk
- **JWT RS256:** JWTs signed with RSA-256, verifiable by any service without calling Keycloak

**Critical HA requirements (R-C6):**
1. Minimum 2 ECS tasks, active-active (not just multi-AZ placement)
2. JWKS (JSON Web Key Set) caching mandatory — JWT validation must not require Keycloak to be up
3. Keycloak database must be separate from application database

**Pros:**
- Significant cost saving vs SaaS alternatives
- Full OIDC/SAML standard compliance
- GDPR-friendly (self-hosted, no data leaves your infrastructure)
- Mature product (10+ years, Red Hat-backed)

**Cons:**
- **Operational burden:** Updates, patches, HA configuration are your responsibility
- Keycloak configuration is complex (realms, clients, identity providers)
- JVM-based: higher memory footprint than Go services (~512MB RAM baseline per Keycloak instance)
- Admin UI complexity can slow onboarding for engineers new to Keycloak

**Trade-offs:**
- Keycloak vs Auth0/WorkOS: At small scale (<1,000 users), Auth0 is operationally simpler. At SMESec's projected scale (50 tenants × 200 users = 10,000 MAU), Keycloak saves ~$25K/yr. The operational cost of running Keycloak is ~1 engineer-day/quarter. The math strongly favors Keycloak.
- Fallback decision point: If DevSecOps capacity is insufficient at v1 launch, re-evaluate WorkOS (~$500-1,000/mo) before v1.5.

**SA likely asks:** *"What happens if Keycloak goes down?"*
Answer: JWKS caching means all services can continue validating JWTs for up to 6 hours without Keycloak. New logins would fail, but authenticated sessions continue working. ECS Fargate min 2 tasks with ALB health checks provide 99.9%+ availability.

---

## 8. Database — PostgreSQL with Row-Level Security

**What it is:**
PostgreSQL is the world's most advanced open-source relational database. Row-Level Security (RLS) is a PostgreSQL feature that attaches security policies directly to tables, filtering rows based on a session variable (`app.tenant_id`).

**Why chosen for multi-tenancy:**
PostgreSQL RLS enforces tenant isolation at the database engine level — not the application level. Even if the Go application code has a bug (missing WHERE clause, incorrect join), the database itself prevents cross-tenant data access.

**Schema pattern:**
```sql
-- Every domain table has these mandatory columns
tenant_id      UUID NOT NULL REFERENCES tenants(id)
data_residency VARCHAR(10) NOT NULL  -- 'US' | 'EU' | 'APAC'

-- RLS policy (auto-applied to every query)
CREATE POLICY tenant_isolation ON assets
    USING (tenant_id = current_setting('app.tenant_id')::UUID);

-- Application middleware sets this per-request
SET LOCAL app.tenant_id = '<uuid>';
```

**Pros:**
- Defense-in-depth: isolation at DB level + application level
- PostgreSQL is battle-tested at large scale (used by Shopify, Instagram)
- JSONB for flexible metadata without sacrificing relational integrity
- Multi-AZ RDS = managed backups, automated failover, no DBA required
- Native support for UUID, TIMESTAMPTZ, ENUM — all needed for audit data

**Cons:**
- RLS policies can be complex to debug (query plan changes)
- `SET LOCAL app.tenant_id` must be called in every transaction — missed transactions would get no data (fail-safe) but need careful implementation
- PostgreSQL horizontal scaling (sharding) is harder than NoSQL at very large scale (>100TB)
- RDS Multi-AZ is more expensive than single-AZ (~2x cost)

**Trade-offs:**
- PostgreSQL vs DynamoDB: DynamoDB has no RLS concept — tenant isolation must be fully implemented in application code. That's a single point of failure. PostgreSQL RLS provides a second isolation layer. For a security product, this extra safety is worth the operational complexity.
- Shared DB + RLS vs separate DB per tenant: Separate DB costs ~$100/mo per tenant at RDS pricing. At 50 tenants = $5,000/mo vs ~$200/mo for shared Multi-AZ RDS. Shared + RLS wins at SME scale.

**SA likely asks:** *"How do you prevent RLS bypass?"*
Answer: (1) `FORCE ROW LEVEL SECURITY` applies even to table owners. (2) The app connects as `app_role`, not a superuser. (3) CI tests assert cross-tenant isolation on every PR. (4) `data_residency` column enforces EU data never goes to US-region queries.

---

## 9. Cache — ElastiCache Redis

**What it is:**
Redis is an in-memory key-value store. AWS ElastiCache manages Redis clusters with automatic failover, backups, and scaling.

**Used for:**
- Session tokens (JWT blacklist, 15-minute TTL)
- Rate limiting counters (per-tenant API rate limits)
- Integration sync state (delta link tokens, last-sync cursors)
- Short-lived cache (user permission sets, asset classification results)

**Pros:**
- Sub-millisecond reads (critical for auth middleware on every request)
- Atomic operations (INCR for rate limiting, SETNX for distributed locks)
- ElastiCache = no Redis ops burden (managed by AWS)
- Cluster mode for horizontal scaling

**Cons:**
- In-memory: data lost on restart unless AOF/RDB persistence configured
- ElastiCache adds cost (~$50-100/mo for cache.t3.small)
- Not a persistent store — cannot be used as a source of truth

**Trade-offs:**
- Redis vs PostgreSQL for rate limiting: PostgreSQL UPDATE operations are ~10ms; Redis INCR operations are ~0.1ms. For per-request rate limiting on every API call, PostgreSQL would add unacceptable latency. Redis wins.

---

## 10. Event Bus — AWS EventBridge

**What it is:**
AWS EventBridge is a serverless event bus. Services publish events; EventBridge routes them to subscribers based on rules. Supports schema registry, replay, and archive.

**Used for:**
- `ThreatDetectionEvent` from Track 2 → triggers Track 1 playbooks
- `AssetDiscovered` → triggers classification workflow
- `AccessRevoked` → triggers notification service
- All domain events between microservices (async coupling)

**Pros:**
- Serverless: no infrastructure to manage
- At-least-once delivery guarantee
- Event Archive and Replay: critical for audit trail and debugging
- Schema Registry: enforces event contract between producers and consumers
- Native integration with Step Functions, Lambda, SQS

**Cons:**
- Maximum event size: 256KB (sufficient for all SMESec events)
- EventBridge has ~0-500ms propagation latency (not for synchronous paths)
- Vendor lock-in to AWS — migrating would require replacing all event publishing code

**Trade-offs:**
- EventBridge vs Kafka: Kafka provides guaranteed ordering and replay, but requires managing a Kafka cluster (or MSK at ~$500/mo). For SMESec's scale (low thousands of events/day), EventBridge at near-zero cost is correct. Kafka becomes relevant at >1M events/day.
- EventBridge vs SQS: SQS is point-to-point; EventBridge is pub/sub (fanout). Multiple consumers of the same `ThreatDetectionEvent` (playbook engine, notification service, audit log) require pub/sub. EventBridge wins.

---

## 11. Workflow Orchestration — AWS Step Functions

**What it is:**
AWS Step Functions is a serverless workflow orchestration service. Workflows are defined as state machines (JSON/YAML). Used for long-running, multi-step processes with retry, error handling, and parallel execution.

**Used for:**
- Incident playbook execution (multi-step: isolate → notify → collect evidence → escalate)
- Automated offboarding (revoke Google access → revoke M365 → revoke Slack → revoke AWS IAM → notify HR → generate compliance record)
- Scheduled compliance evidence collection workflows

**Pros:**
- Visual workflow editor — non-engineers can understand the playbook flow
- Built-in retry logic, error catching, and timeouts — no custom retry code
- Long-running workflows (hours/days) without holding up server resources
- Audit trail of every step execution (Step Functions execution history)
- Parallel execution (concurrent API calls to Google + M365 + Slack simultaneously)

**Cons:**
- Step Functions Express Workflows have 5-minute limit (Standard = no limit, but more expensive)
- State machine JSON/YAML can become verbose for complex branching logic
- Cold-start latency (~100ms) on first execution

**Trade-offs:**
- Step Functions vs custom code loops: Implementing retry, timeout, parallel execution, and error handling in Go code is possible but adds significant complexity and testing burden. Step Functions externalizes this complexity into a managed service with built-in observability.
- Standard vs Express Workflows: Standard workflows (at-least-once execution, full history) for incident playbooks (correctness matters). Express workflows (at-most-once, cheaper) for high-volume scheduled jobs.

**SA likely asks:** *"Why not just call the APIs sequentially from the application service?"*
Answer: A sequential API call approach fails if any step fails (no retry). Step Functions handles partial failure (idempotent retry), long-running steps (offboarding can take 10+ minutes across 4 APIs), and parallel execution (simultaneous revocation across systems). The compliance audit trail from Step Functions execution history is also a requirement.

---

## 12. ML Platform — AWS SageMaker

**What it is:**
AWS SageMaker is a fully managed ML platform: training jobs, model registry, endpoint deployment, A/B testing, and drift monitoring.

**Used for:**
- BERT fine-tuning for prompt injection detection (Track 2)
- Shadow AI risk scoring model training and inference
- Model registry (version, lineage, approval workflow)
- A/B testing between model versions in production

**Pros:**
- Managed: no MLOps infrastructure to build (no Kubeflow, no custom training orchestration)
- Pay-per-use training (no GPU server to maintain)
- Built-in model drift monitoring (critical for ML systems)
- Integrates with EventBridge for inference result routing
- SageMaker Endpoints auto-scale (handle traffic spikes)

**Cons:**
- SageMaker training jobs have ~5-10 minute overhead even for small jobs
- SageMaker Endpoint cold-start: 10-20 seconds (must validate this with load testing in S3 — SA discovery)
- Vendor lock-in: SageMaker uses custom container format and deployment API
- Cost: SageMaker Endpoints are ~4x more expensive than running inference on ECS Fargate directly

**Trade-offs:**
- SageMaker vs self-managed inference on ECS: At Track 2 alpha stage, the managed tooling (model registry, drift monitoring, A/B testing) is worth the cost premium. Once ML patterns are stable and inference load is predictable, migrating hot-path inference to ECS Fargate may reduce cost.
- SageMaker vs Vertex AI (Google): Single cloud vendor (AWS) simplifies compliance scope, cost management, and IAM policy. No multi-cloud for v1.

**Critical: Cold-start validation (Gate 1 prerequisite)**
SageMaker Endpoint cold-start latency (10-20 seconds) must be validated in Sprint 3 before the browser extension ships in Sprint 5. If cold-start is unacceptable, options are: (a) Provisioned Concurrency at higher cost, or (b) move inference to ECS Fargate with a warm container.

---

## 13. Deepfake Detection — Hive Moderation API

**What it is:**
Hive Moderation is an AI API provider specializing in content moderation. Their API provides voice deepfake detection and video deepfake detection as a REST API. Pay-per-use (~$0.01/check).

**Why chosen:**
- Only SME-accessible tool with a production-ready deepfake detection API (voice + video)
- No training data required — Hive maintains and updates models
- Sub-5-second response time (within Track 2 SLA)
- Pay-per-use removes the need to build and maintain a custom deepfake model

**Pros:**
- Immediate capability — no R&D required for deepfake detection
- Vendor handles model updates as deepfake technology evolves
- Pay-per-use pricing aligns with SMESec's usage-based cost model
- Combined accuracy with out-of-vocabulary (OOV) detection ≈ 99% fraud prevention

**Cons:**
- Vendor dependency: if Hive raises prices or goes offline, Track 2 deepfake detection fails
- 1-2 week lead time for API access (must submit request in Sprint 1 Week 1)
- Data sent to Hive: audio/video frames sent to third party (privacy consideration — DPA required)
- Accuracy is vendor-reported, independently must be validated before Gate 3

**Fallback:**
- Resemblyzer (open-source, ~78-82% voice deepfake accuracy) used as fallback until Hive contract signed, and as permanent fallback if Hive becomes unavailable. DeepfakeDetector abstraction interface defined in Week 1 to allow hot-swap.

**Trade-offs:**
- Build vs Buy: Building a custom deepfake detection model requires a dedicated ML research team (6-12 months to match Hive accuracy). At SME pricing ($0.01/check × 1,000 checks/mo × 50 customers = $500/mo), Buy is unambiguously correct.

---

## 14. Prompt Injection Detection — Lakera Guard API

**What it is:**
Lakera Guard is an AI security API that detects prompt injection attacks — attempts by malicious users to override LLM system prompts. Pay-per-use (~$0.001/request).

**Why chosen:**
- Production-validated: Covers known injection patterns + novel variants via continuous red-team updates
- Accuracy: TPR >85%, FPR <2% (validated targets)
- Lead time only 1-2 weeks
- ~$50/mo at 50K daily checks (50 customers × 1K checks/day)

**Pricing hard gate:**
- Target: <$0.05/request
- Decision deadline: Week 2 (Sprint 1 end)
- Fallback: WASM-only BERT model (lower accuracy: TPR ~75% vs >85%, but zero API cost)

**Pros:**
- Immediate capability — skip R&D for prompt injection baseline
- Covers novel injection variants as Lakera updates their model
- Low cost at SME scale

**Cons:**
- Vendor dependency for a security-critical function
- Data leaves customer environment to Lakera API (privacy consideration)
- If FPR climbs (false positives), customer frustration is high
- Internal BERT model (v2 target) will provide vendor independence

**Future:** Internal BERT fine-tuned classifier is on the Sprint 23-24 roadmap (Enterprise-only evaluation). This replaces Lakera Guard for customers with strict data privacy requirements (no third-party API calls).

**Trade-offs:**
- Lakera vs in-house BERT from day 1: BERT fine-tuning requires 3-4 sprints of ML work plus dataset curation. Lakera provides >85% TPR immediately. The right sequence is Buy now → Build later for Enterprise tier.

---

## 15. Compliance Automation — Vanta

**What it is:**
Vanta is a compliance automation SaaS platform. It connects to AWS, GitHub, and other tools to continuously collect evidence for SOC 2, ISO 27001, HIPAA, and GDPR. Cost: $4-6K/yr (Startup plan).

**Why chosen:**
- SOC 2 Type 1 is a v1 launch requirement (customer expectation for a security product)
- Building compliance evidence collection manually would take 3 months of engineering ($60K+) vs $4-6K Vanta
- AWS + GitHub connectors are native (covers SMESec's entire infrastructure)
- Auditor portal: compliance auditors can view evidence directly without PM involvement
- 60-day SOC 2 Type 1 timeline is achievable with Vanta (vs 9-12 months manual)

**Critical timing:**
Vanta must start collecting evidence no later than Week 13 (Sprint 7). Starting Week 8 provides buffer. Must start 13 weeks before v1 launch (Week 26) to have sufficient observation window.

**Pros:**
- Instant compliance visibility dashboard
- 24/7 automated evidence collection
- Native AWS + GitHub integrations
- Significantly faster SOC 2 compared to manual

**Cons:**
- $4-6K/yr is a real cost for an early-stage startup
- Vanta covers evidence collection but not control implementation (SMESec engineers must still implement the actual controls)
- Connector configuration can take 2-3 weeks (lead time)
- Vanta evidence is necessary but not sufficient — still need a compliance consultant for final SOC 2 audit

**Trade-offs:**
- Vanta vs manual SOC 2: Manual SOC 2 evidence collection = 20 hours/week PM workload for 6 months. Vanta eliminates this. At PM fully-loaded cost of $150/hr, manual = $72K vs Vanta $5K. No trade-off — Vanta wins decisively.

---

## 16. Audit Storage — S3 Object Lock (WORM)

**What it is:**
AWS S3 Object Lock enables Write-Once-Read-Many (WORM) storage. Once an object is written with an Object Lock, it cannot be modified or deleted — even by AWS account root users. 7-year retention configured.

**Why chosen:**
- WORM audit logs are a compliance requirement for SOC 2 and ISO 27001 (evidence of security controls cannot be tampered with)
- S3 Object Lock provides WORM at near-zero cost (~$0.023/GB/month)
- Immutable logs prove to customers that SMESec engineers cannot delete audit evidence
- 7-year retention satisfies financial and security compliance requirements

**Pros:**
- True immutability (even AWS support cannot delete Object Lock objects)
- ~$0.023/GB/month — extremely cheap for audit log volume
- S3 Lifecycle policies can transition to Glacier after 1 year ($0.004/GB/month)
- No separate vendor or service required

**Cons:**
- S3 Object Lock Compliance mode cannot be shortened even by administrators — must plan retention period carefully
- Query performance: S3 is object storage, not queryable like a database. Log queries require Athena or CloudWatch Logs integration.
- Accidental Object Lock on wrong objects is irrecoverable

**Trade-offs:**
- S3 Object Lock vs managed logging services (Datadog, Splunk): Managed logging services are excellent for querying and alerting, but they do not provide legal-grade WORM immutability. SMESec uses S3 Object Lock for compliance evidence and CloudWatch Logs for operational alerting — complementary, not alternatives.

---

## 17. Infrastructure — AWS ECS Fargate

**What it is:**
AWS ECS (Elastic Container Service) Fargate is a serverless container execution platform. You define containers, CPU, and memory — AWS manages the underlying servers. No EC2 instances to patch or scale.

**Why chosen:**
- Serverless: no server fleet to manage (critical for a small DevSecOps team)
- Per-task pricing: pay only for running containers, not idle servers
- Native AWS integration: IAM roles per task, CloudWatch Logs, Service Connect
- ECS Service auto-scaling: scale out during sync jobs, scale in during off-peak
- Security: each task has its own IAM role (least privilege), no shared credentials

**Fargate vs EKS (Kubernetes):**

| Factor | ECS Fargate | EKS |
|--------|------------|-----|
| Operational complexity | Low (no cluster management) | High (master nodes, node groups, etcd) |
| Setup time | Hours | Days |
| SME SaaS scale (10-50 services) | Perfect | Overengineered |
| Cost at 10 services | ~$200-400/mo compute | +$150/mo cluster fee + higher overhead |
| Learning curve | Minimal | Steep |

**Pros:**
- Zero server management
- Native IAM integration (task roles = fine-grained permissions per service)
- ECS Service Connect provides built-in service discovery
- CloudWatch Container Insights for container metrics
- Integrates natively with ALB for HTTP routing

**Cons:**
- Fargate cold-start: 30-60 seconds for new container launch (not suitable for sub-second scale-out)
- Limited to AWS (vendor lock-in vs Kubernetes portability)
- Max 4 vCPU / 30GB RAM per task (sufficient for SMESec services)

**Trade-offs:**
- ECS Fargate vs EKS: EKS is the right choice at >100 services or when the team has strong Kubernetes expertise. SMESec has 10-15 services and a small DevSecOps team. ECS Fargate gives 90% of EKS capability at 30% of the operational overhead.
- ECS Fargate vs Lambda: Lambda (serverless functions) is excellent for event-driven jobs but has 15-minute max execution and cold-start issues. Integration sync workers need to run for hours — Lambda is not appropriate.

---

## 18. CDN + Edge — CloudFront + WAF

**What it is:**
AWS CloudFront is a global CDN (Content Delivery Network) with 400+ edge locations. AWS WAF (Web Application Firewall) is integrated with CloudFront to filter malicious requests before they reach the application.

**WAF rules configured:**
- OWASP Top 10 managed rule group (SQLi, XSS, path traversal, etc.)
- AWS managed rules for known bad IPs
- Rate limiting: 100 req/sec per IP (prevents credential stuffing)
- Geo-blocking (if required by compliance)

**Pros:**
- Global CDN reduces latency for EU and APAC customers (~50ms vs ~200ms to us-east-1)
- WAF at edge blocks attacks before they reach VPC
- SSL termination at CloudFront (reduces TLS overhead on ALB)
- Static asset caching (Next.js static files, browser extension update manifests)

**Cons:**
- CloudFront cache invalidation can be slow (up to 15 minutes for global propagation)
- WAF managed rules occasionally produce false positives (need tuning period)
- Additional monthly cost (~$0.0085/10K requests + WAF ~$5/rule group/mo)

**Trade-offs:**
- WAF at CloudFront vs application-level WAF: CloudFront WAF blocks at edge — malicious traffic never hits your VPC, reducing compute load and exposure surface. Application-level input validation is also required but not a substitute for edge WAF.

---

## 19. Backup Storage — Cloudflare R2

**What it is:**
Cloudflare R2 is an S3-compatible object storage service. Key differentiator: **zero egress fees** (vs AWS S3's $0.09/GB egress).

**Why chosen:**
- Zero egress cost for browser extension update files (extension updates = frequent small downloads by many users)
- S3-compatible API: existing AWS SDK code works without changes
- S3 Object Lock compatibility: confirmed for WORM audit log backup
- Cloudflare's network reduces latency for European customers (EU data residency)

**Pros:**
- Zero egress fees (saves ~$50-200/mo vs S3 at 50 customers)
- S3-compatible (drop-in replacement)
- Global network (Cloudflare has more edge locations than AWS)

**Cons:**
- Not as feature-complete as S3 (no S3 Select, no S3 event notifications natively)
- Must test S3 Object Lock compatibility before relying on it for compliance (Sprint 1)
- Adds a second storage vendor (small operational complexity)

**Trade-offs:**
- Cloudflare R2 vs AWS S3 only: Using only S3 simplifies operations but adds egress cost. At 50 customers with regular report downloads, R2 saves ~$50-200/mo. Worth maintaining R2 for static assets and extension distribution.

---

## 20. Observability — CloudWatch + AWS Security Services

**What it is:**
AWS CloudWatch provides metrics, logs, alarms, and dashboards. AWS GuardDuty provides threat detection. AWS Security Hub aggregates security findings. AWS CloudTrail logs all AWS API calls.

**Stack:**

| Service | Purpose |
|---------|---------|
| CloudWatch Metrics | Service health, latency, error rates |
| CloudWatch Logs Insights | Log querying and alerting |
| CloudWatch Container Insights | ECS task metrics (CPU, memory) |
| CloudTrail | All AWS API calls (audit trail) |
| GuardDuty | Threat detection (unusual API access, crypto mining, exfiltration) |
| Security Hub | Centralized security findings from GuardDuty, Config, Inspector |
| AWS Config | Resource configuration compliance checks |
| AWS Inspector | Container image vulnerability scanning (ECR integration) |

**Pros:**
- Zero additional cost for CloudWatch (included in AWS)
- Native integration with all AWS services
- GuardDuty + Security Hub provide managed threat detection without ML expertise needed

**Cons:**
- CloudWatch Logs Insights query syntax is non-standard (not SQL or PromQL)
- CloudWatch dashboards are less polished than Datadog or Grafana
- GuardDuty has false positives in dev environments (needs tuning)
- No distributed tracing out-of-box (would need AWS X-Ray or Datadog APM)

**Datadog (v1.5 optional):**
If budget allows post-v1, Datadog APM would add distributed tracing (critical for debugging multi-service flows like incident playbook execution). ~$200/mo for 10 services.

---

## 21. Multi-Tenancy Strategy — Shared DB + RLS

*(Cross-reference: detailed schema in [02-design-document.md](02-design-document.md) §3)*

**Three approaches evaluated:**

| Approach | Description | Decision |
|----------|-------------|----------|
| **Silo** (DB per tenant) | Each tenant has its own RDS instance | ❌ ~$100/mo per tenant; 50 tenants = $5,000/mo |
| **Shared Schema, App isolation** | One DB, app code filters by tenant_id | ❌ Single bug = data breach |
| **Shared Schema, DB-level RLS** | PostgreSQL RLS enforces isolation | ✅ Defense-in-depth |

**RLS implementation:**
```sql
-- Session variable set by middleware on every request
SET LOCAL app.tenant_id = '<tenant_uuid>';

-- Policy: DB engine rejects any query not matching tenant_id
CREATE POLICY tenant_isolation ON assets
    USING (tenant_id = current_setting('app.tenant_id')::UUID);
FORCE ROW LEVEL SECURITY;  -- Applies even to table owner
```

**Data residency enforcement:**
```sql
-- EU customer in eu-west-1 RDS can never query US tenant rows
data_residency VARCHAR(10) CHECK (data_residency IN ('US', 'EU', 'APAC'))
```

**SA likely asks:** *"Has this been pen tested? What's your blast radius if RLS is bypassed?"*
Answer: (1) RLS is tested in CI on every PR with cross-tenant query tests. (2) If bypassed via SQL injection, the blast radius is limited to the current tenant's `data_residency` region (EU data is in a separate RDS cluster). (3) Third-party pentest scheduled before v1 launch validates this. (4) Defense-in-depth: even if DB-level RLS failed, the application layer also filters by tenant_id.

---

## 22. Local DLP — Microsoft Presidio (WASM)

**What it is:**
Microsoft Presidio is an open-source PII detection library. In the browser extension, it is compiled to WebAssembly (WASM) and runs entirely within the browser's Service Worker — no content ever leaves the user's device for DLP checks.

**Detects:** Credit card numbers, SSNs, email addresses, phone numbers, IBAN codes, IP addresses, and custom company-specific patterns.

**Why chosen:**
- The only open-source PII detection library with WASM compilation support
- Microsoft-backed, actively maintained
- Covers >15 PII entity types out of box
- Custom recognizers can be added for company-specific data patterns

**Pros:**
- Privacy-preserving: text never sent to server for PII analysis
- Sub-200ms detection (WASM is fast)
- Works offline (Service Worker + WASM = no network required for DLP)
- Free (no API cost)

**Cons:**
- WASM binary adds ~8-15MB to extension size (affects install time)
- Custom patterns (trade secrets, proprietary data) cannot be detected without text similarity ML — Presidio only detects structural PII
- False positives on code snippets (e.g., "123-456-7890" in test data)
- No semantic understanding (cannot detect that "our Q3 revenue is $50M" is confidential)

**Trade-offs:**
- Presidio WASM vs server-side DLP: Server-side DLP requires sending text to SMESec servers, creating a privacy risk and requiring a DPA. WASM local DLP avoids both. Accept the limitation (structural PII only, no semantic analysis) in exchange for privacy.
- Presidio vs cloud DLP (AWS Comprehend): Cloud DLP is more accurate but breaks the local-inference guarantee. For the browser extension use case, local inference is non-negotiable.

---

## 23. Rejected Alternatives (Why We Didn't Choose X)

| Technology Considered | Why Rejected |
|-----------------------|--------------|
| **Auth0 / WorkOS** | $23K+/yr at scale vs Keycloak ~$600/yr. Operationally simpler but prohibitively expensive for SME SaaS pricing |
| **Kubernetes (EKS)** | Excessive operational complexity for 10-15 services; ECS Fargate provides 90% capability at 30% overhead |
| **DynamoDB** | No Row-Level Security concept; tenant isolation fully in application code = single point of failure for data breach |
| **MongoDB** | No RLS; weaker consistency guarantees for compliance evidence; PostgreSQL JSONB handles flexible metadata natively |
| **Kafka / MSK** | ~$500/mo for MSK + operational complexity; EventBridge at near-zero cost handles SMESec's event volumes. Revisit at >1M events/day |
| **Silo DB per tenant** | ~$5,000/mo at 50 tenants (vs ~$200/mo shared RDS). Not viable for SME pricing tier |
| **Lambda for sync workers** | 15-minute execution limit; sync workers run for hours. ECS Fargate is appropriate for long-running workloads |
| **Vue.js / SvelteKit** | React ecosystem depth (component libraries, talent pool, shadcn/ui) outweighs framework elegance at v1 speed |
| **React Native** | JS bridge latency vs Flutter's direct GPU rendering. Flutter more reliable for biometric auth and real-time security alerts |
| **Pure in-house BERT (no Lakera)** | 3-4 sprint delay to match Lakera's TPR >85%. Buy Lakera now, build internal BERT for Enterprise v2 |
| **Vertex AI / Azure ML** | Multi-cloud adds compliance complexity and cost management overhead. Single AWS for v1 |
| **Splunk / Datadog from day 1** | CloudWatch is zero additional cost. Add Datadog APM at v1.5 if budget allows |

---

## 24. SA Question Bank — Anticipated Questions & Answers

### Architecture & Design

**Q: How does Track 1 continue working if Track 2 services are down?**
A: Track 1 services never call Track 2 services directly. Communication is one-way: Track 2 publishes `ThreatDetectionEvent` to EventBridge, Track 1 consumes them. If Track 2 is down, Track 1 continues operating — no new threat events are generated, but all access governance, offboarding, inventory, and compliance features remain fully functional.

**Q: How do you handle rate limiting across Google, M365, and Slack APIs?**
A: Each adapter implements token bucket rate limiting with exponential backoff. Google Admin SDK: 100 requests/10 seconds (configurable per API). M365 Graph: delta link sync reduces re-read volume by ~80%. Slack: tier-gated (Business+ for write operations). Rate limit errors trigger backoff + retry, not failure. Integration sync is background, not blocking user-facing requests.

**Q: What is your RTO/RPO for the system?**
A: RDS Multi-AZ: automatic failover in <60 seconds (RTO). Point-in-time recovery to any second in the past 35 days (RPO ~0). ECS Fargate: ALB health checks detect failed tasks in 30 seconds, replacement tasks start in 30-60 seconds. Keycloak: 2 active-active tasks — if one fails, requests route to the other immediately. Overall RTO target: <5 minutes. RPO: ~0 for transactional data.

**Q: How do you guarantee SOC 2 Type 1 certification by v1 launch?**
A: Vanta starts evidence collection at Week 13 — 13 weeks before v1 launch (Week 26). SOC 2 Type 1 requires a minimum observation window (typically 1-3 months). 13 weeks exceeds the minimum. Vanta automates 80% of evidence collection. The gap is control implementation, which is tracked in the Sprint plan. Pentest completes before v1 launch, providing security assessment evidence.

**Q: How do you prevent a tenant from escalating privileges to access another tenant's data?**
A: Three-layer defense: (1) JWT contains `tenant_id` claim, set at login, signed by Keycloak — cannot be forged without the private key. (2) API middleware injects `SET LOCAL app.tenant_id` into every PostgreSQL transaction from the JWT. (3) PostgreSQL RLS policy enforces `tenant_id` at the DB engine level — bypassing layers (1) and (2) still fails at (3). Cross-tenant access tests run in CI on every PR.

### ML/AI Accuracy

**Q: What accuracy can you guarantee for prompt injection detection?**
A: The production gate is TPR >85% AND FPR <2%, validated on a 30-day holdout set independently. Below this threshold, Track 2 prompt injection feature remains in beta (opt-in). Track 1 continues shipping with 100% accuracy for all deterministic features regardless of Track 2 status.

**Q: What happens when Hive Moderation's API goes down?**
A: DeepfakeDetector is an interface. The Hive implementation and the Resemblyzer (open-source, ~78-82% accuracy) implementation are both maintained. If Hive is unreachable (circuit breaker open after 3 failures), traffic automatically routes to Resemblyzer fallback. Customers see a degraded accuracy warning in the dashboard. Hive SLA: 99.9% uptime.

**Q: How do you prevent the BERT model from drifting in production?**
A: SageMaker Model Monitor tracks feature drift and prediction drift on the SageMaker Endpoint. When drift exceeds threshold (e.g., accuracy drops >5% vs baseline), an EventBridge alarm triggers and the ML Engineer is notified. Manual retraining and re-validation against accuracy gates is required before deploying updated model.

### Security

**Q: What is your threat model for the browser extension?**
A: (1) Extension code is reviewed before Chrome Web Store submission. (2) Service Worker runs Presidio WASM locally — no content sent to our servers. (3) Tenant policy configuration fetched over HTTPS from our API with JWT auth. (4) Extension permissions follow least-privilege: only `tabs`, `activeTab`, `storage`, `scripting` — no broad `<all_urls>` access except on explicitly whitelisted AI tool domains. (5) Content Security Policy prevents XSS in extension pages.

**Q: How do you protect customer API secrets (Google, M365 service account keys)?**
A: All secrets stored in AWS Secrets Manager with KMS encryption (CMK, customer-managed, rotated annually). ECS task IAM roles have read-only access to specific secret ARNs — no wildcard access. Secrets never logged, never in environment variables (retrieved via Secrets Manager API at startup). Secret rotation automated via Secrets Manager rotation Lambda.

**Q: Is customer data used to train your ML models?**
A: No. This is a hard architectural guarantee, not a policy: ML models are trained on public datasets (BFCL for prompt injection, FaceForensics++ for deepfake) and synthetic data. Customer data flows through inference endpoints but is never persisted for training. SageMaker training jobs use only data from the `ml-training-datasets` S3 bucket, which has no cross-account access from production. This commitment is in our DPA.

### Scalability

**Q: What is the bottleneck at 500 customers?**
A: Integration Sync service is the likely bottleneck: 500 tenants × 4 providers = 2,000 sync connections. Current design: dedicated sync worker goroutine per tenant. At 500 tenants, we'd need to migrate to a task-queue model (SQS → sync worker pool) to prevent goroutine explosion. This architectural change is planned for v2 (after v1 at 50 customers is validated).

**Q: Does your PostgreSQL RLS approach scale?**
A: RLS adds a negligible overhead (~1-3ms per query in benchmarks). At 500 tenants with 10,000 assets each (5M rows), PostgreSQL with proper indexing on `(tenant_id, asset_type)` sustains >10,000 queries/second. Vertical scaling of RDS (r6g.2xlarge) handles 500 tenants comfortably. Horizontal read replicas for reporting queries. Physical DB per tenant becomes necessary only above ~10,000 tenants.

**Q: Why not use a multi-cloud strategy for resilience?**
A: For v1, single AWS provides: (1) simpler compliance scope (all certifications in one vendor), (2) native service integrations with no glue code, (3) single IAM model (no cross-cloud credential management). Cloudflare R2 is the only non-AWS service, used for cost (zero egress) not resilience. Multi-cloud resilience is a v2+ consideration after product-market fit is established.

---

## Quick Reference: Tech Stack Summary

| Layer | Technology | Cost Model |
|-------|------------|------------|
| Web Frontend | React / Next.js | Engineering cost |
| Mobile | Flutter (iOS + Android) | Engineering cost |
| Browser Extension | Chrome MV3 + Edge | Engineering cost |
| Backend (Track 1) | Go + Echo | Engineering cost |
| Backend (Track 2) | Python + FastAPI | Engineering cost |
| Auth | Keycloak (ECS Fargate) | ~$50/mo compute |
| Database | PostgreSQL + RLS (RDS Multi-AZ) | ~$200-400/mo |
| Cache | ElastiCache Redis | ~$50-100/mo |
| Event Bus | AWS EventBridge | Near-zero ($1.00/M events) |
| Workflows | AWS Step Functions | ~$25/mo at SME scale |
| ML Platform | AWS SageMaker | ~$200-500/mo |
| Deepfake Detection | Hive Moderation API | ~$0.01/check |
| Prompt Injection | Lakera Guard API | ~$0.001/request |
| Compliance | Vanta | $4,800-6,000/yr |
| Audit Storage | S3 Object Lock | ~$10-50/mo |
| Container Runtime | ECS Fargate | ~$1,000-2,000/mo |
| CDN + WAF | CloudFront + WAF | ~$50-100/mo |
| Backup Storage | Cloudflare R2 | ~$50/mo |
| Observability | CloudWatch + GuardDuty + Security Hub | Included/minimal |
| Local DLP | Presidio WASM | Free (open-source) |
| **Total Infrastructure (v1, 50 customers)** | | **~$3,700/mo** |
