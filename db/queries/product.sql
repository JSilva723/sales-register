-- name: CreateProduct :one
INSERT INTO products (
    account_name,
    name    
) VALUES (
    $1, $2
) RETURNING id, name;

-- name: GetProducts :many
SELECT id, name 
FROM products
WHERE account_name = $1 AND is_active = true
LIMIT $2
OFFSET $3;

-- name: ChangeProductName :one
UPDATE products 
SET name = $1
WHERE id = $2 AND is_active = true
RETURNING id, name;

-- name: DeleteProduct :exec
UPDATE products 
SET is_active = false
WHERE id = $1;