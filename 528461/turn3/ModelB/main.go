package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// SessionData represents session information for a user
type SessionData struct {
	UserID     string
	LastActive time.Time
	CartItems  []string
	IsTransact bool // Flag for transaction-related sessions
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

	// If both sessions are transactional, favor the one with more potential
	if local.IsTransact && remote.IsTransact {
		// Implement your logic here to prioritize based on transaction details
		// For simplicity, we'll just pick one at random
		if rand.Float32() < 0.5 {
			return local, nil
		}
		return remote, nil
	}

	// Compare timestamps; choose the session with the most recent activity
	if local.LastActive.After(remote.LastActive) {
		return local, nil
	}
	return remote, nil
}

// DynamicSynchronizer controls the synchronization process
type DynamicSynchronizer struct {
	syncFreq     time.Duration
	minFreq      time.Duration
	maxFreq      time.Duration
	backoffFactor float32
	mutex       sync.Mutex
	stopChan     chan struct{}
}

// NewDynamicSynchronizer initializes a new dynamic synchronizer
func NewDynamicSynchronizer(minFreq, maxFreq time.Duration, backoffFactor float32) *DynamicSynchronizer {
	return &DynamicSynchronizer{
		syncFreq:     minFreq,
		minFreq:      minFreq,
		maxFreq:      maxFreq,
		backoffFactor: backoffFactor,
		stopChan:     make(chan struct{}),
	}
}

// Start starts the dynamic synchronization process
func (ds *DynamicSynchronizer) Start(syncFunc func()) {
	go func() {
		for {
			select {
			case <-ds.stopChan:
				return
			default:
				syncFunc()
				time.Sleep(ds.syncFreq)
			}
		}
	}()
}

// Stop stops the dynamic synchronization process
func (ds *DynamicSynchronizer) Stop() {
	close(ds.stopChan)
}

// AdjustSyncFreq adjusts the synchronization frequency based on success or failure
func (ds *DynamicSynchronizer) AdjustSyncFreq(success bool) {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()

	if success {
		// If synchronization was successful, back off to minimum frequency
		ds.syncFreq = ds.minFreq
	} else {
		// If synchronization failed, increase frequency with backoff
		ds.syncFreq = time.Duration(float32(ds.syncFreq) * ds.backoffFactor)
		if ds.syncFreq > ds.maxFreq {
			ds.syncFreq = ds.maxFreq
		}
	}
}