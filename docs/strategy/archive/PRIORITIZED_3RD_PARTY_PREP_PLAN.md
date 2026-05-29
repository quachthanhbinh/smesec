# SMESec — Prioritized Third-Party Preparation Plan

**Date:** 2026-05-29  
**Purpose:** Action-oriented preparation plan prioritizing long-lead-time integrations  
**Owner:** PM + Tech Lead

---

## Executive Summary

**Critical Finding:** 3 integrations have 2-8 week lead times and MUST start before Sprint 1:
- Google Workspace: 2-4 weeks → Start Week -3
- Microsoft 365: 3-6 weeks → Start Week -3  
- Penetration Test Vendor: 6-8 weeks → Start Week 8

**Total Preparation Time Saved:** Starting early prevents 4-8 weeks of sprint delays.

---

## Priority Matrix

| Priority | Service | Lead Time | Must Start | Blocks | Cost |
|----------|---------|-----------|------------|--------|------|
| 🔴 **P0** | Google Workspace | 2-4 weeks | Week -3 | Sprint 2 (W3-4) | Free |
| 🔴 **P0** | Microsoft 365 | 3-6 weeks | Week -3 | Sprint 3 (W5-6) | Free |
| 🔴 **P0** | Pentest Vendor | 6-8 weeks | Week 8 | Sprint 11 (W21-22) | $10-20K |
| 🟠 **P1** | Vanta | 2-3 weeks | Week 8 | Sprint 7 (W13-14) | $4-6K/yr |
| 🟠 **P1** | Lakera Guard | 1-2 weeks | Week 1 | Sprint 8 (W15-16) | ~$100/mo |
| 🟠 **P1** | Hive Moderation | 1-2 weeks | Week 1 | Sprint 10 (W19-20) | ~$500/mo |
| 🟡 **P2** | Slack Admin API | 1-2 weeks | Week 1 | Sprint 5 (W9-10) | Free |
| 🟡 **P2** | Apple Developer | 1-2 weeks | Week 1 | Sprint 6 (W11-12) | $99/yr |
| 🟢 **P3** | Chrome Web Store | 1 week | Week 10 | Sprint 12 (W23-24) | $5 |
| 🟢 **P3** | Google Play | Immediate | Week 1 | Sprint 6 (W11-12) | $25 |
| 🟢 **P3** | AWS IAM | <1 week | Week 1 | Sprint 5 (W9-10) | Free |
| 🟢 **P3** | Keycloak | <1 week | Week 1 | Sprint 1 (W1-2) | $50/mo |
| 🟢 **P3** | Cloudflare R2 | Immediate | Week 1 | Sprint 1 (W1-2) | $50/mo |

---

## 🔴 PRIORITY 0: CRITICAL PATH (Start BEFORE Sprint 1)

### 1. Google Workspace Admin SDK Access

**Why Critical:** 2-4 week verification process blocks Sprint 2 Google Workspace sync (core MVP feature)

**Timeline:**
```
Week -3: Submit verification → Week 2-4: Approval → Week 3-4: Production-ready
```

**Registration Steps:**

| Step | Time Required | Details |
|------|---------------|---------|
| 1. Create GCP project | 10 minutes | https://console.cloud.google.com |
| 2. Enable Admin SDK API | 5 minutes | APIs & Services → Enable APIs |
| 3. Create service account | 10 minutes | Download JSON key |
| 4. Prepare verification materials | 2-3 hours | See requirements below |
| 5. Submit OAuth consent screen | 30 minutes | **WAIT 2-4 WEEKS** |
| 6. Configure domain-wide delegation | 30 minutes | After approval |
| 7. Test with dev tenant | 1 hour | Verify sync works |

**Verification Requirements (prepare before submission):**
- ✅ Verified domain (DNS TXT record) — **BLOCKER if not ready**
- ✅ Privacy policy URL (public, accessible) — **BLOCKER**
- ✅ Terms of service URL (public, accessible) — **BLOCKER**
- ✅ App logo (512x512 PNG)
- ✅ Support email address
- ✅ Written justification for OAuth scopes:
  - `admin.directory.user.readonly` — "Read user directory for asset inventory"
  - `admin.directory.userschema.readonly` — "Read user schemas for classification"
  - `admin.reports.audit.readonly` — "Read audit logs for shadow IT detection"

**Typical Delays:**
- Google requests additional info: +2 weeks
- Domain verification issues: +1 week
- Privacy policy not accessible: +1 week

**Risk if Delayed:**
- >6 weeks: MVP delayed from W12 to W14-16
- Fallback: Unverified OAuth (100 user limit) for pilot only

