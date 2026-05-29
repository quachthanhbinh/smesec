# SMESec — Third-Party Readiness Checklist (All Timeline Plans)

**Date:** 2026-05-29  
**Purpose:** Executive action plan for critical third-party preparations across all delivery timelines  
**Status:** URGENT — Several items must start BEFORE Sprint 1  
**Applies to:** Original (12mo) · 2x Adjusted (26mo) · Realistic Hiring (36mo+)

---

## 📊 Timeline Plan Selector

This checklist supports three delivery timeline scenarios:

| Plan | MVP | v1 | v1.5 | v2 | ML Eng #1 Joins | Best For |
|------|-----|----|----|----|--------------------|----------|
| **Original (12mo)** | W12 (M3) | W26 (M6) | W38 (M9) | W52 (M12) | Day 1 | Full team available Day 1, aggressive timeline |
| **2x Adjusted (26mo)** | W24 (M6) | W52 (M13) | W76 (M19) | W104 (M26) | Day 1 | Sustainable pace, 50-60% utilization |
| **Realistic Hiring (36mo+)** | M12 | M20 | M28 | M36+ | Month 8 | Solo TL start, progressive hiring |

**How to use this checklist:**
- Lead times are ABSOLUTE (e.g., Google verification = 2-4 weeks regardless of plan)
- Start dates are RELATIVE to your chosen plan
- Track 2 items (Hive, Lakera) timing depends heavily on when ML Eng #1 joins

---

## 🔴 CRITICAL PATH: Identity Provider Integrations

These have the longest lead times (2-8 weeks) and block core sync features.

### 1. Google Workspace Admin SDK Access

**Lead Time:** 2-4 weeks (can extend to 6 weeks if Google requests additional info)  
**Cost:** Free

**When to Start (by plan):**
- **Original (12mo):** Week -3 (3 weeks before Sprint 1) → Blocks Sprint 2 (W3-4)
- **2x Adjusted (26mo):** Week -3 (3 weeks before Sprint 1) → Blocks Sprint 3 (W5-6)
- **Realistic Hiring (36mo+):** Month 3 (when recruiting BE1/BE2) → Blocks Month 6 (when FE1 joins for dashboard)

**Registration Steps:**
1. ✅ Create GCP project (Day 1, 10 minutes)
2. ✅ Enable Admin SDK API (Day 1, 5 minutes)
3. ✅ Create service account + download JSON key (Day 1, 10 minutes)
4. ✅ Configure OAuth consent screen (Day 1, 30 minutes)
   - Requires: Verified domain (DNS TXT record)
   - Requires: Privacy policy URL (public)
   - Requires: Terms of service URL (public)
   - Requires: App logo (512x512 PNG)
   - Requires: Support email
   - Requires: Written justification for OAuth scopes
5. 🔴 **Submit for verification** (Day 2) → **WAIT 2-4 WEEKS**
6. ⏳ Configure domain-wide delegation (after verification)
7. ✅ Test with dev tenant (Week 2)

**Timeline by Plan:**

| Plan | Submit | Approval Target | Production-Ready | Risk if Delayed |
|------|--------|-----------------|------------------|-----------------|
| **Original** | Week -3 | Week 2-4 | Sprint 2 (W3-4) | MVP delayed W12→W14-16 |
| **2x Adjusted** | Week -3 | Week 2-4 | Sprint 3 (W5-6) | MVP delayed W24→W26-28 |
| **Realistic Hiring** | Month 3 | Month 4 | Month 6 (dashboard) | Foundation delayed M6→M8 |

**Owner:** Tech Lead (setup) + PM (verification submission)

**Realistic Hiring Note:** TL can start this solo during Month 3 while recruiting. No team needed for registration.

---

### 2. Microsoft 365 App Registration + Publisher Verification

**Lead Time:** 3-6 weeks (can extend to 8 weeks if Microsoft requests additional docs)  
**Cost:** Free

**When to Start (by plan):**
- **Original (12mo):** Week -3 (3 weeks before Sprint 1) → Blocks Sprint 3 (W5-6)
- **2x Adjusted (26mo):** Week -3 (3 weeks before Sprint 1) → Blocks Sprint 4 (W7-8)
- **Realistic Hiring (36mo+):** Month 3 (parallel with Google) → Blocks Month 6 (dashboard)

