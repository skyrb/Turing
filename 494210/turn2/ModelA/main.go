package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type UserData struct {
	Name string `json:"name"`
}

type ProcessedData struct {
	Greeting string `json:"greeting"`
}

type PipelineError struct {
	Stage   string
	Message string
	Inner   error
}

func wrapError(stage string, err error) PipelineError {
	return PipelineError{stage, err.Error(), err}
}

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err)
}

func validateData(data UserData) error {
	if strings.TrimSpace(data.Name) == "" {
		return fmt.Errorf("name field is required")
	}
	return nil
}

func processData(data UserData) ProcessedData {
	return ProcessedData{Greeting: fmt.Sprintf("Hello, %s!", data.Name)}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/process", processHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func processHandler(w http.ResponseWriter, r *http.Request) {
	var userData UserData
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		handleError(w, wrapError("POST body", err))
		return
	}

	if err := json.Unmarshal(body, &userData); err != nil {
		handleError(w, wrapError("Unmarshaling", err))
		return
	}

	if err := validateData(userData); err != nil {
		handleError(w, wrapError("Validation", err))
		return
	}

	processedData := processData(userData)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(processedData)
}