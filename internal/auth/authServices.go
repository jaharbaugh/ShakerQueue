package auth

import (
	"context"
	//"errors"
	"github.com/jaharbaugh/ShakerQueue/internal/app"
	"github.com/jaharbaugh/ShakerQueue/internal/database"
	"github.com/jaharbaugh/ShakerQueue/internal/models"
	"time"
)

func AuthenticateUser(deps app.Dependencies, creds models.LogInRequest) (database.User, error) {

	dbUser, err := deps.Queries.GetUserByEmail(context.Background(), creds.Email)
	if err != nil {
		return database.User{}, models.ErrInvalidCredentials
	}

	match, err := CheckPasswordHash(creds.Password, dbUser.HashedPassword)
	if match != true {
		return database.User{}, models.ErrInvalidCredentials
	}

	return dbUser, nil
}


func CreateSession (deps app.Dependencies, user database.User) (string, error) {

	tokenString, err := MakeJWT(user.ID, *deps.JWTSecret, time.Hour)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}