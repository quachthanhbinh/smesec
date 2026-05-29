# SMESec Platform — Large-Scale Architecture (100K Tenants)

**Date:** 2026-05-29 | **Version:** 1.0 | **Status:** Draft  
**Audience:** Engineering Leadership · Solution Architects · CTO  
**Related:** [01-system-architecture.md](01-system-architecture.md) · [14-techstack-deep-dive.md](14-techstack-deep-dive.md)

---

## Purpose

The v1 architecture ([01-system-architecture.md](01-system-architecture.md)) is designed for ~1,000 tenants (v1 target). Key v1 components — RDS Proxy and GCP project pool — are provisioned in Sprint 1 precisely because v1 targets 1K tenants, and those components break at 500 and 70 tenants respectively. This document identifies every component that breaks **beyond** the 1K v1 baseline, quantifies the threshold, and specifies the architecture required to reach 100K tenants reliably.

This is a **forward-looking engineering reference** — not an immediate roadmap. The decisions here inform which v1 trade-offs to make carefully (i.e., which ones are expensive to undo) vs which ones are straightforward to evolve later.

---

## Table of Contents

1. [Scale Parameters](#1-scale-parameters)
2. [Current Architecture Bottlenecks](#2-current-architecture-bottlenecks)
3. [Target Architecture Overview](#3-target-architecture-overview)
4. [Database Layer — Sharding Strategy](#4-database-layer--sharding-strategy)
5. [Integration Sync Engine — Queue-Based Architecture](#5-integration-sync-engine--queue-based-architecture)
6. [Event Streaming — Kafka Replacement for EventBridge](#6-event-streaming--kafka-replacement-for-eventbridge)
7. [Authentication — Keycloak Federation at Scale](#7-authentication--keycloak-federation-at-scale)
8. [Multi-Region Active-Active Deployment](#8-multi-region-active-active-deployment)
9. [Google Workspace API — GCP Project Sharding](#9-google-workspace-api--gcp-project-sharding)
10. [Microsoft 365 — Webhook Management at Scale](#10-microsoft-365--webhook-management-at-scale)
11. [Caching Layer — Redis Cluster](#11-caching-layer--redis-cluster)
12. [Observability at Scale](#12-observability-at-scale)
13. [CQRS — Read/Write Separation](#13-cqrs--readwrite-separation)
14. [Migration Path — v1 to 100K Tenants](#14-migration-path--v1-to-100k-tenants)
15. [Cost Model at Scale](#15-cost-model-at-scale)
16. [Summary: What Must Be Correct in v1](#16-summary-what-must-be-correct-in-v1)

---

## 1. Scale Parameters

### 1.1 Target State — 100K Tenants

| Metric | v1 Target (1K tenants) | 100K Tenants Target |
|---|---|---|
| **Tenants** | 1,000 | 100,000 |
| **Users (avg 50/tenant)** | 50,000 | 5,000,000 |
| **OAuth app grants (avg 20/tenant)** | 20,000 | 2,000,000 |
| **Integration sync jobs/day** | 1K × 4 providers × 96 cycles = ~384K | 100K × 4 × 96 = ~38M |
| **EventBridge events/day** | ~100K | ~10M–50M |
| **API requests/sec (peak)** | ~1,000 | ~50,000 |
| **Audit log writes/day** | ~200K | ~10M |
| **PostgreSQL rows (assets table)** | ~1M | ~100M+ |
| **Active WebSocket connections** | ~1,000 | ~50,000 |
| **Offboarding workflows/day** | ~40 | ~4,000 |

### 1.2 Assumptions

- Average tenant size: 50 employees (SME target market unchanged)
- Average 4 integration providers per tenant (Google, M365, Slack, AWS IAM)
- Sync cadence remains 15-min for identity providers (96 cycles/day)
- 5% of tenants trigger at least 1 alert per day
- Peak traffic factor: 3× average (Monday 9AM spikes)

---

## 2. Current Architecture Bottlenecks

Every component below will fail or degrade before reaching 100K tenants. Threshold is the tenant count at which degradation begins.

### 2.1 Bottleneck Map

| Component | v1 Design | Failure Mode | Breaks At |
|---|---|---|---|
| **PostgreSQL (single cluster)** | Shared Multi-AZ RDS, all tenants on one DB | Write throughput saturated; vacuum/bloat on large tables; connection exhaustion | ~2,000 tenants |
| **RDS connection pool** | Direct connections from ECS tasks | `max_connections` (RDS db.r6g.xlarge = 3,200) exhausted: 100K × 4 ECS tasks × 5 connections = 2M | ~500 tenants |
| **Integration Sync (goroutines)** | Single SyncSvc ECS task, parallel goroutines | Memory exhaustion (1 goroutine ≈ 8KB × 100K = 800MB), API rate limits hit without throttling | ~5,000 tenants |
| **Google API quota** | 1,500 req/100s per GCP project, 20 tenants/project | At 100K tenants need 5,000 GCP projects; no architecture for this | ~70 tenants |
| **M365 webhooks** | Webhook renewal job, centralized `subscription_registry` | 100K webhook subscriptions in one table; renewal job serial execution would take hours | ~10,000 tenants |
| **AWS EventBridge** | All domain events on one event bus | EventBridge: 10,000 events/sec soft limit per account; at 100K tenants = 1M+ events/sec potential | ~50,000 tenants (with burst) |
| **Keycloak (single cluster)** | 2 ECS Fargate tasks, shared for all tenants | Single OIDC/SAML issuer; JWKS key rotation affects all tenants simultaneously | ~10,000 concurrent logins |
| **ElastiCache Redis (single node)** | cache.r6g.large (13GB RAM) | Memory ceiling; single point of failure; no horizontal key-space sharding | ~80,000 tenants |
| **S3 audit log (single bucket)** | Object Lock, WORM, all tenants | S3 scales horizontally but ListObjects performance degrades without prefix sharding | ~50,000 tenants |
| **ECS Fargate task sizing** | Min 2 tasks per service | Auto-scaling reacts in 2–3 minutes; 3× traffic spike at 100K tenants needs proactive scaling | ~20,000 tenants |

### 2.2 Critical Path Bottlenecks (Two fixed in v1, two in v2)

Two bottlenecks break **below** the 1K v1 target and are **mandatory Sprint 1 fixes**. Two others require **schema decisions in v1** to avoid expensive retrofits later:

1. **Google GCP project pool (Sprint 1 mandatory)** — breaks at ~70 tenants. 1K tenants requires 50 GCP projects (1K / 20 per project). `gcp_project_id` column in `tenant_config` schema from Sprint 1; SyncScheduler assigns projects at onboarding. Without this, Google quota is exhausted at ~70 tenants.
2. **RDS Proxy (Sprint 1 mandatory)** — breaks at ~500 tenants. At 1K tenants: 1K × 10 ECS tasks × 4 connections = 40K >> RDS max_connections (3,200 for db.r6g.2xlarge). Provision RDS Proxy in Sprint 1 infra before first production tenant.
3. **PostgreSQL horizontal partitioning (schema decision, Sprint 1)** — not needed until ~2,000 tenants but `tenant_id` sharding key must be in the v1 schema. Retrofitting a shard key onto an existing 1M+ row table requires a full migration with downtime.
4. **Kafka migration from EventBridge (v2, at 3K+ tenants)** — EventBridge event schema must be stable before Kafka adoption. Schema changes after migration require dual-publishing during transition. At 1K tenants, EventBridge handles ~100K events/day — well within the 10K/sec limit.

---

## 3. Target Architecture Overview

```
┌───────────────────────────────────────────────────────────────────────────┐
│  GLOBAL EDGE                                                               │
│  Cloudflare (DDoS, WAF, GeoDNS, anycast)                                  │
│  → AWS Global Accelerator → Regional ALBs (us-east-1, eu-west-1, ap-*)    │
└───────────────────────────────────┬───────────────────────────────────────┘
                                     │
┌────────────────────────────────────▼──────────────────────────────────────┐
│  API TIER (per region, active-active)                                      │
│  API Gateway (ECS, 10–100 tasks, HPA) + Kong rate limiting                 │
│  Keycloak Cluster (federated, region-aware, 4–20 tasks)                    │
└────────────────────────────────────┬──────────────────────────────────────┘
                                     │
┌────────────────────────────────────▼──────────────────────────────────────┐
│  APPLICATION TIER (per region, auto-scaled)                                │
│  ┌─────────────────┐  ┌─────────────────┐  ┌──────────────────────────┐  │
│  │ AssetSvc        │  │ AccessSvc        │  │ PlaybookSvc (Step Fn)    │  │
│  │ (read-heavy)    │  │ (write-critical) │  │ (stateful workflows)     │  │
│  │ → read replica  │  │ → shard primary  │  │ → SFN Express+Standard   │  │
│  └─────────────────┘  └─────────────────┘  └──────────────────────────┘  │
│  ┌─────────────────────────────────────────────────────────────────────┐  │
│  │  IntegrationSyncSvc — Queue-based, partitioned by tenant cluster     │  │
│  │  SyncWorker pods consume from Kafka partitions (1 partition/1K tenants)│ │
│  └─────────────────────────────────────────────────────────────────────┘  │
└────────────────────────────────────┬──────────────────────────────────────┘
                                     │
┌────────────────────────────────────▼──────────────────────────────────────┐
│  EVENT STREAMING — Apache Kafka (MSK Serverless)                           │
│  Topics: asset-events · access-events · threat-events · sync-jobs          │
│  Partitioning: by tenant_id hash → guarantees ordering per tenant          │
│  Consumers: PlaybookEngine · AuditLogger · NotificationSvc · AnalyticsSvc  │
└────────────────────────────────────┬──────────────────────────────────────┘
                                     │
┌────────────────────────────────────▼──────────────────────────────────────┐
│  DATA TIER                                                                 │
│  ┌──────────────────────────────────────────────────────────────────────┐ │
│  │  PostgreSQL — Horizontal Shard Clusters (Citus or manual sharding)   │ │
│  │  Shard key: tenant_id (consistent hash → shard node)                 │ │
│  │  Shard 0: tenant_id % 16 = 0  → RDS Cluster 0 (us-east-1)          │ │
│  │  ...                                                                  │ │
│  │  Shard 15: tenant_id % 16 = 15 → RDS Cluster 15 (us-east-1)        │ │
│  │  EU shards: eu-west-1 dedicated shard cluster (data_residency='EU') │ │
│  └──────────────────────────────────────────────────────────────────────┘ │
│  ┌─────────────────────────────┐  ┌───────────────────────────────────┐  │
│  │  Redis Cluster (6 nodes)    │  │  S3 (prefix-sharded by tenant)    │  │
│  │  3 primaries + 3 replicas   │  │  Bucket: audit-logs/{shard}/{tid} │  │
│  │  Keyspace hash slots: 16384 │  │  Object Lock WORM, KMS per-tenant │  │
│  └─────────────────────────────┘  └───────────────────────────────────┘  │
└───────────────────────────────────────────────────────────────────────────┘
```

---

## 4. Database Layer — Sharding Strategy

### 4.1 Problem: PostgreSQL Limits at Scale

A single RDS PostgreSQL Multi-AZ instance breaks in three ways:

| Limit | RDS db.r6g.4xlarge (largest practical single node) | At 100K Tenants |
|---|---|---|
| `max_connections` | 5,000 (with PgBouncer: ~50K) | Need: 100K × avg 5 concurrent = 500K |
| Write throughput | ~10K IOPS | Sync writes: 38M/day = ~440/sec sustained; peak 1,400/sec |
| Storage | 64TB practical limit | 100M assets × 2KB avg = 200GB — manageable but row bloat grows |
| Vacuum performance | Single autovacuum worker | 100M rows across hot tables causes vacuum lag → table bloat → slow queries |

### 4.2 Sharding Strategy — Horizontal Partitioning by `tenant_id`

**Chosen approach: Manual shard clusters (not Citus) at v2, Citus evaluation at v3.**

Reason: Citus adds operational complexity and requires changes to query patterns. At 100K tenants, manual tenant-cluster assignment (round-robin or consistent hash) is operationally simpler and uses standard RDS tooling.

```
Shard assignment:
  shard_id = hash(tenant_id) % NUM_SHARDS

  Phase 1 (0–2K tenants):   1 shard  → 1 RDS Multi-AZ cluster
  Phase 2 (2K–10K tenants): 4 shards → 4 RDS clusters
  Phase 3 (10K–50K tenants): 16 shards → 16 RDS clusters
  Phase 4 (50K–100K tenants): 32 shards → 32 RDS clusters
```

**Shard router (built into API middleware):**

```go
// All queries go through ShardRouter — never directly to a DB connection
type ShardRouter struct {
    shards    []*pgxpool.Pool  // one pool per shard
    numShards int
}

func (r *ShardRouter) ShardFor(tenantID uuid.UUID) *pgxpool.Pool {
    // Consistent hash: same tenant_id always maps to same shard
    h := fnv.New32a()
    h.Write(tenantID[:])
    shardIdx := int(h.Sum32()) % r.numShards
    return r.shards[shardIdx]
}
```

**v1 design requirement:** The `tenant_config` table must include a `shard_id` column from Sprint 1, even when there is only 1 shard. This enables zero-downtime shard migration later:

```sql
CREATE TABLE tenant_config (
    tenant_id      UUID PRIMARY KEY,
    shard_id       SMALLINT NOT NULL DEFAULT 0,  -- v1: always 0; grows with scale
    data_residency VARCHAR(10) NOT NULL,
    plan           VARCHAR(20) NOT NULL,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

### 4.3 Connection Pooling — PgBouncer (Required from Sprint 1)

**Problem:** ECS tasks create new DB connections per request. At 100K tenants, even moderate traffic exhausts `max_connections`.

**Solution:** PgBouncer as a sidecar on every ECS task, plus a shared PgBouncer cluster per shard:

```
ECS Task → PgBouncer (sidecar, transaction pooling) → PgBouncer (shard-level cluster) → RDS
            Pool: 20 connections/task                   Pool: 500 connections/shard
```

Transaction pooling mode: connection is returned to pool after every transaction — not held for the session lifetime. This is critical because `SET LOCAL app.tenant_id` must be re-issued at the start of every transaction (not just session), which Clean Architecture's middleware already does correctly.

**RDS Proxy:** Provision in v1 Sprint 1 even if not needed. RDS Proxy adds IAM auth and connection multiplexing. It can be enabled transparently — no application code change required.

### 4.4 Read Replicas — CQRS Readout

Read-heavy services (`AssetSvc`, `ComplianceSvc`, dashboard queries) must route to RDS read replicas:

```
Write path: Application → ShardRouter → Primary RDS (WRITE)
Read path:  Application → ShardRouter → Read Replica (READ, slight lag tolerable)

Acceptable lag: <5 seconds for asset inventory reads
Not acceptable: access governance checks (must read from primary — freshness critical)
```

---

## 5. Integration Sync Engine — Queue-Based Architecture

### 5.1 Problem: Goroutine-Based Sync Doesn't Scale

The v1 `IntegrationSyncSvc` runs parallel goroutines per tenant per provider. At 100K tenants:

- 100K tenants × 4 providers = 400K concurrent goroutines = ~3.2GB RAM (goroutine stack alone)
- No backpressure: if Google API is slow, all goroutines queue up in memory
- No durability: ECS task crash loses all in-flight sync jobs
- No observability: can't tell which tenants are lagging

### 5.2 Target Design — Kafka-Partitioned Sync Workers

```
┌──────────────────────────────────────────────────────────────┐
│  SyncScheduler (EventBridge cron, every 15 min)               │
│  Publishes SyncJob events to Kafka topic: sync-jobs           │
│  Partitioned by: hash(tenant_id) % 100                        │
│                                                                │
│  SyncJob = { tenant_id, provider, scheduled_at, priority }    │
└─────────────────────────────┬────────────────────────────────┘
                               │ Kafka (MSK Serverless)
                               │ Topic: sync-jobs (100 partitions)
                    ┌──────────┴──────────┐
                    │                     │
          ┌─────────▼──────┐   ┌──────────▼────────────────┐
          │ SyncWorker     │   │ SyncWorker                │
          │ (ECS, 10 pods) │   │ (ECS, 10 pods, EU region) │
          │ Consumes 10    │   │ Consumes EU-scoped         │
          │ partitions each│   │ partitions only            │
          └────────────────┘   └───────────────────────────┘
```

**Each SyncWorker:**
1. Pulls `SyncJob` from assigned Kafka partitions (at most 1,000 jobs/sec per worker)
2. Calls provider API (Google/M365/Slack/IAM) with exponential backoff
3. Writes deltas to PostgreSQL via shard router
4. Publishes `AssetDiscovered` / `AccessRevoked` events to `asset-events` topic
5. Commits Kafka offset only after successful DB write (at-least-once, idempotent writes)

**Why Kafka here (not EventBridge):**
- Kafka provides **durable, replayable job queue** with backpressure. If Google API is throttling, jobs queue in Kafka without blocking memory.
- Partition-based assignment guarantees **ordered processing per tenant** — no two workers process the same tenant simultaneously (prevents race conditions on delta sync cursors).
- EventBridge has no concept of consumer group, partition ordering, or offset commit.

### 5.3 Rate Limit Budget Allocation

```
Provider       | API Limit         | Tenants/Worker  | Max Workers (100K tenants)
───────────────┼───────────────────┼─────────────────┼───────────────────────────
Google Admin   | 1,500 req/100s    | 20 (per project)| 5,000 GCP projects (see §9)
               | per GCP project   |                 |
M365 Graph     | 10,000 req/10s    | 200             | 500 app registrations
               | per app reg       |                 |
Slack Admin    | 100 req/min       | 50              | 2,000 Slack app installs
               | per workspace     |                 |
AWS IAM        | 100 req/sec       | 100             | 1,000 assume-role accounts
               | per IAM call type |                 |
```

---

## 6. Event Streaming — Kafka Replacement for EventBridge

### 6.1 Why EventBridge Breaks at Scale

| Metric | EventBridge Limit | At 100K Tenants |
|---|---|---|
| Events/second per bus | 10,000/sec (soft limit) | Peak: 100K × 5 events/sec = 500K/sec |
| Event retention for replay | 24 hours (archive: unlimited but costly) | Need 7-year audit replay for compliance |
| Ordering guarantees | None | Tenant-level ordering required for CQRS projections |
| Consumer lag visibility | None | Need real-time consumer lag monitoring per tenant group |

### 6.2 Target: Apache Kafka on MSK Serverless

```
Topics and Partitioning Strategy:
───────────────────────────────────────────────────────────
Topic               | Partitions | Key            | Retention
────────────────────┼────────────┼────────────────┼──────────
asset-events        | 100        | tenant_id      | 7 days
access-events       | 100        | tenant_id      | 7 days
threat-events       | 50         | tenant_id      | 30 days
sync-jobs           | 100        | tenant_id      | 1 hour (jobs are ephemeral)
audit-log           | 200        | tenant_id      | 7 years (Tiered Storage → S3)
playbook-triggers   | 50         | tenant_id      | 1 day
notification-events | 20         | tenant_id      | 1 hour
```

**Kafka Tiered Storage** (MSK feature): Hot data (last 7 days) stays on Kafka brokers. Cold data offloads automatically to S3. This replaces the current S3 Object Lock direct-write pattern for audit logs — Kafka becomes the WAL (Write-Ahead Log) for the audit trail.

**Producer pattern (Go):**

```go
type KafkaEventPublisher struct {
    producer *kafka.Writer
    topic    string
}

func (p *KafkaEventPublisher) Publish(ctx context.Context, event DomainEvent) error {
    // Partition key = tenant_id → guaranteed ordering per tenant
    return p.producer.WriteMessages(ctx, kafka.Message{
        Key:   []byte(event.TenantID.String()),
        Value: mustMarshal(event),
        Headers: []kafka.Header{
            {Key: "event_type", Value: []byte(event.Type)},
            {Key: "schema_version", Value: []byte("v1")},
        },
    })
}
```

**Schema Registry (Confluent or AWS Glue):** All events must have registered Avro/Protobuf schemas. Breaking schema changes are prohibited — only backward-compatible evolution allowed. This enforces the `ThreatDetectionEvent` contract that currently relies on documentation alone.

### 6.3 Migration from EventBridge to Kafka

The migration is non-breaking if done via a **dual-publish transition period**:

```
Phase A (current):   Publisher → EventBridge → Consumers
Phase B (migration): Publisher → EventBridge + Kafka → Consumers (consume from both, dedup)
Phase C (complete):  Publisher → Kafka only → Consumers
```

Duration of Phase B: ~1 sprint (2 weeks) of running both in parallel, comparing consumer outputs.

---

## 7. Authentication — Keycloak Federation at Scale

### 7.1 Problem: Single Keycloak Cluster

v1 uses 2 ECS tasks (active-active) with a single PostgreSQL database. At 100K tenants:

- JWKS key rotation: affects all 5M users simultaneously → logout storm
- Single realm for all tenants: Keycloak realm DB row count grows unbounded
- Admin API (tenant provisioning) contends with auth traffic
- One Keycloak DB failure = global auth outage

### 7.2 Target: Federated Multi-Realm Keycloak Clusters

```
┌──────────────────────────────────────────────────────────────┐
│  Keycloak Meta-Realm (Global)                                 │
│  Purpose: Cross-tenant admin, API key issuance only          │
│  Scale: 4 tasks, dedicated RDS, rarely accessed              │
└─────────────────────────────┬────────────────────────────────┘
                               │ JWKS federation
          ┌────────────────────┼─────────────────────┐
          │                    │                     │
┌─────────▼──────┐  ┌──────────▼──────┐  ┌──────────▼──────┐
│ KC Cluster A   │  │ KC Cluster B    │  │ KC Cluster EU   │
│ Tenants 0–25K  │  │ Tenants 25K–50K │  │ EU tenants only  │
│ us-east-1      │  │ us-east-1       │  │ eu-west-1        │
│ 4–20 ECS tasks │  │ 4–20 ECS tasks  │  │ 4–20 ECS tasks   │
│ Own RDS DB     │  │ Own RDS DB      │  │ Own RDS DB       │
└────────────────┘  └─────────────────┘  └─────────────────┘
```

**Tenant-to-cluster assignment** stored in `tenant_config.keycloak_cluster_id`. The API Gateway reads this at JWT validation time to forward to the correct JWKS endpoint.

**JWKS caching (critical):** Each API Gateway instance caches JWKS keys with a 6-hour TTL. JWT validation never requires a live Keycloak call in steady state. A Keycloak cluster going offline only affects new logins, not existing sessions.

### 7.3 WorkOS / Auth0 Evaluation Gate

At v1.5 retrospective: if Keycloak operational overhead exceeds 20% of DevSecOps time, migrate to WorkOS (flat $0.10/tenant/mo = $10K/mo at 100K tenants — acceptable against $30K saved in engineering time).

---

## 8. Multi-Region Active-Active Deployment

### 8.1 v1 vs Large-Scale Deployment Model

| | v1 (1K tenants) | 100K Tenants |
|---|---|---|
| `us-east-1` | Active (all US traffic) | Active (US traffic) |
| `eu-west-1` | DR standby (EU data compliance) | **Active** (EU traffic) |
| `ap-southeast-1` | Not provisioned | Active (APAC traffic) |
| `us-west-2` | Not provisioned | Active (US West failover) |

### 8.2 Active-Active Requirements

**Data sovereignty (hard constraint from v1 schema):**
- EU tenant data: written only to `eu-west-1` RDS + `eu-west-1` S3 + `eu-west-1` KMS
- US tenant data: written only to `us-east-1` RDS + `us-east-1` S3 + `us-east-1` KMS
- The `data_residency` column in every table enforces this. **This constraint must not be violated by any caching layer, CDN, or read replica.**

**Global state that IS cross-region:**
- `tenant_config` table: replicated read-only to all regions (low-write, high-read)
- Kafka: MSK in each region; cross-region replication via MirrorMaker 2 for `threat-events` only (global threat intelligence feeds)
- KMS: regional keys only — never cross-region. Each region has its own CMK per tenant.

```
GeoDNS Routing:
  smesec.com → Cloudflare → Route by client IP
    EU users → eu-west-1 ALB
    APAC users → ap-southeast-1 ALB
    US users → us-east-1 ALB (primary) / us-west-2 (failover)

Failover SLA:
  RTO: < 60 seconds (Route 53 health check interval: 10s + DNS TTL: 30s + warmup: 20s)
  RPO: < 30 seconds (PostgreSQL synchronous replication across AZs; async to standby region)
```

### 8.3 Cross-Region Consistency Model

| Data Type | Consistency | Mechanism |
|---|---|---|
| User sessions (Redis) | Eventually consistent across regions | ElastiCache Global Datastore (async replication, <1s lag) |
| Tenant config | Eventually consistent (seconds) | PostgreSQL logical replication read replica in each region |
| Asset inventory | Regionally consistent (strict within region) | No cross-region writes; each region owns its tenant set |
| Audit logs | Append-only, never cross-region | S3 Object Lock + Kafka Tiered Storage per region |
| Threat events | Eventually consistent (seconds) | Kafka MirrorMaker 2, topic: `threat-events-global` |

---

## 9. Google Workspace API — GCP Project Sharding

### 9.1 Problem: Hard API Quota at 70 Tenants

Google Admin SDK quota: **1,500 requests per 100 seconds per GCP project.** This is a hard quota, not softcappable.

| Tenants | Sync requests per 100s | GCP Projects Required |
|---|---|---|
| 20 | 1,200 (safe) | 1 |
| 50 | 3,000 | 2 |
| 100 | 6,000 | 4 |
| 1,000 | 60,000 | 40 |
| 100,000 | 6,000,000 | 4,000 |

### 9.2 GCP Project Pool Architecture

```
GCP Project Pool Manager (internal service):
  ┌─────────────────────────────────────────────────────────────┐
  │ gcp_project_pool table (PostgreSQL):                        │
  │   project_id  | service_account_id | tenant_count | region  │
  │   gcp-proj-0  | sa-0@proj-0.iam    | 20           | us      │
  │   gcp-proj-1  | sa-1@proj-1.iam    | 20           | us      │
  │   ...                                                        │
  │   gcp-proj-4999 | sa-4999@...      | 20           | eu      │
  └─────────────────────────────────────────────────────────────┘

Tenant onboarding flow:
  1. Find GCP project with tenant_count < 20
  2. Assign tenant to that project
  3. Store gcp_project_id in tenant_config
  4. SyncWorker reads tenant_config.gcp_project_id to select service account credentials
```

**Credential rotation:** Each service account rotates its key every 90 days. The rotation job is batched by project (1 project/minute) to avoid simultaneous key rotation across all projects.

**Automation:** GCP project provisioning is automated via Terraform + GCP Organization-level Service Account. New GCP projects are pre-provisioned in batches of 100 when pool utilization exceeds 80%.

---

## 10. Microsoft 365 — Webhook Management at Scale

### 10.1 Problem: 100K Active Webhook Subscriptions

v1 has a single `subscription_registry` table and a serial renewal job. At 100K tenants:

- 100K webhook subscriptions, each expiring every 3 days
- Serial renewal at 1 req/sec = ~28 hours to renew all — **renewal never completes before next expiry**
- A single `subscription_registry` table becomes a hot write target

### 10.2 Target Architecture

```
Renewal Strategy:
  Partition subscriptions into 72 renewal buckets (= 3 days × 24 hours)
  Each bucket: ~1,400 subscriptions (100K / 72)
  Bucket renewal job runs every hour via EventBridge Scheduler
  Each bucket job renews its 1,400 subscriptions in parallel (async, 100 goroutines)
  Duration per bucket: ~2 minutes at Graph API throttle limits

  subscription_registry:
    + renewal_bucket  SMALLINT  (0–71, assigned at subscription creation: hash(tenant_id) % 72)
    + renewal_shard   SMALLINT  (0–7, for DB write sharding)

  Renewal job query (efficient index scan):
    SELECT * FROM subscription_registry
    WHERE renewal_bucket = $current_bucket
    AND expiry_at < NOW() + INTERVAL '25 hours'
    AND status = 'active'
```

**Per-tenant Graph API app registrations:** At 100K tenants, the single M365 app registration hits per-app throttle limits. Target: 1 app registration per 200 tenants → 500 app registrations, each with its own `client_id / client_secret` stored in Secrets Manager.

---

## 11. Caching Layer — Redis Cluster

### 11.1 Problem: Single Redis Node Memory Ceiling

v1 uses `cache.r6g.large` (~13GB RAM, dedicated CPU — no bursting). At 100K tenants:

| Cache Key Type | Size per Tenant | Total at 100K Tenants |
|---|---|---|
| Session tokens (JWT) | 50 users × 512 bytes = 25KB | 2.5GB |
| Permission sets (user roles) | 50 users × 2KB = 100KB | 10GB |
| Sync state (delta cursors) | 4 providers × 256 bytes = 1KB | 100MB |
| Rate limit counters | 100 keys × 8 bytes = 800 bytes | 80MB |
| Asset classification cache (30s) | 20 apps × 512 bytes = 10KB | 1GB |
| **Total** | | **~14GB** |

### 11.2 Target: Redis Cluster (6-Node)

```
Redis Cluster configuration:
  3 primaries + 3 replicas (1 replica per primary)
  Hash slots: 16,384 (evenly distributed across 3 primaries)
  Each primary handles: ~5,000 hash slots → ~4.7GB working set

  Node sizing: cache.r6g.large (13GB RAM each)
  Total cluster capacity: 3 × 13GB = 39GB (headroom for 100K tenants)

  Keyspace design:
    {tenant_id}:session:{user_id}     → session tokens (15-min TTL)
    {tenant_id}:perm:{user_id}        → permission set cache (5-min TTL)
    {tenant_id}:sync:{provider}:cursor → delta link / page token (no TTL)
    {tenant_id}:ratelimit:{service}   → rate limit counter (1-min window)

  Hash tags: {{ tenant_id }} ensures all keys for one tenant land on same shard
  → Enables MULTI/EXEC transactions within a single tenant without cross-slot errors
```

---

## 12. Observability at Scale

### 12.1 Requirements at 100K Tenants

At v1 scale (1K tenants), CloudWatch metrics with per-tenant tagging and centralized logs are sufficient for standard operations. At 100K tenants:

- Log volume: ~100GB/day (5M users × 20 log lines/day × 1KB avg)
- Trace volume: 50K req/sec × 3 spans each = 150K spans/sec
- Metric cardinality: per-tenant metrics × 100K = cardinality explosion in CloudWatch

### 12.2 OpenTelemetry-First Observability

All services emit traces, metrics, and logs via the **OpenTelemetry SDK** (Go: `go.opentelemetry.io/otel`). The collector layer decides where to ship — not the application.

```
Application Service
    │ OTLP (gRPC)
    ▼
OpenTelemetry Collector (sidecar per ECS task)
    ├─→ Traces → AWS X-Ray (short-term, 30-day retention)
    │            OR Grafana Tempo (self-hosted, 90-day retention at scale)
    ├─→ Metrics → Amazon Managed Prometheus (AMP) + Grafana dashboards
    │             Cardinality reduction: aggregate per-tenant → per-shard
    └─→ Logs → Amazon OpenSearch (hot: 7 days) + S3 (cold: 1 year)
```

**Tenant-aware alerting:** Alerts fire on **per-shard** anomalies (e.g., shard 3 has 40% sync failure rate) rather than per-tenant. Drilling down to specific tenants is a second-level investigation step.

**SLO tracking (per service):**

| Service | SLO | Error Budget (30-day) |
|---|---|---|
| Offboarding workflow | < 5 min p99 | 0.1% (4.3 min/month) |
| Asset sync freshness | < 15 min p95 | 5% (36 hrs/month) |
| API Gateway availability | > 99.9% | 43.2 min/month |
| Auth (Keycloak) | > 99.95% | 21.6 min/month |

---

## 13. CQRS — Read/Write Separation

### 13.1 Why CQRS at 100K Tenants

Dashboard queries at 100K tenants are expensive: "show me all 20 connected OAuth apps for tenant X with their risk scores" requires JOINs across `assets`, `oauth_grants`, `risk_scores`, and `users` tables. These queries compete with write traffic from sync workers.

### 13.2 Materialized Projection Tables (Lightweight CQRS)

Full Event Sourcing is not adopted. Instead, **materialized projection tables** are maintained by Kafka consumers:

```
Write model (normalized):
  assets, oauth_grants, users, risk_scores (normalized, sync workers write here)

Read model (denormalized projections, per-tenant):
  tenant_dashboard_snapshot:
    tenant_id, snapshot_json, updated_at
    → rebuilt by DashboardProjector on every asset-events Kafka message for that tenant
    → dashboard GET /api/v1/dashboard → read from snapshot (< 10ms), no JOIN

  offboarding_readiness:
    tenant_id, user_id, providers_json, estimated_duration_sec, updated_at
    → rebuilt by OffboardingProjector on every access-events message

  compliance_posture:
    tenant_id, framework, controls_json, score, last_evaluated_at
    → rebuilt by ComplianceProjector nightly (expensive, run async)
```

**Eventual consistency window:** Dashboard snapshot may be up to 5 seconds stale (Kafka consumer lag). Acceptable for read-heavy views. Write operations (offboarding, access revocation) always use the normalized write model.

---

## 14. Migration Path — v1 to 100K Tenants

### 14.1 Evolution Stages

```
Stage       | Tenant Range | Key Architecture Change                  | Trigger
────────────┼──────────────┼──────────────────────────────────────────┼─────────────────────────
v1          | 1–1K         | Single RDS + RDS Proxy, GCP project pool (50 projects), EventBridge, bounded goroutine sync (200 workers max) | Baseline — designed for 1K
v1.5        | 1K–3K        | + Redis Cluster, PgBouncer read replicas, concurrent M365 webhook renewal | DB connections > 80% of proxy capacity
v2          | 3K–10K       | + Kafka (MSK Serverless), Sync Workers via Kafka partitions               | Events/day > 500K or sync lag > 15min
v2.5        | 10K–30K      | + DB sharding (4 shard clusters)                                          | p99 query > 200ms
v3          | 30K–100K    | + Multi-region active-active, KC federation, Shard ×16                   | EU revenue > 20% ARR
```

### 14.2 Zero-Downtime Migration Principles

All migrations must meet these constraints:
1. **No global lock migrations.** Use `pg_repack` or blue-green RDS swap.
2. **Dual-write during transitions.** When migrating from EventBridge → Kafka, publish to both for 2 weeks before cutting over.
3. **Shard migration is online.** A tenant can be moved to a new shard by: (1) dual-write to old + new shard, (2) backfill, (3) verify checksum, (4) cut read to new shard, (5) stop old shard writes.
4. **Feature flags for infrastructure.** `SYNC_ENGINE=kafka` or `SYNC_ENGINE=goroutine` are runtime feature flags — rollback is a config change, not a deploy.

---

## 15. Cost Model at Scale

### 15.1 Infrastructure Cost Projection

| Component | v1 (1K tenants) | 10K tenants | 100K tenants |
|---|---|---|---|
| RDS PostgreSQL (shards) | $1,000/mo (r6g.2xlarge Multi-AZ + RDS Proxy) | $2,000/mo (4 shards) | $16,000/mo (32 shards) |
| ECS Fargate (all services) | $2,500/mo | $4,000/mo | $30,000/mo |
| Kafka (MSK Serverless) | $0 (EventBridge at 1K) | $500/mo | $5,000/mo |
| Redis (ElastiCache Cluster) | $200/mo (cache.r6g.large) | $500/mo | $3,000/mo |
| Keycloak (ECS + RDS) | $200/mo | $600/mo | $3,000/mo (federated) |
| S3 + CloudFront | $200/mo | $500/mo | $3,000/mo |
| Secrets Manager (batched, 1K) | $400/mo | $2,000/mo | $20,000/mo |
| Third-party APIs (Vanta, Hive, Lakera) | $2,000/mo | $5,000/mo | $30,000/mo |
| **Total infra** | **~$6,500/mo** | **~$15,000/mo** | **~$110,000/mo** |
| **Revenue (avg $800/mo/tenant)** | **$800K/mo** | **$8M/mo** | **$80M/mo** |
| **Infra as % of Revenue** | **0.8%** | **0.2%** | **0.14%** |

**Unit economics improve significantly with scale.** The architecture choices (shared PostgreSQL + RLS, serverless event bus, managed services) are specifically chosen to keep per-tenant infra cost low as scale grows.

### 15.2 Secrets Manager Cost — Architecture Concern

At 100K tenants × 4 OAuth tokens = 400K secrets × $0.40/secret/month = **$160K/month** — a material cost.

**Mitigation:** Batch secrets per tenant into a single encrypted JSON blob:

```json
// Single secret: smesec/{tenant_id}/oauth
{
  "google_refresh_token": "...",
  "m365_client_secret": "...",
  "slack_bot_token": "...",
  "aws_role_arn": "..."
}
// Cost: $0.40/mo vs $1.60/mo (4 secrets) → 4× savings → $40K/mo instead of $160K/mo
```

This requires a secret accessor update but is a schema change only — no architectural impact.

---

## 16. Summary: What Must Be Correct in v1

These v1 decisions are **load-bearing** — getting them wrong requires expensive migrations. All other decisions are safe to evolve.

| Decision | v1 Requirement | Why It's Hard to Undo |
|---|---|---|
| `tenant_id` as shard key | Add `shard_id` column to `tenant_config` from Sprint 1 | Retrofitting shard key onto 100M-row table = full migration with downtime |
| `data_residency` column on every table | Mandatory from Sprint 1 | Adding it later requires schema migration + backfill + EU data audit |
| `gcp_project_id` in `tenant_config` | Add column in Sprint 1 (value: `default`) | Multi-project Google sync requires per-tenant GCP project assignment; retrofitting = credential re-registration for all tenants |
| `renewal_bucket` in `subscription_registry` | Add column in Sprint 1 | Distributed renewal job design depends on this; adding later = full table rewrite |
| PgBouncer / RDS Proxy | Provision in Sprint 1, enable at ~100 tenants | Connection exhaustion is abrupt, not gradual — you hit the wall suddenly |
| EventBridge event schema stability | Freeze `ThreatDetectionEvent` schema at Sprint 10 | Kafka migration during schema flux = dual-publish complexity × schema evolution |
| OpenTelemetry SDK in all services | Instrument from Sprint 1 | Retrofitting tracing into existing services = rewrite of all middleware |
| OAuth token batching in Secrets Manager | Store as single JSON secret per tenant | Splitting later = 4× cost increase + migration of all existing tenants |

---

*Document owner: Technical Advisor / Solution Architect*  
*Next review: Before v2 planning (Month 7)*
