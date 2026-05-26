# SME AI Security Platform - Decision Record

Date: 2026-05-26  
Status: Approved  
Stakeholders: Client (strategic proposal, sole-source)

---

## Executive Summary

This document records the complete decision-making process for designing the SME AI Security Platform, from initial requirements through architectural choices to final tech stack selection. It captures the brainstorming dialogue, trade-off analysis, and rationale behind key decisions.

---

## 1. Project Context & Requirements

### 1.1 Client Requirements (Original Brief)

**Problem Statement:**
SMEs (10-500 employees) face AI-driven risks but lack dedicated security teams and large budgets:
- Automated spear-phishing, AI-generated disinformation
- Data leakage to public LLMs
- Shadow AI tool adoption by employees
- Supply-chain compromise via AI-powered attacks

**Key Requirements:**
- Asset inventory and classification (data, devices, accounts, third-party integrations)
- AI-specific threat surface: prompt injection, LLM data leakage, deepfake fraud
- Access governance: least-privilege, offboarding automation, shadow IT detection
- Continuous compliance (ISO 27001, GDPR, SOC 2 lite)
- Incident playbooks executable by non-security staff
- Cost model: tiered, pay-as-you-grow pricing
- Integration: Google Workspace, Microsoft 365, Slack, QuickBooks, etc.

**Deliverables Required:**
1. System Architecture Diagram (logical + deployment view)
2. Design Document (max 600 words)
3. Team & Delivery Plan (6 months)
4. AI Governance Module detail

### 1.2 Clarifying Questions & Answers

**Q1: Project type?**
- **A:** Strategic proposal - architecture and plan only (not actual client engagement yet)

**Q2: Competitive context?**
- **A:** Sole-source proposal (direct engagement, not competitive RFP)

**Q3: Client's current security posture?**
- **A:** Greenfield - minimal existing security (basic antivirus, email filtering)

**Q4: Budget for security operations?**
- **A:** $50K/năm - ngân sách rất hạn chế (this is customer subscription price, not build budget)

**Q5: Development team size?**
- **A:** 6-10 người - đội lớn

**Q6: Primary objective?**
- **A:** Toàn diện - cân bằng compliance, threat protection, và AI governance

**Q7: Business model?**
- **A:** Hybrid - linh hoạt theo nhu cầu (SaaS multi-tenant default, on-premise option for larger customers)

**Q8: Target industry?**
- **A:** Cross-industry - tổng quát (not specialized for one vertical)

**Q9: Build budget clarification?**
- **A:** Customer pays $50K/year subscription; build budget is $200K-$500K for 6-month development

**Q10: Target industries?**
- **A:** Cross-industry (tech, professional services, retail, manufacturing - general solution)

**Q11: Deployment preference?**
- **A:** AWS-native deployment for all infrastructure

**Q12: Client applications?**
- **A:** Flutter/Dart for mobile + desktop apps (cross-platform)

**Q13: Riskiest assumption?**
- **A:** Kỹ thuật - AI threat detection accuracy (can we build accurate detection within budget constraints?)

---

## 2. Architectural Options Analysis

### Option 1: Orchestration Platform (Build Light, Integrate Heavy)

**Strategy:** Build thin orchestration layer, maximize use of existing SaaS/OSS tools.

**Build:**
- Unified dashboard & control plane
- AI threat detection engine (prompt injection, LLM leakage detection)
- Policy orchestration engine
- Asset inventory & classification module

**Integrate:**
- Wazuh (OSS SIEM) for log aggregation
- Keycloak (OSS) for identity & access management
- Vanta/Drata API for compliance automation
- Google Workspace/M365 APIs for shadow IT detection
- OpenCTI (OSS) for threat intelligence

**Pros:**
- ✅ Fastest time-to-market (4-5 months for v1)
- ✅ Lowest cost - leverage OSS and SaaS tier thấp
- ✅ Lowest risk - proven solutions
- ✅ Easy to maintain with small team

**Cons:**
- ❌ High vendor dependency (lock-in risk)
- ❌ Hard to customize deeply
- ❌ Lower margin (must pay third-party licenses)
- ❌ Weak differentiation - easy for competitors to copy

