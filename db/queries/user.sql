-- name: CreateUser :one
INSERT INTO users (
    username,
    password,
    account_name,
    rol
) VALUES (
    $1, $2, $3, $4
) RETURNING id, account_name, username, rol, created_at, updated_at;

-- name: GetUser :one
SELECT username, rol, created_at, updated_at
FROM users
WHERE id = $1 AND is_active = true
LIMIT 1;

-- name: GetUsers :many
SELECT id, username, rol
FROM users
WHERE account_name = $1 AND is_active = true
LIMIT $2
OFFSET $3;

-- name: ChangePassword :exec
UPDATE users
SET password = $1, updated_at = (now())
WHERE id = $2;

-- name: ChangeRol :one
UPDATE users
SET rol = $1, updated_at =  (now())
WHERE id = $2
RETURNING username, rol;

-- name: DeleteUser :exec
UPDATE users
SET is_active = false, updated_at = (now())
WHERE id = $1;