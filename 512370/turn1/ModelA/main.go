package main

import (
	"fmt"
	"sort"
	"time"
)

// SalesRecord represents a sales record with Region, Revenue, and Date.
type SalesRecord struct {
	Region   string
	Revenue  float64
	Date     time.Time
}

// SalesRecords is a slice of SalesRecord.
type SalesRecords []SalesRecord

// SortByDateRegionRevenueDesc implements the Sort interface for SalesRecords
// by sorting first by Date, then by Region, and finally by Revenue in descending order.
func (s SalesRecords) SortByDateRegionRevenueDesc() {
	sort.Slice(s, func(i, j int) bool {
		if s[i].Date != s[j].Date {
			return s[i].Date.Before(s[j].Date)
		} else if s[i].Region != s[j].Region {
			return s[i].Region < s[j].Region
		} else {
			return s[i].Revenue > s[j].Revenue
		}
	})
}

func main() {
	// Sample sales data
	salesData := []SalesRecord{
		{Region: "Asia", Revenue: 1234.56, Date: time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC)},
		{Region: "Europe", Revenue: 9876.54, Date: time.Date(2023, 1, 10, 0, 0, 0, 0, time.UTC)},
		{Region: "North America", Revenue: 5678.90, Date: time.Date(2023, 1, 12, 0, 0, 0, 0, time.UTC)},
		{Region: "Asia", Revenue: 3456.78, Date: time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC)},
		{Region: "Europe", Revenue: 1234.56, Date: time.Date(2023, 1, 10, 0, 0, 0, 0, time.UTC)},
		{Region: "North America", Revenue: 9876.54, Date: time.Date(2023, 1, 12, 0, 0, 0, 0, time.UTC)},
		{Region: "Asia", Revenue: 5678.90, Date: time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC)},
		{Region: "Europe", Revenue: 3456.78, Date: time.Date(2023, 1, 10, 0, 0, 0, 0, time.UTC)},
		{Region: "North America", Revenue: 1234.56, Date: time.Date(2023, 1, 12, 0, 0, 0, 0, time.UTC)},
		{Region: "Asia", Revenue: 9876.54, Date: time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC)},
		{Region: "Europe", Revenue: 5678.90, Date: time.Date(2023, 1, 10, 0, 0, 0, 0, time.UTC)},
		{Region: "North America", Revenue: 3456.78, Date: time.Date(2023, 1, 12, 0, 0, 0, 0, time.UTC)},
	}

	// Sort the sales data
	salesData.SortByDateRegionRevenueDesc()

	// Print the sorted sales data
	fmt.Println("Sorted Sales Data:")
	for _, record := range salesData {
		fmt.Printf("Region: %s, Revenue: %.2f, Date: %s\n", record.Region, record.Revenue, record.Date.Format("2006-01-02"))
	}
}