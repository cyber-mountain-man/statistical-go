package hypothesis

import (
	"math"
)

// OneSampleTTest returns a t-test statistic and p-value for a one-sample t-test.
func OneSampleTTest(sampleMean, populationMean, sampleStdDev float64, n int) TestResult {
	if sampleStdDev <= 0 || n <= 1 {
		panic("OneSampleTTest: sample standard deviation must be > 0 and sample size > 1")
	}
	standardError := sampleStdDev / math.Sqrt(float64(n))
	t := (sampleMean - populationMean) / standardError
	df := float64(n - 1)
	p := 2 * (1 - studentTCDF(math.Abs(t), df))
	return TestResult{Statistic: t, PValue: p}
}

// TwoSampleTTestWelch returns a t-test statistic and p-value for Welch's t-test.
func TwoSampleTTestWelch(mean1, mean2, stdDev1, stdDev2 float64, n1, n2 int) TestResult {
	if stdDev1 <= 0 || stdDev2 <= 0 || n1 <= 1 || n2 <= 1 {
		panic("TwoSampleTTestWelch: standard deviations must be > 0 and sample sizes > 1")
	}
	se1 := (stdDev1 * stdDev1) / float64(n1)
	se2 := (stdDev2 * stdDev2) / float64(n2)
	t := (mean1 - mean2) / math.Sqrt(se1+se2)

	// Welch-Satterthwaite approximation
	df := math.Pow(se1+se2, 2) / ((math.Pow(se1, 2) / float64(n1-1)) + (math.Pow(se2, 2) / float64(n2-1)))
	p := 2 * (1 - studentTCDF(math.Abs(t), df))

	return TestResult{Statistic: t, PValue: p}
}

// PairedTTest returns a t-test statistic and p-value for a paired sample t-test.
func PairedTTest(x, y []float64) TestResult {
	if len(x) != len(y) {
		panic("PairedTTest: slices must have the same length")
	}
	n := len(x)
	if n < 2 {
		panic("PairedTTest: need at least 2 paired observations")
	}

	var sumDiff, sumDiffSq float64
	for i := 0; i < n; i++ {
		d := x[i] - y[i]
		sumDiff += d
		sumDiffSq += d * d
	}

	meanDiff := sumDiff / float64(n)
	variance := (sumDiffSq - float64(n)*meanDiff*meanDiff) / float64(n-1)
	stdDev := math.Sqrt(variance)
	standardError := stdDev / math.Sqrt(float64(n))
	t := meanDiff / standardError
	df := float64(n - 1)
	p := 2 * (1 - studentTCDF(math.Abs(t), df))

	return TestResult{Statistic: t, PValue: p}
}

// studentTCDF approximates the CDF for the Student’s t-distribution (for two-tailed tests).
// Currently uses a normal distribution approximation as a placeholder.
// In the future, this can be replaced with a more accurate approximation
// (e.g., using a series expansion or continued fractions) without needing external dependencies like gonum.
func studentTCDF(t, df float64) float64 {
	_ = df // Retain df for future implementation; suppress unused parameter warning.

	// TODO: Replace with a proper Student's t-distribution CDF approximation.
	// For now, use the standard normal CDF as an approximation.
	return standardNormalCDF(t)
}

// standardNormalCDF computes the cumulative distribution function for the standard normal distribution.
func standardNormalCDF(x float64) float64 {
	return 0.5 * (1 + math.Erf(x/math.Sqrt2))
}
