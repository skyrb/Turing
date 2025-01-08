package main  
import (  
    "fmt"
)  

type Address struct {  
    City    string
    ZipCode string
}  

type User struct {  
    Name    string
    Age     int
    Address Address
}  

func main() {  
    users := []User{  
        {  
            Name: "Alice",  
            Age: 25,  
            Address: Address{  
                City:    "New York",  
                ZipCode: "10001",  
            },  
        },  
        {  
            Name: "Bob",  
            Age: 30,  
            Address: Address{  
                City:    "Los Angeles",  
                ZipCode: "90001",  
            },  
        },  
        {  
            Name: "Charlie",  
            Age: 22,  
            Address: Address{  
                City:    "Chicago",  
                ZipCode: "60601",  
            },  
        },  
    }  

    fmt.Println("Users:")
    for _, user := range users {  
        fmt.Printf("%10s: %d\n%20s %s, %s\n\n", user.Name, user.Age, "Address:", user.Address.City, user.Address.ZipCode)  
    }  
}  
  