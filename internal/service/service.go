package service

import (
	"verifyx/config"
	"verifyx/internal/repository"
	"verifyx/internal/storage"
	"verifyx/pkg/logger"
)

type Service struct {
}

func NewService(repos *repository.Repository, storage *storage.Storage, config *config.Config, loggers *logger.Logger) *Service {
	return &Service{}
}
