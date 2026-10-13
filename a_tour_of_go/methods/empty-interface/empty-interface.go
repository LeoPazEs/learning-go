package main

import "fmt"

func main() {
	var i interface{}
	describe(i)
	if i == nil {
		fmt.Println("i is nil")
	} else {
		fmt.Println("i is NOT nil")
	}

	// When you assign int to i, the tuple becomes (nil,*int).
	// Because the Type portion of the tuple is populated with *int, the interface as a whole is no longer nil,
	// even though the data pointer it holds is pointing to nil.
	i = 42
	describe(i)
	if i == nil {
		fmt.Println("i is nil")
	} else {
		fmt.Println("i is NOT nil")
	}
	i = "hello"
	describe(i)
	if i == nil {
		fmt.Println("i is nil")
	} else {
		fmt.Println("i is NOT nil")
	}
}

// An empty interface may hold values of any type. (Every type implements at least zero methods.)
// any is an alias for interface{}, and the two are completely equivalent.

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
