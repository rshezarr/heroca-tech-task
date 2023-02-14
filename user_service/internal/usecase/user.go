package usecase

import (
	"context"
	"user_svc/internal"
	"user_svc/internal/repository"
)

type User interface {
	Create(ctx context.Context, user internal.User, salt string) error
	Get(ctx context.Context, email string) (internal.User, error)
}

type UserUsecase struct {
	repo *repository.Repository
}

func NewUser(repo *repository.Repository) *UserUsecase {
	return &UserUsecase{
		repo: repo,
	}
}
