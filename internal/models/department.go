package models

import (
	"time"

	"github.com/google/uuid"
)

type Department struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type CreateDepartment struct {
	Name string `json:"name" validate:"required"`
}

type UpdateDepartment struct {
	ID   uuid.UUID `json:"-"`
	Name string    `json:"name" validate:"required"`
}

type DepartmentFilter struct {
	Offset int
	Limit  int
	Search string
	SortBy string
	Order  string
}
