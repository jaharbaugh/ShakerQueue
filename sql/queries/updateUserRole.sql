-- name: UpdateUserRole :exec

UPDATE users
SET
    role = $1,
    updated_at = NOW()
WHERE id = $2;