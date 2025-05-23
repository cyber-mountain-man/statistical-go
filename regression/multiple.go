package regression

import (
	"errors"
)

// MultipleLinearRegression performs linear regression with multiple features.
// X is a slice of rows, each containing feature values (without intercept column).
// y is the slice of target values.
// It returns the beta coefficients (including intercept).
func MultipleLinearRegression(X [][]float64, y []float64) ([]float64, error) {
	n := len(X)
	if n == 0 || len(y) != n {
		return nil, errors.New("invalid input dimensions")
	}

	// Add intercept term (column of 1s)
	XIntercept := make([][]float64, n)
	for i := 0; i < n; i++ {
		XIntercept[i] = append([]float64{1}, X[i]...) // prepend 1
	}

	XT := transpose(XIntercept)
	XTX := matMul(XT, XIntercept)
	XTXInv, err := matInverse(XTX)
	if err != nil {
		return nil, err
	}

	XTy := matVecMul(XT, y)
	beta := matVecMul(XTXInv, XTy)
	return beta, nil
}

// PredictMultiple predicts the target using beta coefficients and a feature vector x (no intercept).
func PredictMultiple(x []float64, beta []float64) float64 {
	y := beta[0] // intercept
	for i := 0; i < len(x); i++ {
		y += beta[i+1] * x[i]
	}
	return y
}

// --- Matrix Utilities (standard library only) ---

// transpose returns the transpose of matrix A
func transpose(A [][]float64) [][]float64 {
	rows := len(A)
	cols := len(A[0])
	result := make([][]float64, cols)
	for i := range result {
		result[i] = make([]float64, rows)
		for j := 0; j < rows; j++ {
			result[i][j] = A[j][i]
		}
	}
	return result
}

// matMul performs matrix multiplication A * B
func matMul(A, B [][]float64) [][]float64 {
	m := len(A)
	n := len(B[0])
	p := len(B)
	result := make([][]float64, m)
	for i := range result {
		result[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			sum := 0.0
			for k := 0; k < p; k++ {
				sum += A[i][k] * B[k][j]
			}
			result[i][j] = sum
		}
	}
	return result
}

// matVecMul performs matrix-vector multiplication A * x
func matVecMul(A [][]float64, x []float64) []float64 {
	m := len(A)
	n := len(x)
	result := make([]float64, m)
	for i := 0; i < m; i++ {
		sum := 0.0
		for j := 0; j < n; j++ {
			sum += A[i][j] * x[j]
		}
		result[i] = sum
	}
	return result
}

// matInverse computes the inverse of a square matrix A using Gauss-Jordan elimination.
func matInverse(A [][]float64) ([][]float64, error) {
	n := len(A)
	// Create augmented matrix [A | I]
	aug := make([][]float64, n)
	for i := range aug {
		aug[i] = make([]float64, 2*n)
		for j := 0; j < n; j++ {
			aug[i][j] = A[i][j]
		}
		aug[i][n+i] = 1
	}

	// Perform Gauss-Jordan elimination
	for i := 0; i < n; i++ {
		pivot := aug[i][i]
		if pivot == 0 {
			return nil, errors.New("matrix is singular")
		}
		for j := 0; j < 2*n; j++ {
			aug[i][j] /= pivot
		}
		for k := 0; k < n; k++ {
			if k != i {
				factor := aug[k][i]
				for j := 0; j < 2*n; j++ {
					aug[k][j] -= factor * aug[i][j]
				}
			}
		}
	}

	// Extract the inverse from augmented matrix
	inv := make([][]float64, n)
	for i := range inv {
		inv[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			inv[i][j] = aug[i][j+n]
		}
	}
	return inv, nil
}
