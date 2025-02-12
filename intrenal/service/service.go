package service

import "kurs/intrenal/storage"

type Service struct{}

func NewService(store *storage.Storage) *Service {
	return &Service{}
}
