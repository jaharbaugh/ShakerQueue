package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	//"context"
	"errors"
	//"github.com/jaharbaugh/ShakerQueue/internal/database"
	"github.com/jaharbaugh/ShakerQueue/internal/models"
)

func ConsumerWelcome() (models.LogInRequest, error) {
	fmt.Println("Welcome to ShakerQueue!")
	fmt.Println("Please enter your email:")
	email := GetInput()
	if len(email) == 0 {
		return models.LogInRequest{}, errors.New("you must enter an email. goodbye")
	}
	fmt.Println("Please enter your password:")
	password := GetInput()
	if len(password) == 0 {
		return models.LogInRequest{}, errors.New("you must enter an email. goodbye")
	}
	creds := models.LogInRequest{
		Email: email[0],
		Password: password[0],
	}
	//fmt.Printf("Welcome, %s!\n", username)
	//PrintClientHelp()
	return creds, nil
}

func ConsumerGetNewUsername() (string, error){
	fmt.Println("Please enter your desired username:")
	username := GetInput()
	return username[0], nil
}

func PrintCustomerCommands() {
	fmt.Println("Possible commands:")
	fmt.Println("* menu")
	fmt.Println("* order")
	fmt.Println("* status")
	fmt.Println("* exit")
	fmt.Println("* help")
}

func PrintCustomerHelp() {
	fmt.Println("Possible commands:")
	fmt.Println("* menu:")
	fmt.Println("    view all cocktails available to order")
	fmt.Println("* order: ")
	fmt.Println("    adds a new order to the queue")
	fmt.Println("    Only one drink at a time, please")
	fmt.Println("* status:")
	fmt.Println("    checks the status of all orders under your ID")
	fmt.Println("* exit:")
	fmt.Println(" 	  exits the client")
	fmt.Println("* help:")
	fmt.Println("     prints the help menu")
}

func PrintEmployeeCommands() {
	fmt.Println("Possible commands:")
	fmt.Println("* make")
	fmt.Println("* add")
	fmt.Println("* exit")
	fmt.Println("* help")
}

func PrintEmployeeHelp() {
	fmt.Println("Possible commands:")
	fmt.Println("* make: ")
	fmt.Println("    subscribes to the queue of active orders")
	fmt.Println("    Only one drink at a time, please")
	fmt.Println("* add:")
	fmt.Println("    creates new cocktail recipes to be ordered")
	fmt.Println("* exit:")
	fmt.Println(" 	  exits the client")
	fmt.Println("* help:")
	fmt.Println("     prints the help menu")
}

func PrintAdminCommands() {
	fmt.Println("Possible commands:")
	fmt.Println("* health")
	fmt.Println("* list")
	fmt.Println("* role")
	fmt.Println("* exit")
	fmt.Println("* customer")
	fmt.Println("* employee")
}
func PrintAdminHelp() {
	fmt.Println("Possible commands:")
	fmt.Println("* health: ")
	fmt.Println("    view server health data")
	fmt.Println("* list:")
	fmt.Println("    List all order and their status")
	fmt.Println("* role:")
	fmt.Println(" 	  Change the role of user by their ID")
	fmt.Println("* exit:")
	fmt.Println(" 	  exits the client")
	fmt.Println("* customer:")
	fmt.Println("     prints customer commands")
	fmt.Println("* employee:")
	fmt.Println("     prints employee commands")
}
func GetInput() []string {
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanned := scanner.Scan()
	if !scanned {
		return nil
	}
	line := scanner.Text()
	line = strings.TrimSpace(line)
	return strings.Fields(line)
}

