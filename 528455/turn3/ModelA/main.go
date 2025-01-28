package main

import (
	"fmt"
)

// Person struct with Name and Age fields
type Person struct {
	Name string
	Age  int
}

// String method to satisfy the fmt.Stringer interface
func (p Person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

// FormatElements function formats each element in a slice of mixed data types
func FormatElements(elements []interface{}) string {
	var result string
	for _, element := range elements {
		switch v := element.(type) {
		case int:
			result += fmt.Sprintf("Integer: %d", v)
		case string:
			result += fmt.Sprintf("String: %s", v)
		case Person:
			result += fmt.Sprintf("Person: %s", v) // Calls the String() method of Person
		default:
			result += fmt.Sprintf("Unknown: %v", v)
		}
		result += "\n"
	}
	return result
}

func main() {
	// Create a slice containing mixed data types
	mixedData := []interface{}{
		42,
		"hello",
		Person{Name: "John", Age: 30},
		3.14, // An example of an unsupported type in this example
	}

	// Format and print the elements of the mixedData slice
	formattedOutput := FormatElements(mixedData)
	fmt.Println(formattedOutput)
}