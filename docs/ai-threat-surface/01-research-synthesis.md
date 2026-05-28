# AI-Specific Threat Surface — Research Synthesis
**Date:** May 28, 2026  
**Method:** 3-agent iterated research loop (Product Owner → Technical Advisor → Product Owner Round 2 → Project Manager)  
**Scope:** Key Requirement — AI-specific threat surface for SMEs (10–500 employees)  
**Constraint:** Research conducted without bias from pre-existing Track 1 / Track 2 plans

---

## Executive Summary

The SME AI threat surface is a **genuine, large, and entirely underserved market gap**. Every serious AI security vendor today (HiddenLayer, Wiz AI-SPM, Orca, Protect AI) protects organizations that *deploy* LLMs. SMEs are *consumers* of AI — their threat surface is ChatGPT, Microsoft Copilot, Gemini, and Claude. **No vendor has a product for "SME as AI consumer" security.**

SMESec can be the first platform to address this with a unified, non-technical-staff-operable solution covering:
1. LLM data leakage prevention (browser-based, submit-time)
2. Shadow AI governance (OAuth inventory + risk scoring + attestation)
3. Deepfake fraud defense (voice verification + out-of-band verification workflow)
4. AI phishing integration (M365 Defender + Google Workspace posture)
5. Prompt injection detection (rule-based + ML, for Enterprise tier)

---

## 1. Market & Competitive Landscape

### 1.1 Competitor Feature Matrix

#### Shadow AI Discovery & Governance

| Vendor | Shadow AI Discovery | Policy Enforcement | DLP | SME Pricing | SME UI |
|--------|--------------------|--------------------|-----|-------------|--------|
| **Nudge Security** | ✅ OAuth app discovery, AI categorization | ⚠️ Nudge/alert only — no block | ❌ | ~$4–8/user/mo | ✅ Simple |
| **Obsidian Security** | ✅ SaaS posture + identity risk | ✅ Policy-based | ⚠️ Basic | ❌ $40K+/yr | ❌ |
| **DoControl** | ✅ SaaS data access mapping | ✅ Automated remediation | ✅ Strong | ❌ $30K+/yr | ❌ |
| **Reco AI** | ✅ AI-native SaaS security | ✅ Access anomaly enforcement | ⚠️ Partial | ❌ $25K+/yr | ⚠️ |
| **Metomic** | ⚠️ Data in SaaS only | ❌ Alert only | ✅ PII scanning | ~$5K–15K/yr | ✅ |
| **Nightfall AI** | ❌ DLP focus | ❌ DLP only | ✅ Best-in-class cloud DLP | ⚠️ $10K–20K/yr | ⚠️ |

**Finding:** Nudge Security is the only SME-accessible shadow AI discovery tool, but it only nudges — cannot block, enforce policy, or integrate DLP. No competitor combines shadow AI discovery + DLP + remediation at SME price points.

#### LLM / Prompt Injection Security

| Vendor | Prompt Injection | DLP | Browser Extension | SME Fit |
|--------|-----------------|-----|-------------------|---------|
| **Lakera (Gandalf)** | ✅ Excellent — real-time API | ⚠️ Basic | ❌ | ❌ $20K+/yr, dev/enterprise only |
| **Prompt Security** | ✅ Browser ext + API gateway | ✅ PII redaction | ✅ Chrome | ❌ $15K–30K/yr, IT admin setup |
| **Protect AI** | ✅ ML model security + LLM firewall | ✅ AI Red Team | ❌ | ❌ $50K+/yr |
| **Aporia** | ✅ LLM guardrails (injection, toxicity) | ✅ PII/IP detection | ❌ | ❌ $15K+/yr |
| **Rebuff** | ✅ Heuristic + canary injection | ❌ | ❌ | ❌ Open-source, DIY only |
| **LLM Guard** | ✅ PII anonymization | ✅ | ❌ | ❌ Self-hosted DevOps required |

**Finding:** Every enterprise LLM security product requires developers or security engineers to integrate. None offer a turnkey "install and protect" experience for non-technical SME staff.

#### Deepfake Detection (Voice + Video)

