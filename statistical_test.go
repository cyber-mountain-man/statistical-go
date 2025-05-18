package statistical

import (
	"testing"
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
}
