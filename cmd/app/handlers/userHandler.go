package handlers

import (
	"github.com/labstack/echo/v4"
	"my-budget-planner/internal/postgres/models"
	"my-budget-planner/internal/services"
	"net/http"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func HealthHandler(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Server is running")
}

func (h *UserHandler) CreateUserHandler(ctx echo.Context) error {

	var user models.User
	//parse the request body, with the first_name, last_name, email and password
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Validate that the required fields are not empty
	if user.Username == "" || user.FirstName == "" || user.LastName == "" || user.Email == "" || user.Password == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "All fields (username, first_name, last_name, email, and password) are required"})
	}

	//call the service
	if err := h.UserService.RegisterUser(&user); err != nil {
		return ctx.JSON(http.StatusConflict, map[string]string{"error": err.Error()})
	}
	//handle the response
	return ctx.JSON(http.StatusCreated, map[string]string{"message": "User created successfully"})

}
