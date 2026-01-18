package handler

import (
	"app/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthMessage(ctx *gin.Context) {
	logger := logger.LoggingInit()
	logger.Info("Executing health message handler.")
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
