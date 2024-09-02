package server

import (
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

// address = host:port
func NewServer(address string) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr: address,
		},
	}
}

func (s *Server) AddHandler(mux http.Handler) error {
	s.httpServer.Handler = mux
	return nil
}

func (s *Server) Run() error {

	go func() {
		s.httpServer.ListenAndServe()
		// s.httpServer.Shutdown(ctx)
	}()
	return nil
}
