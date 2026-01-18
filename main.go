package main

import (
	jobhelper "app/JobHelper"
	"app/handler"
	"app/helper"
	"app/logger"
	"app/middleware"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	helper.LoadEnv()
}

func main() {

	logger := logger.LoggingInit()

	logger.Info("Fetching port from enviroment variable.")
	port := ":" + os.Getenv("PORT")

	messageQueue := make(chan jobhelper.Message, 100)
	for i := 0; i < 10; i++ {
		go jobhelper.MessageProcessingWorker(i, messageQueue)
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	logger.Info("Configuring the middleware.")
	router.Use(middleware.HostSecurityMiddleWare(logger), middleware.LoggerMiddleWare(logger), gin.Recovery())

	logger.Info("Configuring the routes:")
	router.GET("/test", handler.MessageHandler)
	router.GET("/health", handler.HealthMessage)
	router.POST("/notify", jobhelper.MessageQueueService(messageQueue))

	logger.Info("Configuraing the server.")
	server := &http.Server{
		Addr:         port,
		Handler:      router,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		IdleTimeout:  30 * time.Second,
	}
	server.ListenAndServe()
}
