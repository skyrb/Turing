package main  
import (  
    "fmt"
)  
type Sale struct {  
    Region  string
    Revenue float64
    Date    string
}  
func partition(sales []Sale, low int, high int) int {  
    pivot := sales[high].Revenue  
    i := low - 1  
    for j := low; j < high; j++ {  
        if sales[j].Revenue <= pivot {  
            i++  
            sales[i], sales[j] = sales[j], sales[i]  
        }  
    }  
    sales[i+1], sales[high] = sales[high], sales[i+1]  
    return i + 1  
}  
func quickSortRevenueDesc(sales []Sale, low int, high int) {  
    if low < high {  
        pi := partition(sales, low, high)  
        quickSortRevenueDesc(sales, low, pi-1)  
        quickSortRevenueDesc(sales, pi+1, high)  
    }  
}  
func sortSalesByDateRegionRevenueDesc(sales []Sale) {  
    quickSortRevenueDesc(sales, 0, len(sales)-1)
}  
func main() {  
    sales := []Sale{  
        {Region: "East", Revenue: 10000.0, Date: "2023-01-01"},  
        {Region: "West", Revenue: 12000.0, Date: "2023-01-01"},  
        {Region: "North", Revenue: 8000.0, Date: "2023-01-02"},  
        {Region: "South", Revenue: 15000.0, Date: "2023-01-01"},  
        {Region: "East", Revenue: 9000.0, Date: "2023-01-02"},  
        {Region: "West", Revenue: 11000.0, Date: "2023-01-02"},  
        {Region: "North", Revenue: 13000.0, Date: "2023-01-03"},  
        {Region: "South", Revenue: 7000.0, Date: "2023-01-03"},  
        {Region: "East", Revenue: 14000.0, Date: "2023-01-03"},  
        {Region: "West", Revenue: 16000.0, Date: "2023-01-03"}, 
		{Region: "East", Revenue: 10000.0, Date: "2023-01-01"},
    }  
  
    fmt.Println("Unsorted Sales Data:")  
    for _, sale := range sales {  
        fmt.Printf("Date: %s, Region: %s, Revenue: %.2f\n", sale.Date, sale.Region, sale.Revenue)  
    }  
  
    sortSalesByDateRegionRevenueDesc(sales)
  
    fmt.Println("\nSorted Sales Data using Quick Sort (Revenue Descending):")  
    for _, sale := range sales {  
        fmt.Printf("Date: %s, Region: %s, Revenue: %.2f\n", sale.Date, sale.Region, sale.Revenue)  
    }  
}  
