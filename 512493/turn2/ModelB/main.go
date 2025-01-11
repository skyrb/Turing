package main

import (
	"fmt"
	"sync"
)

type StockStore struct {
	stocks map[string]float64
	mu     sync.RWMutex
}

// NewStockStore creates a new instance of StockStore
func NewStockStore() *StockStore {
	return &StockStore{stocks: make(map[string]float64)}
}

// AddStock adds a new stock to the store with the given price
func (s *StockStore) AddStock(ticker string, price float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.stocks[ticker] = price
	fmt.Printf("Added stock: %s at price %.2f\n", ticker, price)
}

// UpdateStock updates the price of an existing stock
func (s *StockStore) UpdateStock(ticker string, price float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.stocks[ticker]; exists {
		s.stocks[ticker] = price
		fmt.Printf("Updated stock: %s at new price %.2f\n", ticker, price)
	} else {
		fmt.Printf("Stock %s does not exist\n", ticker)
	}
}

// GetStock retrieves the price of a stock
func (s *StockStore) GetStock(ticker string) (float64, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if price, exists := s.stocks[ticker]; exists {
		return price, nil
	} else {
		return 0, fmt.Errorf("stock %s not found", ticker)
	}
}

func main() {
	store := NewStockStore()

	// Simulate concurrent updates
	go func() {
		store.AddStock("AAPL", 145.25)
		store.UpdateStock("AAPL", 150.00)
		store.UpdateStock("MSFT", 280.00)
	}()

	go func() {
		store.AddStock("GOOGL", 2750.75)
		store.UpdateStock("GOOGL", 2800.00)
		store.UpdateStock("AAPL", 152.50)
	}()

	// Retrieve stock prices concurrently
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		aaplPrice, err := store.GetStock("AAPL")
		if err == nil {
			fmt.Printf("The price of AAPL is: %.2f\n", aaplPrice)
		} else {
			fmt.Println(err)
		}
	}()

	go func() {
		defer wg.Done()
		googlPrice, err := store.GetStock("GOOGL")
		if err == nil {
			fmt.Printf("The price of GOOGL is: %.2f\n", googlPrice)
		} else {
			fmt.Println(err)
		}
	}()

	go func() {
		defer wg.Done()
		msftPrice, err := store.GetStock("MSFT")
		if err == nil {
			fmt.Printf("The price of MSFT is: %.2f\n", msftPrice)
		} else {
			fmt.Println(err)
		}
	}()
	wg.Wait()
}
  