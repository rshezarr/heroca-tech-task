package app

import (
	"salt_srv/internal"
	"salt_srv/internal/handler"
	"salt_srv/internal/usecase"
)

type App struct {
	httpServer  *internal.Server
	httpHandler *handler.Handler
}

func New() (*App, error) {
	usecase := usecase.NewUsecase()

	httpHandler := handler.NewHandler(usecase)

	return &App{
		httpHandler: httpHandler,
		httpServer:  internal.NewServer(httpHandler.InitRoutes()),
	}, nil
}