**Registration Steps:**
1. ✅ Register app in Azure AD (Day 1, 15 minutes)
2. ✅ Configure Graph API permissions (Day 1, 20 minutes)
   - User.Read.All
   - Application.Read.All
   - AuditLog.Read.All
   - DeviceManagementManagedDevices.Read.All
3. ✅ Create client secret (Day 1, 5 minutes)
4. 🔴 **Submit for publisher verification** (Day 2) → **WAIT 3-6 WEEKS**
   - Requires: Verified domain (DNS TXT record)
   - Requires: Privacy policy URL (public)
   - Requires: Terms of service URL (public)
   - Optional but recommended: Microsoft Partner Network enrollment
5. ✅ Test with dev M365 tenant (Week 2, unverified = 10 user limit)
6. ⏳ Configure webhook subscriptions (Week 3)

**Timeline by Plan:**

| Plan | Submit | Approval Target | Production-Ready | Risk if Delayed |
|------|--------|-----------------|------------------|-----------------|
| **Original** | Week -3 | Week 3-6 | Sprint 3 (W5-6) | v1 missing M365 (50% market) |
| **2x Adjusted** | Week -3 | Week 3-6 | Sprint 4 (W7-8) | v1 missing M365 (50% market) |
| **Realistic Hiring** | Month 3 | Month 4-5 | Month 6 (dashboard) | Foundation incomplete |

**Owner:** Tech Lead (setup) + PM (verification submission)

**Realistic Hiring Note:** TL can handle registration solo. No team dependency.

---

### 3. Vanta Account Provisioning

**Lead Time:** 2-3 weeks (account setup + connector config + initial scan)  
**Cost:** $4-6K/year (Startup plan)

**When to Start (by plan):**
- **Original (12mo):** Week 8 (5 weeks before Sprint 7) → Evidence collection starts W13
- **2x Adjusted (26mo):** Week 16 (10 weeks before Sprint 13) → Evidence collection starts W26
- **Realistic Hiring (36mo+):** Month 8 (when DevSecOps joins) → Evidence collection starts Month 10

**Registration Steps:**
1. ✅ Sign up for Vanta Startup plan (Day 1, 30 minutes)
2. ✅ Connect AWS account (Week 1, 1 hour)
   - Requires: AWS IAM role with read-only permissions
   - Requires: CloudTrail enabled
3. ✅ Connect GitHub organization (Week 1, 30 minutes)
4. ✅ Select compliance frameworks (Week 1, 15 minutes)
   - SOC 2 Type 1
   - ISO 27001
5. ✅ Configure evidence collection rules (Week 2, 2-3 hours)
6. 🔴 **Initial evidence scan** (Week 2-3) → **WAIT 2-3 WEEKS**
7. ✅ Review compliance gaps (Week 3, 2 hours)

**Timeline by Plan:**

| Plan | Start Setup | Active Target | Evidence Begins | v1 Launch | Observation Window |
|------|-------------|---------------|-----------------|-----------|-------------------|
| **Original** | Week 8 | Week 11 | Week 13 | Week 26 | 13 weeks (sufficient) |
| **2x Adjusted** | Week 16 | Week 22 | Week 26 | Week 52 | 26 weeks (excellent) |
| **Realistic Hiring** | Month 8 | Month 9 | Month 10 | Month 20 | 10 months (excellent) |

**Owner:** PM (account) + DevSecOps (connectors)

**Realistic Hiring Note:** Wait for DevSecOps (Month 10) to join before starting. TL shouldn't handle this solo.

---

### 4. Penetration Test Vendor Selection + LOI

**Lead Time:** 6-8 weeks (RFP → selection → LOI → scope → kickoff)  
**Cost:** $10-20K

**When to Start (by plan):**
- **Original (12mo):** Week 8 (RFP) → LOI signed Week 14 → Pentest starts Week 21
- **2x Adjusted (26mo):** Week 24 (RFP) → LOI signed Week 32 → Pentest starts Week 40
- **Realistic Hiring (36mo+):** Month 14 (RFP) → LOI signed Month 16 → Pentest starts Month 18

**Registration Steps:**
1. ✅ Send RFP to 3-5 vendors (Week 8, 1 day)
   - Recommended: Bishop Fox, NCC Group, Trail of Bits, Cure53, Cobalt
   - Scope: Multi-tenant SaaS, Track 1 + Track 2, web + mobile + browser extension
