package probability

// UniformPDF returns the probability density for a continuous uniform distribution
// between a and b evaluated at x.
func UniformPDF(x, a, b float64) float64 {
	if a >= b {
		panic("Invalid bounds: a must be less than b")
	}
	if x < a || x > b {
		return 0.0
	}
	return 1.0 / (b - a)
}

// UniformCDF returns the cumulative distribution function value
// for a continuous uniform distribution between a and b at x.
func UniformCDF(x, a, b float64) float64 {
	if a >= b {
		panic("Invalid bounds: a must be less than b")
	}
	if x < a {
		return 0.0
	} else if x > b {
		return 1.0
	}
	return (x - a) / (b - a)
}
