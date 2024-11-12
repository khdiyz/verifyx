package service

import (
	"verifyx/config"
	"verifyx/internal/models"
	"verifyx/internal/repository"
	"verifyx/internal/storage"
	"verifyx/pkg/logger"

	"github.com/google/uuid"
)

type Service struct {
	Authorization
	Department
	User
}

func NewService(repos *repository.Repository, storage *storage.Storage, cfg *config.Config, loggers *logger.Logger) *Service {
	return &Service{
		Authorization: NewAuthService(cfg),
		Department:    NewDepartmentService(repos, loggers),
		User:          NewUserService(repos, loggers),
	}
}

type Authorization interface {
	GenerateToken(role, username string) (string, error)
	ParseToken(token string) (*jwtCustomClaim, error)
}

type Department interface {
	Create(request models.CreateDepartment) (uuid.UUID, error)
	GetDepartments(filter models.DepartmentFilter) ([]models.Department, int, error)
	GetDepartment(id uuid.UUID) (models.Department, error)
	Update(request models.UpdateDepartment) error
	Delete(id uuid.UUID) error
}

type User interface {
	CreateUser(request models.CreateUser) (uuid.UUID, error)
	GetUsers(filter models.UserFilter) ([]models.User, int, error)
}
