package handlers

import (
	"context"
	"errors"
	"net/http"
	"strings"
	//"github.com/google/uuid"
	"github.com/jaharbaugh/ShakerQueue/internal/app"
	"github.com/jaharbaugh/ShakerQueue/internal/auth"
	"github.com/jaharbaugh/ShakerQueue/internal/database"
)

func AuthMiddleware(deps app.Dependencies, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		token, err := GetBearerToken(req.Header)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "Could not retrieve JWT", err)
			return
		}

		userID, err := auth.ValidateJWT(token, *deps.JWTSecret)
		if err != nil {
			RespondWithError(w, http.StatusUnauthorized, "Could not validate JWT", err)
			return
		}

		user, err := deps.Queries.GetUserByID(context.Background(), userID)
		if err != nil {
			RespondWithError(w, http.StatusUnauthorized, "User not found", err)
			return
		}
		ctx := req.Context()
		ctx = context.WithValue(ctx, auth.UserIDKey, userID)
		ctx = context.WithValue(ctx, auth.UserRoleKey, user.Role)
		req = req.WithContext(ctx)

		next.ServeHTTP(w, req)
	}
}

func GetBearerToken(headers http.Header) (string, error) {
	headerString := headers.Get("Authorization")
	if headerString == "" {
		return "", errors.New("Header not set")
	}

	pref := "Bearer "
	tokenString := ""

	if strings.HasPrefix(headerString, pref) {
		tokenString = strings.TrimPrefix(headerString, pref)
		tokenString = strings.TrimSpace(tokenString)
		if tokenString == "" {
			return "", errors.New("Authorization token missing")
		}
	} else {
		return "", errors.New("Authorization header must use Bearer scheme")
	}

	return tokenString, nil

}

func RequireRole(req *http.Request, allowedRoles ...database.UserRole) bool {
	role := req.Context().Value(auth.UserRoleKey).(database.UserRole)
	for _, r := range allowedRoles {
		if role == r {
			return true
		}
	}
	return false
}
