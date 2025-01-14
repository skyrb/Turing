package main

import (
	"fmt"
	"reflect"
)

type UserProfile struct {
	Name     string
	Email    string
	Age      int
	PhoneNumber string // New field added in a later version
}

func SerializeProfile(profile UserProfile) map[string]interface{} {
	profileMap := make(map[string]interface{})
	reflectValue := reflect.ValueOf(profile)

	for i := 0; i < reflectValue.NumField(); i++ {
		field := reflectValue.Type().Field(i)
		profileMap[field.Name] = reflectValue.Field(i).Interface()
	}

	return profileMap
}

func DeserializeProfile(profileMap map[string]interface{}) (UserProfile, error) {
	var profile UserProfile

	// Use reflection to set struct fields from map
	reflectValue := reflect.ValueOf(&profile).Elem()
	for key, value := range profileMap {
		field := reflectValue.FieldByName(key)
		if !field.IsValid() {
			// Field not found in struct, can ignore or handle accordingly
			continue
		}

		if field.CanSet() {
			fieldType := field.Type()
			val := reflect.ValueOf(value)

			// Convert value to the appropriate type if needed
			if val.Type().ConvertibleTo(fieldType) {
				field.Set(val.Convert(fieldType))
			} else {
				return profile, fmt.Errorf("cannot convert %s to %s", val.Type(), fieldType)
			}
		}
	}

	// Dynamically check for missing fields and populate with default values
	for i := 0; i < reflectValue.NumField(); i++ {
		field := reflectValue.Type().Field(i)
		if _, ok := profileMap[field.Name]; !ok {
			// Missing field
			switch field.Type {
			case reflect.TypeOf(""):
				profileMap[field.Name] = "N/A" // String default
			case reflect.TypeOf(0):
				profileMap[field.Name] = 0     // Int default
			default:
				profileMap[field.Name] = nil   // Default to nil for other types
			}
			
			// Set the default value back to the struct
			setDefaultValue(reflectValue, field.Name, reflect.ValueOf(profileMap[field.Name]).Convert(field.Type))
		}
	}

	return profile, nil
}

func setDefaultValue(v reflect.Value, fieldName string, value reflect.Value) {
	if field := v.FieldByName(fieldName); field.IsValid() && field.CanSet() {
		field.Set(value)
	}
}

func main() {
	// Example profile with a new field
	profile := UserProfile{Name: "Alice", Email: "alice@example.com", Age: 30, PhoneNumber: "123-456-7890"}

	// Serialize
	profileMap := SerializeProfile(profile)
	fmt.Println("Serialized Profile:", profileMap)

	// Simulate deserialization with an older version that lacks PhoneNumber
	olderProfileMap := map[string]interface{}{
		"Name":  "Bob",
		"Email": "bob@example.com",
		"Age":   25,
	}

	deserializedProfile, err := DeserializeProfile(olderProfileMap)
	if err != nil {
		fmt.Println("Deserialization error:", err)
	} else {
		fmt.Println("Deserialized Profile:", deserializedProfile)
	}
}