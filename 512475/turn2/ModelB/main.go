package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// Histogram is a simple data structure to count occurrences over time.
type Histogram struct {
	sync.Mutex
	buckets map[int]int
}

func newHistogram() *Histogram {
	return &Histogram{
		buckets: make(map[int]int),
	}
}

func (h *Histogram) Add(value int) {
	h.Lock()
	defer h.Unlock()
	h.buckets[value]++
}

func (h *Histogram) Get(value int) int {
	h.Lock()
	defer h.Unlock()
	return h.buckets[value]
}

func (h *Histogram) Reset() {
	h.Lock()
	defer h.Unlock()
	h.buckets = make(map[int]int)
}

type AnnotatedMap struct {
	sync.RWMutex
	m         map[int]int
	readCount  int64
	writeCount int64
	errors     int64
	// Histograms to track frequency of nil map accesses and concurrent write errors
	nilMapAccessHist *Histogram
	concurrentWriteHist *Histogram
}

// NewAnnotatedMap initializes the map and histograms.
func NewAnnotatedMap() *AnnotatedMap {
	return &AnnotatedMap{
		m:                make(map[int]int),
		nilMapAccessHist: newHistogram(),
		concurrentWriteHist: newHistogram(),
	}
}

// Set adds or updates a key-value pair and logs the write operation.
func (am *AnnotatedMap) Set(key int, value int) {
	am.Lock()
	defer am.Unlock()

	if am.m == nil {
		log.Println("[Error] Attempt to write to a nil map.")
		am.errors++
		am.nilMapAccessHist.Add(1) // Increment nil map access histogram
		return
	}

	am.m[key] = value
	am.writeCount++
	log.Printf("[Write] Key: %d, Value: %d\n", key, value)
}

// Get retrieves a value for a key and logs the read operation.
func (am *AnnotatedMap) Get(key int) (int, bool) {
	am.RLock()
	defer am.RUnlock()

	if am.m == nil {
		log.Println("[Error] Attempt to read from a nil map.")
		am.errors++
		am.nilMapAccessHist.Add(1) // Increment nil map access histogram
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

// Delete removes a key-value pair and logs the delete operation.
func (am *AnnotatedMap) Delete(key int) {
	am.Lock()
	defer am.Unlock()

	if am.m == nil {
		log.Println("[Error] Attempt to delete from a nil map.")
		am.errors++
		am.nilMapAccessHist.Add(1) // Increment nil map access histogram
		return
	}

	delete(am.m, key)
	log.Printf("[Delete] Key: %d\n", key)
}

// MonitorUsage continuously logs usage statistics and analytics.
func (am *AnnotatedMap) MonitorUsage() {
	for {
		time.Sleep(5 * time.Second)