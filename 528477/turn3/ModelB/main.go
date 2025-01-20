package main

import (
	"fmt"
	"sync"
)

type PurchaseHistory map[string]int // key: ProductID, value: Quantity

// Define a struct for storing customer information
type Customer struct {
	ID             string
	PurchaseHistory PurchaseHistory
	Preferences    []string
	Location       string
}
// CustomerManager is extended with segmentation functions and concurrent access control
type CustomerManager struct {
	customers    map[string]Customer // key: CustomerID, value: Customer
	categoryIndex map[string]map[string][]string // key: Category, value: map[Location][]CustomerID
	locationIndex map[string]map[string][]string // key: Location, value: map[Category][]CustomerID
	mutex        sync.RWMutex
}

// NewCustomerManager creates a new CustomerManager instance
func NewCustomerManager() *CustomerManager {
	return &CustomerManager{
		customers:    make(map[string]Customer),
		categoryIndex: make(map[string]map[string][]string),
		locationIndex: make(map[string]map[string][]string),
	}
}

// AddCustomer adds a new customer to the manager
func (cm *CustomerManager) AddCustomer(customer Customer) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	cm.customers[customer.ID] = customer

	// Update indices
	for _, category := range customer.Preferences {
		if cm.categoryIndex[category] == nil {
			cm.categoryIndex[category] = make(map[string][]string)
		}
		cm.categoryIndex[category][customer.Location] = append(cm.categoryIndex[category][customer.Location], customer.ID)
	}

	for _, location := range []string{customer.Location} {
		if cm.locationIndex[location] == nil {
			cm.locationIndex[location] = make(map[string][]string)
		}
		for _, category := range customer.Preferences {
			cm.locationIndex[location][category] = append(cm.locationIndex[location][category], customer.ID)
		}
	}

	fmt.Printf("Added customer: %s\n", customer.ID)
}

// UpdatePurchaseHistory updates a customer's purchase history in real-time
func (cm *CustomerManager) UpdatePurchaseHistory(customerID string, productID string, quantity int) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	customer, exists := cm.customers[customerID]
	if !exists {
		fmt.Printf("Customer with ID %s not found\n", customerID)
		return
	}

	if customer.PurchaseHistory == nil {
		customer.PurchaseHistory = make(PurchaseHistory)
	}

	customer.PurchaseHistory[productID] += quantity
	cm.customers[customerID] = customer
	fmt.Printf("Updated purchase history for customer: %s\n", customerID)
}

// UpdatePreferences updates a customer's preferences in real-time
func (cm *CustomerManager) UpdatePreferences(customerID string, preferences []string) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	customer, exists := cm.customers[customerID]
	if !exists {
		fmt.Printf("Customer with ID %s not found\n", customerID)
		return
	}

	// Remove old preferences from indices
	for _, category := range customer.Preferences {
		if cm.categoryIndex[category] != nil {
			delete(cm.categoryIndex[category], customer.Location)
		}
	}

	customer.Preferences = preferences

	// Add new preferences to indices
	for _, category := range customer.Preferences {
		if cm.categoryIndex[category] == nil {
			cm.categoryIndex[category] = make(map[string][]string)
		}
		cm.categoryIndex[category][customer.Location] = append(cm.categoryIndex[category][customer.Location], customer.ID)
	}

	cm.customers[customerID] = customer
	fmt.Printf("Updated preferences for customer: %s\n", customerID)
}

// GetCustomersByCategoryAndLocation retrieves customers by category and geographic location
func (cm *CustomerManager) GetCustomersByCategoryAndLocation(category, location string) []Customer {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	customerIDs := cm.categoryIndex[category][location]
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

	// Update customer preferences in real-time
	manager.UpdatePreferences("customer1", []string{"electronics", "clothing"})

	// Retrieve and print customers by category and location
	electronicsCustomersInNY := manager.GetCustomersByCategoryAndLocation("electronics", "New York")
	fmt.Println("Customers interested in electronics in New York:")
	for _, c := range electronicsCustomersInNY {
		fmt.Println(c.ID)
	}
}