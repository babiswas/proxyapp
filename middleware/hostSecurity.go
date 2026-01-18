package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func HostSecurityMiddleWare(logger *logrus.Logger) gin.HandlerFunc {
	logger.Info("Executing host security logger middleware.")
	return func(ctx *gin.Context) {
		host := os.Getenv("HOST")
		port := os.Getenv("PORT")
		expected_host := host + ":" + port
		logger.Info("Expected host:", expected_host)
		logger.Info("Host header:", ctx.Request.Host)
		if ctx.Request.Host != expected_host {
			logger.Info("Host mismatch occured.")
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid host header"})
		}
		ctx.Next()
	}
}
