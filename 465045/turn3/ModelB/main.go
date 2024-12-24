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

// Endpoint represents an API endpoint with its rate limits
type Endpoint string

const (
    Endpoint1 Endpoint = "/api/v1/endpoint1"
    Endpoint2 Endpoint = "/api/v1/endpoint2"
)

// RateLimitManager manages rate limits for a user and endpoint
type RateLimitManager struct {
    userRole     UserRole
    endpoint     Endpoint
    rateLimits   map[UserRole]int
    rateLimitMut sync.Mutex
    reqQueue     chan struct{}
}

func (rlm *RateLimitManager) init() {
    rlm.rateLimitMut.Lock()
    defer rlm.rateLimitMut.Unlock()

    rlm.rateLimits = make(map[UserRole]int)
    rlm.rateLimits[Premium] = 10  
    rlm.rateLimits[Standard] = 5  
    rlm.rateLimits[Free] = 2  
    rlm.reqQueue = make(chan struct{}, rlm.getRateLimit())
}

func (rlm *RateLimitManager) getRateLimit() int {
    rlm.rateLimitMut.Lock()
    defer rlm.rateLimitMut.Unlock()
    return rlm.rateLimits[rlm.userRole]
}

func (rlm *RateLimitManager) setRateLimit(limit int) {
    rlm.rateLimitMut.Lock()
    defer rlm.rateLimitMut.Unlock()
    rlm.rateLimits[rlm.userRole] = limit
    // Adjust the reqQueue size if the rate limit changes
    close(rlm.reqQueue)
    rlm.reqQueue = make(chan struct{}, limit)
}

// User represents a user with its current role
type User struct {  
    id        int
    role      UserRole
    roleMutex sync.Mutex
    rlManagers map[Endpoint]*RateLimitManager
}

func (u *User) init() {  
    u.rlManagers = make(map[Endpoint]*RateLimitManager)
    u.addRLManager(Endpoint1)
    u.addRLManager(Endpoint2)
}

func (u *User) addRLManager(endpoint Endpoint) {
    rlm := &RateLimitManager{
        userRole: u.role,
        endpoint: endpoint,
    }
    rlm.init()
    u.rlManagers[endpoint] = rlm
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
    // Update rate limit for all endpoints when role changes
    for _, rlm := range u.rlManagers {
        rlm.setRateLimit(rateLimits[role])
    }
}

// mockAPIRequest simulates an API request for a given user at a specific endpoint
func mockAPIRequest(user *User, endpoint Endpoint) {  
    fmt.Printf("Making an API request for user id: %d with role: %d at endpoint: %s\n", user.id, user.GetRole(), endpoint)
    time.Sleep(time.Millisecond * 50)
}

// simulateUserRequests makes concurrent API requests adhering to the rate limits for a specific endpoint
func simulateUserRequests(user *User, endpoint Endpoint, numRequests int, wg *sync.WaitGroup) {  
    defer wg.Done()