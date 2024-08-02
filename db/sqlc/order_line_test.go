package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomOrderLine(t *testing.T, accountName string) {
	saleOrderID := createRandomSaleOrder(t, accountName)
	productRow := createRandomProduct(t, accountName)

	arg := CreateOrderLineParams{
		AccountName: accountName,
		SaleID:      saleOrderID,
		Ammount:     RandomInt(0, 1000),
		ProductID:   productRow.ID,
	}

	err := testQueries.CreateOrderLine(context.Background(), arg)
	require.NoError(t, err)
}

func TestCreateOrderLine(t *testing.T) {
	accountName := createRandomAccount(t)
	createRandomOrderLine(t, accountName)
}
