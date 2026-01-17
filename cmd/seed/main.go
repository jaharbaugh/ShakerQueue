package main

import (
	"context"
	"database/sql"
	"encoding/json"
	
	"github.com/google/uuid"
	"github.com/jaharbaugh/ShakerQueue/internal/app"
	"github.com/jaharbaugh/ShakerQueue/internal/auth"
	"github.com/jaharbaugh/ShakerQueue/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {

	//Load .env values
	_ = godotenv.Load(".env")

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

	deps := app.Dependencies{
		DB:      db,
		Queries: database.New(db),
	}

	err = deps.Queries.ResetDatabase(context.Background())
	if err != nil{
		log.Fatal("Could not reset database")
	}

	baseCocktails := []database.CreateCocktailRecipeParams{
		{
			ID:          uuid.New(),
			Name:        "Old Fashioned",
			Ingredients: json.RawMessage(`{"bourbon": "2 oz", "simple syrup": "0.25 oz", "Angostura bitters": "4 dashes"}`),
			Build:       database.BuildTypeStirred,
		},
		{
			ID:          uuid.New(),
			Name:        "Daiquiri",
			Ingredients: json.RawMessage(`{"rum": "2 oz", "simple syrup": "1 oz", "lime juice": "1 oz"}`),
			Build:       database.BuildTypeShaken,
		},
		{
			ID:          uuid.New(),
			Name:        "Sidecar",
			Ingredients: json.RawMessage(`{"cognac": "2 oz", "triple sec": "1 oz", "lemon juice": "1 oz"}`),
			Build:       database.BuildTypeShaken,
		},
		{
			ID:          uuid.New(),
			Name:        "Martini",
			Ingredients: json.RawMessage(`{"gin": "2.5 oz", "dry vermouth": "0.5 oz"}`),
			Build:       database.BuildTypeStirred,
		},
		{
			ID:          uuid.New(),
			Name:        "Tom Collins",
			Ingredients: json.RawMessage(`{"gin": "2 oz", "simple syrup": "1 oz", "lemon juice": "1 oz", "soda": "3 oz"}`),
			Build:       database.BuildTypeCollins,
		},
		{
			ID:          uuid.New(),
			Name:        "Clover Club",
			Ingredients: json.RawMessage(`{"gin": "2 oz", "grenadine": "1 oz", "lemon juice": "1 oz", "egg white": "1 oz"}`),
			Build:       database.BuildTypeSour,
		},
	}

	for _, recipe := range baseCocktails {
		_, err := deps.Queries.CreateCocktailRecipe(context.Background(), recipe)
		if err != nil {
			log.Fatalf("Failed to insert %s: %v", recipe.Name, err)
		}
	}

	baseUsers := []database.CreateUserParams{
		{	ID:             uuid.New(),
			Username:       "admin",
			Email:          "admin@admin.com",
			Role:           database.UserRoleAdmin,
			HashedPassword: "admin",
		},
		{	ID:             uuid.New(),
			Username:       "james",
			Email:          "james@james.com",
			Role:           database.UserRoleEmployee,
			HashedPassword: "james",
		},
		{	ID:             uuid.New(),
			Username:       "chunk",
			Email:          "chunk@chunk.com",
			Role:           database.UserRoleCustomer,
			HashedPassword: "chunk",
		},
	}
	for _, user := range baseUsers {
    	hashed, err := auth.HashPassword(user.HashedPassword)
    	if err != nil {
        	log.Fatalf("Failed to hash password for %s: %v", user.Username, err)
    	}
    	user.HashedPassword = hashed

    	_, err = deps.Queries.CreateUser(context.Background(), user)
    	if err != nil {
        	log.Fatalf("Failed to insert %s: %v", user.Username, err)
    	}
	}

	log.Println("Seeded database successfully.")

}
