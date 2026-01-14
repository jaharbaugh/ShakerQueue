package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jaharbaugh/ShakerQueue/internal/app"
	"github.com/jaharbaugh/ShakerQueue/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Setting up the bar now...")
	fmt.Println("Your shift starts now")

	//Load .env values
	_ = godotenv.Load(".env")

	tokenSecret := os.Getenv("JWT_SECRET")
	if tokenSecret == "" {
	log.Fatal("JWT_SECRET must be set")
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

	//Connect to RabbitMQ
	connectionString := os.Getenv("RABBITMQ_URL")
	connection, err := amqp.Dial(connectionString)
	if err != nil {
		log.Fatal("Could not connect to amqp")
	}
	defer connection.Close()

	fmt.Println("Connection to ShakerQueue server successful")

	ch, err := connection.Channel()
	if err != nil {
		log.Fatalf("Could not open new channel on the connection: %v", err)
	}
	defer ch.Close()

	//Connect to a new multiplexer
	deps := app.Dependencies{
		DB:       db,
		Queries:  database.New(db),
		AMQPConn: connection,
		AMQPChan: ch,
		JWTSecret : tokenSecret,
	}
	
	router, err := NewRouter(deps)
	if err != nil{
		log.Fatalf("Could not create a new router: %v", err)
	}
	http.ListenAndServe(":8080", router)

}
