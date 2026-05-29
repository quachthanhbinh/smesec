# SMESec Platform — Team Scope of Work

**Date:** 2026-05-29  
**Status:** Approved  
**Version:** 1.1  
**Related:** [04-delivery-plan-original.md](04-delivery-plan-original.md) · [06-delivery-plan-adjusted-2x.md](06-delivery-plan-adjusted-2x.md) · [07-delivery-plan-realistic-hiring.md](07-delivery-plan-realistic-hiring.md) · [13-metrics-scorecard.md](13-metrics-scorecard.md) · [02-design-document.md](02-design-document.md)

---

## ⚠️ Timeline Context


This document describes team roles and responsibilities. **Timeline references are based on the original 12-month plan.**

For adjusted timelines, see:
- **[2x Adjusted Plan](06-delivery-plan-adjusted-2x.md)** — 26 months (multiply all month references by ~2x)
- **[Realistic Hiring Plan](07-delivery-plan-realistic-hiring.md)** — 36+ months with progressive team build-up

**Key difference in Realistic Hiring Plan:** ML Engineer #1 joins Month 8 (not Day 1), causing an 8-month delay in Track 2 development. All onboarding dates, critical gates, and cross-document dependencies are aligned with the latest delivery plans and integration principles.

---

## Table of Contents

