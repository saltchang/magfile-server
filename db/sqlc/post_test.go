package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/saltchang/magfile-server/util"
	"github.com/stretchr/testify/require"
)

func createRandomPost(t *testing.T, authorID int64) Post {
	params := CreatePostParams{
		SemanticID: util.GetRandomString(6) + "-" + util.GetRandomString(8),
		AuthorID:   authorID,
		Title:      util.GetRandomString(16),
		Abstract:   util.GetRandomString(30),
		Content:    util.GetRandomString(1000),
		Tags:       util.GetRandomStringArray(int(util.GetRandomInt(0, 5))),
		IsArchived: util.GetRandomBoolean(),
		UpdatedAt:  time.Now().UTC(),
	}

	user, err := testQueries.CreatePost(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, params.SemanticID, user.SemanticID)
	require.Equal(t, authorID, user.AuthorID)
	require.Equal(t, params.AuthorID, user.AuthorID)
	require.Equal(t, params.Title, user.Title)
	require.Equal(t, params.Abstract, user.Abstract)
	require.Equal(t, params.Content, user.Content)
	require.Equal(t, params.Tags, user.Tags)
	require.Equal(t, params.IsArchived, user.IsArchived)
	require.Equal(t, params.UpdatedAt, user.UpdatedAt)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreatePost(t *testing.T) {
	user := createRandomBlogUser(t)
	createRandomPost(t, user.ID)
}

func TestGetPostBySemanticID(t *testing.T) {
	user := createRandomBlogUser(t)
	post := createRandomPost(t, user.ID)

	params := GetPostBySemanticIDParams{
		AuthorID:   user.ID,
		SemanticID: post.SemanticID,
	}
	gotPost, err := testQueries.GetPostBySemanticID(context.Background(), params)

	require.NoError(t, err)
	require.NotEmpty(t, gotPost)

	require.Equal(t, gotPost.ID, post.ID)
	require.Equal(t, gotPost.AuthorID, user.ID)
	require.Equal(t, gotPost.AuthorID, params.AuthorID)
	require.Equal(t, gotPost.SemanticID, params.SemanticID)
	require.Equal(t, gotPost.SemanticID, post.SemanticID)
	require.Equal(t, gotPost.Title, post.Title)
	require.Equal(t, gotPost.Abstract, post.Abstract)
	require.Equal(t, gotPost.Content, post.Content)
	require.Equal(t, gotPost.Tags, post.Tags)
	require.Equal(t, gotPost.IsArchived, post.IsArchived)
	require.Equal(t, gotPost.UpdatedAt, post.UpdatedAt)

	require.NotZero(t, gotPost.ID)
	require.NotZero(t, gotPost.CreatedAt)

	require.WithinDuration(t, gotPost.CreatedAt, post.CreatedAt, time.Second)
}

func TestGetPostByPostID(t *testing.T) {
	user := createRandomBlogUser(t)
	post := createRandomPost(t, user.ID)

	gotPost, err := testQueries.GetPostByPostID(context.Background(), post.ID)

	require.NoError(t, err)
	require.NotEmpty(t, gotPost)

	require.Equal(t, gotPost.ID, post.ID)
	require.Equal(t, gotPost.AuthorID, user.ID)
	require.Equal(t, gotPost.AuthorID, post.AuthorID)
	require.Equal(t, gotPost.SemanticID, post.SemanticID)
	require.Equal(t, gotPost.Title, post.Title)
	require.Equal(t, gotPost.Abstract, post.Abstract)
	require.Equal(t, gotPost.Content, post.Content)
	require.Equal(t, gotPost.Tags, post.Tags)
	require.Equal(t, gotPost.IsArchived, post.IsArchived)
	require.Equal(t, gotPost.UpdatedAt, post.UpdatedAt)

	require.NotZero(t, gotPost.ID)
	require.NotZero(t, gotPost.CreatedAt)

	require.WithinDuration(t, gotPost.CreatedAt, post.CreatedAt, time.Second)
}

func TestGetAllPostFromAuthor(t *testing.T) {
	testPostLen := 5

	user := createRandomBlogUser(t)
	for i := 0; i < testPostLen; i++ {
		createRandomPost(t, user.ID)
	}

	allPosts, err := testQueries.GetAllPostFromAuthor(context.Background(), user.ID)

	require.NoError(t, err)
	require.NotEmpty(t, allPosts)

	for i := 0; i < len(allPosts); i++ {
		require.NotZero(t, allPosts[i].ID)
		require.NotZero(t, allPosts[i].CreatedAt)
		require.Equal(t, allPosts[i].AuthorID, user.ID)
	}

	require.GreaterOrEqual(t, len(allPosts), testPostLen)
}

func TestGetAllPostFromAuthorWhoHasNoAnyPost(t *testing.T) {
	user := createRandomBlogUser(t)
	allPosts, err := testQueries.GetAllPostFromAuthor(context.Background(), user.ID)

	require.NoError(t, err)
	require.Empty(t, allPosts)
}

func TestUpdatePost(t *testing.T) {
	user := createRandomBlogUser(t)
	oldPost := createRandomPost(t, user.ID)

	newParams := UpdatePostParams{
		ID:         oldPost.ID,
		SemanticID: util.GetRandomString(6) + "-" + util.GetRandomString(8),
		AuthorID:   oldPost.AuthorID,
		Title:      oldPost.Title,
		Abstract:   util.GetRandomString(30),
		Content:    util.GetRandomString(1000),
		Tags:       util.GetRandomStringArray(int(util.GetRandomInt(0, 5))),
		IsArchived: util.GetRandomBoolean(),
		UpdatedAt:  time.Now().UTC(),
	}

	newPost, err := testQueries.UpdatePost(context.Background(), newParams)
	require.NoError(t, err)
	require.NotEmpty(t, newPost)

	require.Equal(t, oldPost.ID, newPost.ID)
	require.Equal(t, newPost.ID, newParams.ID)

	require.Equal(t, newPost.SemanticID, newParams.SemanticID)
	require.NotEqual(t, newPost.SemanticID, oldPost.SemanticID)

	require.Equal(t, newPost.AuthorID, newParams.AuthorID)
	require.Equal(t, newPost.AuthorID, oldPost.AuthorID) // authorID should be the same

	require.Equal(t, newPost.Title, newParams.Title)
	require.Equal(t, newPost.Title, oldPost.Title) // title should be the same

	require.Equal(t, newPost.Abstract, newParams.Abstract)
	require.NotEqual(t, newPost.Abstract, oldPost.Abstract)

	require.Equal(t, newPost.Content, newParams.Content)
	require.NotEqual(t, newPost.Content, oldPost.Content)

	require.Equal(t, newPost.Tags, newParams.Tags)

	require.Equal(t, newPost.IsArchived, newParams.IsArchived)

	require.WithinDuration(t, newPost.UpdatedAt, newParams.UpdatedAt, time.Second)
	require.WithinDuration(t, oldPost.CreatedAt, newPost.CreatedAt, time.Second)
}

func TestDeletePost(t *testing.T) {
	user := createRandomBlogUser(t)
	post := createRandomPost(t, user.ID)
	err := testQueries.DeletePost(context.Background(), post.ID)

	require.NoError(t, err)

	newPost, err := testQueries.GetPostByPostID(context.Background(), post.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, newPost)

	params := GetPostBySemanticIDParams{
		AuthorID:   user.ID,
		SemanticID: post.SemanticID,
	}
	newPost, err = testQueries.GetPostBySemanticID(context.Background(), params)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, newPost)
}
