package middleware

import (
	"net/http"

	"github.com/MorrisMorrison/retfig/infrastructure/jwt"
	"github.com/gin-gonic/gin"
)

func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		claims, err := jwt.ParseToken(cookie)
		if err != nil || claims.Valid() != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("currentUser", claims.Subject)

		c.Next()
	}
}
