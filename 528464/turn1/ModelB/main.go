
package main  
import (  
    "fmt"
)  

func main() {  
    query := buildFilterQuery(30, "New York")
    fmt.Println(query)
}
func buildFilterQuery(age int, location string) string {
    query := "SELECT * FROM users WHERE 1=1"
    if age > 0 {
        query += fmt.Sprintf(" AND age > %d", age)
    }
    if location != "" {
        query += fmt.Sprintf(" AND location = '%s'", location)
    }
    return query
}
