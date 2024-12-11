package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Function that may panic
	startService()

	fmt.Println("Application continued running after handling panic.")
}

// Function that may panic due to an error
func startService() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Recovered from panic: %v\n", err)
			// Optionally restart the service or perform cleanup
			restartService()
		}
	}()

	fmt.Println("Starting service...")

	// Simulate a situation where the service might panic
	number := rand.Intn(10)
	if number == 0 {
		panic("Service failed: critical error detected!")
	}

	fmt.Println("Service running successfully.")
}

// Function to simulate restarting the service
func restartService() {
	fmt.Println("Service restarted.")
	// You could call startService() again here if necessary
}