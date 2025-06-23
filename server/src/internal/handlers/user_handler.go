package handlers

import (
	"net/http"

	"github.com/garciamendes/small_linkedin/src/internal/middleware"
	"github.com/garciamendes/small_linkedin/src/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	UserUsecase *usecase.UserUseCase
}

func NewUserHandler(uc *usecase.UserUseCase) *UserHandler {
	return &UserHandler{uc}
}

func (userHandler UserHandler) Me(c *gin.Context) {
	userID := c.GetString(middleware.UserIDKey)

	newUserID, _ := uuid.Parse(userID)
	user, _ := userHandler.UserUsecase.UserDetail(newUserID)

	c.JSON(http.StatusOK, user)
}
