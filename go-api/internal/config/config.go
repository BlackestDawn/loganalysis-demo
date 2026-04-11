package config

import (
	"context"
	"os"
)

type Config struct {
	ListenPort string
	DbUrl      string
	DbCtx      context.Context
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

	return &Config{
		ListenPort: ":" + listenPort,
		DbUrl:      dbURL,
		DbCtx:      context.Background(),
	}
}
