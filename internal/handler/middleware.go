package handler

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// Constants for header keys and context keys
const (
	authorizationHeader = "Authorization"
	username            = "username"
	role                = "role"
)

// userIdentity is a middleware function to validate user identity from the Authorization header
func (h *Handler) userIdentity(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Retrieve the Authorization header
		header := c.Request().Header.Get(authorizationHeader)
		if header == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "empty auth header",
			})
		}

		// Split the header into parts
		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid auth header",
			})
		}

		// Validate the token part
		token := headerParts[1]
		if len(token) == 0 {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "token is empty",
			})
		}

		// Parse the token and validate claims
		claims, err := h.service.Authorization.ParseToken(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": err.Error(),
			})
		}

		// Set user information in context for downstream handlers
		c.Set(username, claims.Username)
		c.Set(role, claims.RoleName)

		// Proceed to the next middleware or handler
		return next(c)
	}
}
