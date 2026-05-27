# Hướng Dẫn Triển Khai Kỹ Thuật

## Giới Thiệu

Tài liệu này cung cấp hướng dẫn chi tiết từng bước để triển khai các kiểm soát bảo mật kỹ thuật cho SMESec, đảm bảo tuân thủ ISO 27001, SOC 2, và GDPR.

## A. Quản Lý Mã Nguồn (GitHub)

### 1. Branch Protection Rules

**Mục đích:** Ngăn chặn thay đổi trực tiếp vào production code, đảm bảo code review.

**Cấu hình:**

```bash
# Navigate to GitHub repository settings
# Settings → Branches → Add branch protection rule

Branch name pattern: main
☑ Require a pull request before merging
  ☑ Require approvals: 1
  ☑ Dismiss stale pull request approvals when new commits are pushed
  ☑ Require review from Code Owners (if CODEOWNERS file exists)
☑ Require status checks to pass before merging
  ☑ Require branches to be up to date before merging
  Required status checks:
    - security-scan
    - dependabot
    - codeql
☑ Require conversation resolution before merging
☑ Require signed commits (optional but recommended)
☑ Include administrators (enforce rules for admins too)
☑ Restrict who can push to matching branches
  - Add: Senior Engineers only
```

**Repeat for `production` branch** (if separate from main)

**Validation:**
```bash
# Test that direct push is blocked
git checkout main
echo "test" >> test.txt
git add test.txt
git commit -m "test direct push"
git push origin main
# Expected: Error - protected branch
```

### 2. GitHub Advanced Security

**Enable Dependabot:**

```yaml
# .github/dependabot.yml
version: 2
updates:
  # Enable version updates for npm
  - package-ecosystem: "npm"
    directory: "/"
    schedule:
      interval: "weekly"
    open-pull-requests-limit: 10
    reviewers:
      - "senior-engineer-team"
    labels:
      - "dependencies"
      - "security"
    
  # Enable version updates for pip (if using Python)
  - package-ecosystem: "pip"
    directory: "/"
    schedule:
      interval: "weekly"
    
  # Enable version updates for GitHub Actions
  - package-ecosystem: "github-actions"
    directory: "/"
    schedule:
      interval: "weekly"
```

**Enable CodeQL:**

```yaml
# .github/workflows/codeql.yml
name: "CodeQL Security Scan"

on:
  push:
    branches: [ main, production ]
  pull_request:
    branches: [ main, production ]
  schedule:
    - cron: '0 6 * * 1'  # Weekly on Monday at 6 AM

jobs:
  analyze:
    name: Analyze
    runs-on: ubuntu-latest
    permissions:
      actions: read
      contents: read
      security-events: write

    strategy:
      fail-fast: false
      matrix:
        language: [ 'javascript', 'typescript', 'python' ]
        # Add other languages as needed

    steps:
    - name: Checkout repository
      uses: actions/checkout@v4

    - name: Initialize CodeQL
      uses: github/codeql-action/init@v3
      with:
        languages: ${{ matrix.language }}
        queries: security-extended,security-and-quality

    - name: Autobuild
      uses: github/codeql-action/autobuild@v3

    - name: Perform CodeQL Analysis
      uses: github/codeql-action/analyze@v3
      with:
        category: "/language:${{matrix.language}}"
```

**Enable Secret Scanning:**

```bash
# Navigate to GitHub repository settings
# Settings → Code security and analysis

☑ Dependency graph (should be enabled by default)
☑ Dependabot alerts
☑ Dependabot security updates
☑ CodeQL analysis
☑ Secret scanning
☑ Push protection (prevents pushing secrets)
```

### 3. Security CI/CD Pipeline

**Create comprehensive security workflow:**

