package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) string {
	arg := CreateAccountParams{
		ID:   RandomInt(0, 10000),
		Name: randomString(6),
	}

	name, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, name)

	return name
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}
