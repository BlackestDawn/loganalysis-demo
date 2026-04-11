package services

import (
	"fmt"
	"os"

	"github.com/BlackestDawn/loganalysis-demo/go-api/internal/config"
	"github.com/BlackestDawn/loganalysis-demo/go-api/internal/models"
	"github.com/BlackestDawn/loganalysis-demo/go-api/internal/repository"
)

type LogDataServices interface {
	WriteLogData(data models.LogData) error
	FetchLogData() ([]models.LogData, error)
}

type logService struct {
	logRepo repository.LogRepository
}

func (s *logService) WriteLogData(data models.LogData) (err error) {
	err = s.logRepo.StoreLogData(data)
	return
}

func (s *logService) FetchLogData() (data []models.LogData, err error) {
	data, err = s.logRepo.FetchLogData()
	return
}

func NewLogdataService(conf *config.Config) *logService {
	serv := new(logService)

	repo, err := repository.NewLogPostgresStorage(conf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to database: %s\nFalling back to memory storage", err)
		serv.logRepo = repository.NewLogMemoryStorage(conf)
	} else {
		serv.logRepo = repo
	}

	return serv
}
