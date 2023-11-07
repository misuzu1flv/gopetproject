package core

import (
	"context"
	"errors"
	"homework-8/internal/pkg/repository"
	"net/http"
)

func (s *Server) CreatePost(ctx context.Context, p *PostRequest) int {

	err := s.postrepo.Add(ctx, &repository.Post{
		Id:   p.Id,
		Body: p.Body,
	})

	if err != nil {
		if errors.Is(err, repository.ErrObjectExists) {
			return http.StatusConflict
		}
		return http.StatusInternalServerError
	}

	return http.StatusOK
}
