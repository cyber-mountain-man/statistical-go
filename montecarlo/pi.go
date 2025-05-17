package montecarlo

import (
	"math/rand"
	"time"
)

func EstimatePi(iterations int) float64 {
	rand.Seed(time.Now().UnixNano())
	inside := 0
	for i := 0; i < iterations; i++ {
		x, y := rand.Float64(), rand.Float64()
		if x*x+y*y <= 1 {
			inside++
		}
	}
	return 4.0 * float64(inside) / float64(iterations)
}

func EstimatePiParallel(total int, workers int) float64 {
	if workers <= 0 {
		return EstimatePi(total)
	}

	jobs := total / workers
	results := make(chan int, workers)

	for w := 0; w < workers; w++ {
		go func() {
			count := 0
			for i := 0; i < jobs; i++ {
				x, y := rand.Float64(), rand.Float64()
				if x*x+y*y <= 1 {
					count++
				}
			}
			results <- count
		}()
	}

	totalInside := 0
	for i := 0; i < workers; i++ {
		totalInside += <-results
	}

	return 4.0 * float64(totalInside) / float64(total)
}
