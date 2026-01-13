package main

import (
	"net/http"

	"github.com/jaharbaugh/ShakerQueue/internal/handlers"
	"github.com/jaharbaugh/ShakerQueue/internal/app"
)

func NewRouter(deps app.Dependencies) (http.Handler, error) {
	//Create a new Server Multiplexer
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.HandleHealth(deps))
	mux.HandleFunc("/register", handlers.HandleRegisterUser(deps))
	mux.HandleFunc("/login", handlers.HandleLogIn(deps))
	
	return mux, nil
}
