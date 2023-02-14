package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context) (*mongo.Database, error) {
	cliOptions := options.Client().ApplyURI("mongodb://mongo_db:27017")

	client, err := mongo.Connect(ctx, cliOptions)
	if err != nil {
		return nil, fmt.Errorf("connect: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}

	return client.Database("user-service"), nil
}
