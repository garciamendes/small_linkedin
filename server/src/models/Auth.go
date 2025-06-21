package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Auth struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Email    string    `json:"email" gorm:"not null;unique;size:255"`
	Password string    `json:"-" gorm:"not null"`
	UserID   uuid.UUID `json:"user_id" gorm:"type:uuid;not null;uniqueIndex"`
	*Timestamp
}

func (n *Auth) BeforeCreate(tx *gorm.DB) (err error) {
	n.ID = uuid.New()
	return
}
