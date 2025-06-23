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
	authUC := usecase.NewAuthUseCase(authRepo)
	handlers := handlers.NewAuthController(authUC)

	route := router.Group("/auth")
	{
		route.POST("/register", handlers.Register)
		route.POST("/login", handlers.Login)
	}
}
