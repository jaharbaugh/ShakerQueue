-- name: GetOrderByID :one

SELECT * FROM orders
WHERE id = $1;