```yaml
# .github/workflows/security.yml
name: Security Checks

on:
  pull_request:
    branches: [ main, production ]
  push:
    branches: [ main, production ]

jobs:
  security-scan:
    name: Security Scan
    runs-on: ubuntu-latest
    
    steps:
    - name: Checkout code
      uses: actions/checkout@v4
      with:
        fetch-depth: 0  # Full history for better analysis

    - name: Run Dependabot
      uses: github/dependabot-action@v1
      continue-on-error: false

    - name: Check for secrets
      uses: trufflesecurity/trufflehog@main
      with:
        path: ./
        base: ${{ github.event.repository.default_branch }}
        head: HEAD

    - name: Security audit (npm)
      if: hashFiles('package-lock.json') != ''
      run: |
        npm audit --audit-level=high
        npm audit --audit-level=critical

    - name: Security audit (pip)
      if: hashFiles('requirements.txt') != ''
      run: |
        pip install safety
        safety check --file requirements.txt

    - name: SAST with Semgrep
      uses: returntocorp/semgrep-action@v1
      with:
        config: >-
          p/security-audit
          p/secrets
          p/owasp-top-ten

    - name: Block on High/Critical findings
      if: failure()
      run: |
        echo "❌ Security scan failed - High or Critical vulnerabilities found"
        echo "Please fix security issues before merging"
        exit 1
```

**Validation:**
```bash
# Create a test PR with a known vulnerability
# Verify that the PR is blocked
```

### 4. CODEOWNERS File

**Purpose:** Ensure critical files are reviewed by appropriate team members.

```bash
# .github/CODEOWNERS

# Default owners for everything
* @senior-engineer-team

# Security-sensitive files require security review
/src/auth/** @security-team @senior-engineer-team
/src/api/middleware/auth.ts @security-team
/.github/workflows/** @devops-team @senior-engineer-team

# Infrastructure code requires DevOps review
/terraform/** @devops-team
/infrastructure/** @devops-team
*.tf @devops-team

# Compliance documentation requires compliance lead
/docs/compliance/** @compliance-lead
/docs/security/** @compliance-lead

# Database migrations require careful review
/migrations/** @database-team @senior-engineer-team
```

## B. Quản Lý Hạ Tầng (AWS & Cloudflare R2)

### 1. AWS IAM Identity Center Setup

**Step 1: Enable IAM Identity Center**

```bash
# Via AWS Console
# Navigate to: IAM Identity Center → Enable

# Or via AWS CLI
aws sso-admin create-instance \
  --region us-east-1
```

**Step 2: Create Permission Sets**

```bash
# Create Developer permission set
aws sso-admin create-permission-set \
  --instance-arn arn:aws:sso:::instance/ssoins-xxxxx \
  --name "DeveloperAccess" \
  --description "Read-only access with limited write permissions" \
  --session-duration "PT8H"

# Attach managed policy
aws sso-admin attach-managed-policy-to-permission-set \
  --instance-arn arn:aws:sso:::instance/ssoins-xxxxx \
  --permission-set-arn arn:aws:sso:::permissionSet/xxxxx \
  --managed-policy-arn arn:aws:iam::aws:policy/ReadOnlyAccess
```

**Step 3: Create Custom Policies**

```json
// developer-policy.json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "DeveloperReadAccess",
      "Effect": "Allow",
      "Action": [
        "ec2:Describe*",
        "s3:Get*",
        "s3:List*",
        "rds:Describe*",
        "logs:Get*",
        "logs:Describe*",
        "cloudwatch:Get*",
        "cloudwatch:List*"
      ],
      "Resource": "*"
    },
    {
      "Sid": "DeveloperWriteAccessDev",
      "Effect": "Allow",
      "Action": [
        "s3:PutObject",
        "s3:DeleteObject"
      ],
      "Resource": "arn:aws:s3:::smesec-dev-*/*",
      "Condition": {
        "StringEquals": {
          "aws:RequestedRegion": "us-east-1"
        }
      }
    },
    {
      "Sid": "DenyProductionWrite",
      "Effect": "Deny",
      "Action": [
        "s3:PutObject",
        "s3:DeleteObject",
        "rds:Delete*",
        "rds:Modify*"
      ],
      "Resource": [
        "arn:aws:s3:::smesec-prod-*/*",
        "arn:aws:rds:*:*:db:smesec-prod-*"
      ]
    }
  ]
}
```

**Step 4: Enforce MFA**

