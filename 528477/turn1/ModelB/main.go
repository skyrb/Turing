
package main  
import (  
    "fmt"
)  
type Customer struct {  
    ID int  
    PurchaseHistory map[string]int  
    Preferences []string  
    Location string  
}  
func main() {  
    customers := make(map[int]Customer)  
    customers[100] = Customer{  
        ID: 100,  
        PurchaseHistory: map[string]int{  
            "P1": 3,  
            "P2": 1,  
        },  
        Preferences: []string{"Electronics", "Clothing"},  
        Location: "New York",  
    }  
    customers[101] = Customer{  
        ID: 101,  
        PurchaseHistory: map[string]int{  
            "B1": 2,  
            "B2": 5,  
            "P3": 4,  
        },  
        Preferences: []string{"Books", "Home Decor"},  
        Location: "Los Angeles",  
    }  
  
    // Adding a new product to customer 100's purchase history  
    customers[100].PurchaseHistory["P4"] = 2  
  
    // Adding a new preference to customer 101's preferences  
    customers[101].Preferences = append(customers[101].Preferences, "Sports")  
  
    fmt.Println("Customer Information:")  
    for _, customer := range customers {  
        fmt.Printf("Customer ID: %d\n", customer.ID)  
        fmt.Println("Purchase History:")  
        for productID, quantity := range customer.PurchaseHistory {  
            fmt.Printf("\t%s: %d\n", productID, quantity)  
        }  
        fmt.Println("Preferences:")  
        for _, preference := range customer.Preferences {  
            fmt.Printf("\t%s\n", preference)  
        }  
        fmt.Printf("Location: %s\n", customer.Location)  
        fmt.Println("--------------------")  
    }  
}  
