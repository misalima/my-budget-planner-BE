package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"my-budget-planner/internal/postgres/models"
	"my-budget-planner/internal/services"
	"net/http"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

// RefreshTokenHandler refreshes the access token
func (a *AuthHandler) RefreshTokenHandler(ctx echo.Context) error {
	var refreshToken models.RefreshToken

	//parse the request body, with the token
	if err := ctx.Bind(&refreshToken); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Validate that the required fields are not empty
	if refreshToken.Token == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Token is required"})
	}

	//extract user id data from the jwt token
	userId := ctx.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)["user_id"].(string)

	//call the service
	accessToken, err := a.AuthService.RefreshToken(userId, refreshToken.Token)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	//handle the response
	return ctx.JSON(http.StatusOK, map[string]string{"access_token": accessToken})
}
