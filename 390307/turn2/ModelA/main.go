package main

import (
    "context"
    "fmt"
    "net/http"
    "time"
)

// Define constants for different user roles and fetch rate limits accordingly
const (
    AdminRole  = "admin"
    UserRole   = "user"
    GuestRole  = "guest"
    DefaultLimit = 5
    LimitPerMin = time.Minute
)

// Simple data structure to hold user information
type User struct {
    ID   string
    Role string
}

// RateLimiter to manage requests counts per time window.
type RateLimiter struct {
    mu         sync.Mutex
    Requests    map[string]*rateLimitEntry
    LimitMap    map[string]int
    Window      time.Duration
}

func NewRateLimiter(window time.Duration) *RateLimiter {
    return &RateLimiter{
        Requests: make(map[string]*rateLimitEntry),
        LimitMap: make(map[string]int),
        Window:   window,
    }
}

type rateLimitEntry struct {
    count int
    end   time.Time
}

func (r *RateLimiter) GetLimit(role string) int {
    r.mu.Lock()
    defer r.mu.Unlock()
    if limit, exists := r.LimitMap[role]; exists {
        return limit
    }
    return DefaultLimit
}

func (r *RateLimiter) Allow(ctx context.Context, user *User) bool {
    r.mu.Lock()
    defer r.mu.Unlock()

    key := user.ID
    entry, exists := r.Requests[key]

    if !exists {
        entry = &rateLimitEntry{count: 1, end: time.Now().Add(r.Window)}
        r.Requests[key] = entry
        return true
    }

    if time.Now().After(entry.end) {
        entry.count = 1
        entry.end = time.Now().Add(r.Window)
        r.Requests[key] = entry
        return true
    }

    if entry.count < r.GetLimit(user.Role) {
        entry.count++
        return true
    }

    return false
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
        limiter := NewRateLimiter(LimitPerMin)

        if !limiter.Allow(ctx, user) {
            http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
            return
        }

        // Process next step
        next.ServeHTTP(w, r)
    })
}

// Usage Example 

func main() {
    // Create an example handler that will always operate.
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        user := r.Context().Value("user").(*User)
        fmt.Fprintf(w, "Hello, %s!", user.Role)
    })

    // Wrap example handler in the rate limiting middleware.
    http.Handle("/", rateLimitMiddleware(handler))

    // Example API endpoint to serve user info and verify function.
    http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        user := &User{ID: "123", Role: UserRole}
        ctx := context.WithValue(context.Background(), "user", user)
        http.HandlerFunc(handler).ServeHTTP(w, r.WithContext(ctx))
    })

    fmt.Println("Server started on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Server error:", err)
    }
}