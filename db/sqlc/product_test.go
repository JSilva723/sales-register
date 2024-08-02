package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomProduct(t *testing.T, accountName string) CreateProductRow {
	arg := CreateProductParams{
		AccountName: accountName,
		Name:        randomString(6),
	}

	productRow, err := testQueries.CreateProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, productRow)
	require.Equal(t, productRow.Name, arg.Name)

	return productRow
}

func TestCreateProduct(t *testing.T) {
	accountName := createRandomAccount(t)
	createRandomProduct(t, accountName)
}

func TestGetProducts(t *testing.T) {
	accountName := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomProduct(t, accountName)
	}

	arg := GetProductsParams{
		AccountName: accountName,
		Limit:       5,
		Offset:      5,
	}

	productRows, err := testQueries.GetProducts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, productRows, 5)

	for _, Product := range productRows {
		require.NotEmpty(t, Product)
	}
}

func TestChangeProductName(t *testing.T) {
	accountName := createRandomAccount(t)
	productRow := createRandomProduct(t, accountName)

	productRowUpdated, err := testQueries.ChangeProductName(
		context.Background(),
		ChangeProductNameParams{
			ID:   productRow.ID,
			Name: "Product name updated",
		},
	)
	require.NoError(t, err)
	require.NotEmpty(t, productRowUpdated)
	require.NotEqual(t, productRow.Name, productRowUpdated.Name)
}

func TestDeleteProduct(t *testing.T) {
	accountName := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomProduct(t, accountName)
	}

	arg := GetProductsParams{
		AccountName: accountName,
		Limit:       10,
		Offset:      1,
	}

	productRows, _ := testQueries.GetProducts(context.Background(), arg)

	for _, product := range productRows {
		err := testQueries.DeleteProduct(context.Background(), product.ID)
		require.NoError(t, err)
	}

	productRows, _ = testQueries.GetProducts(context.Background(), arg)
	require.Len(t, productRows, 0)
}
