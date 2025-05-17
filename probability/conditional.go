package probability

func Conditional(probAandB, probB float64) float64 {
	if probB == 0 {
		return 0
	}
	return probAandB / probB
}