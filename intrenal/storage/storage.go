package storage

import "gorm.io/gorm"

type Storage struct{}

func NewStorage(db *gorm.DB) *Storage {
	return &Storage{}
}
