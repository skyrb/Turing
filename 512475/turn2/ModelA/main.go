package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

type AnnotatedMap struct {
	sync.RWMutex
	m           map[int]int
	readCount    int64
	writeCount   int64
	errors       int64
	nilAccesses  int64
	concurrentErrors int64
}

// NewAnnotatedMap initializes the map
func NewAnnotatedMap() *AnnotatedMap {
	return &AnnotatedMap{
		m: make(map[int]int),
	}
}

// Set adds or updates a key-value pair and logs the write operation
func (am *AnnotatedMap) Set(key int, value int) {
	am.Lock()
	defer am.Unlock()

	if am.m == nil {
		log.Println("[Error] Attempt to write to a nil map.")
		am.errors++
		am.nilAccesses++
		return
	}

	am.m[key] = value
	am.writeCount++
	log.Printf("[Write] Key: %d, Value: %d\n", key, value)
}

// Get retrieves a value for a key and logs the read operation
func (am *AnnotatedMap) Get(key int) (int, bool) {
	am.RLock()
	defer am.RUnlock()

	if am.m == nil {
		log.Println("[Error] Attempt to read from a nil map.")
		am.errors++
		am.nilAccesses++
		return 0, false
	}

	value, exists := am.m[key]
	am.readCount++
	if !exists {
		log.Printf("[Warning] Key %d does not exist.\n", key)
	}
	log.Printf("[Read] Key: %d, Value: %d, Exists: %t\n", key, value, exists)
	return value, exists
}

// Delete removes a key-value pair and logs the delete operation
func (am *AnnotatedMap) Delete(key int) {
	am.Lock()
	defer am.Unlock()

	if am.m == nil {
		log.Println("[Error] Attempt to delete from a nil map.")
		am.errors++
		am.nilAccesses++
		return
	}

	delete(am.m, key)
	log.Printf("[Delete] Key: %d\n", key)
}

// MonitorUsage continuously logs usage statistics
func (am *AnnotatedMap) MonitorUsage() {
	for {
		time.Sleep(5 * time.Second)
		am.RLock()
		log.Printf("[Metrics] Reads: %d, Writes: %d, Errors: %d, Nil Accesses: %d, Concurrent Errors: %d\n",
			am.readCount, am.writeCount, am.errors, am.nilAccesses, am.concurrentErrors)
		am.RUnlock()
	}
}

// concurrentMapOperations simulates concurrent access and modifications to the map
func concurrentMapOperations(am *AnnotatedMap) {
	for {
		key := rand.Intn(100)
		value := rand.Intn(100)

		// Randomly decide to perform a read, write, or delete
		operation := rand.Intn(3)
		switch operation {
		case 0: // Write
			am.Set(key, value)
		case 1: // Read
			am.Get(key)
		case 2: // Delete
			am.Delete(key)
		}

		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	// Initialize the annotated map
	am := NewAnnotatedMap()

	// Start monitoring usage metrics
	go am.MonitorUsage()

	// Start multiple goroutines performing map operations
	numGoroutines := 5
	var wg sync.WaitGroup
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			concurrentMapOperations(am)
		}()
	}

	// Let the system run for 15 seconds
	time.Sleep(15 * time.Second)

	// Stop the goroutines and wait for them to finish
	log.Println("Stopping operations...")
	wg.Wait()
	log.Println("All operations stopped.")
}