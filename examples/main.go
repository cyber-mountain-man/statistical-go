package main

import (
	"fmt"

	"github.com/cyber-mountain-man/statistical-go"
)

func main() {
	data := []float64{1.0, 2.0, 3.0, 4.0, 5.0}

	fmt.Println("Mean:", statistical.Mean(data))
	fmt.Println("Std Dev:", statistical.StdDev(data))
	fmt.Println("Variance:", statistical.Variance(data))

	fmt.Println("Estimate Pi:", statistical.EstimatePi(100000))
	fmt.Println("Estimate Pi Parallel:", statistical.EstimatePiParallel(100000, 4))

	fmt.Println("P(A|B):", statistical.ConditionalProbability(0.15, 0.3))
}
