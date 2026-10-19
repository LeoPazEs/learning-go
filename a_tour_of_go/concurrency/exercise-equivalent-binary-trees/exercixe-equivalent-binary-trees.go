package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walking(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walking(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walking(t.Right, ch)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	Walking(t, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	select {
	case vch1 := <-ch1:
		if vch1 != <-ch2 {
			return false
		}
	case vch2 := <-ch2:
		if vch2 != <-ch1 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
