# SMESec — 3rd-Party Integration Principles & Lead Time Management

**Date:** 2026-05-29  
**Purpose:** Master reference for all 3rd-party integration dependencies, lead times, and preparation requirements  
**Status:** Canonical — All other documents must align with this  
**Owner:** PM + Tech Lead

---

## Core Principle

> **3rd-party integrations with >1 week lead time MUST be started BEFORE the sprint that depends on them.**

This principle prevents sprint delays caused by waiting for external approvals, API access, or account verifications.

---

## Lead Time Categories

### Category A: CRITICAL PATH (3-8 weeks lead time)
**Rule:** Must start 3-8 weeks before dependent sprint

| Integration | Lead Time | Reason | Start Timing |
|-------------|-----------|--------|--------------|
| **Google Workspace OAuth** | 2-4 weeks | OAuth consent screen verification by Google | Week -3 (before project start) |
| **Microsoft 365 Publisher** | 3-6 weeks | Publisher verification by Microsoft | Week -3 (before project start) |
| **Vanta Account** | 2-3 weeks | Connector provisioning + initial evidence scan | Week 8 (5 weeks before S7) |
| **Pentest Vendor LOI** | 6-8 weeks | Vendor selection + scheduling + scope agreement | Week 8 (RFP) → Week 14 (LOI signed) |

### Category B: HIGH IMPACT (1-2 weeks lead time)
**Rule:** Must start 1-2 weeks before dependent sprint

| Integration | Lead Time | Reason | Start Timing |
|-------------|-----------|--------|--------------|
| **Slack Admin API** | 1-2 weeks | Admin API scope approval | Week 1 (S1) |
| **Hive Moderation API** | 1-2 weeks | Account approval + API key provisioning | Week 1 (S1) |
| **Lakera Guard API** | 1-2 weeks | Account approval + pricing confirmation | Week 1 (S1) |
| **Chrome Web Store** | 1 week | Developer account verification | Week 10 (13 weeks before S12) |
| **Apple Developer** | 1-2 weeks | Account verification | Week 1 (S1) |

### Category C: MEDIUM IMPACT (<1 week lead time)
**Rule:** Can start in same sprint as dependency

| Integration | Lead Time | Reason | Start Timing |
|-------------|-----------|--------|--------------|
| **AWS IAM Cross-Account** | <1 week | Customer self-service setup | Week 1 (S1) |
| **Google Play Console** | Immediate | No approval process | Week 1 (S1) |
| **Cloudflare R2** | Immediate | No approval process | Week 1 (S1) |
| **Keycloak (self-hosted)** | <1 week | Infrastructure setup only | Week 1 (S1) |

---

## Dependency Mapping: Sprint → 3rd-Party

| Sprint | Deliverable | Required 3rd-Party | Lead Time | Must Start By |
|--------|-------------|-------------------|-----------|---------------|
| **S1 (W1-2)** | Infrastructure + Auth | Keycloak (self-hosted), Cloudflare R2 | <1 week | Week 1 |
| **S2 (W3-4)** | Google Workspace sync | Google Workspace OAuth (verified) | 2-4 weeks | **Week -3** |
| **S3 (W5-6)** | M365 sync | Microsoft 365 Publisher (verified) | 3-6 weeks | **Week -3** |
| **S5 (W9-10)** | Slack + AWS discovery | Slack Admin API, AWS IAM | 1-2 weeks | Week 1 |
| **S6 (W11-12)** | Mobile app beta | Apple Developer, Google Play | 1-2 weeks | Week 1 |
| **S7 (W13-14)** | Vanta evidence collection | Vanta account (active) | 2-3 weeks | **Week 8** |
| **S8 (W15-16)** | Prompt injection | Lakera Guard API | 1-2 weeks | Week 1 |
| **S10 (W19-20)** | Deepfake detection | Hive Moderation API | 1-2 weeks | Week 1 |
| **S11 (W21-22)** | Pentest remediation | Pentest vendor (LOI signed) | 6-8 weeks | **Week 8 (RFP) → Week 14 (LOI)** |
| **S12 (W23-24)** | Chrome extension submission | Chrome Web Store account | 1 week | Week 10 |

---

## Hard Gates & No-Go Conditions

