package main

import (
	"fmt"
	"github.com/cyber-mountain-man/statistical-go"
)

func main() {
	fmt.Println("=== Simple Linear Regression Example ===")

	x := []float64{1, 2, 3, 4, 5}
	y := []float64{2, 4, 5, 4, 5}

	slope, intercept := statistical.SimpleLinearRegression(x, y)
	fmt.Printf("Slope: %.4f, Intercept: %.4f\n", slope, intercept)

	rSquared := statistical.RSquared(x, y, slope, intercept)
	fmt.Printf("R²: %.4f\n", rSquared)

	yPred := make([]float64, len(x))
	for i, val := range x {
		yPred[i] = statistical.PredictY(val, slope, intercept)
	}

	mse := statistical.MSE(y, yPred)
	rmse := statistical.RMSE(y, yPred)
	fmt.Printf("MSE: %.4f, RMSE: %.4f\n", mse, rmse)
}
