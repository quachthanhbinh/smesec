# Phân Bổ Nguồn Lực và Mua Sắm

## Tổng Quan

Tài liệu này chi tiết hóa phân bổ nguồn lực nội bộ và kế hoạch mua sắm công cụ/dịch vụ bên ngoài để đạt được các chứng nhận tuân thủ.

## A. Vai Trò Trong Team Nội Bộ

### 1. Senior Backend Engineer (Lead Compliance)

**Người đảm nhiệm:** Quách Thanh Bình

**Trách nhiệm chính:**
- Thiết kế kiến trúc hạ tầng đạt chuẩn tuân thủ
- Thiết kế data flow và security controls
- Review code với focus vào security
- Quản lý hệ thống Vanta (setup, monitoring, remediation)
- Viết và maintain runbooks kỹ thuật
- Xử lý các lỗi bảo mật phát sinh trên server
- Làm việc trực tiếp với auditors
- Đào tạo team về security best practices

**Time Allocation:**
- Tháng 1-4 (Foundation): 30% time (~12 hours/week)
- Tháng 5-6 (SOC 2 Type 1): 50% time (~20 hours/week)
- Tháng 7-8 (ISO 27001): 40% time (~16 hours/week)
- Tháng 9-12 (Maintenance): 20% time (~8 hours/week)

**Deliverables:**
- Infrastructure as Code templates
- Security CI/CD pipelines
- Technical runbooks
- Vanta configuration và monitoring
- Audit evidence và documentation
- Security training materials

### 2. DevOps/SysAdmin

**Trách nhiệm chính:**
- Cấu hình AWS IAM Identity Center
- Thiết lập CloudTrail và logging infrastructure
- Quản lý phân quyền GitHub
- Đảm bảo IaC (Terraform) pass compliance rules
- Monitor và maintain security configurations
- Implement automated compliance checks
- Backup và disaster recovery procedures

**Time Allocation:**
- Tháng 1-4: 40% time (~16 hours/week)
- Tháng 5-8: 30% time (~12 hours/week)
- Tháng 9-12: 20% time (~8 hours/week)

**Deliverables:**
- AWS security baseline configuration
- CloudTrail và logging setup
- IAM policies và permission sets
- Automated compliance monitoring
- Backup và recovery procedures

**Note:** Nếu chưa có dedicated DevOps, Senior Backend Engineer sẽ đảm nhiệm thêm role này.

### 3. Product Owner / Founder

**Trách nhiệm chính:**
- Duyệt và ký ban hành các bộ tài liệu Chính sách Bảo mật
- Cung cấp ngân sách mua phần mềm và thuê kiểm toán
- Approve major compliance decisions
- Stakeholder communication về compliance status
- Business justification cho compliance investments
- Final approval cho audit reports

**Time Allocation:**
- Tháng 1-4: 10% time (~4 hours/week)
- Tháng 5-8: 20% time (~8 hours/week) - audit periods
- Tháng 9-12: 10% time (~4 hours/week)

**Deliverables:**
- Approved security policies
- Budget allocation decisions
- Stakeholder communications
- Business case documentation

### 4. Tất Cả Nhân Sự

**Trách nhiệm chung:**
- Bật MFA cho mọi tài khoản làm việc (AWS, GitHub, Email, Slack)
- Ký cam kết bảo mật thông tin (NDA)
- Tham gia khóa training bảo mật cơ bản (via Vanta)
- Acknowledge và tuân thủ security policies
- Report security incidents ngay lập tức
- Follow secure coding practices
- Participate in security drills

**Time Allocation:**
- Initial training: 2-4 hours (one-time)
- Monthly security awareness: 30 minutes/month
- Quarterly policy review: 1 hour/quarter
- Annual refresher training: 2 hours/year

**Deliverables:**
- Completed security training
- Signed policy acknowledgments
- MFA enabled on all accounts
- Incident reports (when applicable)

## B. Đơn Vị Cần Thuê và Phần Mềm Cần Mua

### 1. Phần Mềm Compliance Automation

**Sản phẩm:** Vanta (Gói Startups)

**Mục đích:**
- Kết nối API quét AWS, GitHub, Cloudflare
- Quản lý chính sách và policy acknowledgments
- Tự động thu thập bằng chứng 24/7
- Dashboard cho auditors
- Continuous compliance monitoring
- Gap identification và remediation tracking

