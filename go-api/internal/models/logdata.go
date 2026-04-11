package models

import (
	"time"

	"github.com/BlackestDawn/loganalysis-demo/go-api/internal/db"
	"github.com/jackc/pgx/v5/pgtype"
)

type LogData struct {
	Timestamp time.Time `json:"timestamp" binding:"required"`
	Level     string    `json:"level" binding:"required"`
	Service   string    `json:"service" binding:"required"`
	Message   string    `json:"message" binding:"required"`
}

func (l *LogData) ToDbParam() db.StoreLogDataParams {
	return db.StoreLogDataParams{
		Timestamp: pgtype.Timestamptz{
			Time:  l.Timestamp,
			Valid: true,
		},
		Level:   l.Level,
		Service: l.Service,
		Message: l.Message,
	}
}

func FromDbRow(row db.GetLogDataRow) LogData {
	return LogData{
		Timestamp: row.Timestamp.Time,
		Level:     row.Level,
		Service:   row.Service,
		Message:   row.Message,
	}
}
