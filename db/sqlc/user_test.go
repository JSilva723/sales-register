package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAdminUser(t *testing.T) {
	accountName, _ := createAcc()

	arg := CreateUserParams{
		Username:    randomString(6),
		Password:    randomString(6),
		Rol:         "ADMIN",
		AccountName: accountName,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, user.Username, arg.Username)
	require.Equal(t, user.AccountName, arg.AccountName)
	require.Equal(t, user.Rol, arg.Rol)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)
	require.Equal(t, user.CreatedAt, user.UpdatedAt)
}
