package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type TokenBucket struct {
	capacity int
	tokens   int
	mutex    sync.Mutex
	fillRate time.Duration
	lastFill time.Time
}

func NewTokenBucket(capacity int, fillRate time.Duration) *TokenBucket {
	return &TokenBucket{
		capacity:  capacity,
		fillRate: fillRate,
	}
}

func (tb *TokenBucket) GetToken() bool {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()

	// Determine how many tokens have been added since the last time
	now := time.Now()
	tokens := (now.Sub(tb.lastFill) / tb.fillRate).Int()

	tb.lastFill = now
	tb.tokens += tokens

	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}

	if tb.tokens == 0 {
		return false
	}

	tb.tokens--
	return true
}

var buckets map[string]*TokenBucket

func main() {
	buckets = make(map[string]*TokenBucket)
	l := http.ListenAndServe(":8080", nil)
	log.Fatal(l.Err())
}

func userRole(userID string) (string, error) {
	// Simulating user role lookup from a database
	switch userID {
	case "user1":
		return "standard", nil
	case "user2":
		return "admin", nil
	default:
		return "", fmt.Errorf("user not found: %s", userID)
	}
}

func rateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		userID, ok := ctx.Value("userID").(string)
		if !ok {
			log.Printf("Error getting userID from context")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		role, err := userRole(userID)
		if err != nil {
			log.Printf("Error retrieving role for userID %s: %v", userID, err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		var limit int
		var interval time.Duration
		switch role {
		case "standard":
			limit = 100 // 100 requests
			interval = 1 * time.Minute
		case "admin":
			limit = 1000 // 1000 requests
			interval = 1 * time.Minute
		default:
			log.Printf("Unknown role: %s", role)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Get the token bucket for this userID, creating a new one if necessary
		bucketKey := fmt.Sprintf("%s:%s", userID, interval.String())
		bucket, ok := buckets[bucketKey]
		if !ok {
			bucket = NewTokenBucket(limit, interval)
			buckets[bucketKey] = bucket
		}

		if bucket.GetToken() {
			// Proceed with the request
			next.ServeHTTP(w, r)
			return
		}

		// Rate limit exceeded
		http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
		w.Header().Set("X-RateLimit-Limit", fmt.Sprint(limit))
		w.Header().Set("X-RateLimit-Remaining", fmt.Sprint(0))
		w.Header().Set("X-RateLimit-Reset", fmt.Sprint(time.Unix(time.Now().Unix()+int64(interval.Seconds()), 0).Format(time.RFC3339)))
	})
}

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}