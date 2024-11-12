package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	DepartmentId uuid.UUID `json:"-"`
	PhoneNumber  string    `json:"phone_number"`
	ProfileImage string    `json:"profile_image"`
	CreatedAt    time.Time `json:"created_at"`

	Department Department `json:"department"`
}

type CreateUser struct {
	FirstName     string    `json:"first_name" validate:"required"`
	LastName      string    `json:"last_name" validate:"required"`
	DepartmentId  uuid.UUID `json:"department_id"`
	PhoneNumber   string    `json:"phone_number" validate:"required"`
	ProfileImage  string    `json:"profile_image"`
	FaceEmbedding []float64 `json:"-"`
}

type UpdateUser struct {
	ID            uuid.UUID `json:"-"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	DepartmentId  uuid.UUID `json:"department_id"`
	PhoneNumber   string    `json:"phone_number"`
	ProfileImage  string    `json:"profile_image"`
	FaceEmbedding []float64 `json:"face_embedding"`
}

type UserFilter struct {
	Offset int
	Limit  int
	Search string
	SortBy string
	Order  string
}
