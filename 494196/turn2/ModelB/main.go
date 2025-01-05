package main

import (
    "context"
    "errors"
    "fmt"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

// ShutdownOptionTag defines error management strategies for shutdown.
type ShutdownOptionTag string

const (
    // ForcefullyShutdown tag indicates that the server should be stopped immediately without waiting for ongoing requests.
    ForcefullyShutdown ShutdownOptionTag = "forcefully"

    // WaitForRequests tag indicates that the server should wait for ongoing requests to complete before shutting down.
    WaitForRequests ShutdownOptionTag = "wait"
)

// ShutdownOptions struct with struct tags for customizing error management strategies.
type ShutdownOptions struct {
    ShutdownTimeout int             `json:"shutdownTimeout"` // Duration in seconds to wait for server shutdown
    ErrorStrategy  ShutdownOptionTag `json:"errorStrategy"`  // Error management strategy for shutdown
}

func gracefulShutdown(options ShutdownOptions) error {
    // Create a new signal listener context.
    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

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

    // Perform the shutdown based on the error management strategy specified in the options.
    switch options.ErrorStrategy {
    case ForcefullyShutdown:
        // Forcefully shutdown the server without waiting for ongoing requests.
        if err := srv.Shutdown(context.Background()); err != nil {
            return fmt.Errorf("server shutdown error: %w", err)
        }
    case WaitForRequests:
        // Create a new context with a shutdown timeout.
        ctx, cancel := context.WithTimeout(ctx, time.Duration(options.ShutdownTimeout)*time.Second)
        defer cancel()

        // Gracefully shutdown the server by waiting for ongoing requests.
        if err := srv.Shutdown(ctx); err != nil {
            if errors.Is(err, context.DeadlineExceeded) {
                fmt.Println("Some requests were not completed during shutdown.")
            }
            return fmt.Errorf("server shutdown error: %w", err)
        }
    default:
        return fmt.Errorf("unknown error management strategy: %s", options.ErrorStrategy)
    }

    fmt.Println("Server stopped")
    return nil
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    // Example usage with different shutdown strategies.

    // Example 1: Graceful shutdown with a timeout of 5 seconds.
    options1 := ShutdownOptions{
        ShutdownTimeout: 5,
        ErrorStrategy:  WaitForRequests,
    }
    if err := gracefulShutdown(options1); err != nil {
        fmt.Println("Error during graceful shutdown:", err)
        os.Exit(1)
    }

    // Example 2: Forceful shutdown without waiting for ongoing requests.
    options2 := ShutdownOptions{
        ShutdownTimeout: 1,
        ErrorStrategy:  ForcefullyShutdown,
    }
    if err := gracefulShutdown(options2); err != nil {
        fmt.Println("Error during graceful shutdown:", err)
        os.Exit(1)
    }
}