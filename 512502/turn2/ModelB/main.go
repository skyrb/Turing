package main  
import (  
    "fmt"
    "reflect"
)

type UserProfile struct {  
    Name        string
    Email       string
    Age         int
    PhoneNumber string `default:"N/A"`
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

    reflectValue := reflect.ValueOf(&profile).Elem()
    reflectType := reflectValue.Type()

    for i := 0; i < reflectType.NumField(); i++ {  
        field := reflectType.Field(i)
        fieldValue := reflectValue.Field(i)

        // Check if the field exists in the deserialized map
        if val, ok := profileMap[field.Name]; ok {  
            if fieldValue.CanSet() {  
                fieldType := field.Type
                val := reflect.ValueOf(val)
                if val.Type().ConvertibleTo(fieldType) {  
                    fieldValue.Set(val.Convert(fieldType))
                } else {  
                    return profile, fmt.Errorf("cannot convert %s to %s", val.Type(), fieldType)
                }
            }
        } else {  
            // Field is missing from the map, set it to its default value
            if defaultValue, ok := field.Tag.Lookup("default"); ok {  
                if err := setDefaultValue(fieldValue, defaultValue); err != nil {  
                    return profile, err
                }
            }
        }
    }

    return profile, nil
}

func setDefaultValue(field reflect.Value, defaultValue string) error {  
    switch field.Kind() {  
    case reflect.String:  
        field.SetString(defaultValue)
    default:  
        return fmt.Errorf("unsupported default value type for field %s", field.Type())
    }
    return nil
}

func main() {  
    profile := UserProfile{Name: "Alice", Email: "alice@example.com", Age: 30}

    profileMap := SerializeProfile(profile)
    fmt.Println("Serialized Profile:", profileMap)

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
