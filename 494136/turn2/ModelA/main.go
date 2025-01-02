package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
	"hash/fnv"
)

// CustomMap defines a map with dynamic resizing
type CustomMap struct {
	loadFactor float64
	buckets    []bucket
	mu         sync.Mutex
}

type bucket struct {
	data map[interface{}]interface{}
	mu   sync.RWMutex
}

// NewCustomMap initializes a new CustomMap
func NewCustomMap(loadFactor float64, initialSize int) *CustomMap {
	buckets := make([]bucket, initialSize)
	for i := range buckets {
		buckets[i] = bucket{data: make(map[interface{}]interface{})}
	}
	return &CustomMap{
		loadFactor: loadFactor,
		buckets:    buckets,
	}
}

// Set inserts or updates a key-value pair in the map
func (m *CustomMap) Set(key, value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Resize if needed
	m.resizeIfNeeded()

	// Insert into the appropriate bucket
	index := m.hash(key)
	b := &m.buckets[index]
	b.mu.Lock()
	defer b.mu.Unlock()
	b.data[key] = value
}

// Get retrieves the value for a key
func (m *CustomMap) Get(key interface{}) (interface{}, bool) {
	index := m.hash(key)
	b := &m.buckets[index]

	b.mu.RLock()
	defer b.mu.RUnlock()

	value, exists := b.data[key]
	return value, exists
}

// Delete removes a key-value pair from the map
func (m *CustomMap) Delete(key interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()

	index := m.hash(key)
	b := &m.buckets[index]

	b.mu.Lock()
	defer b.mu.Unlock()
	delete(b.data, key)
}

// resizeIfNeeded checks the load factor and resizes if necessary
func (m *CustomMap) resizeIfNeeded() {
	totalKeys := 0
	for _, b := range m.buckets {
		totalKeys += len(b.data)
	}

	if float64(totalKeys)/float64(len(m.buckets)) > m.loadFactor {
		m.resize(len(m.buckets) * 2)
	}
}

// resize changes the number of buckets and redistributes entries
func (m *CustomMap) resize(newSize int) {
	newBuckets := make([]bucket, newSize)
	for i := range newBuckets {
		newBuckets[i] = bucket{data: make(map[interface{}]interface{})}
	}

	for _, b := range m.buckets {
		for key, value := range b.data {
			index := m.hashToIndex(key, newSize)
			newBuckets[index].data[key] = value
		}
	}

	m.buckets = newBuckets
}

// hash computes the bucket index for a key
func (m *CustomMap) hash(key interface{}) int {
	return m.hashToIndex(key, len(m.buckets))
}

// hashToIndex computes a hash index within a given size
func (m *CustomMap) hashToIndex(key interface{}, size int) int {
	hasher := fnv.New32a()
	_, _ = hasher.Write([]byte(fmt.Sprintf("%v", key)))
	return int(hasher.Sum32()) % size
}

func main() {
	customMap := NewCustomMap(0.75, 16)

	// Benchmark Memory Usage
	fmt.Println("Benchmark Memory Usage:")
	benchmarkMemoryUsage("Insertions", func(num int) {
		for i := 0; i < num; i++ {
			customMap.Set(i, i)
		}
	})

	benchmarkMemoryUsage("Deletions", func(num int) {
		for i := num - 1; i >= 0; i-- {
			customMap.Delete(i)
		}
	})
}

func benchmarkMemoryUsage(name string, operation func(num int)) {
	start := time.Now()
	startMem := runtime.MemStats{}
	runtime.ReadMemStats(&startMem)

	operation(1_000_0)

	endMem := runtime.MemStats{}
	runtime.ReadMemStats(&endMem)

	memUsage := endMem.Alloc - startMem.Alloc

	duration := time.Since(start)

	fmt.Printf("%s: Memory usage increased by %d bytes, duration %v\n", name, memUsage, duration)
}