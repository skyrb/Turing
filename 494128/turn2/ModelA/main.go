package main

import (
	"fmt"
	"net/http"
	"strings"

	"go.uber.org/zap"
)

var logger *zap.Logger

func main() {
	// Configure logging
	logger, _ = zap.NewProduction()
	defer logger.Sync()

	// Define handlers and middleware
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserIDFromContext(r.Context())
	if err != nil {
		handleError(w, err)
		return
	}

	fmt.Fprintf(w, "Hello, user %d!", userID)
}

// Middleware function to retrieve user ID from the request context
func userIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate retrieving user ID from the request context or any other source
		userID := 123

		// Create a new context with the user ID
		ctx := r.Context()
		ctx = context.WithValue(ctx, userIDContextKey, userID)

		// Log entry for this middleware
		logger.Info("UserIDMiddleware", zap.String("path", r.URL.Path), zap.Int("userID", userID))

		// Call the next handler with the updated context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Function to retrieve the user ID from the request context
func getUserIDFromContext(ctx context.Context) (userID int, err error) {
	userIDValue := ctx.Value(userIDContextKey)
	if userIDValue == nil {
		err = errors.New("user ID not found in context")
		return
	}

	userID, ok := userIDValue.(int)
	if !ok {
		err = errors.New("invalid user ID type in context")
		return
	}

	return userID, nil
}

// Function to handle errors and send appropriate responses
func handleError(w http.ResponseWriter, err error) {
	// Structured logging for error
	logger.Error("Internal Server Error", zap.Error(err), zap.String("path", w.Request().URL.Path))
	
	// Send a JSON response to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	
	errResponse := map[string]string{"error": err.Error()}
	fmt.Fprintf(w, "%v", errResponse)
}