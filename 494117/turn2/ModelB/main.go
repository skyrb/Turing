package main

import (
	"fmt"
)

// Our stateful application represents a simple queue
type Queue struct {
	items []int
}

// Push an element onto the queue
func (q *Queue) Push(item int) {
	q.items = append(q.items, item)
}

// Pop an element from the queue
func (q *Queue) Pop() int {
	if len(q.items) == 0 {
		return 0 // Return 0 for an empty queue
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func processQueue(q *Queue) {
	// Start a transaction to manage state changes
	defer func() {
		if r := recover(); r != nil {
			// If an error occurs, roll back any state changes
			fmt.Println("Error processing queue:", r)
			// Reset the queue to its initial state
			*q = Queue{}
		}
	}()

	fmt.Println("Processing queue:")
	for q.items != nil {
		item := q.Pop()
		fmt.Println("Processed item:", item)
		if item%3 == 0 {
			// Simulate an error by panicking when processing item 3
			panic("Error processing item 3")
		}
	}
	fmt.Println("Queue processing complete.")
}

func main() {
	q := Queue{items: []int{1, 2, 3, 4, 5}}

	// Start processing the queue
	processQueue(&q)

	// At this point, the queue should be empty (if processing is successful)
	fmt.Println("Final queue state:", q)
}  