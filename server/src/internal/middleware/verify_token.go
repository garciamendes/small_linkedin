package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/garciamendes/small_linkedin/src/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const keyAuth = "Bearer "
const UserIDKey string = "user"

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, keyAuth) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, keyAuth)

		token, err := jwt.ParseWithClaims(tokenString, &utils.Claims{}, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unauthorized")
			}

			secret := os.Getenv("TOKEN_SECRET_KEY")
			if secret == "" {
				return nil, fmt.Errorf("unauthorized")
			}

			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			return
		}

		if claims, ok := token.Claims.(*utils.Claims); ok {
			c.Set(UserIDKey, claims.UserID.String())
		}

		c.Next()
	}
}
