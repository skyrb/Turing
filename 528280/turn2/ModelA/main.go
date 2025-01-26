package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"time"
)

// Department struct to represent department details in XML
type Department struct {
	ID   int    `xml:"ID"`   // use "ID" for XML
	Name string `xml:"Name"` // use "Name" for XML
}

// Employee struct with JSON and XML struct tags to handle different naming conventions
type Employee struct {
	Name       string     `json:"employee_name" xml:"EmployeeName"`
	Age        int        `json:"employee_age" xml:"EmployeeAge"`
	Salary     float64    `json:"employee_salary" xml:"EmployeeSalary"`
	JoiningDate time.Time `json:"joining_date" xml:"JoiningDate"`
	Department Department `json:"-" xml:"Department"` // Department details nested under a single Department tag in XML only
}

func main() {
	// Sample data: Slice of Employee structs
	employees := []Employee{
		{
			Name:        "John Doe",
			Age:         30,
			Salary:      75000,
			JoiningDate: time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
			Department:  Department{ID: 1, Name: "Engineering"},
		},
		{
			Name:        "Jane Smith",
			Age:         40,
			Salary:      82000,
			JoiningDate: time.Date(2017, time.March, 25, 0, 0, 0, 0, time.UTC),
			Department:  Department{ID: 2, Name: "Marketing"},
		},
		{
			Name:        "Emily Johnson",
			Age:         25,
			Salary:      55000,
			JoiningDate: time.Date(2022, time.June, 5, 0, 0, 0, 0, time.UTC),
			Department:  Department{ID: 3, Name: "Sales"},
		},
	}

	// Marshal to JSON with indentation
	jsonData, err := json.MarshalIndent(employees, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling to JSON: %v", err)
	}
	fmt.Println("JSON Output:\n", string(jsonData))

	// Marshal to XML with indentation
	xmlData, err := xml.MarshalIndent(employees, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling to XML: %v", err)
	}
	fmt.Println("\nXML Output:\n", string(xmlData))
}