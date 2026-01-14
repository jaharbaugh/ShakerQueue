package models

import (
	"github.com/google/uuid"
	"github.com/jaharbaugh/ShakerQueue/internal/database"
)

type RegisterUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogInResponse struct {
	User  database.User `json:"user"`
	Token string        `json:"token"`
}

type CreateOrderRequest struct {
	OrderID   uuid.UUID `json:"order_id"`
	RecipeID  uuid.UUID `json:"recipe_id"`
	Quantity  int       `json:"quantity"`
}

type CreateOrderEvent struct {
	OrderID   uuid.UUID 
	UserID    uuid.UUID 
	RecipeID  uuid.UUID 
	Quantity  int       
}

type OrderEvent struct{
	
}