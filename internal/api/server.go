package api

import (
	"net/http"
)

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Start() error {
	http.DefaultServeMux.HandleFunc("/league", s.handleGetLeague)
	http.HandleFunc("/deadline", s.handleGetDeadline)
	return http.ListenAndServe(s.listenAddr, nil)
}
