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
	"strconv"
)

type ShutdownOptions struct {
    ShutdownTimeout int // Duration in seconds to wait for server shutdown
}

type MyError struct {
    error
    Tags map[string]string `json:"tags,omitempty"`
}

func NewMyError(err error, tags map[string]string) error {
    return &MyError{error: err, Tags: tags}
}

func handleError(err error, ctx context.Context) {
    if e, ok := err.(*MyError); ok {
        if e.Tags["strategy"] == "ignore" {
            fmt.Println("Ignoring error:", e.error)
            return
        }
        if e.Tags["strategy"] == "retry" {
            if retries, err := strconv.Atoi(e.Tags["retries"]); err == nil {
                for i := 0; i < retries; i++ {
                    fmt.Println("Retrying error:", e.error)
                    select {
                    case <-ctx.Done():
                        fmt.Println("Retry failed due to context cancellation.")
                        return
                    case <-time.After(time.Second):
                    }
                }
                fmt.Println("Max retries reached:", e.error)
            } else {
                fmt.Println("Invalid retries count specified:", e.error)
            }
            return
        }
    }
    fmt.Println("Default error handling:", err)
}

func gracefulShutdown(options ShutdownOptions) error {
    // Create a new signal listener context.
    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)

    // Set up a channel to listen for OS signals.
    signalChan := make(chan os.Signal, 1)
    signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

    // Create an HTTP server and set up its handler.
    srv := &http.Server{Addr: ":8080"}
    http.HandleFunc("/", helloHandler)

    // Start listening for requests on a separate goroutine.
    go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            handleError(NewMyError(err, map[string]string{"strategy": "retry", "retries": "3"}), ctx)
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
        handleError(NewMyError(err, map[string]string{"strategy": "ignore"}), ctx)
    }

    fmt.Println("Server stopped")

    return nil
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    // Simulate an error with a retry strategy
    if r.Method == "POST" {
        if err := errors.New("simulated post error"); err != nil {
            handleError(NewMyError(err, map[string]string{"strategy": "retry", "retries": "2"}), context.Background())
        }
    }
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