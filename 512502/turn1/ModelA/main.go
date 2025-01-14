package main

import (
	"fmt"
	"encoding/json"
	"reflect"
)

type UserProfile struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	PhoneNumber string `json:"phoneNumber,omitempty"` // New field
}

func unmarshalWithReflection(data []byte, v interface{}) error {
	// Create a new struct instance to fill
	dest := reflect.New(reflect.ValueOf(v).Elem().Type())

	// Unmarshal into the new struct
	if err := json.Unmarshal(data, dest.Interface()); err != nil {
		return err
	}

	// Copy values from the new struct to the original one
	src := reflect.ValueOf(v).Elem()
	for _, field := range src.FieldByNameFunc(func(name string) reflect.StructField {
		return dest.Elem().FieldByName(name)
	}) {
		srcField := field.Field(0)
		destField := dest.Elem().FieldByName(field.Name)

		if srcField.CanSet() && destField.CanInterface() {
			srcField.Set(destField.Interface())
		}
	}

	return nil
}

func main() {
	oldData := `{"name": "John Doe", "email": "johndoe@example.com", "age": 30}`
	newData := `{"name": "Jane Smith", "email": "janesmith@example.com", "age": 25, "phoneNumber": "555-0123"}`

	// Deserialize old and new data using reflection
	var oldProfile, newProfile UserProfile
	unmarshalWithReflection([]byte(oldData), &oldProfile)
	unmarshalWithReflection([]byte(newData), &newProfile)

	// Print profiles
	fmt.Println("Old Profile:", oldProfile)
	fmt.Println("New Profile:", newProfile)
}
