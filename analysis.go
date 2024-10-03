package main

import (
    "fmt"
    "math"
)

// AdvancedAnalysis provides additional technical analysis
type AdvancedAnalysis struct {
    StandardDeviation float64
    UpperBollinger    float64
    LowerBollinger    float64
    RSI               float64
}

// CalculateAdvancedMetrics calculates additional technical indicators
func CalculateAdvancedMetrics(data []StockData, smaResults []MovingAverageResult, window int) (*AdvancedAnalysis, error) {
    if len(smaResults) == 0 {
        return nil, fmt.Errorf("no SMA results available")
    }
    
    analysis := &AdvancedAnalysis{}
    
    // Calculate standard deviation for Bollinger Bands
    latestSMA := smaResults[len(smaResults)-1]
    recentData := data[len(data)-window:]
    
    variance := 0.0
    for _, point := range recentData {
        variance += math.Pow(point.Close-latestSMA.SMA, 2)
    }
    analysis.StandardDeviation = math.Sqrt(variance / float64(window))
    
    // Calculate Bollinger Bands (2 standard deviations)
    analysis.UpperBollinger = latestSMA.SMA + (2 * analysis.StandardDeviation)
    analysis.LowerBollinger = latestSMA.SMA - (2 * analysis.StandardDeviation)
    
    // Simple RSI calculation (simplified)
    analysis.RSI = calculateSimpleRSI(data, 14)
    
    return analysis, nil
}

// calculateSimpleRSI calculates a simplified RSI indicator
func calculateSimpleRSI(data []StockData, period int) float64 {
    if len(data) < period+1 {
        return 50.0 // Default neutral value
    }
    
    gains := 0.0
    losses := 0.0
    
    for i := len(data) - period; i < len(data); i++ {
        change := data[i].Close - data[i-1].Close
        if change > 0 {
            gains += change
        } else {
            losses += math.Abs(change)
        }
    }
    
    avgGain := gains / float64(period)
    avgLoss := losses / float64(period)
    
    if avgLoss == 0 {
        return 100.0
    }
    
    rs := avgGain / avgLoss
    return 100 - (100 / (1 + rs))
}

// PrintAdvancedAnalysis displays advanced technical indicators
func PrintAdvancedAnalysis(analysis *AdvancedAnalysis, latestResult MovingAverageResult) {
    fmt.Println("\n=== ADVANCED TECHNICAL ANALYSIS ===")
    fmt.Printf("Standard Deviation: %.4f\n", analysis.StandardDeviation)
    fmt.Printf("Bollinger Upper Band: %.2f\n", analysis.UpperBollinger)
    fmt.Printf("Bollinger Lower Band: %.2f\n", analysis.LowerBollinger)
    fmt.Printf("RSI (14): %.2f\n", analysis.RSI)
    
    // RSI interpretation
    rsiSignal := "NEUTRAL"
    if analysis.RSI > 70 {
        rsiSignal = "OVERBOUGHT"
    } else if analysis.RSI < 30 {
        rsiSignal = "OVERSOLD"
    }
    fmt.Printf("RSI Signal: %s\n", rsiSignal)
    
    // Bollinger Band position
    bbPosition := "MIDDLE"
    if latestResult.Price >= analysis.UpperBollinger {
        bbPosition = "UPPER BAND"
    } else if latestResult.Price <= analysis.LowerBollinger {
        bbPosition = "LOWER BAND"
    }
    fmt.Printf("Bollinger Band Position: %s\n", bbPosition)
}