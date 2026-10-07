package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	fmt.Println("-------------------------------------------------")

	b[0] = "XXX" // Changed the underlying array
	fmt.Println(a, b)
	fmt.Println(names)
	fmt.Println("-------------------------------------------------")
	pb := &b
	(*pb)[1] = "OPAOPA"
	fmt.Println(b)
	fmt.Println(names)
	fmt.Println("-------------------------------------------------")
	b = names[0:4]
	fmt.Println(b)
	fmt.Println("-------------------------------------------------")
}
