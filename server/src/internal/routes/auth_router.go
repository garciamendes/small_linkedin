package routes

import (
	"github.com/garciamendes/small_linkedin/src/internal/handlers"
	"github.com/garciamendes/small_linkedin/src/internal/repositories"
	"github.com/garciamendes/small_linkedin/src/internal/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoutes(router *gin.RouterGroup, db *gorm.DB) {
	authRepo := repositories.AuthRepository(db)
	userUC := usecase.NewAuthUseCase(authRepo)
	handlers := handlers.NewAuthController(userUC)

	auth := router.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
	}
}
