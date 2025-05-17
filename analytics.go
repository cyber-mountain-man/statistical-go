package statistical

import (
	"github.com/cyber-mountain-man/statistical-go/montecarlo"
	"github.com/cyber-mountain-man/statistical-go/probability"
	"github.com/cyber-mountain-man/statistical-go/stat"
)

// Descriptive Stats
func Mean(data []float64) float64        { return stat.Mean(data) }
func Variance(data []float64) float64    { return stat.Variance(data) }
func StdDev(data []float64) float64      { return stat.StdDev(data) }
func Median(data []float64) float64      { return stat.Median(data) }
func Mode(data []float64) float64        { return stat.Mode(data) }
func Range(data []float64) float64       { return stat.Range(data) }
func Min(data []float64) float64         { return stat.Min(data) }
func Max(data []float64) float64         { return stat.Max(data) }
func Quartiles(data []float64) (float64, float64, float64) {
	return stat.Quartiles(data)
}

// Probability
func ConditionalProbability(probAandB, probB float64) float64 {
	return probability.Conditional(probAandB, probB)
}

// Monte Carlo
func EstimatePi(iterations int) float64 {
	return montecarlo.EstimatePi(iterations)
}

func EstimatePiParallel(total, workers int) float64 {
	return montecarlo.EstimatePiParallel(total, workers)
}
