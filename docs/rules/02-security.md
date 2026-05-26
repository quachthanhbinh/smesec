# 02 Security Rules

## Rule: Validate and sanitize all external input

**When:** Processing data from users, APIs, files, or any external source

**What:** Validate input against expected format and constraints; sanitize data before use in queries, commands, or output

**Why:** Prevents injection attacks, data corruption, and security vulnerabilities

**How to verify:** Check that all external inputs have validation; verify that SQL queries use parameterization; confirm that shell commands don't use unsanitized input

**Example (correct):**
```python
# Parameterized query prevents SQL injection
def get_user_by_email(email: str) -> User:
    if not re.match(r'^[\w\.-]+@[\w\.-]+\.\w+$', email):
        raise ValueError("Invalid email format")
    return db.execute(
        "SELECT * FROM users WHERE email = ?", 
        (email,)
    ).fetchone()
```

**Example (incorrect):**
```python
# String concatenation vulnerable to SQL injection
def get_user_by_email(email: str) -> User:
    query = f"SELECT * FROM users WHERE email = '{email}'"
    return db.execute(query).fetchone()
```
