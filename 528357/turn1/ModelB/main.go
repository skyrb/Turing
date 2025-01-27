package main

import (
	"fmt"
	"sync"
)

// Define a set data structure using a map
type Set map[int]struct{}

// Add a new element to the set
func (s *Set) Add(i int) {
	(*s)[i] = struct{}{}
}

// Check if an element exists in the set
func (s *Set) Contains(i int) bool {
	_, ok := (*s)[i]
	return ok
}

func main() {
	// Create a new set to store unique user IDs
	uniqueUserIDs := Set{}

	// Simulate importing data from multiple sources
	source1 := []int{1, 2, 3, 4, 5}
	source2 := []int{3, 4, 5, 6, 7}
	source3 := []int{8, 9, 10, 11, 12}

	// Use a goroutine to process each source concurrently
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		for _, userID := range source1 {
			uniqueUserIDs.Add(userID)
		}
		wg.Done()
	}()

	go func() {
		for _, userID := range source2 {
			uniqueUserIDs.Add(userID)
		}
		wg.Done()
	}()

	go func() {
		for _, userID := range source3 {
			uniqueUserIDs.Add(userID)
		}
		wg.Done()
	}()

	wg.Wait() // Wait for all goroutines to finish

	// Print the unique user IDs
	fmt.Println("Unique User IDs:")
	for userID := range uniqueUserIDs {
		fmt.Println(userID)
	}
}