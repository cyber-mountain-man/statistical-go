package probability

import "log"

// Conditional calculates P(A | B) = P(A ∩ B) / P(B)
// It panics if P(B) is zero or inputs are out of valid range.
func Conditional(probAandB, probB float64) float64 {
	if probB <= 0 || probB > 1 {
		log.Panic("Conditional: P(B) must be > 0 and ≤ 1")
	}
	if probAandB < 0 || probAandB > 1 {
		log.Panic("Conditional: P(A ∩ B) must be between 0 and 1")
	}
	return probAandB / probB
}
