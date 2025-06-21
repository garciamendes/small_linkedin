package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/garciamendes/small_linkedin/src/infra"
	"github.com/garciamendes/small_linkedin/src/routes"
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

	DB := infra.Init()
	if DB.Error != nil {
		log.Fatal(DB.Error)
		return
	}

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

	routes.AuthRoutes(router)

	router.Run(":8080")
}
