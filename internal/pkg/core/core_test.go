package core_test

import (
	"homework-8/internal/pkg/core"
	mock_repos "homework-8/internal/pkg/repository/mocks"
	mock_sender "homework-8/internal/pkg/sender/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

type TestServer struct {
	Server      *core.Server
	MockComment *mock_repos.MockCommentRepo
	MockPost    *mock_repos.MockPostRepo
	MockSender  *mock_sender.MockSender
	ctrl        *gomock.Controller
}

func SetupTestServer(t *testing.T) *TestServer {
	t.Helper()
	ts := TestServer{}
	ts.ctrl = gomock.NewController(t)
	ts.MockComment = mock_repos.NewMockCommentRepo(ts.ctrl)
	ts.MockPost = mock_repos.NewMockPostRepo(ts.ctrl)
	ts.MockSender = mock_sender.NewMockSender(ts.ctrl)
	ts.Server = core.NewServer(ts.MockPost, ts.MockComment, ts.MockSender)
	return &ts
}
func (ts *TestServer) Teardown(t *testing.T) {
	t.Helper()
	ts.ctrl.Finish()
}
