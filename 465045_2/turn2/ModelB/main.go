package main

import (
	"context"
	"fmt"
	"sync"
	"time"
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

func (rl *RateLimiter) HandleRequest(ctx context.Context, role string, wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if rl.Allow(role) {
				// Simulate request processing
				time.Sleep(time.Millisecond * 100)
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

func main() {
	rl := NewRateLimiter()
	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Simulate dynamic rate limit changes
	go func() {
		for {
			rl.UpdateRateLimit("admin", 5)  // Update rate limit for "admin" role
			time.Sleep(time.Second * 5)
			rl.UpdateRateLimit("admin", 10) // Reset rate limit for "admin" role
			time.Sleep(time.Second * 5)
		}
	}()

	wg.Add(100)
	for i := 0; i < 100; i++ {
		role := "admin" // Assign different roles for each request
		go rl.HandleRequest(ctx, role, &wg)
	}

	wg.Wait()
	fmt.Println("All requests handled.")
}  