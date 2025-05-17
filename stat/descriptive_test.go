package stat

import (
	"math"
	"testing"
)

func almostEqual(a, b, epsilon float64) bool {
	return math.Abs(a-b) <= epsilon
}

func TestMean(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5}
	want := 3.0
	got := Mean(data)
	if !almostEqual(got, want, 1e-9) {
		t.Errorf("Mean() = %v; want %v", got, want)
	}
}

func TestMedian(t *testing.T) {
	data := []float64{5, 1, 3, 4, 2}
	want := 3.0
	got := Median(data)
	if got != want {
		t.Errorf("Median() = %v; want %v", got, want)
	}
}

func TestMode(t *testing.T) {
	data := []float64{1, 2, 2, 3, 4}
	want := 2.0
	got := Mode(data)
	if got != want {
		t.Errorf("Mode() = %v; want %v", got, want)
	}
}

func TestMinMax(t *testing.T) {
	data := []float64{1, 7, 3, -2, 5}
	minWant, maxWant := -2.0, 7.0
	minGot := Min(data)
	maxGot := Max(data)

	if minGot != minWant {
		t.Errorf("Min() = %v; want %v", minGot, minWant)
	}
	if maxGot != maxWant {
		t.Errorf("Max() = %v; want %v", maxGot, maxWant)
	}
}

func TestRange(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5}
	want := 4.0
	got := Range(data)
	if got != want {
		t.Errorf("Range() = %v; want %v", got, want)
	}
}

func TestQuartiles(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6, 7}
	q1, q2, q3 := Quartiles(data)

	if q1 != 2 {
		t.Errorf("Q1 = %v; want 2", q1)
	}
	if q2 != 4 {
		t.Errorf("Q2 = %v; want 4", q2)
	}
	if q3 != 6 {
		t.Errorf("Q3 = %v; want 6", q3)
	}
}

func TestMin(t *testing.T) {
	data := []float64{5, 3, 9, -1, 0}
	want := -1.0
	got := Min(data)
	if got != want {
		t.Errorf("Min() = %v; want %v", got, want)
	}
}

func TestMax(t *testing.T) {
	data := []float64{5, 3, 9, -1, 0}
	want := 9.0
	got := Max(data)
	if got != want {
		t.Errorf("Max() = %v; want %v", got, want)
	}
}

func TestRangeStat(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5}
	want := 4.0 // 5 - 1
	got := Range(data)
	if got != want {
		t.Errorf("Range() = %v; want %v", got, want)
	}
}

func TestQuartilesEven(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6}
	q1, q2, q3 := Quartiles(data)
	if q1 != 2 || q2 != 3.5 || q3 != 5 {
		t.Errorf("Quartiles() = Q1=%v Q2=%v Q3=%v; want Q1=2 Q2=3.5 Q3=5", q1, q2, q3)
	}
}

func TestQuartilesOdd(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6, 7}
	q1, q2, q3 := Quartiles(data)
	if q1 != 2 || q2 != 4 || q3 != 6 {
		t.Errorf("Quartiles() = Q1=%v Q2=%v Q3=%v; want Q1=2 Q2=4 Q3=6", q1, q2, q3)
	}
}

func TestMeanEmpty(t *testing.T) {
	if got := Mean([]float64{}); got != 0 {
		t.Errorf("Mean([]) = %v; want 0", got)
	}
}

func TestVariance(t *testing.T) {
	data := []float64{2, 4, 4, 4, 5, 5, 7, 9}
	want := 4.571428571428571
	got := Variance(data)
	if !almostEqual(got, want, 1e-9) {
		t.Errorf("Variance() = %v; want %v", got, want)
	}
}

func TestVarianceSmall(t *testing.T) {
	data := []float64{1}
	want := 0.0
	got := Variance(data)
	if got != want {
		t.Errorf("Variance(single value) = %v; want 0", got)
	}
}

