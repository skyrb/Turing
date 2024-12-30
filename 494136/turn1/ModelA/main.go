package main

import (
	"sync"
	"time"
	"fmt"
)

type CustomMap struct {
	loadFactor float64
	loadThresh int
	minSize    int
	buckets    []bucket
	mu         sync.Mutex
}

type bucket struct {
	keys   []interface{}
	values []interface{}
}

func NewCustomMap(loadFactor float64, loadThresh, minSize int) *CustomMap {
	return &CustomMap{
		loadFactor: loadFactor,
		loadThresh: loadThresh,
		minSize:    minSize,
		buckets:    make([]bucket, minSize),
	}
}

func (m *CustomMap) Set(key, value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.resizeIfNeeded()

	i := m.hash(key)
	for j := 0; j < len(m.buckets[i].keys); j++ {
		if m.buckets[i].keys[j] == key {
			m.buckets[i].values[j] = value
			return
		}
	}

	m.buckets[i].keys = append(m.buckets[i].keys, key)
	m.buckets[i].values = append(m.buckets[i].values, value)
}

func (m *CustomMap) Get(key interface{}) (interface{}, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	i := m.hash(key)
	for j := 0; j < len(m.buckets[i].keys); j++ {
		if m.buckets[i].keys[j] == key {
			return m.buckets[i].values[j], true
		}
	}

	return nil, false
}

func (m *CustomMap) Delete(key interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()

	i := m.hash(key)
	for j := 0; j < len(m.buckets[i].keys); j++ {
		if m.buckets[i].keys[j] == key {
			// Move the remaining items one position forward
			for k := j + 1; k < len(m.buckets[i].keys); k++ {
				m.buckets[i].keys[k-1] = m.buckets[i].keys[k]
				m.buckets[i].values[k-1] = m.buckets[i].values[k]
			}

			// Remove the last element
			m.buckets[i].keys = m.buckets[i].keys[:len(m.buckets[i].keys)-1]
			m.buckets[i].values = m.buckets[i].values[:len(m.buckets[i].values)-1]

			return
		}
	}
}

func (m *CustomMap) hash(key interface{}) int {
	return len(m.buckets) - 1 & int(hash(key))
}

func hash(key interface{}) uintptr {
	return 0 // Implement your own hashing function
}

func (m *CustomMap) resizeIfNeeded() {
	n := len(m.buckets)
	k := len(m.buckets[0].keys)
	if float64(k)/float64(n) > m.loadFactor {
		m.resizeUp()
	} else if n > m.minSize && float64(k)/float64(n) < m.loadThresh {
		m.resizeDown()
	}
}

func (m *CustomMap) resizeUp() {
	newBuckets := make([]bucket, len(m.buckets)*2)

	for _, b := range m.buckets {
		for i := range b.keys {
			newI := m.hash(b.keys[i])
			newBuckets[newI].keys = append(newBuckets[newI].keys, b.keys[i])
			newBuckets[newI].values = append(newBuckets[newI].values, b.values[i])
		}
	}

	m.buckets = newBuckets
}

func (m *CustomMap) resizeDown() {
	newBuckets := make([]bucket, len(m.buckets)/2)

	for _, b := range m.buckets {
		for i := range b.keys {
			newI := m.hash(b.keys[i])
			newBuckets[newI].keys = append(newBuckets[newI].keys, b.keys[i])
			newBuckets[newI].values = append(newBuckets[newI].values, b.values[i])
		}
	}

	m.buckets = newBuckets
}

func main() {
	b := NewCustomMap(0.75, 0.2, 16)

	benchmarkInsert(b)
	benchmarkLookup(b)
	benchmarkDelete(b)
}

func benchmarkInsert(m *CustomMap) {
	start := time.Now()

	for i := 0; i < 1_000_000; i++ {
		m.Set(i, "value")
	}

	end := time.Now()
	fmt.Println("Custom map insert time:", end.Sub(start))
}

func benchmarkLookup(m *CustomMap) {
	start := time.Now()

	for i := 0; i < 1_000_000; i++ {
		_, _ = m.Get(i)
	}

	end := time.Now()
	fmt.Println("Custom map lookup time:", end.Sub(start))
}

func benchmarkDelete(m *CustomMap) {
	start := time.Now()

	for i := 0; i < 1_000_000; i++ {
		m.Delete(i)
	}

	end := time.Now()
	fmt.Println("Custom map delete time:", end.Sub(start))
}