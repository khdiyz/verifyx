package handler

import (
	"errors"
	"net/http"
	"verifyx/internal/models"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func errorResponse(c echo.Context, status int, err error) error {
	return c.JSON(status, models.ErrorResponse{
		ErrorMessage: err.Error(),
	})
}

func fromError(c echo.Context, serviceError error) error {
	st, _ := status.FromError(serviceError)
	err := st.Message()

	switch st.Code() {
	case codes.NotFound:
		return errorResponse(c, http.StatusNotFound, errors.New(err))
	case codes.InvalidArgument:
		return errorResponse(c, http.StatusBadRequest, errors.New(err))
	case codes.Unavailable:
		return errorResponse(c, http.StatusUnavailableForLegalReasons, errors.New(err))
	case codes.AlreadyExists:
		return errorResponse(c, http.StatusBadRequest, errors.New(err))
	case codes.Unauthenticated:
		return errorResponse(c, http.StatusUnauthorized, errors.New(err))
	}

	return errorResponse(c, http.StatusInternalServerError, errors.New(err))
}
