package logger

import "github.com/sirupsen/logrus"

func LoggingInit() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)
	return logger
}
