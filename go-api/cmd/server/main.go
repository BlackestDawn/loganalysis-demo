package main

import (
	"os"
	"path/filepath"

	"github.com/BlackestDawn/loganalysis-demo/go-api/internal/config"
	"github.com/BlackestDawn/loganalysis-demo/go-api/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(findEnvFile())

	conf := config.NewConfig()
	defer conf.Cleanup()

	router := gin.Default()

	logHandlers := handlers.NewLogDataHandler(conf)

	router.POST("/logs", logHandlers.WriteLogData)
	router.GET("/logs", logHandlers.FetchLogData)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})

	})

	router.Run(conf.ListenPort)
}

func findEnvFile() string {
	dir, _ := os.Getwd()

	for {
		path := filepath.Join(dir, ".env")
		if _, err := os.Stat(path); err == nil {
			return path
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return ""
}
