# SMESec — 3rd-Party Integration Preparation Plan

**Date:** 2026-05-29  
**Purpose:** Prioritize long-lead-time 3rd-party registrations and API access to prevent sprint delays  
**Owner:** PM + Tech Lead  
**Related:** [04-delivery-plan-original.md](04-delivery-plan-original.md) · [06-delivery-plan-adjusted-2x.md](06-delivery-plan-adjusted-2x.md) · [07-delivery-plan-realistic-hiring.md](07-delivery-plan-realistic-hiring.md) · [10-feasibility-assessment-and-remediation-plan.md](10-feasibility-assessment-and-remediation-plan.md)

---

## ⚠️ Timeline Context

This document covers **all three delivery timeline scenarios**. Lead times are ABSOLUTE (e.g., Google verification = 2-4 weeks regardless of plan). Start dates are RELATIVE to the chosen plan.

| Plan | MVP | v1 | v1.5 | v2 | ML Eng #1 Joins | Best For |
|------|-----|----|----|----|--------------------|----------|
| **Original (12mo)** | W12 (M3) | W26 (M6) | W38 (M9) | W52 (M12) | Day 1 | Full team available Day 1, aggressive timeline |
| **2x Adjusted (26mo)** | W24 (M6) | W52 (M13) | W76 (M19) | W104 (M26) | Day 1 | Sustainable pace, 50-60% utilization |
| **Realistic Hiring (36mo+)** | M12 | M20 | M28 | M36+ | Month 8 | Solo TL start, progressive hiring |

**Track 2 items (Hive, Lakera) timing depends heavily on when ML Eng #1 joins.**

---

## Executive Summary

Several 3rd-party integrations have **2-8 week lead times** from registration to production-ready API access. Starting these preparations **before Sprint 1** is critical to avoid blocking sprints S2-S5.

**Critical Finding:** 3 integrations have the longest lead times and MUST start before Sprint 1:
1. Google Workspace Admin SDK (2-4 weeks to production OAuth consent)
2. Microsoft 365 App Registration (3-6 weeks for admin consent + verification)
3. Penetration Test Vendor (6-8 weeks RFP → LOI → kickoff)

**Total Preparation Time Saved:** Starting early prevents **4-8 weeks** of potential sprint delays.

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

## Priority 1: CRITICAL PATH — Must Start Before Project Kick-off

### 1.1 Google Workspace Admin SDK Access

| Item | Details |
|------|---------|
| **What's needed** | GCP Project + Service Account + Domain-wide Delegation + OAuth Consent Screen verification |
| **Lead time** | **2-4 weeks** (OAuth consent screen verification by Google) |
| **Blocks** | Sprint 2 (W3-4) — Google Workspace sync is S2 deliverable |
| **Cost** | Free (GCP project), but requires verified domain |
| **Steps** | 1. Create GCP project (Day 1)<br>2. Enable Admin SDK API (Day 1)<br>3. Create service account + JSON key (Day 1)<br>4. Configure OAuth consent screen (Day 1)<br>5. **Submit for verification** (Day 2) → **2-4 weeks wait**<br>6. Configure domain-wide delegation in Google Workspace Admin (after verification)<br>7. Test with dev tenant (W2) |
| **Risk if delayed** | S2 cannot demo Google Workspace sync → MVP delayed to W14-16 |
| **Owner** | Tech Lead (GCP setup) + PM (Google verification submission) |
| **Documentation** | https://developers.google.com/workspace/guides/create-credentials |

**Action Items:**
- [ ] **Week -2 (before project start):** Create GCP project, enable APIs, create service account
- [ ] **Week -2:** Submit OAuth consent screen for verification (include: app name, logo, privacy policy URL, support email)
- [ ] **Week 1 (S1):** While waiting for verification, test with internal/unverified scope on dev tenant
- [ ] **Week 2-4:** Receive verification approval → configure domain-wide delegation
- [ ] **Week 3 (S2):** Production-ready for S2 sprint

**⚠️ Critical:** Google's verification process can take **up to 6 weeks** if they request additional information. Start this **2 weeks before Sprint 1**.

