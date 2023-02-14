package app

import (
	"salt_srv/internal"
	"salt_srv/internal/handler"
)

type App struct {
	httpServer  *internal.Server
	httpHandler *handler.Handler
}
