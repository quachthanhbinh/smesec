# Incident Playbooks for SMESec Platform

**Date:** 2026-05-28  
**Status:** Approved  
**Key Requirement:** #5 - Incident playbooks executable by non-security staff

---

## Executive Summary

Incident playbooks are a critical component of the SMESec platform, enabling small and medium enterprises to respond effectively to security incidents **without requiring dedicated security expertise**. This document details the design, implementation, and integration of 5 core playbooks that guide non-technical staff through complex security response procedures.

**Core Value Proposition:** Transform security incident response from a specialized skill requiring years of training into a guided, step-by-step process that any IT admin or manager can execute confidently.

---

## Design Principles

### 1. Non-Security Staff First
- **Wizard-driven UI**: Step-by-step guidance with clear instructions in plain language
- **Decision gates**: Yes/No questions instead of open-ended security analysis
- **Progress indicators**: Visual feedback showing completion status and remaining steps
- **Undo support**: Allow reverting actions if mistakes are made
- **No assumed knowledge**: Every step explains WHY and WHAT, not just HOW

### 2. Fault Tolerance & Resumability
- **Stateful execution**: Playbooks can be paused and resumed without losing progress
- **Automatic retry**: Failed API calls retry with exponential backoff
- **Partial failure handling**: If one revocation fails, others continue
- **Audit trail**: Every action logged with timestamp, actor, and result

### 3. Multi-Channel Notification
- **Priority-based routing**:
  - P0 (Critical): Email + Slack + Mobile push (immediate)
  - P1 (High): Email + Slack (within 5 minutes)
  - P2/P3 (Medium/Low): Email only (within 30 minutes)
- **Escalation**: If no response within SLA, escalate to next level

### 4. Integration-Ready
- **Event-driven**: Playbooks can be triggered manually OR automatically via EventBridge
- **Track 2 AI detection**: AI threats (prompt injection, deepfake, DLP violation) auto-trigger appropriate playbooks
- **Evidence collection**: All actions generate compliance-ready reports (PDF + JSON)

---

## The 5 Core Playbooks

### Playbook 1: Account Compromise
**Trigger:** Suspicious login detected, impossible travel, or AI prompt injection with high risk score (>85)

**Target Audience:** IT admin, HR manager, or employee's direct manager

**Steps:**
1. **Verify the alert** (Decision gate: Is this a false positive?)
   - Review login location, device, time
   - Contact user via out-of-band channel (phone call, SMS)
   - If confirmed legitimate → Close playbook
   
2. **Immediate containment** (Automated)
   - Suspend account on Google Workspace, M365, Slack
   - Revoke all active sessions and OAuth tokens
   - Disable AWS IAM user if applicable
   - Estimated time: <2 minutes

3. **Assess impact** (Guided analysis)
   - Which resources did the account access in last 24 hours?
   - Was sensitive data (PII, IP, financial) accessed?
   - Were any changes made (file uploads, config changes)?
   - System generates blast radius report

4. **Notify stakeholders** (Automated)
   - Alert user's manager (P0)
   - Alert IT admin (P0)
   - Alert compliance officer if PII accessed (P1)
   - Generate incident report for legal/audit

5. **Remediation** (Guided)
   - Force password reset
   - Re-enable MFA (if disabled)
   - Review and revoke suspicious OAuth apps
   - Restore account access after verification

6. **Post-incident** (Automated + Manual)
   - Generate PDF incident report with timeline
   - Store evidence in S3 (immutable, 7-year retention)
   - Schedule follow-up review in 7 days
   - Update user security training status

**Success Criteria:**
- Account suspended within 2 minutes of trigger
- Incident report generated within 10 minutes
- 100% of actions logged for audit

---

### Playbook 2: Offboarding Emergency
**Trigger:** Employee terminated immediately (manual trigger from HR system or IT admin)

**Target Audience:** HR manager or IT admin

**Steps:**
1. **Confirm termination** (Decision gate: Is this an emergency offboarding?)
   - Verify employee ID and termination date
   - Confirm this is NOT a planned offboarding (those use automated workflow)
   
2. **Parallel revocation** (Automated via AWS Step Functions)
   - Google Workspace: suspend account, revoke OAuth, transfer Drive ownership
   - Microsoft 365: disable account, revoke sessions, convert mailbox to shared
   - Slack: deactivate account, remove from all channels
   - AWS: disable IAM user, revoke access keys, terminate EC2 instances owned by user
   - Estimated time: <5 minutes

