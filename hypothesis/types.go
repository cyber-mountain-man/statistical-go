package hypothesis

// TestResult holds the outcome of a hypothesis test
type TestResult struct {
	Statistic float64
	PValue    float64
	Err       error
}
