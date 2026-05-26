# 03 Web Development Rules

## Rule: Implement proper error handling and user feedback

**When:** Building web endpoints, forms, or interactive features

**What:** Handle errors gracefully; provide clear, actionable feedback to users; log errors for debugging

**Why:** Improves user experience, aids debugging, and prevents information leakage through error messages

**How to verify:** Test error paths; verify that users see helpful messages; confirm that sensitive details aren't exposed in production errors

**Example (correct):**
```python
@app.route('/api/users', methods=['POST'])
def create_user():
    try:
        data = request.get_json()
        user = UserService.create(data)
        return jsonify(user.to_dict()), 201
    except ValidationError as e:
        return jsonify({"error": "Invalid input", "details": str(e)}), 400
    except Exception as e:
        logger.error(f"User creation failed: {e}")
        return jsonify({"error": "An error occurred"}), 500
```

**Example (incorrect):**
```python
@app.route('/api/users', methods=['POST'])
def create_user():
    data = request.get_json()
    user = UserService.create(data)  # No error handling
    return jsonify(user.to_dict()), 201
```
