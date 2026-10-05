package main

import "fmt"

// Outside functions this does not work.
// k := 3
// But this works, always starts with a keyword
var i, j int = 1, 2

func main() {
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
