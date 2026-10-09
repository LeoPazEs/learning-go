package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prev, sum := 0, 1
	return func() int {
		prev, sum = sum, sum+prev
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 50; i++ {
		fmt.Println(f())
	}
}
