package repository

import (
	"context"
	"user_svc/internal"
)

type User interface {
	Create(ctx context.Context, user internal.User) error
	Get(ctx context.Context, email string) (internal.User, error)
}
