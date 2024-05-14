package db

import (
	util "github.com/QuyenFunc/Goaudio/db/util"

	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashPassWord, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Name:           util.RandomString(7),
		Email:          util.RandomEmail(),
		Phone:          util.RandomPhoneNumber(),
		Address:        util.RandomString(10),
		HashedPassword: hashPassWord,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Phone, user.Phone)
	require.Equal(t, arg.Address, user.Address)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)

	require.NotZero(t, user.ID)

	return user
}

func TestCreatePeoducte(t *testing.T) {
	createRandomProducte(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Name)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Name, user2.Name)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Phone, user2.Phone)
	require.Equal(t, user1.Address, user2.Address)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)

	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
}
