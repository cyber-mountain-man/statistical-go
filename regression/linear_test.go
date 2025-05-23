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

func TestCalculateSlope(t *testing.T) {
	t.Run("Zero denominator", func(t *testing.T) {
		x := []float64{1, 1, 1}
		y := []float64{2, 2, 2}
		if slope := CalculateSlope(x, y); slope != 0 {
			t.Errorf("Expected slope=0 when denominator is zero, got %.2f", slope)
		}
	})

	t.Run("Empty input", func(t *testing.T) {
		x := []float64{}
		y := []float64{}
		if slope := CalculateSlope(x, y); slope != 0 {
			t.Errorf("Expected slope=0 for empty input, got %.2f", slope)
		}
	})
}

func TestCalculateIntercept(t *testing.T) {
	t.Run("Unequal input length", func(t *testing.T) {
		x := []float64{1, 2, 3}
		y := []float64{4, 5}
		slope := 1.0
		if intercept := CalculateIntercept(x, y, slope); intercept != 0 {
			t.Errorf("Expected intercept=0 for unequal input, got %.2f", intercept)
		}
	})

	t.Run("Empty input", func(t *testing.T) {
		x := []float64{}
		y := []float64{}
		slope := 1.0
		if intercept := CalculateIntercept(x, y, slope); intercept != 0 {
			t.Errorf("Expected intercept=0 for empty input, got %.2f", intercept)
		}
	})
}

func TestMeanSquaredError(t *testing.T) {
	t.Run("Valid input", func(t *testing.T) {
		yTrue := []float64{3, -0.5, 2, 7}
		yPred := []float64{2.5, 0.0, 2, 8}
		expected := 0.375
		mse := MeanSquaredError(yTrue, yPred)
		if mse < expected-0.0001 || mse > expected+0.0001 {
			t.Errorf("Expected MSE ≈ %.3f, got %.3f", expected, mse)
		}
	})

	t.Run("Mismatched length", func(t *testing.T) {
		yTrue := []float64{1, 2, 3}
		yPred := []float64{1, 2}
		mse := MeanSquaredError(yTrue, yPred)
		if mse != 0 {
			t.Errorf("Expected MSE = 0 for mismatched lengths, got %.3f", mse)
		}
	})

	t.Run("Empty slices", func(t *testing.T) {
		yTrue := []float64{}
		yPred := []float64{}
		mse := MeanSquaredError(yTrue, yPred)
		if mse != 0 {
			t.Errorf("Expected MSE = 0 for empty input, got %.3f", mse)
		}
	})
}

func TestEvaluateLinearRegression(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{2, 4, 5, 4, 5}

	mse := EvaluateLinearRegression(x, y)
	expected := 0.48

	if mse < expected-0.01 || mse > expected+0.01 {
		t.Errorf("Expected MSE ≈ %.2f, got %.2f", expected, mse)
	}
}

func TestRSquared(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{2, 4, 5, 4, 5}

	slope, intercept := SimpleLinearRegression(x, y)
	r2 := RSquared(x, y, slope, intercept)

	expected := 0.60
	if r2 < expected-0.01 || r2 > expected+0.01 {
		t.Errorf("Expected R² ≈ %.2f, got %.2f", expected, r2)
	}
}

func TestRSquared_EdgeCases(t *testing.T) {
	t.Run("Mismatched input lengths", func(t *testing.T) {
		x := []float64{1, 2, 3}
		y := []float64{1, 2}
		if r2 := RSquared(x, y, 1, 0); r2 != 0 {
			t.Errorf("Expected R² = 0 for mismatched lengths, got %.2f", r2)
		}
	})

	t.Run("Empty input", func(t *testing.T) {
		x := []float64{}
		y := []float64{}
		if r2 := RSquared(x, y, 1, 0); r2 != 0 {
			t.Errorf("Expected R² = 0 for empty input, got %.2f", r2)
		}
	})

	t.Run("Zero total variance", func(t *testing.T) {
		x := []float64{1, 2, 3}
		y := []float64{5, 5, 5}
		if r2 := RSquared(x, y, 0, 5); r2 != 0 {
			t.Errorf("Expected R² = 0 for zero variance, got %.2f", r2)
		}
	})
}

func TestSimpleLinearRegression_ZeroSlopeAndZeroDenominator(t *testing.T) {
	x := []float64{1, 1, 1}
	y := []float64{2, 2, 2}

	slope, intercept := SimpleLinearRegression(x, y)

	if slope != 0 || intercept != 0 {
		t.Errorf("Expected slope and intercept to be 0 for zero denominator case, got slope=%.2f, intercept=%.2f", slope, intercept)
	}
}

