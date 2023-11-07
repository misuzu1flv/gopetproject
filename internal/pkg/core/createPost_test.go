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

func TestCreatePostSuccess(t *testing.T) {
	t.Parallel()

	// arrange
	ts := SetupTestServer(t)
	defer ts.Teardown(t)

	ts.MockPost.EXPECT().Add(gomock.Any(), gomock.Any()).Return(nil)
	//act
	code := ts.Server.CreatePost(test_consts.Ctx, test_consts.PostRequest)
	// assert
	require.Equal(t, http.StatusOK, code)
}

func TestCreatePostConflict(t *testing.T) {
	t.Parallel()

	// arrange
	ts := SetupTestServer(t)
	defer ts.Teardown(t)

	ts.MockPost.EXPECT().Add(gomock.Any(), gomock.Any()).Return(repository.ErrObjectExists)
	//act
	code := ts.Server.CreatePost(test_consts.Ctx, test_consts.PostRequest)
	// assert
	require.Equal(t, http.StatusConflict, code)
}

func TestCreatePostInternalError(t *testing.T) {
	t.Parallel()

	// arrange
	ts := SetupTestServer(t)
	defer ts.Teardown(t)

	ts.MockPost.EXPECT().Add(gomock.Any(), gomock.Any()).Return(errors.New("something"))
	//act
	code := ts.Server.CreatePost(test_consts.Ctx, test_consts.PostRequest)
	// assert
	require.Equal(t, http.StatusInternalServerError, code)
}
