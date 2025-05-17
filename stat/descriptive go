package stat

import "math"

func Mean(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	sum := 0.0
	for _, v := range data {
		sum += v
	}
	return sum / float64(len(data))
}

func Variance(data [] float64) float64 {
	if len(data) < 2 {
		return 0
	}
	mean := Mean(data)
	sum := 0.0
	for _, v := range data {
		d := v - mean
		sum += d * d
	}
	return sum / float64(len(data)-1)
}

func StdDev(data []float64) float64 {
	return math.Sqrt(Variance(data))
}