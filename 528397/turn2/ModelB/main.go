package main  
import (  
        "fmt"
        "sync"
        "time"
)

// RegionData represents data for a region
type RegionData struct {  
        Region   string
        Priority int // Higher value means higher priority
}

// Function to simulate distribution of content to a regional CDN server
func distributeContent(regionData RegionData, wg *sync.WaitGroup, contentCh chan RegionData) {  
        defer wg.Done()
        fmt.Printf("Starting content distribution to %s region with priority %d\n", regionData.Region, regionData.Priority)
        distributeImagesAndBanners(regionData.Region)
        fmt.Printf("Completed content distribution to %s region\n", regionData.Region)
        // Let other regions know that content has been updated
        contentCh <- regionData
}

// Simulate the distribution of product images and banners
func distributeImagesAndBanners(region string) {  
        fmt.Printf("Distributing product images and banners to %s region...\n", region)
        // Simulate network delay
        time.Sleep(2 * time.Second)
        fmt.Printf("Product images and banners distributed to %s region\n", region)
}

func main() {  
        var wg sync.WaitGroup
        // Create buffered channels with a capacity to accommodate multiple regions
        highPriorityCh := make(chan RegionData, 10)
        mediumPriorityCh := make(chan RegionData, 10)
        lowPriorityCh := make(chan RegionData, 10)

        regions := []RegionData{
                {Region: "North America", Priority: 3},
                {Region: "Europe", Priority: 2},
                {Region: "Asia", Priority: 1},
                {Region: "China", Priority: 3},
        }

        // Spawn worker pools for different priority levels
        for i := 0; i < 3; i++ {
                go worker(highPriorityCh, &wg)
        }
        for i := 0; i < 2; i++ {
                go worker(mediumPriorityCh, &wg)
        }
        for i := 0; i < 1; i++ {
                go worker(lowPriorityCh, &wg)
        }

        // Add work to the appropriate channels based on priority
        for _, region := range regions {
                switch region.Priority {
                case 3:
                        highPriorityCh <- region
                case 2:
                        mediumPriorityCh <- region
                default:
                        lowPriorityCh <- region
                }
        }

        // Wait for all goroutines to complete
        wg.Wait()
        fmt.Println("All content distribution completed.")
}

// Worker function to receive region data from channels and initiate distribution
func worker(ch chan RegionData, wg *sync.WaitGroup) {
        for regionData := range ch {
                distributeContent(regionData, wg, make(chan RegionData))
        }
}
