package repository

import (
	"github.com/BlackestDawn/loganalysis-demo/go-api/internal/config"
	"github.com/BlackestDawn/loganalysis-demo/go-api/internal/models"
)

type logMemoryStorage struct {
	logs []models.LogData
}

func (s *logMemoryStorage) StoreLogData(data models.LogData) (err error) {
	s.logs = append(s.logs, data)
	return
}

func (s *logMemoryStorage) FetchLogData() (data []models.LogData, err error) {
	return s.logs, nil
}

func NewLogMemoryStorage(conf *config.Config) *logMemoryStorage {
	return &logMemoryStorage{
		logs: make([]models.LogData, 0),
	}
}
