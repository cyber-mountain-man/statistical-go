package hypothesis

import (
	"math"
	"testing"
)

func TestOneSampleZTest(t *testing.T) {
	res := OneSampleZTest(105, 100, 15, 30)
	expected := (105 - 100) / (15 / math.Sqrt(30))

	if math.Abs(res.Statistic-expected) > 1e-5 {
		t.Errorf("OneSampleZTest Statistic = %f; want %f", res.Statistic, expected)
	}
	if res.PValue <= 0 || res.PValue >= 1 {
		t.Errorf("Unexpected PValue: got %f", res.PValue)
	}
}

func TestOneSampleZTest_ZeroStdDev(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic with zero standard deviation")
		}
	}()
	_ = OneSampleZTest(105, 100, 0, 30)
}

func TestTwoSampleZTest(t *testing.T) {
	mean1, mean2 := 102.0, 100.0
	stdDev1, stdDev2 := 14.0, 15.0
	n1, n2 := 40, 35

	res := TwoSampleZTest(mean1, mean2, stdDev1, stdDev2, n1, n2)
	se := math.Sqrt((stdDev1*stdDev1)/float64(n1) + (stdDev2*stdDev2)/float64(n2))
	expected := (mean1 - mean2) / se

	if math.Abs(res.Statistic-expected) > 1e-5 {
		t.Errorf("TwoSampleZTest Statistic = %f; want %f", res.Statistic, expected)
	}
	if res.PValue <= 0 || res.PValue >= 1 {
		t.Errorf("Unexpected PValue: got %f", res.PValue)
	}
}

func TestTwoSampleZTest_ZeroStdDev(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic with zero standard deviations")
		}
	}()
	_ = TwoSampleZTest(102, 100, 0, 0, 40, 35)
}

func TestOneSampleZTest_InvalidInput(t *testing.T) {
	tests := []struct {
		sampleMean float64
		popMean    float64
		popStdDev  float64
		n          int
	}{
		{5.0, 5.0, 0.0, 10},
		{5.0, 5.0, 1.0, 0},
	}

	for _, tt := range tests {
		func() {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Expected panic with input: %+v", tt)
				}
			}()
			_ = OneSampleZTest(tt.sampleMean, tt.popMean, tt.popStdDev, tt.n)
		}()
	}
}

func TestTwoSampleZTest_InvalidInput(t *testing.T) {
	tests := []struct {
		stdDev1, stdDev2 float64
		n1, n2           int
	}{
		{0.0, 1.0, 30, 30},
		{1.0, -1.0, 30, 30},
		{1.0, 1.0, 0, 30},
		{1.0, 1.0, 30, 0},
	}

	for _, tt := range tests {
		func() {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Expected panic with input: %+v", tt)
				}
			}()
			_ = TwoSampleZTest(5.0, 4.5, tt.stdDev1, tt.stdDev2, tt.n1, tt.n2)
		}()
	}
}
