package service

import "github.com/cirodriiguez/go-db/storage"

type Service struct {
	storage storage.Storage
}

func NewService(s storage.Storage) *Service {
	return &Service{s}
}

func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
