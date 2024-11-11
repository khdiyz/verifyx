package repository

import (
	"verifyx/internal/models"
	"verifyx/pkg/logger"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type departmentRepo struct {
	db     *sqlx.DB
	logger *logger.Logger
}

func NewDepartmentRepo(db *sqlx.DB, logger *logger.Logger) *departmentRepo {
	return &departmentRepo{
		db:     db,
		logger: logger,
	}
}

func (r *departmentRepo) Create(request models.CreateDepartment) (uuid.UUID, error) {
	// Generate a new UUID for the department
	id := uuid.New()

	// Define the SQL query
	query := `
		INSERT INTO departments (
			id,
			name
		) VALUES (:id, :name);`

	// Prepare the data for the query
	data := map[string]interface{}{
		"id":   id,
		"name": request.Name,
	}

	// Execute the query
	_, err := r.db.NamedExec(query, data)
	if err != nil {
		r.logger.Error(err)
		return uuid.Nil, err
	}

	return id, nil
}
