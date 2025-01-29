package main

import (
	"fmt"
	"time"
)

// SessionData represents user session information
type SessionData struct {
	UserID       string
	Timestamp    time.Time
	ShoppingCart []string // Simplified example of shopping cart items
}

// SessionStorage defines the behavior for session storage systems
type SessionStorage interface {
	GetSession(userID string) (SessionData, error)
	SetSession(session SessionData) error
}

// ConflictResolver defines the behavior for resolving conflicts when merging sessions
type ConflictResolver interface {
	ResolveConflict(local SessionData, remote SessionData) SessionData
}

// LatestTimestampResolver implements the ConflictResolver interface
// It resolves conflicts by choosing the session with the latest timestamp
type LatestTimestampResolver struct{}

func (l *LatestTimestampResolver) ResolveConflict(local SessionData, remote SessionData) SessionData {
	if local.Timestamp.After(remote.Timestamp) {
		return local
	}
	return remote
}

// MergeShoppingCartResolver implements the ConflictResolver interface
// It resolves conflicts by merging the shopping carts of both sessions
type MergeShoppingCartResolver struct{}

func (m *MergeShoppingCartResolver) ResolveConflict(local SessionData, remote SessionData) SessionData {
	// Merge shopping carts by appending unique items from both
	mergedCart := append(local.ShoppingCart, remote.ShoppingCart...)
	uniqueCart := make(map[string]bool)
	for _, item := range mergedCart {
		uniqueCart[item] = true
	}
	mergedCart = make([]string, 0, len(uniqueCart))
	for item := range uniqueCart {
		mergedCart = append(mergedCart, item)
	}

	// Update the timestamp with the latest one
	timestamp := local.Timestamp
	if remote.Timestamp.After(timestamp) {
		timestamp = remote.Timestamp
	}

	return SessionData{
		UserID:       local.UserID,
		Timestamp:    timestamp,
		ShoppingCart: mergedCart,
	}
}

// GlobalSessionDistributor simulates synchronizing session data across multiple regions
type GlobalSessionDistributor struct {
	regions       map[string]SessionStorage
	conflictResolver ConflictResolver
}

// NewGlobalSessionDistributor initializes a new GlobalSessionDistributor
func NewGlobalSessionDistributor(conflictResolver ConflictResolver) *GlobalSessionDistributor {
	return &GlobalSessionDistributor{
		regions:       make(map[string]SessionStorage),
		conflictResolver: conflictResolver,
	}
}

// RegisterRegion adds a new session storage region to the distributor
func (dist *GlobalSessionDistributor) RegisterRegion(region string, storage SessionStorage) {
	dist.regions[region] = storage
}

// SyncSession synchronizes session data across all registered regions and resolves conflicts using the configured resolver
func (dist *GlobalSessionDistributor) SyncSession(userID string) error {
	localSession, err := dist.regions["local"].GetSession(userID)
	if err != nil {
		return fmt.Errorf("failed to get local session: %w", err)
	}

	var mergedSession SessionData
	for region, store := range dist.regions {
		if region == "local" {
			continue // Skip local region
		}

		remoteSession, err := store.GetSession(userID)
		if err != nil {
			// Handle error as per your application logic
			continue
		}

		// Resolve conflicts between local and remote sessions
		if mergedSession.UserID == "" {
			mergedSession = localSession
		}
		mergedSession = dist.conflictResolver.ResolveConflict(mergedSession, remoteSession)
	}

	// Update the local session with the merged data
	if err := dist.regions["local"].SetSession(mergedSession); err != nil {
		return fmt.Errorf("failed to update local session: %w", err)
	}

	return nil
}

func main() {
	// Initialize a global session distributor with a conflict resolver strategy
	distributor := NewGlobalSessionDistributor(&MergeShoppingCartResolver{})

	// Initialize regional stores
	localRegion := NewInMemorySessionStore()
	euRegion := NewInMemorySessionStore()
	usRegion := NewInMemorySessionStore()

	// Register regions with the distributor
	distributor.RegisterRegion("local", localRegion)
	distributor.RegisterRegion("EU", euRegion)
	distributor.RegisterRegion("US", usRegion)

	// Simulate different session data in each region
	userID := "user123"
	localSession := SessionData{
		UserID:       userID,
		Timestamp:    time.Now(),
		ShoppingCart: []string{"item1", "item2"},
	}
	euSession := SessionData{
		UserID:       userID,
		Timestamp:    time.Now().Add(-time.Hour), // Older timestamp to simulate conflict
		ShoppingCart: []string{"item3"},
	}
	usSession := SessionData{
		UserID:       userID,
		Timestamp:    time.Now().Add(-2 * time.Hour), // Older timestamp to simulate conflict
		ShoppingCart: []string{"item1"}, // Duplicate item
	}

	_ = localRegion.SetSession(localSession)
	_ = euRegion.SetSession(euSession)
	_ = usRegion.SetSession(usSession)

	// Sync sessions across regions to resolve conflicts
	err := distributor.SyncSession(userID)
	if err != nil {
		fmt.Println("Error syncing sessions:", err)
		return
	}

	// Retrieve the final merged session from the local region
	finalSession, err := localRegion.GetSession(userID)
	if err != nil {
		fmt.Println("Error fetching final session:", err)
		return
	}

	fmt.Printf("Final Merged Session: %+v\n", finalSession)
} 