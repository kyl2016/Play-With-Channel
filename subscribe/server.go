package main

import (
	"time"
)

type Server struct {
	state       string
	subscribers []func(state string)
}

func (s *Server) Subscribe(stateChanged func(state string)) {
	s.subscribers = append(s.subscribers, stateChanged)
}

func (s *Server) Start() {
	time.Sleep(time.Second)
	s.state = "RUNNING"
	for _, sub := range s.subscribers {
		sub(s.state)
	}

	time.Sleep(time.Second)
	s.state = "FINISHED"
	for _, sub := range s.subscribers {
		sub(s.state)
	}
}
