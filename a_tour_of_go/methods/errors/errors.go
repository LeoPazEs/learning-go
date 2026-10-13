package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	// A nil error denotes success; a non-nil error denotes failure.
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

// A nil error denotes success; a non-nil error denotes failure.
func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
