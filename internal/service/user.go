package service

import (
	"verifyx/internal/models"
	"verifyx/internal/repository"
	"verifyx/pkg/logger"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
)

type userService struct {
	repo   *repository.Repository
	logger *logger.Logger
}

func NewUserService(repo *repository.Repository, logger *logger.Logger) *userService {
	return &userService{
		repo:   repo,
		logger: logger,
	}
}

func (s *userService) CreateUser(request models.CreateUser) (uuid.UUID, error) {
	_, err := s.repo.Department.GetById(request.DepartmentId)
	if err != nil {
		return uuid.Nil, serviceError(err, codes.Internal)
	}

	id, err := s.repo.User.Create(request)
	if err != nil {
		return uuid.Nil, serviceError(err, codes.Internal)
	}

	return id, nil
}

func (s *userService) GetUsers(filter models.UserFilter) ([]models.User, int, error) {
	users, total, err := s.repo.User.GetList(filter)
	if err != nil {
		return nil, 0, serviceError(err, codes.Internal)
	}

	departmentIds := make([]uuid.UUID, len(users))
	for i := range users {
		departmentIds[i] = users[i].DepartmentId
	}

	departments, err := s.repo.Department.GetByIds(departmentIds)
	if err != nil {
		return nil, 0, serviceError(err, codes.Internal)
	}

	departmentsMap := make(map[uuid.UUID]models.Department, len(departments))
	for i := range departments {
		departmentsMap[departments[i].ID] = departments[i]
	}

	for i := range users {
		users[i].Department = departmentsMap[users[i].DepartmentId]
	}

	return users, total, nil
}