### Gate 1: Google Workspace Verification (Week -3 → Week 2-4)
- **If delayed >6 weeks:** Use unverified OAuth (limited to 100 users) for pilot. Defer production to W16.
- **Blocks:** S2 (W3-4) Google Workspace sync
- **Fallback:** Pilot-only mode until verification completes

### Gate 2: Microsoft 365 Publisher Verification (Week -3 → Week 3-6)
- **If delayed >8 weeks:** Use unverified app (limited to 10 users) for pilot. Defer production to W18.
- **Blocks:** S3 (W5-6) M365 sync
- **Fallback:** Pilot-only mode until verification completes

### Gate 3: Lakera Guard Pricing Decision (Week 1 → Week 2)
- **If pricing not viable (<$0.05/request):** Fallback to WASM-only BERT model (lower accuracy, no API cost)
- **Blocks:** S8 (W15-16) Prompt injection detection
- **Fallback:** WASM BERT-tiny only (TPR ~75%, FPR ~10% vs Lakera TPR >85%, FPR <2%)

### Gate 4: Pentest Vendor LOI (Week 8 → Week 14)
- **If LOI not signed by W14:** Use backup vendor (2-week delay). Pentest starts W23 instead of W21.
- **Blocks:** S11 (W21-22) Pentest remediation
- **Fallback:** Backup vendor list pre-qualified

### Gate 5: Vanta Setup (Week 8 → Week 11)
- **If delayed >3 weeks:** Manual evidence collection for SOC 2 Type 1. Higher PM workload.
- **Blocks:** S7 (W13-14) Evidence collection start
- **Fallback:** Manual screenshot collection + spreadsheet tracking

---

## Verification Requirements by Integration

### Google Workspace OAuth Consent Screen
**Required for verification:**
- Verified domain (DNS TXT record)
- Privacy policy URL (public, accessible)
- Terms of service URL (public, accessible)
- App logo (512x512 PNG)
- Support email address
- OAuth scopes justification (written explanation)

**Verification timeline:**
- Submit: Day 1
- Google review: 2-4 weeks (can request additional info → +2 weeks)
- Approval: Week 2-6

### Microsoft 365 Publisher Verification
**Required for verification:**
- Verified domain (DNS TXT record)
- Privacy policy URL (public, accessible)
- Terms of service URL (public, accessible)
- Microsoft Partner Network enrollment (optional but recommended)
- App registration in Azure AD
- Publisher domain ownership proof

**Verification timeline:**
- Submit: Day 1
- Microsoft review: 3-6 weeks (can request additional documentation → +2 weeks)
- Approval: Week 3-8

### Slack Admin API Access
**Required for approval:**
- Slack app created
- OAuth scopes defined (admin.users:read, admin.apps:read, admin.users:write)
- App description + use case
- Privacy policy URL

**Approval timeline:**
- Submit: Day 1
- Slack review: 1-2 weeks
- Approval: Week 1-2

**⚠️ Tier constraint:** User management API requires **Business+ tier** ($12.50/user/mo). Free/Pro tiers: read-only access only.

### Hive Moderation API Access
**Required for approval:**
- Account registration
- Use case description (deepfake detection for fraud prevention)
- Estimated volume (checks/month)

**Approval timeline:**
- Submit: Day 1
- Hive review: 1-2 weeks
- API key provisioned: Week 1-2

### Lakera Guard API Access
**Required for approval:**
- Account registration
- Use case description (prompt injection detection)
- Estimated volume (requests/month)
- Pricing tier selection

**Approval timeline:**
- Submit: Day 1
- Lakera review: 1-2 weeks
- API key provisioned: Week 1-2

**⚠️ Pricing gate:** Must confirm <$0.05/request by Week 2 (S1 end) for Go decision.

---

## Cost Summary by Integration

| Integration | Cost Model | Estimated Cost (1K tenants) | Notes |
|-------------|------------|-------------------------------|-------|
| Google Workspace OAuth | Free | $0/mo | GCP project free tier |
| Microsoft 365 Publisher | Free | $0/mo | Azure AD free tier |
| Slack Admin API | Free (app) | $0/mo | But requires Business+ tier for user mgmt ($12.50/user/mo per customer) |
| AWS IAM Cross-Account | Free | $0/mo | CloudTrail + Config storage: ~$10-50/mo per customer |
| Vanta | Subscription | $400/mo ($4,800/yr) | Startup plan |
| Hive Moderation API | Pay-per-use | $200-500/mo | ~$0.01/check, estimated 20K-50K checks/mo |
| Lakera Guard API | Pay-per-use | $50-100/mo | ~$0.001/request, estimated 50K-100K requests/mo |
| Chrome Web Store | One-time | $5 | Developer registration fee |
| Apple Developer | Annual | $99/yr | Developer program membership |
| Google Play Console | One-time | $25 | Developer registration fee |
| Keycloak (self-hosted) | Infrastructure | $50/mo | ECS Fargate compute only |
| Cloudflare R2 | Usage-based | $50/mo | Storage + egress |

