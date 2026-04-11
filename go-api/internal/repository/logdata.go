package repository

import (
	"github.com/BlackestDawn/loganalysis-demo/go-api/internal/models"
)

type LogRepository interface {
	StoreLogData(data models.LogData) error
	FetchLogData() ([]models.LogData, error)
}
