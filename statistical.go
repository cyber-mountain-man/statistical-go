package statistical

import (
	"github.com/cyber-mountain-man/statistical-go/hypothesis"
	"github.com/cyber-mountain-man/statistical-go/montecarlo"
	"github.com/cyber-mountain-man/statistical-go/probability"
	"github.com/cyber-mountain-man/statistical-go/regression"
	"github.com/cyber-mountain-man/statistical-go/stat"
	"github.com/cyber-mountain-man/statistical-go/regression/models"
)

//
// DESCRIPTIVE STATISTICS
//

func Covariance(x, y []float64) float64         { return stat.Covariance(x, y) }
func Kurtosis(data []float64) float64           { return stat.Kurtosis(data) }
func Max(data []float64) float64                { return stat.Max(data) }
func Mean(data []float64) float64               { return stat.Mean(data) }
func Median(data []float64) float64             { return stat.Median(data) }
func Min(data []float64) float64                { return stat.Min(data) }
func Mode(data []float64) float64               { return stat.Mode(data) }
func PearsonCorrelation(x, y []float64) float64 { return stat.PearsonCorrelation(x, y) }
func Quartiles(data []float64) (float64, float64, float64) {
	return stat.Quartiles(data)
}
func Range(data []float64) float64       { return stat.Range(data) }
func Skewness(data []float64) float64    { return stat.Skewness(data) }
func StdDev(data []float64) float64      { return stat.StdDev(data) }
func Variance(data []float64) float64    { return stat.Variance(data) }
func ZScore(x float64, data []float64) float64 { return stat.ZScore(x, data) }

//
// PROBABILITY RULES
//

func AdditionRule(probA, probB, probAandB float64) float64 {
	return probability.AdditionRule(probA, probB, probAandB)
}

func Complement(probA float64) float64 {
	return probability.ComplementRule(probA)
}

func ConditionalProbability(probAandB, probB float64) float64 {
	return probability.Conditional(probAandB, probB)
}

func MultiplicationRuleDependent(probA, probBgivenA float64) float64 {
	return probability.MultiplicationDependent(probA, probBgivenA)
}

func MultiplicationRuleIndependent(probA, probB float64) float64 {
	return probability.MultiplicationIndependent(probA, probB)
}

//
// MONTE CARLO SIMULATIONS
//

func EstimatePi(iterations int) float64 {
	return montecarlo.EstimatePi(iterations)
}

func EstimatePiParallel(total, workers int) float64 {
	return montecarlo.EstimatePiParallel(total, workers)
}

//
// PROBABILITY DISTRIBUTIONS
//

func BinomialCDF(k, n int, p float64) float64 {
	return probability.BinomialCDF(k, n, p)
}

func BinomialPMF(k, n int, p float64) float64 {
	return probability.BinomialPMF(k, n, p)
}

func BinomialQuantile(p float64, n int, prob float64) int {
	return probability.BinomialQuantile(p, n, prob)
}

func NormalCDF(x, mean, stddev float64) float64 {
	return probability.NormalCDF(x, mean, stddev)
}

func NormalPDF(x, mean, stddev float64) float64 {
	return probability.NormalPDF(x, mean, stddev)
}

func NormalInverseCDF(p, mean, stddev float64) float64 {
	return probability.NormalInverseCDF(p, mean, stddev)
}

//
// HYPOTHESIS TESTING
//

func ChiSquareGoodnessOfFit(observed, expected []float64) (float64, float64, error) {
	res := hypothesis.ChiSquareGoodnessOfFit(observed, expected)
	return res.Statistic, res.PValue, res.Err
}

func ChiSquareTestOfIndependence(observed [][]float64) (float64, float64, error) {
	res := hypothesis.ChiSquareTestOfIndependence(observed)
	return res.Statistic, res.PValue, res.Err
}

func OneSampleTTest(sampleMean, populationMean, sampleStdDev float64, n int) (float64, float64) {
	res := hypothesis.OneSampleTTest(sampleMean, populationMean, sampleStdDev, n)
	return res.Statistic, res.PValue
}

func OneSampleZTest(sampleMean, populationMean, stddev float64, n int) (float64, float64) {
	res := hypothesis.OneSampleZTest(sampleMean, populationMean, stddev, n)
	return res.Statistic, res.PValue
}

func OneWayANOVA(groups [][]float64) (float64, float64, error) {
	res := hypothesis.OneWayANOVA(groups)
	return res.Statistic, res.PValue, res.Err
}

func PairedTTest(before, after []float64) (float64, float64) {
	res := hypothesis.PairedTTest(before, after)
	return res.Statistic, res.PValue
}

func TwoSampleTTestWelch(mean1, mean2, stddev1, stddev2 float64, n1, n2 int) (float64, float64) {
	res := hypothesis.TwoSampleTTestWelch(mean1, mean2, stddev1, stddev2, n1, n2)
	return res.Statistic, res.PValue
}

func TwoSampleZTest(mean1, mean2, stddev1, stddev2 float64, n1, n2 int) (float64, float64) {
	res := hypothesis.TwoSampleZTest(mean1, mean2, stddev1, stddev2, n1, n2)
	return res.Statistic, res.PValue
}

//
// LINEAR REGRESSION
//

func EvaluateLinearRegression(x, y []float64) float64 {
	return regression.EvaluateLinearRegression(x, y)
}

func MSE(yTrue, yPred []float64) float64 {
	return regression.MeanSquaredError(yTrue, yPred)
}

func PredictY(x, slope, intercept float64) float64 {
	return regression.PredictY(x, slope, intercept)
}

func RMSE(yTrue, yPred []float64) float64 {
	return regression.RootMeanSquaredError(yTrue, yPred)
}

func RSquared(x, y []float64, slope, intercept float64) float64 {
	return regression.RSquared(x, y, slope, intercept)
}

func SimpleLinearRegression(x, y []float64) (float64, float64) {
	return regression.SimpleLinearRegression(x, y)
}

//
// REGULARIZATION REGRESSION (RIDGE & LASSO)
//

func RidgeRegression(X [][]float64, y []float64, lambda float64) ([]float64, error) {
	return models.RidgeRegression(X, y, lambda)
}

func LassoRegression(X [][]float64, y []float64, lambda float64, maxIter int) ([]float64, error) {
	return models.LassoRegression(X, y, lambda, maxIter)
}
