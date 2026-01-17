-- name: UpdateOrderStatus :exec

UPDATE orders
SET
    status = $1,
    updated_at = NOW()
WHERE id = $2;