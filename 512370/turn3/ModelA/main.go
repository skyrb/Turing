package main

import (
	"fmt"
)

type Sale struct {
	Region  string
	Revenue float64
	Date    string
}

func mergeSort(sales []Sale) []Sale {
	if len(sales) <= 1 {
		return sales
	}

	mid := len(sales) / 2
	leftHalf := mergeSort(sales[:mid])
	rightHalf := mergeSort(sales[mid:])

	return merge(leftHalf, rightHalf)
}

func merge(left, right []Sale) []Sale {
	var sortedSales []Sale
	var leftIndex, rightIndex int

	for leftIndex < len(left) && rightIndex < len(right) {
		if less(left[leftIndex], right[rightIndex]) {
			sortedSales = append(sortedSales, left[leftIndex])
			leftIndex++
		} else {
			sortedSales = append(sortedSales, right[rightIndex])
			rightIndex++
		}
	}

	// Append remaining elements from left or right
	sortedSales = append(sortedSales, left[leftIndex:]...)
	sortedSales = append(sortedSales, right[rightIndex:]...)

	return sortedSales
}

func less(a, b Sale) bool {
	if a.Date != b.Date {
		return a.Date < b.Date
	}
	if a.Region != b.Region {
		return a.Region < b.Region
	}
	return a.Revenue > b.Revenue
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

	sortedSales := mergeSort(sales)

	fmt.Println("Sorted Sales Data:")
	for _, sale := range sortedSales {
		fmt.Printf("Date: %s, Region: %s, Revenue: %.2f\n", sale.Date, sale.Region, sale.Revenue)
	}
}