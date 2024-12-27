package main

import (
	"fmt"
	"strings"
)

type Transaction struct {
	stage    string
	original string
}

func (t *Transaction) Execute() error {
	defer t.handleError() // Schedule to handle any errors
	t.stage = "Starting transaction..."

	if err := t.step1(); err != nil {
		return err
	}

	t.stage = "Executing step 2..."

	if err := t.step2(); err != nil {
		return err
	}

	t.stage = "Transaction completed."
	return nil
}

func (t *Transaction) handleError() {
	// Handle errors by rolling back changes and reporting errors
	if err := recover(); err != nil {
		t.stage = "Rolling back changes due to error."
		fmt.Printf("Error occurred: %v\n", err)

		// Rollback state
		if strings.Contains(t.stage, "starting") || strings.Contains(t.stage, "executing") {
			fmt.Println("Restoring original state.")
			t.stage = "Original state restored."
		}
	}
}

func (t *Transaction) step1() error {
	t.stage = "Executing step 1..."

	fmt.Println("Step 1 completed successfully.")
	return nil
}

func (t *Transaction) step2() error {
	t.stage = "Executing step 2..."

	// Simulate an error by intentionally panicking
	panic("Step 2 failed due to unexpected condition.")
	return nil
}

func main() {
	transaction := Transaction{original: "Initial state"}

	if err := transaction.Execute(); err != nil {
		fmt.Println("Transaction failed:", err)
	} else {
		fmt.Println("Transaction completed successfully.")
	}

	fmt.Println("Final stage:", transaction.stage)
}