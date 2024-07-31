-- name: CreateAccount :one
INSERT INTO accounts (
    name,
    status
) VALUES (
    $1, $2
) RETURNING id;
