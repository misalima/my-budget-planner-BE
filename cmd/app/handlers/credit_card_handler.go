package handlers

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"my-budget-planner/internal/postgres/models"
	"my-budget-planner/internal/services"
	"net/http"
)

type CreditCardHandler struct {
	service *services.CreditCardService
}

func NewCreditCardHandler(service *services.CreditCardService) *CreditCardHandler {
	return &CreditCardHandler{service: service}
}

func (h *CreditCardHandler) GetAllCreditCards(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID, err := uuid.Parse(claims["user_id"].(string))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	creditCards, err := h.service.GetAllByUserID(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, creditCards)
}

func (h *CreditCardHandler) GetCreditCardByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid credit card ID"})
	}

	creditCard, err := h.service.GetByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, creditCard)
}

func (h *CreditCardHandler) CreateCreditCard(c echo.Context) error {
	var cc models.CreditCard
	if err := c.Bind(&cc); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	cc.ID = uuid.New()
	if err := h.service.Create(c.Request().Context(), &cc); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, cc)
}

func (h *CreditCardHandler) DeleteCreditCard(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid credit card ID"})
	}

	if err := h.service.Delete(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
