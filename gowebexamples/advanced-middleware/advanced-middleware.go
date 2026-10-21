package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

// TODO: IMPLEMENT THE CHAIN

func Logging() Middleware {
	start := time.Now()
	defer func() { log.Println(r.URL.Path, time.Since(start)) }()
}

// Method ensures that url can only be requested with a specific method, else returns a 400 Bad Request
func Method(m string) Middleware {
	if r.Method != m {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
}

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	http.HandleFunc("/", Chain(Hello, Method("GET"), Logging()))
	http.ListenAndServe(":8080", nil)
}
