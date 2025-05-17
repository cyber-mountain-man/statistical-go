package main

import (
	"fmt"

	"github.com/cyber-mountain-man/statistical-go"
)

func main() {
	data := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 3.0, 2.0}

	fmt.Println("=== Descriptive Statistics ===")
	fmt.Println("Mean:", statistical.Mean(data))
	fmt.Println("Median:", statistical.Median(data))
	fmt.Println("Mode:", statistical.Mode(data))
	fmt.Println("Range:", statistical.Range(data))
	fmt.Println("Min:", statistical.Min(data))
	fmt.Println("Max:", statistical.Max(data))
	q1, q2, q3 := statistical.Quartiles(data)
	fmt.Printf("Quartiles: Q1 = %.2f, Q2 = %.2f, Q3 = %.2f\n", q1, q2, q3)
	fmt.Println("Variance:", statistical.Variance(data))
	fmt.Println("Standard Deviation:", statistical.StdDev(data))

	fmt.Println("\n=== Probability ===")
	fmt.Println("P(A|B):", statistical.ConditionalProbability(0.15, 0.3))

	fmt.Println("\n=== Monte Carlo Simulations ===")
	fmt.Println("Estimate Pi:", statistical.EstimatePi(100000))
	fmt.Println("Estimate Pi Parallel:", statistical.EstimatePiParallel(100000, 4))
}