---

### 1.2 Microsoft 365 App Registration + Admin Consent

| Item | Details |
|------|---------|
| **What's needed** | Azure AD App Registration + Microsoft Graph API permissions + Admin consent + Publisher verification |
| **Lead time** | **3-6 weeks** (Publisher verification by Microsoft) |
| **Blocks** | Sprint 3 (W5-6) — M365 sync is S3 deliverable |
| **Cost** | Free (Azure AD), but requires verified domain + Microsoft Partner Network enrollment (optional but recommended) |
| **Steps** | 1. Register app in Azure AD (Day 1)<br>2. Configure Graph API permissions (Day 1)<br>3. Create client secret (Day 1)<br>4. **Submit for publisher verification** (Day 2) → **3-6 weeks wait**<br>5. Test with dev M365 tenant (W2)<br>6. Configure webhook subscriptions (W3) |
| **Risk if delayed** | S3 cannot integrate M365 → v1 missing M365 support (50% of SME market) |
| **Owner** | Tech Lead (Azure AD setup) + PM (Microsoft verification submission) |
| **Documentation** | https://learn.microsoft.com/en-us/graph/auth-register-app-v2 |

**Action Items:**
- [ ] **Week -3 (before project start):** Register Azure AD app, configure Graph API permissions
- [ ] **Week -3:** Submit for publisher verification (requires: verified domain, privacy policy, terms of service)
- [ ] **Week 1 (S1):** Test with dev M365 tenant (unverified app, limited to 10 users)
- [ ] **Week 3-6:** Receive verification approval
- [ ] **Week 5 (S3):** Production-ready for S3 sprint

**⚠️ Critical:** Microsoft's publisher verification requires a **verified domain** and can take **up to 8 weeks** if additional documentation is requested. Start this **3 weeks before Sprint 1**.

---

### 1.3 Vanta Account Provisioning

| Item | Details |
|------|---------|
| **What's needed** | Vanta Startup plan account + AWS connector + GitHub connector + compliance framework selection |
| **Lead time** | **2-3 weeks** (account setup + connector configuration + initial evidence collection) |
| **Blocks** | Sprint 7 (W13-14) — Vanta evidence collection must begin W13 for SOC 2 Type 1 at v1 |
| **Cost** | $4-6K/yr (Startup plan) |
| **Steps** | 1. Sign up for Vanta Startup plan (Day 1)<br>2. Connect AWS account (Week 1)<br>3. Connect GitHub organization (Week 1)<br>4. Select compliance frameworks (SOC 2 Type 1, ISO 27001) (Week 1)<br>5. Configure evidence collection rules (Week 2)<br>6. **Initial evidence scan** (Week 2-3) → **2-3 weeks**<br>7. Review compliance gaps (Week 3) |
| **Risk if delayed** | SOC 2 Type 1 evidence window insufficient at v1 (W26) → certification delayed to v1.5 |
| **Owner** | PM (Vanta account) + DevSecOps (connector setup) |
| **Documentation** | https://www.vanta.com/products/soc-2 |

**Action Items:**
- [ ] **Week 8 (before S7):** Sign up for Vanta, connect AWS + GitHub
- [ ] **Week 9-10:** Configure evidence collection, run initial scan
- [ ] **Week 11:** Review compliance gaps, create remediation plan
- [ ] **Week 13 (S7):** Vanta active, evidence collection running continuously

**⚠️ Critical:** Vanta evidence collection must start **no later than W13** to have sufficient evidence for SOC 2 Type 1 audit at v1 (W26). Starting earlier (W8-10) provides buffer for configuration issues.

---

## Priority 2: HIGH IMPACT — Start in Sprint 1-2

### 2.1 Slack App Creation + OAuth Configuration

