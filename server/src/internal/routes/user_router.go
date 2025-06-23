package routes

import (
	"github.com/garciamendes/small_linkedin/src/internal/handlers"
	"github.com/garciamendes/small_linkedin/src/internal/middleware"
	"github.com/garciamendes/small_linkedin/src/internal/repositories"
	"github.com/garciamendes/small_linkedin/src/internal/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRouter(router *gin.RouterGroup, db *gorm.DB) {
	userRepo := repositories.UserRepository(db)
	authRepo := repositories.AuthRepository(db)
	userUC := usecase.NewUserUseCase(userRepo, authRepo)
	handlers := handlers.NewUserHandler(userUC)

	route := router.Group("/user")
	route.Use(middleware.VerifyToken())
	{
		route.GET("/me", handlers.Me)
	}
}
