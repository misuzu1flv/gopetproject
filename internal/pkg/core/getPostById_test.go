package core_test

import (
	"errors"
	"homework-8/internal/pkg/repository"
	test_consts "homework-8/tests/consts"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetPostByIdSucces(t *testing.T) {
	t.Parallel()

	// arrange
	ts := SetupTestServer(t)
	defer ts.Teardown(t)

	ts.MockPost.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(test_consts.Post, nil)
	ts.MockComment.EXPECT().GetByPostId(gomock.Any(), gomock.Any()).Return([]*repository.Comment{test_consts.Comment}, nil)
	//act
	body, code := ts.Server.GetPostById(test_consts.Ctx, int64(test_consts.Id))
	// assert
	require.Equal(t, http.StatusOK, code)
	require.Equal(t, "{\"id\":1,\"comments\":[{\"id\":1,\"postId\":1,\"body\":\"hello\"}],\"body\":\"hello\"}", string(body))
}

func TestGetPostByIdPostDoesntExist(t *testing.T) {
	t.Parallel()

	// arrange
	ts := SetupTestServer(t)
	defer ts.Teardown(t)

	ts.MockPost.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(nil, repository.ErrUknownId)
	//act
	_, code := ts.Server.GetPostById(test_consts.Ctx, int64(test_consts.Id))
	// assert
	require.Equal(t, http.StatusNotFound, code)
}

func TestGetPostByIdErrorGettingPost(t *testing.T) {
	t.Parallel()

	// arrange
	ts := SetupTestServer(t)
	defer ts.Teardown(t)

	ts.MockPost.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(nil, errors.New("Something"))
	//act
	_, code := ts.Server.GetPostById(test_consts.Ctx, int64(test_consts.Id))
	// assert
	require.Equal(t, http.StatusInternalServerError, code)
}

func TestGetPostByIdErrorGettingComment(t *testing.T) {
	t.Parallel()

	// arrange
	ts := SetupTestServer(t)
	defer ts.Teardown(t)

	ts.MockPost.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(test_consts.Post, nil)
	ts.MockComment.EXPECT().GetByPostId(gomock.Any(), gomock.Any()).Return(nil, errors.New("Something"))
	//act
	_, code := ts.Server.GetPostById(test_consts.Ctx, int64(test_consts.Id))
	// assert
	require.Equal(t, http.StatusInternalServerError, code)
}
