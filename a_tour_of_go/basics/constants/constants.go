package main

import "fmt"

// CONST CANNOT BE :=

const Pi = 3.14

func main() {
	const World string = "世界"
	fmt.Println("Hello", World)
	fmt.Printf("Type of World: %T\n", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
