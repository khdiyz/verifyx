package service

import (
	"errors"
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

func (s *DepartmentService) GetDepartments(filter models.DepartmentFilter) ([]models.Department, int, error) {
	if filter.SortBy != "" && filter.SortBy != "created_at" {
		return nil, 0, serviceError(errors.New("invalid sort by filter"), codes.InvalidArgument)
	}
	if filter.Order != "" && filter.Order != "asc" && filter.Order != "desc" {
		return nil, 0, serviceError(errors.New("invalid order filter"), codes.InvalidArgument)

	}

	departments, total, err := s.repo.Department.GetDepartments(filter)
	if err != nil {
		return nil, 0, serviceError(err, codes.Internal)
	}

	return departments, total, nil
}
