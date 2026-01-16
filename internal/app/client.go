package app

import (
	"net/http"
	//"github.com/jaharbaugh/ShakerQueue/internal/database"
)

type Client struct {
	BaseURL     string
	BearerToken string
	HTTPClient  *http.Client
}
