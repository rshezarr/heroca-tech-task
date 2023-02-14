package app

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	db *mongo.Database

	httpServer  *Server
	httpHandler *handler.Handler
}

func New() (*App, error) {
	db, err := mongodb.NewClient(context.Background())
	if err != nil {
		return nil, err
	}

	repo := repository.NewRepository(db)
	usecase := usecase.NewUsecase(repo)

	httpHandler := httpHandler.NewHandler(usecase)

	return &App{
		db:          db,
		httpServer:  server.NewServer(httpHandler.InitRoutes()),
		httpHandler: httpHandler,
	}, nil
}
