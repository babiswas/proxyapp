package handler

import (
	"app/helper"
	"app/logger"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func MessageHandler(ctx *gin.Context) {
	logger := logger.LoggingInit()
	c, cancel := context.WithTimeout(ctx.Request.Context(), 10*time.Second)
	defer cancel()

	logger.Info("validating http request.")
	result, err := helper.ValidateRequest(c, ctx.Request.Header)
	if err != nil {
		if err == context.DeadlineExceeded {
			logger.Error("Request deadline exceeded.")
			ctx.JSON(http.StatusRequestTimeout, gin.H{
				"error": "request timed out",
			})
			return
		}

		logger.Error("Internal error occured.")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	logger.Info("Sucessfully validated the request.")
	data := map[string]string{"user_agent": result}
	ctx.AsciiJSON(http.StatusOK, data)
}
