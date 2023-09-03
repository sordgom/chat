package sqlc

import (
	"chat-app/auth-service/util"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createTestSession(t *testing.T) Session {
	user := createRandomUser(t)
	arg := CreateSessionParams{
		ID:           util.RandomUUID(),
		Username:     user.Username,
		RefreshToken: util.RandomString(10),
		UserAgent:    util.RandomString(10),
		ClientIp:     util.RandomString(10),
		IsBlocked:    false,
		ExpiresAt:    util.RandomTime(),
	}
	session, err := testQueries.CreateSession(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, session)

	require.Equal(t, arg.Username, session.Username)
	require.Equal(t, arg.RefreshToken, session.RefreshToken)
	require.Equal(t, arg.UserAgent, session.UserAgent)
	require.Equal(t, arg.ClientIp, session.ClientIp)
	require.Equal(t, arg.IsBlocked, session.IsBlocked)

	require.NotZero(t, session.ID)
	require.NotZero(t, session.CreatedAt)
	return session
}

func TestSession(t *testing.T) {
	session := createTestSession(t)
	session2, err := testQueries.GetSession(context.Background(), session.ID)

	require.NoError(t, err)
	require.NotEmpty(t, session2)

	require.Equal(t, session.ID, session2.ID)
	require.Equal(t, session.Username, session2.Username)
	require.Equal(t, session.RefreshToken, session2.RefreshToken)
	require.Equal(t, session.UserAgent, session2.UserAgent)
	require.Equal(t, session.ClientIp, session2.ClientIp)
	require.Equal(t, session.IsBlocked, session2.IsBlocked)
	require.Equal(t, session.ExpiresAt, session2.ExpiresAt)
}
