# Feature Debate Analysis: Incident Playbooks for Non-Security Staff

**Date:** 2026-05-28  
**Participants:** PM/Risk Manager (30yr) vs Solution Architect/Cybersecurity (30yr)  
**Objective:** Identify 10 essential feature approaches to achieve key requirement #5

---

## Debate Format

**Round 1:** Independent analysis from each perspective  
**Round 2:** Cross-examination and synthesis  
**Round 3:** Consensus on 10 feature approaches

---

## Round 1: Independent Analysis

### PM/Risk Manager Perspective

**Context:** 30 years managing enterprise software projects, cybersecurity platforms, SaaS for SMEs. Focus on adoption risk, operational feasibility, and business outcomes.

#### Critical Success Factors

1. **User Adoption Risk = Project Failure Risk**
   - If non-security staff abandon playbooks mid-execution, the entire platform value proposition collapses
   - SMEs won't hire security experts to use a "security platform for non-experts"
   - Probability: 40% without proper UX design
   - Impact: Existential (product fails)

2. **Training Time = Barrier to Entry**
   - SME IT admins have 5-10 competing priorities
   - If playbooks require >2 hours training, adoption drops 60%
   - Must be intuitive enough to use in emergency without prior training

3. **False Positive Rate = User Trust**
   - Every false alarm erodes trust
   - At >15% FP rate, users start ignoring all alerts (alert fatigue)
   - Target: <10% FP rate to maintain credibility

4. **Compliance Evidence = Legal Defensibility**
   - SMEs face audits (ISO 27001, GDPR, SOC 2)
   - Missing evidence = failed audit = lost customers
   - Must be 100% automatic (cannot rely on manual documentation)

5. **Failure Recovery = Operational Continuity**
   - Playbooks will fail mid-execution (API timeouts, rate limits, network issues)
   - If failure = data loss or incomplete revocation, SME is exposed
   - Must handle partial failures gracefully

#### Feature Requirements (PM View)

| Feature Category | Why Critical | Consequence if Missing | User Persona |
|-----------------|--------------|----------------------|--------------|
| **1. Wizard UI with Progress Tracking** | Reduces cognitive load; user knows "where am I?" | 35% abandonment rate; users get lost mid-playbook | IT admin, HR manager |
| **2. Decision Gates (Yes/No)** | Eliminates need for security expertise | Users freeze at open-ended questions; call external consultant | IT admin, employee |
| **3. Pause/Resume Capability** | Real-world interruptions (phone call, meeting) | User must restart from beginning; frustration; abandonment | IT admin, manager |
| **4. Undo Support** | Reduces fear of mistakes | Users too scared to take action; escalate everything | IT admin (junior) |
| **5. Multi-Channel Notifications** | Ensures incident response even outside office hours | P0 incidents missed; delayed response; breach escalates | IT admin, manager, employee |
| **6. Automatic Evidence Collection** | Compliance without manual work | Failed audits; lost customers; legal liability | IT admin, compliance officer |
| **7. Plain Language (No Jargon)** | Accessible to non-technical users | Users don't understand instructions; errors; abandonment | HR manager, employee |

**PM Risk Assessment:**

| Risk | Probability | Impact | Mitigation Feature |
|------|------------|--------|-------------------|
| Playbook abandonment | 40% | Critical | Wizard UI + Progress tracking |
| False positive fatigue | 30% | High | Context enrichment + Threshold tuning |
| Compliance evidence gaps | 20% | High | Automatic evidence collection |
| Training bottleneck | 25% | Medium | Zero-training design (intuitive UI) |
| Failure mid-execution | 35% | High | Pause/resume + Retry logic |

---

### Solution Architect / Cybersecurity Perspective

**Context:** 30 years designing cybersecurity platforms, SaaS architecture, cloud-native systems. Focus on security, reliability, scalability, and technical feasibility.

#### Architectural Principles

1. **Playbooks Are High-Privilege Operations**
   - Revoking access, suspending accounts = destructive actions
   - If compromised, attacker can DoS entire organization
   - Must enforce RBAC, audit every action, prevent abuse

