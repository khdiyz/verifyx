package handler

import (
	"os"
	"verifyx/config"
	"verifyx/docs"
	"verifyx/internal/service"
	"verifyx/pkg/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Handler struct {
	service *service.Service
	logger  *logger.Logger
}

func NewHandler(service *service.Service, loggers *logger.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  loggers,
	}
}

func (h *Handler) InitRoutes(cfg *config.Config) *echo.Echo {
	router := echo.New()

	// Echo middleware settings
	router.Use(middleware.Recover())
	router.Use(middleware.LoggerWithConfig(customLoggerConfig()))
	router.Use(middleware.CORS())

	router.GET("/docs/*", func(c echo.Context) error {
		// Dynamically set the Swagger host based on the current request
		docs.SwaggerInfo.Host = c.Request().Host
		if c.Request().TLS != nil {
			docs.SwaggerInfo.Schemes = []string{"https"}
		}
		return echoSwagger.WrapHandler(c)
	})

	router.POST("api/v1/auth/login", h.adminLogin)

	// v1 := router.Group("/api/v1", )

	return router
}

func customLoggerConfig() middleware.LoggerConfig {
	return middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} | ${latency_human} | ${method} ${uri} | Remote IP: ${remote_ip}\n",
		Output: os.Stdout,
	}
}
