package internal

import (
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
			Addr:           ":9091",
			Handler:        router,
			WriteTimeout:   10 * time.Second,
			ReadTimeout:    10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}
