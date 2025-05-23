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

func TestBinomialPMFInvalidInputs(t *testing.T) {
	tests := []struct {
		name string
		n, k int
		p    float64
	}{
		{"Negative n", -1, 0, 0.5},
		{"Negative k", 5, -1, 0.5},
		{"k > n", 5, 6, 0.5},
		{"p < 0", 5, 2, -0.1},
		{"p > 1", 5, 2, 1.1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinomialPMF(tt.n, tt.k, tt.p)
			if got != 0.0 {
				t.Errorf("BinomialPMF(%d, %d, %.2f) = %.6f; want 0.0 for invalid input", tt.n, tt.k, tt.p, got)
			}
		})
	}
}

func TestBinomialCDFInvalidInputs(t *testing.T) {
	tests := []struct {
		name string
		n, k int
		p    float64
	}{
		{"Negative n", -1, 0, 0.5},
		{"Negative k", 5, -1, 0.5},
		{"k > n", 5, 6, 0.5},
		{"p < 0", 5, 2, -0.1},
		{"p > 1", 5, 2, 1.1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinomialCDF(tt.n, tt.k, tt.p)
			if got != 0.0 {
				t.Errorf("BinomialCDF(%d, %d, %.2f) = %.6f; want 0.0 for invalid input", tt.n, tt.k, tt.p, got)
			}
		})
	}
}

func TestBinomialQuantile(t *testing.T) {
	tests := []struct {
		name     string
		targetP  float64
		n        int
		p        float64
		expected int
	}{
		{"Target 0.5", 0.5, 5, 0.5, 2},
		{"Target 0.8", 0.8, 10, 0.6, 7},
		{"Target 0.99", 0.99, 10, 0.8, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinomialQuantile(tt.targetP, tt.n, tt.p)
			if got != tt.expected {
				t.Errorf("BinomialQuantile(%.2f, %d, %.2f) = %d; want %d",
					tt.targetP, tt.n, tt.p, got, tt.expected)
			}
		})
	}
}

func TestBinomialQuantileInvalidInputs(t *testing.T) {
	tests := []struct {
		name    string
		targetP float64
		n       int
		p       float64
	}{
		{"targetP < 0", -0.1, 5, 0.5},
		{"targetP > 1", 1.1, 5, 0.5},
		{"n < 0", 0.5, -1, 0.5},
		{"p < 0", 0.5, 5, -0.1},
		{"p > 1", 0.5, 5, 1.1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinomialQuantile(tt.targetP, tt.n, tt.p)
			expected := tt.n
			if got != expected {
				t.Errorf("expected fallback return %d for invalid input, got %d", expected, got)
			}
		})
	}
}

func TestBinomialQuantileFallback(t *testing.T) {
	n := 10
	p := 1e-9
	targetP := 0.9999

	got := BinomialQuantile(targetP, n, p)

	if got < 0 || got > n {
		t.Errorf("expected fallback within range [0, %d], got %d", n, got)
	}
}

func TestStatisticalEdgeCases(t *testing.T) {
	t.Run("BinomialPMF returns 0.0 for invalid input", func(t *testing.T) {
		got := BinomialPMF(-5, 2, 0.5)
		if got != 0.0 {
			t.Errorf("Expected 0.0 for invalid input, got %.6f", got)
		}
	})

	t.Run("BinomialCDF returns 0.0 for invalid input", func(t *testing.T) {
		got := BinomialCDF(5, 6, 0.5)
		if got != 0.0 {
			t.Errorf("Expected 0.0 for invalid input, got %.6f", got)
		}
	})

	t.Run("BinomialQuantile returns n for invalid input", func(t *testing.T) {
		got := BinomialQuantile(1.5, 10, 0.5) // targetP > 1
		if got != 10 {
			t.Errorf("Expected fallback return of n=10 for invalid input, got %d", got)
		}
	})
}
// probability/binomial_test.go
// … existing tests …

func TestBinomialQuantilePrecisionFallback(t *testing.T) {
    n := 5
    p := 0.4
    // Make the target probability just (numerically) above the true CDF,
    // so the loop never crosses it and the fallback branch is exercised.
    target := BinomialCDF(n, n, p) + 1e-12

    got := BinomialQuantile(target, n, p)
    if got != n {
        t.Errorf("expected fallback %d, got %d", n, got)
    }
}

// ---------------------------------------------------------------------------
// Extra edge-case coverage for BinomialQuantile: targetP = 0 and 1
// ---------------------------------------------------------------------------

func TestBinomialQuantileEdgeCases(t *testing.T) {
	n := 7
	p := 0.3

	// targetP = 0  →  should always return 0
	if got := BinomialQuantile(0.0, n, p); got != 0 {
		t.Errorf("BinomialQuantile(0, %d, %.2f) = %d; want 0", n, p, got)
	}

	// targetP = 1  →  should always return n
	if got := BinomialQuantile(1.0, n, p); got != n {
		t.Errorf("BinomialQuantile(1, %d, %.2f) = %d; want %d", n, p, got, n)
	}
}


