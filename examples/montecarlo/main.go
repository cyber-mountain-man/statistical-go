package main

import (
	"fmt"
	"time"

	"github.com/cyber-mountain-man/statistical-go/montecarlo"
)

func main() {
	fmt.Println("=== Monte Carlo Simulation Demo ===")

	// Estimate Pi using 1 million iterations
	start := time.Now()
	pi := montecarlo.EstimatePi(1_000_000)
	fmt.Printf("Estimated Pi (single-threaded): %.6f\n", pi)
	fmt.Printf("Elapsed: %v\n\n", time.Since(start))

	// Estimate Pi in parallel with 4 workers
	startParallel := time.Now()
	piParallel := montecarlo.EstimatePiParallel(1_000_000, 4)
	fmt.Printf("Estimated Pi (parallel): %.6f\n", piParallel)
	fmt.Printf("Elapsed: %v\n", time.Since(startParallel))
}
