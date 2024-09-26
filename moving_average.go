package main

import "fmt"

// calculateMovingAverages calculates simple moving averages for stock data
func calculateMovingAverages(stockData []StockData, windowSize int) ([]MovingAverageResult, error) {
	if len(stockData) < windowSize {
		return nil, fmt.Errorf("not enough data points for window size %d", windowSize)
	}

	if windowSize <= 0 {
		return nil, fmt.Errorf("window size must be positive")
	}

	var results []MovingAverageResult

	for i := 0; i < len(stockData); i++ {
		// For the first (windowSize-1) elements, we can't calculate moving average
		if i < windowSize-1 {
			results = append(results, MovingAverageResult{
				Date:          stockData[i].Date,
				Close:         stockData[i].Close,
				MovingAverage: 0, // No moving average available
			})
			continue
		}

		// Calculate moving average for current window
		sum := 0.0
		for j := i - windowSize + 1; j <= i; j++ {
			sum += stockData[j].Close
		}
		movingAverage := sum / float64(windowSize)

		results = append(results, MovingAverageResult{
			Date:          stockData[i].Date,
			Close:         stockData[i].Close,
			MovingAverage: movingAverage,
		})
	}

	return results, nil
}