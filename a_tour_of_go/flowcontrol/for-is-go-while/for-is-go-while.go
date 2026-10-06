package main

import "fmt"

func main() {
	sum := 1
	// No while :D
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
