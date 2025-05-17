package probability

// AdditionRule calculates P(A ∪ B) = P(A) + P(B) - P(A ∩ B)
func AdditionRule(probA, probB, probAandB float64) float64 {
	return probA + probB - probAandB
}

// MultiplicationRule calculates P(A ∩ B) = P(A) * P(B) for independent events
func MultiplicationRule(probA, probB float64) float64 {
	return probA * probB
}

// ComplementRule calculates P(A') = 1 - P(A)
func ComplementRule(probA float64) float64 {
	return 1 - probA
}

// Intersection returns P(A ∩ B)
func Intersection(probA, probB float64) float64 {
	return probA * probB
}

// Union returns P(A ∪ B), assuming overlap is passed in
func Union(probA, probB, overlap float64) float64 {
	return probA + probB - overlap
}
