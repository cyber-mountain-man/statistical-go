package statistical

import (
	"github.com/cyber-mountain-man/statistical-go/stat"
	"github.com/cyber-mountain-man/statistical-go/probability"
	"github.com/cyber-mountain-man/statistical-go/montecarlo"
)

// --- Descriptive Stats Wrappers ---
func Mean(data []float64) float64 {
	return stat.Mean(data)
}

func Variance(data []float64) float64 {
	return stat.Variance(data)
}

func StdDev(data []float64) float64 {
	return stat.StdDev(data)
}

// --- Probability Wrappers ---
func ConditionalProbability(probAandB, probB float64) float64 {
	return probability.Conditional(probAandB, probB)
}

// --- Monte Carlo Wrappers ---
func EstimatePi(iterations int) float64 {
	return montecarlo.EstimatePi(iterations)
}

func EstimatePiParallel(total, workers int) float64 {
	return montecarlo.EstimatePiParallel(total, workers)
}
