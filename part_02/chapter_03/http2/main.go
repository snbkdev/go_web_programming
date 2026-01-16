package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/http2"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!!!")
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr: "127.0.0.1:8081",
		Handler: &handler,
	}

	http2.ConfigureServer(&server, &http2.Server{})
	server.ListenAndServeTLS("cert.pem", "key.pem")
}