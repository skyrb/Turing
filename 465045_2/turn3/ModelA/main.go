package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	mu       sync.Mutex
	limits    map[string]*rateLimit
	retries   map[string]int
	retryDelay time.Duration
}

type rateLimit struct {
	limit     int
	remaining int
	reset     time.Time
}

func NewRateLimiter(retryDelay time.Duration) *RateLimiter {
	return &RateLimiter{
		limits:    make(map[string]*rateLimit),
		retries:   make(map[string]int),
		retryDelay: retryDelay,
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

func (rl *RateLimiter) Retry(role string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	retries, ok := rl.retries[role]
	if !ok {
		rl.retries[role] = 3 // Default retry count for new roles
	}

	if retries < 3 {
		rl.retries[role]++
		return true
	}
	return false
}

func (rl *RateLimiter) ResetRetry(role string) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	rl.retries[role] = 0
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
				// Simulate retry logic and request processing
				for i := 0; i < 3; i++ {
					if ok, _ := handleSimulatedRequest(); ok {
						rl.ResetRetry(role)
						fmt.Println("Request handled:", role)
						return
					}
					if i < 2 && rl.Retry(role) {
						fmt.Println("Retrying request:", role)
						time.Sleep(rl.retryDelay)
					} else {
						fmt.Println("Request failed after retries:", role)
						rl.ResetRetry(role)
						return
					}
				}
			} else {
				fmt.Println("Request rejected due to rate limit:", role)
				return
			}
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

func handleSimulatedRequest() (bool, error) {
	if err := simulateAPIRequestError(); err != nil {
		return false, err
	}
	return true, nil
}

func simulateAPIRequestError() error {
	if time.Now().Second()%5 == 0 {
		return fmt.Errorf("simulated API request error")
	}
	return nil
}

func main() {
	rl := NewRateLimiter(time.Second)
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