```json
// mfa-policy.json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "DenyAllExceptListedIfNoMFA",
      "Effect": "Deny",
      "NotAction": [
        "iam:CreateVirtualMFADevice",
        "iam:EnableMFADevice",
        "iam:GetUser",
        "iam:ListMFADevices",
        "iam:ListVirtualMFADevices",
        "iam:ResyncMFADevice",
        "sts:GetSessionToken"
      ],
      "Resource": "*",
      "Condition": {
        "BoolIfExists": {
          "aws:MultiFactorAuthPresent": "false"
        }
      }
    }
  ]
}
```

### 2. AWS CloudTrail Configuration

**Enable CloudTrail with best practices:**

```bash
# Create S3 bucket for logs
aws s3api create-bucket \
  --bucket smesec-cloudtrail-logs \
  --region us-east-1

# Enable versioning
aws s3api put-bucket-versioning \
  --bucket smesec-cloudtrail-logs \
  --versioning-configuration Status=Enabled

# Enable encryption
aws s3api put-bucket-encryption \
  --bucket smesec-cloudtrail-logs \
  --server-side-encryption-configuration '{
    "Rules": [{
      "ApplyServerSideEncryptionByDefault": {
        "SSEAlgorithm": "AES256"
      }
    }]
  }'

# Enable Object Lock (WORM - Write Once Read Many)
aws s3api put-object-lock-configuration \
  --bucket smesec-cloudtrail-logs \
  --object-lock-configuration '{
    "ObjectLockEnabled": "Enabled",
    "Rule": {
      "DefaultRetention": {
        "Mode": "GOVERNANCE",
        "Days": 90
      }
    }
  }'

# Create CloudTrail
aws cloudtrail create-trail \
  --name smesec-main-trail \
  --s3-bucket-name smesec-cloudtrail-logs \
  --is-multi-region-trail \
  --enable-log-file-validation \
  --include-global-service-events

# Start logging
aws cloudtrail start-logging \
  --name smesec-main-trail
```

**Terraform equivalent:**

```hcl
# terraform/cloudtrail.tf
resource "aws_s3_bucket" "cloudtrail_logs" {
  bucket = "smesec-cloudtrail-logs"
  
  versioning {
    enabled = true
  }
  
  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        sse_algorithm = "AES256"
      }
    }
  }
  
  object_lock_configuration {
    object_lock_enabled = "Enabled"
    
    rule {
      default_retention {
        mode = "GOVERNANCE"
        days = 90
      }
    }
  }
  
  lifecycle_rule {
    enabled = true
    
    transition {
      days          = 90
      storage_class = "GLACIER"
    }
    
    expiration {
      days = 365
    }
  }
}

resource "aws_cloudtrail" "main" {
  name                          = "smesec-main-trail"
  s3_bucket_name                = aws_s3_bucket.cloudtrail_logs.id
  is_multi_region_trail         = true
  enable_log_file_validation    = true
  include_global_service_events = true
  
  event_selector {
    read_write_type           = "All"
    include_management_events = true
    
    data_resource {
      type   = "AWS::S3::Object"
      values = ["arn:aws:s3:::smesec-*/*"]
    }
  }
}
```

### 3. Cloudflare R2 Configuration

**Create R2 bucket with encryption:**

```bash
# Via Cloudflare Dashboard
# Storage → R2 → Create bucket

Bucket name: smesec-prod-data
Location: Automatic (closest to users)
Storage class: Standard

# Enable encryption (default in R2)
# Configure access policies
```

**Access Policy (Least Privilege):**

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "ApplicationReadWrite",
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::ACCOUNT_ID:role/smesec-app-role"
      },
      "Action": [
        "s3:GetObject",
        "s3:PutObject",
        "s3:DeleteObject"
      ],
      "Resource": "arn:aws:s3:::smesec-prod-data/user-uploads/*"
    },
    {
      "Sid": "DenyUnencryptedObjectUploads",
      "Effect": "Deny",
      "Principal": "*",
      "Action": "s3:PutObject",
      "Resource": "arn:aws:s3:::smesec-prod-data/*",
      "Condition": {
        "StringNotEquals": {
          "s3:x-amz-server-side-encryption": "AES256"
        }
      }
    },
    {
      "Sid": "DenyPublicAccess",
      "Effect": "Deny",
      "Principal": "*",
      "Action": "s3:*",
      "Resource": [
        "arn:aws:s3:::smesec-prod-data",
        "arn:aws:s3:::smesec-prod-data/*"
      ],
      "Condition": {
        "StringEquals": {
          "s3:x-amz-acl": [
            "public-read",
            "public-read-write"
          ]
        }
      }
    }
  ]
}
```

### 4. Database Security (RDS)

**Terraform configuration for secure RDS:**

```hcl
# terraform/rds.tf
resource "aws_db_subnet_group" "private" {
  name       = "smesec-db-subnet-group"
  subnet_ids = aws_subnet.private[*].id
  
  tags = {
    Name        = "SMESec DB Subnet Group"
    Environment = "production"
  }
}

