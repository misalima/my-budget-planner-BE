package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HealthHandler(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Server is running")
}
