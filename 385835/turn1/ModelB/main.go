package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// handleRequest is a Goroutine that handles each incoming HTTP request
func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	// Create a signal channel to receive termination signals
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// Start the HTTP server on port 8080
	go func() {
		log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(handleRequest)))
	}()

	// Wait for a termination signal
	sig := <-sigCh
	log.Println("Received signal:", sig)

	// Gracefully shut down the server
	fmt.Println("Shutting down...")
	// Add your shutdown logic here, e.g., closing database connections, etc.

	os.Exit(0)
}
