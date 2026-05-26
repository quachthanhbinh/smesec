# 06 Testing Rules

## Rule: Write tests that verify behavior, not implementation

**When:** Writing unit, integration, or end-to-end tests

**What:** Test the observable behavior and outcomes of code, not internal implementation details

**Why:** Implementation-focused tests are brittle and break during refactoring; behavior-focused tests provide confidence while allowing code evolution

**How to verify:** Check that tests don't mock internal methods; verify that tests would pass with alternative implementations; confirm tests describe what the code does, not how

**Example (correct):**
```python
def test_user_activation_sends_welcome_email():
    # Test behavior: activating user sends email
    user = User(email="test@example.com", active=False)
    service.activate_user(user.id)
    
    assert user.active is True
    assert len(email_service.sent_emails) == 1
    assert email_service.sent_emails[0].to == "test@example.com"
```

**Example (incorrect):**
```python
def test_user_activation_calls_internal_method():
    # Test implementation: checking internal method calls
    user = User(email="test@example.com", active=False)
    with patch.object(service, '_set_active_flag') as mock:
        service.activate_user(user.id)
        mock.assert_called_once_with(user.id, True)
```
