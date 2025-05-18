package probability

import (
	"math"
	"testing"
)

func TestNormalPDF(t *testing.T) {
	tests := []struct {
		x, mean, stdDev float64
		expected        float64
	}{
		{0, 0, 1, 1 / math.Sqrt(2*math.Pi)},
		{1, 0, 1, math.Exp(-0.5) / math.Sqrt(2*math.Pi)},
		{0, 0, 0, 0}, // stdDev = 0 should return 0
	}

	for _, tt := range tests {
		result := NormalPDF(tt.x, tt.mean, tt.stdDev)
		if math.Abs(result-tt.expected) > 0.0001 {
			t.Errorf("NormalPDF(%.2f, %.2f, %.2f) = %.5f; want %.5f", tt.x, tt.mean, tt.stdDev, result, tt.expected)
		}
	}
}

func TestNormalCDF(t *testing.T) {
	tests := []struct {
		x, mean, stdDev float64
		expected        float64
	}{
		{0, 0, 1, 0.5},
		{-10, 0, 1, 0}, // Far below mean
		{10, 0, 1, 1},  // Far above mean
		{0, 0, 0, 0},   // stdDev = 0 should return 0
	}

	for _, tt := range tests {
		result := NormalCDF(tt.x, tt.mean, tt.stdDev)
		if math.Abs(result-tt.expected) > 0.0001 {
			t.Errorf("NormalCDF(%.2f, %.2f, %.2f) = %.5f; want %.5f", tt.x, tt.mean, tt.stdDev, result, tt.expected)
		}
	}
}

func TestNormalInverseCDF(t *testing.T) {
	tests := []struct {
		p        float64
		mean     float64
		stddev   float64
		expected float64 // approximate expected value
	}{
		{0.5, 0, 1, 0},
		{0.8413, 0, 1, 1},   // ~1 SD above the mean
		{0.1587, 0, 1, -1},  // ~1 SD below the mean
		{0.975, 0, 1, 1.96}, // ~1.96 SDs above mean
	}

	for _, tt := range tests {
		got := NormalInverseCDF(tt.p, tt.mean, tt.stddev)
		if math.Abs(got-tt.expected) > 0.05 {
			t.Errorf("NormalInverseCDF(%v, %v, %v) = %v; want ~%v",
				tt.p, tt.mean, tt.stddev, got, tt.expected)
		}
	}
}

func TestNormalInverseCDF_AlternateCases(t *testing.T) {
	tests := []struct {
		p        float64
		mean     float64
		stddev   float64
		expected float64
		tol      float64
	}{
		{0.5, 0.0, 1.0, 0.0, 1e-5},
		{0.8413447, 0.0, 1.0, 1.0, 1e-5},
		{0.9772499, 0.0, 1.0, 2.0, 1e-5},
		{0.5, 10.0, 2.0, 10.0, 1e-5},
		{0.8413447, 5.0, 2.0, 7.0, 1e-5},

		// 🧪 Edge cases to trigger all branches of standardNormalInverseCDF
		{0.0001, 0.0, 1.0, -3.719, 0.05}, // far left tail
		{0.9999, 0.0, 1.0, 3.719, 0.05},  // far right tail
	}

	for _, tt := range tests {
		got := NormalInverseCDF(tt.p, tt.mean, tt.stddev)
		if math.Abs(got-tt.expected) > tt.tol {
			t.Errorf("NormalInverseCDF(%.7f, %.1f, %.1f) = %.6f; want %.6f",
				tt.p, tt.mean, tt.stddev, got, tt.expected)
		}
	}
}

func TestNormalInverseCDFPanic(t *testing.T) {
	badInputs := []struct {
		p      float64
		mean   float64
		stddev float64
	}{
		{0.0, 0, 1},
		{1.0, 0, 1},
		{0.5, 0, 0},
	}

	for _, input := range badInputs {
		func() {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Expected panic for NormalInverseCDF(%.2f, %.2f, %.2f)", input.p, input.mean, input.stddev)
				}
			}()
			_ = NormalInverseCDF(input.p, input.mean, input.stddev)
		}()
	}
}
