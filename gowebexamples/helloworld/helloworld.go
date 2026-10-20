package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// w.Write(byte("Hello World\n"))
		size, err := fmt.Fprintf(w, "Hello World")
		fmt.Println(size, err)
	})
	err := http.ListenAndServe(":8000", nil)
	fmt.Println(err)
}