func TestStdDev(t *testing.T) {
	data := []float64{2, 4, 4, 4, 5, 5, 7, 9}
	want := math.Sqrt(4.571428571428571)
	got := StdDev(data)
	if !almostEqual(got, want, 1e-9) {
		t.Errorf("StdDev() = %v; want %v", got, want)
	}
}

func TestMedianEven(t *testing.T) {
	data := []float64{1, 2, 3, 4}
	want := 2.5
	got := Median(data)
	if got != want {
		t.Errorf("Median(even) = %v; want %v", got, want)
	}
}

func TestModeTie(t *testing.T) {
	data := []float64{1, 2, 2, 3, 3}
	got := Mode(data)
	if got != 2 && got != 3 {
		t.Errorf("Mode(tie) = %v; want 2 or 3", got)
	}
}

func TestQuartilesSmallSet(t *testing.T) {
	q1, q2, q3 := Quartiles([]float64{10})
	if q1 != 10 || q2 != 10 || q3 != 10 {
		t.Errorf("Quartiles([10]) = Q1=%v Q2=%v Q3=%v; want 10,10,10", q1, q2, q3)
	}
}

func TestMedianEmpty(t *testing.T) {
	if got := Median([]float64{}); got != 0 {
		t.Errorf("Median([]) = %v; want 0", got)
	}
}

func TestModeAllUnique(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5}
	got := Mode(data)
	if got == 0 {
		t.Errorf("Mode(all unique) = %v; unexpected result", got)
	}
}

func TestMinEmpty(t *testing.T) {
	if got := Min([]float64{}); got != 0 {
		t.Errorf("Min([]) = %v; want 0", got)
	}
}

func TestMaxEmpty(t *testing.T) {
	if got := Max([]float64{}); got != 0 {
		t.Errorf("Max([]) = %v; want 0", got)
	}
}

func TestQuartilesTwoElement(t *testing.T) {
	q1, q2, q3 := Quartiles([]float64{5, 10})
	if q1 != 5 || q2 != 7.5 || q3 != 10 {
		t.Errorf("Quartiles([5,10]) = Q1=%v Q2=%v Q3=%v; want Q1=5 Q2=7.5 Q3=10", q1, q2, q3)
	}
}

func TestModeTieAgain(t *testing.T) {
	data := []float64{4, 4, 5, 5, 6}
	got := Mode(data)
	if got != 4 && got != 5 {
		t.Errorf("Mode(tie) = %v; want 4 or 5", got)
	}
}

func TestQuartilesEvenLength(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6}
	q1, q2, q3 := Quartiles(data)
	if q1 != 2 || q2 != 3.5 || q3 != 5 {
		t.Errorf("Quartiles([1,2,3,4,5,6]) = Q1=%v Q2=%v Q3=%v; want Q1=2 Q2=3.5 Q3=5", q1, q2, q3)
	}
}

func TestModeUnique(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5}
	got := Mode(data)
	if got != 1 && got != 2 && got != 3 && got != 4 && got != 5 {
		t.Errorf("Mode(unique) = %v; expected any value from input", got)
	}
}

func TestQuartilesUnsorted(t *testing.T) {
	data := []float64{12, 6, 16, 2, 10, 4, 14, 8}
	q1, q2, q3 := Quartiles(data)
	if q1 != 5 || q2 != 9 || q3 != 13 {
		t.Errorf("Quartiles(unsorted) = Q1=%v Q2=%v Q3=%v; want 5, 9, 13", q1, q2, q3)
	}
}

func TestModeAllSame(t *testing.T) {
	data := []float64{7, 7, 7, 7}
	want := 7.0
	got := Mode(data)
	if got != want {
		t.Errorf("Mode(all same) = %v; want %v", got, want)
	}
}

func TestQuartilesLongEven(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8}
	q1, q2, q3 := Quartiles(data)
	if q1 != 2.5 || q2 != 4.5 || q3 != 6.5 {
		t.Errorf("Quartiles([1-8]) = Q1=%v Q2=%v Q3=%v; want Q1=2.5 Q2=4.5 Q3=6.5", q1, q2, q3)
	}
}

