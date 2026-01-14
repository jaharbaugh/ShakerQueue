package auth

import (
	"github.com/alexedwards/argon2id"
)

func CheckPasswordHash(password string, hash string) (bool, error) {

	return argon2id.ComparePasswordAndHash(password, hash)

}

func HashPassword(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}

	return hash, nil
}
