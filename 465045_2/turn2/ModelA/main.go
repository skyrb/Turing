package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	mu         sync.Mutex
	limits     map[string]int // Map of role to maximum requests per second
	current     map[string]int // Map of role to current requests in process
	resetTimes map[string]time.Time
}

func NewRateLimiter(defaultLimit int) *RateLimiter {
	return &RateLimiter{
		limits:     make(map[string]int),
		current:    make(map[string]int),
		resetTimes: make(map[string]time.Time),
	}
}

func (rl *RateLimiter) SetLimit(role string, limit int) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	rl.limits[role] = limit
}

func (rl *RateLimiter) Allow(role string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	limit, ok := rl.limits[role]
	if !ok {
		// Default to a low limit if the role is not defined
		limit = 1
	}

	current, ok := rl.current[role]
	if !ok {
		current = 0
	}

	if current >= limit {
		return false
	}

	rl.current[role] = current + 1
	return true
}

func (rl *RateLimiter) HandleRequest(role string, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		return
	default:
	}

	if !rl.Allow(role) {
		fmt.Println("Request rejected due to rate limit:", role)
		return
	}

	// Simulate request processing
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Request handled:", role)
}

func (rl *RateLimiter) ResetLimits() {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	for role := range rl.current {
		rl.current[role] = 0
	}
}

func main() {
	rl := NewRateLimiter(5) // Default limit of 5 requests per second
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start goroutines to handle requests
	for i := 0; i < 20; i++ {
		role := "admin"
		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("Goroutine exiting due to context cancellation.")
					return
				default:
					wg.Add(1)
					rl.HandleRequest(role, ctx, &wg)
					time.Sleep(time.Second / 10) // Make requests more frequent
				}
			}
		}()
	}

	// Wait for all requests to finish or context cancellation
	wg.Wait()
	cancel()

	// Dynamically adjust rate limit
	rl.SetLimit("admin", 10) // Increase limit to 10 requests per second

	// Reset current request counts
	rl.ResetLimits()

	// Start handling requests again with the new limit
	wg.Add(1)
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine exiting due to context cancellation.")
				return
			default:
				wg.Add(1)
				rl.HandleRequest("admin", ctx, &wg)
				time.Sleep(time.Second / 10) // Make requests more frequent
			}
		}
	}()

	// Wait for all requests to finish or context cancellation
	wg.Wait()
	cancel()

	fmt.Println("All requests handled.")
}