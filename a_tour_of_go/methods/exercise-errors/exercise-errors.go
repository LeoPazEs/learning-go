package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// The %f avoids the loop in the Sprint interface
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	minDelta := 1e-5
	z := 1.0
	iterations := 0
	for {
		iterations += 1
		z1 := z
		z -= (z*z - x) / (2 * z)
		if math.Abs(z1-z) <= minDelta {
			return z, nil
		}
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
