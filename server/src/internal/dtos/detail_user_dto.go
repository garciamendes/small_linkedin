package dtos

import (
	"github.com/garciamendes/small_linkedin/src/internal/domain"
	"github.com/google/uuid"
)

type UserDetail struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Headline string    `json:"headline"`
	*domain.Timestamp
}
