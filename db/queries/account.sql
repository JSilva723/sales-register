-- name: CreateAccount :one
INSERT INTO accounts (
    id,
    name
) VALUES (
    $1, $2
) RETURNING name;
