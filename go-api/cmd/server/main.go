package main

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(findEnvFile())
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
