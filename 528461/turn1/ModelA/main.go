package main

import (
	"fmt"
	"sync"
	"time"
)

// SessionData represents user session information
type SessionData struct {
	UserID    string
	Timestamp time.Time
	Data      map[string]interface{}
}

// SessionStorage defines the behavior for session storage systems
type SessionStorage interface {
	GetSession(userID string) (SessionData, error)
	SetSession(session SessionData) error
}

// InMemorySessionStore is a simple in-memory storage for session data
type InMemorySessionStore struct {
	sessions map[string]SessionData
	mu       sync.RWMutex
}

// NewInMemorySessionStore initializes a new InMemorySessionStore
func NewInMemorySessionStore() *InMemorySessionStore {
	return &InMemorySessionStore{
		sessions: make(map[string]SessionData),
	}
}

func (store *InMemorySessionStore) GetSession(userID string) (SessionData, error) {
	store.mu.RLock()
	defer store.mu.RUnlock()

	session, exists := store.sessions[userID]
	if !exists {
		return SessionData{}, fmt.Errorf("session not found")
	}
	return session, nil
}

func (store *InMemorySessionStore) SetSession(session SessionData) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	store.sessions[session.UserID] = session
	return nil
}

// Simulated global distributor for session data synchronization
type GlobalSessionDistributor struct {
	regions map[string]SessionStorage
}

// NewGlobalSessionDistributor initializes a new GlobalSessionDistributor
func NewGlobalSessionDistributor() *GlobalSessionDistributor {
	return &GlobalSessionDistributor{
		regions: make(map[string]SessionStorage),
	}
}

// RegisterRegion adds a new session storage region to the distributor
func (dist *GlobalSessionDistributor) RegisterRegion(region string, storage SessionStorage) {
	dist.regions[region] = storage
}

// SyncSession synchronizes session data across all registered regions
func (dist *GlobalSessionDistributor) SyncSession(session SessionData) {
	for region, store := range dist.regions {
		fmt.Printf("Syncing session to region: %s\n", region)
		_ = store.SetSession(session) // Handle error appropriately in production systems
	}
}

func main() {
	// Initialize a global session distributor
	distributor := NewGlobalSessionDistributor()

	// Initialize regional stores
	usRegion := NewInMemorySessionStore()
	euRegion := NewInMemorySessionStore()

	// Register regions with the distributor
	distributor.RegisterRegion("US", usRegion)
	distributor.RegisterRegion("EU", euRegion)

	// Create a user session
	session := SessionData{
		UserID:    "user123",
		Timestamp: time.Now(),
		Data: map[string]interface{}{
			"lastPage": "home",
		},
	}

	// Sync session across regions
	distributor.SyncSession(session)

	// Retrieve session from a specific region
	retrievedSession, err := usRegion.GetSession(session.UserID)
	if err != nil {
		fmt.Println("Error fetching session:", err)
	}
	fmt.Printf("Session for user %s: %+v\n", retrievedSession.UserID, retrievedSession)
}