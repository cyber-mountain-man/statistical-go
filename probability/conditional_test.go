package probability

import "testing"

func TestConditional(t *testing.T) {
	tests := []struct {
		pAandB float64
		pB     float64
		want   float64
	}{
		{0.15, 0.3, 0.5},
		{0.2, 0.4, 0.5},
		{0.0, 0.5, 0.0},
	}

	for _, tt := range tests {
		got := Conditional(tt.pAandB, tt.pB)
		if !floatEquals(got, tt.want) {
			t.Errorf("Conditional(%.2f, %.2f) = %.2f; want %.2f", tt.pAandB, tt.pB, got, tt.want)
		}
	}
}

func TestConditionalPanic(t *testing.T) {
	tests := [][2]float64{
		{0.3, 0.0},  // Zero denominator
		{0.3, -0.1}, // Negative pB
		{0.3, 1.2},  // pB > 1
		{-0.1, 0.5}, // pAandB < 0
		{1.2, 0.5},  // pAandB > 1
	}

	for _, pair := range tests {
		func(aAndB, b float64) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Conditional(%v, %v) did not panic", aAndB, b)
				}
			}()
			_ = Conditional(aAndB, b)
		}(pair[0], pair[1])
	}
}
