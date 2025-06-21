package handlers

import (
	"github.com/garciamendes/small_linkedin/src/internal/usecase"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	uc *usecase.AuthUseCase
}

func NewAuthController(uc *usecase.AuthUseCase) *AuthController {
	return &AuthController{uc}
}

func (a *AuthController) Register(c *gin.Context) {}
