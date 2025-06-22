package usecase

import (
	"fmt"

	"github.com/garciamendes/small_linkedin/src/internal/domain"
	"github.com/garciamendes/small_linkedin/src/internal/utils"
)

type AuthUseCase struct {
	authRepo domain.AuthRepository
}

func NewAuthUseCase(a domain.AuthRepository) *AuthUseCase {
	return &AuthUseCase{authRepo: a}
}

func (useCase *AuthUseCase) Create(data *domain.Auth) error {
	account, errAccount := useCase.authRepo.GetByEmail(data.Email)

	if errAccount != nil {
		return errAccount
	}

	if account != nil {
		return fmt.Errorf("email already registered")
	}

	passwordhashed, errPassword := utils.HashPassword(data.Password)

	if errPassword != nil {
		return errPassword
	}

	auth := &domain.Auth{
		Email:    data.Email,
		Password: passwordhashed,
	}

	err := useCase.authRepo.Create(auth)
	fmt.Print(err)
	if err != nil {
		return err
	}

	return nil
}

func (u *AuthUseCase) GetByEmail(email string) (*domain.Auth, error) {
	// TODO!: Adicionar toda tratativas
	return u.authRepo.GetByEmail(email)
}