3. **Physical access** (Manual checklist)
   - Retrieve laptop, phone, access cards
   - Disable building access badges
   - Collect company credit cards
   - Wizard provides printable checklist

4. **Data preservation** (Automated)
   - Backup user's email, Drive, OneDrive to compliance archive
   - Generate list of files created/modified in last 30 days
   - Identify any shared files that need ownership transfer

5. **Notification** (Automated)
   - Alert IT admin (P0)
   - Alert manager (P1)
   - Alert finance if user had expense access (P1)
   - Do NOT notify terminated employee

6. **Verification** (Automated)
   - Test: Can user still log in? (Should fail)
   - Test: Are OAuth tokens revoked? (Should be invalid)
   - Test: Can user access AWS? (Should be denied)
   - Generate pass/fail report

**Success Criteria:**
- All digital access revoked within 5 minutes
- Offboarding report (PDF) generated within 1 minute after completion
- Zero failed revocations (or immediate alert if any fail)

---

### Playbook 3: Shadow IT Detected
**Trigger:** Unapproved OAuth app detected OR employee using external AI tool with sensitive data

**Target Audience:** IT admin or security champion

**Steps:**
1. **Assess the tool** (Guided analysis)
   - What is the app/tool? (Auto-populated from OAuth metadata or browser extension)
   - Who authorized it? (User, department)
   - What data does it access? (OAuth scopes, browser intercept data)
   - Risk score: Auto-calculated (data sensitivity × tool reputation × usage frequency)

2. **Determine action** (Decision gate based on risk score)
   - **Low risk (0-30)**: Log and allow, add to monitoring list
   - **Medium risk (31-60)**: Request justification from user, manager approval required
   - **High risk (61-85)**: Block immediately, require IT admin review
   - **Critical risk (86-100)**: Block + revoke + incident report

3. **User notification** (Automated)
   - Email user explaining why tool was flagged
   - Provide approved alternatives (e.g., "Use Google Drive instead of Dropbox")
   - Link to shadow IT policy

4. **Revocation** (Automated if blocked)
   - Revoke OAuth token
   - Add domain to browser extension blocklist
   - Remove app from Google Workspace / M365 allowed apps

5. **Policy update** (Manual)
   - Add to allow-list if approved after review
   - Add to block-list if permanently denied
   - Update DLP policies if needed

6. **Trend analysis** (Automated)
   - Is this tool being used by multiple employees?
   - Is there a legitimate business need?
   - Generate recommendation: "5 employees using Notion — consider enterprise license"

**Success Criteria:**
- Risk assessment completed within 1 minute
- User notified within 5 minutes
- If blocked, access revoked within 2 minutes

---

### Playbook 4: Unauthorized Access
**Trigger:** User attempts to access resource without permission (RBAC denial logged)

**Target Audience:** IT admin or resource owner

**Steps:**
1. **Context gathering** (Automated)
   - Who: User ID, role, department
   - What: Resource attempted (file, database, API endpoint)
   - When: Timestamp, frequency (first time or repeated?)
   - Why: Was this a legitimate business need or suspicious?

2. **Pattern detection** (Automated)
   - Is this user repeatedly trying to access restricted resources?
   - Are multiple users trying to access the same resource? (Possible misconfiguration)
   - Is this a privilege escalation attempt?

3. **Decision gate** (Manual)
   - **Legitimate need**: Grant JIT access (time-limited, logged)
   - **Misconfiguration**: Update RBAC policy to grant permanent access
   - **Suspicious**: Trigger Account Compromise playbook
   - **Policy violation**: Notify manager, log for HR review

4. **JIT access** (if approved)
   - Grant access for specified duration (default: 2 hours)
   - Require justification (logged)
   - Auto-revoke after expiry
   - Send warning 10 minutes before expiry

5. **Audit log** (Automated)
   - Record decision and justification
   - Link to approver
   - Generate compliance report if sensitive resource

**Success Criteria:**
- Decision made within 15 minutes
- If JIT granted, access active within 1 minute
- 100% of decisions logged with justification

---

### Playbook 5: Inactive Account
**Trigger:** Account inactive for >90 days (automated detection)

**Target Audience:** IT admin or HR manager

**Steps:**
1. **Verify inactivity** (Automated)
   - Last login date across all systems (Google, M365, Slack, AWS)
   - Last file modification date
   - Last email sent/received
   - Generate inactivity report

