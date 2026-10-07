package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] // This is a slice and a slice is dynamically-sized
	fmt.Println(s)
	ps := &s
	fmt.Println((*ps)[2]) // Get a index of a slice but a slice is already like a pointer
}
