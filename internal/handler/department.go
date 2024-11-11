package handler

import (
	"math"
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

type getDepartmentsResponse struct {
	Departments []models.Department `json:"data"`
	Pagination  models.Pagination   `json:"pagination"`
}

// @Description Get List Department
// @Summary Get List Department
// @Tags Department
// @Accept json
// @Produce json
// @Param limit query int64 true "limit" default(10)
// @Param page  query int64 true "page" default(1)
// @Param search query string false "search"
// @Param sort_by query string false "sort by"
// @Param order query string false "sort by"
// @Success 200 {object} getDepartmentsResponse
// @Failure 400,401,404,500 {object} models.ErrorResponse
// @Router /api/v1/departments [get]
// @Security ApiKeyAuth
func (h *Handler) getDepartments(c echo.Context) error {
	pagination, err := listPagination(c)
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	var filter models.DepartmentFilter
	filter.Limit = pagination.Limit
	filter.Offset = pagination.Offset

	search := c.QueryParam("search")
	if search != "" {
		filter.Search = search
	}

	sortBy := c.QueryParam("sort_by")
	if sortBy != "" {
		filter.SortBy = sortBy
	}

	order := c.QueryParam("order")
	if order != "" {
		filter.Order = order
	}

	departaments, totalCount, err := h.service.Department.GetDepartments(filter)
	if err != nil {
		return fromError(c, err)
	}
	pagination.TotalCount = totalCount

	pageCount := math.Ceil(float64(pagination.TotalCount) / float64(pagination.Limit))
	pagination.PageCount = int(pageCount)

	return c.JSON(http.StatusOK, getDepartmentsResponse{
		Departments: departaments,
		Pagination:  pagination,
	})
}

// @Description Get Department
// @Summary Get Department
// @Tags Department
// @Accept json
// @Produce json
// @Param id path string true "department id"
// @Success 200 {object} models.Department
// @Failure 400,401,404,500 {object} models.ErrorResponse
// @Router /api/v1/departments/{id} [get]
// @Security ApiKeyAuth
func (h *Handler) getDepartment(c echo.Context) error {
	id, err := getUUIDParam(c, "id")
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	department, err := h.service.Department.GetDepartment(id)
	if err != nil {
		return fromError(c, err)
	}

	return c.JSON(http.StatusOK, models.Department{
		ID:        department.ID,
		Name:      department.Name,
		CreatedAt: department.CreatedAt,
	})
}

// @Description Update Department
// @Summary Update Department
// @Tags Department
// @Accept json
// @Produce json
// @Param id path string true "department id"
// @Param update body models.UpdateDepartment true "update department body"
// @Success 200 {object} createResponse
// @Failure 400,401,404,500 {object} models.ErrorResponse
// @Router /api/v1/departments/{id} [put]
// @Security ApiKeyAuth
func (h *Handler) updateDepartment(c echo.Context) error {
	id, err := getUUIDParam(c, "id")
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	var body models.UpdateDepartment
	if err := c.Bind(&body); err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}
	body.ID = id

	if err := validator.ValidatePayloads(body); err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	if err = h.service.Update(body); err != nil {
		return fromError(c, err)
	}

	return c.JSON(http.StatusOK, createResponse{
		Id: id,
	})
}

// @Description Delete Department
// @Summary Delete Department
// @Tags Department
// @Accept json
// @Produce json
// @Param id path string true "department id"
// @Success 200 {object} createResponse
// @Failure 400,401,404,500 {object} models.ErrorResponse
// @Router /api/v1/departments/{id} [delete]
// @Security ApiKeyAuth
func (h *Handler) deleteDepartment(c echo.Context) error {
	id, err := getUUIDParam(c, "id")
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	if err = h.service.Department.Delete(id); err != nil {
		return fromError(c, err)
	}

	return c.JSON(http.StatusOK, createResponse{
		Id: id,
	})
}
