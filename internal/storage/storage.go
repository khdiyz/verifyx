package storage

import (
	"verifyx/config"
	"verifyx/pkg/logger"

	"github.com/minio/minio-go/v7"
)

type Storage struct {
}

func NewStorage(minio *minio.Client, cfg *config.Config, logger *logger.Logger) *Storage {
	return &Storage{
		
	}
}
