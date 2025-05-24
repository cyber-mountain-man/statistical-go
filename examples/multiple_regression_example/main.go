package main

import (
	"fmt"

	"github.com/cyber-mountain-man/statistical-go/regression"
)

func main() {
	fmt.Println("=== Multiple Linear Regression Example ===")

	// Predictor variables:
	// Columns: [YearsExperience, EducationLevel, HoursPerWeek]
	X := [][]float64{
		{2, 1, 40},
		{3, 2, 42},
		{5, 1, 45},
		{6, 3, 50},
		{8, 2, 52},
		{10, 3, 55},
	}

	// Target variable (Salary in $1000s)
	y := []float64{35, 45, 50, 65, 75, 90}

	// Fit the model
	coefficients, err := regression.MultipleLinearRegression(X, y)
	if err != nil {
		fmt.Println("Error fitting regression:", err)
		return
	}

	fmt.Println("Coefficients:")
	fmt.Printf("Intercept: %.2f\n", coefficients[0])
	for i, coef := range coefficients[1:] {
		fmt.Printf("X%d: %.2f\n", i+1, coef)
	}

	// Predict salary for a new individual
	newX := []float64{1, 5, 2, 48} // intercept + inputs
	predictedSalary := 0.0
	for i := range newX {
		predictedSalary += newX[i] * coefficients[i]
	}
	fmt.Printf("Predicted Salary: $%.2fk\n", predictedSalary)
}
