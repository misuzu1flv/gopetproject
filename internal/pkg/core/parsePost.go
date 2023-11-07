package core

import (
	"encoding/json"
	"net/http"
)

func (s *Server) ParsePost(req *http.Request) (*PostRequest, error) {
	unm := &PostRequest{}
	if err := json.NewDecoder(req.Body).Decode(unm); err != nil {
		return nil, err
	}
	return unm, nil
}