| Item | Details |
|------|---------|
| **What's needed** | Slack App + OAuth scopes + Admin API access + Enterprise Grid verification (if applicable) |
| **Lead time** | **1-2 weeks** (app review for Admin API scopes) |
| **Blocks** | Sprint 5 (W9-10) — Slack integration is S5 deliverable |
| **Cost** | Free (Slack app), but requires Business+ tier for user management API ($12.50/user/mo) |
| **Steps** | 1. Create Slack app (Day 1)<br>2. Configure OAuth scopes (Day 1)<br>3. **Submit for Admin API access** (Day 2) → **1-2 weeks wait**<br>4. Test with dev Slack workspace (W2)<br>5. Implement tier detection (Business+ vs Pro/Free) (S5) |
| **Risk if delayed** | S5 cannot integrate Slack → v1 missing Slack support (30% of SME market) |
| **Owner** | Tech Lead (Slack app setup) |
| **Documentation** | https://api.slack.com/start/quickstart |

**Action Items:**
- [ ] **Week 1 (S1):** Create Slack app, configure OAuth scopes
- [ ] **Week 1:** Submit for Admin API access (if needed)
- [ ] **Week 2-3:** Receive approval (if applicable)
- [ ] **Week 9 (S5):** Production-ready for S5 sprint

**⚠️ Note:** Slack Admin API requires **Business+ tier** ($12.50/user/mo). Free/Pro tiers have read-only access. UI must detect tier and show warning (see R-H10, BS-10).

---

### 2.2 AWS IAM Cross-Account Access Setup

| Item | Details |
|------|---------|
| **What's needed** | AWS account + IAM role for cross-account access + CloudTrail enabled + Config enabled |
| **Lead time** | **1 week** (customer setup + testing) |
| **Blocks** | Sprint 5 (W9-10) — AWS IAM discovery is S5 deliverable |
| **Cost** | Free (IAM), but CloudTrail + Config have storage costs (~$10-50/mo per customer) |
| **Steps** | 1. Create IAM role template (Day 1)<br>2. Document customer setup instructions (Day 1)<br>3. Test with dev AWS account (W2)<br>4. Create CloudFormation template for one-click setup (W3) |
| **Risk if delayed** | S5 cannot integrate AWS → v1 missing AWS support (cloud asset inventory incomplete) |
| **Owner** | Tech Lead (IAM role design) + DevSecOps (CloudFormation template) |
| **Documentation** | https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user.html |

**Action Items:**
- [ ] **Week 1 (S1):** Design IAM role with minimum permissions (read-only)
- [ ] **Week 2:** Create CloudFormation template for customer self-service setup
- [ ] **Week 3:** Test with dev AWS account
- [ ] **Week 9 (S5):** Production-ready for S5 sprint

---

### 2.3 Hive Moderation API Access

| Item | Details |
|------|---------|
| **What's needed** | Hive Moderation API key + rate limits confirmed + pricing tier selected |
| **Lead time** | **1-2 weeks** (account approval + API key provisioning) |
| **Blocks** | Sprint 10 (W19-20) — Deepfake detection is S10 deliverable (Track 2) |
| **Cost** | Pay-per-use (~$0.01/check), estimated $200-500/mo at 50 customers |
| **Steps** | 1. Sign up for Hive Moderation (Day 1)<br>2. **Submit API access request** (Day 2) → **1-2 weeks wait**<br>3. Receive API key (Week 2)<br>4. Test voice/video analysis (W3)<br>5. Confirm rate limits + pricing (W3) |
| **Risk if delayed** | S10 cannot integrate deepfake detection → Track 2 accuracy gate 4 delayed |
| **Owner** | ML Engineer #1 (API integration) + PM (account setup) |
| **Documentation** | https://thehive.ai/apis |

**Action Items:**
- [ ] **Week 1 (S1):** Sign up for Hive Moderation, submit API access request
- [ ] **Week 2-3:** Receive API key, test voice/video analysis
- [ ] **Week 4:** Confirm rate limits + pricing, document in design
- [ ] **Week 19 (S10):** Production-ready for S10 sprint

---

### 2.4 Lakera Guard API Access

