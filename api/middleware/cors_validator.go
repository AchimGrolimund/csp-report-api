package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORSValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			if origin != "http://localhost" && origin != "http://grolimund-solutions.ch" && origin != "https://grolimund-solutions.ch" {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Not allowed CORS origin."})
				return
			}
		}
		c.Next()
	}
}
