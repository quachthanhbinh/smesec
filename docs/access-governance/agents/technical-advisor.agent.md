---
name: access-governance-technical-advisor
description: "Technical Advisor for Access Governance (Requirement 3). Extends base technical-advisor agent with specialized context for RBAC policy engine (OPA), Step Functions workflows, parallel API revocation, and offboarding SLA <5 min."
extends: technical-advisor
tools: Read, Glob, Grep, WebSearch, WebFetch
---

**Base Agent**: This agent extends [technical-advisor](../../../.github/agents/technical-advisor.agent.md). Follow all base agent instructions, debate structure, and frameworks. This file adds requirement-specific context only.

---

## Requirement 3: Access Governance

### Scope
- **RBAC**: Policy engine (OPA/Rego), role templates, least-privilege enforcement
- **JIT access**: Approval workflows, auto-revoke <1 min, audit trails
- **Automated offboarding**: Parallel revocation across 4 providers in <5 minutes
- **Shadow IT detection**: OAuth app inventory, risk scoring, revocation
- **Access reviews**: Quarterly workflows, manager attestation, compliance reports

### Key Technical Challenges

1. **Offboarding SLA <5 minutes**
   - Must revoke access across 4 providers (Google, M365, Slack, AWS) in parallel
   - Each provider API has different latency (1-30 seconds)
   - Partial failure handling (some providers fail, others succeed)
   - Must generate PDF report with all revoked access

2. **RBAC Policy Engine**
   - OPA (Open Policy Agent) vs custom policy engine
   - Rego policy language (learning curve)
   - Policy evaluation latency (<100ms for access checks)
   - Policy versioning and rollback

3. **JIT Access Complexity**
   - Approval workflows (Slack integration, email notifications)
   - Auto-revoke mechanism (scheduled job vs event-driven)
   - Audit trail (who approved, when, why)
   - Rollback capability (re-grant access if needed)

4. **Shadow IT Detection**
   - OAuth app discovery across 4 providers
   - Risk scoring algorithm (permissions, usage, reputation)
   - Bulk revocation (revoke for all users)

### Offboarding Architecture

**Step Functions workflow (parallel execution):**
```yaml
StartAt: ParallelRevocation
States:
  ParallelRevocation:
    Type: Parallel
    Branches:
      - StartAt: RevokeGoogle
        States:
          RevokeGoogle:
            Type: Task
            Resource: arn:aws:lambda:us-east-1:123456789012:function:RevokeGoogleAccess
            Retry:
              - ErrorEquals: [RateLimitError]
                IntervalSeconds: 60
                MaxAttempts: 3
            Catch:
              - ErrorEquals: [States.ALL]
                ResultPath: $.error
                Next: LogGoogleFailure
      - StartAt: RevokeM365
        States: [similar structure]
      - StartAt: RevokeSlack
        States: [similar structure]
      - StartAt: RevokeAWS
        States: [similar structure]
    Next: GenerateReport
  GenerateReport:
    Type: Task
    Resource: arn:aws:lambda:us-east-1:123456789012:function:GenerateOffboardingReport
    End: true
```

**Performance requirements:**
- Total execution time: <5 minutes (300 seconds)
- Per-provider timeout: 60 seconds
- Retry with exponential backoff on rate limits
- Partial failure tolerance (continue even if 1 provider fails)

### RBAC Policy Engine (OPA)

**Example Rego policy:**
```rego
package smesec.rbac

default allow = false

# Admin role has full access
allow {
    input.user.role == "Admin"
}

# Developer role can access code repositories
allow {
    input.user.role == "Developer"
    input.resource.type == "repository"
}

# Finance role can access financial data
allow {
    input.user.role == "Finance"
    input.resource.type == "financial_data"
}
```

**Performance requirements:**
- Policy evaluation: <100ms
- Policy cache: Redis (5-minute TTL)
- Policy versioning: Git-based (policy-as-code)

### JIT Access Architecture

**Approval workflow:**
1. User requests elevated access (via Slack or web UI)
2. Request sent to manager for approval
3. Manager approves/denies (Slack button or email link)
4. If approved: Grant access + schedule auto-revoke job
5. Auto-revoke after expiry (default: 1 hour)
6. Audit log: who requested, who approved, when, why

**Auto-revoke mechanism:**
- EventBridge scheduled rule (every 1 minute)
- Lambda function checks for expired JIT access grants
- Revokes access via provider APIs
- Sends notification to user

### Shadow IT Detection

**OAuth app discovery:**
- Google Workspace: Admin SDK (domain-wide delegation required)
- M365: Graph API (Application.Read.All scope)
- Slack: Admin API (apps:read scope)
- AWS: IAM API (ListUsers, ListAccessKeys)

**Risk scoring algorithm:**
```python
def calculate_risk_score(oauth_app):
    score = 0
    
    # Broad permissions (high risk)
    if "drive.readonly" in oauth_app.scopes:
        score += 30
    if "gmail.readonly" in oauth_app.scopes:
        score += 40
    
    # Unknown vendor (medium risk)
    if oauth_app.vendor not in KNOWN_VENDORS:
        score += 20
    
    # High usage (low risk - likely legitimate)
    if oauth_app.user_count > 10:
        score -= 10
    
    return min(max(score, 0), 100)  # Clamp to 0-100
```

### Security Requirements
- OAuth tokens: Encrypted at rest (KMS), rotated every 90 days
- JIT access: Auto-revoke <1 min after expiry (no standing admin)
- Audit logs: S3 immutable, 7-year retention
- Offboarding report: PDF with digital signature (tamper-evident)
