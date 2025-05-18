package statistical

import (
	"testing"
)

func TestStatisticalWrappers(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5}
	x := []float64{1, 2, 3}
	y := []float64{2, 4, 6}

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

	_ = ConditionalProbability(0.15, 0.3)
	_ = Complement(0.2)
	_ = AdditionRule(0.6, 0.3, 0.1)
	_ = MultiplicationRuleIndependent(0.5, 0.4)
	_ = MultiplicationRuleDependent(0.5, 0.6)

	_ = EstimatePi(1000)
	_ = EstimatePiParallel(1000, 2)
}
