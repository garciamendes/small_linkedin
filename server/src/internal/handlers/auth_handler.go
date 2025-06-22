package handlers

import (
	"net/http"

	"github.com/garciamendes/small_linkedin/src/internal/domain"
	"github.com/garciamendes/small_linkedin/src/internal/dtos"
	"github.com/garciamendes/small_linkedin/src/internal/usecase"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUseCase *usecase.AuthUseCase
}

func NewAuthController(uc *usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{uc}
}

func (authHandler *AuthHandler) Register(c *gin.Context) {
	var body dtos.CreateUserDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	var data = domain.Auth{
		Email:    body.Email,
		Password: body.Password,
	}
	if err := authHandler.authUseCase.Create(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}
