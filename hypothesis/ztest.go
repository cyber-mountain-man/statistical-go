package hypothesis

import (
	"math"
)

// OneSampleZTest performs a one-sample z-test and returns a TestResult.
func OneSampleZTest(sampleMean, populationMean, populationStdDev float64, n int) TestResult {
	if populationStdDev <= 0 {
		panic("OneSampleZTest: population standard deviation must be > 0")
	}
	if n <= 0 {
		panic("OneSampleZTest: sample size must be > 0")
	}
	standardError := populationStdDev / math.Sqrt(float64(n))
	z := (sampleMean - populationMean) / standardError
	p := 2 * (1 - normalCDF(math.Abs(z), 0, 1)) // two-tailed p-value
	return TestResult{Statistic: z, PValue: p}
}

// TwoSampleZTest performs a two-sample z-test and returns a TestResult.
func TwoSampleZTest(mean1, mean2, stdDev1, stdDev2 float64, n1, n2 int) TestResult {
	if stdDev1 <= 0 || stdDev2 <= 0 {
		panic("TwoSampleZTest: both population standard deviations must be > 0")
	}
	if n1 <= 0 || n2 <= 0 {
		panic("TwoSampleZTest: sample sizes must be > 0")
	}
	se := math.Sqrt((stdDev1*stdDev1)/float64(n1) + (stdDev2*stdDev2)/float64(n2))
	z := (mean1 - mean2) / se
	p := 2 * (1 - normalCDF(math.Abs(z), 0, 1)) // two-tailed p-value
	return TestResult{Statistic: z, PValue: p}
}

// --- helper ---

// normalCDF returns the cumulative distribution function for the standard normal distribution.
func normalCDF(x, mean, stddev float64) float64 {
	return 0.5 * (1 + math.Erf((x-mean)/(stddev*math.Sqrt2)))
}
