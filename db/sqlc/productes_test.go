package db

import (
	util "github.com/QuyenFunc/Goaudio/db/util"

	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomProducte(t *testing.T) Producte {
	arg := CreateProducteParams{
		CatalogID: util.RamdomInt(1, 100),
		Name:      util.RandomString(5),
		Price:     util.RandomPrice(),
		View:      util.RamdomInt(1, 1000),
		Discount:  util.RamdomDiscount(),
	}

	producte, err := testQueries.CreateProducte(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, producte)

	require.Equal(t, arg.CatalogID, producte.CatalogID)
	require.Equal(t, arg.Price, producte.Price)
	require.Equal(t, arg.Name, producte.Name)
	require.Equal(t, arg.View, producte.View)
	require.Equal(t, arg.Discount, producte.Discount)

	require.NotZero(t, producte.ID)

	return producte
}

func TestCreateProducte(t *testing.T) {
	createRandomProducte(t)
}

func TestGetProducte(t *testing.T) {
	producte1 := createRandomProducte(t)
	producte2, err := testQueries.GetProducte(context.Background(), producte1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, producte2)

	require.Equal(t, producte1.CatalogID, producte2.CatalogID)
	require.Equal(t, producte1.Name, producte2.Name)
	require.Equal(t, producte1.Price, producte2.Price)
	require.Equal(t, producte1.View, producte2.View)

	require.WithinDuration(t, producte1.CreatedAt, producte2.CreatedAt, time.Second)
}

func TestUpdateProduct(t *testing.T) {
	product1 := createRandomProducte(t)

	arg := UpdateProductesParams{
		ID:        product1.ID,
		CatalogID: util.RamdomInt(5, 1000),
		Price:     util.RandomPrice(),
		View:      util.RandomView(),
		Discount:  util.RamdomDiscount(),
	}

	product2, err := testQueries.UpdateProductes(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product2)
	//cai nao cap nhat thi dung arg con cai nao k thi dung product
	require.Equal(t, product1.ID, product2.ID)
	require.Equal(t, arg.CatalogID, product2.CatalogID)
	require.Equal(t, arg.Price, product2.Price)
	require.Equal(t, arg.View, product2.View)
	require.Equal(t, arg.Discount, product2.Discount)
	require.WithinDuration(t, product1.CreatedAt, product2.CreatedAt, time.Second)
}
func TestDeleteProduct(t *testing.T) {
	product1 := createRandomProducte(t)

	err := testQueries.DeleteProductes(context.Background(), product1.ID)
	require.NoError(t, err)

	product2, err := testQueries.GetProducte(context.Background(), product1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, product2)
}

func TestListProducts(t *testing.T) {
	var LastProduct Producte
	for i := 0; i < 10; i++ {
		LastProduct = createRandomProducte(t)
	}
	arg := ListProductesParams{
		Name:   LastProduct.Name,
		Limit:  5,
		Offset: 0,
	}

	productes, err := testQueries.ListProductes(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, productes)

	for _, product := range productes {
		require.NotEmpty(t, product)
		require.Equal(t, LastProduct.Name, product.Name)
	}
}
