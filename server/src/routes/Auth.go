package routes

import "github.com/gin-gonic/gin"

func AuthRoutes(router *gin.Engine) {
	authApi := router.Group("/api/auth")
	{
		authApi.GET("/health", func(h *gin.Context) {
			h.JSON(200, gin.H{"status": "ok"})
		})

		// authApi.POST("/login", infra.Login)
		// authApi.POST("/register", infra.Register)
		// authApi.GET("/logout", infra.Logout)
		// authApi.GET("/me", infra.GetMe)
		// authApi.PATCH("/me", infra.UpdateMe)
	}
}