func TestSimpleLinearRegression_ZeroSlopeAndZeroDenominatorPath(t *testing.T) {
	// Identical x-values => zero denominator
	x := []float64{2, 2, 2}
	y := []float64{3, 3, 3} // slope would be 0, and denominator is 0

	slope, intercept := SimpleLinearRegression(x, y)

	if slope != 0 || intercept != 0 {
		t.Errorf("Expected slope and intercept to be 0 for zero slope and denominator, got slope=%.2f, intercept=%.2f", slope, intercept)
	}
}

func TestSimpleLinearRegression_ZeroSlope_And_ZeroDenominator_Branch(t *testing.T) {
	x := []float64{5, 5, 5}
	y := []float64{7, 7, 7}

	slope, intercept := SimpleLinearRegression(x, y)

	if slope != 0 || intercept != 0 {
		t.Errorf("Expected slope and intercept = 0 due to zero slope and denominator, got slope=%.2f, intercept=%.2f", slope, intercept)
	}
}

func TestSimpleLinearRegression_SlopeZeroAndDenominatorZero(t *testing.T) {
	x := []float64{5, 5, 5} // identical x-values → zero denominator
	y := []float64{7, 7, 7} // constant y-values → zero slope

	slope, intercept := SimpleLinearRegression(x, y)

	if slope != 0 || intercept != 0 {
		t.Errorf("Expected slope=0 and intercept=0 for zero denominator and slope, got slope=%.2f, intercept=%.2f", slope, intercept)
	}
}

func TestSimpleLinearRegression_SlopeZeroAndDenominatorZero_Explicit(t *testing.T) {
	// Identical x-values => zero denominator
	x := []float64{10, 10, 10}
	// Constant y-values => zero slope
	y := []float64{5, 5, 5}

	slope, intercept := SimpleLinearRegression(x, y)

	if slope != 0 || intercept != 0 {
		t.Errorf("Expected slope=0 and intercept=0 for flat x/y, got slope=%.2f, intercept=%.2f", slope, intercept)
	}
}

func TestSimpleLinearRegression_ForceZeroSlopeAndDenominator(t *testing.T) {
	x := []float64{5, 5, 5} // All x-values same → denominator = 0
	y := []float64{10, 10, 10} // All y-values same → slope = 0

	slope, intercept := SimpleLinearRegression(x, y)

	if slope != 0 || intercept != 0 {
		t.Errorf("Expected slope=0 and intercept=0 for constant x/y, got slope=%.2f, intercept=%.2f", slope, intercept)
	}
}

func TestSimpleLinearRegression_TriggerZeroSlopeAndZeroDenominator(t *testing.T) {
	x := []float64{2, 2, 2}
	y := []float64{5, 5, 5}

	slope, intercept := SimpleLinearRegression(x, y)

	if slope != 0 || intercept != 0 {
		t.Errorf("Expected slope=0 and intercept=0 for constant x and y, got slope=%.2f, intercept=%.2f", slope, intercept)
	}
}

func TestSimpleLinearRegression_ConstantXY(t *testing.T) {
	x := []float64{5, 5, 5, 5, 5}
	y := []float64{7, 7, 7, 7, 7}

	slope, intercept := SimpleLinearRegression(x, y)

	if slope != 0 || intercept != 0 {
		t.Errorf("Expected slope=0 and intercept=0 for constant x and y, got slope=%.2f, intercept=%.2f", slope, intercept)
	}
}

func TestRootMeanSquaredError(t *testing.T) {
	t.Run("Valid input", func(t *testing.T) {
		yTrue := []float64{3, -0.5, 2, 7}
		yPred := []float64{2.5, 0.0, 2, 8}
		expected := 0.612372  // sqrt(0.375)

		rmse := RootMeanSquaredError(yTrue, yPred)
		if rmse < expected-0.0001 || rmse > expected+0.0001 {
			t.Errorf("Expected RMSE ≈ %.6f, got %.6f", expected, rmse)
		}
	})

	t.Run("Mismatched lengths", func(t *testing.T) {
		yTrue := []float64{1, 2, 3}
		yPred := []float64{1, 2}
		rmse := RootMeanSquaredError(yTrue, yPred)
		if rmse != 0 {
			t.Errorf("Expected RMSE = 0 for mismatched lengths, got %.3f", rmse)
		}
	})

	t.Run("Empty input", func(t *testing.T) {
		yTrue := []float64{}
		yPred := []float64{}
		rmse := RootMeanSquaredError(yTrue, yPred)
		if rmse != 0 {
			t.Errorf("Expected RMSE = 0 for empty input, got %.3f", rmse)
		}
	})
}
