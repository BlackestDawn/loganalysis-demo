package config

import (
	"context"
	"os"
)

type Config struct {
	ListenPort string
	DbUrl      string
	DbCtx      context.Context
	closers    []func() error
	Queue      QueueInfo
}

func (c *Config) Cleanup() {
	for _, closer := range c.closers {
		closer()
	}
}

func (c *Config) AddCloser(closer func() error) {
	c.closers = append(c.closers, closer)
}

type QueueInfo struct {
	Server    string
	QueueName string
	Key       string
	Exchange  string
}

func NewConfig() *Config {
	listenPort := os.Getenv("LISTEN_PORT")
	if len(listenPort) == 0 {
		listenPort = defaultListenPort
	}

	dbURL := os.Getenv("DB_URL")
	if len(dbURL) == 0 {
		dbURL = defaultDbUrl
	}

	qServer := os.Getenv("QUEUE_SERVER")
	qName := os.Getenv("QUEUE_NAME")
	qKey := os.Getenv("QUEUE_KEY")
	qExchange := os.Getenv("QUEUE_EXCHANGE")

	if len(qServer) == 0 {
		qServer = defaultLogQueueServer
	}

	if len(qName) == 0 {
		qName = defaultLogQueueName
	}

	if len(qKey) == 0 {
		qKey = defaultLogQueueKey
	}

	if len(qExchange) == 0 {
		qExchange = defaultLogQueueExchange
	}

	return &Config{
		ListenPort: ":" + listenPort,
		DbUrl:      dbURL,
		DbCtx:      context.Background(),
		Queue: QueueInfo{
			Server:    qServer,
			QueueName: qName,
			Key:       qKey,
			Exchange:  qExchange,
		},
	}
}
