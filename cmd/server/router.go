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
	mux.HandleFunc("/health", handlers.AuthMiddleware(deps, handlers.HandleHealth(deps)))
	mux.HandleFunc("/role/set", handlers.AuthMiddleware(deps, handlers.HandleUpdateUserRole(deps)))
	
	//Customer Endpoints
	mux.HandleFunc("/orders/create", handlers.AuthMiddleware(deps, handlers.HandleCreateOrder(deps)))
	mux.HandleFunc("/orders/status", handlers.AuthMiddleware(deps, handlers.HandleOrderStatus(deps)))
	mux.HandleFunc("/menu", handlers.AuthMiddleware(deps, handlers.HandleListRecipes(deps)))

	//Employee Endpoints
	mux.HandleFunc("/orders/complete", handlers.HandleCompleteOrder(deps))


	//Shared Endpoints
	mux.HandleFunc("/register", handlers.HandleRegisterUser(deps))
	mux.HandleFunc("/login", handlers.HandleLogIn(deps))

	
	return mux, nil
}
