package main

import "fmt"

// Result is an interface defining the methods to retrieve data and error.
type Result interface {
	GetValue() interface{}
	GetError() error
	GetMetadata() interface{}
}

// CustomResult struct implements the Result interface to encapsulate value, error, and metadata.
type CustomResult struct {
	value    interface{}
	error    error
	metadata interface{}
}

func (r *CustomResult) GetValue() interface{} {
	return r.value
}

func (r *CustomResult) GetError() error {
	return r.error
}

func (r *CustomResult) GetMetadata() interface{} {
	return r.metadata
}

// CustomError is a custom error type.
type CustomError string

func (e CustomError) Error() string {
	return string(e)
}

// Example function using CustomResult to return data, error, and metadata.
func fetchData(source string) Result {
	switch source {
	case "db":
		// Simulate fetching data from a database
		data := "Data from database"
		return &CustomResult{value: data, error: nil, metadata: map[string]interface{}{"timestamp": "2023-10-28T10:00:00Z"}}
	case "api":
		// Simulate an API request that fails
		return &CustomResult{value: nil, error: CustomError("API call failed"), metadata: map[string]interface{}{"status_code": 404}}
	default:
		return &CustomResult{value: nil, error: fmt.Errorf("unknown source: %s", source), metadata: map[string]interface{}{"status": "error"}}
	}
}

func main() {
	dbResult := fetchData("db")
	if err := dbResult.GetError(); err != nil {
		fmt.Println("Error:", err)
	} else {
		data := dbResult.GetValue()
		metadata := dbResult.GetMetadata()
		fmt.Println("Data:", data)
		fmt.Println("Metadata:", metadata)
	}

	apiResult := fetchData("api")
	if err := apiResult.GetError(); err != nil {
		fmt.Println("Error:", err)
	} else {
		data := apiResult.GetValue()
		metadata := apiResult.GetMetadata()
		fmt.Println("Data:", data)
		fmt.Println("Metadata:", metadata)
	}
}