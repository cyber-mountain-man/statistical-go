package hypothesis

import (
	"math"
	"testing"
)

func TestOneWayANOVA(t *testing.T) {
	tests := []struct {
		name     string
		groups   [][]float64
		expected float64
		tol      float64
		isNaN    bool
	}{
		{
			name: "Normal case with variance",
			groups: [][]float64{
				{4.0, 5.0, 6.0},
				{10.0, 9.0, 11.0},
				{7.0, 8.0, 9.0},
			},
			expected: 19.0,
			tol:      0.1,
			isNaN:    false,
		},
		{
			name: "All groups identical",
			groups: [][]float64{
				{5.0, 5.0, 5.0},
				{5.0, 5.0, 5.0},
				{5.0, 5.0, 5.0},
			},
			isNaN: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OneWayANOVA(tt.groups)

			if tt.isNaN {
				if !math.IsNaN(got.Statistic) {
					t.Errorf("Expected NaN, got Statistic = %.4f", got.Statistic)
				}
				if !math.IsNaN(got.PValue) {
					t.Errorf("Expected NaN, got PValue = %.4f", got.PValue)
				}
			} else if math.Abs(got.Statistic-tt.expected) > tt.tol {
				t.Errorf("OneWayANOVA().Statistic = %.4f; want %.4f", got.Statistic, tt.expected)
			}

			if got.PValue < 0.0 || got.PValue > 1.0 {
				t.Errorf("P-value out of bounds: got %.4f", got.PValue)
			}
		})
	}
}

func TestOneWayANOVAPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for fewer than two groups")
		}
	}()
	_ = OneWayANOVA([][]float64{
		{1.0, 2.0, 3.0},
	}) // only one group
}

func TestFDistributionCDFInvalidDF(t *testing.T) {
	p := fDistributionCDF(1.0, -1, 5) // dfn <= 0
	if p != 0 {
		t.Errorf("Expected 0 for invalid dfn, got %v", p)
	}

	p = fDistributionCDF(1.0, 5, 0) // dfd <= 0
	if p != 0 {
		t.Errorf("Expected 0 for invalid dfd, got %v", p)
	}
}