**Owner:** Tech Lead (GCP setup) + PM (verification submission)

**Cost:** Free (GCP project free tier)

---

### 2. Microsoft 365 App Registration + Publisher Verification

**Why Critical:** 3-6 week verification process blocks Sprint 3 M365 sync (50% of SME market uses M365)

**Timeline:**
```
Week -3: Submit verification → Week 3-6: Approval → Week 5-6: Production-ready
```

**Registration Steps:**

| Step | Time Required | Details |
|------|---------------|---------|
| 1. Register app in Azure AD | 15 minutes | https://portal.azure.com |
| 2. Configure Graph API permissions | 20 minutes | User.Read.All, Application.Read.All, AuditLog.Read.All |
| 3. Create client secret | 5 minutes | Save securely |
| 4. Prepare verification materials | 2-3 hours | See requirements below |
| 5. Submit publisher verification | 30 minutes | **WAIT 3-6 WEEKS** |
| 6. Configure webhook subscriptions | 1 hour | After approval |
| 7. Test with dev M365 tenant | 1 hour | Unverified = 10 user limit |

**Verification Requirements (prepare before submission):**
- ✅ Verified domain (DNS TXT record) — **BLOCKER if not ready**
- ✅ Privacy policy URL (public, accessible) — **BLOCKER**
- ✅ Terms of service URL (public, accessible) — **BLOCKER**
- ✅ Microsoft Partner Network enrollment (optional but recommended)
- ✅ Publisher domain ownership proof

**Graph API Permissions Needed:**
- `User.Read.All` — Read all users
- `Application.Read.All` — Read all OAuth apps
- `AuditLog.Read.All` — Read audit logs
- `DeviceManagementManagedDevices.Read.All` — Read device inventory

**Typical Delays:**
- Microsoft requests additional documentation: +2 weeks
- Domain verification issues: +1 week
- Partner Network enrollment required: +1 week

**Risk if Delayed:**
- >8 weeks: v1 missing M365 support (50% market loss)
- Fallback: Unverified app (10 user limit) for pilot only

**Owner:** Tech Lead (Azure AD setup) + PM (verification submission)

**Cost:** Free (Azure AD free tier)

---

### 3. Penetration Test Vendor Selection + LOI

**Why Critical:** 6-8 week process from RFP to kickoff; LOI must be signed by Week 14 (hard gate)

**Timeline:**
```
Week 8: RFP sent → Week 10: Vendor selected → Week 14: LOI signed → Week 21: Pentest starts
```

**Selection Process:**

| Step | Time Required | Details |
|------|---------------|---------|
| 1. Prepare RFP document | 1 day | Scope: multi-tenant SaaS, Track 1+2, web+mobile+extension |
| 2. Send RFP to 3-5 vendors | 1 day | Recommended: Bishop Fox, NCC Group, Trail of Bits, Cure53, Cobalt |
| 3. Vendor responses | 1-2 weeks | **WAIT** |
| 4. Vendor evaluation | 1 week | Compare pricing, timeline, expertise |
| 5. Vendor selection | 1 day | Decision + notification |
| 6. LOI negotiation | 2 weeks | Scope, pricing, timeline |
| 7. **LOI signing** | 1 day | **HARD DEADLINE: Week 14** |
| 8. Scope agreement | 2 weeks | Detailed test plan |
| 9. Pentest kickoff | 1 day | Week 21 (Sprint 11) |

**RFP Requirements:**
- Scope: Multi-tenant SaaS application
- Components: Web app, mobile app (iOS/Android), browser extension
- Focus areas: Track 1 (deterministic) + Track 2 (ML/AI)
- Timeline: Must complete by Week 25 (before v1 launch Week 26)
- Deliverables: Findings report, remediation guidance, retest

**Vendor Evaluation Criteria:**
- Multi-tenant SaaS experience
- AI/ML security testing capability
- Browser extension security expertise
- Timeline fit (must start Week 21)
- Pricing ($10-20K range)
- Retest included

**Risk if Delayed:**
- LOI not signed by Week 14: Use backup vendor (2-week delay, pentest starts Week 23)
- Pentest not complete by Week 25: v1 delayed or ships with unresolved findings

**Owner:** PM (vendor selection) + Tech Lead (scope definition)

**Cost:** $10-20K

---

## 🟠 PRIORITY 1: HIGH IMPACT (Start Sprint 1)

### 4. Vanta Account Provisioning

**Why Important:** 2-3 week setup blocks Sprint 7 evidence collection (SOC 2 Type 1 requires 6-month observation window)

**Timeline:**
```
Week 8: Sign up → Week 11: Active → Week 13: Evidence collection begins
```

