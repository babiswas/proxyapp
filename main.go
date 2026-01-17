package main

import (
	"app/handler"
	"app/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(middleware.LoggerMiddleWare(logger), gin.Recovery())
	router.GET("/test", handler.MessageHandler)
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		IdleTimeout:  30 * time.Second,
	}
	server.ListenAndServe()
}
