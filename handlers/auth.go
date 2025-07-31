package handlers

import (
	"cbi-2-be/middleware"
	"cbi-2-be/models"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var users = map[string]string{}

func LoadAdmin() {
	adminUser := os.Getenv("ADMIN_USERNAME")
	adminPass := os.Getenv("ADMIN_PASSWORD")
	if adminUser != "" && adminPass != "" {
		users[adminUser] = adminPass
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	if _, exists := users[user.Username]; exists {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}

	users[user.Username] = user.Password
	w.Write([]byte("Register successful"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	if pw, ok := users[user.Username]; !ok || pw != user.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	})

	tokenString, _ := token.SignedString(middleware.JwtKey)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