| Vendor | Voice | Video | Real-Time | SME Pricing |
|--------|-------|-------|-----------|-------------|
| **Reality Defender** | ✅ | ✅ | ✅ | ❌ $30K+/yr, enterprise only |
| **Pindrop** | ✅ Best-in-class | ❌ Voice only | ✅ Telephony | ❌ $50K+/yr, call center |
| **Hive Moderation** | ✅ | ✅ Video+image | ⚠️ Async | ✅ $0.001–0.01/req, pay-per-use |
| **Resemble Detect** | ✅ | ❌ Audio only | ⚠️ Near real-time | ✅ Pay-per-use |
| **Azure AI** | ✅ Speaker recognition | ⚠️ Via Video Indexer | ⚠️ Async | ✅ M365 customers |

**Finding:** No deepfake detection vendor has an SME-targeted product. The core SME use case — "Is this voice/video call real before we wire $50K?" — has NO turnkey solution.

#### Integrated AI Security / Compliance Platforms

| Vendor | AI-Specific Coverage | SME Fit |
|--------|---------------------|---------|
| HiddenLayer | ML model attacks — for orgs *deploying* AI models | ❌ |
| Wiz AI-SPM | Cloud AI pipeline misconfigs | ❌ |
| CrowdStrike Charlotte AI | AI-assisted SOC | ❌ Requires SOC |
| Vanta / Drata / Secureframe | Zero AI threat coverage | ✅ SME pricing, compliance only |

**Key Finding:** HiddenLayer, Wiz, Orca solve AI security for companies *building* AI. SMEs are *using* AI. These are different threat models. SME compliance tools (Vanta, Drata) have zero AI threat detection.

---

### 1.2 Market Gap Analysis

| Gap | Severity | Description |
|-----|----------|-------------|
| **AI User vs AI Builder Blind Spot** | 🔴 Critical | All serious vendors target AI builders. SMEs are AI consumers. No vendor covers this. |
| **No Unified Platform at SME Price** | 🔴 Critical | Full coverage requires 4–6 separate tools ($60K+/yr minimum). Unaffordable for 90% of SMEs. |
| **Deepfake-to-Fraud Correlation** | 🔴 Critical | The `deepfake call → impersonate executive → wire transfer` chain has no turnkey prevention solution. |
| **SME-Executable Response Workflows** | 🟠 High | Discovery tools find threats but don't provide non-security-staff-executable response playbooks. |
| **Shadow AI Risk Scoring** | 🟠 High | Nudge Security discovers apps, but doesn't score risk based on what data those employees access. |
| **Zero-IT-Config Employee Protection** | 🟠 High | Prompt Security comes closest but requires IT admin/developer setup. No "one-click protect" solution. |

---

### 1.3 Customer Pain Point Ranking

| Rank | Pain Point | Severity | Evidence |
|------|-----------|----------|----------|
| **1** | Employees pasting confidential data into ChatGPT/Copilot | 🔴 Critical | Samsung data leak incident; Gartner: 55% of orgs report uncontrolled LLM data sharing |
| **2** | AI-powered CEO/CFO voice impersonation for wire fraud | 🔴 Critical | FBI IC3: BEC losses $2.9B in 2023; voice cloning now $5 with ElevenLabs; avg SME loss $140K/incident |
| **3** | Shadow AI app sprawl — no visibility | 🟠 High | Nudge Security: average SME has 20+ unapproved AI tools |
| **4** | AI-powered hyper-personalized spear-phishing | 🟠 High | IBM X-Force: AI phishing 40% more effective; Huntress: 3x increase in AI-crafted BEC at SMEs in 2025 |
| **5** | Deepfake video in board/investor calls | 🟡 Medium-High | Growing, primarily affects 100–500 employee companies |
| **6** | Prompt injection in internal AI tools | 🟡 Medium | Relevant for only ~15–20% of SMEs deploying internal LLM tools |
| **7** | AI-generated disinformation targeting brand | 🟡 Medium | Growing concern, lacks immediate financial loss trigger |
| **8** | Adversarial ML / model poisoning | 🟢 Low | Irrelevant for SMEs not deploying their own ML models |

---

## 2. Feature Set — Finalized v1

> Iterated through: PO research → TA feasibility assessment → PO acceptance of counter-proposals → PM sprint sequencing.  
> Counter-proposals accepted: CP-1 (AI Submission Gate), CP-2 (Redaction Review), CP-3 (Lakera Guard partnership), CP-4 (AI Tool Attestation), CP-5 (M365 Defender first).

---

### Module A — AI Submission Gate (Browser DLP)

