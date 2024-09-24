package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// readStockData reads and parses stock data from CSV file
func readStockData(filename string) ([]StockData, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("cannot open file %s: %v", filename, err)
	}
	defer file.Close()

	var stockData []StockData
	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		
		// Skip header line
		if lineNumber == 1 {
			continue
		}

		// Parse CSV line
		fields := strings.Split(line, ",")
		if len(fields) < 2 {
			continue // Skip invalid lines
		}

		// Parse date and close price
		date := strings.TrimSpace(fields[0])
		closePrice, err := strconv.ParseFloat(strings.TrimSpace(fields[1]), 64)
		if err != nil {
			continue // Skip lines with invalid price data
		}

		stockData = append(stockData, StockData{
			Date:  date,
			Close: closePrice,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	if len(stockData) == 0 {
		return nil, fmt.Errorf("no valid stock data found in file")
	}

	return stockData, nil
}