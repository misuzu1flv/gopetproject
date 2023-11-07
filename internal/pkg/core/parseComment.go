package core

import (
	"encoding/json"
	"net/http"
)

func (s *Server) ParseComment(req *http.Request) (*CommentRequest, error) {
	unm := &CommentRequest{}
	if err := json.NewDecoder(req.Body).Decode(unm); err != nil {
		return nil, err
	}
	return unm, nil
}
