package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	accountID, _ := createAcc()

	arg := CreateUserParams{
		Username:  randomString(6),
		Password:  randomString(6),
		AccountID: accountID,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, user.Username, arg.Username)
	require.Equal(t, user.AccountID, arg.AccountID)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)
	require.Equal(t, user.CreatedAt, user.UpdatedAt)
}
