package core

import (
	"context"
	"homework-8/internal/pkg/repository"
	"homework-8/internal/pkg/sender"
)

type PostRequest struct {
	Id   int64  `json:"id"`
	Body string `json:"body"`
}

type postGetRequest struct {
	Id       int64            `json:"id"`
	Comments []CommentRequest `json:"comments"`
	Body     string           `json:"body"`
}

type CommentRequest struct {
	Id     int64  `json:"id"`
	PostId int64  `json:"postId"`
	Body   string `json:"body"`
}

type PostRepo interface {
	GetById(ctx context.Context, id int64) (*repository.Post, error)
	Add(ctx context.Context, post *repository.Post) error
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, post *repository.Post) error
}

type CommentRepo interface {
	GetById(ctx context.Context, id int64) (*repository.Comment, error)
	GetByPostId(ctx context.Context, post_id int64) ([]*repository.Comment, error)
	DeleteByPostId(ctx context.Context, post_id int64) error
	Add(ctx context.Context, comment *repository.Comment) error
}

type Sender interface {
	SendMessage(sender.Message) error
}

type Server struct {
	postrepo    PostRepo
	commentrepo CommentRepo
	sender      Sender
}

func NewServer(pr PostRepo, cr CommentRepo, s Sender) *Server {
	return &Server{pr, cr, s}
}
