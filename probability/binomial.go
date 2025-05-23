package probability

import (
	"math"
)

// BinomialPMF returns the probability of getting exactly k successes in n trials
// with success probability p (P(X = k)).
// Returns 0.0 for invalid inputs.
func BinomialPMF(n, k int, p float64) float64 {
	if n < 0 || k < 0 || k > n || p < 0 || p > 1 {
		return 0.0
	}
	coef := float64(binomialCoefficient(n, k))
	return coef * math.Pow(p, float64(k)) * math.Pow(1-p, float64(n-k))
}

// BinomialCDF returns the cumulative probability of getting at most k successes (P(X ≤ k)).
// Returns 0.0 for invalid inputs.
func BinomialCDF(n, k int, p float64) float64 {
	if n < 0 || k < 0 || k > n || p < 0 || p > 1 {
		return 0.0
	}
	sum := 0.0
	for i := 0; i <= k; i++ {
		sum += BinomialPMF(n, i, p)
	}
	return sum
}

// BinomialQuantile returns the smallest integer k such that P(X ≤ k) ≥ targetP
// for a binomial distribution with parameters n and p.
// **Returns n** when the inputs are invalid or the target is unreachable
// so that callers/tests can recognise “fallback” with the same signature.
func BinomialQuantile(targetP float64, n int, p float64) int {
	// ── invalid-input guard ───────────────────────────────────────────────
	if targetP < 0 || targetP > 1 || n < 0 || p < 0 || p > 1 {
		return n            // ← restore previous behaviour
	}

	cumulative := 0.0
	for k := 0; k <= n; k++ {
		cumulative += BinomialPMF(n, k, p)
		if cumulative >= targetP {
			return k
		}
	}
	return n                // ← fallback remains ‘n’, not 0
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
