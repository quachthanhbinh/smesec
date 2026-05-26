# 00 Universal Rules

## Rule: Follow established patterns

**When:** Writing any code in the project

**What:** Maintain consistency with existing code patterns, naming conventions, and architectural decisions

**Why:** Consistency reduces cognitive load, makes code easier to understand, and prevents fragmentation

**How to verify:** Compare new code against similar existing implementations; check that naming and structure match project conventions

**Example (correct):**
```python
# Following project's snake_case convention
def validate_user_input(data: dict) -> bool:
    return all(key in data for key in ["name", "email"])
```

**Example (incorrect):**
```python
# Mixing camelCase in a snake_case project
def validateUserInput(data: dict) -> bool:
    return all(key in data for key in ["name", "email"])
```
