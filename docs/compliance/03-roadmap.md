# Lộ Trình Triển Khai Tuân Thủ

## Tổng Quan

**Thời gian:** 6-9 tháng (song song với phát triển v1)  
**Chiến lược:** Compliance-by-Design - xây dựng tuân thủ ngay từ đầu  
**Mục tiêu:** ISO 27001 + SOC 2 Type 1 (Tháng 6-8) → SOC 2 Type 2 (Tháng 12)

## Timeline Tổng Thể

```
Tháng 1-4: FOUNDATION (Nền móng)
    ├─ Setup công cụ automation
    ├─ Cấu hình hạ tầng tuân thủ
    └─ Tích hợp security scanning

Tháng 5-6: AUTOMATION & SOC 2 TYPE 1
    ├─ Triển khai Vanta
    ├─ Hoàn thiện documentation
    └─ Milestone 1: SOC 2 Type 1 + V1 Release

Tháng 7-8: ISO 27001 CERTIFICATION
    ├─ Internal audit
    ├─ External certification audit
    └─ Milestone 2: ISO 27001 Certificate

Tháng 9-11: MAINTENANCE & MONITORING
    ├─ Continuous compliance
    ├─ Quarterly reviews
    └─ Evidence collection

Tháng 12: SOC 2 TYPE 2
    └─ Milestone 3: SOC 2 Type 2 Report
```

## Phase 1: Foundation (Tháng 1-4)

### Mục Tiêu
Thiết lập nền móng kỹ thuật và tổ chức cho tuân thủ, song song với việc phát triển v1.

### Tháng 1: Infrastructure Setup

**Week 1-2: AWS Security Baseline**
- [ ] Setup AWS IAM Identity Center
  - Tạo user groups: Developers, DevOps, Admins
  - Cấu hình MFA bắt buộc cho tất cả accounts
  - Implement least privilege access policies
- [ ] Enable AWS CloudTrail
  - Tạo dedicated S3 bucket cho logs
  - Enable Object Lock (WORM) để prevent deletion
  - Configure log file validation
- [ ] Setup AWS Config
  - Enable compliance rules (encryption, public access, etc.)
  - Configure SNS notifications cho violations

**Week 3-4: GitHub Security**
- [ ] Enable GitHub Advanced Security
  - Activate Dependabot alerts
  - Enable CodeQL scanning
  - Configure secret scanning
- [ ] Configure branch protection
  - Require PR reviews (min 1 Senior Engineer)
  - Block direct push to main/production
  - Require status checks to pass
- [ ] Create security workflow
  ```yaml
  # .github/workflows/security.yml
  name: Security Checks
  on: [pull_request]
  jobs:
    security:
      - Dependabot scan
      - CodeQL analysis
      - Secret scanning
      → Block merge if High/Critical
  ```

**Deliverables:**
- ✅ AWS security baseline configured
- ✅ GitHub security features enabled
- ✅ Security CI/CD pipeline active

### Tháng 2: Storage & Database Security

**Week 1-2: Cloudflare R2 Configuration**
- [ ] Setup R2 buckets với encryption
  - Enable server-side encryption (AES-256)
  - Configure bucket policies (block public access)
  - Setup access logging
- [ ] Implement access controls
  - Create service accounts với least privilege
  - Configure API tokens với expiration
  - Document access procedures

**Week 3-4: Database Security**
- [ ] Configure RDS/Database
  - Place in private subnets (no public IP)
  - Enable encryption at rest
  - Enable automated backups (7-day retention)
  - Configure SSL/TLS for connections
- [ ] Setup database access controls
  - Create application-specific DB users
  - Implement connection pooling với credentials rotation
  - Document connection procedures

**Deliverables:**
- ✅ All storage encrypted at rest
- ✅ Databases in private subnets
- ✅ Access logging enabled

### Tháng 3: Infrastructure as Code

**Week 1-2: IaC Templates**
- [ ] Create Terraform/IaC templates
  - S3 bucket template (với encryption + logging)
  - RDS template (với private subnet + encryption)
  - IAM role template (least privilege)
  - Security group template (restrictive by default)
- [ ] Implement compliance checks
  - Pre-commit hooks for IaC validation
  - CI/CD checks for compliance rules
  - Automated testing for security configs

