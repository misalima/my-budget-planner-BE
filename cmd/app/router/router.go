package router

import (
	"github.com/labstack/echo/v4"
	"my-budget-planner/cmd/app/handlers"
	"net/http"
)

func LoadRoutes(e *echo.Echo, userHandler *handlers.UserHandler, authHandler *handlers.AuthHandler) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/health", handlers.HealthHandler)
	e.POST("/user", userHandler.CreateUserHandler)

	//auth routes
	e.GET("/auth/refresh", authHandler.RefreshTokenHandler)
}
