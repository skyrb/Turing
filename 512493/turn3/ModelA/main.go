package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

type StockStore struct {
	stocks  map[string]float64
	mutex   sync.RWMutex
	batch   map[string]float64
	batchCh chan struct{}
}

// NewStockStore creates a new instance of StockStore with a batching mechanism
func NewStockStore() *StockStore {
	s := &StockStore{
		stocks:  make(map[string]float64),
		batch:   make(map[string]float64),
		batchCh: make(chan struct{}, 1),
	}
	go s.batchProcessor()
	return s
}

// AddStock adds a new stock to the batch with the given price
func (s *StockStore) AddStock(ticker string, price float64) {
	s.batch[ticker] = price
	select {
	case s.batchCh <- struct{}{}:
	default:
	}
	fmt.Printf("Added stock to batch: %s at price %.2f\n", ticker, price)
}

// UpdateStock updates the price of an existing stock in the batch
func (s *StockStore) UpdateStock(ticker string, price float64) {
	s.batch[ticker] = price
	select {
	case s.batchCh <- struct{}{}:
	default:
	}
	fmt.Printf("Updated stock in batch: %s at new price %.2f\n", ticker, price)
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

func (s *StockStore) batchProcessor() {
	tick := time.NewTicker(100 * time.Millisecond) // Process batch every 100ms
	defer tick.Stop()

	for range tick.C {
		s.mutex.Lock()
		defer s.mutex.Unlock()

		for ticker, price := range s.batch {
			s.stocks[ticker] = price
			delete(s.batch, ticker)
		}
	}
}

func main() {
	store := NewStockStore()

	// Simulate high-frequency updates
	go simulateUpdates(store)

	// Wait for some time to allow updates to accumulate
	time.Sleep(5 * time.Second)

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

func simulateUpdates(store *StockStore) {
	tick := time.NewTicker(20 * time.Millisecond) // Simulate updates every 20ms
	defer tick.Stop()

	stocks := []string{"AAPL", "GOOGL", "AMZN", "TSLA", "FB"}
	for _, ticker := range stocks {
		store.AddStock(ticker, 100.0)
	}

	for range tick.C {
		for _, ticker := range stocks {
			newPrice := float64(rand.Intn(100)) + 100.0
			store.UpdateStock(ticker, newPrice)
		}
	}
}