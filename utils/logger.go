package utils

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() *zap.Logger {
	today := time.Now().Format("2006-01-02")
	logFileName := "logs/" + today + ".log"

	err := os.MkdirAll("logs", os.ModePerm)
	if err != nil {
		panic("Error creating log folder: " + err.Error())
	}

	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	fileWriter := zapcore.AddSync(openLogFile(logFileName))

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, fileWriter, zap.LevelEnablerFunc(func(l zapcore.Level) bool {
			return l >= zapcore.InfoLevel
		})),
		zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zap.DebugLevel),
	)

	logger := zap.New(core)

	return logger
}

func openLogFile(logFileName string) *os.File {
	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("Error opening log file: " + err.Error())
	}
	return file
}
