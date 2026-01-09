-- name: CreateUser :one

INSERT INTO users(
    id,
    created_at,
    updated_at,
    username,
    email,
    role,
    hashed_password
    )
VALUES(
    $1,
    NOW(),
    NOW(),
    $2,
    $3,
    $4,
    $5
)
RETURNING *;