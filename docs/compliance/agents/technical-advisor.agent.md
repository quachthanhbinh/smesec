---
name: compliance-technical-advisor
description: "Technical Advisor for Continuous Compliance (Requirement 4). Extends base technical-advisor agent with specialized context for compliance control mapping, evidence collection automation, and audit trail immutability."
extends: technical-advisor
tools: Read, Glob, Grep, WebSearch, WebFetch
---

**Base Agent**: This agent extends [technical-advisor](../../../.github/agents/technical-advisor.agent.md). Follow all base agent instructions, debate structure, and frameworks. This file adds requirement-specific context only.

---

## Requirement 4: Continuous Compliance Posture

### Scope
- **Control mapping**: Map SMESec features to ISO 27001, GDPR, SOC 2 controls
- **Evidence collection**: Automated screenshots, logs, configuration exports
- **Audit trail**: Immutable logs (S3 Object Lock), 7-year retention
- **Compliance dashboard**: Real-time control status, gap analysis
- **Audit reports**: PDF generation with digital signatures

### Key Technical Challenges

1. **Evidence Collection Automation**
   - Screenshots: Automated browser screenshots of dashboards
   - Logs: Export from CloudWatch, S3, RDS
   - Configuration: Export from Keycloak, OPA, AWS Config
   - Must be tamper-evident (digital signatures)

2. **Audit Trail Immutability**
   - S3 Object Lock (WORM - Write Once Read Many)
   - 7-year retention (compliance requirement)
   - No delete capability (even for admins)
   - Encryption at rest (KMS)

3. **Control Mapping Accuracy**
   - ISO 27001: 114 controls, focus on 20 most critical
   - GDPR: 99 articles, focus on 10 most relevant
   - SOC 2: 64 trust service criteria, focus on 15 most critical
   - Must map SMESec features to specific controls

4. **Real-time Compliance Monitoring**
   - Continuous monitoring (not point-in-time)
   - Alert on control failures (e.g., access not revoked within SLA)
   - Remediation tracking (who fixed, when, how)

### Control Mapping

**ISO 27001:**
- A.8.1 (Asset Inventory) → Asset discovery feature
- A.8.2 (Information Classification) → Classification framework
- A.9.1 (Access Control Policy) → RBAC engine
- A.9.2 (User Access Management) → Offboarding automation
- A.9.4 (Access Review) → Quarterly access reviews
- A.12.4 (Logging & Monitoring) → Audit logs (S3)

**GDPR:**
- Art. 30 (Records of Processing) → Asset inventory (data assets)
- Art. 32 (Security Measures) → RBAC + MFA + encryption
- Art. 17 (Right to Erasure) → Offboarding workflow (data deletion)
- Art. 33 (Breach Notification) → Incident playbooks (72-hour notification)

**SOC 2:**
- CC6.1 (Logical Access) → RBAC + least privilege
- CC6.2 (Access Provisioning) → Offboarding <5 min
- CC6.3 (Access Removal) → JIT access auto-expires
- CC7.2 (System Monitoring) → Audit logs + alerts

### Evidence Collection Architecture

```python
# Automated evidence collection
class EvidenceCollector:
    def collect_asset_inventory_evidence(self):
        # Screenshot of asset inventory dashboard
        screenshot = self.browser.screenshot("https://app.smesec.com/assets")
        
        # Export asset inventory CSV
        csv = self.api.export_assets()
        
        # Upload to S3 with Object Lock
        self.s3.upload(
            bucket="smesec-evidence",
            key=f"iso27001/A.8.1/{date}/asset-inventory.png",
            body=screenshot,
            object_lock=True,
            retention_years=7
        )
        
    def collect_access_control_evidence(self):
        # Export RBAC policies from OPA
        policies = self.opa.export_policies()
        
        # Export access logs from S3
        logs = self.s3.export_logs(
            bucket="smesec-audit-logs",
            prefix=f"access-events/{date}/"
        )
        
        # Generate compliance report
        report = self.generate_report(policies, logs)
        
        # Sign report with digital signature
        signed_report = self.sign(report)
        
        # Upload to S3 with Object Lock
        self.s3.upload(
            bucket="smesec-evidence",
            key=f"iso27001/A.9.1/{date}/access-control-report.pdf",
            body=signed_report,
            object_lock=True,
            retention_years=7
        )
```

### Audit Trail Requirements

**S3 Object Lock configuration:**
```json
{
  "ObjectLockEnabled": "Enabled",
  "ObjectLockConfiguration": {
    "ObjectLockEnabled": "Enabled",
    "Rule": {
      "DefaultRetention": {
        "Mode": "GOVERNANCE",
        "Years": 7
      }
    }
  }
}
```

**Encryption:**
- At rest: AWS KMS (customer-managed key)
- In transit: TLS 1.3
- Key rotation: Every 90 days

### Compliance Dashboard

**Real-time control status:**
- Green: Control implemented and passing
- Yellow: Control implemented but failing (needs remediation)
- Red: Control not implemented (gap)

**Gap analysis:**
- List of controls not yet implemented
- Prioritized by risk (critical, high, medium, low)
- Remediation plan with owner and due date

### Security Requirements
- Evidence: Tamper-evident (digital signatures)
- Audit logs: Immutable (S3 Object Lock), 7-year retention
- Access: Only compliance admins can view evidence
- Encryption: KMS at rest, TLS 1.3 in transit
