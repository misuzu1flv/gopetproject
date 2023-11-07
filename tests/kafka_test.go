package tests

import (
	"homework-8/internal/pkg/core"
	"homework-8/internal/pkg/repository/postgresql"
	"homework-8/internal/pkg/sender"
	test_consts "homework-8/tests/consts"
	loggertest "homework-8/tests/logger"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEvent(t *testing.T) {
	db.SetUp(t, test_consts.CommentRequest)
	defer db.TearDown()
	rch := loggertest.SetUpTest(cons)
	defer close(rch)
	//arrange
	server := core.NewServer(postgresql.NewPostRepo(db.DB), postgresql.NewCommentRepo(db.DB), sender.NewKafkaSender(pr, test_consts.Topic))
	req, err := http.NewRequest(http.MethodGet, "/post?id=1", nil)
	server.SendEvent(req.Method, req.RequestURI)
	res := <-rch
	//assert

	require.NoError(t, err)
	assert.Equal(t, "GET", res)
}