resource "aws_db_instance" "main" {
  identifier     = "smesec-prod-db"
  engine         = "postgres"
  engine_version = "15.4"
  instance_class = "db.t3.medium"
  
  allocated_storage     = 100
  max_allocated_storage = 500
  storage_encrypted     = true
  kms_key_id            = aws_kms_key.rds.arn
  
  db_name  = "smesec"
  username = "smesec_admin"
  password = random_password.db_password.result
  
  db_subnet_group_name   = aws_db_subnet_group.private.name
  vpc_security_group_ids = [aws_security_group.rds.id]
  publicly_accessible    = false
  
  backup_retention_period = 7
  backup_window           = "03:00-04:00"
  maintenance_window      = "mon:04:00-mon:05:00"
  
  enabled_cloudwatch_logs_exports = ["postgresql", "upgrade"]
  
  deletion_protection = true
  skip_final_snapshot = false
  final_snapshot_identifier = "smesec-prod-db-final-snapshot"
  
  tags = {
    Name        = "SMESec Production DB"
    Environment = "production"
    Compliance  = "ISO27001,SOC2"
  }
}

resource "aws_security_group" "rds" {
  name        = "smesec-rds-sg"
  description = "Security group for RDS database"
  vpc_id      = aws_vpc.main.id
  
  ingress {
    description     = "PostgreSQL from application"
    from_port       = 5432
    to_port         = 5432
    protocol        = "tcp"
    security_groups = [aws_security_group.app.id]
  }
  
  egress {
    description = "No outbound traffic"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  
  tags = {
    Name = "SMESec RDS Security Group"
  }
}
```

## C. Vận Hành (Runbooks)

### 1. Incident Response Runbook

**File:** `docs/runbooks/incident-response.md`

```markdown
# Incident Response Runbook

## Detection
- CloudWatch alarms
- Vanta security alerts
- User reports
- Monitoring dashboards

## Severity Classification

### Critical (P0)
- Data breach or suspected breach
- Complete service outage
- Security vulnerability being actively exploited

### High (P1)
- Partial service outage
- Performance degradation affecting >50% users
- Security vulnerability discovered (not yet exploited)

### Medium (P2)
- Minor service degradation
- Security misconfiguration
- Compliance violation

### Low (P3)
- Cosmetic issues
- Documentation errors

## Response Procedure

### Step 1: Acknowledge (Within 15 minutes)
```bash
# Post in #incidents Slack channel
Incident detected: [Brief description]
Severity: [P0/P1/P2/P3]
Incident Commander: [Your name]
Status page: [Link if applicable]
```

### Step 2: Assess (Within 30 minutes)
- Determine scope and impact
- Identify affected systems
- Estimate user impact
- Classify severity

### Step 3: Contain (Within 1 hour for P0/P1)
- Isolate affected systems
- Prevent further damage
- Preserve evidence (for security incidents)

### Step 4: Investigate
- Review logs (CloudTrail, application logs)
- Identify root cause
- Document findings

### Step 5: Remediate
- Apply fix
- Verify fix effectiveness
- Monitor for recurrence

### Step 6: Communicate
- Update stakeholders
- Post-mortem (within 48 hours for P0/P1)
- Update documentation

### Step 7: Learn
- Conduct blameless post-mortem
- Identify preventive measures
- Update runbooks
- Implement improvements

## Security Incident Specific Steps

### Data Breach Response
1. **Immediate containment** (within 1 hour)
   - Revoke compromised credentials
   - Block suspicious IP addresses
   - Isolate affected systems

2. **Assessment** (within 4 hours)
   - Determine what data was accessed
   - Identify affected users
   - Estimate breach scope

3. **Notification** (within 72 hours for GDPR)
   - Notify affected users
   - Report to authorities if required
   - Update compliance team

4. **Evidence preservation**
   - Save all logs
   - Document timeline
   - Preserve system state

## Contact Information
- Incident Commander: [Phone]
- Security Lead: [Phone]
- Compliance Lead: [Phone]
- Legal: [Phone]
- PR/Communications: [Phone]

## Tools
- AWS Console: https://console.aws.amazon.com
- CloudTrail: https://console.aws.amazon.com/cloudtrail
- Vanta: https://app.vanta.com
- Status Page: [URL]
```

### 2. GDPR Deletion Runbook

**File:** `docs/runbooks/gdpr-deletion.md`

```markdown
# GDPR Data Deletion Runbook

## Purpose
Process user data deletion requests per GDPR Article 17 (Right to Erasure).

## SLA
Complete deletion within 30 days of verified request.

## Prerequisites
- User identity verified
- Request authenticated
- Legal review completed (if applicable)

## Deletion Procedure

### Step 1: Verify Request
```bash
# Verify user identity
# Check request authenticity
# Confirm no legal hold
```

### Step 2: Export Data (Optional)
```bash
# If user requested data portability first
node scripts/export-user-data.js --user-id=USER_ID --output=export.json
```

### Step 3: Database Deletion
```sql
-- Start transaction
BEGIN;

-- Delete user data (cascade)
DELETE FROM users WHERE id = 'USER_ID';

-- Verify deletion
SELECT COUNT(*) FROM users WHERE id = 'USER_ID';
-- Expected: 0

-- Commit
COMMIT;
```

### Step 4: Object Storage Deletion
```bash
# Delete from Cloudflare R2
aws s3 rm s3://smesec-prod-data/users/USER_ID/ --recursive --endpoint-url=https://ACCOUNT_ID.r2.cloudflarestorage.com

# Verify deletion
aws s3 ls s3://smesec-prod-data/users/USER_ID/ --endpoint-url=https://ACCOUNT_ID.r2.cloudflarestorage.com
# Expected: empty
```

### Step 5: Backup Deletion
```bash
# Mark for deletion in backups
# Note: Backups may retain data for retention period
# Document this in deletion confirmation
```

### Step 6: Third-Party Services
```bash
# Delete from analytics (if PII stored)
# Delete from email service
# Delete from support system
```

### Step 7: Verification
```bash
# Run verification script
node scripts/verify-user-deletion.js --user-id=USER_ID

# Expected output:
# ✓ Database: No records found
# ✓ R2 Storage: No files found
# ✓ Analytics: User anonymized
# ✓ Email service: User removed
```

### Step 8: Documentation
```bash
# Log deletion in compliance system
# Send confirmation to user
# Update deletion log
```

## Deletion Log
Maintain log at: `logs/gdpr-deletions.csv`

```csv
timestamp,user_id,request_date,completion_date,verified_by
2026-05-27T10:00:00Z,user_123,2026-05-20,2026-05-27,admin@smesec.com
```

## Exceptions
- Legal hold: Cannot delete
- Active contract: Consult legal
- Financial records: Retain per regulations (7 years)
```

### 3. Backup & Recovery Runbook

**File:** `docs/runbooks/backup-recovery.md`

```markdown
# Backup & Recovery Runbook

## Backup Schedule
- **Database:** Daily automated backups (7-day retention)
- **Object Storage:** Versioning enabled (30-day retention)
- **Configuration:** Git-tracked (infinite retention)

## Recovery Time Objective (RTO)
- Critical systems: 4 hours
- Non-critical systems: 24 hours

## Recovery Point Objective (RPO)
- Database: 24 hours (daily backups)
- Object storage: Real-time (versioning)

## Database Restore Procedure

### Step 1: Identify Backup
```bash
# List available backups
aws rds describe-db-snapshots \
  --db-instance-identifier smesec-prod-db \
  --query 'DBSnapshots[*].[DBSnapshotIdentifier,SnapshotCreateTime]' \
  --output table
```

### Step 2: Restore to New Instance
```bash
# Restore snapshot to new instance
aws rds restore-db-instance-from-db-snapshot \
  --db-instance-identifier smesec-prod-db-restored \
  --db-snapshot-identifier smesec-prod-db-snapshot-2026-05-27 \
  --db-subnet-group-name smesec-db-subnet-group \
  --publicly-accessible false

# Wait for instance to be available
aws rds wait db-instance-available \
  --db-instance-identifier smesec-prod-db-restored
```

### Step 3: Verify Data
```bash
# Connect and verify
psql -h smesec-prod-db-restored.xxxxx.us-east-1.rds.amazonaws.com \
     -U smesec_admin \
     -d smesec

# Run verification queries
SELECT COUNT(*) FROM users;
SELECT MAX(created_at) FROM users;
```

### Step 4: Cutover (if needed)
```bash
# Update application configuration
# Point to restored instance
# Monitor for issues
```

## Object Storage Restore

### Restore Deleted File
```bash
# List versions
aws s3api list-object-versions \
  --bucket smesec-prod-data \
  --prefix users/user_123/document.pdf

# Restore specific version
aws s3api copy-object \
  --bucket smesec-prod-data \
  --copy-source smesec-prod-data/users/user_123/document.pdf?versionId=VERSION_ID \
  --key users/user_123/document.pdf
```

## Testing
- Monthly restore test
- Document test results
- Update runbook based on learnings
```

### 4. Offboarding Runbook

**File:** `docs/runbooks/employee-offboarding.md`

```markdown
# Employee Offboarding Runbook

## Timeline
Complete within 24 hours of departure notification.

## Checklist

### Immediate (Within 1 hour)
- [ ] Disable AWS IAM Identity Center access
- [ ] Revoke GitHub access
- [ ] Disable email account
- [ ] Revoke VPN access (if applicable)
- [ ] Disable Slack account
- [ ] Revoke Vanta access

### Same Day
- [ ] Remove from all AWS IAM groups
- [ ] Remove from GitHub teams
- [ ] Remove from shared drives
- [ ] Collect company devices
- [ ] Reset shared passwords they had access to
- [ ] Review and transfer ownership of resources

### Within 1 Week
- [ ] Conduct exit interview
- [ ] Document knowledge transfer
- [ ] Update documentation with new owners
- [ ] Archive employee files
- [ ] Update compliance records in Vanta

## AWS Access Revocation
```bash
# Disable IAM Identity Center user
aws identitystore delete-user \
  --identity-store-id d-xxxxxxxxxx \
  --user-id USER_ID

# List and remove from all groups
aws identitystore list-group-memberships-for-member \
  --identity-store-id d-xxxxxxxxxx \
  --member-id USER_ID

# Verify no active sessions
aws sts get-caller-identity
```

## GitHub Access Revocation
```bash
# Via GitHub UI
# Organization → People → [User] → Remove from organization

# Or via API
curl -X DELETE \
  -H "Authorization: token GITHUB_TOKEN" \
  https://api.github.com/orgs/smesec/members/USERNAME
```

## Verification
```bash
# Run verification script
node scripts/verify-offboarding.js --employee-id=EMP_ID

# Expected output:
# ✓ AWS: No active access
# ✓ GitHub: Removed from organization
# ✓ Email: Disabled
# ✓ Slack: Deactivated
# ✓ Vanta: Access revoked
```

## Documentation
- Update team roster
- Update CODEOWNERS file
- Update on-call rotation
- Update compliance documentation
```

## Validation & Testing

### Security Validation Checklist

```bash
# Run comprehensive security check
./scripts/security-audit.sh

# Expected checks:
# ✓ All S3 buckets encrypted
# ✓ All RDS instances in private subnets
# ✓ CloudTrail enabled and logging
# ✓ MFA enabled for all users
# ✓ No public S3 buckets
# ✓ Security groups follow least privilege
# ✓ No hardcoded secrets in code
# ✓ All dependencies up to date
```

### Compliance Validation

```bash
# Vanta compliance check
# Login to Vanta dashboard
# Review compliance score (target: >95%)
# Address any findings
```

---

**Document Owner:** Quách Thanh Bình  
**Last Updated:** 2026-05-27  
**Next Review:** 2026-08-27