| Item | Details |
|------|---------|
| **What's needed** | Lakera Guard API key + pricing confirmation + rate limits confirmed |
| **Lead time** | **1-2 weeks** (account approval + API key provisioning) |
| **Blocks** | Sprint 8 (W15-16) — Prompt injection detection is S8 deliverable (Track 2) |
| **Cost** | ~$0.001/request, estimated $50-100/mo at 50 customers × 50K daily checks |
| **Steps** | 1. Sign up for Lakera Guard (Day 1)<br>2. **Submit API access request** (Day 2) → **1-2 weeks wait**<br>3. Receive API key (Week 2)<br>4. Test prompt injection detection (W3)<br>5. **Confirm pricing viability** (W3) → **Go/No-go decision by S1 end (W2)** |
| **Risk if delayed** | S8 cannot integrate prompt injection detection → Track 2 feature delayed |
| **Owner** | ML Engineer #1 (API integration) + PM (pricing negotiation) |
| **Documentation** | https://www.lakera.ai/lakera-guard |

**Action Items:**
- [ ] **Week 1 (S1):** Sign up for Lakera Guard, submit API access request
- [ ] **Week 2:** Receive API key, test prompt injection detection
- [ ] **Week 2 (S1 end):** **Go/No-go decision:** Confirm pricing <$0.05/request is viable at SME scale (see OQ-2)
- [ ] **Week 15 (S8):** Production-ready for S8 sprint

**⚠️ Critical:** Lakera Guard pricing decision is a **hard gate at S1 end (W2)** (see OQ-2). If pricing is not viable, fallback to WASM-only BERT model (lower accuracy, but no API cost).

---

## Priority 3: MEDIUM IMPACT — Start in Sprint 3-5

### 3.1 Chrome Web Store Developer Account

| Item | Details |
|------|---------|
| **What's needed** | Chrome Web Store developer account + $5 registration fee + privacy policy + terms of service |
| **Lead time** | **1 week** (account verification) |
| **Blocks** | Sprint 12 (W23-24) — Browser extension submission is S12 deliverable |
| **Cost** | $5 one-time registration fee |
| **Steps** | 1. Register Chrome Web Store developer account (Day 1)<br>2. Pay $5 registration fee (Day 1)<br>3. **Account verification** (Day 2-7) → **1 week wait**<br>4. Prepare privacy policy + terms of service (W2)<br>5. Test extension submission with dummy extension (W3) |
| **Risk if delayed** | S12 cannot submit extension → v1.5 (W38) missing browser extension |
| **Owner** | Frontend Engineer #2 (extension developer) + PM (legal docs) |
| **Documentation** | https://developer.chrome.com/docs/webstore/register/ |

**Action Items:**
- [ ] **Week 10 (before S12):** Register Chrome Web Store developer account
- [ ] **Week 11:** Account verified, privacy policy + terms of service ready
- [ ] **Week 12:** Test dummy extension submission
- [ ] **Week 23 (S12):** Production extension submission

**⚠️ Note:** Chrome Web Store review for security extensions (with `tabs`, `webRequest`, `scripting` permissions) can take **2-6 weeks** (see R-H10). Submit stripped-down v0 extension at W18 (S8) for early review buffer.

---

### 3.2 Apple Developer Program + Google Play Console

| Item | Details |
|------|---------|
| **What's needed** | Apple Developer Program membership ($99/yr) + Google Play Console account ($25 one-time) |
| **Lead time** | **1-2 weeks** (Apple: account verification, Google: immediate) |
| **Blocks** | Sprint 6 (W11-12) — Mobile app beta is S6 deliverable (MVP) |
| **Cost** | Apple: $99/yr, Google: $25 one-time |
| **Steps** | 1. Register Apple Developer Program (Day 1)<br>2. **Apple account verification** (Day 2-14) → **1-2 weeks wait**<br>3. Register Google Play Console (Day 1) → **immediate**<br>4. Create app listings (W2)<br>5. Prepare app store assets (screenshots, descriptions) (W3) |
| **Risk if delayed** | S6 cannot publish mobile app beta → MVP delayed |
| **Owner** | Flutter Engineer (mobile developer) + PM (app store listings) |
| **Documentation** | https://developer.apple.com/programs/ · https://play.google.com/console/ |

