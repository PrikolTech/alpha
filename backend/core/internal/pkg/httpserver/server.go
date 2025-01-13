package httpserver

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/PrikolTech/alpha/backend/core/pkg/env"
)

type Server struct {
	server *http.Server
	errCh  chan error
}

func New(handler http.Handler, opts ...OptionFunc) *Server {
	host := env.GetString("SERVICE_HOST", "localhost")
	port := env.GetString("SERVICE_PORT", "3000")

	httpServer := &http.Server{
		Handler:      handler,
		Addr:         net.JoinHostPort(host, port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	server := &Server{
		server: httpServer,
	}

	for _, opt := range opts {
		opt(server)
	}

	return server
}

func (s *Server) Addr() string { return s.server.Addr }

func (s *Server) Start(ctx context.Context) {
	s.server.BaseContext = func(_ net.Listener) context.Context {
		return ctx
	}

	go func() {
		s.errCh <- s.server.ListenAndServe()
		close(s.errCh)
	}()
}

func (s *Server) Err() <-chan error { return s.errCh }

func (s *Server) Stop(ctx context.Context) error { return s.server.Shutdown(ctx) }
