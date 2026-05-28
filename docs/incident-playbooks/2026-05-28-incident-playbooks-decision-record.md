# Decision Record: Incident Playbooks Architecture

**Date:** 2026-05-28  
**Status:** Approved  
**Deciders:** Solution Architect (30 years cybersecurity), PM/Risk Manager (30 years), Tech Lead  
**Related Documents:** [incident-playbooks-analysis.md](incident-playbooks-analysis.md), [Track 1 Requirements](../track1-foundation/requirements.md)

---

## Context and Problem Statement

SMEs (10-500 employees) lack dedicated security teams but face increasing security incidents requiring rapid, expert-level response. Traditional incident response requires specialized knowledge that SMEs cannot afford to hire or train. We need a system that enables **non-security staff** (IT admins, HR managers, employees) to execute complex security incident response procedures with confidence and compliance.

**Key Constraints:**
- Target users have minimal security training
- Response time critical: containment must happen in <5 minutes for P0 incidents
- Must generate compliance-ready evidence (ISO 27001, GDPR, SOC 2)
- Must integrate with both Track 1 (deterministic controls) and Track 2 (AI detection)
- Cannot require manual security expertise for 90%+ of incidents

**Success Criteria:**
- Non-security staff can execute playbooks without external help
- 80% reduction in incident response time vs manual process
- 100% of incidents generate audit-ready evidence
- <10% false positive rate (user frustration)

---

## Decision Drivers

1. **Accessibility**: Must be usable by non-technical staff without security background
2. **Reliability**: Cannot fail mid-execution; must handle partial failures gracefully
3. **Compliance**: Every action must be logged and evidence must be immutable
4. **Integration**: Must work seamlessly with Track 1 (access control) and Track 2 (AI detection)
5. **Speed**: Containment actions must complete in <5 minutes
6. **Auditability**: Complete forensic trail for post-incident review and compliance audits

---

## Decisions

### Decision 1: AWS Step Functions as Playbook Engine

**Chosen:** AWS Step Functions (state machines)

**Rationale:**
- **Stateful execution**: Playbooks can pause/resume without losing progress
- **Built-in retry logic**: Exponential backoff for transient failures
- **Visual workflow**: Non-technical users can understand execution flow
- **Fault tolerance**: Partial failures don't crash entire playbook
- **Event-driven**: Native EventBridge integration for auto-triggering
- **Audit trail**: Every state transition logged automatically

**Alternatives Considered:**

| Alternative | Why Rejected |
|------------|--------------|
| **Temporal.io** | Requires dedicated infrastructure; overkill for SME scale; team lacks Go expertise |
| **Airflow** | Designed for data pipelines, not real-time incident response; 5-10s latency unacceptable |
| **Custom orchestrator (Python/Celery)** | Would need to build retry, state persistence, monitoring from scratch; 3-4 sprints of work |
| **GitHub Actions** | Not designed for real-time execution; no sub-minute SLA; poor failure handling |

**Consequences:**
- ✅ Proven AWS service with 99.9% SLA
- ✅ Visual workflow editor for debugging
- ✅ Native integration with Lambda, SNS, SES, EventBridge
- ⚠️ Vendor lock-in to AWS (mitigated: entire platform is AWS-based)
- ⚠️ Learning curve for state machine JSON (mitigated: templates provided)

**Implementation Timeline:** Sprint 8 (W15-16)

---

### Decision 2: Wizard UI Pattern (Not CLI or Chatbot)

**Chosen:** Step-by-step wizard UI with decision gates (web + mobile)

**Rationale:**
- **Guided experience**: User cannot skip critical steps or get lost
- **Decision gates**: Yes/No questions instead of open-ended analysis
- **Progress visibility**: User knows how far along they are
- **Undo support**: Reversible actions reduce fear of mistakes
- **Mobile-first**: P0 incidents require response from anywhere (Sprint 9-10)

**Alternatives Considered:**

| Alternative | Why Rejected |
|------------|--------------|
| **CLI-based playbooks** | Requires terminal comfort; error-prone for non-technical users; no mobile support |
| **Chatbot interface** | Ambiguous input; hard to enforce required steps; poor for compliance audit trail |
| **Dashboard with buttons** | No guidance; user must know what to do; high error rate in testing |
| **Runbook documents (PDF)** | Manual execution; no automation; no evidence collection; 10x slower |

**Consequences:**
- ✅ Lowest cognitive load for non-security staff
- ✅ Enforces best practices (cannot skip steps)
- ✅ Mobile app enables response from anywhere
- ⚠️ More UI development work (mitigated: reusable wizard component)
- ⚠️ Less flexible than free-form interface (acceptable: consistency > flexibility)

**Implementation Timeline:** Sprint 8 (web), Sprint 9-10 (mobile)

---

### Decision 3: 5 Core Playbooks (Not Customizable in V1)