**Week 3-4: Documentation Foundation**
- [ ] Create initial policy documents
  - Information Security Policy (draft)
  - Access Control Policy (draft)
  - Incident Response Plan (draft)
  - Data Retention Policy (draft)
- [ ] Setup documentation repository
  - Create `docs/compliance/` structure
  - Version control for policies
  - Review workflow for policy changes

**Deliverables:**
- ✅ IaC templates với compliance-by-design
- ✅ Initial policy documentation
- ✅ Compliance validation automated

### Tháng 4: Operational Procedures

**Week 1-2: Runbooks**
- [ ] Create operational runbooks
  - Incident Response Runbook
  - GDPR Deletion Runbook
  - Backup & Recovery Runbook
  - Offboarding Runbook
- [ ] Test runbooks
  - Dry-run incident response
  - Test backup restoration
  - Validate deletion procedures

**Week 3-4: Training & Awareness**
- [ ] Security awareness program
  - Create training materials
  - Schedule initial training sessions
  - Document training completion
- [ ] Access review procedures
  - Quarterly access review process
  - Automated access reports
  - Offboarding checklist

**Deliverables:**
- ✅ Operational runbooks documented
- ✅ Initial security training completed
- ✅ Access review process established

**Phase 1 Checkpoint:**
- Infrastructure security baseline: ✅
- Security automation: ✅
- Documentation foundation: ✅
- Team trained: ✅

## Phase 2: Automation & SOC 2 Type 1 (Tháng 5-6)

### Mục Tiêu
Triển khai Vanta, hoàn thiện documentation, và đạt SOC 2 Type 1 cùng với v1 release.

### Tháng 5: Vanta Integration

**Week 1: Vanta Setup**
- [ ] Purchase Vanta subscription (Startups plan)
- [ ] Create Vanta account và invite team
- [ ] Complete initial Vanta questionnaire
- [ ] Configure company information

**Week 2: System Integration**
- [ ] Connect AWS to Vanta
  - Grant Vanta read-only IAM role
  - Verify CloudTrail integration
  - Test resource discovery
- [ ] Connect GitHub to Vanta
  - Install Vanta GitHub App
  - Verify repository scanning
  - Test security findings sync
- [ ] Connect HR systems (if applicable)
  - Employee roster sync
  - Onboarding/offboarding automation

**Week 3: Policy Upload**
- [ ] Upload policies to Vanta
  - Information Security Policy
  - Access Control Policy
  - Incident Response Plan
  - Data Retention Policy
  - Acceptable Use Policy
  - Vendor Management Policy
- [ ] Configure policy acknowledgment workflow
- [ ] Collect employee acknowledgments

**Week 4: Evidence Collection**
- [ ] Configure automated evidence collection
  - AWS security configurations
  - GitHub security settings
  - Access logs
  - Training records
- [ ] Review Vanta dashboard
  - Identify gaps
  - Prioritize remediation
  - Track progress

**Deliverables:**
- ✅ Vanta fully integrated
- ✅ Policies uploaded và acknowledged
- ✅ Automated evidence collection active

### Tháng 6: SOC 2 Type 1 Preparation

**Week 1: Gap Remediation**
- [ ] Address Vanta-identified gaps
  - Fix missing security controls
  - Complete documentation
  - Implement missing procedures
- [ ] Mock audit với Vanta
  - Review readiness score
  - Validate evidence completeness
  - Test auditor access

**Week 2: Auditor Selection**
- [ ] Select SOC 2 auditor
  - Get recommendations from Vanta
  - Compare quotes (target: $5K-8K)
  - Check auditor credentials
- [ ] Kickoff meeting với auditor
  - Define scope (Security TSC only)
  - Agree on timeline
  - Provide Vanta access

**Week 3-4: SOC 2 Type 1 Audit**
- [ ] Auditor fieldwork
  - Provide Vanta dashboard access
  - Answer auditor questions
  - Submit additional evidence as needed
- [ ] Review draft report
  - Address any findings
  - Provide clarifications
  - Approve final report

**Week 4: V1 Release + SOC 2 Type 1**
- [ ] **MILESTONE 1: V1 Release**
- [ ] **MILESTONE 1: SOC 2 Type 1 Report Issued**
- [ ] Publish SOC 2 report to customers
- [ ] Update marketing materials

