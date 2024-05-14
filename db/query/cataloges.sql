-- name: CreateCatalog :one
INSERT INTO cataloges (
    name, parent_id, sort_order
) VALUES (
             $1,$2,$3
         )
RETURNING *;

-- name: GetCatalog :one
SELECT * FROM cataloges
WHERE id = $1 LIMIT 1;

-- name: UpdateCatalog :one
UPDATE cataloges
SET sort_order = $2
WHERE id = $1
RETURNING *;

-- name: GetCatalogesForUpdate :one
SELECT * FROM cataloges
WHERE id = $1 LIMIT 1
    FOR NO KEY UPDATE
;