func TestModeEmpty(t *testing.T) {
	data := []float64{}
	want := 0.0
	got := Mode(data)
	if got != want {
		t.Errorf("Mode([]) = %v; want %v", got, want)
	}
}

func TestModeEqualFrequency(t *testing.T) {
	data := []float64{1, 2, 3, 4}
	got := Mode(data)

	// Any value is technically valid here because all have the same frequency
	valid := false
	for _, v := range data {
		if got == v {
			valid = true
			break
		}
	}
	if !valid {
		t.Errorf("Mode(equal freq) = %v; expected one of %v", got, data)
	}
}

func TestQuartilesEmpty(t *testing.T) {
	q1, q2, q3 := Quartiles([]float64{})
	if q1 != 0 || q2 != 0 || q3 != 0 {
		t.Errorf("Quartiles([]) = Q1=%v Q2=%v Q3=%v; want 0,0,0", q1, q2, q3)
	}
}

func TestZScore(t *testing.T) {
	data := []float64{10, 20, 30, 40, 50}
	z := ZScore(30, data)
	if !almostEqual(z, 0, 1e-9) {
		t.Errorf("ZScore() = %v; want 0", z)
	}
}

func TestCovariance(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{2, 4, 6, 8, 10}
	cov := Covariance(x, y)
	if !almostEqual(cov, 5.0, 1e-9) {
		t.Errorf("Covariance() = %v; want 5.0", cov)
	}
}

func TestPearsonCorrelation(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{2, 4, 6, 8, 10}
	r := PearsonCorrelation(x, y)
	if !almostEqual(r, 1.0, 1e-9) {
		t.Errorf("PearsonCorrelation() = %v; want 1.0", r)
	}
}

func TestSkewness(t *testing.T) {
	data := []float64{2, 4, 6, 8, 10}
	skew := Skewness(data)
	if math.Abs(skew) > 1e-9 {
		t.Errorf("Skewness() = %v; want 0", skew)
	}
}

func TestKurtosis(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6}
	k := Kurtosis(data)
	if math.IsNaN(k) {
		t.Errorf("Kurtosis() = NaN; want real number")
	}
}

func TestZScore_EdgeCases(t *testing.T) {
	data := []float64{10, 10, 10}
	z := ZScore(10, data)
	if z != 0 {
		t.Errorf("ZScore with zero std dev = %v; want 0", z)
	}
	if ZScore(10, []float64{}) != 0 {
		t.Error("ZScore with empty slice should return 0")
	}
}

func TestCovariance_EdgeCases(t *testing.T) {
	a := []float64{1, 2}
	b := []float64{3}
	if Covariance(a, b) != 0 {
		t.Error("Covariance with unequal lengths should return 0")
	}
}

func TestPearsonCorrelation_EdgeCases(t *testing.T) {
	a := []float64{1, 2, 3}
	b := []float64{1}
	if PearsonCorrelation(a, b) != 0 {
		t.Error("PearsonCorrelation with unequal lengths should return 0")
	}
	flat := []float64{2, 2, 2}
	if PearsonCorrelation(flat, flat) != 0 {
		t.Error("PearsonCorrelation with zero std dev should return 0")
	}
}

func TestSkewness_EdgeCases(t *testing.T) {
	short := []float64{1, 2}
	if Skewness(short) != 0 {
		t.Error("Skewness with len < 3 should return 0")
	}
	flat := []float64{5, 5, 5}
	if Skewness(flat) != 0 {
		t.Error("Skewness with zero std dev should return 0")
	}
}

func TestKurtosis_EdgeCases(t *testing.T) {
	short := []float64{1, 2, 3}
	if Kurtosis(short) != 0 {
		t.Error("Kurtosis with len < 4 should return 0")
	}
	flat := []float64{4, 4, 4, 4}
	if Kurtosis(flat) != 0 {
		t.Error("Kurtosis with zero std dev should return 0")
	}
}
