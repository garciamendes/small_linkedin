package repositories

import (
	"github.com/garciamendes/small_linkedin/src/internal/domain"
	"gorm.io/gorm"
)

type authRepo struct {
	db *gorm.DB
}

func AuthRepository(db *gorm.DB) domain.AuthRepository {
	return &authRepo{db}
}

func (a *authRepo) Create(auth *domain.Auth) error {
	if err := a.db.Create(auth).Error; err != nil {
		return err
	}

	return nil
}

func (a *authRepo) GetByEmail(email string) (*domain.Auth, error) {
	var auth domain.Auth

	if err := a.db.Where("email = ?", email).First(&auth).Error; err != nil {
		return nil, err
	}

	return &auth, nil
}
