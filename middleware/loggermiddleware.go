package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoggerMiddleWare(logger *logrus.Logger) gin.HandlerFunc {
	logger.Info("Executing logger middleware.")
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		duration := time.Since(start)
		logger.WithFields(logrus.Fields{
			"status":   ctx.Writer.Status(),
			"method":   ctx.Request.Method,
			"path":     ctx.Request.URL.Path,
			"duration": duration.String(),
			"client":   ctx.ClientIP(),
		}).Info("request completed")
	}
}
