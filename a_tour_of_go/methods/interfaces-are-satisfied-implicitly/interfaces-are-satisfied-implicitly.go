package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	// This method means type T implements the interface I,
	// but we don't need to explicitly declare that it does so.
	fmt.Println(t.S)
}

// Why this matters?
// Zero Upfront Coordination
// Trivial Testing
// Standard Library Compatibility
func main() {
	var i I = T{"hello"}
	i.M()
}
