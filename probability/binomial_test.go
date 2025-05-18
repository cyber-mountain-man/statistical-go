package probability

import (
	"math"
	"testing"
)

func TestBinomialPMF(t *testing.T) {
	tests := []struct {
		name     string
		n, k     int
		p        float64
		expected float64
	}{
		{"Symmetric case", 5, 2, 0.5, 0.3125},
		{"All failures", 10, 0, 0.1, 0.348678},
		{"All successes", 10, 10, 0.9, 0.348678},
		{"Mid-range", 7, 3, 0.3, 0.226895},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinomialPMF(tt.n, tt.k, tt.p)
			if math.Abs(got-tt.expected) > 1e-6 {
				t.Errorf("BinomialPMF(%d, %d, %.2f) = %.6f; want %.6f", tt.n, tt.k, tt.p, got, tt.expected)
			}
		})
	}
}

func TestBinomialCDF(t *testing.T) {
	tests := []struct {
		name     string
		n, k     int
		p        float64
		expected float64
	}{
		{"Halfway", 5, 2, 0.5, 0.5},
		{"Full cumulative", 10, 10, 0.9, 1.0},
		{"Only first term", 10, 0, 0.1, 0.348678},
		{"Mid-range", 7, 3, 0.3, 0.873964},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinomialCDF(tt.n, tt.k, tt.p)
			if math.Abs(got-tt.expected) > 1e-6 {
				t.Errorf("BinomialCDF(%d, %d, %.2f) = %.6f; want %.6f", tt.n, tt.k, tt.p, got, tt.expected)
			}
		})
	}
}

func TestBinomialPMFPanic(t *testing.T) {
	tests := []struct {
		name    string
		n, k    int
		p       float64
	}{
		{"Negative n", -1, 0, 0.5},
		{"Negative k", 5, -1, 0.5},
		{"k > n", 5, 6, 0.5},
		{"p < 0", 5, 2, -0.1},
		{"p > 1", 5, 2, 1.1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("BinomialPMF(%d, %d, %.2f) did not panic", tt.n, tt.k, tt.p)
				}
			}()
			_ = BinomialPMF(tt.n, tt.k, tt.p)
		})
	}
}

func TestBinomialCDFPanic(t *testing.T) {
	tests := []struct {
		name    string
		n, k    int
		p       float64
	}{
		{"Negative n", -1, 0, 0.5},
		{"Negative k", 5, -1, 0.5},
		{"k > n", 5, 6, 0.5},
		{"p < 0", 5, 2, -0.1},
		{"p > 1", 5, 2, 1.1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("BinomialCDF(%d, %d, %.2f) did not panic", tt.n, tt.k, tt.p)
				}
			}()
			_ = BinomialCDF(tt.n, tt.k, tt.p)
		})
	}
}
