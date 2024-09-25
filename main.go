package main

import (
    "encoding/csv"
    "fmt"
    "io"
    "log"
    "os"
    "strconv"
    "time"
)

// StockData represents a single stock data point
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
    Price float64
    SMA   float64
}

func main() {
    fmt.Println("Stock Data Analysis - Simple Moving Average Calculator")
    fmt.Println("======================================================")
    
    // Check if filename is provided
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <csv_file> [window_size]")
        fmt.Println("Example: go run main.go stock_data.csv 20")
        return
    }
    
    filename := os.Args[1]
    windowSize := 20 // default window size
    
    // Parse window size if provided
    if len(os.Args) >= 3 {
        if ws, err := strconv.Atoi(os.Args[2]); err == nil && ws > 0 {
            windowSize = ws
        } else {
            fmt.Printf("Invalid window size. Using default: %d\n", windowSize)
        }
    }
    
    fmt.Printf("Analyzing file: %s with window size: %d\n\n", filename, windowSize)
    
    // Read and parse CSV data
    stockData, err := readCSV(filename)
    if err != nil {
        log.Fatalf("Error reading CSV file: %v", err)
    }
    
    // Calculate SMA
    smaResults, err := calculateSMA(stockData, windowSize)
    if err != nil {
        log.Fatalf("Error calculating SMA: %v", err)
    }
    
    // Display results
    displayResults(smaResults, windowSize)
    
    // Generate summary report
    generateSummary(smaResults)
}

// readCSV reads and parses stock data from CSV file
func readCSV(filename string) ([]StockData, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    
    reader := csv.NewReader(file)
    
    // Skip header
    _, err = reader.Read()
    if err != nil {
        return nil, err
    }
    
    var stockData []StockData
    
    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            return nil, err
        }
        
        // Parse date (assuming format: YYYY-MM-DD)
        date, err := time.Parse("2006-01-02", record[0])
        if err != nil {
            return nil, fmt.Errorf("error parsing date %s: %v", record[0], err)
        }
        
        // Parse numeric values
        open, _ := strconv.ParseFloat(record[1], 64)
        high, _ := strconv.ParseFloat(record[2], 64)
        low, _ := strconv.ParseFloat(record[3], 64)
        closePrice, _ := strconv.ParseFloat(record[4], 64)
        volume, _ := strconv.ParseInt(record[5], 10, 64)
        
        stockData = append(stockData, StockData{
            Date:   date,
            Open:   open,
            High:   high,
            Low:    low,
            Close:  closePrice,
            Volume: volume,
        })
    }
    
    // Reverse data to have oldest first (CSV usually has newest first)
    for i, j := 0, len(stockData)-1; i < j; i, j = i+1, j-1 {
        stockData[i], stockData[j] = stockData[j], stockData[i]
    }
    
    return stockData, nil
}

// calculateSMA calculates Simple Moving Average for closing prices
func calculateSMA(data []StockData, window int) ([]MovingAverageResult, error) {
    if len(data) < window {
        return nil, fmt.Errorf("not enough data points for window size %d", window)
    }
    
    var results []MovingAverageResult
    
    for i := window - 1; i < len(data); i++ {
        sum := 0.0
        for j := i - window + 1; j <= i; j++ {
            sum += data[j].Close
        }
        
        sma := sum / float64(window)
        
        results = append(results, MovingAverageResult{
            Date:  data[i].Date,
            Price: data[i].Close,
            SMA:   sma,
        })
    }
    
    return results, nil
}

// displayResults shows the SMA calculation results
func displayResults(results []MovingAverageResult, window int) {
    fmt.Printf("Simple Moving Average (SMA-%d) Results:\n", window)
    fmt.Println("Date\t\t\tPrice\t\tSMA\t\tSignal")
    fmt.Println("----\t\t\t-----\t\t---\t\t------")
    
    for _, result := range results {
        signal := "HOLD"
        if result.Price > result.SMA {
            signal = "BUY"
        } else if result.Price < result.SMA {
            signal = "SELL"
        }
        
        fmt.Printf("%s\t%.2f\t\t%.2f\t\t%s\n", 
            result.Date.Format("2006-01-02"), 
            result.Price, 
            result.SMA, 
            signal)
    }
}

// generateSummary provides a summary of the analysis
func generateSummary(results []MovingAverageResult) {
    if len(results) == 0 {
        return
    }
    
    fmt.Println("\n=== ANALYSIS SUMMARY ===")
    fmt.Printf("Total data points with SMA: %d\n", len(results))
    
    buySignals := 0
    sellSignals := 0
    latestResult := results[len(results)-1]
    
    for _, result := range results {
        if result.Price > result.SMA {
            buySignals++
        } else if result.Price < result.SMA {
            sellSignals++
        }
    }
    
    fmt.Printf("Buy signals: %d\n", buySignals)
    fmt.Printf("Sell signals: %d\n", sellSignals)
    fmt.Printf("Hold signals: %d\n", len(results)-buySignals-sellSignals)
    
    currentSignal := "HOLD"
    if latestResult.Price > latestResult.SMA {
        currentSignal = "BUY"
    } else if latestResult.Price < latestResult.SMA {
        currentSignal = "SELL"
    }
    
    fmt.Printf("\nLatest signal (%s): %s\n", 
        latestResult.Date.Format("2006-01-02"), 
        currentSignal)
    fmt.Printf("Price: %.2f, SMA: %.2f\n", latestResult.Price, latestResult.SMA)
}