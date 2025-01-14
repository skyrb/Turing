package main

import (
	"fmt"
	"reflect"
	"strings"
)

type UserProfile struct {
	Name     string
	Email    string
	Age      int
	Address  Address
	PhoneNumber string // New field added in a later version
}

// Address is now a nested struct
type Address struct {
	Street string
	City   string
	ZIP    string
}

func SerializeProfile(profile UserProfile) map[string]interface{} {
	profileMap := make(map[string]interface{})
	reflectValue := reflect.ValueOf(profile)

	for i := 0; i < reflectValue.NumField(); i++ {
		field := reflectValue.Type().Field(i)
		val := reflectValue.Field(i)
		profileMap[field.Name] = serializeField(val)
	}

	return profileMap
}

func serializeField(val reflect.Value) interface{} {
	switch val.Kind() {
	case reflect.Struct:
		// For nested structs, serialize each field recursively
		result := make(map[string]interface{})
		for i := 0; i < val.NumField(); i++ {
			field := val.Type().Field(i)
			result[field.Name] = serializeField(val.Field(i))
		}
		return result
	default:
		// For non-struct fields, return the value directly
		return val.Interface()
	}
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
			err := deserializeField(field, reflect.ValueOf(value))
			if err != nil {
				return profile, err
			}
		}
	}

	// Handle missing fields and populate with default values (as before)
	// ...

	return profile, nil
}

func deserializeField(field reflect.Value, value reflect.Value) error {
	switch field.Kind() {
	case reflect.Struct:
		// For nested structs, deserialize each field recursively
		switch value.Kind() {
		case reflect.String:
			// Handle old format with single string Address
			addressParts := strings.Split(value.String(), ",")
			if len(addressParts) != 3 {
				return fmt.Errorf("invalid address format: %s", value.String())
			}
			address := Address{
				Street: addressParts[0],
				City:   addressParts[1],
				ZIP:    addressParts[2],
			}
			field.Set(reflect.ValueOf(address))
		case reflect.Map:
			// Handle new format with nested map Address
			address := Address{}
			for _, subKey := range value.MapKeys() {
				subField := field.FieldByName(subKey.String())
				if !subField.IsValid() {
					continue
				}
				if subField.CanSet() {
					subVal := value.MapIndex(subKey)
					err := deserializeField(subField, subVal)
					if err != nil {
						return err
					}
				}
			}
			field.Set(reflect.ValueOf(address))
		default:
			return fmt.Errorf("invalid value type for struct field: %s", value.Kind())
		}
	default:
		// For non-struct fields, convert and set the value directly
		fieldType := field.Type()
		if value.Type().ConvertibleTo(fieldType) {
			field.Set(value.Convert(fieldType))
		} else {
			return fmt.Errorf("cannot convert %s to %s", value.Type(), fieldType)
		}
	}
	return nil
}

func main() {
	// Example profile with both old and new address formats
	profile := UserProfile{
		Name:    "Alice",
		Email:   "alice@example.com",
		Age:     30,
		Address: Address{Street: "Main St", City: "New York", ZIP: "10001"},
		PhoneNumber: "123-456-7890",
	}

	// Serialize
	profileMap := SerializeProfile(profile)
	fmt.Println("Serialized Profile:", profileMap)

	// Simulate deserialization with an older version that uses a single string for Address
	olderProfileMap := map[string]interface{}{
		"Name":  "Bob",
		"Email": "bob@example.com",
		"Age":   25,
		"Address": "Oak St,Los Angeles,90001",
	}

	deserializedProfile, err := DeserializeProfile(olderProfileMap)
	if err != nil {
		fmt.Println("Deserialization error:", err)
	} else {
		fmt.Println("Deserialized Profile:", deserializedProfile)
	}
}
