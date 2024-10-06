# Stock Data Analysis Tool in Go

A simple Go application for analyzing stock data and calculating technical indicators, particularly Simple Moving Averages (SMA).

## Features

- **Simple Moving Average (SMA)** calculation with customizable window sizes
- **Basic trading signals** (BUY/SELL/HOLD) based on price vs SMA position
- **Advanced technical analysis** including:
  - Bollinger Bands
  - Relative Strength Index (RSI)
  - Standard deviation
- **CSV data parsing** for stock price data
- **Sample data generator** for testing

## Requirements

- Go 1.21 or later

## Installation

1. Clone or download the source files
2. Ensure all files are in the same directory
3. Run `go mod tidy` to initialize the module

## Usage

### Basic Usage

```bash
# Analyze a CSV file with default 20-day SMA
go run main.go stock_data.csv

# Analyze with custom window size (e.g., 50-day SMA)
go run main.go stock_data.csv 50
```

### CSV Format

The CSV file should have the following columns:
```
Date,Open,High,Low,Close,Volume
2024-01-01,100.0,102.5,99.5,101.2,1000000
2024-01-02,101.5,103.0,100.5,102.8,1200000
...
```

### Generating Sample Data

To test the application, you can use the included sample data generator:

```go
// Add this to main function temporarily to generate sample data
func main() {
    // ... existing code ...
    
    // Uncomment to generate sample data
    // if err := GenerateSampleData(); err != nil {
    //     log.Fatalf("Error generating sample data: %v", err)
    // }
    // fmt.Println("Sample data generated: sample_stock_data.csv")
}
```

## Output

The application provides:
- Daily SMA calculations with trading signals
- Summary statistics (total signals, buy/sell/hold counts)
- Latest trading signal
- Advanced technical indicators (Bollinger Bands, RSI)

## File Structure

- `main.go` - Main application entry point and core logic
- `go.mod` - Go module definition
- `sample_data.go` - Sample data generator for testing
- `analysis.go` - Advanced technical analysis functions

## Trading Signals

- **BUY**: When closing price crosses above SMA
- **SELL**: When closing price crosses below SMA  
- **HOLD**: When price is close to SMA or no clear trend

## Example Output

```
Stock Data Analysis - Simple Moving Average Calculator
======================================================
Analyzing file: sample_stock_data.csv with window size: 20

Simple Moving Average (SMA-20) Results:
Date                    Price           SMA             Signal
----                    -----           ---             ------
2024-01-20      121.20          113.55          BUY
2024-01-21      122.50          114.40          BUY
...

=== ANALYSIS SUMMARY ===
Total data points with SMA: 6
Buy signals: 6
Sell signals: 0
Hold signals: 0

Latest signal (2024-01-25): BUY
Price: 126.80, SMA: 116.13
```

This Go implementation provides a robust, type-safe alternative to Python for stock data analysis with better performance and compile-time error checking.