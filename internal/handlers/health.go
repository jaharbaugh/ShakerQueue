package handlers

import (
	"github.com/jaharbaugh/ShakerQueue/internal/app"
	"github.com/jaharbaugh/ShakerQueue/internal/database"
	"net/http"
)

func HandleHealth(deps app.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		if !RequireRole(req, database.UserRoleAdmin) {
			RespondWithError(w, http.StatusForbidden, "Insufficient permissions", nil)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
}
