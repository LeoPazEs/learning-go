package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	y := make([][]uint8, dy)
	for i := range y {
		x := make([]uint8, dx)
		for j := range x {
			x[j] = uint8(((j + i) / 2) - (j * i) + (j ^ i))
		}
		y[i] = x
	}
	return y
}

func main() {
	pic.Show(Pic)
}
