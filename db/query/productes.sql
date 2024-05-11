-- name: CreateProducte :one
INSERT INTO productes (
    catalog_id, name, price,view,
    content,discount,image_link,image_list
) VALUES (
             $1,$2,$3,$4,$5,$6,$7,$8
         )
RETURNING *;

-- name: GetProducte :one
SELECT * FROM productes
WHERE catalog_id = $1 LIMIT 1;

-- name: GetProductesForUpdate :one
SELECT * FROM productes
WHERE catalog_id = $1 LIMIT 1
    FOR NO KEY UPDATE
;

-- name: ListProductes :many
SELECT * FROM productes
WHERE name = $1
ORDER BY catalog_id
LIMIT $2
OFFSET $3;


-- name: UpdateProductes :one
UPDATE productes
SET price = $3,
    view = $4,
    discount = $6
WHERE catalog_id = $1
RETURNING *;

-- name: DeleteProductes :exec
DELETE FROM productes WHERE catalog_id = $1;

