// forms.go
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func ShowForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("forms.html"))
	tmpl.Execute(w, nil)
}

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func RecieveForm(w http.ResponseWriter, r *http.Request) {
	details := ContactDetails{
		r.FormValue("email"),
		r.FormValue("subject"),
		r.FormValue("message"),
	}
	fmt.Println(details)
	tmpl := template.Must(template.ParseFiles("forms.html"))
	tmpl.Execute(w, struct {
		Success bool
		Details *ContactDetails
	}{true, &details})
}

func main() {
	http.HandleFunc("GET /", ShowForm)
	http.HandleFunc("POST /", RecieveForm)
	http.ListenAndServe(":8000", nil)
}
