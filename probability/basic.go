package probability

import "log"

// AdditionRule calculates P(A ∪ B) = P(A) + P(B) - P(A ∩ B)
func AdditionRule(probA, probB, probAandB float64) float64 {
	if probA < 0 || probA > 1 || probB < 0 || probB > 1 || probAandB < 0 || probAandB > 1 {
		log.Panic("AdditionRule: all probabilities must be between 0 and 1")
	}
	return probA + probB - probAandB
}

// MultiplicationIndependent calculates P(A ∩ B) = P(A) * P(B) for independent events
func MultiplicationIndependent(probA, probB float64) float64 {
	if probA < 0 || probA > 1 || probB < 0 || probB > 1 {
		log.Panic("MultiplicationIndependent: probabilities must be between 0 and 1")
	}
	return probA * probB
}

// MultiplicationDependent calculates P(A ∩ B) = P(A) * P(B|A) for dependent events
func MultiplicationDependent(probA, probBgivenA float64) float64 {
	if probA < 0 || probA > 1 || probBgivenA < 0 || probBgivenA > 1 {
		log.Panic("MultiplicationDependent: probabilities must be between 0 and 1")
	}
	return probA * probBgivenA
}

// ComplementRule calculates P(A') = 1 - P(A)
func ComplementRule(probA float64) float64 {
	if probA < 0 || probA > 1 {
		log.Panic("ComplementRule: probability must be between 0 and 1")
	}
	return 1 - probA
}

// Intersection calculates P(A ∩ B). Assumes independence unless noted.
func Intersection(probA, probB float64) float64 {
	if probA < 0 || probA > 1 || probB < 0 || probB > 1 {
		log.Panic("Intersection: probabilities must be between 0 and 1")
	}
	return probA * probB
}

// Union calculates P(A ∪ B) = P(A) + P(B) - P(A ∩ B)
func Union(probA, probB, overlap float64) float64 {
	if probA < 0 || probA > 1 || probB < 0 || probB > 1 || overlap < 0 || overlap > 1 {
		log.Panic("Union: all probabilities must be between 0 and 1")
	}
	return probA + probB - overlap
}
