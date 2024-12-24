
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
)

// User represents a user with its current role
type User struct {
	id        int
	role      UserRole
	roleMutex sync.Mutex
}

// GetRole returns the current role of the user
func (u *User) GetRole() UserRole {
	u.roleMutex.Lock()
	defer u.roleMutex.Unlock()
	return u.role
}

// SetRole sets the current role of the user
func (u *User) SetRole(role UserRole) {
	u.roleMutex.Lock()
	defer u.roleMutex.Unlock()
	u.role = role
}

// mockAPIRequest simulates an API request for a given user
func mockAPIRequest(user *User) {
	fmt.Printf("Making an API request for user id: %d with role: %d\n", user.id, user.GetRole())
	time.Sleep(time.Millisecond * 50)
}

// simulateUserRequests makes concurrent API requests adhering to the rate limits
func simulateUserRequests(user *User, numRequests int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < numRequests; i++ {
		userRole := user.GetRole()
		limits := rateLimits[userRole]
		if limits == 0 {
			fmt.Printf("User id: %d with role: %d has no rate limit.\n", user.id, userRole)
			continue
		}

		reqQueue := make(chan struct{}, limits) // Semaphore with limit of concurrent requests

		go func() {
			reqQueue <- struct{}{}
			defer func() {
				<-reqQueue // Release the slot after request
			}()
			mockAPIRequest(user)
		}()

		// Sleep to throttle the rate of requests per the rate limit
		time.Sleep(time.Second / time.Duration(limits))
	}
}

func main() {
	wg := sync.WaitGroup{}

	user1 := &User{id: 1, role: Premium}

	// Simulate requests for user1 with changing roles
	wg.Add(1)
	go func() {
		defer wg.Done()
		simulateUserRequests(user1, 10, &wg)
		time.Sleep(time.Second) // Simulate role update after some time
		user1.SetRole(Standard)
		simulateUserRequests(user1, 10, &wg)
	}()

	wg.Add(1)
	go simulateUserRequests(&User{id: 2, role: Standard}, 15, &wg)

	wg.Add(1)
	go simulateUserRequests(&User{id: 3, role: Free}, 8, &wg)

	// Wait for all requests to be completed before exiting
	wg.Wait()
	fmt.Println("All requests completed.")
}
  