package usecase

import (
	"github.com/garciamendes/small_linkedin/src/internal/domain"
)

type AuthUseCase struct {
	repo domain.AuthRepository
}

func NewAuthUseCase(r domain.AuthRepository) *AuthUseCase {
	return &AuthUseCase{repo: r}
}

func (u *AuthUseCase) Create(data *domain.Auth) error {
	// TODO!: Adicionar toda tratativas
	return u.repo.Create(data)
}

func (u *AuthUseCase) GetByEmail(email string) (*domain.Auth, error) {
	// TODO!: Adicionar toda tratativas
	return u.repo.GetByEmail(email)
}