**Economics:**
- Development: $250K-$300K
- Third-party tools: ~$10K-$15K/year per customer
- Gross margin: ~60-65%

---

### Option 2: Custom Platform (Build Heavy)

**Strategy:** Build everything from scratch for full IP ownership.

**Build Everything:**
- Custom SIEM engine
- Identity & access management system
- AI threat detection (ML models, behavioral analysis)
- Compliance engine with rule builder
- Asset discovery & inventory
- Incident response automation
- Integration framework for third-party tools

**Pros:**
- ✅ Full control & IP ownership
- ✅ Highly customizable per customer
- ✅ Strong differentiation
- ✅ Higher margin (no third-party licenses)
- ✅ Strong competitive moat

**Cons:**
- ❌ **HIGHEST RISK** - many complex components
- ❌ Slowest time-to-market (6 months = MVP only)
- ❌ Requires very strong team (security experts, ML engineers)
- ❌ High cost, likely exceeds budget
- ❌ Hard to maintain and scale

**Economics:**
- Development: $450K-$600K (exceeds budget)
- Infrastructure: $5K-$8K/year per customer
- Gross margin: ~80-85% (but high risk)

**Assessment:** **NOT RECOMMENDED** - too risky for budget and timeline.

---

### Option 3: Hybrid Modular Platform ⭐ **SELECTED**

**Strategy:** Build core differentiators, integrate SaaS for commodity functions.

**Build (Core IP):**
- **AI Threat Detection Engine** (main differentiation):
  - Prompt injection detection
  - LLM data leakage monitoring
  - Deepfake detection (voice/video)
  - Shadow AI discovery
- **Unified Asset Inventory & Classification**
- **Policy Engine & Orchestration**
- **Unified Dashboard & Analytics**
- **Incident Playbook Automation**

**Integrate (Commodity):**
- Wazuh (OSS SIEM) for log collection
- Keycloak (OSS) for SSO/IAM
- Compliance-as-a-Service API (Vanta/Drata) for reporting
- Cloud provider APIs (AWS/GCP/Azure) for cloud asset discovery

**Architecture:**
- Microservices architecture
- Plugin system for integrations
- API-first design
- Multi-tenant by default, single-tenant capable

**Pros:**
- ✅ **BEST BALANCE** of speed, cost, and differentiation
- ✅ Focus resources on AI detection (highest risk area)
- ✅ Reasonable time-to-market (5-6 months for v1)
- ✅ Modular = can swap components later
- ✅ Scalable architecture
- ✅ Good margin (~70-75%)

**Cons:**
- ⚠️ More complex architecture design
- ⚠️ Needs skilled architect for module boundaries
- ⚠️ Some vendor dependency remains

**Economics:**
- Development: $350K-$450K (within budget)
- Third-party tools: ~$8K-$12K/year per customer
- Gross margin: ~70-75%

---

## 3. Decision Rationale: Why Option 3?

### 3.1 Risk Mitigation

**Addresses the riskiest assumption first:**
- Focus development resources on AI threat detection engine
- This is both the highest technical risk AND the main differentiation
- Validates feasibility early (Milestone M3 at Week 5-6)

### 3.2 Budget & Timeline Fit

- $350K-$450K fits within $200K-$500K budget
- 5-6 months achievable with 6-10 person team
- Not too aggressive (Option 2) or too conservative (Option 1)

### 3.3 Competitive Position

- AI detection engine = proprietary IP, hard to copy
- Modular architecture allows future expansion
- Can compete on innovation, not just integration

### 3.4 Scalability Path

- Modular design supports future growth
- Can replace integrated components with custom builds later
- Multi-tenant architecture scales economically

### 3.5 Trade-off Acceptance

**Accepted trade-off:** More complex architecture requiring skilled architect.

**Why acceptable:** This complexity is investment in long-term platform health. The alternative (Option 1) has lower upfront complexity but creates technical debt and vendor lock-in. Option 2's complexity is in breadth (too many components); Option 3's complexity is in depth (well-designed boundaries).

---

## 4. Tech Stack Decisions

