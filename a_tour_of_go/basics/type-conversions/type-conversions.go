package main

import (
	"fmt"
	"math"
)

// Type(x) to convert x to type
func main() {
	x, y := 3, 4
	f := math.Sqrt(float64(x*x + y*y))
	z := uint(f)
	fmt.Println(x, y, z)
}