2. **Multi-Tenant Isolation = Non-Negotiable**
   - One tenant's playbook must NEVER affect another tenant
   - Cross-tenant data leakage = existential security failure
   - Must enforce at DB row level + API middleware + Step Functions

3. **Fault Tolerance = Reliability**
   - Playbooks execute in distributed systems (AWS, Google, M365, Slack)
   - Any component can fail (network, API rate limit, timeout)
   - Must handle partial failures without data loss or inconsistent state

4. **Immutable Audit Trail = Forensic Integrity**
   - Post-incident investigations require complete timeline
   - Mutable logs = tampering risk = inadmissible in legal proceedings
   - Must use append-only storage (S3 Object Lock)

5. **Event-Driven Integration = Loose Coupling**
   - Track 1 and Track 2 evolve independently
   - Tight coupling = deployment dependencies = slower velocity
   - Must use EventBridge for async, decoupled communication

#### Technical Feature Requirements (SA View)

| Feature Category | Technical Approach | Security/Reliability Property | Architecture Layer |
|-----------------|-------------------|------------------------------|-------------------|
| **1. Stateful Execution Engine** | AWS Step Functions (state machines) | Fault-tolerant, resumable, audit trail | Orchestration |
| **2. RBAC Policy Engine** | OPA/Rego with role-based policies | Least-privilege, prevent unauthorized playbook execution | Authorization |
| **3. Immutable Evidence Storage** | S3 Object Lock (WORM mode), 7-year retention | Tamper-proof, compliance-ready, forensic integrity | Storage |
| **4. Parallel Revocation** | Step Functions parallel states | <5 min containment, fault isolation | Orchestration |
| **5. Context Enrichment API** | Query Track 1 for user role, data sensitivity | Reduce false positives, risk-aware decisions | Integration |
| **6. Event-Driven Triggering** | EventBridge + Lambda router | Loose coupling, auto-trigger from AI detection | Integration |
| **7. Multi-Tenant Isolation** | Tenant-scoped encryption (KMS), row-level security | Zero cross-tenant leakage | Security |

**SA Security Threat Model:**

| Threat | Attack Vector | Mitigation Feature |
|--------|--------------|-------------------|
| Unauthorized playbook execution | Compromised IT admin account | RBAC + MFA + Audit log |
| Cross-tenant data leakage | Bug in tenant isolation | Row-level security + CI tests |
| Evidence tampering | Malicious insider | S3 Object Lock (immutable) |
| Playbook abuse (DoS) | Attacker triggers mass offboarding | Rate limiting + Manual approval gates |
| API credential theft | Stolen OAuth tokens | Short-lived tokens + Auto-rotation |

---

## Round 2: Cross-Examination

### PM Challenges SA

**PM:** "Your Step Functions approach is technically sound, but what about the user experience? A state machine JSON is not something an IT admin can debug when a playbook hangs. How do they know what went wrong?"

**SA:** "Valid concern. Step Functions has a visual workflow editor in AWS Console — IT admin can see which step failed and why. We'll also expose this in our UI: a timeline view showing each action (pending/in-progress/completed/failed) with plain-language error messages. The state machine is internal; users never see JSON."

**PM:** "What if AWS Step Functions is down? You've created a single point of failure."

**SA:** "Step Functions has 99.9% SLA. For the 0.1% downtime, we have two mitigations: (1) Manual playbook mode — UI shows checklist, IT admin executes manually, we log their actions. (2) Critical actions (suspend account) can be triggered via direct API calls bypassing Step Functions. This is a degraded mode, not ideal, but prevents total outage."

**PM:** "Acceptable. But I'm concerned about the learning curve for your 'context enrichment API.' If it adds 200ms latency and requires IT admin to understand 'risk multipliers,' we've lost the simplicity advantage."

**SA:** "Context enrichment is invisible to the user. It happens server-side before the playbook even starts. User sees: 'Risk score: 92/100 — High risk, recommend immediate action.' They don't see the math. Latency: 50-100ms, not 200ms. We cache user context for 5 minutes to avoid repeated queries."

---

### SA Challenges PM

**SA:** "Your 'undo support' feature sounds great for UX, but it's technically dangerous. If we allow undoing an account suspension, and the account was actually compromised, we've just re-enabled the attacker. How do you reconcile this?"

