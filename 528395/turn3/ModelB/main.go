package main  
import (  
    "context"
    "fmt"
    "log"
    "math/rand"
    "sync"
    "time"
)

const (
    // Maximum number of retries for each API call
    maxRetries = 3
    // Initial delay for exponential backoff
    backoffDelay = 500 * time.Millisecond
    // Maximum number of goroutines in the pool
    maxGoroutines = 50
)

// Define a struct for the data source
type DataSource struct {
    Name       string
    FetchFunc   func(context.Context) (interface{}, error)
    MaxRetries int
    BackoffDelay time.Duration
}

// Define a struct for the data
type StockData struct {
    Source   DataSource
    Data     interface{}
    Timestamp time.Time
    Err      error
}

// Health status struct
type DataSourceStatus struct {
    Name        string
    IsHealthy   bool
    LastError   error
    NumRetries  int
    LastFetch   time.Time
    LastSuccess time.Time
}

// Goroutine function with timeout, retries, and health checks
func fetchData(ctx context.Context, dataCh chan<- StockData, healthCh chan<- DataSourceStatus, source DataSource) {
    for {
        select {
        case <-ctx.Done():
            return
        default:
            // Perform API call with timeout
            ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
            defer cancel()

            data, err := fetchWithRetry(ctxWithTimeout, source)

            // Update health status
            updateHealthStatus(healthCh, source, err, data)

            if err != nil {
                log.Printf("Error fetching data from %s: %v\n", source.Name, err)
                continue
            }

            dataCh <- StockData{
                Source:   source,
                Data:     data,
                Timestamp: time.Now(),
            }
        }
    }
}

// Function to perform API call with retries
func fetchWithRetry(ctx context.Context, source DataSource) (interface{}, error) {
    for i := 0; i <= source.MaxRetries; i++ {
        data, err := source.FetchFunc(ctx)
        if err == nil {
            return data, nil
        }

        // Exponential backoff on retry
        delay := source.BackoffDelay * time.Duration(rand.Intn(100)+1)
        time.Sleep(delay)
    }

    return nil, fmt.Errorf("max retries exceeded for %s", source.Name)
}

// Function to update health status
func updateHealthStatus(healthCh chan<- DataSourceStatus, source DataSource, err error, data interface{}) {
    status := DataSourceStatus{
        Name:        source.Name,
        IsHealthy:   err == nil,
        LastError:   err,
        NumRetries:  0,
        LastFetch:   time.Now(),
        LastSuccess: time.Now(),
    }

    if err != nil {
        status.IsHealthy = false
        status.NumRetries++
    } else {
        status.IsHealthy = true
    }

    healthCh <- status
}

// Goroutine pool management
type GoroutinePool struct {  
    wg sync.WaitGroup
    dataCh chan StockData
    healthCh chan DataSourceStatus
    sources []DataSource
    numActive int
    maxActive int
}

func NewGoroutinePool(dataCh chan StockData, healthCh chan DataSourceStatus, sources []DataSource, maxActive int) *GoroutinePool {  
    return &GoroutinePool{  
        dataCh: dataCh,
        healthCh: healthCh,
        sources: sources,
        maxActive: maxActive,
    }
}

func (p *GoroutinePool) Start() {  
    for i := 0; i < p.maxActive; i++ {  
        p.startGoroutine()
    }

    // Goroutine to monitor health checks and restart unhealthy goroutines
    go p.healthCheckMonitor()
}

func (p *GoroutinePool) startGoroutine() {  
    p.wg.Add(1)
    go func() {  
        defer p.wg.Done()
        for {  
            select {  
            case <-p.ctx.Done():  
                return
            default:  
                if p.numActive < p.maxActive {  
                    source := p.sources[rand.Intn(len(p.sources))]
                    fetchData(p.ctx, p.dataCh, p.healthCh, source)
                    p.numActive--
                }
            }
        }
    }()
}

func (p *GoroutinePool) healthCheckMonitor() {  
    for status := range p.healthCh {  
        if !status.IsHealthy {  
            log.Printf("Data source %s is unhealthy. Error: %v\n", status.Name, status.LastError)
            // You can add logic to restart the goroutine based on your requirements
        }
    }
}

func main() {  
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    
    dataCh := make(chan StockData)  
    healthCh := make(chan DataSourceStatus)  
    
    // Define your data sources
    sources := []DataSource{
        {Name: "Source1", FetchFunc: fetchRealTimePrices},
        {Name: "Source2", FetchFunc: fetchMarketNews},
        // Add more data sources here
    }
    
    pool := NewGoroutinePool(dataCh, healthCh, sources, maxGoroutines)  
    pool.Start()  
    
    // Main processing loop  
    go func() {  
        for data := range dataCh {  
            if data.Err != nil {  
                log.Printf("Error received: %v\n", data.Err)  
                continue  
            }  
            fmt.Printf("Received data from %s at %v: %v\n", data.Source.Name, data.Timestamp, data.Data)  
        }  
    }()  
    
    // Run for a specific duration  
    time.Sleep(10 * time.Second)  
    cancel()  
    
    pool.wg.Wait()  
    close(dataCh)  
    close(healthCh)  
} 