2. **Determine status** (Decision gate)
   - **On leave**: Extend review period, set reminder for return date
   - **Contractor ended**: Trigger Offboarding Emergency playbook
   - **Zombie account**: Proceed to deactivation
   - **False positive**: Update activity detection rules

3. **Manager confirmation** (Manual)
   - Email manager: "Employee X has been inactive for 95 days. Still employed?"
   - If no response in 48 hours, escalate to HR
   - If confirmed inactive, proceed to step 4

4. **Graceful deactivation** (Automated)
   - Suspend account (do not delete)
   - Revoke active sessions
   - Archive data to compliance storage
   - Generate deactivation report

5. **Retention policy** (Automated)
   - Account remains suspended for 30 days (recovery window)
   - After 30 days, convert to "decommissioned" state
   - Data retained per compliance requirements (7 years)
   - Account metadata retained indefinitely for audit

**Success Criteria:**
- Manager notified within 24 hours of 90-day threshold
- If no response, escalation within 48 hours
- Deactivation completed within 1 hour of approval

---

## Technical Architecture

### Playbook Engine: AWS Step Functions
**Why Step Functions?**
- **Stateful**: Execution state persists across restarts
- **Fault-tolerant**: Automatic retry with exponential backoff
- **Visual workflow**: Easy to debug and audit
- **Scalable**: Handles concurrent playbook executions
- **Event-driven**: Integrates with EventBridge for auto-triggering

**Execution Model:**
```
EventBridge Event → Lambda (Playbook Router) → Step Functions State Machine → Actions (parallel)
                                                                           ↓
                                                                    Notification Service
                                                                           ↓
                                                                    Evidence Collection (S3)
```

### Wizard UI (Web + Mobile)
**Web (React/Next.js):**
- Step-by-step wizard with progress bar
- Decision gates rendered as radio buttons (Yes/No/Skip)
- Real-time status updates via WebSocket
- Undo button for reversible actions
- PDF report download

**Mobile (Flutter - iOS + Android):**
- Sprint 9: Asset inventory + JIT approval + incident wizard (read-only)
- Sprint 10: Full incident wizard (execute playbooks from mobile)
- Push notifications for P0/P1 incidents (FCM for Android, APNs for iOS)

### Notification System
**Architecture:**
- AWS SNS for email (via SES)
- Slack API for channel/DM notifications
- FCM/APNs for mobile push
- Priority-based routing (P0/P1/P2/P3)

**Delivery SLA:**
- P0: <1 minute (email + Slack + push)
- P1: <5 minutes (email + Slack)
- P2/P3: <30 minutes (email only)

### Evidence Collection
**Storage:** AWS S3 with Object Lock (immutable, 7-year retention)

**Evidence Types:**
- Incident reports (PDF + JSON)
- Action logs (who, what, when, result)
- Screenshots (if applicable)
- API responses (OAuth revocation confirmations)
- User justifications (for JIT access, shadow IT approvals)

**Compliance Mapping:**
- ISO 27001 A.16.1.4 (Assessment of information security events)
- ISO 27001 A.16.1.5 (Response to information security incidents)
- GDPR Article 33 (Notification of a personal data breach)
- SOC 2 CC7.3 (Incident response)

---

## Integration with Track 2 (AI Detection)

### Event-Driven Triggering
Track 2 publishes events to EventBridge when AI threats are detected:

**Event Schema:**
```json
{
  "source": "smesec.ai-detection",
  "detail-type": "ai.threat.detected",
  "detail": {
    "threat_type": "prompt_injection | dlp_violation | deepfake | shadow_ai",
    "risk_score": 0-100,
    "user_id": "alice@company.com",
    "timestamp": "2026-05-28T10:30:00Z",
    "evidence_url": "s3://smesec-evidence/...",
    "metadata": { ... }
  }
}
```

**Playbook Routing:**
| AI Threat Type | Risk Score | Triggered Playbook |
|---------------|------------|-------------------|
| Prompt Injection | >85 | Account Compromise |
| DLP Violation (critical data) | >60 | Account Compromise |
| Deepfake (voice/video) | >90 | Account Compromise |
| Shadow AI (unapproved tool) | Any | Shadow IT Detected |
| Repeated violations | >3 in 24h | Account Compromise |

### Enrichment from Track 1
Playbooks leverage Track 1 data for context:
- **User role** (admin/employee/contractor) → Adjust response severity
- **Data sensitivity** (Restricted/Confidential/Internal) → Determine notification priority
- **Historical patterns** (baseline behavior) → Reduce false positives
- **Asset dependencies** (blast radius) → Assess impact

---

