package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type TodoItem struct {
	Title string
	Done  bool
}

type TodoList struct {
	Title string
	Items []TodoItem
}

func main() {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, TodoList{"Todo List", []TodoItem{{"Cagar", true}, {"Mijar", false}}})
		fmt.Println(err)
	})
	err := http.ListenAndServe(":8000", nil)
	fmt.Println(err)
}