**Setup Steps:**

| Step | Time Required | Details |
|------|---------------|---------|
| 1. Sign up for Vanta Startup plan | 30 minutes | https://www.vanta.com |
| 2. Connect AWS account | 1 hour | IAM role with read-only permissions |
| 3. Connect GitHub organization | 30 minutes | OAuth app installation |
| 4. Select compliance frameworks | 15 minutes | SOC 2 Type 1, ISO 27001 |
| 5. Configure evidence collection | 2-3 hours | Define control mappings |
| 6. **Initial evidence scan** | **2-3 weeks** | **WAIT** |
| 7. Review compliance gaps | 2 hours | Create remediation plan |

**Prerequisites:**
- AWS account with CloudTrail enabled
- GitHub organization with admin access
- Compliance framework selection (SOC 2 Type 1, ISO 27001)

**Evidence Collection Requirements:**
- AWS: IAM policies, CloudTrail logs, Config snapshots
- GitHub: Branch protection, code review policies, commit signing
- Infrastructure: ECS task definitions, RDS encryption, KMS keys

**Risk if Delayed:**
- >3 weeks: Manual evidence collection required (20h/week PM workload)
- Evidence window insufficient: SOC 2 Type 1 delayed to v1.5

**Owner:** PM (account) + DevSecOps (connectors)

**Cost:** $4-6K/year (Startup plan)

---

### 5. Lakera Guard API Access (Prompt Injection Detection)

**Why Important:** 1-2 week approval blocks Sprint 8 prompt injection detection; pricing decision is hard gate at Sprint 1 end

**Timeline:**
```
Week 1: Submit → Week 2: API key + pricing decision → Week 15: Production-ready
```

**Registration Steps:**

| Step | Time Required | Details |
|------|---------------|---------|
| 1. Sign up for Lakera Guard | 20 minutes | https://www.lakera.ai/lakera-guard |
| 2. Submit API access request | 30 minutes | Use case: prompt injection detection for SME AI governance |
| 3. **Wait for approval** | **1-2 weeks** | **WAIT** |
| 4. Receive API key | Immediate | After approval |
| 5. Test prompt injection detection | 2 hours | Verify TPR >85%, FPR <2% |
| 6. **Pricing confirmation** | 1 day | **GO/NO-GO GATE: Week 2 (Sprint 1 end)** |

**API Access Request Details:**
- Use case: "Prompt injection detection for SME AI governance platform"
- Estimated volume: 50K-100K requests/month (50 customers × 1K-2K daily checks)
- Pricing tier: Pay-per-use (~$0.001/request)

**Pricing Decision (Hard Gate):**
- Target: <$0.05/request
- If >$0.05/request: Fallback to WASM-only BERT model (lower accuracy: TPR ~75% vs >85%)
- Decision deadline: Week 2 (Sprint 1 end)

**Risk if Delayed:**
- Pricing not viable: Sprint 8 prompt injection uses WASM BERT (lower accuracy)
- API access denied: No prompt injection detection in v1

**Owner:** ML Engineer #1 (integration) + PM (pricing negotiation)

**Cost:** ~$50-100/month (estimated at 50 customers)

---

### 6. Hive Moderation API Access (Deepfake Detection)

**Why Important:** 1-2 week approval blocks Sprint 10 deepfake detection (Track 2 feature)

**Timeline:**
```
Week 1: Submit → Week 2-3: API key → Week 19: Production-ready
```

**Registration Steps:**

| Step | Time Required | Details |
|------|---------------|---------|
| 1. Sign up for Hive Moderation | 20 minutes | https://thehive.ai/apis |
| 2. Submit API access request | 30 minutes | Use case: deepfake detection for fraud prevention |
| 3. **Wait for approval** | **1-2 weeks** | **WAIT** |
| 4. Receive API key | Immediate | After approval |
| 5. Test voice/video analysis | 2 hours | Verify >80% deepfake detection |
| 6. Confirm rate limits + pricing | 1 hour | Pay-per-use model |

**API Access Request Details:**
- Use case: "Deepfake voice/video detection for SME fraud prevention"
- Estimated volume: 20K-50K checks/month (50 customers × 400-1K monthly checks)
- Pricing: Pay-per-use (~$0.01/check)

**Testing Requirements:**
- Voice deepfake detection: >80% accuracy on test dataset
- Video deepfake detection: >80% accuracy on test dataset
- Latency: <5 seconds per check

**Risk if Delayed:**
- API access denied: Fallback to Resemble Detect (voice only) or defer to v1.5
- Pricing not viable: Defer deepfake detection to v1.5

**Owner:** ML Engineer #1 (integration) + PM (account)

