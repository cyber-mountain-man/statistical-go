package main

import (
	"fmt"

	"github.com/cyber-mountain-man/statistical-go"
)

func main() {
	fmt.Println("=== Hypothesis Testing Examples ===")

	// One-Sample Z-Test
	zStat, zP := statistical.OneSampleZTest(102, 100, 15, 30)
	fmt.Printf("Z-Test -> Z = %.4f, P = %.4f\n", zStat, zP)

	// One-Sample T-Test
	tStat, tP := statistical.OneSampleTTest(98, 100, 10, 25)
	fmt.Printf("T-Test -> T = %.4f, P = %.4f\n", tStat, tP)

	// Paired T-Test
	before := []float64{10, 12, 9, 11, 13}
	after := []float64{12, 14, 10, 13, 15}
	ptStat, ptP := statistical.PairedTTest(before, after)
	fmt.Printf("Paired T-Test -> T = %.4f, P = %.4f\n", ptStat, ptP)

	// Two-Sample Welch T-Test
	t2Stat, t2P := statistical.TwoSampleTTestWelch(100, 105, 15, 12, 30, 35)
	fmt.Printf("Two-Sample T-Test (Welch) -> T = %.4f, P = %.4f\n", t2Stat, t2P)

	// Chi-Square Goodness of Fit
	observed := []float64{18, 22, 20}
	expected := []float64{20, 20, 20}
	chiStat, chiP, _ := statistical.ChiSquareGoodnessOfFit(observed, expected)
	fmt.Printf("Chi-Square GoF -> χ² = %.4f, P = %.4f\n", chiStat, chiP)

	// One-Way ANOVA
	groupA := []float64{12, 14, 13, 15}
	groupB := []float64{22, 21, 23, 24}
	groupC := []float64{32, 33, 31, 30}
	anStat, anP, _ := statistical.OneWayANOVA([][]float64{groupA, groupB, groupC})
	fmt.Printf("One-Way ANOVA -> F = %.4f, P = %.4f\n", anStat, anP)
}
