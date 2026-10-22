package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Logging() Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()
			hf(w, r)
		}
	}
}

// Method ensures that url can only be requested with a specific method, else returns a 405 Method Not Allowed
func Method(m string) Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusBadRequest)
				return
			}
			hf(w, r)
		}
	}
}

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
		// LOOP 1
		// f = Method(GET)(Hello)
		// f = func(Hello)
		// LOOP 2
		// f = Logging(func(Hello))
		// f = func(w,r)(func(w,r)(Hello(w,r)))
	}
	return f
}

func Hello(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "hello world")
	fmt.Println(err)
}

func main() {
	http.HandleFunc("/", Chain(Hello, Method("GET"), Logging()))
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}
