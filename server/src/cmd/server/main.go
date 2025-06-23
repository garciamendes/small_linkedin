package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/garciamendes/small_linkedin/src/internal/infra"
	"github.com/garciamendes/small_linkedin/src/internal/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	DB := infra.NewDB()
	if DB.Error != nil {
		log.Fatal(DB.Error)
		return
	}

	infra.Migrate(DB)

	allowHost := os.Getenv("ALLOW_HOST")
	corsHandler := cors.New(cors.Config{
		AllowOrigins:     strings.Split(allowHost, ","),
		AllowMethods:     []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
	router.Use(corsHandler)
	router.Use(gin.Logger(), gin.Recovery())

	apiGroup := router.Group("/api")

	routes.AuthRoutes(apiGroup, DB)
	routes.UserRouter(apiGroup, DB)

	router.Run(":8080")
}
