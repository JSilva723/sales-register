-- name: CreateUser :one
INSERT INTO users (
    username,
    password,
    account_id
) VALUES (
    $1, $2, $3
) RETURNING *;