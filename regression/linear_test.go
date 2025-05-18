package regression

import (
	"testing"
)

func TestSimpleLinearRegression(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{2, 4, 5, 4, 5}

	slope, intercept := SimpleLinearRegression(x, y)

	expectedSlope := 0.6
	expectedIntercept := 2.2

	if (slope < expectedSlope-0.01 || slope > expectedSlope+0.01) ||
		(intercept < expectedIntercept-0.01 || intercept > expectedIntercept+0.01) {
		t.Errorf("Unexpected regression result: got slope=%.2f, intercept=%.2f", slope, intercept)
	}
}

func TestSimpleLinearRegression_EdgeCases(t *testing.T) {
	t.Run("Unequal length input", func(t *testing.T) {
		x := []float64{1, 2, 3}
		y := []float64{4, 5}
		slope, intercept := SimpleLinearRegression(x, y)
		if slope != 0 || intercept != 0 {
			t.Errorf("Expected zero return for unequal length, got slope=%.2f, intercept=%.2f", slope, intercept)
		}
	})

	t.Run("Empty input", func(t *testing.T) {
		x := []float64{}
		y := []float64{}
		slope, intercept := SimpleLinearRegression(x, y)
		if slope != 0 || intercept != 0 {
			t.Errorf("Expected zero return for empty input, got slope=%.2f, intercept=%.2f", slope, intercept)
		}
	})

	t.Run("Zero denominator case", func(t *testing.T) {
		x := []float64{1, 1, 1}
		y := []float64{2, 2, 2}
		slope, intercept := SimpleLinearRegression(x, y)
		if slope != 0 || intercept != 0 {
			t.Errorf("Expected zero return for zero denominator, got slope=%.2f, intercept=%.2f", slope, intercept)
		}
	})
}

func TestPredictY(t *testing.T) {
	x := 4.0
	slope := 2.0
	intercept := 1.0
	expected := 9.0

	result := PredictY(x, slope, intercept)
	if result != expected {
		t.Errorf("Expected PredictY to return %.2f, got %.2f", expected, result)
	}
}

func TestIsZeroDenominator(t *testing.T) {
	t.Run("Identical x-values", func(t *testing.T) {
		x := []float64{3, 3, 3, 3, 3}
		if !isZeroDenominator(x) {
			t.Errorf("Expected true for identical x-values, got false")
		}
	})

	t.Run("Varying x-values", func(t *testing.T) {
		x := []float64{1, 2, 3, 4, 5}
		if isZeroDenominator(x) {
			t.Errorf("Expected false for varying x-values, got true")
		}
	})
}

func TestCalculateSlope_DivisionByZero(t *testing.T) {
	x := []float64{1, 1, 1}
	y := []float64{2, 2, 2}

	slope := CalculateSlope(x, y)
	if slope != 0 {
		t.Errorf("Expected slope to be 0 when denominator is zero, got %.2f", slope)
	}
}

func TestCalculateIntercept_UnequalLength(t *testing.T) {
	x := []float64{1, 2, 3}
	y := []float64{4, 5}
	slope := 1.0

	intercept := CalculateIntercept(x, y, slope)
	if intercept != 0 {
		t.Errorf("Expected intercept to be 0 for unequal length input, got %.2f", intercept)
	}
}

func TestCalculateIntercept_EmptyInput(t *testing.T) {
	x := []float64{}
	y := []float64{}
	slope := 1.0

	intercept := CalculateIntercept(x, y, slope)
	if intercept != 0 {
		t.Errorf("Expected intercept to be 0 for empty input, got %.2f", intercept)
	}
}

func TestCalculateSlope_EmptyInput(t *testing.T) {
	x := []float64{}
	y := []float64{}

	slope := CalculateSlope(x, y)
	if slope != 0 {
		t.Errorf("Expected slope to be 0 for empty input, got %.2f", slope)
	}
}
