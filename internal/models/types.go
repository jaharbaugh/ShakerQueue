package models

import (
	"github.com/google/uuid"
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