package server

import (
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func (s *Server) Run(addr string, handler http.Handler) error {
	s.server = &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}
	return s.server.ListenAndServe()
}