**Cost:** ~$200-500/month (estimated at 50 customers)

---

## 🟡 PRIORITY 2: MEDIUM IMPACT (Start Sprint 1-2)

### 7. Slack App + Admin API Access

**Lead Time:** 1-2 weeks  
**Blocks:** Sprint 5 (W9-10) Slack integration  
**Cost:** Free (app), but requires Business+ tier for user management ($12.50/user/mo)

**Registration Steps:**
1. Create Slack app (30 minutes)
2. Configure OAuth scopes (20 minutes): `admin.users:read`, `admin.apps:read`, `admin.users:write`
3. Submit for Admin API access (30 minutes) → **WAIT 1-2 WEEKS**
4. Test with dev workspace (1 hour)
5. Implement tier detection (Sprint 5)

**⚠️ Important:** Admin API requires Business+ tier. Free/Pro tiers: read-only only.

---

### 8. Apple Developer Program + Google Play Console

**Lead Time:** 1-2 weeks (Apple), immediate (Google)  
**Blocks:** Sprint 6 (W11-12) mobile app beta (MVP)  
**Cost:** Apple $99/year, Google $25 one-time

**Apple Registration:**
1. Register Apple Developer Program (30 minutes) → **WAIT 1-2 WEEKS**
2. Create app listing (1 hour)
3. Prepare app store assets (2-3 hours)

**Google Registration:**
1. Register Google Play Console (20 minutes) → **IMMEDIATE**
2. Create app listing (1 hour)
3. Prepare app store assets (2-3 hours)

---

## 🟢 PRIORITY 3: LOW IMPACT (Start Sprint 3-5)

### 9. Chrome Web Store Developer Account

**Lead Time:** 1 week  
**Blocks:** Sprint 12 (W23-24) browser extension submission  
**Cost:** $5 one-time

**Registration:**
1. Register account (15 minutes)
2. Pay $5 fee (5 minutes) → **WAIT 1 WEEK**
3. Prepare privacy policy + ToS (2-3 hours)

**⚠️ Important:** Chrome Web Store review for security extensions can take 2-6 weeks. Submit v0 extension at Week 18 for early review buffer.

---

### 10. AWS IAM Cross-Account Access

**Lead Time:** <1 week (customer self-service)  
**Blocks:** Sprint 5 (W9-10) AWS discovery  
**Cost:** Free (IAM), CloudTrail + Config ~$10-50/mo per customer

**Setup:**
1. Design IAM role template (2 hours)
2. Create CloudFormation template (4 hours)
3. Test with dev account (1 hour)

---

### 11. Keycloak Self-Hosted + Cloudflare R2

**Lead Time:** <1 week / Immediate  
**Blocks:** Sprint 1 (W1-2) auth + audit logs  
**Cost:** $50/mo (Keycloak) + $50/mo (R2)

**Setup:**
- Keycloak: Deploy to ECS Fargate (1 day)
- Cloudflare R2: Sign up + enable (30 minutes)

---

## 📅 Critical Path Timeline

```
BEFORE PROJECT START:
Week -3:
  🔴 Google Workspace OAuth consent screen verification submission
  🔴 Microsoft 365 publisher verification submission
  ⏰ Prepare: verified domain, privacy policy, ToS, app logo

Week -2:
  ⏳ Follow up on Google + Microsoft verifications

Week -1:
  ✅ Confirm verifications in progress (no blockers)

SPRINT 1 (Week 1-2):
Week 1 Day 1-2:
  🟠 Lakera Guard API access request
  🟠 Hive Moderation API access request
  🟡 Slack app + Admin API access request
  🟡 Apple Developer Program registration
  🟡 Google Play Console registration
  🟢 Keycloak deployment
  🟢 Cloudflare R2 setup
  🟢 AWS IAM role template design

Week 2 (Sprint 1 end):
  🔴 Lakera Guard pricing decision (GO/NO-GO GATE)
  ⏳ Google Workspace verification approved (target)
  ⏳ Lakera API key received (target)
  ⏳ Hive API key received (target)

SPRINT 2-3 (Week 3-6):
Week 3-4:
  ✅ Google Workspace production-ready (Sprint 2)
  ✅ Apple Developer verified

Week 5-6:
  ✅ Microsoft 365 verification approved (target)
  ✅ Microsoft 365 production-ready (Sprint 3)

SPRINT 7-11 (Week 8-22):
Week 8:
  🔴 Pentest vendor RFP sent
  🟠 Vanta account setup begins

Week 10:
  🔴 Pentest vendor selected
  🟢 Chrome Web Store registration

Week 11:
  🟠 Vanta evidence collection active

Week 13 (Sprint 7):
  🟠 Vanta evidence collection running continuously

Week 14 (HARD DEADLINE):
  🔴 Pentest vendor LOI signed (NO EXCEPTIONS)

Week 21 (Sprint 11):
  🔴 Pentest kickoff
```

