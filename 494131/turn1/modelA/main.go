package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// Simulates a function that processes a number and may fail.
func processNumber(num int) (result int, err error) {
	if num <= 0 {
		return 0, fmt.Errorf("number must be positive")
	}
	// Simulate work
	time.Sleep(100 * time.Millisecond)
	result = num * num
	return
}

func BenchmarkProcessNumber(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < b.N; i++ {
		num := rand.Intn(10) + 1 // Generate a random positive number
		_, _ = processNumber(num) // Discard the return values
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(10) + 1 // Generate a random positive number

	result, err := processNumber(num)
	if err != nil {
		fmt.Println("Error processing number:", err)
	} else {
		fmt.Println("Result:", result)
	}

	testing.Main()
}