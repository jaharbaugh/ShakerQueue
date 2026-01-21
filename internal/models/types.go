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
	Name string `json:"name"`
}

type UpdateUserRoleRequest struct {
	Email   string `json:"email"`
	NewRole string `json:"new_role"`
}

type UpdateUserRoleResponse struct {
	Status string `json:"status"`
}

type ProcessOrderRequest struct {
	OrderID uuid.UUID `json:"order_id"`
}
type ListOrderResponse struct {
	Orders []database.Order `json:"orders"`
}

type ListMenuResponse struct {
	Menu []database.CocktailRecipe `json:"menu"`
}

type OrderStatusResponse struct {
	Orders []database.Order `json:"orders"`
}

type OrderEvent struct {
	OrderID  uuid.UUID
	UserID   uuid.UUID
	RecipeID uuid.UUID
	Priority uint8
	Delay uint8
}

type CreateCocktailRecipeRequest struct {
	Name        string            `json:"name"`
	Ingredients map[string]string `json:"ingredients"`
	BuildType   string            `json:"build_type"`
}

type CreateCocktailRecipeResponse struct {
	Recipe database.CocktailRecipe `json:"recipe"`
}

type CreateOrderResponse struct {
}
