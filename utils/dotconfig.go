package utils

import (
	"path/filepath"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var Logger *zap.Logger

func Dotconfig() {
	Logger = NewLogger()
	envPath, err := filepath.Abs(".env")
	if err != nil {
		Logger.Error("Error getting absolute path for .env file: " + err.Error())
	}

	err = godotenv.Load(envPath)
	if err != nil {
		Logger.Error("Error while opening .env", zap.Error(err))
	}
}
