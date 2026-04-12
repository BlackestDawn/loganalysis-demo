package queue

import (
	"github.com/BlackestDawn/loganalysis-demo/go-api/internal/config"
	"github.com/BlackestDawn/loganalysis-demo/go-api/internal/utils"
)

type LogsQueue interface {
	PublishJSON(key string, val any) error
}

func NewQueue(conf *config.Config) (queue LogsQueue) {
	queue, err := initRabbitMQ(conf)
	utils.FailOnError(err, "Failed to initialize message queue")

	return
}
