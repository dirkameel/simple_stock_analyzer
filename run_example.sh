#!/bin/bash

echo "Stock Analyzer - Go Implementation"
echo "=================================="

# Generate sample data if it doesn't exist
if [ ! -f "sample_stock_data.csv" ]; then
    echo "Generating sample stock data..."
    go run data_generator.go main.go
fi

# Run the analyzer
echo ""
echo "Running stock analyzer..."
go run main.go sample_stock_data.csv