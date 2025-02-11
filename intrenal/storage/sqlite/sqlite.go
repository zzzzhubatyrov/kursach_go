package storage

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SQLiteStorageInit() (*gorm.DB, error) {
	file, err := os.OpenFile("gorm.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %v", err)
	}

	newLogger := logger.New(
		log.New(file, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      false,
		},
	)

	customLogger := &CustomLogger{logger: newLogger, level: logger.Info, file: file}

	db, err := gorm.Open(sqlite.Open("kurs.db"), &gorm.Config{
		Logger: customLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("err: %v", err)
	}
	return db, err
}
