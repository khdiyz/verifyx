package handler

import (
	"fmt"
	"strconv"
	"verifyx/config"
	"verifyx/internal/models"

	"github.com/labstack/echo/v4"
)

func listPagination(c echo.Context) (models.Pagination, error) {
	page, err := getPageQuery(c)
	if err != nil {
		return models.Pagination{}, err
	}
	limit, err := getLimitQuery(c)
	if err != nil {
		return models.Pagination{}, err
	}

	offset, limit := calculatePagination(page, limit)

	return models.Pagination{
		Page:   page,
		Limit:  limit,
		Offset: offset,
	}, nil
}

func getPageQuery(c echo.Context) (int, error) {
	// Get the page query parameter with a default value
	pageStr := c.QueryParam("page")
	if pageStr == "" {
		pageStr = config.DefaultPage // Ensure DefaultPage is a string
	}

	// Convert the page string to an integer
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return 0, fmt.Errorf("invalid page value: %v", err)
	}

	// Validate that the page number is positive
	if page < 1 {
		return 0, fmt.Errorf("page must be greater than or equal to 1")
	}

	return page, nil
}

func getLimitQuery(c echo.Context) (int, error) {
	limitStr := c.QueryParam("limit")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return 0, fmt.Errorf("error while parsing query: %v", err.Error())
	}

	if limit < 1 {
		return 0, fmt.Errorf("invalid limit query")
	}

	return limit, nil
}

// it returns offset and limit
func calculatePagination(page, limit int) (int, int) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	return (page - 1) * limit, limit
}