**Total 3rd-party costs (Year 1 at 1K tenant capacity):** ~$7,500/mo = $90,000/yr. Infrastructure is pre-provisioned for 1K tenants from Sprint 1; costs scale sub-linearly relative to revenue ($800K/mo MRR at 1K tenants = ~0.9% infra-to-revenue ratio.

---

## Risk Mitigation: Backup Plans

| Risk | Probability | Impact | Backup Plan | Cost of Backup |
|------|-------------|--------|-------------|----------------|
| Google verification delayed >6 weeks | Medium | High | Unverified OAuth (100 user limit) for pilot | Delayed production launch (W12 → W16) |
| Microsoft verification delayed >8 weeks | Medium | High | Unverified app (10 user limit) for pilot | Delayed production launch (W12 → W18) |
| Lakera Guard pricing not viable | Low | Medium | WASM-only BERT model | Lower accuracy (TPR ~75% vs >85%) |
| Pentest vendor LOI not signed by W14 | Low | High | Backup vendor (pre-qualified) | 2-week delay (W21 → W23) |
| Vanta setup delayed >3 weeks | Low | Medium | Manual evidence collection | Higher PM workload (~20h/week) |
| Hive API access denied | Low | Medium | Resemble Detect (voice only) or defer to v1.5 | Reduced feature scope |
| Slack Admin API denied | Low | Low | Read-only Slack integration | No user deactivation capability |

---

## Checklist: Pre-Sprint 1 Preparation

**Week -3 (3 weeks before Sprint 1):**
- [ ] Google Workspace: Create GCP project, enable Admin SDK API, create service account
- [ ] Google Workspace: Submit OAuth consent screen for verification (with privacy policy, ToS, logo, support email)
- [ ] Microsoft 365: Register Azure AD app, configure Graph API permissions, create client secret
- [ ] Microsoft 365: Submit for publisher verification (with verified domain, privacy policy, ToS)

**Week -2:**
- [ ] Google Workspace: Follow up on verification status (check for additional info requests)
- [ ] Microsoft 365: Follow up on verification status (check for additional documentation requests)

**Week -1:**
- [ ] Confirm Google + Microsoft verifications are in progress (no blockers)
- [ ] Prepare Keycloak deployment plan (ECS Fargate, PostgreSQL, OIDC/SAML config)

**Week 1 (Sprint 1 Day 1):**
- [ ] Slack: Create Slack app, configure OAuth scopes, submit for Admin API access
- [ ] Hive Moderation: Sign up, submit API access request with use case
- [ ] Lakera Guard: Sign up, submit API access request with use case
- [ ] Apple Developer: Register Apple Developer Program ($99/yr)
- [ ] Google Play Console: Register Google Play Console ($25 one-time)
- [ ] AWS IAM: Design IAM role template with minimum permissions
- [ ] Keycloak: Deploy to ECS Fargate, configure OIDC/SAML
- [ ] Cloudflare R2: Sign up, enable R2 storage, create S3-compatible API credentials

**Week 2 (Sprint 1 end):**
- [ ] **Lakera Guard pricing decision (Go/No-go gate):** Confirm <$0.05/request viable
- [ ] Google Workspace: Verification approved (target) → configure domain-wide delegation
- [ ] Hive Moderation: API key received (target) → test voice/video analysis
- [ ] Lakera Guard: API key received (target) → test prompt injection detection

**Week 3-4 (Sprint 2):**
- [ ] Google Workspace: Production-ready for S2 sprint
- [ ] Apple Developer: Account verified

**Week 5-6 (Sprint 3):**
- [ ] Microsoft 365: Verification approved (target) → production-ready for S3 sprint

**Week 8 (before Sprint 7):**
- [ ] Pentest vendor: Send RFP to 3-5 vendors
- [ ] Vanta: Sign up for Vanta Startup plan, connect AWS + GitHub

**Week 10:**
- [ ] Pentest vendor: Selection complete
- [ ] Chrome Web Store: Register developer account ($5)
- [ ] Vanta: Configure evidence collection rules

**Week 11:**
- [ ] Vanta: Initial evidence scan complete, review compliance gaps

**Week 13 (Sprint 7):**
- [ ] Vanta: Evidence collection running continuously (SOC 2 Type 1 window begins)

**Week 14 (hard deadline):**
- [ ] **Pentest vendor: LOI signed** (no exceptions — this is a hard gate)

**Week 21 (Sprint 11):**
- [ ] Pentest: Kickoff (vendor onboarded, scope agreed, testing begins)

---

## Enforcement Rules

### Rule 1: No Sprint Starts Without Dependencies Ready
**Enforcement:** PM reviews 3rd-party readiness at sprint planning (Day 1 of each sprint). If dependency not ready → defer sprint deliverable or activate backup plan.

### Rule 2: Hard Gates Are Non-Negotiable
**Enforcement:** The following gates CANNOT be bypassed:
- Google Workspace verification (S2 blocker)
- Microsoft 365 verification (S3 blocker)
- Lakera Guard pricing decision (S1 end blocker)
- Pentest vendor LOI (W14 hard deadline)

If a hard gate is missed → milestone date slips automatically (no negotiation).

### Rule 3: Backup Plans Must Be Pre-Qualified
**Enforcement:** All backup plans (backup pentest vendor, WASM-only BERT, manual Vanta evidence) must be validated BEFORE the primary plan fails. No "figure it out later" backups.

### Rule 4: Cost Overruns Require Re-Approval
**Enforcement:** If any 3rd-party cost exceeds estimate by >50% → PM escalates for re-approval before proceeding. Example: Lakera Guard pricing >$0.075/request → escalate before committing.

---

## Document Synchronization Requirements

**This document is canonical.** All other documents must align with the lead times, gates, and timelines defined here.

**Documents that must be synchronized:**
1. [04-delivery-plan-original.md](04-delivery-plan-original.md) — Sprint dependencies must match 3rd-party readiness
2. [10-feasibility-assessment-and-remediation-plan.md](10-feasibility-assessment-and-remediation-plan.md) — Risk register must include 3rd-party gates
3. [02-design-document.md](02-design-document.md) — Build vs Buy decisions must reference lead times
4. [01-system-architecture.md](01-system-architecture.md) — Integration touchpoints must note verification requirements
5. [03-two-track-approach.md](03-two-track-approach.md) — Track 2 timeline must account for API access lead times
6. [09-ai-governance-module.md](09-ai-governance-module.md) — Module delivery must account for Hive/Lakera lead times

**Synchronization cadence:** Weekly review during Phase 1 (S1-S6). Any change to lead times or gates → immediate update to all dependent documents.

---

## Appendix: Verification Status Tracking Template

```markdown
## 3rd-Party Verification Status (Week X)

### Google Workspace OAuth
- Status: [Pending / Approved / Blocked]
- Submitted: [Date]
- Expected approval: [Date]
- Blockers: [None / Additional info requested / ...]

### Microsoft 365 Publisher
- Status: [Pending / Approved / Blocked]
- Submitted: [Date]
- Expected approval: [Date]
- Blockers: [None / Additional documentation requested / ...]

### Slack Admin API
- Status: [Pending / Approved / Blocked]
- Submitted: [Date]
- Expected approval: [Date]
- Blockers: [None / ...]

### Hive Moderation API
- Status: [Pending / Approved / Blocked]
- Submitted: [Date]
- API key received: [Date]
- Blockers: [None / ...]

### Lakera Guard API
- Status: [Pending / Approved / Blocked]
- Submitted: [Date]
- API key received: [Date]
- Pricing confirmed: [Yes / No / Pending]
- Blockers: [None / ...]

### Pentest Vendor
- Status: [RFP sent / Vendor selected / LOI signed]
- RFP sent: [Date]
- Vendor selected: [Date]
- LOI signed: [Date]
- Blockers: [None / ...]

### Vanta
- Status: [Not started / Setup in progress / Active]
- Account created: [Date]
- Connectors configured: [Date]
- Evidence collection active: [Date]
- Blockers: [None / ...]
```

**Update frequency:** Weekly during Phase 1 (S1-S6), bi-weekly during Phase 2-4.
