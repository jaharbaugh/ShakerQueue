package main

import (
	"fmt"
	"os"
	"log"
	//"net/http"

	"github.com/joho/godotenv"
)

func main(){
	fmt.Println("Setting up the bar now...")
	fmt.Println("Welcome to ShakerQueue!")
	
	//Load .env values
	godotenv.Load(".env")

	//Connect to the database
	pathToDB := os.Getenv("DATABASE_URL")
	if pathToDB == "" {
		log.Fatal("DATABASE_URL must be set")
	}

}