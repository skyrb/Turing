package main  
import (  
    "fmt"
    "sync"
    "time"
)

type RateLimiter struct {
    mu sync.Mutex
    limits map[string]*rateLimit
}

type rateLimit struct {
    limit     int
    remaining int
    reset     time.Time
}

func NewRateLimiter() *RateLimiter {
    return &RateLimiter{
        limits: make(map[string]*rateLimit),
    }
}

func (rl *RateLimiter) Allow(role string) bool {
    rl.mu.Lock()
    defer rl.mu.Unlock()

    limit, ok := rl.limits[role]
    if !ok {
        // Initialize rate limit for the role if it doesn't exist
        limit = &rateLimit{
            limit:     10, // Default rate limit for new roles
            remaining: 10,
            reset:     time.Now().Add(time.Second),
        }
        rl.limits[role] = limit
    }

    if time.Now().After(limit.reset) {
        // Reset rate limit if it has expired
        limit.remaining = limit.limit
        limit.reset = time.Now().Add(time.Second)
    }

    if limit.remaining > 0 {
        limit.remaining--
        return true
    }
    return false
}

func (rl *RateLimiter) HandleRequest(role string, wg *sync.WaitGroup) {
    defer wg.Done()

    if !rl.Allow(role) {
        fmt.Println("Request rejected due to rate limit:", role)
        return
    }

    // Simulate request processing
    time.Sleep(time.Millisecond * 100)
    fmt.Println("Request handled:", role)
}

func main() {
    rl := NewRateLimiter()
    var wg sync.WaitGroup

    wg.Add(100)
    for i := 0; i < 100; i++ {
        role := "admin" // Assign different roles for each request
        go rl.HandleRequest(role, &wg)
    }

    wg.Wait()
    fmt.Println("All requests handled.")
}