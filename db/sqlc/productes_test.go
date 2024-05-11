package db

import (
	util "GoAudio/db/util"
	"context"
	"go/printer"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomProducte(t *testing.T) Producte {
	arg := CreateProducteParams{
		Name:  util.RandomString(5),
		Price: int32(util.RandomPrice()),
		View:  int32(util.RandomView()),
		Discount: util.RamdomDiscount(),
	}

	producte, err := testQueries.CreateProducte(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, producte)

	require.Equal(t, arg.Price, producte.Price)
	require.Equal(t, arg.Name, producte.Name)
	require.Equal(t, arg.View, producte.View)
	require.Equal(t, arg.Discount, producte.Discount)

	require.NotZero(t, producte.ID)

	return producte
}

func TestCreatePeoducte(t *testing.T) {
	createRandomProducte(t)
}

func TestGetProducte(t *testing.T){
	producte1 := createRandomProducte(t)
	producte2, err := testQueries.CreateProducte(context.Background(), producte1.CatalogID)

	require.NoError(t, err)
	require.NotEmpty(t, producte2)

	require.Equal(t, producte1.CatalogID, producte2.CatalogID)
	require.Equal(t, producte1.Name, producte2.Name)
	require.Equal(t, producte1.Price, producte2.Price)
	require.Equal(t, producte1.View, producte2.View)

	require.WithinDuration(t, producte1.Created)
}

func TestUpdateProduct(t *testing.T){
	product1 := createRandomProducte(t)

	arg := Upa
}