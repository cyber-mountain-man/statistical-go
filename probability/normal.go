package probability

import (
	"math"
)

// NormalPDF computes the probability density function for a normal distribution.
// μ = mean, σ = standard deviation
func NormalPDF(x, mean, stdDev float64) float64 {
	if stdDev <= 0 {
		return math.NaN()
	}
	expPart := math.Exp(-math.Pow(x-mean, 2) / (2 * math.Pow(stdDev, 2)))
	coef := 1 / (stdDev * math.Sqrt(2*math.Pi))
	return coef * expPart
}

// NormalCDF computes the cumulative distribution function using the error function approximation.
func NormalCDF(x, mean, stdDev float64) float64 {
	if stdDev <= 0 {
		return math.NaN()
	}
	z := (x - mean) / (stdDev * math.Sqrt2)
	return 0.5 * (1 + math.Erf(z))
}


