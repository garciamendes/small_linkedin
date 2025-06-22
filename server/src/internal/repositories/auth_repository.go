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

func (a *authRepo) Create(authData *domain.Auth) error {
	a.db.Begin()

	var user domain.User
	if err := a.db.Create(&user).Error; err != nil {
		a.db.Rollback()
		return err
	}

	var auth = &domain.Auth{
		Email:    authData.Email,
		UserID:   user.ID,
		Password: authData.Password,
	}
	if err := a.db.Create(auth).Error; err != nil {
		a.db.Rollback()
		return err
	}

	a.db.Commit()
	return nil
}

func (a *authRepo) GetByEmail(email string) (*domain.Auth, error) {
	var auth domain.Auth

	if err := a.db.Where("email = ?", email).First(&auth).Error; err != nil {
		return nil, err
	}

	return &auth, nil
}
