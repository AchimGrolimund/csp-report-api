package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORSValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			if origin != "http://localhost" && origin != "https://groltech.ch" && origin != "https://grolimund-solutions.ch" && origin != "https://csp.groltech.ch" {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Not allowed CORS origin."})
				return
			}
		}
		c.Next()
	}
}
