package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

// StockData represents a single day's stock data
type StockData struct {
	Date   time.Time
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume int64
}

// MovingAverageResult represents SMA calculation result
type MovingAverageResult struct {
	Date  time.Time
	Close float64
	SMA   float64
}

func main() {
	fmt.Println("Stock Data Analyzer - Simple Moving Average Calculator")
	fmt.Println("======================================================")
	
	// Check if CSV file exists
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <csv_file>")
		fmt.Println("Example: go run main.go stock_data.csv")
		return
	}
	
	filename := os.Args[1]
	
	// Load stock data
	stockData, err := loadStockData(filename)
	if err != nil {
		fmt.Printf("Error loading stock data: %v\n", err)
		return
	}
	
	fmt.Printf("Loaded %d days of stock data\n", len(stockData))
	
	// Calculate SMA for different periods
	periods := []int{5, 10, 20}
	
	for _, period := range periods {
		fmt.Printf("\n=== %d-Day Simple Moving Average ===\n", period)
		smaResults := calculateSMA(stockData, period)
		
		// Display last 10 results
		startIdx := len(smaResults) - 10
		if startIdx < 0 {
			startIdx = 0
		}
		
		for i := startIdx; i < len(smaResults); i++ {
			result := smaResults[i]
			fmt.Printf("%s: Close: $%.2f, SMA%d: $%.2f\n", 
				result.Date.Format("2006-01-02"), 
				result.Close, 
				period, 
				result.SMA)
		}
	}
	
	// Generate analysis report
	generateAnalysisReport(stockData)
}

// loadStockData reads CSV file and parses stock data
func loadStockData(filename string) ([]StockData, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	
	var stockData []StockData
	
	// Skip header row and parse data
	for i := 1; i < len(records); i++ {
		record := records[i]
		
		// Parse date
		date, err := time.Parse("2006-01-02", record[0])
		if err != nil {
			continue // Skip invalid dates
		}
		
		// Parse numeric values
		open, _ := strconv.ParseFloat(record[1], 64)
		high, _ := strconv.ParseFloat(record[2], 64)
		low, _ := strconv.ParseFloat(record[3], 64)
		closePrice, _ := strconv.ParseFloat(record[4], 64)
		volume, _ := strconv.ParseInt(record[5], 10, 64)
		
		data := StockData{
			Date:   date,
			Open:   open,
			High:   high,
			Low:    low,
			Close:  closePrice,
			Volume: volume,
		}
		
		stockData = append(stockData, data)
	}
	
	return stockData, nil
}

// calculateSMA calculates Simple Moving Average for given period
func calculateSMA(data []StockData, period int) []MovingAverageResult {
	var results []MovingAverageResult
	
	for i := period - 1; i < len(data); i++ {
		var sum float64
		count := 0
		
		// Calculate sum of closing prices for the period
		for j := i - period + 1; j <= i; j++ {
			sum += data[j].Close
			count++
		}
		
		// Calculate average
		average := sum / float64(count)
		
		result := MovingAverageResult{
			Date:  data[i].Date,
			Close: data[i].Close,
			SMA:   average,
		}
		
		results = append(results, result)
	}
	
	return results
}

// generateAnalysisReport provides basic analysis of the stock data
func generateAnalysisReport(data []StockData) {
	if len(data) == 0 {
		return
	}
	
	fmt.Println("\n=== Stock Analysis Report ===")
	
	// Basic statistics
	latest := data[len(data)-1]
	oldest := data[0]
	
	priceChange := latest.Close - oldest.Close
	percentChange := (priceChange / oldest.Close) * 100
	
	fmt.Printf("Analysis Period: %s to %s\n", 
		oldest.Date.Format("2006-01-02"), 
		latest.Date.Format("2006-01-02"))
	fmt.Printf("Price Change: $%.2f (%.2f%%)\n", priceChange, percentChange)
	fmt.Printf("Latest Price: $%.2f\n", latest.Close)
	
	// Find highest and lowest prices
	highest := data[0].High
	lowest := data[0].Low
	
	for _, day := range data {
		if day.High > highest {
			highest = day.High
		}
		if day.Low < lowest {
			lowest = day.Low
		}
	}
	
	fmt.Printf("52-Week High: $%.2f\n", highest)
	fmt.Printf("52-Week Low: $%.2f\n", lowest)
	fmt.Printf("Current vs High: %.2f%% below high\n", 
		((highest - latest.Close) / highest) * 100)
}