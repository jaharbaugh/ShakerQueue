package app

import (
	"net/http"
	amqp "github.com/rabbitmq/amqp091-go"
	//"github.com/jaharbaugh/ShakerQueue/internal/database"
)

type Client struct {
	BaseURL     string
	BearerToken string
	HTTPClient  *http.Client
	AMQPConn  *amqp.Connection
	AMQPChan  *amqp.Channel
}