**Tính năng chính:**
- AWS integration (CloudTrail, IAM, S3, RDS, etc.)
- GitHub integration (repos, security, access)
- HR integration (employee lifecycle)
- Policy management và distribution
- Evidence collection và storage
- Audit readiness dashboard
- Compliance score tracking
- Automated questionnaire responses

**Ngân sách:**
- **Gói Startups:** $4,000 - $6,000/năm
- **Payment:** Annual upfront (discount available)
- **Setup fee:** Usually waived for startups

**Alternatives Considered:**
- Drata: $8,000 - $12,000/năm (more expensive)
- Secureframe: $6,000 - $10,000/năm (fewer integrations)
- Manual compliance: $0 but 10x time investment

**ROI:**
- Saves ~80% of manual compliance effort
- Reduces audit preparation time by 60%
- Continuous monitoring prevents gaps
- Estimated time savings: 20-30 hours/month

**Procurement Process:**
1. Sign up at vanta.com/startups
2. Schedule demo và onboarding call
3. Provide company information
4. Sign contract (annual)
5. Setup integrations (Week 1)
6. Configure policies (Week 2)
7. Start evidence collection (Week 3)

### 2. Quét Bảo Mật Code

**Sản phẩm:** GitHub Dependabot + GitHub Advanced Security (CodeQL)

**Mục đích:**
- Kiểm tra lỗ hổng thư viện bên thứ 3 (SCA)
- Quét mã nguồn tự viết (SAST)
- Secret scanning
- Automated security updates
- PR blocking cho High/Critical vulnerabilities

**Tính năng:**
- Dependabot alerts và updates
- CodeQL analysis (multi-language)
- Secret scanning với push protection
- Security advisories
- Dependency graph
- Integration với GitHub Actions

**Ngân sách:**
- **Miễn phí** cho public repositories
- **GitHub Team:** $4/user/month (includes Advanced Security)
- **GitHub Enterprise:** $21/user/month

**Estimated Cost:**
- Team of 5: $20/month ($240/year)
- Team of 10: $40/month ($480/year)

**Note:** Nếu đã có GitHub Team/Enterprise subscription, Advanced Security đã included.

### 3. Công Ty Kiểm Toán (CPA Firm) - SOC 2

**Mục đích:**
- Kiểm tra dashboard Vanta
- Phỏng vấn team members
- Test control effectiveness
- Cấp báo cáo SOC 2 Type 1 và Type 2

**Scope:**
- **Type 1:** Point-in-time audit (Security TSC only)
- **Type 2:** 6-12 month operational effectiveness audit

**Ngân sách:**
- **SOC 2 Type 1:** $5,000 - $8,000
- **SOC 2 Type 2:** $8,000 - $15,000
- **Combined discount:** ~10-15% if same auditor

**Timeline:**
- Type 1: 2-3 months from kickoff
- Type 2: 6-12 months observation period + 1-2 months audit

**Recommended Auditors:**
- Johanson Group (Vanta partner)
- A-LIGN (Vanta partner)
- Sensiba San Filippo
- Armanino
- Moss Adams

**Selection Criteria:**
- Vanta integration experience
- Startup-friendly pricing
- Fast turnaround time
- Good communication
- References from similar companies

**Procurement Process:**
1. Get recommendations from Vanta
2. Request quotes from 3 auditors
3. Compare pricing và timeline
4. Check references
5. Sign engagement letter
6. Kickoff meeting
7. Provide Vanta access
8. Fieldwork (2-4 weeks)
9. Draft report review
10. Final report issuance

### 4. Tổ Chức Chứng Nhận ISO 27001

**Đơn vị đề xuất:** BSI, TÜV, hoặc SGS

**Mục đích:**
- Kiểm tra hệ thống quản lý tài liệu
- Verify control implementations
- Test effectiveness
- Cấp chứng chỉ ISO 27001 (valid 3 years)
- Annual surveillance audits

**Scope:**
- Stage 1: Document review (1-2 weeks)
- Stage 2: Implementation audit (1-2 weeks)
- Surveillance: Annual audits (1 day each)

