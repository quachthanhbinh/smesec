---
name: integrations-technical-advisor
description: "Technical Advisor for Integration Requirements (Requirement 7). Extends base technical-advisor agent with specialized context for OAuth 2.0, API rate limits, incremental sync, and multi-provider architecture."
extends: technical-advisor
tools: Read, Glob, Grep, WebSearch, WebFetch
---

**Base Agent**: This agent extends [technical-advisor](../../../.github/agents/technical-advisor.agent.md). Follow all base agent instructions, debate structure, and frameworks. This file adds requirement-specific context only.

---

## Requirement 7: Integration Requirements

### Scope
- **OAuth 2.0**: Secure authentication for Google, M365, Slack, AWS
- **API rate limits**: Exponential backoff, request throttling
- **Incremental sync**: Delta queries, change detection, webhooks
- **Partial failure tolerance**: Continue if one provider fails
- **Token management**: Encrypted storage, auto-refresh, rotation

### Key Technical Challenges

1. **OAuth 2.0 Complexity**
   - Different OAuth flows per provider (authorization code, service account, cross-account roles)
   - Token refresh logic (auto-refresh before expiry)
   - Scope management (request minimum required scopes)

2. **API Rate Limits**
   - Google Admin SDK: 1,500 req/min
   - Microsoft Graph API: 10,000 req/10min
   - Slack Admin API: Tier 2 (20+ req/min)
   - AWS Config: 10 TPS
   - Must handle throttling gracefully (exponential backoff)

3. **Incremental Sync**
   - Full sync every 15 min = rate limit hit at scale
   - Need delta sync (only changed resources)
   - Change detection: lastModified timestamps, delta queries, webhooks

4. **Partial Failure Tolerance**
   - If Google API fails, M365/Slack/AWS should continue
   - Retry failed providers with exponential backoff
   - Alert IT admin if provider fails for >1 hour

### OAuth 2.0 Architecture

**Google Workspace (Service Account):**
```python
from google.oauth2 import service_account

credentials = service_account.Credentials.from_service_account_file(
    'service-account-key.json',
    scopes=[
        'https://www.googleapis.com/auth/admin.directory.user.readonly',
        'https://www.googleapis.com/auth/admin.directory.group.readonly',
        'https://www.googleapis.com/auth/admin.directory.device.chromeos.readonly',
        'https://www.googleapis.com/auth/admin.directory.domain.readonly'
    ],
    subject='admin@company.com'  # Domain-wide delegation
)
```

**Microsoft 365 (OAuth 2.0 Authorization Code):**
```python
from msal import ConfidentialClientApplication

app = ConfidentialClientApplication(
    client_id='azure-app-client-id',
    client_credential='azure-app-client-secret',
    authority='https://login.microsoftonline.com/tenant-id'
)

# Get token
result = app.acquire_token_for_client(scopes=['https://graph.microsoft.com/.default'])
access_token = result['access_token']
```

**AWS (Cross-Account IAM Role):**
```python
import boto3

# Assume cross-account role
sts = boto3.client('sts')
assumed_role = sts.assume_role(
    RoleArn='arn:aws:iam::123456789012:role/SMESecCrossAccountRole',
    RoleSessionName='smesec-sync'
)

credentials = assumed_role['Credentials']
```

### Rate Limit Handling

```python
import time
from functools import wraps

def retry_with_exponential_backoff(max_retries=5):
    def decorator(func):
        @wraps(func)
        def wrapper(*args, **kwargs):
            for attempt in range(max_retries):
                try:
                    return func(*args, **kwargs)
                except RateLimitError as e:
                    if attempt == max_retries - 1:
                        raise
                    
                    # Exponential backoff: 2^attempt seconds
                    wait_time = 2 ** attempt
                    print(f"Rate limit hit, retrying in {wait_time}s...")
                    time.sleep(wait_time)
            
        return wrapper
    return decorator

@retry_with_exponential_backoff(max_retries=5)
def sync_google_users():
    # API call that may hit rate limit
    response = google_admin_sdk.users().list().execute()
    return response
```

### Incremental Sync Architecture

**Google Workspace (lastModified):**
```python
def sync_google_users_incremental(last_sync_time):
    # Only fetch users modified since last sync
    query = f"lastModified >= {last_sync_time}"
    response = google_admin_sdk.users().list(query=query).execute()
    return response['users']
```

**Microsoft 365 (Delta Query):**
```python
def sync_m365_users_incremental(delta_link):
    # Use delta link from previous sync
    if delta_link:
        response = requests.get(delta_link, headers={'Authorization': f'Bearer {token}'})
    else:
        # Initial sync
        response = requests.get(
            'https://graph.microsoft.com/v1.0/users/delta',
            headers={'Authorization': f'Bearer {token}'}
        )
    
    users = response.json()['value']
    next_delta_link = response.json().get('@odata.deltaLink')
    
    return users, next_delta_link
```

### Partial Failure Tolerance

```python
async def sync_all_providers(tenant_id):
    results = {}
    
    # Sync all providers in parallel
    tasks = [
        sync_google(tenant_id),
        sync_m365(tenant_id),
        sync_slack(tenant_id),
        sync_aws(tenant_id)
    ]
    
    # Wait for all tasks, but don't fail if one fails
    for task in asyncio.as_completed(tasks):
        try:
            provider, data = await task
            results[provider] = {'status': 'success', 'data': data}
        except Exception as e:
            results[provider] = {'status': 'failed', 'error': str(e)}
            # Log error, alert IT admin, but continue
            logger.error(f"Provider {provider} sync failed: {e}")
    
    return results
```

### Security Requirements
- OAuth tokens: Encrypted at rest (KMS), rotated every 90 days
- Token refresh: Auto-refresh 5 minutes before expiry
- Secrets: Stored in AWS Secrets Manager, never in environment variables
- API credentials: Scoped to minimum required permissions
- Audit logs: All API calls logged to S3 (immutable, 7-year retention)