## Success Metrics

### Operational Metrics
- **Time to containment**: <5 minutes for P0 incidents
- **Playbook completion rate**: >95% (not abandoned mid-execution)
- **False positive rate**: <10% (user reports "this was legitimate")
- **User satisfaction**: >4.0/5.0 (post-incident survey)

### Compliance Metrics
- **Incident report generation**: 100% within 10 minutes
- **Evidence retention**: 100% stored immutably for 7 years
- **Audit trail completeness**: 100% of actions logged

### Business Metrics
- **Reduction in incident response time**: Target 80% reduction vs manual process
- **Non-security staff capability**: >90% of incidents handled without security expert
- **Cost avoidance**: Prevent data breaches that would cost $50K-$500K per incident

---

## Rollout Plan

### Sprint 8 (W15-16): Foundation + 3 Playbooks
- Playbook engine (Step Functions)
- Wizard UI (web)
- Notification system (Email + Slack)
- Playbooks: Account Compromise, Offboarding Emergency, Shadow IT Detected

**Deliverable:** IT admin can run 3 playbooks end-to-end via web UI

### Sprint 9 (W17-18): 2 Additional Playbooks + Mobile
- Playbooks: Unauthorized Access, Inactive Account
- Mobile app: Asset inventory + JIT approval + incident wizard (basic)
- Push notifications (FCM for Android)

**Deliverable:** All 5 playbooks available; mobile app can approve JIT and view incidents

### Sprint 10 (W19-20): Track 2 Integration
- EventBridge event schema finalized
- AI threat events auto-trigger playbooks
- Mobile incident wizard (full execution capability)
- APNs push for iOS

**Deliverable:** AI detection → automatic playbook trigger → incident resolved

### Sprint 11-13: Pilot + Hardening
- Pilot customers test playbooks in real incidents
- Tune thresholds based on false positive feedback
- Performance optimization (playbook execution <5 min)
- Security review (pen-test playbook endpoints)

**Deliverable:** Production-ready playbooks validated by 5-10 pilot customers

---

## Risk Mitigation

### Risk 1: Non-technical users abandon playbooks mid-execution
**Mitigation:**
- Wizard UI with clear progress indicators
- Ability to pause and resume
- "Get help" button connects to IT admin via Slack
- Post-incident survey to identify UX friction

### Risk 2: Playbooks trigger false positives, causing user frustration
**Mitigation:**
- Context enrichment from Track 1 (user role, historical patterns)
- Decision gates allow manual override
- "Report false positive" button → feeds ML model tuning
- Threshold tuning during pilot phase (Sprint 11-12)

### Risk 3: Playbook actions fail (API rate limits, network issues)
**Mitigation:**
- Automatic retry with exponential backoff
- Partial failure handling (continue other actions)
- Immediate alert to IT admin if critical action fails
- Manual fallback instructions in wizard UI

### Risk 4: Compliance evidence is incomplete or lost
**Mitigation:**
- S3 Object Lock (immutable storage)
- Append-only logs (cannot be deleted or modified)
- Automated evidence collection (no manual steps)
- Daily backup verification

---

## Future Enhancements (Post-V1)

### V1.1 (6-12 months post-launch)
- **Custom playbook builder**: IT admin can create org-specific playbooks via UI
- **Voice call escalation**: Twilio integration for P0 incidents
- **Multi-language support**: Playbooks in Spanish, French, German
- **Behavioral anomaly detection**: Trigger playbooks based on user baseline deviation

### V1.2 (12-18 months post-launch)
- **Automated remediation**: Playbooks execute without human approval for low-risk incidents
- **Federated learning**: Improve playbook accuracy by learning across customers (privacy-preserving)
- **Integration with SIEM**: Export playbook events to Splunk, Datadog, etc.
- **Playbook marketplace**: Share community-contributed playbooks

---

## Conclusion

Incident playbooks are the bridge between **detection** (Track 1 + Track 2) and **response**. By making security incident response accessible to non-security staff, SMESec empowers SMEs to protect themselves without hiring expensive security teams.

**Key Success Factors:**
1. **Simplicity**: Wizard UI that anyone can follow
2. **Automation**: Reduce manual steps to <20% of total workflow
3. **Integration**: Seamless connection between AI detection and response
4. **Compliance**: Every action generates audit-ready evidence

**Next Steps:**
- Sprint 8 kickoff: Playbook engine architecture review
- Week 1: Finalize EventBridge schema with Track 2 team
- Week 2: Wizard UI mockups for user testing
