package test_consts

import (
	"context"
	"homework-8/internal/pkg/core"
	"homework-8/internal/pkg/repository"
	"time"
)

var (
	Brokers = []string{"127.0.0.1:9091", "127.0.0.1:9092", "127.0.0.1:9093"}

	Topic = "logs"

	Ctx = context.Background()

	CommentRequest = &core.CommentRequest{
		Id:     1,
		Body:   "test",
		PostId: 1,
	}

	PostRequest = &core.PostRequest{
		Id:   1,
		Body: "test",
	}

	Comment = &repository.Comment{
		Id:        1,
		PostId:    1,
		Body:      "hello",
		CreatedAt: time.Time{},
	}

	Post = &repository.Post{
		Id:        1,
		Body:      "hello",
		CreatedAt: time.Time{},
	}

	PostRequestInvalid = &core.PostRequest{
		Id:   1,
		Body: "GOODBYE!",
	}

	Id = 1
)
