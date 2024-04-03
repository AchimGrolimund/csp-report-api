package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func ZapLogger(log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Calculate response time
		elapsed := time.Since(start)

		// Log request details
		log.Info("Incoming request",
			zap.String("method", c.Request.Method),
			zap.String("proto", c.Request.Proto),
			zap.String("host", c.Request.Host),
			zap.String("uri", c.Request.RequestURI),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("elapsed", elapsed),
			zap.String("clientIP", c.ClientIP()),
			zap.Int64("contentLength", c.Request.ContentLength),
			zap.String("referer", c.Request.Referer()),
			zap.String("userAgent", c.Request.UserAgent()),
		)
	}
}
