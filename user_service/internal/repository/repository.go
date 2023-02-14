package repository

import "go.mongodb.org/mongo-driver/mongo"

type Repository struct {
	User User
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		User: NewUser(db, "users"),
	}
}
