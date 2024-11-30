package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

// NewServer - init new server. address is `host:port`.
func NewServer(address string) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:              address,
			ReadHeaderTimeout: 0,
		},
	}
}

func (s *Server) AddHandler(mux http.Handler) error {
	s.httpServer.Handler = mux
	return nil
}

func (s *Server) Run(ctx context.Context) error {

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Print(fmt.Errorf(`lister server have err: %w`, err))
		}

		if err := s.httpServer.Shutdown(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Print(fmt.Errorf(`stop server have err: %w`, err))
		}
	}()
	return nil
}
