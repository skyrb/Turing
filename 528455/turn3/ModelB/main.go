package main

import (
	"fmt"
)

// Person struct
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

// InterfaceToStrings takes a slice of interface{} and returns a slice of formatted strings
func InterfaceToStrings(slice []interface{}) []string {
	var formattedStrings []string

	for _, element := range slice {
		var str string
		switch val := element.(type) {
		case string:
			str = val
		case int:
			str = fmt.Sprintf("%d", val)
		case Person:
			str = val.String()
		default:
			str = fmt.Sprintf("%v", val)
		}
		formattedStrings = append(formattedStrings, str)
	}
	return formattedStrings
}

func main() {
	slice := []interface{}{42, "hello", Person{Name: "John", Age: 30}}
	for _, str := range InterfaceToStrings(slice) {
		fmt.Println(str)
	}
}