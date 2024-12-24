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

type User struct {
	id        int
	role      UserRole
	roleMutex sync.Mutex
}

func (u *User) GetRole() UserRole {
	u.roleMutex.Lock()
	defer u.roleMutex.Unlock()
	return u.role
}

func (u *User) SetRole(role UserRole) {
	u.roleMutex.Lock()
	defer u.roleMutex.Unlock()
	u.role = role
}

// Endpoint represents an API endpoint with a rate limit per user role
type Endpoint struct {
	name        string
	rateLimits map[UserRole]int
	channels   map[UserRole]chan struct{}
}

func newEndpoint(name string, rateLimits map[UserRole]int) *Endpoint {
	e := &Endpoint{
		name:        name,
		rateLimits: rateLimits,
		channels:   make(map[UserRole]chan struct{}),
	}
	for role, limit := range rateLimits {
		if limit > 0 {
			e.channels[role] = make(chan struct{}, limit)
		}
	}
	return e
}

func (e *Endpoint) getChannel(role UserRole) chan struct{} {
	if limits, ok := e.channels[role]; ok {
		return limits
	}
	return nil
}

func mockAPIRequest(user *User, endpoint *Endpoint) {
	fmt.Printf("Making an API request to endpoint: %s for user id: %d with role: %d\n", endpoint.name, user.id, user.GetRole())
	time.Sleep(time.Millisecond * 50)
}

func simulateUserRequests(user *User, endpoints []*Endpoint, numRequests int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < numRequests; i++ {
		userRole := user.GetRole()
		for _, endpoint := range endpoints {
			if limits := endpoint.getChannel(userRole); limits != nil {
				go func(endpoint *Endpoint) {
					limits <- struct{}{}
					defer func() {
						<-limits // Release the slot after request
					}()
					mockAPIRequest(user, endpoint)
				}(endpoint)

				// Sleep to throttle the rate of requests per the rate limit
				if limit := endpoint.rateLimits[userRole]; limit > 0 {
					time.Sleep(time.Second / time.Duration(limit))
				}
			}
		}
	}
}

func main() {
	wg := sync.WaitGroup{}

	user1 := &User{id: 1, role: Premium}

	// Define endpoints with rate limits per user role
	endpoints := []*Endpoint{
		newEndpoint("/v1/users", map[UserRole]int{Premium: 10, Standard: 5, Free: 2}),
		newEndpoint("/v1/orders", map[UserRole]int{Premium: 8, Standard: 4, Free: 1}),
	}

	// Simulate requests for user1 with changing roles
	wg.Add(1)
	go func() {
		defer wg.Done()
		simulateUserRequests(user1, endpoints, 20, &wg)
		time.Sleep(time.Second) // Simulate role update after some time
		user1.SetRole(Standard)
		simulateUserRequests(user1, endpoints, 20, &wg)
	}()

	wg.Add(1)
	go simulateUserRequests(&User{id: 2, role: Standard}, endpoints, 15, &wg)

	wg.Add(1)
	go simulateUserRequests(&User{id: 3, role: Free}, endpoints, 8, &wg)

	// Wait for all requests to be completed before exiting
	wg.Wait()
	fmt.Println("All requests completed.")
}