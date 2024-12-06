package handler

import (
	"encoding/json"
	"net/http"
)

// GetUsers returns a list of users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := []map[string]string{
		{"username": "user1"},
		{"username": "user2"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
