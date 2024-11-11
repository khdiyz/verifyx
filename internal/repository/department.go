package repository

import (
	"errors"
	"fmt"
	"strings"
	"verifyx/internal/models"
	"verifyx/pkg/logger"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

var (
	errNoRowsAffected = errors.New("no rows affected")
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

func (r *departmentRepo) GetList(filter models.DepartmentFilter) ([]models.Department, int, error) {

	// Base query and dynamic conditions
	baseQuery := `SELECT id, name, created_at FROM departments WHERE deleted_at IS NULL `
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
		whereClause := strings.Join(conditions, " AND ")
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

func (r *departmentRepo) GetById(id uuid.UUID) (models.Department, error) {
	var department models.Department

	query := `
	SELECT
		id,
		name,
		created_at
	FROM departments
	WHERE id = $1 AND deleted_at IS NULL;`

	if err := r.db.QueryRow(query, id).Scan(
		&department.ID,
		&department.Name,
		&department.CreatedAt,
	); err != nil {
		r.logger.Error(err)
		return models.Department{}, err
	}

	return department, nil
}

func (r *departmentRepo) Update(request models.UpdateDepartment) error {
	query := `
	UPDATE departments
	SET
		name = :name,
		updated_at = now()
	WHERE
		id = :id
		AND deleted_at IS NULL;`

	// Prepare the data for the query
	data := map[string]interface{}{
		"id":   request.ID,
		"name": request.Name,
	}

	// Execute the query
	row, err := r.db.NamedExec(query, data)
	if err != nil {
		r.logger.Error(err)
		return err
	}

	rowAffected, err := row.RowsAffected()
	if err != nil {
		r.logger.Error(err)
		return err
	}

	if rowAffected == 0 {
		return errNoRowsAffected
	}

	return nil
}

func (r *departmentRepo) Delete(id uuid.UUID) error {
	query := `
	UPDATE departments
	SET
		deleted_at = now()
	WHERE
		id = $1
		AND deleted_at IS NULL;`

	// Execute the query
	row, err := r.db.Exec(query, id)
	if err != nil {
		r.logger.Error(err)
		return err
	}

	rowAffected, err := row.RowsAffected()
	if err != nil {
		r.logger.Error(err)
		return err
	}

	if rowAffected == 0 {
		return errNoRowsAffected
	}

	return nil
}
