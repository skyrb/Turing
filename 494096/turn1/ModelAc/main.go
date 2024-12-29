package main

import (
	"fmt"
	"time"
)

var currentState string

func transitionState(newState string) {
	oldState := currentState
	currentState = newState
	fmt.Printf("Transitioning from %s to %s\n", oldState, currentState)

	// Simulate work
	time.Sleep(1 * time.Second)

	// Return to original state for demonstration
	defer func() {
		fmt.Printf("Returning to state %s\n", oldState)
		currentState = oldState
	}()
}

func main() {
	currentState = "Idle"
	transitionState("Processing")
	transitionState("Complete")
}