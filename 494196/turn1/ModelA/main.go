package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"syscall"
	"time"
)

// GracefulError struct to hold details about the error
type GracefulError struct {
	Message string   `json:"message"` // Brief description of the error
	Reason  string   `json:"reason"`  // Detailed reason for the error
	Logs    []string `json:"logs"`    // Array of logs associated with the error
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a listener for incoming connections
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fatalError("Failed to create listener", err)
	}
	defer listener.Close()

	// Simulate a graceful shutdown by sending a signal
	go func() {
		<-time.After(2 * time.Second) // After 2 seconds, simulate a shutdown signal
		os.Kill(os.Getpid(), syscall.SIGINT)
	}()

	// Handle incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			if operationCanceled(err) {
				fmt.Println("Gracefully shutting down...")
				break
			}
			fatalError("Failed to accept connection", err)
		}
		go handleConnection(ctx, conn)
	}

	fmt.Println("Server shutting down...")
	select {
	case <-ctx.Done():
		fmt.Println("Context timeout, shutting down forcefully...")
	default:
	}
}

func handleConnection(ctx context.Context, conn net.Conn) {
	defer conn.Close()
	// Simulate a long-running operation
	time.Sleep(5 * time.Second)
	fmt.Println("Connected:", conn.RemoteAddr())
}

func operationCanceled(err error) bool {
	return context.Canceled == err || context.DeadlineExceeded == err
}

func fatalError(message, err error) {
	log.Printf("Fatal error: %s; %v\n", message, err)
	var gracefulError GracefulError
	if ge, ok := err.(*GracefulError); ok {
		gracefulError = *ge
	} else {
		gracefulError = GracefulError{
			Message: message,
			Reason:  err.Error(),
			Logs:    append(gracefulError.Logs, "No additional logs available"),
		}
	}
	// You can now handle this structured error further, such as
	// logging to a file, sending to a monitoring system, etc.
	// For demonstration, we'll just print it
	fmt.Printf("Structured error details: %+v\n", gracefulError)
	os.Exit(1)
}