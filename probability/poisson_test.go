package probability

import (
	"math"
	"testing"
)

func TestPoissonPMF(t *testing.T) {
	tests := []struct {
		k        int
		lambda   float64
		expected float64
	}{
		{0, 2.0, 0.135335},
		{1, 2.0, 0.270670},
		{2, 2.0, 0.270670},
		{3, 2.0, 0.180447},
	}

	for _, tt := range tests {
		got := PoissonPMF(tt.k, tt.lambda)
		if math.Abs(got-tt.expected) > 1e-5 {
			t.Errorf("PoissonPMF(%d, %.2f) = %.6f; want %.6f", tt.k, tt.lambda, got, tt.expected)
		}
	}
}

func TestPoissonCDF(t *testing.T) {
	tests := []struct {
		k        int
		lambda   float64
		expected float64
	}{
		{2, 2.0, 0.676676},
		{3, 2.0, 0.857123},
	}

	for _, tt := range tests {
		got := PoissonCDF(tt.k, tt.lambda)
		if math.Abs(got-tt.expected) > 1e-5 {
			t.Errorf("PoissonCDF(%d, %.2f) = %.6f; want %.6f", tt.k, tt.lambda, got, tt.expected)
		}
	}
}

func TestPoissonPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("PoissonPMF or CDF should panic on invalid input")
		}
	}()
	_ = PoissonPMF(-1, 1.0)
	_ = PoissonCDF(2, -1.0)
}

func TestPoissonPMFPanicOnNegativeLambda(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for negative lambda in PoissonPMF")
		}
	}()
	_ = PoissonPMF(2, -1.0)
}

func TestFactorialPanicOnNegativeInput(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for negative input to factorial")
		}
	}()
	_ = factorial(-1)
}

func TestFactorialOverflowPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for large input to factorial due to overflow risk")
		}
	}()
	_ = factorial(200) // Should overflow float64 range
}

func TestPoissonCDFPanicOnNegativeK(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for negative k in PoissonCDF")
		}
	}()
	_ = PoissonCDF(-1, 2.0)
}

func TestPoissonCDFPanicOnNonPositiveLambda(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for non-positive lambda in PoissonCDF")
		}
	}()
	_ = PoissonCDF(3, 0.0)
}
