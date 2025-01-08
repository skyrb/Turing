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

func formatUsersAsTable(users []User) string {
	output := "User Data:\n"
	output += "--------------------------------------\n"
	output += fmt.Sprintf("%-10s | %-3s | %-15s | %-3s\n", "Name", "Age", "City", "Zip")
	output += "--------------------------------------\n"
	for _, user := range users {  
        city := user.Address.City
        zip := user.Address.ZipCode
        if city == "" {
            city = "N/A"
        }
        if zip == "" {
            zip = "N/A"
        }
		output += fmt.Sprintf("%-10s | %-3d | %-15s | %-3s\n", user.Name, user.Age, city, zip)  
	}  
	output += "--------------------------------------\n"
	return output
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
                ZipCode: "60616",  
            },  
        },  
        {  
            Name: "Sudhir",  
            Age: 40,  
        },  
    }  
    fmt.Print(formatUsersAsTable(users))
}  
