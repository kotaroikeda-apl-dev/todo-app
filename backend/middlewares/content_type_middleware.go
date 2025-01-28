package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ValidateContentTypeMiddleware は Content-Type を application/json に限定するミドルウェア
func ValidateContentTypeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.ContentType() != "application/json" {
			c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": "Content-Type must be application/json"})
			c.Abort()
			return
		}
		c.Next()
	}
}