**Chosen:** Ship 5 pre-built playbooks; custom playbook builder deferred to V1.1

**Playbooks:**
1. Account Compromise (AI injection, suspicious login, impossible travel)
2. Offboarding Emergency (immediate termination)
3. Shadow IT Detected (unapproved OAuth app or AI tool)
4. Unauthorized Access (RBAC violation, privilege escalation attempt)
5. Inactive Account (>90 days unused)

**Rationale:**
- **80/20 rule**: These 5 cover 85%+ of SME security incidents (validated with 12 SME CISOs)
- **Time to market**: Custom builder adds 4-6 sprints; delays launch
- **Quality over quantity**: 5 polished playbooks > 20 half-baked ones
- **Validation**: Pilot customers can test thoroughly with limited scope

**Alternatives Considered:**

| Alternative | Why Rejected |
|------------|--------------|
| **10+ playbooks in V1** | Scope creep; testing burden; pilot customers overwhelmed |
| **Custom playbook builder in V1** | 4-6 additional sprints; complex UI; high bug risk; delays launch |
| **Only 3 playbooks** | Insufficient coverage; misses critical scenarios (unauthorized access, inactive accounts) |
| **Industry-standard playbooks (NIST, SANS)** | Too complex for non-security staff; 50+ steps; require security expertise |

**Consequences:**
- ✅ Focused scope enables thorough testing
- ✅ Pilot customers can validate all playbooks
- ✅ Faster time to market (launch Sprint 13)
- ⚠️ Some customers may need custom playbooks (mitigated: V1.1 roadmap item; can manually adapt existing playbooks)
- ⚠️ Edge cases may not be covered (mitigated: "Get help" button connects to IT admin)

**Implementation Timeline:** Sprint 8 (3 playbooks), Sprint 9 (2 playbooks)

---

### Decision 4: Multi-Channel Notifications with Priority-Based Routing

**Chosen:** Email + Slack + Mobile Push, routed by incident priority (P0/P1/P2/P3)

**Routing Rules:**
- **P0 (Critical)**: Email + Slack + Mobile Push (immediate, <1 min)
- **P1 (High)**: Email + Slack (within 5 min)
- **P2/P3 (Medium/Low)**: Email only (within 30 min)

**Rationale:**
- **Attention management**: P0 incidents need immediate response; P2/P3 should not interrupt
- **Multi-channel redundancy**: If Slack is down, email still works
- **Mobile push**: Enables response outside office hours
- **Cost efficiency**: SMS expensive ($0.01-0.05 per message); push notifications free

**Alternatives Considered:**

| Alternative | Why Rejected |
|------------|--------------|
| **Email only** | Delayed response; users don't check email constantly; no mobile urgency |
| **SMS for all incidents** | Expensive at scale ($500-2000/month for 100 employees); alert fatigue |
| **Slack only** | Single point of failure; not all SMEs use Slack; no mobile push |
| **Phone calls (Twilio)** | Expensive; requires voice menu; deferred to V1.1 for P0 escalation |
| **Single priority level** | Alert fatigue; users ignore all notifications; P0 incidents lost in noise |

**Consequences:**
- ✅ Balances urgency with cost
- ✅ Redundancy prevents missed critical incidents
- ✅ Mobile push enables 24/7 response
- ⚠️ Requires 3 notification integrations (mitigated: AWS SNS/SES + Slack API + FCM/APNs)
- ⚠️ Users must configure Slack/mobile (mitigated: onboarding wizard)

**Implementation Timeline:** Sprint 8 (Email + Slack), Sprint 9 (FCM Android), Sprint 10 (APNs iOS)

---

### Decision 5: Immutable Evidence Storage (S3 Object Lock, 7-Year Retention)

**Chosen:** AWS S3 with Object Lock (WORM mode), 7-year retention, KMS encryption

**Rationale:**
- **Compliance requirement**: ISO 27001, GDPR, SOC 2 require immutable audit logs
- **Forensic integrity**: Evidence cannot be tampered with post-incident
- **Legal defensibility**: Proves actions taken during incident response
- **Regulatory alignment**: 7 years matches most industry retention requirements

**Alternatives Considered:**

| Alternative | Why Rejected |
|------------|--------------|
| **PostgreSQL audit log** | Mutable; DBA can delete/modify; not compliance-ready |
| **CloudWatch Logs** | 90-day default retention; expensive for 7 years; not immutable |
| **Blockchain-based storage** | Overkill; expensive; slow writes; no SME demand |
| **Local file storage** | Not immutable; vulnerable to ransomware; no disaster recovery |

**Consequences:**
- ✅ Compliance-ready out of the box
- ✅ Cannot be deleted or modified (even by AWS root account)
- ✅ Automatic encryption at rest (KMS)
- ⚠️ Storage cost: ~$0.023/GB/month × 7 years (mitigated: evidence is small, <1GB/year per tenant)
- ⚠️ Cannot delete evidence even if customer requests (acceptable: compliance requirement)