### 4.1 Monorepo Strategy: Nx

**Decision:** Use Nx monorepo for all services and apps.

**Rationale:**
- Single source of truth for shared types, rules, fixtures
- Unified build/test/deploy pipeline
- Better code reuse across TypeScript services
- Proven at scale (used by Google, Microsoft, etc.)

**Alternatives considered:**
- Turborepo: Less mature, fewer features
- Polyrepo: Too much overhead for 6-10 person team

---

### 4.2 Backend Services

**API Gateway & Web Services:** Node.js + TypeScript + Express/Next.js

**Rationale:**
- Team likely has TypeScript expertise
- Fast development velocity
- Rich ecosystem for web/API development
- Easy to hire for

**AI Threat Service:** Python + FastAPI

**Rationale:**
- Python is standard for ML/AI work
- FastAPI provides async performance
- Easy integration with ML libraries (PyTorch, TensorFlow)
- SageMaker native support

**Policy & Asset Services:** Go

**Rationale:**
- Performance-critical services
- Strong concurrency model
- Low memory footprint
- Good for long-running services

---

### 4.3 Frontend Applications

**Web Admin Dashboard:** React + Next.js + TypeScript

**Rationale:**
- Modern, well-supported stack
- Server-side rendering for performance
- TypeScript for type safety
- Large talent pool

**Mobile + Desktop:** Flutter/Dart

**Rationale:**
- **Single codebase** for iOS, Android, Windows, macOS, Linux
- Native performance
- Consistent UX across platforms
- Lower maintenance cost than separate native apps
- Growing ecosystem and Google backing

