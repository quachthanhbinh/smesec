# SMESec — AI Governance Module: Phân Tích Chuyên Sâu

**Ngày tạo:** 2026-05-28  
**Phiên bản:** 1.0  
**Trạng thái:** Approved  
**Nguồn:** Multi-agent research synthesis (Product Owner × Technical Advisor × Project Manager × 4 iterations)  
**Deliverable gốc:** Topic.md #4 — "Detail your approach to detecting and governing employee use of external AI tools (e.g., ChatGPT, Copilot) and the risks this introduces."  
**Liên quan:** [design-document.md](design-document.md) · [system-architecture.md](system-architecture.md) · [delivery-plan.md](delivery-plan.md)

---

## Mục Lục

1. [Bối Cảnh & Mức Độ Khẩn Cấp](#1-bối-cảnh--mức-độ-khẩn-cấp)
2. [Phân Tích Thị Trường & Khoảng Trống Giải Pháp](#2-phân-tích-thị-trường--khoảng-trống-giải-pháp)
3. [Threat Model: 8 Rủi Ro AI SME Phải Đối Mặt](#3-threat-model-8-rủi-ro-ai-sme-phải-đối-mặt)
4. [AI Governance Framework: 3 Tầng Kiểm Soát](#4-ai-governance-framework-3-tầng-kiểm-soát)
5. [Module A — AI Submission Gate (Browser DLP)](#5-module-a--ai-submission-gate-browser-dlp)
6. [Module B — Prompt Injection Detection Engine](#6-module-b--prompt-injection-detection-engine)
7. [Module C — Shadow AI Governance](#7-module-c--shadow-ai-governance)
8. [Module D — Deepfake Fraud Defense](#8-module-d--deepfake-fraud-defense)
9. [Module E — AI Phishing Defense](#9-module-e--ai-phishing-defense)
10. [Module F — Employee Privacy & Transparency](#10-module-f--employee-privacy--transparency)
11. [Kiến Trúc Kỹ Thuật: Zero-Knowledge DLP](#11-kiến-trúc-kỹ-thuật-zero-knowledge-dlp)
12. [Build vs Buy: Quyết Định Từng Component](#12-build-vs-buy-quyết-định-từng-component)
13. [Risk Register: 7 Rủi Ro Triển Khai Chính](#13-risk-register-7-rủi-ro-triển-khai-chính)
14. [Delivery Sequence: 13-Sprint Roadmap](#14-delivery-sequence-13-sprint-roadmap)
15. [Key Performance Indicators](#15-key-performance-indicators)
16. [Open Questions & Hard Gates](#16-open-questions--hard-gates)

---

## 1. Bối Cảnh & Mức Độ Khẩn Cấp

### 1.1 Tại Sao AI Governance Là Cấp Thiết Ngay Hôm Nay

Năm 2026, AI tools đã trở thành phần thiết yếu của công việc hàng ngày. Nhưng cũng chính vì điều này, tuyến phòng thủ bảo mật truyền thống — tường lửa, endpoint protection, email filtering — không còn đủ. Các mối đe dọa lớn nhất bây giờ đến **từ trong tổ chức**, khi nhân viên vô tình hoặc cố ý:

- Paste source code, dữ liệu khách hàng, thông tin tài chính vào ChatGPT
- Cấp OAuth access cho AI tools không được IT phê duyệt
- Thực hiện chuyển khoản sau một cuộc gọi video giả mạo CEO

**Quy mô vấn đề (dữ liệu thực tế 2025–2026):**

| Số liệu | Nguồn | Mức Độ |
|---|---|---|
| 55% tổ chức báo cáo nhân viên chia sẻ dữ liệu không kiểm soát qua LLM | Gartner 2025 | 🔴 Nghiêm trọng |
| Samsung mất source code khi kỹ sư paste vào ChatGPT | Samsung IR 2023 | 🔴 Ví dụ điển hình |
| Average SME có >20 AI tools không được IT biết đến | Nudge Security 2024 | 🔴 Nghiêm trọng |
| BEC (Business Email Compromise) losses: $2.9B năm 2023 | FBI IC3 2024 | 🔴 Nghiêm trọng |
| Voice cloning tool giá ~$5/month (ElevenLabs, thấp hơn) | Open market 2025 | 🟠 Đáng lo ngại |
| Average SME mất $140K/incident từ deepfake wire fraud | FBI IC3 2024 | 🔴 Nghiêm trọng |
| 78% knowledge workers dùng AI tools tại nơi làm việc | Microsoft Work Trend Index 2025 | Baseline |
| 52% dùng tools mà employer không cung cấp | Cyberhaven Research 2025 | 🟠 Shadow AI scope |
| 11% content paste vào ChatGPT là thông tin mật công ty | Cyberhaven Research 2025 | 🔴 IP leakage |

### 1.2 Tại Sao SME Đặc Biệt Dễ Bị Tổn Thương

```
SME vs Enterprise — AI Governance Readiness Gap:

┌────────────────────────────────────────────────────────────────┐
│                        ENTERPRISE                              │
│  ● Dedicated CISO + Security team (5–50 người)                 │
│  ● Ngân sách $200K–$2M/yr cho security tools                  │
│  ● CrowdStrike Charlotte AI, Netskope, Zscaler DLP             │
│  ● Policy framework, training programs, enforcement            │
└────────────────────────────────────────────────────────────────┘

                ╔══════════════════════════════╗
                ║    THE PROTECTION GAP         ║  ← SMESec fills this
                ╚══════════════════════════════╝

┌────────────────────────────────────────────────────────────────┐
│                          SME                                   │
│  ● IT "admin" = developer kiêm thêm IT (20% thời gian)        │
│  ● Ngân sách bảo mật $5K–50K/yr (nếu có)                      │
│  ● Không có DLP, không có AI governance, không có SOC          │
│  ● "Chúng tôi tin tưởng nhân viên" = chiến lược duy nhất       │
│  ● Nhân viên dùng ChatGPT như dùng Google — không phân biệt   │
└────────────────────────────────────────────────────────────────┘
```

**Conclusion:** SME cần bảo vệ AI-grade nhưng không thể vận hành enterprise-grade tools. Đây là khoảng trống mà tất cả vendor hiện tại đều bỏ ngỏ.

---

## 2. Phân Tích Thị Trường & Khoảng Trống Giải Pháp

### 2.1 Competitor Feature Matrix — Shadow AI & DLP

| Vendor | Shadow AI Discovery | Policy Enforcement | DLP | SME Pricing | SME UX | Verdict |
|---|---|---|---|---|---|---|
| **Nudge Security** | ✅ OAuth app discovery, AI categorization | ⚠️ Nudge/alert only — không thể block | ❌ | ~$4–8/user/mo | ✅ | Gần nhất với SME nhưng không có enforcement |
| **Obsidian Security** | ✅ SaaS posture + identity risk | ✅ Policy-based | ⚠️ Cơ bản | ❌ $40K+/yr | ❌ | Enterprise only |
| **DoControl** | ✅ SaaS data access mapping | ✅ Automated remediation | ✅ Mạnh | ❌ $30K+/yr | ❌ | Enterprise only |
| **Metomic** | ⚠️ Data in SaaS only | ❌ Alert only | ✅ PII scanning | ~$5K–15K/yr | ✅ | Compliance-focused, không real-time |
| **Nightfall AI** | ❌ DLP-only | ❌ | ✅ Best-in-class cloud DLP | ~$10K–20K/yr | ⚠️ | Không có AI governance |
| **Prompt Security** | ❌ | ✅ Browser ext + API gateway | ✅ PII redaction | ❌ $15K–30K/yr | ❌ | Đòi hỏi IT admin/developer setup |
| **SMESec (target)** | ✅ OAuth + extension telemetry | ✅ Block + attestation + policy | ✅ Zero-knowledge local | **✅ $3–5/user/mo** | **✅ Non-expert** | **Unified, SME-native** |

### 2.2 LLM Security Vendors — Prompt Injection & DLP

| Vendor | Prompt Injection | DLP | Browser Extension | SME Fit | Vấn Đề |
|---|---|---|---|---|---|
| **Lakera Guard** | ✅ Xuất sắc — real-time API | ⚠️ Cơ bản | ❌ | ❌ $20K+/yr | Developer/enterprise only |
| **Prompt Security** | ✅ Browser ext + API gateway | ✅ PII redaction | ✅ Chrome | ❌ $15K–30K/yr | Cần IT setup |
| **Protect AI** | ✅ ML model security + LLM firewall | ✅ AI Red Team | ❌ | ❌ $50K+/yr | Enterprise only |
| **Aporia** | ✅ LLM guardrails | ✅ PII/IP detection | ❌ | ❌ $15K+/yr | API integration required |
| **LLM Guard** | ✅ PII anonymization | ✅ | ❌ | ❌ DIY | Self-hosted DevOps required |
| **SMESec** | ✅ Rule-based + WASM BERT | ✅ Zero-knowledge | **✅ Chrome MV3 + Edge** | **✅** | **Install and protect** |

### 2.3 Deepfake Detection — Vendor Landscape

| Vendor | Voice | Video | Real-Time | SME Pricing | Vấn Đề |
|---|---|---|---|---|---|
| **Reality Defender** | ✅ | ✅ | ✅ | ❌ $30K+/yr | Enterprise only |
| **Pindrop** | ✅ Best-in-class | ❌ | ✅ Telephony | ❌ $50K+/yr | Call center specific |
| **Hive Moderation** | ✅ | ✅ Video + image | ⚠️ Async | **✅ $0.001–0.01/req** | **SME viable** |
| **Resemble Detect** | ✅ | ❌ | ⚠️ Near real-time | **✅ Pay-per-use** | Audio only |
| **Azure AI Speaker** | ✅ | ⚠️ Video Indexer | ⚠️ Async | ✅ M365 customers | Cần Azure integration |

**Key Finding:** Không có vendor nào có **sản phẩm AI governance cho SME**. Khoảng trống thị trường là $2–4B tại thị trường SME toàn cầu (ước tính dựa trên TAM).

### 2.4 Khoảng Trống Thị Trường — 6 Gaps Quan Trọng

| Khoảng Trống | Mức Độ | Mô Tả |
|---|---|---|
| **AI Consumer vs AI Builder Blind Spot** | 🔴 Critical | Tất cả vendor serious đều protect AI builders. SME là AI consumers. Đây là threat model khác biệt hoàn toàn. |
| **Không có unified platform tại SME pricing** | 🔴 Critical | Cover đầy đủ cần 4–6 tools riêng lẻ ($60K+/yr). 90% SME không thể afford. |
| **Deepfake → Fraud correlation chain** | 🔴 Critical | Chuỗi `deepfake call → impersonate exec → wire transfer` chưa có giải pháp turnkey nào. |
| **SME-executable response workflows** | 🟠 High | Discovery tools tìm thấy threats nhưng không cung cấp response playbooks cho non-security staff. |
| **Shadow AI risk scoring dựa theo context** | 🟠 High | Nudge Security discover apps, nhưng không score risk dựa trên loại data employees đó có access. |
| **Zero-IT-config employee protection** | 🟠 High | Prompt Security gần nhất nhưng cần IT admin/developer setup. Không có "one-click protect". |

---

## 3. Threat Model: 8 Rủi Ro AI SME Phải Đối Mặt

### 3.1 Threat Ranking theo Xác Suất × Tác Động

| Rank | Mối Đe Dọa | Severity | Xác Suất | ROI Prevention | Evidence |
|---|---|---|---|---|---|
| **#1** | **Nhân viên paste confidential data vào ChatGPT/Copilot** | 🔴 Critical | Đang xảy ra — 55% orgs | Trực tiếp: IP, GDPR violation, client data breach | Samsung incident; Cyberhaven 11% figure |
| **#2** | **AI-powered CEO/CFO voice impersonation → wire fraud** | 🔴 Critical | Tăng 3x/năm | $140K avg SME loss/incident | FBI IC3 $2.9B BEC 2023; ElevenLabs $5/mo voice clone |
| **#3** | **Shadow AI app sprawl — không có visibility** | 🟠 High | >80% SME đã có | GDPR/ISO liability, data sovereignty | Nudge Security: avg 20+ unapproved AI tools |
| **#4** | **AI-powered hyper-personalized spear-phishing** | 🟠 High | Tăng 40%/năm | Avg phishing loss $136K/SME | IBM X-Force; Huntress 3x BEC 2025 |
| **#5** | **Deepfake video trong board/investor calls** | 🟡 Medium-High | 100–500 employee firms | Reputational + financial | Growing, >$1M recorded cases in 2025 |
| **#6** | **Prompt injection trong internal AI tools** | 🟡 Medium | ~15–20% SME deploying internal LLMs | Data exfil, incorrect AI outputs | OWASP LLM Top 10; Enterprise tier |
| **#7** | **AI-generated disinformation nhắm vào brand** | 🟡 Medium | Selective targeting | Reputational, share price | Growing, không có immediate loss trigger |
| **#8** | **Adversarial ML / model poisoning** | 🟢 Low | Không liên quan SME | Không applicable | SME không deploy custom ML models |

### 3.2 Attack Patterns Chi Tiết

#### Pattern 1: The Confidential Data Paste (Mối đe dọa #1)

```
ATTACK CHAIN:
  Developer → Opens ChatGPT.com
  Developer → Pastes: database schema + API keys + customer PII
  (Intent: "Debug this code for me")
  
  → ChatGPT processes request
  → Data potentially used in training (if user hasn't opted out)
  → Data visible to OpenAI employees in abuse monitoring
  → GDPR violation: customer PII sent outside EU without DPA agreement
  → IP exposure: proprietary business logic disclosed to third party

SMESEC INTERCEPT POINT:
  Browser Extension Content Script monitors textarea on chatgpt.com
  → Pre-submit scan: detects API_KEY regex + PII patterns
  → Blocks submission
  → Shows Redaction Review UI: "[API_KEY_1] [EMAIL_1] [PHONE_1]"
  → Employee can send redacted version OR override with justification
  → Override event logged → IT admin dashboard within 60 seconds
```

#### Pattern 2: The Deepfake CEO Call (Mối đe dọa #2)

```
ATTACK CHAIN:
  Attacker: Collects 30s of CEO voice from YouTube/LinkedIn videos
  Attacker: Generates clone via ElevenLabs/Resemble ($5/month)
  Attacker: Calls CFO/Finance: "I need an urgent wire transfer, $85K"
  CFO: Voice sounds exactly like CEO, trusts it
  CFO: Initiates wire transfer
  → $85K lost, recovery rate <10% for SME (FBI)

SMESEC INTERCEPT POINT:
  CFO opens SMESec mobile app → "Verify this call"
  → SMESec initiates OOB verification:
      a. Sends verification code to CEO's registered phone (separate channel)
      b. Sends audio sample to Hive Moderation API for deepfake analysis
  → Two independent signals:
      Code NOT received + deepfake_score > 0.7 → "LIKELY SYNTHETIC — Do NOT proceed"
      Code received + deepfake_score < 0.3    → "VERIFIED — Likely authentic"
  → CFO makes informed decision. Event logged as compliance evidence.
```

#### Pattern 3: Shadow AI OAuth Creep (Mối đe dọa #3)

```
ATTACK CHAIN:
  Employee (Sales): Discovers Jasper AI for email writing
  Employee: Clicks "Connect with Google" → grants Jasper access to Gmail + Drive
  (Intent: automate outreach emails)
  
  Jasper now has READ access to:
  → All company emails (confidential deals, M&A discussions)
  → All Drive files (financial models, client contracts)
  
  IT admin: has NO idea Jasper exists in company ecosystem
  Jasper: data may be used for AI training per their ToS

SMESEC INTERCEPT POINT:
  Track 1 Google Workspace Admin API sync (every 15 min)
  → Detects new OAuth grant: "Jasper AI" for user "alice@company.com"
  → AI Tool Classifier: category = 'ai_writing', scopes = ['gmail.readonly', 'drive.readonly']
  → Risk scoring: HIGH (broad scopes + email access + unknown data retention policy)
  → IT admin receives alert + in-app notification
  → Alice receives attestation request: "Confirm your use of Jasper AI"
  → IT admin can approve (add to allow-list) OR revoke OAuth grant (one-click)
```

---

## 4. AI Governance Framework: 3 Tầng Kiểm Soát

### 4.1 Kiến Trúc Tầng

```
╔═══════════════════════════════════════════════════════════════╗
║  TIER 3 — PROTECT (Real-time Prevention)                      ║
║  Browser Extension: intercepts AI tool submissions            ║
║  ● Block sensitive data before it leaves browser              ║
║  ● Local PII inference (zero-knowledge)                       ║
║  ● Deepfake detection workflow                                 ║
╠═══════════════════════════════════════════════════════════════╣
║  TIER 2 — GOVERN (Policy Enforcement)                         ║
║  Track 2 AI Services: policy engine + attestation             ║
║  ● Block/allow AI tools by policy tier                        ║
║  ● Require attestation for HIGH risk tools                    ║
║  ● Manager approval for override events                       ║
╠═══════════════════════════════════════════════════════════════╣
║  TIER 1 — DISCOVER (Passive Inventory)                        ║
║  Track 1 Integration Sync: OAuth inventory                    ║
║  ● Discover all AI tools via OAuth API grants                 ║
║  ● Classify and risk-score automatically                      ║
║  ● Close the shadow AI visibility gap                         ║
╚═══════════════════════════════════════════════════════════════╝

Dependency:  Tier 1 → feeds → Tier 2 → feeds → Tier 3 context
Independence: Tier 3 (browser ext) works even if Tier 1/2 is unavailable
              (fails closed — blocks rather than passes silently)
```

### 4.2 Coverage Matrix: Tier × Threat

| Mối Đe Dọa | Tier 1 (Discover) | Tier 2 (Govern) | Tier 3 (Protect) | Coverage |
|---|---|---|---|---|
| Confidential data paste | ❌ | ⚠️ Context enrichment | ✅ Real-time block | **✅ Tier 3** |
| Shadow AI OAuth sprawl | ✅ OAuth inventory | ✅ Policy + attestation | ⚠️ Domain telemetry | **✅ Tier 1+2** |
| Deepfake wire fraud | ❌ | ❌ | ✅ OOB verification | **✅ Tier 3** |
| Prompt injection (internal LLM) | ❌ | ❌ | ✅ Rule-based + ML | **✅ Tier 3** |
| AI phishing (M365) | ⚠️ Asset context | ✅ Alert routing + playbook | ❌ | **✅ Tier 1+2** |
| Browser-only AI usage (no OAuth) | ❌ | ❌ | ✅ Extension telemetry | **⚠️ Partial** |
| API-based AI usage (no browser) | ⚠️ OAuth grant | ✅ Policy | ❌ No browser hook | **⚠️ Partial** |

**Coverage at v1:** >80% của các AI threat vectors được cover bởi Tier 1 + Tier 2 + Tier 3 kết hợp.  
**Blind spots được thừa nhận:** API-direct usage không qua browser, non-Chrome browsers, mobile AI apps.

---

## 5. Module A — AI Submission Gate (Browser DLP)

### 5.1 Mục Tiêu & User Story

> **User Story (IT Admin):** "Tôi muốn các prompt được scan trước khi gửi tới AI tools, để tôi có thể enforce data protection mà không cần thay đổi workflow của nhân viên."

> **User Story (Employee):** "Tôi muốn thấy chính xác phần nào trong prompt của tôi bị flag trước khi gửi đi, để tôi có thể đưa ra quyết định có hiểu biết và gửi phiên bản an toàn mà không cần viết lại từ đầu."

### 5.2 A1 — Prompt Content Scanner (Must-have)

```
SCANNING PIPELINE: 3 Tầng Trong Browser

┌─────────────────────────────────────────────────────────────┐
│  TIER 1: Regex / Rule Engine              Latency: <1ms     │
│  ─────────────────────────────────────────────────────────  │
│  Patterns (server-push updatable, versioned):               │
│  ● Credit card numbers (Luhn algorithm + regex)             │
│  ● SSN / Tax ID (US: 123-45-6789, VN: CCCD format)         │
│  ● Email addresses + phone numbers                          │
│  ● API keys: AWS (AKIA...), GitHub (ghp_...), Stripe (sk_)  │
│  ● JWT tokens                                               │
│  ● IBAN / SWIFT codes                                       │
│  ● Source code patterns (class declarations, db schemas)    │
│  ● Company-specific keywords (configurable per tenant)      │
│                                                             │
│  Accuracy: >99% for CRITICAL PII (credit card, API keys)   │
│  False positive rate: <1% on legitimate prompts             │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│  TIER 2: WASM BERT-tiny Semantic Scanner  Latency: 50–80ms  │
│  ─────────────────────────────────────────────────────────  │
│  Model: Microsoft Presidio → compiled to WASM via ONNX      │
│  Size: 17MB (lazy-loaded post-install via CDN, integrity hash)│
│  Detects: Semantic confidential patterns that regex misses: │
│  ● "Our Q3 revenue forecast is..." (no PII but CONFIDENTIAL) │
│  ● "The acquisition target is..." (M&A discussion)         │
│  ● Employee performance discussions                         │
│  ● Client-specific proprietary information                  │
│                                                             │
│  Accuracy: >85% for semantic confidential data              │
│  False positive rate: <10% (warns, not always blocks)       │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│  TIER 3: Context-Aware Risk Scoring       Async (non-blocking)│
│  ─────────────────────────────────────────────────────────  │
│  Server-side: FastAPI → Lakera Guard API                    │
│  Enriches with Track 1 context:                             │
│  ● User's role (e.g. CFO → higher risk multiplier)         │
│  ● Data sensitivity level of user's assets                  │
│  ● Target AI tool risk tier                                 │
│  ● Recurrence (same user, 3rd time today)                   │
│                                                             │
│  Score 0–100 → Risk Tier → Response Action                  │
│  Không block trong Tier 3 (async) — Tier 1+2 là gatekeeper │
└─────────────────────────────────────────────────────────────┘
```

**Supported AI Tools (v1 launch — expandable via server-push config):**
- chatgpt.com, chat.openai.com
- copilot.microsoft.com, bing.com/chat
- gemini.google.com, bard.google.com
- claude.ai
- perplexity.ai
- github.com/copilot
- notion.so (AI features)

**Fail-closed Architecture:**
> Nếu extension không load được WASM (cold start), WASM scan bị skip — Tier 1 regex vẫn chạy. Nếu extension hoàn toàn offline → submission **bị block** với notice rõ ràng: "SMESec extension offline — submission blocked per company policy". Không bao giờ silent pass-through.

### 5.3 A2 — Pre-send Redaction Review (Must-have)

```
REDACTION REVIEW UI FLOW:

User types prompt in ChatGPT → Presses Enter/Submit
  ↓
[Browser Extension intercepts]
  ↓
Tier 1 + Tier 2 scan completes
  ↓
┌─────────────────────────────────────────────────────────────┐
│  REDACTION REVIEW MODAL (cannot be dismissed by Esc)        │
│                                                             │
│  ⚠️ Sensitive data detected in your prompt                  │
│                                                             │
│  Your original prompt:                                      │
│  "Debug this function for [API_KEY_1] that calls our DB     │
│   for user [EMAIL_1] at [PHONE_1]. The key is [SECRET_1]." │
│                                                             │
│  Detected:                                                  │
│  [API_KEY_1]  → AWS API Key         🔴 CRITICAL             │
│  [EMAIL_1]    → Email address       🟠 HIGH                 │
│  [PHONE_1]    → Phone number        🟠 HIGH                 │
│  [SECRET_1]   → Secret/credential  🔴 CRITICAL             │
│                                                             │
│  ┌─────────────────────────────────────────────────────┐   │
│  │  ✅ RECOMMENDED: Send with redactions applied       │   │
│  │  (sensitive data replaced with placeholders)        │   │
│  └─────────────────────────────────────────────────────┘   │
│                                                             │
│  ┌─────────────────────────────────────────────────────┐   │
│  │  ⚠️ Override: Send original (requires justification)│   │
│  │  [__________________________] Business reason       │   │
│  └─────────────────────────────────────────────────────┘   │
│                                                             │
│  📋 Cancel — I'll rewrite my prompt                         │
└─────────────────────────────────────────────────────────────┘

ON "SEND WITH REDACTIONS":
  → Prompt text modified: sensitive tokens replaced with placeholders
  → Modified prompt submitted to AI tool
  → Event logged: {type: 'redaction_applied', categories: ['api_key', 'email'], tenant_id}
  → IT admin visible: "1 submission auto-redacted today"

ON "SEND ORIGINAL WITH JUSTIFICATION":
  → Prompt submitted unchanged
  → Event logged: {type: 'bypass_override', justification: "...", categories: [...], tenant_id, user_hash}
  → IT admin ALERT (within 60 seconds): "User bypassed DLP — review required"
  → Logged to incident timeline for compliance evidence
```

**Acceptance Criteria:**
1. Sensitive tokens highlighted với category label trong compose area
2. Default action = "Send with redactions"; explicit override required
3. Override yêu cầu justification capture (logged to incident timeline)
4. Redaction giữ nguyên ngữ nghĩa của prompt (grammatically coherent token substitution)
5. Bypass events xuất hiện trong IT admin dashboard trong <60 giây
6. Metric: <5% Redaction Review events dẫn đến employee gửi unredacted original

### 5.4 A3 — Extension Health Monitoring (Must-have)

```
CANARY HEALTH CHECK SYSTEM:

Every 60 minutes, per supported AI domain:
  Extension Content Script → synthetic navigation check
  Verifies: DOM hook active, intercept pipeline responding
  
Health status per domain:
  GREEN: Checked <30 min ago, all hooks active
  AMBER: Last check >30 min ago, or last check had warnings
  RED:   Last check failed, or >15 min since failed check

Alert escalation:
  RED for >15 minutes → P2 Alert in IT admin dashboard
  RED for >30 minutes → Optional Slack/email notification
  
Rationale: Silent extension failure = false sense of security.
Admin thinks data is protected. Extension broke 2 days ago. Samsung moment.
```

---

## 6. Module B — Prompt Injection Detection Engine

### 6.1 Phạm Vi & Use Case

> Prompt injection áp dụng cho **~15–20% SME** — cụ thể là những SME đang triển khai internal AI assistants (Slack bots sử dụng ChatGPT API, customer service bots, internal search). Không áp dụng cho SME chỉ dùng AI tools như người dùng cuối.

**Tại sao vẫn build cho v1:** Là competitive differentiator, chi phí thấp (build on OWASP open-source), và là foundation cho Enterprise tier. Được scoped là **Layer 1 rule-based** trong v1 — ML classifier (BERT) là v2.

### 6.2 B1 — OWASP Regex Library (Must-have)

```go
// infrastructure/dlp/injection_detector.go

// InjectionPattern represents a single detection rule
type InjectionPattern struct {
    ID          string
    Category    string  // 'jailbreak'|'data_exfil'|'system_prompt'|'role_override'
    Pattern     *regexp.Regexp
    Risk        RiskLevel
    Description string
    Source      string  // 'OWASP_LLM_TOP10'|'custom'|'community'
}

// 50+ patterns covering OWASP LLM Top 10:
var builtinPatterns = []InjectionPattern{
    // LLM01: Prompt Injection
    {ID: "LLM01-001", Category: "jailbreak",
     Pattern: regexp.MustCompile(`(?i)ignore (all |previous |prior |above |your )?(instructions?|prompts?|rules?|constraints?)`),
     Risk: HIGH, Source: "OWASP_LLM_TOP10"},
    
    {ID: "LLM01-002", Category: "role_override",
     Pattern: regexp.MustCompile(`(?i)(you are now|act as|pretend (to be|you are)|you('re| are) (now )?a?n? )(unrestricted|jailbroken|DAN|evil|unethical)`),
     Risk: CRITICAL, Source: "OWASP_LLM_TOP10"},
    
    {ID: "LLM01-003", Category: "system_prompt",
     Pattern: regexp.MustCompile(`(?i)(print|reveal|show|output|display|repeat|tell me) (your |the )?(system prompt|instructions|initial prompt|configuration)`),
     Risk: HIGH, Source: "OWASP_LLM_TOP10"},
    
    // LLM06: Sensitive Information Disclosure
    {ID: "LLM06-001", Category: "data_exfil",
     Pattern: regexp.MustCompile(`(?i)(extract|exfiltrate|send|forward|email|transmit) (all |the )?(data|information|records|database|files)`),
     Risk: CRITICAL, Source: "OWASP_LLM_TOP10"},
    
    // ... 45+ more patterns
}

// Server-push updates: patterns versioned as JSON, pushed via CDN
// Extension loads new patterns without reinstall
// Pattern version tracked in localStorage for rollback capability
```

### 6.3 B2 — Lakera Guard API Integration (Must-have cho v1)

```
DECISION: White-label Lakera Guard cho server-side prompt injection

Lý do:
  ● Lakera = production-hardened, continuously updated với novel attacks
  ● SMESec differentiator là Layer 3 context-aware scoring, không phải injection patterns
  ● Build custom = 6+ months để đạt Lakera accuracy level
  ● Unit economics: <$0.05/request viable tại SME scale
     (10 employees × 20 checks/day = 200 req/day = $3–10/mo/company)

Architecture:
  Browser extension (Tier 1 + 2 local) → FastAPI proxy → Lakera Guard API
  
  FastAPI proxy purpose:
  1. Tenant authentication + rate limiting
  2. Context enrichment (add user role, asset sensitivity from Track 1)
  3. Audit logging (log request metadata, NOT prompt content)
  4. Fallback to WASM-only if Lakera API unavailable

WASM Fallback (always built, regardless of Lakera):
  Model: BERT-tiny quantized → ONNX → WASM (17MB)
  Loaded: Lazy post-install, served from CDN with integrity hash
  Chrome Web Store: Extension package <2MB (WASM not bundled)
  Chrome 10MB limit: met (2MB package + 17MB lazy = within policy)
  Accuracy: >75% known patterns (vs >90% Lakera)
  Latency: 50–80ms (vs <300ms Lakera API)
```

### 6.4 B4 — Risk Response Action Engine (Must-have)

| Score Range | Action | User Experience | Admin Notification |
|---|---|---|---|
| **0–30** | Log only | No interruption | Weekly digest |
| **31–60** | Advisory + justification | Warning toast, can dismiss | Daily summary |
| **61–85** | Block + manager approval | Hard block modal, requires manager approve/deny | Real-time alert + manager email |
| **86–100** | Hard block + admin alert | Hard block, no override available | Immediate P1 alert |

```
Manager Approval Workflow (Score 61–85):

Employee's submission blocked
  ↓
SMESec sends Slack DM to manager (from Track 1 org chart)
  
  "🔴 Security Review Required
   [Employee name] attempted to send content to [AI tool]
   that may contain sensitive data.
   
   Risk score: 72/100
   Detection: Source code pattern + credentials
   
   [✅ APPROVE — Employee can re-submit]  [❌ DENY — Block access]
   
   Expires in: 30 minutes"
  
  Manager one-click approve/deny
  → Result relayed to employee browser within seconds
  → Full audit trail logged
```

---

## 7. Module C — Shadow AI Governance

### 7.1 C1 — OAuth AI Tool Inventory (Must-have)

```
DATA SOURCES:

1. Google Workspace Admin API
   Endpoint: GET /admin/directory/v1/users/{userId}/tokens
   Scope: admin.directory.user.security.readonly (service account)
   Cadence: Every 15 minutes (delta sync)
   What we get: app_name, scopes_granted, created_date, last_used
   
2. Microsoft 365 Graph API
   Endpoint: GET /v1.0/users/{userId}/oauth2PermissionGrants
   Scope: DelegatedPermissionGrant.ReadWrite.All (application permission)
   Cadence: Every 15 minutes + webhook for new grants
   What we get: clientId, scope, consentType, principalId
   
3. Slack Admin API (Enterprise Grid)
   Endpoint: GET /admin.apps.approved.list
   Scope: admin.apps:read
   Covers: Apps installed in workspace channels

CLASSIFICATION PIPELINE:

OAuth App Discovered
  ↓
AI Tool Registry lookup (maintained catalog, v1: 100+ AI tools, quarterly update)
  ├─ Match found: categorize (ai_assistant/ai_coding/ai_image/ai_search/ai_writing)
  └─ No match found: Unknown → flag for manual review
  ↓
Risk Scoring (C4):
  Risk = f(oauth_scopes, vendor_dpa, data_residency, certifications, incidents, app_age)
  ↓
IT Admin Dashboard update (<24h for new discoveries)
  + Alert for HIGH/CRITICAL risk tools
```

**AI Tool Registry (v1 — curated, 100+ tools):**

```
TIER: CRITICAL (auto-review on discovery)
  ● Tools with gmail.modify / drive.readwrite OAuth scopes
  ● Tools with no public DPA / privacy policy
  ● Tools with known data breach incidents
  Examples: [Various new AI startups with broad OAuth requests]

TIER: HIGH (attestation required)
  ● AI writing tools with email read access (Jasper + Gmail)
  ● AI coding assistants with repo access (unknown vendors)
  ● AI meeting transcription tools
  Examples: Jasper, Otter.ai (unverified configs), Loom AI

TIER: MEDIUM (logged, monthly report)
  ● AI writing (text only, no file access)
  ● AI image generators
  ● AI search tools
  Examples: ChatGPT (without broad OAuth), Midjourney, Perplexity

TIER: LOW / PRE-APPROVED (inventory only)
  ● Microsoft Copilot (M365 tenant = IT-controlled, data stays in tenant)
  ● GitHub Copilot (code only)
  ● Google Duet AI (Google Workspace tenant = IT-controlled)
  Examples: Microsoft Copilot, Google Duet, GitHub Copilot
```

### 7.2 C2 — AI Tool Attestation Workflow (Must-have)

```
PROBLEM: OAuth discovery covers "apps with OAuth" but misses:
  ● Employees using ChatGPT directly (login with email, no OAuth to company account)
  ● AI tools accessed via API keys employees created themselves
  ● AI tools used on personal devices for work purposes

SOLUTION: Quarterly attestation survey closes the "no-OAuth" blind spot

WORKFLOW:

Quarterly trigger (configurable: monthly for Enterprise):
  → IT admin reviews C1 OAuth inventory
  → System generates personalized survey per employee based on their role

Survey: "Please confirm your AI tool usage (last quarter)"

Section 1 — Detected tools (from OAuth inventory):
  "We detected the following AI tools linked to your company account.
   Please confirm your current usage:"
   ☑ Microsoft Copilot — [still using] / [no longer using]
   ☑ Jasper AI — [still using] / [no longer using]
   
Section 2 — Undetected tools (self-report):
  "Do you use any AI tools NOT listed above for work purposes?"
  [Text field — tool name + frequency + type of content]

Section 3 — Understanding check:
  "Have you read and understood the company AI usage policy?"
  [✅ Yes, I understand] [📄 View Policy]

Non-response after 5 business days:
  → Compliance gap finding created
  → Included in next SOC 2 review cycle

Discrepancy detection:
  Self-reported tool NOT in OAuth inventory → "Browser-only AI usage detected"
  → Added to Shadow AI findings
  → C3 telemetry cross-referenced
```

### 7.3 C4 — Risk Scoring & Policy Engine (Should-have)

```go
// application/shadow_ai/risk_scorer.go

type AIToolRiskScore struct {
    AppID              string
    AppName            string
    Category           AIToolCategory
    RiskScore          float64  // 0.0 – 1.0
    RiskTier           RiskTier // LOW | MEDIUM | HIGH | CRITICAL
    Factors            []RiskFactor
    PolicyAction       PolicyAction  // LOG | ALERT | ATTESTATION | AUTO_BLOCK
    ScoreUpdatedAt     time.Time
}

type RiskFactor struct {
    Name   string
    Weight float64
    Value  float64
    Score  float64 // Weight × Value
}

// Risk scoring formula:
// score = Σ(factor_weight × factor_value) / Σ(factor_weights)
//
// Factors & weights:
// ● oauth_scopes_sensitivity:     30% weight (mail.readonly = 0.6, drive.readwrite = 1.0)
// ● vendor_dpa_available:         20% weight (DPA found + signed = 0.0, none = 1.0)
// ● data_residency_compliance:    15% weight (EU cert = 0.0, unknown = 1.0)
// ● security_certifications:      15% weight (SOC 2 Type 2 + ISO 27001 = 0.0, none = 1.0)
// ● known_incidents:              10% weight (no incidents = 0.0, recent breach = 1.0)
// ● app_age_days:                  5% weight (>2yr = 0.0, <3mo = 1.0)
// ● user_count_in_tenant:          5% weight (higher = slightly higher risk)

// IT admin policy: "Auto-block any new AI tool with CRITICAL risk score"
type TenantAIPolicy struct {
    TenantID         string
    AutoBlockCritical bool    // default: false (dry-run first, then 2-step confirm)
    AutoAlertHigh    bool    // default: true
    RequireAttestation []RiskTier  // default: [HIGH, CRITICAL]
    AllowedTools     []string // pre-approved tool IDs
    BlockedTools     []string // explicitly blocked tool IDs
}
```

---

## 8. Module D — Deepfake Fraud Defense

### 8.1 D1 — Voice Deepfake Detection (Must-have, với legal gate cho EU)

**Đặc điểm kỹ thuật:**

```
INPUT:   Employee uploads audio: MP3/WAV/M4A, ≤60 giây
         OR records trực tiếp qua MediaRecorder API trong browser
         
ANALYSIS: Hive Moderation API (primary) / Resemble Detect (fallback)
         ● Audio NOT transmitted as raw file — sent as hash + metadata
         ● Raw audio: deleted within 60 seconds của analysis
         ● Zero audio retention on SMESec servers
         
OUTPUT:  3-tier result (NOT binary):
         ● "Highly likely authentic" (score 0.0–0.3)
         ● "Inconclusive — verify through other means" (score 0.3–0.7)
         ● "Likely synthetic — exercise caution" (score 0.7–1.0)
         
Rationale for non-binary output:
  Deepfake detection is probabilistic, not definitive.
  "FAKE" output = employee refuses to comply → potential real CEO annoyed.
  "Likely synthetic — exercise caution" + OOB verification = informed decision.
```

**GDPR Hard Gate:**
```
⚠️ LEGAL GATE: Voice = Biometric Data under GDPR Article 9

EU Deployment Requirements:
  ● GDPR Article 9 legal opinion commissioned Day 1 (Sprint 1, Week 1)
  ● Opinion must address: biometric processing basis, consent vs legitimate interest
  ● Employee-initiated only (never employer-triggered) — reduces legal risk
  ● Zero retention architecture (audio deleted <60s) — reduces processing scope
  
Deployment sequence:
  v1 (Month 6): US + UK + Australia + Singapore (no GDPR constraint)
  v1.1 (Month 8): EU deployment ONLY AFTER legal opinion clearance
  If opinion negative → EU employees use D2 (OOB verification) only, no audio analysis
```

**Accuracy Limitations (transparent to customers):**
```
Current deepfake detection accuracy benchmarks (2025):
  ● High-quality studio audio: >90% detection rate
  ● Compressed VoIP calls (<64kbps): ~70–75% detection rate
  ● Non-English speakers: limited benchmark data
  ● Short clips (<10 seconds): reduced accuracy (<65%)

SMESec commitments:
  1. Never present results as "definitely fake" — probabilistic framing always
  2. Always pair with D2 (OOB verification) — two independent signals
  3. Accuracy benchmarks published in transparency report (quarterly)
  4. Employee can view accuracy caveats in F1 transparency dashboard
```

### 8.2 D2 — Out-of-Band Verification Workflow (Must-have)

```
OUT-OF-BAND VERIFICATION ARCHITECTURE:

TRIGGER: Employee suspects a call/request is fraudulent
  (e.g., "My CEO just called me via Zoom to ask for a $50K wire transfer")
  
STEP 1: Employee opens SMESec Mobile App → "Verify Identity" → 3 taps
  Enter: "CEO - John Smith" / Select from contacts / Enter phone number
  
STEP 2: SMESec initiates dual-channel verification
  Channel A: Email to john.smith@company.com (from Track 1 HR/Directory integration)
             "Did you just contact [employee name] at [time] about [topic]?"
             [✅ YES, this was me]  [❌ NO, this was not me]
             (No SMESec account required for recipient — anonymous link)
             
  Channel B: SMS to CEO's registered phone (from Track 1 People directory)
             "SMESec Security: Verification code 847-291"
             (Code must be read back by alleged caller to employee)
  
STEP 3: Combined result within 5 minutes
  Case A: Email = "NOT ME" + Code NOT provided → "⚠️ LIKELY FRAUD — Do NOT proceed"
  Case B: Email = "YES" + Code PROVIDED correctly → "✅ VERIFIED — Identity confirmed"
  Case C: Ambiguous (no email response, code partial) → "⚠️ INCONCLUSIVE — Escalate to IT admin"
  
STEP 4: Audit log
  Full verification timeline: who triggered, what alleged caller, channels used, outcome
  Stored in PostgreSQL (tenant-scoped) + S3 Object Lock (90-day retention minimum)
  
STEP 5: If FRAUD confirmed → one-tap trigger Incident Playbook #6 (Deepfake Fraud Response)
  Playbook #6: Notify IT admin + Document in incident log + Email security team +
               Report to bank (payment fraud section) + 30-day monitoring flag on account

```

### 8.3 D3 — Video Deepfake Detection (Should-have, Sprint 8 go/no-go)

```
SCOPE: Upload-based (NOT real-time stream interception)
  ● MP4/MOV, up to 2 minutes
  ● Admin-only (not employee-facing)
  ● Use case: "Verify this Zoom recording before sharing in investor meeting"
  
VENDOR: Hive Moderation API (video analysis endpoint)
  ● Analysis: face boundaries, lighting consistency, blink patterns, temporal artifacts
  ● Response: deepfake_score + highlighted artifact timestamps
  ● Latency: <30 seconds for 2-minute clip
  
SPRINT 8 GO/NO-GO CRITERIA:
  ✅ GO:   Hive API accuracy >75% on internal test set + contract signed
  ❌ NO-GO: Defer to v1.1 (Month 8) — D2 OOB verification covers the gap

COMPLIANCE:
  Evidence package per verification: timestamp, analyzed file hash, result, artifacts
  Auto-delete after 90 days (unless pinned by IT admin as compliance evidence)
```

---

## 9. Module E — AI Phishing Defense

### 9.1 E1 — M365 Defender Phishing Alerts (Must-have cho M365 customers)

```
INTEGRATION ARCHITECTURE:

Microsoft Security Graph API
  Endpoint: GET /v1.0/security/alerts_v2
  Scope: SecurityEvents.Read.All
  Polling: Every 5 minutes (webhook preferred, fallback to polling)
  
Filter criteria:
  ● category: 'Phishing' OR 'Malware' OR 'Suspicious email activity'
  ● severity: 'high' OR 'medium'
  
Enrichment from Track 1 (Asset Inventory):
  Alert contains: affected_user_email → Look up in Asset DB
  Get: user's role, data_sensitivity_level, direct_reports, financial_access
  
Result: Enriched alert:
  "🎣 Phishing email detected for Alice Johnson (CFO — Financial data access: CRITICAL)
   Payload: AI-crafted spear-phish impersonating [Vendor name]
   Subject: Urgent: Invoice #INV-2847 overdue
   AI confidence: 94% AI-generated content
   
   [▶️ Launch Incident Playbook #3: Phishing Response]"

Value proposition vs raw M365 Defender alert:
  ● "Alice is your CFO with financial access" context = higher urgency prioritization
  ● One-click playbook launch (vs manual response coordination)
  ● Evidence logged to compliance record automatically
```

### 9.2 E2 — Email Authentication Posture (Should-have, Google Workspace)

```
SIGNALS COLLECTED (weekly via Google Workspace Admin SDK):
  ● DMARC: policy (none/quarantine/reject) + pass rate %
  ● DKIM: signing enabled + algorithm (RSA-1024 deprecation check)
  ● SPF: record present + alignment status
  
Remediation guidance (non-expert, action-oriented):
  "⚠️ Your DMARC policy is set to 'none' — emails can be spoofed.
   Action: Update DNS record to p=quarantine. Step-by-step guide →"
  
This is email POSTURE (prevention), not phishing DETECTION.
Framing to customers: "Make it harder for attackers to send fake emails pretending to be your company"
```

---

## 10. Module F — Employee Privacy & Transparency

### 10.1 F1 — Employee Transparency Dashboard (Must-have, required cho EU)

```
CONTENT (plain language, employee-facing):

╔══════════════════════════════════════════════════════════════╗
║     WHAT SMESEC DOES AND DOES NOT DO                         ║
╠══════════════════════════════════════════════════════════════╣
║                                                              ║
║  ✅ WE DO:                                                   ║
║  ● Scan text BEFORE you send it to AI tools (in browser)     ║
║  ● Alert you if your prompt contains sensitive company data  ║
║  ● Log which AI websites you visited (domain + date only)    ║
║  ● Check which AI tools have access to your company account  ║
║  ● Store that list of AI tool visits for 12 months           ║
║                                                              ║
║  ❌ WE DO NOT:                                               ║
║  ● Read, store, or transmit what you type into AI tools      ║
║  ● Monitor your personal browsing                            ║
║  ● Record your screen or keystrokes                          ║
║  ● Share your individual data with other companies           ║
║  ● Train AI models using your data                           ║
║                                                              ║
║  📊 YOUR RECENT ACTIVITY (last 10 events):                   ║
║  • May 28: 1 prompt auto-redacted (email address detected)  ║
║  • May 27: Jasper AI detected — awaiting IT attestation     ║
║  • May 26: No flagged activity                              ║
║                                                              ║
║  🔐 DATA RETENTION: 12 months for DLP events                 ║
║     Contact: security@yourcompany.com to request deletion    ║
╚══════════════════════════════════════════════════════════════╝

Accessibility:
  ● Always available from extension popup (one click)
  ● Always available from SMESec mobile app (settings > privacy)
  ● IT admin CANNOT hide or disable this page
  ● Available in: English, Vietnamese, French, Spanish, German (v1)
```

### 10.2 F2 — Employee Opt-in / Pause Capability (Must-have, EU consent model)

```
PAUSE FEATURE DESIGN:

Available durations: 15 / 30 / 60 minutes
(IT admin can restrict max duration for specific roles)

When paused:
  ● ZERO scanning, ZERO telemetry, ZERO backend transmission
  ● Extension icon shows grey (paused state)
  ● No logging of what employee does during pause

IT admin notification:
  ● "Employee X has paused monitoring for 30 min" — time only, NO reason logged
  ● Reason is NOT logged (employee privacy)
  ● IT admin can see pause frequency (weekly count) for coverage dashboard

EU Legal Basis:
  ● Monitoring employees = requires GDPR Article 6(1)(f) legitimate interest
  ● Pause capability = demonstrates proportionality
  ● Disclosed in Works Council / employee handbook
  ● France/Germany: May require specific consent or Works Council agreement (OQ-3)

IT Admin Configuration:
  ● Restrict pause for: CFO, CEO roles (financial + sensitive roles)
  ● Restriction is disclosed in F1 transparency dashboard ("Your role does not support pause")
  ● Maximum pause: configurable per role (default 60 min; some roles: 0 min)
```

---

## 11. Kiến Trúc Kỹ Thuật: Zero-Knowledge DLP

### 11.1 Nguyên Tắc Cốt Lõi

> **Cam kết thiết kế không thể nhượng bộ:** Nội dung mà nhân viên gõ vào AI tools **không bao giờ rời khỏi browser của họ**.

```
┌─────────────────────────────────────────────────────────────────────┐
│                    BROWSER (Sandboxed per Tab)                      │
│                                                                     │
│  Content Script (lives for full tab lifetime):                      │
│  ┌──────────────────────────────────────────────────────────────┐   │
│  │  ChatGPT.com page                                            │   │
│  │  ┌────────────────────────────────────────────────────────┐  │   │
│  │  │  [User textarea] ← Content Script intercepts here     │  │   │
│  │  │  "Our Q3 revenue is $4M, can you help me write..."    │  │   │
│  │  └────────────────────────────────────────────────────────┘  │   │
│  │                             │                                 │   │
│  │                    On submit event:                           │   │
│  │                             ▼                                 │   │
│  │                  Tier 1: Regex scan (<1ms)                    │   │
│  │                  Tier 2: WASM BERT scan (50–80ms)             │   │
│  │                             │                                 │   │
│  │              ┌──────────────┴──────────────┐                 │   │
│  │              ▼                             ▼                  │   │
│  │         SAFE → allow                  RISK → block           │   │
│  │                                   Show Redaction Review       │   │
│  └──────────────────────────────────────────────────────────────┘   │
│                                                                     │
│  What is sent to SMESec backend:                                    │
│  ✅ risk_tier: 'HIGH'                                               │
│  ✅ pattern_category: ['email_address', 'financial_data']           │
│  ✅ target_domain: 'chatgpt.com'                                    │
│  ✅ timestamp: 2026-05-28T10:23:41Z                                 │
│  ✅ tenant_id: UUID (hashed, not cleartext)                         │
│  ✅ action_taken: 'redaction_applied' | 'bypass_override'           │
│  ❌ Actual prompt content — NEVER transmitted                        │
│  ❌ Redacted tokens — NEVER transmitted                              │
└─────────────────────────────────────────────────────────────────────┘
                               │ Only metadata POST
                               ▼
┌─────────────────────────────────────────────────────────────────────┐
│                      SMESEC BACKEND (AWS)                           │
│                                                                     │
│  POST /api/v1/events/dlp → EventBridge → RDS (tenant-scoped)        │
│  POST /api/v1/detect/tier3 → Lakera Guard (async, no PII)           │
│  GET  /api/v1/ai-tools → C1 OAuth inventory                        │
│  POST /api/v1/verify/oob → D2 out-of-band verification             │
│                                                                     │
│  Data retention:                                                    │
│  ● Risk metadata: 7 years (S3 Object Lock — compliance evidence)    │
│  ● Prompt content: NEVER stored (zero-knowledge by architecture)    │
└─────────────────────────────────────────────────────────────────────┘
```

### 11.2 Chrome MV3 Architecture — Critical Technical Decision

```
PROBLEM: Chrome MV3 removes persistent background pages
         Service workers terminate after 5 minutes of inactivity
         → DLP scanning could silently stop working

SOLUTION: Content Script Architecture (không dùng Service Worker cho scanning)

Content Scripts:
  ● Run in page context (not in service worker)
  ● Active for entire tab lifetime
  ● No termination issue (lives with the page)
  ● Site-specific adapter pattern per AI tool domain

Service Worker role (minimal):
  ● Coordination only (extension icon, popup communication)
  ● Keepalive from content script every 25 seconds
  ● Content script does ALL scanning — service worker does NOT scan

Contingency architectures (if needed):
  Option A: Chrome Offscreen Documents (experimental) — fallback for WASM loading
  Option B: Server-side proxy mode — request goes through SMESec proxy (higher privacy concern)
  Option C: API-only mode — no real-time blocking, only logging (reduced effectiveness)

HARD GATE: MV3 persistence prototype must work by Sprint 1, Week 2 (June 6)
           If fails → pivot to Option B/C or delay launch
```

### 11.3 Site-Specific Adapter Pattern

```javascript
// extension/adapters/chatgpt.js
// ChatGPT DOM changes monthly — adapter pattern isolates the fragility

export class ChatGPTAdapter {
    // DOM selectors — version-tagged for change tracking
    // Version: 2026-05-01 (checked monthly in automated E2E test)
    
    getSubmitButton() {
        return document.querySelector('[data-testid="send-button"]')
               ?? document.querySelector('button[aria-label="Send prompt"]')
               ?? this.fallbackSubmitDetect();
    }
    
    getTextarea() {
        return document.querySelector('#prompt-textarea')
               ?? document.querySelector('[contenteditable="true"]');
    }
    
    injectRedactionUI(flaggedContent) {
        // Inject Redaction Review modal above ChatGPT's compose area
        // Positioned as overlay (not replacing ChatGPT's native UI)
    }
    
    // Canary health check — called hourly
    isHookActive() {
        return !!this.getSubmitButton() && !!this.getTextarea();
    }
}

// Automated E2E test suite (runs nightly against live AI tool URLs):
// → If ChatGPT changes DOM → test fails → alert extension engineer
// → Time to detect breakage: <30 minutes (via A3 canary)
```

---

## 12. Build vs Buy: Quyết Định Từng Component

| Component | Decision | Rationale | Cost |
|---|---|---|---|
| **Voice deepfake detection** | **Buy: Hive Moderation + Resemble Detect** | Training custom model: 50K+ voice samples needed. API cost viable: 5–20 checks/day/company = $0.30–0.60/mo | ~$0.01/check |
| **LLM DLP patterns (PII)** | **Augment open-source: Presidio + custom rules** | Presidio = production-grade, MIT license, 50+ entity types. Thêm: AWS credentials, GitHub tokens, source code patterns | $0 (FOSS) |
| **Shadow AI app catalog** | **Build + curate** | Không có off-the-shelf AI-specific risk scoring cho SME context. 100 top AI tools, quarterly maintenance | 0.5 engineer-day/quarter |
| **Prompt injection (browser)** | **Build: ONNX BERT-tiny WASM** | WASM viable for browser scanning. No server round-trip = better privacy + latency | Engineering cost only |
| **Prompt injection (server-side)** | **White-label: Lakera Guard** | Production-hardened, continuously updated. SMESec differentiator = Layer 3 context scoring, not pattern library | <$0.05/request |
| **AI phishing detection** | **Partner: M365 Defender (Graph API)** | Enterprise-grade detection already in M365. SMESec adds: context enrichment + playbook trigger | Included in M365 |
| **OAuth inventory** | **Build on Google/M365 APIs** | Core differentiator — AI-specific risk scoring không có competitor | 1 engineer × 2 sprints |
| **Attestation workflow** | **Build** | Domain-specific UX requirement. SurveyMonkey/Typeform không có HR/OAuth cross-reference | 0.5 engineer × 2 sprints |

---

## 13. Risk Register: 7 Rủi Ro Triển Khai Chính

| # | Rủi Ro | Xác Suất | Tác Động | Mitigation | Gate |
|---|---|---|---|---|---|
| **R1** | Chrome MV3 service worker termination | 🔴 60% | 🔴 Critical | Content Script architecture (không dùng SW cho scanning). Contingency: 3 alternatives defined. | **Sprint 1 Week 2 (June 6) — Hard gate** |
| **R2** | ChatGPT / Copilot DOM changes break extension hook | 🔴 High (monthly) | 🔴 High | Site-specific adapter pattern; nightly E2E tests; canary health checks; alert <30 min | Ongoing — automated |
| **R3** | Employee privacy trust erosion (EU adoption) | 🔴 High in EU | 🔴 High | Zero-knowledge architecture; open-source extension code; F1 transparency dashboard; F2 pause; Works Council engagement | Sprint 5 consent flow |
| **R4** | WASM 17MB exceeds Chrome Web Store 10MB limit | 🔴 65% | 🟠 Medium | Lazy-load WASM post-install via CDN với integrity hash; extension package <2MB | Sprint 12 Store submission |
| **R5** | GDPR Article 9 (voice = biometric) legal risk | 🟠 Medium | 🔴 High | Legal opinion Day 1; employee-initiated only; zero audio retention; US launch first | Legal opinion by Sprint 4 |
| **R6** | Lakera Guard unit economics not viable at SME scale | 🟠 35% | 🟠 Medium | Validate pricing Sprint 1; WASM-only fallback always built regardless of Lakera decision | **Sprint 1 decision (June 13)** |
| **R7** | Track 1 data dependency cho B3 context scoring | 🟠 45% | 🟠 Medium | Ship B3 với directory-role approximation in v1; full Track 1 integration Sprint 10 | Sprint 10 integration |

---

## 14. Delivery Sequence: 13-Sprint Roadmap

> Đây là roadmap riêng cho AI Governance Module, tích hợp với [delivery-plan.md](delivery-plan.md) tổng thể.

| Sprint | Tuần | Focus | Deliverables Chính | Critical Gate |
|---|---|---|---|---|
| **S1** | W1–2 | Foundation | B1 regex engine; MV3 prototype; FastAPI scaffold | **MV3 persistence by June 6** |
| **S2** | W3–4 | Extension scaffold | B1 server-push updates; extension site adapter (ChatGPT); C1 OAuth inventory start | Lakera Guard pricing decision (June 13) |
| **S3** | W5–6 | **A1 complete** | Prompt interception 100% ChatGPT + Copilot + Gemini + Claude; C1 complete | A1 intercepts 100% of ChatGPT/Copilot submits |
| **S4** | W7–8 | **A2 complete** | Redaction Review UI; B2 WASM BERT-tiny integration; F2 Pause capability | A2 default-redact flow end-to-end |
| **S5** | W9–10 | B3 + B4 + C2 | Context risk scoring; Action engine (4 tiers); Attestation workflow; E1 M365 alerts | B4 injection → correct response tier |
| **S6** | W11–12 | C4 + F1 + D2 | AI tool risk scoring; Transparency dashboard; OOB verification workflow | C4 policy enforcement live; D2 OOB verification |
| **S7** | W13–14 | **D1 Voice** | Voice deepfake detection (non-EU); Full pipeline integration test | D1: audio analysis <5s; legal opinion check |
| **S8** | W15–16 | D3 go/no-go; Enterprise deployment | D3 decision; Chrome Enterprise MDM push; Pen-test commissioned | Chrome Enterprise confirmed; D3 vendor LOI or defer |
| **S9** | W17–18 | **Pilot (2–3 SME)** | Real-world pilot; measure FP rates; collect feedback | FP <10% injection, <5% DLP in production |
| **S10** | W19–20 | B3 Track 1 integration; Pen-test remediation | Full context scoring with asset sensitivity; All Critical/High findings resolved | Pen-test clean |
| **S11** | W21–22 | Hardening + load test | 500-user load test; GDPR erasure automation; Documentation | Load test pass; 0 new Critical findings |
| **S12** | W23–24 | Chrome Web Store | Store submission; Go/no-go checklist | Extension approved for publication |
| **S13** | W25–26 | **Launch** | Graduated rollout 5 → 10 → full; D1/D2/B1 all live | Day 1/3/7 customer success check-ins |

**Critical Path:**
```
[S1: MV3 Gate] → [S3: A1 Interception] → [S4: A2 Redaction] → [S5: B4 Action Engine]
→ [S7: D1 Voice] → [S9: Pilot] → [S10: Pen-test] → [S12: Store Submission] → LAUNCH

Slack on critical path: ZERO — any slip cascades to launch date.
Single point of failure: Extension Engineer (owns A1, A2, A3, B4-ext, C3, F1, F2)
Contingency: If Extension Engineer behind at Sprint 3 → simplify F1 to static HTML, defer D3 + C2 UI
```

---

## 15. Key Performance Indicators

### 15.1 KPIs Kỹ Thuật (Platform Health)

| Metric | Target | Measurement |
|---|---|---|
| Extension availability (canary health) | >99.5% per AI domain | A3 health check, daily report |
| Tier 1 scan latency (p95) | <1ms | Extension performance.now() telemetry |
| Tier 2 WASM scan latency (p95) | <80ms | Extension performance.now() telemetry |
| OAuth discovery lag (new app → alert) | <24h | C1 delta sync |
| DLP bypass event → IT dashboard | <60s | Event pipeline latency |
| OOB verification result | <5 min | D2 step function duration |
| False positive rate (Tier 1 regex) | <1% on legitimate prompts | Pilot measurement (S9) |
| False positive rate (Tier 2 WASM) | <10% for CONFIDENTIAL data | Pilot measurement (S9) |

### 15.2 KPIs Sản Phẩm (Customer Value)

| Metric | Target | Timeline |
|---|---|---|
| Sensitive data submissions blocked / warned (% of total) | >5% (indicates real risk exists, not just security theater) | S9 pilot |
| % users engaging with Redaction Review (vs dismissing) | >80% | S9 pilot |
| % DLP bypass events using override vs cancelling | <15% override | S9 pilot, then monitor |
| Shadow AI tools discovered per tenant (avg) | >5 trong 30 ngày đầu (thị trường thực tế) | S9 pilot |
| IT admin Time-to-Respond to HIGH risk AI tool alert | <2h | S9 pilot |
| Employee satisfaction with transparency (F1) | NPS >30 | S9 pilot survey |
| EU adoption rate post-legal opinion | >60% EU pilot customers opt-in | v1.1 |

### 15.3 Business KPIs

| Metric | Target | Timeline |
|---|---|---|
| Customers using AI Governance module (% of total) | >60% (bundled in all tiers) | v1 |
| AI Governance as stated reason for purchase | >25% deal notes | v1 Q3 2026 |
| Monthly DLP events per customer (activity proxy) | >50 events/mo/customer = healthy usage | 90 days post-launch |
| Churn from customers with active DLP usage | <5%/year | v1 retrospective |

---

## 16. Open Questions & Hard Gates

### 16.1 Hard Gates (Block Progress nếu Không Resolve)

| # | Gate | Deadline | Block Quá | Owner |
|---|---|---|---|---|
| **OQ-1** | Chrome MV3 persistence prototype hoạt động với content script architecture | **June 6 (Sprint 1 W2)** | Toàn bộ extension delivery model | Extension Engineer |
| **OQ-2** | Lakera Guard pricing decision: viable <$0.05/request tại SME scale | **June 13 (Sprint 1 end)** | B2 architecture: Lakera vs WASM-only | PM + ML Engineer |
| **OQ-3** | EU employee consent model for extension monitoring (Works Council requirements) | Sprint 5 (August 8) | EU extension onboarding flow design | PM + Legal |
| **OQ-4** | GDPR Article 9 legal opinion: voice analysis for deepfake detection | Sprint 4–5 | D1 EU deployment timeline | PM + Legal |

### 16.2 Open Questions (Không Block Ngay Nhưng Cần Resolve)

| # | Question | Blocks | Deadline |
|---|---|---|---|
| **OQ-5** | Resemble Detect accuracy on non-English VoIP clips (<15 giây) | D1 accuracy claims in EU/APAC markets | Sprint 7 planning |
| **OQ-6** | Chrome Enterprise MDM managed deployment confirmation từ pilot customers | Onboarding at scale | End of Sprint 5 |
| **OQ-7** | Google Workspace phishing gap: Sublime Security partnership timeline | Go/no-go messaging to non-M365 customers | Before launch go/no-go |
| **OQ-8** | Vietnam-specific: PDPL (Personal Data Protection Law) compliance for D1 voice analysis | Vietnam market launch | Sprint 8 |

---

## 17. Tóm Tắt Điểm Khác Biệt

SMESec AI Governance Module là **giải pháp đầu tiên trên thị trường** kết hợp cả 5 capability sau tại SME pricing:

```
╔═════════════════════════════════════════════════════════════════════╗
║  SMESEC AI GOVERNANCE — 5 CAPABILITIES BUNDLED, SME-NATIVE          ║
╠═════════════════════════════════════════════════════════════════════╣
║                                                                     ║
║  1. REAL-TIME DLP (zero-knowledge, browser-side)                    ║
║     → Content never leaves browser                                  ║
║     → Blocks before data leaks, not after                           ║
║     → No IT setup required                                          ║
║                                                                     ║
║  2. SHADOW AI GOVERNANCE (OAuth inventory + risk scoring)           ║
║     → Discovers AI tools employees actually use                     ║
║     → Risk-scores based on data access + vendor posture             ║
║     → Policy enforcement: block/allow/attest                        ║
║                                                                     ║
║  3. DEEPFAKE FRAUD DEFENSE (OOB verification + audio analysis)      ║
║     → Protects against CEO voice impersonation fraud                ║
║     → Employee-triggered, 3 taps, <5 min result                     ║
║     → No deepfake product exists at SME pricing                     ║
║                                                                     ║
║  4. AI PHISHING CONTEXT ENRICHMENT (M365 Defender integration)      ║
║     → M365 security alerts enriched with asset context              ║
║     → One-click playbook trigger                                     ║
║                                                                     ║
║  5. EMPLOYEE TRANSPARENCY (privacy-respecting, EU-compliant)        ║
║     → Clear visibility into what is/is not monitored                ║
║     → Pause capability builds employee trust                        ║
║     → GDPR Article 9 + Works Council ready                          ║
║                                                                     ║
║  PRICING: ~$3–5/user/month (bundled — not 6 separate tools)         ║
║  SETUP: No IT expertise required                                     ║
║  COMPLIANCE: GDPR + SOC 2 + ISO 27001 ready                         ║
╚═════════════════════════════════════════════════════════════════════╝
```

**Closest competitor (Nudge Security)** chỉ có shadow AI discovery (không có DLP, không có deepfake, không có prompt injection) tại $4–8/user/mo.  
**Next closest (Prompt Security)** có browser DLP nhưng $15–30K/yr và cần IT/developer setup.

**SMESec gap:** Unified, affordable, non-expert-operable. Đây là vị trí không có competitor nào chiếm giữ vào tháng 5/2026.
