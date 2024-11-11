package handler

import (
	"net/http"
	"verifyx/internal/models"
	"verifyx/pkg/validator"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type createResponse struct {
	Id uuid.UUID `json:"id"`
}

// @Description Create Department
// @Summary Create Department
// @Tags Department
// @Accept json
// @Produce json
// @Param create body models.CreateDepartment true "Create Department"
// @Success 201 {object} createResponse
// @Failure 400,401,404,500 {object} models.ErrorResponse
// @Router /api/v1/departments [post]
// @Security ApiKeyAuth
func (h *Handler) createDepartment(c echo.Context) error {
	var body models.CreateDepartment

	if err := c.Bind(&body); err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	if err := validator.ValidatePayloads(body); err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	id, err := h.service.Department.Create(body)
	if err != nil {
		return fromError(c, err)
	}

	return c.JSON(http.StatusCreated, createResponse{
		Id: id,
	})
}
