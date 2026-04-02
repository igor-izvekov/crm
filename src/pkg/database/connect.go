package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(dbPath string) Error {
	var err error

	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info)
	})

	if err != nil {
		return err
	}

	log.Println("Подключено к SQLite:", dbPath)
	return nil
}

func GetDB() *gorm.DB {
	return DB
}

func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		return
	}
	sqlDB.Close()
}
