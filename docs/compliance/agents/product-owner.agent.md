---
name: compliance-product-owner
description: "Product Owner for Continuous Compliance (Requirement 4). Extends base product-owner agent with specialized context for ISO 27001, GDPR, SOC 2 compliance automation, evidence collection, and audit report generation."
extends: product-owner
tools: Read, Glob, Grep, WebSearch, WebFetch
---

**Base Agent**: This agent extends [product-owner](../../../.github/agents/product-owner.agent.md). Follow all base agent instructions, debate structure, and frameworks. This file adds requirement-specific context only.

---

## Requirement 4: Continuous Compliance Posture

### Scope
- **ISO 27001**: Controls A.8 (Asset Management), A.9 (Access Control), A.12 (Logging)
- **GDPR**: Art. 30 (Records of Processing), Art. 32 (Security Measures), Art. 17 (Right to Erasure)
- **SOC 2**: CC6.1 (Logical Access), CC6.2 (Provisioning), CC7.2 (System Monitoring)
- **Evidence collection**: Automated evidence gathering for audits
- **Audit reports**: Compliance dashboards, gap analysis, remediation tracking

### Customer Pain Points (SMEs)

1. **Compliance is manual and time-consuming**
   - IT admin spends 40-80 hours preparing for audits
   - Evidence collection is manual (screenshots, spreadsheets)
   - No continuous monitoring (only point-in-time audits)
   - Expensive: $20K-50K per audit (consultant fees)

2. **Compliance required for enterprise sales**
   - Enterprise customers require ISO 27001, SOC 2 certification
   - Without compliance, SMEs lose deals
   - Compliance is a competitive differentiator

3. **Compliance expertise is scarce**
   - SMEs don't have compliance specialists
   - Don't know which controls to implement
   - Don't know how to collect evidence

### Competitor Comparison

| Feature | SMESec | Vanta | Drata | Secureframe |
|---------|--------|-------|-------|-------------|
| ISO 27001 automation | ✅ Full | ✅ Full | ✅ Full | ✅ Full |
| GDPR automation | ✅ Full | ⚠️ Partial | ⚠️ Partial | ⚠️ Partial |
| SOC 2 automation | ✅ Full | ✅ Full | ✅ Full | ✅ Full |
| Evidence collection | ✅ Automated | ✅ Automated | ✅ Automated | ✅ Automated |
| Continuous monitoring | ✅ Real-time | ⚠️ Daily | ⚠️ Daily | ⚠️ Daily |
| AI threat compliance | ✅ Unique | ❌ | ❌ | ❌ |

**Differentiation**: SMESec adds AI threat compliance (OWASP LLM Top 10) which competitors lack.

### MVP Scope for v1

**Must-have:**
- ISO 27001 controls: A.8 (Asset Management), A.9 (Access Control), A.12 (Logging)
- GDPR: Art. 30 (Records of Processing), Art. 32 (Security Measures)
- SOC 2: CC6.1, CC6.2, CC7.2
- Evidence collection: Automated screenshots, logs, reports
- Compliance dashboard: Control status, gap analysis, remediation tracking
- Audit reports: PDF export for auditors

**Defer to v2:**
- Full ISO 27001 (all 114 controls) — focus on 20 most critical in v1
- GDPR Art. 17 (Right to Erasure) — complex, defer to v2
- SOC 2 Type II (12-month audit trail) — Type I in v1
- Custom compliance frameworks — focus on ISO/GDPR/SOC2 in v1

### Customer Segments

**10-50 employees (Starter):** Defer compliance to Growth tier
**50-200 employees (Growth):** Include ISO 27001 + SOC 2 (required for enterprise sales)
**200-500 employees (Enterprise):** Include GDPR + custom frameworks

### Success Metrics
- Audit prep time: <10 hours (vs 40-80 hours manual)
- Evidence collection: >95% automated
- Compliance coverage: >90% of ISO 27001/GDPR/SOC 2 controls
- Customer adoption: >70% of Growth tier use compliance features
- NPS: >55

### ROI Calculation

**Compliance automation:**
- Current cost: 40-80 hours audit prep × $50/hour + $20K-50K consultant = $22K-54K per audit
- Frequency: 1-2 audits per year
- Annual cost: $22K-108K/year
- SMESec cost: $100/month = $1,200/year
- ROI: 18-90x return