**Ngân sách:**
- **Initial Certification:** $3,000 - $5,000
- **Annual Surveillance:** $1,500 - $2,500/year
- **Recertification (Year 3):** $2,500 - $4,000

**Total 3-year cost:** ~$10,000 - $15,000

**Certification Bodies:**

| Organization | Pros | Cons | Est. Cost |
|--------------|------|------|-----------|
| **BSI** | Most recognized globally | More expensive | $4,000 - $5,000 |
| **TÜV** | Strong in EU/Asia | Less known in US | $3,500 - $4,500 |
| **SGS** | Good balance | Longer timeline | $3,000 - $4,000 |

**Procurement Process:**
1. Request quotes from 3 certification bodies
2. Compare pricing, timeline, reputation
3. Check accreditation (ANAB, UKAS, etc.)
4. Review sample certificates
5. Sign certification agreement
6. Schedule Stage 1 audit
7. Address Stage 1 findings
8. Schedule Stage 2 audit
9. Address Stage 2 findings
10. Receive certificate

### 5. Additional Tools (Optional)

**Security Scanning Tools:**

| Tool | Purpose | Cost | Priority |
|------|---------|------|----------|
| **Snyk** | Advanced dependency scanning | $500-1K/year | Medium |
| **SonarQube** | Code quality + security | $150-500/year | Low |
| **Semgrep** | SAST rules | Free - $500/year | Medium |

**Infrastructure Tools:**

| Tool | Purpose | Cost | Priority |
|------|---------|------|----------|
| **Terraform Cloud** | IaC state management | Free - $20/month | Medium |
| **AWS Config** | Compliance monitoring | ~$50-100/month | High |
| **CloudWatch** | Logging và monitoring | ~$50-150/month | High |

**Recommendation:** Start with free/included tools (GitHub, AWS native), add paid tools only if gaps identified.

## C. Tổng Ngân Sách

### Year 1 (2026) - Initial Certification

| Hạng mục | Q1 | Q2 | Q3 | Q4 | Total |
|----------|----|----|----|----|-------|
| **Vanta Subscription** | $1,000 | $1,000 | $1,000 | $1,000 | $4,000 |
| **GitHub Advanced Security** | $60 | $60 | $60 | $60 | $240 |
| **SOC 2 Type 1 Audit** | - | $6,500 | - | - | $6,500 |
| **ISO 27001 Certification** | - | - | $4,000 | - | $4,000 |
| **SOC 2 Type 2 Audit** | - | - | - | $10,000 | $10,000 |
| **Contingency (20%)** | $212 | $1,512 | $1,012 | $2,212 | $4,948 |
| **Quarterly Total** | $1,272 | $9,072 | $6,072 | $13,272 | **$29,688** |

### Year 2+ (2027+) - Maintenance

| Hạng mục | Annual Cost |
|----------|-------------|
| **Vanta Subscription** | $4,000 - $6,000 |
| **GitHub Advanced Security** | $240 - $480 |
| **SOC 2 Type 2 Renewal** | $10,000 - $12,000 |
| **ISO 27001 Surveillance** | $1,500 - $2,500 |
| **Contingency (15%)** | $2,361 - $3,147 |
| **Annual Total** | **$18,101 - $24,127** |

### Cost Optimization Strategies

**Year 1 Savings:**
- Vanta Startups discount: Save $2K-4K vs standard pricing
- Combined SOC 2 auditor: Save 10-15% on Type 2
- GitHub Team vs Enterprise: Save $17/user/month
- Self-service implementation: Save $10K-20K consultant fees

**Ongoing Savings:**
- Annual Vanta payment: 10-15% discount
- Multi-year ISO contract: 5-10% discount
- Same auditor loyalty: 5-10% discount on renewals
- Automation: Reduce manual effort by 80%

**Total Estimated Savings:** $15,000 - $25,000 in Year 1

## D. ROI Analysis

### Compliance Investment

**Total Year 1 Investment:** ~$30,000
- Direct costs: $24,740
- Internal time (estimated): ~$5,000 equivalent

**Total Year 2+ Investment:** ~$20,000/year

### Business Value

