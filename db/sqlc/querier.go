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
	DeleteAdmin(ctx context.Context, id int64) error
	DeleteProductes(ctx context.Context, id int64) error
	DeleteUser(ctx context.Context, id int64) error
	GetAdmin(ctx context.Context, id int64) (Admin, error)
	GetCatalog(ctx context.Context, id int64) (Cataloge, error)
	GetCatalogesForUpdate(ctx context.Context, id int64) (Cataloge, error)
	GetProducte(ctx context.Context, id int64) (Producte, error)
	GetProductesForUpdate(ctx context.Context, id int64) (Producte, error)
	GetUser(ctx context.Context, name string) (User, error)
	ListProductes(ctx context.Context, arg ListProductesParams) ([]Producte, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	UpdateCatalog(ctx context.Context, arg UpdateCatalogParams) (Cataloge, error)
	UpdateProductes(ctx context.Context, arg UpdateProductesParams) (Producte, error)
}

var _ Querier = (*Queries)(nil)
