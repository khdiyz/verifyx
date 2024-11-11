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
// @Success 200 {object} models.ListResponse
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

	return c.JSON(http.StatusOK, models.ListResponse{
		Data:       departaments,
		Pagination: pagination,
	})
}
