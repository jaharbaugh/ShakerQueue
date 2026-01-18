package handlers

import (
	"net/http"
	"context"

	"github.com/jaharbaugh/ShakerQueue/internal/database"
	"github.com/jaharbaugh/ShakerQueue/internal/app"
	//"github.com/jaharbaugh/ShakerQueue/internal/auth"
	//"github.com/jaharbaugh/ShakerQueue/internal/models"
)

func HandleListOrders(deps app.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		if !RequireRole(req, database.UserRoleAdmin) {
			RespondWithError(w, http.StatusForbidden, "Insufficient permissions", nil)
			return
		}

		orders, err := deps.Queries.GetAllOrders(context.Background())
		if err != nil{
			RespondWithError(w, http.StatusInternalServerError, "Could not retrieve orders", err)
			return
		}

		RespondWithJSON(w, http.StatusOK, orders)

	}
}