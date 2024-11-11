package repository

import (
	"verifyx/internal/models"
	"verifyx/pkg/logger"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Department
}

func NewRepository(db *sqlx.DB, logger *logger.Logger) *Repository {
	return &Repository{
		Department: NewDepartmentRepo(db, logger),
	}
}

type Department interface {
	Create(request models.CreateDepartment) (uuid.UUID, error)
	GetList(filter models.DepartmentFilter) ([]models.Department, int, error)
	GetById(id uuid.UUID) (models.Department, error)
	Update(request models.UpdateDepartment) error
	Delete(id uuid.UUID) error
}
