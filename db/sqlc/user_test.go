package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/wizlif/simplebank/util"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	args := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomCurrency(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.Username, args.Username)
	require.Equal(t, user.HashedPassword, args.HashedPassword)
	require.Equal(t, user.FullName, args.FullName)
	require.Equal(t, user.Email, args.Email)

	require.NotZero(t, user.CreatedAt)
	require.True(t, user.PasswordChangedAt.IsZero())

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
}

func TestUpdateUserOnlyFullName(t *testing.T) {
	oldUser := createRandomUser(t)

	newFullName := util.RandomOwner()
	newUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		FullName: sql.NullString{String: newFullName, Valid: true},
		Username: oldUser.Username,
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.FullName, newUser.FullName)
	require.Equal(t, newUser.FullName, newFullName)
	require.Equal(t, newUser.Email, oldUser.Email)
	require.Equal(t, newUser.HashedPassword, oldUser.HashedPassword)
}

func TestUpdateUserOnlyEmail(t *testing.T) {
	oldUser := createRandomUser(t)

	newEmail := util.RandomEmail()
	newUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		Email:    sql.NullString{String: newEmail, Valid: true},
		Username: oldUser.Username,
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.Email, newUser.Email)
	require.Equal(t, newUser.Email, newEmail)
	require.Equal(t, newUser.FullName, oldUser.FullName)
	require.Equal(t, newUser.HashedPassword, oldUser.HashedPassword)
}

func TestUpdateUserOnlyPassword(t *testing.T) {
	oldUser := createRandomUser(t)

	newPassword := util.RandomString(6)
	newHashedPassword, err := util.HashPassword(newPassword)
	require.NoError(t,err)

	newUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		HashedPassword: sql.NullString{String: newHashedPassword, Valid: true},
		Username:       oldUser.Username,
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.HashedPassword, newUser.HashedPassword)
	require.Equal(t, newUser.HashedPassword, newHashedPassword)
	require.Equal(t, newUser.FullName, oldUser.FullName)
	require.Equal(t, newUser.Email, oldUser.Email)
}

func TestUpdateUserAllFields(t *testing.T) {
	oldUser := createRandomUser(t)

	newEmail := util.RandomEmail()
	newFullName := util.RandomOwner()
	newPassword := util.RandomString(6)
	newHashedPassword, err := util.HashPassword(newPassword)
	require.NoError(t,err)

	newUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		HashedPassword: sql.NullString{String: newHashedPassword, Valid: true},
		Email: sql.NullString{String: newEmail, Valid: true},
		FullName: sql.NullString{String: newFullName, Valid: true},
		Username:       oldUser.Username,
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.HashedPassword, newUser.HashedPassword)
	require.Equal(t, newUser.HashedPassword, newHashedPassword)
	require.NotEqual(t, oldUser.FullName, newUser.FullName)
	require.Equal(t, newUser.FullName, newFullName)
	require.NotEqual(t, oldUser.Email, newUser.Email)
	require.Equal(t, newUser.Email, newEmail)
}
