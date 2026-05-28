---
name: incident-playbooks-technical-advisor
description: "Technical Advisor for Incident Playbooks (Requirement 5). Extends base technical-advisor agent with specialized context for Step Functions workflows, playbook state machines, and mobile app architecture."
extends: technical-advisor
tools: Read, Glob, Grep, WebSearch, WebFetch
---

**Base Agent**: This agent extends [technical-advisor](../../../.github/agents/technical-advisor.agent.md). Follow all base agent instructions, debate structure, and frameworks. This file adds requirement-specific context only.

---

## Requirement 5: Incident Playbooks

### Scope
- **Playbook engine**: Step Functions state machines for workflow orchestration
- **Pre-built playbooks**: 4 playbooks (account compromise, unauthorized access, shadow IT, offboarding emergency)
- **Mobile/Desktop app**: Flutter app for incident response on mobile
- **Notification system**: Slack, email, push notifications (SNS)
- **Audit trail**: All playbook executions logged to S3

### Key Technical Challenges

1. **Step Functions Complexity**
   - State machine design (sequential, parallel, choice states)
   - Error handling and retries
   - Human-in-the-loop (wait for user input)
   - Timeout handling (playbook expires after 24 hours)

2. **Mobile App Architecture**
   - Flutter app (iOS + Android)
   - Push notifications (Firebase Cloud Messaging)
   - Offline support (local state, sync when online)
   - Deep linking (open specific playbook from notification)

3. **Playbook State Management**
   - Current step tracking
   - User input persistence
   - Rollback capability (undo previous step)
   - Resume capability (continue from where left off)

### Playbook Architecture

**Step Functions state machine example (Account Compromise):**
```json
{
  "Comment": "Account Compromise Playbook",
  "StartAt": "NotifyITAdmin",
  "States": {
    "NotifyITAdmin": {
      "Type": "Task",
      "Resource": "arn:aws:lambda:us-east-1:123456789012:function:SendNotification",
      "Parameters": {
        "channel": "slack",
        "message": "Account compromise detected for user ${userId}"
      },
      "Next": "WaitForConfirmation"
    },
    "WaitForConfirmation": {
      "Type": "Task",
      "Resource": "arn:aws:states:::lambda:invoke.waitForTaskToken",
      "Parameters": {
        "FunctionName": "WaitForUserInput",
        "Payload": {
          "taskToken.$": "$$.Task.Token",
          "question": "Confirm account compromise?"
        }
      },
      "TimeoutSeconds": 3600,
      "Next": "DisableAccount"
    },
    "DisableAccount": {
      "Type": "Parallel",
      "Branches": [
        {
          "StartAt": "DisableGoogle",
          "States": {
            "DisableGoogle": {
              "Type": "Task",
              "Resource": "arn:aws:lambda:us-east-1:123456789012:function:DisableGoogleAccount",
              "End": true
            }
          }
        },
        {
          "StartAt": "DisableM365",
          "States": {
            "DisableM365": {
              "Type": "Task",
              "Resource": "arn:aws:lambda:us-east-1:123456789012:function:DisableM365Account",
              "End": true
            }
          }
        }
      ],
      "Next": "GenerateReport"
    },
    "GenerateReport": {
      "Type": "Task",
      "Resource": "arn:aws:lambda:us-east-1:123456789012:function:GenerateIncidentReport",
      "End": true
    }
  }
}
```

### Mobile App Architecture (Flutter)

**Key features:**
- Push notifications (FCM)
- Offline support (local SQLite database)
- Deep linking (open playbook from notification)
- Biometric authentication (Face ID, Touch ID)

**Performance requirements:**
- App launch time: <2 seconds
- Playbook load time: <1 second
- Push notification delivery: <5 seconds

### Security Requirements
- Playbook execution: Requires MFA (biometric on mobile)
- Audit trail: All steps logged to S3 (immutable, 7-year retention)
- Notification encryption: TLS 1.3 in transit
- Mobile app: Code signing, certificate pinning