**Deliverables:**
- ✅ SOC 2 Type 1 Report
- ✅ V1 released với compliance baseline
- ✅ Customer-facing compliance page

**Phase 2 Checkpoint:**
- Vanta operational: ✅
- SOC 2 Type 1 achieved: ✅
- V1 released: ✅

## Phase 3: ISO 27001 Certification (Tháng 7-8)

### Mục Tiêu
Đạt chứng chỉ ISO 27001 để bổ sung cho SOC 2 và mở rộng thị trường quốc tế.

### Tháng 7: ISO 27001 Preparation

**Week 1-2: Documentation Enhancement**
- [ ] Enhance policies cho ISO 27001
  - Risk Assessment & Treatment Plan
  - Statement of Applicability (SoA)
  - Asset Inventory
  - Business Continuity Plan
  - Supplier Management Policy
- [ ] Map existing controls to ISO 27001 Annex A
  - Identify coverage gaps
  - Document control implementations
  - Prepare evidence

**Week 3: Internal Audit**
- [ ] Conduct internal ISO 27001 audit
  - Review all 93 Annex A controls
  - Test control effectiveness
  - Document findings
- [ ] Remediate internal audit findings
  - Fix non-conformities
  - Update documentation
  - Re-test controls

**Week 4: Certification Body Selection**
- [ ] Select ISO 27001 certification body
  - Options: BSI, TÜV, SGS
  - Compare quotes (target: $3K-5K)
  - Check accreditation
- [ ] Schedule Stage 1 audit
  - Provide documentation
  - Agree on timeline
  - Prepare team

**Deliverables:**
- ✅ ISO 27001 documentation complete
- ✅ Internal audit passed
- ✅ Certification body selected

### Tháng 8: ISO 27001 Certification Audit

**Week 1-2: Stage 1 Audit (Document Review)**
- [ ] Certification body reviews documentation
  - ISMS manual
  - Policies và procedures
  - Risk assessment
  - Statement of Applicability
- [ ] Address Stage 1 findings
  - Update documentation
  - Clarify implementations
  - Prepare for Stage 2

**Week 3-4: Stage 2 Audit (Implementation Review)**
- [ ] On-site/remote audit
  - Auditor interviews team
  - Reviews system implementations
  - Tests control effectiveness
  - Validates evidence
- [ ] Address Stage 2 findings
  - Fix minor non-conformities
  - Provide additional evidence
  - Submit corrective actions

**Week 4: Certification Decision**
- [ ] **MILESTONE 2: ISO 27001 Certificate Issued**
- [ ] Receive certificate (valid 3 years)
- [ ] Schedule annual surveillance audits
- [ ] Update customer communications

**Deliverables:**
- ✅ ISO 27001 Certificate
- ✅ Surveillance audit schedule
- ✅ Updated compliance marketing

**Phase 3 Checkpoint:**
- ISO 27001 certified: ✅
- Dual certification (SOC 2 + ISO): ✅
- International credibility: ✅

## Phase 4: Maintenance & Monitoring (Tháng 9-11)

### Mục Tiêu
Duy trì tuân thủ, thu thập bằng chứng cho SOC 2 Type 2, và cải tiến liên tục.

### Tháng 9: Continuous Compliance

**Ongoing Activities:**
- [ ] Monitor Vanta dashboard daily
  - Review new findings
  - Track remediation progress
  - Maintain >95% compliance score
- [ ] Quarterly access reviews
  - Review AWS IAM permissions
  - Review GitHub access
  - Revoke unnecessary access
- [ ] Security awareness training
  - Monthly security tips
  - Phishing simulation tests
  - Track completion rates

**Week 1-2: GDPR Compliance Review**
- [ ] Assess GDPR requirements
  - Review data processing activities
  - Update privacy policy
  - Implement data subject rights APIs
- [ ] Create GDPR documentation
  - Records of Processing Activities (RoPA)
  - Data Processing Agreements (DPAs)
  - Data Breach Response Plan

**Week 3-4: Vendor Risk Management**
- [ ] Review third-party vendors
  - AWS, Cloudflare, GitHub
  - Collect SOC 2 reports
  - Document vendor assessments
- [ ] Update vendor management process
  - Vendor onboarding checklist
  - Annual vendor reviews
  - Contract security requirements

