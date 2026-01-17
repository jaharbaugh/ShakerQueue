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
	fmt.Println("* create")
	fmt.Println("* status")
	fmt.Println("* exit")
	fmt.Println("* help")
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

func PrintCustomerHelp() {
	fmt.Println("Possible commands:")
	fmt.Println("* create: ")
	fmt.Println("    adds a new order to the queue")
	fmt.Println("    Only one drink at a time, please")
	fmt.Println("* status:")
	fmt.Println("    checks the status of all orders under your ID")
	fmt.Println("* exit:")
	fmt.Println(" 	  exits the client")
	fmt.Println("* help:")
	fmt.Println("     prints the help menu")
}