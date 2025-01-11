package main

import (
	"fmt"
	"sync"
)

type StockStore struct {
	stocks  map[string]float64
	mutex   sync.RWMutex
}

// NewStockStore creates a new instance of StockStore
func NewStockStore() *StockStore {
	return &StockStore{stocks: make(map[string]float64)}
}

// AddStock adds a new stock to the store with the given price
func (s *StockStore) AddStock(ticker string, price float64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.stocks[ticker] = price
	fmt.Printf("Added stock: %s at price %.2f\n", ticker, price)
}

// UpdateStock updates the price of an existing stock
func (s *StockStore) UpdateStock(ticker string, price float64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if _, exists := s.stocks[ticker]; exists {
		s.stocks[ticker] = price
		fmt.Printf("Updated stock: %s at new price %.2f\n", ticker, price)
	} else {
		fmt.Printf("Stock %s does not exist\n", ticker)
	}
}

// GetStock retrieves the price of a stock
func (s *StockStore) GetStock(ticker string) (float64, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if price, exists := s.stocks[ticker]; exists {
		return price, nil
	} else {
		return 0, fmt.Errorf("stock %s not found", ticker)
	}
}

func main() {
	store := NewStockStore()

	// Add stocks
	store.AddStock("AAPL", 145.25)
	store.AddStock("GOOGL", 2750.75)

	// Update stock price
	store.UpdateStock("AAPL", 150.00)

	// Retrieve stock prices
	aaplPrice, err := store.GetStock("AAPL")
	if err == nil {
		fmt.Printf("The price of AAPL is: %.2f\n", aaplPrice)
	} else {
		fmt.Println(err)
	}

	googlPrice, err := store.GetStock("GOOGL")
	if err == nil {
		fmt.Printf("The price of GOOGL is: %.2f\n", googlPrice)
	} else {
		fmt.Println(err)
	}

	// Attempt to get a non-existent stock
	_, err = store.GetStock("MSFT")
	if err != nil {
		fmt.Println(err)
	}
}