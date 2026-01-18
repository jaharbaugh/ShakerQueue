package handlers

import (
	"net/http"
	//"context"

	"github.com/jaharbaugh/ShakerQueue/internal/database"
	"github.com/jaharbaugh/ShakerQueue/internal/app"
	//"github.com/jaharbaugh/ShakerQueue/internal/auth"
	"github.com/jaharbaugh/ShakerQueue/internal/models"
)

func HandleListRecipes(deps app.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		if !RequireRole(req, database.UserRoleAdmin, database.UserRoleCustomer) {
			RespondWithError(w, http.StatusForbidden, "Insufficient permissions", nil)
			return
		}

		recipes, err := deps.Queries.GetAllRecipes(req.Context())
		if err != nil{
			RespondWithError(w, http.StatusInternalServerError, "Could not retrieve recipes", err)
			return
		}

		RespondWithJSON(w, http.StatusOK, models.ListMenuResponse{
			Menu: recipes,
		})

	}
}