package probability

import (
	"math"
)

// PoissonPMF returns the probability of observing k events
// given the average rate λ using the Poisson distribution.
func PoissonPMF(k int, lambda float64) float64 {
	if k < 0 || lambda <= 0 {
		panic("PoissonPMF: k must be >= 0 and lambda must be > 0")
	}
	return math.Pow(lambda, float64(k)) * math.Exp(-lambda) / float64(factorial(k))
}

// PoissonCDF returns the cumulative probability of observing
// 0 through k events in a Poisson distribution with mean λ.
func PoissonCDF(k int, lambda float64) float64 {
	if k < 0 || lambda <= 0 {
		panic("PoissonCDF: k must be >= 0 and lambda must be > 0")
	}
	var sum float64
	for i := 0; i <= k; i++ {
		sum += PoissonPMF(i, lambda)
	}
	return sum
}

// Helper: integer factorial
func factorial(n int) float64 {
	if n < 0 {
		panic("factorial: n must be non-negative")
	}
	result := 1.0
	for i := 2; i <= n; i++ {
		result *= float64(i)
		if math.IsInf(result, 0) {
			panic("factorial: result overflowed float64")
		}
	}
	return result
}
