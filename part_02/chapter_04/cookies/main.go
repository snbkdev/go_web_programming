package main

import "net/http"

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name: "first_ccokie",
		Value: "Go Web Programming",
		HttpOnly: true,
	}

	c2 := http.Cookie{
		Name: "second_cookie",
		Value: "Manning Publications Go",
		HttpOnly: true,
	}

	w.Header().Set("Set-Cookie", c1.String())
	w.Header().Set("Set-Cookie", c2.String())
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8081",
	}

	http.HandleFunc("/set_cookie", setCookie)
	server.ListenAndServe()
}