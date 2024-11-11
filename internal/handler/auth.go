package handler

import (
	"errors"
	"net/http"
	"verifyx/config"
	"verifyx/internal/models"

	"github.com/labstack/echo/v4"
)

type loginResponse struct {
	Token string `json:"token"`
}

// @Description Sign In Admin
// @Summary Sign In Admin
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body models.Login true "Login"
// @Success 200 {object} loginResponse
// @Failure 400,404,500 {object} models.ErrorResponse
// @Router /api/v1/auth/login [post]
func (h *Handler) adminLogin(c echo.Context) error {
	var login models.Login

	cfg := config.GetConfig()

	if err := c.Bind(&login); err != nil {
		return errorResponse(c, http.StatusBadRequest, errors.New("invalid input"))
	}

	// Check admin credentials (replace with DB check)
	if login.Username != cfg.Username || login.Password != cfg.Password {
		return errorResponse(c, http.StatusUnauthorized, errors.New("incorrect username or password"))
	}

	// Generate JWT token
	token, err := h.service.Authorization.GenerateToken("admin", login.Username)
	if err != nil {
		return errorResponse(c, http.StatusInternalServerError, errors.New("failed to generate token"))
	}

	return c.JSON(http.StatusOK, loginResponse{
		Token: token,
	})
}
