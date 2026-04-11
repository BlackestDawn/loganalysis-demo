package config

const defaultListenPort = "8080"
const defaultDbUrl = "postgres://postgres:postgres@localhost:5432/loganalysis?sslmode=disable"

const defaultLogQueueServer = "amqp://guest:guest@localhost:5672/"
const defaultLogQueueName = "logs"
const defaultLogQueueKey = "logs"
const defaultLogQueueExchange = "logs"
