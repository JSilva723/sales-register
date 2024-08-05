package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOrderTx(t *testing.T) {
	store := NewStore(testDB)

	accountName := createRandomAccount(t)
	user := createRandomUser(t, accountName)
	payment := createRandomPayment(t, accountName)
	product1 := createRandomProduct(t, accountName)
	product2 := createRandomProduct(t, accountName)
	lines := []Line{
		{Ammount: RandomInt(0, 1000), ProductID: product1.ID},
		{Ammount: RandomInt(0, 1000), ProductID: product2.ID},
		{Ammount: RandomInt(0, 1000), ProductID: product1.ID},
		{Ammount: RandomInt(0, 1000), ProductID: product1.ID},
	}

	n := 5
	errs := make(chan error)
	results := make(chan OrderTxResult)
	for i := 0; i < n; i++ {
		go func() {
			result, err := store.OrderTx(context.Background(), OrderTxParams{
				AccountName: accountName,
				UserID:      user.ID,
				PaymentID:   payment.ID,
				Lines:       lines,
			})

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
		result := <-results
		require.NotEmpty(t, result)
	}
}
