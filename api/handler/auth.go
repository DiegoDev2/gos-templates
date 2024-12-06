package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("secret")

// Login handles user login and returns a JWT token
func Login(w http.ResponseWriter, r *http.Request) {
	// In a real application, validate user credentials
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = "user"
	claims["exp"] = 15000

	// Generate the token
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	// Return token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
