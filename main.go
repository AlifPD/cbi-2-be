package main

import (
	"cbi-2-be/handlers"
	"cbi-2-be/middleware"
	"net/http"
	"os"
)

func init() {
	middleware.JwtKey = []byte(os.Getenv("JWT_SECRET"))
}

func main() {
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)

	http.HandleFunc("/todos", middleware.Auth(handlers.GetTodos))
	http.HandleFunc("/todos/create", middleware.Auth(handlers.CreateTodo))
	http.HandleFunc("/todos/update", middleware.Auth(handlers.UpdateTodo))
	http.HandleFunc("/todos/delete", middleware.Auth(handlers.DeleteTodo))

	http.ListenAndServe(":8080", nil)
}
