package repository

import (
	"fmt"
	"strings"
	"verifyx/internal/models"
	"verifyx/pkg/logger"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db     *sqlx.DB
	logger *logger.Logger
}

func NewUserRepo(db *sqlx.DB, logger *logger.Logger) *userRepo {
	return &userRepo{
		db:     db,
		logger: logger,
	}
}

func (r *userRepo) Create(request models.CreateUser) (uuid.UUID, error) {
	id := uuid.New()

	query := squirrel.Insert("users").Columns(
		"id", "first_name", "last_name", "department_id", "phone_number", "profile_image",
	).Values(
		id, request.FirstName, request.LastName, request.DepartmentId, request.PhoneNumber, request.ProfileImage,
	).PlaceholderFormat(squirrel.Dollar)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		r.logger.Error(err)
		return uuid.Nil, err
	}

	_, err = r.db.Exec(sqlQuery, args...)
	if err != nil {
		r.logger.Error(err)
		return uuid.Nil, err
	}

	return id, nil
}

func (r *userRepo) GetList(filter models.UserFilter) ([]models.User, int, error) {
	// Base query and dynamic conditions
	baseQuery := `
	SELECT 
		id, 
		first_name, 
		last_name, 
		department_id, 
		phone_number, 
		profile_image, 
		created_at 
	FROM users 
	WHERE 
		deleted_at IS NULL `

	countQuery := `SELECT COUNT(*) FROM users WHERE deleted_at IS NULL `
	conditions := []string{}
	params := map[string]any{
		"limit":  filter.Limit,
		"offset": filter.Offset,
	}

	// Add search condition
	if filter.Search != "" {
		conditions = append(conditions, "(first_name || last_name || phone_number) ILIKE :search")
		params["search"] = "%" + filter.Search + "%"
	}

	// Add WHERE clause if conditions exist
	if len(conditions) > 0 {
		whereClause := " AND " + strings.Join(conditions, " AND ")
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

	users := []models.User{}
	rows, err := r.db.NamedQuery(baseQuery, params)
	if err != nil {
		r.logger.Error(err)
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			user models.User
		)
		if err = rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.DepartmentId,
			&user.PhoneNumber,
			&user.ProfileImage,
			&user.CreatedAt,
		); err != nil {
			r.logger.Error(err)
			return nil, 0, err
		}

		users = append(users, user)
	}

	// Execute the count query
	var total int
	countQuery, countArgs, err := sqlx.Named(countQuery, params)
	if err != nil {
		r.logger.Error(err)
		return nil, 0, err
	}
	countQuery = r.db.Rebind(countQuery)

	if err := r.db.Get(&total, countQuery, countArgs...); err != nil {
		r.logger.Error(err)
		return nil, 0, err
	}

	return users, total, nil
}
