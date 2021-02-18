package test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	db "github.com/saltchang/magfile-server/db/sqlc"
	"github.com/saltchang/magfile-server/util"
	"github.com/stretchr/testify/require"
)

const testSalt string = "testSalt_123+456/abcd"

func createRandomBlogUser(t *testing.T) db.BlogUser {
	params := db.CreateBlogUserParams{
		Username:        util.GetRandomString(8),
		Email:           util.GetRandomEmail(),
		FullName:        util.GetRandomString(8),
		Gender:          util.GetRandomStringOption(util.GetGenderList()),
		CurrentLocation: util.GetRandomString(6) + ", " + util.GetRandomString(6),
		PasswordHash:    util.GetPasswordHash(util.GetRandomString(16), testSalt),
		LoginedAt:       time.Now().UTC(),
	}

	user, err := testQueries.CreateBlogUser(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, params.Username, user.Username)
	require.Equal(t, params.Email, user.Email)
	require.Equal(t, params.FullName, user.FullName)
	require.Equal(t, params.Gender, user.Gender)
	require.Equal(t, params.CurrentLocation, user.CurrentLocation)
	require.Equal(t, params.PasswordHash, user.PasswordHash)
	require.Equal(t, params.LoginedAt, user.LoginedAt)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateBlogUser(t *testing.T) {
	createRandomBlogUser(t)
}

func TestGetBlogUser(t *testing.T) {
	user1 := createRandomBlogUser(t)
	user2, err := testQueries.GetBlogUser(context.Background(), user1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Gender, user2.Gender)
	require.Equal(t, user1.CurrentLocation, user2.CurrentLocation)
	require.Equal(t, user1.PasswordHash, user2.PasswordHash)
	require.WithinDuration(t, user1.LoginedAt, user2.LoginedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateBlogUser(t *testing.T) {
	oldUser := createRandomBlogUser(t)

	newParams := db.UpdateBlogUserParams{
		ID:              oldUser.ID,
		Username:        util.GetRandomString(8),
		Email:           util.GetRandomEmail(),
		FullName:        oldUser.FullName,
		Gender:          util.GetRandomStringOption(util.GetGenderList()),
		CurrentLocation: util.GetRandomString(6) + ", " + util.GetRandomString(6),
		PasswordHash:    util.GetPasswordHash(util.GetRandomString(16), testSalt),
		LoginedAt:       time.Now().UTC(),
	}

	newUser, err := testQueries.UpdateBlogUser(context.Background(), newParams)
	require.NoError(t, err)
	require.NotEmpty(t, newUser)

	require.Equal(t, oldUser.ID, newUser.ID)
	require.Equal(t, newUser.ID, newParams.ID)

	require.Equal(t, newUser.Username, newParams.Username)
	require.NotEqual(t, newUser.Username, oldUser.Username)

	require.Equal(t, newUser.Email, newParams.Email)
	require.NotEqual(t, newUser.Email, oldUser.Email)

	require.Equal(t, newUser.FullName, newParams.FullName)
	require.Equal(t, newUser.FullName, oldUser.FullName) // only fullname should be the same

	require.Equal(t, newUser.Gender, newParams.Gender)

	require.Equal(t, newUser.CurrentLocation, newParams.CurrentLocation)
	require.NotEqual(t, newUser.CurrentLocation, oldUser.CurrentLocation)

	require.Equal(t, newUser.PasswordHash, newParams.PasswordHash)
	require.NotEqual(t, newUser.PasswordHash, oldUser.PasswordHash)

	require.WithinDuration(t, newUser.LoginedAt, newParams.LoginedAt, time.Second)
	require.WithinDuration(t, oldUser.CreatedAt, newUser.CreatedAt, time.Second)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomBlogUser(t)
	err := testQueries.DeleteBlogUser(context.Background(), user1.ID)

	require.NoError(t, err)

	user2, err := testQueries.GetBlogUser(context.Background(), user1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}
