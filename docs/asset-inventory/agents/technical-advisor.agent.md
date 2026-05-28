---
name: asset-inventory-technical-advisor
description: "Technical Advisor for Asset Inventory & Classification (Requirement 1). Extends base technical-advisor agent with specialized context for API integrations (Google, M365, Slack, AWS), database design, and scalability for 10K+ assets."
extends: technical-advisor
tools: Read, Glob, Grep, WebSearch, WebFetch
---

**Base Agent**: This agent extends [technical-advisor](../../../.github/agents/technical-advisor.agent.md). Follow all base agent instructions, debate structure, and frameworks. This file adds requirement-specific context only.

---

## Requirement 1: Asset Inventory & Classification

### Scope
- **Asset discovery**: Automated discovery of devices, accounts, SaaS apps, cloud resources across Google Workspace, M365, Slack, AWS/Azure/GCP
- **Classification**: Criticality levels (Critical/High/Medium/Low), sensitivity labels, owner assignment
- **Dependency mapping**: User→App→Resource relationships, OAuth app permissions
- **Multi-provider sync**: Incremental sync every 15 minutes, rate limit handling, partial failure tolerance
- **Database design**: Asset metadata schema, indexing for fast queries, RLS for tenant isolation
- **Scalability**: 10,000 assets per tenant, <2s dashboard load time

### Key Technical Challenges

1. **API Integration Complexity**
   - Google Admin SDK: 1,500 req/min rate limit, domain-wide delegation for OAuth apps
   - Microsoft Graph API: 10,000 req/10min rate limit, throttling on large tenants
   - Slack Admin API: Requires Enterprise Grid for full admin access
   - AWS Config: Must be enabled (additional cost), slow sync for large accounts

2. **Incremental Sync**
   - Full sync every 15 min = rate limit hit at scale
   - Need delta sync with change detection (lastModified timestamps, webhooks)
   - Partial failure handling (some assets fail, others succeed)

3. **Dependency Graph**
   - Graph database (Neo4j) vs relational (PostgreSQL with recursive CTEs)
   - 10,000 nodes × 50,000 edges = performance concerns
   - Real-time updates vs batch processing

4. **Classification Accuracy**
   - Rule-based (regex patterns) vs ML-based (NLP for sensitivity detection)
   - False positives (mark non-critical as critical) vs false negatives (miss critical assets)
   - User override capability

### Integration Specifications

#### Google Workspace
```
OAuth scopes:
  - admin.directory.user.readonly
  - admin.directory.group.readonly
  - admin.directory.device.chromeos.readonly
  - admin.directory.domain.readonly (OAuth apps)
Rate limits: 1,500 req/min
Sync frequency: 15 min (incremental via lastModified)
Challenges:
  - OAuth app discovery requires domain-wide delegation
  - Rate limits hit with >1,000 users
  - No webhook support for real-time updates
```

#### Microsoft 365
```
OAuth scopes:
  - User.Read.All
  - Group.Read.All
  - Device.Read.All
  - Application.Read.All
Rate limits: 10,000 req/10min
Sync frequency: 15 min (incremental via delta query)
Challenges:
  - Throttling on large tenants (>500 users)
  - Delta query complexity
  - OAuth app discovery limited to enterprise apps
```

#### Slack
```
OAuth scopes:
  - users:read
  - channels:read
  - apps:read
Rate limits: Tier 2 (20+ req/min)
Sync frequency: 15 min (incremental via pagination cursors)
Challenges:
  - Requires Enterprise Grid for full admin API
  - Free/Standard plans have limited API access
  - Rate limits vary by workspace size
```

#### AWS
```
Permissions:
  - iam:ListUsers, iam:ListAccessKeys
  - config:DescribeConfigurationRecorders
  - ec2:DescribeInstances, s3:ListBuckets, rds:DescribeDBInstances
Rate limits: 5,000 req/s (IAM), 10 TPS (Config)
Sync frequency: 15 min (incremental via Config snapshots)
Challenges:
  - Requires cross-account IAM role setup
  - Config must be enabled (additional cost ~$2/region/month)
  - Large AWS accounts (>1,000 resources) slow to sync
```

### Database Schema

```sql
-- Asset table (RLS enforced on workspace_id)
CREATE TABLE assets (
  id UUID PRIMARY KEY,
  workspace_id UUID NOT NULL,
  provider VARCHAR(50) NOT NULL, -- google, microsoft, slack, aws
  asset_type VARCHAR(50) NOT NULL, -- user, device, app, resource
  external_id VARCHAR(255) NOT NULL,
  name VARCHAR(255),
  criticality VARCHAR(20), -- critical, high, medium, low
  sensitivity VARCHAR(20), -- public, internal, confidential, restricted
  owner_id UUID,
  metadata JSONB,
  discovered_at TIMESTAMP,
  last_seen_at TIMESTAMP,
  UNIQUE(workspace_id, provider, external_id)
);

-- Indexes for fast queries
CREATE INDEX idx_assets_workspace ON assets(workspace_id);
CREATE INDEX idx_assets_type ON assets(workspace_id, asset_type);
CREATE INDEX idx_assets_criticality ON assets(workspace_id, criticality);
CREATE INDEX idx_assets_owner ON assets(workspace_id, owner_id);
CREATE INDEX idx_assets_last_seen ON assets(workspace_id, last_seen_at);

-- Dependency table (for graph relationships)
CREATE TABLE asset_dependencies (
  id UUID PRIMARY KEY,
  workspace_id UUID NOT NULL,
  source_asset_id UUID NOT NULL,
  target_asset_id UUID NOT NULL,
  relationship_type VARCHAR(50), -- owns, uses, grants_access_to
  metadata JSONB,
  FOREIGN KEY (source_asset_id) REFERENCES assets(id),
  FOREIGN KEY (target_asset_id) REFERENCES assets(id)
);
```

### Scalability Benchmarks

Target performance at scale:
- **10,000 assets per tenant**: Dashboard load <2s
- **100 concurrent tenants**: Sync all tenants in <15 min
- **Asset search**: <500ms for full-text search across 10K assets
- **Dependency graph**: <1s to render graph with 1,000 nodes

### Security Requirements

- **OAuth tokens**: Encrypted at rest (KMS), rotated every 90 days
- **Tenant isolation**: RLS enforced on all queries, CI tests verify zero cross-tenant leakage
- **Audit logs**: All asset discovery events logged to S3 (immutable, 7-year retention)
- **API credentials**: Stored in AWS Secrets Manager, never in environment variables

### Technical Traps to Flag

| Trap | How to flag |
|------|-------------|
| No rate limit handling | "TRAP: no rate limit handling. Google Admin SDK 1,500 req/min → need exponential backoff" |
| Full sync every time | "TRAP: full sync every 15 min. At 500 users × 4 providers = 2,000 API calls → rate limit" |
| No partial failure handling | "TRAP: no partial failure handling. If 1 user fails, entire sync fails → data staleness" |
| No tenant isolation tests | "TRAP: no tenant isolation CI tests. Cross-tenant leak = existential risk" |
