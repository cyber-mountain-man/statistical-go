package probability

import (
	"testing"
)

func TestAdditionRule(t *testing.T) {
	tests := []struct {
		probA, probB, probAandB, want float64
	}{
		{0.6, 0.3, 0.2, 0.7},
		{0.5, 0.5, 0.0, 1.0},
		{1.0, 0.0, 0.0, 1.0},
		{0.0, 0.0, 0.0, 0.0},
	}
	for _, tt := range tests {
		got := AdditionRule(tt.probA, tt.probB, tt.probAandB)
		if !floatEquals(got, tt.want) {
			t.Errorf("AdditionRule(%v, %v, %v) = %v; want %v", tt.probA, tt.probB, tt.probAandB, got, tt.want)
		}
	}
}

func TestMultiplicationIndependent(t *testing.T) {
	tests := []struct {
		probA, probB, want float64
	}{
		{0.5, 0.4, 0.2},
		{1.0, 1.0, 1.0},
		{0.0, 1.0, 0.0},
		{0.3, 0.3, 0.09},
	}
	for _, tt := range tests {
		got := MultiplicationIndependent(tt.probA, tt.probB)
		if !floatEquals(got, tt.want) {
			t.Errorf("MultiplicationIndependent(%v, %v) = %v; want %v", tt.probA, tt.probB, got, tt.want)
		}
	}
}

func TestComplementRule(t *testing.T) {
	tests := []struct {
		probA, want float64
	}{
		{0.8, 0.2},
		{0.0, 1.0},
		{1.0, 0.0},
		{0.33, 0.67},
	}
	for _, tt := range tests {
		got := ComplementRule(tt.probA)
		if !floatEquals(got, tt.want) {
			t.Errorf("ComplementRule(%v) = %v; want %v", tt.probA, got, tt.want)
		}
	}
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		probA, probB, want float64
	}{
		{0.7, 0.2, 0.14},
		{0.0, 0.5, 0.0},
		{1.0, 1.0, 1.0},
		{0.25, 0.4, 0.1},
	}
	for _, tt := range tests {
		got := Intersection(tt.probA, tt.probB)
		if !floatEquals(got, tt.want) {
			t.Errorf("Intersection(%v, %v) = %v; want %v", tt.probA, tt.probB, got, tt.want)
		}
	}
}

func TestUnion(t *testing.T) {
	tests := []struct {
		probA, probB, overlap, want float64
	}{
		{0.5, 0.4, 0.1, 0.8},
		{0.7, 0.3, 0.0, 1.0},
		{0.6, 0.5, 0.2, 0.9},
		{0.2, 0.2, 0.05, 0.35},
	}
	for _, tt := range tests {
		got := Union(tt.probA, tt.probB, tt.overlap)
		if !floatEquals(got, tt.want) {
			t.Errorf("Union(%v, %v, %v) = %v; want %v", tt.probA, tt.probB, tt.overlap, got, tt.want)
		}
	}
}

func TestMultiplicationDependent(t *testing.T) {
	tests := []struct {
		probA, probBgivenA, want float64
	}{
		{0.5, 0.6, 0.3},
		{1.0, 0.4, 0.4},
		{0.0, 0.5, 0.0},
		{0.75, 0.8, 0.6},
	}
	for _, tt := range tests {
		got := MultiplicationDependent(tt.probA, tt.probBgivenA)
		if !floatEquals(got, tt.want) {
			t.Errorf("MultiplicationDependent(%v, %v) = %v; want %v", tt.probA, tt.probBgivenA, got, tt.want)
		}
	}
}

func TestComplementRulePanic(t *testing.T) {
	tests := []float64{-0.1, 1.1}
	for _, prob := range tests {
		func() {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("ComplementRule(%v) did not panic", prob)
				}
			}()
			_ = ComplementRule(prob)
		}()
	}
}

func TestIntersectionPanic(t *testing.T) {
	tests := [][2]float64{
		{-0.1, 0.5},
		{0.5, 1.5},
	}
	for _, pair := range tests {
		func() {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Intersection(%v, %v) did not panic", pair[0], pair[1])
				}
			}()
			_ = Intersection(pair[0], pair[1])
		}()
	}
}

func TestUnionPanic(t *testing.T) {
	tests := [][3]float64{
		{0.5, 0.4, -0.1},
		{1.1, 0.4, 0.2},
		{0.3, 1.2, 0.1},
	}
	for _, triplet := range tests {
		func() {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Union(%v, %v, %v) did not panic", triplet[0], triplet[1], triplet[2])
				}
			}()
			_ = Union(triplet[0], triplet[1], triplet[2])
		}()
	}
}

func TestMultiplicationIndependentPanic(t *testing.T) {
	tests := [][2]float64{
		{-0.1, 0.5},
		{1.1, 0.5},
		{0.3, -0.2},
		{0.6, 1.5},
	}
	for _, pair := range tests {
		func() {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("MultiplicationIndependent(%v, %v) did not panic", pair[0], pair[1])
				}
			}()
			_ = MultiplicationIndependent(pair[0], pair[1])
		}()
	}
}

func TestMultiplicationDependentPanic(t *testing.T) {
	tests := [][2]float64{
		{-0.1, 0.5},
		{1.1, 0.5},
		{0.5, -0.2},
		{0.6, 1.8},
	}
	for _, pair := range tests {
		func() {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("MultiplicationDependent(%v, %v) did not panic", pair[0], pair[1])
				}
			}()
			_ = MultiplicationDependent(pair[0], pair[1])
		}()
	}
}

func TestAdditionRulePanic(t *testing.T) {
	tests := [][3]float64{
		{-0.1, 0.5, 0.1},
		{0.5, -0.2, 0.1},
		{0.5, 0.4, -0.3},
		{1.2, 0.1, 0.1},
	}
	for _, triple := range tests {
		func() {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("AdditionRule(%v, %v, %v) did not panic", triple[0], triple[1], triple[2])
				}
			}()
			_ = AdditionRule(triple[0], triple[1], triple[2])
		}()
	}
}
