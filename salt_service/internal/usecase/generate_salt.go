package usecase

type Salt interface {
	GenerateSalt() string
}
