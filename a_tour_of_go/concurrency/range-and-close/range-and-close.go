package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		if i == 12 {
			time.Sleep(2 * time.Second)
		}
		c <- x
		x, y = y, x+y
	}
	// Only the sender should close a channel, never the receiver.
	// Sending on a closed channel will cause a panic.
	close(c)
}

func main() {
	c := make(chan int)
	go fibonacci(50, c)
	// Goes until closed
	for i := range c {
		fmt.Println(i)
	}
}
