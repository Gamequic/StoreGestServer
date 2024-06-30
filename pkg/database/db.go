package database

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Logger *zap.Logger

func InitDB() error {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		Logger.Error(fmt.Sprintf("Error getting DB instance: %v", err))
		return err
	}
	Logger.Info("Connected with database")

	DB = db

	return nil
}

func CloseDB() {
	if DB != nil {
		db, err := DB.DB()
		if err != nil {
			Logger.Error(fmt.Sprintf("Error getting DB instance: %v", err))
			return
		}
		err = db.Close()
		if err != nil {
			Logger.Error(fmt.Sprintf("Error closing DB instance: %v", err))
			return
		}
		Logger.Info("Database connection closed")
	}
}