**Deliverables:**
- ✅ GDPR compliance baseline
- ✅ Vendor risk assessments complete
- ✅ Continuous monitoring active

### Tháng 10: Process Optimization

**Week 1-2: Automation Enhancement**
- [ ] Enhance security automation
  - Automated compliance checks
  - Self-healing configurations
  - Alerting improvements
- [ ] Optimize Vanta integration
  - Custom integrations
  - Automated evidence collection
  - Reporting automation

**Week 3-4: Documentation Updates**
- [ ] Review và update policies
  - Incorporate lessons learned
  - Address process gaps
  - Improve clarity
- [ ] Update runbooks
  - Add new procedures
  - Refine existing processes
  - Test effectiveness

**Deliverables:**
- ✅ Enhanced automation
- ✅ Updated documentation
- ✅ Improved processes

### Tháng 11: SOC 2 Type 2 Preparation

**Week 1-2: Evidence Review**
- [ ] Review 6-month evidence collection
  - Verify completeness
  - Identify gaps
  - Collect missing evidence
- [ ] Prepare for Type 2 audit
  - Review Vanta timeline
  - Validate control operation
  - Document exceptions

**Week 3-4: Pre-Audit Readiness**
- [ ] Mock Type 2 audit
  - Review evidence với auditor lens
  - Test control effectiveness
  - Identify potential findings
- [ ] Remediate pre-audit findings
  - Fix control gaps
  - Enhance documentation
  - Prepare explanations

**Deliverables:**
- ✅ 6-month evidence complete
- ✅ Type 2 readiness validated
- ✅ Pre-audit findings addressed

**Phase 4 Checkpoint:**
- Continuous compliance: ✅
- GDPR baseline: ✅
- Type 2 ready: ✅

## Phase 5: SOC 2 Type 2 (Tháng 12)

### Mục Tiêu
Đạt SOC 2 Type 2 để chứng minh vận hành bảo mật liên tục trong 6 tháng.

### Tháng 12: SOC 2 Type 2 Audit

**Week 1: Audit Kickoff**
- [ ] Engage SOC 2 auditor (same as Type 1)
- [ ] Define audit period (June - December)
- [ ] Provide Vanta access
- [ ] Schedule interviews

**Week 2-3: Audit Fieldwork**
- [ ] Auditor reviews 6-month evidence
  - Control operation testing
  - Sample selection và testing
  - Exception analysis
  - Interviews với team
- [ ] Respond to auditor requests
  - Provide additional evidence
  - Clarify implementations
  - Explain exceptions

**Week 4: Report Issuance**
- [ ] Review draft Type 2 report
  - Address findings
  - Provide management responses
  - Approve final report
- [ ] **MILESTONE 3: SOC 2 Type 2 Report Issued**
- [ ] Distribute to customers
- [ ] Update compliance page

**Deliverables:**
- ✅ SOC 2 Type 2 Report
- ✅ 6-month operational effectiveness proven
- ✅ Customer trust enhanced

**Phase 5 Checkpoint:**
- SOC 2 Type 2 achieved: ✅
- Full compliance stack: ✅
- Ready for growth: ✅

## Ongoing Maintenance (2027+)

### Annual Activities

**Q1 (Jan-Mar):**
- [ ] ISO 27001 surveillance audit
- [ ] Annual risk assessment
- [ ] Policy review và updates
- [ ] Security awareness training refresh

**Q2 (Apr-Jun):**
- [ ] SOC 2 Type 2 renewal audit
- [ ] Quarterly access reviews
- [ ] Vendor risk assessments
- [ ] Incident response drill

**Q3 (Jul-Sep):**
- [ ] GDPR compliance review
- [ ] Data retention cleanup
- [ ] Security tool evaluation
- [ ] Penetration testing

**Q4 (Oct-Dec):**
- [ ] Annual compliance planning
- [ ] Budget review
- [ ] Team training updates
- [ ] Year-end reporting

### Continuous Activities

**Daily:**
- Monitor Vanta dashboard
- Review security alerts
- Track compliance score

**Weekly:**
- Review new Dependabot/CodeQL findings
- Process access requests
- Update documentation

**Monthly:**
- Security awareness communications
- Compliance metrics reporting
- Management review meetings

