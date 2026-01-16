package app

import (
	"database/sql"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/jaharbaugh/ShakerQueue/internal/database"
)

type Dependencies struct {
	DB        *sql.DB
	Queries   *database.Queries
	AMQPConn  *amqp.Connection
	AMQPChan  *amqp.Channel
	JWTSecret *string
}
