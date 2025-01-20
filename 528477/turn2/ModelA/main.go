package main

import (
	"fmt"
)

type PurchaseHistory map[string]int // key: ProductID, value: Quantity

// Define a struct for storing customer information
type Customer struct {
	ID             string
	PurchaseHistory PurchaseHistory
	Preferences    []string
	Location       string
}


// CustomerManager is extended with segmentation functions
type CustomerManager struct {
	customers    map[string]Customer // key: CustomerID, value: Customer
	categoryIndex map[string][]string // key: Category, value: []CustomerID
	locationIndex map[string][]string // key: Location, value: []CustomerID
}

// NewCustomerManager creates a new CustomerManager instance
func NewCustomerManager() *CustomerManager {
	return &CustomerManager{
		customers:    make(map[string]Customer),
		categoryIndex: make(map[string][]string),
		locationIndex: make(map[string][]string),
	}
}

// AddCustomer adds a new customer to the manager
func (cm *CustomerManager) AddCustomer(customer Customer) {
	cm.customers[customer.ID] = customer

	// Update indices
	for _, category := range customer.Preferences {
		cm.categoryIndex[category] = append(cm.categoryIndex[category], customer.ID)
	}
	cm.locationIndex[customer.Location] = append(cm.locationIndex[customer.Location], customer.ID)

	fmt.Printf("Added customer: %s\n", customer.ID)
}

// GetCustomersByCategory retrieves customers by purchase category preference
func (cm *CustomerManager) GetCustomersByCategory(category string) []Customer {
	customerIDs := cm.categoryIndex[category]
	customers := make([]Customer, 0, len(customerIDs))
	for _, id := range customerIDs {
		customers = append(customers, cm.customers[id])
	}
	return customers
}

// GetCustomersByLocation retrieves customers by geographic location
func (cm *CustomerManager) GetCustomersByLocation(location string) []Customer {
	customerIDs := cm.locationIndex[location]
	customers := make([]Customer, 0, len(customerIDs))
	for _, id := range customerIDs {
		customers = append(customers, cm.customers[id])
	}
	return customers
}

func main() {
	// Create a new CustomerManager
	manager := NewCustomerManager()

	// Add customers
	manager.AddCustomer(Customer{
		ID:             "customer1",
		PurchaseHistory: make(PurchaseHistory),
		Preferences:    []string{"electronics", "books"},
		Location:       "New York",
	})

	manager.AddCustomer(Customer{
		ID:             "customer2",
		PurchaseHistory: make(PurchaseHistory),
		Preferences:    []string{"clothing", "books"},
		Location:       "San Francisco",
	})

	// Retrieve and print customers by category
	electronicsCustomers := manager.GetCustomersByCategory("electronics")
	fmt.Println("Customers interested in electronics:")
	for _, c := range electronicsCustomers {
		fmt.Println(c.ID)
	}

	// Retrieve and print customers by location
	nyCustomers := manager.GetCustomersByLocation("New York")
	fmt.Println("Customers in New York:")
	for _, c := range nyCustomers {
		fmt.Println(c.ID)
	}
}