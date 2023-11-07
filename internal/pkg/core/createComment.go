package core

import (
	"context"
	"errors"
	"homework-8/internal/pkg/repository"
	"log"
	"net/http"
)

func (s *Server) CreateComment(ctx context.Context, comment *CommentRequest) int {
	log.Println("CreateComment")

	err := s.commentrepo.Add(ctx, &repository.Comment{
		Id:     comment.Id,
		PostId: comment.PostId,
		Body:   comment.Body,
	})

	if err != nil {
		if errors.Is(err, repository.ErrObjectExists) {
			return http.StatusConflict
		}
		return http.StatusInternalServerError
	}

	return http.StatusOK
}