**Implementation Timeline:** Sprint 10 (evidence collection), Sprint 11 (compliance reports)

---

### Decision 6: Event-Driven Integration with Track 2 (EventBridge)

**Chosen:** Track 2 publishes events to EventBridge → Lambda router → Step Functions playbook

**Event Schema:**
```json
{
  "source": "smesec.ai-detection",
  "detail-type": "ai.threat.detected | ai.dlp.violation | ai.deepfake.detected | ai.shadow_tool.detected",
  "detail": {
    "threat_type": "prompt_injection | dlp_violation | deepfake | shadow_ai",
    "risk_score": 0-100,
    "user_id": "alice@company.com",
    "timestamp": "ISO8601",
    "evidence_url": "s3://...",
    "metadata": { ... }
  }
}
```

**Rationale:**
- **Loose coupling**: Track 1 and Track 2 can evolve independently
- **Automatic triggering**: AI threats trigger playbooks without manual intervention
- **Scalability**: EventBridge handles 10K+ events/second
- **Audit trail**: Every event logged automatically
- **Flexibility**: Easy to add new event types or playbooks

**Alternatives Considered:**

| Alternative | Why Rejected |
|------------|--------------|
| **Direct API calls (Track 2 → Track 1)** | Tight coupling; Track 2 must know Track 1 API; synchronous blocking |
| **Polling (Track 1 polls Track 2 DB)** | Latency (15-60s); inefficient; scales poorly |
| **Kafka/RabbitMQ** | Requires dedicated infrastructure; overkill for SME scale; team lacks expertise |
| **Webhooks** | Requires public endpoint; security risk; no built-in retry; manual scaling |

**Consequences:**
- ✅ Decoupled architecture (Track 1 and Track 2 can deploy independently)
- ✅ Sub-second latency (event → playbook trigger)
- ✅ Built-in retry and dead-letter queue
- ⚠️ Requires schema coordination between teams (mitigated: joint T1-T2 schema session Sprint 1, frozen by Sprint 2)
- ⚠️ Debugging distributed events harder than direct calls (mitigated: CloudWatch event tracing)

**Implementation Timeline:** Sprint 1 (schema definition), Sprint 10 (integration)

---

### Decision 7: Context Enrichment from Track 1 (Reduce False Positives)

**Chosen:** Playbooks query Track 1 Asset Inventory + Access Governance for user context before taking action

**Context Used:**
- User role (admin / employee / contractor)
- Data sensitivity level user has access to
- Historical behavior baseline (7-day window)
- Current access patterns

**Risk Score Adjustment:**
- Admin user: 0.5x multiplier (less suspicious)
- Employee with PII access: 2.0x multiplier (higher risk)
- First-time AI tool user: 1.5x multiplier
- Repeated similar prompts: 0.7x multiplier (likely legitimate workflow)

**Rationale:**
- **Reduce false positives**: Same action by admin vs contractor has different risk
- **User experience**: Fewer false alarms = less alert fatigue
- **Accuracy**: Context improves detection precision from ~85% to >95% (validated in Track 2 Gate 2)

**Alternatives Considered:**

| Alternative | Why Rejected |
|------------|--------------|
| **No context (treat all users equally)** | 25-30% false positive rate; user frustration; playbooks ignored |
| **Manual context gathering (ask user)** | Slow; interrupts workflow; users don't know their own risk profile |
| **ML-based context (behavioral anomaly)** | Requires 30-90 days training data; deferred to V1.1; V1 uses rule-based |

**Consequences:**
- ✅ False positive rate reduced from ~25% to <10% (Track 2 Gate 2 target)
- ✅ Better user experience (fewer false alarms)
- ✅ Leverages existing Track 1 data (no new data collection)
- ⚠️ Playbooks depend on Track 1 API availability (mitigated: fallback to no-context mode if API down)
- ⚠️ Adds 50-100ms latency per playbook (acceptable: still <1s total)

**Implementation Timeline:** Sprint 6 (Track 2 context enrichment), Sprint 10 (Track 1 integration)

---

## Validation Gates

### Gate 1: Playbook Usability (End of Sprint 8)
**Criteria:**
- Non-security staff can complete playbook in <10 minutes without help
- 0 critical bugs (playbook hangs, data loss, incorrect revocation)
- User satisfaction >3.5/5.0 in internal testing

**Result:** PASS (internal testing with 5 non-technical employees)

### Gate 2: Integration Testing (End of Sprint 10)
**Criteria:**
- AI threat event → playbook trigger → action completed in <2 minutes
- 100% of events logged with correct schema
- 0 cross-tenant data leakage

**Result:** TBD (Sprint 10)

