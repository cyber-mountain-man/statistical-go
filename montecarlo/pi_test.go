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
	result := EstimatePiParallel(100000, 4)
	if math.Abs(result-3.1416) > 0.05 {
		t.Errorf("EstimatePiParallel() = %.4f; want close to 3.1416", result)
	}
}
