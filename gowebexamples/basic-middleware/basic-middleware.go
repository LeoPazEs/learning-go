package main

import (
	"fmt"
	"log"
	"net/http"
)

func Logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func ShowValue(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", r.PathValue("value"))
}

func main() {
	r := http.NewServeMux()
	r.HandleFunc("GET /{value}", Logging(ShowValue))
	http.ListenAndServe(":8000", r)
}
