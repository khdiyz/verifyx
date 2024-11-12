package handler

import (
	"math"
	"net/http"
	"verifyx/internal/models"
	"verifyx/pkg/validator"

	"github.com/labstack/echo/v4"
)

// @Description Create User
// @Summary Create User
// @Tags User
// @Accept json
// @Produce json
// @Param create body models.CreateUser true "Create User"
// @Success 201 {object} createResponse
// @Failure 400,401,404,500 {object} models.ErrorResponse
// @Router /api/v1/users [post]
// @Security ApiKeyAuth
func (h *Handler) createUser(c echo.Context) error {
	var body models.CreateUser

	if err := c.Bind(&body); err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}
	body.FaceEmbedding = nil

	if err := validator.ValidatePayloads(body); err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	id, err := h.service.User.CreateUser(body)
	if err != nil {
		return fromError(c, err)
	}

	return c.JSON(http.StatusCreated, createResponse{
		Id: id,
	})
}

type getUsersResponse struct {
	Users      []models.User     `json:"data"`
	Pagination models.Pagination `json:"pagination"`
}

// @Description Get List User
// @Summary Get List User
// @Tags User
// @Accept json
// @Produce json
// @Param limit query int64 true "limit" default(10)
// @Param page  query int64 true "page" default(1)
// @Param search query string false "search"
// @Param sort_by query string false "sort by"
// @Param order query string false "sort by"
// @Success 200 {object} getUsersResponse
// @Failure 400,401,404,500 {object} models.ErrorResponse
// @Router /api/v1/users [get]
// @Security ApiKeyAuth
func (h *Handler) getUsers(c echo.Context) error {
	pagination, err := listPagination(c)
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	var filter models.UserFilter
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

	users, totalCount, err := h.service.User.GetUsers(filter)
	if err != nil {
		return fromError(c, err)
	}
	pagination.TotalCount = totalCount

	pageCount := math.Ceil(float64(pagination.TotalCount) / float64(pagination.Limit))
	pagination.PageCount = int(pageCount)

	return c.JSON(http.StatusOK, getUsersResponse{
		Users:      users,
		Pagination: pagination,
	})
}
