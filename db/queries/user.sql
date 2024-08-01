-- name: CreateUser :one
INSERT INTO users (
    username,
    password,
    account_name,
    rol
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetUser :one
SELECT username, account_name, rol, created_at, updated_at
FROM users
WHERE account_name = $1 AND id = $2
LIMIT 1;

-- name: ChangePassword :exec
UPDATE users
SET password = $1, updated_at = (now())
WHERE account_name = $2 AND id = $3;

-- name: ChangeRol :one
UPDATE users
SET rol = $1, updated_at =  (now())
WHERE account_name = $2 AND id = $3
RETURNING username, rol, account_name;