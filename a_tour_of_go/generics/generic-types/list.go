package main

import (
	"errors"
	"fmt"
)

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}
type ErrListWithoutNext[T any] struct {
	Node *List[T]
}

func (e *ErrListWithoutNext[T]) Error() string {
	return fmt.Sprintf("Node with value: %v, in address %p without next", e.Node.val, e.Node)
}

var ErrListNotInitialized = errors.New("List not initialized")

func (node *List[T]) Add(val T) (*List[T], error) {
	if node == nil {
		return nil, ErrListNotInitialized
	}
	currentNode := node
	for {
		if currentNode.next == nil {
			newNode := List[T]{nil, val}
			node.next = &newNode
			return &newNode, nil
		}
		currentNode = currentNode.next
	}
}

func (node *List[T]) Pop() (List[T], error) {
	if node == nil {
		return List[T]{}, ErrListNotInitialized
	}
	if node.next == nil {
		return List[T]{}, &ErrListWithoutNext[T]{node}
	}
	currentNode := node
	var previousNode *List[T] = nil
	for {
		if currentNode.next == nil {
			previousNode.next = nil
			return *currentNode, nil
		}
		previousNode = currentNode
		currentNode = currentNode.next
	}
}

func (node *List[T]) Print() error {
	if node == nil {
		return ErrListNotInitialized
	}
	currentNode := node
	for {
		if currentNode.next == nil {
			fmt.Printf("%v -> nil\n", currentNode.val)
			break
		}
		fmt.Printf("%v -> ", currentNode.val)
		currentNode = currentNode.next
	}
	return nil
}

func main() {
	a := List[int]{nil, 10}
	a.Print()
	a.Add(15)
	a.Print()
	deletedNode, err := a.Pop()
	a.Print()
	fmt.Println(deletedNode, err)
	deletedNode, err = a.Pop()
	a.Print()
	fmt.Println(deletedNode, err)
	var list *List[int]
	fmt.Println(list.Pop())
	fmt.Println(list.Add(15))
}
