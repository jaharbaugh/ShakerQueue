package handlers

import (
	//"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/jaharbaugh/ShakerQueue/internal/app"
	"github.com/jaharbaugh/ShakerQueue/internal/auth"
	"github.com/jaharbaugh/ShakerQueue/internal/database"
	"github.com/jaharbaugh/ShakerQueue/internal/models"
	"github.com/jaharbaugh/ShakerQueue/internal/queue"
)

func HandleCreateOrder(deps app.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		userID, ok := req.Context().Value(auth.UserIDKey).(uuid.UUID)
		if !ok {
			RespondWithError(w, http.StatusUnauthorized, "User not authenticated", nil)
			return
		}

		if !RequireRole(req, database.UserRoleAdmin, database.UserRoleCustomer) {
			RespondWithError(w, http.StatusForbidden, "Insufficient permissions", nil)
			return
		}

		newOrderRequest := models.CreateOrderRequest{}

		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&newOrderRequest)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
			return
		}
		
		cocktailRecipe, err := deps.Queries.GetRecipeByName(req.Context(), newOrderRequest.Name)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "Could not find cocktail recipe", err)
			return
		}
		
		var priority int
		var delay int
		switch cocktailRecipe.Build{
		case database.BuildTypeShaken:
			priority = 5
			delay = 45
		case database.BuildTypeStirred:
			priority = 3
			delay = 30
		case database.BuildTypeCollins:
			priority = 1
			delay = 15
		case database.BuildTypeSour:
			priority = 8
			delay = 60
		default:
			RespondWithError(w, http.StatusBadRequest, "Unknown build type", nil)
			return
		} 

		newOrder, err := deps.Queries.CreateOrder(req.Context(), database.CreateOrderParams{
			ID:       uuid.New(),
			UserID:   userID,
			RecipeID: cocktailRecipe.ID,
		})
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Could not create new order", err)
			return
		}

		event := models.OrderEvent{
			OrderID:  newOrder.ID,
			UserID:   newOrder.UserID,
			RecipeID: newOrder.RecipeID,
			Priority: uint8(priority),
			Delay: uint8(delay),
		}

		err = queue.PublishJSON(
			deps.AMQPChan,
			queue.ExchangeDirect,
			"order.created",
			event,
			uint8(priority),
			uint8(delay),
		)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Could not publish order event", err)
			return
		}

		RespondWithJSON(w, http.StatusCreated, newOrder)
	}
}
