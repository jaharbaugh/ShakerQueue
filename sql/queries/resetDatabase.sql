-- name: ResetDatabase :exec

TRUNCATE TABLE orders, cocktail_recipes, users RESTART IDENTITY CASCADE;