package core

import (
	"context"
	"encoding/json"
	"errors"
	"homework-8/internal/pkg/repository"
	"net/http"
)

func (s *Server) GetPostById(ctx context.Context, id int64) ([]byte, int) {
	post, err := s.postrepo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrUknownId) {
			return nil, http.StatusNotFound
		} else {
			return nil, http.StatusInternalServerError
		}
	}

	comments, err := s.commentrepo.GetByPostId(ctx, id)
	if err != nil {
		return nil, http.StatusInternalServerError
	}

	commentRequests := make([]CommentRequest, 0)
	for _, comment := range comments {
		commentRequests = append(commentRequests, CommentRequest{Id: comment.Id, PostId: comment.PostId, Body: comment.Body})
	}

	json, err := json.Marshal(postGetRequest{Id: post.Id, Comments: commentRequests, Body: post.Body})
	if err != nil {
		return nil, http.StatusInternalServerError
	}

	return json, http.StatusOK
}
