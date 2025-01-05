package main  
import (  
    "context"
    "fmt"
	"os"
    "database/sql"
    _ "github.com/lib/pq"  
)  
// Function that takes a context and an event as input and returns an error  
func HandleRequest(ctx context.Context, event interface{}) error {  
    // Open a database connection using your preferred database driver  
    db, err := sql.Open("postgres", "user=postgres password=password dbname=postgres sslmode=disable")  
    if err != nil {  
        return fmt.Errorf("error opening database: %w", err)  
    }  
    // Close the database connection at the end of the function execution using 'defer'  
    defer db.Close()  
    // Perform database operations here  
    _, err = db.Exec("INSERT INTO users (name) VALUES ($1)", "John Doe")  
    if err != nil {  
        return fmt.Errorf("error executing database query: %w", err)  
    }  
    return nil  
}


func main() {  
    file, err := os.Create("example.txt")  
    if err != nil {  
        fmt.Println("Error creating file:", err)  
        return  
    }  
    defer file.Close()  
    _, err = file.WriteString("Hello, World!")  
    if err != nil {  
        fmt.Println("Error writing to file:", err)  
        return  
    }  
    fmt.Println("Data written successfully to example.txt")  
}