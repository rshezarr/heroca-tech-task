package repository

import (
	"context"
	"user_svc/internal"

	"go.mongodb.org/mongo-driver/mongo"
)

type User interface {
	Create(ctx context.Context, user internal.User) error
	Get(ctx context.Context, email string) (internal.User, error)
}

type UserRepo struct {
	db *mongo.Collection
}

func NewUser(database *mongo.Database, collection string) *UserRepo {
	return &UserRepo{
		db: database.Collection(collection),
	}
}
