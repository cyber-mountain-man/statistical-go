package hypothesis

import (
	"math"
	"testing"
)

func TestOneSampleTTest_Valid(t *testing.T) {
	sampleMean := 105.0
	popMean := 100.0
	stdDev := 15.0
	n := 30

	expected := (sampleMean - popMean) / (stdDev / math.Sqrt(float64(n)))
	result := OneSampleTTest(sampleMean, popMean, stdDev, n)

	if math.Abs(result.Statistic-expected) > 1e-6 {
		t.Errorf("OneSampleTTest() = %f; want %f", result.Statistic, expected)
	}
	if result.PValue < 0 || result.PValue > 1 {
		t.Errorf("PValue should be between 0 and 1; got %f", result.PValue)
	}
}

func TestOneSampleTTest_SampleSizeTooSmall(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for n <= 1")
		}
	}()
	_ = OneSampleTTest(100, 100, 10, 1)
}

func TestOneSampleTTest_NonPositiveStdDev(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for stdDev <= 0")
		}
	}()
	_ = OneSampleTTest(100, 95, 0, 10)
}

func TestTwoSampleTTestWelch(t *testing.T) {
	mean1, mean2 := 100.0, 95.0
	stdDev1, stdDev2 := 10.0, 12.0
	n1, n2 := 30, 25

	got := TwoSampleTTestWelch(mean1, mean2, stdDev1, stdDev2, n1, n2)
	se := math.Sqrt((stdDev1*stdDev1)/float64(n1) + (stdDev2*stdDev2)/float64(n2))
	expected := (mean1 - mean2) / se

	if math.Abs(got.Statistic-expected) > 1e-5 {
		t.Errorf("TwoSampleTTestWelch() = %.5f; want %.5f", got.Statistic, expected)
	}
	if got.PValue < 0 || got.PValue > 1 {
		t.Errorf("PValue should be between 0 and 1; got %f", got.PValue)
	}
}

func TestTwoSampleTTestWelch_InvalidSampleSize(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for n1 or n2 <= 1")
		}
	}()
	_ = TwoSampleTTestWelch(100, 95, 10, 12, 1, 25)
}

func TestTwoSampleTTestWelch_InvalidStdDev(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for stdDev1 or stdDev2 <= 0")
		}
	}()
	_ = TwoSampleTTestWelch(100, 95, 0, 12, 30, 25)
}

func TestPairedTTest(t *testing.T) {
	x := []float64{85, 90, 78, 92, 88}
	y := []float64{80, 88, 75, 89, 85}

	got := PairedTTest(x, y)

	n := float64(len(x))
	diffs := make([]float64, len(x))
	sum := 0.0
	for i := range x {
		diffs[i] = x[i] - y[i]
		sum += diffs[i]
	}
	meanDiff := sum / n

	var sqDiffs float64
	for _, d := range diffs {
		sqDiffs += math.Pow(d-meanDiff, 2)
	}
	stdDev := math.Sqrt(sqDiffs / (n - 1))
	expected := meanDiff / (stdDev / math.Sqrt(n))

	if math.Abs(got.Statistic-expected) > 1e-5 {
		t.Errorf("PairedTTest() = %.5f; want %.5f", got.Statistic, expected)
	}
	if got.PValue < 0 || got.PValue > 1 {
		t.Errorf("PValue should be between 0 and 1; got %f", got.PValue)
	}
}

func TestPairedTTest_LengthMismatch(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for unequal lengths")
		}
	}()
	_ = PairedTTest([]float64{1, 2}, []float64{1})
}

func TestPairedTTest_TooShort(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for length < 2")
		}
	}()
	_ = PairedTTest([]float64{1}, []float64{1})
}
