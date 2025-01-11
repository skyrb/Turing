package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	batchSize      = 1000
	updateInterval = 100 * time.Millisecond
)

type StockUpdate struct {
	Ticker string
	Price  float64
}

type StockStore struct {
	stocks  map[string]float64
	mutex   sync.RWMutex
	updates chan StockUpdate
}

// NewStockStore creates a new instance of StockStore
func NewStockStore() *StockStore {
	s := &StockStore{
		stocks:  make(map[string]float64),
		updates: make(chan StockUpdate, batchSize),
	}
	go s.processUpdates()
	return s
}

// AddStock adds a new stock to the store with the given price
func (s *StockStore) AddStock(ticker string, price float64) {
	s.updates <- StockUpdate{ticker, price}
}

// UpdateStock updates the price of an existing stock
func (s *StockStore) UpdateStock(ticker string, price float64) {
	s.updates <- StockUpdate{ticker, price}
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

func (s *StockStore) processUpdates() {
	var batch []StockUpdate
	ticker := time.NewTicker(updateInterval)
	defer ticker.Stop()

	for {
		select {
		case update := <-s.updates:
			batch = append(batch, update)
			if len(batch) >= batchSize {
				s.applyBatch(batch)
				batch = nil
			}
		case <-ticker.C:
			if len(batch) > 0 {
				s.applyBatch(batch)
				batch = nil
			}
		}
	}
}

func (s *StockStore) applyBatch(batch []StockUpdate) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for _, update := range batch {
		s.stocks[update.Ticker] = update.Price
	}
}

func main() {
	store := NewStockStore()

	// Generate high-frequency updates
	go func() {
		for {
			ticker := fmt.Sprintf("TICKER_%d", time.Now().UnixNano())
			price := float64(time.Now().UnixNano()) / 1e6
			store.AddStock(ticker, price)
			time.Sleep(1 * time.Microsecond)
		}
	}()

	// Retrieve stock prices periodically
	go func() {
		for {
			ticker := fmt.Sprintf("TICKER_%d", time.Now().UnixNano())
			price, err := store.GetStock(ticker)
			if err == nil {
				fmt.Printf("The price of %s is: %.2f\n", ticker, price)
			} else {
				fmt.Println(err)
			}
			time.Sleep(10 * time.Millisecond)
		}
	}()

	// Wait for updates to process
	time.Sleep(10 * time.Second)
}