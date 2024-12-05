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

// RateLimiter manages requests counts per time window.
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

// Allow checks if a request can be processed.
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

// cleanup removes old request records.
func (r *RateLimiter) cleanup(key string) {
    time.Sleep(r.Window)
    r.mu.Lock()
    defer r.mu.Unlock()
    delete(r.Requests, key)
}

// RateLimitMiddleware applies rate limiting based on user roles.
func RateLimitMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx := r.Context()

        user, ok := ctx.Value("user").(*User)
        if !ok || user == nil {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // Determine rate limits based on the user's role
        limiter := getUserRateLimiter(user.Role)

        if !limiter.Allow(ctx, user) {
            http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
            return
        }

        // Allow the request to proceed
        next.ServeHTTP(w, r)
    })
}

// getUserRateLimiter returns a RateLimiter configured for a specific role.
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

func main() {
    // Create an example handler that will always operate.
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        user, _ := r.Context().Value("user").(*User)
        fmt.Fprintf(w, "Hello, %s!", user.Role)
    })

    // A sample user to simulate authenticated user data
    exampleUser := &User{ID: "123", Role: UserRole}
    ctx := context.WithValue(context.Background(), "user", exampleUser)

    // Create a new server mux and use the RateLimitMiddleware
    mux := http.NewServeMux()
    mux.Handle("/", RateLimitMiddleware(handler))

    // Create a server that doesn't overwrite the user context (simulating that users are passed to handlers)
    server := &http.Server{
        Addr:    ":8080",
        Handler: contextClearMiddleware(ctx, mux),
    }

    // Run the server
    fmt.Println("Server starting at :8080")
    if err := server.ListenAndServe(); err != nil {
        fmt.Printf("Failed to start server: %v\n", err)
    }
}

// contextClearMiddleware helps to re-inject context on each request with preserved user 'cookie'.
func contextClearMiddleware(baseCtx context.Context, h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        r = r.WithContext(baseCtx)
        h.ServeHTTP(w, r)
    })
}