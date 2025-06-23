package usecase

import (
	"fmt"

	"github.com/garciamendes/small_linkedin/src/internal/domain"
	"github.com/garciamendes/small_linkedin/src/internal/dtos"
	"github.com/google/uuid"
)

type UserUseCase struct {
	userRepo domain.UserRepository
	authRepo domain.AuthRepository
}

func NewUserUseCase(userRepo domain.UserRepository, authRepo domain.AuthRepository) *UserUseCase {
	return &UserUseCase{userRepo, authRepo}
}

func (useCase *UserUseCase) UserDetail(userID uuid.UUID) (*dtos.UserDetail, error) {
	user, err := useCase.userRepo.GetById(userID)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	auth, _ := useCase.authRepo.GetByUserId(user.ID)

	return &dtos.UserDetail{
		ID:        user.ID,
		Name:      user.Name,
		Email:     auth.Email,
		Headline:  user.Headline,
		Timestamp: user.Timestamp,
	}, nil
}
