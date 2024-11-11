package service

import (
	"verifyx/internal/models"
	"verifyx/internal/repository"
	"verifyx/pkg/logger"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
)

type DepartmentService struct {
	repo   *repository.Repository
	logger *logger.Logger
}

func NewDepartmentService(repo *repository.Repository, logger *logger.Logger) *DepartmentService {
	return &DepartmentService{
		repo:   repo,
		logger: logger,
	}
}

func (s *DepartmentService) Create(request models.CreateDepartment) (uuid.UUID, error) {
	id, err := s.repo.Department.Create(request)
	if err != nil {
		return uuid.Nil, serviceError(err, codes.Internal)
	}

	return id, nil
}
