package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// generateSampleData creates a sample CSV file with fake stock data
func generateSampleData() {
	fmt.Println("Generating sample stock data...")
	
	// Create sample data
	records := [][]string{
		{"Date", "Open", "High", "Low", "Close", "Volume"},
	}
	
	// Start from 90 days ago
	startDate := time.Now().AddDate(0, 0, -90)
	basePrice := 100.0
	
	for i := 0; i < 90; i++ {
		currentDate := startDate.AddDate(0, 0, i)
		
		// Generate realistic price movements
		change := (rand.Float64() - 0.5) * 4.0 // Random change between -2 and +2
		openPrice := basePrice
		closePrice := basePrice + change
		
		// Ensure high is higher than both open and close
		highPrice := max(openPrice, closePrice) + rand.Float64()*2
		// Ensure low is lower than both open and close
		lowPrice := min(openPrice, closePrice) - rand.Float64()*2
		
		// Generate volume (in thousands)
		volume := rand.Int63n(10000) + 1000
		
		record := []string{
			currentDate.Format("2006-01-02"),
			fmt.Sprintf("%.2f", openPrice),
			fmt.Sprintf("%.2f", highPrice),
			fmt.Sprintf("%.2f", lowPrice),
			fmt.Sprintf("%.2f", closePrice),
			fmt.Sprintf("%d", volume),
		}
		
		records = append(records, record)
		
		// Update base price for next day
		basePrice = closePrice
	}
	
	// Write to CSV file
	file, err := os.Create("sample_stock_data.csv")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()
	
	writer := csv.NewWriter(file)
	err = writer.WriteAll(records)
	if err != nil {
		fmt.Printf("Error writing CSV: %v\n", err)
		return
	}
	
	fmt.Println("Sample data generated: sample_stock_data.csv")
}

// Helper functions
func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}