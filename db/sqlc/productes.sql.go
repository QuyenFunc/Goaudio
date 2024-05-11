// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: productes.sql

package db

import (
	"context"
)

const createProducte = `-- name: CreateProducte :one
INSERT INTO productes (
    catalog_id, name, price,created,view,
    content,discount,image_link,image_list
) VALUES (
             $1,$2,$3, $4,$5,$6, $7,$8,$9
         )
RETURNING id, catalog_id, name, price, content, discount, image_link, image_list, created, view
`

type CreateProducteParams struct {
	CatalogID int32  `json:"catalog_id"`
	Name      string `json:"name"`
	Price     int32  `json:"price"`
	Created   int32  `json:"created"`
	View      int32  `json:"view"`
	Content   string `json:"content"`
	Discount  int32  `json:"discount"`
	ImageLink string `json:"image_link"`
	ImageList string `json:"image_list"`
}

func (q *Queries) CreateProducte(ctx context.Context, arg CreateProducteParams) (Producte, error) {
	row := q.db.QueryRowContext(ctx, createProducte,
		arg.CatalogID,
		arg.Name,
		arg.Price,
		arg.Created,
		arg.View,
		arg.Content,
		arg.Discount,
		arg.ImageLink,
		arg.ImageList,
	)
	var i Producte
	err := row.Scan(
		&i.ID,
		&i.CatalogID,
		&i.Name,
		&i.Price,
		&i.Content,
		&i.Discount,
		&i.ImageLink,
		&i.ImageList,
		&i.Created,
		&i.View,
	)
	return i, err
}

const deleteProductes = `-- name: DeleteProductes :exec

DELETE FROM productes WHERE id = $1
`

// -- name: UpdateProductes :one
// UPDATE productes
// SET balance = $2
// WHERE id = $1
// RETURNING *;
func (q *Queries) DeleteProductes(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteProductes, id)
	return err
}

const getProducte = `-- name: GetProducte :one
SELECT id, catalog_id, name, price, content, discount, image_link, image_list, created, view FROM productes
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProducte(ctx context.Context, id int32) (Producte, error) {
	row := q.db.QueryRowContext(ctx, getProducte, id)
	var i Producte
	err := row.Scan(
		&i.ID,
		&i.CatalogID,
		&i.Name,
		&i.Price,
		&i.Content,
		&i.Discount,
		&i.ImageLink,
		&i.ImageList,
		&i.Created,
		&i.View,
	)
	return i, err
}

const getProductesForUpdate = `-- name: GetProductesForUpdate :one
SELECT id, catalog_id, name, price, content, discount, image_link, image_list, created, view FROM productes
WHERE id = $1 LIMIT 1
    FOR NO KEY UPDATE
`

func (q *Queries) GetProductesForUpdate(ctx context.Context, id int32) (Producte, error) {
	row := q.db.QueryRowContext(ctx, getProductesForUpdate, id)
	var i Producte
	err := row.Scan(
		&i.ID,
		&i.CatalogID,
		&i.Name,
		&i.Price,
		&i.Content,
		&i.Discount,
		&i.ImageLink,
		&i.ImageList,
		&i.Created,
		&i.View,
	)
	return i, err
}

const listProductes = `-- name: ListProductes :many
SELECT id, catalog_id, name, price, content, discount, image_link, image_list, created, view FROM productes
ORDER BY id
LIMIT $1
    OFFSET $2
`

type ListProductesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListProductes(ctx context.Context, arg ListProductesParams) ([]Producte, error) {
	rows, err := q.db.QueryContext(ctx, listProductes, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Producte{}
	for rows.Next() {
		var i Producte
		if err := rows.Scan(
			&i.ID,
			&i.CatalogID,
			&i.Name,
			&i.Price,
			&i.Content,
			&i.Discount,
			&i.ImageLink,
			&i.ImageList,
			&i.Created,
			&i.View,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}