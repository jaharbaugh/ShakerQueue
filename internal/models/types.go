package models

import (
	"github.com/google/uuid"
	"github.com/jaharbaugh/ShakerQueue/internal/database"
)

type RegisterUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	//Role	string	`json:"role"`
}

type RegisterUserResponse struct {
	User  database.User `json:"user"`
	Token string        `json:"token"`
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
	//OrderID   uuid.UUID `json:"order_id"`
	Name string `json:"name"`
	//UserID: uuid.UUID `json:"user_id"`
}

type UpdateUserRoleRequest struct {
	Email    string `json:"email"`
	NewRole string        `json:"new_role"`
}

type UpdateUserRoleResponse struct {
	Status string `json:"status"`
}

type ProcessOrderRequest struct {
	OrderID	uuid.UUID `json:"order_id"`
}


type OrderEvent struct {
	OrderID  uuid.UUID
	UserID   uuid.UUID
	RecipeID uuid.UUID
	Quantity int
}

type CreateOrderResponse struct {
}

