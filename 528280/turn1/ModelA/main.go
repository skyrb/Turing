package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Employee struct with JSON struct tags to ensure custom field names
type Employee struct {
	Name       string  `json:"employee_name"`
	Age        int     `json:"employee_age"`
	Department string  `json:"employee_department"`
	Salary     float64 `json:"employee_salary"`
}

func main() {
	// Sample data: slice of Employee structs
	employees := []Employee{
		{Name: "John Doe", Age: 30, Department: "Engineering", Salary: 75000},
		{Name: "Jane Smith", Age: 40, Department: "Marketing", Salary: 82000},
		{Name: "Emily Johnson", Age: 25, Department: "Sales", Salary: 55000},
	}

	// Marshal struct to JSON with indentation for human readability
	jsonData, err := json.MarshalIndent(employees, "", "    ") // 4 spaces indentation
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	// Print the JSON output
	fmt.Println(string(jsonData))
}