package main

import (
    "context"
    "fmt"
    "net/http"
    "sync"
    "time"
)

// Define constants for different user roles and fetch rate limits accordingly
const (
    AdminRole = "admin"
    UserRole  = "user"
    GuestRole = "guest"
)

// Simple data structure to hold user information
type User struct {
    ID   string
    Role string
}

// RateLimiter to manage requests counts per time window.
type RateLimiter struct {
    mu       sync.Mutex
    Requests map[string]int
    Limit    int
    Window   time.Duration
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
    return &RateLimiter{
        Requests: make(map[string]int),
        Limit:    limit,
        Window:   window,
    }
}

// Rate limit checking function.
func (r *RateLimiter) Allow(ctx context.Context, user *User) bool {
    r.mu.Lock()
    defer r.mu.Unlock()

    key := user.ID

    if _, exists := r.Requests[key]; !exists {
        r.Requests[key] = 1
        go r.cleanup(key)
        return true
    }

    if r.Requests[key] < r.Limit {
        r.Requests[key]++
        return true
    }

    return false
}

// Clean resources beyond the time window.
func (r *RateLimiter) cleanup(key string) {
    time.Sleep(r.Window)
    r.mu.Lock()
    defer r.mu.Unlock()
    delete(r.Requests, key)
}

// Middleware getting the user info from context and applying rate rules
func rateLimitMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx := r.Context()

        user, ok := ctx.Value("user").(*User)
        if !ok || user == nil {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // Restrict dynamically based on role
        limiter := getUserRateLimiter(user.Role)

        if !limiter.Allow(ctx, user) {
            http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
            return
        }

        // Process next step
        next.ServeHTTP(w, r)
    })
}

// Helper function to configure limiters based on user roles
func getUserRateLimiter(role string) *RateLimiter {
    var (
        limit  int
        window time.Duration = time.Minute
    )
    switch role {
    case AdminRole:
        limit = 100
    case UserRole:
        limit = 10
    case GuestRole:
        limit = 5
    default:
        limit = 1
    }
    return NewRateLimiter(limit, window)
}

// Usage Example 

func main() {
    // Create an example handler that will always operate.
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        user := r.Context().Value("user").(*User)
        fmt.Fprintf(w, "Hello, %s!", user.Role)
    })
    
    // Wrap example handler in the rate limiting middleware.
    http.Handle("/", RateLimitMiddleware(handler))
    
    exampleUser := &User{ID: "123", Role: UserRole}
    ctx := context.WithValue(context.Background(), "user", exampleUser)

    http.ListenAndServe(":8080", http.DefaultServeMux)
}

// RateLimitMiddleware refers back to the created `rateLimitMiddleware` to handle flows via injected user scope
