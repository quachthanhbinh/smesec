# 01 Architecture Rules

## Rule: Maintain clear separation of concerns

**When:** Designing or modifying system components

**What:** Keep business logic, data access, and presentation layers separate; avoid mixing concerns within a single module

**Why:** Separation of concerns improves testability, maintainability, and allows independent evolution of components

**How to verify:** Check that modules have single, well-defined responsibilities; verify that data access code doesn't contain business logic and vice versa

**Example (correct):**
```python
# Clear separation: repository handles data access
class UserRepository:
    def get_user(self, user_id: int) -> User:
        return db.query(User).filter_by(id=user_id).first()

# Service handles business logic
class UserService:
    def __init__(self, repo: UserRepository):
        self.repo = repo
    
    def activate_user(self, user_id: int) -> bool:
        user = self.repo.get_user(user_id)
        user.active = True
        return True
```

**Example (incorrect):**
```python
# Mixed concerns: data access and business logic together
class UserManager:
    def activate_user(self, user_id: int) -> bool:
        user = db.query(User).filter_by(id=user_id).first()
        user.active = True
        # Business logic mixed with data access
        if user.subscription_expired():
            send_renewal_email(user)
        return True
```
