package main

import (
    "log"
    "net/http"
)

// RecoveryMiddleware is a middleware that recovers from panics and logs them.
func RecoveryMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                log.Printf("Recovered from panic: %v", err)
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            }
        }()
        next.ServeHTTP(w, r)
    })
}

// DangerousHandler simulates a handler that may panic.
func DangerousHandler(w http.ResponseWriter, r *http.Request) {
    // Simulating a panic for demonstration purposes
    panic("something went wrong!")
}

func main() {
    mux := http.NewServeMux()
    mux.Handle("/", RecoveryMiddleware(http.HandlerFunc(DangerousHandler)))

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", mux); err != nil {
        log.Fatalf("Could not start server: %v", err)
    }
}