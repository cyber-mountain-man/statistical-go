package hypothesis

import (
	"math"
)

// ChiSquareGoodnessOfFit computes the Chi-Square statistic and p-value for a goodness-of-fit test.
// observed: slice of observed frequencies
// expected: slice of expected frequencies
func ChiSquareGoodnessOfFit(observed, expected []float64) TestResult {
	if len(observed) != len(expected) || len(observed) == 0 {
		panic("ChiSquareGoodnessOfFit: observed and expected slices must be of equal non-zero length")
	}

	var chi2 float64
	df := float64(len(observed) - 1)
	for i := range observed {
		if expected[i] <= 0 {
			panic("ChiSquareGoodnessOfFit: expected frequencies must be > 0")
		}
		diff := observed[i] - expected[i]
		chi2 += (diff * diff) / expected[i]
	}

	p := 1 - chiSquareCDF(chi2, df)
	return TestResult{Statistic: chi2, PValue: p}
}

// ChiSquareTestOfIndependence computes the Chi-Square statistic and p-value for a contingency table.
func ChiSquareTestOfIndependence(table [][]float64) TestResult {
	numRows := len(table)
	if numRows < 2 {
		panic("ChiSquareTestOfIndependence: table must have at least 2 rows")
	}
	numCols := len(table[0])
	if numCols < 2 {
		panic("ChiSquareTestOfIndependence: table must have at least 2 columns")
	}

	rowSums := make([]float64, numRows)
	colSums := make([]float64, numCols)
	var total float64

	for i := 0; i < numRows; i++ {
		if len(table[i]) != numCols {
			panic("ChiSquareTestOfIndependence: all rows must have equal number of columns")
		}
		for j := 0; j < numCols; j++ {
			val := table[i][j]
			if val < 0 {
				panic("ChiSquareTestOfIndependence: all values must be ≥ 0")
			}
			rowSums[i] += val
			colSums[j] += val
			total += val
		}
	}

	for i := range rowSums {
		if rowSums[i] == 0 {
			panic("ChiSquareTestOfIndependence: row total is zero")
		}
	}
	for j := range colSums {
		if colSums[j] == 0 {
			panic("ChiSquareTestOfIndependence: column total is zero")
		}
	}

	var chi2 float64
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			expected := (rowSums[i] * colSums[j]) / total
			diff := table[i][j] - expected
			chi2 += (diff * diff) / expected
		}
	}

	df := float64((numRows - 1) * (numCols - 1))
	p := 1 - chiSquareCDF(chi2, df)
	return TestResult{Statistic: chi2, PValue: p}
}

// chiSquareCDF approximates the cumulative distribution function for the Chi-Square distribution.
// This uses the regularized incomplete gamma function approximation for simplicity.
func chiSquareCDF(x, k float64) float64 {
	// Using a basic series expansion (Gamma lower incomplete approximation)
	return gammaLowerIncomplete(k/2, x/2) / gamma(k/2)
}

// gammaLowerIncomplete approximates the lower incomplete gamma function using a series expansion.
func gammaLowerIncomplete(s, x float64) float64 {
	const maxIter = 100
	const epsilon = 1e-8
	sum := 1.0 / s
	term := sum
	for n := 1; n < maxIter; n++ {
		term *= x / (s + float64(n))
		sum += term
		if term < epsilon {
			break
		}
	}
	return math.Exp(-x) * math.Pow(x, s) * sum
}

// gamma approximates the gamma function using the Lanczos approximation.
func gamma(z float64) float64 {
	// Lanczos coefficients for g=7, n=9
	p := []float64{
		0.99999999999980993, 676.5203681218851, -1259.1392167224028,
		771.32342877765313, -176.61502916214059, 12.507343278686905,
		-0.13857109526572012, 9.9843695780195716e-6, 1.5056327351493116e-7,
	}
	g := 7.0
	if z < 0.5 {
		// Reflection formula
		return math.Pi / (math.Sin(math.Pi*z) * gamma(1-z))
	}
	z -= 1
	x := p[0]
	for i := 1; i < len(p); i++ {
		x += p[i] / (z + float64(i))
	}
	t := z + g + 0.5
	return math.Sqrt(2*math.Pi) * math.Pow(t, z+0.5) * math.Exp(-t) * x
}
