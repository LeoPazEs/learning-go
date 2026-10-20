package main

import (
	"fmt"
	"net/http"
)

func main() {
	// gorilla/mux is deprecated
	r := http.NewServeMux()
	r.HandleFunc("GET /books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		title := r.PathValue("title")
		page := r.PathValue("page")
		_, err := fmt.Fprintf(w, "The title is %s in page %s\n", title, page)
		fmt.Println(err)
	})
	err := http.ListenAndServe(":8000", r)
	fmt.Println(err)
}
