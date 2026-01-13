package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/jaharbaugh/ShakerQueue/internal/app"
	"github.com/jaharbaugh/ShakerQueue/internal/database"
	"github.com/joho/godotenv"
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

	log.Println("Seeded cocktail recipes successfully.")

}
