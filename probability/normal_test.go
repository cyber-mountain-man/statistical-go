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
		{0, 0, 1, 1 / math.Sqrt(2 * math.Pi)},
		{1, 0, 1, math.Exp(-0.5) / math.Sqrt(2 * math.Pi)},
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
		{-10, 0, 1, 0},   // Far below mean
		{10, 0, 1, 1},    // Far above mean
		{0, 0, 0, 0},     // stdDev = 0 should return 0
	}

	for _, tt := range tests {
		result := NormalCDF(tt.x, tt.mean, tt.stdDev)
		if math.Abs(result-tt.expected) > 0.0001 {
			t.Errorf("NormalCDF(%.2f, %.2f, %.2f) = %.5f; want %.5f", tt.x, tt.mean, tt.stdDev, result, tt.expected)
		}
	}
}
