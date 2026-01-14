package auth

import (
	"context"
	"errors"
	"github.com/jaharbaugh/ShakerQueue/internal/app"
	"github.com/jaharbaugh/ShakerQueue/internal/database"
	"github.com/jaharbaugh/ShakerQueue/internal/models"
	"time"
)

var ErrInvalidCredentials = errors.New("invalid email or password")

func AuthService(deps app.Dependencies, creds models.LogInRequest) (database.User, string, error) {

	dbUser, err := deps.Queries.GetUserByEmail(context.Background(), creds.Email)
	if err != nil {
		return database.User{}, "", ErrInvalidCredentials
	}

	match, err := CheckPasswordHash(creds.Password, dbUser.HashedPassword)
	if match != true {
		return database.User{}, "", ErrInvalidCredentials
	}

	tokenString, err := MakeJWT(dbUser.ID, *deps.JWTSecret, time.Hour)
	if err != nil {
		return database.User{}, "", err
	}

	return dbUser, tokenString, nil
}
