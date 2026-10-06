package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	p := &v.X
	fmt.Println(v.X)
	fmt.Println(*p)
	*p = 14
	fmt.Println(v.X)
	fmt.Println(*p)
}
