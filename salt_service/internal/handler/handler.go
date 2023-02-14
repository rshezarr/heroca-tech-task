package handler

import (
	"salt_srv/internal/usecase"

	"github.com/go-chi/chi"
)

type Handler struct {
	router  *chi.Mux
	service *usecase.Usecase
}

func NewHandler(service *usecase.Usecase) *Handler {
	return &Handler{
		router:  chi.NewRouter(),
		service: service,
	}
}
