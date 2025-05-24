package main

import (
	"fmt"
	"log"

	"github.com/cyber-mountain-man/statistical-go"
)

func main() {
	fmt.Println("=== Ridge & Lasso Regression Example ===")

	// Example dataset (X = features, y = labels)
	X := [][]float64{
		{1, 2},
		{2, 3},
		{4, 5},
		{3, 6},
		{5, 7},
	}
	y := []float64{3, 5, 9, 8, 11}

	// Ridge Regression
	ridgeCoeffs, err := statistical.RidgeRegression(X, y, 0.5)
	if err != nil {
		log.Fatalf("Ridge regression error: %v", err)
	}
	fmt.Printf("Ridge Coefficients (λ=0.5): %.4f\n", ridgeCoeffs)

	// Lasso Regression
	lassoCoeffs, err := statistical.LassoRegression(X, y, 0.1, 100)
	if err != nil {
		log.Fatalf("Lasso regression error: %v", err)
	}
	fmt.Printf("Lasso Coefficients (λ=0.1): %.4f\n", lassoCoeffs)
}