1. [Team Overview](#1-team-overview)
2. [Tech Lead / Architect](#2-tech-lead--architect)
3. [Backend Engineer #1 — Go (Track 1 Core)](#3-backend-engineer-1--go-track-1-core)
4. [Backend Engineer #2 — Go/Python (Track 1 Integration)](#4-backend-engineer-2--gopython-track-1-integration)
5. [Frontend Engineer #1 — React (Web Platform)](#5-frontend-engineer-1--react-web-platform)
6. [Flutter / Mobile Engineer](#6-flutter--mobile-engineer)
7. [ML Engineer #1 — Track 2 Lead](#7-ml-engineer-1--track-2-lead)
8. [ML Engineer #2 — Track 2 (Optional, Month 8)](#8-ml-engineer-2--track-2-optional-month-8)
9. [Backend Engineer #3 — Python/FastAPI (Track 2 API)](#9-backend-engineer-3--pythonfastapi-track-2-api)
10. [Frontend Engineer #2 — Browser Extension](#10-frontend-engineer-2--browser-extension)
11. [DevSecOps Engineer](#11-devsecops-engineer)
12. [Customer Success Engineer](#12-customer-success-engineer)
13. [PM (Part-time 0.5 FTE)](#13-pm-part-time-05-fte)
14. [BD Consultant (Contract, 3 days/week)](#14-bd-consultant-contract-3-daysweek)
15. [Compliance Consultant (Contract, Month 10–12)](#15-compliance-consultant-contract-month-1012)
16. [Cross-Role Responsibilities](#16-cross-role-responsibilities)
17. [Phase-by-Phase Ownership Summary](#17-phase-by-phase-ownership-summary)

---

## 1. Team Overview

```
PHASE 1 (Month 1–3)         PHASE 2 (Month 4–6)         PHASE 3–4 (Month 7–12)
────────────────────         ────────────────────         ────────────────────────
Tech Lead                    + Backend Eng #3             + Customer Success Eng
Backend Eng #1               + Frontend Eng #2            + ML Eng #2 (opt.)
Backend Eng #2                                            DevSecOps → FTE
Frontend Eng #1
Flutter / Mobile Eng
ML Engineer #1               ─────────────────────────────────────────────────────
DevSecOps (contract 0.5)     DevSecOps (contract 0.5)    DevSecOps (FTE 1.0)
PM (0.5)                     PM (0.5)                    PM (0.5)
BD Consultant (contract)     BD Consultant (contract)     BD Consultant (contract)
                                                          Compliance Consultant (M10–12)
```

| Role | Track | Onboard | FTE | Hard Deadline? |
|---|---|---|---|---|
| Tech Lead / Architect | Shared | Week 1 | 1.0 | — |
| Backend Eng #1 | Track 1 | Week 1 | 1.0 | — |
| Backend Eng #2 | Track 1 | Week 1 | 1.0 | — |
| Frontend Eng #1 | Track 1 | Week 1 | 1.0 | — |
| Flutter / Mobile Eng | Track 1 | Week 1 | 1.0 | — |
| ML Engineer #1 | Track 2 | **Day 1, Week 1** | 1.0 | ⚠️ Day 1 hard requirement |
| DevSecOps | Shared | Week 1 (contract 0.5) → FTE Month 7 | 0.5 → 1.0 | — |
| PM | Shared | Week 1 | 0.5 | — |
| BD Consultant | Business | **Week 1** | Contract | ⚠️ Week 1 — not after launch |
| Backend Eng #3 | Track 2 | Month 4 | 1.0 | Required before S7 |
| Frontend Eng #2 | Track 2 | Month 4.5 | 1.0 | Required before S8 |
| Customer Success Eng | Customer | Month 7 | 1.0 | Required at v1 launch |
| ML Engineer #2 | Track 2 | Month 8 (optional) | 1.0 | Depends on v1 velocity |
| Compliance Consultant | Compliance | Month 10 | Contract | Required before S21 |

---

## 2. Tech Lead / Architect

**Onboard:** Week 1  
**Track:** Shared (owns both Track 1 and Track 2 architecture decisions)  
**The single hardest role to replace.** Loss of Tech Lead mid-project = 4–6 week recovery minimum.

### Primary Ownership

| Domain | Responsibility |
|---|---|
| **System architecture** | All architectural decisions — Clean Architecture layers, domain model, API contracts, database schema design |
| **T1-T2 integration contract** | `ThreatDetectionEvent` schema (owns the joint T1+T2 contract, finalized S10) — this schema gates the S11 integration sprint |
| **Multi-tenancy model** | RLS policy design, `tenant_id`/`data_residency` enforcement, cross-tenant isolation CI |
| **AWS infrastructure design** | VPC topology, ECS Fargate task definitions, RDS Multi-AZ design, EventBridge routing, Step Functions workflow design |
| **Security architecture** | Keycloak OIDC/SAML configuration, KMS envelope encryption design, RLS policy, OWASP Top 10 compliance |
| **Code review authority** | Final approval on all PRs touching domain layer, schema migrations, auth, or multi-tenancy |
| **Tech debt governance** | Owns the 20% sprint capacity allocation for tech debt — decides what gets paid down and when |

### Phase-by-Phase Focus

**Phase 1 (S1–S6 / Month 1–3):**
- Design and validate entire infrastructure stack (AWS VPC, ECS, RDS, Keycloak)
- Finalize `ThreatDetectionEvent` schema v0.1 jointly with ML Eng #1 (Week 1)
- Resolve: Keycloak self-hosted vs Auth0/Cognito (must be Day 1 decision — no delay)
- Google rate limit strategy: per-cluster GCP service account layout (R-C2)
- Unblock BE1 and BE2 when domain decisions are unclear

**Phase 2 (S7–S13 / Month 4–6):**
- **Full-time on S11 T1-T2 integration** — this is flagged as highest technical risk in the entire plan; Tech Lead must not be context-switched during this sprint
- `ThreatDetectionEvent` schema v1 final approval (end S10)
- Pentest findings: own all infrastructure-level Critical/High remediation (S12)

**Phase 3–4 (S14–S26 / Month 7–12):**
- Stream A technical lead — pre-planned roadmap features
- Enterprise tier architecture: multi-tenant enterprise model, SIEM webhook design
- SOC 2 Type 2 + ISO 27001 evidence architecture (control mappings for infrastructure controls)
- BERT ensemble design (S23–24) if triggered

### What This Role Does NOT Own
- Sprint planning (PM owns)
- Customer demos (BD Consultant + PM)
- BERT model training (ML Eng #1 owns)
- Browser extension implementation (FE2 owns)

### Hiring Profile
- 5+ years production Go or equivalent backend + cloud architecture experience
- Deep AWS fluency (ECS Fargate, RDS, EventBridge, Step Functions, IAM, KMS)
- Has shipped a multi-tenant SaaS product before — not just designed one
- Security mindset: OWASP, RLS, encryption-at-rest are defaults, not afterthoughts

---

## 3. Backend Engineer #1 — Go (Track 1 Core)

**Onboard:** Week 1  
**Track:** Track 1  
**Core focus:** The engine that powers asset inventory, offboarding, and integration sync.

### Primary Ownership

| Domain | Responsibility |
|---|---|
| **Automated offboarding workflow** | AWS Step Functions orchestration — the most critical piece of Track 1. Grace period, rollback, idempotency key (R-C1). This is the product's core promise. |
| **Asset inventory engine** | Google Workspace Admin SDK sync (users, OAuth apps, devices, groups), 15-min delta sync, rate limit handling |
| **M365 sync** | Microsoft Graph API integration, delta link, webhook renewal service (`subscription_registry` + EventBridge Scheduler, R-C3) |
| **Cross-provider identity matching** | Canonical email matching, composite identity graph across Google + M365 + Slack + AWS |
| **Incident playbook engine** | AWS Step Functions playbook execution, step logging, audit trail |

### Phase-by-Phase Focus

**Phase 1 (S1–S6):**
- S1: Multi-tenant schema, RLS enforcement, CI test harness
- S2: Google Workspace sync — pagination, rate limits, 15-min incremental delta
- S3: M365 sync + **M365 webhook renewal service** (R-C3 — cannot be skipped, 410 Gone handler + DLQ + polling fallback)
- S4: Asset classification engine, OAuth scope risk scoring, alert pipeline
- S5: Slack Admin API integration + composite identity graph
- S6: **Automated offboarding workflow** — Step Functions with grace period, rollback, idempotency, dry-run confirmation

**Phase 2 (S7–S13):**
- S7: JIT access workflow (request → approval → time-boxed grant → auto-revoke)
- S8: AWS Step Functions playbook engine (generic execution framework for all 5 playbooks)
- S9: Mobile push notifications backend (FCM/APNs integration)
- S10: Compliance dashboard backend — Vanta evidence collection hooks, cross-provider composite risk score
- S11: Track 1 side of T1-T2 integration (EventBridge routing of `ThreatDetectionEvent`)
- S12: SaaS dependency mapping, zombie app detection, pentest Critical/High remediation (code-layer)
- S13: Hardening, performance, launch runbook

**Phase 3–4:**
- Stream A: AWS v1.1 (CloudTrail, S3 audit, IAM deep integration)
- Enterprise tier backend: custom RBAC policies, SIEM webhook integration (Splunk/QRadar)
- Multi-region DR drill (RTO/RPO)

### What This Role Does NOT Own
- Database schema design (Tech Lead owns, BE1 implements)
- ML model training (ML Eng #1)
- Frontend/UI (FE1)

### Hiring Profile
- 3+ years Go in production (not just learning Go)
- AWS Step Functions or equivalent workflow orchestration experience
- Has integrated at least one OAuth-based SaaS API (Google, M365, Slack, etc.)
- Understands rate limiting, idempotency, and retry patterns — these are daily concerns

---

## 4. Backend Engineer #2 — Go/Python (Track 1 Integration)

**Onboard:** Week 1  
**Track:** Track 1 (with Python capability for light Track 2 support in Phase 2+)  
**Core focus:** Integrations, security, GDPR, and auth infrastructure.

### Primary Ownership

| Domain | Responsibility |
|---|---|
| **Keycloak SSO** | OIDC/SAML configuration, JWKS cache (6-hour TTL, serve-stale-on-failure), MFA TOTP enforcement |
| **AWS IAM integration** | IAM inventory (users, roles, policies), RBAC permission diff engine, least-privilege recommendations |
| **Slack integration** | Slack Admin API, tier detection, shadow IT detection, Slack-based deactivation |
| **GDPR data subject requests** | `/api/v1/gdpr/erasure` endpoint, KMS key destruction workflow, per-tenant data export |
| **Access review workflows** | Scheduled access reviews, periodic reminder system |
| **Shadow IT alert pipeline** | New OAuth app detection → alert in <15 min → email + Slack notification |

### Phase-by-Phase Focus

**Phase 1 (S1–S6):**
- S1: Keycloak ECS deployment (2 tasks minimum HA), JWKS cache, MFA TOTP, Secrets Manager integration
- S2: Shadow IT detection rules v1 (high-risk OAuth scopes)
- S3: Dashboard API polish (filter, sort, search), cross-provider risk indicators
- S4: Email + Slack notification system, alert pipeline backend
- S5: AWS IAM inventory + RBAC model + least-privilege recommendations engine
- S6: GDPR automation, audit log API, immutable log design

**Phase 2 (S7–S13):**
- S7: Access reviews scheduling backend
- S8–S9: 5 playbook definitions (Steps, conditions, rollback states)
- S10: ISO 27001 + SOC 2 control mapping backend, Vanta evidence API hooks
- S11: GDPR export automation, compliance reports (PDF)
- S12: Vanta evidence dry run validation

**Phase 3–4:**
- Stream A: Advanced compliance reporting, custom playbook builder backend
- ISO 27001 Statement of Applicability technical inputs
- Peer group anomaly detection backend (insider threat signal)

### What This Role Does NOT Own
- AWS Step Functions orchestration (BE1 owns)
- ML inference (BE3/ML Eng)
- Mobile (Flutter Eng)

### Hiring Profile
- 3+ years Go; Python comfortable (will own Python services in Phase 2+)
- Security engineering experience — OAuth, OIDC, JWKS, KMS are not new concepts
- GDPR or data privacy compliance implementation experience is a strong plus

---

## 5. Frontend Engineer #1 — React (Web Platform)

**Onboard:** Week 1  
**Track:** Track 1  
**Core focus:** The web dashboard — what every customer sees every day.

### Primary Ownership

| Domain | Responsibility |
|---|---|
| **Main dashboard** | Asset inventory view, risk indicators per user/app, cross-provider unified view |
| **RBAC dashboard** | Role assignment display, permission diff visualization, least-privilege recommendations UI |
| **Incident playbook wizard UI** | Step-by-step wizard for non-security staff — clear language, no jargon |
| **Compliance dashboard** | ISO 27001 + SOC 2 control status, evidence link view, compliance %, export report trigger |
| **Offboarding workflow UI** | Grace period countdown, cancel button, dry-run confirmation screen, rollback trigger |
| **JIT access request UI** | Request form, approval workflow UI, time-boxed grant status view |
| **Customer onboarding wizard** | OAuth wizard for Google + M365 setup — must complete in <30 min for non-developer IT admin |

### Phase-by-Phase Focus

**Phase 1 (S1–S6):**
- S1: App shell, routing, auth flow (Keycloak PKCE redirect)
- S2: Google Workspace dashboard — user list, OAuth app list, shadow IT flags
- S3: Unified dashboard (Google + M365), risk indicators, filter/search/sort, CSV export
- S4: Asset classification UI, shadow IT alert feed
- S5: RBAC dashboard, least-privilege recommendations, composite identity graph view
- S6: Offboarding wizard (grace period countdown, cancel, dry-run confirmation), 2 playbook wizard UIs, audit log viewer

**Phase 2 (S7–S13):**
- S7: JIT access request + approval UI
- S8: Playbook wizard (generic engine — all 5 playbooks share same UI component)
- S10: Compliance dashboard (ISO 27001 + SOC 2 control mapping view, evidence links)
- S11: GDPR export UI, audit trail viewer
- S12: SaaS dependency map visualization, zombie app detection UI

**Phase 3–4:**
- Stream A: Custom playbook builder UI (coordinated with BE on Step Functions schema)
- Advanced analytics dashboard (SOC-level insights, peer group anomaly visualization)
- Stream B: Dashboard UX redesign from pilot feedback, custom alert rules UI

### What This Role Does NOT Own
- Browser extension (FE2 owns)
- Mobile (Flutter Eng owns)
- API contracts (Tech Lead defines, FE1 consumes)

### Hiring Profile
- 3+ years React (hooks, context, async patterns)
- Has built data-heavy dashboards — tables, charts, filters with real API data
- Security product UX experience is a plus — understands that security UI must be clear for non-security staff
- No need for ML/Python — purely frontend

---

## 6. Flutter / Mobile Engineer

**Onboard:** Week 1  
**Track:** Track 1  
**Core focus:** iOS + Android app — alerts, offboarding trigger, on-the-go access control.

### Primary Ownership

| Domain | Responsibility |
|---|---|
| **Flutter mobile app** | iOS + Android native builds (single Dart codebase via Flutter) |
| **Keycloak PKCE auth** | Mobile auth flow — PKCE code flow (no client secret), secure token storage |
| **Push notifications** | FCM (Android) + APNs (iOS) — security alert delivery, offboarding trigger notification |
| **Read-only inventory** | Asset inventory browsing, risk score view, user/app lookup |
| **Offboarding trigger** | One-tap offboarding initiation from mobile (with 2-step confirmation) |
| **Incident alert feed** | Real-time security alerts with action buttons (acknowledge, escalate, trigger playbook) |
| **App Store + Play Console** | Manage submission, review process, versioning, crash reporting |

### Phase-by-Phase Focus

**Phase 1 (S1–S6):**
- S1: Flutter project setup, Keycloak PKCE flow, navigation shell
- S4: Push notification skeleton (FCM + APNs), auth flow complete
- S6: MVP mobile: alert feed, offboarding trigger, read-only inventory → TestFlight + Play Console beta

**Phase 2 (S7–S13):**
- S9: Full push notifications (FCM/APNs production, alert from playbooks → mobile notification)
- S12: App Store + Google Play submission (target S12 submit, ~1 week review cycle to meet v1 gate)
- S13: v1 mobile in production (iOS App Store + Google Play)

**Phase 3–4 (Stream B, 40% time):**
- UX improvements from pilot feedback (Stream B)
- Mobile crash fixes
- Onboarding wizard mobile-specific flows
- Track 2 mobile views: Shadow AI alerts on mobile, deepfake OOV notification flow

### What This Role Does NOT Own
- Backend APIs (BE1 owns)
- Web dashboard (FE1 owns)
- Push notification backend integration (BE2 owns FCM/APNs server-side)

### Hiring Profile
- 2+ years Flutter production (shipped at least one app to both App Store + Play Store)
- Keycloak or OAuth PKCE mobile flow experience
- Comfortable with FCM + APNs setup (not just Dart — the infra plumbing too)
- Understanding of mobile security: secure storage (flutter_secure_storage), certificate pinning desirable

---

## 7. ML Engineer #1 — Track 2 Lead

**Onboard:** **Day 1, Week 1 — no exceptions**  
**Track:** Track 2 (lead)  
**The highest-leverage hire in the project after Tech Lead.** Track 2 is the primary product differentiator — no other SME security platform does AI threat detection. Every day this role is unfilled = one day of R&D lost that cannot be recovered.

### Primary Ownership

| Domain | Responsibility |
|---|---|
| **Shadow AI risk scoring** | SageMaker training pipeline — feature vector design (OAuth scopes, vendor DPA, app age), model training, evaluation, deployment |
| **Prompt injection detection** | Lakera Guard API integration (v1), BERT-tiny fine-tuning (v2 Enterprise), accuracy gate evaluation on 30-day holdout |
| **LLM DLP — BERT-tiny ONNX** | Fine-tune BERT-tiny for semantic PII detection (Tier 2), export to ONNX, validate in WASM browser environment |
| **Deepfake defense** | Hive Moderation API integration, OOV verification workflow model inputs, independent accuracy evaluation |
| **Accuracy gate evaluation** | Monthly re-evaluation — 500 labeled samples, TPR/FPR measurement, drift alerting |
| **ML infrastructure** | SageMaker workspace, training pipelines, experiment tracking, model registry, endpoint deployment |
| **Dataset management** | Curate, label, and manage all training/holdout datasets (PromptBench, PII-Bench, internal production samples) |
| **`ThreatDetectionEvent` schema** | Joint co-author with Tech Lead (S1). This schema is the contract between Track 1 and Track 2 — it must be right the first time. |

### Phase-by-Phase Breakdown

**Phase 1 (S1–S6 / Month 1–3) — R&D Sprint**

| Sprint | Work |
|---|---|
| S1 (W1–2) | `ThreatDetectionEvent` schema v0.1 — co-design with Tech Lead · OWASP LLM Top 10 + PromptBench literature review · Dataset collection plan: PromptBench, LLM Attacks repo, PII-Bench · SageMaker workspace setup · Shadow AI tool registry v0.1 (100+ known AI tools) |
| S2 (W3–4) | Dataset labeling: prompt injection test cases + PII benchmark · Baseline evaluation: BERT-tiny + regex vs labeled dataset · Record baseline F1, precision, recall — sets improvement targets · Shadow AI risk scoring rubric design |
| S3 (W5–6) | Prompt injection prototype v0.1: fine-tune BERT-tiny on labeled dataset · Evaluate TPR/FPR vs baseline · Lakera Guard API: account setup, cost baseline, first test calls · **Lakera Guard designated as primary v1 implementation** |
| S4 (W7–8) | Shadow AI risk scoring model v0.1: SageMaker training job with OAuth scope feature vector · Browser extension WASM validate (support FE2 — confirm ONNX model loads in service worker) |
| S5 (W9–10) | **Accuracy Gate 1 (W10):** Lakera Guard API: TPR >85%, FPR <2% on 30-day holdout (independent evaluation) · **Accuracy Gate 2 (W10):** Critical PII >99%, FP <5% on Presidio benchmark · Shadow AI classification >95% on top-100 tools · Hive API account live |
| S6 (W11–12) | LLM DLP end-to-end test on real ChatGPT/Gemini (staging) · Track 2 Phase 1 retrospective: accuracy vs targets, revised plan for S7–S13 |

**Phase 2 (S7–S13 / Month 4–6) — Integration**

| Sprint | Work |
|---|---|
| S7 | Shadow AI governance: connect v1 model to live `oauth_application` table from Track 1 · OAuth risk score model v0.2 training on live data |
| S9 | Shadow AI governance v1: AI tool classification live · Shadow AI attestation workflow inputs · **Accuracy Gate 3 (W18):** >95% AI tool classification on top-200 tool list |
| S10 | **Accuracy Gate 4 (W20):** Hive Moderation API >80% deepfake on labeled test dataset (independent evaluation) · OOV workflow design inputs |
| S11 | T1-T2 integration test: `ThreatDetectionEvent` flow end-to-end validated · Prompt injection v1 (Lakera Guard) production integration |

**Phase 3 (S14–S20 / Month 7–9) — Production Scale**

- LLM DLP v2: real-time semantic DLP (beyond regex — BERT-tiny ONNX for novel PII patterns)
- Deepfake v2: Hive API + OOV production combined effectiveness measurement
- ML Eng #2 onboarded (M8): onboarding, task handoff, collaborative model improvement
- SageMaker model monitoring: drift detection active, monthly re-evaluation protocol running
- v1.5 accuracy gate: AI detection combined accuracy >90%

**Phase 4 (S21–S26 / Month 10–12) — Enterprise ML**

- **BERT fine-tuning for Enterprise prompt injection** (S23–24): triggered only if (a) Lakera Guard cost prohibitive at Enterprise volume AND (b) ≥50K labeled production samples available
- Ensemble approach: Lakera Guard + fine-tuned BERT → target >95% TPR, FPR <2%
- Advanced analytics: SOC-level insights, anomaly pattern feeds
- All Track 2 features graduate from beta: SLA guarantees require monthly accuracy re-evaluation to be running

### Accuracy Roadmap Ownership

| Version | Approach | Target TPR | ML Eng #1 Action |
|---|---|---|---|
| v1 | Lakera Guard API alone | 85–90% | Integrate + independently validate on holdout |
| v1.5 | Lakera Guard + regex pre-filter | ~92–93% | Add OWASP + custom pre-filter rules (no new model needed) |
| v2 Enterprise | Lakera Guard + BERT ensemble | >95% | Only if ≥50K samples; fine-tune BERT-tiny on production data |

### What This Role Does NOT Own
- Browser extension implementation (FE2 owns the Chrome MV3 code)
- ONNX runtime integration in browser (FE2 owns, ML Eng #1 exports the model)
- SageMaker infrastructure setup (Tech Lead/DevSecOps own the infra, ML Eng #1 owns the training jobs)
- Lakera Guard API key management (DevSecOps / Secrets Manager)

### Hiring Profile
- **Must have:** SageMaker or equivalent managed ML platform (Vertex AI, Azure ML) — not a pure research background
- BERT fine-tuning in HuggingFace Trainer (not just calling OpenAI API)
- ONNX export and runtime experience — this is non-negotiable for the browser extension architecture
- Python, scikit-learn, PyTorch or TensorFlow
- Comfortable reading security research papers (OWASP LLM Top 10, PromptBench, adversarial ML)
- **Does NOT need:** LLM research background, PhD, or prior security product experience

---

## 8. ML Engineer #2 — Track 2 (Optional, Month 8)

**Onboard:** Month 8 (start of Phase 3, mid-cycle)  
**Track:** Track 2  
**Trigger:** Hire if v1 velocity is on track AND Track 2 backlog has more work than ML Eng #1 can complete alone in Phase 3.

### Primary Ownership (after onboarding)

Owns parallel workstreams that ML Eng #1 would otherwise sequentially queue:

| Domain | Responsibility |
|---|---|
| **Shadow AI model improvement** | Expand registry from top-200 to top-500 tools, retrain model, evaluate accuracy improvement |
| **LLM DLP semantic accuracy** | Systematic evaluation of BERT-tiny ONNX false negatives — edge cases, multilingual PII, novel patterns |
| **Deepfake adversarial testing** | New sample types (VC/voice clone tools not yet in Hive's training data), measure combined Hive + OOV effectiveness against them |
| **Monthly re-evaluation pipeline** | Automate the 500-sample monthly re-evaluation protocol (currently manual ML Eng #1 task) |
| **BERT fine-tuning preparation** | Dataset curation for Enterprise BERT fine-tuning (S23–24) — labeling, deduplication, quality checks on production samples |

### What This Role Does NOT Own
- Architecture decisions (ML Eng #1 + Tech Lead own)
- Primary Lakera Guard integration (already done by ML Eng #1)
- SageMaker infrastructure design

---

## 9. Backend Engineer #3 — Python/FastAPI (Track 2 API)

**Onboard:** Month 4 (Sprint 7 start)  
**Track:** Track 2  
**Core focus:** The production API layer for all ML-powered features — the bridge between ML Eng models and the frontend/extension.

### Primary Ownership

| Domain | Responsibility |
|---|---|
| **ML inference API** | FastAPI service that wraps SageMaker endpoints — prompt injection scoring, shadow AI classification, deepfake verdict |
| **`ThreatDetectionEvent` producer** | Produces events to EventBridge when AI threats are detected — this is the integration contract with Track 1 |
| **Lakera Guard API client** | Production Go/Python client: rate limiting, retry logic, fallback to WASM BERT, cost tracking |
| **Hive Moderation API client** | Deepfake detection client: audio/video submission, async callback, OOV trigger logic |
| **OOV verification workflow** | Out-of-band verification Step Functions workflow — employee confirmation via SMS/Slack second channel |
| **Shadow AI governance API** | REST endpoints for shadow AI risk scores, attestation workflow, allow-list management |
| **Model serving** | SageMaker endpoint management: health checks, fallback routing, inference latency monitoring |

### Phase-by-Phase Focus

**Phase 2 (S7–S13 / Month 4–6) — Build from scratch**

- S7: Environment setup, codebase walkthrough, SageMaker endpoint consumption
- S7: Shadow AI governance API v1 — connect ML model to live data from Track 1
- S8: LLM DLP browser extension backend (FastAPI endpoint for Tier 2 BERT semantic check)
- S9: Shadow AI attestation API, allow-list management
- S10: Deepfake API client (Hive Moderation), OOV workflow design
- S11: T1-T2 integration — `ThreatDetectionEvent` EventBridge publisher, prompt injection API production

**Phase 3–4 (S14–S26 / Month 7–12) — Scale and harden**

- Stream A: ML inference performance (latency <300ms p95), load testing, SageMaker autoscaling
- Advanced analytics API (aggregated threat signals for SOC-level dashboard)
- BERT inference API for Enterprise tier (S23–24)
- SIEM integration backend: Splunk/QRadar webhook format

### What This Role Does NOT Own
- ML model training (ML Eng #1 trains, BE3 serves)
- SageMaker training pipelines (ML Eng #1 owns)
- Browser extension Chrome code (FE2 owns)
- Track 1 Go services (BE1/BE2 own)

### Hiring Profile
- 3+ years Python in production (FastAPI or Flask, not just scripts)
- AWS SageMaker inference endpoint experience (real-time endpoints, async inference)
- API design fluency — rate limiting, retry, circuit breaker patterns
- EventBridge or similar event-driven architecture experience is a strong plus
- Understanding of ML model serving concepts (not training — just serving)

---

## 10. Frontend Engineer #2 — Browser Extension

**Onboard:** Month 4.5 (Sprint 8 start — must be onboarded before W15)  
**Track:** Track 2  
**Core focus:** Chrome extension for LLM DLP — the most technically unusual frontend in the entire project.

### Primary Ownership

| Domain | Responsibility |
|---|---|
| **Chrome MV3 extension** | Manifest V3 architecture — content scripts, service worker, background messaging |
| **LLM DLP intercept** | Content script intercepts textarea submit events on AI platforms (ChatGPT, Gemini, Claude, Copilot, etc.) |
| **Presidio WASM (Tier 1)** | OWASP regex patterns + WASM module loaded in service worker — runs locally, no API call |
| **BERT-tiny ONNX (Tier 2)** | 17MB ONNX model loaded in service worker via onnxruntime-web — lazy-loaded, ~50–80ms inference |
| **Fallback logic** | If Lakera Guard API unavailable → WASM BERT fallback (no user-visible degradation) |
| **Tenant-scoped allow-list** | Per-tenant allow-list managed by admin dashboard (FE1) — extension enforces it locally |
| **Block UI** | User-facing block notification: clear message, override with justification, audit log entry |
| **Chrome Web Store** | Own the publish pipeline — manifest, privacy policy, Web Store listing, review cycle |
| **CI/CD for extension** | GitHub Actions build pipeline for extension packaging, version bumping, Web Store upload |

### Phase-by-Phase Focus

**Phase 2 (S7–S13 / Month 4–6):**
- S8: Extension scaffold (Chrome MV3), CI/CD, Presidio WASM Tier 1 active, first PII block in dev Chrome
- S9: Tier 2 BERT-tiny ONNX integrated in service worker, allow-list enforcement, block UI
- S10: Tenant scoping (per-tenant allow-list from BE3 API)
- S11: Lakera Guard fallback logic (WASM BERT activates when Lakera API unavailable)
- S12: Real-world testing on chatgpt.com, gemini.google.com, claude.ai, GitHub Copilot

**Phase 3 (S14–S20 / Month 7–9):**
- S14–S15: Chrome Web Store submission (production — not sideloaded)
- v1.5 gate: Chrome Web Store published and approved
- LLM DLP v2: real-time semantic detection improvements from ML Eng #1 model updates
- New AI platform coverage (Grok, Perplexity, custom enterprise LLMs)

**Phase 4 (S21–S26 / Month 10–12):**
- Stream B: Bug fixes, UX polish from customer feedback
- Enterprise tier extension features: custom policy enforcement, admin override audit

### Key Technical Constraints

| Constraint | Impact |
|---|---|
| **Chrome MV3 service worker lifecycle** | Service worker can be terminated between events — WASM model must reload. Must validate this in S4 before committing to architecture. |
| **ONNX model size (17MB)** | Lazy-loaded — first PII check after cold start will be slow (~200ms). Cache strategy required. |
| **Content Security Policy** | Chrome MV3 CSP restrictions may block certain WASM loading patterns — validate early |
| **Cross-origin isolation** | COOP/COEP headers needed for SharedArrayBuffer (ONNX WASM threading) — test against real LLM sites |

### What This Role Does NOT Own
- BERT-tiny model training (ML Eng #1 trains, FE2 exports and runs inference)
- Backend API for Tier 2 semantic check (BE3 owns the server-side fallback path)
- Web dashboard (FE1 owns)

### Hiring Profile
- 2+ years Chrome extension development (MV3 specifically — MV2 experience does not transfer cleanly)
- WebAssembly / ONNX runtime in browser experience — this is a hard requirement, not a nice-to-have
- Comfortable with Chrome extension CI/CD and Web Store submission process
- JavaScript/TypeScript, no framework needed (vanilla JS in content scripts is fine)
- No backend required — purely browser-side

---

## 11. DevSecOps Engineer

**Onboard:** Week 1 (contract 0.5 FTE) → Month 7 FTE (1.0 FTE)  
**Track:** Shared  
**Core focus:** The project cannot ship SOC 2 or ISO 27001 without this role. Security is not bolt-on.

### Primary Ownership

| Domain | Responsibility |
|---|---|
| **CI/CD pipeline** | GitHub Actions: build, test, SAST scan, Docker build, ECS deploy — every PR |
| **SAST / DAST** | Static analysis (Semgrep, GoSec), dynamic analysis, dependency scanning (Snyk/Dependabot) — zero Critical/High open policy |
| **Tenant isolation CI test** | Automated cross-tenant RLS breach test — runs on every PR, failures block merge |
| **AWS security posture** | IAM least-privilege for all ECS tasks, VPC security groups, WAF rules, Shield config |
| **Secrets management** | Secrets Manager rotation (90-day), no plaintext secrets in env vars, scan for secrets in code |
| **Vanta evidence collection** | Configure Vanta integrations for SOC 2 + ISO 27001 evidence collection — owns 100% evidence coverage |
| **Pentest coordination** | LOI signed before W14 · Pentest W21–W23 · Remediation tracking Critical/High before v1 gate |
| **KMS + encryption** | Per-tenant KMS key lifecycle, S3 Object Lock configuration, encryption-at-rest verification |
| **SageMaker security** | ML model artifact encryption, SageMaker endpoint IAM roles, VPC endpoint for SageMaker |

### Phase-by-Phase Focus

**Phase 1 (S1–S6, contract 0.5 FTE):**
- S1: CI/CD pipeline green from Day 1, tenant isolation test running, S3 Object Lock + KMS configured
- S2–S5: SAST gate in CI (block Critical/High), WAF configured, secrets scan active
- S6: Pre-MVP security checklist — zero Critical/High in SAST, RLS verified, Secrets Manager all secrets rotated

**Phase 2 (S7–S13, contract 0.5 FTE):**
- S7: Vanta account provisioned — evidence collection starts immediately (this is non-negotiable for SOC 2 Type 2)
- S11: Pentest vendor engaged (LOI signed by W14)
- S12: Pentest Critical/High remediation (owns the security sprint work)

**Phase 3–4 (Month 7+, FTE 1.0):**
- Become FTE — full focus on SOC 2 Type 2 evidence coverage (target >95% in Vanta by S20)
- ISO 27001 evidence preparation (Statement of Applicability technical controls)
- SageMaker model security: adversarial input rate limiting, model artifact integrity
- Multi-region DR: RTO/RPO testing, failover drill (v2 gate requirement)
- OWASP DAST against production environment

### What This Role Does NOT Own
- Application security design (Tech Lead owns architecture, DevSecOps enforces CI gates)
- Compliance framework selection (already decided: Vanta for SOC 2 + ISO 27001)
- BERT model security (ML Eng #1 owns adversarial robustness)

### Hiring Profile
- 3+ years DevSecOps or Platform Engineering with security focus
- AWS security deep expertise: IAM, VPC, WAF, Shield, CloudTrail, GuardDuty
- CI/CD security tooling: Semgrep, Snyk, Dependabot, SAST/DAST pipelines
- SOC 2 or ISO 27001 evidence collection experience (Vanta or Drata preferred)
- Go and Python familiarity for reading code in security reviews

---

## 12. Customer Success Engineer

**Onboard:** Month 7 (start of Phase 3 — at v1 launch)  
**Track:** Customer  
**Core focus:** Owns retention. The product's long-term survival depends on this role as much as engineering.

### Primary Ownership

| Domain | Responsibility |
|---|---|
| **Onboarding** | Customer onboarding from pilot → paying: walkthrough, OAuth wizard support, first-insight validation |
| **Retention** | Weekly customer health check, proactive outreach when health score drops |
| **Stream B triage** | Monday triage of customer-reported issues — classify as P1 bug / UX polish / feature request |
| **Bug queue** | Own Stream B work (with BE3, FE2, Flutter 40%) — fix and polish, not new features |
| **Customer feedback loop** | Synthesize customer feedback into Product backlog — weekly summary to PM |
| **Support SLA** | P1 <24h, P2 <5 days — owns the SLA |
| **NPS** | Run monthly NPS collection, own response to detractors |
| **Churn prevention** | Identify at-risk accounts (health score <70), escalate to PM + BD Consultant |

### Phase-by-Phase Focus

**Phase 3 (S14–S20 / Month 7–9):**
- Week 1: Full onboarding audit of all existing pilot customers — identify top friction points
- Establish Stream B triage process with BE3 + FE2
- Own customer feedback synthesis — top-10 issues, top-5 feature requests, by W32
- Dashboard UX redesign inputs (Stream B S16–S17): customer-reported friction → FE1 changes
- Pilot → paid conversion support: billing flow walkthrough, pricing tier explanation

**Phase 4 (S21–S26 / Month 10–12):**
- Enterprise tier onboarding design — custom pricing, dedicated CSM process
- SOC 2 Type 2 customer-facing communications (what the certificate means for customers)
- All Track 2 features graduating from beta: communicate changes to customers, update documentation

### What This Role Does NOT Own
- Engineering implementation (routes bugs to Stream B team)
- Pricing decisions (PM + CPO)
- Sales pipeline (BD Consultant owns)

### Hiring Profile
- 3+ years in SaaS customer success or technical account management
- Technical enough to read logs, explain API errors, and triage bugs — not a pure relationship role
- Security product experience is a strong plus (customers are IT admins, not executives)
- Has managed churn prevention at a SaaS company with <$1K ACV — knows the playbook

---

## 13. PM (Part-time 0.5 FTE)

**Onboard:** Week 1  
**Track:** Shared  
**Core focus:** Keeps the project running without blocking engineers. A bad PM at 0.5 FTE is worse than no PM.

### Primary Ownership

| Domain | Responsibility |
|---|---|
| **Sprint planning** | 2-week sprint cadence, backlog grooming, story sizing with Tech Lead |
| **Gate checklists** | Owns the MVP / v1 / v1.5 / v2 gate checklists — tracks status, escalates blockers |
| **Pilot funnel** | Works with BD Consultant — tracks outreach → qualified → demo → pilot → paid |
| **Risk register** | Maintains R-C1 through R-C6 risk status, escalates when risks become issues |
| **External dependencies** | Pentest LOI (W14), Vanta account (W13), Chrome Web Store (review timeline), auditor engagement |
| **Stream B coordination** | Monday triage with Customer Success Eng — decides Stream B vs backlog |
| **Metrics reporting** | Weekly: sprint completion rate, burn rate vs plan, accuracy gate status. Monthly: full Tier 1–5 metrics review |

### Critical PM Hard Deadlines

| Deadline | Week | If Missed |
|---|---|---|
| BD Consultant onboarded | W1 | Pilot pipeline is empty at MVP |
| Pentest vendor LOI signed | W14 (end S7) | Pentest cannot start W21 → v1 gate at risk |
| Vanta account provisioned | W13 (start S7) | SOC 2 Type 2 evidence collection delayed → v2 certification slips |
| 3+ pilot customers on staging | W12 (end S6) | MVP has no customers → launch without real feedback |
| 5+ paying customers by v1 | W26 | ARR $0 at v1 → investor / runway concern |
| 10+ paying customers by v1.5 | W38 | Revenue growth below model |

---

## 14. BD Consultant (Contract, 3 days/week)

**Onboard:** Week 1 — not after product is ready  
**Track:** Business  
**Rate:** $60–80/hr, 3 days/week  
**Core focus:** Build the pilot pipeline before there is a product to show. No pilot = no feedback = no v1.

### Primary Ownership

| Domain | Responsibility |
|---|---|
| **Pilot pipeline** | 100 outreach → 30 qualified → 15 demo → 5 pilot by W12. This is the target, not a stretch goal. |
| **MSP partner program** | Sign 3 MSP/IT consultant firm agreements by M6 (CAC via MSP: $500–$800 vs $3,000–$5,000 direct) |
| **Freemium funnel** | "Security Health Check" free tier (5 users, read-only, 14 days) → conversion funnel setup |
| **Demo delivery** | Customer-facing demos alongside PM — own the business narrative, PM owns technical depth |
| **Pricing validation** | Test pricing tiers ($399/$799/$1,499) with prospects — report objections to PM weekly |
| **ICP definition** | Refine Ideal Customer Profile: company size (50–200 employees), industry (professional services, fintech, healthtech most likely), IT admin decision-maker |

### Funnel Targets by Phase

| Phase | Target | Metric |
|---|---|---|
| M1–M3 | Pipeline building | 30+ qualified leads, 5+ demos |
| M3 (MVP) | Pilots live | 3–5 pilot customers on staging |
| M4–M6 (v1) | Paid conversion | 5+ paying customers |
| M7–M9 (v1.5) | Scale | 10+ paying customers |

---

## 15. Compliance Consultant (Contract, Month 10–12)

**Onboard:** Month 10 (start of Phase 4)  
**Track:** Compliance  
**Core focus:** SOC 2 Type 2 audit + ISO 27001 Stage 2 — the credentials that unlock Enterprise sales.

### Primary Ownership

| Domain | Responsibility |
|---|---|
| **SOC 2 Type 2 audit** | Own the auditor relationship, evidence package submission, response to auditor queries |
| **ISO 27001 Stage 2** | Statement of Applicability finalization, Stage 2 audit coordination, ISMS documentation review |
| **Vanta evidence review** | Final review of Vanta evidence coverage before audit — identify gaps, assign remediation to DevSecOps |
| **Risk assessment** | Formal ISO 27001 risk assessment documentation (input from Tech Lead for technical risks) |
| **Control descriptions** | Translate technical controls (RLS, KMS, S3 Object Lock) into auditor-readable control descriptions |
| **Audit fieldwork** | Present during W46–W48 SOC 2 auditor fieldwork — answer auditor questions about controls |

### What This Role Does NOT Own
- Vanta configuration (DevSecOps owns)
- Technical control implementation (BE1/BE2/DevSecOps own)
- ISMS policy drafting before Month 10 (Tech Lead + PM own v1 compliance posture)

---

## 16. Cross-Role Responsibilities

Some responsibilities are shared and must be explicitly assigned — the most common source of "I thought someone else was doing it."

| Responsibility | Owner | Reviewers |
|---|---|---|
| Tenant isolation CI test (write + maintain) | BE1 + Tech Lead | DevSecOps (gate) |
| `ThreatDetectionEvent` schema evolution | Tech Lead + ML Eng #1 | BE3 (producer), BE1 (consumer) |
| GDPR erasure endpoint | BE2 | Tech Lead (audit), DevSecOps (secrets) |
| Monthly accuracy re-evaluation (500 samples) | ML Eng #1 | Tech Lead (review results), PM (track in metrics) |
| Vendor API contracts (Lakera, Hive, Vanta) | PM | Tech Lead (technical SLA review) |
| On-call (P1 incidents) | Rotating: Tech Lead → BE1 → BE2 | DevSecOps (security P1) |
| Pentest findings triage | DevSecOps | Tech Lead (architecture issues), BE1/BE2 (code issues) |
| Accuracy gate report (each gate) | ML Eng #1 | Tech Lead, PM (go/no-go decision) |
| AWS Cost Explorer review (weekly) | DevSecOps | PM (budget gate) |

---

## 17. Phase-by-Phase Ownership Summary

### Phase 1 — Foundation → MVP (Month 1–3)

| Who | Focus |
|---|---|
| Tech Lead | Infrastructure architecture, schema design, T1-T2 contract, unblock BE1/BE2 |
| BE1 | Offboarding workflow, Google/M365/Slack/AWS sync, Step Functions |
| BE2 | Keycloak, security, GDPR, alert pipeline, access reviews |
| FE1 | Dashboard, onboarding wizard, playbook wizard, offboarding UI |
| Flutter | Mobile scaffold, auth, push notification skeleton |
| ML Eng #1 | Dataset collection, baseline models, Lakera integration, prototype prompt injection |
| DevSecOps | CI/CD, SAST, tenant isolation test, S3 Object Lock, KMS |
| PM | Sprint planning, pilot outreach coordination, gate tracking |
| BD Consultant | Pilot pipeline (100 outreach → 5 pilots by W12) |

### Phase 2 — MVP → v1 (Month 4–6)

| Who | Focus |
|---|---|
| Tech Lead | S11 T1-T2 integration (full-time), pentest architecture remediation |
| BE1 | JIT access, playbook engine, mobile notifications, compliance backend |
| BE2 | Access reviews, GDPR automation, Vanta evidence hooks, compliance reports |
| FE1 | JIT UI, playbook wizard, compliance dashboard, GDPR export UI |
| Flutter | Push notifications production, App Store submission |
| ML Eng #1 | Shadow AI governance live, deepfake gate evaluation, T1-T2 integration validation |
| BE3 | Shadow AI API, LLM DLP backend, EventBridge producer, Hive API client |
| FE2 | Browser extension (WASM Tier 1 + ONNX Tier 2), Chrome CI/CD |
| DevSecOps | Vanta evidence collection, pentest coordination, remediation |
| PM | Pentest LOI (W14), Vanta (W13), 5+ pilots → paying customers |
| BD Consultant | Pilot → paid conversion, MSP partner program |

### Phase 3–4 — v1 → v2 (Month 7–12)

| Who | Focus (Stream A) | Focus (Stream B) |
|---|---|---|
| Tech Lead | AWS v1.1, Enterprise architecture, SOC 2 Type 2 prep | — |
| BE1 | AWS CloudTrail, SIEM webhook, Enterprise RBAC | (escalation support) |
| BE2 | Advanced compliance, ISO 27001 controls, peer anomaly | — |
| FE1 | Advanced analytics dashboard, custom playbook builder | Dashboard UX redesign |
| Flutter | Track 2 mobile views | Crash fixes, UX polish |
| ML Eng #1 | LLM DLP v2, deepfake v2, drift monitoring, BERT prep | — |
| ML Eng #2 | Shadow AI expansion, evaluation pipeline automation | — |
| BE3 | ML inference performance, SIEM format, BERT serving | Bug fixes, API polish |
| FE2 | Web Store publish, new LLM platform coverage | Extension bug fixes |
| DevSecOps | SOC 2 Type 2 evidence (>95% Vanta), ISO 27001, DR drill | — |
| Customer Success | Onboarding, retention, churn prevention | Stream B triage (Monday) |
| PM | Gate checklists, ARR tracking, hiring plan | — |
| Compliance Consultant (M10–12) | SOC 2 Type 2 audit, ISO 27001 Stage 2 | — |

---

*This document is a living reference. Update when roles change or new hires join. Last updated: 2026-05-28.*
