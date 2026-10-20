package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, `
			<head>
				<link rel="stylesheet" href="/static/css/style.css">
			</head>
			<body>
			<h1>Welcome! :D</h1>
			</body>
			`)
		fmt.Print(err)
	})
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	err := http.ListenAndServe(":8000", nil)
	fmt.Println(err)
}