**A1 — Prompt Content Scanner** `Must-have`
- Extension intercepts at submit time (paste event + submit button click)
- 3-tier scanning: Tier 1 Regex (<1ms) → Tier 2 WASM BERT-tiny (~50–80ms) → Tier 3 Server API (async, non-blocking)
- Scans for: credit cards, SSNs, API keys, source code patterns, company entities
- Fails closed: if scanner offline, submission blocked with explicit notice
- Zero-knowledge architecture: prompt content never leaves the browser

**User Story:** As an IT admin, I want prompts to be scanned for sensitive data at submit time before they reach any AI tool, so that I can enforce data protection without changing employee workflow.

**Acceptance Criteria:**
1. Extension intercepts submissions on chatgpt.com, copilot.microsoft.com, gemini.google.com, claude.ai with <80ms latency (p95)
2. Tier 1 regex detects >99% of critical PII (credit cards, SSNs, API keys) with <1% false positive on legitimate prompts
3. Tier 2 WASM BERT-tiny detects >85% of semantic confidential data with <10% false positive
4. On trigger: submission blocked and Redaction Review UI presented (not a dismissable toast)
5. Extension unavailability fails closed (not silently pass-through)

**Success Metric:** >95% of sensitive data submissions result in Redaction Review engagement within 30 days; scanner availability >99.5%

---

**A2 — Pre-send Redaction Review** `Must-have`
- Sensitive tokens highlighted inline in the prompt editor with category label
- Default action: "Send with redactions applied"
- To include flagged content: deliberate secondary action + one-line justification (captured in audit log)
- Redaction uses placeholder tokens (e.g., `[CARD_1]`, `[PERSON_1]`) that preserve prompt grammatical coherence
- Bypass events visible to IT admin in dashboard within 60 seconds

**User Story:** As an employee, I want to see exactly which parts of my prompt were flagged before it's sent, so I can make an informed decision and send a safe version without rewriting from scratch.

**Acceptance Criteria:**
1. Sensitive tokens highlighted with category label in compose area
2. Default action = "Send with redactions"; explicit override required to send original
3. Override requires justification capture (logged to incident timeline)
4. Redaction preserves prompt meaning (grammatically coherent token substitution)
5. Bypass events appear in IT admin dashboard within 60 seconds

**Success Metric:** <5% of Redaction Review events result in employee sending unredacted original

---

**A3 — Extension Health Monitoring (Canary)** `Must-have`
- Synthetic health check per AI domain every hour
- Admin alert if health check fails >15 minutes
- Dashboard shows per-domain health status (green/amber/red) with last-verified timestamp
- Rationale: Prevent silent failure → false sense of security

**Acceptance Criteria:**
1. Extension performs synthetic health check on each supported domain once per hour
2. Health check failure >15 min raises P2 alert in dashboard + optional Slack/email notification
3. Dashboard shows per-domain health status with last-verified timestamp
4. Mean time to detect extension breakage <30 minutes

---

### Module B — Prompt Injection Detection Engine

**B1 — Tier 1 Pattern Detection (OWASP Regex Library)** `Must-have`
- ≥50 OWASP LLM Top 10 injection patterns
- Detection latency <1ms (client-side)
- Versioned library, updatable via server-push without extension reinstall
- Open-sourced as E2 community edition

**B2 — Tier 2 ML Detection** `Must-have`
- Server-side: Lakera Guard API (white-labeled, <300ms response, p95)
- Client-side fallback: WASM BERT-tiny (17MB, lazy-loaded, ~50–80ms) via ONNX Runtime Web
- Automatic fallback if Lakera API unavailable — no user-visible degradation
- Combined Tier 1+2 precision target: >90% on novel injection attempts

**B3 — Layer 3 Context-Aware Risk Scoring** `Must-have`
- Risk score multipliers: user role × data sensitivity level (Track 1) × AI tool × recurrence
- Response action tiers: 0–30 log only, 31–60 advisory + justification, 61–85 block + manager approval, 86–100 hard block + admin alert
- Fails secure: if Track 1 API unavailable, defaults to highest risk tier
- Partial v1: directory-sourced roles (job title from M365/Google); full Track 1 integration in Sprint 10

**B4 — Risk Response Action Engine** `Must-have`
- 4-tier enforcement: log / warn / block+approve / hard block
- Manager approval workflow via Slack/email with one-click approve/deny
- All risk events retained 90 days in audit log
- IT admin can replay any event (full risk score breakdown + action taken)

---

### Module C — Shadow AI Governance

