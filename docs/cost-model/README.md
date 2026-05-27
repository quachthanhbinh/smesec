# Tài Liệu Chiến Lược: Cost Model

## Tổng Quan

Bộ tài liệu này mô tả chiến lược xây dựng cost model với tiered, pay-as-you-grow pricing aligned với SME constraints cho hệ thống SMESec.

**Thông tin dự án:**
- **Người phụ trách:** Quách Thanh Bình - Senior Backend Engineer
- **Giai đoạn hệ thống:** Build from scratch (Dự kiến v1 sau 6 tháng)
- **Ngày cập nhật:** Tháng 5/2026
- **Thời gian triển khai:** 6 tháng (song song với v1)

## Mục Tiêu

Xây dựng cost model bền vững để:
- Phù hợp với ngân sách hạn chế của SMEs (10-500 employees)
- Tăng trưởng theo quy mô doanh nghiệp (pay-as-you-grow)
- Cạnh tranh với enterprise security solutions
- Đảm bảo profitability cho SMESec platform
- Transparent pricing không có hidden costs

## Phạm Vi Cost Model

### 1. Pricing Tiers
- **Starter Tier**: 10-50 employees
- **Growth Tier**: 51-150 employees
- **Business Tier**: 151-300 employees
- **Enterprise Tier**: 301-500 employees

### 2. Cost Components
- **Base Platform Fee**: Core security features
- **Per-User Pricing**: Scalable với số lượng nhân viên
- **Feature Add-ons**: Optional advanced features
- **Usage-Based Pricing**: API calls, storage, bandwidth

### 3. Cost Drivers
- **Infrastructure Costs**: AWS, Cloudflare, databases
- **Third-Party Services**: Vanta, security tools, APIs
- **Development & Maintenance**: Team costs, support
- **Compliance & Certification**: Audit costs, certifications

### 4. Revenue Streams
- **Subscription Revenue**: Monthly/annual recurring
- **Professional Services**: Implementation, training, consulting
- **Partner Revenue**: Referral fees, integrations
- **Premium Support**: Dedicated support tiers

## Cấu Trúc Tài Liệu

### [01. Architecture Decision Record (ADR)](01-adr.md)
Ghi nhận các quyết định về:
- Pricing model selection (tiered vs usage-based vs hybrid)
- Billing system architecture
- Payment gateway integration
- Revenue recognition strategy

### [02. Pricing Strategy Framework](02-pricing-strategy.md)
Chi tiết về:
- Market analysis và competitive pricing
- Value-based pricing methodology
- Customer segmentation
- Pricing psychology và anchoring

### [03. Tiered Pricing Structure](03-tiered-pricing.md)
Cấu trúc giá chi tiết:
- Feature matrix per tier
- Per-user pricing calculations
- Volume discounts
- Annual vs monthly pricing

### [04. Cost Analysis & Unit Economics](04-cost-analysis.md)
Phân tích chi phí:
- Infrastructure cost per customer
- Customer Acquisition Cost (CAC)
- Lifetime Value (LTV)
- Gross margin analysis
- Break-even analysis

### [05. Lộ Trình Triển Khai](05-roadmap.md)
Timeline chi tiết 6 tháng:
- **Tháng 1-2**: Cost modeling + pricing research
- **Tháng 3-4**: Billing system implementation
- **Tháng 5-6**: Payment integration + pricing optimization

### [06. Technical Implementation Guide](06-technical-guide.md)
Hướng dẫn kỹ thuật:
- Billing system architecture
- Subscription management
- Usage metering và tracking
- Payment gateway integration

### [07. Phân Bổ Nguồn Lực](07-resources.md)
Kế hoạch nguồn lực:
- Team roles (Product Manager, Finance, Engineering)
- Tools và services cần thiết
- Ngân sách ước tính

## Công Nghệ & Công Cụ

