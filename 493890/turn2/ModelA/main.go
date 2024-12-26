package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var db *sql.DB
var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Simulate connecting to a database
	// Replace "driver" and "dataSourceName" with actual database details
	db, err := sql.Open("driver", "dataSourceName")
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Capture SIGINT (Ctrl+C) and SIGTERM
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go handleShutdown(sigChan, cancel)

	// Launch concurrent worker goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(ctx, i)
	}

	// Wait for all workers to complete
	wg.Wait()
	fmt.Println("All workers have completed. Shutting down...")
}

func handleShutdown(sigChan <-chan os.Signal, cancel context.CancelFunc) {
	select {
	case <-sigChan:
		log.Println("Shutdown signal received. Canceling context...")
		cancel() // Trigger context cancellation
	}
}

func worker(ctx context.Context, id int) {
	defer wg.Done()

	// Use context.WithTimeout to set a deadline
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			// Handle context cancellation (shutdown)
			log.Printf("Worker %d: Shutting down due to context deadline...\n", id)
			// Rollback any ongoing transactions here
			return
		default:
			// Simulate work
			log.Printf("Worker %d: Processing transaction...\n", id)

			// Simulate a transaction (Replace with real database logic)
			tx, err := db.BeginTx(ctx, nil)
			if err != nil {
				log.Printf("Worker %d: Error starting transaction: %v\n", id, err)
				return
			}
			defer tx.Rollback() // Default to rollback

			// Simulate transaction work (e.g., updating a record)
			time.Sleep(2*time.Second)

			// Commit the transaction if no error occurred
			if err := tx.Commit(); err == nil {
				log.Printf("Worker %d: Transaction committed successfully.\n", id)
			} else {
				log.Printf("Worker %d: Transaction failed: %v\n", id, err)
			}
		}
	}
}