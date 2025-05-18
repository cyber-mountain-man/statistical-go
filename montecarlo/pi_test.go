package montecarlo

import (
	"math"
	"testing"
)

func TestEstimatePi(t *testing.T) {
	result := EstimatePi(100000)
	if math.Abs(result-3.1416) > 0.05 {
		t.Errorf("EstimatePi() = %.4f; want close to 3.1416", result)
	}
}

func TestEstimatePiParallel(t *testing.T) {
	tests := []struct {
		points      int
		workers     int
		description string
	}{
		{100000, 4, "normal case"},
		{100000, 1, "single worker"},
		{100000, 3, "non-divisible workers"},
		{100000, 0, "zero workers - fallback"},
		{100000, -2, "negative workers - fallback"},
	}

	for _, tt := range tests {
		result := EstimatePiParallel(tt.points, tt.workers)
		if math.Abs(result-3.1416) > 0.05 {
			t.Errorf("EstimatePiParallel(%d, %d) = %.4f; want close to 3.1416 [%s]",
				tt.points, tt.workers, result, tt.description)
		}
	}
}
