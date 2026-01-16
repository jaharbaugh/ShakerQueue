package handlers


import (
	"net/http"
	//"context"
	"encoding/json"

	//"github.com/google/uuid"
	"github.com/jaharbaugh/ShakerQueue/internal/app"
	//"github.com/jaharbaugh/ShakerQueue/internal/auth"
	"github.com/jaharbaugh/ShakerQueue/internal/database"
	"github.com/jaharbaugh/ShakerQueue/internal/models"
	//"github.com/jaharbaugh/ShakerQueue/internal/queue"
)

func HandleUpdateUserRole(deps app.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		/*adminID, ok := req.Context().Value(auth.UserIDKey).(uuid.UUID)
		if !ok {
			RespondWithError(w, http.StatusUnauthorized, "User not authenticated", nil)
			return
		}*/

		if !RequireRole(req, database.UserRoleAdmin) {
			RespondWithError(w, http.StatusForbidden, "Insufficient permissions", nil)
			return
		}

		newUserRoleRequest := models.UpdateUserRoleRequest{}

		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&newUserRoleRequest)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
			return
		}

		switch newUserRoleRequest.NewRole{
		case "employee": 
				err := deps.Queries.UpdateUserRole(
					req.Context(), 
					database.UpdateUserRoleParams{
						Role: database.UserRoleEmployee,
						ID: newUserRoleRequest.UserID,
					})
				if err != nil {
					RespondWithError(w, http.StatusInternalServerError, "Could not alter user role", err)
					return
				}

		case "admin": 
				err := deps.Queries.UpdateUserRole(
					req.Context(), 
					database.UpdateUserRoleParams{
						Role: database.UserRoleAdmin,
						ID: newUserRoleRequest.UserID,
					})
				if err != nil {
					RespondWithError(w, http.StatusInternalServerError, "Could not alter user role", err)
					return
				}
		default:
			RespondWithError(w, http.StatusInternalServerError, "Could not alter user role", nil)
			return
		}

		RespondWithJSON(w, http.StatusOK, models.UpdateUserRoleResponse{
			Status: "User role updated",
		})
	}
}