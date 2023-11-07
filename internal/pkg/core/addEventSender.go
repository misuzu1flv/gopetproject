package core

import (
	"log"
	"net/http"
)

func (s *Server) AddEventSender(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			if err := s.SendEvent(req.Method, req.RequestURI); err != nil {
				log.Println(err)
			}
			handler.ServeHTTP(w, req)
		},
	)
}
