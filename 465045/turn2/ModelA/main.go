package main

import (
	"fmt"
	"sync"
	"time"
)

type UserRole int

const (
	Premium  UserRole = 0
	Standard UserRole = 1
	Free     UserRole = 2
)

var (
	// Define rate limits per second for each user role
	rateLimits = map[UserRole]int{
		Premium:  10,
		Standard: 5,
		Free:     2,
	}

	// Map to store the current role for each user
	userRoles = map[string]UserRole{}

	// Mutex to ensure thread safety when updating user roles
	roleMutex sync.Mutex
)

// getCurrentRole retrieves the current role for a given user
func getCurrentRole(user string) UserRole {
	roleMutex.Lock()
	defer roleMutex.Unlock()
	return userRoles[user]
}

// updateUserRole updates the role for a given user
func updateUserRole(user string, newRole UserRole) {
	roleMutex.Lock()
	defer roleMutex.Unlock()
	userRoles[user] = newRole
}

// mockAPIRequest simulates an API request
func mockAPIRequest(user string) {
	fmt.Printf("Making an API request for user %s with role %d\n", user, getCurrentRole(user))
	time.Sleep(time.Millisecond * 50)
}

// simulateUserRequests makes concurrent API requests adhering to the rate limits
func simulateUserRequests(user string, numRequests int, wg *sync.WaitGroup) {
	defer wg.Done()

	limits := rateLimits[getCurrentRole(user)]
	if limits == 0 {
		fmt.Printf("User role %d has no rate limit.\n", getCurrentRole(user))
		return
	}

	reqQueue := make(chan struct{}, limits) // Semaphore with limit of concurrent requests

	for i := 0; i < numRequests; i++ {
		// Ensure that we don't exceed the rate limit
		reqQueue <- struct{}{}
		go func() {
			defer func() {
				<-reqQueue // Release the slot after request
			}()

			currentRole := getCurrentRole(user)
			for {
				if limits == 0 {
					break
				}
				// Retrieve the current rate limit for the user's role
				updatedLimits := rateLimits[currentRole]

				// Check if the role has changed
				if updatedLimits != limits {
					limits = updatedLimits
					reqQueue = make(chan struct{}, limits) // Reset semaphore
					fmt.Printf("User %s's role changed from %d to %d; resetting request queue to %d slots\n", user, currentRole, getCurrentRole(user), limits)
					continue
				}

				// Proceed with the request if the role hasn't changed
				mockAPIRequest(user)
				break
			}
		}()

		// Sleep to throttle the rate of requests per the rate limit
		time.Sleep(time.Second / time.Duration(limits))
	}
}

func main() {
	wg := sync.WaitGroup{}

	// Simulate requests for users with different roles
	user := "alice"
	userRoles[user] = Standard
	wg.Add(1)
	go simulateUserRequests(user, 20, &wg)

	user = "bob"
	userRoles[user] = Premium
	wg.Add(1)
	go simulateUserRequests(user, 15, &wg)

	user = "charlie"
	userRoles[user] = Free
	wg.Add(1)
	go simulateUserRequests(user, 8, &wg)

	// Simulate a dynamic role change for user "alice"
	time.Sleep(5 * time.Second)
	updateUserRole(user, Free) // Change Alice's role to Free

	// Wait for all requests to be completed before exiting
	wg.Wait()
	fmt.Println("All requests completed.")
}