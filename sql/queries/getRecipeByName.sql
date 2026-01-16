-- name: GetRecipeByName :one

SELECT * FROM cocktail_recipes
WHERE name = $1;