**PM:** "Fair point. Undo should only be available for reversible, low-risk actions. For example: undo adding a tool to the block-list (can re-allow it). But NOT undo account suspension (that requires a separate 'restore account' playbook with manager approval). We need to define which actions are undo-able in the spec."

**SA:** "Agreed. Let's make undo explicit per action type, not a blanket feature. Also, every undo must be logged as a separate audit event, not just 'reverting' the original action."

**PM:** "Accepted. Now, your 'parallel revocation' via Step Functions — what happens if Google Workspace revocation succeeds but M365 fails? User sees 'partially completed' status. What do they do next?"

**SA:** "Step Functions continues executing other branches even if one fails. At the end, we generate a report: 'Google: Success. M365: Failed (rate limit). Slack: Success. AWS: Success.' IT admin sees: 'Action required: Manually disable M365 account' with a link to M365 admin console and step-by-step instructions. We also send a P0 alert to IT admin immediately when any critical revocation fails."

**PM:** "Good. But this means your '5-minute containment' SLA is conditional. If M365 fails, containment is incomplete. We need to be honest with customers: '5 minutes for automated actions; manual fallback may take 10-30 minutes.'"

**SA:** "Agreed. We'll document this in the SLA: 'Best effort 5 minutes; 95th percentile 8 minutes; manual fallback required in <5% of cases.'"

---

## Round 3: Consensus on 10 Feature Approaches

### Synthesis Methodology

1. **User-Centric + Technically Feasible**: Features must satisfy both PM (adoption, usability) and SA (security, reliability) requirements
2. **Risk-Driven Prioritization**: Features that mitigate highest-probability, highest-impact risks come first
3. **Layered Defense**: Combine UX features (reduce user error) with technical features (handle system failures)

---

## 10 Essential Feature Approaches

---

### Feature 1: Wizard-Driven UI with Step-by-Step Guidance

#### Approach (PM)
- **Problem:** Non-security staff don't know what to do in an incident
- **Solution:** Wizard UI that guides user through each step with clear instructions
- **UX Pattern:** Progress bar (Step 3 of 7), current step highlighted, previous steps grayed out, next step disabled until current completes

#### Approach (SA)
- **Technical Implementation:** React component library with reusable wizard framework
- **State Management:** Step Functions state machine drives UI state (backend is source of truth)
- **Accessibility:** WCAG 2.1 AA compliant, keyboard navigation, screen reader support

#### Conclusion: Required Features
1. **Progress indicator**: Visual bar showing "Step X of Y" with percentage
2. **Step validation**: Cannot proceed to next step until current step completes
3. **Contextual help**: "?" icon on each step with plain-language explanation
4. **Mobile responsive**: Works on phone (manager approving from home)
5. **Real-time status updates**: WebSocket connection shows live progress

**Success Metric:** >95% playbook completion rate (not abandoned mid-execution)

---

### Feature 2: Decision Gates with Binary Choices (Yes/No)

#### Approach (PM)
- **Problem:** Open-ended questions require security expertise
- **Solution:** Convert all decisions to Yes/No or multiple-choice (max 3 options)
- **Example:** Instead of "Assess the risk level" → "Is this user an admin?" [Yes/No]

