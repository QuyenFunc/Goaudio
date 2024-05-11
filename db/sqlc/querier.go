// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
)

type Querier interface {
	CreateAdmin(ctx context.Context, arg CreateAdminParams) (Admin, error)
	CreateCatalog(ctx context.Context, arg CreateCatalogParams) (Cataloge, error)
	CreateProducte(ctx context.Context, arg CreateProducteParams) (Producte, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteAdmin(ctx context.Context, id int32) error
	DeleteCataloges(ctx context.Context, id int32) error
	// -- name: UpdateProductes :one
	// UPDATE productes
	// SET balance = $2
	// WHERE id = $1
	// RETURNING *;
	DeleteProductes(ctx context.Context, id int32) error
	DeleteUser(ctx context.Context, id int32) error
	GetAccount(ctx context.Context, id int32) (Cataloge, error)
	GetAdmin(ctx context.Context, id int32) (Admin, error)
	GetCatalogesForUpdate(ctx context.Context, id int32) (Cataloge, error)
	GetProducte(ctx context.Context, id int32) (Producte, error)
	GetProductesForUpdate(ctx context.Context, id int32) (Producte, error)
	GetUser(ctx context.Context, id int32) (User, error)
	ListCataloges(ctx context.Context, arg ListCatalogesParams) ([]Cataloge, error)
	ListProductes(ctx context.Context, arg ListProductesParams) ([]Producte, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
}

var _ Querier = (*Queries)(nil)
