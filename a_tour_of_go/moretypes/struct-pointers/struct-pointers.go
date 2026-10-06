package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	// When struct pointers does not need the explicit deference
	p.X = 1e9
	fmt.Println(v)
}
