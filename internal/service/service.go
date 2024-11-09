package service

import (
	"verifyx/config"
	"verifyx/internal/repository"
	"verifyx/internal/storage"
	"verifyx/pkg/logger"
)

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository, storage *storage.Storage, cfg *config.Config, loggers *logger.Logger) *Service {
	return &Service{
		Authorization: NewAuthService(cfg),
	}
}

type Authorization interface {
	GenerateToken(role, username string) (string, error)
}
