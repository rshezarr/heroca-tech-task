package app

import (
	"context"
	"user_svc/internal"
	"user_svc/internal/handler"
	"user_svc/internal/repository"
	"user_svc/internal/usecase"

	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	db *mongo.Database

	httpServer  *internal.Server
	httpHandler *handler.Handler
}

func New() (*App, error) {
	db, err := mongodb.NewClient(context.Background())
	if err != nil {
		return nil, err
	}

	repo := repository.NewRepository(db)
	usecase := usecase.NewUsecase(repo)

	httpHandler := handler.NewHandler(usecase)

	return &App{
		db:          db,
		httpServer:  internal.NewServer(httpHandler.InitRoutes()),
		httpHandler: httpHandler,
	}, nil
}
