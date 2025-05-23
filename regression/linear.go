package regression

import "math"

// SimpleLinearRegression calculates the slope (m) and intercept (b)
// for the equation: y = m*x + b
func SimpleLinearRegression(x, y []float64) (slope float64, intercept float64) {
	if len(x) != len(y) || len(x) == 0 {
		return 0, 0 // Return zero if input is invalid
	}

	slope = CalculateSlope(x, y)
	if slope == 0 && isZeroDenominator(x) {
		return 0, 0 // Ensure intercept is also zero if slope is invalid due to constant x-values
	}

	intercept = CalculateIntercept(x, y, slope)
	return
}

// CalculateSlope computes the slope of the regression line.
func CalculateSlope(x, y []float64) float64 {
	if len(x) != len(y) || len(x) == 0 {
		return 0
	}

	n := float64(len(x))
	var sumX, sumY, sumXY, sumX2 float64

	for i := 0; i < len(x); i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumX2 += x[i] * x[i]
	}

	denominator := n*sumX2 - sumX*sumX
	if denominator == 0 {
		return 0 // Avoid division by zero
	}

	return (n*sumXY - sumX*sumY) / denominator
}

// CalculateIntercept computes the y-intercept given x, y, and the slope.
func CalculateIntercept(x, y []float64, slope float64) float64 {
	if len(x) != len(y) || len(x) == 0 {
		return 0
	}

	var sumX, sumY float64
	for i := 0; i < len(x); i++ {
		sumX += x[i]
		sumY += y[i]
	}

	n := float64(len(x))
	return (sumY - slope*sumX) / n
}

// PredictY returns the predicted y-value given x, using slope and intercept.
func PredictY(x float64, slope float64, intercept float64) float64 {
	return slope*x + intercept
}

// isZeroDenominator checks whether the denominator of the slope formula is zero.
// This usually happens when all x values are identical.
func isZeroDenominator(x []float64) bool {
	n := float64(len(x))
	var sumX, sumX2 float64
	for _, xi := range x {
		sumX += xi
		sumX2 += xi * xi
	}
	return n*sumX2-sumX*sumX == 0
}

// MeanSquaredError calculates the average squared difference between actual and predicted y-values.
func MeanSquaredError(yTrue, yPred []float64) float64 {
	if len(yTrue) != len(yPred) || len(yTrue) == 0 {
		return 0
	}

	var sum float64
	for i := range yTrue {
		diff := yTrue[i] - yPred[i]
		sum += diff * diff
	}

	return sum / float64(len(yTrue))
}

// RootMeanSquaredError computes the RMSE between true and predicted values.
func RootMeanSquaredError(yTrue, yPred []float64) float64 {
	if len(yTrue) != len(yPred) || len(yTrue) == 0 {
		return 0
	}
	mse := MeanSquaredError(yTrue, yPred)
	return math.Sqrt(mse)
}

// EvaluateLinearRegression runs simple linear regression and returns the MSE.
func EvaluateLinearRegression(x, y []float64) float64 {
	slope, intercept := SimpleLinearRegression(x, y)

	// Predict y values
	yPred := make([]float64, len(x))
	for i := range x {
		yPred[i] = PredictY(x[i], slope, intercept)
	}

	return MeanSquaredError(y, yPred)
}

// RSquared calculates the coefficient of determination (R²)
// to evaluate the fit of the regression model.
func RSquared(x, y []float64, slope, intercept float64) float64 {
	if len(x) != len(y) || len(x) == 0 {
		return 0
	}

	var ssTotal, ssResidual float64
	var meanY float64
	for _, val := range y {
		meanY += val
	}
	meanY /= float64(len(y))

	for i := range x {
		predicted := PredictY(x[i], slope, intercept)
		ssResidual += (y[i] - predicted) * (y[i] - predicted)
		ssTotal += (y[i] - meanY) * (y[i] - meanY)
	}

	if ssTotal == 0 {
		return 0
	}

	return 1 - (ssResidual / ssTotal)
}
