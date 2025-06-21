package usecase

import (
	"github.com/garciamendes/small_linkedin/src/internal/domain"
)

type UserUseCase struct {
	repo domain.UserRepository
}

func NewUserUseCase(r domain.UserRepository) *UserUseCase {
	// TODO!: Adicionar toda tratativas
	return &UserUseCase{repo: r}
}

func (u *UserUseCase) Create(user *domain.User) error {
	// TODO!: Adicionar toda tratativas
	return u.repo.Create(user)
}
