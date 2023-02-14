package internal

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type Server struct {
	server *http.Server
}

func NewServer(router *chi.Mux) *Server {
	return &Server{
		server: &http.Server{
			Addr:           ":9090",
			Handler:        router,
			WriteTimeout:   10 * time.Second,
			ReadTimeout:    10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func (s *Server) StartServer() error {
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	defer cancel()

	return s.server.Shutdown(ctx)
}
