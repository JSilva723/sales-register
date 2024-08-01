package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomUser() (CreateUserParams, User, error) {
	accountName, _ := createAcc()

	arg := CreateUserParams{
		Username:    randomString(6),
		Password:    randomString(6),
		Rol:         "ADMIN",
		AccountName: accountName,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	return arg, user, err
}

func TestCreateAdminUser(t *testing.T) {
	arg, user, err := createRandomUser()
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, user.Username, arg.Username)
	require.Equal(t, user.AccountName, arg.AccountName)
	require.Equal(t, user.Rol, arg.Rol)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)
	require.Equal(t, user.CreatedAt, user.UpdatedAt)
}

func TestGetUser(t *testing.T) {
	_, user, _ := createRandomUser()
	userFounded, err := testQueries.GetUser(context.Background(), GetUserParams{
		ID:          user.ID,
		AccountName: user.AccountName,
	})
	require.NoError(t, err)
	require.Equal(t, user.Rol, userFounded.Rol)
	require.Equal(t, user.Username, userFounded.Username)
	require.Equal(t, user.AccountName, userFounded.AccountName)
	require.Equal(t, user.CreatedAt, userFounded.UpdatedAt)
	require.Equal(t, user.UpdatedAt, userFounded.UpdatedAt)
}

func TestChangePassword(t *testing.T) {
	_, user, _ := createRandomUser()

	err := testQueries.ChangePassword(context.Background(), ChangePasswordParams{
		Password:    "NEW-PASS",
		ID:          user.ID,
		AccountName: user.AccountName,
	})
	require.NoError(t, err)

	userFounded, _ := testQueries.GetUser(context.Background(), GetUserParams{
		ID:          user.ID,
		AccountName: user.AccountName,
	})

	require.NotEqual(t, userFounded.CreatedAt, userFounded.UpdatedAt)
}

func TestChangeRol(t *testing.T) {
	_, user, _ := createRandomUser()

	userArg := ChangeRolParams{
		Rol:         "EMPLOYEE",
		ID:          user.ID,
		AccountName: user.AccountName,
	}

	rolRow, err := testQueries.ChangeRol(context.Background(), userArg)
	require.NoError(t, err)
	require.NotEmpty(t, rolRow)
	require.Equal(t, user.Username, rolRow.Username)
	require.Equal(t, userArg.Rol, rolRow.Rol)
	require.Equal(t, user.AccountName, rolRow.AccountName)

	userFounded, _ := testQueries.GetUser(context.Background(), GetUserParams{
		ID:          user.ID,
		AccountName: user.AccountName,
	})

	require.NotEqual(t, userFounded.CreatedAt, userFounded.UpdatedAt)
}
