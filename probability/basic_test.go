package probability

import "testing"

func floatEquals(a, b float64) bool {
	const epsilon = 1e-9
	return (a-b) < epsilon && (b-a) < epsilon
}

func TestAdditionRule(t *testing.T) {
	got := AdditionRule(0.6, 0.3, 0.2)
	want := 0.7
	if !floatEquals(got, want) {
		t.Errorf("AdditionRule(0.6, 0.3, 0.2) = %v; want %v", got, want)
	}
}

func TestMultiplicationRule(t *testing.T) {
	got := MultiplicationRule(0.5, 0.4)
	want := 0.2
	if !floatEquals(got, want) {
		t.Errorf("MultiplicationRule(0.5, 0.4) = %v; want %v", got, want)
	}
}

func TestComplementRule(t *testing.T) {
	got := ComplementRule(0.8)
	want := 0.2
	if !floatEquals(got, want) {
		t.Errorf("ComplementRule(0.8) = %v; want %v", got, want)
	}
}

func TestIntersection(t *testing.T) {
	got := Intersection(0.7, 0.2)
	want := 0.14
	if !floatEquals(got, want) {
		t.Errorf("Intersection(0.7, 0.2) = %v; want %v", got, want)
	}
}

func TestUnion(t *testing.T) {
	got := Union(0.5, 0.4, 0.1)
	want := 0.8
	if !floatEquals(got, want) {
		t.Errorf("Union(0.5, 0.4, 0.1) = %v; want %v", got, want)
	}
}