#### Approach (SA)
- **Technical Implementation:** Decision tree logic in Step Functions (Choice state)
- **Branching:** Each decision creates a different execution path
- **Audit Trail:** Log which branch was taken and why (user's answer + timestamp)

#### Conclusion: Required Features
1. **Binary decision gates**: Yes/No radio buttons (no free text)
2. **Contextual guidance**: Each option explains consequence ("If Yes, account will be suspended immediately")
3. **Default recommendations**: System suggests answer based on context ("Recommended: Yes, based on risk score 92")
4. **Override capability**: User can choose opposite of recommendation (logged as manual override)
5. **Decision history**: Show past decisions for similar incidents (learning from history)

**Success Metric:** <5% of users request external help during playbook execution

---

### Feature 3: Pause/Resume with State Persistence

#### Approach (PM)
- **Problem:** Real-world interruptions (phone call, meeting, emergency)
- **Solution:** User can pause playbook, close browser, resume later from exact same step
- **Use Case:** IT admin starts offboarding, gets called into meeting, resumes 30 minutes later

#### Approach (SA)
- **Technical Implementation:** Step Functions state machines are inherently resumable
- **State Storage:** Current step, user inputs, partial results stored in RDS (encrypted)
- **Session Management:** JWT token with 24-hour expiry; user can resume from any device

#### Conclusion: Required Features
1. **Pause button**: Visible on every step, saves state immediately
2. **Resume from dashboard**: "In Progress" section shows paused playbooks
3. **Timeout handling**: If paused >24 hours, playbook expires (security risk)
4. **Notification reminder**: Email after 2 hours if playbook still paused ("You have an incomplete incident response")
5. **Handoff capability**: IT admin can assign paused playbook to colleague

**Success Metric:** <2% playbook abandonment due to interruptions

---

### Feature 4: Selective Undo for Reversible Actions

#### Approach (PM)
- **Problem:** Fear of mistakes prevents action; users escalate everything
- **Solution:** Allow undo for low-risk, reversible actions (not account suspension)
- **Example:** Undo adding tool to block-list; cannot undo account suspension

#### Approach (SA)
- **Technical Implementation:** Action metadata includes `reversible: true/false`
- **Undo Logic:** Separate Step Functions workflow for reversal (not just "revert")
- **Audit Trail:** Undo logged as distinct event (who, when, why, what was undone)

#### Conclusion: Required Features
1. **Undo button**: Only visible for reversible actions (grayed out for irreversible)
2. **Undo confirmation**: "Are you sure? This will re-allow [tool name]"
3. **Undo time limit**: Can only undo within 1 hour (prevents abuse)
4. **Undo audit log**: Separate entry in evidence collection (not merged with original action)
5. **Undo restrictions**: Cannot undo if downstream actions already executed

**Reversible Actions:** Add to allow-list, add to block-list, send notification  
**Irreversible Actions:** Suspend account, revoke OAuth token, delete data

**Success Metric:** 30% reduction in escalations to external security consultants

---

### Feature 5: Multi-Channel, Priority-Based Notifications

#### Approach (PM)
- **Problem:** P0 incidents missed because user didn't check email
- **Solution:** Route notifications by priority: P0 = email + Slack + mobile push
- **Use Case:** Deepfake CEO voice detected at 11pm → IT admin gets push notification on phone

#### Approach (SA)
- **Technical Implementation:** AWS SNS for email (SES), Slack API, FCM/APNs for mobile
- **Priority Routing:** Lambda function routes based on incident severity
- **Delivery Tracking:** Log delivery status (sent, delivered, opened, clicked)

#### Conclusion: Required Features
1. **Priority levels**: P0 (Critical), P1 (High), P2 (Medium), P3 (Low)
2. **Channel routing**:
   - P0: Email + Slack + Mobile Push (immediate, <1 min)
   - P1: Email + Slack (within 5 min)
   - P2/P3: Email only (within 30 min)
3. **Escalation**: If no response to P0 within 10 minutes, escalate to manager
4. **Notification preferences**: User can configure (e.g., "No Slack after 10pm, only mobile")
5. **Delivery confirmation**: Track if notification was received (read receipt)

**Success Metric:** <2 minutes median response time for P0 incidents

---

### Feature 6: Automatic Evidence Collection with Immutable Storage

#### Approach (PM)
- **Problem:** Manual documentation = incomplete evidence = failed audits
- **Solution:** Every action automatically generates evidence, stored immutably
- **Use Case:** ISO 27001 auditor asks "Prove you revoked access within 24 hours" → One-click PDF report

#### Approach (SA)
- **Technical Implementation:** S3 Object Lock (WORM mode), 7-year retention, KMS encryption
- **Evidence Types:** Action logs, API responses, screenshots, user justifications
- **Compliance Mapping:** Each evidence item tagged with control (ISO 27001 A.16.1.5)

#### Conclusion: Required Features
1. **Automatic capture**: Every playbook action generates evidence (no manual steps)
2. **Immutable storage**: S3 Object Lock prevents deletion or modification (even by AWS root)
3. **Structured format**: JSON + PDF (machine-readable + human-readable)
4. **Compliance tagging**: Each evidence item mapped to ISO 27001, GDPR, SOC 2 controls
5. **One-click reports**: Generate audit report for specific incident or time period

**Evidence Collected:**
- Action logs (who, what, when, result)
- API responses (OAuth revocation confirmation)
- User inputs (justifications, decisions)
- Incident timeline (start, steps, completion)
- Blast radius report (affected resources)

**Success Metric:** 100% of incidents have complete evidence within 10 minutes

---

### Feature 7: Context Enrichment for False Positive Reduction

#### Approach (PM)
- **Problem:** False positives cause alert fatigue; users ignore all alerts
- **Solution:** Use context (user role, historical behavior) to adjust risk scores
- **Example:** Admin user doing unusual action = lower risk than contractor doing same action

#### Approach (SA)
- **Technical Implementation:** Query Track 1 API for user context (role, data access, history)
- **Risk Multipliers**: Admin 0.5x, Employee with PII access 2.0x, First-time user 1.5x
- **Caching**: Cache user context for 5 minutes to reduce latency

#### Conclusion: Required Features
1. **User role context**: Query Track 1 for user role (admin/employee/contractor)
2. **Data sensitivity context**: What data does user have access to? (Restricted/Confidential/Internal)
3. **Historical baseline**: Compare current action to user's 7-day behavior pattern
4. **Risk score adjustment**: Multiply base risk score by context multipliers
5. **Transparency**: Show user why risk score was adjusted ("Risk reduced because you are an admin")

**Risk Score Formula:**
```
Final Risk Score = Base Risk Score × Role Multiplier × Data Sensitivity Multiplier × Behavior Multiplier
```

**Success Metric:** False positive rate <10% (down from 25% without context)

---

### Feature 8: Stateful Execution Engine with Fault Tolerance

#### Approach (PM)
- **Problem:** Playbooks fail mid-execution due to API timeouts, rate limits
- **Solution:** Execution engine that retries failures, handles partial completion
- **Use Case:** Google Workspace API times out → System retries 3 times → If still fails, alerts IT admin

#### Approach (SA)
- **Technical Implementation:** AWS Step Functions (state machines)
- **Retry Logic**: Exponential backoff (1s, 2s, 4s, 8s, 16s) up to 5 retries
- **Partial Failure Handling**: Continue other actions even if one fails (parallel execution)

#### Conclusion: Required Features
1. **State machine orchestration**: AWS Step Functions for stateful, resumable execution
2. **Automatic retry**: Exponential backoff for transient failures (network, rate limit)
3. **Parallel execution**: Revoke Google, M365, Slack, AWS in parallel (not sequential)
4. **Partial failure reporting**: "3 of 4 actions succeeded; M365 failed (rate limit)"
5. **Manual fallback**: If automated action fails after 5 retries, show manual instructions

**Failure Scenarios Handled:**
- Network timeout → Retry
- API rate limit → Exponential backoff
- Invalid credentials → Alert IT admin (cannot auto-fix)
- Service outage → Manual fallback instructions

**Success Metric:** >99% of playbooks complete without manual intervention

---

### Feature 9: Event-Driven Integration with Track 2 (AI Detection)

#### Approach (PM)
- **Problem:** Manual playbook triggering = delayed response = breach escalates
- **Solution:** AI detection automatically triggers appropriate playbook
- **Use Case:** Prompt injection detected (risk score 92) → Account Compromise playbook auto-starts

#### Approach (SA)
- **Technical Implementation:** EventBridge for async, decoupled communication
- **Event Schema**: Standardized format (threat_type, risk_score, user_id, evidence_url)
- **Routing Logic**: Lambda function maps event type to playbook

#### Conclusion: Required Features
1. **EventBridge integration**: Track 2 publishes events, Track 1 consumes
2. **Event schema**: Standardized JSON format (frozen by Sprint 2)
3. **Playbook routing**: Map event type to playbook (prompt_injection → Account Compromise)
4. **Risk threshold**: Only trigger if risk_score > threshold (configurable per tenant)
5. **Manual override**: IT admin can disable auto-trigger (require manual approval)

**Event-to-Playbook Mapping:**
| Event Type | Risk Score Threshold | Triggered Playbook |
|-----------|---------------------|-------------------|
| `ai.threat.detected` (prompt injection) | >85 | Account Compromise |
| `ai.dlp.violation` (critical data) | >60 | Account Compromise |
| `ai.deepfake.detected` | >90 | Account Compromise |
| `ai.shadow_tool.detected` | Any | Shadow IT Detected |

**Success Metric:** <2 minutes from AI detection to playbook start (automated)

---

### Feature 10: RBAC Policy Engine for Playbook Authorization

#### Approach (PM)
- **Problem:** Not all users should execute all playbooks (security risk)
- **Solution:** Role-based access control: only authorized users can trigger playbooks
- **Example:** HR manager can trigger Offboarding, but NOT Account Compromise (requires IT admin)

#### Approach (SA)
- **Technical Implementation:** OPA/Rego policy engine (same as Track 1 RBAC)
- **Policy Evaluation**: <100ms, cached for 5 minutes
- **Audit Log**: Every authorization decision logged (who, what, result, timestamp)

#### Conclusion: Required Features
1. **Role-based policies**: Define which roles can execute which playbooks
2. **Policy engine**: OPA/Rego for centralized, auditable authorization
3. **Least-privilege**: Default deny; explicit allow per role
4. **Emergency override**: IT admin can grant temporary access (logged, expires in 1 hour)
5. **Audit trail**: Every playbook execution logs who authorized it

**Default RBAC Policies:**
| Role | Allowed Playbooks |
|------|------------------|
| IT Admin | All playbooks |
| HR Manager | Offboarding Emergency, Inactive Account |
| Manager | Unauthorized Access (for their team only) |
| Employee | Shadow IT Detected (self-service) |
| Contractor | None (must request IT admin) |

**Success Metric:** Zero unauthorized playbook executions (enforced by RBAC)

---

## Feature Prioritization Matrix

| Feature | PM Priority | SA Priority | Implementation Complexity | Sprint |
|---------|------------|------------|--------------------------|--------|
| 1. Wizard UI | P0 | P1 | Medium | S8 |
| 2. Decision Gates | P0 | P1 | Low | S8 |
| 3. Pause/Resume | P1 | P0 | Medium | S8 |
| 4. Selective Undo | P2 | P1 | Medium | S9 |
| 5. Multi-Channel Notifications | P0 | P1 | Medium | S8-S10 |
| 6. Automatic Evidence | P0 | P0 | High | S10 |
| 7. Context Enrichment | P1 | P0 | High | S6, S10 |
| 8. Stateful Execution | P0 | P0 | High | S8 |
| 9. Event-Driven Integration | P1 | P0 | Medium | S10 |
| 10. RBAC Policy Engine | P1 | P0 | Medium | S5, S8 |

**Legend:**
- P0 = Must-have for V1 launch
- P1 = Important for V1, can be simplified
- P2 = Nice-to-have, can defer to V1.1

---

## Debate Outcomes: Key Agreements

### Agreement 1: UX Simplicity > Technical Elegance
**PM:** "If users can't execute playbooks without training, the product fails."  
**SA:** "Agreed. Technical sophistication must be invisible. User sees wizard, we handle complexity."

### Agreement 2: Security Cannot Be Compromised for UX
**PM:** "We need undo support for user confidence."  
**SA:** "Only for reversible actions. Account suspension is irreversible by design."  
**Consensus:** Selective undo with explicit reversibility metadata per action type.

### Agreement 3: Fault Tolerance > Performance
**PM:** "If playbooks fail 10% of the time, users lose trust."  
**SA:** "Step Functions with retry logic ensures >99% success rate. Worth the 100ms overhead."  
**Consensus:** Reliability is more important than sub-second latency.

### Agreement 4: Evidence Collection Must Be Automatic
**PM:** "Manual documentation = compliance failure."  
**SA:** "S3 Object Lock ensures immutability. No shortcuts."  
**Consensus:** 100% automatic evidence collection, no manual steps.

### Agreement 5: Context Enrichment Is Non-Negotiable
**PM:** "False positives kill adoption."  
**SA:** "Track 1 integration adds 50-100ms but reduces FP by 60%."  
**Consensus:** Accept latency cost for accuracy gain.

---

## Debate Outcomes: Key Disagreements (Resolved)

### Disagreement 1: Undo Scope
**PM:** "Users need undo for everything to feel safe."  
**SA:** "Undoing account suspension re-enables attacker. Security risk."  
**Resolution:** Selective undo only for reversible actions (allow-list, block-list). Irreversible actions require separate restore workflow with manager approval.

### Disagreement 2: Manual Fallback
**PM:** "If automation fails, user must have manual instructions."  
**SA:** "Manual instructions = inconsistent execution = compliance gap."  
**Resolution:** Manual fallback only for <5% edge cases (vendor API down). Instructions are step-by-step, logged as manual execution, generate same evidence format.

### Disagreement 3: Auto-Trigger vs Manual Approval
**PM:** "Auto-triggering playbooks is risky. What if false positive?"  
**SA:** "Manual approval adds 5-10 minutes delay. Breach escalates."  
**Resolution:** Auto-trigger for risk_score >85 (high confidence). Risk_score 60-85 requires manual approval. Configurable per tenant.

---

## Implementation Roadmap

### Sprint 8 (W15-16): Foundation
- Feature 1: Wizard UI (web)
- Feature 2: Decision Gates
- Feature 3: Pause/Resume
- Feature 5: Multi-Channel Notifications (Email + Slack)
- Feature 8: Stateful Execution Engine (Step Functions)

### Sprint 9 (W17-18): Mobile + Refinement
- Feature 4: Selective Undo
- Feature 5: Mobile Push Notifications (FCM Android)

### Sprint 10 (W19-20): Integration
- Feature 6: Automatic Evidence Collection
- Feature 7: Context Enrichment (Track 1 integration)
- Feature 9: Event-Driven Integration (Track 2)
- Feature 10: RBAC Policy Engine (playbook authorization)

### Sprint 11-13: Pilot + Hardening
- Threshold tuning (Feature 7)
- False positive reduction
- Performance optimization
- Security review

---

## Success Metrics Summary

| Feature | Success Metric | Target | Measurement Method |
|---------|---------------|--------|-------------------|
| Wizard UI | Completion rate | >95% | Telemetry: playbooks started vs completed |
| Decision Gates | External help requests | <5% | Support tickets tagged "playbook help" |
| Pause/Resume | Abandonment due to interruption | <2% | User survey post-incident |
| Selective Undo | Escalations to consultants | -30% | Compare pre/post launch |
| Multi-Channel Notifications | P0 response time | <2 min median | Notification sent → playbook started |
| Automatic Evidence | Evidence completeness | 100% | Audit: incidents with complete evidence |
| Context Enrichment | False positive rate | <10% | User reports "false alarm" |
| Stateful Execution | Playbook success rate | >99% | Playbooks completed without manual intervention |
| Event-Driven Integration | Auto-trigger latency | <2 min | AI detection → playbook start |
| RBAC Policy Engine | Unauthorized executions | 0 | Audit log: denied playbook attempts |

---

## Conclusion

**10 Essential Features Identified:**

1. **Wizard-Driven UI** — Guides non-security staff step-by-step
2. **Decision Gates** — Binary choices eliminate need for expertise
3. **Pause/Resume** — Handles real-world interruptions
4. **Selective Undo** — Reduces fear of mistakes (reversible actions only)
5. **Multi-Channel Notifications** — Ensures P0 incidents never missed
6. **Automatic Evidence Collection** — Compliance without manual work
7. **Context Enrichment** — Reduces false positives via user/data context
8. **Stateful Execution Engine** — Fault-tolerant, resumable playbooks
9. **Event-Driven Integration** — Auto-trigger from AI detection
10. **RBAC Policy Engine** — Least-privilege playbook authorization

**Debate Consensus:**
- PM and SA agree on all 10 features as essential for V1
- Disagreements resolved through selective implementation (e.g., selective undo, configurable auto-trigger)
- Features prioritized by risk mitigation (highest probability × impact first)

**Next Steps:**
- Sprint 8 kickoff: Implement features 1, 2, 3, 5, 8
- Joint T1-T2 schema session (Week 1) for feature 9
- Pilot validation (Sprint 11-12) to tune feature 7 thresholds
