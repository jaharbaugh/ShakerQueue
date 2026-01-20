package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	//"context"
	"errors"
	//"github.com/jaharbaugh/ShakerQueue/internal/database"
	"github.com/jaharbaugh/ShakerQueue/internal/models"
)

func ConsumerWelcome() (models.LogInRequest, error) {
	fmt.Println("🍸 Welcome to ShakerQueue!")
	fmt.Println("The bar is open. Let's get you a seat.")
	fmt.Println("What's your email?")
	email := GetInput()
	if len(email) == 0 {
		return models.LogInRequest{}, errors.New("no tab, no service — you need to enter an email")
	}
	fmt.Println("🍸Please enter your password:")
	password := GetInput()
	if len(password) == 0 {
		return models.LogInRequest{}, errors.New("no tab, no service — you need to enter an email")
	}
	creds := models.LogInRequest{
		Email:    email[0],
		Password: password[0],
	}

	return creds, nil
}

func ConsumerGetNewUsername() (string, error) {
	fmt.Println("What should we call you?")
	username := GetInput()
	return username[0], nil
}

func PrintCustomerCommands() {
	fmt.Println("🍹 Customer commands:")
	fmt.Println("* menu   — see what's on tap")
	fmt.Println("* order  — place a drink order")
	fmt.Println("* status — check on your drinks")
	fmt.Println("* help   — how this place works")
	fmt.Println("* exit   — close your tab")
}

func PrintCustomerHelp() {
	fmt.Println("🍹 Customer commands:")
	fmt.Println("* menu:")
	fmt.Println("    browse the cocktail menu")
	fmt.Println("")
	fmt.Println("* order:")
	fmt.Println("    place a drink order with the bar")
	fmt.Println("    one drink at a time — we're classy like that")
	fmt.Println("")
	fmt.Println("* status:")
	fmt.Println("    check the status of all drinks under your tab")
	fmt.Println("")
	fmt.Println("* exit:")
	fmt.Println("    close your tab and leave the bar")
	fmt.Println("")
	fmt.Println("* help:")
	fmt.Println("    show this menu again")
}

func PrintEmployeeCommands() {
	fmt.Println("🍺 Bartender commands:")
	fmt.Println("* make  — start mixing drinks and getting paid")
	fmt.Println("* add   — add a new cocktail recipe to the menu")
	fmt.Println("* help  — refresher on bar duties")
	fmt.Println("* exit  — clock out")
}

func PrintEmployeeHelp() {
	fmt.Println("🍺 Bartender commands:")
	fmt.Println("* make:")
	fmt.Println("    hop behind the bar and start serving orders")
	fmt.Println("    one drink at a time — no spills")
	fmt.Println("")
	fmt.Println("* add:")
	fmt.Println("    create a new cocktail recipe for the menu")
	fmt.Println("")
	fmt.Println("* exit:")
	fmt.Println("    clock out and leave the bar")
	fmt.Println("")
	fmt.Println("* help:")
	fmt.Println("    show this menu again")
}

func PrintAdminCommands() {
	fmt.Println("🧾 Manager commands:")
	fmt.Println("* health   — check the bar's vitals")
	fmt.Println("* list     — see all orders")
	fmt.Println("* role     — change a user's role")
	fmt.Println("* customer — view customer commands")
	fmt.Println("* employee — view employee commands")
	fmt.Println("* exit     — lock up for the night")
}
func PrintAdminHelp() {
	fmt.Println("🧾 Manager commands:")
	fmt.Println("* health:")
	fmt.Println("    check server health and system status")
	fmt.Println("")
	fmt.Println("* list:")
	fmt.Println("    view all orders and their current state")
	fmt.Println("")
	fmt.Println("* role:")
	fmt.Println("    change a user's role by their ID")
	fmt.Println("")
	fmt.Println("* customer:")
	fmt.Println("    show customer command list")
	fmt.Println("")
	fmt.Println("* employee:")
	fmt.Println("    show employee command list")
	fmt.Println("")
	fmt.Println("* exit:")
	fmt.Println("    shut things down and close the bar")
}
func GetInput() []string {
	fmt.Print("🍸> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanned := scanner.Scan()
	if !scanned {
		return nil
	}
	line := scanner.Text()
	line = strings.TrimSpace(line)
	return strings.Fields(line)
}

func Section(title string) {
	fmt.Println()
	fmt.Println("===", title, "===")
}

func Divider() {
	fmt.Println("----------------------------")
}


func SplitAndTrim(s string, sep string) []string {
	raw := []string{}
	for _, part := range Split(s, sep) {
		raw = append(raw, Trim(part))
	}
	return raw
}

func Split(s, sep string) []string {
	return []string(strings.Split(s, sep))
}
func Trim(s string) string {
	return strings.TrimSpace(s)
}