-- name: CreateCatalog :one
INSERT INTO cataloges (
    name, parent_id, sort_order
) VALUES (
             $1,$2,$3
         )
RETURNING *;

-- name: GetAccount :one
SELECT * FROM cataloges
WHERE id = $1 LIMIT 1;

-- name: GetCatalogesForUpdate :one
SELECT * FROM cataloges
WHERE id = $1 LIMIT 1
    FOR NO KEY UPDATE
;

-- name: ListCataloges :many
SELECT * FROM cataloges
ORDER BY id
LIMIT $1
    OFFSET $2;

-- name: DeleteCataloges :exec
DELETE FROM cataloges WHERE id = $1;
