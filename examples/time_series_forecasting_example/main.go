package main

import (
	"fmt"
	"log"

	"github.com/cyber-mountain-man/statistical-go/regression"
	"github.com/cyber-mountain-man/statistical-go/stat"
)

// Simulated monthly revenue data (in thousands of dollars)
var months = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
var revenue = []float64{52.3, 55.1, 54.7, 57.8, 60.4, 61.2, 64.5, 65.1, 67.3, 69.0, 70.5, 72.8}

func main() {
	fmt.Println("=== Financial Time Series Forecasting ===")

	// 1. Moving Average Forecast
	windowSize := 3
	smaForecast := simpleMovingAverage(revenue, windowSize)
	fmt.Printf("Simple Moving Average (%d-month): %.2f\n", windowSize, smaForecast)

	// 2. Linear Trend Forecast using regression
	slope, intercept := regression.SimpleLinearRegression(months, revenue)
	nextMonth := float64(len(months) + 1)
	trendForecast := intercept + slope*nextMonth
	fmt.Printf("Linear Trend Forecast for Month %d: $%.2fk\n", int(nextMonth), trendForecast)

	// 3. Display underlying trend model
	fmt.Printf("Trend Line: Revenue ≈ %.2f + %.2f × Month\n", intercept, slope)

	// Optional: Check model quality (R²)
	rSquared := regression.RSquared(months, revenue, slope, intercept)
	fmt.Printf("Model R²: %.4f (1.0 = perfect fit)\n", rSquared)
}

// simpleMovingAverage returns the average of the last n values
func simpleMovingAverage(data []float64, n int) float64 {
	if n <= 0 || n > len(data) {
		log.Fatalf("invalid window size for moving average: %d", n)
	}
	return stat.Mean(data[len(data)-n:])
}