**Revenue Impact:**
- **Enterprise deals enabled:** $50K - $200K ARR per deal
- **Deal velocity improvement:** 30-50% faster security reviews
- **Win rate improvement:** 10-20% higher for enterprise deals
- **Market expansion:** EU and US enterprise markets accessible

**Cost Avoidance:**
- **Security incidents prevented:** $50K - $500K potential cost
- **Compliance violations avoided:** $10K - $100K potential fines
- **Manual effort saved:** 20-30 hours/month (~$10K-15K/year)

**Estimated ROI:**
- **Break-even:** 1-2 enterprise deals
- **Year 1 ROI:** 200-500% (if 2-3 enterprise deals closed)
- **Year 2+ ROI:** 400-1000% (maintenance costs lower)

### Intangible Benefits

- **Brand credibility:** Trust signal for enterprise customers
- **Competitive advantage:** Differentiation from competitors
- **Team confidence:** Better security practices
- **Investor appeal:** Due diligence readiness
- **Partnership opportunities:** Required by many partners

## E. Resource Planning Timeline

### Q1 2026 (Jan-Mar): Foundation

**Team Focus:**
- Senior Backend Engineer: 30% time
- DevOps: 40% time
- All staff: Initial training (4 hours)

**Spending:**
- Vanta: $1,000
- GitHub: $60
- Contingency: $212
- **Total: $1,272**

**Deliverables:**
- AWS security baseline
- GitHub security enabled
- Initial policies drafted
- Vanta account setup

### Q2 2026 (Apr-Jun): SOC 2 Type 1

**Team Focus:**
- Senior Backend Engineer: 50% time (peak effort)
- DevOps: 30% time
- Product Owner: 20% time
- All staff: Monthly training (30 min)

**Spending:**
- Vanta: $1,000
- GitHub: $60
- SOC 2 Type 1 Audit: $6,500
- Contingency: $1,512
- **Total: $9,072**

**Deliverables:**
- Vanta fully integrated
- SOC 2 Type 1 Report
- V1 released với compliance

### Q3 2026 (Jul-Sep): ISO 27001

**Team Focus:**
- Senior Backend Engineer: 40% time
- DevOps: 30% time
- Product Owner: 20% time
- All staff: Quarterly review (1 hour)

**Spending:**
- Vanta: $1,000
- GitHub: $60
- ISO 27001 Certification: $4,000
- Contingency: $1,012
- **Total: $6,072**

**Deliverables:**
- ISO 27001 Certificate
- Enhanced documentation
- GDPR compliance baseline

### Q4 2026 (Oct-Dec): SOC 2 Type 2

**Team Focus:**
- Senior Backend Engineer: 30% time
- DevOps: 20% time
- Product Owner: 10% time
- All staff: Annual refresher (2 hours)

**Spending:**
- Vanta: $1,000
- GitHub: $60
- SOC 2 Type 2 Audit: $10,000
- Contingency: $2,212
- **Total: $13,272**

**Deliverables:**
- SOC 2 Type 2 Report
- Full compliance stack
- Continuous monitoring established

## F. Vendor Management

### Vendor Risk Assessment

**Critical Vendors (Require SOC 2):**
- AWS (Infrastructure)
- Cloudflare (Storage)
- GitHub (Source control)
- Vanta (Compliance automation)

**Standard Vendors (Security review):**
- Email provider
- Communication tools (Slack, etc.)
- Analytics tools
- Payment processors

**Vendor Onboarding Checklist:**
- [ ] Security questionnaire completed
- [ ] SOC 2 report reviewed (if applicable)
- [ ] Data Processing Agreement signed (for GDPR)
- [ ] Access controls documented
- [ ] Incident notification process established
- [ ] Annual review scheduled

### Vendor Contracts

**Key Terms to Negotiate:**
- **Data ownership:** Customer owns all data
- **Data deletion:** 30-day deletion upon termination
- **Security incidents:** 24-hour notification requirement
- **Audit rights:** Right to audit vendor security
- **Liability:** Clear liability terms for breaches
- **Termination:** 30-60 day termination clause

**Contract Review Process:**
1. Legal review (if available)
2. Security review by Senior Engineer
3. Compliance review (GDPR, ISO 27001)
4. Approval by Product Owner
5. Signature và storage

## G. Success Metrics

### Compliance Metrics

