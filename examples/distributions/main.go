package main

import (
	"fmt"
	"github.com/cyber-mountain-man/statistical-go/probability"
)

func main() {
	fmt.Println("=== Distributions Demo ===")

	// Normal Distribution
	fmt.Printf("NormalPDF(0, 0, 1): %.5f\n", probability.NormalPDF(0, 0, 1))
	fmt.Printf("NormalCDF(1.96, 0, 1): %.5f\n", probability.NormalCDF(1.96, 0, 1))

	// Binomial Distribution
	fmt.Printf("BinomialPMF(n=10, k=3, p=0.5): %.5f\n", probability.BinomialPMF(10, 3, 0.5))
	fmt.Printf("BinomialCDF(n=10, k=3, p=0.5): %.5f\n", probability.BinomialCDF(10, 3, 0.5))

	// Uniform Distribution
	fmt.Printf("UniformPDF(x=3, a=0, b=5): %.5f\n", probability.UniformPDF(3, 0, 5))
	fmt.Printf("UniformCDF(x=3, a=0, b=5): %.5f\n", probability.UniformCDF(3, 0, 5))

	// Poisson Distribution
	fmt.Printf("PoissonPMF(k=2, λ=3): %.5f\n", probability.PoissonPMF(2, 3))
	fmt.Printf("PoissonCDF(k=2, λ=3): %.5f\n", probability.PoissonCDF(2, 3))

	// Exponential Distribution
	fmt.Printf("ExponentialPDF(x=1, λ=2): %.5f\n", probability.ExponentialPDF(1, 2))
	fmt.Printf("ExponentialCDF(x=1, λ=2): %.5f\n", probability.ExponentialCDF(1, 2))
}
