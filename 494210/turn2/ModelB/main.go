
package main  
import (  
    "encoding/json"  
    "errors"  
    "log"  
    "net/http"  
    "time"  
  
    "github.com/gorilla/mux"  
)  

// Define custom error types with error messages
var (
    ErrValueRequired = errors.New("value field is required")
    ErrInvalidData  = errors.New("invalid data provided")
    ErrProcessingFailed = errors.New("processing failed")
)

type Data struct {  
    Value string `json:"value"`  
}  

type ErrorResponse struct {  
    Error   string `json:"error"`  
    Message string `json:"message"`  
}  
  
// ProcessStep functions will return an error if any issue occurs during processing

// Processing step 1: Validate the data
func validateData(data Data) error {  
    if data.Value == "" {  
        return ErrValueRequired
    }  
    return nil  
}  
  
// Processing step 2: Simulate a long-running task and return error if processing fails
func processData(data Data) (Data, error) {  
    time.Sleep(2 * time.Second) // Simulate work
    if data.Value == "invalid" {
        return Data{}, ErrInvalidData
    }
    processedData := data  
    processedData.Value = processedData.Value + "_processed"  
    return processedData, nil  
}  
  
// Final processing step: Respond with the processed data
func respondWithData(w http.ResponseWriter, data Data) {  
    w.WriteHeader(http.StatusOK)  
    json.NewEncoder(w).Encode(data)  
}  
  
// handleError function will be used to wrap around the processing steps and log errors
func handleError(w http.ResponseWriter, err error) {
    log.Printf("Error: %v", err)
    var errResponse ErrorResponse

    // Use type switch to handle custom errors and provide meaningful error messages
    switch err {
    case ErrValueRequired:
        errResponse = ErrorResponse{Error: ErrValueRequired.Error(), Message: "Value field is required."}
    case ErrInvalidData:
        errResponse = ErrorResponse{Error: ErrInvalidData.Error(), Message: "Invalid data provided."}
    case ErrProcessingFailed:
        errResponse = ErrorResponse{Error: ErrProcessingFailed.Error(), Message: "Processing failed."}
    default:
        errResponse = ErrorResponse{Error: err.Error(), Message: "An internal server error occurred."}
    }

    w.WriteHeader(http.StatusInternalServerError)  
    json.NewEncoder(w).Encode(errResponse)  
}  
  
func main() {  
    router := mux.NewRouter()  
    router.HandleFunc("/process", processHandler).Methods("POST")  
  
    log.Fatal(http.ListenAndServe(":8080", router))  
}  
  
func processHandler(w http.ResponseWriter, r *http.Request) {  
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
