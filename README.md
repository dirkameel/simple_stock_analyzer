# Stock Data Analyzer - Go Implementation

A simple Go application that analyzes stock data and calculates Simple Moving Averages (SMA).

## Features

- Load stock data from CSV files
- Calculate Simple Moving Averages (5, 10, 20 days)
- Generate basic stock analysis reports
- Generate sample data for testing

## Requirements

- Go 1.16 or higher

## Usage

### 1. Using Sample Data

The easiest way to get started is to use the included sample data generator:

```bash
# Make the script executable (Unix/Linux/Mac)
chmod +x run_example.sh

# Run the example
./run_example.sh
```

### 2. Using Your Own Data

Prepare a CSV file with the following columns:
- Date (YYYY-MM-DD)
- Open
- High  
- Low
- Close
- Volume

Example CSV format:
```csv
Date,Open,High,Low,Close,Volume
2023-01-01,100.50,102.30,99.80,101.20,1500000
2023-01-02,101.30,103.10,100.50,102.80,1600000
```

Run the analyzer with your CSV file:
```bash
go run main.go your_stock_data.csv
```

### 3. Generate Sample Data

If you want to generate sample data separately:

```bash
go run data_generator.go main.go
```

## Output

The program will display:
- 5, 10, and 20-day Simple Moving Averages
- Last 10 days of SMA calculations for each period
- Basic stock analysis including price changes and high/low statistics

## Building

To build an executable:

```bash
go build -o stock-analyzer main.go data_generator.go
./stock-analyzer sample_stock_data.csv
```

## CSV Format Requirements

Your CSV file must have this exact header:
```
Date,Open,High,Low,Close,Volume
```

Dates should be in `YYYY-MM-DD` format, and all price values should be numeric.

## Example Output

```
Stock Data Analyzer - Simple Moving Average Calculator
======================================================
Loaded 90 days of stock data

=== 5-Day Simple Moving Average ===
2024-01-15: Close: $105.30, SMA5: $104.82
2024-01-16: Close: $106.10, SMA5: $105.44
...

=== Stock Analysis Report ===
Analysis Period: 2023-10-18 to 2024-01-15
Price Change: $5.30 (5.30%)
Latest Price: $105.30
52-Week High: $108.50
52-Week Low: $95.20
Current vs High: 2.95% below high
```

This Go implementation provides a fast, compiled alternative to Python for stock data analysis with similar functionality for calculating moving averages and basic technical analysis.