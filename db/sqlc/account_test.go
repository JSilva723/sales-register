package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

// Create Test account and return id for other test
func createAcc() (string, error) {
	arg := CreateAccountParams{
		Name:   randomString(6),
		Status: "ACTIVE",
	}

	return testQueries.CreateAccount(context.Background(), arg)
}

func TestCreateAccount(t *testing.T) {
	name, err := createAcc()
	require.NoError(t, err)
	require.NotEmpty(t, name)
}