---

## 🚨 Hard Gates & Fallback Plans

| Gate | Deadline | Consequence | Fallback |
|------|----------|-------------|----------|
| **Google Workspace verification** | Week 2-4 | MVP delayed W12→W14-16 | Unverified OAuth (100 user limit) |
| **Microsoft 365 verification** | Week 3-6 | v1 missing M365 (50% market) | Unverified app (10 user limit) |
| **Lakera Guard pricing** | Week 2 | Lower accuracy prompt injection | WASM BERT (TPR ~75% vs >85%) |
| **Pentest vendor LOI** | Week 14 | Pentest delayed 2 weeks | Backup vendor pre-qualified |
| **Vanta setup** | Week 11 | Manual evidence collection | 20h/week PM workload |

---

## 💰 Total Cost Summary (Year 1, 50 customers)

| Category | Annual Cost |
|----------|-------------|
| **Critical Path** | $14,800-26,000 |
| - Pentest vendor | $10,000-20,000 |
| - Vanta | $4,800-6,000 |
| **Track 2 APIs** | $3,000-7,200 |
| - Hive Moderation | $2,400-6,000 |
| - Lakera Guard | $600-1,200 |
| **Infrastructure** | $1,200 |
| - Keycloak | $600 |
| - Cloudflare R2 | $600 |
| **One-time/Annual** | $129 |
| - Apple Developer | $99 |
| - Chrome Web Store | $5 |
| - Google Play | $25 |
| **TOTAL** | **$19,129-35,329** |

**Note:** Excludes customer-side costs (Slack Business+ tier, AWS CloudTrail/Config storage).

---

## ✅ Action Checklist

**PM Responsibilities:**
- [ ] **Week -3:** Submit Google Workspace OAuth verification
- [ ] **Week -3:** Submit Microsoft 365 publisher verification
- [ ] **Week 1:** Lakera Guard pricing negotiation
- [ ] **Week 2:** Lakera Guard GO/NO-GO decision
- [ ] **Week 8:** Send pentest vendor RFP
- [ ] **Week 8:** Sign up for Vanta
- [ ] **Week 14:** Sign pentest vendor LOI (HARD DEADLINE)

**Tech Lead Responsibilities:**
- [ ] **Week -3:** Create GCP project + service account
- [ ] **Week -3:** Register Azure AD app
- [ ] **Week 1:** Create Slack app
- [ ] **Week 1:** Design AWS IAM role template
- [ ] **Week 1:** Deploy Keycloak
- [ ] **Week 1:** Set up Cloudflare R2

**ML Engineer #1 Responsibilities:**
- [ ] **Week 1:** Submit Hive Moderation API request
- [ ] **Week 1:** Submit Lakera Guard API request
- [ ] **Week 2-3:** Test Hive + Lakera APIs
- [ ] **Week 2:** Provide Lakera pricing assessment

**DevSecOps Responsibilities:**
- [ ] **Week 8-10:** Configure Vanta AWS + GitHub connectors
- [ ] **Week 1-2:** Configure Keycloak HA (min 2 ECS tasks)
- [ ] **Week 1:** Test Cloudflare R2 S3 Object Lock

---

## 📊 Risk Summary

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| Google verification >6 weeks | Medium | Critical | Start Week -3, use unverified for pilot |
| Microsoft verification >8 weeks | Medium | Critical | Start Week -3, use unverified for pilot |
| Lakera pricing not viable | Low | Medium | WASM BERT fallback (lower accuracy) |
| Pentest LOI not signed by W14 | Low | High | Backup vendor pre-qualified |
| Vanta setup >3 weeks | Low | Medium | Manual evidence collection |

---

## 🎯 Success Criteria

**Week 2 (Sprint 1 end):**
- ✅ All API access requests submitted
- ✅ Lakera Guard pricing decision made
- ✅ Google + Microsoft verifications in progress

**Week 6 (Sprint 3 end):**
- ✅ Google Workspace production-ready
- ✅ Microsoft 365 production-ready
- ✅ All Track 2 API keys received

**Week 14 (Hard Gate):**
- ✅ Pentest vendor LOI signed

**Week 26 (v1 Launch):**
- ✅ All integrations production-ready
- ✅ Pentest complete, all Critical/High findings resolved
- ✅ SOC 2 Type 1 evidence collection active for 13 weeks
