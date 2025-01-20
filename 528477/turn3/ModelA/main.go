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

// CustomerManager now uses RWMutex for thread safety
type CustomerManager struct {
	mu            sync.RWMutex
	customers     map[string]Customer
	categoryIndex map[string]map[string]bool // map[Category]map[CustomerID]bool
	locationIndex map[string]map[string]bool // map[Location]map[CustomerID]bool
}

// NewCustomerManager initializes a new CustomerManager
func NewCustomerManager() *CustomerManager {
	return &CustomerManager{
		customers:     make(map[string]Customer),
		categoryIndex: make(map[string]map[string]bool),
		locationIndex: make(map[string]map[string]bool),
	}
}

// AddCustomer safely adds a customer
func (cm *CustomerManager) AddCustomer(customer Customer) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	cm.customers[customer.ID] = customer

	// Update indices
	for _, category := range customer.Preferences {
		if _, exists := cm.categoryIndex[category]; !exists {
			cm.categoryIndex[category] = make(map[string]bool)
		}
		cm.categoryIndex[category][customer.ID] = true
	}

	if _, exists := cm.locationIndex[customer.Location]; !exists {
		cm.locationIndex[customer.Location] = make(map[string]bool)
	}
	cm.locationIndex[customer.Location][customer.ID] = true

	fmt.Printf("Added customer: %s\n", customer.ID)
}

// UpdatePurchaseHistory safely updates a customer’s purchase history
func (cm *CustomerManager) UpdatePurchaseHistory(customerID string, productID string, quantity int) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

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
}

// UpdatePreferences safely updates a customer’s preferences
func (cm *CustomerManager) UpdatePreferences(customerID string, preferences []string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	customer, exists := cm.customers[customerID]
	if !exists {
		fmt.Printf("Customer with ID %s not found\n", customerID)
		return
	}

	// Remove old preferences from index
	for _, category := range customer.Preferences {
		delete(cm.categoryIndex[category], customerID)
	}

	// Add new preferences to index
	customer.Preferences = preferences
	for _, category := range customer.Preferences {
		if _, exists := cm.categoryIndex[category]; !exists {
			cm.categoryIndex[category] = make(map[string]bool)
		}
		cm.categoryIndex[category][customerID] = true
	}

	cm.customers[customerID] = customer
}

// GetCustomersByCategoryAndLocation retrieves customers by category and location
func (cm *CustomerManager) GetCustomersByCategoryAndLocation(category, location string) []Customer {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	customers := []Customer{}
	if categoryCustomers, exists := cm.categoryIndex[category]; exists {
		if locationCustomers, lExists := cm.locationIndex[location]; lExists {
			for id := range categoryCustomers {
				if _, ok := locationCustomers[id]; ok {
					customers = append(customers, cm.customers[id])
				}
			}
		}
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

	// Example query for customers who are interested in electronics and located in New York
	customers := manager.GetCustomersByCategoryAndLocation("electronics", "New York")
	fmt.Println("Electronics customers in New York:")
	for _, c := range customers {
		fmt.Println(c.ID)
	}
}