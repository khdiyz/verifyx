package models

import (
	"time"

	"github.com/google/uuid"
)

type Department struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateDepartment struct {
	Name string `json:"name" validate:"required"`
}

type UpdateDepartment struct {
	ID   uuid.UUID `json:"-"`
	Name string    `json:"name"`
}