**Action Items:**
- [ ] **Week 1 (S1):** Register Apple Developer Program + Google Play Console
- [ ] **Week 2-3:** Apple account verified
- [ ] **Week 4:** Create app listings, prepare app store assets
- [ ] **Week 11 (S6):** Mobile app beta published to TestFlight + Play Console

---

### 3.3 Penetration Test Vendor Selection + LOI

| Item | Details |
|------|---------|
| **What's needed** | Pentest vendor selected + Letter of Intent (LOI) signed + scope agreed |
| **Lead time** | **6-8 weeks** (vendor selection + scheduling + pentest execution) |
| **Blocks** | Sprint 11 (W21-22) — Pentest must start W21 for v1 (W26) |
| **Cost** | $10-20K for multi-tenant SaaS pentest |
| **Steps** | 1. RFP to 3-5 pentest vendors (W8)<br>2. Vendor selection (W10)<br>3. **Sign LOI** (W14) → **hard deadline**<br>4. Scope agreement (W15-16)<br>5. Pentest kickoff (W21) |
| **Risk if delayed** | Pentest does not complete before v1 (W26) → v1 launch delayed or ships with unresolved Critical/High findings |
| **Owner** | PM (vendor selection) + Tech Lead (scope definition) |
| **Documentation** | Internal RFP template (to be created) |

**Action Items:**
- [ ] **Week 8:** Send RFP to 3-5 pentest vendors (recommendations: Bishop Fox, NCC Group, Trail of Bits, Cure53, Cobalt)
- [ ] **Week 10:** Vendor selection complete
- [ ] **Week 14 (hard deadline):** **Sign LOI** (see R-H9, external dependency)
- [ ] **Week 15-16:** Scope agreement (multi-tenant SaaS, Track 1 + Track 2, web + mobile + browser extension)
- [ ] **Week 21 (S11):** Pentest kickoff

**⚠️ Critical:** Pentest vendor LOI must be signed **no later than W14** (see external dependencies). Starting vendor selection at W8 provides 6-week buffer.

---

## Priority 4: LOW IMPACT — Can Start Later

### 4.1 Keycloak Self-Hosted Setup

| Item | Details |
|------|---------|
| **What's needed** | Keycloak ECS Fargate deployment + PostgreSQL database + OIDC/SAML configuration |
| **Lead time** | **1 week** (infrastructure setup + testing) |
| **Blocks** | Sprint 1 (W1-2) — Keycloak SSO is S1 deliverable |
| **Cost** | ~$50/mo (ECS compute) |
| **Steps** | 1. Deploy Keycloak to ECS Fargate (S1)<br>2. Configure PostgreSQL database (S1)<br>3. Configure OIDC/SAML (S1)<br>4. Test Google + M365 federation (S1) |
| **Risk if delayed** | S1 cannot deliver auth → all sprints delayed |
| **Owner** | Tech Lead (Keycloak deployment) + DevSecOps (infrastructure) |
| **Documentation** | https://www.keycloak.org/getting-started/getting-started-docker |

**Action Items:**
- [ ] **Week 1 (S1):** Deploy Keycloak to ECS Fargate, configure OIDC/SAML
- [ ] **Week 2 (S1):** Test Google + M365 federation, configure MFA TOTP

**⚠️ Note:** Keycloak is self-hosted, so no external approval process. However, see R-C6 for HA requirements (min 2 ECS tasks, JWKS caching).

---

### 4.2 Cloudflare R2 Account Setup

| Item | Details |
|------|---------|
| **What's needed** | Cloudflare account + R2 storage enabled + S3-compatible API credentials |
| **Lead time** | **Immediate** (no approval process) |
| **Blocks** | Sprint 1 (W1-2) — S3 Object Lock is S1 deliverable (audit log storage) |
| **Cost** | ~$50/mo (storage + egress) |
| **Steps** | 1. Sign up for Cloudflare (Day 1)<br>2. Enable R2 storage (Day 1)<br>3. Create S3-compatible API credentials (Day 1)<br>4. Test S3 Object Lock compatibility (W1) |
| **Risk if delayed** | S1 cannot deliver audit log storage → compliance evidence collection delayed |
| **Owner** | DevSecOps (Cloudflare setup) |
| **Documentation** | https://developers.cloudflare.com/r2/ |

