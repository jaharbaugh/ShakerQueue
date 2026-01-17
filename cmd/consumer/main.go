package main

import(
	"fmt"
	"log"
	//"os"
	//"database/sql"
	"errors"
	"net/http"
	"time"

	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/jaharbaugh/ShakerQueue/internal/queue"
	//"github.com/jaharbaugh/ShakerQueue/internal/handlers"
	"github.com/jaharbaugh/ShakerQueue/internal/app"
	"github.com/jaharbaugh/ShakerQueue/internal/database"
	"github.com/jaharbaugh/ShakerQueue/internal/models"
)
const urlAddress = "http://localhost:8080"
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
/*
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

	deps := app.ConsumerDependencies{
		//DB:       db,
		//Queries:  database.New(db),
		AMQPConn: connection,
		AMQPChan: ch,
		//JWTSecret : nil,
	}
*/
	//Prompt Login
	loginCreds, err := ConsumerWelcome()
	if err != nil {
		log.Fatalf("Could not find credentials: %v", err)
	}

	//Request Login from Server
	loginResponse, err := Login(urlAddress, loginCreds)
	if err != nil{
		if errors.Is(err, models.ErrUserNotFound){
			fmt.Println("No account found. Registering new user:")
			username, err := ConsumerGetNewUsername()
			registerCreds := models.RegisterUserRequest{
				Username:	username,
				Email:	loginCreds.Email,
				Password: loginCreds.Password,
			}
			regResp, err := RegisterUser(urlAddress, registerCreds)
			if err != nil {
				log.Fatalf("registration failed: %v", err)
			}
			loginResponse = &models.LogInResponse{
				User:  regResp.User,
				Token: regResp.Token,
			}
		} else {
		log.Fatalf("login failed: %v", err)
		}
	
	}
	//Create Session State
	sessionClient := app.Client{
		BaseURL: urlAddress,
		BearerToken: loginResponse.Token,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		AMQPChan: ch,
		AMQPConn: connection,
	}

	
	//Infinite Block
	
		switch loginResponse.User.Role{
		case database.UserRoleCustomer:
			for {
				PrintCustomerCommands()
				newRequest := GetInput()
				switch newRequest[0]{
				case "create":
					fmt.Println("What drink would you like?")
					cocktail := GetInput()
					CreateOrder(sessionClient, cocktail[0])
					fmt.Println("Order Created Successfully")
				case "status":
				case "exit":
				case "help":
					PrintCustomerHelp()
					continue
				}
			}
		case database.UserRoleEmployee:
			fmt.Println("Employee ready. Waiting for orders...")
			//Subscribe to Queue
			err = queue.SubscribeJSON(
				connection,
				queue.ExchangeDirect,
				//"orders."+loginResponse.User.ID.String(),
				"orders.employee",      
				"orders.created",
				queue.SimpleQueueDurable,
				ProcessOrder(sessionClient),
			)
			if err != nil {
				log.Fatalf("Failed to subscribe: %v", err)
			}
			select {}
		}

}