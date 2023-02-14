package repository

import (
	"context"
	"errors"
	"user_svc/internal"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var ErrUserNotFound = errors.New("user not found")

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

func (r *UserRepo) Create(ctx context.Context, user internal.User) error {
	_, err := r.db.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) Get(ctx context.Context, email string) (internal.User, error) {
	filter := bson.M{"email": email}

	var user internal.User
	if err := r.db.FindOne(ctx, filter).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return internal.User{}, ErrUserNotFound
		}
		return internal.User{}, err
	}

	return user, nil
}
