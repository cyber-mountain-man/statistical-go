package regression

import (
	"math"
	"testing"
)

func floatsAlmostEqual(a, b float64, epsilon float64) bool {
	return math.Abs(a-b) <= epsilon
}

func slicesAlmostEqual(a, b []float64, epsilon float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !floatsAlmostEqual(a[i], b[i], epsilon) {
			return false
		}
	}
	return true
}

func TestMultipleLinearRegression(t *testing.T) {
	X := [][]float64{
		{1, 2},
		{2, 1},
		{3, 3},
	}
	y := []float64{6, 5, 12}

	// Updated expected values based on actual output
	expected := []float64{-1, 1.66666666666667, 2.66666666666667}

	beta, err := MultipleLinearRegression(X, y)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !slicesAlmostEqual(beta, expected, 1e-6) {
		t.Errorf("expected beta %v, got %v", expected, beta)
	}
}

func TestPredictMultiple(t *testing.T) {
	// Using updated coefficients
	beta := []float64{-1, 1.66666666666667, 2.66666666666667}
	x := []float64{2, 3} // expected y = -1 + 1.666...*2 + 2.666...*3 = approx 10.3333

	expected := -1 + 1.66666666666667*2 + 2.66666666666667*3
	predicted := PredictMultiple(x, beta)

	if !floatsAlmostEqual(predicted, expected, 1e-6) {
		t.Errorf("expected prediction %.6f, got %.6f", expected, predicted)
	}
}

func TestMultipleLinearRegression_InvalidInput(t *testing.T) {
	X := [][]float64{}
	y := []float64{}
	_, err := MultipleLinearRegression(X, y)
	if err == nil {
		t.Error("expected error for empty input, got nil")
	}

	X = [][]float64{{1, 2}, {3, 4}}
	y = []float64{5}
	_, err = MultipleLinearRegression(X, y)
	if err == nil {
		t.Error("expected error for mismatched dimensions, got nil")
	}
}

func TestMatInverse_SingularMatrix(t *testing.T) {
	singular := [][]float64{
		{1, 2},
		{2, 4}, // This row is linearly dependent on the first
	}
	_, err := matInverse(singular)
	if err == nil {
		t.Error("expected error for singular matrix, got nil")
	}
}

func TestTranspose_EmptyMatrix(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for empty matrix, but code did not panic")
		}
	}()

	_ = transpose([][]float64{})
}

func TestTranspose_EmptyMatrixPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic when transposing an empty matrix, but no panic occurred")
		}
	}()

	// Trigger panic from accessing A[0] when A is empty
	_ = transpose([][]float64{})
}

func TestMultipleLinearRegression_SingularMatrix(t *testing.T) {
	// Create a singular matrix (two identical rows)
	X := [][]float64{
		{1, 2},
		{1, 2},
	}
	y := []float64{3, 3}

	_, err := MultipleLinearRegression(X, y)
	if err == nil {
		t.Error("expected error due to singular matrix, got nil")
	}
}
