package statistical

import (
	"github.com/cyber-mountain-man/statistical-go/montecarlo"
	"github.com/cyber-mountain-man/statistical-go/probability"
	"github.com/cyber-mountain-man/statistical-go/stat"
)

//
// DESCRIPTIVE STATISTICS
//

func Mean(data []float64) float64     { return stat.Mean(data) }
func Variance(data []float64) float64 { return stat.Variance(data) }
func StdDev(data []float64) float64   { return stat.StdDev(data) }
func Median(data []float64) float64   { return stat.Median(data) }
func Mode(data []float64) float64     { return stat.Mode(data) }
func Range(data []float64) float64    { return stat.Range(data) }
func Min(data []float64) float64      { return stat.Min(data) }
func Max(data []float64) float64      { return stat.Max(data) }
func Quartiles(data []float64) (float64, float64, float64) {
	return stat.Quartiles(data)
}
func ZScore(x float64, data []float64) float64  { return stat.ZScore(x, data) }
func Covariance(x, y []float64) float64         { return stat.Covariance(x, y) }
func PearsonCorrelation(x, y []float64) float64 { return stat.PearsonCorrelation(x, y) }
func Skewness(data []float64) float64           { return stat.Skewness(data) }
func Kurtosis(data []float64) float64           { return stat.Kurtosis(data) }

//
// PROBABILITY
//

func ConditionalProbability(probAandB, probB float64) float64 {
	return probability.Conditional(probAandB, probB)
}

func Complement(probA float64) float64 {
	return probability.ComplementRule(probA)
}

func AdditionRule(probA, probB, probAandB float64) float64 {
	return probability.AdditionRule(probA, probB, probAandB)
}

func MultiplicationRuleIndependent(probA, probB float64) float64 {
	return probability.MultiplicationIndependent(probA, probB)
}

func MultiplicationRuleDependent(probA, probBgivenA float64) float64 {
	return probability.MultiplicationDependent(probA, probBgivenA)
}

//
// MONTE CARLO SIMULATIONS
//

func EstimatePi(iterations int) float64 {
	return montecarlo.EstimatePi(iterations)
}

func EstimatePiParallel(total, workers int) float64 {
	return montecarlo.EstimatePiParallel(total, workers)
}
