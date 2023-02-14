package usecase

import "user_svc/internal/repository"

type Usecase struct {
	User User
}

func NewUsecase(repo *repository.Repository) *Usecase {
	return &Usecase{
		User: NewUser(repo),
	}
}
