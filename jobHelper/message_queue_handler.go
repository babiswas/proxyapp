package jobhelper

import (
	"app/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MessageQueueService(msgQueue chan<- Message) func(ctx *gin.Context) {
	logger := logger.LoggingInit()
	logger.Info("Queueing message in message queue.")
	return func(ctx *gin.Context) {
		var tempMsg struct {
			Msg string `json:"message"`
		}
		if err := ctx.ShouldBindJSON(&tempMsg); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid request",
			})
			return
		}
		job := Message{Msg: tempMsg.Msg}

		select {
		case msgQueue <- job:
			logger.Info("Sucessfully queued the message.")
			ctx.JSON(http.StatusAccepted, gin.H{
				"status": "message queued.",
			})
		default:
			logger.Info("Queue service is busy:")
			ctx.JSON(http.StatusAccepted, gin.H{
				"error": "queueing service busy.",
			})
		}
	}
}
