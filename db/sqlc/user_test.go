package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T, accountName string) User {
	arg := CreateUserParams{
		Username:    randomString(8),
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

	return user
}

func TestCreateAdminUser(t *testing.T) {
	accountName := createRandomAccount(t)
	createRandomUser(t, accountName)
}

func TestGetUser(t *testing.T) {
	accountName := createRandomAccount(t)
	user := createRandomUser(t, accountName)
	userFounded, err := testQueries.GetUser(context.Background(), user.ID)
	require.NoError(t, err)
	require.Equal(t, user.Rol, userFounded.Rol)
	require.Equal(t, user.Username, userFounded.Username)
	require.Equal(t, user.CreatedAt, userFounded.UpdatedAt)
	require.Equal(t, user.UpdatedAt, userFounded.UpdatedAt)
}

func TestGetUsers(t *testing.T) {
	accountName := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomUser(t, accountName)
	}

	arg := GetUsersParams{
		AccountName: accountName,
		Limit:       5,
		Offset:      5,
	}

	users, err := testQueries.GetUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, userRow := range users {
		require.NotEmpty(t, userRow)
	}
}

func TestChangePassword(t *testing.T) {
	accountName := createRandomAccount(t)
	user := createRandomUser(t, accountName)

	err := testQueries.ChangePassword(context.Background(), ChangePasswordParams{
		Password: "NEW-PASS",
		ID:       user.ID,
	})
	require.NoError(t, err)

	userFounded, _ := testQueries.GetUser(context.Background(), user.ID)
	require.NotEqual(t, userFounded.CreatedAt, userFounded.UpdatedAt)
}

func TestChangeRol(t *testing.T) {
	accountName := createRandomAccount(t)
	user := createRandomUser(t, accountName)

	userArg := ChangeRolParams{Rol: "EMPLOYEE", ID: user.ID}

	rolRow, err := testQueries.ChangeRol(context.Background(), userArg)
	require.NoError(t, err)
	require.NotEmpty(t, rolRow)
	require.Equal(t, user.Username, rolRow.Username)
	require.Equal(t, userArg.Rol, rolRow.Rol)

	userFounded, _ := testQueries.GetUser(context.Background(), user.ID)
	require.NotEqual(t, userFounded.CreatedAt, userFounded.UpdatedAt)
}

func TestDeleteUser(t *testing.T) {
	accountName := createRandomAccount(t)
	user := createRandomUser(t, accountName)

	err := testQueries.DeleteUser(context.Background(), user.ID)
	require.NoError(t, err)

	userFounded, err := testQueries.GetUser(context.Background(), user.ID)
	require.Error(t, err)
	require.Empty(t, userFounded)
}
