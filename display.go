package main

import "fmt"

// displayResults prints the moving average results in a formatted table
func displayResults(results []MovingAverageResult) {
	fmt.Printf("\n%-12s %-10s %-15s\n", "Date", "Close", "Moving Average")
	fmt.Println("-------------------------------------")

	for _, result := range results {
		if result.MovingAverage == 0 {
			fmt.Printf("%-12s %-10.2f %-15s\n", result.Date, result.Close, "N/A")
		} else {
			fmt.Printf("%-12s %-10.2f %-15.2f\n", result.Date, result.Close, result.MovingAverage)
		}
	}

	// Print summary statistics
	validResults := 0
	for _, result := range results {
		if result.MovingAverage != 0 {
			validResults++
		}
	}
	fmt.Printf("\nSummary: %d data points processed, %d moving averages calculated\n", 
		len(results), validResults)
}