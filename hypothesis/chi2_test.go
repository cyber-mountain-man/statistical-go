package hypothesis

import (
	"math"
	"testing"
)

func TestChiSquareGoodnessOfFit(t *testing.T) {
	observed := []float64{16, 18, 16, 14, 12, 12}
	expected := []float64{16, 16, 16, 16, 16, 8}
	expectedChi2 := 3.5

	result := ChiSquareGoodnessOfFit(observed, expected)
	if math.Abs(result.Statistic-expectedChi2) > 0.1 {
		t.Errorf("ChiSquareGoodnessOfFit().Statistic = %.4f; want %.4f", result.Statistic, expectedChi2)
	}
	if result.PValue <= 0 || result.PValue > 1 {
		t.Errorf("ChiSquareGoodnessOfFit().PValue = %.4f; expected valid p-value in (0, 1]", result.PValue)
	}
}

func TestChiSquareGoodnessOfFit_InvalidLength(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic due to mismatched lengths")
		}
	}()
	ChiSquareGoodnessOfFit([]float64{10, 20}, []float64{15})
}

func TestChiSquareGoodnessOfFit_ZeroExpected(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic due to zero expected frequency")
		}
	}()
	ChiSquareGoodnessOfFit([]float64{10}, []float64{0})
}

func TestChiSquareTestOfIndependence(t *testing.T) {
	table := [][]float64{
		{90, 60, 104, 95},
		{30, 50, 51, 20},
		{30, 40, 45, 35},
	}
	expectedChi2 := 24.57

	result := ChiSquareTestOfIndependence(table)
	if math.Abs(result.Statistic-expectedChi2) > 0.1 {
		t.Errorf("ChiSquareTestOfIndependence().Statistic = %.4f; want %.4f", result.Statistic, expectedChi2)
	}
	if result.PValue <= 0 || result.PValue > 1 {
		t.Errorf("ChiSquareTestOfIndependence().PValue = %.4f; expected valid p-value in (0, 1]", result.PValue)
	}
}

func TestChiSquareTestOfIndependence_InvalidRowSize(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic due to inconsistent row sizes")
		}
	}()
	ChiSquareTestOfIndependence([][]float64{
		{10, 20},
		{5},
	})
}

func TestChiSquareTestOfIndependence_NegativeValue(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic due to negative value")
		}
	}()
	ChiSquareTestOfIndependence([][]float64{
		{10, -5},
		{15, 20},
	})
}

func TestChiSquareTestOfIndependence_InvalidTableDimensions(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for malformed contingency table, but no panic occurred")
		}
	}()
	ChiSquareTestOfIndependence([][]float64{
		{10, 20},
		{5},
	})
}

func TestChiSquareTestOfIndependence_OneRow(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic due to insufficient table rows")
		}
	}()
	table := [][]float64{{10, 20, 30}}
	ChiSquareTestOfIndependence(table)
}

func TestChiSquareTestOfIndependence_OneColumn(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic due to insufficient table columns")
		}
	}()
	table := [][]float64{
		{10},
		{20},
	}
	ChiSquareTestOfIndependence(table)
}

func TestChiSquareTestOfIndependence_ZeroMarginals(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic due to zero marginal totals")
		}
	}()
	table := [][]float64{
		{0, 0},
		{0, 0},
	}
	ChiSquareTestOfIndependence(table)
}

func TestChiSquareTestOfIndependence_ZeroExpectedCell(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected a panic due to zero row total, but function did not panic")
		}
	}()

	table := [][]float64{
		{0, 0},
		{0, 10},
	}

	_ = ChiSquareTestOfIndependence(table)
}

func TestChiSquareTestOfIndependence_ZeroColTotal(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic due to zero column total, but no panic occurred")
		}
	}()

	table := [][]float64{
		{0, 1},
		{0, 1},
	}

	_ = ChiSquareTestOfIndependence(table)
}

func TestGammaReflectionBranch(t *testing.T) {
	// z < 0.5 triggers the reflection formula path
	val := gamma(0.4)
	if math.IsNaN(val) || val <= 0 {
		t.Errorf("Expected valid gamma value for z < 0.5, got %v", val)
	}
}
