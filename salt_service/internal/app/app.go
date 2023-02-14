package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"salt_svc/internal"
	"salt_svc/internal/handler"
	"salt_svc/internal/usecase"
	"syscall"
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

func (a *App) Start() error {
	go func() {
		if err := a.httpServer.StartServer(); err != nil {
			log.Fatal(err)
		}
	}()
	log.Println("generate-salt server started on :9090")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	fmt.Println()
	log.Println("Received terminate, graceful shutdown", sig)

	if err := a.httpServer.Shutdown(); err != nil {
		return err
	}

	return nil
}
