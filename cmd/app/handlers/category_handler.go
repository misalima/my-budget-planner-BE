package handlers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"my-budget-planner/internal/postgres/models"
	"my-budget-planner/internal/services"
	"net/http"
	"strconv"
)

type CategoryHandler struct {
	categoryService *services.CategoryServices
}

func NewCategoryHandler(categoryService *services.CategoryServices) *CategoryHandler {
	return &CategoryHandler{categoryService: categoryService}
}

func (h *CategoryHandler) CreateCategory(ctx echo.Context) error {
	var category models.Category
	err := ctx.Bind(&category)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request data"})
	}
	if category.Name == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "category name is required"})
	}
	if category.UserID.String() == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "user id is required"})
	}

	err = h.categoryService.CreateCategory(&category)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return ctx.JSON(http.StatusCreated, category)
}

func (h *CategoryHandler) GetCategoriesByUserID(ctx echo.Context) error {
	userIdStr := ctx.Param("user_id")
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid user id"})
	}

	categories, err := h.categoryService.GetCategoriesByUserID(userId)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return ctx.JSON(http.StatusOK, categories)
}

func (h *CategoryHandler) DeleteCategory(ctx echo.Context) error {
	categoryId := ctx.Param("id")
	intCategoryId, err := strconv.Atoi(categoryId)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid category id"})
	}
	err = h.categoryService.DeleteCategory(intCategoryId)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return ctx.JSON(http.StatusNoContent, nil)
}
