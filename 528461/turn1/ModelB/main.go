package main

import (
    "fmt"
    "sync"
)

// Session represents user session data
type Session struct {
    UserID   string
    CartItems []string
}

// RegionCache represents an in-memory cache for a specific region
type RegionCache interface {
    Get(key string) (Session, bool)
    Set(key string, session Session)
    Delete(key string)
}

// Broker represents a central broker to synchronize cache data across regions
type Broker interface {
    RegisterRegion(region string, cache RegionCache)
    GetSession(userID string) (Session, bool)
    SetSession(userID string, session Session)
    DeleteSession(userID string)
}

// InMemoryCache is a simple in-memory implementation of RegionCache
type InMemoryCache struct {
    cache map[string]Session
    mu    sync.RWMutex
}

func NewInMemoryCache() *InMemoryCache {
    return &InMemoryCache{cache: make(map[string]Session)}
}

func (c *InMemoryCache) Get(key string) (Session, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    session, ok := c.cache[key]
    return session, ok
}

func (c *InMemoryCache) Set(key string, session Session) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.cache[key] = session
}

func (c *InMemoryCache) Delete(key string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    delete(c.cache, key)
}

// SimpleBroker is a basic implementation of Broker using a map to store regions
type SimpleBroker struct {
    regions map[string]RegionCache
    mu      sync.RWMutex
}

func NewSimpleBroker() *SimpleBroker {
    return &SimpleBroker{regions: make(map[string]RegionCache)}
}

func (b *SimpleBroker) RegisterRegion(region string, cache RegionCache) {
    b.mu.Lock()
    defer b.mu.Unlock()
    b.regions[region] = cache
}

func (b *SimpleBroker) GetSession(userID string) (Session, bool) {
    b.mu.RLock()
    defer b.mu.RUnlock()
    // For simplicity, we'll assume the user ID is the same as the cache key
    for _, cache := range b.regions {
        if session, ok := cache.Get(userID); ok {
            return session, true
        }
    }
    return Session{}, false
}

func (b *SimpleBroker) SetSession(userID string, session Session) {
    b.mu.Lock()
    defer b.mu.Unlock()
    for _, cache := range b.regions {
        cache.Set(userID, session)
    }
}

func (b *SimpleBroker) DeleteSession(userID string) {
    b.mu.Lock()
    defer b.mu.Unlock()
    for _, cache := range b.regions {
        cache.Delete(userID)
    }
}

func main() {
    // Create regional caches and the broker
    usCache := NewInMemoryCache()
    euCache := NewInMemoryCache()
    broker := NewSimpleBroker()

    // Register the regions with the broker
    broker.RegisterRegion("us", usCache)
    broker.RegisterRegion("eu", euCache)

    // Create and set a user session
    userID := "user123"
    session := Session{UserID: userID, CartItems: []string{"ItemA", "ItemB"}}
    broker.SetSession(userID, session)

    // Get the session from another region (simulating user switching locations)
    if session, ok := broker.GetSession(userID); ok {
        fmt.Println("Session found in EU:", session)
    } else {
        fmt.Println("Session not found in EU.")
    }

    // Update and delete the session
    session.CartItems = append(session.CartItems, "ItemC")
    broker.SetSession(userID, session)
    broker.DeleteSession(userID)
}  