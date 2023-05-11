package middleware

import (
	"net/http"
	"os"
	// "test-golang/config/mongdb"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")
	if authorizationHeader == os.Getenv("SECRET_KEY") {
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}
}