**C1 — OAuth AI Tool Inventory** `Must-have`
- Pulls OAuth app grants from M365 Graph API (user-consented apps) + Google Workspace Admin API
- Identifies ≥30 known AI tools from maintained taxonomy
- Shows: tool name, category, data scopes granted, user count, first-seen date, risk classification
- New AI tool grant detected within 24 hours
- IT admin can mark tools Approved / Under Review / Blocked + trigger Track 1 revocation workflow

**C2 — AI Tool Attestation Workflow** `Must-have`
- Quarterly self-survey delivered in-app + email fallback
- Configurable: monthly by default for Enterprise tier
- Cross-references self-reported tools against C1 OAuth inventory
- Discrepancies (browser-only AI use not captured by OAuth) flagged as findings
- Non-response after 5 business days = compliance gap finding

**C3 — Extension Domain Usage Telemetry** `Should-have`
- Aggregate counts per domain per user per day (domain + count only, zero content)
- Distinguishes known AI domains from general traffic
- Displayed alongside OAuth inventory in unified Shadow AI view
- Employees can view what's being counted (privacy transparency)
- Can be disabled per user group by IT admin

**C4 — AI Tool Risk Scoring & Policy Enforcement** `Should-have`
- Risk score: OAuth scopes × vendor DPA availability × data residency × certifications × incidents
- Risk bands: Low / Medium / High / Critical
- IT admin can set automated policies: "Auto-block any new AI tool with Critical risk score"
- Scores updated when vendor posture changes (not static)

---

### Module D — Deepfake Defense

**D1 — Voice Deepfake Detection** `Must-have (with legal gate for EU)`
- Employee uploads audio (MP3/WAV/M4A, ≤60s) or records via MediaRecorder in browser
- Vendor API (Resemble Detect) returns deepfake probability score within 5 seconds
- Results shown as: "likely authentic" / "likely synthetic" / "inconclusive" — NOT binary fake/real
- Zero audio retention: audio deleted within 60 seconds of analysis
- **EU Legal Gate:** Deploy EU only after GDPR Article 9 legal opinion (commission Day 1; US/UK/AU ship first)
- Employee-initiated only; never employer-triggered

**D2 — Out-of-Band Verification Workflow** `Must-have`
- Employee triggers verification from app in ≤3 clicks after receiving a suspicious request
- System sends verification request via both email + SMS to purported sender (contacts from Track 1 HR integration)
- Purported sender receives confirm/deny link — no SMESec account required
- Employee receives result within 5 minutes
- Full verification timeline stored in audit log

**D3 — Video Deepfake Detection** `Should-have (Sprint 8 go/no-go)`
- Upload-based: MP4/MOV up to 2 minutes, admin-only
- Vendor API analysis within 30 seconds
- Results highlight specific artifacts (face boundary, lighting, blink patterns)
- Compliance evidence package, 90-day auto-delete unless pinned
- **Decision gate:** If vendor not contracted by Sprint 8 start → defer to v1.1

---

### Module E — AI Phishing Defense

