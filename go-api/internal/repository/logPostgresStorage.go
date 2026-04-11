package repository

import (
	"context"

	"github.com/BlackestDawn/loganalysis-demo/go-api/internal/config"
	"github.com/BlackestDawn/loganalysis-demo/go-api/internal/db"
	"github.com/BlackestDawn/loganalysis-demo/go-api/internal/models"
)

type logPostgresStorage struct {
	queries *db.Queries
	ctx     context.Context
}

func (s *logPostgresStorage) StoreLogData(data models.LogData) (err error) {
	err = s.queries.StoreLogData(s.ctx, db.StoreLogDataParams(data.ToDbParam()))
	return
}

func (s *logPostgresStorage) FetchLogData() (data []models.LogData, err error) {
	rows, err := s.queries.GetLogData(s.ctx)
	if err != nil {
		return
	}

	for _, row := range rows {
		data = append(data, models.FromDbRow(row))
	}

	return
}

func NewLogPostgresStorage(conf *config.Config) (storage *logPostgresStorage, err error) {
	conn, err := db.Init(conf)
	if err != nil {
		return
	}
	conf.AddCloser(func() error {
		return conn.Close(conf.DbCtx)
	})

	storage = &logPostgresStorage{
		queries: db.New(conn),
		ctx:     conf.DbCtx,
	}

	return
}
