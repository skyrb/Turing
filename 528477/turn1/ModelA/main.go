package main

import (
	"fmt"
)

// Define a struct for storing purchase history
type PurchaseHistory map[string]int // key: ProductID, value: Quantity

// Define a struct for storing customer information
type Customer struct {
	ID             string
	PurchaseHistory PurchaseHistory
	Preferences    []string
	Location       string
}

// CustomerManager is responsible for managing the customers
type CustomerManager struct {
	customers map[string]Customer // key: CustomerID, value: Customer
}

// NewCustomerManager creates a new CustomerManager instance
func NewCustomerManager() *CustomerManager {
	return &CustomerManager{
		customers: make(map[string]Customer),
	}
}

// AddCustomer adds a new customer to the manager
func (cm *CustomerManager) AddCustomer(customer Customer) {
	cm.customers[customer.ID] = customer
	fmt.Printf("Added customer: %s\n", customer.ID)
}

// UpdatePurchaseHistory updates a customer's purchase history
func (cm *CustomerManager) UpdatePurchaseHistory(customerID string, productID string, quantity int) {
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

// UpdatePreferences updates a customer's preferences
func (cm *CustomerManager) UpdatePreferences(customerID string, preferences []string) {
	customer, exists := cm.customers[customerID]
	if !exists {
		fmt.Printf("Customer with ID %s not found\n", customerID)
		return
	}

	customer.Preferences = preferences
	cm.customers[customerID] = customer
	fmt.Printf("Updated preferences for customer: %s\n", customerID)
}

// GetCustomer retrieves a customer by ID
func (cm *CustomerManager) GetCustomer(customerID string) (Customer, bool) {
	customer, exists := cm.customers[customerID]
	if !exists {
		fmt.Printf("Customer with ID %s not found\n", customerID)
	}
	return customer, exists
}

func main() {
	// Create a new CustomerManager
	manager := NewCustomerManager()

	// Add a new customer
	customer := Customer{
		ID:             "customer1",
		PurchaseHistory: make(PurchaseHistory),
		Preferences:    []string{"electronics", "books"},
		Location:       "New York",
	}
	manager.AddCustomer(customer)

	// Update purchase history
	manager.UpdatePurchaseHistory("customer1", "product1", 3)
	manager.UpdatePurchaseHistory("customer1", "product2", 1)

	// Update preferences
	manager.UpdatePreferences("customer1", []string{"books", "clothing"})

	// Retrieve and print customer information
	if customer, found := manager.GetCustomer("customer1"); found {
		fmt.Printf("Customer ID: %s\n", customer.ID)
		fmt.Printf("Location: %s\n", customer.Location)
		fmt.Printf("Preferences: %v\n", customer.Preferences)
		fmt.Printf("Purchase History: %v\n", customer.PurchaseHistory)
	}
}