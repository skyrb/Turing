package main

import (
	"fmt"
	"sync"
	"time"
)

type UserRole int
const (
	Premium UserRole = 0
	Standard UserRole = 1
	Free UserRole = 2
)

var (
	// Define rate limits per second for each user role
	rateLimits = map[UserRole]int{
		Premium:   10,
		Standard:  5,
		Free:      2,
	}
)

// mockAPIRequest simulates an API request
func mockAPIRequest(userRole UserRole) {
	fmt.Printf("Making an API request for user role %d\n", userRole)
	time.Sleep(time.Millisecond * 50)
}

// simulateUserRequests makes concurrent API requests adhering to the rate limits
func simulateUserRequests(userRole UserRole, numRequests int, wg *sync.WaitGroup) {
	defer wg.Done()

	limits := rateLimits[userRole]
	if limits == 0 {
		fmt.Printf("User role %d has no rate limit.\n", userRole)
		return
	}

	reqQueue := make(chan struct{}, limits)

	for i := 0; i < numRequests; i++ {
		select {
		case reqQueue <- struct{}{}:
			go func() {
				mockAPIRequest(userRole)
				reqQueue <- struct{}{}
			}()
		default:
			time.Sleep(time.Duration(1/float64(limits)) * time.Second)
			reqQueue <- struct{}{}
			go func() {
				mockAPIRequest(userRole)
				reqQueue <- struct{}{}
			}()
		}
	}
}

func main() {
	wg := sync.WaitGroup{}

	// Simulate requests for different user roles
	wg.Add(1)
	go simulateUserRequests(Premium, 20, &wg)

	wg.Add(1)
	go simulateUserRequests(Standard, 15, &wg)

	wg.Add(1)
	go simulateUserRequests(Free, 8, &wg)

	wg.Wait()
	fmt.Println("All requests completed.")
}