### Gate 3: Pilot Validation (End of Sprint 12)
**Criteria:**
- 5-10 pilot customers onboarded
- >90% of incidents resolved without escalation to security expert
- Customer satisfaction >4.0/5.0
- False positive rate <10%

**Result:** TBD (Sprint 12)

---

## Risks and Mitigations

### Risk 1: Users abandon playbooks mid-execution (high impact)
**Likelihood:** Medium (30%)  
**Impact:** High (incident not contained, compliance gap)

**Mitigations:**
- Wizard UI with clear progress indicators
- Ability to pause and resume
- "Get help" button connects to IT admin via Slack
- Post-incident survey to identify UX friction
- Sprint 11: Tune based on pilot feedback

**Residual Risk:** Low (5%) after mitigations

---

### Risk 2: Step Functions state machine bugs cause incorrect revocations (critical impact)
**Likelihood:** Low (10%)  
**Impact:** Critical (wrong account suspended, business disruption)

**Mitigations:**
- Extensive integration testing (Sprint 8-10)
- Dry-run mode for testing (no actual revocations)
- Manual approval gate for high-risk actions (optional per tenant)
- Undo support for reversible actions
- Pen-test validation (Sprint 12-13)

**Residual Risk:** Very Low (<2%) after mitigations

---

### Risk 3: False positives cause user frustration and playbook abandonment (high impact)
**Likelihood:** Medium (25% without context enrichment)  
**Impact:** High (users ignore all alerts, real incidents missed)

**Mitigations:**
- Context enrichment from Track 1 (Decision 7)
- Threshold tuning during pilot (Sprint 11-12)
- "Report false positive" button → feeds ML tuning
- Decision gates allow manual override
- Track 2 Gate 2: <10% false positive rate

**Residual Risk:** Low (8-10%) after mitigations

---

### Risk 4: EventBridge schema incompatibility between Track 1 and Track 2 (medium impact)
**Likelihood:** Medium (20%)  
**Impact:** Medium (integration broken, manual playbook triggering required)

**Mitigations:**
- Joint T1-T2 schema session in Sprint 1 Week 1 (Decision from Debate)
- Schema frozen by end of Sprint 2
- OpenAPI spec + shared TypeScript types (owned by Extension Eng)
- Integration testing in Sprint 10
- Schema versioning (v1, v2) for future changes

**Residual Risk:** Low (5%) after mitigations

---

## Compliance Mapping

| Standard | Control | How Playbooks Address |
|---------|---------|----------------------|
| **ISO 27001** | A.16.1.4 (Assessment of security events) | All incidents assessed via wizard decision gates; risk score calculated |
| **ISO 27001** | A.16.1.5 (Response to incidents) | 5 playbooks cover 85%+ of SME incidents; <5 min containment |
| **ISO 27001** | A.16.1.7 (Collection of evidence) | S3 Object Lock immutable storage; 7-year retention; PDF reports |
| **GDPR** | Article 33 (Breach notification) | Auto-generate breach report if PII accessed; 72-hour timeline tracked |
| **GDPR** | Article 32 (Security measures) | Automated containment (suspend accounts, revoke tokens) within 5 min |
| **SOC 2** | CC7.3 (Incident response) | Documented playbooks; audit trail; evidence collection; post-incident review |
| **SOC 2** | CC7.4 (Incident communication) | Multi-channel notifications (email, Slack, mobile); priority-based routing |

---

## Success Metrics (Post-Launch)

### Operational Metrics
- **Time to containment**: <5 minutes for P0 incidents (target: 95th percentile)
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
- **Cost avoidance**: Prevent data breaches ($50K-$500K per incident)

---

## Future Iterations

### V1.1 (6-12 months post-launch)
- Custom playbook builder (UI-based, no code)
- Voice call escalation (Twilio) for P0 incidents
- Multi-language support (Spanish, French, German)
- Behavioral anomaly detection (ML-based context)

### V1.2 (12-18 months post-launch)
- Automated remediation (no human approval for low-risk incidents)
- Federated learning (improve accuracy across customers, privacy-preserving)
- Integration with SIEM (Splunk, Datadog, Sumo Logic)
- Playbook marketplace (community-contributed playbooks)

---

## References

- [Incident Playbooks Analysis](incident-playbooks-analysis.md)
- [Track 1 Requirements](../track1-foundation/requirements.md)
- [Track 2 Requirements](../track2-ai-detection/requirements.md)
- [2-Track Approach Strategy](../strategy/2-track-approach.md)
- [Compliance Roadmap](../compliance/03-roadmap.md)

---

## Approval

**Approved by:**
- Solution Architect (Cybersecurity, 30 years) — 2026-05-28
- PM/Risk Manager (30 years) — 2026-05-28
- Tech Lead (Track 1) — 2026-05-28

**Next Review:** End of Sprint 12 (W24) — Pilot validation results
