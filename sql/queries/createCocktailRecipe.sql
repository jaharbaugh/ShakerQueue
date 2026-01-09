-- name: CreateCocktailRecipe :one
INSERT INTO cocktail_recipes(
    id,
    created_at,
    updated_at,
    name,
    ingredients,
    build
)
VALUES(
    $1,
    NOW(),
    NOW(),
    $2,
    $3,
    $4
)

RETURNING *;