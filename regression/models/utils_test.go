package models

import "math"

// floatsAlmostEqual checks if two floats are approximately equal within epsilon.
func floatsAlmostEqual(a, b float64, epsilon float64) bool {
	return math.Abs(a-b) <= epsilon
}

// slicesAlmostEqual checks if two slices of floats are approximately equal.
func slicesAlmostEqual(a, b []float64, epsilon float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !floatsAlmostEqual(a[i], b[i], epsilon) {
			return false
		}
	}
	return true
}
