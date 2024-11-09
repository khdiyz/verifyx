package handler

import (
	"verifyx/internal/models"

	"github.com/labstack/echo/v4"
)

func errorResponse(c echo.Context, status int, err error) error {
	return c.JSON(status, models.ErrorResponse{
		ErrorMessage: err.Error(),
	})
}
