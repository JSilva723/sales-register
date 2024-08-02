package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomSaleOrder(t *testing.T, accountName string) int32 {
	user := createRandomUser(t, accountName)
	paymentRow := createRandomPayment(t, accountName)

	arg := CreateSaleOrderParams{
		AccountName: accountName,
		UserID:      user.ID,
		PaymentID:   paymentRow.ID,
	}

	orderID, err := testQueries.CreateSaleOrder(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, orderID)

	return orderID
}

func TestCreateSaleOrder(t *testing.T) {
	accountName := createRandomAccount(t)
	createRandomSaleOrder(t, accountName)
}
