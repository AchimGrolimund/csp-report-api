// File path: api/middleware/auth.go

package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Disable auth for now, can not be used with the CSP report endpoint
		//authHeader := c.GetHeader("Authorization")
		//if authHeader == "" {
		//	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		//	return
		//}

		c.Next()
	}
}