**Action Items:**
- [ ] **Week 1 (S1):** Sign up for Cloudflare, enable R2 storage, create API credentials
- [ ] **Week 1 (S1):** Test S3 Object Lock compatibility (WORM mode)

---

## Summary: Critical Path Timeline

```
Week -3 (before project start):
  ✅ Microsoft 365 App Registration + Publisher verification submission
  ✅ Google Workspace GCP project + OAuth consent screen verification submission

Week -2:
  ✅ Google Workspace verification follow-up (if needed)

Week -1:
  ✅ Confirm Google + Microsoft verifications in progress

Week 1 (Sprint 1):
  ✅ Slack app creation + Admin API access request
  ✅ Hive Moderation API access request
  ✅ Lakera Guard API access request
  ✅ Apple Developer Program registration
  ✅ Google Play Console registration
  ✅ Keycloak deployment
  ✅ Cloudflare R2 setup

Week 2 (Sprint 1 end):
  ✅ Lakera Guard pricing decision (Go/No-go gate)
  ✅ Google Workspace verification approved (target)
  ✅ Hive Moderation API key received (target)
  ✅ Lakera Guard API key received (target)

Week 3-4 (Sprint 2):
  ✅ Google Workspace production-ready for S2 sprint
  ✅ Apple Developer account verified

Week 5-6 (Sprint 3):
  ✅ Microsoft 365 verification approved (target)
  ✅ Microsoft 365 production-ready for S3 sprint

Week 8 (before Sprint 7):
  ✅ Pentest vendor RFP sent
  ✅ Vanta account setup begins

Week 10:
  ✅ Pentest vendor selected
  ✅ Chrome Web Store developer account registration

Week 11:
  ✅ Vanta evidence collection active

Week 13 (Sprint 7):
  ✅ Vanta evidence collection running continuously (SOC 2 Type 1 window begins)

Week 14 (hard deadline):
  ✅ Pentest vendor LOI signed

Week 21 (Sprint 11):
  ✅ Pentest kickoff
```

---

## Risk Mitigation: Backup Plans

| Risk | Backup Plan |
|------|-------------|
| **Google Workspace verification delayed >6 weeks** | Use unverified OAuth consent screen for pilot customers (limited to 100 users). Defer production launch to W16. |
| **Microsoft 365 publisher verification delayed >8 weeks** | Use unverified app for pilot customers (limited to 10 users). Defer production launch to W18. |
| **Lakera Guard pricing not viable** | Fallback to WASM-only BERT model (lower accuracy, but no API cost). Accept lower TPR/FPR in v1. |
| **Pentest vendor LOI not signed by W14** | Use backup vendor (pre-qualified list). Accept 2-week delay in pentest start (W23 instead of W21). |
| **Vanta setup delayed >3 weeks** | Manual evidence collection for SOC 2 Type 1. Accept higher PM workload. |
| **Hive Moderation API access denied** | Fallback to Resemble Detect (voice only) or defer deepfake detection to v1.5. |

---

## Checklist: Pre-Sprint 1 Preparation

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

## Conclusion

**Critical Actions Before Sprint 1:**
1. **Week -3:** Google Workspace + Microsoft 365 verification submissions (longest lead time: 3-6 weeks)
2. **Week 1:** Slack, Hive, Lakera API access requests (1-2 weeks lead time)
3. **Week 2:** Lakera Guard pricing decision (Go/No-go gate)

**Critical Actions During Phase 1:**
1. **Week 8:** Pentest vendor RFP + Vanta account setup
2. **Week 14:** Pentest vendor LOI signing (hard deadline)

**Total Lead Time Savings:** Starting these preparations early saves **4-8 weeks** of potential sprint delays.
