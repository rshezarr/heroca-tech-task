package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
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

func (a *App) RunApp() {
	go func() {
		if err := a.httpServer.StartServer(); err != nil {
			log.Println(err)
			return
		}
	}()
	log.Println("http server started on :9091")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	fmt.Println()
	log.Println("Received terminate, graceful shutdown", sig)

	if err := a.httpServer.Shutdown(); err != nil {
		log.Println(err)
		return
	}
}