### Billing & Subscription Management
- **Billing Platform**: Stripe Billing, Chargebee, Recurly
- **Subscription Logic**: Custom subscription engine
- **Invoicing**: Stripe Invoicing, QuickBooks integration
- **Tax Calculation**: Stripe Tax, Avalara, TaxJar

### Usage Metering & Analytics
- **Metering System**: Custom usage tracking
- **Analytics**: Mixpanel, Amplitude, custom dashboards
- **Cost Allocation**: AWS Cost Explorer, Cloudflare Analytics
- **Reporting**: Metabase, Looker, custom reports

### Payment Processing
- **Payment Gateway**: Stripe, PayPal, Braintree
- **Payment Methods**: Credit cards, ACH, wire transfer
- **Multi-Currency**: Stripe multi-currency support
- **Fraud Prevention**: Stripe Radar, custom rules

### Financial Management
- **Accounting**: QuickBooks, Xero, NetSuite
- **Revenue Recognition**: Stripe Revenue Recognition, custom logic
- **Financial Reporting**: Custom dashboards, accounting software
- **Forecasting**: Spreadsheets, financial modeling tools

## Nguyên Tắc Chính

### 1. Transparent Pricing
Không có hidden fees, tất cả costs được công khai rõ ràng.

### 2. Value-Based Pricing
Giá dựa trên value delivered, không chỉ cost-plus pricing.

### 3. Predictable Costs
SMEs cần predictable monthly costs để budget planning.

### 4. Fair Growth Pricing
Pricing scales fairly với company growth, không có sudden jumps.

## Proposed Pricing Structure

### Starter Tier (10-50 employees)
**Base Price**: $299/month ($249/month annual)
- Core security features
- Asset inventory & classification
- Basic AI threat protection
- Access governance (SSO + MFA)
- 5 incident playbooks
- Email support

**Per-User**: $5/user/month above 25 users

### Growth Tier (51-150 employees)
**Base Price**: $799/month ($699/month annual)
- All Starter features
- Advanced AI threat detection
- Shadow IT detection
- Automated offboarding
- 15 incident playbooks
- Priority email + chat support
- Quarterly compliance reports

**Per-User**: $4/user/month above 75 users

### Business Tier (151-300 employees)
**Base Price**: $1,999/month ($1,799/month annual)
- All Growth features
- Custom playbooks
- Advanced analytics & reporting
- API access
- Dedicated account manager
- Phone support
- Monthly compliance reports
- SOC 2 Type 2 support

**Per-User**: $3/user/month above 200 users

### Enterprise Tier (301-500 employees)
**Base Price**: Custom pricing (starting $4,999/month)
- All Business features
- Custom integrations
- On-premise deployment option
- 24/7 support
- Custom SLAs
- Professional services included
- Dedicated security consultant

**Per-User**: Custom pricing

## Add-On Features (All Tiers)

| Feature | Price |
|---------|-------|
| Additional Playbooks (5-pack) | $99/month |
| Advanced Deepfake Detection | $199/month |
| Premium Threat Intelligence | $299/month |
| Custom Integration Development | $2,000 one-time |
| Professional Services (per hour) | $200/hour |
| Dedicated Training Session | $500/session |

## Cost Analysis

### Infrastructure Costs (per customer)

| Component | Starter | Growth | Business | Enterprise |
|-----------|---------|--------|----------|------------|
| AWS Infrastructure | $50 | $150 | $400 | $1,000+ |
| Cloudflare R2 Storage | $10 | $30 | $80 | $200+ |
| Third-Party APIs | $30 | $80 | $200 | $500+ |
| Compliance Tools (Vanta) | $20 | $40 | $80 | $150+ |
| **Total COGS** | **$110** | **$300** | **$760** | **$1,850+** |

### Unit Economics

