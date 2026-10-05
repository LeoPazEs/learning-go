package main

import "fmt"

// This is a naked return! You just say return and he solves it by the x,y definition at the beginning of the func
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
