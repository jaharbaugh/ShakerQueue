-- name: GetAllOrders :many

SELECT * FROM orders
ORDER BY 
    status,
    created_at;