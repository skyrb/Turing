package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// ExternalServiceResponse represents the response from an external service
type ExternalServiceResponse struct {
	Message string `json:"message"`
}

// ExternalServiceCall makes an HTTP request to an external service
func ExternalServiceCall(ctx context.Context, url string) (*ExternalServiceResponse, error) {
	client := &http.Client{
		Timeout: 5 * time.Second, // Set timeout for the client
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("external service returned status %d", res.StatusCode)
	}

	var response ExternalServiceResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Handler for the /example endpoint
func exampleHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second) // Set server-side timeout
	defer cancel()

	url := "https://api.example.com/slow" // Replace with your external service URL
	response, err := ExternalServiceCall(ctx, url)

	if err != nil {
		switch err {
		case context.DeadlineExceeded:
			http.Error(w, "Request timed out", http.StatusGatewayTimeout)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/example", exampleHandler)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}