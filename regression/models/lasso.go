package models

import (
	"errors"
)

// LassoRegression performs Lasso regression using coordinate descent.
func LassoRegression(X [][]float64, y []float64, lambda float64, maxIter int) ([]float64, error) {
	n := len(X)
	if n == 0 || len(y) != n {
		return nil, errors.New("invalid input dimensions")
	}

	p := len(X[0])

	// Center X and y
	XCentered, xMeans := centerColumns(X)
	yCentered, yMean := centerVector(y)

	// Initialize coefficients to 0
	beta := make([]float64, p)

	for iter := 0; iter < maxIter; iter++ {
		for j := 0; j < p; j++ {
			var rho float64
			for i := 0; i < n; i++ {
				pred := 0.0
				for k := 0; k < p; k++ {
					if k != j {
						pred += XCentered[i][k] * beta[k]
					}
				}
				rho += XCentered[i][j] * (yCentered[i] - pred)
			}

			norm := 0.0
			for i := 0; i < n; i++ {
				norm += XCentered[i][j] * XCentered[i][j]
			}

			if rho < -lambda/2 {
				beta[j] = (rho + lambda/2) / norm
			} else if rho > lambda/2 {
				beta[j] = (rho - lambda/2) / norm
			} else {
				beta[j] = 0
			}
		}
	}

	// Add intercept
	intercept := yMean
	for j := 0; j < p; j++ {
		intercept -= beta[j] * xMeans[j]
	}

	return append([]float64{intercept}, beta...), nil
}

// --- Centering utilities ---

func centerVector(v []float64) ([]float64, float64) {
	n := len(v)
	sum := 0.0
	for _, val := range v {
		sum += val
	}
	mean := sum / float64(n)

	centered := make([]float64, n)
	for i := range v {
		centered[i] = v[i] - mean
	}
	return centered, mean
}

func centerColumns(X [][]float64) ([][]float64, []float64) {
	n := len(X)
	p := len(X[0])
	means := make([]float64, p)
	for j := 0; j < p; j++ {
		sum := 0.0
		for i := 0; i < n; i++ {
			sum += X[i][j]
		}
		means[j] = sum / float64(n)
	}

	XCentered := make([][]float64, n)
	for i := 0; i < n; i++ {
		XCentered[i] = make([]float64, p)
		for j := 0; j < p; j++ {
			XCentered[i][j] = X[i][j] - means[j]
		}
	}
	return XCentered, means
}
