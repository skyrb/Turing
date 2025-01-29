package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Metrics struct {
	NetworkLatency float64
	ServerLoad     float64
}

func monitorSystem() Metrics {
	// Placeholder for actual monitoring code
	// In real scenarios, this could involve pinging servers, checking CPU load, etc.
	return Metrics{
		NetworkLatency: 50.0, // in milliseconds
		ServerLoad:     0.5,  // CPU load, 0.0 to 1.0
	}
}

type TransactionType int

const (
	NormalTransaction TransactionType = iota
	HighValueTransaction
)

// SessionData represents session information for a user
type SessionData struct {
	UserID     string
	LastActive time.Time
	CartItems  []string
}

// ConflictResolver interface defines methods for resolving session conflicts
type ConflictResolver interface {
	ResolveConflict(local, remote SessionData) (SessionData, error)
}

// LastActiveConflictResolver prioritizes sessions by last activity timestamp
type LastActiveConflictResolver struct{}

func (r *LastActiveConflictResolver) ResolveConflict(local, remote SessionData) (SessionData, error) {
	if local.UserID != remote.UserID {
		return SessionData{}, errors.New("cannot resolve conflict: different users")
	}

	// Compare timestamps; choose the session with the most recent activity
	if local.LastActive.After(remote.LastActive) {
		return local, nil
	}
	return remote, nil
}

// RegionState represents a regional view of user sessions
type RegionState struct {
	sessions map[string]SessionData
	mu       sync.RWMutex
}

// NewRegionState initializes a new regional state
func NewRegionState() *RegionState {
	return &RegionState{
		sessions: make(map[string]SessionData),
	}
}

// SetSession sets session data in the region
func (rs *RegionState) SetSession(session SessionData) {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	rs.sessions[session.UserID] = session
}

// GetSession retrieves session data for a given user, if exists
func (rs *RegionState) GetSession(userID string) (SessionData, bool) {
	rs.mu.RLock()
	defer rs.mu.RUnlock()
	session, exists := rs.sessions[userID]
	return session, exists
}

// GlobalSessionManager manages sessions across multiple regions
type GlobalSessionManager struct {
	regions          map[string]*RegionState
	conflictResolver ConflictResolver
}

// NewGlobalSessionManager creates a new global session manager
func NewGlobalSessionManager(resolver ConflictResolver) *GlobalSessionManager {
	return &GlobalSessionManager{
		regions:          make(map[string]*RegionState),
		conflictResolver: resolver,
	}
}

// RegisterRegion adds a new region to the manager
func (gsm *GlobalSessionManager) RegisterRegion(regionName string, state *RegionState) {
	gsm.regions[regionName] = state
}


func (gsm *GlobalSessionManager) AdaptiveSyncSessionAcrossRegions() {
	metrics := monitorSystem()

	// Example adaptive logic
	var syncInterval time.Duration
	if metrics.NetworkLatency > 100 || metrics.ServerLoad > 0.8 {
		syncInterval = 10 * time.Second // Reduce frequency due to high congestion/load
	} else {
		syncInterval = 2 * time.Second // Increase frequency when possible
	}

	fmt.Printf("Sync interval set to %v based on metrics\n", syncInterval)

	ticker := time.NewTicker(syncInterval)
	defer ticker.Stop()

	for range ticker.C {
		gsm.SyncSessionAcrossRegions()
	}
}

func (gsm *GlobalSessionManager) SyncSessionAcrossRegions() {
	userSessions := make(map[string]SessionData)

	// Detect conflict and prioritize based on TransactionType
	for _, region := range gsm.regions {
		region.mu.RLock()
		for userID, session := range region.sessions {
			localSession, exists := userSessions[userID]
			if !exists || (session.TransactionType == HighValueTransaction &&
				localSession.TransactionType != HighValueTransaction) {
				userSessions[userID] = session
			} else if exists {
				resolvedSession, err := gsm.conflictResolver.ResolveConflict(localSession, session)
				if err != nil {
					fmt.Printf("Error resolving conflict for user %s: %v\n", userID, err)
				} else {
					userSessions[userID] = resolvedSession
				}
			}
		}
		region.mu.RUnlock()
	}

	// Apply resolved states back to each regional state
	for userID, resolvedSession := range userSessions {
		for _, region := range gsm.regions {
			region.SetSession(resolvedSession)
		}
	}
}

func main() {
	// Initialize conflict resolver
	resolver := &LastActiveConflictResolver{}

	// Setup global session manager
	manager := NewGlobalSessionManager(resolver)

	// Create regional states
	usRegion := NewRegionState()
	euRegion := NewRegionState()

	// Register regions
	manager.RegisterRegion("US", usRegion)
	manager.RegisterRegion("EU", euRegion)

	// Simulate adding sessions in each region
	usRegion.SetSession(SessionData{
		UserID:     "user123",
		LastActive: time.Now().Add(-5 * time.Minute),
		CartItems:  []string{"item1", "item2"},
	})

	euRegion.SetSession(SessionData{
		UserID:     "user123",
		LastActive: time.Now(),
		CartItems:  []string{"item3"},
	})

	// Synchronize sessions information across regions
	manager.SyncSessionAcrossRegions()

	if session, exists := usRegion.GetSession("user123"); exists {
		fmt.Printf("Resolved session in US region: %+v\n", session)
	}

	if session, exists := euRegion.GetSession("user123"); exists {
		fmt.Printf("Resolved session in EU region: %+v\n", session)
	}
}