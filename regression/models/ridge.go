package models

import (
	"errors"
)

// RidgeRegression performs ridge regression with L2 regularization.
// It centers X and y before applying the closed-form ridge regression.
func RidgeRegression(X [][]float64, y []float64, lambda float64) ([]float64, error) {
	n := len(X)
	if n == 0 || len(y) != n {
		return nil, errors.New("invalid input dimensions")
	}

	m := len(X[0]) // number of features

	// Compute means of each feature column
	featureMeans := make([]float64, m)
	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			featureMeans[j] += X[i][j]
		}
		featureMeans[j] /= float64(n)
	}

	// Center X
	XCentered := make([][]float64, n)
	for i := 0; i < n; i++ {
		XCentered[i] = make([]float64, m)
		for j := 0; j < m; j++ {
			XCentered[i][j] = X[i][j] - featureMeans[j]
		}
	}

	// Compute mean of y
	yMean := 0.0
	for _, val := range y {
		yMean += val
	}
	yMean /= float64(n)

	// Center y
	yCentered := make([]float64, n)
	for i := 0; i < n; i++ {
		yCentered[i] = y[i] - yMean
	}

	// Add intercept term (column of 1s)
	XIntercept := make([][]float64, n)
	for i := 0; i < n; i++ {
		XIntercept[i] = append([]float64{1}, XCentered[i]...) // prepend 1
	}

	// X^T * X
	XT := transpose(XIntercept)
	XTX := matMul(XT, XIntercept)

	// Add lambda * I to X^T * X (except intercept)
	for i := 1; i < len(XTX); i++ {
		XTX[i][i] += lambda
	}

	// X^T * y
	XTy := matVecMul(XT, yCentered)

	// Inverse and multiply
	XTXInv, err := matInverse(XTX)
	if err != nil {
		return nil, err
	}

	betaCentered := matVecMul(XTXInv, XTy)

	// Adjust intercept to account for original data shift
	adjustedIntercept := yMean
	for j := 0; j < m; j++ {
		adjustedIntercept -= betaCentered[j+1] * featureMeans[j]
	}
	betaCentered[0] = adjustedIntercept

	return betaCentered, nil
}

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

func matVecMul(A [][]float64, x []float64) []float64 {
	m := len(A)
	result := make([]float64, m)
	for i := 0; i < m; i++ {
		for j := range x {
			result[i] += A[i][j] * x[j]
		}
	}
	return result
}

func matInverse(A [][]float64) ([][]float64, error) {
	n := len(A)
	aug := make([][]float64, n)
	for i := range aug {
		aug[i] = make([]float64, 2*n)
		for j := 0; j < n; j++ {
			aug[i][j] = A[i][j]
		}
		aug[i][n+i] = 1
	}

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

	inv := make([][]float64, n)
	for i := range inv {
		inv[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			inv[i][j] = aug[i][n+j]
		}
	}
	return inv, nil
}
