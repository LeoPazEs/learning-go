package main

import "fmt"

func main() {
	fmt.Println("counting")

	// Goes to stack! LIFO
	for i := 0; i < 10; i++ {
		// Defer is evaluated in the defer line! Not after the return of the main function.
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

// This returns 2
// func c() (i int) {
//     defer func() { i++ }()
//     return 1
// }
