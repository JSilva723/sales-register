-- name: CreatePayment :one
INSERT INTO payments (
    account_name,
    name    
) VALUES (
    $1, $2
) RETURNING id, name;

-- name: GetPayments :many
SELECT id, name 
FROM payments
WHERE account_name = $1 AND is_active = true
LIMIT $2
OFFSET $3;

-- name: ChangePaymentName :one
UPDATE payments 
SET name = $1
WHERE id = $2 AND is_active = true
RETURNING id, name;

-- name: DeletePayment :exec
UPDATE payments 
SET is_active = false
WHERE id = $1;