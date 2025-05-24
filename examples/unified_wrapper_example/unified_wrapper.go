package main

import (
	"fmt"
	"github.com/cyber-mountain-man/statistical-go"
)

func main() {
	fmt.Println("=== Unified API Example ===")

	// Data sample
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Descriptive statistics
	fmt.Printf("Mean: %.2f\n", statistical.Mean(data))
	fmt.Printf("Median: %.2f\n", statistical.Median(data))
	fmt.Printf("Standard Deviation: %.2f\n", statistical.StdDev(data))
	fmt.Printf("Skewness: %.2f\n", statistical.Skewness(data))
	fmt.Printf("Kurtosis: %.2f\n", statistical.Kurtosis(data))

	// Probability example
	fmt.Printf("Binomial P(X=3) [n=5, p=0.5]: %.4f\n", statistical.BinomialPMF(3, 5, 0.5))
	fmt.Printf("Normal CDF(x=1.96, μ=0, σ=1): %.4f\n", statistical.NormalCDF(1.96, 0, 1))
	fmt.Printf("Normal Inverse CDF(p=0.975, μ=0, σ=1): %.4f\n", statistical.NormalInverseCDF(0.975, 0, 1))

	// Monte Carlo simulation
	fmt.Printf("Estimated Pi: %.6f\n", statistical.EstimatePi(1_000_000))
}
