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

func TestUpdatePost(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		// arrange
		ts := SetupTestServer(t)
		defer ts.Teardown(t)

		ts.MockPost.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
		//act
		code := ts.Server.UpdatePost(test_consts.Ctx, test_consts.PostRequest)
		// assert
		require.Equal(t, http.StatusOK, code)
	})

	t.Run("status not found", func(t *testing.T) {
		t.Parallel()

		// arrange
		ts := SetupTestServer(t)
		defer ts.Teardown(t)

		ts.MockPost.EXPECT().Update(gomock.Any(), gomock.Any()).Return(repository.ErrUknownId)
		//act
		code := ts.Server.UpdatePost(test_consts.Ctx, test_consts.PostRequest)
		// assert
		require.Equal(t, http.StatusNotFound, code)
	})

	t.Run("comment delete error", func(t *testing.T) {
		t.Parallel()

		// arrange
		ts := SetupTestServer(t)
		defer ts.Teardown(t)

		ts.MockPost.EXPECT().Update(gomock.Any(), gomock.Any()).Return(errors.New("something"))
		//act
		code := ts.Server.UpdatePost(test_consts.Ctx, test_consts.PostRequest)
		// assert
		require.Equal(t, http.StatusInternalServerError, code)
	})

}
