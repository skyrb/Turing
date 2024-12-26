package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

var db *sql.DB
var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Simulate connecting to a database
	db, err := sql.Open("driver", "dataSourceName")
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Capture SIGINT (Ctrl+C) and SIGTERM
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, os.Kill)

	go handleShutdown(ctx, sigChan)

	// Launch concurrent worker goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(ctx, i)
	}

	// Wait for all workers to complete
	wg.Wait()
	fmt.Println("All workers have completed. Shutting down...")
}

func handleShutdown(ctx context.Context, sigChan <-chan os.Signal) {
	select {
	case <-sigChan:
		log.Println("Shutdown signal received. Canceling context...")
		ctx.Cancel()
	case <-ctx.Done():
		log.Println("Shutdown initiated from context cancellation.")
	}
}

func worker(ctx context.Context, id int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			// Handle context cancellation (shutdown)
			log.Printf("Worker %d: Shutting down gracefully...\n", id)
			// Commit or rollback any ongoing transactions here
			return
		default:
			// Simulate work
			log.Printf("Worker %d: Processing transaction...\n", id)
			// Start a transaction and handle it
			tx, err := db.BeginTx(ctx, nil)
			if err != nil {
				log.Printf("Worker %d: Error starting transaction: %v\n", id, err)
				return
			}
			defer tx.Rollback() // Default to rollback

			// Simulate transaction work
			time.Sleep(time.Second)

			// Commit the transaction if no error occurred
			if err := tx.Commit(); err == nil {
				log.Printf("Worker %d: Transaction committed successfully.\n", id)
			} else {
				log.Printf("Worker %d: Transaction failed: %v\n", id, err)
			}
		}
	}
}