2. ✅ Vendor selection (Week 10, 1 week)
3. 🔴 **Sign LOI** (Week 14) → **HARD DEADLINE, NO EXCEPTIONS**
4. ✅ Scope agreement (Week 15-16, 2 weeks)
5. ✅ Pentest kickoff (Week 21)

**Timeline by Plan:**

| Plan | RFP Sent | Vendor Selected | LOI Signed (HARD GATE) | Pentest Kickoff | v1 Launch |
|------|----------|-----------------|------------------------|-----------------|-----------|
| **Original** | Week 8 | Week 10 | Week 14 | Week 21 | Week 26 (5-week buffer) |
| **2x Adjusted** | Week 24 | Week 28 | Week 32 | Week 40 | Week 52 (12-week buffer) |
| **Realistic Hiring** | Month 14 | Month 15 | Month 16 | Month 18 | Month 20 (2-month buffer) |

**Owner:** PM (vendor selection) + Tech Lead (scope)

**Realistic Hiring Note:** PM joins Month 11, can handle RFP from Month 14. TL provides technical scope.

---

## 🟠 HIGH PRIORITY: Track 1 Integrations

These have 1-2 week lead times and block Track 1 features.

### 5. Slack App + Admin API Access

**Lead Time:** 1-2 weeks  
**Cost:** Free (app), but requires Business+ tier for user management ($12.50/user/mo)

**When to Start (by plan):**
- **Original (12mo):** Week 1 (Sprint 1) → Blocks Sprint 5 (W9-10)
- **2x Adjusted (26mo):** Week 1 (Sprint 1) → Blocks Sprint 6 (W11-12)
- **Realistic Hiring (36mo+):** Month 6 (when FE1 joins) → Blocks Month 9 (Slack integration)

**Registration Steps:**
1. ✅ Create Slack app (Day 1, 30 minutes)
2. ✅ Configure OAuth scopes (Day 1, 20 minutes)
   - admin.users:read
   - admin.apps:read
   - auditlogs:read
   - admin.users:write (Business+ only)
3. 🔴 **Submit for Admin API access** (Day 2) → **WAIT 1-2 WEEKS**
4. ✅ Test with dev Slack workspace (Week 2)
5. ✅ Implement tier detection (Sprint 5)

**Timeline:**
- Submit: Week 1 (Sprint 1)
- Approval: Week 2-3
- Production-ready: Week 9 (Sprint 5)

**⚠️ Important:** Slack Admin API requires Business+ tier. Free/Pro tiers: read-only only. UI must detect tier and show warning.

**Owner:** Tech Lead

---

## 🔵 TRACK 2 PRIORITY: ML/AI Integrations

**⚠️ CRITICAL:** These depend on ML Engineer #1 joining. Timing varies significantly by plan.

### 6. Hive Moderation API Access (Deepfake Detection)

**Lead Time:** 1-2 weeks  
**Cost:** Pay-per-use (~$0.01/check), estimated $200-500/mo at 50 customers

