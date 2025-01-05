package main

import (
	"fmt"
	"sync"
	"time"
)

// ReadWriteQueue is a concurrent queue that supports multiple readers and writers.
type ReadWriteQueue struct {
	queue []string
	rwmu  sync.RWMutex
	wg    sync.WaitGroup
}

// NewReadWriteQueue initializes a new ReadWriteQueue instance.
func NewReadWriteQueue() *ReadWriteQueue {
	return &ReadWriteQueue{}
}

// Write adds an element to the queue.
func (q *ReadWriteQueue) Write(item string) {
	q.rwmu.Lock()
	defer q.rwmu.Unlock()

	q.queue = append(q.queue, item)
}

// Read removes and returns an element from the front of the queue.
func (q *ReadWriteQueue) Read() (string, bool) {
	q.rwmu.RLock()
	defer q.rwmu.RUnlock()

	if len(q.queue) == 0 {
		return "", false
	}

	item := q.queue[0]
	q.queue = q.queue[1:]
	return item, true
}

// StartReadGoroutines initiates goroutines to simulate reading from the queue.
func (q *ReadWriteQueue) StartReadGoroutines(count int, interval time.Duration) {
	q.wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			defer q.wg.Done()
			for {
				item, ok := q.Read()
				if !ok {
					// Queue is empty
					break
				}
				fmt.Println("Read:", item)
				time.Sleep(interval)
			}
		}()
	}
}

// WaitForCompletion blocks until all reader goroutines have finished.
func (q *ReadWriteQueue) WaitForCompletion() {
	q.wg.Wait()
}

func main() {
	queue := NewReadWriteQueue()

	// Starting some read goroutines
	queue.StartReadGoroutines(10, time.Second/4)

	// Adding items to the queue concurrently
	for i := 0; i < 10; i++ {
		queue.Write(fmt.Sprintf("Item %d", i))
		time.Sleep(time.Second / 2)
	}

	queue.WaitForCompletion()
	fmt.Println("All read operations completed.")
}