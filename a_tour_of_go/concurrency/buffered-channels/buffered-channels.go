package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 3 // DEADLOCK
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
