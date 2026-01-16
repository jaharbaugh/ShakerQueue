package main

import (
	"net/http"

	"github.com/jaharbaugh/ShakerQueue/internal/handlers"
	"github.com/jaharbaugh/ShakerQueue/internal/app"
)

func NewRouter(deps app.Dependencies) (http.Handler, error) {
	//Create a new Server Multiplexer
	mux := http.NewServeMux()

	//Admin Endpoints
	mux.HandleFunc("/health", handlers.HandleHealth(deps))
	
	//Customer Endpoints
	mux.HandleFunc("/createorder", handlers.HandleCreateOrder(deps))
	
	//Employee Endpoints
	mux.HandleFunc("/processorder", handlers.HandleProcessOrder(deps))


	//Shared Endpoints
	mux.HandleFunc("/register", handlers.HandleRegisterUser(deps))
	mux.HandleFunc("/login", handlers.HandleLogIn(deps))

	
	return mux, nil
}
