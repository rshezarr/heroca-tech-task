package usecase

type Usecase struct {
	Salt Salt
}

func NewUsecase() *Usecase {
	return &Usecase{
		Salt: NewSaltUsecase(),
	}
}
