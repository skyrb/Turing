package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"time"
)

// Employee struct with JSON and XML struct tags
type Employee struct {
	Name         string  `json:"employee_name" xml:"EmployeeName"`
	Age          int     `json:"employee_age" xml:"EmployeeAge"`
	Department   Department `json:"employee_department" xml:"Department"`
	Salary       float64 `json:"employee_salary" xml:"EmployeeSalary"`
	JoiningDate time.Time `json:"joining_date" xml:"JoiningDate"`
}

// Department struct for nested XML tags
type Department struct {
	ID   int    `xml:"ID"`
	Name string `xml:"Name"`
}

func main() {
	// Sample data: slice of Employee structs
	employees := []Employee{
		{
			Name:         "John Doe",
			Age:          30,
			Department:   Department{ID: 1, Name: "Engineering"},
			Salary:       75000,
			JoiningDate: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			Name:         "Jane Smith",
			Age:          40,
			Department:   Department{ID: 2, Name: "Marketing"},
			Salary:       82000,
			JoiningDate: time.Date(2015, 6, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			Name:         "Emily Johnson",
			Age:          25,
			Department:   Department{ID: 3, Name: "Sales"},
			Salary:       55000,
			JoiningDate: time.Date(2020, 3, 10, 0, 0, 0, 0, time.UTC),
		},
	}

	// Marshal struct to JSON with indentation
	jsonData, err := json.MarshalIndent(employees, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}
	fmt.Println("JSON Output:")
	fmt.Println(string(jsonData))

	// Marshal struct to XML with indentation
	xmlData, err := xml.MarshalIndent(employees, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling XML: %v", err)
	}
	fmt.Println("\nXML Output:")
	fmt.Println(string(xmlData))
}