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

func TestDeletePostSuccess(t *testing.T) {
	t.Parallel()

	// arrange
	ts := SetupTestServer(t)
	defer ts.Teardown(t)

	ts.MockPost.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)
	ts.MockComment.EXPECT().DeleteByPostId(gomock.Any(), gomock.Any()).Return(nil)
	//act
	code := ts.Server.DeletePost(test_consts.Ctx, int64(test_consts.Id))
	// assert
	require.Equal(t, http.StatusOK, code)
}

func TestDeleteNotFound(t *testing.T) {
	t.Parallel()

	// arrange
	ts := SetupTestServer(t)
	defer ts.Teardown(t)

	ts.MockPost.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(repository.ErrUknownId)
	ts.MockComment.EXPECT().DeleteByPostId(gomock.Any(), gomock.Any()).Return(nil)
	//act
	code := ts.Server.DeletePost(test_consts.Ctx, int64(test_consts.Id))
	// assert
	require.Equal(t, http.StatusNotFound, code)
}

func TestDeleteCouldntDeleteComment(t *testing.T) {
	t.Parallel()

	// arrange
	ts := SetupTestServer(t)
	defer ts.Teardown(t)

	ts.MockComment.EXPECT().DeleteByPostId(gomock.Any(), gomock.Any()).Return(errors.New("something"))
	//act
	code := ts.Server.DeletePost(test_consts.Ctx, int64(test_consts.Id))
	// assert
	require.Equal(t, http.StatusInternalServerError, code)
}

func TestDeleteCouldntDeletePost(t *testing.T) {
	t.Parallel()

	// arrange
	ts := SetupTestServer(t)
	defer ts.Teardown(t)

	ts.MockPost.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(errors.New("something"))
	ts.MockComment.EXPECT().DeleteByPostId(gomock.Any(), gomock.Any()).Return(nil)
	//act
	code := ts.Server.DeletePost(test_consts.Ctx, int64(test_consts.Id))
	// assert
	require.Equal(t, http.StatusInternalServerError, code)
}