**Alternatives considered:**
- React Native: Less mature desktop support
- Native (Swift/Kotlin/C#): 3x development cost
- Electron: Poor performance, large bundle size

---

### 4.4 Infrastructure: AWS-Native

**Core Services:**
- ECS Fargate (container orchestration)
- Aurora PostgreSQL (primary database)
- ElastiCache Redis (caching, sessions)
- S3 (object storage, evidence)
- SQS + EventBridge (async messaging)
- Step Functions (workflow orchestration)
- Lambda (event-driven functions)
- SageMaker (ML model serving)
- CloudWatch + SNS (monitoring, alerting)

**Rationale:**
- Fully managed services = lower ops overhead
- Proven at scale
- Strong security posture
- Pay-as-you-grow pricing
- Team likely has AWS experience

**Alternatives considered:**
- GCP: Less mature enterprise tooling
- Azure: Client base is cross-cloud
- Self-hosted Kubernetes: Too much ops overhead for v1

---

### 4.5 V1 vs V2+ Tech Stack

**V1 (Weeks 1-24): Optimize for speed and simplicity**
- EventBridge + SQS + Lambda (async processing)
- Aurora PostgreSQL only (no graph DB yet)
- Step Functions (workflow orchestration)
- ECS Fargate (container hosting)
- SageMaker endpoints (ML serving)

**V2+ (Post-pilot, scale phase): Optimize for performance and scale**
- Amazon MSK (Kafka) + Managed Flink (stream processing)
- EKS (Kubernetes) for high-density services
- OpenSearch (search + threat hunting)
- DynamoDB (high-throughput state)
- S3 Data Lake + Glue/Athena (analytics)
- Redshift (BI/reporting)
- Multi-region DR (active-passive)
- SageMaker Pipelines + Model Registry (MLOps)

**Rationale for phasing:**
- V1 validates product-market fit with minimal ops complexity
- V2+ adds scale infrastructure only when needed
- Avoids premature optimization
- Keeps v1 team lean (6-10 people)

---

## 5. AI Detection Strategy (Phased Approach)

### 5.1 V1 Scope (Weeks 1-24)

**Prompt Injection Detection:**
- Rule-based patterns (regex)
- Lightweight ML classifier (fine-tuned BERT)
- Real-time scoring (0-100 risk score)

**LLM Data Leakage Monitoring:**
- DLP pattern checks at endpoint/browser
- Detect PII, credentials, IP in prompts
- Alert on sensitive data patterns

**Shadow AI Discovery:**
- DNS/domain detection (ChatGPT, Copilot, etc.)
- OAuth app inventory from Google/M365
- Browser telemetry (optional extension)
- Usage analytics (frequency, data volume)

**Deepfake Detection:**
- **Vendor API integration** (Sensity, Reality Defender)
- Out-of-band callback verification workflow
- Confidence scoring

**Rationale:**
- Focus on **high-impact, achievable** controls
- Avoid building complex ML models from scratch in v1
- Validate accuracy early (Milestone M3: precision ≥85%, FP <15%)
- Defer model-heavy detection to v2 after pilot feedback

### 5.2 V2+ Enhancements (Post-pilot)

- Custom deepfake models (voice + video)
- Behavioral analysis (user baseline + anomaly detection)
- Advanced LLM leakage (network payload inspection)
- Threat intelligence integration
- Automated response actions

---

## 6. Team & Delivery Plan

### 6.1 Team Structure (8 FTE)

- 1 Product Manager / Security Analyst
- 1 Solution Architect / Tech Lead
- 2 Backend Engineers (Go/Python)
- 1 Frontend Engineer (React/Next.js)
- 2 Flutter Engineers (mobile + desktop)
- 1 DevSecOps/QA Engineer

**Rationale:**
- Balanced across backend, frontend, mobile
- Architect ensures module boundaries stay clean
- DevSecOps ensures security from day 1
- 8 people fits $350K-$450K budget (~$44K-$56K per person for 6 months)

### 6.2 6-Month Roadmap

**Month 1:** AWS foundation, tenant model, auth baseline, integration skeletons  
**Month 2:** Asset inventory + classification, policy engine v1, offboarding automation  
**Month 3:** AI governance v1 (shadow AI, prompt guard, DLP pattern controls)  
**Month 4:** Incident playbooks + compliance control mappings  
**Month 5:** Unified dashboard + Flutter mobile/desktop workflows  
**Month 6:** Hardening, pilot with 2-3 SMEs, false-positive tuning, launch readiness

### 6.3 Milestones & Gates

**M1 (Week 1-2):** Engineering foundation ready  
- Gate: All apps boot locally; CI runs lint+unit tests green

**M2 (Week 3-4):** AI risk scoring vertical slice  
- Gate: End-to-end "ingest prompt → score → alert in UI" works

**M3 (Week 5-6):** Accuracy validation gate ⚠️ **CRITICAL**  
- Gate: Precision ≥85% on severe class; false-positive <15% on severe alerts

**M4 (Week 7-8):** AWS dev deployment baseline  
- Gate: Services deploy to dev AWS, traces/logs/alerts working

### 6.4 Riskiest Assumption Validation

**Assumption:** AI detection quality is actionable for SMEs without overwhelming false positives.

**Validation Plan (First 6 weeks):**
1. Build minimal detection engine (M1-M2)
2. Run pilot dataset + limited real traffic
3. Measure precision and false-positive rate
4. **Decision gate at M3:**
   - If metrics pass: proceed to full build
   - If metrics fail: pivot to vendor API-only approach or adjust thresholds

**Success Criteria:**
- Precision (severe alerts) ≥ 85%
- False positive rate (severe) < 15%
- Alert response time < 5 minutes

---

## 7. AI Governance Module Detail

### 7.1 Detection Mechanisms

**Shadow AI Discovery:**
- Domain/API endpoint detection (ChatGPT, Copilot, Gemini, etc.)
- OAuth app inventory from Google Workspace / M365
- Browser/desktop telemetry (opt-in agent)

**Usage Risk Scoring:**
- Data type sent (PII, financial, source code, secrets)
- AI tool used (approved vs unapproved)
- User role + device context

**Policy Enforcement Levels:**
- **Advisory:** Warning + guidance (educate user)
- **Justification required:** User must provide business reason
- **Block:** Hard block on severe violations

### 7.2 Governance Workflow

1. **Approved AI Catalog:** Per-department approved tools list
2. **Prompt/Data Guards:** Templates to prevent sensitive data leakage
3. **Offboarding Automation:** Revoke AI app tokens + sessions on employee exit
4. **Audit Trail:** Full compliance evidence for GDPR/ISO 27001

### 7.3 Privacy Guarantees

- **Metadata-first collection:** Minimize raw prompt storage
- **Data minimization:** Only collect what's needed for detection
- **Encryption:** At rest (KMS) and in transit (TLS)
- **Tenant isolation:** Row-level security + schema separation
- **Retention policies:** Configurable by tier
- **GDPR controls:** Export/delete workflows, redaction

---

## 8. Cost Model & Pricing Strategy

### 8.1 Tiered Pricing

**Starter (10-50 users):**
- Core asset inventory + policy engine
- Basic incident playbooks
- Community support
- Price: ~$15K-$25K/year

**Growth (51-200 users):**
- Full AI governance module
- Compliance automation (ISO/GDPR/SOC2)
- Advanced integrations
- Email support
- Price: ~$35K-$55K/year

**Scale (201-500 users):**
- Dedicated controls
- Extended retention
- Custom playbooks
- Priority support + CSM
- Price: ~$60K-$100K/year

**Hybrid Deployment Surcharge:**
- On-premise/private cloud: +30-50% for deployment complexity

### 8.2 Unit Economics

**Build Cost:** $350K-$450K (one-time)

**Per-Customer Costs (annual):**
- Infrastructure (AWS): $5K-$8K
- Third-party tools: $8K-$12K
- Support/ops: $3K-$5K
- **Total COGS:** ~$16K-$25K per customer

**Gross Margin:** 70-75% at $50K average selling price

**Break-even:** ~15-20 customers to recover build cost in Year 1

---

## 9. Scope Boundaries (Plan A vs B/C/D)

This design and plan cover **Plan A only** - the foundation and AI detection vertical slice.

**Plan A (This Document) - 12 Tasks:**
- Monorepo foundation (Tasks 1-2)
- AI threat detection engine (Task 3)
- API Gateway + Web Dashboard + Flutter App (Tasks 4-6)
- Accuracy validation gate (Task 7)
- AWS deployment baseline (Task 8)
- CI/CD pipeline (Task 9)
- Pilot readiness package (Task 10)
- **Browser Extension - Prompt Interceptor (Task 11)** ← Added after architecture review
- **Browser Extension - DLP Scanner (Task 12)** ← Added after architecture review

**Critical Architecture Decision (2026-05-26):**

After reviewing the design spec, we identified a gap: the backend API cannot collect real prompts without client-side components. We evaluated three approaches:

1. **Add both Browser Extension + Desktop Agent to Plan A** → Too much scope, risks 6-month timeline
2. **Defer all client-side monitoring to Plan B** → Plan A becomes passive dashboard, cannot validate AI accuracy
3. **Hybrid: Browser Extension in Plan A, Desktop Agent in Plan B** ← **SELECTED**

**Rationale for Hybrid Approach:**
- Browser extension covers 90% of AI usage (ChatGPT, Copilot web)
- Simpler than desktop agent (no admin rights, no kernel hooks)
- Critical for M3 accuracy gate (need real prompts to test)
- Desktop agent deferred to Plan B (covers edge cases: desktop apps, clipboard)

**Deferred to Plan B (Post-V1):**
- **Desktop Monitoring Agent:** Clipboard monitoring, desktop app traffic inspection, kernel-level hooks
- **Endpoint DLP Agent:** File operation monitoring, screen capture detection, USB controls
- **Network-Level Inspection:** Corporate proxy/firewall integration, SSL/TLS decryption, DNS blocking
- **MDM Integration:** Mobile device management for BYOD scenarios (Intune, Workspace ONE)
- **Deep asset discovery:** Network scanning, agent-based discovery
- **Advanced policy orchestration:** Complex rule engine, automated offboarding workflows

**Deferred to Plan C:**
- Full compliance engine
- Evidence collection automation
- Audit report generation

**Deferred to Plan D:**
- Incident playbook library
- Production hardening
- Pilot customer onboarding
- SLA/monitoring setup

**BYOD Considerations:**

Personal mobile/tablet devices present unique challenges:
- Cannot install browser extension (mobile browsers don't support extensions)
- Cannot install desktop agent (privacy concerns + no admin rights)
- Cannot intercept network traffic (requires corporate certificate)

**V1 Approach (Plan A):**
- OAuth token monitoring via Google Workspace / M365 APIs (already in Task 3)
- Detect when employees authorize AI apps on any device
- IT can revoke tokens remotely
- Limitation: Reactive (detect after authorization), not proactive (block before)

**V2 Approach (Plan B):**
- MDM integration (Intune, Workspace ONE, MobileIron)
- Enforce work profile policies on BYOD devices
- Conditional access: require MDM enrollment to access work email
- Block unapproved AI apps within work profile
- Preserve personal app privacy (IT cannot access personal data)

**Rationale:**
- Plan A validates the riskiest assumption (AI detection accuracy) with browser extension
- Keeps scope tight for 6-month timeline
- Each plan produces working, testable software
- Avoids "big bang" integration risk
- BYOD monitoring deferred to Plan B based on pilot feedback

---

## 10. Key Decisions Summary

| Decision | Choice | Rationale |
|----------|--------|-----------|
| **Architecture** | Hybrid Modular (Option 3) | Best balance of speed, cost, differentiation |
| **Monorepo** | Nx | Unified tooling, code reuse, proven at scale |
| **Backend** | Go (services) + Python (AI) + Node (gateway) | Right tool for each job |
| **Frontend** | React/Next.js (web) + Flutter (mobile/desktop) | Modern stack, cross-platform efficiency |
| **Infrastructure** | AWS-native | Managed services, lower ops overhead |
| **AI Detection** | Phased (rules + lightweight ML in v1) | Validate accuracy early, defer complexity |
| **Deployment** | Multi-tenant SaaS default, on-prem capable | Economic efficiency + flexibility |
| **Team** | 8 FTE (balanced across stack) | Fits budget, covers all areas |
| **Timeline** | 6 months to pilot-ready v1 | Aggressive but achievable |
| **Riskiest Assumption** | AI detection accuracy | Validate at M3 (Week 5-6) |

---

## 11. Success Criteria

**Technical:**
- AI detection precision ≥85% (severe class)
- False positive rate <15% (severe alerts)
- End-to-end alert latency <5 minutes
- All services deploy to AWS dev
- CI/CD pipeline green

**Business:**
- Pilot with 2-3 SME customers
- Customer feedback validates value proposition
- Pricing model accepted by pilot customers
- Path to 15-20 customers in Year 1 clear

**Team:**
- No critical team member dependencies
- Documentation sufficient for new hires
- Architecture decisions documented (this file)

---

## 12. Open Questions & Risks

### 12.1 Open Questions

1. **Compliance certification timeline:** How long to get ISO 27001/SOC 2 certified?
2. **Pilot customer selection:** Which SMEs are best fit for pilot?
3. **Pricing validation:** Will SMEs actually pay $50K/year?
4. **Integration complexity:** How hard is Google/M365/Slack integration in practice?

### 12.2 Known Risks

| Risk | Impact | Mitigation |
|------|--------|------------|
| AI detection accuracy insufficient | High | Early validation at M3; fallback to vendor APIs |
| Timeline slips | Medium | Strict scope boundaries; defer non-critical features |
| Team skill gaps | Medium | Hire architect early; pair programming |
| Third-party API changes | Low | Abstract integrations behind adapters |
| AWS cost overruns | Low | Monitor spend weekly; use cost alerts |

---

## 13. Next Steps

1. ✅ **Design approved** - This document
2. ✅ **Implementation plan created** - [2026-05-26-sme-ai-security-platform-v1-plan.md](2026-05-26-sme-ai-security-platform-v1-plan.md)
3. ⏳ **Set up worktree** - Isolated workspace for implementation
4. ⏳ **Execute Task 1** - Initialize Nx monorepo
5. ⏳ **Execute Tasks 2-10** - Follow plan sequentially
6. ⏳ **M3 validation gate** - Week 5-6 accuracy check
7. ⏳ **Pilot deployment** - Week 24

---

## Document History

- 2026-05-26: Initial version - full decision record from brainstorming to final design
