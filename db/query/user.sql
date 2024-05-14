-- name: CreateUser :one
INSERT INTO users (
    name, email, phone, address, hashed_password
) VALUES (
             $1,$2,$3,$4,$5
         )
RETURNING *;

-- name: GetUser :one
SELECT * FROM Users
WHERE name = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM Users
ORDER BY id
LIMIT $1
    OFFSET $2;

-- name: DeleteUser :exec
DELETE FROM Users WHERE id = $1;