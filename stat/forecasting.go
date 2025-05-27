package stat

// ExponentialSmoothing applies single exponential smoothing to a time series.
// alpha is the smoothing factor (0 < alpha <= 1).
// Returns the smoothed value at the end of the series.
func ExponentialSmoothing(data []float64, alpha float64) float64 {
	if len(data) == 0 {
		panic("ExponentialSmoothing: input data cannot be empty")
	}
	if alpha <= 0 || alpha > 1 {
		panic("ExponentialSmoothing: alpha must be in (0, 1]")
	}

	smoothed := data[0]
	for i := 1; i < len(data); i++ {
		smoothed = alpha*data[i] + (1-alpha)*smoothed
	}
	return smoothed
}
