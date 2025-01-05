package main

import (
    "context"
    "fmt"
    "net/http"
    "os"
    "os/signal"
    "syscall"
	"time"
)

type ShutdownOptions struct {
    ShutdownTimeout int // Duration in seconds to wait for server shutdown
}

func gracefulShutdown(options ShutdownOptions) error {
    // Create a new signal listener context.
    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)

    // Set up a channel to listen for OS signals.
    signalChan := make(chan os.Signal, 1)
    signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

    // Create an HTTP server and set up its handle.
    srv := &http.Server{Addr: ":8080"}
    http.HandleFunc("/", helloHandler)

    // Start listening for requests on a separate goroutine.
    go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            fmt.Println("HTTP server error:", err)
        }
    }()

    fmt.Println("Server is running on port 8080")

    // Wait for an OS signal.
    <-signalChan

    // Cancel the signal listener context.
    cancel()

    // Create a new context with a shutdown timeout.
    ctx, cancel = context.WithTimeout(ctx, time.Duration(options.ShutdownTimeout)*time.Second)
    defer cancel()

    // Gracefully shutdown the server.
    if err := srv.Shutdown(ctx); err != nil {
        fmt.Println("Server shutdown error:", err)
    }

    fmt.Println("Server stopped")

    return nil
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    // Example usage with configurable options
    options := ShutdownOptions{
        ShutdownTimeout: 5,
    }
    if err := gracefulShutdown(options); err != nil {
        fmt.Println("Error during graceful shutdown:", err)
        os.Exit(1)
    }
}  