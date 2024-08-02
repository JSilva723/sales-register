package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomPayment(t *testing.T, accountName string) CreatePaymentRow {
	arg := CreatePaymentParams{
		AccountName: accountName,
		Name:        randomString(6),
	}

	paymentRow, err := testQueries.CreatePayment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, paymentRow)
	require.Equal(t, paymentRow.Name, arg.Name)

	return paymentRow
}

func TestCreatePayment(t *testing.T) {
	accountName := createRandomAccount(t)
	createRandomPayment(t, accountName)
}

func TestGetPayments(t *testing.T) {
	accountName := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomPayment(t, accountName)
	}

	arg := GetPaymentsParams{
		AccountName: accountName,
		Limit:       5,
		Offset:      5,
	}

	paymentRows, err := testQueries.GetPayments(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, paymentRows, 5)

	for _, payment := range paymentRows {
		require.NotEmpty(t, payment)
	}
}

func TestChangePaymnetName(t *testing.T) {
	accountName := createRandomAccount(t)
	paymentRow := createRandomPayment(t, accountName)

	paymentRowUpdated, err := testQueries.ChangePaymentName(
		context.Background(),
		ChangePaymentNameParams{
			ID:   paymentRow.ID,
			Name: "Payment name updated",
		},
	)
	require.NoError(t, err)
	require.NotEmpty(t, paymentRowUpdated)
	require.NotEqual(t, paymentRow.Name, paymentRowUpdated.Name)
}

func TestDeletePayment(t *testing.T) {
	accountName := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomPayment(t, accountName)
	}

	arg := GetPaymentsParams{
		AccountName: accountName,
		Limit:       10,
		Offset:      1,
	}

	paymentRows, _ := testQueries.GetPayments(context.Background(), arg)

	for _, payment := range paymentRows {
		err := testQueries.DeletePayment(context.Background(), payment.ID)
		require.NoError(t, err)
	}

	paymentRows, _ = testQueries.GetPayments(context.Background(), arg)
	require.Len(t, paymentRows, 0)
}
