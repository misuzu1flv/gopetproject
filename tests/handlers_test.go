package tests

import (
	"homework-8/internal/pkg/core"
	"homework-8/internal/pkg/repository/postgresql"
	"homework-8/internal/pkg/sender"
	test_consts "homework-8/tests/consts"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreatePostSuccess(t *testing.T) {
	db.SetUp(t, test_consts.PostRequest)
	defer db.TearDown()
	//arrange
	server := core.NewServer(postgresql.NewPostRepo(db.DB), postgresql.NewCommentRepo(db.DB), sender.NewKafkaSender(pr, test_consts.Topic))
	code := server.CreatePost(test_consts.Ctx, test_consts.PostRequest)

	//assert
	require.Equal(t, code, http.StatusOK)
}

func TestCreatePostConflict(t *testing.T) {
	db.SetUp(t, test_consts.PostRequest)
	defer db.TearDown()
	//arrange
	server := core.NewServer(postgresql.NewPostRepo(db.DB), postgresql.NewCommentRepo(db.DB), sender.NewKafkaSender(pr, test_consts.Topic))
	code1 := server.CreatePost(test_consts.Ctx, test_consts.PostRequest)
	code2 := server.CreatePost(test_consts.Ctx, test_consts.PostRequest)

	//assert
	assert.Equal(t, code1, http.StatusOK)
	require.Equal(t, code2, http.StatusConflict)
}

func TestDeletePostSucces(t *testing.T) {
	db.SetUp(t, test_consts.PostRequest)
	defer db.TearDown()
	//arrange
	server := core.NewServer(postgresql.NewPostRepo(db.DB), postgresql.NewCommentRepo(db.DB), sender.NewKafkaSender(pr, test_consts.Topic))
	code1 := server.CreatePost(test_consts.Ctx, test_consts.PostRequest)

	code2 := server.DeletePost(test_consts.Ctx, test_consts.PostRequest.Id)
	//assert
	assert.Equal(t, code1, http.StatusOK)
	require.Equal(t, code2, http.StatusOK)
}

func TestDeletePostNotFound(t *testing.T) {
	db.SetUp(t, test_consts.PostRequest)
	defer db.TearDown()
	//arrange
	server := core.NewServer(postgresql.NewPostRepo(db.DB), postgresql.NewCommentRepo(db.DB), sender.NewKafkaSender(pr, test_consts.Topic))

	code := server.DeletePost(test_consts.Ctx, test_consts.PostRequest.Id)

	//assert
	require.Equal(t, code, http.StatusNotFound)
}

func TestUpdatePostSucces(t *testing.T) {
	db.SetUp(t, test_consts.PostRequest)
	defer db.TearDown()
	//arrange
	server := core.NewServer(postgresql.NewPostRepo(db.DB), postgresql.NewCommentRepo(db.DB), sender.NewKafkaSender(pr, test_consts.Topic))
	code1 := server.CreatePost(test_consts.Ctx, test_consts.PostRequest)
	code2 := server.UpdatePost(test_consts.Ctx, test_consts.PostRequestInvalid)
	//assert
	assert.Equal(t, code1, http.StatusOK)
	require.Equal(t, code2, http.StatusOK)
}

func TestUpdateNotFound(t *testing.T) {
	db.SetUp(t, test_consts.PostRequest)
	defer db.TearDown()
	//arrange
	server := core.NewServer(postgresql.NewPostRepo(db.DB), postgresql.NewCommentRepo(db.DB), sender.NewKafkaSender(pr, test_consts.Topic))

	code := server.UpdatePost(test_consts.Ctx, test_consts.PostRequestInvalid)

	//assert
	require.Equal(t, code, http.StatusNotFound)
}

func TestGetPostSuccces(t *testing.T) {
	db.SetUp(t, test_consts.PostRequest)
	defer db.TearDown()
	//arrange
	server := core.NewServer(postgresql.NewPostRepo(db.DB), postgresql.NewCommentRepo(db.DB), sender.NewKafkaSender(pr, test_consts.Topic))
	code1 := server.CreatePost(test_consts.Ctx, test_consts.PostRequest)
	post, code2 := server.GetPostById(test_consts.Ctx, test_consts.PostRequest.Id)

	//assert
	assert.Equal(t, code1, http.StatusOK)
	assert.Equal(t, code2, http.StatusOK)
	require.Equal(t, string(post), "{\"id\":1,\"comments\":[],\"body\":\"test\"}")

}

func TestGetPostNotFound(t *testing.T) {
	db.SetUp(t, test_consts.PostRequest)
	defer db.TearDown()
	//arrange
	server := core.NewServer(postgresql.NewPostRepo(db.DB), postgresql.NewCommentRepo(db.DB), sender.NewKafkaSender(pr, test_consts.Topic))

	post, code := server.GetPostById(test_consts.Ctx, test_consts.PostRequest.Id)

	//assert
	assert.Equal(t, code, http.StatusNotFound)
	require.Equal(t, string(post), "")
}

func TestCreateCommentSuccess(t *testing.T) {
	db.SetUp(t, test_consts.CommentRequest)
	defer db.TearDown()
	//arrange
	server := core.NewServer(postgresql.NewPostRepo(db.DB), postgresql.NewCommentRepo(db.DB), sender.NewKafkaSender(pr, test_consts.Topic))
	code1 := server.CreatePost(test_consts.Ctx, test_consts.PostRequest)
	code2 := server.CreateComment(test_consts.Ctx, test_consts.CommentRequest)

	//assert
	assert.Equal(t, code1, http.StatusOK)
	require.Equal(t, code2, http.StatusOK)

}

func TestCreateCommentConflict(t *testing.T) {
	db.SetUp(t, test_consts.CommentRequest)
	defer db.TearDown()
	//arrange
	server := core.NewServer(postgresql.NewPostRepo(db.DB), postgresql.NewCommentRepo(db.DB), sender.NewKafkaSender(pr, test_consts.Topic))
	code1 := server.CreatePost(test_consts.Ctx, test_consts.PostRequest)
	code2 := server.CreateComment(test_consts.Ctx, test_consts.CommentRequest)

	code3 := server.CreateComment(test_consts.Ctx, test_consts.CommentRequest)

	//assert
	assert.Equal(t, code1, http.StatusOK)
	assert.Equal(t, code2, http.StatusOK)
	require.Equal(t, code3, http.StatusConflict)
}
