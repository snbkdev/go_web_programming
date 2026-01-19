package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/t1.html", "templates/t2.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, "hello")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8081",
	}

	http.HandleFunc("/process", process)
	server.ListenAndServe()
}