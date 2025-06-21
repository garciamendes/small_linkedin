package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name     string    `json:"name" gorm:"not null;size:100"`
	Headline string    `json:"headline,omitempty" gorm:"size:200"`
	*Timestamp
}

func (n *User) BeforeCreate(tx *gorm.DB) (err error) {
	n.ID = uuid.New()
	return
}
