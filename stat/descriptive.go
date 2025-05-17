package stat

import (
	"math"
	"sort"
)

// Mean calculates the average of a slice of float64 values.
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

// Variance calculates the sample variance of a slice of float64.
func Variance(data []float64) float64 {
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

// StdDev calculates the sample standard deviation.
func StdDev(data []float64) float64 {
	return math.Sqrt(Variance(data))
}

// Median returns the middle value in a sorted data set.
func Median(data []float64) float64 {
	n := len(data)
	if n == 0 {
		return 0
	}
	sorted := append([]float64{}, data...) // copy
	sort.Float64s(sorted)
	mid := n / 2
	if n%2 == 0 {
		return (sorted[mid-1] + sorted[mid]) / 2.0
	}
	return sorted[mid]
}

// Mode returns the most frequent value(s) in the dataset.
// If multiple modes exist, it returns the first one.
func Mode(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	freq := make(map[float64]int)
	for _, v := range data {
		freq[v]++
	}
	maxFreq := 0
	mode := data[0]
	for val, count := range freq {
		if count > maxFreq {
			maxFreq = count
			mode = val
		}
	}
	return mode
}

// Range returns the difference between the max and min values.
func Range(data []float64) float64 {
	min, max := Min(data), Max(data)
	return max - min
}

// Min returns the smallest value in the data set.
func Min(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	min := data[0]
	for _, v := range data {
		if v < min {
			min = v
		}
	}
	return min
}

// Max returns the largest value in the data set.
func Max(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	return max
}

func Quartiles(data []float64) (float64, float64, float64) {
	n := len(data)
	if n == 0 {
		return 0, 0, 0
	}

	sorted := append([]float64{}, data...) // copy to avoid mutating original
	sort.Float64s(sorted)

	q2 := Median(sorted)

	var q1, q3 float64
	mid := n / 2

	if n == 1 {
		return sorted[0], sorted[0], sorted[0]
	}

	if n == 2 {
		q1 = sorted[0]
		q2 = (sorted[0] + sorted[1]) / 2
		q3 = sorted[1]
		return q1, q2, q3
	}

	if n%2 == 0 {
		q1 = Median(sorted[:mid])
		q3 = Median(sorted[mid:])
	} else {
		q1 = Median(sorted[:mid])
		q3 = Median(sorted[mid+1:])
	}

	return q1, q2, q3
}

// ZScore calculates the z-score of a value given the dataset.
// It returns 0 if standard deviation is zero or data is empty.
func ZScore(x float64, data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	std := StdDev(data)
	if std == 0 {
		return 0
	}
	return (x - Mean(data)) / std
}

// Covariance returns the sample covariance between two datasets.
// Returns 0 if the slices are unequal or too small.
func Covariance(x, y []float64) float64 {
	n := len(x)
	if n != len(y) || n < 2 {
		return 0
	}
	meanX := Mean(x)
	meanY := Mean(y)
	sum := 0.0
	for i := 0; i < n; i++ {
		sum += (x[i] - meanX) * (y[i] - meanY)
	}
	return sum / float64(n-1)
}

// PearsonCorrelation returns the Pearson correlation coefficient between x and y.
// Returns 0 if input lengths mismatch, std dev is 0, or not enough data.
func PearsonCorrelation(x, y []float64) float64 {
	if len(x) != len(y) || len(x) < 2 {
		return 0
	}
	stdX := StdDev(x)
	stdY := StdDev(y)
	if stdX == 0 || stdY == 0 {
		return 0
	}
	return Covariance(x, y) / (stdX * stdY)
}

// Skewness calculates the sample skewness of a dataset.
// Returns 0 for slices smaller than 3.
func Skewness(data []float64) float64 {
	n := len(data)
	if n < 3 {
		return 0
	}
	mean := Mean(data)
	std := StdDev(data)
	if std == 0 {
		return 0
	}
	var sum float64
	for _, x := range data {
		sum += math.Pow((x-mean)/std, 3)
	}
	return float64(n) / float64((n-1)*(n-2)) * sum
}

// Kurtosis returns the excess kurtosis (compared to normal distribution).
// Returns 0 for slices smaller than 4.
func Kurtosis(data []float64) float64 {
	n := len(data)
	if n < 4 {
		return 0
	}
	mean := Mean(data)
	std := StdDev(data)
	if std == 0 {
		return 0
	}
	var sum float64
	for _, x := range data {
		sum += math.Pow((x-mean)/std, 4)
	}
	n1, n2, n3 := float64(n-1), float64(n-2), float64(n-3)
	term1 := (float64(n) * (float64(n) + 1)) / (n1 * n2 * n3)
	term2 := 3 * math.Pow(n1, 2) / (n2 * n3)
	return term1*sum - term2
}
