package main  
import (  
        "fmt"
)

// Customer struct  
type Customer struct {  
        ID             string  
        PurchaseHistory PurchaseHistory  
        Preferences    []string  
        Location       string  
}  

// PurchaseHistory struct  
type PurchaseHistory map[string]int  

// CustomerManager struct
type CustomerManager struct {  
        customers         map[string]Customer  
        customersByCategory map[string]map[string]bool  
        customersByRegion   map[string]map[string]bool  
}

// NewCustomerManager function 
func NewCustomerManager() *CustomerManager {  
        return &CustomerManager{  
                customers:         make(map[string]Customer),  
                customersByCategory: make(map[string]map[string]bool),  
                customersByRegion:   make(map[string]map[string]bool),  
        }  
}  

func (cm *CustomerManager) AddCustomer(customer Customer) {  
        cm.customers[customer.ID] = customer  

        // Update segmentation maps  
        for _, category := range customer.Preferences {  
                if cm.customersByCategory[category] == nil {  
                        cm.customersByCategory[category] = make(map[string]bool)  
                }  
                cm.customersByCategory[category][customer.ID] = true  
        }  
        if cm.customersByRegion[customer.Location] == nil {  
                cm.customersByRegion[customer.Location] = make(map[string]bool)  
        }  
        cm.customersByRegion[customer.Location][customer.ID] = true  
} 

 // FindCustomersByCategory: This method returns the customer IDs of all customers who frequently purchase products in a specific category.
func (cm *CustomerManager) FindCustomersByCategory(category string) []string {
        var customersInCategory []string
        if customers, ok := cm.customersByCategory[category]; ok {
                for customerID := range customers {
                        customersInCategory = append(customersInCategory, customerID)
                }
        }
        return customersInCategory
}

// FindCustomersByRegion: This method returns the customer IDs of all customers living in a specific region.
func (cm *CustomerManager) FindCustomersByRegion(region string) []string {
        var customersInRegion []string
        if customers, ok := cm.customersByRegion[region]; ok {
                for customerID := range customers {
                        customersInRegion = append(customersInRegion, customerID)
                }
        }
        return customersInRegion
}
        

func main() {  
        manager := NewCustomerManager()  

        customer1 := Customer{  
                ID: "customer1",  
                PurchaseHistory: PurchaseHistory{  
                        "product1": 2,  
                        "product2": 1,  
                },  
                Preferences: []string{"Electronics", "Books"},  
                Location:   "New York",  
        }  
        
        manager.AddCustomer(customer1)
  
        // Add more customers and call FindCustomersByCategory and FindCustomersByRegion functions to get the customers quickly based on category and region
        
         // ....

         customersInElectronics := manager.FindCustomersByCategory("Electronics")
         customersInNewYork := manager.FindCustomersByRegion("New York")
         fmt.Println("Customers in Electronics category:", customersInElectronics)
         fmt.Println("Customers in New York region:", customersInNewYork)
}  