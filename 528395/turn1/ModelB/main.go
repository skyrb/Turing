package main  
import (  
    "fmt"
    "sync"
    "time"
	"math/rand"
)

// Data structures to represent different APIs
type StockPrice struct {
    Symbol string
    Price  float64
}

type MarketNews struct {
    Headline string
    Content  string
}

type TradingVolume struct {
    Symbol string
    Volume int
}

// Functions to fetch data from different APIs (simulated with random delays)
func fetchStockPrice() (StockPrice, error) {
    // Simulate a random delay
    time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
    return StockPrice{"AAPL", 165.78}, nil
}

func fetchMarketNews() (MarketNews, error) {
    // Simulate a random delay
    time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
    return MarketNews{"Apple Unveils New iPhone 14 Pro", "The latest iPhone 14 Pro features a powerful A16 Bionic chip and advanced camera system."}, nil
}

func fetchTradingVolume() (TradingVolume, error) {
    // Simulate a random delay
    time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
    return TradingVolume{"GOOGL", 12345678}, nil
}

// Channel type to handle data from each API
type APIData chan interface{}

func ingestStockPrices(data APIData, wg *sync.WaitGroup) {
    defer wg.Done()
    for {
        price, err := fetchStockPrice()
        if err != nil {
            // Handle error
            fmt.Println("Error fetching stock price:", err)
            continue
        }
        data <- price
    }
}

func ingestMarketNews(data APIData, wg *sync.WaitGroup) {
    defer wg.Done()
    for {
        news, err := fetchMarketNews()
        if err != nil {
            // Handle error
            fmt.Println("Error fetching market news:", err)
            continue
        }
        data <- news
    }
}

func ingestTradingVolumes(data APIData, wg *sync.WaitGroup) {
    defer wg.Done()
    for {
        volume, err := fetchTradingVolume()
        if err != nil {
            // Handle error
            fmt.Println("Error fetching trading volume:", err)
            continue
        }
        data <- volume
    }
}

func processData(data APIData) {
    for {
        select {
        case price := <-data:
            switch p := price.(type) {
            case StockPrice:
                fmt.Printf("Stock Price: %s - %.2f\n", p.Symbol, p.Price)
            case MarketNews:
                fmt.Printf("Market News: %s\n", p.Headline)
            case TradingVolume:
                fmt.Printf("Trading Volume: %s - %d\n", p.Symbol, p.Volume)
            }
        }
    }
}

func main() {
    data := make(APIData)
    var wg sync.WaitGroup

    // Start goroutines to ingest data from APIs
    wg.Add(1)
    go ingestStockPrices(data, &wg)
    wg.Add(1)
    go ingestMarketNews(data, &wg)
    wg.Add(1)
    go ingestTradingVolumes(data, &wg)

    // Start goroutine to process data
    go processData(data)

    // Wait for all goroutines to finish
    wg.Wait()
}