**Quarterly:**
- Access reviews
- Policy acknowledgments
- Vendor assessments
- Compliance committee meetings

## Success Metrics

### Compliance Metrics
- **Vanta Score:** Maintain >95%
- **Audit Findings:** Zero critical findings
- **Training Completion:** 100% within 30 days
- **Access Reviews:** 100% quarterly completion

### Security Metrics
- **Vulnerability Remediation:** <7 days for High/Critical
- **Incident Response Time:** <1 hour detection, <4 hours containment
- **Backup Success Rate:** >99%
- **MFA Adoption:** 100%

### Business Metrics
- **Customer Trust:** SOC 2 + ISO 27001 in sales materials
- **Deal Velocity:** Reduced security questionnaire time
- **Market Access:** EU và US enterprise customers
- **Audit Efficiency:** <40 hours team effort per audit

## Risk Management

### Timeline Risks

| Risk | Impact | Mitigation |
|------|--------|------------|
| V1 delay affects SOC 2 timeline | High | Decouple compliance from feature work |
| Audit findings require rework | Medium | Mock audits, early remediation |
| Resource constraints | Medium | Prioritize automation, use Vanta |
| Vendor delays (auditors) | Low | Book auditors early, have backups |

### Technical Risks

| Risk | Impact | Mitigation |
|------|--------|------------|
| Vanta integration issues | Medium | Test thoroughly, have manual fallback |
| AWS misconfigurations | High | IaC validation, automated checks |
| Security tool false positives | Low | Tune tools, document exceptions |
| Evidence collection gaps | Medium | Regular Vanta dashboard reviews |

### Organizational Risks

| Risk | Impact | Mitigation |
|------|--------|------------|
| Team turnover | Medium | Document everything, cross-train |
| Compliance fatigue | Low | Automate heavily, celebrate wins |
| Budget overruns | Low | Track monthly, 20% buffer |
| Scope creep | Medium | Stick to Security TSC only for Type 1 |

## Budget Tracking

### Planned Spending

| Quarter | Item | Amount |
|---------|------|--------|
| Q1 2026 | Vanta subscription | $1,000 |
| Q2 2026 | SOC 2 Type 1 audit | $6,500 |
| Q2 2026 | Vanta (Q2) | $1,000 |
| Q3 2026 | ISO 27001 certification | $4,000 |
| Q3 2026 | Vanta (Q3) | $1,000 |
| Q4 2026 | SOC 2 Type 2 audit | $10,000 |
| Q4 2026 | Vanta (Q4) | $1,000 |
| **Total 2026** | | **$24,500** |

### Cost Optimization

- **Vanta Startups Plan:** $4K/year vs $12K+ for standard
- **Combined Auditor:** Use same firm for Type 1 & Type 2 (discount)
- **Automation:** Reduce manual effort by 80%
- **Self-Service:** Vanta dashboard reduces consultant needs

## Communication Plan

### Internal Communications

**Weekly:**
- Compliance progress updates in team standup
- Vanta dashboard review với Senior Engineer

**Monthly:**
- Compliance metrics to leadership
- Security awareness tips to all staff

**Quarterly:**
- Compliance committee meetings
- Board reporting (if applicable)

### External Communications

**Milestones:**
- SOC 2 Type 1: Press release, customer email
- ISO 27001: Website update, sales enablement
- SOC 2 Type 2: Customer success stories

**Ongoing:**
- Compliance page on website
- Security questionnaire responses
- RFP compliance sections

## Next Steps

### Immediate Actions (This Week)
1. [ ] Review và approve this roadmap
2. [ ] Assign compliance lead (Quách Thanh Bình)
3. [ ] Schedule Phase 1 kickoff meeting
4. [ ] Begin AWS IAM Identity Center setup

### Month 1 Priorities
1. [ ] Complete AWS security baseline
2. [ ] Enable GitHub security features
3. [ ] Create initial policy drafts
4. [ ] Schedule Vanta demo

### Success Criteria
- [ ] All milestones achieved on time
- [ ] Zero critical audit findings
- [ ] <$25K total spend in 2026
- [ ] >95% Vanta score maintained

---

**Document Owner:** Quách Thanh Bình  
**Last Updated:** 2026-05-27  
**Next Review:** 2026-06-27
