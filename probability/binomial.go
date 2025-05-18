package probability

import (
	"log"
	"math"
)

// BinomialPMF returns the probability of getting exactly k successes in n trials
// with success probability p (P(X = k)). It panics on invalid inputs.
func BinomialPMF(n, k int, p float64) float64 {
	if n < 0 || k < 0 || k > n || p < 0 || p > 1 {
		log.Panic("BinomialPMF: invalid inputs (n, k, p must be in valid ranges)")
	}
	coef := binomialCoefficient(n, k)
	return coef * math.Pow(p, float64(k)) * math.Pow(1-p, float64(n-k))
}

// BinomialCDF returns the cumulative probability of getting at most k successes (P(X ≤ k)).
// It panics on invalid inputs.
func BinomialCDF(n, k int, p float64) float64 {
	if n < 0 || k < 0 || k > n || p < 0 || p > 1 {
		log.Panic("BinomialCDF: invalid inputs (n, k, p must be in valid ranges)")
	}
	sum := 0.0
	for i := 0; i <= k; i++ {
		sum += BinomialPMF(n, i, p)
	}
	return sum
}

// binomialCoefficient calculates "n choose k" using an iterative approach.
func binomialCoefficient(n, k int) float64 {
	if k > n-k {
		k = n - k
	}
	result := 1
	for i := 0; i < k; i++ {
		result *= (n - i)
		result /= (i + 1)
	}
	return float64(result)
}
