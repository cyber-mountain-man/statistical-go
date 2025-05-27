package statistical

import (
	"math"
	"testing"

	"github.com/cyber-mountain-man/statistical-go/stat"
)


func TestStatisticalWrappers(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5}
	x := []float64{1, 2, 3}
	y := []float64{2, 4, 6}

	// Descriptive statistics
	_ = Mean(data)
	_ = Variance(data)
	_ = StdDev(data)
	_ = Median(data)
	_ = Mode(data)
	_ = Range(data)
	_ = Min(data)
	_ = Max(data)
	_, _, _ = Quartiles(data)
	_ = ZScore(3.0, data)
	_ = Covariance(x, y)
	_ = PearsonCorrelation(x, y)
	_ = Skewness(data)
	_ = Kurtosis(data)

	// Probability
	_ = ConditionalProbability(0.15, 0.3)
	_ = Complement(0.2)
	_ = AdditionRule(0.6, 0.3, 0.1)
	_ = MultiplicationRuleIndependent(0.5, 0.4)
	_ = MultiplicationRuleDependent(0.5, 0.6)

	// Monte Carlo
	_ = EstimatePi(1000)
	_ = EstimatePiParallel(1000, 2)

	// Normal Distribution
	_ = NormalPDF(1.0, 0.0, 1.0)
	_ = NormalCDF(1.0, 0.0, 1.0)

	// Exercise the inverse CDF wrapper
    _ = NormalInverseCDF(0.975, 0.0, 1.0) // 97.5 % quantile of N(0,1)

	// Binomial Distribution
	_ = BinomialPMF(5, 2, 0.5)
	_ = BinomialCDF(5, 2, 0.5)

	// Z-Tests
	_, _ = OneSampleZTest(100, 95, 10, 30)
	_, _ = TwoSampleZTest(100, 95, 10, 12, 30, 30)

	// T-Tests
	_, _ = OneSampleTTest(100, 95, 10, 30)
	_, _ = TwoSampleTTestWelch(100, 95, 10, 12, 30, 30)
	_, _ = PairedTTest([]float64{1, 2, 3}, []float64{1, 2, 2})

	// Chi-Square
	_, _, _ = ChiSquareGoodnessOfFit([]float64{10, 20}, []float64{15, 15})
	_, _, _ = ChiSquareTestOfIndependence([][]float64{{10, 20}, {30, 40}})

	// ANOVA
	_, _, _ = OneWayANOVA([][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})

	// Linear Regression
	slope, intercept := SimpleLinearRegression([]float64{1, 2, 3}, []float64{2, 4, 6})
	_ = PredictY(4, slope, intercept)
	_ = RMSE([]float64{3, -0.5, 2, 7}, []float64{2.5, 0.0, 2, 8})
	_ = MSE([]float64{3, -0.5, 2, 7}, []float64{2.5, 0.0, 2, 8})
	_ = EvaluateLinearRegression([]float64{1, 2, 3, 4, 5}, []float64{2, 4, 5, 4, 5})
	_ = RSquared([]float64{1, 2, 3, 4, 5}, []float64{2, 4, 5, 4, 5}, slope, intercept)

	// Regularized Regression
	X := [][]float64{
		{1, 2},
		{2, 3},
		{3, 4},
	}
	yVec := []float64{5, 8, 11}

	_, _ = RidgeRegression(X, yVec, 0.1)
	_, _ = LassoRegression(X, yVec, 0.1, 100)
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
		expected := 10                        // fallback value for invalid input
		if got != expected {
			t.Errorf("Expected fallback return of %d for invalid input, got %d", expected, got)
		}
	})
}

func TestStatisticalWrapper_ExponentialSmoothing(t *testing.T) {
	data := []float64{50, 52, 58, 60, 63}
	alpha := 0.5

	// Call the unified wrapper
	got := ExponentialSmoothing(data, alpha)


	// Expected from the stat package directly
	expected := stat.ExponentialSmoothing(data, alpha)

	if math.Abs(got-expected) > 1e-6 {
		t.Errorf("ExponentialSmoothing wrapper mismatch: got %.6f, want %.6f", got, expected)
	}
}