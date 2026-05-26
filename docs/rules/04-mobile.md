# 04 Mobile Development Rules

## Rule: Handle network failures and offline scenarios

**When:** Implementing features that require network connectivity

**What:** Gracefully handle network failures, timeouts, and offline states; provide appropriate feedback and fallback behavior

**Why:** Mobile networks are unreliable; users expect apps to work in degraded conditions

**How to verify:** Test with airplane mode, slow connections, and intermittent connectivity; verify that app doesn't crash and provides useful feedback

**Example (correct):**
```kotlin
suspend fun fetchUserData(userId: String): Result<User> {
    return try {
        val response = api.getUser(userId)
        Result.success(response)
    } catch (e: IOException) {
        // Network error - try cache
        val cached = cache.getUser(userId)
        if (cached != null) {
            Result.success(cached)
        } else {
            Result.failure(NetworkException("No connection and no cached data"))
        }
    }
}
```

**Example (incorrect):**
```kotlin
suspend fun fetchUserData(userId: String): User {
    return api.getUser(userId)  // Crashes on network failure
}
```
