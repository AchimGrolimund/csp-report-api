package middleware

import (
	"github.com/achimgrolimund/csp-report-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func CORSValidator() gin.HandlerFunc {
	log, err := logger.NewLogger()
	if err != nil {
		panic(err)
	}
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			if origin != "http://localhost" && origin != "https://groltech.ch" && origin != "https://grolimund-solutions.ch" && origin != "https://wiki.groltech.ch" && origin != "https://hoas.groltech.ch" && origin != "https://ansible.groltech.ch" {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Not allowed CORS origin."})
				log.Error("Not allowed CORS origin", zap.String("origin", origin))
				return
			}
		} else {
			log.Error("No CORS origin header found")
		}
		c.Next()
	}
}
