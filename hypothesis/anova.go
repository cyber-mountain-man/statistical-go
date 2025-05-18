package hypothesis

import (
	"math"
)

// OneWayANOVA calculates the F-statistic and p-value for one-way ANOVA.
// Each inner slice in `groups` represents the data for one group.
func OneWayANOVA(groups [][]float64) TestResult {
	if len(groups) < 2 {
		panic("OneWayANOVA: requires at least two groups")
	}

	totalCount := 0
	totalSum := 0.0
	for _, group := range groups {
		for _, val := range group {
			totalSum += val
			totalCount++
		}
	}
	grandMean := totalSum / float64(totalCount)

	// Between-group variability (SSB)
	ssb := 0.0
	for _, group := range groups {
		groupMean := mean(group)
		n := float64(len(group))
		ssb += n * math.Pow(groupMean-grandMean, 2)
	}

	// Within-group variability (SSW)
	ssw := 0.0
	for _, group := range groups {
		groupMean := mean(group)
		for _, val := range group {
			ssw += math.Pow(val-groupMean, 2)
		}
	}

	dfBetween := float64(len(groups) - 1)
	dfWithin := float64(totalCount - len(groups))

	msb := ssb / dfBetween
	msw := ssw / dfWithin

	if msw == 0 {
		return TestResult{Statistic: math.NaN(), PValue: math.NaN()}
	}

	f := msb / msw
	p := 1 - fDistributionCDF(f, dfBetween, dfWithin)

	return TestResult{Statistic: f, PValue: p}
}

func mean(data []float64) float64 {
	sum := 0.0
	for _, val := range data {
		sum += val
	}
	return sum / float64(len(data))
}

// fDistributionCDF is a placeholder using a normal approximation for demonstration.
// In production, use a proper F-distribution CDF (e.g., from gonum/stat/distuv).
func fDistributionCDF(f, dfn, dfd float64) float64 {
	// Approximate F with standard normal CDF for large df
	if dfn <= 0 || dfd <= 0 {
		return 0
	}
	approxZ := (f - 1) / (2 / math.Sqrt(dfd))
	return standardNormalCDF(approxZ)
}
