package handlers

import (
	"net/http"
	//"context"

	"github.com/google/uuid"
	"github.com/jaharbaugh/ShakerQueue/internal/database"
	"github.com/jaharbaugh/ShakerQueue/internal/app"
	"github.com/jaharbaugh/ShakerQueue/internal/auth"
	"github.com/jaharbaugh/ShakerQueue/internal/models"
)

func HandleOrderStatus(deps app.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		if !RequireRole(req, database.UserRoleAdmin, database.UserRoleCustomer) {
			RespondWithError(w, http.StatusForbidden, "Insufficient permissions", nil)
			return
		}

		userID, ok := req.Context().Value(auth.UserIDKey).(uuid.UUID)
		if !ok {
			RespondWithError(w, http.StatusUnauthorized, "User not authenticated", nil)
			return
		}

		orders, err := deps.Queries.GetOrdersByUserID(req.Context(), userID)
		if err != nil{
			RespondWithError(w, http.StatusInternalServerError, "Could not retrieve orders", err)
			return
		}

		RespondWithJSON(w, http.StatusOK, models.OrderStatusResponse{
			Orders: orders,
		})
	}
}
