package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// RateLimit represents the rate limit for a user role
type RateLimit struct {
	limit  int   // Maximum requests allowed
	current int   // Current number of requests
	last    int64 // Last request time in seconds
	mutex   sync.Mutex
}

// IsAllowed checks if the user role can make another request
func (rl *RateLimit) IsAllowed() bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	now := time.Now().Unix()
	if now-rl.last >= 60 { // Reset the limit every minute
		rl.current = 0
		rl.last = now
	}

	if rl.current < rl.limit {
		rl.current++
		return true
	}
	return false
}

// Decrease decrements the current request count
func (rl *RateLimit) Decrease() {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()
	rl.current--
}

// WaitGroupWrapper is a wrapper around sync.WaitGroup to manage Goroutines
type WaitGroupWrapper struct {
	wg *sync.WaitGroup
}

// Add adds one to the wait group counter
func (wgw *WaitGroupWrapper) Add() {
	wgw.wg.Add(1)
}

// Done decrements the wait group counter and signals it's done
func (wgw *WaitGroupWrapper) Done() {
	wgw.wg.Done()
}

// Wait waits for all Goroutines to finish
func (wgw *WaitGroupWrapper) Wait() {
	wgw.wg.Wait()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Define rate limits for different user roles
	rateLimits := map[string]*RateLimit{
		"admin":    &RateLimit{limit: 100, current: 0, last: 0},
		"user":     &RateLimit{limit: 20, current: 0, last: 0},
		"guest":    &RateLimit{limit: 10, current: 0, last: 0},
	}

	// Create a wait group wrapper
	wgw := &WaitGroupWrapper{wg: &sync.WaitGroup{}}

	// Function to simulate a request
	handleRequest := func(role string) {
		wgw.Add()
		defer wgw.Done()

		if rateLimits[role].IsAllowed() {
			fmt.Printf("Handling request from %s\n", role)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			rateLimits[role].Decrease()
		} else {
			fmt.Printf("Rate limit exceeded for %s\n", role)
		}
	}

	// Simulate multiple requests from different user roles
	for i := 0; i < 100; i++ {
		role := []string{"admin", "user", "guest"}[rand.Intn(3)]
		go handleRequest(role)
	}

	// Wait for all Goroutines to finish
	wgw.Wait()
}