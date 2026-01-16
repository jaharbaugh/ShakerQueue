package models

import(
	"errors"
)

var ErrInvalidCredentials = errors.New("Invalid Email or Password")

var ErrUserNotFound = errors.New("User Not Found")

