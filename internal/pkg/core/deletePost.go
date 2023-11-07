package core

import (
	"context"
	"errors"
	"homework-8/internal/pkg/repository"
	"net/http"
)

func (s *Server) DeletePost(ctx context.Context, id int64) int {
	if err := s.commentrepo.DeleteByPostId(ctx, id); err != nil {
		return http.StatusInternalServerError
	}

	if err := s.postrepo.Delete(ctx, id); err != nil {
		if errors.Is(err, repository.ErrUknownId) {
			return http.StatusNotFound
		}
		return http.StatusInternalServerError
	}

	return http.StatusOK
}
