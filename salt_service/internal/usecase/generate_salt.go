package usecase

type Salt interface {
	GenerateSalt() string
}

type SaltUsecase struct{}

func NewSaltUsecase() *SaltUsecase {
	return &SaltUsecase{}
}
