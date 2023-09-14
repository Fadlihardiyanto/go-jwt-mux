package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Fadlihardiyanto/go-jwt-mux/controllers/authcontroller"
	"github.com/Fadlihardiyanto/go-jwt-mux/controllers/productcontroller"
	"github.com/Fadlihardiyanto/go-jwt-mux/middlewares"
	"github.com/Fadlihardiyanto/go-jwt-mux/models"
	"github.com/gorilla/mux"
)

func main() {

	// connect database
	models.ConnectDatabase()

	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/products", productcontroller.Index).Methods("GET")
	api.Use(middlewares.JWTMiddleware)

	fmt.Println("Server running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
