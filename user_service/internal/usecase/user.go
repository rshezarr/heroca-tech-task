package usecase

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"regexp"
	"user_svc/internal"
	"user_svc/internal/repository"
)

var ErrUserExists = errors.New("user already exists")

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

func (s *UserUsecase) Create(ctx context.Context, user internal.User, salt string) error {
	if _, err := s.Get(ctx, user.Email); err == nil {
		return ErrUserExists
	}

	matched, err := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, user.Email)
	if err != nil {
		return err
	}

	if !matched {
		return err
	}

	hash := md5.Sum([]byte(user.Password + salt))
	user.Password = hex.EncodeToString(hash[:])

	if err := s.repo.User.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *UserUsecase) Get(ctx context.Context, email string) (internal.User, error) {
	user, err := s.repo.User.Get(ctx, email)
	if err != nil {
		return internal.User{}, err
	}

	return user, nil
}
