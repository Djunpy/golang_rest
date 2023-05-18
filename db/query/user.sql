-- name: CreateUser :one
INSERT INTO users (
    name,
    email,
    verified,
    password,
    role
) VALUES (
             $1, $2, $3, $4, $5
         )
    RETURNING *;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
    LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET name = $2,
    email = $3,
    verified = $4,
    password = $5,
    role = $6
WHERE id = $1
    RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;