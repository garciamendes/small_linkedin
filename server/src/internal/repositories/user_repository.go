package repositories

import (
	"github.com/garciamendes/small_linkedin/src/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func UserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepo{db}
}

func (u *userRepo) Create(user *domain.User) error {
	if err := u.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (u *userRepo) GetById(id uuid.UUID) (*domain.User, error) {
	var user domain.User

	if err := u.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
