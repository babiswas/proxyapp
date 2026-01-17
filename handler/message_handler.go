package handler

import (
	"app/helper"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func MessageHandler(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx.Request.Context(), 10*time.Second)
	defer cancel()

	result, err := helper.ValidateRequest(c, ctx.Request.Header)
	if err != nil {
		if err == context.DeadlineExceeded {
			ctx.JSON(http.StatusRequestTimeout, gin.H{
				"error": "request timed out",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	data := map[string]string{"user_agent": result}
	ctx.AsciiJSON(http.StatusOK, data)
}
