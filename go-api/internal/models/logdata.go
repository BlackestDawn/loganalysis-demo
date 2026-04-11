package models

import "time"

type LogData struct {
	Timestamp time.Time `json:"timestamp" binding:"required"`
	Level     string    `json:"level" binding:"required"`
	Service   string    `json:"service" binding:"required"`
	Message   string    `json:"message" binding:"required"`
}
