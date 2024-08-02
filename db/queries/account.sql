-- name: CreateAccount :one
INSERT INTO accounts (
    id,
    name,
    status
) VALUES (
    $1, $2, $3
) RETURNING name;
