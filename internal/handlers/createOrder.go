package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/jaharbaugh/ShakerQueue/internal/app"
	//"github.com/jaharbaugh/ShakerQueue/internal/auth"
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

		newOrderRequest := models.CreateOrderRequest{}

		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&newOrderRequest)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
			return
		}

		newOrder, err := deps.Queries.CreateOrder(context.Background(), database.CreateOrderParams{
			ID:       uuid.New(),
			UserID:   uuid.New(), //Place holder
			RecipeID: newOrderRequest.RecipeID,
			//Quantity: newOrderRequest.Quantity,
			//Status: database.StatusPending,
		})
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Could not create new order", err)
			return
		}

		event := models.CreateOrderEvent{
			OrderID:  newOrder.ID,
			UserID:   newOrder.UserID,
			RecipeID: newOrder.RecipeID,
			//Quantity: newOrder.Quantity,
		}

		err = queue.PublishJSON(
			deps.AMQPChan,
			"orders",
			"order.created",
			event,
		)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Could not publish order event", err)
			return
		}

		RespondWithJSON(w, http.StatusCreated, newOrder)
	}
}
