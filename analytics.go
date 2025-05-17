package statistical

import (
	"github.com/cyber-mountain-man/statistical-go/stat"
	"github.com/cyber-mountain-man/statistical-go/probability"
	"github.com/cyber-mountain-man/statistical-go/montecarlo"
)

// --- Descriptive Statistics ----
var Mean = stat.Mean
var Variance = stat.Variance
var StdDev = stat.StdDev

// --- Probability ---
var ConditionalProbability = probability.ConditionalProbability

// --- Monte Carlo ---
var EstimatePi = montecarlo.EstimatePi
var EstimatePiParallel = montecarlo.EstimatePiParallel