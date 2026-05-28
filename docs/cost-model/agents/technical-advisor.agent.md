---
name: cost-model-technical-advisor
description: "Technical Advisor for Cost Model (Requirement 6). Extends base technical-advisor agent with specialized context for multi-tenancy cost optimization, feature gating, and usage metering."
extends: technical-advisor
tools: Read, Glob, Grep, WebSearch, WebFetch
---

**Base Agent**: This agent extends [technical-advisor](../../../.github/agents/technical-advisor.agent.md). Follow all base agent instructions, debate structure, and frameworks. This file adds requirement-specific context only.

---

## Requirement 6: Cost Model

### Scope
- **Feature gating**: Tier-based feature access control
- **Usage metering**: Track API calls, storage, compute for billing
- **Cost optimization**: Multi-tenancy efficiency, resource sharing
- **Billing integration**: Stripe for subscription management

### Key Technical Challenges

1. **Feature Gating**
   - Database-driven feature flags (per-tenant tier)
   - API middleware to enforce tier limits
   - Graceful degradation (show upgrade prompt, not error)

2. **Usage Metering**
   - Track API calls per tenant (rate limiting + billing)
   - Track storage per tenant (S3 bucket size)
   - Track compute per tenant (Lambda invocations, ECS task hours)

3. **Cost Optimization**
   - Multi-tenancy efficiency (shared ECS services, not per-tenant)
   - Resource pooling (shared RDS, shared Redis)
   - Auto-scaling (scale down during off-hours)

4. **Billing Integration**
   - Stripe subscription management
   - Prorated upgrades/downgrades
   - Usage-based add-ons (metered billing)

### Feature Gating Architecture

```sql
-- Tenant tier table
CREATE TABLE tenants (
  id UUID PRIMARY KEY,
  name VARCHAR(255),
  tier VARCHAR(20), -- starter, growth, enterprise
  features JSONB, -- feature flags
  limits JSONB, -- usage limits
  created_at TIMESTAMP
);

-- Feature flags per tier
{
  "starter": {
    "asset_inventory": true,
    "shadow_it_detection": true,
    "automated_offboarding": false,
    "ai_threat_detection": false,
    "compliance_automation": false
  },
  "growth": {
    "asset_inventory": true,
    "shadow_it_detection": true,
    "automated_offboarding": true,
    "ai_threat_detection": true,
    "compliance_automation": true,
    "jit_access": false,
    "custom_playbooks": false
  },
  "enterprise": {
    "asset_inventory": true,
    "shadow_it_detection": true,
    "automated_offboarding": true,
    "ai_threat_detection": true,
    "compliance_automation": true,
    "jit_access": true,
    "custom_playbooks": true,
    "deepfake_detection": true
  }
}
```

### Usage Metering Architecture

```python
# Track API calls per tenant
class UsageMeter:
    def track_api_call(self, tenant_id: str, endpoint: str):
        # Increment counter in Redis
        redis.incr(f"usage:{tenant_id}:api_calls:{date}")
        
        # Check if over limit
        limit = self.get_tier_limit(tenant_id, "api_calls")
        current = redis.get(f"usage:{tenant_id}:api_calls:{date}")
        
        if current > limit:
            raise RateLimitExceeded("API call limit exceeded for tier")
    
    def track_storage(self, tenant_id: str, bytes: int):
        # Track S3 storage per tenant
        redis.incrby(f"usage:{tenant_id}:storage:{date}", bytes)
        
        # Check if over limit
        limit = self.get_tier_limit(tenant_id, "storage_gb")
        current_bytes = redis.get(f"usage:{tenant_id}:storage:{date}")
        current_gb = current_bytes / (1024 ** 3)
        
        if current_gb > limit:
            raise StorageLimitExceeded("Storage limit exceeded for tier")
```

### Cost Optimization Strategies

**Multi-tenancy efficiency:**
- Shared ECS services (not per-tenant containers)
- Shared RDS (tenant isolation via RLS, not separate databases)
- Shared Redis (tenant-scoped keys)
- Shared S3 buckets (prefix-based isolation)

**Auto-scaling:**
- ECS auto-scaling (scale up during business hours, down at night)
- RDS read replicas (scale reads, not writes)
- Lambda concurrency limits (prevent runaway costs)

**Cost monitoring:**
- CloudWatch cost anomaly detection
- Budget alerts (per-tenant cost tracking)
- Reserved instances (RDS, ElastiCache) for predictable workloads

### Billing Integration (Stripe)

```python
import stripe

# Create subscription
def create_subscription(tenant_id: str, tier: str):
    customer = stripe.Customer.create(
        email=tenant.email,
        metadata={"tenant_id": tenant_id}
    )
    
    price_id = {
        "starter": "price_starter_monthly",
        "growth": "price_growth_monthly",
        "enterprise": "price_enterprise_monthly"
    }[tier]
    
    subscription = stripe.Subscription.create(
        customer=customer.id,
        items=[{"price": price_id}],
        metadata={"tenant_id": tenant_id}
    )
    
    return subscription

# Upgrade/downgrade (prorated)
def change_tier(tenant_id: str, new_tier: str):
    subscription = stripe.Subscription.retrieve(tenant.stripe_subscription_id)
    
    new_price_id = {
        "starter": "price_starter_monthly",
        "growth": "price_growth_monthly",
        "enterprise": "price_enterprise_monthly"
    }[new_tier]
    
    stripe.Subscription.modify(
        subscription.id,
        items=[{
            "id": subscription["items"]["data"][0].id,
            "price": new_price_id
        }],
        proration_behavior="always_invoice"  # Prorated billing
    )
```

### Security Requirements
- Feature flags: Cached in Redis (5-minute TTL), not per-request DB query
- Usage metering: Async (don't block API requests)
- Billing webhooks: Verify Stripe signature (prevent fraud)
- Tenant isolation: RLS enforced even for shared resources
