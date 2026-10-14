package main

import "fmt"

// Index returns the index of x in s, or -1 if not found.
// [T comparable]: This tells Go that T is a placeholder for any type. The comparable constraint ensures that whatever type T becomes (like an int or a string), it must support the == operator so the function can accurately compare values.
// s []T, x T: The function takes a slice s containing elements of type T, and a target value x which is exactly the same type T. It returns an int representing the index position.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))

	// Index does not work here because is not comparable
	// fmt.Println(Index(ss, 15))
}
