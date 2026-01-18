package main

import(
	"fmt"
	"log"
	"errors"
	"net/http"
	"time"

	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
	//"github.com/jaharbaugh/ShakerQueue/internal/queue"
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
		case database.UserRoleAdmin:
			for {
				PrintAdminCommands()
				newRequest:= GetInput()
				switch newRequest[0]{
				case "health":
					fmt.Println("Checking Server Health")
					health, err := Health(sessionClient)
					if err != nil{
						log.Fatalf("Could not get health info: %v", err)
					}
					fmt.Sprintf(health)
					fmt.Println("------")
					
				case "list":
					fmt.Println("Fetching all orders")
					orders, err := ListOrders(sessionClient)
					if err != nil{
						fmt.Printf("Could not retrieve orders: %v \n", err)
						return
					}

					for _, order := range orders.Orders{
						fmt.Printf("%v\n", order.ID)
						fmt.Printf("%v\n", order.Status)
						fmt.Printf("%v\n", order.CreatedAt)
						fmt.Printf("%v\n", order.RecipeID)
						fmt.Printf("%v\n", order.UserID)
						fmt.Println("-------")
					}
				case "role":
					fmt.Println("What is the user's email?")
					email := GetInput()
					fmt.Println("What is their new role?")
					newRole := GetInput()
					err := UpdateUserRole(sessionClient, email[0], newRole[0])
					if err != nil{
						fmt.Printf("Could not upadate user role: %v \n", err)
						return
					}
					fmt.Println("User role updated successfuly")
				case "exit":
					return
				case "customer":
					PrintCustomerHelp()
					
				case "employee":
					PrintEmployeeHelp()
					
			}
		}
		case database.UserRoleCustomer:
			for {
				PrintCustomerCommands()
				newRequest := GetInput()
				switch newRequest[0]{
				case "menu":
					fmt.Println("Fetching all menu items")
					recipes, err := GetRecpies(sessionClient)
					if err != nil{
						fmt.Printf("Could not retrieve menu: %v \n", err)
						return
					}

					for _, recipe := range recipes.Menu{
						//fmt.Printf("%v\n", recipe.ID)
						fmt.Printf("%v\n", recipe.Name)
						fmt.Println("-------")
					}
				case "order":
					fmt.Println("What drink would you like?")
					cocktail := GetInput()
					CreateOrder(sessionClient, cocktail[0])
					fmt.Println("Order Created Successfully")
				case "status":
					// TODO: implement
				case "exit":
					return
				case "help":
					PrintCustomerHelp()
					
				}
			}
		case database.UserRoleEmployee:
			for {
				PrintEmployeeCommands()
				newRequest := GetInput()
				switch newRequest[0]{
				case "make":
					err := JoinQueue(sessionClient)
					if err != nil{
						log.Fatalf("Could not subscribe: %v", err)
					}
				case "add":
					// TODO: implement
				case "status":
					// TODO: implement
				case "exit":
					return
				case "help":
					PrintEmployeeHelp()
					
				}
			}
		default: select {}
		}
}
