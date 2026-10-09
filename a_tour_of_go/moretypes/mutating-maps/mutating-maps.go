package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// V is VALUE and ok is the True if exists and False if doesnt, if ok is false the v will be the zero value.
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	fmt.Println("The value:", m["XRARASLDKFJASLDJ"])
}
