package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, template.HTML(r.FormValue("comment")))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8081",
	}

	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
