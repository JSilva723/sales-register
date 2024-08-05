-- name: CreateOrderLine :exec
INSERT INTO order_lines (
    account_name,
    sale_order_id,
    ammount,
    product_id
) VALUES (
    $1, $2, $3, $4
);