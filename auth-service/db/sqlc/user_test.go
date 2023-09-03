package sqlc

import (
	"chat-app/auth-service/util"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:     util.RandomString(10),
		PasswordHash: util.RandomString(10),
		Role:         "user",
		Email:        util.RandomString(10) + "@test.com",
		Verified:     false,
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.PasswordHash, user.PasswordHash)
	require.Equal(t, arg.Role, user.Role)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Verified, user.Verified)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	return user
}
func TestUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUserByUsername(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.PasswordHash, user2.PasswordHash)
	require.Equal(t, user1.Role, user2.Role)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Verified, user2.Verified)
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)

	arg := UpdateUserParams{
		ID:           user1.ID,
		Username:     util.RandomString(10),
		PasswordHash: util.RandomString(10),
		Role:         "user",
		Email:        util.RandomString(10) + "@test.com",
		Verified:     false,
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, arg.Username, user2.Username)
	require.Equal(t, arg.PasswordHash, user2.PasswordHash)
	require.Equal(t, arg.Role, user2.Role)
	require.Equal(t, arg.Email, user2.Email)
	require.Equal(t, arg.Verified, user2.Verified)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.Error(t, err)
	require.Empty(t, user2)
}
