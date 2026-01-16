package handlers

import (
	//"context"
	"net/http"
	//"log"
	"encoding/json"

	//"github.com/alexedwards/argon2id"
	//"github.com/jaharbaugh/ShakerQueue/internal/database"
	"github.com/jaharbaugh/ShakerQueue/internal/app"
	"github.com/jaharbaugh/ShakerQueue/internal/auth"
	"github.com/jaharbaugh/ShakerQueue/internal/models"
)

func HandleLogIn(deps app.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		unverifiedUser := models.LogInRequest{}

		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&unverifiedUser)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
			return
		}

		user, err := auth.AuthenticateUser(deps, unverifiedUser)
		if err != nil {
			RespondWithError(w, http.StatusUnauthorized, "Invalid Email or Password", err)
			return
		}

		token, err := auth.CreateSession(deps, user)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Could not create new user session", err)
			return
		}
		RespondWithJSON(w, http.StatusOK, models.LogInResponse{
			User:  user,
			Token: token,
		})
	}
}
