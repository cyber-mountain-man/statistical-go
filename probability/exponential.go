package probability

import "math"

// ExponentialPDF returns the probability density at x for a given rate λ.
func ExponentialPDF(x, lambda float64) float64 {
	if x < 0 || lambda <= 0 {
		panic("ExponentialPDF: x must be >= 0 and lambda must be > 0")
	}
	return lambda * math.Exp(-lambda*x)
}

// ExponentialCDF returns the cumulative probability for a given x and rate λ.
func ExponentialCDF(x, lambda float64) float64 {
	if x < 0 || lambda <= 0 {
		panic("ExponentialCDF: x must be >= 0 and lambda must be > 0")
	}
	return 1 - math.Exp(-lambda*x)
}
