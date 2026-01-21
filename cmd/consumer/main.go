package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
	"os"
	//"strings"

	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/joho/godotenv"
	//"github.com/jaharbaugh/ShakerQueue/internal/queue"
	"github.com/jaharbaugh/ShakerQueue/internal/app"
	"github.com/jaharbaugh/ShakerQueue/internal/database"
	"github.com/jaharbaugh/ShakerQueue/internal/models"
)

//const urlAddress = "http://localhost:8080"

func main() {
	fmt.Println("Welcome to the ShakerQueue")
	//Load .env values
	if err := godotenv.Load(); err != nil {
	log.Println("No .env file found, relying on environment variables")
	}
	urlAddress := os.Getenv("BASE_URL")
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
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			fmt.Println("No account found. Registering new user:")
			username, err := ConsumerGetNewUsername()
			registerCreds := models.RegisterUserRequest{
				Username: username,
				Email:    loginCreds.Email,
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
		BaseURL:     urlAddress,
		BearerToken: loginResponse.Token,
		HTTPClient:  &http.Client{Timeout: 10 * time.Second},
		AMQPChan:    ch,
		AMQPConn:    connection,
	}

	//Infinite Block

	switch loginResponse.User.Role {
	case database.UserRoleAdmin:
		for {
			fmt.Println()
			fmt.Println("🧾 Manager Console — ShakerQueue")
			PrintAdminCommands()
			newRequest := GetInput()
			switch newRequest[0] {
			case "health":
				fmt.Println("🩺 Running a systems check...")
				health, err := Health(sessionClient)
				if err != nil {
					log.Fatalf("❌ Could not get health info: %v", err)
				}
				fmt.Println("All systems reporting in:")
				fmt.Println(health)
				Divider()

			case "list":
				fmt.Println("📋 Pulling the full order ledger...")
				orders, err := ListOrders(sessionClient)
				if err != nil {
					fmt.Printf("❌ Could not retrieve orders: %v \n", err)
					return
				}
				for _, order := range orders.Orders {
					fmt.Println("Order ID:   ", order.ID)
					fmt.Println("Status:     ", order.Status)
					fmt.Println("Created At: ", order.CreatedAt)
					fmt.Println("Recipe ID:  ", order.RecipeID)
					fmt.Println("User ID:    ", order.UserID)
					Divider()
				}

			case "role":
				fmt.Println("🔐 User management")
				fmt.Println("Whose role are we changing? (email)")
				email := GetInput()
				fmt.Println("What’s their new role?")
				newRole := GetInput()
				err := UpdateUserRole(sessionClient, email[0], newRole[0])
				if err != nil {
					fmt.Printf("❌ Could not update user role: %v \n", err)
					return
				}
				fmt.Println("✅ Role updated. Paperwork signed and filed.")

			case "exit":
				fmt.Println("Locking the office. See you tomorrow.")
				return

			case "customer":
				PrintCustomerHelp()

			case "employee":
				PrintEmployeeHelp()
			}
		}

	case database.UserRoleCustomer:
		for {
			fmt.Println()
			fmt.Println("🍹 Welcome back to the bar")
			PrintCustomerCommands()
			newRequest := GetInput()
			switch newRequest[0] {
			case "menu":
				fmt.Println("📖 Checking what’s on the menu...")
				recipes, err := GetRecpies(sessionClient)
				if err != nil {
					fmt.Printf("❌ Could not retrieve menu: %v \n", err)
					return
				}
				for _, recipe := range recipes.Menu {
					fmt.Println("•", recipe.Name)
				}
				Divider()

			case "order":
				fmt.Println("🍸 What’ll it be?")
				cocktail := GetInput()
				CreateOrder(sessionClient, cocktail[0])
				fmt.Println("✅ Order sent to the bar. The bartender’s on it.")

			case "status":
				fmt.Println("🧾 Checking your tab...")
				orders, err := OrderStatus(sessionClient)
				if err != nil {
					fmt.Printf("❌ Could not retrieve orders: %v \n", err)
					return
				}
				for _, order := range orders.Orders {
					fmt.Println("Order ID:   ", order.ID)
					fmt.Println("Status:     ", order.Status)
					fmt.Println("Placed At:  ", order.CreatedAt)
					fmt.Println("Drink ID:   ", order.RecipeID)
					Divider()
				}

			case "exit":
				fmt.Println("Thanks for stopping by — tab closed 🍻")
				return

			case "help":
				PrintCustomerHelp()
			}
		}

	case database.UserRoleEmployee:
		for {
			fmt.Println()
			fmt.Println("🍺 Bartender Console — Clocked In")
			PrintEmployeeCommands()
			newRequest := GetInput()
			switch newRequest[0] {
			case "make":
				fmt.Println("🍹 Grabbing the shaker and hopping behind the bar...")
				err := JoinQueue(sessionClient)
				if err != nil {
					log.Fatalf("❌ Could not join the queue: %v", err)
				}

			case "add":
				//fmt.Println("🧪 Recipe creation coming soon — stay tuned.")
				fmt.Println("What is the name of the cocktail?")
				name := GetInput()
				//name = name[0]
				fmt.Println("How do you build it?")
				buildType := GetInput()
				//buildType = buildType[0]
				fmt.Println("What are the ingredients and how much do you need of each?")
				ingredients := GetInput()
				ingredientsString := ingredients[0]
				parts := make(map[string]string)
				for _, item := range SplitAndTrim(ingredientsString, ",") {
					kv := SplitAndTrim(item, ":")
					if len(kv) == 2 {
						parts[kv[0]] = kv[1]
					}
				}
				
				err := AddCocktailRecipe(sessionClient, name[0], parts, buildType[0])
				if err != nil {
    				fmt.Printf("❌ Could not add recipe: %v\n", err)
					return
				}
				fmt.Println("✅ Cocktail recipe added successfully!")

			case "exit":
				fmt.Println("Clocking out. Bar's in good hands.")
				return

			case "help":
				PrintEmployeeHelp()
			}
		}

	default:
		select {}
	}
}
