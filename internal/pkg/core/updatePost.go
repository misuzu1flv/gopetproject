package core

import (
	"context"
	"errors"
	"homework-8/internal/pkg/repository"
	"net/http"
)

func (s *Server) UpdatePost(ctx context.Context, p *PostRequest) int {
	err := s.postrepo.Update(ctx, &repository.Post{
		Id:   p.Id,
		Body: p.Body,
	})

	if err != nil {
		if errors.Is(err, repository.ErrUknownId) {
			return http.StatusNotFound
		}
		return http.StatusInternalServerError
	}
	return http.StatusOK
}
