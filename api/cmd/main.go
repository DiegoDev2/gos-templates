package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DiegoDev2/gos-templates/api/handler"
	"github.com/DiegoDev2/gos-templates/api/middleware"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.Use(middleware.JWTMiddleware)

	// Register routes
	r.HandleFunc("/api/v1/users", handler.GetUsers).Methods("GET")
	r.HandleFunc("/api/v1/login", handler.Login).Methods("POST")

	// Start server
	http.Handle("/", r)
	fmt.Println("API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
