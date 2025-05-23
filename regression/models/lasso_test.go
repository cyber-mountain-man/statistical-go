package models

import (
	"math"
	"testing"
)

func TestLassoRegression(t *testing.T) {
	X := [][]float64{
		{1, 2},
		{2, 1},
		{3, 3},
	}
	y := []float64{6, 5, 12}
	lambda := 0.5
	maxIter := 100

	beta, err := LassoRegression(X, y, lambda, maxIter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(beta) != 3 {
		t.Errorf("expected 3 coefficients (intercept + 2 features), got %d", len(beta))
	}

	for i, b := range beta {
		if math.IsNaN(b) || math.IsInf(b, 0) {
			t.Errorf("coefficient %d is invalid: %v", i, b)
		}
	}
}

func TestLassoRegression_InvalidInput(t *testing.T) {
	t.Run("Empty input", func(t *testing.T) {
		X := [][]float64{}
		y := []float64{}
		_, err := LassoRegression(X, y, 0.1, 100)
		if err == nil {
			t.Error("expected error for empty input, got nil")
		}
	})

	t.Run("Mismatched input", func(t *testing.T) {
		X := [][]float64{{1, 2}, {3, 4}}
		y := []float64{5}
		_, err := LassoRegression(X, y, 0.1, 100)
		if err == nil {
			t.Error("expected error for mismatched input lengths, got nil")
		}
	})
}

func TestLassoRegression_ZeroCoefficientCase(t *testing.T) {
	X := [][]float64{
		{1, 0},
		{0, 1},
		{1, 1},
	}
	y := []float64{1, 1, 2}
	lambda := 10.0
	maxIter := 1000

	beta, err := LassoRegression(X, y, lambda, maxIter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if beta[1] != 0 || beta[2] != 0 {
		t.Errorf("expected coefficients to shrink to zero, got %v", beta)
	}
}

func TestLassoRegression_ZeroBetaCase(t *testing.T) {
	X := [][]float64{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	y := []float64{2, 2, 2}
	lambda := 1e6
	maxIter := 1000

	beta, err := LassoRegression(X, y, lambda, maxIter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for i := 1; i < len(beta); i++ {
		if math.Abs(beta[i]) > 1e-6 {
			t.Errorf("expected beta[%d] ≈ 0 due to heavy regularization, got %.6f", i, beta[i])
		}
	}
}

func TestLassoRegression_NegativeRhoUpdate(t *testing.T) {
	X := [][]float64{
		{1, 10},
		{1, 20},
		{1, 30},
	}
	y := []float64{30, 20, 10} // decreasing output as feature 2 increases
	lambda := 1.0
	maxIter := 1000

	beta, err := LassoRegression(X, y, lambda, maxIter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	t.Logf("Coefficients: %v", beta)

	// We're testing for negative rho: beta[2] should be negative
	if beta[2] >= 0 {
		t.Errorf("expected beta[2] to be negative (rho < -lambda/2 branch), got %.6f", beta[2])
	}
}

