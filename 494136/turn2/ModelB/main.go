package main  
import (  
    "fmt"
    "runtime"
    "sync"
)

// customMap defines a map with custom key type (int32) and value type (string)
type customMap struct {
    data map[int32]string
    mu   sync.RWMutex
}

func newCustomMap() *customMap {
    return &customMap{data: make(map[int32]string)}
}

// insert inserts a key-value pair into the customMap
func (m *customMap) insert(key int32, value string) {
    m.mu.Lock()
    defer m.mu.Unlock()
    m.data[key] = value
}

// delete deletes a key-value pair from the customMap
func (m *customMap) delete(key int32) {
    m.mu.Lock()
    defer m.mu.Unlock()
    delete(m.data, key)
}

// benchmarkMemoryUsage runs the specified benchmark function and prints memory usage statistics
func benchmarkMemoryUsage(benchmarkFunc func()) {
    var memStats runtime.MemStats
    runtime.ReadMemStats(&memStats)
    beforeAlloc := memStats.Alloc
    benchmarkFunc()
    runtime.ReadMemStats(&memStats)
    afterAlloc := memStats.Alloc
    fmt.Printf("Memory Allocated: %d bytes\n", afterAlloc-beforeAlloc)
}

func main() {
    // Benchmark memory usage during insertions
    benchmarkMemoryUsage(func() {
        customMap := newCustomMap()
        for i := 0; i < 1_000_000; i++ {
            customMap.insert(int32(i), fmt.Sprintf("value-%d", i))
        }
    })

    // Benchmark memory usage during deletions
    benchmarkMemoryUsage(func() {
        customMap := newCustomMap()
        for i := 0; i < 1_000_000; i++ {
            customMap.insert(int32(i), fmt.Sprintf("value-%d", i))
        }
        for i := 0; i < 1_000_000; i++ {
            customMap.delete(int32(i))
        }
    })

    // Strategy 1: Use smaller key type (int32) instead of int
    benchmarkMemoryUsage(func() {
        customMap := newCustomMap()
        for i := 0; i < 1_000_000; i++ {
            customMap.insert(int32(i), fmt.Sprintf("value-%d", i))
        }
    })

    // Strategy 2: Avoid unnecessary string allocations
    benchmarkMemoryUsage(func() {
        customMap := newCustomMap()
        var valueBuffer [10]byte // Assuming each value has maximum length of 9 characters
        for i := 0; i < 1_000_000; i++ {
            n := copy(valueBuffer[:], fmt.Sprintf("value-%d", i))
            customMap.insert(int32(i), string(valueBuffer[:n]))
        }
    })
}