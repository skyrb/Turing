package main  
import "fmt"

type StockStore struct {  
    store map[string]float64
}

func NewStockStore() *StockStore {  
    return &StockStore{store: make(map[string]float64)}
}

// Add a new stock ticker or update its price
func (s *StockStore) AddOrUpdate(ticker string, price float64) {  
    s.store[ticker] = price
}

// Retrieve the price of a stock ticker
func (s *StockStore) GetPrice(ticker string) (float64, bool) {  
    price, ok := s.store[ticker]
    return price, ok
}

// Delete a stock ticker from the store
func (s *StockStore) Delete(ticker string) {  
    delete(s.store, ticker)
}

func main() {  
    store := NewStockStore()

    // Add some stock prices
    store.AddOrUpdate("GOOG", 1000.0)
    store.AddOrUpdate("AAPL", 250.0)
    store.AddOrUpdate("MSFT", 180.0)

    // Retrieve and print prices
    googPrice, found := store.GetPrice("GOOG")
    if found {  
        fmt.Println("GOOG price:", googPrice)
    } else {  
        fmt.Println("GOOG not found.")
    }

    aaplPrice, found := store.GetPrice("AAPL")  
    if found {  
        fmt.Println("AAPL price:", aaplPrice)
    } else {  
        fmt.Println("AAPL not found.")
    }

    // Update the price of MSFT
    store.AddOrUpdate("MSFT", 190.0)

    // Delete MSFT
    store.Delete("MSFT")

    // Try to retrieve MSFT again
    msftPrice, found := store.GetPrice("MSFT")  
    if found {  
        fmt.Println("MSFT price:", msftPrice)
    } else {  
        fmt.Println("MSFT not found.")
    }
}  