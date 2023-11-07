package core

import (
	"homework-8/internal/pkg/sender"
	"time"
)

func (s *Server) SendEvent(eventType, request string) error {
	err := s.sender.SendMessage(sender.Message{
		EventType: eventType,
		Request:   request,
		Time:      time.Now(),
	})

	return err
}
