-- name: CreateUser :one
INSERT INTO users (
    username,
    password,
    account_name,
    rol
) VALUES (
    $1, $2, $3, $4
) RETURNING *;
