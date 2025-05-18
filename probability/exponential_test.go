package probability

import (
	"math"
	"testing"
)

func TestExponentialPDF(t *testing.T) {
	result := ExponentialPDF(1, 1.5)
	expected := 1.5 * math.Exp(-1.5)
	if math.Abs(result-expected) > 1e-5 {
		t.Errorf("ExponentialPDF(1, 1.5) = %.6f; want %.6f", result, expected)
	}
}

func TestExponentialCDF(t *testing.T) {
	result := ExponentialCDF(2, 0.5)
	expected := 1 - math.Exp(-1.0)
	if math.Abs(result-expected) > 1e-5 {
		t.Errorf("ExponentialCDF(2, 0.5) = %.6f; want %.6f", result, expected)
	}
}

// --- Panic Tests for ExponentialPDF ---

func TestExponentialPDFPanicOnNegativeX(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for negative x in ExponentialPDF")
		}
	}()
	_ = ExponentialPDF(-1, 1.0)
}

func TestExponentialPDFPanicOnZeroLambda(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for zero lambda in ExponentialPDF")
		}
	}()
	_ = ExponentialPDF(1, 0.0)
}

func TestExponentialPDFPanicOnNegativeLambda(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for negative lambda in ExponentialPDF")
		}
	}()
	_ = ExponentialPDF(1, -2.0)
}

// --- Panic Tests for ExponentialCDF ---

func TestExponentialCDFPanicOnNegativeX(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for negative x in ExponentialCDF")
		}
	}()
	_ = ExponentialCDF(-1, 1.0)
}

func TestExponentialCDFPanicOnZeroLambda(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for zero lambda in ExponentialCDF")
		}
	}()
	_ = ExponentialCDF(2, 0.0)
}

func TestExponentialCDFPanicOnNegativeLambda(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for negative lambda in ExponentialCDF")
		}
	}()
	_ = ExponentialCDF(2, -0.5)
}