| Metric | Target | Measurement |
|--------|--------|-------------|
| **Vanta Compliance Score** | >95% | Weekly dashboard review |
| **Audit Findings** | 0 critical | Post-audit report |
| **Training Completion** | 100% within 30 days | Vanta tracking |
| **Policy Acknowledgment** | 100% within 14 days | Vanta tracking |
| **Access Review Completion** | 100% quarterly | Manual tracking |
| **Incident Response Time** | <1 hour detection | Incident logs |

### Efficiency Metrics

| Metric | Target | Measurement |
|--------|--------|-------------|
| **Time to Audit Ready** | <6 months | Project timeline |
| **Manual Effort Reduction** | >80% | Time tracking |
| **Audit Preparation Time** | <40 hours | Time tracking |
| **Security Questionnaire Time** | <2 hours | Sales tracking |
| **Evidence Collection** | 100% automated | Vanta dashboard |

### Business Metrics

| Metric | Target | Measurement |
|--------|--------|-------------|
| **Enterprise Deals Enabled** | 2-3 in Year 1 | Sales pipeline |
| **Deal Velocity** | 30% improvement | Sales cycle time |
| **Win Rate** | 10% improvement | Sales conversion |
| **Customer Trust Score** | >4.5/5 | Customer surveys |
| **Security Incidents** | 0 major | Incident logs |

## H. Risk Mitigation

### Budget Risks

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| Audit cost overrun | Medium | $2K-5K | 20% contingency buffer |
| Additional tools needed | Low | $1K-3K | Evaluate free alternatives first |
| Consultant fees | Low | $5K-10K | Self-service via Vanta |
| Remediation costs | Medium | $2K-5K | Early gap identification |

### Resource Risks

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| Team overload | High | Timeline delay | Prioritize automation |
| Key person dependency | Medium | Project delay | Cross-training |
| Competing priorities | High | Timeline slip | Executive sponsorship |
| Skill gaps | Medium | Quality issues | Vanta support + training |

### Vendor Risks

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| Vanta service issues | Low | Workflow disruption | Manual fallback procedures |
| Auditor delays | Medium | Timeline slip | Book early, have backup |
| AWS outages | Low | Evidence gaps | Multi-region logging |
| GitHub unavailability | Low | Development delay | Local git backups |

## I. Procurement Checklist

### Immediate (This Month)
- [ ] Approve compliance budget ($30K Year 1)
- [ ] Sign up for Vanta Startups program
- [ ] Enable GitHub Advanced Security
- [ ] Assign compliance lead (Quách Thanh Bình)
- [ ] Schedule kickoff meeting

### Month 1-2
- [ ] Complete Vanta onboarding
- [ ] Connect AWS, GitHub, Cloudflare to Vanta
- [ ] Upload initial policies to Vanta
- [ ] Request SOC 2 auditor quotes
- [ ] Request ISO 27001 certification quotes

### Month 3-4
- [ ] Select SOC 2 auditor
- [ ] Select ISO 27001 certification body
- [ ] Sign audit engagement letters
- [ ] Complete vendor risk assessments
- [ ] Finalize all procurement

### Ongoing
- [ ] Monthly Vanta dashboard review
- [ ] Quarterly vendor reviews
- [ ] Annual contract renewals
- [ ] Continuous cost optimization

## J. Contact Information

### Internal Contacts
- **Compliance Lead:** Quách Thanh Bình - [email] - [phone]
- **Product Owner:** [Name] - [email] - [phone]
- **DevOps Lead:** [Name] - [email] - [phone]

### Vendor Contacts
- **Vanta Support:** support@vanta.com
- **GitHub Support:** support@github.com
- **AWS Support:** [Account manager]
- **SOC 2 Auditor:** [TBD after selection]
- **ISO 27001 Auditor:** [TBD after selection]

### Emergency Contacts
- **Security Incidents:** [On-call rotation]
- **Compliance Issues:** Quách Thanh Bình
- **Legal:** [Legal counsel]
- **Executive Sponsor:** [Product Owner/Founder]

---

**Document Owner:** Quách Thanh Bình  
**Last Updated:** 2026-05-27  
**Next Review:** 2026-08-27  
**Budget Approval:** [Pending/Approved]