| Metric | Starter | Growth | Business | Enterprise |
|--------|---------|--------|----------|------------|
| Average Monthly Revenue | $350 | $1,000 | $2,500 | $6,000+ |
| COGS | $110 | $300 | $760 | $1,850 |
| Gross Margin | $240 (69%) | $700 (70%) | $1,740 (70%) | $4,150+ (69%) |
| Target CAC | $1,050 | $3,000 | $7,500 | $18,000 |
| Payback Period | 3 months | 3 months | 3 months | 3 months |
| LTV (3 years) | $12,600 | $36,000 | $90,000 | $216,000+ |
| LTV:CAC Ratio | 12:1 | 12:1 | 12:1 | 12:1 |

## Metrics & KPIs

| Metric | Target | Measurement |
|--------|--------|-------------|
| Monthly Recurring Revenue (MRR) | Growth 15%/month | Total subscription revenue |
| Customer Acquisition Cost (CAC) | <$3,000 avg | Sales + marketing costs / new customers |
| Customer Lifetime Value (LTV) | >$36,000 avg | Avg revenue per customer * avg lifetime |
| LTV:CAC Ratio | >10:1 | LTV / CAC |
| Gross Margin | >65% | (Revenue - COGS) / Revenue |
| Net Revenue Retention | >110% | Revenue from existing customers YoY |
| Churn Rate | <5%/year | % customers canceling annually |
| Average Revenue Per User (ARPU) | >$15/user/month | Total revenue / total users |

## Ngân Sách Ước Tính (Platform Development)

| Hạng mục | Chi phí |
|----------|---------|
| Billing System Development | $15,000 - $25,000 |
| Stripe Integration | $5,000 - $8,000 |
| Usage Metering System | $10,000 - $15,000 |
| Admin Dashboard | $8,000 - $12,000 |
| Financial Reporting | $5,000 - $8,000 |
| Stripe Fees (2.9% + $0.30) | Variable |
| Accounting Software | $500 - $1,000/year |

**Tổng ước tính:** ~$43,500 - $69,000 one-time + ongoing fees

## Milestone Chính

- **Milestone 1 (Tháng 2)**: Pricing model finalized + cost analysis complete
- **Milestone 2 (Tháng 4)**: Billing system operational + Stripe integrated
- **Milestone 3 (Tháng 6)**: Full pricing tiers live + usage metering working

## Competitive Analysis

### Competitor Pricing (SME Security Platforms)

| Competitor | Target | Starting Price | Notes |
|------------|--------|----------------|-------|
| Huntress | SMBs | $5-7/endpoint/month | Endpoint focus |
| Datto | MSPs | $3-5/endpoint/month | MSP channel |
| Cisco Umbrella | SMBs | $3-5/user/month | DNS security |
| Okta | All sizes | $2-15/user/month | Identity focus |
| Vanta | Startups | $3,000-5,000/year | Compliance focus |

**SMESec Positioning**: Comprehensive platform at competitive pricing with AI-specific protections.

## Pricing Optimization Strategies

### 1. Annual Discount
Offer 15-20% discount for annual prepayment to improve cash flow.

### 2. Volume Discounts
Automatic discounts at tier boundaries to encourage growth.

### 3. Bundling
Bundle popular add-ons into higher tiers for better value perception.

### 4. Freemium/Trial
14-day free trial, no credit card required, to reduce friction.

### 5. Referral Program
$100 credit for referrals that convert to paid customers.

## Liên Hệ & Hỗ Trợ

**Người phụ trách:** Quách Thanh Bình  
**Email:** [Thêm email]  
**Slack:** [Thêm channel]

## Tài Liệu Tham Khảo

- [SaaS Pricing Strategy Guide - ProfitWell](https://www.profitwell.com/recur/all/saas-pricing-strategy)
- [Stripe Billing Documentation](https://stripe.com/docs/billing)
- [SaaS Metrics 2.0 - David Skok](https://www.forentrepreneurs.com/saas-metrics-2/)
- [Price Intelligently - Pricing Strategy](https://www.priceintelligently.com/)
- [OpenView SaaS Benchmarks](https://openviewpartners.com/saas-benchmarks/)

---

**Lưu ý:** Tài liệu này là living document và sẽ được cập nhật thường xuyên dựa trên market feedback và financial performance.
