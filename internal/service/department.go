package service

import (
	"errors"
	"verifyx/internal/models"
	"verifyx/internal/repository"
	"verifyx/pkg/logger"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
)

type departmentService struct {
	repo   *repository.Repository
	logger *logger.Logger
}

func NewDepartmentService(repo *repository.Repository, logger *logger.Logger) *departmentService {
	return &departmentService{
		repo:   repo,
		logger: logger,
	}
}

func (s *departmentService) Create(request models.CreateDepartment) (uuid.UUID, error) {
	id, err := s.repo.Department.Create(request)
	if err != nil {
		return uuid.Nil, serviceError(err, codes.Internal)
	}

	return id, nil
}

func (s *departmentService) GetDepartments(filter models.DepartmentFilter) ([]models.Department, int, error) {
	if filter.SortBy != "" && filter.SortBy != "created_at" {
		return nil, 0, serviceError(errors.New("invalid sort by filter"), codes.InvalidArgument)
	}
	if filter.Order != "" && filter.Order != "asc" && filter.Order != "desc" {
		return nil, 0, serviceError(errors.New("invalid order filter"), codes.InvalidArgument)

	}

	departments, total, err := s.repo.Department.GetList(filter)
	if err != nil {
		return nil, 0, serviceError(err, codes.Internal)
	}

	return departments, total, nil
}

func (s *departmentService) GetDepartment(id uuid.UUID) (models.Department, error) {
	department, err := s.repo.Department.GetById(id)
	if err != nil {
		return models.Department{}, serviceError(err, codes.Internal)
	}

	return department, nil
}

func (s *departmentService) Update(request models.UpdateDepartment) error {
	err := s.repo.Department.Update(request)
	if err != nil {
		return serviceError(err, codes.Internal)
	}

	return nil
}

func (s *departmentService) Delete(id uuid.UUID) error {
	err := s.repo.Department.Delete(id)
	if err != nil {
		return serviceError(err, codes.Internal)
	}

	return nil
}
