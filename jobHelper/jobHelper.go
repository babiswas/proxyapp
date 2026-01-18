package jobhelper

import (
	"app/logger"
	"fmt"
)

type Message struct {
	Msg string `json:"message"`
}

func MessageProcessingWorker(id int, msgQueue chan Message) {
	logger := logger.LoggingInit()
	logger.Info("Processing message from the queue.")
	for job := range msgQueue {
		fmt.Println("Processed message from queue:", job.Msg)
	}
}
