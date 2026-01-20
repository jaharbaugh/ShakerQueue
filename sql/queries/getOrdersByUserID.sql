-- name: GetOrdersByUserID :many

SELECT * FROM orders
WHERE user_id = $1
ORDER BY 
    status,
    created_at;