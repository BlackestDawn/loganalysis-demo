package config

import "os"

type Config struct {
	ListenPort string
}

func NewConfig() *Config {
	listenPort := os.Getenv("LISTEN_PORT")
	if len(listenPort) == 0 {
		listenPort = defaultListenPort
	}

	return &Config{
		ListenPort: ":" + listenPort,
	}
}
