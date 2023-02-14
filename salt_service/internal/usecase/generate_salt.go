package usecase

import "math/rand"

type Salt interface {
	GenerateSalt() string
}

type SaltUsecase struct{}

func NewSaltUsecase() *SaltUsecase {
	return &SaltUsecase{}
}

func (s *SaltUsecase) GenerateSalt() string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, 12)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
