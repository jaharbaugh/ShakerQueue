package handlers

import (
	"context"
	"encoding/json"
	//"github.com/alexedwards/argon2id"
	"github.com/google/uuid"
	"github.com/jaharbaugh/ShakerQueue/internal/app"
	"github.com/jaharbaugh/ShakerQueue/internal/auth"
	"github.com/jaharbaugh/ShakerQueue/internal/database"
	"github.com/jaharbaugh/ShakerQueue/internal/models"
	"net/http"
)

func HandleRegisterUser(deps app.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		if req.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		params := models.RegisterUserRequest{}

		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&params)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
			return
		}

		hash, err := auth.HashPassword(params.Password)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Could not hash password", err)
			return
		}
		/*
			var role database.UserRole
			switch params.Role{
			case "customer":
				role = database.UserRoleCustomer
			case "employee":
				role = database.UserRoleEmployee
			case "admin":
				role = database.UserRoleAdmin
			}*/
		newUserParams := database.CreateUserParams{
			ID:             uuid.New(),
			Username:       params.Username,
			Email:          params.Email,
			Role:           database.UserRoleCustomer,
			HashedPassword: hash,
		}
		newUser, err := deps.Queries.CreateUser(context.Background(), newUserParams)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Could not register new user", err)
			return
		}

		token, err := auth.CreateSession(deps, newUser)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Could not create new user session", err)
			return
		}
		registrationResp := models.RegisterUserResponse{
			User:  newUser,
			Token: token,
		}

		RespondWithJSON(w, http.StatusOK, registrationResp)
	}
}
