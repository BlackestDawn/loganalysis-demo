package handlers

import (
	"github.com/BlackestDawn/loganalysis-demo/go-api/internal/config"
	"github.com/BlackestDawn/loganalysis-demo/go-api/internal/models"
	"github.com/BlackestDawn/loganalysis-demo/go-api/internal/services"
	"github.com/gin-gonic/gin"
)

type LogDataHandler interface {
	WriteLogData(ctx *gin.Context)
	FetchLogData(ctx *gin.Context)
}

type logHandlers struct {
	logService services.LogDataServices
}

func (h *logHandlers) WriteLogData(ctx *gin.Context) {
	var logData models.LogData
	if err := ctx.ShouldBindJSON(&logData); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.logService.WriteLogData(logData); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{"status": "success"})
}

func (h *logHandlers) FetchLogData(ctx *gin.Context) {
	data, err := h.logService.FetchLogData()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, data)
}

func NewLogDataHandler(conf *config.Config) LogDataHandler {
	return &logHandlers{
		logService: services.NewLogdataService(conf),
	}
}
