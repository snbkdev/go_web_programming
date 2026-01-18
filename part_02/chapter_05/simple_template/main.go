package main

import (
	"html/template"
	"log"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Ошибка загрузки шаблона: %v", err)
		return
	}
	err = t.Execute(w, "Hello chapter 05")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Ошибка выполнения шаблона: %v", err)
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8081",
	}

	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
