package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, int) {
	minDelta := 1e-5
	z := 1.0
	iterations := 0
	for {
		iterations += 1
		z1 := z
		z -= (z*z - x) / (2 * z)
		if math.Abs(z1-z) <= minDelta {
			return z, iterations
		}
	}
}

func main() {
	value := 1175928123049817142.0
	mine, iterations := Sqrt(value)
	maths := math.Sqrt(value)
	fmt.Printf("Mine: %g With %d Iterations, Math: %g, Diff: %g\n", mine, iterations, maths, math.Abs(mine-maths))
}
