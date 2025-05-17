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
		{0.3, 0.0, 0.0}, // handle division by zero
	}

	for _, tt := range tests {
		got := Conditional(tt.pAandB, tt.pB)
		if got != tt.want {
			t.Errorf("Conditional(%.2f, %.2f) = %.2f; want %.2f", tt.pAandB, tt.pB, got, tt.want)
		}
	}
}
