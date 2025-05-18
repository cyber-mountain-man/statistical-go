package probability

import "math"

func floatEquals(a, b float64) bool {
	const epsilon = 1e-9
	return math.Abs(a-b) < epsilon
}
