-- name: CreateAdmin :one
INSERT INTO Admin (
    username, password, name
) VALUES (
             $1,$2,$3
         )
RETURNING *;

-- name: GetAdmin :one
SELECT * FROM Admin
WHERE id = $1 LIMIT 1;

-- name: DeleteAdmin :exec
DELETE FROM Admin WHERE id = $1;