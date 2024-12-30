package main

import (
	"fmt"
	"sync"
)

type MyMap struct {
	lock  sync.RWMutex
	data  map[string]int
	load  int
	cap   int
	ratio float64
}

func NewMyMap(capacity int, loadRatio float64) *MyMap {
	return &MyMap{
		data:  make(map[string]int, capacity),
		cap:   capacity,
		ratio: loadRatio,
	}
}

func (m *MyMap) Put(key string, value int) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.data[key] = value
	m.load++

	if m.load >= int(float64(m.cap)*m.ratio) {
		m.resize(2 * m.cap)
	}
}

func (m *MyMap) Get(key string) int {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.data[key]
}

func (m *MyMap) Delete(key string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	delete(m.data, key)
	m.load--

	if m.load < int(float64(m.cap)/4.0) && m.cap > 1 {
		m.resize(m.cap / 2)
	}
}

func (m *MyMap) resize(newCap int) {
	newData := make(map[string]int, newCap)
	for key, value := range m.data {
		newData[key] = value
	}
	m.data = newData
	m.cap = newCap
}

func main() {
	// Original implementation with fixed capacity
	originalMap := make(map[string]int)

	// Optimized implementation with dynamic resizing
	optimizedMap := NewMyMap(100, 0.75)

	const numKeys = 1000000
	const numIterations = 10

	// Benchmark original map
	fmt.Println("Benchmarking original map...")
	for i := 0; i < numIterations; i++ {
		for j := 0; j < numKeys; j++ {
			key := fmt.Sprintf("key%d", j)
			originalMap[key] = j
		}
		for j := 0; j < numKeys; j++ {
			key := fmt.Sprintf("key%d", j)
			_ = originalMap[key]
		}
		for j := 0; j < numKeys; j++ {
			key := fmt.Sprintf("key%d", j)
			delete(originalMap, key)
		}
	}

	// Benchmark optimized map
	fmt.Println("Benchmarking optimized map...")
	for i := 0; i < numIterations; i++ {
		for j := 0; j < numKeys; j++ {
			key := fmt.Sprintf("key%d", j)
			optimizedMap.Put(key, j)
		}
		for j := 0; j < numKeys; j++ {
			key := fmt.Sprintf("key%d", j)
			_ = optimizedMap.Get(key)
		}
		for j := 0; j < numKeys; j++ {
			key := fmt.Sprintf("key%d", j)
			optimizedMap.Delete(key)
		}
	}
}