**E1 — M365 Defender AI Phishing Alerts** `Must-have (M365 customers)`
- Phishing/malware alerts from Microsoft Security Graph API (`/security/alerts_v2`) within ≤5 minutes
- Alerts enriched with Track 1 asset context (affected user's role, data access)
- One-click trigger to Track 1 incident playbook from alert
- Applies to M365-licensed tenants only; graceful "not available" for non-M365

**E2 — Email Authentication Posture (Google Workspace)** `Should-have`
- Google Workspace Admin SDK: DMARC pass/fail rates, DKIM signing, SPF alignment
- Actionable remediation guidance for misconfigurations
- Weekly email authentication digest to IT admin
- Explicitly scoped as email authentication posture only (not positioned as phishing detection equivalent)

---

### Module F — Privacy, Trust & Transparency

**F1 — Employee Transparency Dashboard** `Must-have (required for EU)`
- What is monitored, what is NOT, how long data is retained, who can see it — in plain language
- Always accessible from extension popup and mobile app; IT admin cannot hide it
- Employee can view their own last 10 flagged/redacted events
- Explicit "What SMESec does NOT do" section

**F2 — Employee Opt-in / Pause Capability** `Must-have (required for EU consent model)`
- Employee can pause extension monitoring for 15/30/60 minutes
- When paused: zero scanning, zero telemetry, zero backend transmission
- IT admin notified when pause is activated (reason not logged — private)
- IT admin can restrict pause duration for specific roles (restriction disclosed in F1)

**F3 — Admin Coverage & Canary Monitoring** `Must-have`
- Dashboard: total employees, extension installed, active (checked in <24h), paused, offline >48h, health status
- Employees without extension = one-click bulk invite
- Automated alert if coverage drops below configurable threshold (default: <80% active)
- Coverage report exportable for compliance evidence

---

### Sub-module: Open-Source Contribution

**E_OS — Prompt Injection Rules Library** `Should-have`
- Fork OWASP LLM Top 10 rules + Jailbreak-LLMs dataset
- Publish as open-source library (community edition of B1)
- Rationale: brand authority, community improvements, inbound leads

---

## 3. Technical Architecture Decisions

### 3.1 Core Architecture: Zero-Knowledge DLP

The central architectural decision: **prompt content never leaves the browser.**

```
BROWSER EXTENSION (sandboxed content script per tab)
├── Content Script — lives for full tab lifetime (not service worker)
│   ├── Tier 1: Regex/rule scan  → <1ms
│   ├── Tier 2: WASM BERT-tiny  → 50–80ms (lazy-loaded, 17MB via ONNX Runtime Web)
│   └── Submit Gate UI          → Block/Warn modal
│
│   What is sent to backend:
│   ✅ Risk classification (HIGH/MED/LOW)
│   ✅ Pattern type detected (PII/SECRET/IP)
│   ✅ Target AI tool (chatgpt.com)
│   ✅ Timestamp + tenant ID
│   ❌ Actual prompt content — NEVER transmitted
│
└── Service Worker (coordination only, not scanning)
    └── Keepalive ping from content script every 25 seconds

BACKEND (FastAPI / AWS)
├── POST /api/v1/events/dlp       → EventBridge → Incident DB
├── POST /api/v1/detect/tier3     → Lakera Guard API (server-side only)
├── GET  /api/v1/ai-tools         → C1 OAuth inventory
└── POST /api/v1/verify/oob       → D2 out-of-band verification

DATA RETENTION:
├── Incident log (risk class + metadata): retained 7 years (S3 Object Lock)
└── Prompt content: NEVER retained (zero-knowledge by architecture)
```

### 3.2 Build vs Buy Decisions

| Capability | Decision | Rationale |
|------------|----------|-----------|
| Voice deepfake detection | **Buy (Resemble Detect API)** | Training custom model requires 50K+ voice samples; API cost viable at realistic usage (5–20 checks/day/company = $0.30–0.60/mo API cost) |
| LLM DLP patterns | **Augment open-source** (Microsoft Presidio + custom rules) | Presidio is production-grade, MIT licensed, covers 50+ entity types. Add: AWS credentials regex, GitHub tokens, source code patterns, API key formats |
| Shadow AI app catalog | **Build + curate** | No off-the-shelf AI-specific risk scoring for SME context. Manual catalog of 100 top AI tools with quarterly maintenance |
| Prompt injection (browser) | **Build with ONNX BERT-tiny** | WASM BERT-tiny viable for browser-side scanning. No server round-trip = better privacy + latency |
| Prompt injection (API scanner) | **White-label Lakera Guard** | Production-hardened, continuously updated. SMESec differentiator is context-aware scoring (Layer 3), not the injection pattern library |
| AI phishing | **Partner** | M365 Defender via Graph API (existing integration), Sublime Security in v2 for Google Workspace customers |

### 3.3 Key Technical Risks

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| Chrome MV3 service worker termination | High (60%) | Critical | Prototype offscreen documents / keepalive solution in Sprint 1, Week 1. Hard gate: June 6 decision. Three contingency architectures defined (offscreen docs / server-side proxy / API-only mode) |
| React SPA DOM fragility (ChatGPT UI changes) | High | High | Site-specific adapter pattern; daily automated E2E tests against live AI tool URLs; canary health checks (A3) alert admin within 30 min of failure |
| Employee privacy trust erosion | High in EU | High | Zero-knowledge architecture; open-source extension code; F1 transparency dashboard; F2 pause capability; employee opt-in consent flow |
| WASM 17MB extension size (Chrome Web Store 10MB limit) | High (65%) | Medium | Lazy-load WASM post-install via CDN with integrity hash; extension package <2MB at install time |
| GDPR Article 9 (voice = biometric data) | Medium | High | Commission legal opinion Day 1; ship D1 for US/UK/AU first; EU deployment gated on legal clearance |
| Lakera Guard unit economics | Medium (35%) | Medium | Validate pricing by end of Sprint 1; WASM-only fallback always built regardless |
| B3 Track 1 data dependency | Medium (45%) | Medium | Ship B3 with directory-role approximation in v1; full integration Sprint 10 |

---

## 4. Build vs Defer Decisions

### Must-Build for v1 Launch
- A1 (Prompt Content Scanner), A2 (Redaction Review), A3 (Extension Health Monitoring)
- B1 (OWASP Regex Library), B2 (Lakera Guard + WASM), B3 (Context Risk Scoring), B4 (Action Engine)
- C1 (OAuth AI App Inventory), C2 (AI Tool Attestation), C4 (Risk Scoring + Policy)
- D1 (Voice Deepfake — non-EU first), D2 (Out-of-Band Verification)
- E1 (M365 Defender Phishing Alerts)
- F1 (Employee Transparency), F2 (Pause Capability), F3 (Admin Coverage Dashboard)

### Defer to v1.1
- D3 (Video Deepfake Detection) — Sprint 8 go/no-go; defer if vendor not contracted
- C2 (Attestation Workflow) — deferrable if Backend Eng overloaded in S6–S7
- E2 (Google Workspace DMARC visibility) — informational feature, deferrable
- B3 full Track 1 integration — partial version ships in v1

### Do Not Build / Partner Instead
- AI email phishing detection engine → M365 Defender (existing), Sublime Security (v2 partner)
- Adversarial ML / model poisoning protection → Not an SME problem; HiddenLayer covers this for enterprises
- Dynamic proxy-based redaction → Replaced by Redaction Review (CP-2); proxy architecture is fragile and creates SSL trust issues
- DNS/network traffic-based shadow AI detection → Replaced by AI Tool Attestation (CP-4) + C3 telemetry

---

## 5. Delivery Plan (6-Month Sprint Sequence)

**Team:** 3 engineers (1 ML, 1 Frontend/Extension, 1 Full-stack Backend) | 13 × 2-week sprints

| Sprint | Dates | Key Deliverables | Critical Gate |
|--------|-------|-----------------|---------------|
| S1 | Jun 2–13 | B1 regex engine, MV3 prototype, FastAPI scaffold | **HARD GATE: MV3 persistence by June 6** |
| S2 | Jun 16–27 | B1 server-push updates, extension scaffold, C1 start | Lakera Guard pricing decision |
| S3 | Jun 30–Jul 11 | **A1 Prompt Interception complete**, C1 complete, B2 start | A1 intercepts 100% of ChatGPT/Copilot submits |
| S4 | Jul 14–25 | **A2 Redaction Review**, B2 complete, F2 Pause | A2 default-redact flow working end-to-end |
| S5 | Jul 28–Aug 8 | B3 context scoring, B4 action engine, C2, E1 | B4: injection → correct response tier; **OQ-6 managed deploy confirmed** |
| S6 | Aug 11–22 | C4 risk scoring, F1 transparency, F3 coverage, D2 | C4 policy enforcement live; D2 out-of-band verification |
| S7 | Aug 25–Sep 5 | **D1 Voice Deepfake**, end-to-end pipeline integration | D1: audio analysis <5s; legal opinion status check |
| S8 | Sep 8–19 | D3 go/no-go, Chrome Enterprise deployment, pen-test commissioned | Chrome Enterprise push confirmed; pen-test vendor LOI signed |
| S9 | Sep 22–Oct 3 | **Pilot execution** (2–3 SME customers) | Real-world FP <10% injection, <5% DLP |
| S10 | Oct 6–17 | B3 full Track 1 integration, pen-test remediation | All Critical/High pen-test findings resolved |
| S11 | Oct 20–31 | Hardening, load testing (500 users), documentation | Load test passes; no new Critical findings |
| S12 | Nov 3–14 | Chrome Web Store submission, go/no-go checklist | Extension approved for publication |
| S13 | Nov 17–28 | **Launch** — graduated rollout (5 → 10 → full) | Day 1/3/7 customer success check-ins |

### Critical Path
```
[S1: MV3 Gate] → [S3: A1 Interception] → [S4: A2 Redaction] → [S5: B4 Action Engine] 
→ [S7: Integration Testing] → [S9: Pilot] → [S10: Pen-test Remediation] → [S12: Store Submission] → Launch

Slack on critical path: ZERO — any slip cascades to launch date.
```

### Team Bottlenecks
- **Extension Engineer** is the single point of failure (owns A1, A2, A3, B4 extension-side, C3, F1, F2, F3, D1 UI)
- Sprint 3 = 100% utilization for Extension Engineer; no buffer
- **Drop/defer priority** if behind: keep A1 > A2 > B4 enforcement; simplify F1 to static HTML; defer D3 and C2 UI

---

## 6. Open Questions & Blockers

| # | Question | Blocks | Owner | Deadline |
|---|----------|--------|-------|----------|
| **OQ-1** | GDPR Article 9 legal opinion for D1 voice deepfake analysis | D1 EU deployment | Legal / PM | Commission Day 1; opinion by Sprint 4–5 |
| **OQ-2** | Lakera Guard pricing — viable at SME scale (<$0.05/req)? | B2 API vs WASM-only architecture | PM + ML Engineer | End of Sprint 1 (June 13) |
| **OQ-3** | EU employee consent model for extension monitoring | Extension onboarding flow design | PM + Legal | Sprint 5 (August 8) |
| **OQ-4** | Resemble Detect accuracy on non-English, compressed VoIP clips (<15s) | D1 accuracy claims in EU markets | ML Engineer | Sprint 7 planning |
| **OQ-5** | Google Workspace phishing gap — Sublime Security timeline | Go/no-go messaging to GWS customers | PM | Before launch go/no-go |
| **OQ-6** | Chrome Enterprise managed deployment confirmation | Pilot onboarding at scale | Extension Engineer | End of Sprint 5 (August 8) |

### PM Mandatory Conditions (from PM agent)
1. **Chrome MV3 prototype = hard gate by June 6** — no exceptions, no extensions
2. **Lakera Guard pricing decision by June 13** — no extensions
3. **D3 (video deepfake) = out-of-scope for planning until Sprint 8 go/no-go**
4. **C2 and E2 = planned sacrifice surface** if team hits 3+ simultaneous risk materializations
5. **Pilot customer outreach begins June 2** (not Week 3)

---

## 7. Differentiated Value Proposition

> "SMESec is the only security platform built around how SMEs actually work in 2026 — where your biggest threats aren't firewall breaches but employees inadvertently feeding confidential data into ChatGPT, a deepfake 'CEO' voice call authorizing a wire transfer, or a dozen unsanctioned AI tools quietly holding access to your company email. Unlike Vanta or Drata, which generate compliance reports after the fact, SMESec actively stops these threats in real time — blocking sensitive data before it leaves your browser, flagging deepfake fraud before money moves, and mapping your entire AI tool footprint automatically. And unlike enterprise security tools that require a dedicated security team to operate, every feature in SMESec is designed to be understood and acted on by the same IT admin who also manages your Google Workspace."

---

## 8. Agent Iteration Summary

This research was produced through a structured 4-pass iteration loop:

| Pass | Agent | Output | Key Contribution |
|------|-------|--------|-----------------|
| Round 1 | Product Owner | Competitive landscape, market gap analysis, pain point ranking, initial feature set, open questions T1–T8 | Identified the "AI consumer vs AI builder" market gap; ranked deepfake wire fraud as #2 pain point with hard evidence |
| Round 1 | Technical Advisor | Feasibility ratings for all features, answered T1–T8, architecture recommendation (zero-knowledge DLP), 5 counter-proposals | Zero-knowledge architecture; Chrome MV3 service worker clarification; GDPR Article 9 hard stop on D1 EU; Lakera Guard white-label recommendation |
| Round 2 | Product Owner | Accepted/rejected counter-proposals, finalized feature list with user stories and acceptance criteria, remaining open questions, value proposition | Accepted all 5 counter-proposals with modifications; added F module (employee transparency + pause) as hard requirements for EU adoption |
| Round 2 | Project Manager | 13-sprint delivery plan, critical path analysis, 7-risk register, bottleneck analysis, mandatory decision gates | Identified Extension Engineer as single point of failure; MV3 gate June 6 as the make-or-break decision for the entire delivery model |

**Combined agent confidence: 7.5/10**  
Strong product shape with genuine market differentiation. Real blockers: MV3 prototype (architecture), Lakera unit economics (architecture), GDPR Article 9 (legal), managed deployment confirmation (pilot readiness). None are insurmountable, but all must be resolved on the sprint schedule above.
