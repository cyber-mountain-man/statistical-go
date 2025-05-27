package stat

import (
	"math"
	"testing"
)

func TestExponentialSmoothing(t *testing.T) {
	data := []float64{50, 55, 53, 60, 62, 65}
	alpha := 0.3

	expected := 59.158850
	result := ExponentialSmoothing(data, alpha)

	if math.Abs(result-expected) > 1e-4 {
		t.Errorf("ExponentialSmoothing() = %.6f; want %.6f", result, expected)
	}
}

func TestExponentialSmoothing_EmptyData(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for empty data but did not get one")
		}
	}()
	_ = ExponentialSmoothing([]float64{}, 0.3)
}

func TestExponentialSmoothing_InvalidAlpha(t *testing.T) {
	tests := []struct {
		alpha float64
	}{
		{0.0},
		{-0.5},
		{1.1},
	}

	for _, tt := range tests {
		func(a float64) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Expected panic for alpha=%.2f but did not get one", a)
				}
			}()
			_ = ExponentialSmoothing([]float64{10, 20, 30}, a)
		}(tt.alpha)
	}
}
