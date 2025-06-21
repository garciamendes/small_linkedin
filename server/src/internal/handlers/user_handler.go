package handlers

import "github.com/garciamendes/small_linkedin/src/internal/usecase"

type UserHandler struct {
	uc *usecase.UserUseCase
}

func NewUserHandler(uc *usecase.UserUseCase) *UserHandler {
	return &UserHandler{uc}
}
