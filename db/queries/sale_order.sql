-- name: CreateSaleOrder :one
INSERT INTO sale_orders (
    account_name,
    user_id,
    payment_id
) VALUES (
    $1, $2, $3
) RETURNING id;