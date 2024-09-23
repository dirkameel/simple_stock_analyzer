package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// StockData represents a single stock data point
type StockData struct {
	Date  string
	Close float64
}

// MovingAverageResult represents the result of moving average calculation
type MovingAverageResult struct {
	Date          string
	Close         float64
	MovingAverage float64
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <filename> <window_size>")
		fmt.Println("Example: go run main.go stock_data.csv 5")
		return
	}

	filename := os.Args[1]
	windowSize, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("Invalid window size. Please provide a valid integer.")
	}

	// Read and parse stock data
	stockData, err := readStockData(filename)
	if err != nil {
		log.Fatal("Error reading stock data:", err)
	}

	// Calculate moving averages
	results, err := calculateMovingAverages(stockData, windowSize)
	if err != nil {
		log.Fatal("Error calculating moving averages:", err)
	}

	// Display results
	displayResults(results)
}