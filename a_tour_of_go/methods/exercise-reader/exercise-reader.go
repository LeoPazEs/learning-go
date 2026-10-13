package main

import (
	"fmt"

	"golang.org/x/tour/reader"
)

type (
	MyReader            struct{}
	BufferTooShortError int
)

func (b BufferTooShortError) Error() string {
	return fmt.Sprintf("Buffer too short with cap= %v", uint(b))
}

func (mr MyReader) Read(b []byte) (int, error) {
	if cap(b) < 1 {
		return 0, BufferTooShortError(cap(b))
	}
	b[0] = 65
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
	x := MyReader{}
	b := make([]byte, 0)
	fmt.Print(x.Read(b))
}
