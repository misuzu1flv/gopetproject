package core

import (
	"net/http"
	"strconv"
)

func (s *Server) ParseId(req *http.Request) (int64, error) {
	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	return int64(id), err
}
