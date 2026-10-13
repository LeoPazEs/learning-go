package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 *rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.r.Read(b)
	for i, v := range b {
		if b[i] >= 65 && b[i] <= 90 {
			b[i] = ((v-65)+13)%26 + 65
		}
		if b[i] >= 97 && b[i] <= 122 {
			b[i] = ((v-97)+13)%26 + 97
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
