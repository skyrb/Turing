package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func logRequest(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	defer func() {
		end := time.Now()
		duration := end.Sub(start)
		log.Printf("Request completed for %s in %s. Method: %s, Status: %d",
			r.RemoteAddr, duration, r.Method, http.StatusOK)
	}()

	// Parse the "delay" query parameter
	delayStr := r.URL.Query().Get("delay")
	delay := 0 // Default delay
	if delayStr != "" {
		if parsedDelay, err := strconv.Atoi(delayStr); err == nil {
			delay = parsedDelay
		} else {
			log.Printf("Invalid delay value: %s. Using default delay of 0 seconds.", delayStr)
		}
	}

	// Simulate some work
	time.Sleep(time.Duration(delay) * time.Second)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", logRequest)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
