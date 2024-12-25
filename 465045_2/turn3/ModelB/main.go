package main  
import (  
    "context"
    "errors"
    "fmt"
    "math/rand"
    "sync"
    "time"
)

const (
    maxRetries     = 3
    retryDelay    = 100 * time.Millisecond
    maxBackoffDelay = 5 * time.Second
)

type RateLimiter struct {
    mu     sync.Mutex
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

func (rl *RateLimiter) HandleRequest(ctx context.Context, role string, requestCh chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()

    for {
        select {
        case <-ctx.Done():
            return
        default:
            if rl.Allow(role) {
                // Simulate request processing
                time.Sleep(time.Millisecond * 100)
                if rand.Intn(10) < 3 { // Simulate random failures
                    fmt.Println("Request failed:", role)
                    requestCh <- 1 // Send request failure to channel
                    continue
                }
                fmt.Println("Request handled:", role)
                return
            }
            fmt.Println("Request rejected due to rate limit:", role)
        }
    }
}

func (rl *RateLimiter) UpdateRateLimit(role string, newLimit int) {
    rl.mu.Lock()
    defer rl.mu.Unlock()

    limit, ok := rl.limits[role]
    if !ok {
        limit = &rateLimit{}
        rl.limits[role] = limit
    }
    limit.limit = newLimit
    limit.remaining = newLimit
}

func handleFailedRequests(ctx context.Context, rl *RateLimiter, requestCh <-chan int) {
    for {
        select {
        case <-ctx.Done():
            return
        case <-requestCh:
            // Handle failed request
            // You can retry the request here based on your retry strategy
            time.Sleep(retryDelay)
        }
    }
}

func main() {
    rl := NewRateLimiter()
    var wg sync.WaitGroup

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    requestCh := make(chan int)
    go handleFailedRequests(ctx, rl, requestCh)

    // Simulate dynamic rate limit changes
    go func() {
        for {
            rl.UpdateRateLimit("admin", 5)  // Update rate limit for "admin" role
            time.Sleep(time.Second * 5)
            rl.UpdateRateLimit("admin", 10) // Reset rate limit for "admin" role
            time.Sleep(time.Second * 5)