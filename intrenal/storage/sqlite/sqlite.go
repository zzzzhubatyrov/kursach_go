package storage

import (
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SQLiteStorageInit() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("kurs.db"), &gorm.Config{})
	if err != nil {
		log.Error(err)
	}
	return db, err
}
