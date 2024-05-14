package db

import (
	util "github.com/QuyenFunc/Goaudio/db/util"

	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomCatalog(t *testing.T) Cataloge {
	arg := CreateCatalogParams{
		Name:      util.RandomString(6),
		ParentID:  util.RamdomInt(0, 1000),
		SortOrder: int16(util.RamdomInt(1, 1000)),
	}

	catalog, err := testQueries.CreateCatalog(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, catalog)

	require.Equal(t, arg.Name, catalog.Name)
	require.Equal(t, arg.ParentID, catalog.ParentID)
	require.Equal(t, arg.SortOrder, catalog.SortOrder)

	require.NotZero(t, catalog.ID)

	return catalog
}

func TestCreateCatalog(t *testing.T) {
	createRandomCatalog(t)
}

func TestGetCatalog(t *testing.T) {
	catalog1 := createRandomCatalog(t)
	catalog2, err := testQueries.GetCatalog(context.Background(), catalog1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, catalog2)

	require.Equal(t, catalog1.Name, catalog2.Name)
	require.Equal(t, catalog1.ParentID, catalog2.ParentID)
	require.Equal(t, catalog1.SortOrder, catalog2.SortOrder)
}

func TestUpdateCatalog(t *testing.T) {
	catalog1 := createRandomCatalog(t)

	arg := UpdateCatalogParams{
		ID:        catalog1.ID,
		SortOrder: int16(util.RamdomInt(1, 1000)),
	}

	catalog2, err := testQueries.UpdateCatalog(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, catalog2)
	//cai nao cap nhat thi dung arg con cai nao k thi dung catalog
	require.Equal(t, arg.SortOrder, catalog2.SortOrder)
	require.Equal(t, catalog1.Name, catalog2.Name)
	require.Equal(t, catalog1.ParentID, catalog2.ParentID)
}
