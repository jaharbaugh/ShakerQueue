-- name: CreateOrder :one

INSERT INTO orders(
    id,
    created_at,
    updated_at,
    user_id,
    recipe_id,
    status
)
VALUES(
    $1,
    NOW(),
    NOW(),
    $2,
    $3,
    'pending'::order_status
)

RETURNING *;