package probability

import (
	"math"
	"testing"
)

func TestUniformPDF(t *testing.T) {
	tests := []struct {
		x, a, b     float64
		expected    float64
		shouldPanic bool
	}{
		{5, 0, 10, 0.1, false},
		{0, 0, 10, 0.1, false},
		{10, 0, 10, 0.1, false},
		{-1, 0, 10, 0, false},
		{11, 0, 10, 0, false},
		{5, 10, 0, 0, true},
	}

	for _, tt := range tests {
		func() {
			defer func() {
				if r := recover(); (r != nil) != tt.shouldPanic {
					t.Errorf("UniformPDF(%v, %v, %v) panic = %v; want panic = %v", tt.x, tt.a, tt.b, r != nil, tt.shouldPanic)
				}
			}()
			got := UniformPDF(tt.x, tt.a, tt.b)
			if math.Abs(got-tt.expected) > 1e-6 && !tt.shouldPanic {
				t.Errorf("UniformPDF(%v, %v, %v) = %v; want %v", tt.x, tt.a, tt.b, got, tt.expected)
			}
		}()
	}
}

func TestUniformCDF(t *testing.T) {
	tests := []struct {
		x, a, b     float64
		expected    float64
		shouldPanic bool
	}{
		{5, 0, 10, 0.5, false},
		{0, 0, 10, 0.0, false},
		{10, 0, 10, 1.0, false},
		{-1, 0, 10, 0.0, false},
		{11, 0, 10, 1.0, false},
		{5, 10, 0, 0, true},
	}

	for _, tt := range tests {
		func() {
			defer func() {
				if r := recover(); (r != nil) != tt.shouldPanic {
					t.Errorf("UniformCDF(%v, %v, %v) panic = %v; want panic = %v", tt.x, tt.a, tt.b, r != nil, tt.shouldPanic)
				}
			}()
			got := UniformCDF(tt.x, tt.a, tt.b)
			if math.Abs(got-tt.expected) > 1e-6 && !tt.shouldPanic {
				t.Errorf("UniformCDF(%v, %v, %v) = %v; want %v", tt.x, tt.a, tt.b, got, tt.expected)
			}
		}()
	}
}
