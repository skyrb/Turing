package main  
import (  
    "encoding/json"  
    "fmt"  
    "reflect"  
)  

// UserProfile represents the user profile data  
type UserProfile struct {  
    Name        string `json:"name"`  
    Email       string `json:"email"`  
    Age         int    `json:"age"`  
    PhoneNumber string `json:"phone_number,omitempty"` // New field  
}  

// MarshalJSON implements the json.Marshaler interface.  
func (up UserProfile) MarshalJSON() ([]byte, error) {  
    m := make(map[string]interface{})  
    v := reflect.ValueOf(up)  
    t := reflect.TypeOf(up)  

    for i := 0; i < v.NumField(); i++ {  
        field := t.Field(i)  
        tag := field.Tag.Get("json")  
        if tag == "-" {  
            continue // Skip fields with "-" json tag  
        }  

        if tag == "omitempty" && v.Field(i).IsZero() {  
            continue // Skip fields with "omitempty" json tag and zero value  
        }  

        key := tag  
        if key == "" {  
            key = field.Name  
        }  

        value := v.Field(i).Interface()  
        m[key] = value  
    }  

    return json.Marshal(m)  
}  

// UnmarshalJSON implements the json.Unmarshaler interface.  
func (up *UserProfile) UnmarshalJSON(data []byte) error {  
    var m map[string]interface{}  
    if err := json.Unmarshal(data, &m); err != nil {  
        return err  
    }  

    v := reflect.ValueOf(up).Elem()  
    t := reflect.TypeOf(*up)  

    for i := 0; i < v.NumField(); i++ {  
        field := t.Field(i)  
        tag := field.Tag.Get("json")  
        if tag == "-" {  
            continue // Skip fields with "-" json tag  
        }  

        key := tag  
        if key == "" {  
            key = field.Name  
        }  

        if value, ok := m[key]; ok {  
            // Found the field in the decoded JSON object  
            if err := convertAndSetField(v.Field(i), value); err != nil {  
                return err  
            }  
        }  
    }  

    return nil  
}  

func convertAndSetField(field reflect.Value, value interface{}) error {  
    // Check if the value can be directly assigned to the field  
    if field.Type().AssignableTo(reflect.TypeOf(value)) {  
        field.Set(reflect.ValueOf(value))  
        return nil  
    }  

    // Try converting the value to the field's type  
    if convertedValue, ok := convertValue(value, field.Type()); ok {  
        field.Set(convertedValue)  
        return nil  
    }  

    return fmt.Errorf("cannot convert value to field type")  
}  

func convertValue(value interface{}, toType reflect.Type) (reflect.Value, bool) {  
    switch toType.Kind() {  
    case reflect.Int:  
        switch v := value.(type) {  
        case float64:  
            return reflect.ValueOf(int(v)), true  
        }  
    case reflect.String:  
        switch v := value.(type) {  
        case string:  
            return reflect.ValueOf(v), true  
        }  
    }  

    return reflect.Value{}, false // Conversion failed  
}  