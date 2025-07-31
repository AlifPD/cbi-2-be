package main

import (
	"cbi-2-be/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)

	http.ListenAndServe(":8080", nil)
}
