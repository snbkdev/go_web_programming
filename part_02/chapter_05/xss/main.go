package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, r.FormValue("comment"))
}

func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/form.html")
	t.Execute(w, nil)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8081",
	}

	http.HandleFunc("/process", process)
	http.HandleFunc("/form", form)
	server.ListenAndServe()
}
