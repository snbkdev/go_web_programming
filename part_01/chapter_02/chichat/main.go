package main

import (
	"net/http"
	"web_programming/part_01/chapter_02/chichat/data"
	"log"
)

func main() {
	if err := data.Init(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer data.Close()

	db := data.GetDB()
	
	var version string
	err := db.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		log.Printf("Query error: %v", err)
	} else {
		log.Printf("Database version: %s\n", version)
	}

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)
	//mux.HandleFunc("/err", err)

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr: "0.0.0.0:8082",
		Handler: mux,
	}

	server.ListenAndServe()
}
