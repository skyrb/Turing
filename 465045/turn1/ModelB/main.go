package main

import (
    "fmt"
    "sync"
    "time"
)

// UserRole represents a user's role with an associated rate limit.
type UserRole struct {
    Name      string
    RateLimit int // Requests per second
}

// RateLimiter manages API rate limits for different user roles.
type RateLimiter struct {
    limits  map[string]*UserRole
    mutex   sync.Mutex
    wg      sync.WaitGroup
    done    chan struct{}
    requests chan struct{}
}

// NewRateLimiter creates a new RateLimiter with default user roles.
func NewRateLimiter() *RateLimiter {
    rl := &RateLimiter{
        limits: map[string]*UserRole{
            "user":  &UserRole{Name: "user", RateLimit: 2},
            "admin": &UserRole{Name: "admin", RateLimit: 10},
        },
        done:    make(chan struct{}),
        requests: make(chan struct{}),
    }
    return rl
}

// Start begins the rate limiting process.
func (rl *RateLimiter) Start() {
    rl.wg.Add(1)
    go rl.run()
}

// Stop stops the rate limiting process.
func (rl *RateLimiter) Stop() {
    close(rl.done)
    rl.wg.Wait()
}

// LimitUser applies the rate limit for the specified user role to the current request.
func (rl *RateLimiter) LimitUser(role string) {
    rl.wg.Add(1)
    go rl.limitUserGoroutine(role)
}

func (rl *RateLimiter) limitUserGoroutine(role string) {
    defer rl.wg.Done()

    rl.mutex.Lock()
    limit, ok := rl.limits[role]
    rl.mutex.Unlock()

    if !ok {
        fmt.Println("Invalid user role:", role)
        return
    }

    ticker := time.NewTicker(time.Second / time.Duration(limit.RateLimit))
    defer ticker.Stop()

    for {
        select {
        case <-rl.done:
            return
        case <-rl.requests:
            <-ticker.C
        }
    }
}

// AddRequest indicates a new API request.
func (rl *RateLimiter) AddRequest() {
    rl.requests <- struct{}{}
}

func main() {
    rl := NewRateLimiter()
    rl.Start()
    defer rl.Stop()

    go func() {
        for {
            rl.LimitUser("user")
            rl.LimitUser("admin")
            time.Sleep(time.Second * 5)
        }
    }()

    for i := 0; i < 10; i++ {
        rl.AddRequest()
        fmt.Println("API request processed.")
        time.Sleep(time.Second)
    }
}