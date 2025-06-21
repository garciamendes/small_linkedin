package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name     string    `json:"name,omitempty" gorm:"null;"`
	Headline string    `json:"headline,omitempty" gorm:"size:200"`
	*Timestamp
}

func (n *User) BeforeCreate(tx *gorm.DB) (err error) {
	n.ID = uuid.New()
	return
}

type UserRepository interface {
	Create(user *User) error
	GetById(id uuid.UUID) (*User, error)
}
