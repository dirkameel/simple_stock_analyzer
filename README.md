## Stock Data Analyzer - Go Implementation

This is a Go implementation of a stock data analyzer that calculates simple moving averages from CSV stock data.

### Prerequisites

- Go 1.16 or higher installed on your system

### Installation

1. Clone or download the project files
2. Navigate to the project directory
3. The `go.mod` file is already included for module support

### Usage

1. **Prepare your stock data CSV file**:
   - Format: `Date,Close` (header required)
   - One data point per line
   - See `stock_data.csv` for example format

2. **Run the analyzer**:
   ```bash
   go run main.go data_reader.go moving_average.go display.go stock_data.csv 5
   ```

   Where:
   - `stock_data.csv` is your input file
   - `5` is the moving average window size

### Example Output

```
Date         Close      Moving Average
-------------------------------------
2024-01-01   100.50     N/A
2024-01-02   102.25     N/A
2024-01-03   101.75     N/A
2024-01-04   103.00     N/A
2024-01-05   104.50     102.40
2024-01-06   103.25     102.95
2024-01-07   105.00     103.50
2024-01-08   106.25     104.60
2024-01-09   107.50     105.50
2024-01-10   108.00     106.00

Summary: 10 data points processed, 6 moving averages calculated
```

### Features

- **CSV Data Reading**: Reads stock data from CSV files with date and closing price
- **Simple Moving Average**: Calculates SMA for specified window size
- **Error Handling**: Robust error handling for file operations and data parsing
- **Formatted Output**: Clean, tabular display of results
- **Input Validation**: Validates window size and data sufficiency

### File Structure

- `main.go` - Entry point and command-line interface
- `data_reader.go` - CSV file reading and parsing
- `moving_average.go` - Moving average calculation logic
- `display.go` - Results formatting and display
- `go.mod` - Go module definition

### Building

To create an executable binary:

```bash
go build -o stockanalyzer
./stockanalyzer stock_data.csv 5
```

### Error Handling

The program handles various error conditions:
- Missing or invalid command-line arguments
- File not found or unreadable
- Invalid CSV format
- Insufficient data for moving average calculation
- Invalid window size

This implementation provides a robust, efficient solution for stock data analysis with moving averages in Go.