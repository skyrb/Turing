

package main

import (
        "encoding/json"
        "fmt"
        "log"
        "net/http"
        "time"

        "github.com/gorilla/mux"
)

type Data struct {
        Value string `json:"value"`
}

type ErrorResponse struct {
        Error   string `json:"error"`
        Message string `json:"message"`
}

func handleError(w http.ResponseWriter, err error) {
        if err != nil {
                w.WriteHeader(http.StatusInternalServerError)
                json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error(), Message: "Internal Server Error"})
        }
}

// Processing step 1: Validate the data
func validateData(data Data) error {
        if data.Value == "" {
                return fmt.Errorf("value field is required")
        }
        return nil
}

// Processing step 2: Simulate a long-running task
func processData(data Data) (Data, error) {
        time.Sleep(2 * time.Second) // Simulate work
        processedData := data
        processedData.Value = processedData.Value + "_processed"
        return processedData, nil
}

// Final processing step: Respond with the processed data
func respondWithData(w http.ResponseWriter, data Data) {
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(data)
}

// Main function to set up the API and routing
func main() {
        router := mux.NewRouter()
        router.HandleFunc("/process", processHandler).Methods("POST")

        log.Fatal(http.ListenAndServe(":8080", router))
}

func processHandler1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Processing request")
}

// processHandler handles incoming requests, processes the data through the pipeline, and handles errors.
func processHandler(w http.ResponseWriter, r *http.Request) {
        // Decode the request body into a Data struct
        var data Data
        if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
                handleError(w, err)
                return
        }

        // Step 1: Validate the data
        if err := validateData(data); err != nil {
                handleError(w, err)
                return
        }

        // Step 2: Simulate processing the data
        processedData, err := processData(data)
        if err != nil {
                handleError(w, err)
                return
        }

        // Step 3: Respond with the processed data
        respondWithData(w, processedData)
}

