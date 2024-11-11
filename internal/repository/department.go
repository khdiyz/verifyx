package repository

import (
	"fmt"
	"strings"
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

func (r *departmentRepo) GetDepartments(filter models.DepartmentFilter) ([]models.Department, int, error) {

	// Base query and dynamic conditions
	baseQuery := `SELECT id, name, created_at FROM departments`
	countQuery := `SELECT COUNT(*) FROM departments`
	conditions := []string{}
	params := map[string]interface{}{
		"limit":  filter.Limit,
		"offset": filter.Offset,
	}

	// Add search condition
	if filter.Search != "" {
		conditions = append(conditions, "name ILIKE :search")
		params["search"] = "%" + filter.Search + "%"
	}

	// Add WHERE clause if conditions exist
	if len(conditions) > 0 {
		whereClause := " WHERE " + strings.Join(conditions, " AND ")
		baseQuery += whereClause
		countQuery += whereClause
	}

	// Add sorting
	if filter.SortBy != "" {
		order := "ASC"
		if strings.ToUpper(filter.Order) == "DESC" {
			order = "DESC"
		}
		baseQuery += fmt.Sprintf(" ORDER BY %s %s", filter.SortBy, order)
	}

	// Add pagination
	baseQuery += " LIMIT :limit OFFSET :offset"

	// Execute the main query
	var departments []models.Department
	rows, err := r.db.NamedQuery(baseQuery, params)
	if err != nil {
		r.logger.Error(err)
		return nil, 0, fmt.Errorf("failed to fetch departments: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var department models.Department
		if err := rows.StructScan(&department); err != nil {
			r.logger.Error(err)
			return nil, 0, fmt.Errorf("failed to scan department row: %w", err)
		}
		departments = append(departments, department)
	}

	// Execute the count query
	var total int
	countQuery, countArgs, err := sqlx.Named(countQuery, params)
	if err != nil {
		r.logger.Error(err)
		return nil, 0, fmt.Errorf("failed to build count query: %w", err)
	}
	countQuery = r.db.Rebind(countQuery)

	if err := r.db.Get(&total, countQuery, countArgs...); err != nil {
		r.logger.Error(err)
		return nil, 0, fmt.Errorf("failed to count total departments: %w", err)
	}

	return departments, total, nil
}
