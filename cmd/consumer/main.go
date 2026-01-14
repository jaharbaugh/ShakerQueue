package main

import(
	"fmt"
	"log"
	"os"
	"database/sql"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/jaharbaugh/ShakerQueue/internal/queue"
	"github.com/jaharbaugh/ShakerQueue/internal/handlers"
	"github.com/jaharbaugh/ShakerQueue/internal/app"
	"github.com/jaharbaugh/ShakerQueue/internal/database"
)

func main() {
	fmt.Println("Welcome to the ShakerQueue")
	
	//Connect to Rabbitmq
	connectionString := "amqp://guest:guest@localhost:5672/"
	connection, err := amqp.Dial(connectionString)
	if err != nil {
		log.Fatalf("Could not connect to amqp: %v", err)
	}
	defer connection.Close()

	ch, err := connection.Channel()
	if err != nil {
		log.Fatalf("Could not open new channel on the connection: %v", err)
	}

	//Connect to the database
	pathToDB := os.Getenv("DATABASE_URL")
	if pathToDB == "" {
		log.Fatal("DATABASE_URL must be set")
	}
	//Open PSQL database
	db, err := sql.Open("postgres", pathToDB)
	if err != nil {
		log.Fatalf("Could not open Postgresql database: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	defer db.Close()
	empty := ""
	deps := app.Dependencies{
		DB:       db,
		Queries:  database.New(db),
		AMQPConn: connection,
		AMQPChan: ch,
		JWTSecret : &empty,
	}
	//Prompt Login
	creds, err := ConsumerWelcome()
	if err != nil {
		log.Fatalf("Could not find credentials: %v", err)
	}

	//Request Login from Server
	loginResponse, err := Login("http://localhost:8080", creds)
	if err != nil{
		log.Fatal(err)
	}

	//Subscribe to Queue
	err = queue.SubscribeJSON(
		connection,
		queue.ExchangeDirect,
		"orders."+loginResponse.User.ID.String(),
		"orders.created",
		queue.SimpleQueueDurable,
		handlers.HandleProcessOrder(deps),
	)
	//Infinite Block
	for{}
}