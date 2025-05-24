// examples/basic_stats.go
package main

import (
	"fmt"
	"github.com/cyber-mountain-man/statistical-go/stat"
)

func main() {
	data := []float64{2, 4, 4, 4, 5, 5, 7, 9}

	fmt.Println("Basic Descriptive Statistics")
	fmt.Printf("Data: %v\n", data)
	fmt.Printf("Mean: %.2f\n", stat.Mean(data))
	fmt.Printf("Median: %.2f\n", stat.Median(data))
	fmt.Printf("Mode: %.2f\n", stat.Mode(data))
	fmt.Printf("StdDev: %.2f\n", stat.StdDev(data))
	fmt.Printf("Variance: %.2f\n", stat.Variance(data))
	fmt.Printf("Range: %.2f\n", stat.Range(data))
	fmt.Printf("Min: %.2f\n", stat.Min(data))
	fmt.Printf("Max: %.2f\n", stat.Max(data))

	q1, q2, q3 := stat.Quartiles(data)
	fmt.Printf("Q1, Q2, Q3: %.2f, %.2f, %.2f\n", q1, q2, q3)

	fmt.Printf("Skewness: %.4f\n", stat.Skewness(data))
	fmt.Printf("Kurtosis: %.4f\n", stat.Kurtosis(data))
}
