package handler

import (
	"user_svc/internal/usecase"

	"github.com/go-chi/chi"
)

type Handler struct {
	router  *chi.Mux
	service *usecase.Usecase
}

func NewHandler(usecase *usecase.Usecase) *Handler {
	return &Handler{
		router:  chi.NewRouter(),
		service: usecase,
	}
}
