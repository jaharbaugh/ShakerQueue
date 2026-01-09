package handlers

import (
	"context"
	"encoding/json"
	"github.com/alexedwards/argon2id"
	"github.com/google/uuid"
	"github.com/jaharbaugh/ShakerQueue/internal/app"
	"github.com/jaharbaugh/ShakerQueue/internal/database"
	"net/http"
)

func HandleRegisterUser(deps app.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		if req.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		type parameters struct {
			username string `json:"username"`
			email    string `json:"email"`
			password string `json:"password"`
		}
		params := parameters{}

		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&params)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
			return
		}

		hash, err := HashPassword(params.password)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Could not hash password", err)
			return
		}

		newUserParams := database.CreateUserParams{
			ID:             uuid.New(),
			Username:       params.username,
			Email:          params.email,
			Role:           database.UserRoleCustomer,
			HashedPassword: hash,
		}

		newUser, err := deps.Queries.CreateUser(context.Background(), newUserParams)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Could not register new user", err)
			return
		}

		RespondWithJSON(w, http.StatusOK, newUser)
	}
}

func HashPassword(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}

	return hash, nil
}
