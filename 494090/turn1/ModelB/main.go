package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Sample error handling
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(errorMessage(err, "Failed to start HTTP server"))
	}
}

func errorMessage(err error, message string) string {
	if err != nil {
		return fmt.Sprintf("%s: %v", message, err)
	}
	return message
}