**When to Start (by plan):**
- **Original (12mo):** Week 1 (ML Eng #1 joins Day 1) → Blocks Sprint 10 (W19-20)
- **2x Adjusted (26mo):** Week 1 (ML Eng #1 joins Day 1) → Blocks Sprint 16 (W31-32)
- **Realistic Hiring (36mo+):** Month 8 (when ML Eng #1 joins) → Blocks Month 16 (deepfake feature)

**Registration Steps:**
1. ✅ Sign up for Hive Moderation (Day 1, 20 minutes)
2. 🔴 **Submit API access request** (Day 2) → **WAIT 1-2 WEEKS**
   - Requires: Use case description (deepfake detection for fraud prevention)
   - Requires: Estimated volume (checks/month)
3. ⏳ Receive API key (Week 2)
4. ✅ Test voice/video analysis (Week 3, 2 hours)
5. ✅ Confirm rate limits + pricing (Week 3)

**Timeline by Plan:**

| Plan | ML Eng #1 Joins | Submit | API Key | Production-Ready |
|------|-----------------|--------|---------|------------------|
| **Original** | Day 1 | Week 1 | Week 2-3 | Sprint 10 (W19-20) |
| **2x Adjusted** | Day 1 | Week 1 | Week 2-3 | Sprint 16 (W31-32) |
| **Realistic Hiring** | **Month 8** | **Month 8** | **Month 9** | **Month 16** |

**Owner:** ML Engineer #1 (integration) + PM (account)

**Realistic Hiring Note:** CANNOT start until ML Eng #1 joins Month 8. This is an 8-month delay vs original plan.

---

### 7. Lakera Guard API Access (Prompt Injection Detection)

**Lead Time:** 1-2 weeks  
**Cost:** ~$0.001/request, estimated $50-100/mo at 50 customers

**When to Start (by plan):**
- **Original (12mo):** Week 1 (ML Eng #1 joins Day 1) → Blocks Sprint 8 (W15-16)
- **2x Adjusted (26mo):** Week 1 (ML Eng #1 joins Day 1) → Blocks Sprint 14 (W27-28)
- **Realistic Hiring (36mo+):** Month 8 (when ML Eng #1 joins) → Blocks Month 14 (prompt injection)

**Registration Steps:**
1. ✅ Sign up for Lakera Guard (Day 1, 20 minutes)
2. 🔴 **Submit API access request** (Day 2) → **WAIT 1-2 WEEKS**
   - Requires: Use case description (prompt injection detection)
   - Requires: Estimated volume (requests/month)
   - Requires: Pricing tier selection
3. ⏳ Receive API key (Week 2)
4. ✅ Test prompt injection detection (Week 3, 2 hours)
5. 🔴 **Confirm pricing viability** (Week 3) → **GO/NO-GO DECISION BY SPRINT 1 END (WEEK 2)**

**Timeline by Plan:**

| Plan | ML Eng #1 Joins | Submit | Pricing Decision (HARD GATE) | Production-Ready |
|------|-----------------|--------|------------------------------|------------------|
| **Original** | Day 1 | Week 1 | Week 2 (S1 end) | Sprint 8 (W15-16) |
| **2x Adjusted** | Day 1 | Week 1 | Week 2 (S1 end) | Sprint 14 (W27-28) |
| **Realistic Hiring** | **Month 8** | **Month 8** | **Month 8** | **Month 14** |

**⚠️ Critical:** If pricing >$0.05/request → fallback to WASM-only BERT model (lower accuracy, no API cost)

**Owner:** ML Engineer #1 (integration) + PM (pricing negotiation)

**Realistic Hiring Note:** CANNOT start until ML Eng #1 joins Month 8. Pricing decision delayed 8 months.

---

### 8. Apple Developer Program + Google Play Console

**Lead Time:** 1-2 weeks (Apple), immediate (Google)  
**Cost:** Apple $99/year, Google $25 one-time

**When to Start (by plan):**
- **Original (12mo):** Week 1 (Sprint 1) → Blocks Sprint 6 (W11-12, MVP)
- **2x Adjusted (26mo):** Week 1 (Sprint 1) → Blocks Sprint 10 (W19-20, MVP)
- **Realistic Hiring (36mo+):** Month 7 (when Flutter joins) → Blocks Month 12 (MVP)

**Registration Steps:**

**Apple:**
1. ✅ Register Apple Developer Program (Day 1, 30 minutes)
2. 🔴 **Apple account verification** (Day 2-14) → **WAIT 1-2 WEEKS**
3. ✅ Create app listing (Week 2, 1 hour)
4. ✅ Prepare app store assets (Week 3, 2-3 hours)

**Google:**
1. ✅ Register Google Play Console (Day 1, 20 minutes) → **IMMEDIATE**
2. ✅ Create app listing (Week 2, 1 hour)
3. ✅ Prepare app store assets (Week 3, 2-3 hours)

**Timeline by Plan:**

| Plan | Flutter Joins | Submit | Apple Verified | Production-Ready |
|------|---------------|--------|----------------|------------------|
| **Original** | Week 1 | Week 1 | Week 2-3 | Sprint 6 (W11-12, MVP) |
| **2x Adjusted** | Week 1 | Week 1 | Week 2-3 | Sprint 10 (W19-20, MVP) |
| **Realistic Hiring** | **Month 7** | **Month 7** | **Month 8** | **Month 12 (MVP)** |

**Owner:** Flutter Engineer + PM (listings)

**Realistic Hiring Note:** Wait for Flutter Engineer (Month 7). TL shouldn't handle mobile app store setup.

---

## 🟡 MEDIUM PRIORITY: Start Sprint 3-5

### 9. Chrome Web Store Developer Account
**Lead Time:** 1 week  
**Blocks:** Sprint 12 (Week 23-24) — Browser extension submission  
**Cost:** $5 one-time

**Registration Steps:**
1. ✅ Register Chrome Web Store developer account (Day 1, 15 minutes)
2. ✅ Pay $5 registration fee (Day 1, 5 minutes)
3. 🔴 **Account verification** (Day 2-7) → **WAIT 1 WEEK**
4. ✅ Prepare privacy policy + terms of service (Week 2, 2-3 hours)
5. ✅ Test dummy extension submission (Week 3, 1 hour)

**Timeline:**
- Register: Week 10 (before Sprint 12)
- Verified: Week 11
- Production submission: Week 23 (Sprint 12)

**⚠️ Important:** Chrome Web Store review for security extensions can take 2-6 weeks. Submit stripped-down v0 extension at Week 18 (Sprint 8) for early review buffer.

**Owner:** Frontend Engineer #2 + PM (legal docs)

---

### 10. AWS IAM Cross-Account Access
**Lead Time:** <1 week (customer self-service)  
**Blocks:** Sprint 5 (Week 9-10) — AWS IAM discovery  
**Cost:** Free (IAM), CloudTrail + Config storage ~$10-50/mo per customer

**Registration Steps:**
1. ✅ Design IAM role template (Day 1, 2 hours)
   - Read-only permissions
   - CloudTrail access
   - Config access
2. ✅ Document customer setup instructions (Day 1, 2 hours)
3. ✅ Create CloudFormation template (Week 2, 4 hours)
4. ✅ Test with dev AWS account (Week 3, 1 hour)

**Timeline:**
- Design: Week 1 (Sprint 1)
- Template: Week 2
- Production-ready: Week 9 (Sprint 5)

**Owner:** Tech Lead (design) + DevSecOps (CloudFormation)

---

## 🟢 LOW PRIORITY: Can Start Later

### 11. Keycloak Self-Hosted Setup
**Lead Time:** <1 week  
**Blocks:** Sprint 1 (Week 1-2) — Auth  
**Cost:** ~$50/mo (ECS compute)

**Setup Steps:**
1. ✅ Deploy Keycloak to ECS Fargate (Sprint 1, 1 day)
2. ✅ Configure PostgreSQL database (Sprint 1, 2 hours)
3. ✅ Configure OIDC/SAML (Sprint 1, 4 hours)
4. ✅ Test Google + M365 federation (Sprint 1, 2 hours)

**Timeline:**
- Deploy: Week 1 (Sprint 1)
- Production-ready: Week 2 (Sprint 1)

**⚠️ Important:** Keycloak requires HA config (min 2 ECS tasks, JWKS caching). See R-C6.

**Owner:** Tech Lead + DevSecOps

---

### 12. Cloudflare R2 Account
**Lead Time:** Immediate  
**Blocks:** Sprint 1 (Week 1-2) — Audit log storage  
**Cost:** ~$50/mo

**Setup Steps:**
1. ✅ Sign up for Cloudflare (Day 1, 15 minutes)
2. ✅ Enable R2 storage (Day 1, 5 minutes)
3. ✅ Create S3-compatible API credentials (Day 1, 10 minutes)
4. ✅ Test S3 Object Lock compatibility (Week 1, 1 hour)

**Timeline:**
- Setup: Week 1 (Sprint 1)
- Production-ready: Week 1 (Sprint 1)

**Owner:** DevSecOps

---

## 📋 Summary: Critical Path Timeline

```
Week -3 (BEFORE PROJECT START):
  🔴 Google Workspace OAuth consent screen verification submission
  🔴 Microsoft 365 App Registration + Publisher verification submission
  
Week 1 (Sprint 1 Day 1-2):
  🟠 Slack app + Admin API access request
  🟠 Hive Moderation API access request
  🟠 Lakera Guard API access request
  🟠 Apple Developer Program registration
  🟠 Google Play Console registration
  🟢 Keycloak deployment
  🟢 Cloudflare R2 setup
  🟡 AWS IAM role template design

Week 2 (Sprint 1 end):
  🔴 Lakera Guard pricing decision (GO/NO-GO GATE)
  ⏳ Google Workspace verification approved (target)
  ⏳ Hive API key received (target)
  ⏳ Lakera API key received (target)

Week 3-4 (Sprint 2):
  ✅ Google Workspace production-ready
  ✅ Apple Developer verified

Week 5-6 (Sprint 3):
  ✅ Microsoft 365 verification approved (target)
  ✅ Microsoft 365 production-ready

Week 8 (before Sprint 7):
  🔴 Pentest vendor RFP sent
  🔴 Vanta account setup begins

Week 10:
  🔴 Pentest vendor selected
  🟡 Chrome Web Store developer account registration

Week 11:
  🔴 Vanta evidence collection active

Week 13 (Sprint 7):
  🔴 Vanta evidence collection running continuously (SOC 2 Type 1 window begins)

Week 14 (HARD DEADLINE):
  🔴 Pentest vendor LOI signed (NO EXCEPTIONS)

Week 21 (Sprint 11):
  🔴 Pentest kickoff
```

---

## ⚠️ Hard Gates & No-Go Conditions

| Gate | Deadline | Consequence if Missed | Fallback |
|------|----------|----------------------|----------|
| **Google Workspace verification** | Week 2-4 (target) | S2 blocked → MVP delayed to W14-16 | Unverified OAuth (100 user limit) for pilot |
| **Microsoft 365 publisher verification** | Week 3-6 (target) | S3 blocked → v1 missing M365 (50% market) | Unverified app (10 user limit) for pilot |
| **Lakera Guard pricing decision** | Week 2 (S1 end) | S8 prompt injection delayed or lower accuracy | WASM-only BERT (TPR ~75% vs >85%) |
| **Pentest vendor LOI signed** | Week 14 | Pentest delayed → v1 delayed or ships with findings | Backup vendor (2-week delay) |
| **Vanta setup active** | Week 11 | SOC 2 Type 1 evidence insufficient | Manual evidence collection (higher PM workload) |

---

## 💰 Cost Summary (Year 1, 50 customers)

| Service | Cost Model | Annual Cost |
|---------|------------|-------------|
| Google Workspace OAuth | Free | $0 |
| Microsoft 365 Publisher | Free | $0 |
| Slack Admin API | Free (app) | $0 |
| AWS IAM | Free | $0 |
| Vanta | Subscription | $4,800 |
| Hive Moderation | Pay-per-use | $2,400-6,000 |
| Lakera Guard | Pay-per-use | $600-1,200 |
| Chrome Web Store | One-time | $5 |
| Apple Developer | Annual | $99 |
| Google Play Console | One-time | $25 |
| Keycloak (self-hosted) | Infrastructure | $600 |
| Cloudflare R2 | Usage-based | $600 |
| **TOTAL** | | **~$9,329-13,329** |

**Note:** This excludes customer-side costs (Slack Business+ tier, AWS CloudTrail/Config storage).

---

## ✅ Pre-Sprint 1 Checklist

**PM Responsibilities:**
- [ ] Google Workspace OAuth consent screen verification submission (Week -3)
- [ ] Microsoft 365 publisher verification submission (Week -3)
- [ ] Vanta account signup (Week 8)
- [ ] Pentest vendor RFP (Week 8)
- [ ] Pentest vendor LOI signing (Week 14 hard deadline)
- [ ] Lakera Guard pricing negotiation (Week 1-2)

**Tech Lead Responsibilities:**
- [ ] GCP project + service account creation (Week -3)
- [ ] Azure AD app registration (Week -3)
- [ ] Slack app creation (Week 1)
- [ ] AWS IAM role template design (Week 1)
- [ ] Keycloak deployment (Week 1)
- [ ] Cloudflare R2 setup (Week 1)

**ML Engineer #1 Responsibilities:**
- [ ] Hive Moderation API access request (Week 1)
- [ ] Lakera Guard API access request (Week 1)
- [ ] Hive + Lakera API testing (Week 2-3)

**DevSecOps Responsibilities:**
- [ ] Vanta AWS + GitHub connector setup (Week 8-10)
- [ ] Keycloak HA configuration (Week 1-2)
- [ ] Cloudflare R2 S3 Object Lock testing (Week 1)

---

## 🚨 Action Required NOW

**If you are reading this before Sprint 1 starts:**

1. **IMMEDIATELY (Week -3):**
   - Submit Google Workspace OAuth consent screen verification
   - Submit Microsoft 365 publisher verification
   - Both require: verified domain, privacy policy URL, terms of service URL

2. **Week 1 Day 1-2:**
   - Submit all API access requests (Slack, Hive, Lakera)
   - Register Apple Developer Program
   - Register Google Play Console

3. **Week 2 (Sprint 1 end):**
   - Make Lakera Guard pricing decision (GO/NO-GO)

**Total lead time savings by starting early:** 4-8 weeks of potential sprint delays avoided.
