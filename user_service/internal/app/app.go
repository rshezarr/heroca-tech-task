package app

import "go.mongodb.org/mongo-driver/mongo"

type App struct {
	db *mongo.Database

	httpServer  *Server
	httpHandler *handler.Handler
}
