package main

import (
	"fmt"
	"sort"
)

type Sale struct {
	Region  string
	Revenue float64
	Date    string
}

type ByDateThenRegionThenRevenue []Sale

func (a ByDateThenRegionThenRevenue) Len() int           { return len(a) }
func (a ByDateThenRegionThenRevenue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDateThenRegionThenRevenue) Less(i, j int) bool {
	if a[i].Date != a[j].Date {
		return a[i].Date < a[j].Date
	}
	if a[i].Region != a[j].Region {
		return a[i].Region < a[j].Region
	}
	return a[i].Revenue > a[j].Revenue
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

	// Sort using Quick Sort
	sort.QuickSort(ByDateThenRegionThenRevenue(sales))

	fmt.Println("Sorted Sales Data:")
	for _, sale := range sales {
		fmt.Printf("Date: %s, Region: %s, Revenue: %.2f\n", sale.Date, sale.Region, sale.Revenue)
	}
}