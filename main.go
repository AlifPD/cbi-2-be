package main

import (
	"cbi-2-be/handlers"
	"cbi-2-be/middleware"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func withCORS(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		h(w, r)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Gagal membaca file .env")
	}

	middleware.JwtKey = []byte(os.Getenv("JWT_SECRET"))
	handlers.LoadAdmin()

	http.HandleFunc("/register", withCORS(handlers.Register))
	http.HandleFunc("/login", withCORS(handlers.Login))

	http.HandleFunc("/todos", withCORS(middleware.Auth(handlers.GetTodos)))
	http.HandleFunc("/todos/create", withCORS(middleware.Auth(handlers.CreateTodo)))
	http.HandleFunc("/todos/update", withCORS(middleware.Auth(handlers.UpdateTodo)))
	http.HandleFunc("/todos/delete", withCORS(middleware.Auth(handlers.DeleteTodo)))

	http.ListenAndServe(":8080", nil)
}
