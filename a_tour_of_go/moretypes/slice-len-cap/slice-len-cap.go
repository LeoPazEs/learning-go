package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	// Here CAP is 6 because cap only goes down when the first elements are dropped not the last
	s = s[:4]
	printSlice(s)

	// This does not break the code
	s = s[:6]
	printSlice(s)

	// This breaks the code at runtime
	// s = s[:8]
	// printSlice(s)

	// Drop its first two values.
	// Here CAP is 4 because cap only goes down when the first elements are dropped not the last
	s = s[2:]
	printSlice(s)

	// This breaks the code at runtime
	// s = s[:6]
	// printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
